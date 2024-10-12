# \AvailabilityZoneApiAPI

All URIs are relative to *https://api.omnistrate.cloud*

Method | HTTP request | Description
------------- | ------------- | -------------
[**AvailabilityZoneApiDescribeAvailabilityZone**](AvailabilityZoneApiAPI.md#AvailabilityZoneApiDescribeAvailabilityZone) | **Get** /2022-09-01-00/availability-zone/{id} | DescribeAvailabilityZone availability-zone-api
[**AvailabilityZoneApiGetAvailabilityZoneByCode**](AvailabilityZoneApiAPI.md#AvailabilityZoneApiGetAvailabilityZoneByCode) | **Get** /2022-09-01-00/availability-zone/code/{code}/cloud-provider/{cloudProviderName} | GetAvailabilityZoneByCode availability-zone-api
[**AvailabilityZoneApiListAvailabilityZone**](AvailabilityZoneApiAPI.md#AvailabilityZoneApiListAvailabilityZone) | **Get** /2022-09-01-00/availability-zone/cloud-provider/{cloudProviderName} | ListAvailabilityZone availability-zone-api
[**AvailabilityZoneApiListAvailabilityZonesByRegionCode**](AvailabilityZoneApiAPI.md#AvailabilityZoneApiListAvailabilityZonesByRegionCode) | **Get** /2022-09-01-00/availability-zone/region/code/{regionCode}/cloud-provider/{cloudProviderName} | ListAvailabilityZonesByRegionCode availability-zone-api



## AvailabilityZoneApiDescribeAvailabilityZone

> DescribeAvailabilityZoneResult AvailabilityZoneApiDescribeAvailabilityZone(ctx, id).Execute()

DescribeAvailabilityZone availability-zone-api

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
	id := "az-12345678" // string | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.AvailabilityZoneApiAPI.AvailabilityZoneApiDescribeAvailabilityZone(context.Background(), id).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `AvailabilityZoneApiAPI.AvailabilityZoneApiDescribeAvailabilityZone``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `AvailabilityZoneApiDescribeAvailabilityZone`: DescribeAvailabilityZoneResult
	fmt.Fprintf(os.Stdout, "Response from `AvailabilityZoneApiAPI.AvailabilityZoneApiDescribeAvailabilityZone`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiAvailabilityZoneApiDescribeAvailabilityZoneRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

[**DescribeAvailabilityZoneResult**](DescribeAvailabilityZoneResult.md)

### Authorization

[api_key_header_Authorization](../README.md#api_key_header_Authorization)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json, application/vnd.goa.error

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## AvailabilityZoneApiGetAvailabilityZoneByCode

> string AvailabilityZoneApiGetAvailabilityZoneByCode(ctx, code, cloudProviderName).Execute()

GetAvailabilityZoneByCode availability-zone-api

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
	code := "us-east-1a" // string | Cloud-provider native availability zone code
	cloudProviderName := "aws" // string | The cloud provider for this compute instance type config

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.AvailabilityZoneApiAPI.AvailabilityZoneApiGetAvailabilityZoneByCode(context.Background(), code, cloudProviderName).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `AvailabilityZoneApiAPI.AvailabilityZoneApiGetAvailabilityZoneByCode``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `AvailabilityZoneApiGetAvailabilityZoneByCode`: string
	fmt.Fprintf(os.Stdout, "Response from `AvailabilityZoneApiAPI.AvailabilityZoneApiGetAvailabilityZoneByCode`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**code** | **string** | Cloud-provider native availability zone code | 
**cloudProviderName** | **string** | The cloud provider for this compute instance type config | 

### Other Parameters

Other parameters are passed through a pointer to a apiAvailabilityZoneApiGetAvailabilityZoneByCodeRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------



### Return type

**string**

### Authorization

[api_key_header_Authorization](../README.md#api_key_header_Authorization)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json, application/vnd.goa.error

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## AvailabilityZoneApiListAvailabilityZone

> ListServiceEnvironmentsResult AvailabilityZoneApiListAvailabilityZone(ctx, cloudProviderName).Execute()

ListAvailabilityZone availability-zone-api

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
	cloudProviderName := "aws" // string | The cloud provider for this compute instance type config

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.AvailabilityZoneApiAPI.AvailabilityZoneApiListAvailabilityZone(context.Background(), cloudProviderName).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `AvailabilityZoneApiAPI.AvailabilityZoneApiListAvailabilityZone``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `AvailabilityZoneApiListAvailabilityZone`: ListServiceEnvironmentsResult
	fmt.Fprintf(os.Stdout, "Response from `AvailabilityZoneApiAPI.AvailabilityZoneApiListAvailabilityZone`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**cloudProviderName** | **string** | The cloud provider for this compute instance type config | 

### Other Parameters

Other parameters are passed through a pointer to a apiAvailabilityZoneApiListAvailabilityZoneRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

[**ListServiceEnvironmentsResult**](ListServiceEnvironmentsResult.md)

### Authorization

[api_key_header_Authorization](../README.md#api_key_header_Authorization)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json, application/vnd.goa.error

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## AvailabilityZoneApiListAvailabilityZonesByRegionCode

> ListServiceEnvironmentsResult AvailabilityZoneApiListAvailabilityZonesByRegionCode(ctx, regionCode, cloudProviderName).Execute()

ListAvailabilityZonesByRegionCode availability-zone-api

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
	regionCode := "us-east-1" // string | 
	cloudProviderName := "aws" // string | The cloud provider for this compute instance type config

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.AvailabilityZoneApiAPI.AvailabilityZoneApiListAvailabilityZonesByRegionCode(context.Background(), regionCode, cloudProviderName).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `AvailabilityZoneApiAPI.AvailabilityZoneApiListAvailabilityZonesByRegionCode``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `AvailabilityZoneApiListAvailabilityZonesByRegionCode`: ListServiceEnvironmentsResult
	fmt.Fprintf(os.Stdout, "Response from `AvailabilityZoneApiAPI.AvailabilityZoneApiListAvailabilityZonesByRegionCode`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**regionCode** | **string** |  | 
**cloudProviderName** | **string** | The cloud provider for this compute instance type config | 

### Other Parameters

Other parameters are passed through a pointer to a apiAvailabilityZoneApiListAvailabilityZonesByRegionCodeRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------



### Return type

[**ListServiceEnvironmentsResult**](ListServiceEnvironmentsResult.md)

### Authorization

[api_key_header_Authorization](../README.md#api_key_header_Authorization)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json, application/vnd.goa.error

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

