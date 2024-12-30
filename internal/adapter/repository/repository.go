package repository

import (
	"context"

	"server/internal/infrastructure/database"

	"gorm.io/gorm"
)

type Repository struct {
	DB *database.Kind[*gorm.DB]
}

func (r *Repository) getTransaction(db *gorm.DB) *gorm.DB {
	return db.WithContext(context.Background()).Begin()
}

func (r *Repository) GetTx(db *gorm.DB) *gorm.DB {
	return db.WithContext(context.Background()).Begin()
}

func NewRepository(db *database.Kind[*gorm.DB]) *Repository {
	return &Repository{
		DB: db,
	}
}

func Paginate(page int, size int) func(db *gorm.DB) *gorm.DB {
	return func(db *gorm.DB) *gorm.DB {

		if page <= 0 {
			page = 1
		}
		if size <= 0 {
			size = 10
		}

		offset := (page - 1) * size
		return db.Offset(offset).Limit(size)
	}
}

func (r *Repository) ParseFilterOperator(operator string) string {
	var result string
	switch operator {
	case "eq", "equal", "equals":
		result = "="
	case "not_eq", "not_equal", "not_equals", "=":
		result = "!="
	case "gte", "greater_than_or_equal", ">=":
		result = ">="
	case "lte", "less_than_or_equal", "<=":
		result = "<="
	case "gt", "greater_than", ">":
		result = ">"
	case "lt", "less_than", "<":
		result = "<"
	case "contains":
		result = "like"
	default:
		result = operator
	}

	return result
}
