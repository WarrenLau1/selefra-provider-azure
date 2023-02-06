package azure_client

import (
	"context"
	"errors"
	"fmt"
	"os"
	"strings"

	// Import all autorest modules
	"github.com/Azure/azure-sdk-for-go/sdk/azcore"
	"github.com/Azure/azure-sdk-for-go/sdk/azcore/arm"
	"github.com/Azure/azure-sdk-for-go/sdk/azidentity"
	"github.com/Azure/azure-sdk-for-go/sdk/resourcemanager/resources/armresources"
	"github.com/Azure/azure-sdk-for-go/sdk/resourcemanager/subscription/armsubscription"
	"github.com/Azure/azure-sdk-for-go/services/subscription/mgmt/2020-09-01/subscription"
	"github.com/Azure/go-autorest/autorest"
	_ "github.com/Azure/go-autorest/autorest"
	"github.com/Azure/go-autorest/autorest/azure/auth"
)

type Client struct {
	subscriptions []string
	// This is to cache full objects returned from ListSubscriptions on initialisation
	SubscriptionsObjects []*armsubscription.Subscription
	registeredNamespaces map[string]map[string]bool
	resourceGroups       map[string][]*armresources.GenericResourceExpanded
	// this is set by table client multiplexer
	SubscriptionId string
	Creds          azcore.TokenCredential
	Options        *arm.ClientOptions
	ClientID       string
}

func (c *Client) discoverSubscriptions(ctx context.Context) error {
	c.subscriptions = make([]string, 0)
	subscriptionClient, err := armsubscription.NewSubscriptionsClient(c.Creds, nil)
	if err != nil {
		return err
	}
	pager := subscriptionClient.NewListPager(nil)
	for pager.More() {
		page, err := pager.NextPage(ctx)
		if err != nil {
			return err
		}
		// we record all returned values, even disabled
		c.SubscriptionsObjects = append(c.SubscriptionsObjects, page.Value...)
		for _, sub := range page.Value {
			if *sub.State == armsubscription.SubscriptionStateEnabled {
				c.subscriptions = append(c.subscriptions, strings.TrimPrefix(*sub.ID, "/subscriptions/"))
			}
		}
	}

	return nil
}

func (c *Client) disocverResourceGroups(ctx context.Context) error {
	c.resourceGroups = make(map[string][]*armresources.GenericResourceExpanded, len(c.subscriptions))
	filter := "resourceType eq 'Microsoft.Resources/resourceGroups'"
	c.registeredNamespaces = make(map[string]map[string]bool, len(c.subscriptions))

	for _, subID := range c.subscriptions {
		c.registeredNamespaces[subID] = make(map[string]bool)

		cl, err := armresources.NewClient(subID, c.Creds, nil)
		if err != nil {
			return fmt.Errorf("failed to create resource group client: %w", err)
		}

		pager := cl.NewListPager(&armresources.ClientListOptions{
			Filter: &filter,
		})
		for pager.More() {
			page, err := pager.NextPage(ctx)
			if err != nil {
				return fmt.Errorf("failed to list resource groups: %w", err)
			}
			if len(page.Value) == 0 {
				continue
			}
			c.resourceGroups[subID] = append(c.resourceGroups[subID], page.Value...)
		}

		providerClient, err := armresources.NewProvidersClient(subID, c.Creds, nil)
		if err != nil {
			return fmt.Errorf("failed to create provider client: %w", err)
		}
		providerPager := providerClient.NewListPager(nil)
		for providerPager.More() {
			providerPage, err := providerPager.NextPage(ctx)
			if err != nil {
				return fmt.Errorf("failed to list providers: %w", err)
			}
			if len(providerPage.Value) == 0 {
				continue
			}
			for _, p := range providerPage.Value {
				if p.RegistrationState != nil && *p.RegistrationState == "Registered" {
					c.registeredNamespaces[subID][strings.ToLower(*p.Namespace)] = true
				}
			}
		}
	}
	return nil
}

func NewClients(config Config) ([]*Client, error) {
	client, err := newClient(config)
	if err != nil {
		return nil, err
	}
	return []*Client{client}, nil
}

func newClient(config Config) (*Client, error) {
	var azureAuth autorest.Authorizer
	var err error

	if config.FromFile != "" {
		_ = os.Setenv("AZURE_AUTH_LOCATION", config.FromFile)
		azureAuth, err = auth.NewAuthorizerFromFile(config.ResourceBaseURI)
	} else {
		if config.TenantID != "" {
			_ = os.Setenv(auth.TenantID, config.TenantID)
		}

		if config.ClientID != "" {
			_ = os.Setenv(auth.ClientID, config.ClientID)
		}

		if config.ClientSecret != "" {
			_ = os.Setenv(auth.ClientSecret, config.ClientSecret)
		}

		azureAuth, err = auth.NewAuthorizerFromEnvironment()
		if err != nil {
			azureAuth, err = auth.NewAuthorizerFromCLI()
		}
	}

	if err != nil {
		return nil, err
	}

	client := NewAzureClient(config.Subscriptions)

	client.Creds, err = azidentity.NewDefaultAzureCredential(nil)
	if err != nil {
		client.Creds, err = azidentity.NewEnvironmentCredential(nil)
		if err != nil {
			return nil, err
		}
	}

	client.ClientID = os.Getenv(auth.ClientID)

	if len(config.Subscriptions) == 0 {
		ctx := context.Background()
		svc := subscription.NewSubscriptionsClient()
		svc.Authorizer = azureAuth
		res, err := svc.List(ctx)
		if err != nil {
			return nil, err
		}
		subscriptions := make([]string, 0)
		for res.NotDone() {
			for _, sub := range res.Values() {
				switch sub.State {
				case subscription.Disabled:
					fmt.Println("Not fetching from subscription because it is disabled subscription", *sub.SubscriptionID)
				case subscription.Deleted:
					fmt.Println("Not fetching from subscription because it is deleted subscription", *sub.SubscriptionID)
				default:
					subscriptions = append(subscriptions, *sub.SubscriptionID)
				}
			}
			err := res.NextWithContext(ctx)
			if err != nil {
				return nil, err
			}
			client.subscriptions = subscriptions
		}

		if len(client.subscriptions) == 0 {
			return nil, errors.New("could not find any subscription")
		}
	}

	//var ctx context.Context
	//if err := client.disocverResourceGroups(ctx); err != nil {
	//	return nil, err
	//}

	return client, nil
}

func NewAzureClient(subscriptions []string) *Client {
	return &Client{
		subscriptions: subscriptions,
	}
}

func (c *Client) ID() string {
	return c.SubscriptionId
}

// withSubscription allows multiplexer to create a new client with given subscriptionId
func (c *Client) withSubscription(subscriptionId string) *Client {
	newC := *c
	newC.SubscriptionId = subscriptionId
	return &newC
}
