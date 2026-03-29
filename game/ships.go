package game

import (
	"errors"
	"log"
	"gorm.io/gorm"
)

type Ships struct {
	ID           uint `gorm:"primaryKey;default:1"`
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
		ID:           1,
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
	ships.ID = 1
	err := g.db.Save(ships).Error
	if err != nil {
		log.Fatal(err)
	}
}

func (g *Game) loadShipsFromDB() *Ships {
	var ships Ships
	err := g.db.First(&ships, 1).Error
	if err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return NewShips()
		}
	}
	return &ships
}
