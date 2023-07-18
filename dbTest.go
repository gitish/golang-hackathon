package main

import (
	"database/sql"
	"fmt"

	_ "github.com/lib/pq"
)

const (
	host     = "localhost"
	port     = 5432
	user     = "shail"
	password = "shail"
	dbname   = "postgres"
)

func main() {
	// connection string
	psqlconn := fmt.Sprintf("host=%s port=%d user=%s password=%s dbname=%s sslmode=disable", host, port, user, password, dbname)

	// open database
	db, err := sql.Open("postgres", psqlconn)
	CheckError(err)

	// close database
	defer db.Close()

	// check db
	err = db.Ping()
	CheckError(err)

	fmt.Println("Connected!")

	// insert
	// hardcoded
	insertStmt := `insert into "Students"("Name", "Roll") values('John', 1)`
	_, e := db.Exec(insertStmt)
	CheckError(e)

	// dynamic
	insertDynStmt := `insert into "Students"("Name", "Roll") values($1, $2)`
	_, e = db.Exec(insertDynStmt, "Jane", 2)
	CheckError(e)

	rows, err := db.Query(`SELECT "Name", "Roll" FROM "Students"`)
	CheckError(err)

	defer rows.Close()
	for rows.Next() {
		var name string
		var roll int

		err = rows.Scan(&name, &roll)
		CheckError(err)

		fmt.Println(name, roll)
	}

	CheckError(err)
}

func CheckError(err error) {
	if err != nil {
		panic(err)
	}
}
