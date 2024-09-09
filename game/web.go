package game

func ResetGalaxy(game *Game) {
	game.ResetPlanets()
	game.ResetManagers()
}

func GenerateHTMLTable(game *Game, steps int) interface{} {
	// Simulate the upgrades
	SimulateUpgrades(game, steps)

	// Create table data
	data := CreateTableData(game)

	return data
}
