package home

import (
	"net/http"

	"github.com/labstack/echo/v4"
)

func Home(c echo.Context) error {
	userids := c.Param("userid")

	return c.Render(http.StatusOK, "form.html", map[string]interface{}{
		"userid": userids,
	})

}
