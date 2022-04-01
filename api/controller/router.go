package controller

import (
	"fmt"
	"net/http"
)

// type router struct {
// 	tc TodoController
// }

func HandleTodoRequest(w http.ResponseWriter, r *http.Request) {
	fmt.Println("listen on controller")
	fmt.Fprintf(w, "Hello, World")
}
