package game

import (
	"errors"
	"gorm.io/gorm"
)

func (g *Game) GetDB() *gorm.DB {
	return g.db
}

// Centralized Save methods

func (g *Game) SavePlanet(p *Planet) error {
	return g.db.Save(p).Error
}

func (g *Game) SaveManager(m *Manager) error {
	return g.db.Save(m).Error
}

func (g *Game) SaveProjects(p *Projects) error {
	return g.db.Save(p).Error
}

func (g *Game) SaveRooms(r *Rooms) error {
	return g.db.Save(r).Error
}

func (g *Game) SaveShips(s *Ships) error {
	return g.db.Save(s).Error
}

func (g *Game) SaveStation(s *Station) error {
	return g.db.Save(s).Error
}

func (g *Game) SaveBeacon(b *Beacon) error {
	return g.db.Save(b).Error
}

func (g *Game) SaveGameData(gd *GameData) error {
	return g.db.Save(gd).Error
}

func (g *Game) SaveUpgradeHistory(uh *UpgradeHistory) error {
	return g.db.Save(uh).Error
}

func (g *Game) SaveStar(s *Star) error {
	return g.db.Save(s).Error
}

func (g *Game) DeleteStar(name string) error {
	return g.db.Where("name = ?", name).Delete(&Star{}).Error
}

// Loader methods

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
