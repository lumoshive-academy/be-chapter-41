package handler

import (
	"golang-chapter-41/implem-redis/service"

	"go.uber.org/zap"
)

type AllHandler struct {
	ShippingHandler ShippingHadler
}

func NewAllHandler(service service.AllService, log *zap.Logger) AllHandler {
	return AllHandler{
		ShippingHandler: NewShippingHandler(service, log),
	}
}
