package routes

import (
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
	"sirlon.org/iph-sim/game"
)

func RegisterManagerRoutes(r *gin.Engine, gameInstance *game.Game) {
	r.GET("/managers", func(c *gin.Context) {
		managers := game.GetManagers(gameInstance)
		planets := gameInstance.Planets
		c.HTML(http.StatusOK, "managers.html", gin.H{
			"Managers": managers,
			"Planets":  planets,
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
		game.AddManager(gameInstance, manager)
		c.Redirect(http.StatusFound, "/managers")
	})

	r.POST("/delete-manager", func(c *gin.Context) {
		managerID, err := strconv.Atoi(c.PostForm("manager_id"))
		if err != nil {
			c.String(http.StatusBadRequest, "Invalid manager ID")
			return
		}
		game.DeleteManager(gameInstance, managerID)
		c.Redirect(http.StatusFound, "/managers")
	})

	r.POST("/update-manager-planet", func(c *gin.Context) {
		managerID, err := strconv.Atoi(c.PostForm("manager_id"))
		if err != nil {
			c.String(http.StatusBadRequest, "Invalid manager ID")
			return
		}
		planetName := c.PostForm("planet")

		err = game.UpdateManagerPlanet(gameInstance, managerID, planetName)
		if err != nil {
			c.String(http.StatusInternalServerError, "Failed to update manager planet")
			return
		}
		c.Redirect(http.StatusFound, "/managers")
	})

	r.GET("/assign-managers", func(c *gin.Context) {
		gameInstance.AssignManagers()
		c.Redirect(http.StatusFound, "/")
	})

}
