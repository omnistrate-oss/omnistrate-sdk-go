# \ConsumptionUsageApiAPI

All URIs are relative to *https://api.omnistrate.cloud*

Method | HTTP request | Description
------------- | ------------- | -------------
[**ConsumptionUsageApiGetConsumptionUsagePerDay**](ConsumptionUsageApiAPI.md#ConsumptionUsageApiGetConsumptionUsagePerDay) | **Get** /2022-09-01-00/resource-instance/usage-per-day | GetConsumptionUsagePerDay consumption-usage-api
[**ConsumptionUsageApiGetCurrentConsumptionUsage**](ConsumptionUsageApiAPI.md#ConsumptionUsageApiGetCurrentConsumptionUsage) | **Get** /2022-09-01-00/resource-instance/usage | GetCurrentConsumptionUsage consumption-usage-api



## ConsumptionUsageApiGetConsumptionUsagePerDay

> GetConsumptionUsageResult ConsumptionUsageApiGetConsumptionUsagePerDay(ctx).StartDate(startDate).EndDate(endDate).SubscriptionID(subscriptionID).Execute()

GetConsumptionUsagePerDay consumption-usage-api

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
	startDate := time.Now() // time.Time | Start date of the usage report (optional)
	endDate := time.Now() // time.Time | End date of the usage report (optional)
	subscriptionID := "sub-12345678" // string | The subscription ID to get usage for (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.ConsumptionUsageApiAPI.ConsumptionUsageApiGetConsumptionUsagePerDay(context.Background()).StartDate(startDate).EndDate(endDate).SubscriptionID(subscriptionID).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `ConsumptionUsageApiAPI.ConsumptionUsageApiGetConsumptionUsagePerDay``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `ConsumptionUsageApiGetConsumptionUsagePerDay`: GetConsumptionUsageResult
	fmt.Fprintf(os.Stdout, "Response from `ConsumptionUsageApiAPI.ConsumptionUsageApiGetConsumptionUsagePerDay`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiConsumptionUsageApiGetConsumptionUsagePerDayRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **startDate** | **time.Time** | Start date of the usage report | 
 **endDate** | **time.Time** | End date of the usage report | 
 **subscriptionID** | **string** | The subscription ID to get usage for | 

### Return type

[**GetConsumptionUsageResult**](GetConsumptionUsageResult.md)

### Authorization

[api_key_header_Authorization](../README.md#api_key_header_Authorization)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json, application/vnd.goa.error

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## ConsumptionUsageApiGetCurrentConsumptionUsage

> GetConsumptionUsageResult ConsumptionUsageApiGetCurrentConsumptionUsage(ctx).SubscriptionID(subscriptionID).Execute()

GetCurrentConsumptionUsage consumption-usage-api

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
	subscriptionID := "sub-12345678" // string | The subscription ID to get usage for (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.ConsumptionUsageApiAPI.ConsumptionUsageApiGetCurrentConsumptionUsage(context.Background()).SubscriptionID(subscriptionID).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `ConsumptionUsageApiAPI.ConsumptionUsageApiGetCurrentConsumptionUsage``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `ConsumptionUsageApiGetCurrentConsumptionUsage`: GetConsumptionUsageResult
	fmt.Fprintf(os.Stdout, "Response from `ConsumptionUsageApiAPI.ConsumptionUsageApiGetCurrentConsumptionUsage`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiConsumptionUsageApiGetCurrentConsumptionUsageRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **subscriptionID** | **string** | The subscription ID to get usage for | 

### Return type

[**GetConsumptionUsageResult**](GetConsumptionUsageResult.md)

### Authorization

[api_key_header_Authorization](../README.md#api_key_header_Authorization)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json, application/vnd.goa.error

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

