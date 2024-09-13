package routes

import (
	"fmt"
	"strconv"

	"github.com/gin-gonic/gin"
	"sirlon.org/iph-sim/game"
)

func RegisterMarketRoutes(r *gin.Engine, gameInstance *game.Game) {

	r.GET("/market", func(c *gin.Context) {
		data := gameInstance.GenerateMarketHTML()
		c.HTML(200, "market.html", data)
	})

	r.POST("/updateOreStars", func(c *gin.Context) {
		c.PostForm("oreName")
		stars, err := strconv.Atoi(c.PostForm("stars"))
		if err != nil {
			fmt.Println("Error converting stars to int")
		}
		gameInstance.SetStars(c.PostForm("oreName"), stars)
		data := gameInstance.GenerateMarketHTML()
		c.HTML(200, "market.html", data)
	})
	r.POST("/updateOreMarket", func(c *gin.Context) {
		c.PostForm("oreName")
		trend, err := strconv.ParseFloat(c.PostForm("marketTrend"), 64)
		if err != nil {
			fmt.Println("Error converting stars to int")
		}
		fmt.Println(c.PostForm("oreName"))
		fmt.Println(trend)
		data := gameInstance.GenerateMarketHTML()
		c.HTML(200, "market.html", data)
	})
	r.POST("/updateAlloyStars", func(c *gin.Context) {
		data := gameInstance.GenerateMarketHTML()
		c.HTML(200, "market.html", data)
	})
	r.POST("/updateAlloyMarket", func(c *gin.Context) {
		data := gameInstance.GenerateMarketHTML()
		c.HTML(200, "market.html", data)
	})
	r.POST("/updateItemMarket", func(c *gin.Context) {
		data := gameInstance.GenerateMarketHTML()
		c.HTML(200, "market.html", data)
	})
}
