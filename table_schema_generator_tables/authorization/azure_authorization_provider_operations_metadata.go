package authorization

import (
	"context"
	"github.com/Azure/azure-sdk-for-go/sdk/resourcemanager/authorization/armauthorization/v2"

	"github.com/selefra/selefra-provider-azure/azure_client"
	"github.com/selefra/selefra-provider-azure/table_schema_generator"
	"github.com/selefra/selefra-provider-sdk/provider/schema"
	"github.com/selefra/selefra-provider-sdk/provider/transformer/column_value_extractor"
)

type TableAzureAuthorizationProviderOperationsMetadataGenerator struct {
}

var _ table_schema_generator.TableSchemaGenerator = &TableAzureAuthorizationProviderOperationsMetadataGenerator{}

func (x *TableAzureAuthorizationProviderOperationsMetadataGenerator) GetTableName() string {
	return "azure_authorization_provider_operations_metadata"
}

func (x *TableAzureAuthorizationProviderOperationsMetadataGenerator) GetTableDescription() string {
	return ""
}

func (x *TableAzureAuthorizationProviderOperationsMetadataGenerator) GetVersion() uint64 {
	return 0
}

func (x *TableAzureAuthorizationProviderOperationsMetadataGenerator) GetOptions() *schema.TableOptions {
	return &schema.TableOptions{}
}

func (x *TableAzureAuthorizationProviderOperationsMetadataGenerator) GetDataSource() *schema.DataSource {
	return &schema.DataSource{
		Pull: func(ctx context.Context, clientMeta *schema.ClientMeta, client any, task *schema.DataSourcePullTask, resultChannel chan<- any) *schema.Diagnostics {
			cl := client.(*azure_client.Client)
			svc, err := armauthorization.NewProviderOperationsMetadataClient(cl.Creds, cl.Options)
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

func (x *TableAzureAuthorizationProviderOperationsMetadataGenerator) GetExpandClientTask() func(ctx context.Context, clientMeta *schema.ClientMeta, client any, task *schema.DataSourcePullTask) []*schema.ClientTaskContext {
	return azure_client.ExpandSubscriptionMultiplexRegisteredNamespace("azure_authorization_provider_operations_metadata", azure_client.Namespacemicrosoft_authorization)
}

func (x *TableAzureAuthorizationProviderOperationsMetadataGenerator) GetColumns() []*schema.Column {
	return []*schema.Column{
		table_schema_generator.NewColumnBuilder().ColumnName("display_name").ColumnType(schema.ColumnTypeString).
			Extractor(column_value_extractor.StructSelector("DisplayName")).Build(),
		table_schema_generator.NewColumnBuilder().ColumnName("id").ColumnType(schema.ColumnTypeString).
			Extractor(column_value_extractor.StructSelector("ID")).Build(),
		table_schema_generator.NewColumnBuilder().ColumnName("name").ColumnType(schema.ColumnTypeString).
			Extractor(column_value_extractor.StructSelector("Name")).Build(),
		table_schema_generator.NewColumnBuilder().ColumnName("operations").ColumnType(schema.ColumnTypeJSON).
			Extractor(column_value_extractor.StructSelector("Operations")).Build(),
		table_schema_generator.NewColumnBuilder().ColumnName("resource_types").ColumnType(schema.ColumnTypeJSON).
			Extractor(column_value_extractor.StructSelector("ResourceTypes")).Build(),
		table_schema_generator.NewColumnBuilder().ColumnName("type").ColumnType(schema.ColumnTypeString).
			Extractor(column_value_extractor.StructSelector("Type")).Build(),
		table_schema_generator.NewColumnBuilder().ColumnName("selefra_id").ColumnType(schema.ColumnTypeString).SetUnique().Description("random id").
			Extractor(column_value_extractor.UUID()).Build(),
	}
}

func (x *TableAzureAuthorizationProviderOperationsMetadataGenerator) GetSubTables() []*schema.Table {
	return nil
}
