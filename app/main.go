package main

import (
	"os"

	"github.com/gin-gonic/gin"
	"github.com/yuuLab/go-validation/presentation/handler"
	"github.com/yuuLab/go-validation/presentation/validation"
	"golang.org/x/exp/slog"
)

func main() {
	r := gin.Default()
	// set logger
	logger := slog.New(slog.NewJSONHandler(os.Stdout, nil))
	slog.SetDefault(logger)

	v := validation.NewValidator()
	r.POST("/books", handler.NewBookHandler(v).Create)
	r.Run()
}
