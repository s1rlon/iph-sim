package game

func GenerateHTMLTable(game *Game, steps int) interface{} {
	// Create table data
	data := CreateTableData(game)

	return data
}

func GameSim(game *Game, steps int) interface{} {
	// Simulate the upgrades
	SimulateUpgrades(game, steps)

	// Create table data
	data := CreateTableData(game)

	return data
}
