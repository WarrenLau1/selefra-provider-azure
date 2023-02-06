package devops

import (
	"context"

	"github.com/Azure/azure-sdk-for-go/sdk/resourcemanager/devops/armdevops"
	"github.com/selefra/selefra-provider-azure/azure_client"
	"github.com/selefra/selefra-provider-azure/table_schema_generator"
	"github.com/selefra/selefra-provider-sdk/provider/schema"
	"github.com/selefra/selefra-provider-sdk/provider/transformer/column_value_extractor"
)

type TableAzureDevopsPipelineTemplateDefinitionsGenerator struct {
}

var _ table_schema_generator.TableSchemaGenerator = &TableAzureDevopsPipelineTemplateDefinitionsGenerator{}

func (x *TableAzureDevopsPipelineTemplateDefinitionsGenerator) GetTableName() string {
	return "azure_devops_pipeline_template_definitions"
}

func (x *TableAzureDevopsPipelineTemplateDefinitionsGenerator) GetTableDescription() string {
	return ""
}

func (x *TableAzureDevopsPipelineTemplateDefinitionsGenerator) GetVersion() uint64 {
	return 0
}

func (x *TableAzureDevopsPipelineTemplateDefinitionsGenerator) GetOptions() *schema.TableOptions {
	return &schema.TableOptions{}
}

func (x *TableAzureDevopsPipelineTemplateDefinitionsGenerator) GetDataSource() *schema.DataSource {
	return &schema.DataSource{
		Pull: func(ctx context.Context, clientMeta *schema.ClientMeta, client any, task *schema.DataSourcePullTask, resultChannel chan<- any) *schema.Diagnostics {
			cl := client.(*azure_client.Client)
			svc, err := armdevops.NewPipelineTemplateDefinitionsClient(cl.Creds, cl.Options)
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

func (x *TableAzureDevopsPipelineTemplateDefinitionsGenerator) GetExpandClientTask() func(ctx context.Context, clientMeta *schema.ClientMeta, client any, task *schema.DataSourcePullTask) []*schema.ClientTaskContext {
	return azure_client.ExpandSubscriptionMultiplexRegisteredNamespace("azure_devops_pipeline_template_definitions", azure_client.Namespacemicrosoft_devops)
}

func (x *TableAzureDevopsPipelineTemplateDefinitionsGenerator) GetColumns() []*schema.Column {
	return []*schema.Column{
		table_schema_generator.NewColumnBuilder().ColumnName("inputs").ColumnType(schema.ColumnTypeJSON).
			Extractor(column_value_extractor.StructSelector("Inputs")).Build(),
		table_schema_generator.NewColumnBuilder().ColumnName("selefra_id").ColumnType(schema.ColumnTypeString).SetUnique().Description("random id").
			Extractor(column_value_extractor.UUID()).Build(),
		table_schema_generator.NewColumnBuilder().ColumnName("id").ColumnType(schema.ColumnTypeString).
			Extractor(column_value_extractor.StructSelector("ID")).Build(),
		table_schema_generator.NewColumnBuilder().ColumnName("description").ColumnType(schema.ColumnTypeString).
			Extractor(column_value_extractor.StructSelector("Description")).Build(),
	}
}

func (x *TableAzureDevopsPipelineTemplateDefinitionsGenerator) GetSubTables() []*schema.Table {
	return nil
}
