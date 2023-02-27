package handler

import (
	"github.com/gomodule/redigo/redis"
	"gorm.io/gorm"
)

type handler struct {
	Rc redis.Conn
	DB *gorm.DB
}

func New(db *gorm.DB, con redis.Conn) handler {
	return handler{Rc: con, DB: db}
}
