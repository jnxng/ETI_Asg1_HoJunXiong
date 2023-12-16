package user

import (
	"log"

	"github.com/gin-gonic/gin"
	"github.com/jnxng/eti_asg1_hojunxiong/user"
)
func StartServer() {
	router := gin.Default()

	router.POST("/users", createUserHandler)
	router.PUT("/users/:id", updateUserHandler)
	router.DELETE("/users/:id", deleteUserHandler)
	router.GET("/users/:id/trips", getUserTripsHandler)

	port := "8081"
	log.Printf("User microservice is running on :%s", port)
	err := router.Run(":" + port)
	if err != nil {
		log.Fatal("Error starting user microservice: ", err)
	}
}
func main() {
	user.StartServer()
}

