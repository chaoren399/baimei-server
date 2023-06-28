package k8s

import (
	svc_req "baimei-server/model/svc/request"
	"baimei-server/response"
	"github.com/gin-gonic/gin"
)

//@Author: morris

type SvcApi struct {
}

func (SvcApi) CreateOrUpdateSvc(c *gin.Context) {
	var serviceReq svc_req.Service
	err := c.ShouldBind(&serviceReq)
	if err != nil {
		response.FailWithMessage(c, err.Error())
		return
	}
	err = svcService.CreateOrUpdateSvc(serviceReq)
	if err != nil {
		response.FailWithMessage(c, err.Error())
		return
	}
	response.Success(c)
}

func (SvcApi) DeleteSvc(c *gin.Context) {
	namespace := c.Param("namespace")
	name := c.Param("name")
	err := svcService.DeleteSvc(namespace, name)
	if err != nil {
		response.FailWithMessage(c, err.Error())
		return
	}
	response.Success(c)
}

func (SvcApi) GetSvcDetailOrList(c *gin.Context) {
	//1. 从k8s查询
	//2. 转换为req格式(详情) || 转换为res格式（列表）
	namespace := c.Param("namespace")
	keyword := c.Query("keyword")
	name := c.Query("name")
	if name == "" {
		list, err := svcService.GetSvcList(namespace, keyword)
		if err != nil {
			response.FailWithMessage(c, err.Error())
			return
		}
		response.SuccessWithDetailed(c, "查询service列表成功！", list)
	} else {
		detail, err := svcService.GetSvcDetail(namespace, name)
		if err != nil {
			response.FailWithMessage(c, err.Error())
			return
		}
		response.SuccessWithDetailed(c, "查询service成功！", detail)
	}

}
