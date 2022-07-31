package route

import (
	"os"

	_ "github.com/duyanghao/sample_apiserver/docs"
	"github.com/duyanghao/sample_apiserver/pkg/controller"
	"github.com/duyanghao/sample_apiserver/pkg/log"
	"github.com/gin-gonic/gin"
	ginSwagger "github.com/swaggo/gin-swagger"
	"github.com/swaggo/gin-swagger/swaggerFiles"
)

// @title Swagger sample_apiserver
// @version 0.1.0
// @description This is a sample_apiserver.
// @contact.name gaoyujian
// @contact.url github.com/YJ546
// @contact.email 1194373959@qq.com
// @license.name Apache 2.0
// @license.url http://www.apache.org/licenses/LICENSE-2.0.html
// @BasePath /api/v1
func InstallRoutes(r *gin.Engine) {
	// Recovery middleware recovers from any panics and writes a 500 if there was one.
	r.Use(gin.Recovery())

	// docs.SwaggerInfo.BasePath = "/api/v1"
	r.GET("/swagger/*any", ginSwagger.WrapHandler(swaggerFiles.Handler))

	// a ping api test
	r.GET("/ping", controller.Ping)

	// get sample_apiserver version
	r.GET("/version", controller.Version)

	// config reload
	r.Any("/-/reload", func(c *gin.Context) {
		log.Info("===== Server Stop! Cause: Config Reload. =====")
		os.Exit(1)
	})

	rootGroup := r.Group("/api/v1")
	// AuthRequired middleware that provide basic auth
	//rootGroup.Use(middleware.BasicAuthMiddleware())

	{
		// a ping api to test basic auth
		rootGroup.GET("/ping", controller.Ping)
		clusterController := controller.NewClusterController()
		rootGroup.GET("/clusters/:id", clusterController.GetCluster)
		rootGroup.GET("/clusters", clusterController.ListClusters)
		rootGroup.POST("/clusters", clusterController.CreateCluster)
		rootGroup.PUT("/clusters/:id", clusterController.UpdateCluster)
		rootGroup.DELETE("/clusters/:id", clusterController.DeleteCluster)
	}

}
