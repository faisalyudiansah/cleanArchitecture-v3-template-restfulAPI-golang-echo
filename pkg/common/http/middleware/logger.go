package httpmiddleware

import (
	"github.com/labstack/echo/v4"
	"github.com/sirupsen/logrus"
)

// Logger is an echo's logger middleware with configs
func Logger(logger *logrus.Logger) echo.MiddlewareFunc {
	return func(next echo.HandlerFunc) echo.HandlerFunc {
		return func(c echo.Context) error {
			// Log the incoming request
			logger.WithFields(logrus.Fields{
				"id":         c.Get(echo.HeaderXRequestID),
				"method":     c.Request().Method,
				"path":       c.Request().URL.Path,
				"query":      c.Request().URL.Query(),
				"ip":         c.RealIP(),
				"user_agent": c.Request().UserAgent(),
			}).Info("Received a new request")

			// Continue to the next handler
			err := next(c)

			// Log the response
			status := c.Response().Status
			logFields := logrus.Fields{
				"id":     c.Get(echo.HeaderXRequestID),
				"status": status,
				"url":    c.Request().URL.String(),
			}

			if status >= 300 && status < 400 {
				logger.WithFields(logFields).Info("Request resulted in a redirection")
			} else if status >= 200 && status < 300 {
				logger.WithFields(logFields).Info("Request processed successfully")
			} else if status >= 400 && status < 500 {
				logger.WithFields(logFields).Error("Client error occurred during the request")
			} else if status >= 500 {
				logger.WithFields(logFields).Error("Server error occurred while processing the request")
			}

			return err
		}
	}
}
