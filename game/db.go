package game

import (
	"database/sql"
	"log"
)

func resetPlanetDB(db *sql.DB, planet *Planet) {
	_, err := db.Exec("UPDATE planets SET mining_level = 1, ship_speed_level = 1, ship_cargo_level = 1, colony_level = 0, locked = 1, alchemy_level = 1 WHERE name = ?", planet.Name)
	if err != nil {
		log.Fatal(err)
	}
}

func updatePlanetDB(db *sql.DB, planet *Planet) {
	var locked int
	if planet.Locked {
		locked = 1
	} else {
		locked = 0
	}

	query := `
        INSERT OR REPLACE INTO planets (
            id, name, mining_level, ship_speed_level, ship_cargo_level, colony_level, locked, alchemy_level
        ) VALUES (
            (SELECT id FROM planets WHERE name = ?), ?, ?, ?, ?, ?, ?, ?
        )
    `
	_, err := db.Exec(query, planet.Name, planet.Name, planet.MiningLevel, planet.ShipSpeedLeve1, planet.ShipCargoLevel, planet.ColonyLevel, locked, planet.AlchemyLevel)
	if err != nil {
		log.Fatal(err)
	}
}

func getManagersFromDB(db *sql.DB) []*Manager {
	rows, err := db.Query("SELECT id, stars, primary_role, secondary_role FROM managers")
	if err != nil {
		log.Fatal(err)
	}
	defer rows.Close()

	var managers []*Manager
	for rows.Next() {
		var id int
		var stars int
		var primaryRole string
		var secondaryRole string

		err = rows.Scan(&id, &stars, &primaryRole, &secondaryRole)
		if err != nil {
			log.Fatal(err)
		}

		manager := &Manager{
			ID:        id,
			Stars:     stars,
			Primary:   Role(primaryRole),
			Secondary: SecondaryRole(secondaryRole),
		}
		managers = append(managers, manager)
	}

	return managers
}

func getPlanetsFromDB(db *sql.DB) ([]Planet, error) {
	rows, err := db.Query("SELECT id, name, mining_level, ship_speed_level, ship_cargo_level, colony_level, locked, alchemy_level FROM planets")
	if err != nil {
		log.Fatal(err)
	}
	defer rows.Close()
	var dbPlanets []Planet

	for rows.Next() {
		var id int
		var name string
		var mining_level int
		var ship_speed_level int
		var ship_cargo_level int
		var colony_level int
		var locked int
		var alchemy_level int

		err = rows.Scan(&id, &name, &mining_level, &ship_speed_level, &ship_cargo_level, &colony_level, &locked, &alchemy_level)
		if err != nil {
			log.Fatal(err)
		}
		var lockedBool bool
		if locked == 1 {
			lockedBool = true
		} else {
			lockedBool = false
		}

		planet := Planet{
			Name:           name,
			MiningLevel:    mining_level,
			ShipSpeedLeve1: ship_speed_level,
			ShipCargoLevel: ship_cargo_level,
			ColonyLevel:    colony_level,
			Locked:         lockedBool,
			AlchemyLevel:   alchemy_level,
		}
		dbPlanets = append(dbPlanets, planet)
	}
	return dbPlanets, nil
}

func makeTables(db *sql.DB) error {
	managerSQL := `CREATE TABLE IF NOT EXISTS managers (
        id INTEGER PRIMARY KEY AUTOINCREMENT,
        stars INTEGER,
        primary_role TEXT,
        secondary_role TEXT
    );`

	_, err := db.Exec(managerSQL)
	if err != nil {
		return err
	}
	planetSQL := `CREATE TABLE IF NOT EXISTS planets (
				id INTEGER PRIMARY KEY AUTOINCREMENT,
				name TEXT,
				mining_level INTEGER,
				ship_speed_level INTEGER,
				ship_cargo_level INTEGER,
				locked INTEGER,
				colony_level INTEGER,
				alchemy_level INTEGER
			);`
	_, err = db.Exec(planetSQL)
	if err != nil {
		return err
	}
	return nil
}
