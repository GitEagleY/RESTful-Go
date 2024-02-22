# RESTful Api in Golang


## Prerequisites
- [Golang](https://go.dev/doc/install)

## Running

	git clone https://github.com/GitEagleY/RESTful-Go.git
	cd RESTful-Go
	go run main.go

## Usage

Once the server is running, you can interact with the API using HTTP requests. Here are some example endpoints:

    GET /tasks/display: Retrieve all tasks.
    POST /tasks/addById/{id}: Add a new task with the specified ID.
    PUT /tasks/update/{id}: Update an existing task.
    DELETE /tasks/delete/{id}: Delete a task.
    GET /tasks/getTaskById/{id}: Retrieve a task by ID.
