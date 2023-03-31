package fromproto

import (
	"github.com/hashicorp/terraform-plugin-go/tfprotov6"
	"github.com/pulumi/terraform/pkg/tfplugin6"
)

func StringKind(in tfplugin6.StringKind) tfprotov6.StringKind {
	return tfprotov6.StringKind(in)
}
