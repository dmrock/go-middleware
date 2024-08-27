package main

import (
	"fmt"
	"log"
	"net/http"

	"github.com/labstack/echo/v4"
)

func main() {
	fmt.Println("Server is running")

	s := echo.New()

	s.GET("/status", Handler)

	err := s.Start(":8080")
	if err != nil {
		log.Fatal(err)
	}
}

func Handler(ctx echo.Context) error {
	err := ctx.String(http.StatusOK, "Test")
	if err != nil {
		return err
	}

	return nil
}