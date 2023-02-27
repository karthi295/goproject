package handler

import (
	"encoding/json"
	"fmt"
	"karthikeyan/books/model"
	"log"
	"net/http"

	"github.com/gomodule/redigo/redis"
)

func (h handler) GetAllBooks(w http.ResponseWriter, r *http.Request) {

	var books []model.Book

	res, _ := redis.StringMap(h.Rc.Do("HGETALL", "podcast1"))

	if res != nil {
		for k, v := range res {
			fmt.Println("key: ", k)
			var tag model.Book
			_ = json.Unmarshal([]byte(v), &tag)
			books = append(books, tag)
			fmt.Println("value: ", v)
		}
		fmt.Println("value: ", books)
	}

	if books == nil {
		fmt.Println("inside if from DB")
		if result := h.DB.Find(&books); result.Error != nil {
			log.Fatal(result.Error)
		}

		for _, book := range books {
			mar, _ := json.Marshal(book)
			str := string(mar)
			_, err := h.Rc.Do("HMSET", "podcast1", book.Id, str)
			if err != nil {
				log.Fatal(err)
			}
		}
		fmt.Println("inside if from Db")
		w.Header().Add("Content-Type", "application/json")
		w.WriteHeader(http.StatusOK)
		json.NewEncoder(w).Encode(books)

	} else {
		fmt.Println("inside else from redis/cache")
		w.Header().Add("Content-Type", "application/json")
		w.WriteHeader(http.StatusOK)
		json.NewEncoder(w).Encode(books)
	}

}
