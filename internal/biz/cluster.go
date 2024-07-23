package biz

import (
	"time"

	"github.com/ranryl/go-ddd/internal/data"
	log "github.com/sirupsen/logrus"
)

type ClusterCase struct {
	clusterRepo ClusterRepo
}

func NewClusterCase(clusterRepo *data.ClusterRepo) *ClusterCase {
	return &ClusterCase{
		clusterRepo: clusterRepo,
	}
}

type Cluster struct {
	Name            string
	Mode            string
	Provider        string   `json:"provider,omitempty"`
	Region          string   `json:"region,omitempty"`
	Zone            string   `json:"zone,omitempty"`
	Zones           []string `json:"zones,omitempty"`
	CreateTimeStamp time.Time
}

func (c *ClusterCase) List() []Cluster {
	clusterList, err := c.clusterRepo.List()
	if err != nil {
		log.Error(err)
		return nil
	}
	result := make([]Cluster, 0)
	for _, value := range clusterList {
		var c Cluster
		c.Name = value.Name
		c.Mode = value.Mode
		c.Provider = value.Provider
		c.Region = value.Region
		c.Zone = value.Zone
		c.Zones = value.Zones
		c.CreateTimeStamp = value.CreateTimeStamp
		result = append(result, c)
	}
	return result
}

type ClusterRepo interface {
	List() ([]data.Cluster, error)
}
