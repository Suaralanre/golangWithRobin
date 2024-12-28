package main

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"os"
	"time"
)

type DamageReport struct {
	Description string    `json:"description"`
	CreatedAt   time.Time `json:"created_at"`
	Files       []File  `json:"files"`
}

type File struct {
	Mime string `json:"mime"`
	Data string `json:"data"`
}

func main() {

	route := http.NewServeMux()
	route.HandleFunc("POST /report", reportHandler)

	log.Println("starting server on :4000")

	err := http.ListenAndServe(":4000", route)
	log.Fatal(err)

}

func reportHandler(w http.ResponseWriter, r *http.Request) {

	var report DamageReport
	err := json.NewDecoder(r.Body).Decode(&report)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	jsonReport, err := json.MarshalIndent(report, "", "")
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	// write json report to a file

	err = os.WriteFile("damage_report.json", jsonReport, 0700)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	w.WriteHeader(http.StatusCreated)
	fmt.Fprint(w, "Damage report created successfully!")

}
