package routes

import (
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
	"sirlon.org/iph-sim/game"
)

func RegisterPlanetRoutes(r *gin.Engine, gameInstance *game.Game) {
	r.GET("/", func(c *gin.Context) {
		stepsStr := c.DefaultQuery("steps", "0")
		steps, err := strconv.Atoi(stepsStr)
		if err != nil || steps < 1 {
			steps = gameInstance.LastSteps
		}
		data := game.GenerateHTMLTable(gameInstance, steps)
		c.HTML(http.StatusOK, "planets.html", data)
	})

	r.GET("/reset", func(c *gin.Context) {
		gameInstance.ResetPlanets()
		c.Redirect(http.StatusFound, "/")
	})
}
