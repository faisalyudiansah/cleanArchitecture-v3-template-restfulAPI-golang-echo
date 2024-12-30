package repository

import (
	"fmt"
	"server/internal/infrastructure/database"
	"server/internal/infrastructure/logger"
	httprequest "server/pkg/common/http/request"

	"gorm.io/gorm"
)

type ExampleRepositoryInterface interface {
	Ping() error
	Get(request *httprequest.ListRequest) error
	// List(request *httprequest.ListRequest, result *[]entity.Example) error
}

type ExampleRepository struct {
	Repository
	Logger *logger.Logger
}

func NewExampleRepository(l *logger.Logger, db *database.Kind[*gorm.DB]) *ExampleRepository {
	repository := &ExampleRepository{
		Logger: l,
	}
	repository.DB = db
	return repository
}

// GetDB returns *database.Kind
func (r *ExampleRepository) GetDB() *database.Kind[*gorm.DB] {
	return r.DB
}

// Ping Check repository health
func (r *ExampleRepository) Ping() error {
	readDB, err := r.DB.Read.DB()
	if err != nil {
		return fmt.Errorf("failed to get read database instance: %w", err)
	}
	if err := readDB.Ping(); err != nil {
		return fmt.Errorf("read database is not available: %w", err)
	}
	return nil
}

func (cr *ExampleRepository) logError(err error, msg string) {
	cr.Logger.Logger.WithError(err).Error(msg)
}

func (r *ExampleRepository) applyFilterScope(request httprequest.FilteredRequestInterface) func(db *gorm.DB) *gorm.DB {
	return func(trx *gorm.DB) *gorm.DB {
		if request.GetFilters() == nil {
			return trx
		}
		for _, filter := range *request.GetFilters() {
			operator := r.ParseFilterOperator(filter.Operator)
			value := filter.Value

			if operator == "like" {
				value = fmt.Sprintf("%%%v%%", value)
			}
			trx = trx.Where(fmt.Sprintf("%s %s ?", filter.Field, operator), value)
		}
		return trx
	}
}

func (cr *ExampleRepository) Get(request *httprequest.ListRequest) error {
	return nil
}

// func (cr *ExampleRepository) List(request *httprequest.ListRequest, result *[]entity.Example) error {
// 	trx := cr.getTransaction(cr.DB.Read)
// 	defer trx.Rollback()
// 	if err := trx.Scopes(
// 		Paginate(request.Page, request.PerPage),
// 		cr.applyFilterScope(request),
// 	).Find(result).Error; err != nil {
// 		cr.logError(err, "failed while querying list")
// 		return err
// 	}
// 	if err := trx.Commit().Error; err != nil {
// 		cr.logError(err, "failed while committing")
// 		return err
// 	}
// 	return nil
// }
