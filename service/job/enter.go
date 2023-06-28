package job

import "baimei-server/convert"

// @Author: morris
type ServiceGroup struct {
	JobSetService
}

var podConvert = convert.ConvertGroupApp.PodConvert
