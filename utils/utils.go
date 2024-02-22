package utils

import (
	"encoding/json"
	"log"
	"net/http"
	"strconv"
	"strings"
)

// respondWithJSON is a utility function that sends a JSON response with the specified payload.
func RespondWithJSON(w http.ResponseWriter, payload interface{}) {
	// SETTING

	w.Header().Set("Content-Type", "application/json") // setting header settings
	response, err := json.Marshal(payload)             // marshalling payload to json response
	// ERROR CHECKING
	if err != nil { // if err return err
		log.Fatalf("Error marshaling JSON: %v", err)
		RespondWithError(w, "Internal Server Error")
		return
	}

	// SENDING
	w.Write(response) // write json response to writer
}

// respondWithError is a utility function that sends a JSON response with an error message.
func RespondWithError(w http.ResponseWriter, message string) {
	// SETTING
	w.Header().Set("Content-Type", "application/json") // setting header settings

	response, err := json.Marshal(map[string]string{"error": message}) // marshalling error message to json response
	// ERROR CHECKING
	if err != nil {
		log.Fatalf("Error marshaling JSON: %v", err)
		http.Error(w, "Internal Server Error", http.StatusInternalServerError)
		return
	}

	// writing error
	http.Error(w, string(response), http.StatusMethodNotAllowed)
}

// getTaskNumFromRequest is a utility function to extract the task number from the URL.
func GetTaskNumFromRequest(r *http.Request) int {
	// Extract the task number from the URL
	parts := strings.Split(r.URL.Path, "/") // splitting to parts
	if len(parts) < 4 {                     // checking url for validness
		log.Printf("Invalid URL format: %s", r.URL.Path)
		return 0
	}

	id := parts[len(parts)-1]      // getting last part
	idInt, err := strconv.Atoi(id) // converting to int
	if err != nil {                // error checking
		log.Printf("Error converting task number to integer: %v", err)
		return 0
	}
	return idInt // returning id
}
