//go:build wireinject
// +build wireinject

package api

import (
	"github.com/gin-gonic/gin"
	"github.com/google/wire"
	"github.com/ranryl/go-ddd/internal/biz"
	"github.com/ranryl/go-ddd/internal/data"
	"github.com/ranryl/go-ddd/internal/service"
	"gorm.io/gorm"
)

func Initialize(dsn data.DBDSN, config *gorm.Config) *gin.Engine {
	wire.Build(data.ProviderSet, biz.ProviderSet, service.ProviderSet, NewRouter)
	return nil
}
