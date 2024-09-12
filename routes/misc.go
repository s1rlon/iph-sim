package routes

import (
	"github.com/gin-gonic/gin"
	"sirlon.org/iph-sim/game"
)

func RegisterMiscRoutes(r *gin.Engine, gameInstance *game.Game) {

	r.GET("/items", func(c *gin.Context) {
		c.HTML(200, "items.html", nil)
	})

	r.GET("/rooms", func(c *gin.Context) {
		c.HTML(200, "rooms.html", nil)
	})

	r.GET("/ships", func(c *gin.Context) {
		c.HTML(200, "ships.html", nil)
	})

	r.GET("/beacon", func(c *gin.Context) {
		c.HTML(200, "beacon.html", nil)
	})
}
