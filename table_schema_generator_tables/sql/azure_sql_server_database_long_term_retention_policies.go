package sql

import (
	"context"

	"github.com/selefra/selefra-provider-azure/azure_client"

	"github.com/Azure/azure-sdk-for-go/sdk/resourcemanager/sql/armsql"
	"github.com/selefra/selefra-provider-azure/table_schema_generator"
	"github.com/selefra/selefra-provider-sdk/provider/schema"
	"github.com/selefra/selefra-provider-sdk/provider/transformer/column_value_extractor"
)

type TableAzureSqlServerDatabaseLongTermRetentionPoliciesGenerator struct {
}

var _ table_schema_generator.TableSchemaGenerator = &TableAzureSqlServerDatabaseLongTermRetentionPoliciesGenerator{}

func (x *TableAzureSqlServerDatabaseLongTermRetentionPoliciesGenerator) GetTableName() string {
	return "azure_sql_server_database_long_term_retention_policies"
}

func (x *TableAzureSqlServerDatabaseLongTermRetentionPoliciesGenerator) GetTableDescription() string {
	return ""
}

func (x *TableAzureSqlServerDatabaseLongTermRetentionPoliciesGenerator) GetVersion() uint64 {
	return 0
}

func (x *TableAzureSqlServerDatabaseLongTermRetentionPoliciesGenerator) GetOptions() *schema.TableOptions {
	return &schema.TableOptions{}
}

func (x *TableAzureSqlServerDatabaseLongTermRetentionPoliciesGenerator) GetDataSource() *schema.DataSource {
	return &schema.DataSource{
		Pull: func(ctx context.Context, clientMeta *schema.ClientMeta, client any, task *schema.DataSourcePullTask, resultChannel chan<- any) *schema.Diagnostics {
			pp := task.ParentTask.ParentRawResult.(*armsql.Server)
			p := task.ParentRawResult.(*armsql.Database)
			cl := client.(*azure_client.Client)
			svc, err := armsql.NewLongTermRetentionPoliciesClient(cl.SubscriptionId, cl.Creds, cl.Options)
			if err != nil {
				return schema.NewDiagnosticsErrorPullTable(task.Table, err)

			}
			group, err := azure_client.ParseResourceGroup(*p.ID)
			if err != nil {
				return schema.NewDiagnosticsErrorPullTable(task.Table, err)

			}
			pager := svc.NewListByDatabasePager(group, *pp.Name, *p.Name, nil)
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

func (x *TableAzureSqlServerDatabaseLongTermRetentionPoliciesGenerator) GetExpandClientTask() func(ctx context.Context, clientMeta *schema.ClientMeta, client any, task *schema.DataSourcePullTask) []*schema.ClientTaskContext {
	return nil
}

func (x *TableAzureSqlServerDatabaseLongTermRetentionPoliciesGenerator) GetColumns() []*schema.Column {
	return []*schema.Column{
		table_schema_generator.NewColumnBuilder().ColumnName("properties").ColumnType(schema.ColumnTypeJSON).
			Extractor(column_value_extractor.StructSelector("Properties")).Build(),
		table_schema_generator.NewColumnBuilder().ColumnName("id").ColumnType(schema.ColumnTypeString).
			Extractor(column_value_extractor.StructSelector("ID")).Build(),
		table_schema_generator.NewColumnBuilder().ColumnName("name").ColumnType(schema.ColumnTypeString).
			Extractor(column_value_extractor.StructSelector("Name")).Build(),
		table_schema_generator.NewColumnBuilder().ColumnName("type").ColumnType(schema.ColumnTypeString).
			Extractor(column_value_extractor.StructSelector("Type")).Build(),
		table_schema_generator.NewColumnBuilder().ColumnName("selefra_id").ColumnType(schema.ColumnTypeString).SetUnique().Description("random id").
			Extractor(column_value_extractor.UUID()).Build(),
		table_schema_generator.NewColumnBuilder().ColumnName("azure_sql_server_databases_selefra_id").ColumnType(schema.ColumnTypeString).SetNotNull().Description("fk to azure_sql_server_databases.selefra_id").
			Extractor(column_value_extractor.ParentColumnValue("selefra_id")).Build(),
	}
}

func (x *TableAzureSqlServerDatabaseLongTermRetentionPoliciesGenerator) GetSubTables() []*schema.Table {
	return nil
}
