package login

import (
	"net/http"

	jwts "github.com/golangast/jwttoken/jwts"
	"github.com/labstack/echo/v4"
)

func Login(c echo.Context) error {
	fname := c.FormValue("fname")
	lname := c.FormValue("lname")

	jwts.JwtLogin(c)

	return c.Render(http.StatusOK, "index.html", map[string]interface{}{
		"fname": fname,
		"lname": lname,
	})

}
