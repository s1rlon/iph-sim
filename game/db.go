package game

import (
	"database/sql"
	"log"
)

func resetPlanetDB(db *sql.DB, planet *Planet) {
	_, err := db.Exec("UPDATE planets SET mining_level = 1, ship_speed_level = 1, ship_cargo_level = 1, colony_level = 0, locked = 1, alchemy_level = 1, rover = 0 WHERE name = ?", planet.Name)
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

	var rover int
	if planet.Rover {
		rover = 1
	} else {
		rover = 0
	}

	query := `
			INSERT OR REPLACE INTO planets (
					id, name, mining_level, ship_speed_level, ship_cargo_level, colony_level, locked, alchemy_level, rover
			) VALUES (
					(SELECT id FROM planets WHERE name = ?), ?, ?, ?, ?, ?, ?, ?, ?
			)
	`
	_, err := db.Exec(query, planet.Name, planet.Name, planet.MiningLevel, planet.ShipSpeedLeve1, planet.ShipCargoLevel, planet.ColonyLevel, locked, planet.AlchemyLevel, rover)
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
	rows, err := db.Query("SELECT id, name, mining_level, ship_speed_level, ship_cargo_level, colony_level, locked, alchemy_level, rover FROM planets")
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
		var rover int

		err = rows.Scan(&id, &name, &mining_level, &ship_speed_level, &ship_cargo_level, &colony_level, &locked, &alchemy_level, &rover)
		if err != nil {
			log.Fatal(err)
		}
		var lockedBool bool
		if locked == 1 {
			lockedBool = true
		} else {
			lockedBool = false
		}

		var roverBool bool
		if rover == 1 {
			roverBool = true
		} else {
			roverBool = false
		}

		planet := Planet{
			Name:           name,
			MiningLevel:    mining_level,
			ShipSpeedLeve1: ship_speed_level,
			ShipCargoLevel: ship_cargo_level,
			ColonyLevel:    colony_level,
			Locked:         lockedBool,
			AlchemyLevel:   alchemy_level,
			Rover:          roverBool,
		}
		dbPlanets = append(dbPlanets, planet)
	}
	return dbPlanets, nil
}

func loadProjectsFromDB(db *sql.DB) *Projects {
	query := `SELECT telescope_level, mining_level, ship_speed_level, ship_cargo_level, beacon, tax_level, smelt_speed, smelt_eff, alloy_value, smelt_spec ,craft_speed, craft_eff, item_value, craft_spec ,pref_vendor, ore_targeting, man_training, man_straing, leader_training FROM projects ORDER BY id DESC LIMIT 1`
	row := db.QueryRow(query)

	p := newProjects()
	err := row.Scan(&p.TelescopeLevel, &p.MiningLevel, &p.ShipSpeedLevel, &p.ShipCargoLevel, &p.Beacon, &p.TaxLevel, &p.SmeltSpeed, &p.SmeltEff, &p.AlloyValue, &p.SmeltSpec, &p.CraftSpeed, &p.CraftEff, &p.ItemValue, &p.CraftSpec, &p.PrefVendor, &p.OreTargeting, &p.ManTraining, &p.ManSTraing, &p.LeaderTraining)
	if err != nil {
		if err == sql.ErrNoRows {
			return newProjects()
		}
		log.Fatal(err)
	}
	return p
}

func makeTables(db *sql.DB) error {
	queries := []string{
		`CREATE TABLE IF NOT EXISTS planets (id INTEGER PRIMARY KEY AUTOINCREMENT, name TEXT, mining_level INTEGER, ship_speed_level INTEGER, ship_cargo_level INTEGER, colony_level INTEGER, locked INTEGER, alchemy_level INTEGER, rover INTEGER)`,
		`CREATE TABLE IF NOT EXISTS managers ( id INTEGER PRIMARY KEY AUTOINCREMENT, stars INTEGER, primary_role TEXT, secondary_role TEXT)`,
		`CREATE TABLE IF NOT EXISTS projects (id INTEGER PRIMARY KEY AUTOINCREMENT, telescope_level INTEGER, mining_level INTEGER, ship_speed_level INTEGER, ship_cargo_level INTEGER, beacon INTEGER, tax_level INTEGER, smelt_speed INTEGER, smelt_eff INTEGER, alloy_value INTEGER, smelt_spec INTEGER ,craft_speed INTEGER, craft_eff INTEGER, item_value INTEGER, craft_spec INTEGER, pref_vendor INTEGER, ore_targeting INTEGER, man_training INTEGER, man_straing INTEGER, leader_training INTEGER)`,
		`CREATE TABLE IF NOT EXISTS upgrade_history (id INTEGER PRIMARY KEY AUTOINCREMENT, stepnum INTEGER, planet TEXT, upgradecost REAL, roitime REAL, valueincrease REAL, totalspend REAL)`,
		`CREATE TABLE IF NOT EXISTS rooms (id INTEGER PRIMARY KEY AUTOINCREMENT,engineering INTEGER,aeronautical INTEGER,packaging INTEGER,forge INTEGER,workshop INTEGER,astronomy INTEGER,laboratory INTEGER,terrarium INTEGER,lounge INTEGER,robotics INTEGER,backup_generator INTEGER,underforge INTEGER,dorm INTEGER,sales INTEGER,classroom INTEGER,marketing INTEGER)`,
		`CREATE TABLE IF NOT EXISTS stars (id INTEGER PRIMARY KEY AUTOINCREMENT,name TEXT,stars INTEGER)`,
	}

	for _, query := range queries {
		_, err := db.Exec(query)
		if err != nil {
			return err
		}
	}

	return nil
}
