# \OperationsApiAPI

All URIs are relative to *https://api.omnistrate.cloud*

Method | HTTP request | Description
------------- | ------------- | -------------
[**OperationsApiDeploymentCellHealth**](OperationsApiAPI.md#OperationsApiDeploymentCellHealth) | **Get** /2022-09-01-00/fleet/operations/deployment-cell-health | DeploymentCellHealth operations-api
[**OperationsApiListEvents**](OperationsApiAPI.md#OperationsApiListEvents) | **Get** /2022-09-01-00/fleet/operations/events | ListEvents operations-api
[**OperationsApiServiceHealth**](OperationsApiAPI.md#OperationsApiServiceHealth) | **Get** /2022-09-01-00/fleet/operations/service-health | ServiceHealth operations-api



## OperationsApiDeploymentCellHealth

> DeploymentCellHealthDetail OperationsApiDeploymentCellHealth(ctx).HostClusterID(hostClusterID).ServiceID(serviceID).ServiceEnvironmentID(serviceEnvironmentID).Execute()

DeploymentCellHealth operations-api

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
	hostClusterID := "hc-12345678" // string | The host cluster ID to get the health for (optional)
	serviceID := "s-12345678" // string | The service ID to get the health for (optional)
	serviceEnvironmentID := "se-12345678" // string | The service environment ID to get the health for (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.OperationsApiAPI.OperationsApiDeploymentCellHealth(context.Background()).HostClusterID(hostClusterID).ServiceID(serviceID).ServiceEnvironmentID(serviceEnvironmentID).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `OperationsApiAPI.OperationsApiDeploymentCellHealth``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `OperationsApiDeploymentCellHealth`: DeploymentCellHealthDetail
	fmt.Fprintf(os.Stdout, "Response from `OperationsApiAPI.OperationsApiDeploymentCellHealth`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiOperationsApiDeploymentCellHealthRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **hostClusterID** | **string** | The host cluster ID to get the health for | 
 **serviceID** | **string** | The service ID to get the health for | 
 **serviceEnvironmentID** | **string** | The service environment ID to get the health for | 

### Return type

[**DeploymentCellHealthDetail**](DeploymentCellHealthDetail.md)

### Authorization

[api_key_header_Authorization](../README.md#api_key_header_Authorization)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json, application/vnd.goa.error

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## OperationsApiListEvents

> ListServiceProviderEventsResult OperationsApiListEvents(ctx).NextPageToken(nextPageToken).PageSize(pageSize).EnvironmentType(environmentType).EventTypes(eventTypes).ServiceID(serviceID).ServiceEnvironmentID(serviceEnvironmentID).InstanceID(instanceID).StartDate(startDate).EndDate(endDate).ProductTierID(productTierID).Execute()

ListEvents operations-api

### Example

```go
package main

import (
	"context"
	"fmt"
	"os"
    "time"
	openapiclient "github.com/omnistrate-oss/omnistrate-sdk-go/fleet"
)

func main() {
	nextPageToken := "token" // string |  (optional)
	pageSize := int64(10) // int64 |  (optional)
	environmentType := "PROD|PRIVATE|CANARY|STAGING|QA|DEV" // string |  (optional)
	eventTypes := []string{"UnhealthyInstance|FailedDeployment|ScaleOut|UserSignUp|UserSubscription"} // []string | The event types to filter by (optional)
	serviceID := "s-123456" // string | The service ID to list events for (optional)
	serviceEnvironmentID := "se-123456" // string | The service environment ID to list events for (optional)
	instanceID := "instance-12345678" // string | The instance ID to list events for (optional)
	startDate := time.Now() // time.Time | Start date of the events (optional)
	endDate := time.Now() // time.Time | End date of the events (optional)
	productTierID := "Eum officiis et." // string |  (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.OperationsApiAPI.OperationsApiListEvents(context.Background()).NextPageToken(nextPageToken).PageSize(pageSize).EnvironmentType(environmentType).EventTypes(eventTypes).ServiceID(serviceID).ServiceEnvironmentID(serviceEnvironmentID).InstanceID(instanceID).StartDate(startDate).EndDate(endDate).ProductTierID(productTierID).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `OperationsApiAPI.OperationsApiListEvents``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `OperationsApiListEvents`: ListServiceProviderEventsResult
	fmt.Fprintf(os.Stdout, "Response from `OperationsApiAPI.OperationsApiListEvents`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiOperationsApiListEventsRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **nextPageToken** | **string** |  | 
 **pageSize** | **int64** |  | 
 **environmentType** | **string** |  | 
 **eventTypes** | **[]string** | The event types to filter by | 
 **serviceID** | **string** | The service ID to list events for | 
 **serviceEnvironmentID** | **string** | The service environment ID to list events for | 
 **instanceID** | **string** | The instance ID to list events for | 
 **startDate** | **time.Time** | Start date of the events | 
 **endDate** | **time.Time** | End date of the events | 
 **productTierID** | **string** |  | 

### Return type

[**ListServiceProviderEventsResult**](ListServiceProviderEventsResult.md)

### Authorization

[api_key_header_Authorization](../README.md#api_key_header_Authorization)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json, application/vnd.goa.error

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## OperationsApiServiceHealth

> ServiceHealthSummary OperationsApiServiceHealth(ctx).ServiceID(serviceID).ServiceEnvironmentID(serviceEnvironmentID).Execute()

ServiceHealth operations-api

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
	serviceID := "s-123456" // string | The service ID to get the health for
	serviceEnvironmentID := "se-123456" // string | The service environment ID to get the health for

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.OperationsApiAPI.OperationsApiServiceHealth(context.Background()).ServiceID(serviceID).ServiceEnvironmentID(serviceEnvironmentID).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `OperationsApiAPI.OperationsApiServiceHealth``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `OperationsApiServiceHealth`: ServiceHealthSummary
	fmt.Fprintf(os.Stdout, "Response from `OperationsApiAPI.OperationsApiServiceHealth`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiOperationsApiServiceHealthRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **serviceID** | **string** | The service ID to get the health for | 
 **serviceEnvironmentID** | **string** | The service environment ID to get the health for | 

### Return type

[**ServiceHealthSummary**](ServiceHealthSummary.md)

### Authorization

[api_key_header_Authorization](../README.md#api_key_header_Authorization)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json, application/vnd.goa.error

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

