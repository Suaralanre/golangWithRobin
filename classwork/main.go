package main

import (
	"encoding/base64"
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
	Files       []File      `json:"files"`
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

	for i, file := range report.Files {
	decodedData, err := base64.StdEncoding.DecodeString(file.Data)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	fileName := fmt.Sprintf("file_%d%s", i, getFileExtension(file.Mime))
	err = os.WriteFile(fileName, decodedData, 0700)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	report.Files[i].Data = fileName
}
	

	jsonReport, err := json.MarshalIndent(&report, "", "    ")
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	reportName := fmt.Sprintf("damage_report_%s.json", report.CreatedAt.Format("2006-01-02"))

	// write json report to a file
	err = os.WriteFile(reportName, jsonReport, 0700)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	w.WriteHeader(http.StatusCreated)
	fmt.Fprint(w, "Damage report created successfully!")

}

func getFileExtension(mime string) string {
	switch mime {
	case "image/jpeg":
		return ".jpg"
	case "image/png":
		return ".png"
	// just fooling around here
	case "image/gif":
		return ".gif"
	case "application/pdf":
		return ".pdf"
	default:
		return ".bin"
	}
}
