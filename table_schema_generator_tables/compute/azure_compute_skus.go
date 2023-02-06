package compute

import (
	"context"

	"github.com/Azure/azure-sdk-for-go/sdk/resourcemanager/compute/armcompute/v4"

	"github.com/Azure/azure-sdk-for-go/sdk/azcore/to"
	"github.com/selefra/selefra-provider-azure/azure_client"
	"github.com/selefra/selefra-provider-azure/table_schema_generator"
	"github.com/selefra/selefra-provider-sdk/provider/schema"
	"github.com/selefra/selefra-provider-sdk/provider/transformer/column_value_extractor"
)

type TableAzureComputeSkusGenerator struct {
}

var _ table_schema_generator.TableSchemaGenerator = &TableAzureComputeSkusGenerator{}

func (x *TableAzureComputeSkusGenerator) GetTableName() string {
	return "azure_compute_skus"
}

func (x *TableAzureComputeSkusGenerator) GetTableDescription() string {
	return ""
}

func (x *TableAzureComputeSkusGenerator) GetVersion() uint64 {
	return 0
}

func (x *TableAzureComputeSkusGenerator) GetOptions() *schema.TableOptions {
	return &schema.TableOptions{}
}

func (x *TableAzureComputeSkusGenerator) GetDataSource() *schema.DataSource {
	return &schema.DataSource{
		Pull: func(ctx context.Context, clientMeta *schema.ClientMeta, client any, task *schema.DataSourcePullTask, resultChannel chan<- any) *schema.Diagnostics {
			cl := client.(*azure_client.Client)
			svc, err := armcompute.NewResourceSKUsClient(cl.SubscriptionId, cl.Creds, cl.Options)
			if err != nil {
				return schema.NewDiagnosticsErrorPullTable(task.Table, err)

			}
			pager := svc.NewListPager(&armcompute.ResourceSKUsClientListOptions{IncludeExtendedLocations: to.Ptr("true")})
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

func (x *TableAzureComputeSkusGenerator) GetExpandClientTask() func(ctx context.Context, clientMeta *schema.ClientMeta, client any, task *schema.DataSourcePullTask) []*schema.ClientTaskContext {
	return azure_client.ExpandSubscriptionMultiplexRegisteredNamespace("azure_compute_skus", azure_client.Namespacemicrosoft_compute)
}

func (x *TableAzureComputeSkusGenerator) GetColumns() []*schema.Column {
	return []*schema.Column{
		table_schema_generator.NewColumnBuilder().ColumnName("kind").ColumnType(schema.ColumnTypeString).
			Extractor(column_value_extractor.StructSelector("Kind")).Build(),
		table_schema_generator.NewColumnBuilder().ColumnName("location_info").ColumnType(schema.ColumnTypeJSON).
			Extractor(column_value_extractor.StructSelector("LocationInfo")).Build(),
		table_schema_generator.NewColumnBuilder().ColumnName("locations").ColumnType(schema.ColumnTypeStringArray).
			Extractor(column_value_extractor.StructSelector("Locations")).Build(),
		table_schema_generator.NewColumnBuilder().ColumnName("size").ColumnType(schema.ColumnTypeString).
			Extractor(column_value_extractor.StructSelector("Size")).Build(),
		table_schema_generator.NewColumnBuilder().ColumnName("costs").ColumnType(schema.ColumnTypeJSON).
			Extractor(column_value_extractor.StructSelector("Costs")).Build(),
		table_schema_generator.NewColumnBuilder().ColumnName("capacity").ColumnType(schema.ColumnTypeJSON).
			Extractor(column_value_extractor.StructSelector("Capacity")).Build(),
		table_schema_generator.NewColumnBuilder().ColumnName("selefra_id").ColumnType(schema.ColumnTypeString).SetUnique().Description("random id").
			Extractor(column_value_extractor.UUID()).Build(),
		table_schema_generator.NewColumnBuilder().ColumnName("api_versions").ColumnType(schema.ColumnTypeStringArray).
			Extractor(column_value_extractor.StructSelector("APIVersions")).Build(),
		table_schema_generator.NewColumnBuilder().ColumnName("family").ColumnType(schema.ColumnTypeString).
			Extractor(column_value_extractor.StructSelector("Family")).Build(),
		table_schema_generator.NewColumnBuilder().ColumnName("resource_type").ColumnType(schema.ColumnTypeString).
			Extractor(column_value_extractor.StructSelector("ResourceType")).Build(),
		table_schema_generator.NewColumnBuilder().ColumnName("tier").ColumnType(schema.ColumnTypeString).
			Extractor(column_value_extractor.StructSelector("Tier")).Build(),
		table_schema_generator.NewColumnBuilder().ColumnName("capabilities").ColumnType(schema.ColumnTypeJSON).
			Extractor(column_value_extractor.StructSelector("Capabilities")).Build(),
		table_schema_generator.NewColumnBuilder().ColumnName("restrictions").ColumnType(schema.ColumnTypeJSON).
			Extractor(column_value_extractor.StructSelector("Restrictions")).Build(),
		table_schema_generator.NewColumnBuilder().ColumnName("name").ColumnType(schema.ColumnTypeString).
			Extractor(column_value_extractor.StructSelector("Name")).Build(),
	}
}

func (x *TableAzureComputeSkusGenerator) GetSubTables() []*schema.Table {
	return nil
}
