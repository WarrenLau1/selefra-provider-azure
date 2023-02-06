package subscription

import (
	"context"

	"github.com/Azure/azure-sdk-for-go/sdk/resourcemanager/subscription/armsubscription"
	"github.com/selefra/selefra-provider-azure/azure_client"
	"github.com/selefra/selefra-provider-azure/table_schema_generator"
	"github.com/selefra/selefra-provider-sdk/provider/schema"
	"github.com/selefra/selefra-provider-sdk/provider/transformer/column_value_extractor"
)

type TableAzureSubscriptionTenantsGenerator struct {
}

var _ table_schema_generator.TableSchemaGenerator = &TableAzureSubscriptionTenantsGenerator{}

func (x *TableAzureSubscriptionTenantsGenerator) GetTableName() string {
	return "azure_subscription_tenants"
}

func (x *TableAzureSubscriptionTenantsGenerator) GetTableDescription() string {
	return ""
}

func (x *TableAzureSubscriptionTenantsGenerator) GetVersion() uint64 {
	return 0
}

func (x *TableAzureSubscriptionTenantsGenerator) GetOptions() *schema.TableOptions {
	return &schema.TableOptions{}
}

func (x *TableAzureSubscriptionTenantsGenerator) GetDataSource() *schema.DataSource {
	return &schema.DataSource{
		Pull: func(ctx context.Context, clientMeta *schema.ClientMeta, client any, task *schema.DataSourcePullTask, resultChannel chan<- any) *schema.Diagnostics {
			cl := client.(*azure_client.Client)
			svc, err := armsubscription.NewTenantsClient(cl.Creds, cl.Options)
			if err != nil {
				return schema.NewDiagnosticsErrorPullTable(task.Table, err)

			}
			pager := svc.NewListPager(nil)
			for pager.More() {
				p, err := pager.NextPage(ctx)
				if err != nil {
					return schema.NewDiagnosticsErrorPullTable(task.Table, err)

				}
				resultChannel <- p.Value
			}
			return nil
		},
	}
}

func (x *TableAzureSubscriptionTenantsGenerator) GetExpandClientTask() func(ctx context.Context, clientMeta *schema.ClientMeta, client any, task *schema.DataSourcePullTask) []*schema.ClientTaskContext {
	return azure_client.ExpandSingleSubscription()
}

func (x *TableAzureSubscriptionTenantsGenerator) GetColumns() []*schema.Column {
	return []*schema.Column{
		table_schema_generator.NewColumnBuilder().ColumnName("id").ColumnType(schema.ColumnTypeString).
			Extractor(column_value_extractor.StructSelector("ID")).Build(),
		table_schema_generator.NewColumnBuilder().ColumnName("tenant_id").ColumnType(schema.ColumnTypeString).
			Extractor(column_value_extractor.StructSelector("TenantID")).Build(),
		table_schema_generator.NewColumnBuilder().ColumnName("selefra_id").ColumnType(schema.ColumnTypeString).SetUnique().Description("random id").
			Extractor(column_value_extractor.UUID()).Build(),
	}
}

func (x *TableAzureSubscriptionTenantsGenerator) GetSubTables() []*schema.Table {
	return nil
}
