package routes

import (
	"fmt"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
	"sirlon.org/iph-sim/game"
)

func RegisterMarketRoutes(r *gin.Engine, gameInstance *game.Game) {

	r.GET("/market", func(c *gin.Context) {
		data := gameInstance.GenerateMarketHTML()
		c.HTML(200, "market.html", data)
	})

	r.POST("/updateCraftableMarket", func(c *gin.Context) {
		stars, err := strconv.Atoi(c.PostForm("stars"))
		if err != nil {
			fmt.Println("Error converting stars to int")
		}
		trend, err := strconv.ParseFloat(c.PostForm("marketTrend"), 64)
		if err != nil {
			fmt.Println("Error converting stars to int")
		}
		fmt.Println(trend)
		gameInstance.SetStars(c.PostForm("CraftableName"), stars)
		gameInstance.SetTrend(c.PostForm("CraftableName"), trend)
		c.Redirect(http.StatusFound, "/market")
	})
}
