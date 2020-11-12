// model.go

package main

import (
    "database/sql"
)

type ArticleID struct {
	ID int `json:"id"`
}

type CustomResponse struct {
	Status int `json:"status"`
	Message string `json:"message"`
	Data interface{} `json:"data"`
}

// Article struct (Model)
type Article struct {
	ID int `json:"id"`
	Title  string `json:"title"`
	Content string `json:"content"`
	Author string `json:"author"`
}

func (p *Article) getArticle(db *sql.DB) error {
    return db.QueryRow("SELECT * FROM articles WHERE id=$1", p.ID).Scan(&p.ID, &p.Title, &p.Content, &p.Author)
}

func (p *Article) createArticle(db *sql.DB) error {
    err := db.QueryRow(
        "INSERT INTO articles(title, content, author) VALUES($1, $2, $3) RETURNING id",
        p.Title, p.Content, p.Author).Scan(&p.ID)

    if err != nil {
        return err
    }

    return nil
}

func getAllArticles(db *sql.DB) ([]Article, error) {
    rows, err := db.Query("SELECT id, title, content, author FROM articles")

    if err != nil {
        return nil, err
    }

    defer rows.Close()
    var articleArr []Article

    for rows.Next() {
        var p Article
        if err := rows.Scan(&p.ID, &p.Title, &p.Content, &p.Author); err != nil {
            return nil, err
        }
        articleArr = append(articleArr, p)
    }

    return articleArr, nil
}