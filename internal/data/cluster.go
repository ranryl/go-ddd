package data

import (
	"time"
)

type ClusterRepo struct {
	Data
}

type Cluster struct {
	ID              int64
	Name            string
	Mode            string
	Provider        string   `json:"provider,omitempty"`
	Region          string   `json:"region,omitempty"`
	Zone            string   `json:"zone,omitempty"`
	Zones           []string `json:"zones,omitempty"`
	CreateTimeStamp time.Time
}

func NewClusterRepo(data *Data) *ClusterRepo {
	return &ClusterRepo{
		Data: *data,
	}
}

func (c *ClusterRepo) List() ([]Cluster, error) {
	clusters := make([]Cluster, 0)
	tx := c.db.Find(clusters, []int{1, 2, 3})
	if tx.Error != nil {
		return nil, tx.Error
	}
	return clusters, nil
}
