# \RegionApiAPI

All URIs are relative to *https://api.omnistrate.cloud*

Method | HTTP request | Description
------------- | ------------- | -------------
[**RegionApiDescribeRegion**](RegionApiAPI.md#RegionApiDescribeRegion) | **Get** /2022-09-01-00/region/{id} | DescribeRegion region-api
[**RegionApiGetRegionByCode**](RegionApiAPI.md#RegionApiGetRegionByCode) | **Get** /2022-09-01-00/region/code/{code}/cloud-provider/{cloudProviderName} | GetRegionByCode region-api
[**RegionApiListRegion**](RegionApiAPI.md#RegionApiListRegion) | **Get** /2022-09-01-00/region/cloudprovider/{cloudProviderName} | ListRegion region-api



## RegionApiDescribeRegion

> DescribeRegionResult RegionApiDescribeRegion(ctx, id).Execute()

DescribeRegion region-api

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
	id := "region-12345678" // string | ID of the Region

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.RegionApiAPI.RegionApiDescribeRegion(context.Background(), id).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `RegionApiAPI.RegionApiDescribeRegion``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `RegionApiDescribeRegion`: DescribeRegionResult
	fmt.Fprintf(os.Stdout, "Response from `RegionApiAPI.RegionApiDescribeRegion`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **string** | ID of the Region | 

### Other Parameters

Other parameters are passed through a pointer to a apiRegionApiDescribeRegionRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

[**DescribeRegionResult**](DescribeRegionResult.md)

### Authorization

[api_key_header_Authorization](../README.md#api_key_header_Authorization)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json, application/vnd.goa.error

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## RegionApiGetRegionByCode

> string RegionApiGetRegionByCode(ctx, code, cloudProviderName).Execute()

GetRegionByCode region-api

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
	code := "us-east-1" // string | Cloud-provider native region code
	cloudProviderName := "aws" // string | The cloud provider for this compute instance type config

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.RegionApiAPI.RegionApiGetRegionByCode(context.Background(), code, cloudProviderName).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `RegionApiAPI.RegionApiGetRegionByCode``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `RegionApiGetRegionByCode`: string
	fmt.Fprintf(os.Stdout, "Response from `RegionApiAPI.RegionApiGetRegionByCode`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**code** | **string** | Cloud-provider native region code | 
**cloudProviderName** | **string** | The cloud provider for this compute instance type config | 

### Other Parameters

Other parameters are passed through a pointer to a apiRegionApiGetRegionByCodeRequest struct via the builder pattern


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


## RegionApiListRegion

> ListRegionsResult RegionApiListRegion(ctx, cloudProviderName).ServiceId(serviceId).ServiceModelId(serviceModelId).ProductTierId(productTierId).ModelType(modelType).Execute()

ListRegion region-api

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
	serviceId := "s-12345678" // string | Service ID. If specified together with serviceModelId,list the regions of the service model. Otherwise list all regions. (optional)
	serviceModelId := "sm-12345678" // string | Service model ID. If specified together with serviceId,list the regions of the service model. Otherwise list all regions. (optional)
	productTierId := "pt-12345678" // string | Product tier ID. If specified together with serviceId,list the regions of the product tier. Otherwise list all regions. (optional)
	modelType := "CUSTOMER_HOSTED" // string | Model type. If specified, list regions for the specified model type. (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.RegionApiAPI.RegionApiListRegion(context.Background(), cloudProviderName).ServiceId(serviceId).ServiceModelId(serviceModelId).ProductTierId(productTierId).ModelType(modelType).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `RegionApiAPI.RegionApiListRegion``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `RegionApiListRegion`: ListRegionsResult
	fmt.Fprintf(os.Stdout, "Response from `RegionApiAPI.RegionApiListRegion`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**cloudProviderName** | **string** | The cloud provider for this compute instance type config | 

### Other Parameters

Other parameters are passed through a pointer to a apiRegionApiListRegionRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **serviceId** | **string** | Service ID. If specified together with serviceModelId,list the regions of the service model. Otherwise list all regions. | 
 **serviceModelId** | **string** | Service model ID. If specified together with serviceId,list the regions of the service model. Otherwise list all regions. | 
 **productTierId** | **string** | Product tier ID. If specified together with serviceId,list the regions of the product tier. Otherwise list all regions. | 
 **modelType** | **string** | Model type. If specified, list regions for the specified model type. | 

### Return type

[**ListRegionsResult**](ListRegionsResult.md)

### Authorization

[api_key_header_Authorization](../README.md#api_key_header_Authorization)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json, application/vnd.goa.error

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

