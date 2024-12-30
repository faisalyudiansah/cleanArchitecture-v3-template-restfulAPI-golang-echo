package httpmiddleware

import (
	"net/http"

	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
	"github.com/sirupsen/logrus"
)

// Recover enable echo's ability to recover after system failures.
func Recovery(l *logrus.Logger) echo.MiddlewareFunc {
	return middleware.RecoverWithConfig(middleware.RecoverConfig{
		StackSize: 1 << 10, // 1 KB
		LogLevel:  0,       // Set log level (0 = Print)
		LogErrorFunc: func(c echo.Context, err error, stack []byte) error {
			// Log recovery
			l.WithFields(logrus.Fields{
				"id":    c.Get(echo.HeaderXRequestID),
				"trace": string(stack),
			}).WithError(err).Errorf("Recovering from error: %v", err)

			return c.JSON(http.StatusInternalServerError, echo.Map{
				"message": "Internal Server Error",
			})
		},
	})
}
