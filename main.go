package main

import (
	"fmt"
	"log"
	"net/http"
	"time"

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
	date := time.Date(2025, time.January, 1, 0, 0, 0, 0, time.UTC)
	duration := time.Until(date)
	solution := fmt.Sprintf("Days until 2025: %d", int64(duration.Hours() / 24))

	err := ctx.String(http.StatusOK, solution)
	if err != nil {
		return err
	}

	return nil
}