package game

import (
	"database/sql"
	"encoding/json"
	"log"
)

type Beacon struct {
	Levels []float64
}

func newBeacon() *Beacon {
	levels := make([]float64, 21)
	for i := range levels {
		levels[i] = 1
	}
	return &Beacon{Levels: levels}
}

func loadBeaconDataFromDB(db *sql.DB) *Beacon {
	var jsonString string
	querySQL := `SELECT beacon FROM beacon WHERE id = 1`
	err := db.QueryRow(querySQL).Scan(&jsonString)
	if err != nil {
		if err == sql.ErrNoRows {
			return newBeacon()
		}
		log.Fatal(err)
	}
	var b Beacon
	// Deserialize the JSON string back into an array
	err = json.Unmarshal([]byte(jsonString), &b.Levels)
	if err != nil {
		log.Fatal(err)
	}
	return &b
}

func (g *Game) saveBeaconLevelsToDB(levels []float64) {
	// Serialize the array into a JSON string
	jsonString, err := json.Marshal(levels)
	if err != nil {
		log.Fatal(err)
	}
	querySQL := `INSERT OR REPLACE INTO beacon (id, beacon) VALUES (1, ?)`
	_, err = g.db.Exec(querySQL, jsonString)
	if err != nil {
		log.Fatal(err)
	}
}

func (g *Game) UpdateBeacon(levels []float64) {
	g.Beacon.Levels = levels
	g.saveBeaconLevelsToDB(levels)
}
