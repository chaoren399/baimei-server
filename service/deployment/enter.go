package deployment

import "baimei-server/convert"

// @Author: morris
type ServiceGroup struct {
	DeploymentService
}

var podConvert = convert.ConvertGroupApp.PodConvert
