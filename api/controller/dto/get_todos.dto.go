package dto

type TodoResponse struct {
	Id      int    `json:"id"`
	Content string `json:"content"`
}

type TodosResponse struct {
	Todos []TodoResponse `json:"todos"`
}
