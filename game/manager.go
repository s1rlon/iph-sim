package game

import (
	"fmt"
	"log"
)

type Role string

const (
	Pilot    Role = "Pilot"
	Miner    Role = "Miner"
	Packager Role = "Packager"
)

type SecondaryRole string

const (
	Smelt SecondaryRole = "smelt"
	Craft SecondaryRole = "craft"
	Mine  SecondaryRole = "mine"
	Speed SecondaryRole = "speed"
	Cargo SecondaryRole = "cargo"
)

type Manager struct {
	ID        int
	Stars     int
	Primary   Role
	Secondary SecondaryRole
	Planet    *Planet
}

func GetManagers(game *Game) []*Manager {
	return game.Managers
}

func AddManager(game *Game, manager *Manager) {
	insertManagerSQL := `INSERT INTO managers (stars, primary_role, secondary_role) VALUES (?, ?, ?)`
	statement, err := game.db.Prepare(insertManagerSQL)
	if err != nil {
		log.Fatal(err)
	}
	defer statement.Close()

	_, err = statement.Exec(manager.Stars, manager.Primary, manager.Secondary)
	if err != nil {
		log.Fatal(err)
	}
	game.Managers = append(game.Managers, manager)

}

func DeleteManager(game *Game, managerID int) {
	deleteManagerSQL := `DELETE FROM managers WHERE id = ?`
	statement, err := game.db.Prepare(deleteManagerSQL)
	if err != nil {
		log.Fatal(err)
	}
	defer statement.Close()

	_, err = statement.Exec(managerID)
	if err != nil {
		log.Fatal(err)
	}

	// Delete from the game.Managers slice
	for i, manager := range game.Managers {
		if manager.ID == managerID {
			game.Managers = append(game.Managers[:i], game.Managers[i+1:]...)
			break
		}
	}
}

func UpdateManagerPlanet(game *Game, managerID int, planetName string) error {
	var manager *Manager
	for _, m := range game.Managers {
		if m.ID == managerID {
			manager = m
			break
		}
	}

	if manager == nil {
		return fmt.Errorf("manager with ID %d not found", managerID)
	}

	for _, planet := range game.Planets {
		if planet.Name == planetName {
			manager.Planet = planet
			return nil
		}
	}

	return fmt.Errorf("planet with name %s not found", planetName)
}
