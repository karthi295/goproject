package handler

import (
	"encoding/json"
	"io/ioutil"
	"karthikeyan/books/model"
	"log"
	"net/http"
	"strconv"

	"github.com/gorilla/mux"
)

func (h handler) UpdateBook(w http.ResponseWriter, r *http.Request) {
	defer r.Body.Close()
	body, err := ioutil.ReadAll(r.Body)
	if err != nil {
		log.Fatal(err)

	}
	var UpdateBook model.Book
	json.Unmarshal(body, &UpdateBook)

	vars := mux.Vars(r)
	id, _ := strconv.Atoi(vars["id"])

	var book model.Book
	if result := h.DB.First(&book, id); result.Error != nil {
		log.Fatal(result.Error)
	}

	book.Author = UpdateBook.Author
	book.Desc = UpdateBook.Desc
	book.Title = UpdateBook.Title

	h.DB.Save(&book).Order(id)

	// for index, book := range mocks.Books {
	// 	if book.Id == id {
	// 		book.Author = UpdateBook.Author
	// 		book.Desc = UpdateBook.Desc
	// 		book.Title = UpdateBook.Title

	// 	}
	// 	mocks.Books[index] = book
	w.Header().Add("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode("updated")

}
