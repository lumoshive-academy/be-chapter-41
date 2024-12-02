package handler

import (
	"golang-chapter-41/implem-redis/helper"
	"golang-chapter-41/implem-redis/model"
	"golang-chapter-41/implem-redis/service"
	"net/http"

	"github.com/gin-gonic/gin"
	"go.uber.org/zap"
)

type ShippingHadler struct {
	Service service.AllService
	Log     *zap.Logger
}

func NewShippingHandler(service service.AllService, log *zap.Logger) ShippingHadler {
	return ShippingHadler{
		Service: service,
		Log:     log,
	}
}

func (shippingHadler *ShippingHadler) Create(c *gin.Context) {

}

func (shippingHadler *ShippingHadler) GetAllShipping(c *gin.Context) {
	data, err := shippingHadler.Service.CustomerService.GetAll()
	if err != nil {
		shippingHadler.Log.Error("get-all-shipping", zap.Error(err))
		helper.BadResponse(c, err.Error(), http.StatusBadRequest)
	}

	helper.SuccessResponseWithData(c, "success get data", http.StatusOK, *data)
}

func (shippingHadler *ShippingHadler) ShippingCost(c *gin.Context) {
	var data model.RequestDestination
	if err := c.ShouldBindQuery(&data); err != nil {
		helper.BadResponse(c, err.Error(), http.StatusBadRequest)
		return
	}

	cost, err := shippingHadler.Service.CustomerService.ShippingCost(data)
	if err != nil {
		helper.BadResponse(c, err.Error(), http.StatusBadRequest)
	}

	result := gin.H{
		"cost": cost,
	}

	helper.SuccessResponseWithData(c, "success", http.StatusOK, result)
}
