package controller

import (
	"errors"
	"net/http"
	validatormodel "server/internal/adapter/validator/model"
	"server/internal/infrastructure/logger"
	"server/internal/usecase"
	httprequest "server/pkg/common/http/request"
	httpresponse "server/pkg/common/http/response"

	"github.com/labstack/echo/v4"
	"gorm.io/gorm"
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
		errResponse := &httpresponse.ErrorResponse{}
		c.logError(err, errResponse.Message)
		return errResponse.EchoJsonResponse(ctx)
	}
	res, err := c.ExampleUsecase.Get(request)
	if err != nil {
		errResponse := &httpresponse.ErrorResponse{}
		errResponse.Message = "something error"
		if errs, ok := err.(validatormodel.ValidationErrors); ok {
			errResponse.Error = errs.ToResponseErrors()
			errResponse.Code = http.StatusBadRequest
		} else if errors.Is(err, gorm.ErrRecordNotFound) {
			errResponse.Code = http.StatusNotFound
			errResponse.Message = "NOT FOUND"
			errResponse.Error = err.Error()
		} else {
			errResponse.Error = err.Error()
			errResponse.Code = http.StatusInternalServerError
		}
		c.Logger.Logger.WithError(err).Error(errResponse.Message)
		return errResponse.EchoJsonResponse(ctx)
	}
	return ctx.JSON(http.StatusOK, res)
}
