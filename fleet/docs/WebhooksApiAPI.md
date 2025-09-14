# \WebhooksApiAPI

All URIs are relative to *https://api.omnistrate.cloud*

Method | HTTP request | Description
------------- | ------------- | -------------
[**WebhooksApiReceiveWebhook**](WebhooksApiAPI.md#WebhooksApiReceiveWebhook) | **Post** /2022-09-01-00/fleet/hooks/{id} | ReceiveWebhook webhooks-api



## WebhooksApiReceiveWebhook

> WebhooksApiReceiveWebhook(ctx, id).Body(body).Execute()

ReceiveWebhook webhooks-api

### Example

```go
package main

import (
	"context"
	"fmt"
	"os"
	openapiclient "github.com/omnistrate-oss/omnistrate-sdk-go/fleet"
)

func main() {
	id := "Veniam enim nihil." // string | The unique id per producer.
	body := "Rerum earum consequuntur omnis itaque." // string | The event data

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	r, err := apiClient.WebhooksApiAPI.WebhooksApiReceiveWebhook(context.Background(), id).Body(body).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `WebhooksApiAPI.WebhooksApiReceiveWebhook``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **string** | The unique id per producer. | 

### Other Parameters

Other parameters are passed through a pointer to a apiWebhooksApiReceiveWebhookRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **body** | **string** | The event data | 

### Return type

 (empty response body)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/vnd.goa.error

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

