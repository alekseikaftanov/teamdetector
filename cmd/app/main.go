package main

import (
	"context"
	"os"
	"os/signal"
	"syscall"

	"github.com/gin-gonic/gin"
	"github.com/teamdetected/internal/config"
	"github.com/teamdetected/internal/handler"
	"github.com/teamdetected/internal/repository"
	"github.com/teamdetected/internal/server"
	"github.com/teamdetected/internal/service"
)

func main() {
	// Инициализация конфигурации
	cfg, err := config.Init()
	if err != nil {
		panic(err)
	}

	// Подключение к базе данных
	dsn := "host=" + cfg.DB.Host + " port=" + cfg.DB.Port + " user=" + cfg.DB.Username + " dbname=" + cfg.DB.DBName + " password=" + os.Getenv("DB_PASSWORD") + " sslmode=" + cfg.DB.SSLMode
	db, err := repository.NewPostgresDB(dsn)
	if err != nil {
		panic(err)
	}

	// Создаем репозитории
	repos := repository.NewRepository(db)

	// Создаем email сервис
	emailServ := service.NewEmailService(
		cfg.Email.From,
		cfg.Email.Password,
		cfg.Email.Host,
		cfg.Email.Port,
	)

	// Создаем сервисы
	services := service.NewService(repos, emailServ)

	// Создаем обработчики
	handlers := handler.NewHandler(services)

	// Инициализация сервера
	router := gin.Default()
	api := router.Group("/api/v1")
	{
		auth := api.Group("/auth")
		{
			auth.POST("/register", handlers.Register)
			auth.POST("/login", handlers.Login)
			auth.DELETE("/users/:id", handlers.UserIdentity, handlers.DeleteUser)
		}

		companies := api.Group("/companies")
		{
			companies.POST("", handlers.UserIdentity, handlers.CreateCompany)
			companies.GET("", handlers.UserIdentity, handlers.GetCompanies)
			companies.GET("/:id", handlers.UserIdentity, handlers.GetCompany)
			companies.DELETE("/:id", handlers.UserIdentity, handlers.DeleteCompany)
		}

		teams := api.Group("/teams")
		{
			teams.POST("", handlers.UserIdentity, handlers.CreateTeam)
			teams.GET("/company/:company_id", handlers.UserIdentity, handlers.GetTeams)
			teams.GET("/team/:id", handlers.UserIdentity, handlers.GetTeam)
			teams.DELETE("/team/:id", handlers.UserIdentity, handlers.DeleteTeam)
		}

		// Survey routes
		survey := api.Group("/surveys", handlers.UserIdentity)
		{
			// General endpoints
			survey.GET("/questions", handlers.GetSurveyQuestions)
			survey.GET("/options", handlers.GetSurveyOptions)
			survey.POST("", handlers.CreateSurvey)
			survey.GET("/team/:team_id", handlers.GetSurveysByTeam)
			survey.GET("/:survey_id", handlers.GetSurvey)
			survey.DELETE("/:survey_id", handlers.DeleteSurvey)

			// Survey responses as a nested resource
			survey.POST("/:survey_id/responses", handlers.CreateSurveyResponse)
			survey.GET("/:survey_id/responses", handlers.GetSurveyResponses)
		}
	}

	srv := server.NewServer(cfg.Port, router)

	// Запуск сервера
	go func() {
		if err := srv.Run(); err != nil {
			panic(err)
		}
	}()

	// Graceful Shutdown
	quit := make(chan os.Signal, 1)
	signal.Notify(quit, syscall.SIGTERM, syscall.SIGINT)
	<-quit

	if err := srv.Shutdown(context.Background()); err != nil {
		panic(err)
	}
}
