package routes

import (
	"net/http"
	"sort"
	"strconv"

	"github.com/gin-gonic/gin"
	"sirlon.org/iph-sim/game"
)

func RegisterManagerRoutes(r *gin.Engine, gameInstance *game.Game) {
	r.GET("/managers", func(c *gin.Context) {
		managers := gameInstance.GetManagers()

		sort.Slice(managers, func(i, j int) bool {
			return managers[i].Stars > managers[j].Stars
		})
		planets := gameInstance.Planets
		c.HTML(http.StatusOK, "managers.html", gin.H{
			"Managers": managers,
			"Planets":  planets,
			"GameData": gameInstance.GameData,
		})
	})

	r.POST("/add-manager", func(c *gin.Context) {
		stars, err := strconv.Atoi(c.PostForm("stars"))
		if err != nil {
			c.String(http.StatusBadRequest, "Invalid stars value")
			return
		}
		primary := c.PostForm("primary")
		secondary := c.PostForm("secondary")

		manager := &game.Manager{
			Stars:     stars,
			Primary:   game.Role(primary),
			Secondary: game.SecondaryRole(secondary),
		}
		gameInstance.AddManager(manager)
		c.Redirect(http.StatusFound, "/managers")
	})

	r.POST("/delete-manager", func(c *gin.Context) {
		managerID, err := strconv.Atoi(c.PostForm("manager_id"))
		if err != nil {
			c.String(http.StatusBadRequest, "Invalid manager ID")
			return
		}
		gameInstance.DeleteManager(managerID)
		c.Redirect(http.StatusFound, "/managers")
	})

	r.POST("/update-manager-planet", func(c *gin.Context) {
		managerID, err := strconv.Atoi(c.PostForm("manager_id"))
		if err != nil {
			c.String(http.StatusBadRequest, "Invalid manager ID")
			return
		}
		planetName := c.PostForm("planet")

		err = gameInstance.UpdateManagerPlanet(managerID, planetName)
		if err != nil {
			c.String(http.StatusInternalServerError, "Failed to update manager planet")
			return
		}
		c.Redirect(http.StatusFound, "/managers")
	})

	r.POST("/update-manager-slots", func(c *gin.Context) {
		managerSlots, err := strconv.Atoi(c.PostForm("managerSlots"))
		if err != nil {
			c.String(http.StatusBadRequest, "Invalid manager slots value")
			return
		}
		gameInstance.GameData.ManagerSlots = managerSlots
		gameInstance.GameData.SyncDB(gameInstance.GetDB())
		c.Redirect(http.StatusFound, "/managers")
	})

	r.GET("/assign-managers", func(c *gin.Context) {
		gameInstance.AssignManagers()
		c.Redirect(http.StatusFound, "/")
	})

}
