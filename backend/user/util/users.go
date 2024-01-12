package util

import (
	"net/http"
	"time"

	"github.com/labstack/echo/v4"
)

func SetCookies(c echo.Context, kv map[string]string) {
	for k, v := range kv {
		cookie := new(http.Cookie)
		cookie.Name = k
		cookie.Value = v
		cookie.Expires = time.Now().Add(24 * time.Hour)
		cookie.Path = "/"
		c.SetCookie(cookie)
	}

}
