package trip

import (
	"log"

	"github.com/gin-gonic/gin"
	"github.com/jnxng/ETI_Asg1_HoJunXiong/trip"
)
func StartServer() {
	router := gin.Default()

	router.POST("/trips", createTripHandler)
	router.POST("/trips/enroll", enrollTripHandler)
	router.PUT("/trips/:id/start", startTripHandler)
	router.DELETE("/trips/:id", cancelTripHandler)

	port := "8082"
	log.Printf("Trip microservice is running on :%s", port)
	err := router.Run(":" + port)
	if err != nil {
		log.Fatal("Error starting trip microservice: ", err)
	}
}

func main() {
	trip.StartServer()
}
