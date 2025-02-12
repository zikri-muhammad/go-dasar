package go_database

import (
	"context"
	"database/sql"
	"fmt"
	"strconv"
	"testing"
	"time"
)

func TestExecXContext(t *testing.T) {
	db := GetConnection()
	defer db.Close()

	ctx := context.Background()
	row := "INSERT INTO customer(id, name) values('zikri', 'zikri')"
	_, err := db.ExecContext(ctx, row)

	if err != nil {
		panic(err)
	}

	fmt.Println("Success insert new customer")
}

func TestQuerySql(t *testing.T) {
	db := GetConnection()
	defer db.Close()

	ctx := context.Background()
	row := "SELECT id, name from customer"
	data, err := db.QueryContext(ctx, row)

	if err != nil {
		panic(err)
	}

	for data.Next() {
		var id, name string
		err := data.Scan(&id, &name)
		if err != nil {
			panic(err)
		}

		fmt.Println("ID", id)
		fmt.Println("Name", name)
	}

	defer data.Close()
}

func TestQuerySqlComplex(t *testing.T) {
	db := GetConnection()
	defer db.Close()

	ctx := context.Background()
	row := "SELECT id, name, email, balance, rating, birth_date, married, created_at from customer"
	data, err := db.QueryContext(ctx, row)

	if err != nil {
		panic(err)
	}

	for data.Next() {
		var id, name string
		var email sql.NullString
		var balance int32
		var rating float64
		var birthDate sql.NullTime
		var createdAt time.Time
		var married bool
		err := data.Scan(&id, &name, &email, &balance, &rating, &birthDate, &married, &createdAt)
		if err != nil {
			panic(err)
		}

		fmt.Println("ID", id)
		fmt.Println("Name", name)
		if email.Valid {
			fmt.Println("Email", email)
		}
		fmt.Println("balance", balance)
		fmt.Println("Rating", rating)
		if birthDate.Valid {
			fmt.Println("Birth Date", birthDate)
		}
		fmt.Println("Created At", createdAt)
		fmt.Println("Married", married)
	}

	defer data.Close()
}

func TestSqlInjection(t *testing.T) {
	db := GetConnection()
	defer db.Close()

	ctx := context.Background()

	username := "admin"
	password := "admin"

	row := "SELECT username from user WHERE username = '" + username + "' AND password = '" + password + "' LIMIT 1"
	data, err := db.QueryContext(ctx, row)

	if err != nil {
		panic(err)
	}

	if data.Next() {
		var username string
		err := data.Scan(&username)
		if err != nil {
			panic(err)
		}

		fmt.Println("Success Login", username)
	} else {
		fmt.Println("Galal Login")
	}

	defer data.Close()
}

func TestSqlInjectionSafe(t *testing.T) {
	db := GetConnection()
	defer db.Close()

	ctx := context.Background()

	username := "admin"
	password := "admin"

	row := "SELECT username from user WHERE username = ? AND password = ? LIMIT 1"
	data, err := db.QueryContext(ctx, row, username, password)

	if err != nil {
		panic(err)
	}

	if data.Next() {
		var username string
		err := data.Scan(&username)
		if err != nil {
			panic(err)
		}

		fmt.Println("Success Login", username)
	} else {
		fmt.Println("Galal Login")
	}

	defer data.Close()
}

func TestExecParameter(t *testing.T) {
	db := GetConnection()
	defer db.Close()

	ctx := context.Background()

	username := "ocu"
	password := "ocu"

	row := "INSERT INTO user(username, password) values(?, ?)"
	_, err := db.ExecContext(ctx, row, username, password)

	if err != nil {
		panic(err)
	}

	fmt.Println("Success insert new user")
}

func TestAutoIncreament(t *testing.T) {
	db := GetConnection()
	defer db.Close()

	ctx := context.Background()

	email := "zikri@ocudevs.com"
	comment := "This is commnet"

	row := "INSERT INTO comments(email, comment) values(?, ?)"
	result, err := db.ExecContext(ctx, row, email, comment)

	if err != nil {
		panic(err)
	}

	insertId, err := result.LastInsertId()
	if err != nil {
		panic(err)
	}

	fmt.Println("Success insert new commnet", insertId)
}

func TestPrepareStatment(t *testing.T) {
	db := GetConnection()
	defer db.Close()

	ctx := context.Background()

	row := "INSERT INTO comments(email, comment) values(?, ?)"
	stmt, err := db.PrepareContext(ctx, row)

	if err != nil {
		panic(err)
	}

	defer stmt.Close()

	for i := 0; i < 10; i++ {
		email := "zikri" + strconv.Itoa(i) + "@ocudevs.com"
		comment := "This is commnet " + strconv.Itoa(i)

		result, err := stmt.ExecContext(ctx, email, comment)
		if err != nil {
			panic(err)
		}
		insertId, err := result.LastInsertId()

		if err != nil {
			panic(err)
		}

		fmt.Println("Success insert new commnet", insertId)
	}

}

func TestTransaction(t *testing.T) {
	db := GetConnection()
	defer db.Close()

	ctx := context.Background()
	tx, err := db.Begin()
	if err != nil {
		panic(err)
	}

	row := "INSERT INTO comments(email, comment) values(?, ?)"

	for i := 0; i < 10; i++ {
		email := "zikri" + strconv.Itoa(i) + "@ocudevs.com"
		comment := "This is commnet " + strconv.Itoa(i)

		result, err := tx.ExecContext(ctx, row, email, comment)
		if err != nil {
			panic(err)
		}
		insertId, err := result.LastInsertId()

		if err != nil {
			panic(err)
		}

		fmt.Println("Success insert new commnet", insertId)
	}

	// err = tx.Rollback() // rolback
	err = tx.Commit() // Commit

	if err != nil {
		panic(err)
	}

}
