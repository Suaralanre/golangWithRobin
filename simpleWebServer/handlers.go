package main

import (
	"fmt"
	"log"
	"net/http"
	"time"
)

func GetTime(w http.ResponseWriter, r *http.Request) {

	if r.Method != http.MethodGet {
		w.Header().Set("Allow", http.MethodGet)
		http.Error(w, "Method Not Allowed", http.StatusMethodNotAllowed)
		return
	}

	_, err := fmt.Fprintf(w, "The current date and time is %s", currentTime.Format(time.RFC850))

	if err != nil {
		WriteError(w, "Internal Server Error", http.StatusInternalServerError)
		log.Print("Error with printing out time of the day")
		return
	}
}

func SetTime(w http.ResponseWriter, r *http.Request) {

	if r.Method != http.MethodPost {
		w.Header().Set("Allow", http.MethodPost)
		WriteError(w, "Method Not Allowed", http.StatusMethodNotAllowed)
		return
	}

	err := r.ParseForm()

	if err != nil {
		WriteError(w, "Invalid Request", http.StatusBadRequest)
		log.Print("Error Parsing request")
		return
	}

	newTime := r.FormValue("time")
	if newTime == "" {
		WriteError(w, "A valid time value is required", http.StatusBadRequest)
		log.Print("Time is required")
		return
	}

	parsedTime, err := time.Parse(time.RFC3339, newTime)
	if err != nil {
		WriteError(w, "Invalid time format: return time in RFC3339", http.StatusBadRequest)
		log.Print("Error parsing time")
		return
	}
	currentTime = parsedTime

	w.WriteHeader(http.StatusOK)
	fmt.Fprintf(w, "Time updated to: %s", currentTime.Format(time.RFC850))
}
