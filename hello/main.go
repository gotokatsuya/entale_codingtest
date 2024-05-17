package main

import (
	"database/sql"
	"fmt"
	"net/http"
	"os"

	_ "github.com/go-sql-driver/mysql"
)

type Article struct {
	ID     int     `json:"id"`
	Title  string  `json:"title"`
	Body   string  `json:"body"`
	Medias []Media `json:"medias"`
}

type Media struct {
	ID          int    `json:"id"`
	ContentURL  string `json:"contentUrl"`
	ContentType string `json:"contentType"`
}

func main() {
	pswd := os.Getenv("MYSQL_ROOT_PASSWORD")
	db, err := sql.Open("mysql", "root:"+pswd+"@tcp(localhost:3306)/entale")
	if err != nil {
		fmt.Println("Not connected to the database")
		panic(err.Error())
	}
	defer db.Close()
	err = db.Ping()
	if err != nil {
		fmt.Println("error in connection")
		panic(err.Error())
	}
	// Fetch articles from the URL
	response, err := http.Get("https://gist.githubusercontent.com/gotokatsuya/cc78c04d3af15ebe43afe5ad970bc334/raw/dc39bacb834105c81497ba08940be5432ed69848/articles.json")
	if err != nil {
		fmt.Print(err.Error())
		os.Exit(1)
	}
	// defer response.Body.Close()

	// // Read the response body
	// responseData, err := ioutil.ReadAll(response.Body)
	// if err != nil {
	// 	log.Fatal(err)
	// }

	// // Parse JSON into articles
	// var articles []Article
	// err = json.Unmarshal(responseData, &articles)
	// if err != nil {
	// 	log.Fatal(err)
	// }

	// // Print details of each article
	// for _, article := range articles {
	// 	fmt.Println("Title:", article.Title)
	// 	fmt.Println("Body:", article.Body)
	// 	fmt.Println("Media Count:", len(article.Medias))
	// 	fmt.Println("--------------")
	// }

	// // Prepare the SQL query
	// query, params := prepareQuery(articles)

	// // Print the generated SQL query and its parameters
	// fmt.Println("Generated SQL Query:")
	// fmt.Println(query)
	// fmt.Println("Query Parameters:")
	// fmt.Println(params)

}
