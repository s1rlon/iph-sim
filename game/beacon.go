package game

import (
	"encoding/json"
	"errors"
	"log"
	"gorm.io/gorm"
)

type Beacon struct {
	ID     uint      `gorm:"primaryKey;default:1"`
	Levels []float64 `gorm:"type:text"`
}

func (b *Beacon) BeforeSave(tx *gorm.DB) error {
	return nil
}

func newBeacon() *Beacon {
	levels := make([]float64, 21)
	for i := range levels {
		levels[i] = 1
	}
	return &Beacon{ID: 1, Levels: levels}
}

func (g *Game) loadBeaconDataFromDB() *Beacon {
	var b struct {
		ID     uint
		Levels string
	}
	err := g.db.Table("beacons").First(&b, 1).Error
	if err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return newBeacon()
		}
	}

	var levels []float64
	err = json.Unmarshal([]byte(b.Levels), &levels)
	if err != nil {
		return newBeacon()
	}
	return &Beacon{ID: b.ID, Levels: levels}
}

func (g *Game) saveBeaconLevelsToDB(levels []float64) {
	jsonString, err := json.Marshal(levels)
	if err != nil {
		log.Fatal(err)
	}

	err = g.db.Table("beacons").Save(map[string]interface{}{"id": 1, "levels": string(jsonString)}).Error
	if err != nil {
		log.Fatal(err)
	}
}

func (g *Game) UpdateBeacon(levels []float64) {
	g.Beacon.Levels = levels
	g.saveBeaconLevelsToDB(levels)
}
