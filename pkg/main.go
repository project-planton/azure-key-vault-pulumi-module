package pkg

import (
	"github.com/pkg/errors"
	"github.com/plantoncloud/planton-cloud-apis/zzgo/cloud/planton/apis/code2cloud/v1/azure/azurekeyvault"
	"github.com/plantoncloud/pulumi-module-golang-commons/pkg/provider/azure/pulumiazureprovider"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

type ResourceStack struct {
	Input  *azurekeyvault.AzureKeyVaultStackInput
	Labels map[string]string
}

func (s *ResourceStack) Resources(ctx *pulumi.Context) error {
	//create azure provider using the credentials from the input
	_, err := pulumiazureprovider.Get(ctx, s.Input.AzureCredential)
	if err != nil {
		return errors.Wrap(err, "failed to create azure provider")
	}

	return nil
}
