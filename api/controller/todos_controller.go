package controller

import (
	"encoding/json"
	"net/http"

	"github.com/kouki-iwahara/golang-todos/api/controller/dto"
	"github.com/kouki-iwahara/golang-todos/api/models/repository"
)

// 外部に公開するのは実態ではなくインターフェース
type TodoController interface {
	GetTodos(w http.ResponseWriter, r *http.Request)
}

// controllerの実態は非公開
type todoController struct {
	tr repository.TodoRepository
}

// todoControllerのコンストラクタ。repositoryを受け取り、controllerのポインタを返す。
func NewTodoController(todoRepository repository.TodoRepository) TodoController {
	return &todoController{
		todoRepository,
	}
}

// todo取得処理
func (tc *todoController) GetTodos(w http.ResponseWriter, r *http.Request) {
	/*
		初期化処理でrepositoryを受け取ったので、repositoryのメソッドが呼び出せる。
		Classで言うとthisで参照しているイメージ。
	*/
	result, err := tc.tr.GetTodos()
	if err != nil {
		w.WriteHeader(500)
		return
	}

	// DBから取得したtodoをDTOに詰める
	var todos []dto.TodoResponse
	for _, v := range result {
		todos = append(todos, dto.TodoResponse{Id: v.Id, Content: v.Content})
	}
	var todosResponse dto.TodosResponse
	todosResponse.Todos = todos

	// JSONに変換
	output, _ := json.MarshalIndent(todosResponse, "", "\t\t")

	// JSONを返却
	w.Header().Set("Contet-Type", "application/json")
	w.Write(output)
}
