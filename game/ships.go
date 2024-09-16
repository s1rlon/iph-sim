package game

import (
	"database/sql"
	"log"
)

type Ships struct {
	AdShip       bool
	Daugtership  bool
	Eldership    bool
	Aurora       bool
	Enigma       bool
	Exodus       bool
	Merchant     bool
	Thunderhorse bool
}

func NewShips() *Ships {
	return &Ships{
		AdShip:       false,
		Daugtership:  false,
		Eldership:    false,
		Aurora:       false,
		Enigma:       false,
		Exodus:       false,
		Merchant:     false,
		Thunderhorse: false,
	}
}

func (g *Game) UpdateShips(ships *Ships) {
	g.Ships = ships
	err := saveShipsToDB(g.db, ships)
	if err != nil {
		log.Fatal(err)
	}
}

func loadShipsFromDB(db *sql.DB) *Ships {
	query := `SELECT ad_ship, daugtership, eldership, aurora, enigma, exodus, merchant, thunderhorse FROM ships ORDER BY id DESC LIMIT 1`
	row := db.QueryRow(query)

	ships := NewShips()
	err := row.Scan(&ships.AdShip, &ships.Daugtership, &ships.Eldership, &ships.Aurora, &ships.Enigma, &ships.Exodus, &ships.Merchant, &ships.Thunderhorse)
	if err != nil {
		if err == sql.ErrNoRows {
			return NewShips()
		}
		log.Fatal(err)
	}
	return ships
}

func saveShipsToDB(db *sql.DB, ships *Ships) error {
	_, err := db.Exec("DELETE FROM ships")
	if err != nil {
		log.Fatal(err)
	}
	query := `INSERT INTO ships (ad_ship, daugtership, eldership, aurora, enigma, exodus, merchant, thunderhorse) VALUES (?, ?, ?, ?, ?, ?, ?, ?)`
	_, err = db.Exec(query, ships.AdShip, ships.Daugtership, ships.Eldership, ships.Aurora, ships.Enigma, ships.Exodus, ships.Merchant, ships.Thunderhorse)
	return err
}
