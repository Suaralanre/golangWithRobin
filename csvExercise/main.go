package main

import (
	"fmt"
	"log"
	"net/http"
	"encoding/csv"
)

var headers = []string{"FirstName", "LastName", "Age", "Number of Children"}

var data = [][]string{
	{"John", "Doe", "30", "2"},
	{"Jane", "Smith", "25", "1"},
	{"Bob", "Johnson", "40", "3"},
	{"Alice", "Williams", "35", "2"},
	{"Mike", "Brown", "28", "0"},
	{"Emma", "Davis", "32", "1"},
	{"Oliver", "Miller", "38", "2"},
	{"Sophia", "Wilson", "29", "1"},
	{"William", "Moore", "42", "3"},
	{"Olivia", "Taylor", "31", "2"},
}

func main() {
	http.HandleFunc("/", handleCSV)
	fmt.Print("Starting Server...")
	log.Fatal(http.ListenAndServe(":8080", nil))
}

func handleCSV(w http.ResponseWriter, r *http.Request) {

	w.Header().Set("Content-Type", "text/csv")
	w.Header().Set("Content-Disposition", "attachment; filename=data.csv")

	query := r.URL.Query().Get("delim")

	var delim rune;

	if query == "" {
		delim = ','
	} else if len(query) > 0 {
		delim = rune(query[0])
	} else {
		http.Error(w, "Invalid delimiter", http.StatusBadRequest)
		return
	}

	writer := csv.NewWriter(w)
	writer.Comma = delim
	fmt.Println(writer.Comma)

	headerErr := writer.Write(headers)
	if headerErr != nil {
		log.Printf("CSV data not written: %v", headerErr)
	}
	defer writer.Flush()

	writerErr := writer.WriteAll(data)
	if writerErr != nil {
		log.Printf("CSV data not written: %v", writerErr)
	}
	}
