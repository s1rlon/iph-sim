package game

import (
	"errors"
	"gorm.io/gorm"
)

func (g *Game) GetDB() *gorm.DB {
	return g.db
}

func (g *Game) resetPlanetDB(planet *Planet) {
	planet.MiningLevel = 1
	planet.ShipSpeedLeve1 = 1
	planet.ShipCargoLevel = 1
	planet.ColonyLevel = 0
	planet.Locked = true
	planet.AlchemyLevel = 1
	planet.Rover = false
	g.db.Save(planet)
}

func (g *Game) updatePlanetDB(planet *Planet) {
	g.db.Save(planet)
}

func (g *Game) getManagersFromDB() []*Manager {
	var managers []*Manager
	g.db.Preload("Planet").Find(&managers)
	return managers
}

func (g *Game) getPlanetsFromDB() ([]Planet, error) {
	var planets []Planet
	err := g.db.Find(&planets).Error
	return planets, err
}

func (g *Game) loadProjectsFromDB() *Projects {
	var p Projects
	err := g.db.First(&p, 1).Error
	if err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return newProjects()
		}
	}
	return &p
}
