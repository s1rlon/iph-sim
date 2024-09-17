package routes

import (
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
	"sirlon.org/iph-sim/game"
)

func RegisterProjectRoutes(r *gin.Engine, gameInstance *game.Game) {
	r.GET("/projects", func(c *gin.Context) {
		c.HTML(http.StatusOK, "projects.html", gameInstance.Projects)
	})

	r.POST("/updateProjects", func(c *gin.Context) {
		projects := &game.Projects{
			TelescopeLevel: parseFormValue(c, "TelescopeLevel"),
			MiningLevel:    parseFormValue(c, "MiningLevel"),
			ShipSpeedLevel: parseFormValue(c, "ShipSpeedLevel"),
			ShipCargoLevel: parseFormValue(c, "ShipCargoLevel"),
			Beacon:         parseFormValue(c, "Beacon"),
			TaxLevel:       parseFormValue(c, "TaxLevel"),
			SmeltSpeed:     parseFormValue(c, "SmeltSpeed"),
			SmeltEff:       parseFormValue(c, "SmeltEff"),
			AlloyValue:     parseFormValue(c, "AlloyValue"),
			SmeltSpec:      parseFormValue(c, "SmeltSpec"),
			CraftSpeed:     parseFormValue(c, "CraftSpeed"),
			CraftEff:       parseFormValue(c, "CraftEff"),
			ItemValue:      parseFormValue(c, "ItemValue"),
			CraftSpec:      parseFormValue(c, "CraftSpec"),
			PrefVendor:     parseFormValue(c, "PrefVendor"),
			OreTargeting:   parseFormValue(c, "OreTargeting"),
			ManTraining:    parseFormValue(c, "ManTraining"),
			ManSTraing:     parseFormValue(c, "ManSTraing"),
			LeaderTraining: parseFormValue(c, "LeaderTraining"),
		}

		// Assuming you have a function to update the projects
		gameInstance.UpdateProjects(projects)

		c.Redirect(http.StatusSeeOther, "/projects")
	})
}

func parseFormValue(c *gin.Context, key string) int {
	valueStr := c.PostForm(key)
	value, err := strconv.Atoi(valueStr)
	if err != nil {
		return 0
	}
	return value
}

func parseFormFloat(c *gin.Context, field string) float64 {
	valueStr := c.PostForm(field)
	value, err := strconv.ParseFloat(valueStr, 64)
	if err != nil {
		return 0.0
	}
	return value
}
