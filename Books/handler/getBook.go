package handler

import (
	"encoding/json"
	"fmt"
	"karthikeyan/books/model"

	"log"
	"net/http"

	"github.com/gomodule/redigo/redis"
	"github.com/gorilla/mux"
)

// ID := string(id)
// t = &cache.RedisCache{}
// book := t.Get(id)
func (h handler) GetBook(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	id, _ := (vars["id"])

	var tag model.Book

	res, err := redis.String(h.Rc.Do("HGET", "podcast1", id))
	var raw *model.Book
	if res != "" {
		err = json.Unmarshal([]byte(res), &raw)
		if err != nil {
			log.Fatal(err)
		}
	}
	if raw == nil {
		fmt.Println("inside if from DB")

		if result := h.DB.First(&tag, id); result.Error != nil {
			log.Fatal(result.Error)
		}
		w.Header().Add("Content-Type", "application/json")
		w.WriteHeader(http.StatusOK)
		json.NewEncoder(w).Encode(tag)

		mar, _ := json.Marshal(tag)

		str := string(mar)

		_, err := h.Rc.Do("HMSET",
			"podcast1",
			tag.Id,
			str,
		)
		if err != nil {
			log.Fatal(err)
		}
		// _, err1 := h.Rc.Do("EXPIRE",
		// 	"podcast1",
		// 	10,
		// )
		// if err1 != nil {
		// 	log.Fatal(err)
		// }

	} else {
		_, err1 := h.Rc.Do("DEL",
			"podcast1",
		)
		if err1 != nil {
			log.Fatal(err)
		}
		fmt.Println("inside else from redis/cache")
		w.Header().Add("Content-Type", "application/json")
		w.WriteHeader(http.StatusOK)
		json.NewEncoder(w).Encode(raw)

	}

}
