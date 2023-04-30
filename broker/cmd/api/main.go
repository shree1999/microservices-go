package main

import (
	"fmt"

	"github.com/gin-gonic/gin"
	"github.com/shree1999/broker/cmd/api/middlewares"
	"github.com/shree1999/broker/cmd/api/routes"
)

const webPort = "4000"

func main() {
	router := gin.Default()

	router.Use(middlewares.CORSMiddleware())

	routes.BrokerRoutes(router)

	err := router.Run(fmt.Sprintf(":%s", webPort))
	if err != nil {
		panic(fmt.Sprintf("Error from creating server %s", err.Error()))
	}
}
