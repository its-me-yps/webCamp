package data

import (
	"fmt"
	model "wingiesOrNot/models"
)

func Group(s model.Students) map[string]model.Hall {
	groupedData := make(map[string]model.Hall)

	for _, student := range s {
		hall, hallExists := groupedData[student.Hall]
		if !hallExists {
			hall = make(model.Hall)
		}
		if student.Room == "NA" {
			continue
		}
		wingName := student.Room[:3]
		roomNumber := student.Room[3:]

		wing, wingExists := hall[wingName]
		if !wingExists {
			wing = make(model.Wing)
		}

		room, roomExists := wing[roomNumber]
		if !roomExists {
			room = make(model.Room, 0)
		}
		room = append(room, student)
		wing[roomNumber] = room
		hall[wingName] = wing
		groupedData[student.Hall] = hall
	}

	return groupedData
}

func PrintGroupedData(data map[string]model.Hall) {
	for hallName, hall := range data {
		fmt.Printf("%s\n", hallName)

		for wingName, wing := range hall {
			fmt.Printf("  Wing: %s\n", wingName)

			for roomNumber, room := range wing {
				fmt.Printf("    Room: %s\n", roomNumber)

				for i, student := range room {
					fmt.Printf("      %d. %s\n", i+1, student.Name)
				}
			}
		}
	}
}
