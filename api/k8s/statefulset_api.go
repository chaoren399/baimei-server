package k8s

import (
	statefulset_req "baimei-server/model/statefulset/request"
	"baimei-server/response"
	"github.com/gin-gonic/gin"
)

//@Author: morris

type StatefulSetApi struct {
}

func (StatefulSetApi) CreateOrUpdateStatefulSet(c *gin.Context) {
	var statefulSetReq statefulset_req.StatefulSet
	err := c.ShouldBind(&statefulSetReq)
	if err != nil {
		response.FailWithMessage(c, err.Error())
		return
	}
	err = statefulSetService.CreateOrUpdateStatefulSet(statefulSetReq)
	if err != nil {
		response.FailWithMessage(c, err.Error())
		return
	}
	response.Success(c)
}

func (StatefulSetApi) DeleteStatefulSet(c *gin.Context) {
	err := statefulSetService.DeleteStatefulSet(c.Param("namespace"), c.Param("name"))
	if err != nil {
		response.FailWithMessage(c, err.Error())
		return
	}
	response.Success(c)
}

func (StatefulSetApi) GetStatefulSetDetailOrList(c *gin.Context) {
	keyword := c.Query("keyword")
	namespace := c.Param("namespace")
	name := c.Query("name")
	if name == "" {
		resList, err := statefulSetService.GetStatefulSetList(namespace, keyword)
		if err != nil {
			response.FailWithMessage(c, err.Error())
			return
		}
		response.SuccessWithDetailed(c, "查询statefulset列表成功！", resList)
	} else {
		statefulSetReq, err := statefulSetService.GetStatefulSetDetail(namespace, name)
		if err != nil {
			response.FailWithMessage(c, err.Error())
			return
		}
		response.SuccessWithDetailed(c, "查询statefulset详情成功！", statefulSetReq)
	}

}
