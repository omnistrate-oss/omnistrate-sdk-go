# \FleetFeaturesApiAPI

All URIs are relative to *https://api.omnistrate.cloud*

Method | HTTP request | Description
------------- | ------------- | -------------
[**FleetFeaturesApiDescribeFleetFeature**](FleetFeaturesApiAPI.md#FleetFeaturesApiDescribeFleetFeature) | **Get** /2022-09-01-00/fleet/feature/{feature} | DescribeFleetFeature fleet-features-api
[**FleetFeaturesApiDisableFleetFeature**](FleetFeaturesApiAPI.md#FleetFeaturesApiDisableFleetFeature) | **Delete** /2022-09-01-00/fleet/feature | DisableFleetFeature fleet-features-api
[**FleetFeaturesApiEnableFleetFeature**](FleetFeaturesApiAPI.md#FleetFeaturesApiEnableFleetFeature) | **Put** /2022-09-01-00/fleet/feature | EnableFleetFeature fleet-features-api
[**FleetFeaturesApiListFleetFeatures**](FleetFeaturesApiAPI.md#FleetFeaturesApiListFleetFeatures) | **Get** /2022-09-01-00/fleet/features | ListFleetFeatures fleet-features-api



## FleetFeaturesApiDescribeFleetFeature

> FleetFeature FleetFeaturesApiDescribeFleetFeature(ctx, feature).Execute()

DescribeFleetFeature fleet-features-api

### Example

```go
package main

import (
	"context"
	"fmt"
	"os"
	openapiclient "github.com/omnistrate/omnistrate-sdk-go/fleet"
)

func main() {
	feature := "CUSTOMER_BILLING" // string | The feature to describe.

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.FleetFeaturesApiAPI.FleetFeaturesApiDescribeFleetFeature(context.Background(), feature).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `FleetFeaturesApiAPI.FleetFeaturesApiDescribeFleetFeature``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `FleetFeaturesApiDescribeFleetFeature`: FleetFeature
	fmt.Fprintf(os.Stdout, "Response from `FleetFeaturesApiAPI.FleetFeaturesApiDescribeFleetFeature`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**feature** | **string** | The feature to describe. | 

### Other Parameters

Other parameters are passed through a pointer to a apiFleetFeaturesApiDescribeFleetFeatureRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

[**FleetFeature**](FleetFeature.md)

### Authorization

[api_key_header_Authorization](../README.md#api_key_header_Authorization)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json, application/vnd.goa.error

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## FleetFeaturesApiDisableFleetFeature

> FleetFeaturesApiDisableFleetFeature(ctx).EnableFleetFeatureRequestBody(enableFleetFeatureRequestBody).Execute()

DisableFleetFeature fleet-features-api

### Example

```go
package main

import (
	"context"
	"fmt"
	"os"
	openapiclient "github.com/omnistrate/omnistrate-sdk-go/fleet"
)

func main() {
	enableFleetFeatureRequestBody := *openapiclient.NewEnableFleetFeatureRequestBody() // EnableFleetFeatureRequestBody | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	r, err := apiClient.FleetFeaturesApiAPI.FleetFeaturesApiDisableFleetFeature(context.Background()).EnableFleetFeatureRequestBody(enableFleetFeatureRequestBody).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `FleetFeaturesApiAPI.FleetFeaturesApiDisableFleetFeature``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiFleetFeaturesApiDisableFleetFeatureRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **enableFleetFeatureRequestBody** | [**EnableFleetFeatureRequestBody**](EnableFleetFeatureRequestBody.md) |  | 

### Return type

 (empty response body)

### Authorization

[api_key_header_Authorization](../README.md#api_key_header_Authorization)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/vnd.goa.error

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## FleetFeaturesApiEnableFleetFeature

> FleetFeaturesApiEnableFleetFeature(ctx).EnableFleetFeatureRequestBody(enableFleetFeatureRequestBody).Execute()

EnableFleetFeature fleet-features-api

### Example

```go
package main

import (
	"context"
	"fmt"
	"os"
	openapiclient "github.com/omnistrate/omnistrate-sdk-go/fleet"
)

func main() {
	enableFleetFeatureRequestBody := *openapiclient.NewEnableFleetFeatureRequestBody() // EnableFleetFeatureRequestBody | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	r, err := apiClient.FleetFeaturesApiAPI.FleetFeaturesApiEnableFleetFeature(context.Background()).EnableFleetFeatureRequestBody(enableFleetFeatureRequestBody).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `FleetFeaturesApiAPI.FleetFeaturesApiEnableFleetFeature``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiFleetFeaturesApiEnableFleetFeatureRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **enableFleetFeatureRequestBody** | [**EnableFleetFeatureRequestBody**](EnableFleetFeatureRequestBody.md) |  | 

### Return type

 (empty response body)

### Authorization

[api_key_header_Authorization](../README.md#api_key_header_Authorization)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/vnd.goa.error

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## FleetFeaturesApiListFleetFeatures

> ListFleetFeaturesResult FleetFeaturesApiListFleetFeatures(ctx).Execute()

ListFleetFeatures fleet-features-api

### Example

```go
package main

import (
	"context"
	"fmt"
	"os"
	openapiclient "github.com/omnistrate/omnistrate-sdk-go/fleet"
)

func main() {

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.FleetFeaturesApiAPI.FleetFeaturesApiListFleetFeatures(context.Background()).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `FleetFeaturesApiAPI.FleetFeaturesApiListFleetFeatures``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `FleetFeaturesApiListFleetFeatures`: ListFleetFeaturesResult
	fmt.Fprintf(os.Stdout, "Response from `FleetFeaturesApiAPI.FleetFeaturesApiListFleetFeatures`: %v\n", resp)
}
```

### Path Parameters

This endpoint does not need any parameter.

### Other Parameters

Other parameters are passed through a pointer to a apiFleetFeaturesApiListFleetFeaturesRequest struct via the builder pattern


### Return type

[**ListFleetFeaturesResult**](ListFleetFeaturesResult.md)

### Authorization

[api_key_header_Authorization](../README.md#api_key_header_Authorization)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json, application/vnd.goa.error

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)
