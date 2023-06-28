package k8s

import (
	pvc_req "baimei-server/model/pvc/request"
	"baimei-server/response"
	"github.com/gin-gonic/gin"
)

// @Author: morris
type PVCApi struct {
}

func (PVCApi) CreatePVC(c *gin.Context) {
	var pvcReq pvc_req.PersistentVolumeClaim
	if err := c.ShouldBind(&pvcReq); err != nil {
		response.FailWithMessage(c, err.Error())
		return
	}
	err := pvcService.CreatePVC(pvcReq)
	if err != nil {
		response.FailWithMessage(c, err.Error())
		return
	}
	response.Success(c)
}
func (PVCApi) DeletePVC(c *gin.Context) {
	namespace := c.Param("namespace")
	name := c.Param("name")
	err := pvcService.DeletePVC(namespace, name)
	if err != nil {
		response.FailWithMessage(c, err.Error())
		return
	}
	response.Success(c)
}

func (PVCApi) GetPVCList(c *gin.Context) {
	list, err := pvcService.GetPVCList(c.Param("namespace"), c.Query("keyword"))
	if err != nil {
		response.FailWithMessage(c, err.Error())
		return
	}
	response.SuccessWithDetailed(c, "获取数据成功", list)
}
