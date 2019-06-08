package server

import (
	"github.com/Viktor19931/books_api/log"
	"encoding/json"
	"fmt"
	"net/http"
	"github.com/Viktor19931/books_api/db/dbm"
	"github.com/Viktor19931/books_api/db"
	"github.com/gorilla/mux"
)
//----------------------------------------------------------------------------------------------------------------------
// Routers
//----------------------------------------------------------------------------------------------------------------------
func InitRouters() {
	router := mux.NewRouter()
	router.HandleFunc("/", home).Methods("GET","POST","PUT","DELETE")
	router.HandleFunc("/book", getBook).Methods("GET")
	router.HandleFunc("/books", getBooks).Methods("GET")
	router.HandleFunc("/books", addBook).Methods("POST")
	router.HandleFunc("/books", upadateBook).Methods("PUT")
	router.HandleFunc("/books", removeBook).Methods("DELETE")
	//...other Handlers
	http.Handle("/", router)
}

//----------------------------------------------------------------------------------------------------------------------
// Implementation
//----------------------------------------------------------------------------------------------------------------------
func home(w http.ResponseWriter, r *http.Request) {
	w.WriteHeader(http.StatusOK)
	fmt.Fprintf(w, "Hello Viktor!!!") // отправляем данные на клиентскую сторону
}

func getBook(w http.ResponseWriter, r *http.Request) {
	var err error
	// id any error return http.StatusBadRequest(400) else http.StatusOK(200)
	defer func() {
		if err == nil {
			w.WriteHeader(http.StatusOK)
		} else {
			log.Error(err)
			//w.Write([]byte(err.Error()))
			http.Error(w, err.Error(), http.StatusBadRequest)
		}
	}()
	defer log.AppRecover(getBook, &err)	// перехватчик аварийного завершения

	// read form data
	id := r.FormValue("id")

	book := dbm.Book{}

	query := fmt.Sprintf("select * from %s where id = ?", book.TableName())
	log.Debug("query:%s", query)
	if err = db.Instance().Get(&book, db.Instance().Rebind(query), id);	err != nil {
		return
	}
	log.Info("book:`%+v`", book)

	json.NewEncoder(w).Encode(book)
	return
}

func getBooks(w http.ResponseWriter, r *http.Request) {
	var err error
	// id any error return http.StatusBadRequest(400) else http.StatusOK(200)
	defer func() {
		if err == nil {
			w.WriteHeader(http.StatusOK)
		} else {
			log.Error(err)
			//w.Write([]byte(err.Error()))
			http.Error(w, err.Error(), http.StatusBadRequest)
		}
	}()
	defer log.AppRecover(getBook, &err)	// перехватчик аварийного завершения

	book := dbm.Book{}
	books := []*dbm.Book{}

	if err := db.Instance().Select(&books, "select * from " + book.TableName() ); err != nil {
		log.Error(err)
		return
	}
	json.NewEncoder(w).Encode(books)
	return
}

func addBook(w http.ResponseWriter, r *http.Request) {
	var err error
	// id any error return http.StatusBadRequest(400) else http.StatusOK(200)
	defer func() {
		if err == nil {
			w.WriteHeader(http.StatusOK)
		} else {
			log.Error(err)
			//w.Write([]byte(err.Error()))
			http.Error(w, err.Error(), http.StatusBadRequest)
		}
	}()
	defer log.AppRecover(getBook, &err)	// перехватчик аварийного завершения

	// read request parameters
	id 	   := r.FormValue("id") // TODO shoud increment automatically
	title  := r.FormValue("title")
	author := r.FormValue("author")
	year   := r.FormValue("year")

	book := dbm.Book{}

	query := fmt.Sprintf("insert %s set id=?, title=?, author=?, relise_year=?", book.TableName())

	res, err := db.Instance().Exec(db.Instance().Rebind(query), id, title, author, year )
	if err != nil {
		return
	}

	rec_id, err := res.LastInsertId() // Last id

	// return data to client
	r.Form.Set("id",  fmt.Sprint("%d",rec_id))
	getBook(w, r) // read book record

	return
}

func upadateBook(w http.ResponseWriter, r *http.Request) {
	var err error
	// id any error return http.StatusBadRequest(400) else http.StatusOK(200)
	defer func() {
		if err == nil {
			w.WriteHeader(http.StatusOK)
		} else {
			log.Error(err)
			//w.Write([]byte(err.Error()))
			http.Error(w, err.Error(), http.StatusBadRequest)
		}
	}()
	defer log.AppRecover(getBook, &err)	// перехватчик аварийного завершения

	// read request parameters
	id 		:= r.FormValue("id")
	title 	:= r.FormValue("title")
	author 	:= r.FormValue("author")
	year 	:= r.FormValue("year")

	book := dbm.Book{}

	query := fmt.Sprintf("update %s set title=?, author=?, relise_year=? where id=?", book.TableName())
	res, err := db.Instance().Exec(db.Instance().Rebind(query), id, title, author, year )
	if err != nil {
		return
	}

	num, err := res.RowsAffected() // Last id
	if err != nil {
		return
	}
	log.Debug("RowsAffected:%d", num)

	// return data to client
	getBook(w, r) // read book record
	return
}

func removeBook(w http.ResponseWriter, r *http.Request) {
	var err error
	// id any error return http.StatusBadRequest(400) else http.StatusOK(200)
	defer func() {
		if err == nil {
			w.WriteHeader(http.StatusOK)
		} else {
			log.Error(err)
			//w.Write([]byte(err.Error()))
			http.Error(w, err.Error(), http.StatusBadRequest)
		}
	}()
	defer log.AppRecover(getBook, &err)	// перехватчик аварийного завершения

	// read request parameters
	id := r.FormValue("id")

	book := dbm.Book{}

	query := fmt.Sprintf("delete from %s where id=?", book.TableName())

	res, err := db.Instance().Exec(db.Instance().Rebind(query), id)
	if err != nil {
		return
	}

	num, err := res.RowsAffected() // Last id
	if err != nil {
		return
	}
	log.Debug("RowsAffected:%d", num)

	if num == 0 {
		err = fmt.Errorf("book with id:%d not found", id)
	}
	return
}
