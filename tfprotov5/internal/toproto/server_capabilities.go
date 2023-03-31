package toproto

import (
	"github.com/pulumi/terraform-plugin-go/tfprotov5"
	"github.com/pulumi/terraform/pkg/tfplugin5"
)

func GetProviderSchema_ServerCapabilities(in *tfprotov5.ServerCapabilities) *tfplugin5.GetProviderSchema_ServerCapabilities {
	if in == nil {
		return nil
	}

	return &tfplugin5.GetProviderSchema_ServerCapabilities{
		PlanDestroy: in.PlanDestroy,
	}
}
