package game

import (
	"fmt"
	"log"
	"sort"
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
			manager.Planet.Manager = nil
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
			planet.Manager = manager
			return nil
		}
	}
	if manager.Planet != nil {
		manager.Planet.Manager = nil
	}
	manager.Planet = nil
	return nil
}

type PlanetManagerValue struct {
	PlanetName string
	ManagerID  int
	AddedValue float64
}

func unassignAllManagers(game *Game) {
	for _, manager := range game.Managers {
		if manager.Planet != nil {
			manager.Planet.Manager = nil
		}
		manager.Planet = nil
	}
}

func assignManagers(game *Game) {
	// Create a slice to store the added values for each planet-manager combination
	var planetManagerValues []PlanetManagerValue

	unassignAllManagers(game)

	// Calculate the added value for each planet-manager combination
	for _, manager := range game.Managers {
		for _, planet := range game.Planets {
			currentValue := planet.getMinedOresValue(planet.MiningLevel)
			planet.Manager = manager
			newValue := planet.getMinedOresValue(planet.MiningLevel)
			planet.Manager = nil
			addedValue := newValue - currentValue

			planetManagerValues = append(planetManagerValues, PlanetManagerValue{
				PlanetName: planet.Name,
				ManagerID:  manager.ID,
				AddedValue: addedValue,
			})
		}
	}

	// Sort the planet-manager combinations by added value in descending order
	sort.Slice(planetManagerValues, func(i, j int) bool {
		return planetManagerValues[i].AddedValue > planetManagerValues[j].AddedValue
	})

	// Create a map to keep track of assigned planets and managers
	assignedPlanets := make(map[string]bool)
	assignedManagers := make(map[int]bool)

	// Assign managers to the most profitable planets
	for _, pmv := range planetManagerValues {
		if assignedPlanets[pmv.PlanetName] || assignedManagers[pmv.ManagerID] {
			continue
		}

		// Assign the manager to the planet
		for _, planet := range game.Planets {
			if planet.Name == pmv.PlanetName {
				planet.Manager = nil
				for _, manager := range game.Managers {
					if manager.ID == pmv.ManagerID {
						planet.Manager = manager
						manager.Planet = planet
						assignedPlanets[planet.Name] = true
						assignedManagers[manager.ID] = true
						fmt.Printf("Assigning manager %d to planet: %s with value add of %f\n", manager.ID, planet.Name, pmv.AddedValue)
						break
					}
				}
				break
			}
		}
	}
}
