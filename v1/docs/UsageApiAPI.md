# \UsageApiAPI

All URIs are relative to *https://api.omnistrate.cloud*

Method | HTTP request | Description
------------- | ------------- | -------------
[**UsageApiGetCurrentUsage**](UsageApiAPI.md#UsageApiGetCurrentUsage) | **Get** /2022-09-01-00/usage | GetCurrentUsage usage-api
[**UsageApiGetUsagePerDay**](UsageApiAPI.md#UsageApiGetUsagePerDay) | **Get** /2022-09-01-00/usage-per-day | GetUsagePerDay usage-api



## UsageApiGetCurrentUsage

> GetUsageResult UsageApiGetCurrentUsage(ctx).Execute()

GetCurrentUsage usage-api

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

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.UsageApiAPI.UsageApiGetCurrentUsage(context.Background()).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `UsageApiAPI.UsageApiGetCurrentUsage``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `UsageApiGetCurrentUsage`: GetUsageResult
	fmt.Fprintf(os.Stdout, "Response from `UsageApiAPI.UsageApiGetCurrentUsage`: %v\n", resp)
}
```

### Path Parameters

This endpoint does not need any parameter.

### Other Parameters

Other parameters are passed through a pointer to a apiUsageApiGetCurrentUsageRequest struct via the builder pattern


### Return type

[**GetUsageResult**](GetUsageResult.md)

### Authorization

[api_key_header_Authorization](../README.md#api_key_header_Authorization)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json, application/vnd.goa.error

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## UsageApiGetUsagePerDay

> GetUsageResult UsageApiGetUsagePerDay(ctx).StartDate(startDate).EndDate(endDate).Execute()

GetUsagePerDay usage-api

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

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.UsageApiAPI.UsageApiGetUsagePerDay(context.Background()).StartDate(startDate).EndDate(endDate).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `UsageApiAPI.UsageApiGetUsagePerDay``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `UsageApiGetUsagePerDay`: GetUsageResult
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

### Return type

[**GetUsageResult**](GetUsageResult.md)

### Authorization

[api_key_header_Authorization](../README.md#api_key_header_Authorization)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json, application/vnd.goa.error

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

