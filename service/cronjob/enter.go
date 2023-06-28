package cronjob

import "baimei-server/convert"

// @Author: morris
type ServiceGroup struct {
	CronJobSetService
}

var podConvert = convert.ConvertGroupApp.PodConvert
