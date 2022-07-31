package service

import (
	"errors"
	"fmt"

	"github.com/duyanghao/sample_apiserver/pkg/models"
)

var clusterFakeList []models.Cluster

type ClusterServce interface {
	CreateCluster(cluster models.Cluster) ([]models.Cluster, error)
	DeleteCluster(id int) ([]models.Cluster, error)
	UpdateCluster(id int, cluster models.Cluster) ([]models.Cluster, error)
	GetCluster(id int) (models.Cluster, error)
	ListClusters() ([]models.Cluster, error)
}

type clusterService struct {
}

func NewClusterService() ClusterServce {
	return &clusterService{}
}

func (dm *clusterService) CreateCluster(cluster models.Cluster) (table []models.Cluster, err error) {
	for _, c := range clusterFakeList {
		if c.ID == cluster.ID {
			return nil, errors.New("Error: ID exists.")
		}
	}
	clusterFakeList = append(clusterFakeList, cluster)
	return clusterFakeList, nil

}

func (dm *clusterService) DeleteCluster(id int) (table []models.Cluster, err error) {
	for index, c := range clusterFakeList {
		if c.ID == id {
			clusterFakeList = append(clusterFakeList[:index], clusterFakeList[index+1:]...)
			return clusterFakeList, nil
		}
	}
	return nil, errors.New("Error: ID doesn't exixt.")
}

func (dm *clusterService) UpdateCluster(id int, cluster models.Cluster) (table []models.Cluster, err error) {
	for index, c := range clusterFakeList {
		if c.ID == id {
			clusterFakeList[index] = cluster
			return clusterFakeList, nil
		}
	}
	return nil, errors.New("Error: ID doesn't exit.")
}

func (dm *clusterService) GetCluster(id int) (cluster models.Cluster, err error) {
	for index, c := range clusterFakeList {
		if c.ID == id {
			return clusterFakeList[index], nil
		}
	}
	return models.Cluster{}, errors.New("Error: ID doesn't exist.")
}

func (dm *clusterService) ListClusters() (table []models.Cluster, err error) {
	if len(clusterFakeList) == 0 {
		return nil, errors.New("Error: Clusters's tabel is empty.")
	}
	return clusterFakeList, nil
}

func init() {
	fmt.Println("Init Clusterstable.")
	str := fmt.Sprintf("tableClusters:\n%v", clusterFakeList)
	fmt.Println(str)
}
