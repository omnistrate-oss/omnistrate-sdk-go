# \ReportApiAPI

All URIs are relative to *https://api.omnistrate.cloud*

Method | HTTP request | Description
------------- | ------------- | -------------
[**ReportApiDescribeEnvironmentReportStats**](ReportApiAPI.md#ReportApiDescribeEnvironmentReportStats) | **Get** /2022-09-01-00/fleet/report/stats/env/{environmentType} | DescribeEnvironmentReportStats report-api



## ReportApiDescribeEnvironmentReportStats

> EnvironmentReportStatsSummary ReportApiDescribeEnvironmentReportStats(ctx, environmentType).Execute()

DescribeEnvironmentReportStats report-api

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
	environmentType := "PROD" // string | The environment type to retrieve report statistics for

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.ReportApiAPI.ReportApiDescribeEnvironmentReportStats(context.Background(), environmentType).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `ReportApiAPI.ReportApiDescribeEnvironmentReportStats``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `ReportApiDescribeEnvironmentReportStats`: EnvironmentReportStatsSummary
	fmt.Fprintf(os.Stdout, "Response from `ReportApiAPI.ReportApiDescribeEnvironmentReportStats`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**environmentType** | **string** | The environment type to retrieve report statistics for | 

### Other Parameters

Other parameters are passed through a pointer to a apiReportApiDescribeEnvironmentReportStatsRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

[**EnvironmentReportStatsSummary**](EnvironmentReportStatsSummary.md)

### Authorization

[api_key_header_Authorization](../README.md#api_key_header_Authorization)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json, application/vnd.goa.error

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

