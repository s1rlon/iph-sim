package main

import (
	"github.com/gin-gonic/gin"
	"sirlon.org/iph-sim/game"
)

func main() {
	// Initialize game with planets
	gameInstance := game.NewGame()

	r := gin.Default()

	r.GET("/", func(c *gin.Context) {
		html := game.GenerateHTMLTable(gameInstance)
		c.Data(200, "text/html; charset=utf-8", []byte(html))
	})

	r.Run(":8080") // Start the server on port 8080
}
