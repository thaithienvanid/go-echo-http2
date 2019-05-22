package main

import (
	"net/http"
	"github.com/labstack/echo/v4"
)

type Response struct {
	Code string `json:"code"`
	Message string `json:"message"`
	Payload interface{} `json:"payload"` 
}

func main() {
	e := echo.New()
	
	res := Response{Code: "0", Message: "OK", Payload: "Hello"}

	e.GET("/", func(c echo.Context) error {
		return c.JSON(http.StatusOK, res)
	})

	e.Logger.Fatal(e.StartTLS(":8080", "./certs/localhost.cert", "./certs/localhost.key"))
}
