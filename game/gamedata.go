package game

import "database/sql"

type GameData struct {
	UpgradeHistory []UpgradeHistory
}

type UpgradeHistory struct {
	Stepnum       int
	Planet        string
	Upgradecost   float64
	Roitime       float64
	ValueIncrease float64
	TotalSpend    float64
}

func (uh *UpgradeHistory) saveToDB(db *sql.DB) error {
	query := `INSERT INTO upgrade_history (stepnum, planet, upgradecost, roitime, valueincrease, totalspend) VALUES (?, ?, ?, ?, ?, ?)`
	_, err := db.Exec(query, uh.Stepnum, uh.Planet, uh.Upgradecost, uh.Roitime, uh.ValueIncrease, uh.TotalSpend)
	return err
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

}
