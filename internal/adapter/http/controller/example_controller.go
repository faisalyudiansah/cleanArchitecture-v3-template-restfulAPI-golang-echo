package controller

import (
	"net/http"
	"server/internal/infrastructure/logger"
	"server/internal/usecase"
	"server/pkg/apperror"
	httprequest "server/pkg/common/http/request"
	"server/pkg/constant"

	"github.com/go-playground/validator/v10"
	"github.com/labstack/echo/v4"
)

type ExampleController struct {
	Logger         *logger.Logger
	ExampleUsecase usecase.ExampleUsecaseInterface
}

func NewExampleController(
	logger *logger.Logger,
	ExampleUsecase usecase.ExampleUsecaseInterface,
) *ExampleController {
	return &ExampleController{
		Logger:         logger,
		ExampleUsecase: ExampleUsecase,
	}
}

func (c *ExampleController) logError(err error, msg string) {
	c.Logger.Logger.WithError(err).Error(msg)
}

func (c *ExampleController) Get(ctx echo.Context) error {
	request := &httprequest.ListRequest{}
	if err := ctx.Bind(request); err != nil {
		c.logError(err, "error binding")
		return err
	}
	res, err := c.ExampleUsecase.Get(request)
	if err != nil {
		if errs, ok := err.(validator.ValidationErrors); ok {
			return echo.NewHTTPError(http.StatusUnprocessableEntity, errs).SetInternal(err)
		}
		if appErr, ok := err.(*apperror.AppError); ok {
			return ctx.JSON(appErr.Code, appErr)
		}
		return echo.NewHTTPError(http.StatusInternalServerError, constant.InternalServerErrorMessage).SetInternal(err)
	}
	return ctx.JSON(http.StatusOK, res)
}
