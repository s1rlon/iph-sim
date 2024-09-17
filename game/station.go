package game

import (
	"database/sql"
	"encoding/json"
	"log"
)

type Station struct {
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

func loadStationDataFromDB(db *sql.DB) *Station {
	var jsonString string
	querySQL := `SELECT station FROM station WHERE id = 1`
	err := db.QueryRow(querySQL).Scan(&jsonString)
	if err != nil {
		if err == sql.ErrNoRows {
			return newStation()
		}
		log.Fatal(err)
	}
	var s Station
	err = json.Unmarshal([]byte(jsonString), &s)
	if err != nil {
		log.Fatal(err)
	}
	return &s
}

func (g *Game) saveStationDataToDB(station *Station) {
	jsonString, err := json.Marshal(station)
	if err != nil {
		log.Fatal(err)
	}
	querySQL := `INSERT OR REPLACE INTO station (id, station) VALUES (1, ?)`
	_, err = g.db.Exec(querySQL, jsonString)
	if err != nil {
		log.Fatal(err)
	}
}

func (g *Game) UpdateStation(station *Station) {
	g.Station = station
	g.saveStationDataToDB(station)
}
