package advisor

import (
	"context"

	"github.com/Azure/azure-sdk-for-go/sdk/resourcemanager/advisor/armadvisor"
	"github.com/selefra/selefra-provider-azure/azure_client"
	"github.com/selefra/selefra-provider-azure/table_schema_generator"
	"github.com/selefra/selefra-provider-sdk/provider/schema"
	"github.com/selefra/selefra-provider-sdk/provider/transformer/column_value_extractor"
)

type TableAzureAdvisorSuppressionsGenerator struct {
}

var _ table_schema_generator.TableSchemaGenerator = &TableAzureAdvisorSuppressionsGenerator{}

func (x *TableAzureAdvisorSuppressionsGenerator) GetTableName() string {
	return "azure_advisor_suppressions"
}

func (x *TableAzureAdvisorSuppressionsGenerator) GetTableDescription() string {
	return ""
}

func (x *TableAzureAdvisorSuppressionsGenerator) GetVersion() uint64 {
	return 0
}

func (x *TableAzureAdvisorSuppressionsGenerator) GetOptions() *schema.TableOptions {
	return &schema.TableOptions{}
}

func (x *TableAzureAdvisorSuppressionsGenerator) GetDataSource() *schema.DataSource {
	return &schema.DataSource{
		Pull: func(ctx context.Context, clientMeta *schema.ClientMeta, client any, task *schema.DataSourcePullTask, resultChannel chan<- any) *schema.Diagnostics {
			cl := client.(*azure_client.Client)
			svc, err := armadvisor.NewSuppressionsClient(cl.SubscriptionId, cl.Creds, cl.Options)
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

func (x *TableAzureAdvisorSuppressionsGenerator) GetExpandClientTask() func(ctx context.Context, clientMeta *schema.ClientMeta, client any, task *schema.DataSourcePullTask) []*schema.ClientTaskContext {
	return azure_client.ExpandSubscriptionMultiplexRegisteredNamespace("azure_advisor_suppressions", azure_client.Namespacemicrosoft_advisor)
}

func (x *TableAzureAdvisorSuppressionsGenerator) GetColumns() []*schema.Column {
	return []*schema.Column{
		table_schema_generator.NewColumnBuilder().ColumnName("id").ColumnType(schema.ColumnTypeString).
			Extractor(column_value_extractor.StructSelector("ID")).Build(),
		table_schema_generator.NewColumnBuilder().ColumnName("name").ColumnType(schema.ColumnTypeString).
			Extractor(column_value_extractor.StructSelector("Name")).Build(),
		table_schema_generator.NewColumnBuilder().ColumnName("type").ColumnType(schema.ColumnTypeString).
			Extractor(column_value_extractor.StructSelector("Type")).Build(),
		table_schema_generator.NewColumnBuilder().ColumnName("selefra_id").ColumnType(schema.ColumnTypeString).SetUnique().Description("random id").
			Extractor(column_value_extractor.UUID()).Build(),
		table_schema_generator.NewColumnBuilder().ColumnName("properties").ColumnType(schema.ColumnTypeJSON).
			Extractor(column_value_extractor.StructSelector("Properties")).Build(),
	}
}

func (x *TableAzureAdvisorSuppressionsGenerator) GetSubTables() []*schema.Table {
	return nil
}
