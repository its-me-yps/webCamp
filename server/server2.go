package server

import (
	model "wingiesOrNot/models"

	"github.com/gin-gonic/gin"
)

// Server2( using gin framework )
// utility of framework
func Server2(groupedData map[string]model.Hall, rawData model.Students, port string) {
	r := gin.Default()

	r.GET("/:hall/:wing/:room", func(c *gin.Context) {
		h := c.Param("hall")
		w := c.Param("wing")
		r := c.Param("room")

		if h == "" {
			c.JSON(200, groupedData)
		}
		if hall, ok := groupedData[h]; ok {
			if w == "" {
				c.JSON(200, hall)
			}
			if wing, ok := hall[w]; ok {
				if r == "" {
					c.JSON(200, wing)
				}
				if room, ok := wing[r]; ok {
					c.JSON(200, room)
					return
				}
			}
		}

		c.JSON(404, gin.H{"error": "Not Found"})
	})
	
	r.Run(":8080")
}
