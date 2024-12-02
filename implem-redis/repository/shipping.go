package repository

import (
	"encoding/json"
	"errors"
	"fmt"
	"golang-chapter-41/implem-redis/helper"
	"golang-chapter-41/implem-redis/model"
	"net/http"

	"go.uber.org/zap"
	"gorm.io/gorm"
)

type ShippingRepoInterface interface {
	Create(customer *model.Shipping) error
	GetAll() (*[]model.Shipping, error)
	GetDestination(customer model.RequestDestination) (*float64, error)
	GetByID(id int) (*model.Shipping, error)
}

type ShippingRepository struct {
	DB     *gorm.DB
	Logger *zap.Logger
}

func NewShippingRepository(db *gorm.DB, log *zap.Logger) ShippingRepoInterface {
	return &ShippingRepository{
		DB:     db,
		Logger: log,
	}
}

func (shippingRepo *ShippingRepository) Create(shipping *model.Shipping) error {
	return nil
}

func (shippingRepo *ShippingRepository) GetAll() (*[]model.Shipping, error) {
	var shippings []model.Shipping

	// Menjalankan query untuk mengambil semua data dari tabel shipping
	if err := shippingRepo.DB.Find(&shippings).Error; err != nil {
		return nil, err
	}

	return &shippings, nil
}

func (shippingRepo *ShippingRepository) GetByID(id int) (*model.Shipping, error) {
	var shipping model.Shipping
	query := "SELECT id, name, price FROM shippings WHERE id= ?"

	err := shippingRepo.DB.Raw(query, id).Scan(&shipping).Error
	if err != nil {
		return nil, errors.New("error get data by id")
	}

	return &shipping, nil
}

func (shippingRepo *ShippingRepository) GetDestination(destination model.RequestDestination) (*float64, error) {

	url := fmt.Sprintf("https://router.project-osrm.org/route/v1/driving/%s;%s?overview=false",
		destination.OriginLongLat, destination.DestinationLongLat)
	var header http.Header
	data, err := helper.HTTPRequest("GET", header, url, nil)
	if err != nil {
		return nil, errors.New("error http request direction")
	}

	var dataMap map[string]interface{}

	err = json.Unmarshal(data, &dataMap)
	if err != nil {
		return nil, errors.New("error decode data")
	}

	routes, ok := dataMap["routes"].([]interface{})
	if !ok {
		return nil, errors.New("error decode routes")
	}

	// Periksa jika routes ada dan tidak kosong
	if len(routes) == 0 {
		return nil, errors.New("routes array is empty")
	}

	// Ambil elemen pertama dari slice routes
	route, ok := routes[0].(map[string]interface{})
	if !ok {
		return nil, errors.New("error decoding route")
	}

	distance, ok := route["distance"].(float64)
	if !ok {
		return nil, errors.New("error decode distance")
	}

	return &distance, nil
}
