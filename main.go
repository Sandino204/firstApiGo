package main

import (
	"net/http"
	"github.com/gin-gonic/gin"
	"errors"
	"strconv"
)

type todo struct {
	ID int `json:"id"`
	Item string `json:"item"`
	Completed bool `json:"completed"`
}

var todos = []todo {
	{ID: 1, Item: "Clear Room", Completed: false },
	{ID: 1, Item: "Clear Room", Completed: false },
	{ID: 1, Item: "Clear Room", Completed: false },
}

func getTodos(context * gin.Context) {
	context.IndentedJSON(http.StatusOK, todos)
}

func addTodo(context * gin.Context) {
	var newTodo todo

	if err := context.BindJSON(&newTodo); err != nil {
		return 
	}

	todos = append(todos, newTodo)

	context.IndentedJSON(http.StatusCreated, newTodo)
}

func getTodoById(id int) (*todo, error) {
	for i, t := range todos {
		if t.ID == id {
			return &todos[i], nil
		}
	}	

	return nil, errors.New("todo not found")
}

func getTodo(context *gin.Context) {
	strId := context.Param("id")
	
	id, err := strconv.Atoi(strId); 

	if err != nil {
		return
	}

	todo, err := getTodoById(id)

	if err != nil {
		context.IndentedJSON(http.StatusNotFound, gin.H{"message": "todo not found"})
	}

	context.IndentedJSON(http.StatusOK, todo)
}

func toggleTodoStatus(context *gin.Context) {
	strId := context.Param("id")
	
	id, err := strconv.Atoi(strId); 

	if err != nil {
		return
	}

	todo, err := getTodoById(id)

	if err != nil {
		context.IndentedJSON(http.StatusNotFound, gin.H{"message": "todo not found"})
	}

	todo.Completed = !todo.Completed

	context.IndentedJSON(http.StatusOK, todo)
}

func main() {
	router := gin.Default()
	router.GET("/todos/:id", getTodo)
	router.GET("/todos", getTodos)
	router.POST("/todos", addTodo)
	router.PATCH("/todos/:id", toggleTodoStatus)
	router.Run("localhost:9090")
}

