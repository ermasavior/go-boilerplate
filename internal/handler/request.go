package handler

import (
	"encoding/json"

	"github.com/gin-gonic/gin"
)

func getRequestBody(c *gin.Context, req interface{}) error {
	if err := json.NewDecoder(c.Request.Body).Decode(req); err != nil {
		return err
	}

	return nil
}
