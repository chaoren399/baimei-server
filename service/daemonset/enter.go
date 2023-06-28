package daemonset

import "baimei-server/convert"

// @Author: morris
type ServiceGroup struct {
	DaemonSetService
}

var podConvert = convert.ConvertGroupApp.PodConvert
