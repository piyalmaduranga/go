package main

import (	
"fmt"
"log"
"net/http"
"encoding/json"
"database/sql"

)

type Article struct {
    Title string `json:"Title"`
    Desc string `json:"desc"`
    Content string `json:"content"`
}

type Articles []Article
 
func homePage(w http.ResponseWriter, r *http.Request){

	db, err := sql.Open("mysql", "root:@/nsoft")
    if err != nil {
        panic(err.Error())  // Just for example purpose. You should use proper error handling instead of panic
    }
    defer db.Close()

    // Execute the query
    rows, err := db.Query("SELECT * FROM table")
    if err != nil {
        panic(err.Error()) // proper error handling instead of panic in your app
    }

    // Get column names
    columns, err := rows.Columns()
    if err != nil {
        panic(err.Error()) // proper error handling instead of panic in your app
    }
	fmt.Println(columns)
    fmt.Fprintf(w, "Welcome to the HomePage!")
    fmt.Println("Endpoint Hit: homePage")
}

func returnArticle(w http.ResponseWriter, r *http.Request) {
    fmt.Fprintf(w, "returns a specific article")
    fmt.Println("Endpoint Hit: returnArticle")
}

func returnAllArticles(w http.ResponseWriter, r *http.Request){
    articles := Articles{
        Article{Title: "Hello", Desc: "Article Description", Content:"Content of hello"},
        Article{Title: "Hello 2", Desc: "Article Description", Content: "Content of hello 2"},
    }    
    fmt.Println("Endpoint Hit: returnAllArticles")

    json.NewEncoder(w).Encode(articles)
}

func addArticle(w http.ResponseWriter, r *http.Request){
    fmt.Fprintf(w, "Adds an article to list of articles")
    fmt.Println("Endpoint Hit: addArticle")
}

func delArticle(w http.ResponseWriter, r *http.Request){
    fmt.Fprintf(w, "deletes a specific article")
    fmt.Println("Endpoint Hit: delArticle")
}

func handleRequests() {
    http.HandleFunc("/", homePage)
    http.HandleFunc("/all", returnAllArticles)
    http.HandleFunc("/single", returnArticle)
    http.HandleFunc("/delete", delArticle)
    http.HandleFunc("/add", addArticle)
    log.Fatal(http.ListenAndServe(":8081", nil))
}

func main() {

    handleRequests()
	
}