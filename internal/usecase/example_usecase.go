package usecase

import (
	"server/internal/adapter/repository"
	"server/internal/adapter/validator"
	"server/internal/infrastructure/logger"
	httprequest "server/pkg/common/http/request"
	httpresponse "server/pkg/common/http/response"

	"github.com/pkg/errors"
)

type ExampleUsecaseInterface interface {
	Get(req *httprequest.ListRequest) (*httpresponse.PaginatedReponse[*httpresponse.ResponseData], error)
}

type ExampleUsecase struct {
	Logger            *logger.Logger
	Validator         *validator.CustomValidator
	Repository        *repository.Repository
	ExampleRepository repository.ExampleRepositoryInterface
}

func NewExampleUsecase(
	l *logger.Logger,
	v *validator.CustomValidator,
	r *repository.Repository,
	er repository.ExampleRepositoryInterface,
) *ExampleUsecase {
	return &ExampleUsecase{
		Logger:            l,
		Validator:         v,
		Repository:        r,
		ExampleRepository: er,
	}
}

func (tu *ExampleUsecase) logError(err error, msg string) {
	tu.Logger.Logger.WithError(err).Error(msg)
}

func (tu *ExampleUsecase) Get(req *httprequest.ListRequest) (*httpresponse.PaginatedReponse[*httpresponse.ResponseData], error) {
	if err := tu.Validator.Validate(req); err != nil {
		errMessage := "failed to validate list request"
		validationErrors := tu.Validator.ParseValidationErrors(err)
		tu.logError(validationErrors, errMessage)
		return nil, validationErrors
	}
	if err := req.DecodeFilters(); err != nil {
		errMessage := "failed to decode list filters"
		tu.logError(err, errMessage)
		return nil, err
	}
	if err := tu.ExampleRepository.Get(req); err != nil {
		errMessage := "failed to get data"
		tu.logError(err, errMessage)
		return nil, errors.Wrap(err, errMessage)
	}
	res := []*httpresponse.ResponseData{
		{
			Data: "hello world",
		},
	}
	paginated := &httpresponse.PaginatedReponse[*httpresponse.ResponseData]{
		Data:        res,
		Total:       1,
		CurrentPage: req.Page,
		PerPage:     req.PerPage,
	}

	return paginated, nil
}
