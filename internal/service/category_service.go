package service

import (
	"context"
	"inventory-api/internal/model"
	"inventory-api/internal/repository/interfaces"

	"gorm.io/gorm"
)

type CategoryService struct {
	categoryRepository interfaces.CategoryRepository
	db                 *gorm.DB
}

func NewCategoryService(categoryRepository interfaces.CategoryRepository, db *gorm.DB) *CategoryService{
	return &CategoryService{
		categoryRepository: categoryRepository,
		db:db,
	}
}

func (service *CategoryService) SaveCategory(
	ctx context.Context,
	category *model.Categories,
) error {
	return service.db.WithContext(ctx).Transaction(func(tx *gorm.DB) error {
		return service.categoryRepository.Save(ctx, tx, category)
	})
}

func (service *CategoryService) UpdateCategory(
	ctx context.Context,
	category *model.Categories,
) error {
	return service.db.WithContext(ctx).Transaction(func(tx *gorm.DB) error {
		return service.categoryRepository.Update(ctx, tx, category)
	})
}

func (service *CategoryService) DeleteCategory(
	ctx context.Context,
	category *model.Categories,
) error {
	return service.db.WithContext(ctx).Transaction(func(tx *gorm.DB) error {
		return service.categoryRepository.Delete(ctx, tx, category)
	})
}

func (service *CategoryService) FindByIDCategory(
	ctx context.Context,
	categoryID uint64,
) (*model.Categories, error){
	category, err := service.categoryRepository.FindByID(
		ctx,
		service.db,
		categoryID,
	)

	if err != nil {
		return nil, err
	}
	
	return category, nil
}

func (service *CategoryService) FindAllCategory(
	ctx context.Context,
) ([]*model.Categories, error){
	category, err := service.categoryRepository.FindAll(
		ctx,
		service.db,
	)

	if err != nil {
		return nil, err		
	}

	return category, nil
}