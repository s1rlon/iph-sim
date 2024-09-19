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
	Smelt SecondaryRole = "Smelt"
	Craft SecondaryRole = "Craft"
	Mine  SecondaryRole = "Mine"
	Speed SecondaryRole = "Speed"
	Cargo SecondaryRole = "Cargo"
)

type Manager struct {
	ID        int
	Stars     int
	Primary   Role
	Secondary SecondaryRole
	Planet    *Planet
}

func (g *Game) GetManagers() []*Manager {
	return g.Managers
}

func (g *Game) AddManager(manager *Manager) {
	insertManagerSQL := `INSERT INTO managers (stars, primary_role, secondary_role) VALUES (?, ?, ?)`
	statement, err := g.db.Prepare(insertManagerSQL)
	if err != nil {
		log.Fatal(err)
	}
	defer statement.Close()

	_, err = statement.Exec(manager.Stars, manager.Primary, manager.Secondary)
	if err != nil {
		log.Fatal(err)
	}
	g.Managers = append(g.Managers, manager)

}

func (m *Manager) unassignManager() {
	if m.Planet != nil {
		m.Planet.Manager = nil
	}
	m.Planet = nil
}

func (g *Game) DeleteManager(managerID int) {
	deleteManagerSQL := `DELETE FROM managers WHERE id = ?`
	statement, err := g.db.Prepare(deleteManagerSQL)
	if err != nil {
		log.Fatal(err)
	}
	defer statement.Close()

	_, err = statement.Exec(managerID)
	if err != nil {
		log.Fatal(err)
	}

	// Delete from the game.Managers slice
	for i, manager := range g.Managers {
		if manager.ID == managerID {
			manager.unassignManager()
			g.Managers = append(g.Managers[:i], g.Managers[i+1:]...)
			break
		}
	}
}

func (g *Game) UpdateManagerPlanet(managerID int, planetName string) error {
	var manager *Manager
	for _, m := range g.Managers {
		if m.ID == managerID {
			manager = m
			break
		}
	}

	if manager == nil {
		return fmt.Errorf("manager with ID %d not found", managerID)
	}

	for _, planet := range g.Planets {
		if planet.Name == planetName {
			manager.unassignManager()
			manager.Planet = planet
			planet.Manager = manager
			return nil
		}
	}
	if manager.Planet != nil {
		manager.unassignManager()
	}
	return nil
}

type PlanetManagerValue struct {
	PlanetName string
	ManagerID  int
	AddedValue float64
}

func (game *Game) unassignAllManagers() {
	for _, manager := range game.Managers {
		manager.unassignManager()
	}
}

func (g *Game) AssignManagers() {
	// Create slices to store the added values for each planet-manager combination
	var minerManagerValues []PlanetManagerValue
	var otherManagerValues []PlanetManagerValue

	g.unassignAllManagers()

	// Separate managers by role and calculate the added value for each planet-manager combination
	for _, manager := range g.Managers {
		for _, planet := range g.Planets {
			currentValue := planet.getMinedOresValue(planet.MiningLevel)
			planet.Manager = manager
			newValue := planet.getMinedOresValue(planet.MiningLevel)
			planet.Manager = nil
			addedValue := newValue - currentValue

			pmv := PlanetManagerValue{
				PlanetName: planet.Name,
				ManagerID:  manager.ID,
				AddedValue: addedValue,
			}

			if manager.Primary == Miner {
				minerManagerValues = append(minerManagerValues, pmv)
			} else {
				otherManagerValues = append(otherManagerValues, pmv)
			}
		}
	}

	// Sort the planet-manager combinations by added value in descending order
	sort.Slice(minerManagerValues, func(i, j int) bool {
		return minerManagerValues[i].AddedValue > minerManagerValues[j].AddedValue
	})

	// Sort the other managers by added value in descending order, then by distance (index) in descending order, and then by stars in descending order
	sort.Slice(otherManagerValues, func(i, j int) bool {
		if otherManagerValues[i].AddedValue == otherManagerValues[j].AddedValue {
			indexI := g.getPlanetIndexByName(otherManagerValues[i].PlanetName)
			indexJ := g.getPlanetIndexByName(otherManagerValues[j].PlanetName)
			if indexI == indexJ {
				managerI := g.getManagerByID(otherManagerValues[i].ManagerID)
				managerJ := g.getManagerByID(otherManagerValues[j].ManagerID)
				return managerI.Stars > managerJ.Stars
			}
			return indexI > indexJ
		}
		return otherManagerValues[i].AddedValue > otherManagerValues[j].AddedValue
	})

	// Create a map to keep track of assigned planets and managers
	assignedPlanets := make(map[string]bool)
	assignedManagers := make(map[int]bool)

	// Assign managers to the most profitable planets up to the ManagerSlots limit
	managerSlots := g.GameData.ManagerSlots
	assignedCount := 0

	// Assign "Miner" managers first
	for _, pmv := range minerManagerValues {
		if assignedCount >= managerSlots {
			break
		}
		if assignedPlanets[pmv.PlanetName] || assignedManagers[pmv.ManagerID] {
			continue
		}

		// Assign the manager to the planet
		for _, planet := range g.Planets {
			if planet.Name == pmv.PlanetName && !planet.Locked {
				planet.Manager = nil
				for _, manager := range g.Managers {
					if manager.ID == pmv.ManagerID {
						planet.Manager = manager
						manager.Planet = planet
						assignedPlanets[planet.Name] = true
						assignedManagers[manager.ID] = true
						assignedCount++
						fmt.Printf("Assigning Miner manager %d to planet: %s with value add of %f\n", manager.ID, planet.Name, pmv.AddedValue)
						break
					}
				}
				break
			}
		}
	}

	// Assign other managers
	for _, pmv := range otherManagerValues {
		if assignedCount >= managerSlots {
			break
		}
		if assignedPlanets[pmv.PlanetName] || assignedManagers[pmv.ManagerID] {
			continue
		}

		// Assign the manager to the planet
		for _, planet := range g.Planets {
			if planet.Name == pmv.PlanetName && !planet.Locked {
				planet.Manager = nil
				for _, manager := range g.Managers {
					if manager.ID == pmv.ManagerID {
						planet.Manager = manager
						manager.Planet = planet
						assignedPlanets[planet.Name] = true
						assignedManagers[manager.ID] = true
						assignedCount++
						fmt.Printf("Assigning non-Miner manager %d to planet: %s with value add of %f\n", manager.ID, planet.Name, pmv.AddedValue)
						break
					}
				}
				break
			}
		}
	}
}

func (g *Game) getManagerByID(managerID int) *Manager {
	for _, manager := range g.Managers {
		if manager.ID == managerID {
			return manager
		}
	}
	return nil
}
