package game

import (
	"database/sql"
	"log"
)

type Rooms struct {
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
	_, err := g.db.Exec("DELETE FROM rooms")
	if err != nil {
		log.Fatal(err)
	}

	query := `
			INSERT INTO rooms (
					engineering, aeronautical, packaging, forge, workshop, astronomy, laboratory, terrarium, lounge, robotics, backup_generator, underforge, dorm, sales, classroom, marketing
			) VALUES (?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?)
	`
	_, err = g.db.Exec(query, r.Engineering, r.Aeronautical, r.Packaging, r.Forge, r.Workshop, r.Astronomy, r.Laboratory, r.Terrarium, r.Lounge, r.Robotics, r.BackupGenerator, r.Underforge, r.Dorm, r.Sales, r.Classroom, r.Marketing)
	if err != nil {
		log.Fatal(err)
	}
}

func loadRoomsFromDB(db *sql.DB) *Rooms {
	query := `
			SELECT engineering, aeronautical, packaging, forge, workshop, astronomy, laboratory, terrarium, lounge, robotics, backup_generator, underforge, dorm, sales, classroom, marketing
			FROM rooms
			ORDER BY id DESC LIMIT 1
	`
	row := db.QueryRow(query)

	var r Rooms
	err := row.Scan(&r.Engineering, &r.Aeronautical, &r.Packaging, &r.Forge, &r.Workshop, &r.Astronomy, &r.Laboratory, &r.Terrarium, &r.Lounge, &r.Robotics, &r.BackupGenerator, &r.Underforge, &r.Dorm, &r.Sales, &r.Classroom, &r.Marketing)
	if err != nil {
		if err == sql.ErrNoRows {
			return createRooms()
		}
		log.Fatal(err)
	}

	return &r
}

func (g *Game) UpdateRooms(rooms *Rooms) {
	g.Rooms = rooms
	g.saveRoomsToDB(rooms)
}
