package routes

import (
	"github.com/gin-gonic/gin"
	"github.com/shree1999/broker/cmd/api/handlers"
)

func BrokerRoutes(route *gin.Engine) {
	brokerRoutes := route.Group("/api/v1/broker")

	brokerRoutes.GET("/ping", handlers.BrokerPingHandler())
}
