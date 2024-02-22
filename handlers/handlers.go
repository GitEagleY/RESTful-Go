package handlers

import (
	"encoding/json"
	"net/http"

	models "github.com/GitEagleY/RESTful-Go.git/models"
	utils "github.com/GitEagleY/RESTful-Go.git/utils"
)

var db = make(map[int]models.Task)

// GET
func DisplayAllTasks(w http.ResponseWriter, r *http.Request) {
	tempTasksDB := make([]models.Task, 0, len(db))

	switch r.Method {
	case http.MethodGet: // if http method that used to access /tasks/display route GET

		for _, task := range db {
			tempTasksDB = append(tempTasksDB, task) // adding to temp task db tasks from actual db
		}
		utils.RespondWithJSON(w, tempTasksDB) // sending json
	default: // if http not GET
		utils.RespondWithError(w, "Method not allowed") // if err sending err json
	}
}

// POST
func AddByID(w http.ResponseWriter, r *http.Request) {
	var taskToAdd models.Task // Temporary variable to hold the task data.

	switch r.Method {
	case http.MethodPost: // Handle only HTTP POST requests.
		id := utils.GetTaskNumFromRequest(r) // Extract the task ID from the request.

		// Decode the JSON payload from the request body.
		err := json.NewDecoder(r.Body).Decode(&taskToAdd)
		if err != nil {
			utils.RespondWithError(w, "Invalid request payload")
			return
		}

		// No error while decoding, so add the task to the database.
		db[id] = taskToAdd

		// responding with success json
		utils.RespondWithJSON(w, taskToAdd)
	default:
		// Respond with an error for methods other than POST.
		utils.RespondWithError(w, "Method not allowed")
	}
}

// PUT
func UpdateTask(w http.ResponseWriter, r *http.Request) {
	switch r.Method { // checking for http method
	case http.MethodPut: // if PUT
		taskID := utils.GetTaskNumFromRequest(r) // getting id of task
		task, exists := db[taskID]         // taking task from db and checking for existence
		if exists {                        // checking if task exists in db
			var updatedTask models.Task                                                 // temp task for holding updated task data
			if err := json.NewDecoder(r.Body).Decode(&updatedTask); err != nil { // trying to decode json from body
				utils.RespondWithError(w, "Invalid request payload") // return eror json if there an err
				return
			}

			// success
			task.TaskContent = updatedTask.TaskContent // updating TaskContent from just received updatedTask
			db[taskID] = task                          // updating db[taskID] by new task with updated data
			// responding with success json
			utils.RespondWithJSON(w, task)
		} else {
			utils.RespondWithError(w, "Task not found") // if task not found respond with error json
		}
	default:
		utils.RespondWithError(w, "Method not allowed") // if inapropriate method respond with error json
	}
}

// DELETE
func DeleteTask(w http.ResponseWriter, r *http.Request) {
	switch r.Method {
	case http.MethodDelete:
		taskID := utils.GetTaskNumFromRequest(r)
		_, exists := db[taskID] // checking if task existing
		if exists {             // if true
			delete(db, taskID)                                               // deleting from db
			utils.RespondWithJSON(w, map[string]string{"message": "Task deleted"}) // respond with success message
		} else {
			utils.RespondWithError(w, "Task not found") // respond error json if there no such task
		}
	default:
		utils.RespondWithError(w, "Method not allowed") // respond with error if inapropriate method
	}
}

// GET
func GetByID(w http.ResponseWriter, r *http.Request) {
	switch r.Method { // checking for method
	case http.MethodGet: // if GET
		taskID := utils.GetTaskNumFromRequest(r) // getting id of task from request
		task, exists := db[taskID]         // taking task by id and checking for existance
		if exists {                        // if exists return it
			utils.RespondWithJSON(w, task)
		} else { // if not - return not found
			utils.RespondWithError(w, "Task not found")
		}
	default: // return err if not appropriate method
		utils.RespondWithError(w, "Method not allowed")
	}
}
