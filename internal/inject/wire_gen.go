// Code generated by Wire. DO NOT EDIT.

//go:generate go run github.com/google/wire/cmd/wire
//go:build !wireinject
// +build !wireinject

package inject

import (
	"context"
	"github.com/LyricTian/gin-admin/v9/internal/api"
	"github.com/LyricTian/gin-admin/v9/internal/dao/repo"
	"github.com/LyricTian/gin-admin/v9/internal/dao/util"
	"github.com/LyricTian/gin-admin/v9/internal/router"
	"github.com/LyricTian/gin-admin/v9/internal/service"
)

// Injectors from wire.go:

func BuildInjector(ctx context.Context) (*Injector, func(), error) {
	cacher, cleanup, err := InitCache(ctx)
	if err != nil {
		return nil, nil, err
	}
	auther, cleanup2, err := InitJWTAuth()
	if err != nil {
		cleanup()
		return nil, nil, err
	}
	enforcer, cleanup3, err := InitCasbin(ctx)
	if err != nil {
		cleanup2()
		cleanup()
		return nil, nil, err
	}
	db, cleanup4, err := InitGormDB(ctx)
	if err != nil {
		cleanup3()
		cleanup2()
		cleanup()
		return nil, nil, err
	}
	trans := &util.Trans{
		DB: db,
	}
	demoRepo := &repo.DemoRepo{
		DB: db,
	}
	demoSrv := &service.DemoSrv{
		TransRepo: trans,
		DemoRepo:  demoRepo,
	}
	demoAPI := &api.DemoAPI{
		DemoSrv: demoSrv,
	}
	routerRouter := &router.Router{
		Cache:          cacher,
		JWTAuth:        auther,
		CasbinEnforcer: enforcer,
		DemoAPI:        demoAPI,
	}
	engine := InitEngine(routerRouter)
	injector := &Injector{
		Engine: engine,
	}
	return injector, func() {
		cleanup4()
		cleanup3()
		cleanup2()
		cleanup()
	}, nil
}