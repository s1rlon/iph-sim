package routes

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"sirlon.org/iph-sim/game"
)

func RegisterProjectRoutes(r *gin.Engine, gameInstance *game.Game) {
	r.GET("/projects", func(c *gin.Context) {
		c.HTML(http.StatusOK, "projects.html", gameInstance.Projects)
	})
}
