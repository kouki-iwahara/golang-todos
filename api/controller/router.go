package controller

import (
	"net/http"
)

type Router interface {
	HandleTodoRequest(w http.ResponseWriter, r *http.Request)
}

type router struct {
	tc TodoController
}

func (ro *router) HandleTodoRequest(w http.ResponseWriter, r *http.Request) {
	switch r.Method {
	case "GET":
		ro.tc.GetTodos(w, r)
	default:
		w.WriteHeader(405)
	}
}

// routerのコンストラクタ。TodoControllerを受け取りrouterのポインタを返却する
func NewRouter(tc TodoController) Router {
	return &router{tc}
}
