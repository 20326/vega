package handler

import (
	"github.com/20326/vega/app/handler/api"
	"github.com/20326/vega/app/handler/console"

	"github.com/gin-gonic/gin"
)

func NewHandlers(r *gin.Engine) {
	api.NewHandlers(r)
	console.NewHandlers(r)

	// portal

	// theme

	// static
}
