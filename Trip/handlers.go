package trip

import (
	"net/http"
	"strconv"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/jnxng/eti_asg1_hojunxiong/common"
)

func createTripHandler(c *gin.Context) {
	var trip Trip
	if err := c.ShouldBindJSON(&trip); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	trip.StartTime, _ = time.Parse(common.TimeFormat, c.PostForm("start_time"))

	if err := createTrip(&trip); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Error creating trip"})
		return
	}

	c.JSON(http.StatusCreated, trip)
}

func enrollTripHandler(c *gin.Context) {
	userID, err := strconv.Atoi(c.PostForm("user_id"))
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid user ID"})
		return
	}

	tripID, err := strconv.Atoi(c.PostForm("trip_id"))
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid trip ID"})
		return
	}

	if err := enrollTrip(userID, tripID); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Error enrolling in the trip"})
		return
	}

	c.JSON(http.StatusOK, gin.H{"message": "Enrolled in the trip successfully"})
}

func startTripHandler(c *gin.Context) {
	tripID, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid trip ID"})
		return
	}

	if err := startTrip(tripID); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Error starting the trip"})
		return
	}

	c.JSON(http.StatusOK, gin.H{"message": "Trip started successfully"})
}

func cancelTripHandler(c *gin.Context) {
	tripID, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid trip ID"})
		return
	}

	if err := cancelTrip(tripID); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Error canceling the trip"})
		return
	}

	c.JSON(http.StatusOK, gin.H{"message": "Trip canceled successfully"})
}
