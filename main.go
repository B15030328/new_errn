package main

import (
	"error_test/errn"
	"fmt"
	"net/http"

	"git.dz11.com/vega/minerva/server/echo"
)

type data struct {
	name string
	age  int
}

func main() {
	err := errn.Error{
		Errs:    2001,
		Message: "333",
	}
	fmt.Println(err)
}

func JsonTest(c echo.Context) error {
	err := errn.BadRequest
	return c.JSON(http.StatusOK, errn.ConvertJson(err, "", ""))
}
