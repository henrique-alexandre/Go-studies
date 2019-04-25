package main

import (
	"database/sql"
	"encoding/json"
	"github.com/gorilla/mux"
	"github.com/lib/pq"
	"github.com/subosito/gotenv"
	"log"
	"net/http"
	"os"
)

type Book struct {
	ID     int    `json:id`
	Title  string `json:title`
	Author string `json:author`
	Year   string `json:year`
}

var books []Book
var db *sql.DB

func init() {
	gotenv.Load()
}

func logPrint(err error) {
	if err != nil {
		log.Fatal(err)
	}
}

func main() {
	//function used to connect to the database
	pgURL, e := pq.ParseURL(os.Getenv("ELEPHANTSQL_URL"))
	logPrint(e)

	log.Println(pgURL)
	//function used to open the connection agains the database
	db, e = sql.Open("postgres", pgURL)
	logPrint(e)

	//testing the connection
	e = db.Ping()
	log.Println(e)

	router := mux.NewRouter()

	router.HandleFunc("/books", getBooks).Methods("GET")
	router.HandleFunc("/books/{id}", getBook).Methods("GET")
	router.HandleFunc("/books", addBook).Methods("POST")
	router.HandleFunc("/books", updateBook).Methods("PUT")
	router.HandleFunc("/books/{id}", removeBook).Methods("DELETE")

	error := http.ListenAndServe(":8000", router)
	logPrint(error)

}

func getBooks(w http.ResponseWriter, r *http.Request) {
	var b Book
	books = []Book{}

	rows, err := db.Query("select * from books")
	logPrint(err)

	defer rows.Close()

	for rows.Next() {
		err := rows.Scan(&b.ID, &b.Title, &b.Author, &b.Year)
		logPrint(err)

		books = append(books, b)
	}

	json.NewEncoder(w).Encode(books)

}

func getBook(w http.ResponseWriter, r *http.Request) {

	var b Book
	p := mux.Vars(r)

	rows := db.QueryRow("select * from books where id=$1", p["id"])

	err := rows.Scan(&b.ID, &b.Title, &b.Author, &b.Year)
	logPrint(err)

	json.NewEncoder(w).Encode(b)

}

func addBook(w http.ResponseWriter, r *http.Request) {

	var b Book
	var bid int

	json.NewDecoder(r.Body).Decode(&b)

	err := db.QueryRow("inserting books (title, author, year), values($1, $2, $3) RETURNING bid",
		b.Title, b.Author, b.Year).Scan(&bid)
	logPrint(err)
}

func updateBook(w http.ResponseWriter, r *http.Request) {

	var b Book

	json.NewDecoder(r.Body).Decode(&b)

	result, _ := db.Exec("update books set title=$1, author=$2, year=$3 where id=$4 RETURNING id",
		&b.Title, &b.Author, &b.Year, &b.ID)

	rowsUpdated, err := result.RowsAffected()
	logPrint(err)

	json.NewEncoder(w).Encode(rowsUpdated)

}
func removeBook(w http.ResponseWriter, r *http.Request) {
	p := mux.Vars(r)

	result, err := db.Exec("delete from books where id = $1", p["id"])
	logPrint(err)

	rowsDeleted, e := result.RowsAffected()
	logPrint(e)

	json.NewEncoder(w).Encode(rowsDeleted)
}
