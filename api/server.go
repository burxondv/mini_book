package api

import (
	"book/storage"

	"github.com/gin-gonic/gin"
)

type handler struct {
	storage *storage.DBManager
}

func NewServer(storage *storage.DBManager) *gin.Engine {
	r := gin.Default()

	h := handler{
		storage: storage,
	}

	r.GET("/book/:id", h.GetBook)
	r.POST("/book", h.CreateBook)

	return r
}
