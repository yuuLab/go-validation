package handler

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/yuuLab/go-validation/presentation/request"
	"github.com/yuuLab/go-validation/presentation/validation"
	"golang.org/x/exp/slog"
)

type bookHandler struct {
}

// NewBookHandler creates a new BookHandler.
func NewBookHandler() *bookHandler {
	return &bookHandler{}
}

// Create creates a new book from request information.
func (h *bookHandler) Create(c *gin.Context) {
	var req request.Book
	if err := c.ShouldBindJSON(&req); err != nil {
		if verr, ok := err.(validation.ValidationError); ok {
			slog.Info(verr.Error())
			c.JSON(verr.Status(), verr.Response())
			return
		}
		// json decord errors, etc.
		c.Status(http.StatusBadRequest)
		return
	}

	//TODO: create a new book...

	c.Status(http.StatusCreated)
}
