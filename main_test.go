// main_test.go

package main

import (
    "testing"   
    "log"
    "os"
    "bytes"
    "net/http"
    "net/http/httptest"
    "encoding/json"
)

type ReturnResponse struct {
    Status int `json:"status"`
    Message string `json:"message"`
    Data []Article `json:"data"`
}

var a App

func TestMain(m *testing.M) {
    a.Initialize()

    // Ensure Proper Connection to Database
    ensureTableExists()
    os.Exit(m.Run())
}

func ensureTableExists() {
    _, err := a.DB.Exec(`CREATE TABLE IF NOT EXISTS articles (
    id SERIAL PRIMARY KEY,
    title TEXT,
    content TEXT,
    author TEXT);`)
    
    if err != nil {
        log.Fatal(err)
    }
}

func TestGetArticles(t *testing.T) {

    req, _ := http.NewRequest("GET", "/articles", nil)
    response := executeRequest(req)

    checkResponseCode(t, http.StatusOK, response.Code)

    var m ReturnResponse
    json.Unmarshal(response.Body.Bytes(), &m)
    if len(m.Data) != 2 {
        t.Errorf("Expected 2 Dummy Data Records. Got '%v'", len(m.Data))
    }

}

func TestGetArticle(t *testing.T) {
    // Invalid Article ID
    req, _ := http.NewRequest("GET", "/articles/hh", nil)
    response := executeRequest(req)

    checkResponseCode(t, http.StatusBadRequest, response.Code)

    // Valid Article ID
    req, _ = http.NewRequest("GET", "/articles/1", nil)
    response = executeRequest(req)

    checkResponseCode(t, http.StatusOK, response.Code)
}

func TestPostArticle(t *testing.T) {
    var jsonStr = []byte(`{"title": "Friend Foreever", "Content": "My Personal Content", "Author": "Richmond Goh2"}`)
    req, _ := http.NewRequest("POST", "/articles", bytes.NewBuffer(jsonStr))
    req.Header.Set("Content-Type", "application/json")

    response := executeRequest(req)
    checkResponseCode(t, http.StatusCreated, response.Code)
}

func executeRequest(req *http.Request) *httptest.ResponseRecorder {
    rr := httptest.NewRecorder()
    a.Router.ServeHTTP(rr, req)

    return rr
}

func checkResponseCode(t *testing.T, expected, actual int) {
    if expected != actual {
        t.Errorf("Expected response code %d. Got %d\n", expected, actual)
    }
}
