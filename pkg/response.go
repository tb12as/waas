package pkg

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

type JsonResponse struct {
	Status  int    `json:"status"`
	Message string `json:"message"`
	Data    any    `json:"data"`
}

func Ok(c *gin.Context, data any) {
	response := JsonResponse{
		Status:  http.StatusOK,
		Message: "Success",
		Data:    data,
	}
	c.JSON(http.StatusOK, response)
}

func Fail(c *gin.Context, message string) {
	response := JsonResponse{
		Status:  http.StatusBadRequest,
		Message: message,
	}
	c.JSON(http.StatusBadRequest, response)
}
