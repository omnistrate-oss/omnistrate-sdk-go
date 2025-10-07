# \AuditEventsApiAPI

All URIs are relative to *https://api.omnistrate.cloud*

Method | HTTP request | Description
------------- | ------------- | -------------
[**AuditEventsApiDeprecatedDescribeAuditEvent**](AuditEventsApiAPI.md#AuditEventsApiDeprecatedDescribeAuditEvent) | **Get** /2022-09-01-00/resource-instance/event/{id} | DeprecatedDescribeAuditEvent audit-events-api
[**AuditEventsApiDeprecatedListAuditEventsForInstance**](AuditEventsApiAPI.md#AuditEventsApiDeprecatedListAuditEventsForInstance) | **Get** /2022-09-01-00/resource-instance/{instanceId}/event | DeprecatedListAuditEventsForInstance audit-events-api
[**AuditEventsApiDeprecatedListAuditEventsForServicePlan**](AuditEventsApiAPI.md#AuditEventsApiDeprecatedListAuditEventsForServicePlan) | **Get** /2022-09-01-00/resource-instance/{serviceProviderId}/{serviceKey}/{serviceAPIVersion}/{serviceEnvironmentKey}/{serviceModelKey}/{productTierKey}/event | DeprecatedListAuditEventsForServicePlan audit-events-api
[**AuditEventsApiDescribeAuditEvent**](AuditEventsApiAPI.md#AuditEventsApiDescribeAuditEvent) | **Get** /2022-09-01-00/resource-instance/audit-events/{id} | DescribeAuditEvent audit-events-api
[**AuditEventsApiListAllAuditEvents**](AuditEventsApiAPI.md#AuditEventsApiListAllAuditEvents) | **Get** /2022-09-01-00/resource-instance/audit-events | ListAllAuditEvents audit-events-api
[**AuditEventsApiListAuditEventsForInstance**](AuditEventsApiAPI.md#AuditEventsApiListAuditEventsForInstance) | **Get** /2022-09-01-00/resource-instance/{instanceId}/audit-events | ListAuditEventsForInstance audit-events-api
[**AuditEventsApiListAuditEventsForServicePlan**](AuditEventsApiAPI.md#AuditEventsApiListAuditEventsForServicePlan) | **Get** /2022-09-01-00/resource-instance/{serviceProviderId}/{serviceKey}/{serviceAPIVersion}/{serviceEnvironmentKey}/{serviceModelKey}/{productTierKey}/audit-events | ListAuditEventsForServicePlan audit-events-api



## AuditEventsApiDeprecatedDescribeAuditEvent

> DescribeAuditEventResult AuditEventsApiDeprecatedDescribeAuditEvent(ctx, id).SubscriptionId(subscriptionId).Execute()

DeprecatedDescribeAuditEvent audit-events-api

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
	resp, r, err := apiClient.AuditEventsApiAPI.AuditEventsApiDeprecatedDescribeAuditEvent(context.Background(), id).SubscriptionId(subscriptionId).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `AuditEventsApiAPI.AuditEventsApiDeprecatedDescribeAuditEvent``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `AuditEventsApiDeprecatedDescribeAuditEvent`: DescribeAuditEventResult
	fmt.Fprintf(os.Stdout, "Response from `AuditEventsApiAPI.AuditEventsApiDeprecatedDescribeAuditEvent`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **string** | The ID of the event | 

### Other Parameters

Other parameters are passed through a pointer to a apiAuditEventsApiDeprecatedDescribeAuditEventRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **subscriptionId** | **string** | Subscription Id | 

### Return type

[**DescribeAuditEventResult**](DescribeAuditEventResult.md)

### Authorization

[api_key_header_Authorization](../README.md#api_key_header_Authorization)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json, application/vnd.goa.error

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## AuditEventsApiDeprecatedListAuditEventsForInstance

> ListAuditEventsResult AuditEventsApiDeprecatedListAuditEventsForInstance(ctx, instanceId).SubscriptionId(subscriptionId).Execute()

DeprecatedListAuditEventsForInstance audit-events-api

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
	resp, r, err := apiClient.AuditEventsApiAPI.AuditEventsApiDeprecatedListAuditEventsForInstance(context.Background(), instanceId).SubscriptionId(subscriptionId).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `AuditEventsApiAPI.AuditEventsApiDeprecatedListAuditEventsForInstance``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `AuditEventsApiDeprecatedListAuditEventsForInstance`: ListAuditEventsResult
	fmt.Fprintf(os.Stdout, "Response from `AuditEventsApiAPI.AuditEventsApiDeprecatedListAuditEventsForInstance`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**instanceId** | **string** | The ID of the resource instance | 

### Other Parameters

Other parameters are passed through a pointer to a apiAuditEventsApiDeprecatedListAuditEventsForInstanceRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **subscriptionId** | **string** | Subscription Id | 

### Return type

[**ListAuditEventsResult**](ListAuditEventsResult.md)

### Authorization

[api_key_header_Authorization](../README.md#api_key_header_Authorization)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json, application/vnd.goa.error

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## AuditEventsApiDeprecatedListAuditEventsForServicePlan

> ListAuditEventsResult AuditEventsApiDeprecatedListAuditEventsForServicePlan(ctx, serviceProviderId, serviceKey, serviceAPIVersion, serviceEnvironmentKey, serviceModelKey, productTierKey).SubscriptionId(subscriptionId).Execute()

DeprecatedListAuditEventsForServicePlan audit-events-api

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
	resp, r, err := apiClient.AuditEventsApiAPI.AuditEventsApiDeprecatedListAuditEventsForServicePlan(context.Background(), serviceProviderId, serviceKey, serviceAPIVersion, serviceEnvironmentKey, serviceModelKey, productTierKey).SubscriptionId(subscriptionId).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `AuditEventsApiAPI.AuditEventsApiDeprecatedListAuditEventsForServicePlan``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `AuditEventsApiDeprecatedListAuditEventsForServicePlan`: ListAuditEventsResult
	fmt.Fprintf(os.Stdout, "Response from `AuditEventsApiAPI.AuditEventsApiDeprecatedListAuditEventsForServicePlan`: %v\n", resp)
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

Other parameters are passed through a pointer to a apiAuditEventsApiDeprecatedListAuditEventsForServicePlanRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------






 **subscriptionId** | **string** | Subscription Id | 

### Return type

[**ListAuditEventsResult**](ListAuditEventsResult.md)

### Authorization

[api_key_header_Authorization](../README.md#api_key_header_Authorization)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json, application/vnd.goa.error

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## AuditEventsApiDescribeAuditEvent

> DescribeAuditEventResult AuditEventsApiDescribeAuditEvent(ctx, id).SubscriptionId(subscriptionId).Execute()

DescribeAuditEvent audit-events-api

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
	resp, r, err := apiClient.AuditEventsApiAPI.AuditEventsApiDescribeAuditEvent(context.Background(), id).SubscriptionId(subscriptionId).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `AuditEventsApiAPI.AuditEventsApiDescribeAuditEvent``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `AuditEventsApiDescribeAuditEvent`: DescribeAuditEventResult
	fmt.Fprintf(os.Stdout, "Response from `AuditEventsApiAPI.AuditEventsApiDescribeAuditEvent`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **string** | The ID of the event | 

### Other Parameters

Other parameters are passed through a pointer to a apiAuditEventsApiDescribeAuditEventRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **subscriptionId** | **string** | Subscription Id | 

### Return type

[**DescribeAuditEventResult**](DescribeAuditEventResult.md)

### Authorization

[api_key_header_Authorization](../README.md#api_key_header_Authorization)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json, application/vnd.goa.error

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## AuditEventsApiListAllAuditEvents

> ListAuditEventsResult AuditEventsApiListAllAuditEvents(ctx).NextPageToken(nextPageToken).PageSize(pageSize).ServiceID(serviceID).EnvironmentType(environmentType).EventSourceTypes(eventSourceTypes).InstanceID(instanceID).StartDate(startDate).EndDate(endDate).Execute()

ListAllAuditEvents audit-events-api

### Example

```go
package main

import (
	"context"
	"fmt"
	"os"
    "time"
	openapiclient "github.com/omnistrate-oss/omnistrate-sdk-go/v1"
)

func main() {
	nextPageToken := "token" // string |  (optional)
	pageSize := int64(10) // int64 |  (optional)
	serviceID := "s-123456" // string | The service ID to list events for (optional)
	environmentType := "PROD|PRIVATE|CANARY|STAGING|QA|DEV|GLOBAL" // string |  (optional)
	eventSourceTypes := []string{"Officiis aut."} // []string | The event types to filter by (optional)
	instanceID := "instance-12345678" // string | The instance ID to list events for (optional)
	startDate := time.Now() // time.Time | Start date of the events (optional)
	endDate := time.Now() // time.Time | End date of the events (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.AuditEventsApiAPI.AuditEventsApiListAllAuditEvents(context.Background()).NextPageToken(nextPageToken).PageSize(pageSize).ServiceID(serviceID).EnvironmentType(environmentType).EventSourceTypes(eventSourceTypes).InstanceID(instanceID).StartDate(startDate).EndDate(endDate).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `AuditEventsApiAPI.AuditEventsApiListAllAuditEvents``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `AuditEventsApiListAllAuditEvents`: ListAuditEventsResult
	fmt.Fprintf(os.Stdout, "Response from `AuditEventsApiAPI.AuditEventsApiListAllAuditEvents`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiAuditEventsApiListAllAuditEventsRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **nextPageToken** | **string** |  | 
 **pageSize** | **int64** |  | 
 **serviceID** | **string** | The service ID to list events for | 
 **environmentType** | **string** |  | 
 **eventSourceTypes** | **[]string** | The event types to filter by | 
 **instanceID** | **string** | The instance ID to list events for | 
 **startDate** | **time.Time** | Start date of the events | 
 **endDate** | **time.Time** | End date of the events | 

### Return type

[**ListAuditEventsResult**](ListAuditEventsResult.md)

### Authorization

[api_key_header_Authorization](../README.md#api_key_header_Authorization)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json, application/vnd.goa.error

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## AuditEventsApiListAuditEventsForInstance

> ListAuditEventsResult AuditEventsApiListAuditEventsForInstance(ctx, instanceId).SubscriptionId(subscriptionId).Execute()

ListAuditEventsForInstance audit-events-api

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
	resp, r, err := apiClient.AuditEventsApiAPI.AuditEventsApiListAuditEventsForInstance(context.Background(), instanceId).SubscriptionId(subscriptionId).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `AuditEventsApiAPI.AuditEventsApiListAuditEventsForInstance``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `AuditEventsApiListAuditEventsForInstance`: ListAuditEventsResult
	fmt.Fprintf(os.Stdout, "Response from `AuditEventsApiAPI.AuditEventsApiListAuditEventsForInstance`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**instanceId** | **string** | The ID of the resource instance | 

### Other Parameters

Other parameters are passed through a pointer to a apiAuditEventsApiListAuditEventsForInstanceRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **subscriptionId** | **string** | Subscription Id | 

### Return type

[**ListAuditEventsResult**](ListAuditEventsResult.md)

### Authorization

[api_key_header_Authorization](../README.md#api_key_header_Authorization)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json, application/vnd.goa.error

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## AuditEventsApiListAuditEventsForServicePlan

> ListAuditEventsResult AuditEventsApiListAuditEventsForServicePlan(ctx, serviceProviderId, serviceKey, serviceAPIVersion, serviceEnvironmentKey, serviceModelKey, productTierKey).SubscriptionId(subscriptionId).Execute()

ListAuditEventsForServicePlan audit-events-api

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
	resp, r, err := apiClient.AuditEventsApiAPI.AuditEventsApiListAuditEventsForServicePlan(context.Background(), serviceProviderId, serviceKey, serviceAPIVersion, serviceEnvironmentKey, serviceModelKey, productTierKey).SubscriptionId(subscriptionId).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `AuditEventsApiAPI.AuditEventsApiListAuditEventsForServicePlan``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `AuditEventsApiListAuditEventsForServicePlan`: ListAuditEventsResult
	fmt.Fprintf(os.Stdout, "Response from `AuditEventsApiAPI.AuditEventsApiListAuditEventsForServicePlan`: %v\n", resp)
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

Other parameters are passed through a pointer to a apiAuditEventsApiListAuditEventsForServicePlanRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------






 **subscriptionId** | **string** | Subscription Id | 

### Return type

[**ListAuditEventsResult**](ListAuditEventsResult.md)

### Authorization

[api_key_header_Authorization](../README.md#api_key_header_Authorization)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json, application/vnd.goa.error

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

