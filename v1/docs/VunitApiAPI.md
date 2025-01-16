# \VunitApiAPI

All URIs are relative to *https://api.omnistrate.cloud*

Method | HTTP request | Description
------------- | ------------- | -------------
[**VunitApiDescribeNetwork**](VunitApiAPI.md#VunitApiDescribeNetwork) | **Get** /2022-09-01-00/service/{serviceId}/network/{id} | DescribeNetwork vunit-api
[**VunitApiDescribeVUnit**](VunitApiAPI.md#VunitApiDescribeVUnit) | **Get** /2022-09-01-00/service/{serviceId}/vunit/{id} | DescribeVUnit vunit-api
[**VunitApiListVUnits**](VunitApiAPI.md#VunitApiListVUnits) | **Get** /2022-09-01-00/service/{serviceId}/model/{serviceModelId}/vunit | ListVUnits vunit-api



## VunitApiDescribeNetwork

> DescribeNetworkResult VunitApiDescribeNetwork(ctx, serviceId, id).Execute()

DescribeNetwork vunit-api

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
	serviceId := "s-12345678" // string | Service ID context
	id := "n-12345678" // string | Network to operate on

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.VunitApiAPI.VunitApiDescribeNetwork(context.Background(), serviceId, id).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `VunitApiAPI.VunitApiDescribeNetwork``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `VunitApiDescribeNetwork`: DescribeNetworkResult
	fmt.Fprintf(os.Stdout, "Response from `VunitApiAPI.VunitApiDescribeNetwork`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**serviceId** | **string** | Service ID context | 
**id** | **string** | Network to operate on | 

### Other Parameters

Other parameters are passed through a pointer to a apiVunitApiDescribeNetworkRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------



### Return type

[**DescribeNetworkResult**](DescribeNetworkResult.md)

### Authorization

[api_key_header_Authorization](../README.md#api_key_header_Authorization)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json, application/vnd.goa.error

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## VunitApiDescribeVUnit

> DescribeVUnitResult VunitApiDescribeVUnit(ctx, serviceId, id).Execute()

DescribeVUnit vunit-api

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
	serviceId := "s-12345678" // string | Service ID for the VUnit
	id := "vu-12345678" // string | VUnit to operate on

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.VunitApiAPI.VunitApiDescribeVUnit(context.Background(), serviceId, id).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `VunitApiAPI.VunitApiDescribeVUnit``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `VunitApiDescribeVUnit`: DescribeVUnitResult
	fmt.Fprintf(os.Stdout, "Response from `VunitApiAPI.VunitApiDescribeVUnit`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**serviceId** | **string** | Service ID for the VUnit | 
**id** | **string** | VUnit to operate on | 

### Other Parameters

Other parameters are passed through a pointer to a apiVunitApiDescribeVUnitRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------



### Return type

[**DescribeVUnitResult**](DescribeVUnitResult.md)

### Authorization

[api_key_header_Authorization](../README.md#api_key_header_Authorization)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json, application/vnd.goa.error

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## VunitApiListVUnits

> ListVUnitsResult VunitApiListVUnits(ctx, serviceId, serviceModelId).ListVUnitsRequest2(listVUnitsRequest2).Execute()

ListVUnits vunit-api

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
	serviceId := "s-12345678" // string | Service ID for the VUnit
	serviceModelId := "sm-12345678" // string | Service Model ID for the VUnit
	listVUnitsRequest2 := *openapiclient.NewListVUnitsRequest2("aws", "us-west-2") // ListVUnitsRequest2 | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.VunitApiAPI.VunitApiListVUnits(context.Background(), serviceId, serviceModelId).ListVUnitsRequest2(listVUnitsRequest2).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `VunitApiAPI.VunitApiListVUnits``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `VunitApiListVUnits`: ListVUnitsResult
	fmt.Fprintf(os.Stdout, "Response from `VunitApiAPI.VunitApiListVUnits`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**serviceId** | **string** | Service ID for the VUnit | 
**serviceModelId** | **string** | Service Model ID for the VUnit | 

### Other Parameters

Other parameters are passed through a pointer to a apiVunitApiListVUnitsRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


 **listVUnitsRequest2** | [**ListVUnitsRequest2**](ListVUnitsRequest2.md) |  | 

### Return type

[**ListVUnitsResult**](ListVUnitsResult.md)

### Authorization

[api_key_header_Authorization](../README.md#api_key_header_Authorization)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json, application/vnd.goa.error

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

