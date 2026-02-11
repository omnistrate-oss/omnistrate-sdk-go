# \SubscriptionApiAPI

All URIs are relative to *https://api.omnistrate.cloud*

Method | HTTP request | Description
------------- | ------------- | -------------
[**SubscriptionApiCreateSubscription**](SubscriptionApiAPI.md#SubscriptionApiCreateSubscription) | **Post** /2022-09-01-00/subscription | CreateSubscription subscription-api
[**SubscriptionApiDeleteSubscription**](SubscriptionApiAPI.md#SubscriptionApiDeleteSubscription) | **Delete** /2022-09-01-00/subscription/{id} | DeleteSubscription subscription-api
[**SubscriptionApiDescribeSubscription**](SubscriptionApiAPI.md#SubscriptionApiDescribeSubscription) | **Get** /2022-09-01-00/subscription/{id} | DescribeSubscription subscription-api
[**SubscriptionApiListSubscriptions**](SubscriptionApiAPI.md#SubscriptionApiListSubscriptions) | **Get** /2022-09-01-00/subscription | ListSubscriptions subscription-api



## SubscriptionApiCreateSubscription

> string SubscriptionApiCreateSubscription(ctx).CreateSubscriptionRequest2(createSubscriptionRequest2).Execute()

CreateSubscription subscription-api

### Example

```go
package main

import (
	"context"
	"fmt"
	"os"
	openapiclient "github.com/omnistrate-oss/omnistrate-sdk-go/v1"
)

func main() {
	createSubscriptionRequest2 := *openapiclient.NewCreateSubscriptionRequest2("Ad fugit.", "s-123456") // CreateSubscriptionRequest2 | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.SubscriptionApiAPI.SubscriptionApiCreateSubscription(context.Background()).CreateSubscriptionRequest2(createSubscriptionRequest2).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `SubscriptionApiAPI.SubscriptionApiCreateSubscription``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `SubscriptionApiCreateSubscription`: string
	fmt.Fprintf(os.Stdout, "Response from `SubscriptionApiAPI.SubscriptionApiCreateSubscription`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiSubscriptionApiCreateSubscriptionRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **createSubscriptionRequest2** | [**CreateSubscriptionRequest2**](CreateSubscriptionRequest2.md) |  | 

### Return type

**string**

### Authorization

[api_key_header_Authorization](../README.md#api_key_header_Authorization)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json, application/vnd.goa.error

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## SubscriptionApiDeleteSubscription

> SubscriptionApiDeleteSubscription(ctx, id).Execute()

DeleteSubscription subscription-api

### Example

```go
package main

import (
	"context"
	"fmt"
	"os"
	openapiclient "github.com/omnistrate-oss/omnistrate-sdk-go/v1"
)

func main() {
	id := "sub-12345678" // string | The subscription ID

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	r, err := apiClient.SubscriptionApiAPI.SubscriptionApiDeleteSubscription(context.Background(), id).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `SubscriptionApiAPI.SubscriptionApiDeleteSubscription``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **string** | The subscription ID | 

### Other Parameters

Other parameters are passed through a pointer to a apiSubscriptionApiDeleteSubscriptionRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

 (empty response body)

### Authorization

[api_key_header_Authorization](../README.md#api_key_header_Authorization)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/vnd.goa.error

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## SubscriptionApiDescribeSubscription

> DescribeSubscriptionResult SubscriptionApiDescribeSubscription(ctx, id).Execute()

DescribeSubscription subscription-api

### Example

```go
package main

import (
	"context"
	"fmt"
	"os"
	openapiclient "github.com/omnistrate-oss/omnistrate-sdk-go/v1"
)

func main() {
	id := "sub-12345678" // string | The subscription ID

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.SubscriptionApiAPI.SubscriptionApiDescribeSubscription(context.Background(), id).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `SubscriptionApiAPI.SubscriptionApiDescribeSubscription``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `SubscriptionApiDescribeSubscription`: DescribeSubscriptionResult
	fmt.Fprintf(os.Stdout, "Response from `SubscriptionApiAPI.SubscriptionApiDescribeSubscription`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **string** | The subscription ID | 

### Other Parameters

Other parameters are passed through a pointer to a apiSubscriptionApiDescribeSubscriptionRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

[**DescribeSubscriptionResult**](DescribeSubscriptionResult.md)

### Authorization

[api_key_header_Authorization](../README.md#api_key_header_Authorization)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json, application/vnd.goa.error

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## SubscriptionApiListSubscriptions

> ListSubscriptionsResult SubscriptionApiListSubscriptions(ctx).ServiceId(serviceId).EnvironmentType(environmentType).IncludeInactive(includeInactive).Execute()

ListSubscriptions subscription-api

### Example

```go
package main

import (
	"context"
	"fmt"
	"os"
	openapiclient "github.com/omnistrate-oss/omnistrate-sdk-go/v1"
)

func main() {
	serviceId := "service-12345678" // string | Service Id (optional)
	environmentType := "DEV" // string | The environment type to filter by (optional)
	includeInactive := false // bool | Flag indicating whether to include inactive (suspended, cancelled, terminated) subscriptions (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.SubscriptionApiAPI.SubscriptionApiListSubscriptions(context.Background()).ServiceId(serviceId).EnvironmentType(environmentType).IncludeInactive(includeInactive).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `SubscriptionApiAPI.SubscriptionApiListSubscriptions``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `SubscriptionApiListSubscriptions`: ListSubscriptionsResult
	fmt.Fprintf(os.Stdout, "Response from `SubscriptionApiAPI.SubscriptionApiListSubscriptions`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiSubscriptionApiListSubscriptionsRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **serviceId** | **string** | Service Id | 
 **environmentType** | **string** | The environment type to filter by | 
 **includeInactive** | **bool** | Flag indicating whether to include inactive (suspended, cancelled, terminated) subscriptions | 

### Return type

[**ListSubscriptionsResult**](ListSubscriptionsResult.md)

### Authorization

[api_key_header_Authorization](../README.md#api_key_header_Authorization)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json, application/vnd.goa.error

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

