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
}

type GameData struct {
	UpgradeHistory  []UpgradeHistory
	TotalMoneySpent float64
	CurrentStep     int
}

type UpgradeHistory struct {
	Stepnum       int
	Planet        string
	Upgradecost   float64
	Roitime       float64
	ValueIncrease float64
	TotalSpend    float64
}

var GlobalCalcer *Calcer
var DB *sql.DB

func NewGame() *Game {
	db, err := sql.Open("sqlite3", "ipm.sql")
	if err != nil {
		panic(err)
	}

	gameData := &GameData{UpgradeHistory: []UpgradeHistory{}, TotalMoneySpent: 0, CurrentStep: 1}

	makeTables(db)

	managers := getManagersFromDB(db)
	projects := loadProjectsFromDB(db)

	return &Game{
		Planets: []*Planet{
			NewPlanet("Balor", []Ore{{Name: "Copper"}}, []float64{1.0}, 100, 15),
			NewPlanet("Drasta", []Ore{{Name: "Copper"}, {Name: "Iron"}}, []float64{0.8, 0.2}, 200, 15.4),
			NewPlanet("Anadius", []Ore{{Name: "Copper"}, {Name: "Iron"}}, []float64{0.5, 0.5}, 500, 15.8),
			NewPlanet("Dholen", []Ore{{Name: "Iron"}, {Name: "Lead"}}, []float64{0.8, 0.2}, 1250, 15.8),
			NewPlanet("Verr", []Ore{{Name: "Lead"}, {Name: "Iron"}, {Name: "Copper"}}, []float64{0.5, 0.3, 0.2}, 5000, 16.4),
			NewPlanet("Newton", []Ore{{Name: "Lead"}}, []float64{1.0}, 9000, 17.2),
			NewPlanet("Widow", []Ore{{Name: "Iron"}, {Name: "Copper"}, {Name: "Silica"}}, []float64{0.4, 0.4, 0.2}, 15000, 19),
			NewPlanet("Acheron", []Ore{{Name: "Silica"}, {Name: "Copper"}}, []float64{0.6, 0.4}, 25000, 19.5),
			NewPlanet("Yangtze", []Ore{{Name: "Silica"}, {Name: "Aluminium"}}, []float64{0.8, 0.2}, 40000, 19.5),
			NewPlanet("Solveig", []Ore{{Name: "Aluminium"}, {Name: "Silica"}, {Name: "Lead"}}, []float64{0.5, 0.3, 0.2}, 75000, 21),
			NewPlanet("Imir", []Ore{{Name: "Aluminium"}}, []float64{1.0}, 150000, 22.5),
			NewPlanet("Relic", []Ore{{Name: "Lead"}, {Name: "Silica"}, {Name: "Silver"}}, []float64{0.45, 0.35, 0.2}, 250000, 24.5),
			NewPlanet("Nith", []Ore{{Name: "Silver"}, {Name: "Aluminium"}}, []float64{0.8, 0.2}, 400000, 26),
			NewPlanet("Batalla", []Ore{{Name: "Copper"}, {Name: "Iron"}, {Name: "Gold"}}, []float64{0.4, 0.4, 0.2}, 800000, 29),
			NewPlanet("Micah", []Ore{{Name: "Gold"}, {Name: "Silver"}}, []float64{0.5, 0.5}, 1500000, 30.5),
			NewPlanet("Pranas", []Ore{{Name: "Gold"}}, []float64{1.0}, 3000000, 32.5),
			NewPlanet("Castellus", []Ore{{Name: "Aluminium"}, {Name: "Silica"}, {Name: "Diamond"}}, []float64{0.4, 0.35, 0.25}, 6000000, 33.5),
			NewPlanet("Gorgon", []Ore{{Name: "Diamond"}, {Name: "Lead"}}, []float64{0.8, 0.2}, 12000000, 35),
			NewPlanet("Parnitha", []Ore{{Name: "Gold"}, {Name: "Platinum"}}, []float64{0.7, 0.3}, 25000000, 38),
			NewPlanet("Orisoni", []Ore{{Name: "Platinum"}, {Name: "Diamond"}}, []float64{0.7, 0.3}, 50000000, 40),
			NewPlanet("Theseus", []Ore{{Name: "Platinum"}}, []float64{1.0}, 100000000, 44),
			NewPlanet("Zelene", []Ore{{Name: "Silver"}, {Name: "Titanium"}}, []float64{0.7, 0.3}, 200000000, 47.5),
			NewPlanet("Han", []Ore{{Name: "Titanium"}, {Name: "Diamond"}, {Name: "Gold"}}, []float64{0.7, 0.2, 0.1}, 400000000, 50),
			NewPlanet("Strennus", []Ore{{Name: "Titanium"}, {Name: "Platinum"}}, []float64{0.7, 0.3}, 800000000, 55),
			NewPlanet("Osun", []Ore{{Name: "Aluminium"}, {Name: "Iridium"}}, []float64{0.6, 0.4}, 1600000000, 58),
			NewPlanet("Ploitari", []Ore{{Name: "Iridium"}, {Name: "Diamond"}}, []float64{0.5, 0.5}, 3200000000, 60),
			NewPlanet("Elysta", []Ore{{Name: "Iridium"}}, []float64{1.0}, 6400000000, 63),
			NewPlanet("Tikkuun", []Ore{{Name: "Iridium"}, {Name: "Titanium"}, {Name: "Palladium"}}, []float64{0.4, 0.35, 0.25}, 12500000000, 67),
			NewPlanet("Satent", []Ore{{Name: "Palladium"}, {Name: "Titanium"}}, []float64{0.6, 0.4}, 25000000000, 70),
			NewPlanet("Urla Rast", []Ore{{Name: "Palladium"}, {Name: "Diamond"}}, []float64{0.9, 0.1}, 50000000000, 73),
			NewPlanet("Vular", []Ore{{Name: "Palladium"}, {Name: "Osmium"}}, []float64{0.7, 0.3}, 100000000000, 75),
			NewPlanet("Nibiru", []Ore{{Name: "Osmium"}, {Name: "Iridium"}}, []float64{0.6, 0.4}, 250000000000, 76),
			NewPlanet("Xena", []Ore{{Name: "Osmium"}}, []float64{1.0}, 600000000000, 78),
			NewPlanet("Rupert", []Ore{{Name: "Palladium"}, {Name: "Osmium"}, {Name: "Rhodium"}}, []float64{0.55, 0.3, 0.15}, 1500000000000, 78),
			NewPlanet("Pax", []Ore{{Name: "Rhodium"}, {Name: "Platinum"}}, []float64{0.5, 0.5}, 4000000000000, 80),
			NewPlanet("Ivyra", []Ore{{Name: "Rhodium"}}, []float64{1.0}, 10000000000000, 81),
			NewPlanet("Utrits", []Ore{{Name: "Rhodium"}, {Name: "Inerton"}}, []float64{0.8, 0.2}, 25000000000000, 82),
			NewPlanet("Doosie", []Ore{{Name: "Inerton"}, {Name: "Osmium"}}, []float64{0.5, 0.5}, 62000000000000, 84),
			NewPlanet("Zulu", []Ore{{Name: "Inerton"}}, []float64{1.0}, 160000000000000, 84),
			NewPlanet("Unicae", []Ore{{Name: "Inerton"}, {Name: "Quadium"}}, []float64{0.8, 0.2}, 400000000000000, 85),
			NewPlanet("Dune", []Ore{{Name: "Osmium"}}, []float64{0.6}, 1000000000000000, 87),
		},
		LastSteps: 0,
		db:        db,
		Managers:  managers,
		Projects:  projects,
		GamdeData: gameData,
	}
}

func (g *Game) InitData() {
	GlobalCalcer = NewCalcer(g)
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
	g.GamdeData.CurrentStep = 1
	g.GamdeData.UpgradeHistory = []UpgradeHistory{}
	g.Projects = newProjects()
}

func (g *Game) ResetPlanets() {
	for _, planet := range g.Planets {
		planet.resetPlanet()
	}
	g.GamdeData.TotalMoneySpent = 0
}

func (g *Game) ResetManagers() {
	for _, manager := range g.Managers {
		manager.unassignManager()
	}
}
