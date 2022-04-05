package main

import (
	"fmt"
	"net/http"

	"github.com/kouki-iwahara/golang-todos/api/controller"
	"github.com/kouki-iwahara/golang-todos/api/models/repository"
)

// DI
var todoRepository = repository.NewTodoRepository()
var todoController = controller.NewTodoController(todoRepository)
var todoRouter = controller.NewRouter(todoController)

func main() {
	server := http.Server{
		Addr: ":8080",
	}

	http.HandleFunc("/todos", todoRouter.HandleTodoRequest)
	fmt.Println("server listening on port 8080")
	server.ListenAndServe()
}
