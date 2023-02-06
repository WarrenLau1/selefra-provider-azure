package portal

import (
	"context"

	"github.com/Azure/azure-sdk-for-go/sdk/resourcemanager/portal/armportal"
	"github.com/selefra/selefra-provider-azure/azure_client"
	"github.com/selefra/selefra-provider-azure/table_schema_generator"
	"github.com/selefra/selefra-provider-sdk/provider/schema"
	"github.com/selefra/selefra-provider-sdk/provider/transformer/column_value_extractor"
)

type TableAzurePortalListTenantConfigurationViolationsGenerator struct {
}

var _ table_schema_generator.TableSchemaGenerator = &TableAzurePortalListTenantConfigurationViolationsGenerator{}

func (x *TableAzurePortalListTenantConfigurationViolationsGenerator) GetTableName() string {
	return "azure_portal_list_tenant_configuration_violations"
}

func (x *TableAzurePortalListTenantConfigurationViolationsGenerator) GetTableDescription() string {
	return ""
}

func (x *TableAzurePortalListTenantConfigurationViolationsGenerator) GetVersion() uint64 {
	return 0
}

func (x *TableAzurePortalListTenantConfigurationViolationsGenerator) GetOptions() *schema.TableOptions {
	return &schema.TableOptions{}
}

func (x *TableAzurePortalListTenantConfigurationViolationsGenerator) GetDataSource() *schema.DataSource {
	return &schema.DataSource{
		Pull: func(ctx context.Context, clientMeta *schema.ClientMeta, client any, task *schema.DataSourcePullTask, resultChannel chan<- any) *schema.Diagnostics {
			cl := client.(*azure_client.Client)
			svc, err := armportal.NewListTenantConfigurationViolationsClient(cl.Creds, cl.Options)
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

func (x *TableAzurePortalListTenantConfigurationViolationsGenerator) GetExpandClientTask() func(ctx context.Context, clientMeta *schema.ClientMeta, client any, task *schema.DataSourcePullTask) []*schema.ClientTaskContext {
	return azure_client.ExpandSubscriptionMultiplexRegisteredNamespace("azure_portal_list_tenant_configuration_violations", azure_client.Namespacemicrosoft_portal)
}

func (x *TableAzurePortalListTenantConfigurationViolationsGenerator) GetColumns() []*schema.Column {
	return []*schema.Column{
		table_schema_generator.NewColumnBuilder().ColumnName("selefra_id").ColumnType(schema.ColumnTypeString).SetUnique().Description("random id").
			Extractor(column_value_extractor.UUID()).Build(),
		table_schema_generator.NewColumnBuilder().ColumnName("error_message").ColumnType(schema.ColumnTypeString).
			Extractor(column_value_extractor.StructSelector("ErrorMessage")).Build(),
		table_schema_generator.NewColumnBuilder().ColumnName("id").ColumnType(schema.ColumnTypeString).
			Extractor(column_value_extractor.StructSelector("ID")).Build(),
		table_schema_generator.NewColumnBuilder().ColumnName("user_id").ColumnType(schema.ColumnTypeString).
			Extractor(column_value_extractor.StructSelector("UserID")).Build(),
	}
}

func (x *TableAzurePortalListTenantConfigurationViolationsGenerator) GetSubTables() []*schema.Table {
	return nil
}
