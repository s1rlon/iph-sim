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

	// Define the template function map
	funcMap := template.FuncMap{
		"formatNumber": game.FormatNumber,
		"add":          game.Add,
	}

	// Load HTML templates with the function map
	r.SetFuncMap(funcMap)
	r.LoadHTMLGlob("templates/*")

	// Register routes
	routes.RegisterPlanetRoutes(r, gameInstance)
	routes.RegisterManagerRoutes(r, gameInstance)
	routes.RegisterProjectRoutes(r, gameInstance)

	r.Run(":8080") // Start the server on port 8080
}
