package httproute

import (
	"server/internal/adapter/http/controller"

	"github.com/labstack/echo/v4"
)

type Route struct {
	App                   *echo.Echo
	HealthCheckController *controller.HealthCheckController
	LogController         *controller.LogController
	ExampleController     *controller.ExampleController
}

func (r *Route) Setup() {
	r.App.GET("/health", r.HealthCheckController.Ping)
	r.App.GET("/logs/:year/:month/:day/:f", r.LogController.Read)

	r.yourRoute()
}

func (r *Route) yourRoute() {
	r.App.GET("/", r.ExampleController.Get)
}
