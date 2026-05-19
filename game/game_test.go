package game

import (
	"os"
	"testing"
)

func TestNewGame(t *testing.T) {
	// Change working directory to project root so it can find data/ and ipm2.sql
	// In a real test, we'd use a temporary DB or mock data.
	// For this smoke test, we'll just check if it loads.
	err := os.Chdir("..")
	if err != nil {
		t.Fatalf("Failed to change directory: %v", err)
	}

	game := NewGame()
	game.InitData()

	if len(game.Ores) != 20 {
		t.Errorf("Expected 20 ores, got %d", len(game.Ores))
	}

	if len(game.Planets) != 41 {
		t.Errorf("Expected 41 planets, got %d", len(game.Planets))
	}

	if len(game.Alloys) != 6 {
		t.Errorf("Expected 6 alloys, got %d", len(game.Alloys))
	}

	if len(game.Items) != 13 {
		t.Errorf("Expected 13 items, got %d", len(game.Items))
	}

	if len(game.Recepies) != 11 {
		t.Errorf("Expected 11 recipes, got %d", len(game.Recepies))
	}

	// Verify planet-game pointer
	if game.Planets[0].game != game {
		t.Errorf("Planet[0] game pointer is not correctly set")
	}

	// Verify ore-game pointer
	if game.Ores[0].game != game {
		t.Errorf("Ore[0] game pointer is not correctly set")
	}
}

func TestSimulateUpgradesSmoke(t *testing.T) {
	// Already in root from previous test if run in sequence, but let's be safe
	// Note: go test runs tests in a specific way, but here we're just checking loading.
	
	game := NewGame()
	game.InitData()

	// Unlock first planet
	game.Planets[0].Locked = false
	game.SavePlanet(game.Planets[0])

	// Run a small simulation
	SimulateUpgrades(game, 5)

	if len(game.GameData.UpgradeHistory) < 5 {
		t.Errorf("Expected at least 5 history entries, got %d", len(game.GameData.UpgradeHistory))
	}
}
