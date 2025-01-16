# \ImageRegistryApiAPI

All URIs are relative to *https://api.omnistrate.cloud*

Method | HTTP request | Description
------------- | ------------- | -------------
[**ImageRegistryApiCreateImageRegistry**](ImageRegistryApiAPI.md#ImageRegistryApiCreateImageRegistry) | **Post** /2022-09-01-00/image-registry | CreateImageRegistry image-registry-api
[**ImageRegistryApiDeleteImageRegistry**](ImageRegistryApiAPI.md#ImageRegistryApiDeleteImageRegistry) | **Delete** /2022-09-01-00/image-registry/{id} | DeleteImageRegistry image-registry-api
[**ImageRegistryApiDescribeImageRegistry**](ImageRegistryApiAPI.md#ImageRegistryApiDescribeImageRegistry) | **Get** /2022-09-01-00/image-registry/{id} | DescribeImageRegistry image-registry-api
[**ImageRegistryApiListImageRegistry**](ImageRegistryApiAPI.md#ImageRegistryApiListImageRegistry) | **Get** /2022-09-01-00/image-registry | ListImageRegistry image-registry-api
[**ImageRegistryApiUpdateImageRegistry**](ImageRegistryApiAPI.md#ImageRegistryApiUpdateImageRegistry) | **Patch** /2022-09-01-00/image-registry/{id} | UpdateImageRegistry image-registry-api



## ImageRegistryApiCreateImageRegistry

> string ImageRegistryApiCreateImageRegistry(ctx).CreateImageRegistryRequest2(createImageRegistryRequest2).Execute()

CreateImageRegistry image-registry-api

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
	createImageRegistryRequest2 := *openapiclient.NewCreateImageRegistryRequest2("DockerHub is a public Docker Image Registry", "docker.io", "DockerHub") // CreateImageRegistryRequest2 | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.ImageRegistryApiAPI.ImageRegistryApiCreateImageRegistry(context.Background()).CreateImageRegistryRequest2(createImageRegistryRequest2).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `ImageRegistryApiAPI.ImageRegistryApiCreateImageRegistry``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `ImageRegistryApiCreateImageRegistry`: string
	fmt.Fprintf(os.Stdout, "Response from `ImageRegistryApiAPI.ImageRegistryApiCreateImageRegistry`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiImageRegistryApiCreateImageRegistryRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **createImageRegistryRequest2** | [**CreateImageRegistryRequest2**](CreateImageRegistryRequest2.md) |  | 

### Return type

**string**

### Authorization

[api_key_header_Authorization](../README.md#api_key_header_Authorization)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json, application/vnd.goa.error

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## ImageRegistryApiDeleteImageRegistry

> ImageRegistryApiDeleteImageRegistry(ctx, id).Execute()

DeleteImageRegistry image-registry-api

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
	id := "ir-12345678" // string | The ID of the Image Registry

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	r, err := apiClient.ImageRegistryApiAPI.ImageRegistryApiDeleteImageRegistry(context.Background(), id).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `ImageRegistryApiAPI.ImageRegistryApiDeleteImageRegistry``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **string** | The ID of the Image Registry | 

### Other Parameters

Other parameters are passed through a pointer to a apiImageRegistryApiDeleteImageRegistryRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

 (empty response body)

### Authorization

[api_key_header_Authorization](../README.md#api_key_header_Authorization)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/vnd.goa.error

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## ImageRegistryApiDescribeImageRegistry

> DescribeImageRegistryResult ImageRegistryApiDescribeImageRegistry(ctx, id).Execute()

DescribeImageRegistry image-registry-api

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
	id := "ir-12345678" // string | The ID of the Image Registry

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.ImageRegistryApiAPI.ImageRegistryApiDescribeImageRegistry(context.Background(), id).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `ImageRegistryApiAPI.ImageRegistryApiDescribeImageRegistry``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `ImageRegistryApiDescribeImageRegistry`: DescribeImageRegistryResult
	fmt.Fprintf(os.Stdout, "Response from `ImageRegistryApiAPI.ImageRegistryApiDescribeImageRegistry`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **string** | The ID of the Image Registry | 

### Other Parameters

Other parameters are passed through a pointer to a apiImageRegistryApiDescribeImageRegistryRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

[**DescribeImageRegistryResult**](DescribeImageRegistryResult.md)

### Authorization

[api_key_header_Authorization](../README.md#api_key_header_Authorization)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json, application/vnd.goa.error

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## ImageRegistryApiListImageRegistry

> ListImageRegistryResult ImageRegistryApiListImageRegistry(ctx).Execute()

ListImageRegistry image-registry-api

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
	resp, r, err := apiClient.ImageRegistryApiAPI.ImageRegistryApiListImageRegistry(context.Background()).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `ImageRegistryApiAPI.ImageRegistryApiListImageRegistry``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `ImageRegistryApiListImageRegistry`: ListImageRegistryResult
	fmt.Fprintf(os.Stdout, "Response from `ImageRegistryApiAPI.ImageRegistryApiListImageRegistry`: %v\n", resp)
}
```

### Path Parameters

This endpoint does not need any parameter.

### Other Parameters

Other parameters are passed through a pointer to a apiImageRegistryApiListImageRegistryRequest struct via the builder pattern


### Return type

[**ListImageRegistryResult**](ListImageRegistryResult.md)

### Authorization

[api_key_header_Authorization](../README.md#api_key_header_Authorization)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json, application/vnd.goa.error

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## ImageRegistryApiUpdateImageRegistry

> ImageRegistryApiUpdateImageRegistry(ctx, id).UpdateImageRegistryRequest2(updateImageRegistryRequest2).Execute()

UpdateImageRegistry image-registry-api

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
	id := "ir-12345678" // string | The ID of the Image Registry
	updateImageRegistryRequest2 := *openapiclient.NewUpdateImageRegistryRequest2() // UpdateImageRegistryRequest2 | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	r, err := apiClient.ImageRegistryApiAPI.ImageRegistryApiUpdateImageRegistry(context.Background(), id).UpdateImageRegistryRequest2(updateImageRegistryRequest2).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `ImageRegistryApiAPI.ImageRegistryApiUpdateImageRegistry``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **string** | The ID of the Image Registry | 

### Other Parameters

Other parameters are passed through a pointer to a apiImageRegistryApiUpdateImageRegistryRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **updateImageRegistryRequest2** | [**UpdateImageRegistryRequest2**](UpdateImageRegistryRequest2.md) |  | 

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

