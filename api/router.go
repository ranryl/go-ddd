package api

import (
	"github.com/gin-gonic/gin"
	"github.com/ranryl/go-ddd/internal/service"
)

func NewRouter(cs *service.ClusterService) *gin.Engine {
	r := gin.New()
	apiV1 := r.Group("/api/v1")
	apiV1.GET("/clusters", cs.List)
	return r
}
