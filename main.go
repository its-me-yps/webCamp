package main

import (
	"log"
	"os"
	"wingiesOrNot/data"
	model "wingiesOrNot/models"
	"wingiesOrNot/server"
	"wingiesOrNot/utils"
)

var raw model.Students
var grouped map[string]model.Hall

func main() {
	// Checking if data has already been fetched and grouped on basis of wing
	_, err := os.Stat("./data/groupedData.json")
	if err != nil {
		fetchAndGroup()
	} else {
		err := utils.RetrieveFromFile("./data/groupedData.json", &grouped)
		if err != nil {
			log.Fatal("Error retrieving:", err)
		}

		err = utils.RetrieveFromFile("./data/rawData.json", &raw)
		if err != nil {
			log.Fatal("Error retrieving:", err)
		}
	}

	// Command Line arguments
	args := os.Args

	os.Setenv("SECRET", "SPRING_CAMP")

	if args[1] == "1" {
		server.Server1(grouped, raw, "8080")
	} else if args[1] == "2" {
		server.Server2(grouped, raw, "3000")
	} else {
		log.Println("Invalid Argument")
	}
}

func fetchAndGroup() {
	// Fetching Raw Data
	raw = data.Fetch("XvhvZNBWObiDyf651zDE8LsSx59zssBKVMlTHSftn566l7rXoVrbQxnW0L2p6L5A", "23")
	data.PrintData(raw)
	utils.SaveToFile(&raw, "./data/rawData.json")

	// Grouping based on wing
	grouped = data.Group(raw)
	data.PrintGroupedData(grouped)
	utils.SaveToFile(&grouped, "./data/groupedData.json")
}
