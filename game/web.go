package game

func (game *Game) GenerateHTMLTable(steps int) interface{} {
	// Create table data
	data := game.CreateTableData()

	return data
}

func (g *Game) SimAndTable(steps int) interface{} {
	// Simulate the upgrades
	SimulateUpgrades(g, steps)

	// Create table data
	data := g.CreateTableData()

	return data
}
