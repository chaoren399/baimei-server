package service

import (
	"baimei-server/service/configmap"
	"baimei-server/service/cronjob"
	"baimei-server/service/daemonset"
	"baimei-server/service/deployment"
	"baimei-server/service/ingress"
	"baimei-server/service/ingroute"
	"baimei-server/service/job"
	"baimei-server/service/node"
	"baimei-server/service/pod"
	"baimei-server/service/pv"
	"baimei-server/service/pvc"
	"baimei-server/service/sc"
	"baimei-server/service/secret"
	"baimei-server/service/statefulset"
	"baimei-server/service/svc"
)

// @Author: morris
type ServiceGroup struct {
	PodServiceGroup         pod.PodServiceGroup
	NodeServiceGroup        node.Group
	ConfigMapServiceGroup   configmap.ServiceGroup
	SecretServiceGroup      secret.SeviceGroup
	PVServiceGroup          pv.ServiceGroup
	PVCServiceGroup         pvc.ServiceGroup
	SCServiceGroup          sc.SCServiceGroup
	SvcServiceGroup         svc.ServiceGroup
	IngressServiceGroup     ingress.ServiceGroup
	IngRouteServiceGroup    ingroute.ServiceGroup
	StatefulSetServiceGroup statefulset.ServiceGroup
	DeploymentServiceGroup  deployment.ServiceGroup
	DaemonSetServiceGroup   daemonset.ServiceGroup
	JobServiceGroup         job.ServiceGroup
	CronJobServiceGroup     cronjob.ServiceGroup
}

var ServiceGroupApp = new(ServiceGroup)
