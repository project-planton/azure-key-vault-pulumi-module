package pkg

import (
	"github.com/pkg/errors"
	"github.com/plantoncloud/planton-cloud-apis/zzgo/cloud/planton/apis/code2cloud/v1/azure/azurekeyvault"
	"github.com/pulumi/pulumi-azure/sdk/v5/go/azure"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

func Resources(ctx *pulumi.Context, stackInput *azurekeyvault.AzureKeyVaultStackInput) error {
	azureCredential := stackInput.AzureCredential
	//create azure provider using the credentials from the input
	_, err := azure.NewProvider(ctx,
		"azure",
		&azure.ProviderArgs{
			ClientId:       pulumi.String(azureCredential.Spec.ClientId),
			ClientSecret:   pulumi.String(azureCredential.Spec.ClientSecret),
			SubscriptionId: pulumi.String(azureCredential.Spec.SubscriptionId),
			TenantId:       pulumi.String(azureCredential.Spec.TenantId),
		})
	if err != nil {
		return errors.Wrap(err, "failed to create azure provider")
	}
	return nil
}
