# \ComputeConfigApiAPI

All URIs are relative to *https://api.omnistrate.cloud*

Method | HTTP request | Description
------------- | ------------- | -------------
[**ComputeConfigApiAddComputeInstanceType**](ComputeConfigApiAPI.md#ComputeConfigApiAddComputeInstanceType) | **Post** /2022-09-01-00/service/{serviceId}/compute-config/{id}/instance-type | AddComputeInstanceType compute-config-api
[**ComputeConfigApiCreateComputeConfig**](ComputeConfigApiAPI.md#ComputeConfigApiCreateComputeConfig) | **Post** /2022-09-01-00/service/{serviceId}/compute-config | CreateComputeConfig compute-config-api
[**ComputeConfigApiDeleteComputeConfig**](ComputeConfigApiAPI.md#ComputeConfigApiDeleteComputeConfig) | **Delete** /2022-09-01-00/service/{serviceId}/compute-config/{id} | DeleteComputeConfig compute-config-api
[**ComputeConfigApiDescribeComputeConfig**](ComputeConfigApiAPI.md#ComputeConfigApiDescribeComputeConfig) | **Get** /2022-09-01-00/service/{serviceId}/compute-config/{id} | DescribeComputeConfig compute-config-api
[**ComputeConfigApiListComputeConfig**](ComputeConfigApiAPI.md#ComputeConfigApiListComputeConfig) | **Get** /2022-09-01-00/service/{serviceId}/compute-config | ListComputeConfig compute-config-api
[**ComputeConfigApiListComputeInstanceTypes**](ComputeConfigApiAPI.md#ComputeConfigApiListComputeInstanceTypes) | **Get** /2022-09-01-00/service/{serviceId}/compute-config/cloud-provider/{cloudProviderName}/instance-types | ListComputeInstanceTypes compute-config-api
[**ComputeConfigApiRemoveComputeInstanceType**](ComputeConfigApiAPI.md#ComputeConfigApiRemoveComputeInstanceType) | **Delete** /2022-09-01-00/service/{serviceId}/compute-config/{id}/instance-type | RemoveComputeInstanceType compute-config-api
[**ComputeConfigApiUpdateComputeConfig**](ComputeConfigApiAPI.md#ComputeConfigApiUpdateComputeConfig) | **Patch** /2022-09-01-00/service/{serviceId}/compute-config/{id} | UpdateComputeConfig compute-config-api



## ComputeConfigApiAddComputeInstanceType

> ComputeConfigApiAddComputeInstanceType(ctx, serviceId, id).AddComputeInstanceTypeRequestBody(addComputeInstanceTypeRequestBody).Execute()

AddComputeInstanceType compute-config-api

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
	id := "cc-12345678" // string | ID of the compute config
	addComputeInstanceTypeRequestBody := *openapiclient.NewAddComputeInstanceTypeRequestBody("aws", "t3.micro") // AddComputeInstanceTypeRequestBody | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	r, err := apiClient.ComputeConfigApiAPI.ComputeConfigApiAddComputeInstanceType(context.Background(), serviceId, id).AddComputeInstanceTypeRequestBody(addComputeInstanceTypeRequestBody).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `ComputeConfigApiAPI.ComputeConfigApiAddComputeInstanceType``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**serviceId** | **string** | The service ID | 
**id** | **string** | ID of the compute config | 

### Other Parameters

Other parameters are passed through a pointer to a apiComputeConfigApiAddComputeInstanceTypeRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


 **addComputeInstanceTypeRequestBody** | [**AddComputeInstanceTypeRequestBody**](AddComputeInstanceTypeRequestBody.md) |  | 

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


## ComputeConfigApiCreateComputeConfig

> string ComputeConfigApiCreateComputeConfig(ctx, serviceId).CreateComputeConfigRequestBody(createComputeConfigRequestBody).Execute()

CreateComputeConfig compute-config-api

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
	createComputeConfigRequestBody := *openapiclient.NewCreateComputeConfigRequestBody("my compute config description", "my compute config") // CreateComputeConfigRequestBody | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.ComputeConfigApiAPI.ComputeConfigApiCreateComputeConfig(context.Background(), serviceId).CreateComputeConfigRequestBody(createComputeConfigRequestBody).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `ComputeConfigApiAPI.ComputeConfigApiCreateComputeConfig``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `ComputeConfigApiCreateComputeConfig`: string
	fmt.Fprintf(os.Stdout, "Response from `ComputeConfigApiAPI.ComputeConfigApiCreateComputeConfig`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**serviceId** | **string** | The service ID | 

### Other Parameters

Other parameters are passed through a pointer to a apiComputeConfigApiCreateComputeConfigRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **createComputeConfigRequestBody** | [**CreateComputeConfigRequestBody**](CreateComputeConfigRequestBody.md) |  | 

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


## ComputeConfigApiDeleteComputeConfig

> ComputeConfigApiDeleteComputeConfig(ctx, serviceId, id).Execute()

DeleteComputeConfig compute-config-api

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
	id := "cc-12345678" // string | ID of the compute config

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	r, err := apiClient.ComputeConfigApiAPI.ComputeConfigApiDeleteComputeConfig(context.Background(), serviceId, id).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `ComputeConfigApiAPI.ComputeConfigApiDeleteComputeConfig``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**serviceId** | **string** | The service ID | 
**id** | **string** | ID of the compute config | 

### Other Parameters

Other parameters are passed through a pointer to a apiComputeConfigApiDeleteComputeConfigRequest struct via the builder pattern


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


## ComputeConfigApiDescribeComputeConfig

> DescribeComputeConfigResult ComputeConfigApiDescribeComputeConfig(ctx, serviceId, id).ProductTierVersion(productTierVersion).ProductTierId(productTierId).Execute()

DescribeComputeConfig compute-config-api

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
	id := "cc-12345678" // string | ID of the compute config
	productTierVersion := "Eum dignissimos natus quo sit qui qui." // string | Product tier version of the compute config to describe. If not specified, the latest version is described. (optional)
	productTierId := "Beatae beatae." // string | ProductTierId of the compute config to describe. Needs to specified in combination with the product tier version (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.ComputeConfigApiAPI.ComputeConfigApiDescribeComputeConfig(context.Background(), serviceId, id).ProductTierVersion(productTierVersion).ProductTierId(productTierId).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `ComputeConfigApiAPI.ComputeConfigApiDescribeComputeConfig``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `ComputeConfigApiDescribeComputeConfig`: DescribeComputeConfigResult
	fmt.Fprintf(os.Stdout, "Response from `ComputeConfigApiAPI.ComputeConfigApiDescribeComputeConfig`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**serviceId** | **string** | The service ID | 
**id** | **string** | ID of the compute config | 

### Other Parameters

Other parameters are passed through a pointer to a apiComputeConfigApiDescribeComputeConfigRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


 **productTierVersion** | **string** | Product tier version of the compute config to describe. If not specified, the latest version is described. | 
 **productTierId** | **string** | ProductTierId of the compute config to describe. Needs to specified in combination with the product tier version | 

### Return type

[**DescribeComputeConfigResult**](DescribeComputeConfigResult.md)

### Authorization

[api_key_header_Authorization](../README.md#api_key_header_Authorization)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json, application/vnd.goa.error

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## ComputeConfigApiListComputeConfig

> ListServiceEnvironmentsResult ComputeConfigApiListComputeConfig(ctx, serviceId).Managed(managed).Execute()

ListComputeConfig compute-config-api

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
	managed := false // bool | Is compute config managed by omnistrate (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.ComputeConfigApiAPI.ComputeConfigApiListComputeConfig(context.Background(), serviceId).Managed(managed).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `ComputeConfigApiAPI.ComputeConfigApiListComputeConfig``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `ComputeConfigApiListComputeConfig`: ListServiceEnvironmentsResult
	fmt.Fprintf(os.Stdout, "Response from `ComputeConfigApiAPI.ComputeConfigApiListComputeConfig`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**serviceId** | **string** | The service ID | 

### Other Parameters

Other parameters are passed through a pointer to a apiComputeConfigApiListComputeConfigRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **managed** | **bool** | Is compute config managed by omnistrate | 

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


## ComputeConfigApiListComputeInstanceTypes

> ListComputeInstanceTypesResult ComputeConfigApiListComputeInstanceTypes(ctx, serviceId, cloudProviderName).Execute()

ListComputeInstanceTypes compute-config-api

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
	cloudProviderName := "aws" // string | The cloud provider for this compute instance type config

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.ComputeConfigApiAPI.ComputeConfigApiListComputeInstanceTypes(context.Background(), serviceId, cloudProviderName).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `ComputeConfigApiAPI.ComputeConfigApiListComputeInstanceTypes``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `ComputeConfigApiListComputeInstanceTypes`: ListComputeInstanceTypesResult
	fmt.Fprintf(os.Stdout, "Response from `ComputeConfigApiAPI.ComputeConfigApiListComputeInstanceTypes`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**serviceId** | **string** | The service ID | 
**cloudProviderName** | **string** | The cloud provider for this compute instance type config | 

### Other Parameters

Other parameters are passed through a pointer to a apiComputeConfigApiListComputeInstanceTypesRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------



### Return type

[**ListComputeInstanceTypesResult**](ListComputeInstanceTypesResult.md)

### Authorization

[api_key_header_Authorization](../README.md#api_key_header_Authorization)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json, application/vnd.goa.error

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## ComputeConfigApiRemoveComputeInstanceType

> ComputeConfigApiRemoveComputeInstanceType(ctx, serviceId, id).AddComputeInstanceTypeRequestBody(addComputeInstanceTypeRequestBody).Execute()

RemoveComputeInstanceType compute-config-api

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
	id := "cc-12345678" // string | ID of the compute config
	addComputeInstanceTypeRequestBody := *openapiclient.NewAddComputeInstanceTypeRequestBody("aws", "t3.micro") // AddComputeInstanceTypeRequestBody | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	r, err := apiClient.ComputeConfigApiAPI.ComputeConfigApiRemoveComputeInstanceType(context.Background(), serviceId, id).AddComputeInstanceTypeRequestBody(addComputeInstanceTypeRequestBody).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `ComputeConfigApiAPI.ComputeConfigApiRemoveComputeInstanceType``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**serviceId** | **string** | The service ID | 
**id** | **string** | ID of the compute config | 

### Other Parameters

Other parameters are passed through a pointer to a apiComputeConfigApiRemoveComputeInstanceTypeRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


 **addComputeInstanceTypeRequestBody** | [**AddComputeInstanceTypeRequestBody**](AddComputeInstanceTypeRequestBody.md) |  | 

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


## ComputeConfigApiUpdateComputeConfig

> ComputeConfigApiUpdateComputeConfig(ctx, serviceId, id).UpdateComputeConfigRequestBody(updateComputeConfigRequestBody).Execute()

UpdateComputeConfig compute-config-api

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
	id := "cc-12345678" // string | ID of the compute config
	updateComputeConfigRequestBody := *openapiclient.NewUpdateComputeConfigRequestBody() // UpdateComputeConfigRequestBody | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	r, err := apiClient.ComputeConfigApiAPI.ComputeConfigApiUpdateComputeConfig(context.Background(), serviceId, id).UpdateComputeConfigRequestBody(updateComputeConfigRequestBody).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `ComputeConfigApiAPI.ComputeConfigApiUpdateComputeConfig``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**serviceId** | **string** | The service ID | 
**id** | **string** | ID of the compute config | 

### Other Parameters

Other parameters are passed through a pointer to a apiComputeConfigApiUpdateComputeConfigRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


 **updateComputeConfigRequestBody** | [**UpdateComputeConfigRequestBody**](UpdateComputeConfigRequestBody.md) |  | 

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

