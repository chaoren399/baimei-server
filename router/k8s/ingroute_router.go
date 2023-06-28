package k8s

import (
	"baimei-server/api"
	"github.com/gin-gonic/gin"
)

// @Author: morris
func intIngRouteRouter(group *gin.RouterGroup) {
	ingRtApiGroup := api.ApiGroupApp.K8SApiGroup.IngRouteApi
	group.POST("/ingroute", ingRtApiGroup.CreateOrUpdateIngRoute)
	group.GET("/ingroute/:namespace", ingRtApiGroup.GetIngRouteDetailOrList)
	group.GET("/ingroute/:namespace/middleware", ingRtApiGroup.GetIngRouteMiddlewareList)
	group.DELETE("/ingroute/:namespace/:name", ingRtApiGroup.DeleteIngRoute)
}
