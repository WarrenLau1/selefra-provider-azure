package monitor

import (
	"context"
	"errors"

	"github.com/Azure/azure-sdk-for-go/sdk/azcore"
	"github.com/Azure/azure-sdk-for-go/sdk/resourcemanager/monitor/armmonitor"
	"github.com/Azure/azure-sdk-for-go/sdk/resourcemanager/resources/armresources"
	"github.com/selefra/selefra-provider-azure/azure_client"
	"github.com/selefra/selefra-provider-azure/table_schema_generator"
	"github.com/selefra/selefra-provider-sdk/provider/schema"
	"github.com/selefra/selefra-provider-sdk/provider/transformer/column_value_extractor"
)

type TableAzureMonitorDiagnosticSettingsGenerator struct {
}

var _ table_schema_generator.TableSchemaGenerator = &TableAzureMonitorDiagnosticSettingsGenerator{}

func (x *TableAzureMonitorDiagnosticSettingsGenerator) GetTableName() string {
	return "azure_monitor_diagnostic_settings"
}

func (x *TableAzureMonitorDiagnosticSettingsGenerator) GetTableDescription() string {
	return ""
}

func (x *TableAzureMonitorDiagnosticSettingsGenerator) GetVersion() uint64 {
	return 0
}

func (x *TableAzureMonitorDiagnosticSettingsGenerator) GetOptions() *schema.TableOptions {
	return &schema.TableOptions{}
}

func (x *TableAzureMonitorDiagnosticSettingsGenerator) GetDataSource() *schema.DataSource {
	return &schema.DataSource{
		Pull: func(ctx context.Context, clientMeta *schema.ClientMeta, client any, task *schema.DataSourcePullTask, resultChannel chan<- any) *schema.Diagnostics {
			cl := client.(*azure_client.Client)
			svc, err := armresources.NewClient(cl.SubscriptionId, cl.Creds, cl.Options)
			if err != nil {
				return schema.NewDiagnosticsErrorPullTable(task.Table, err)

			}
			pager := svc.NewListPager(nil)
			for pager.More() {
				p, err := pager.NextPage(ctx)
				if err != nil {
					return schema.NewDiagnosticsErrorPullTable(task.Table, err)

				}
				for _, r := range p.Value {
					svc, err := armmonitor.NewDiagnosticSettingsClient(cl.Creds, cl.Options)
					if err != nil {
						return schema.NewDiagnosticsErrorPullTable(task.Table, err)

					}
					pager := svc.NewListPager(*r.ID, nil)
					for pager.More() {
						p, err := pager.NextPage(ctx)
						if err != nil {
							if isResourceTypeNotSupported(err) {
								break
							}
							return schema.NewDiagnosticsErrorPullTable(task.Table, err)

						}
						for _, ds := range p.Value {
							resultChannel <- diagnosticSettingsWrapper{
								ds,
								*r.ID,
							}
						}
					}
				}
			}
			return nil
		},
	}
}

type diagnosticSettingsWrapper struct {
	*armmonitor.DiagnosticSettingsResource
	ResourceId string
}

func isResourceTypeNotSupported(err error) bool {
	var azureErr *azcore.ResponseError
	if errors.As(err, &azureErr) {
		return azureErr != nil && azureErr.ErrorCode == "ResourceTypeNotSupported"
	}
	return false
}

func (x *TableAzureMonitorDiagnosticSettingsGenerator) GetExpandClientTask() func(ctx context.Context, clientMeta *schema.ClientMeta, client any, task *schema.DataSourcePullTask) []*schema.ClientTaskContext {
	return azure_client.ExpandSubscriptionMultiplexRegisteredNamespace("azure_monitor_diagnostic_settings", azure_client.Namespacemicrosoft_insights)
}

func (x *TableAzureMonitorDiagnosticSettingsGenerator) GetColumns() []*schema.Column {
	return []*schema.Column{
		table_schema_generator.NewColumnBuilder().ColumnName("diagnostic_settings_resource").ColumnType(schema.ColumnTypeJSON).
			Extractor(column_value_extractor.StructSelector("DiagnosticSettingsResource")).Build(),
		table_schema_generator.NewColumnBuilder().ColumnName("resource_id").ColumnType(schema.ColumnTypeString).
			Extractor(column_value_extractor.StructSelector("ResourceId")).Build(),
		table_schema_generator.NewColumnBuilder().ColumnName("selefra_id").ColumnType(schema.ColumnTypeString).SetUnique().Description("random id").
			Extractor(column_value_extractor.UUID()).Build(),
	}
}

func (x *TableAzureMonitorDiagnosticSettingsGenerator) GetSubTables() []*schema.Table {
	return nil
}
