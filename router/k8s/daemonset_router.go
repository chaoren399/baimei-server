package k8s

import (
	"baimei-server/api"
	"github.com/gin-gonic/gin"
)

// @Author: morris
func initDaemonSetRouter(group *gin.RouterGroup) {
	daemonbsetApiGroup := api.ApiGroupApp.K8SApiGroup.DaemonsetApi
	group.POST("/daemonset", daemonbsetApiGroup.CreateOrUpdateDaemonSet)
	group.GET("/daemonset/:namespace", daemonbsetApiGroup.GetDaemonSetDetailOrList)
	group.DELETE("/daemonset/:namespace/:name", daemonbsetApiGroup.DeleteDaemonSet)
}
