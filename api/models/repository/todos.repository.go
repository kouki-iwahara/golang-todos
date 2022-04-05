package repository

import (
	"log"

	_ "github.com/go-sql-driver/mysql"
	"github.com/kouki-iwahara/golang-todos/api/models/entity"
)

// 外部に公開するのは実態ではなくインターフェース
type TodoRepository interface {
	GetTodos() (todos []entity.TodoEntity, err error)
}

// repositoryのコンストラクタ。何も受け取らずポインタを返却。
func NewTodoRepository() TodoRepository {
	return &todoRepository{}
}

// repositoryの実態は非公開
type todoRepository struct{}

// todoの取得処理
func (tr *todoRepository) GetTodos() (todos []entity.TodoEntity, err error) {
	const sql = "select id, content from todos order by id ASC"
	rows, err := Db.Query(sql)
	if err != nil {
		log.Fatalln(err)
		return
	}
	// Db.Query()で返るrowsはコネクションを占領し続けるので適切にrows.Close()する
	defer rows.Close()

	// 返却用のスライスを定義
	todos = []entity.TodoEntity{}
	// 1レコードずつtodo entityにマッピング。返却用のスライスに詰める。
	for rows.Next() {
		todo := entity.TodoEntity{}
		// 取得したレコードのデータを型変換する。型変換時にエラーがある時のみ戻り値がある。
		err = rows.Scan(&todo.Id, &todo.Content)
		if err != nil {
			log.Fatalln(err)
		}
		todos = append(todos, todo)
	}

	// rows.Next()最中で発生したエラーを拾う
	err = rows.Err()
	if err != nil {
		log.Fatalln(err)
	}

	// 名前付き戻り値を定義しているので、明示的な戻り値の記述は不要。
	return
}
