package convert

import (
	"baimei-server/convert/configmap"
	"baimei-server/convert/node"
	"baimei-server/convert/pod"
	"baimei-server/convert/secret"
)

//@Author: morris

type ConvertGroup struct {
	PodConvert       pod.PodConvertGroup
	NodeConvert      node.Group
	ConfigMapConvert configmap.ConvertGroup
	SecretConvert    secret.ConvertGroup
}

var ConvertGroupApp = new(ConvertGroup)
