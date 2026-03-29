package game

import (
	"errors"
	"gorm.io/gorm"
)

type GameData struct {
	ID             uint             `gorm:"primaryKey;default:1"`
	UpgradeHistory []UpgradeHistory `gorm:"foreignKey:GameDataID"`
	Smelters       int              `gorm:"default:1"`
	Crafters       int              `gorm:"default:1"`
	ManagerSlots   int              `gorm:"default:2"`
}

type UpgradeHistory struct {
	ID            uint    `gorm:"primaryKey;autoIncrement"`
	GameDataID    uint    `gorm:"index"`
	Stepnum       int     `gorm:"not null"`
	PlanetName    string  `gorm:"not null"`
	Planet        *Planet `gorm:"foreignKey:PlanetName;references:Name"`
	Upgradecost   float64
	Roitime       float64
	ValueIncrease float64
	TotalSpend    float64
}

func NewGameData() *GameData {
	return &GameData{UpgradeHistory: []UpgradeHistory{}, Smelters: 1, Crafters: 1, ManagerSlots: 2}
}

func (g *Game) loadGameDataFromDB() *GameData {
	var gd GameData
	err := g.db.Preload("UpgradeHistory").First(&gd, 1).Error
	if err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return NewGameData()
		}
	}
	return &gd
}

func (uh *UpgradeHistory) saveToDB(g *Game) error {
	return g.db.Save(uh).Error
}

func (gd *GameData) SyncDB(g *Game) {
	g.db.Save(gd)
}

func (gd *GameData) LoadUpgradeHistoryFromDB(g *Game) error {
	return g.db.Model(gd).Association("UpgradeHistory").Find(&gd.UpgradeHistory)
}

func (gd *GameData) resetGameData() {
	gd.UpgradeHistory = []UpgradeHistory{}
	DB.Where("game_data_id = ?", gd.ID).Delete(&UpgradeHistory{})
	gd.Crafters = 1
	gd.Smelters = 1
	DB.Save(gd)
}
