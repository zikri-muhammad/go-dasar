package go_database

import (
	"database/sql"
	"testing"

	_ "github.com/go-sql-driver/mysql"
)

func TestOpenConnection(t *testing.T) {
	db, err := sql.Open("mysql", "root:@tcp(localhost:3306)/data_go")
	if err != nil {
		panic(err)
	}

	defer db.Close()

}

func HelloWorl() string {
	return "Hello World"
}
