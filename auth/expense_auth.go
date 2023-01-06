package auth

import (
	"net/http"

	"github.com/labstack/echo"
)

type ResAuth struct {
	HTTPStatus int    `json:"HTTPStatus"`
	Message    string `json:"Message"`
}

func Authentication(next echo.HandlerFunc) echo.HandlerFunc {
	return func(c echo.Context) error {
		c.Response().Header().Set(echo.HeaderContentType, echo.MIMEApplicationJSONCharsetUTF8)

		token := c.Request().Header.Get("Authorization")
		if token == "" || token != "November 10, 2009wrong_token" {
			return c.JSON(http.StatusUnauthorized, ResAuth{Message: "Unauthorization token is invalid or expired.", HTTPStatus: http.StatusUnauthorized})
		}

		// username, password, ok := c.Request().BasicAuth()
		// if !ok {
		// 	return c.JSON(http.StatusUnauthorized, ResAuth{Message: "Can't parse the basic auth", HTTPStatus: http.StatusUnauthorized})
		// }

		// if username != "apidesign" || password != "45678" {
		// 	return c.JSON(http.StatusUnauthorized, ResAuth{Message: "Username or Password is invalid", HTTPStatus: http.StatusUnauthorized})
		// }

		return next(c)
	}
}