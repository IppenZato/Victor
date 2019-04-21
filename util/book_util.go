package util

import (
	"database/sql"
	"encoding/json"
	"fmt"
	"log"
	"net/http"

	"github.com/gorilla/mux"
)

// Book type
type Book struct {
	ID     int    `json:"id"`
	Title  string `json:"title"`
	Author string `json:"author"`
	Year   int    `json:"year"`
}

func getBooks(w http.ResponseWriter, r *http.Request) {
	var books []Book
	var book Book

	rows, err := DB.Query("select * from " + TableBook)
	CheckErr(err)

	defer rows.Close()

	for rows.Next() {
		err := rows.Scan(&book.ID, &book.Title, &book.Author, &book.Year)
		CheckErr(err)

		books = append(books, book)
	}
	json.NewEncoder(w).Encode(books)
}

func getBook(w http.ResponseWriter, r *http.Request) {
	var book Book
	id := r.FormValue("id")

	sqlStmt := fmt.Sprintf("select * from %s where id=?", TableBook)
	log.Println(sqlStmt)

	err := DB.QueryRow(sqlStmt, id).Scan(&book.ID, &book.Title, &book.Author, &book.Year)
	CheckErr(err)

	json.NewEncoder(w).Encode(book)
}

func addBook(w http.ResponseWriter, r *http.Request) {

	// read request parameters
	ID := r.FormValue("id") // TODO shoud increment automatically
	title := r.FormValue("title")
	author := r.FormValue("author")
	year := r.FormValue("year")

	var book Book

	sqlStmt := fmt.Sprintf("insert %s set id=?, title=?, author=?, relise_year=?", TableBook)

	stmt, err := DB.Prepare(sqlStmt)
	CheckErr(err)

	var res sql.Result
	res, err = stmt.Exec(ID, title, author, year)
	CheckErr(err)

	id, err := res.LastInsertId()
	CheckErr(err)

	// return data to client
	sql := fmt.Sprintf("select * from %s where id=?", TableBook)

	err = DB.QueryRow(sql, id).Scan(&book.ID, &book.Title, &book.Author, &book.Year)
	CheckErr(err)

	json.NewEncoder(w).Encode(book)
}

func upadateBook(w http.ResponseWriter, r *http.Request) {
	
	// read request parameters
	ID := r.FormValue("id")
	title := r.FormValue("title")
	author := r.FormValue("author")
	year := r.FormValue("year")

	var book Book

	sqlStmt := fmt.Sprintf("update %s set title=?, author=?, relise_year=? where id=?", TableBook)
	stmt, err := DB.Prepare(sqlStmt)
	CheckErr(err)


	res, err := stmt.Exec(title, author, year, ID)
	CheckErr(err)

	affect, err := res.RowsAffected()
	CheckErr(err)

	fmt.Println(affect)

	// return data to client
	sql := fmt.Sprintf("select * from %s where id=?", TableBook)

	err = DB.QueryRow(sql, ID).Scan(&book.ID, &book.Title, &book.Author, &book.Year)
	CheckErr(err)

	json.NewEncoder(w).Encode(book)
}

func removeBook(w http.ResponseWriter, r *http.Request) {
	// read request parameters
	ID := r.FormValue("id")

	sqlStmt := fmt.Sprintf("delete from %s where id=?", TableBook)

	stmt, err := DB.Prepare(sqlStmt)
	CheckErr(err)

	res, err := stmt.Exec(ID)
	CheckErr(err)

	affect, err := res.RowsAffected()
	CheckErr(err)
	fmt.Println(affect)
}

// BookService
func BookService() {
	router := mux.NewRouter()

	router.HandleFunc("/book", getBook).Methods("GET")
	router.HandleFunc("/books", getBooks).Methods("GET")
	router.HandleFunc("/books", addBook).Methods("POST")
	router.HandleFunc("/books", upadateBook).Methods("PUT")
	router.HandleFunc("/books", removeBook).Methods("DELETE")

	log.Fatal(http.ListenAndServe(":8000", router))
}
