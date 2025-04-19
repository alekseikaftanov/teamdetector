package handler

import (
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
	"github.com/teamdetected/internal/model"
)

func (h *Handler) AddUsersToTeam(c *gin.Context) {
	teamID, err := strconv.Atoi(c.Param("team_id"))
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "invalid team id"})
		return
	}

	var input struct {
		Users []model.AddUserToTeamInput `json:"users" binding:"required"`
	}

	if err := c.ShouldBindJSON(&input); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	if len(input.Users) == 0 {
		c.JSON(http.StatusBadRequest, gin.H{"error": "users list cannot be empty"})
		return
	}

	err = h.services.Team.AddUsersToTeam(teamID, input.Users)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{"message": "users added to team successfully"})
}
