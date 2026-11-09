package interfaces

import (
	"context"
	"inventory-api/internal/model"

	"gorm.io/gorm"
)

type UserRepository interface {
	Save(ctx context.Context, tx *gorm.DB, user *model.Users) error
	Update(ctx context.Context, tx *gorm.DB, user *model.Users) error
	FindById(ctx context.Context, tx *gorm.DB, userId uint64) (*model.Users, error)
	FindByEmail(ctx context.Context, tx *gorm.DB, email string) (*model.Users, error)
}
