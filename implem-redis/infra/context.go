package infra

import (
	"golang-chapter-41/implem-redis/database"
	"golang-chapter-41/implem-redis/handler"
	"golang-chapter-41/implem-redis/middleware"
	"golang-chapter-41/implem-redis/repository"
	"golang-chapter-41/implem-redis/service"
	"golang-chapter-41/implem-redis/util"

	"go.uber.org/zap"
)

type Context struct {
	Log        *zap.Logger
	Config     util.Configuration
	Handler    handler.AllHandler
	Cacher     database.Cacher
	Middleware middleware.Middleware
}

func NewContext() (Context, error) {
	logger, err := util.LoggerInit()
	if err != nil {
		return Context{}, err
	}

	config, err := util.ReadConfig()
	if err != nil {
		return Context{
			Log: logger,
		}, err
	}

	db, err := database.InitDB(config)
	if err != nil {
		return Context{
			Log: logger,
		}, err
	}

	rdb := database.NewCacher(config, 60*60)

	repo := repository.NewAllRepository(db, logger)
	service := service.NewAllService(repo, logger)
	handler := handler.NewAllHandler(service, logger)
	middleware := middleware.NewMiddleware(rdb)
	return Context{Log: logger, Config: config, Handler: handler, Cacher: rdb, Middleware: middleware}, nil
}
