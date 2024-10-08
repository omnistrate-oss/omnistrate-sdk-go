# \ServiceOfferingApiAPI

All URIs are relative to *https://api.omnistrate.cloud*

Method | HTTP request | Description
------------- | ------------- | -------------
[**ServiceOfferingApiDescribeServiceOffering**](ServiceOfferingApiAPI.md#ServiceOfferingApiDescribeServiceOffering) | **Get** /2022-09-01-00/service-offering/{serviceId} | DescribeServiceOffering service-offering-api
[**ServiceOfferingApiDescribeServiceOfferingResource**](ServiceOfferingApiAPI.md#ServiceOfferingApiDescribeServiceOfferingResource) | **Get** /2022-09-01-00/service-offering/{serviceId}/resource/{resourceId}/instance/{instanceId} | DescribeServiceOfferingResource service-offering-api
[**ServiceOfferingApiListServiceOffering**](ServiceOfferingApiAPI.md#ServiceOfferingApiListServiceOffering) | **Get** /2022-09-01-00/service-offering | ListServiceOffering service-offering-api



## ServiceOfferingApiDescribeServiceOffering

> DescribeServiceOfferingResult ServiceOfferingApiDescribeServiceOffering(ctx, serviceId).Visibility(visibility).EnvironmentType(environmentType).Execute()

DescribeServiceOffering service-offering-api

### Example

```go
package main

import (
	"context"
	"fmt"
	"os"
	openapiclient "github.com/omnistrate/omnistrate-sdk-go/v1"
)

func main() {
	serviceId := "s-12345678" // string | The service ID
	visibility := "PRIVATE" // string | The visibility of service offering (optional)
	environmentType := "DEV" // string | The environment type to filter by (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.ServiceOfferingApiAPI.ServiceOfferingApiDescribeServiceOffering(context.Background(), serviceId).Visibility(visibility).EnvironmentType(environmentType).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `ServiceOfferingApiAPI.ServiceOfferingApiDescribeServiceOffering``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `ServiceOfferingApiDescribeServiceOffering`: DescribeServiceOfferingResult
	fmt.Fprintf(os.Stdout, "Response from `ServiceOfferingApiAPI.ServiceOfferingApiDescribeServiceOffering`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**serviceId** | **string** | The service ID | 

### Other Parameters

Other parameters are passed through a pointer to a apiServiceOfferingApiDescribeServiceOfferingRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **visibility** | **string** | The visibility of service offering | 
 **environmentType** | **string** | The environment type to filter by | 

### Return type

[**DescribeServiceOfferingResult**](DescribeServiceOfferingResult.md)

### Authorization

[api_key_header_Authorization](../README.md#api_key_header_Authorization)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json, application/vnd.goa.error

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## ServiceOfferingApiDescribeServiceOfferingResource

> DescribeServiceOfferingResourceResult ServiceOfferingApiDescribeServiceOfferingResource(ctx, serviceId, resourceId, instanceId).Execute()

DescribeServiceOfferingResource service-offering-api

### Example

```go
package main

import (
	"context"
	"fmt"
	"os"
	openapiclient "github.com/omnistrate/omnistrate-sdk-go/v1"
)

func main() {
	serviceId := "s-12345678" // string | The service ID
	resourceId := "r-12345678" // string | The resource ID
	instanceId := "instance-12345678" // string | The instance ID (default to "none")

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.ServiceOfferingApiAPI.ServiceOfferingApiDescribeServiceOfferingResource(context.Background(), serviceId, resourceId, instanceId).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `ServiceOfferingApiAPI.ServiceOfferingApiDescribeServiceOfferingResource``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `ServiceOfferingApiDescribeServiceOfferingResource`: DescribeServiceOfferingResourceResult
	fmt.Fprintf(os.Stdout, "Response from `ServiceOfferingApiAPI.ServiceOfferingApiDescribeServiceOfferingResource`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**serviceId** | **string** | The service ID | 
**resourceId** | **string** | The resource ID | 
**instanceId** | **string** | The instance ID | [default to &quot;none&quot;]

### Other Parameters

Other parameters are passed through a pointer to a apiServiceOfferingApiDescribeServiceOfferingResourceRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------




### Return type

[**DescribeServiceOfferingResourceResult**](DescribeServiceOfferingResourceResult.md)

### Authorization

[api_key_header_Authorization](../README.md#api_key_header_Authorization)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json, application/vnd.goa.error

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## ServiceOfferingApiListServiceOffering

> ListServiceOfferingsResult ServiceOfferingApiListServiceOffering(ctx).OrgId(orgId).Visibility(visibility).EnvironmentType(environmentType).Execute()

ListServiceOffering service-offering-api

### Example

```go
package main

import (
	"context"
	"fmt"
	"os"
	openapiclient "github.com/omnistrate/omnistrate-sdk-go/v1"
)

func main() {
	orgId := "org-12345678" // string | Org Id (optional)
	visibility := "PRIVATE" // string | The visibility of service offering (optional)
	environmentType := "DEV" // string | The environment type to filter by (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.ServiceOfferingApiAPI.ServiceOfferingApiListServiceOffering(context.Background()).OrgId(orgId).Visibility(visibility).EnvironmentType(environmentType).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `ServiceOfferingApiAPI.ServiceOfferingApiListServiceOffering``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `ServiceOfferingApiListServiceOffering`: ListServiceOfferingsResult
	fmt.Fprintf(os.Stdout, "Response from `ServiceOfferingApiAPI.ServiceOfferingApiListServiceOffering`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiServiceOfferingApiListServiceOfferingRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **orgId** | **string** | Org Id | 
 **visibility** | **string** | The visibility of service offering | 
 **environmentType** | **string** | The environment type to filter by | 

### Return type

[**ListServiceOfferingsResult**](ListServiceOfferingsResult.md)

### Authorization

[api_key_header_Authorization](../README.md#api_key_header_Authorization)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json, application/vnd.goa.error

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

