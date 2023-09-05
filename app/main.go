package main

import (
	"os"

	"github.com/gin-gonic/gin"
	"github.com/gin-gonic/gin/binding"
	"github.com/yuuLab/go-validation/presentation/handler"
	"github.com/yuuLab/go-validation/presentation/validation"
	"golang.org/x/exp/slog"
)

func main() {
	r := gin.Default()
	// set validator
	binding.Validator = validation.NewOzzoValidator()
	// set logger
	slog.SetDefault(slog.New(slog.NewJSONHandler(os.Stdout, nil)))

	r.POST("/books", handler.NewBookHandler().Create)
	r.Run()
}
