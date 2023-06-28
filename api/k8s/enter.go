package k8s

import (
	"baimei-server/service"
	"baimei-server/validate"
)

type ApiGroup struct {
	PodApi
	NamespaceApi
	NodeApi
	ConfigMapApi
	SecretApi
	PVApi
	PVCApi
	SCApi
	SvcApi
	IngressApi
	IngRouteApi
	StatefulSetApi
	DeploymentApi
	DaemonsetApi
	JobApi
	CronJobApi
}

var podValidate = validate.ValidateGroupApp.PodValidate
var podService = service.ServiceGroupApp.PodServiceGroup.PodService
var nodeService = service.ServiceGroupApp.NodeServiceGroup.NodeService
var configMapService = service.ServiceGroupApp.ConfigMapServiceGroup.ConfigMapService
var secretService = service.ServiceGroupApp.SecretServiceGroup.SecretService
var pvService = service.ServiceGroupApp.PVServiceGroup.PVService
var pvcService = service.ServiceGroupApp.PVCServiceGroup.PVCService
var scService = service.ServiceGroupApp.SCServiceGroup.SCService
var svcService = service.ServiceGroupApp.SvcServiceGroup.SvcService
var ingressService = service.ServiceGroupApp.IngressServiceGroup.IngressService
var ingRouteService = service.ServiceGroupApp.IngRouteServiceGroup.IngRouteService
var statefulSetService = service.ServiceGroupApp.StatefulSetServiceGroup.StatefulSetService
var deploymentService = service.ServiceGroupApp.DeploymentServiceGroup.DeploymentService
var daemonsetService = service.ServiceGroupApp.DaemonSetServiceGroup.DaemonSetService
var jobService = service.ServiceGroupApp.JobServiceGroup.JobSetService
var cronJobService = service.ServiceGroupApp.CronJobServiceGroup.CronJobSetService
