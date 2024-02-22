// main.go
package main

import (
	"log"
	"net/http"

	"github.com/GitEagleY/RESTful-Go.git/handlers"
)

func main() {
	// Set up the server routes
	http.HandleFunc("/tasks/display", handlers.DisplayAllTasks)
	http.HandleFunc("/tasks/addById/", handlers.AddByID)
	http.HandleFunc("/tasks/update/", handlers.UpdateTask)
	http.HandleFunc("/tasks/delete/", handlers.DeleteTask)
	http.HandleFunc("/tasks/getTaskById/", handlers.GetByID)

	// Start the server
	log.Fatal(http.ListenAndServe("localhost:8080", nil))
}
