package implementations

import (
	"context"
	"inventory-api/internal/model"
	"inventory-api/internal/repository/interfaces"

	"gorm.io/gorm"
)

type categoryRepository struct {
	db *gorm.DB
}

func NewCategoryRepository(db *gorm.DB) interfaces.CategoryRepository  {
	return &categoryRepository{
		db: db,
	}
}

func (repository *categoryRepository) Save(
	ctx context.Context,
	tx *gorm.DB,
	category *model.Categories,
) error {
	return tx.WithContext(ctx).Create(category).Error
}

func (repository *categoryRepository) Update (
	ctx context.Context,
	tx *gorm.DB,
	category *model.Categories,
) error {
	return tx.WithContext(ctx).Save(category).Error
}

func (repository *categoryRepository) Delete (
	ctx context.Context,
	tx *gorm.DB,
	category *model.Categories,
) error {
	return tx.WithContext(ctx).Delete(category).Error
}

func (repository *categoryRepository) FindByID(
	ctx context.Context,
	tx *gorm.DB,
	categoryID uint64,
)(*model.Categories, error){
	var category model.Categories

	err := tx.WithContext(ctx).First(&category, categoryID).Error
	if err != nil {
		return nil, err
	}
	return &category, nil
}

func (repository *categoryRepository) FindAll(
	ctx context.Context,
	tx *gorm.DB,
)([]*model.Categories, error){
	var category []*model.Categories

	err := tx.WithContext(ctx).Find(&category).Error
	if err != nil {
		return nil, err
	}

	return category, nil
}