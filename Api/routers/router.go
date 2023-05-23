package routers

import (
	"net/http"
	"strconv"

	"github.com/Rajendro1/AccuKnox/Api/controllers/notes"
	"github.com/Rajendro1/AccuKnox/Api/controllers/users"
	"github.com/Rajendro1/AccuKnox/config"
	"github.com/gin-gonic/gin"
)

func CORS(c *gin.Context) {

	c.Header("Access-Control-Allow-Origin", "*")
	c.Header("Access-Control-Allow-Methods", "*")
	c.Header("Access-Control-Allow-Headers", "*")

	if c.Request.Method != "OPTIONS" {

		c.Next()

	} else {
		c.AbortWithStatus(http.StatusOK)
	}
}
func HandleRequest() {
	r := gin.Default()
	gin.SetMode(gin.ReleaseMode)
	r.Use(CORS)

	r.POST("/signup", users.CreateUsers())
	r.POST("/login", users.PostLogin())

	r.GET("/notes", notes.GetNotes())
	r.POST("/notes", notes.CreateNotes())
	r.DELETE("/notes", notes.DeleteNote())

	r.Run(":" + strconv.Itoa(config.APP_HTTP_PORT))
}
