package httpmiddleware

import (
	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
)

var RequestID = middleware.RequestIDWithConfig(middleware.RequestIDConfig{
	RequestIDHandler: func(c echo.Context, id string) {
		c.Set(echo.HeaderXRequestID, id)
	},
})
