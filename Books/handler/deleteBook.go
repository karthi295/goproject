package handler

import (
	"encoding/json"
	"karthikeyan/books/model"
	"log"
	"net/http"

	"github.com/gorilla/mux"
)

func (h handler) DeleteBook(w http.ResponseWriter, r *http.Request) {

	vars := mux.Vars(r)
	id := vars["id"]

	var book model.Book
	if result := h.DB.First(&book, id); result.Error != nil {
		log.Fatal(result.Error)
	}

	h.DB.Delete(&book)

	w.Header().Add("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)

	json.NewEncoder(w).Encode(book)

}
