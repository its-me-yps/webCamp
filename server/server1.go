package server

import (
	"encoding/json"
	"log"
	"net/http"
	"strings"
	model "wingiesOrNot/models"
)

// Server1( using standard http package )
func Server1(groupedData map[string]model.Hall, raw model.Students, port string) {
	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
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
	})

	log.Printf("Server1 starting on port %s ...", port)
	log.Fatal(http.ListenAndServe(":"+port, nil))
	log.Println("YESS")
}
