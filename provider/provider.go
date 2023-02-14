package provider

import (
	"context"

	"github.com/selefra/selefra-provider-sdk/provider"
	"github.com/selefra/selefra-provider-sdk/provider/schema"
	"github.com/spf13/viper"

	"github.com/selefra/selefra-provider-azure/azure_client"
)

const Version = "v0.0.6"

func GetProvider() *provider.Provider {
	return &provider.Provider{
		Name:      "azure",
		Version:   Version,
		TableList: GenTables(),
		ClientMeta: schema.ClientMeta{
			InitClient: func(ctx context.Context, clientMeta *schema.ClientMeta, config *viper.Viper) ([]any, *schema.Diagnostics) {
				var azureConfig azure_client.Config
				err := config.Unmarshal(&azureConfig)
				if err != nil {
					return nil, schema.NewDiagnostics().AddErrorMsg("analysis config err: %s", err.Error())
				}

				clients, err := azure_client.NewClients(azureConfig)

				if err != nil {
					clientMeta.ErrorF("new clients err: %s", err.Error())
					return nil, schema.NewDiagnostics().AddError(err)
				}

				if len(clients) == 0 {
					return nil, schema.NewDiagnostics().AddErrorMsg("account information not found")
				}

				hash := make(map[string]bool)
				res := make([]interface{}, 0, len(clients))
				for i := range clients {
					if hash[clients[i].ClientID] {
						continue
					}
					hash[clients[i].ClientID] = true
					res = append(res, clients[i])
				}
				return res, nil
			},
		},
		ConfigMeta: provider.ConfigMeta{
			GetDefaultConfigTemplate: func(ctx context.Context) string {
				return `# subscriptions:
#   - <Your Azure subscriptions>	
# tenant_id: <Your Azure tenant id>
# client_id: <Your Azure client id>
# client_secret: <Your Azure client secret>`
			},
			Validation: func(ctx context.Context, config *viper.Viper) *schema.Diagnostics {
				var azureConfig azure_client.Config
				err := config.Unmarshal(&azureConfig)
				if err != nil {
					return schema.NewDiagnostics().AddErrorMsg("analysis config err: %s", err.Error())
				}
				return nil
			},
		},
		TransformerMeta: schema.TransformerMeta{
			DefaultColumnValueConvertorBlackList: []string{
				"",
				"N/A",
				"not_supported",
			},
			DataSourcePullResultAutoExpand: true,
		},
		ErrorsHandlerMeta: schema.ErrorsHandlerMeta{

			IgnoredErrors: []schema.IgnoredError{schema.IgnoredErrorOnSaveResult},
		},
	}
}
