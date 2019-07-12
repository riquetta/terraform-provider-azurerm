package media

import (
	"github.com/Azure/azure-sdk-for-go/services/mediaservices/mgmt/2018-07-01/media"
	"github.com/terraform-providers/terraform-provider-azurerm/azurerm/internal/common"
)

type Client struct {
	ServicesClient media.MediaservicesClient
}

func BuildClient(o *common.ClientOptions) *Client {
	c := Client{}

	c.ServicesClient = media.NewMediaservicesClientWithBaseURI(o.ResourceManagerEndpoint, o.SubscriptionId)
	o.ConfigureClient(&c.ServicesClient.Client, o.ResourceManagerAuthorizer)

	return &c
}
