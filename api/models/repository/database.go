package repository

import (
	"database/sql"
	"fmt"
	"os"

	_ "github.com/go-sql-driver/mysql"
)

var Db *sql.DB

func init() {
	var err error
	var username = os.Getenv("MYSQL_USER")
	var password = os.Getenv("MYSQL_ROOT_PASSWORD")
	var address = fmt.Sprintf("%s:%s", os.Getenv("MYSQL_DATABASE"), os.Getenv("MYSQL_PORT"))
	var dbname = os.Getenv("MYSQL_DATABASE")
	// [username[:password]@][protocol[(address)]]/dbname[?param1=value1&...&paramN=valueN]
	source := fmt.Sprintf("%s:%s@tcp(%s)/%s?charset=utf8mb4", username, password, address, dbname)
	fmt.Println(source)

	Db, err = sql.Open("mysql", source)
	if err != nil {
		panic(err)
	}
}
