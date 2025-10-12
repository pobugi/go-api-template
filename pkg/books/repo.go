package books

import (
	"context"

	"gorm.io/gorm"
)

type Repo struct{ db *gorm.DB }

func NewRepo(db *gorm.DB) *Repo { return &Repo{db: db} }

func (r *Repo) Create(ctx context.Context, b *Book) error {
	return r.db.WithContext(ctx).Create(b).Error
}

func (r *Repo) GetAll(ctx context.Context) ([]Book, error) {
	var out []Book
	err := r.db.WithContext(ctx).Order("id").Find(&out).Error
	return out, err
}

func (r *Repo) GetByID(ctx context.Context, id uint) (*Book, error) {
	var b Book
	if err := r.db.WithContext(ctx).First(&b, id).Error; err != nil {
		return nil, err
	}
	return &b, nil
}

func (r *Repo) Save(ctx context.Context, b *Book) error {
	return r.db.WithContext(ctx).Save(b).Error
}

func (r *Repo) Delete(ctx context.Context, id uint) error {
	return r.db.WithContext(ctx).Delete(&Book{}, id).Error
}
