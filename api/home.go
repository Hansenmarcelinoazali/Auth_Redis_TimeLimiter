package api

import (
	"net/http"

	"github.com/labstack/echo"
)

func Home(c echo.Context) error {

	return c.String(http.StatusOK, "Hello Welcome to My Page")

}
