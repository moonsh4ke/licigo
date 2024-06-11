package routes

import (
	"context"
	"example/licigo/config"

	"github.com/gin-gonic/gin"
)

var router = gin.Default()

var client = config.GetMongoDbClient()

// Run will start the server
func Run() {
	getRoutes()
	router.Run("localhost:8080")
	defer func() {
		if err := client.Disconnect(context.TODO()); err != nil {
			panic(err)
		}
	}()
}

func getRoutes() {
	v1 := router.Group("/v1")
	addTenderRoutes(v1)

	// v2 := router.Group("/v2")
	// addPingRoutes(v2)
}
