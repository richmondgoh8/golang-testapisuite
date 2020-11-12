package main

import (
	"encoding/json"
	"log"
	"net/http"
	"strconv"
	"github.com/gorilla/mux"
	_ "github.com/lib/pq"
	//"fmt"
	"database/sql"
)

// Init articles var as a slice Article struct
var articles []Article

// Get all articles
func (a *App)getArticles(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	_Articles, err := getAllArticles(a.DB)
	if err != nil {
		panic(err)
	}
	res := CustomResponse{Status:http.StatusOK,Message:"Success",Data:_Articles}
	json.NewEncoder(w).Encode(res)

	//json.NewEncoder(w).Encode(articles)
}

// Get single article
func (a *App)getArticle(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	params := mux.Vars(r) // Gets params

	value, err := strconv.Atoi(params["article_id"])
	//w.WriteHeader(http.StatusOK)
	if err != nil {
     	// handle error
     	w.WriteHeader(http.StatusBadRequest)
		res := CustomResponse{Status:http.StatusBadRequest,Message:"Article ID should be a number",Data:nil}
		json.NewEncoder(w).Encode(res)
     	return
   	}

   	p := Article{ID: value}
   	if err = p.getArticle(a.DB); err != nil {
   		switch err {
   			case sql.ErrNoRows:
   				w.WriteHeader(http.StatusNotFound)
   				res := CustomResponse{Status:http.StatusNotFound,Message:"Article does not exist",Data:nil}
				json.NewEncoder(w).Encode(res)
   			default:
   				w.WriteHeader(http.StatusInternalServerError)
   				res := CustomResponse{Status:http.StatusInternalServerError,Message:err.Error(),Data:nil}
				json.NewEncoder(w).Encode(res)
   		}
   		return
   	}
   	res := CustomResponse{Status:http.StatusOK,Message:"Success",Data:p}
   	json.NewEncoder(w).Encode(res)
}

// Add new Article
func (a *App)createArticle(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	
	var article Article
	_ = json.NewDecoder(r.Body).Decode(&article)
	for _, data := range []interface{}{article.Title, article.Content, article.Author} {
	//for _, data := range []Article {
		if (data == "" || data == nil || data == 0) {

			w.WriteHeader(http.StatusBadRequest)
			res := CustomResponse{Status:http.StatusBadRequest,Message:"Please provide article title, content and author",Data:nil}
			json.NewEncoder(w).Encode(res)
     		return
		}
	}

	if err := article.createArticle(a.DB); err != nil {
		res := CustomResponse{Status:http.StatusInternalServerError,Message:err.Error(),Data:nil}
		json.NewEncoder(w).Encode(res)
		return
	}

	data := struct {ID int `json:"id"`}{article.ID}
//
	w.WriteHeader(http.StatusCreated)
	res := CustomResponse{Status:http.StatusCreated,Message:"Success",Data:data}
	json.NewEncoder(w).Encode(res)

	//json.NewEncoder(w).Encode(article)
}


// Main function
func main() {
	// Init router
	a := App{}
	a.Initialize()

	// Start server
	log.Fatal(http.ListenAndServe(":8080", a.Router))
}