package connectedvmware

import (
	"context"

	"github.com/Azure/azure-sdk-for-go/sdk/resourcemanager/connectedvmware/armconnectedvmware"
	"github.com/selefra/selefra-provider-azure/azure_client"
	"github.com/selefra/selefra-provider-azure/table_schema_generator"
	"github.com/selefra/selefra-provider-sdk/provider/schema"
	"github.com/selefra/selefra-provider-sdk/provider/transformer/column_value_extractor"
)

type TableAzureConnectedvmwareClustersGenerator struct {
}

var _ table_schema_generator.TableSchemaGenerator = &TableAzureConnectedvmwareClustersGenerator{}

func (x *TableAzureConnectedvmwareClustersGenerator) GetTableName() string {
	return "azure_connectedvmware_clusters"
}

func (x *TableAzureConnectedvmwareClustersGenerator) GetTableDescription() string {
	return ""
}

func (x *TableAzureConnectedvmwareClustersGenerator) GetVersion() uint64 {
	return 0
}

func (x *TableAzureConnectedvmwareClustersGenerator) GetOptions() *schema.TableOptions {
	return &schema.TableOptions{}
}

func (x *TableAzureConnectedvmwareClustersGenerator) GetDataSource() *schema.DataSource {
	return &schema.DataSource{
		Pull: func(ctx context.Context, clientMeta *schema.ClientMeta, client any, task *schema.DataSourcePullTask, resultChannel chan<- any) *schema.Diagnostics {
			cl := client.(*azure_client.Client)
			svc, err := armconnectedvmware.NewClustersClient(cl.SubscriptionId, cl.Creds, cl.Options)
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

func (x *TableAzureConnectedvmwareClustersGenerator) GetExpandClientTask() func(ctx context.Context, clientMeta *schema.ClientMeta, client any, task *schema.DataSourcePullTask) []*schema.ClientTaskContext {
	return azure_client.ExpandSubscriptionMultiplexRegisteredNamespace("azure_connectedvmware_clusters", azure_client.Namespacemicrosoft_connectedvmwarevsphere)
}

func (x *TableAzureConnectedvmwareClustersGenerator) GetColumns() []*schema.Column {
	return []*schema.Column{
		table_schema_generator.NewColumnBuilder().ColumnName("name").ColumnType(schema.ColumnTypeString).
			Extractor(column_value_extractor.StructSelector("Name")).Build(),
		table_schema_generator.NewColumnBuilder().ColumnName("selefra_id").ColumnType(schema.ColumnTypeString).SetUnique().Description("random id").
			Extractor(column_value_extractor.UUID()).Build(),
		table_schema_generator.NewColumnBuilder().ColumnName("properties").ColumnType(schema.ColumnTypeJSON).
			Extractor(column_value_extractor.StructSelector("Properties")).Build(),
		table_schema_generator.NewColumnBuilder().ColumnName("kind").ColumnType(schema.ColumnTypeString).
			Extractor(column_value_extractor.StructSelector("Kind")).Build(),
		table_schema_generator.NewColumnBuilder().ColumnName("id").ColumnType(schema.ColumnTypeString).
			Extractor(column_value_extractor.StructSelector("ID")).Build(),
		table_schema_generator.NewColumnBuilder().ColumnName("system_data").ColumnType(schema.ColumnTypeJSON).
			Extractor(column_value_extractor.StructSelector("SystemData")).Build(),
		table_schema_generator.NewColumnBuilder().ColumnName("type").ColumnType(schema.ColumnTypeString).
			Extractor(column_value_extractor.StructSelector("Type")).Build(),
		table_schema_generator.NewColumnBuilder().ColumnName("location").ColumnType(schema.ColumnTypeString).
			Extractor(column_value_extractor.StructSelector("Location")).Build(),
		table_schema_generator.NewColumnBuilder().ColumnName("extended_location").ColumnType(schema.ColumnTypeJSON).
			Extractor(column_value_extractor.StructSelector("ExtendedLocation")).Build(),
		table_schema_generator.NewColumnBuilder().ColumnName("tags").ColumnType(schema.ColumnTypeJSON).
			Extractor(column_value_extractor.StructSelector("Tags")).Build(),
	}
}

func (x *TableAzureConnectedvmwareClustersGenerator) GetSubTables() []*schema.Table {
	return nil
}
