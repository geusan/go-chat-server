package rest

import "github.com/labstack/echo/v4"

func ParseBody[T interface{}](c echo.Context, body T) T {
	if err := c.Bind(&body); err != nil {
		panic(err)
	}
	return body
}
