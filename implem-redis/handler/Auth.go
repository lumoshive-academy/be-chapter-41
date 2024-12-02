package handler

import (
	"golang-chapter-41/implem-redis/database"
	"golang-chapter-41/implem-redis/helper"
	"golang-chapter-41/implem-redis/service"

	"github.com/gin-gonic/gin"
	"go.uber.org/zap"
)

type AuthHadler struct {
	Service service.AllService
	Log     *zap.Logger
	Cacher  database.Cacher
}

func NewUserHandler(service service.AllService, log *zap.Logger, rdb database.Cacher) AuthHadler {
	return AuthHadler{
		Service: service,
		Log:     log,
		Cacher:  rdb,
	}
}

func (auth *AuthHadler) Login(c *gin.Context) {

	// get user form database
	token := "2323232"
	IDKEY := "username-1"

	err := auth.Cacher.Set(IDKEY, token)
	if err != nil {
		helper.BadResponse(c, "server error", 500)
	}

}
