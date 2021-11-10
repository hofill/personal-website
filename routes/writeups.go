package routes

import (
	"github.com/gin-gonic/gin"
	"hofill/models"
	"net/http"
)

func getAllWriteUps(c *gin.Context) {

	writeUpPreview := models.WriteUpPreview{}
	writeUpPreview.Event = c.Query("event_name")
	c.IndentedJSON(http.StatusOK, writeUpPreview)

}

func SetupWriteUpRoutes(r *gin.Engine) {
	r.GET("/writeups", getAllWriteUps)
}
