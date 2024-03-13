package server

import (
	"net/http"
	model "wingiesOrNot/models"
	"wingiesOrNot/utils"

	"github.com/gin-gonic/gin"
)

func getReq2(c *gin.Context, groupedData map[string]model.Hall) {
	h := c.Param("hall")
	w := c.Param("wing")
	r := c.Param("room")

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
}

func postReq2(c *gin.Context, raw model.Students) {
	// Expected body struct of post req
	var reqBody model.WingiesOrNot
	if err := c.ShouldBindJSON(&reqBody); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid JSON format"})
		return
	}

	result, err := utils.WingiesOrNot(reqBody.Id1, reqBody.Id2, raw)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	if result {
		c.String(http.StatusOK, "YES")
	} else {
		c.String(http.StatusOK, "NO")
	}
}
func findRoommate(c *gin.Context, groupedData map[string]model.Hall) {
	var reqBody struct {
		ID string `json:"id"`
	}
	if err := c.ShouldBindJSON(&reqBody); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid JSON format"})
		return
	}
	studentID := reqBody.ID

	var studentDetails model.Student

	for _, hall := range groupedData {

		for _, wing := range hall {
			for _, room := range wing {

				for _, student := range room {
					if student.Id == studentID {
						studentDetails = student
						break
					}
				}
			}
		}
	}

	if studentDetails == (model.Student{}) {
		c.JSON(http.StatusNotFound, gin.H{"error": "Student not found"})
		return
	}
	var roommates []model.Student

	hall := studentDetails.Hall
	room := studentDetails.Room[3:]
	wing := studentDetails.Room[:3]
	fmt.Printf("h-%v , r-%v,w-%v\n", hall, room, wing)

	for _, student := range groupedData[hall][wing][room] {
		if student.Id != studentID {
			roommates = append(roommates, student)
		}
	}
	c.JSON(http.StatusOK, roommates)
}
