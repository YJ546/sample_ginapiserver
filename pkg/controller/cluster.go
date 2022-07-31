package controller

import (
	"strconv"

	"github.com/duyanghao/sample_apiserver/pkg/models"
	"github.com/duyanghao/sample_apiserver/pkg/service"
	"github.com/gin-gonic/gin"
)

type ClusterController interface {
	GetCluster(*gin.Context)
	CreateCluster(*gin.Context)
	DeleteCluster(*gin.Context)
	UpdateCluster(*gin.Context)
	ListClusters(*gin.Context)
}

type clusterController struct {
	clusterService service.ClusterServce
}

func NewClusterController() ClusterController {
	return &clusterController{
		clusterService: service.NewClusterService(),
	}
}

// @Summary 查询集群
// @Description to get a cluster's message by id
// @Param ID path int true "集群id"
// @Success 200 {object} models.Cluster OK
// @Failure 400 {object} models.Cluster Bad Requests
// @Failure 403 {object} models.Cluster Forbidden
// @Failure 500 {object} models.Cluster Internal Server Error
// @router /api/v1/clusters/{id} [get]
func (cc *clusterController) GetCluster(c *gin.Context) {
	StrId := c.Params.ByName("id")
	id, _ := strconv.Atoi(StrId) //if arg is invalid(like float,string,) ,Atoi(arg) return 0

	//valid id must be positive integer
	if id <= 0 {
		c.JSON(400, models.Response{Code: -1, Message: "Invalid Id, Get Cluster Failed", Data: struct{}{}})
		return
	}
	cluster, err := cc.clusterService.GetCluster(id)
	if err != nil {
		c.JSON(404, models.Response{Code: -1, Message: "Error:" + err.Error(), Data: struct{}{}})
		return
	}
	c.JSON(200, models.Response{Code: 0, Message: "Get Cluster Sucessful", Data: cluster})
}

// @Summary 展示集群列表
// @Description List all of clusters' message
// @Success 200 {object} models.Cluster OK
// @Failure 400 {object} models.Cluster Bad Requests
// @Failure 403 {object} models.Cluster Forbidden
// @Failure 500 {object} models.Cluster Internal Server Error
// @router /api/v1/clusters [get]
func (cc *clusterController) ListClusters(c *gin.Context) {
	table, err := cc.clusterService.ListClusters()
	if err != nil {
		c.JSON(404, models.Response{Code: -1, Message: "Error:" + err.Error(), Data: struct{}{}})
		return
	}
	c.JSON(200, models.Response{Code: 0, Message: "List Clusters Sucessful", Data: table})
}

// @Summary 更新集群信息
// @Description to modify a cluster's message by id
// @Param ID path int true "集群id"
// @Success 200 {object} models.Cluster OK
// @Failure 400 {object} models.Cluster Bad Requests
// @Failure 403 {object} models.Cluster Forbidden
// @Failure 500 {object} models.Cluster Internal Server Error
// @router /api/v1/clusters/{id} [put]
func (cc *clusterController) UpdateCluster(c *gin.Context) {
	StrId := c.Params.ByName("id")
	id, _ := strconv.Atoi(StrId)
	//valid id must be positive integer
	if id <= 0 {
		c.JSON(400, models.Response{Code: -1, Message: "Invalid Id,Update Failed", Data: struct{}{}})
		return
	}

	var cluster models.Cluster
	c.BindJSON(&cluster)
	_, err := cc.clusterService.UpdateCluster(id, cluster)
	if err != nil {
		c.JSON(404, models.Response{Code: -1, Message: "Error:" + err.Error(), Data: struct{}{}})
		return
	}
	c.JSON(200, models.Response{Code: 0, Message: "Update Cluster Sucessful", Data: struct{}{}})

}

// @Summary 创建集群
// @Description to add a cluster's message
// @Success 200 {object} models.Cluster OK
// @Failure 400 {object} models.Cluster Bad Requests
// @Failure 403 {object} models.Cluster Forbidden
// @Failure 500 {object} models.Cluster Internal Server Error
// @router /api/v1/clusters [post]
func (cc *clusterController) CreateCluster(c *gin.Context) {
	var cluster models.Cluster
	c.BindJSON(&cluster)
	_, err := cc.clusterService.CreateCluster(cluster)
	if err != nil {
		c.JSON(404, models.Response{Code: -1, Message: "Error:" + err.Error(), Data: struct{}{}})
	}
	c.JSON(200, models.Response{Code: 0, Message: "Create Cluster Sucessful", Data: struct{}{}})
}

// @Summary 删除集群
// @Description to delete a cluster's message by id
// @Param ID path int true "集群id"
// @Success 200 {object} models.Cluster OK
// @Failure 400 {object} models.Cluster Bad Requests
// @Failure 403 {object} models.Cluster Forbidden
// @Failure 500 {object} models.Cluster Internal Server Error
// @router /api/v1/clusters/{id} [delete]
func (cc *clusterController) DeleteCluster(c *gin.Context) {
	StrId := c.Params.ByName("id")
	id, _ := strconv.Atoi(StrId)
	_, err := cc.clusterService.DeleteCluster(id)
	if err != nil {
		c.JSON(404, models.Response{Code: -1, Message: "Error:" + err.Error(), Data: struct{}{}})
	}
	c.JSON(200, models.Response{Code: 0, Message: "Delete Cluster Sucessful", Data: struct{}{}})
}
