package main

import (
	"log"
    "net/http"

    "github.com/labstack/echo/v5"
)

func main() {
    e := echo.New()

    e.GET("/api/health", func(c *echo.Context) error {
        return c.JSON(http.StatusOK, map[string]string{
            "status": "ok",
        })
    })

    log.Fatal(e.Start(":8080"))
}