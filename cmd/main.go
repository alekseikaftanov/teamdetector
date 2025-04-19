package main

import (
	"log"
	"os"

	"github.com/gin-gonic/gin"
	"github.com/joho/godotenv"
	"github.com/teamdetected/internal/handler"
	"github.com/teamdetected/internal/repository"
	"github.com/teamdetected/internal/service"
)

func main() {
	if err := godotenv.Load(); err != nil {
		log.Fatal("Error loading .env file")
	}

	db, err := repository.NewPostgresDB()
	if err != nil {
		log.Fatal(err)
	}

	repos := repository.NewRepository(db)
	services := service.NewService(repos)
	handlers := handler.NewHandler(services)

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
	}

	port := os.Getenv("PORT")
	if port == "" {
		port = "8080"
	}

	if err := router.Run(":" + port); err != nil {
		log.Fatal(err)
	}
}
