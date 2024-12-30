package usecase

import (
	"net/http"
	"server/internal/adapter/repository"
	"server/internal/adapter/validator"
	"server/internal/infrastructure/logger"
	"server/pkg/apperror"
	httprequest "server/pkg/common/http/request"
	httpresponse "server/pkg/common/http/response"
	"server/pkg/constant"

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
		tu.logError(err, constant.ErrValidateRequest)
		return nil, err
	}
	if err := req.DecodeFilters(); err != nil {
		formatErr := errors.Wrap(err, constant.ErrDecodeFilterRequest)
		tu.logError(formatErr, constant.ErrDecodeFilterRequest)
		return nil, apperror.NewAppError(http.StatusUnprocessableEntity, formatErr.Error(), formatErr)
	}
	if err := tu.ExampleRepository.Get(req); err != nil {
		errMessage := "failed to get data"
		tu.logError(err, errMessage)
		return nil, apperror.NewAppError(http.StatusInternalServerError, err.Error(), err)
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
