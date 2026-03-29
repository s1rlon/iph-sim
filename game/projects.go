package game

import (
	"log"
)

type Projects struct {
	ID             uint `gorm:"primaryKey;default:1"`
	TelescopeLevel int
	MiningLevel    int
	ShipSpeedLevel int
	ShipCargoLevel int
	Beacon         int
	TaxLevel       int
	SmeltSpeed     int // speed 0-1-2
	SmeltEff       int // input reduc 20
	AlloyValue     int
	SmeltSpec      int
	CraftSpeed     int
	CraftEff       int
	ItemValue      int
	CraftSpec      int
	PrefVendor     int
	OreTargeting   int
	ManTraining    int
	ManSTraing     int
	LeaderTraining int
}

func newProjects() *Projects {
	return &Projects{
		ID:             1,
		TelescopeLevel: 0,
		MiningLevel:    0,
		ShipSpeedLevel: 0,
		ShipCargoLevel: 0,
		Beacon:         0,
		TaxLevel:       0,
		SmeltSpeed:     0,
		SmeltEff:       0,
		AlloyValue:     0,
		SmeltSpec:      0,
		CraftSpeed:     0,
		CraftEff:       0,
		ItemValue:      0,
		CraftSpec:      0,
		PrefVendor:     0,
		OreTargeting:   0,
		ManTraining:    0,
		ManSTraing:     0,
		LeaderTraining: 0,
	}
}

func (g *Game) saveProjectsToDB(p *Projects) {
	p.ID = 1
	err := g.db.Save(p).Error
	if err != nil {
		log.Fatal(err)
	}
}

func (g *Game) UpdateProjects(projects *Projects) {
	g.Projects = projects
	g.saveProjectsToDB(projects)
}

func (p *Projects) telescopeRange() int {
	return 4 + p.TelescopeLevel*3
}
