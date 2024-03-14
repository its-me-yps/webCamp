package server

import (
	model "wingiesOrNot/models"

	"github.com/gin-gonic/gin"
)

// Server2( using gin framework )
// utility of framework
func Server2(groupedData map[string]model.Hall, rawData model.Students, port string) {
	r := gin.Default()

	r.POST("/wingiesOrNot", func(c *gin.Context) {
		postReq2(c, rawData)
	})

	r.GET("/getWingies/:id", func(c *gin.Context) {
		student := getWing(c, rawData)
		if (student != model.Student{}) {
			fetchWingies(c, groupedData, student)
		}
	})
	r.GET("/", func(c *gin.Context) {
		c.JSON(200, groupedData)
	})

	r.GET("/:hall", func(c *gin.Context) {
		c.Set("wing", "")
		c.Set("room", "")
		getReq2(c, groupedData)
	})

	r.GET("/:hall/:wing", func(c *gin.Context) {
		c.Set("room", "")
		getReq2(c, groupedData)
	})

	r.GET("/:hall/:wing/:room", func(c *gin.Context) {
		getReq2(c, groupedData)
	})

	r.Run(":" + port)
}
