# \EventApiAPI

All URIs are relative to *https://api.omnistrate.cloud*

Method | HTTP request | Description
------------- | ------------- | -------------
[**EventApiDescribeEvent**](EventApiAPI.md#EventApiDescribeEvent) | **Get** /2022-09-01-00/resource-instance/event/{id} | DescribeEvent event-api
[**EventApiListAllEvents**](EventApiAPI.md#EventApiListAllEvents) | **Get** /2022-09-01-00/resource-instance/event | ListAllEvents event-api
[**EventApiListEventsForInstance**](EventApiAPI.md#EventApiListEventsForInstance) | **Get** /2022-09-01-00/resource-instance/{instanceId}/event | ListEventsForInstance event-api
[**EventApiListEventsForServicePlan**](EventApiAPI.md#EventApiListEventsForServicePlan) | **Get** /2022-09-01-00/resource-instance/{serviceProviderId}/{serviceKey}/{serviceAPIVersion}/{serviceEnvironmentKey}/{serviceModelKey}/{productTierKey}/event | ListEventsForServicePlan event-api



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
	openapiclient "github.com/omnistrate-oss/omnistrate-sdk-go/v1"
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


## EventApiListAllEvents

> ListEventsResult EventApiListAllEvents(ctx).EnvironmentType(environmentType).StartTime(startTime).EndTime(endTime).ServiceId(serviceId).ProductTierId(productTierId).SubscriptionId(subscriptionId).EventSource(eventSource).Execute()

ListAllEvents event-api

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
	environmentType := "DEV" // string | The environment type to filter by (optional)
	startTime := "2023-01-10T00:00:00Z" // string | Filter events that occurred after this time (optional)
	endTime := "2023-01-10T00:00:00Z" // string | Filter events that occurred before this time (optional)
	serviceId := "s-12345678" // string | The service ID to filter by (optional)
	productTierId := "pt-12345678" // string | The product tier ID to filter by (optional)
	subscriptionId := "sub-abcd1234" // string | The subscription ID to filter by (optional)
	eventSource := "Customer" // string | The event source to filter by (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.EventApiAPI.EventApiListAllEvents(context.Background()).EnvironmentType(environmentType).StartTime(startTime).EndTime(endTime).ServiceId(serviceId).ProductTierId(productTierId).SubscriptionId(subscriptionId).EventSource(eventSource).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `EventApiAPI.EventApiListAllEvents``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `EventApiListAllEvents`: ListEventsResult
	fmt.Fprintf(os.Stdout, "Response from `EventApiAPI.EventApiListAllEvents`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiEventApiListAllEventsRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **environmentType** | **string** | The environment type to filter by | 
 **startTime** | **string** | Filter events that occurred after this time | 
 **endTime** | **string** | Filter events that occurred before this time | 
 **serviceId** | **string** | The service ID to filter by | 
 **productTierId** | **string** | The product tier ID to filter by | 
 **subscriptionId** | **string** | The subscription ID to filter by | 
 **eventSource** | **string** | The event source to filter by | 

### Return type

[**ListEventsResult**](ListEventsResult.md)

### Authorization

[api_key_header_Authorization](../README.md#api_key_header_Authorization)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json, application/vnd.goa.error

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## EventApiListEventsForInstance

> ListEventsResult EventApiListEventsForInstance(ctx, instanceId).SubscriptionId(subscriptionId).Execute()

ListEventsForInstance event-api

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
	instanceId := "instance-12345678" // string | The ID of the resource instance
	subscriptionId := "sub-abcd1234" // string | Subscription Id (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.EventApiAPI.EventApiListEventsForInstance(context.Background(), instanceId).SubscriptionId(subscriptionId).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `EventApiAPI.EventApiListEventsForInstance``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `EventApiListEventsForInstance`: ListEventsResult
	fmt.Fprintf(os.Stdout, "Response from `EventApiAPI.EventApiListEventsForInstance`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**instanceId** | **string** | The ID of the resource instance | 

### Other Parameters

Other parameters are passed through a pointer to a apiEventApiListEventsForInstanceRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **subscriptionId** | **string** | Subscription Id | 

### Return type

[**ListEventsResult**](ListEventsResult.md)

### Authorization

[api_key_header_Authorization](../README.md#api_key_header_Authorization)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json, application/vnd.goa.error

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## EventApiListEventsForServicePlan

> ListEventsResult EventApiListEventsForServicePlan(ctx, serviceProviderId, serviceKey, serviceAPIVersion, serviceEnvironmentKey, serviceModelKey, productTierKey).SubscriptionId(subscriptionId).Execute()

ListEventsForServicePlan event-api

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
	serviceProviderId := "omnistrate" // string | The service provider ID
	serviceKey := "service-orchestration" // string | The service name
	serviceAPIVersion := "v1" // string | The service API version
	serviceEnvironmentKey := "dev" // string | The service environment name
	serviceModelKey := "hosted" // string | The service model name
	productTierKey := "premium" // string | The product tier name
	subscriptionId := "sub-abcd1234" // string | Subscription Id (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.EventApiAPI.EventApiListEventsForServicePlan(context.Background(), serviceProviderId, serviceKey, serviceAPIVersion, serviceEnvironmentKey, serviceModelKey, productTierKey).SubscriptionId(subscriptionId).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `EventApiAPI.EventApiListEventsForServicePlan``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `EventApiListEventsForServicePlan`: ListEventsResult
	fmt.Fprintf(os.Stdout, "Response from `EventApiAPI.EventApiListEventsForServicePlan`: %v\n", resp)
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

Other parameters are passed through a pointer to a apiEventApiListEventsForServicePlanRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------






 **subscriptionId** | **string** | Subscription Id | 

### Return type

[**ListEventsResult**](ListEventsResult.md)

### Authorization

[api_key_header_Authorization](../README.md#api_key_header_Authorization)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json, application/vnd.goa.error

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

