# \UsageApiAPI

All URIs are relative to *https://api.omnistrate.cloud*

Method | HTTP request | Description
------------- | ------------- | -------------
[**UsageApiGetCurrentUsage**](UsageApiAPI.md#UsageApiGetCurrentUsage) | **Get** /2022-09-01-00/fleet/usage | GetCurrentUsage usage-api
[**UsageApiGetUsagePerDay**](UsageApiAPI.md#UsageApiGetUsagePerDay) | **Get** /2022-09-01-00/fleet/usage-per-day | GetUsagePerDay usage-api



## UsageApiGetCurrentUsage

> FleetGetUsageResult UsageApiGetCurrentUsage(ctx).StartDate(startDate).EndDate(endDate).ServiceID(serviceID).EnvironmentID(environmentID).ProductTierID(productTierID).SubscriptionIDs(subscriptionIDs).Execute()

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
	startDate := time.Now() // time.Time | Start date of the usage report (optional)
	endDate := time.Now() // time.Time | End date of the usage report (optional)
	serviceID := "s-12345678" // string | Filter usage by service ID (optional)
	environmentID := "env-12345678" // string | Filter usage by environment ID (optional)
	productTierID := "pt-12345678" // string | Filter usage by product tier ID (optional)
	subscriptionIDs := []string{"Ea ut ex unde rerum nihil."} // []string | Filter usage by subscription IDs (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.UsageApiAPI.UsageApiGetCurrentUsage(context.Background()).StartDate(startDate).EndDate(endDate).ServiceID(serviceID).EnvironmentID(environmentID).ProductTierID(productTierID).SubscriptionIDs(subscriptionIDs).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `UsageApiAPI.UsageApiGetCurrentUsage``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `UsageApiGetCurrentUsage`: FleetGetUsageResult
	fmt.Fprintf(os.Stdout, "Response from `UsageApiAPI.UsageApiGetCurrentUsage`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiUsageApiGetCurrentUsageRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **startDate** | **time.Time** | Start date of the usage report | 
 **endDate** | **time.Time** | End date of the usage report | 
 **serviceID** | **string** | Filter usage by service ID | 
 **environmentID** | **string** | Filter usage by environment ID | 
 **productTierID** | **string** | Filter usage by product tier ID | 
 **subscriptionIDs** | **[]string** | Filter usage by subscription IDs | 

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

> FleetGetUsageResult UsageApiGetUsagePerDay(ctx).StartDate(startDate).EndDate(endDate).ServiceID(serviceID).EnvironmentID(environmentID).ProductTierID(productTierID).SubscriptionIDs(subscriptionIDs).Execute()

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
	startDate := time.Now() // time.Time | Start date of the usage report (optional)
	endDate := time.Now() // time.Time | End date of the usage report (optional)
	serviceID := "s-12345678" // string | Filter usage by service ID (optional)
	environmentID := "env-12345678" // string | Filter usage by environment ID (optional)
	productTierID := "pt-12345678" // string | Filter usage by product tier ID (optional)
	subscriptionIDs := []string{"Vel qui repudiandae a."} // []string | Filter usage by subscription IDs (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.UsageApiAPI.UsageApiGetUsagePerDay(context.Background()).StartDate(startDate).EndDate(endDate).ServiceID(serviceID).EnvironmentID(environmentID).ProductTierID(productTierID).SubscriptionIDs(subscriptionIDs).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `UsageApiAPI.UsageApiGetUsagePerDay``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `UsageApiGetUsagePerDay`: FleetGetUsageResult
	fmt.Fprintf(os.Stdout, "Response from `UsageApiAPI.UsageApiGetUsagePerDay`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiUsageApiGetUsagePerDayRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **startDate** | **time.Time** | Start date of the usage report | 
 **endDate** | **time.Time** | End date of the usage report | 
 **serviceID** | **string** | Filter usage by service ID | 
 **environmentID** | **string** | Filter usage by environment ID | 
 **productTierID** | **string** | Filter usage by product tier ID | 
 **subscriptionIDs** | **[]string** | Filter usage by subscription IDs | 

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

