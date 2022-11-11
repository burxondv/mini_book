package api

import (
	"book/storage"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
)

func (h *handler) GetBook(ctx *gin.Context) {
	id, err := strconv.ParseInt(ctx.Param("id"), 10, 64)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{
			"message": "failed to parse param",
		})
		return
	}

	blog, err := h.storage.GetBook(int(id))
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{
			"message": err.Error(),
		})
		return
	}

	ctx.JSON(http.StatusOK, blog)
}

func (h *handler) CreateBook(ctx *gin.Context) {
	var b storage.Book

	err := ctx.ShouldBindJSON(&b)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{
			"message": "failed bind request body",
		})
		return
	}

	blog, err := h.storage.CreateBook(&b)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{
			"message": "failed to create book",
		})
		return
	}

	ctx.JSON(http.StatusOK, blog)
}
