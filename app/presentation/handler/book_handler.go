package handler

import (
	"fmt"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/yuuLab/go-validation/presentation/request"
	"github.com/yuuLab/go-validation/presentation/validation"
	"golang.org/x/exp/slog"
)

type bookHandler struct {
	validator *validation.Validator
}

// NewBookHandler creates a new BookHandler.
func NewBookHandler(validator *validation.Validator) *bookHandler {
	return &bookHandler{validator: validator}
}

// Create creates a new book from request information.
func (h *bookHandler) Create(c *gin.Context) {
	var req request.Book
	if err := c.ShouldBindJSON(&req); err != nil {
		c.Status(http.StatusBadRequest)
		return
	}

	if err := h.validator.Validate(req); err != nil {
		if verr, ok := err.(validation.ValidationError); ok {
			c.JSON(verr.Status(), verr.Response())
			return
		}
		// unexpected error occurred...
		slog.Warn(fmt.Sprintf("an unexpected error occurred during validation: %v", err))
		c.Status(http.StatusInternalServerError)
		return
	}

	//TODO: create a new book...

	c.Status(http.StatusCreated)
}
