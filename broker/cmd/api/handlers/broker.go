package handlers

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/shree1999/broker/cmd/api/models"
)

func BrokerPingHandler() gin.HandlerFunc {
	return func(ctx *gin.Context) {
		payload := models.JsonPayload{
			Error: false,
			Message: "Broker Test",
		}

		ctx.Header("Content-Type", "application/json")
		ctx.JSON(http.StatusOK, payload)
	}
}