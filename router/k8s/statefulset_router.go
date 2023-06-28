package k8s

import (
	"baimei-server/api"
	"github.com/gin-gonic/gin"
)

// @Author: morris
func initStatefulSetRouter(group *gin.RouterGroup) {
	statefulsetApiGroup := api.ApiGroupApp.K8SApiGroup.StatefulSetApi
	group.POST("/statefulset", statefulsetApiGroup.CreateOrUpdateStatefulSet)
	group.GET("/statefulset/:namespace", statefulsetApiGroup.GetStatefulSetDetailOrList)
	group.DELETE("/statefulset/:namespace/:name", statefulsetApiGroup.DeleteStatefulSet)
}
