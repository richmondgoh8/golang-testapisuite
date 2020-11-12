// +build pro
package main

import (
  "database/sql"
  "fmt"

  _ "github.com/lib/pq"
)

type User struct {
  ID        int
  Age       int
  FirstName string
  LastName  string
  Email     string
}

func test() {
	psqlInfo := fmt.Sprintf("host=%s port=%d user=%s "+
	"password=%s dbname=%s sslmode=disable",
	host, port, user, password, dbname)

	db, err := sql.Open("postgres", psqlInfo)
	if err != nil {
		panic(err)
	}
	defer db.Close()

	err = db.Ping()
	if err != nil {
		panic(err)
	}

	_, err = db.Exec("DROP TABLE IF EXISTS users")
	if err != nil {
		panic(err)
	}
	_, err = db.Exec(`CREATE TABLE users (
	id SERIAL PRIMARY KEY,
	age INT,
	first_name TEXT,
	last_name TEXT,
	email TEXT UNIQUE NOT NULL);`)
	
	if err != nil {
		panic(err)
	}

	// Creates default data
	_, err = db.Exec(`
	INSERT INTO users (age, email, first_name, last_name)
	VALUES 
	(52, 'bob@smith.io', 'Bob', 'Smith'),
	(15, 'jerryjr123@gmail.com', 'Jerry', 'Seinfeld');`)
	if err != nil {
		panic(err)
	}
	// Prevents SQL Injection if using fmt
	sqlStatement := `
	INSERT INTO users (age, email, first_name, last_name)
	VALUES ($1, $2, $3, $4)
	RETURNING id`

	id := 0
	err = db.QueryRow(sqlStatement, 30, "jon@calhoun.io", "Jonathan", "Calhoun").Scan(&id)
	if err != nil {
		panic(err)
	}
	fmt.Println("New record ID is:", id)

	sqlStatement = `SELECT * FROM users WHERE id=$1;`
	var user User

	row := db.QueryRow(sqlStatement, 2)
	err = row.Scan(&user.ID, &user.Age, &user.FirstName,&user.LastName, &user.Email)
	switch err {
	case sql.ErrNoRows:
		fmt.Println("No rows were returned!")
		return
	case nil:
		fmt.Println(user)
	default:
		panic(err)
	}

	rows, err := db.Query("SELECT id, first_name FROM users LIMIT $1", 3)
	if err != nil {
		panic(err)
	}

	defer rows.Close()
	for rows.Next() {
		var id int
		var firstName string
		err = rows.Scan(&id, &firstName)
		if err != nil {
			panic(err)
		}
		fmt.Println(id, firstName)
	}

	err = rows.Err()
	if err != nil {
		panic(err)
	}


	fmt.Println("Successfully connected!")
}