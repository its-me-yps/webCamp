package server

import (
	"encoding/json"
	"log"
	"net/http"
	"strings"
	model "wingiesOrNot/models"
	"wingiesOrNot/utils"
)

func getReq1(w http.ResponseWriter, r *http.Request, groupedData map[string]model.Hall) {
	path := r.URL.Path
	var res interface{}
	log.Println(path)
	if path == "/" {
		res = groupedData
	} else {
		path = strings.TrimPrefix(path, "/")
		req := strings.Split(path, "/")

		switch len(req) {
		case 1:
			// Return Hall Data
			if hall, ok := groupedData[req[0]]; ok {
				res = hall
			} else {
				http.NotFound(w, r)
				return
			}
		case 2:
			// Return Wing Data
			if hall, ok := groupedData[req[0]]; ok {
				if wing, ok2 := hall[req[1]]; ok2 {
					res = wing
				} else {
					http.NotFound(w, r)
					return
				}
			} else {
				http.NotFound(w, r)
				return
			}
		case 3:
			// Return Room Data
			if hall, ok := groupedData[req[0]]; ok {
				if wing, ok2 := hall[req[1]]; ok2 {
					if room, ok3 := wing[req[2]]; ok3 {
						res = room
					} else {
						http.NotFound(w, r)
						return
					}
				} else {
					http.NotFound(w, r)
					return
				}
			} else {
				http.NotFound(w, r)
				return
			}
		default:
			http.NotFound(w, r)
			return
		}
	}

	jsonRes, err := json.Marshal(res)
	if err != nil {
		log.Println("Error:", err)
		http.Error(w, "Internal Server Error", http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	w.Write(jsonRes)
}

func postReq1(w http.ResponseWriter, r *http.Request, raw model.Students) {
	// Expected body struct of post req
	var reqBody model.WingiesOrNot
	err := json.NewDecoder(r.Body).Decode(&reqBody)
	if err != nil {
		http.Error(w, "Invalid JSON format", http.StatusBadRequest)
		return
	}

	result, err := utils.WingiesOrNot(reqBody.Id1, reqBody.Id2, raw)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	if result {
		w.Write([]byte("YES"))
	} else {
		w.Write([]byte("NO"))
	}
}
