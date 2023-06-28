package api

import (
	"baimei-server/api/example"
	"baimei-server/api/k8s"
)

type ApiGroup struct {
	ExampleApiGroup example.ApiGroup
	K8SApiGroup     k8s.ApiGroup
}

var ApiGroupApp = new(ApiGroup)
