package service

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/ranryl/go-ddd/internal/biz"
)

type ClusterService struct {
	cluster *biz.ClusterCase
}

func NewClusterService(cluster *biz.ClusterCase) *ClusterService {
	return &ClusterService{
		cluster: cluster,
	}
}
func (cs *ClusterService) List(c *gin.Context) {
	c.JSON(http.StatusOK, cs.cluster.List())
}
