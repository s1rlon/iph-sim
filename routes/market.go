package routes

import (
	"github.com/gin-gonic/gin"
	"sirlon.org/iph-sim/game"
)

func RegisterMarketRoutes(r *gin.Engine, gameInstance *game.Game) {

	r.GET("/market", func(c *gin.Context) {
		c.HTML(200, "market.html", nil)
	})
}
