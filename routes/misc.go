package routes

import (
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
	"sirlon.org/iph-sim/game"
)

func RegisterMiscRoutes(r *gin.Engine, gameInstance *game.Game) {

	r.GET("/items", func(c *gin.Context) {
		data := struct {
			CraftingData []*game.CraftingData
			GameData     *game.GameData
		}{
			CraftingData: gameInstance.MakeCraftingData(),
			GameData:     gameInstance.GameData,
		}
		c.HTML(200, "items.html", data)
	})

	r.POST("/updateCrafters", func(c *gin.Context) {
		// Parse the form values
		smelters := parseFormValue(c, "smelters")
		crafters := parseFormValue(c, "crafters")
		gameInstance.GameData.Smelters = smelters
		gameInstance.GameData.Crafters = crafters
		c.Redirect(http.StatusSeeOther, "/items")
	})

	r.GET("/rooms", func(c *gin.Context) {
		c.HTML(200, "rooms.html", gameInstance.Rooms)
	})

	r.POST("/updateRooms", func(c *gin.Context) {
		rooms := &game.Rooms{
			Engineering:     parseFormValue(c, "Engineering"),
			Aeronautical:    parseFormValue(c, "Aeronautical"),
			Packaging:       parseFormValue(c, "Packaging"),
			Forge:           parseFormValue(c, "Forge"),
			Workshop:        parseFormValue(c, "Workshop"),
			Astronomy:       parseFormValue(c, "Astronomy"),
			Laboratory:      parseFormValue(c, "Laboratory"),
			Terrarium:       parseFormValue(c, "Terrarium"),
			Lounge:          parseFormValue(c, "Lounge"),
			Robotics:        parseFormValue(c, "Robotics"),
			BackupGenerator: parseFormValue(c, "BackupGenerator"),
			Underforge:      parseFormValue(c, "Underforge"),
			Dorm:            parseFormValue(c, "Dorm"),
			Sales:           parseFormValue(c, "Sales"),
			Classroom:       parseFormValue(c, "Classroom"),
			Marketing:       parseFormValue(c, "Marketing"),
		}

		gameInstance.UpdateRooms(rooms)

		c.Redirect(http.StatusSeeOther, "/rooms")
	})

	r.GET("/ships", func(c *gin.Context) {
		c.HTML(200, "ships.html", gameInstance.Ships)
	})

	r.GET("/beacon", func(c *gin.Context) {
		c.HTML(200, "beacon.html", gameInstance.Beacon)
	})

	r.POST("/updateBeaconLevels", func(c *gin.Context) {
		var levels []float64
		for i := 0; i < 21; i++ {
			levelStr := c.PostForm("levels[" + strconv.Itoa(i) + "]")
			level, err := strconv.ParseFloat(levelStr, 64)
			if err != nil {
				c.String(http.StatusBadRequest, "Invalid level value")
				return
			}
			levels = append(levels, level)
		}
		gameInstance.UpdateBeacon(levels)
		c.Redirect(http.StatusFound, "/beacon")
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

	r.GET("/station", func(c *gin.Context) {
		c.HTML(200, "station.html", gameInstance.Station)
	})
}
