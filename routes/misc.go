package routes

import (
	"net/http"

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
		c.HTML(200, "ships.html", gameInstance.Ships)
	})

	r.GET("/beacon", func(c *gin.Context) {
		c.HTML(200, "beacon.html", nil)
	})

	r.POST("/update-ships", func(c *gin.Context) {
		ships := &game.Ships{
			AdShip:       c.PostForm("AdShip") == "on",
			Daugtership:  c.PostForm("Daugtership") == "on",
			Eldership:    c.PostForm("Eldership") == "on",
			Aurora:       c.PostForm("Aurora") == "on",
			Enigma:       c.PostForm("Enigma") == "on",
			Exodus:       c.PostForm("Exodus") == "on",
			Merchant:     c.PostForm("Merchant") == "on",
			Thunderhorse: c.PostForm("Thunderhorse") == "on",
		}

		// Update the game instance with the new ships data
		gameInstance.UpdateShips(ships)

		// Redirect back to the ships page
		c.Redirect(http.StatusSeeOther, "/ships")
	})
}
