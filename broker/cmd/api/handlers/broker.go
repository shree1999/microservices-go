package handlers

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/shree1999/broker/cmd/api/models"
	"github.com/shree1999/broker/cmd/api/utils"
)

func BrokerPingHandler() gin.HandlerFunc {
	return func(ctx *gin.Context) {
		payload := models.JsonPayload{
			Error: false,
			Message: "Broker Test",
		}

		_ = utils.WriteJSON(ctx, http.StatusOK, payload)
	}
}