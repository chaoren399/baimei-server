package configmap

import "baimei-server/convert"

// @Author: morris
type ServiceGroup struct {
	ConfigMapService
}

var configConvert = convert.ConvertGroupApp.ConfigMapConvert
