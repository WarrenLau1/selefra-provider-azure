package monitor

import (
	"context"

	"github.com/Azure/azure-sdk-for-go/sdk/resourcemanager/monitor/armmonitor"
	"github.com/selefra/selefra-provider-azure/azure_client"
	"github.com/selefra/selefra-provider-azure/table_schema_generator"
	"github.com/selefra/selefra-provider-sdk/provider/schema"
	"github.com/selefra/selefra-provider-sdk/provider/transformer/column_value_extractor"
)

type TableAzureMonitorTenantActivityLogsGenerator struct {
}

var _ table_schema_generator.TableSchemaGenerator = &TableAzureMonitorTenantActivityLogsGenerator{}

func (x *TableAzureMonitorTenantActivityLogsGenerator) GetTableName() string {
	return "azure_monitor_tenant_activity_logs"
}

func (x *TableAzureMonitorTenantActivityLogsGenerator) GetTableDescription() string {
	return ""
}

func (x *TableAzureMonitorTenantActivityLogsGenerator) GetVersion() uint64 {
	return 0
}

func (x *TableAzureMonitorTenantActivityLogsGenerator) GetOptions() *schema.TableOptions {
	return &schema.TableOptions{}
}

func (x *TableAzureMonitorTenantActivityLogsGenerator) GetDataSource() *schema.DataSource {
	return &schema.DataSource{
		Pull: func(ctx context.Context, clientMeta *schema.ClientMeta, client any, task *schema.DataSourcePullTask, resultChannel chan<- any) *schema.Diagnostics {
			cl := client.(*azure_client.Client)
			svc, err := armmonitor.NewTenantActivityLogsClient(cl.Creds, cl.Options)
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

func (x *TableAzureMonitorTenantActivityLogsGenerator) GetExpandClientTask() func(ctx context.Context, clientMeta *schema.ClientMeta, client any, task *schema.DataSourcePullTask) []*schema.ClientTaskContext {
	return azure_client.ExpandSubscriptionMultiplexRegisteredNamespace("azure_monitor_tenant_activity_logs", azure_client.Namespacemicrosoft_insights)
}

func (x *TableAzureMonitorTenantActivityLogsGenerator) GetColumns() []*schema.Column {
	return []*schema.Column{
		table_schema_generator.NewColumnBuilder().ColumnName("caller").ColumnType(schema.ColumnTypeString).
			Extractor(column_value_extractor.StructSelector("Caller")).Build(),
		table_schema_generator.NewColumnBuilder().ColumnName("category").ColumnType(schema.ColumnTypeJSON).
			Extractor(column_value_extractor.StructSelector("Category")).Build(),
		table_schema_generator.NewColumnBuilder().ColumnName("http_request").ColumnType(schema.ColumnTypeJSON).
			Extractor(column_value_extractor.StructSelector("HTTPRequest")).Build(),
		table_schema_generator.NewColumnBuilder().ColumnName("resource_provider_name").ColumnType(schema.ColumnTypeJSON).
			Extractor(column_value_extractor.StructSelector("ResourceProviderName")).Build(),
		table_schema_generator.NewColumnBuilder().ColumnName("sub_status").ColumnType(schema.ColumnTypeJSON).
			Extractor(column_value_extractor.StructSelector("SubStatus")).Build(),
		table_schema_generator.NewColumnBuilder().ColumnName("properties").ColumnType(schema.ColumnTypeJSON).
			Extractor(column_value_extractor.StructSelector("Properties")).Build(),
		table_schema_generator.NewColumnBuilder().ColumnName("subscription_id").ColumnType(schema.ColumnTypeString).
			Extractor(column_value_extractor.StructSelector("SubscriptionID")).Build(),
		table_schema_generator.NewColumnBuilder().ColumnName("selefra_id").ColumnType(schema.ColumnTypeString).SetUnique().Description("random id").
			Extractor(column_value_extractor.UUID()).Build(),
		table_schema_generator.NewColumnBuilder().ColumnName("claims").ColumnType(schema.ColumnTypeJSON).
			Extractor(column_value_extractor.StructSelector("Claims")).Build(),
		table_schema_generator.NewColumnBuilder().ColumnName("id").ColumnType(schema.ColumnTypeString).
			Extractor(column_value_extractor.StructSelector("ID")).Build(),
		table_schema_generator.NewColumnBuilder().ColumnName("level").ColumnType(schema.ColumnTypeString).
			Extractor(column_value_extractor.StructSelector("Level")).Build(),
		table_schema_generator.NewColumnBuilder().ColumnName("operation_id").ColumnType(schema.ColumnTypeString).
			Extractor(column_value_extractor.StructSelector("OperationID")).Build(),
		table_schema_generator.NewColumnBuilder().ColumnName("operation_name").ColumnType(schema.ColumnTypeJSON).
			Extractor(column_value_extractor.StructSelector("OperationName")).Build(),
		table_schema_generator.NewColumnBuilder().ColumnName("authorization").ColumnType(schema.ColumnTypeJSON).
			Extractor(column_value_extractor.StructSelector("Authorization")).Build(),
		table_schema_generator.NewColumnBuilder().ColumnName("event_data_id").ColumnType(schema.ColumnTypeString).
			Extractor(column_value_extractor.StructSelector("EventDataID")).Build(),
		table_schema_generator.NewColumnBuilder().ColumnName("event_timestamp").ColumnType(schema.ColumnTypeTimestamp).
			Extractor(column_value_extractor.StructSelector("EventTimestamp")).Build(),
		table_schema_generator.NewColumnBuilder().ColumnName("resource_group_name").ColumnType(schema.ColumnTypeString).
			Extractor(column_value_extractor.StructSelector("ResourceGroupName")).Build(),
		table_schema_generator.NewColumnBuilder().ColumnName("resource_id").ColumnType(schema.ColumnTypeString).
			Extractor(column_value_extractor.StructSelector("ResourceID")).Build(),
		table_schema_generator.NewColumnBuilder().ColumnName("submission_timestamp").ColumnType(schema.ColumnTypeTimestamp).
			Extractor(column_value_extractor.StructSelector("SubmissionTimestamp")).Build(),
		table_schema_generator.NewColumnBuilder().ColumnName("tenant_id").ColumnType(schema.ColumnTypeString).
			Extractor(column_value_extractor.StructSelector("TenantID")).Build(),
		table_schema_generator.NewColumnBuilder().ColumnName("correlation_id").ColumnType(schema.ColumnTypeString).
			Extractor(column_value_extractor.StructSelector("CorrelationID")).Build(),
		table_schema_generator.NewColumnBuilder().ColumnName("description").ColumnType(schema.ColumnTypeString).
			Extractor(column_value_extractor.StructSelector("Description")).Build(),
		table_schema_generator.NewColumnBuilder().ColumnName("event_name").ColumnType(schema.ColumnTypeJSON).
			Extractor(column_value_extractor.StructSelector("EventName")).Build(),
		table_schema_generator.NewColumnBuilder().ColumnName("resource_type").ColumnType(schema.ColumnTypeJSON).
			Extractor(column_value_extractor.StructSelector("ResourceType")).Build(),
		table_schema_generator.NewColumnBuilder().ColumnName("status").ColumnType(schema.ColumnTypeJSON).
			Extractor(column_value_extractor.StructSelector("Status")).Build(),
	}
}

func (x *TableAzureMonitorTenantActivityLogsGenerator) GetSubTables() []*schema.Table {
	return nil
}
