package routes

import (
	"github.com/gin-gonic/gin"
	"net/http"
)

func SetupDefaultRoutes(r *gin.Engine) {

	invalidRoute := struct {
		Message string `json:"message"`
		Status  string `json:"status"`
	}{
		"invalid route",
		"OK",
	}

	r.NoRoute(func(c *gin.Context) {
		c.IndentedJSON(http.StatusNotFound, invalidRoute)
	})
}
