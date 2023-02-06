package eventhub

import (
	"context"

	"github.com/Azure/azure-sdk-for-go/sdk/resourcemanager/eventhub/armeventhub"
	"github.com/selefra/selefra-provider-azure/azure_client"
	"github.com/selefra/selefra-provider-azure/table_schema_generator"
	"github.com/selefra/selefra-provider-sdk/provider/schema"
	"github.com/selefra/selefra-provider-sdk/provider/transformer/column_value_extractor"
)

type TableAzureEventhubNamespaceNetworkRuleSetsGenerator struct {
}

var _ table_schema_generator.TableSchemaGenerator = &TableAzureEventhubNamespaceNetworkRuleSetsGenerator{}

func (x *TableAzureEventhubNamespaceNetworkRuleSetsGenerator) GetTableName() string {
	return "azure_eventhub_namespace_network_rule_sets"
}

func (x *TableAzureEventhubNamespaceNetworkRuleSetsGenerator) GetTableDescription() string {
	return ""
}

func (x *TableAzureEventhubNamespaceNetworkRuleSetsGenerator) GetVersion() uint64 {
	return 0
}

func (x *TableAzureEventhubNamespaceNetworkRuleSetsGenerator) GetOptions() *schema.TableOptions {
	return &schema.TableOptions{}
}

func (x *TableAzureEventhubNamespaceNetworkRuleSetsGenerator) GetDataSource() *schema.DataSource {
	return &schema.DataSource{
		Pull: func(ctx context.Context, clientMeta *schema.ClientMeta, client any, task *schema.DataSourcePullTask, resultChannel chan<- any) *schema.Diagnostics {
			p := task.ParentRawResult.(*armeventhub.EHNamespace)
			cl := client.(*azure_client.Client)
			svc, err := armeventhub.NewNamespacesClient(cl.SubscriptionId, cl.Creds, cl.Options)
			if err != nil {
				return schema.NewDiagnosticsErrorPullTable(task.Table, err)

			}
			group, err := azure_client.ParseResourceGroup(*p.ID)
			if err != nil {
				return schema.NewDiagnosticsErrorPullTable(task.Table, err)

			}

			resp, err := svc.ListNetworkRuleSet(ctx, group, *p.Name, nil)
			if err != nil {
				return schema.NewDiagnosticsErrorPullTable(task.Table, err)

			}
			resultChannel <- resp.Value

			return nil
		},
	}
}

func (x *TableAzureEventhubNamespaceNetworkRuleSetsGenerator) GetExpandClientTask() func(ctx context.Context, clientMeta *schema.ClientMeta, client any, task *schema.DataSourcePullTask) []*schema.ClientTaskContext {
	return nil
}

func (x *TableAzureEventhubNamespaceNetworkRuleSetsGenerator) GetColumns() []*schema.Column {
	return []*schema.Column{
		table_schema_generator.NewColumnBuilder().ColumnName("location").ColumnType(schema.ColumnTypeString).
			Extractor(column_value_extractor.StructSelector("Location")).Build(),
		table_schema_generator.NewColumnBuilder().ColumnName("name").ColumnType(schema.ColumnTypeString).
			Extractor(column_value_extractor.StructSelector("Name")).Build(),
		table_schema_generator.NewColumnBuilder().ColumnName("system_data").ColumnType(schema.ColumnTypeJSON).
			Extractor(column_value_extractor.StructSelector("SystemData")).Build(),
		table_schema_generator.NewColumnBuilder().ColumnName("type").ColumnType(schema.ColumnTypeString).
			Extractor(column_value_extractor.StructSelector("Type")).Build(),
		table_schema_generator.NewColumnBuilder().ColumnName("selefra_id").ColumnType(schema.ColumnTypeString).SetUnique().Description("random id").
			Extractor(column_value_extractor.UUID()).Build(),
		table_schema_generator.NewColumnBuilder().ColumnName("azure_eventhub_namespaces_selefra_id").ColumnType(schema.ColumnTypeString).SetNotNull().Description("fk to azure_eventhub_namespaces.selefra_id").
			Extractor(column_value_extractor.ParentColumnValue("selefra_id")).Build(),
		table_schema_generator.NewColumnBuilder().ColumnName("properties").ColumnType(schema.ColumnTypeJSON).
			Extractor(column_value_extractor.StructSelector("Properties")).Build(),
		table_schema_generator.NewColumnBuilder().ColumnName("id").ColumnType(schema.ColumnTypeString).
			Extractor(column_value_extractor.StructSelector("ID")).Build(),
	}
}

func (x *TableAzureEventhubNamespaceNetworkRuleSetsGenerator) GetSubTables() []*schema.Table {
	return nil
}
