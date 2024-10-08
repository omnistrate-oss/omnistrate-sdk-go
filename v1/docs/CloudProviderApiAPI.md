# \CloudProviderApiAPI

All URIs are relative to *https://api.omnistrate.cloud*

Method | HTTP request | Description
------------- | ------------- | -------------
[**CloudProviderApiDescribeCloudProvider**](CloudProviderApiAPI.md#CloudProviderApiDescribeCloudProvider) | **Get** /2022-09-01-00/cloud-provider/{id} | DescribeCloudProvider cloud-provider-api
[**CloudProviderApiGetCloudProviderByName**](CloudProviderApiAPI.md#CloudProviderApiGetCloudProviderByName) | **Get** /2022-09-01-00/cloud-provider/name/{name} | GetCloudProviderByName cloud-provider-api
[**CloudProviderApiListCloudProvider**](CloudProviderApiAPI.md#CloudProviderApiListCloudProvider) | **Get** /2022-09-01-00/cloud-provider | ListCloudProvider cloud-provider-api



## CloudProviderApiDescribeCloudProvider

> DescribeCloudProviderResult CloudProviderApiDescribeCloudProvider(ctx, id).Execute()

DescribeCloudProvider cloud-provider-api

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
	id := "infra-12345678" // string | ID of the CloudProvider

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.CloudProviderApiAPI.CloudProviderApiDescribeCloudProvider(context.Background(), id).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `CloudProviderApiAPI.CloudProviderApiDescribeCloudProvider``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `CloudProviderApiDescribeCloudProvider`: DescribeCloudProviderResult
	fmt.Fprintf(os.Stdout, "Response from `CloudProviderApiAPI.CloudProviderApiDescribeCloudProvider`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **string** | ID of the CloudProvider | 

### Other Parameters

Other parameters are passed through a pointer to a apiCloudProviderApiDescribeCloudProviderRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

[**DescribeCloudProviderResult**](DescribeCloudProviderResult.md)

### Authorization

[api_key_header_Authorization](../README.md#api_key_header_Authorization)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json, application/vnd.goa.error

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## CloudProviderApiGetCloudProviderByName

> string CloudProviderApiGetCloudProviderByName(ctx, name).Execute()

GetCloudProviderByName cloud-provider-api

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
	name := "aws" // string | Name of the CloudProvider

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.CloudProviderApiAPI.CloudProviderApiGetCloudProviderByName(context.Background(), name).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `CloudProviderApiAPI.CloudProviderApiGetCloudProviderByName``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `CloudProviderApiGetCloudProviderByName`: string
	fmt.Fprintf(os.Stdout, "Response from `CloudProviderApiAPI.CloudProviderApiGetCloudProviderByName`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**name** | **string** | Name of the CloudProvider | 

### Other Parameters

Other parameters are passed through a pointer to a apiCloudProviderApiGetCloudProviderByNameRequest struct via the builder pattern


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


## CloudProviderApiListCloudProvider

> ListServiceEnvironmentsResult CloudProviderApiListCloudProvider(ctx).ServiceId(serviceId).ServiceModelId(serviceModelId).ProductTierId(productTierId).Execute()

ListCloudProvider cloud-provider-api

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
	serviceId := "s-12345678" // string | Service ID. If specified together with serviceModelId,list the cloud provider of the service model. Otherwise list all cloud provider. (optional)
	serviceModelId := "sm-12345678" // string | Service model ID. If specified together with serviceId,list the cloud provider of the service model. Otherwise list all cloud provider. (optional)
	productTierId := "pt-12345678" // string | Product tier ID. If specified together with serviceId,list the cloud provider of the product tier. Otherwise list all cloud provider. (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.CloudProviderApiAPI.CloudProviderApiListCloudProvider(context.Background()).ServiceId(serviceId).ServiceModelId(serviceModelId).ProductTierId(productTierId).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `CloudProviderApiAPI.CloudProviderApiListCloudProvider``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `CloudProviderApiListCloudProvider`: ListServiceEnvironmentsResult
	fmt.Fprintf(os.Stdout, "Response from `CloudProviderApiAPI.CloudProviderApiListCloudProvider`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiCloudProviderApiListCloudProviderRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **serviceId** | **string** | Service ID. If specified together with serviceModelId,list the cloud provider of the service model. Otherwise list all cloud provider. | 
 **serviceModelId** | **string** | Service model ID. If specified together with serviceId,list the cloud provider of the service model. Otherwise list all cloud provider. | 
 **productTierId** | **string** | Product tier ID. If specified together with serviceId,list the cloud provider of the product tier. Otherwise list all cloud provider. | 

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

