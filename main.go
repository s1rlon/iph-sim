package main

import (
	"strconv"

	"github.com/gin-gonic/gin"
	"sirlon.org/iph-sim/game"
)

func main() {
	// Initialize game with planets
	gameInstance := game.NewGame()

	r := gin.Default()

	r.GET("/", func(c *gin.Context) {
		stepsStr := c.DefaultQuery("steps", "0")
		steps, err := strconv.Atoi(stepsStr)
		if err != nil || steps < 1 {
			steps = 0
		}
		html := game.GenerateHTMLTable(gameInstance, steps)
		c.Data(200, "text/html; charset=utf-8", []byte(html))
	})

	r.GET("/reset", func(c *gin.Context) {
		game.ResetMiningLevels(gameInstance)
		c.Redirect(302, "/")
	})

	r.Run(":8080") // Start the server on port 8080
}
