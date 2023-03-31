package fromproto

import (
	"github.com/pulumi/terraform-plugin-go/tfprotov6"
	"github.com/pulumi/terraform/pkg/tfplugin6"
)

func DynamicValue(in *tfplugin6.DynamicValue) *tfprotov6.DynamicValue {
	return &tfprotov6.DynamicValue{
		MsgPack: in.Msgpack,
		JSON:    in.Json,
	}
}
