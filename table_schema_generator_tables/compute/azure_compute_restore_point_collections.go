package compute

import (
	"context"

	"github.com/Azure/azure-sdk-for-go/sdk/resourcemanager/compute/armcompute/v4"

	"github.com/selefra/selefra-provider-azure/azure_client"
	"github.com/selefra/selefra-provider-azure/table_schema_generator"
	"github.com/selefra/selefra-provider-sdk/provider/schema"
	"github.com/selefra/selefra-provider-sdk/provider/transformer/column_value_extractor"
)

type TableAzureComputeRestorePointCollectionsGenerator struct {
}

var _ table_schema_generator.TableSchemaGenerator = &TableAzureComputeRestorePointCollectionsGenerator{}

func (x *TableAzureComputeRestorePointCollectionsGenerator) GetTableName() string {
	return "azure_compute_restore_point_collections"
}

func (x *TableAzureComputeRestorePointCollectionsGenerator) GetTableDescription() string {
	return ""
}

func (x *TableAzureComputeRestorePointCollectionsGenerator) GetVersion() uint64 {
	return 0
}

func (x *TableAzureComputeRestorePointCollectionsGenerator) GetOptions() *schema.TableOptions {
	return &schema.TableOptions{}
}

func (x *TableAzureComputeRestorePointCollectionsGenerator) GetDataSource() *schema.DataSource {
	return &schema.DataSource{
		Pull: func(ctx context.Context, clientMeta *schema.ClientMeta, client any, task *schema.DataSourcePullTask, resultChannel chan<- any) *schema.Diagnostics {
			cl := client.(*azure_client.Client)
			svc, err := armcompute.NewRestorePointCollectionsClient(cl.SubscriptionId, cl.Creds, cl.Options)
			if err != nil {
				return schema.NewDiagnosticsErrorPullTable(task.Table, err)

			}
			pager := svc.NewListAllPager(nil)
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

func (x *TableAzureComputeRestorePointCollectionsGenerator) GetExpandClientTask() func(ctx context.Context, clientMeta *schema.ClientMeta, client any, task *schema.DataSourcePullTask) []*schema.ClientTaskContext {
	return azure_client.ExpandSubscriptionMultiplexRegisteredNamespace("azure_compute_restore_point_collections", azure_client.Namespacemicrosoft_compute)
}

func (x *TableAzureComputeRestorePointCollectionsGenerator) GetColumns() []*schema.Column {
	return []*schema.Column{
		table_schema_generator.NewColumnBuilder().ColumnName("id").ColumnType(schema.ColumnTypeString).
			Extractor(column_value_extractor.StructSelector("ID")).Build(),
		table_schema_generator.NewColumnBuilder().ColumnName("name").ColumnType(schema.ColumnTypeString).
			Extractor(column_value_extractor.StructSelector("Name")).Build(),
		table_schema_generator.NewColumnBuilder().ColumnName("type").ColumnType(schema.ColumnTypeString).
			Extractor(column_value_extractor.StructSelector("Type")).Build(),
		table_schema_generator.NewColumnBuilder().ColumnName("selefra_id").ColumnType(schema.ColumnTypeString).SetUnique().Description("random id").
			Extractor(column_value_extractor.UUID()).Build(),
		table_schema_generator.NewColumnBuilder().ColumnName("location").ColumnType(schema.ColumnTypeString).
			Extractor(column_value_extractor.StructSelector("Location")).Build(),
		table_schema_generator.NewColumnBuilder().ColumnName("properties").ColumnType(schema.ColumnTypeJSON).
			Extractor(column_value_extractor.StructSelector("Properties")).Build(),
		table_schema_generator.NewColumnBuilder().ColumnName("tags").ColumnType(schema.ColumnTypeJSON).
			Extractor(column_value_extractor.StructSelector("Tags")).Build(),
	}
}

func (x *TableAzureComputeRestorePointCollectionsGenerator) GetSubTables() []*schema.Table {
	return nil
}
