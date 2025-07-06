package repository

import (
	"context"
	"database/sql"
	"hermawansafrin/belajar-golang-restful-api/model/domain"
)

// bikin contract nya dulu
type CategoryRepository interface {
	// utamakan menggunakan context dulu
	// kemudian ini tx adalah penggunaan transactional query

	Save(ctx context.Context, tx *sql.Tx, category domain.Category) domain.Category
	Update(ctx context.Context, tx *sql.Tx, category domain.Category) domain.Category
	Delete(ctx context.Context, tx *sql.Tx, category domain.Category)
	FindById(ctx context.Context, tx *sql.Tx, categoryId int) (domain.Category, error)
	FindAll(ctx context.Context, tx *sql.Tx) []domain.Category // list slices category
}
