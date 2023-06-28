package k8s

import (
	"baimei-server/api"
	"github.com/gin-gonic/gin"
)

// @Author: morris
func initSvcRouter(group *gin.RouterGroup) {
	svcApiGroup := api.ApiGroupApp.K8SApiGroup.SvcApi
	group.POST("/svc", svcApiGroup.CreateOrUpdateSvc)
	group.GET("/svc/:namespace", svcApiGroup.GetSvcDetailOrList)
	group.DELETE("/svc/:namespace/:name", svcApiGroup.DeleteSvc)
}
