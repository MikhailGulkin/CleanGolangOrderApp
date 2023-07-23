package healthcheck

import (
	"github.com/gin-gonic/gin"
	"net/http"
)

type OkMessage struct {
	Status string `json:"status"`
}
type Handler struct {
}

func (c *Handler) GetStatus(context *gin.Context) {
	context.JSON(http.StatusOK, OkMessage{Status: "OK"})
}
