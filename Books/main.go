package main

import (
	"karthikeyan/books/db"
	"karthikeyan/books/handler"
	"log"
	"net/http"

	"github.com/gorilla/mux"
)

func main() {

	DB := db.Init()
	con := db.Connection()
	h := handler.New(DB, con)

	router := mux.NewRouter()
	router.HandleFunc("/books", h.GetAllBooks).Methods("GET")
	router.HandleFunc("/books", h.Addbook).Methods("POST")
	router.HandleFunc("/books/{id}", h.GetBook).Methods("GET")
	router.HandleFunc("/books/{id}", h.UpdateBook).Methods("PUT")
	router.HandleFunc("/books/{id}", h.DeleteBook).Methods("DELETE")
	log.Print("API running")
	log.Fatal(http.ListenAndServe(":8081", router))

	// con := db.Connection()

	// book := &model.Book{
	// 	Id:     "1",
	// 	Title:  "wings of fire",
	// 	Author: "APJ abdul kalam",
	// 	Desc:   "biography of APJ",
	// }
	// mar, _ := json.Marshal(book)

	// str := string(mar)

	// _, err := con.Do("HMSET",
	// 	"podcast",
	// 	"title",
	// 	"teach over",
	// 	"creator",
	// 	"Brodie",
	// 	"age",
	// 	1.1,
	// 	"book",
	// 	str,
	// )

	// if err != nil {
	// 	log.Fatal(err)
	// }

	// res, err := redis.String(con.Do("HGET", "podcast", "book"))

	// var raw *model.Book
	// json.Unmarshal([]byte(res), &raw)

	// if err != nil {
	// 	log.Fatal(err)
	// }
	// fmt.Println("podcast title: ", raw.Title)
	// fmt.Printf("%T", raw)
}
