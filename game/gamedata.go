package game

import (
	"database/sql"
	"log"
)

type GameData struct {
	UpgradeHistory []UpgradeHistory
	Smelters       int
	Crafters       int
	ManagerSlots   int
}

type UpgradeHistory struct {
	Stepnum       int
	Planet        string
	Upgradecost   float64
	Roitime       float64
	ValueIncrease float64
	TotalSpend    float64
}

func NewGameData() *GameData {
	return &GameData{UpgradeHistory: []UpgradeHistory{}, Smelters: 1, Crafters: 1, ManagerSlots: 2}
}

func loadGameDataFromDB(db *sql.DB) *GameData {
	gd := NewGameData()

	// Load Smelters, Crafters, and ManagerSlots from the database
	query := `SELECT smelters, crafters, managerslots FROM gamedata WHERE id = 1`
	row := db.QueryRow(query)
	err := row.Scan(&gd.Smelters, &gd.Crafters, &gd.ManagerSlots)
	if err != nil {
		if err == sql.ErrNoRows {
			// No rows found, initialize with default values
			gd.Smelters = 1
			gd.Crafters = 1
			gd.ManagerSlots = 2
		} else {
			log.Fatal(err)
		}
	}
	err = gd.LoadUpgradeHistoryFromDB(db)
	if err != nil {
		log.Fatal(err)
	}
	return gd
}

func (uh *UpgradeHistory) saveToDB(db *sql.DB) error {
	query := `INSERT INTO upgrade_history (stepnum, planet, upgradecost, roitime, valueincrease, totalspend) VALUES (?, ?, ?, ?, ?, ?)`
	_, err := db.Exec(query, uh.Stepnum, uh.Planet, uh.Upgradecost, uh.Roitime, uh.ValueIncrease, uh.TotalSpend)
	return err
}

func (gd *GameData) SyncDB(db *sql.DB) {
	_, err := db.Exec("DELETE FROM gamedata")
	if err != nil {
		log.Fatal(err)
	}
	query := `INSERT OR REPLACE INTO gamedata (id, smelters, crafters, managerslots) VALUES (1, ?, ?, ?)`
	_, err = db.Exec(query, gd.Smelters, gd.Crafters, gd.ManagerSlots)
	if err != nil {
		log.Fatal(err)
	}
}

func (gd *GameData) LoadUpgradeHistoryFromDB(db *sql.DB) error {
	query := `SELECT stepnum, planet, upgradecost, roitime, valueincrease, totalspend FROM upgrade_history ORDER BY stepnum`
	rows, err := db.Query(query)
	if err != nil {
		return err
	}
	defer rows.Close()

	var history []UpgradeHistory
	for rows.Next() {
		var uh UpgradeHistory
		err := rows.Scan(&uh.Stepnum, &uh.Planet, &uh.Upgradecost, &uh.Roitime, &uh.ValueIncrease, &uh.TotalSpend)
		if err != nil {
			return err
		}
		history = append(history, uh)
	}

	if err = rows.Err(); err != nil {
		return err
	}

	gd.UpgradeHistory = history
	return nil
}

func (gd *GameData) resetGameData() {
	gd.UpgradeHistory = []UpgradeHistory{}
	_, err := DB.Exec("DELETE FROM upgrade_history")
	if err != nil {
		log.Fatal(err)
	}
	gd.Crafters = 1
	gd.Smelters = 1
}
