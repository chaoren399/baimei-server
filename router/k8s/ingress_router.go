package k8s

import (
	"baimei-server/api"
	"github.com/gin-gonic/gin"
)

// @Author: morris
func initIngressRouter(group *gin.RouterGroup) {
	ingressApiGroup := api.ApiGroupApp.K8SApiGroup.IngressApi
	group.POST("/ingress", ingressApiGroup.CreateOrUpdateIngress)
	group.GET("/ingress/:namespace", ingressApiGroup.GetIngressDetailOrList)
	group.DELETE("/ingress/:namespace/:name", ingressApiGroup.DeleteIngress)
}
