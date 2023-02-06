package azure_client

import "context"
import "github.com/selefra/selefra-provider-sdk/provider/schema"

func ExpandSubscription() func(ctx context.Context, clientMeta *schema.ClientMeta, client any, task *schema.DataSourcePullTask) []*schema.ClientTaskContext {
	return func(ctx context.Context, clientMeta *schema.ClientMeta, client any, task *schema.DataSourcePullTask) []*schema.ClientTaskContext {
		c := client.(*Client)
		var cs = make([]*schema.ClientTaskContext, 0)
		for _, subID := range c.subscriptions {
			cs = append(cs, &schema.ClientTaskContext{Client: c.withSubscription(subID)})
		}
		return cs
	}
}

func ExpandSingleSubscription() func(ctx context.Context, clientMeta *schema.ClientMeta, client any, task *schema.DataSourcePullTask) []*schema.ClientTaskContext {
	return func(ctx context.Context, clientMeta *schema.ClientMeta, client any, task *schema.DataSourcePullTask) []*schema.ClientTaskContext {
		c := client.(*Client)
		if len(c.subscriptions) == 0 {
			return []*schema.ClientTaskContext{}
		}
		return []*schema.ClientTaskContext{
			&schema.ClientTaskContext{Client: c.withSubscription(c.subscriptions[0])},
		}
	}
}

func ExpandSubscriptionMultiplexRegisteredNamespace(table, namespace string) func(ctx context.Context, clientMeta *schema.ClientMeta, client any, task *schema.DataSourcePullTask) []*schema.ClientTaskContext {
	//return func(ctx context.Context, clientMeta *schema.ClientMeta, client any, task *schema.DataSourcePullTask) []*schema.ClientTaskContext {
	//	c := client.(*Client)
	//	var cs = make([]*schema.ClientTaskContext, 0)
	//	for _, subId := range c.subscriptions {
	//		if _, ok := c.registeredNamespaces[subId][namespace]; ok {
	//			cs = append(cs, &schema.ClientTaskContext{Client: c.withSubscription(subId)})
	//		}
	//	}
	//	return cs
	//}
	return func(ctx context.Context, clientMeta *schema.ClientMeta, client any, task *schema.DataSourcePullTask) []*schema.ClientTaskContext {
		c := client.(*Client)
		var cs = make([]*schema.ClientTaskContext, 0)
		for _, subID := range c.subscriptions {
			cs = append(cs, &schema.ClientTaskContext{Client: c.withSubscription(subID)})
		}
		return cs
	}
}
