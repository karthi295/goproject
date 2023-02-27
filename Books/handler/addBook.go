package handler

import (
	"encoding/json"
	"io/ioutil"
	"karthikeyan/books/model"
	"log"
	"net/http"

	_ "github.com/gorilla/mux"
)

func (h handler) Addbook(w http.ResponseWriter, r *http.Request) {
	defer r.Body.Close()

	body, err := ioutil.ReadAll(r.Body)
	if err != nil {
		log.Fatal(err)
	}

	var book model.Book

	json.Unmarshal(body, &book)

	if result := h.DB.Create(&book); result.Error != nil {
		log.Fatal(result.Error)
	}
	//book.Id = rand.Intn(100)
	//mocks.Books = append(mocks.Books, book)

	w.Header().Add("Content-Type", "application/json")
	w.WriteHeader(http.StatusCreated)
	json.NewEncoder(w).Encode("Created")

}
