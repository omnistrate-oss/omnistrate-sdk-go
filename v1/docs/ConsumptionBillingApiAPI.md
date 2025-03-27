# \ConsumptionBillingApiAPI

All URIs are relative to *https://api.omnistrate.cloud*

Method | HTTP request | Description
------------- | ------------- | -------------
[**ConsumptionBillingApiDescribeConsumptionBillingDetails**](ConsumptionBillingApiAPI.md#ConsumptionBillingApiDescribeConsumptionBillingDetails) | **Get** /2022-09-01-00/resource-instance/billing-details | DescribeConsumptionBillingDetails consumption-billing-api
[**ConsumptionBillingApiDescribeConsumptionBillingStatus**](ConsumptionBillingApiAPI.md#ConsumptionBillingApiDescribeConsumptionBillingStatus) | **Get** /2022-09-01-00/resource-instance/billing-status | DescribeConsumptionBillingStatus consumption-billing-api



## ConsumptionBillingApiDescribeConsumptionBillingDetails

> DescribeConsumptionBillingDetailsResult ConsumptionBillingApiDescribeConsumptionBillingDetails(ctx).ReturnUrl(returnUrl).Execute()

DescribeConsumptionBillingDetails consumption-billing-api

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
	returnUrl := "https://mysaasportal.com" // string | Return Url used to configure payment methods links (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.ConsumptionBillingApiAPI.ConsumptionBillingApiDescribeConsumptionBillingDetails(context.Background()).ReturnUrl(returnUrl).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `ConsumptionBillingApiAPI.ConsumptionBillingApiDescribeConsumptionBillingDetails``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `ConsumptionBillingApiDescribeConsumptionBillingDetails`: DescribeConsumptionBillingDetailsResult
	fmt.Fprintf(os.Stdout, "Response from `ConsumptionBillingApiAPI.ConsumptionBillingApiDescribeConsumptionBillingDetails`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiConsumptionBillingApiDescribeConsumptionBillingDetailsRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **returnUrl** | **string** | Return Url used to configure payment methods links | 

### Return type

[**DescribeConsumptionBillingDetailsResult**](DescribeConsumptionBillingDetailsResult.md)

### Authorization

[api_key_header_Authorization](../README.md#api_key_header_Authorization)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json, application/vnd.goa.error

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## ConsumptionBillingApiDescribeConsumptionBillingStatus

> DescribeConsumptionBillingStatusResult ConsumptionBillingApiDescribeConsumptionBillingStatus(ctx).Execute()

DescribeConsumptionBillingStatus consumption-billing-api

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
	resp, r, err := apiClient.ConsumptionBillingApiAPI.ConsumptionBillingApiDescribeConsumptionBillingStatus(context.Background()).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `ConsumptionBillingApiAPI.ConsumptionBillingApiDescribeConsumptionBillingStatus``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `ConsumptionBillingApiDescribeConsumptionBillingStatus`: DescribeConsumptionBillingStatusResult
	fmt.Fprintf(os.Stdout, "Response from `ConsumptionBillingApiAPI.ConsumptionBillingApiDescribeConsumptionBillingStatus`: %v\n", resp)
}
```

### Path Parameters

This endpoint does not need any parameter.

### Other Parameters

Other parameters are passed through a pointer to a apiConsumptionBillingApiDescribeConsumptionBillingStatusRequest struct via the builder pattern


### Return type

[**DescribeConsumptionBillingStatusResult**](DescribeConsumptionBillingStatusResult.md)

### Authorization

[api_key_header_Authorization](../README.md#api_key_header_Authorization)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json, application/vnd.goa.error

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

