# \ImageConfigApiAPI

All URIs are relative to *https://api.omnistrate.cloud*

Method | HTTP request | Description
------------- | ------------- | -------------
[**ImageConfigApiCreateImageConfig**](ImageConfigApiAPI.md#ImageConfigApiCreateImageConfig) | **Post** /2022-09-01-00/service/{serviceId}/image-config | CreateImageConfig image-config-api
[**ImageConfigApiDeleteImageConfig**](ImageConfigApiAPI.md#ImageConfigApiDeleteImageConfig) | **Delete** /2022-09-01-00/service/{serviceId}/image-config/{id} | DeleteImageConfig image-config-api
[**ImageConfigApiDescribeImageConfig**](ImageConfigApiAPI.md#ImageConfigApiDescribeImageConfig) | **Get** /2022-09-01-00/service/{serviceId}/image-config/{id} | DescribeImageConfig image-config-api
[**ImageConfigApiListImageConfigs**](ImageConfigApiAPI.md#ImageConfigApiListImageConfigs) | **Get** /2022-09-01-00/service/{serviceId}/serviceenvironment/{serviceEnvironmentId}/image-config | ListImageConfigs image-config-api
[**ImageConfigApiReleaseImageConfig**](ImageConfigApiAPI.md#ImageConfigApiReleaseImageConfig) | **Post** /2022-09-01-00/service/{serviceId}/image-config/{id}/release | ReleaseImageConfig image-config-api
[**ImageConfigApiRolloutFleetImageConfig**](ImageConfigApiAPI.md#ImageConfigApiRolloutFleetImageConfig) | **Post** /2022-09-01-00/service/{serviceId}/image-config/{id}/rollout | RolloutFleetImageConfig image-config-api
[**ImageConfigApiRolloutFleetImageStatus**](ImageConfigApiAPI.md#ImageConfigApiRolloutFleetImageStatus) | **Get** /2022-09-01-00/service/{serviceId}/image-config/{id}/rollout | RolloutFleetImageStatus image-config-api
[**ImageConfigApiUpdateImageConfig**](ImageConfigApiAPI.md#ImageConfigApiUpdateImageConfig) | **Patch** /2022-09-01-00/service/{serviceId}/image-config/{id} | UpdateImageConfig image-config-api



## ImageConfigApiCreateImageConfig

> string ImageConfigApiCreateImageConfig(ctx, serviceId).CreateImageConfigRequest2(createImageConfigRequest2).Execute()

CreateImageConfig image-config-api

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
	serviceId := "s-12345678" // string | The service ID to use for the infra
	createImageConfigRequest2 := *openapiclient.NewCreateImageConfigRequest2("A image configuration for my new entity", "mysql", "se-123456") // CreateImageConfigRequest2 | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.ImageConfigApiAPI.ImageConfigApiCreateImageConfig(context.Background(), serviceId).CreateImageConfigRequest2(createImageConfigRequest2).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `ImageConfigApiAPI.ImageConfigApiCreateImageConfig``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `ImageConfigApiCreateImageConfig`: string
	fmt.Fprintf(os.Stdout, "Response from `ImageConfigApiAPI.ImageConfigApiCreateImageConfig`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**serviceId** | **string** | The service ID to use for the infra | 

### Other Parameters

Other parameters are passed through a pointer to a apiImageConfigApiCreateImageConfigRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **createImageConfigRequest2** | [**CreateImageConfigRequest2**](CreateImageConfigRequest2.md) |  | 

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


## ImageConfigApiDeleteImageConfig

> ImageConfigApiDeleteImageConfig(ctx, serviceId, id).Execute()

DeleteImageConfig image-config-api

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
	serviceId := "s-12345678" // string | The service ID
	id := "imgc-12345678" // string | The image configuration ID

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	r, err := apiClient.ImageConfigApiAPI.ImageConfigApiDeleteImageConfig(context.Background(), serviceId, id).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `ImageConfigApiAPI.ImageConfigApiDeleteImageConfig``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**serviceId** | **string** | The service ID | 
**id** | **string** | The image configuration ID | 

### Other Parameters

Other parameters are passed through a pointer to a apiImageConfigApiDeleteImageConfigRequest struct via the builder pattern


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


## ImageConfigApiDescribeImageConfig

> DescribeImageConfigResult ImageConfigApiDescribeImageConfig(ctx, serviceId, id).ProductTierVersion(productTierVersion).ProductTierId(productTierId).Execute()

DescribeImageConfig image-config-api

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
	serviceId := "s-12345678" // string | The service ID
	id := "imgc-12345678" // string | The image configuration ID
	productTierVersion := "Aut iste doloremque minima fuga assumenda numquam." // string | Product tier version of the image config to describe. If not specified, the latest version is described. (optional)
	productTierId := "Beatae beatae." // string | ProductTierId of the image config to describe. Needs to specified in combination with the product tier version (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.ImageConfigApiAPI.ImageConfigApiDescribeImageConfig(context.Background(), serviceId, id).ProductTierVersion(productTierVersion).ProductTierId(productTierId).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `ImageConfigApiAPI.ImageConfigApiDescribeImageConfig``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `ImageConfigApiDescribeImageConfig`: DescribeImageConfigResult
	fmt.Fprintf(os.Stdout, "Response from `ImageConfigApiAPI.ImageConfigApiDescribeImageConfig`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**serviceId** | **string** | The service ID | 
**id** | **string** | The image configuration ID | 

### Other Parameters

Other parameters are passed through a pointer to a apiImageConfigApiDescribeImageConfigRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


 **productTierVersion** | **string** | Product tier version of the image config to describe. If not specified, the latest version is described. | 
 **productTierId** | **string** | ProductTierId of the image config to describe. Needs to specified in combination with the product tier version | 

### Return type

[**DescribeImageConfigResult**](DescribeImageConfigResult.md)

### Authorization

[api_key_header_Authorization](../README.md#api_key_header_Authorization)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json, application/vnd.goa.error

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## ImageConfigApiListImageConfigs

> ListImageConfigsResult ImageConfigApiListImageConfigs(ctx, serviceId, serviceEnvironmentId).Execute()

ListImageConfigs image-config-api

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
	serviceId := "s-12345678" // string | The service ID to use for the infra
	serviceEnvironmentId := "se-12345678" // string | The service environment ID

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.ImageConfigApiAPI.ImageConfigApiListImageConfigs(context.Background(), serviceId, serviceEnvironmentId).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `ImageConfigApiAPI.ImageConfigApiListImageConfigs``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `ImageConfigApiListImageConfigs`: ListImageConfigsResult
	fmt.Fprintf(os.Stdout, "Response from `ImageConfigApiAPI.ImageConfigApiListImageConfigs`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**serviceId** | **string** | The service ID to use for the infra | 
**serviceEnvironmentId** | **string** | The service environment ID | 

### Other Parameters

Other parameters are passed through a pointer to a apiImageConfigApiListImageConfigsRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------



### Return type

[**ListImageConfigsResult**](ListImageConfigsResult.md)

### Authorization

[api_key_header_Authorization](../README.md#api_key_header_Authorization)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json, application/vnd.goa.error

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## ImageConfigApiReleaseImageConfig

> ImageConfigApiReleaseImageConfig(ctx, serviceId, id).ReleaseImageConfigRequest2(releaseImageConfigRequest2).Execute()

ReleaseImageConfig image-config-api

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
	serviceId := "s-12345678" // string | The service ID
	id := "imgc-12345678" // string | The image configuration ID
	releaseImageConfigRequest2 := *openapiclient.NewReleaseImageConfigRequest2() // ReleaseImageConfigRequest2 | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	r, err := apiClient.ImageConfigApiAPI.ImageConfigApiReleaseImageConfig(context.Background(), serviceId, id).ReleaseImageConfigRequest2(releaseImageConfigRequest2).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `ImageConfigApiAPI.ImageConfigApiReleaseImageConfig``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**serviceId** | **string** | The service ID | 
**id** | **string** | The image configuration ID | 

### Other Parameters

Other parameters are passed through a pointer to a apiImageConfigApiReleaseImageConfigRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


 **releaseImageConfigRequest2** | [**ReleaseImageConfigRequest2**](ReleaseImageConfigRequest2.md) |  | 

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


## ImageConfigApiRolloutFleetImageConfig

> ImageConfigApiRolloutFleetImageConfig(ctx, serviceId, id).Execute()

RolloutFleetImageConfig image-config-api

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
	serviceId := "s-12345678" // string | The service ID
	id := "imgc-12345678" // string | The image configuration ID

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	r, err := apiClient.ImageConfigApiAPI.ImageConfigApiRolloutFleetImageConfig(context.Background(), serviceId, id).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `ImageConfigApiAPI.ImageConfigApiRolloutFleetImageConfig``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**serviceId** | **string** | The service ID | 
**id** | **string** | The image configuration ID | 

### Other Parameters

Other parameters are passed through a pointer to a apiImageConfigApiRolloutFleetImageConfigRequest struct via the builder pattern


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


## ImageConfigApiRolloutFleetImageStatus

> RolloutFleetImageStatusResult ImageConfigApiRolloutFleetImageStatus(ctx, serviceId, id).Execute()

RolloutFleetImageStatus image-config-api

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
	serviceId := "s-12345678" // string | The service ID
	id := "imgc-12345678" // string | The image configuration ID

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.ImageConfigApiAPI.ImageConfigApiRolloutFleetImageStatus(context.Background(), serviceId, id).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `ImageConfigApiAPI.ImageConfigApiRolloutFleetImageStatus``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `ImageConfigApiRolloutFleetImageStatus`: RolloutFleetImageStatusResult
	fmt.Fprintf(os.Stdout, "Response from `ImageConfigApiAPI.ImageConfigApiRolloutFleetImageStatus`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**serviceId** | **string** | The service ID | 
**id** | **string** | The image configuration ID | 

### Other Parameters

Other parameters are passed through a pointer to a apiImageConfigApiRolloutFleetImageStatusRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------



### Return type

[**RolloutFleetImageStatusResult**](RolloutFleetImageStatusResult.md)

### Authorization

[api_key_header_Authorization](../README.md#api_key_header_Authorization)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json, application/vnd.goa.error

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## ImageConfigApiUpdateImageConfig

> ImageConfigApiUpdateImageConfig(ctx, serviceId, id).UpdateImageConfigRequest2(updateImageConfigRequest2).Execute()

UpdateImageConfig image-config-api

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
	serviceId := "s-12345678" // string | The service ID
	id := "imgc-12345678" // string | The image configuration ID
	updateImageConfigRequest2 := *openapiclient.NewUpdateImageConfigRequest2() // UpdateImageConfigRequest2 | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	r, err := apiClient.ImageConfigApiAPI.ImageConfigApiUpdateImageConfig(context.Background(), serviceId, id).UpdateImageConfigRequest2(updateImageConfigRequest2).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `ImageConfigApiAPI.ImageConfigApiUpdateImageConfig``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**serviceId** | **string** | The service ID | 
**id** | **string** | The image configuration ID | 

### Other Parameters

Other parameters are passed through a pointer to a apiImageConfigApiUpdateImageConfigRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


 **updateImageConfigRequest2** | [**UpdateImageConfigRequest2**](UpdateImageConfigRequest2.md) |  | 

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

