package main

import (
	"database/sql"
	_ "github.com/lib/pq"
	"fmt"
	"github.com/gorilla/mux"
)

const (
  host     = "localhost"
  //host     = "database"
  port     = 5432
  user     = "postgres"
  password = "Abcd@1234"
  dbname   = "test_demo"
)

// Holds references to router and db
type App struct {
	DB     *sql.DB
	Router *mux.Router
}

func (a *App) Initialize() {
	psqlInfo := fmt.Sprintf("host=%s port=%d user=%s "+
	"password=%s dbname=%s sslmode=disable",
	host, port, user, password, dbname)

	var err error
	a.DB, err = sql.Open("postgres", psqlInfo)
	if err != nil {
		panic(err)
		// log.Fatal(err)
	}
	//defer a.DB.Close()

	err = a.DB.Ping()
	if err != nil {
		panic(err)
	}
	fmt.Println("Successfully connected!")

	_, err = a.DB.Exec("DROP TABLE IF EXISTS articles")
	if err != nil {
		panic(err)
	}

	_, err = a.DB.Exec(`CREATE TABLE articles (
	id SERIAL PRIMARY KEY,
	title TEXT,
	content TEXT,
	author TEXT);`)
	
	if err != nil {
		panic(err)
	}

	// Creates dummy data
	_, err = a.DB.Exec(`
	INSERT INTO articles (title, content, author)
	VALUES 
	('Book Title 1', 'Rando Text 1', 'Smith'),
	('Book Title 2', 'Rando Text 2', 'Seinfeld');`)
	if err != nil {
		panic(err)
	}

	a.Router = mux.NewRouter()
	a.initializeRoutes()

}

func (a *App) initializeRoutes() {
    a.Router.HandleFunc("/articles", a.getArticles).Methods("GET")
	a.Router.HandleFunc("/articles/{article_id}", a.getArticle).Methods("GET")
	a.Router.HandleFunc("/articles", a.createArticle).Methods("POST")
}