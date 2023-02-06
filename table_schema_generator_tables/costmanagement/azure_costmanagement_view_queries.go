package costmanagement

import (
	"context"
	"encoding/json"

	"github.com/Azure/azure-sdk-for-go/sdk/resourcemanager/costmanagement/armcostmanagement"
	"github.com/selefra/selefra-provider-azure/azure_client"
	"github.com/selefra/selefra-provider-azure/table_schema_generator"
	"github.com/selefra/selefra-provider-sdk/provider/schema"
	"github.com/selefra/selefra-provider-sdk/provider/transformer/column_value_extractor"
)

type TableAzureCostmanagementViewQueriesGenerator struct {
}

var _ table_schema_generator.TableSchemaGenerator = &TableAzureCostmanagementViewQueriesGenerator{}

func (x *TableAzureCostmanagementViewQueriesGenerator) GetTableName() string {
	return "azure_costmanagement_view_queries"
}

func (x *TableAzureCostmanagementViewQueriesGenerator) GetTableDescription() string {
	return ""
}

func (x *TableAzureCostmanagementViewQueriesGenerator) GetVersion() uint64 {
	return 0
}

func (x *TableAzureCostmanagementViewQueriesGenerator) GetOptions() *schema.TableOptions {
	return &schema.TableOptions{}
}

func (x *TableAzureCostmanagementViewQueriesGenerator) GetDataSource() *schema.DataSource {
	return &schema.DataSource{
		Pull: func(ctx context.Context, clientMeta *schema.ClientMeta, client any, task *schema.DataSourcePullTask, resultChannel chan<- any) *schema.Diagnostics {
			cl := client.(*azure_client.Client)
			item := task.ParentRawResult.(*armcostmanagement.View)
			if item.Properties == nil {
				return nil
			}

			svc, err := armcostmanagement.NewQueryClient(cl.Creds, cl.Options)
			if err != nil {
				return schema.NewDiagnosticsErrorPullTable(task.Table, err)

			}

			b, err := json.Marshal(item.Properties.Query)
			if err != nil {
				return schema.NewDiagnosticsErrorPullTable(task.Table, err)

			}
			var qd armcostmanagement.QueryDefinition
			if err := json.Unmarshal(b, &qd); err != nil {
				return schema.NewDiagnosticsErrorPullTable(task.Table, err)

			}

			data, err := svc.Usage(ctx, *item.Properties.Scope, qd, nil)
			if err != nil {
				return schema.NewDiagnosticsErrorPullTable(task.Table, err)

			}

			resultChannel <- data
			return nil
		},
	}
}

func (x *TableAzureCostmanagementViewQueriesGenerator) GetExpandClientTask() func(ctx context.Context, clientMeta *schema.ClientMeta, client any, task *schema.DataSourcePullTask) []*schema.ClientTaskContext {
	return nil
}

func (x *TableAzureCostmanagementViewQueriesGenerator) GetColumns() []*schema.Column {
	return []*schema.Column{
		table_schema_generator.NewColumnBuilder().ColumnName("e_tag").ColumnType(schema.ColumnTypeString).
			Extractor(column_value_extractor.StructSelector("ETag")).Build(),
		table_schema_generator.NewColumnBuilder().ColumnName("location").ColumnType(schema.ColumnTypeString).
			Extractor(column_value_extractor.StructSelector("Location")).Build(),
		table_schema_generator.NewColumnBuilder().ColumnName("name").ColumnType(schema.ColumnTypeString).
			Extractor(column_value_extractor.StructSelector("Name")).Build(),
		table_schema_generator.NewColumnBuilder().ColumnName("azure_costmanagement_views_selefra_id").ColumnType(schema.ColumnTypeString).SetNotNull().Description("fk to azure_costmanagement_views.selefra_id").
			Extractor(column_value_extractor.ParentColumnValue("selefra_id")).Build(),
		table_schema_generator.NewColumnBuilder().ColumnName("type").ColumnType(schema.ColumnTypeString).
			Extractor(column_value_extractor.StructSelector("Type")).Build(),
		table_schema_generator.NewColumnBuilder().ColumnName("selefra_id").ColumnType(schema.ColumnTypeString).SetUnique().Description("random id").
			Extractor(column_value_extractor.UUID()).Build(),
		table_schema_generator.NewColumnBuilder().ColumnName("properties").ColumnType(schema.ColumnTypeJSON).
			Extractor(column_value_extractor.StructSelector("Properties")).Build(),
		table_schema_generator.NewColumnBuilder().ColumnName("id").ColumnType(schema.ColumnTypeString).
			Extractor(column_value_extractor.StructSelector("ID")).Build(),
		table_schema_generator.NewColumnBuilder().ColumnName("sku").ColumnType(schema.ColumnTypeString).
			Extractor(column_value_extractor.StructSelector("SKU")).Build(),
		table_schema_generator.NewColumnBuilder().ColumnName("tags").ColumnType(schema.ColumnTypeJSON).
			Extractor(column_value_extractor.StructSelector("Tags")).Build(),
	}
}

func (x *TableAzureCostmanagementViewQueriesGenerator) GetSubTables() []*schema.Table {
	return nil
}
