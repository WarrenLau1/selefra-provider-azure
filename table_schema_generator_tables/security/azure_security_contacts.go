package security

import (
	"context"

	"github.com/Azure/azure-sdk-for-go/sdk/resourcemanager/security/armsecurity"
	"github.com/selefra/selefra-provider-azure/azure_client"
	"github.com/selefra/selefra-provider-azure/table_schema_generator"
	"github.com/selefra/selefra-provider-sdk/provider/schema"
	"github.com/selefra/selefra-provider-sdk/provider/transformer/column_value_extractor"
)

type TableAzureSecurityContactsGenerator struct {
}

var _ table_schema_generator.TableSchemaGenerator = &TableAzureSecurityContactsGenerator{}

func (x *TableAzureSecurityContactsGenerator) GetTableName() string {
	return "azure_security_contacts"
}

func (x *TableAzureSecurityContactsGenerator) GetTableDescription() string {
	return ""
}

func (x *TableAzureSecurityContactsGenerator) GetVersion() uint64 {
	return 0
}

func (x *TableAzureSecurityContactsGenerator) GetOptions() *schema.TableOptions {
	return &schema.TableOptions{}
}

func (x *TableAzureSecurityContactsGenerator) GetDataSource() *schema.DataSource {
	return &schema.DataSource{
		Pull: func(ctx context.Context, clientMeta *schema.ClientMeta, client any, task *schema.DataSourcePullTask, resultChannel chan<- any) *schema.Diagnostics {
			cl := client.(*azure_client.Client)
			svc, err := armsecurity.NewContactsClient(cl.SubscriptionId, cl.Creds, cl.Options)
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

func (x *TableAzureSecurityContactsGenerator) GetExpandClientTask() func(ctx context.Context, clientMeta *schema.ClientMeta, client any, task *schema.DataSourcePullTask) []*schema.ClientTaskContext {
	return azure_client.ExpandSubscriptionMultiplexRegisteredNamespace("azure_security_contacts", azure_client.Namespacemicrosoft_security)
}

func (x *TableAzureSecurityContactsGenerator) GetColumns() []*schema.Column {
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
	}
}

func (x *TableAzureSecurityContactsGenerator) GetSubTables() []*schema.Table {
	return nil
}
