package utils

import (
	"encoding/json"
	"net/http"

	"github.com/gin-gonic/gin"
)

func ReadJSON(ctx *gin.Context, data any) error {
	maxBytes := 1048576

	ctx.Request.Body = http.MaxBytesReader(ctx.Writer, ctx.Request.Body, int64(maxBytes))

	dec := json.NewDecoder(ctx.Request.Body)
	err := dec.Decode(data)
	if err != nil {
		return err
	}

	err = dec.Decode(&struct{}{})
	if err != nil {
		return err
	}

	return nil
}

func WriteJSON(ctx *gin.Context, status int, data any, headers ...http.Header) error {
	out, err := json.Marshal(data)
	if err != nil {
		return err
	}

	if len(headers) > 0 {
		for key, value := range headers[0] {
			ctx.Writer.Header()[key] = value
		}
	}

	ctx.Header("Content-Type", "application/json")
	ctx.Status(status)
	ctx.Writer.Write(out)
	return nil
}