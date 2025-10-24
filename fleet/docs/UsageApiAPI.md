# \UsageApiAPI

All URIs are relative to *https://api.omnistrate.cloud*

Method | HTTP request | Description
------------- | ------------- | -------------
[**UsageApiGetCurrentUsage**](UsageApiAPI.md#UsageApiGetCurrentUsage) | **Get** /2022-09-01-00/fleet/usage/{subscriptionId} | GetCurrentUsage usage-api
[**UsageApiGetUsagePerDay**](UsageApiAPI.md#UsageApiGetUsagePerDay) | **Get** /2022-09-01-00/fleet/usage-per-day/{subscriptionId} | GetUsagePerDay usage-api



## UsageApiGetCurrentUsage

> FleetGetUsageResult UsageApiGetCurrentUsage(ctx, subscriptionId).StartDate(startDate).EndDate(endDate).Execute()

GetCurrentUsage usage-api

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
	subscriptionId := "sub-12345678" // string | The ID of the subscription to get usage for
	startDate := time.Now() // time.Time | Start date of the usage report (optional)
	endDate := time.Now() // time.Time | End date of the usage report (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.UsageApiAPI.UsageApiGetCurrentUsage(context.Background(), subscriptionId).StartDate(startDate).EndDate(endDate).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `UsageApiAPI.UsageApiGetCurrentUsage``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `UsageApiGetCurrentUsage`: FleetGetUsageResult
	fmt.Fprintf(os.Stdout, "Response from `UsageApiAPI.UsageApiGetCurrentUsage`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**subscriptionId** | **string** | The ID of the subscription to get usage for | 

### Other Parameters

Other parameters are passed through a pointer to a apiUsageApiGetCurrentUsageRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **startDate** | **time.Time** | Start date of the usage report | 
 **endDate** | **time.Time** | End date of the usage report | 

### Return type

[**FleetGetUsageResult**](FleetGetUsageResult.md)

### Authorization

[api_key_header_Authorization](../README.md#api_key_header_Authorization)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json, application/vnd.goa.error

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## UsageApiGetUsagePerDay

> FleetGetUsageResult UsageApiGetUsagePerDay(ctx, subscriptionId).StartDate(startDate).EndDate(endDate).Execute()

GetUsagePerDay usage-api

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
	subscriptionId := "sub-12345678" // string | The ID of the subscription to get usage for
	startDate := time.Now() // time.Time | Start date of the usage report (optional)
	endDate := time.Now() // time.Time | End date of the usage report (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.UsageApiAPI.UsageApiGetUsagePerDay(context.Background(), subscriptionId).StartDate(startDate).EndDate(endDate).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `UsageApiAPI.UsageApiGetUsagePerDay``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `UsageApiGetUsagePerDay`: FleetGetUsageResult
	fmt.Fprintf(os.Stdout, "Response from `UsageApiAPI.UsageApiGetUsagePerDay`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**subscriptionId** | **string** | The ID of the subscription to get usage for | 

### Other Parameters

Other parameters are passed through a pointer to a apiUsageApiGetUsagePerDayRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **startDate** | **time.Time** | Start date of the usage report | 
 **endDate** | **time.Time** | End date of the usage report | 

### Return type

[**FleetGetUsageResult**](FleetGetUsageResult.md)

### Authorization

[api_key_header_Authorization](../README.md#api_key_header_Authorization)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json, application/vnd.goa.error

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

