package main

import (
	"html/template"

	"github.com/gin-gonic/gin"
	"sirlon.org/iph-sim/game"
	"sirlon.org/iph-sim/routes"
)

func main() {
	// Initialize game with planets
	gameInstance := game.NewGame()

	r := gin.Default()

	funcMap := template.FuncMap{
		"formatNumber": game.FormatNumber,
	}

	// Load HTML templates with the function map
	r.SetFuncMap(funcMap)
	r.LoadHTMLGlob("templates/*")

	routes.RegisterPlanetRoutes(r, gameInstance)
	routes.RegisterManagerRoutes(r, gameInstance)

	r.Run(":8080") // Start the server on port 8080
}
