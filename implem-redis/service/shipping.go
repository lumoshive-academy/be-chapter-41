package service

import (
	"errors"
	"fmt"
	"golang-chapter-41/implem-redis/model"
	"golang-chapter-41/implem-redis/repository"

	"go.uber.org/zap"
)

type ShippingServiceInterface interface {
	Create(customer *model.Shipping) error
	GetAll() (*[]model.Shipping, error)
	ShippingCost(payload model.RequestDestination) (*float64, error)
}

type ShippingService struct {
	Repo repository.AllRepository
	Log  *zap.Logger
}

func NewShippingService(repo repository.AllRepository, log *zap.Logger) ShippingServiceInterface {
	return &ShippingService{
		Repo: repo,
		Log:  log,
	}
}

func (shippingService *ShippingService) Create(customer *model.Shipping) error {
	return nil
}

func (shippingService *ShippingService) GetAll() (*[]model.Shipping, error) {
	return shippingService.Repo.ShippingRepo.GetAll()
}

func (shippingService *ShippingService) ShippingCost(payload model.RequestDestination) (*float64, error) {

	// Ambil data shipping berdasarkan ID
	data, err := shippingService.Repo.ShippingRepo.GetByID(payload.ShippingID)
	if err != nil {
		return nil, fmt.Errorf("failed to get shipping data: %w", err)
	}

	// Ambil jarak ke tujuan
	distance, err := shippingService.Repo.ShippingRepo.GetDestination(payload)
	if err != nil {
		return nil, fmt.Errorf("failed to get destination distance: %w", err)
	}
	if distance == nil {
		return nil, errors.New("distance is nil")
	}

	// Perhitungan biaya
	priceOfShipping := data.Price
	priceByQty := priceOfShipping * float64(payload.Qty)
	priceOfDistance := priceByQty * *distance

	return &priceOfDistance, nil
}
