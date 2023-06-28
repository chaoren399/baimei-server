package statefulset

import "baimei-server/convert"

// @Author: morris
type ServiceGroup struct {
	StatefulSetService
}

var podConvert = convert.ConvertGroupApp.PodConvert
