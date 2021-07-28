package routing

import (
	"net/http"

	"github.com/dgrijalva/jwt-go"
	home "github.com/golangast/jwttoken/routing/home"
	login "github.com/golangast/jwttoken/routing/process/login"

	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
)

func Routing(e *echo.Echo) {
	e.GET("/", home.Home)
	e.GET("/login", login.Login)

	r := e.Group("/restricted")
	config := middleware.JWTConfig{
		Claims:     &jwtCustomClaims{},
		SigningKey: []byte("secret"),
	}
	r.Use(middleware.JWTWithConfig(config))
	r.GET("/t", restricted)

}

type jwtCustomClaims struct {
	Name  string `json:"name"`
	Admin bool   `json:"admin"`
	jwt.StandardClaims
}

func restricted(c echo.Context) error {
	user := c.Get("user").(*jwt.Token)
	claims := user.Claims.(*jwtCustomClaims)
	name := claims.Name
	return c.String(http.StatusOK, "Welcome "+name+"!")
}
