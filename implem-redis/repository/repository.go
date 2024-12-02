package repository

import (
	"go.uber.org/zap"
	"gorm.io/gorm"
)

type AllRepository struct {
	ShippingRepo ShippingRepoInterface
}

func NewAllRepository(db *gorm.DB, log *zap.Logger) AllRepository {
	return AllRepository{
		ShippingRepo: NewShippingRepository(db, log),
	}
}
