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
		data := gameInstance.GenerateHTMLTable(steps)
		c.HTML(http.StatusOK, "planets.html", data)
	})

	r.GET("/simulate", func(c *gin.Context) {
		stepsStr := c.DefaultQuery("steps", "0")
		steps, err := strconv.Atoi(stepsStr)
		if err != nil || steps < 1 {
			steps = gameInstance.LastSteps
		}
		data := gameInstance.SimAndTable(steps)
		c.HTML(http.StatusOK, "planets.html", data)
	})

	r.GET("/reset", func(c *gin.Context) {
		gameInstance.ResetGalaxy()
		c.Redirect(http.StatusFound, "/")
	})

	r.POST("/update-colony-level", func(c *gin.Context) {
		planetName := c.PostForm("planet")
		colonyLevel, err := strconv.Atoi(c.PostForm("colonyLevel"))
		if err != nil {
			c.String(http.StatusBadRequest, "Invalid stars value")
			return
		}
		gameInstance.UpdateColonyLevel(planetName, colonyLevel)
		c.Redirect(http.StatusFound, "/")
	})

	r.POST("/update-alchemy-level", func(c *gin.Context) {
		planetName := c.PostForm("planet")
		colonyLevel, err := strconv.Atoi(c.PostForm("alchemyLevel"))
		if err != nil {
			c.String(http.StatusBadRequest, "Invalid stars value")
			return
		}
		gameInstance.UpdateAlchemyLevel(planetName, colonyLevel)
		c.Redirect(http.StatusFound, "/")
	})
}
