package subscription

import (
	"context"

	"github.com/Azure/azure-sdk-for-go/sdk/resourcemanager/subscription/armsubscription"
	"github.com/selefra/selefra-provider-azure/azure_client"
	"github.com/selefra/selefra-provider-azure/table_schema_generator"
	"github.com/selefra/selefra-provider-sdk/provider/schema"
	"github.com/selefra/selefra-provider-sdk/provider/transformer/column_value_extractor"
)

type TableAzureSubscriptionSubscriptionsGenerator struct {
}

var _ table_schema_generator.TableSchemaGenerator = &TableAzureSubscriptionSubscriptionsGenerator{}

func (x *TableAzureSubscriptionSubscriptionsGenerator) GetTableName() string {
	return "azure_subscription_subscriptions"
}

func (x *TableAzureSubscriptionSubscriptionsGenerator) GetTableDescription() string {
	return ""
}

func (x *TableAzureSubscriptionSubscriptionsGenerator) GetVersion() uint64 {
	return 0
}

func (x *TableAzureSubscriptionSubscriptionsGenerator) GetOptions() *schema.TableOptions {
	return &schema.TableOptions{}
}

func (x *TableAzureSubscriptionSubscriptionsGenerator) GetDataSource() *schema.DataSource {
	return &schema.DataSource{
		Pull: func(ctx context.Context, clientMeta *schema.ClientMeta, client any, task *schema.DataSourcePullTask, resultChannel chan<- any) *schema.Diagnostics {
			cl := client.(*azure_client.Client)

			if len(cl.SubscriptionsObjects) != 0 {
				resultChannel <- cl.SubscriptionsObjects
				return nil
			}

			svc, err := armsubscription.NewSubscriptionsClient(cl.Creds, cl.Options)
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

func (x *TableAzureSubscriptionSubscriptionsGenerator) GetExpandClientTask() func(ctx context.Context, clientMeta *schema.ClientMeta, client any, task *schema.DataSourcePullTask) []*schema.ClientTaskContext {
	return azure_client.ExpandSingleSubscription()
}

func (x *TableAzureSubscriptionSubscriptionsGenerator) GetColumns() []*schema.Column {
	return []*schema.Column{
		table_schema_generator.NewColumnBuilder().ColumnName("selefra_id").ColumnType(schema.ColumnTypeString).SetUnique().Description("random id").
			Extractor(column_value_extractor.UUID()).Build(),
		table_schema_generator.NewColumnBuilder().ColumnName("authorization_source").ColumnType(schema.ColumnTypeString).
			Extractor(column_value_extractor.StructSelector("AuthorizationSource")).Build(),
		table_schema_generator.NewColumnBuilder().ColumnName("subscription_policies").ColumnType(schema.ColumnTypeJSON).
			Extractor(column_value_extractor.StructSelector("SubscriptionPolicies")).Build(),
		table_schema_generator.NewColumnBuilder().ColumnName("display_name").ColumnType(schema.ColumnTypeString).
			Extractor(column_value_extractor.StructSelector("DisplayName")).Build(),
		table_schema_generator.NewColumnBuilder().ColumnName("id").ColumnType(schema.ColumnTypeString).
			Extractor(column_value_extractor.StructSelector("ID")).Build(),
		table_schema_generator.NewColumnBuilder().ColumnName("state").ColumnType(schema.ColumnTypeString).
			Extractor(column_value_extractor.StructSelector("State")).Build(),
		table_schema_generator.NewColumnBuilder().ColumnName("subscription_id").ColumnType(schema.ColumnTypeString).
			Extractor(column_value_extractor.StructSelector("SubscriptionID")).Build(),
	}
}

func (x *TableAzureSubscriptionSubscriptionsGenerator) GetSubTables() []*schema.Table {
	return []*schema.Table{}
}
