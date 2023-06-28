package k8s

import (
	"baimei-server/api"
	"github.com/gin-gonic/gin"
)

// @Author: morris
func initDeloymentRouter(group *gin.RouterGroup) {
	deploymentApiGroup := api.ApiGroupApp.K8SApiGroup.DeploymentApi
	group.POST("/deployment", deploymentApiGroup.CreateOrUpdateDeployment)
	group.GET("/deployment/:namespace", deploymentApiGroup.GetDeploymentDetailOrList)
	group.DELETE("/deployment/:namespace/:name", deploymentApiGroup.DeleteDeployment)
}
