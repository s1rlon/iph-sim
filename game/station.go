package game

import (
	"errors"
	"log"
	"gorm.io/gorm"
)

type Station struct {
	ID                uint `gorm:"primaryKey;default:1"`
	MineBoost         float64
	SpeedBoost        float64
	CargoBoost        float64
	SmeltBoost        float64
	CraftBoost        float64
	ItemBoost         float64
	MarketBoost       float64
	ManagerBoost      float64
	AsteroidValue     float64
	Colonizing        float64
	ProductionBoost   float64
	PlanetUpgradeCost float64
	ColonyCostRedux   float64
}

func newStation() *Station {
	return &Station{
		ID:                1,
		MineBoost:         1,
		SpeedBoost:        1,
		CargoBoost:        1,
		SmeltBoost:        1,
		CraftBoost:        1,
		ItemBoost:         1,
		MarketBoost:       1,
		ManagerBoost:      1,
		AsteroidValue:     1,
		Colonizing:        1,
		ProductionBoost:   1,
		PlanetUpgradeCost: 1,
		ColonyCostRedux:   1,
	}
}

func (g *Game) loadStationDataFromDB() *Station {
	var s Station
	err := g.db.First(&s, 1).Error
	if err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return newStation()
		}
	}
	return &s
}

func (g *Game) saveStationDataToDB(station *Station) {
	station.ID = 1
	err := g.SaveStation(station)
	if err != nil {
		log.Fatal(err)
	}
}

func (g *Game) UpdateStation(station *Station) {
	g.Station = station
	g.saveStationDataToDB(station)
}
