package usecase

import (
	"fmt"
	"net/http"
	"server/internal/adapter/repository"
	"server/internal/infrastructure/logger"
	httpresponse "server/pkg/common/http/response"
)

type HealthCheckUseCase struct {
	Logger                 *logger.Logger
	CollectionRepositories map[string]RepositoryInterface
}

type RepositoryInterface interface {
	Ping() error
}

func NewHealthCheckUseCase(
	l *logger.Logger,
	ExampleRepository *repository.ExampleRepository,
) *HealthCheckUseCase {
	return &HealthCheckUseCase{
		Logger: l,
		CollectionRepositories: map[string]RepositoryInterface{
			"Example Repository": ExampleRepository,
		},
	}
}

func (u *HealthCheckUseCase) Ping() *httpresponse.HealthCheckResponse {
	response := &httpresponse.HealthCheckResponse{
		Code:    http.StatusOK,
		Message: "OK",
	}
	for name, repo := range u.CollectionRepositories {
		if err := repo.Ping(); err != nil {
			response.Code = http.StatusServiceUnavailable
			response.Message = fmt.Sprintf("%s health check failed: %s", name, err.Error())
			u.Logger.Logger.Error(response.Message)
			return response
		}
		u.Logger.Logger.Info(fmt.Sprintf("%v has successfully passed the health check", name))
	}
	return response
}
