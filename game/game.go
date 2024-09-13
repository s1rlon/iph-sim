package game

import (
	"database/sql"

	_ "github.com/mattn/go-sqlite3"
)

type Game struct {
	Planets   []*Planet
	LastSteps int
	Managers  []*Manager
	Projects  *Projects
	db        *sql.DB
	GamdeData *GameData
	Ships     *Ships
	Ores      *[]*Ore
}

var GlobalCalcer *Calcer
var DB *sql.DB
var MarketSVC *Market

func NewGame() *Game {
	db, err := sql.Open("sqlite3", "ipm.sql")
	if err != nil {
		panic(err)
	}

	gameData := &GameData{UpgradeHistory: []UpgradeHistory{}}
	gameData.LoadUpgradeHistoryFromDB(db)

	makeTables(db)
	ores := createOres()

	managers := getManagersFromDB(db)
	projects := loadProjectsFromDB(db)
	ships := NewShips()
	planets := makeNewPlanets(ores)

	return &Game{
		Planets:   planets,
		LastSteps: 0,
		db:        db,
		Managers:  managers,
		Projects:  projects,
		GamdeData: gameData,
		Ships:     ships,
		Ores:      ores,
	}
}

func (g *Game) InitData() {
	GlobalCalcer = NewCalcer(g)
	MarketSVC = NewMarket(g)
	DB = g.db
	dbPlanets, _ := getPlanetsFromDB(g.db)
	for _, planet := range g.Planets {
		for _, dbPlanet := range dbPlanets {
			if planet.Name == dbPlanet.Name {
				planet.MiningLevel = dbPlanet.MiningLevel
				planet.ShipSpeedLeve1 = dbPlanet.ShipSpeedLeve1
				planet.ShipCargoLevel = dbPlanet.ShipCargoLevel
				planet.Locked = dbPlanet.Locked
				planet.ColonyLevel = dbPlanet.ColonyLevel
				planet.AlchemyLevel = dbPlanet.AlchemyLevel
			}
		}
	}
}

func (g *Game) ResetGalaxy() {
	g.ResetPlanets()
	g.ResetManagers()
	g.GamdeData.resetGameData()
	g.Projects = newProjects()
}

func (g *Game) ResetPlanets() {
	for _, planet := range g.Planets {
		planet.resetPlanet()
	}
}

func (g *Game) ResetManagers() {
	for _, manager := range g.Managers {
		manager.unassignManager()
	}
}

func (g *Game) moneySpent() float64 {
	total := 0.0
	for _, upgrade := range g.GamdeData.UpgradeHistory {
		total += upgrade.Upgradecost
	}
	return total
}

func (g *Game) currentStep() int {
	return len(g.GamdeData.UpgradeHistory)
}
