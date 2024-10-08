# \EventApiAPI

All URIs are relative to *https://api.omnistrate.cloud*

Method | HTTP request | Description
------------- | ------------- | -------------
[**EventApiDescribeEvent**](EventApiAPI.md#EventApiDescribeEvent) | **Get** /2022-09-01-00/resource-instance/event/{id} | DescribeEvent event-api
[**EventApiListAllEvent**](EventApiAPI.md#EventApiListAllEvent) | **Get** /2022-09-01-00/resource-instance/{serviceProviderId}/{serviceKey}/{serviceAPIVersion}/{serviceEnvironmentKey}/{serviceModelKey}/{productTierKey}/event | ListAllEvent event-api
[**EventApiListEvent**](EventApiAPI.md#EventApiListEvent) | **Get** /2022-09-01-00/resource-instance/{instanceId}/event | ListEvent event-api



## EventApiDescribeEvent

> DescribeEventResult EventApiDescribeEvent(ctx, id).SubscriptionId(subscriptionId).Execute()

DescribeEvent event-api

### Example

```go
package main

import (
	"context"
	"fmt"
	"os"
	openapiclient "github.com/omnistrate/omnistrate-sdk-go/registration"
)

func main() {
	id := "event-12345678" // string | The ID of the event
	subscriptionId := "sub-abcd1234" // string | Subscription Id (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.EventApiAPI.EventApiDescribeEvent(context.Background(), id).SubscriptionId(subscriptionId).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `EventApiAPI.EventApiDescribeEvent``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `EventApiDescribeEvent`: DescribeEventResult
	fmt.Fprintf(os.Stdout, "Response from `EventApiAPI.EventApiDescribeEvent`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **string** | The ID of the event | 

### Other Parameters

Other parameters are passed through a pointer to a apiEventApiDescribeEventRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **subscriptionId** | **string** | Subscription Id | 

### Return type

[**DescribeEventResult**](DescribeEventResult.md)

### Authorization

[api_key_header_Authorization](../README.md#api_key_header_Authorization)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json, application/vnd.goa.error

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## EventApiListAllEvent

> ListEventResult EventApiListAllEvent(ctx, serviceProviderId, serviceKey, serviceAPIVersion, serviceEnvironmentKey, serviceModelKey, productTierKey).SubscriptionId(subscriptionId).Execute()

ListAllEvent event-api

### Example

```go
package main

import (
	"context"
	"fmt"
	"os"
	openapiclient "github.com/omnistrate/omnistrate-sdk-go/registration"
)

func main() {
	serviceProviderId := "omnistrate" // string | The service provider ID
	serviceKey := "service-orchestration" // string | The service name
	serviceAPIVersion := "v1" // string | The service API version
	serviceEnvironmentKey := "dev" // string | The service environment name
	serviceModelKey := "hosted" // string | The service model name
	productTierKey := "premium" // string | The product tier name
	subscriptionId := "sub-abcd1234" // string | Subscription Id (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.EventApiAPI.EventApiListAllEvent(context.Background(), serviceProviderId, serviceKey, serviceAPIVersion, serviceEnvironmentKey, serviceModelKey, productTierKey).SubscriptionId(subscriptionId).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `EventApiAPI.EventApiListAllEvent``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `EventApiListAllEvent`: ListEventResult
	fmt.Fprintf(os.Stdout, "Response from `EventApiAPI.EventApiListAllEvent`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**serviceProviderId** | **string** | The service provider ID | 
**serviceKey** | **string** | The service name | 
**serviceAPIVersion** | **string** | The service API version | 
**serviceEnvironmentKey** | **string** | The service environment name | 
**serviceModelKey** | **string** | The service model name | 
**productTierKey** | **string** | The product tier name | 

### Other Parameters

Other parameters are passed through a pointer to a apiEventApiListAllEventRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------






 **subscriptionId** | **string** | Subscription Id | 

### Return type

[**ListEventResult**](ListEventResult.md)

### Authorization

[api_key_header_Authorization](../README.md#api_key_header_Authorization)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json, application/vnd.goa.error

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## EventApiListEvent

> ListEventResult EventApiListEvent(ctx, instanceId).SubscriptionId(subscriptionId).Execute()

ListEvent event-api

### Example

```go
package main

import (
	"context"
	"fmt"
	"os"
	openapiclient "github.com/omnistrate/omnistrate-sdk-go/registration"
)

func main() {
	instanceId := "instance-12345678" // string | The ID of the resource instance
	subscriptionId := "sub-abcd1234" // string | Subscription Id (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.EventApiAPI.EventApiListEvent(context.Background(), instanceId).SubscriptionId(subscriptionId).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `EventApiAPI.EventApiListEvent``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `EventApiListEvent`: ListEventResult
	fmt.Fprintf(os.Stdout, "Response from `EventApiAPI.EventApiListEvent`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**instanceId** | **string** | The ID of the resource instance | 

### Other Parameters

Other parameters are passed through a pointer to a apiEventApiListEventRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **subscriptionId** | **string** | Subscription Id | 

### Return type

[**ListEventResult**](ListEventResult.md)

### Authorization

[api_key_header_Authorization](../README.md#api_key_header_Authorization)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json, application/vnd.goa.error

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

