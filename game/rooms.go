package game

import (
	"errors"
	"log"
	"gorm.io/gorm"
)

type Rooms struct {
	ID              uint `gorm:"primaryKey;default:1"`
	Engineering     int // Mine speed
	Aeronautical    int // Ship speed
	Packaging       int // Cargo
	Forge           int // Smelt speed
	Workshop        int // Craft speed
	Astronomy       int // Upgrade discount
	Laboratory      int // Project discount
	Terrarium       int // Colonization discount
	Lounge          int // Credits
	Robotics        int // Rover speed
	BackupGenerator int // Idle time
	Underforge      int // Smelter efficiency
	Dorm            int // Crafter efficiency
	Sales           int // Alloy & item value
	Classroom       int // Managers
	Marketing       int // Market multiplier
}

func createRooms() *Rooms {
	return &Rooms{
		ID:              1,
		Engineering:     0,
		Aeronautical:    0,
		Packaging:       0,
		Forge:           0,
		Workshop:        0,
		Astronomy:       0,
		Laboratory:      0,
		Terrarium:       0,
		Lounge:          0,
		Robotics:        0,
		BackupGenerator: 0,
		Underforge:      0,
		Dorm:            0,
		Sales:           0,
		Classroom:       0,
		Marketing:       0,
	}
}

func (g *Game) saveRoomsToDB(r *Rooms) {
	r.ID = 1
	err := g.SaveRooms(r)
	if err != nil {
		log.Fatal(err)
	}
}

func (g *Game) loadRoomsFromDB() *Rooms {
	var r Rooms
	err := g.db.First(&r, 1).Error
	if err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return createRooms()
		}
	}
	return &r
}

func (g *Game) UpdateRooms(rooms *Rooms) {
	g.Rooms = rooms
	g.saveRoomsToDB(rooms)
}
