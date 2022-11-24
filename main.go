// run the app [go run main.go]
package main

import (
	"errors"
	"net/http" //native to Golang

	"github.com/gin-gonic/gin" //to have gin working inside this app
)

// This data structure is different from json, we need to convert in the form of json, so server/client will understand
type todo struct {
	ID        string `json:"id"`
	Item      string `json:"item"`
	Completed bool   `json:"completed"`
}

// This data structure is understand by Golang
var todos = []todo{
	{ID: "1", Item: "Clean Room", Completed: false},
	{ID: "2", Item: "Read Book", Completed: false},
	{ID: "3", Item: "Wash Dishes", Completed: false},
}

// main function is the entrypoint for the execution or a "starting point". A program can only have 1 main function
func main() {
	//create a server called [router]
	router := gin.Default()
	//create endpoint
	// first: endpoint name , second: function
	router.GET("/todos", getTodos) //for get multiple todos
	router.POST("/todos", addTodo)
	router.GET("/todos/:id", getTodo) //for get todo by Id
	router.PATCH("/todos/:id", toggleTodoStatus)
	router.Run("localhost:9090") //our app will run on port 9090
}

func addTodo(context *gin.Context) {
	var newTodo todo

	if err := context.BindJSON(&newTodo); err != nil {
		// if everything is okay, it'll bind a new todo, but if err is not nil and we have and error
		// it'll throw an error for us, but we don't want to continue if there's an error, so we'll simply return nothing
		return
	}

	todos = append(todos, newTodo)
	context.IndentedJSON(http.StatusCreated, newTodo)
}

// context contains information about incoming http request i.e. header, body etc.
func getTodos(context *gin.Context) {
	//convert above data structure into JSON. Here we pass two arguments which are http status and todos items
	context.IndentedJSON(http.StatusOK, todos)
}

func getTodo(context *gin.Context) {
	id := context.Param("id")    //extract "id" from the parameter
	todo, err := getTodoById(id) //get todo by "id"

	if err != nil {
		context.IndentedJSON(http.StatusNotFound, gin.H{"message": "Todo not found"})
	}

	context.IndentedJSON(http.StatusOK, todo)
}

func getTodoById(id string) (*todo, error) {
	for i, t := range todos {
		if t.ID == id {
			return &todos[i], nil
		}
	}

	return nil, errors.New("Todo Not Found")
}

func toggleTodoStatus(context *gin.Context) {
	id := context.Param("id")    //extract "id" from the parameter
	todo, err := getTodoById(id) //get todo by "id"

	if err != nil {
		context.IndentedJSON(http.StatusNotFound, gin.H{"message": "Todo not found"})
	}

	todo.Completed = !todo.Completed //flip the boolean value
	context.IndentedJSON(http.StatusOK, todo)
}
