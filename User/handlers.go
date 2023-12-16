package user

import (
	"net/http"
	"strconv"
	"time"

	"github.com/gin-gonic/gin"
)

func createUserHandler(c *gin.Context) {
	var user User
	if err := c.ShouldBindJSON(&user); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	user.CreatedAt = time.Now()

	if err := createUser(&user); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Error creating user"})
		return
	}

	c.JSON(http.StatusCreated, user)
}

func updateUserHandler(c *gin.Context) {
	userID, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid user ID"})
		return
	}

	var updatedUser User
	if err := c.ShouldBindJSON(&updatedUser); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	updatedUser.ID = userID
	if err := updateUser(&updatedUser); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Error updating user"})
		return
	}

	c.JSON(http.StatusOK, updatedUser)
}

func deleteUserHandler(c *gin.Context) {
	userID, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid user ID"})
		return
	}

	if err := deleteUser(userID); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Error deleting user"})
		return
	}

	c.JSON(http.StatusNoContent, nil)
}

func getUserTripsHandler(c *gin.Context) {
	userID, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid user ID"})
		return
	}

	trips, err := getUserTrips(userID)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Error retrieving user trips"})
		return
	}

	c.JSON(http.StatusOK, trips)
}
