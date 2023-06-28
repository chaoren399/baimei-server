package k8s

import (
	"baimei-server/api"
	"github.com/gin-gonic/gin"
)

// @Author: morris
func initJobRouter(group *gin.RouterGroup) {
	jobApiGroup := api.ApiGroupApp.K8SApiGroup.JobApi
	group.POST("/job", jobApiGroup.CreateOrUpdateJob)
	group.GET("/job/:namespace", jobApiGroup.GetJobDetailOrList)
	group.DELETE("/job/:namespace/:name", jobApiGroup.DeleteJob)
}
