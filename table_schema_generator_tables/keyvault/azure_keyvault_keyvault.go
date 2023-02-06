package keyvault

import (
	"context"

	"github.com/Azure/azure-sdk-for-go/sdk/resourcemanager/keyvault/armkeyvault"
	"github.com/selefra/selefra-provider-azure/azure_client"
	"github.com/selefra/selefra-provider-azure/table_schema_generator"
	"github.com/selefra/selefra-provider-sdk/provider/schema"
	"github.com/selefra/selefra-provider-sdk/provider/transformer/column_value_extractor"
)

type TableAzureKeyvaultKeyvaultGenerator struct {
}

var _ table_schema_generator.TableSchemaGenerator = &TableAzureKeyvaultKeyvaultGenerator{}

func (x *TableAzureKeyvaultKeyvaultGenerator) GetTableName() string {
	return "azure_keyvault_keyvault"
}

func (x *TableAzureKeyvaultKeyvaultGenerator) GetTableDescription() string {
	return ""
}

func (x *TableAzureKeyvaultKeyvaultGenerator) GetVersion() uint64 {
	return 0
}

func (x *TableAzureKeyvaultKeyvaultGenerator) GetOptions() *schema.TableOptions {
	return &schema.TableOptions{}
}

func (x *TableAzureKeyvaultKeyvaultGenerator) GetDataSource() *schema.DataSource {
	return &schema.DataSource{
		Pull: func(ctx context.Context, clientMeta *schema.ClientMeta, client any, task *schema.DataSourcePullTask, resultChannel chan<- any) *schema.Diagnostics {
			cl := client.(*azure_client.Client)
			svc, err := armkeyvault.NewVaultsClient(cl.SubscriptionId, cl.Creds, cl.Options)
			if err != nil {
				return schema.NewDiagnosticsErrorPullTable(task.Table, err)

			}
			pager := svc.NewListPager(nil)
			for pager.More() {
				p, err := pager.NextPage(ctx)
				if err != nil {
					return schema.NewDiagnosticsErrorPullTable(task.Table, err)

				}

				for _, r := range p.Value {
					azure_client.SendResults(resultChannel, &armkeyvault.Vault{
						ID:   r.ID,
						Name: r.Name,
					}, func(result any) (any, error) {
						r := result.(*armkeyvault.Vault)
						cl := client.(*azure_client.Client)
						svc, err := armkeyvault.NewVaultsClient(cl.SubscriptionId, cl.Creds, cl.Options)
						if err != nil {
							return nil, err
						}
						group, err := azure_client.ParseResourceGroup(*r.ID)
						if err != nil {
							return nil, err
						}
						resp, err := svc.Get(ctx, group, *r.Name, nil)
						if err != nil {
							return nil, err
						}
						return resp.Vault, nil
					})

				}
			}
			return nil
		},
	}
}

func (x *TableAzureKeyvaultKeyvaultGenerator) GetExpandClientTask() func(ctx context.Context, clientMeta *schema.ClientMeta, client any, task *schema.DataSourcePullTask) []*schema.ClientTaskContext {
	return azure_client.ExpandSubscription()
}

func (x *TableAzureKeyvaultKeyvaultGenerator) GetColumns() []*schema.Column {
	return []*schema.Column{
		table_schema_generator.NewColumnBuilder().ColumnName("system_data").ColumnType(schema.ColumnTypeJSON).
			Extractor(column_value_extractor.StructSelector("SystemData")).Build(),
		table_schema_generator.NewColumnBuilder().ColumnName("type").ColumnType(schema.ColumnTypeString).
			Extractor(column_value_extractor.StructSelector("Type")).Build(),
		table_schema_generator.NewColumnBuilder().ColumnName("selefra_id").ColumnType(schema.ColumnTypeString).SetUnique().Description("random id").
			Extractor(column_value_extractor.UUID()).Build(),
		table_schema_generator.NewColumnBuilder().ColumnName("properties").ColumnType(schema.ColumnTypeJSON).
			Extractor(column_value_extractor.StructSelector("Properties")).Build(),
		table_schema_generator.NewColumnBuilder().ColumnName("location").ColumnType(schema.ColumnTypeString).
			Extractor(column_value_extractor.StructSelector("Location")).Build(),
		table_schema_generator.NewColumnBuilder().ColumnName("tags").ColumnType(schema.ColumnTypeJSON).
			Extractor(column_value_extractor.StructSelector("Tags")).Build(),
		table_schema_generator.NewColumnBuilder().ColumnName("id").ColumnType(schema.ColumnTypeString).
			Extractor(column_value_extractor.StructSelector("ID")).Build(),
		table_schema_generator.NewColumnBuilder().ColumnName("name").ColumnType(schema.ColumnTypeString).
			Extractor(column_value_extractor.StructSelector("Name")).Build(),
	}
}

func (x *TableAzureKeyvaultKeyvaultGenerator) GetSubTables() []*schema.Table {
	return []*schema.Table{
		table_schema_generator.GenTableSchema(&TableAzureKeyvaultKeyvaultKeysGenerator{}),
		table_schema_generator.GenTableSchema(&TableAzureKeyvaultKeyvaultSecretsGenerator{}),
	}
}
