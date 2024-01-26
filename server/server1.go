package server

import (
	"log"
	"net/http"
	model "wingiesOrNot/models"
)

// Server1( using standard http package )
func Server1(groupedData map[string]model.Hall, raw model.Students, port string) {
	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		if r.Method == http.MethodGet {
			getReq1(w, r, groupedData)
		} else if r.Method == http.MethodPost {
			if r.URL.Path == "/wingiesOrNot" {
				postReq1(w, r, raw)
			} else {
				http.NotFound(w, r)
			}
		} else {
			http.Error(w, "Method Not Allowed", http.StatusMethodNotAllowed)
		}
	})

	log.Printf("Server1 starting on port %s ...", port)
	log.Fatal(http.ListenAndServe(":"+port, nil))
}
