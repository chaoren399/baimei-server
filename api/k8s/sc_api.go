package k8s

import (
	sc_req "baimei-server/model/sc/request"
	"baimei-server/response"
	"github.com/gin-gonic/gin"
)

// @Author: morris
type SCApi struct {
}

func (SCApi) GetSCList(c *gin.Context) {
	list, err := scService.GetSCList(c.Query("keyword"))
	if err != nil {
		response.FailWithMessage(c, err.Error())
		return
	}
	response.SuccessWithDetailed(c, "获取数据成功", list)
}

func (SCApi) CreateSC(c *gin.Context) {
	var scReq sc_req.StorageClass
	if err := c.ShouldBind(&scReq); err != nil {
		response.FailWithMessage(c, err.Error())
		return
	}
	err := scService.CreateSC(scReq)
	if err != nil {
		response.FailWithMessage(c, err.Error())
		return
	}
	response.Success(c)
}

func (SCApi) DeleteSC(c *gin.Context) {
	namespace := c.Param("namespace")
	name := c.Param("name")
	err := scService.DeleteSC(namespace, name)
	if err != nil {
		response.FailWithMessage(c, err.Error())
		return
	}
	response.Success(c)
}
