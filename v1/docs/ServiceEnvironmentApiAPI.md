# \ServiceEnvironmentApiAPI

All URIs are relative to *https://api.omnistrate.cloud*

Method | HTTP request | Description
------------- | ------------- | -------------
[**ServiceEnvironmentApiCreateServiceEnvironment**](ServiceEnvironmentApiAPI.md#ServiceEnvironmentApiCreateServiceEnvironment) | **Post** /2022-09-01-00/service/{serviceId}/environment | CreateServiceEnvironment service-environment-api
[**ServiceEnvironmentApiDeleteServiceEnvironment**](ServiceEnvironmentApiAPI.md#ServiceEnvironmentApiDeleteServiceEnvironment) | **Delete** /2022-09-01-00/service/{serviceId}/environment/{id} | DeleteServiceEnvironment service-environment-api
[**ServiceEnvironmentApiDescribeServiceEnvironment**](ServiceEnvironmentApiAPI.md#ServiceEnvironmentApiDescribeServiceEnvironment) | **Get** /2022-09-01-00/service/{serviceId}/environment/{id} | DescribeServiceEnvironment service-environment-api
[**ServiceEnvironmentApiListServiceEnvironment**](ServiceEnvironmentApiAPI.md#ServiceEnvironmentApiListServiceEnvironment) | **Get** /2022-09-01-00/service/{serviceId}/environment | ListServiceEnvironment service-environment-api
[**ServiceEnvironmentApiPromoteServiceEnvironment**](ServiceEnvironmentApiAPI.md#ServiceEnvironmentApiPromoteServiceEnvironment) | **Post** /2022-09-01-00/service/{serviceId}/environment/{id}/promote | PromoteServiceEnvironment service-environment-api
[**ServiceEnvironmentApiPromoteServiceEnvironmentStatus**](ServiceEnvironmentApiAPI.md#ServiceEnvironmentApiPromoteServiceEnvironmentStatus) | **Get** /2022-09-01-00/service/{serviceId}/environment/{id}/promote | PromoteServiceEnvironmentStatus service-environment-api
[**ServiceEnvironmentApiUpdateServiceEnvironment**](ServiceEnvironmentApiAPI.md#ServiceEnvironmentApiUpdateServiceEnvironment) | **Patch** /2022-09-01-00/service/{serviceId}/environment/{id} | UpdateServiceEnvironment service-environment-api



## ServiceEnvironmentApiCreateServiceEnvironment

> string ServiceEnvironmentApiCreateServiceEnvironment(ctx, serviceId).CreateServiceEnvironmentRequest2(createServiceEnvironmentRequest2).Execute()

CreateServiceEnvironment service-environment-api

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
	serviceId := "s-12345678" // string | The ID of the service this environment belongs to
	createServiceEnvironmentRequest2 := *openapiclient.NewCreateServiceEnvironmentRequest2("Ut quod alias.", "The production environment for the MySQL multi-writer service", "Production") // CreateServiceEnvironmentRequest2 | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.ServiceEnvironmentApiAPI.ServiceEnvironmentApiCreateServiceEnvironment(context.Background(), serviceId).CreateServiceEnvironmentRequest2(createServiceEnvironmentRequest2).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `ServiceEnvironmentApiAPI.ServiceEnvironmentApiCreateServiceEnvironment``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `ServiceEnvironmentApiCreateServiceEnvironment`: string
	fmt.Fprintf(os.Stdout, "Response from `ServiceEnvironmentApiAPI.ServiceEnvironmentApiCreateServiceEnvironment`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**serviceId** | **string** | The ID of the service this environment belongs to | 

### Other Parameters

Other parameters are passed through a pointer to a apiServiceEnvironmentApiCreateServiceEnvironmentRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **createServiceEnvironmentRequest2** | [**CreateServiceEnvironmentRequest2**](CreateServiceEnvironmentRequest2.md) |  | 

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


## ServiceEnvironmentApiDeleteServiceEnvironment

> ServiceEnvironmentApiDeleteServiceEnvironment(ctx, serviceId, id).Execute()

DeleteServiceEnvironment service-environment-api

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
	serviceId := "s-12345678" // string | The ID of the service this environment belongs to
	id := "se-12345678" // string | The ID of the service environment

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	r, err := apiClient.ServiceEnvironmentApiAPI.ServiceEnvironmentApiDeleteServiceEnvironment(context.Background(), serviceId, id).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `ServiceEnvironmentApiAPI.ServiceEnvironmentApiDeleteServiceEnvironment``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**serviceId** | **string** | The ID of the service this environment belongs to | 
**id** | **string** | The ID of the service environment | 

### Other Parameters

Other parameters are passed through a pointer to a apiServiceEnvironmentApiDeleteServiceEnvironmentRequest struct via the builder pattern


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


## ServiceEnvironmentApiDescribeServiceEnvironment

> DescribeServiceEnvironmentResult ServiceEnvironmentApiDescribeServiceEnvironment(ctx, serviceId, id).Execute()

DescribeServiceEnvironment service-environment-api

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
	serviceId := "s-12345678" // string | The ID of the service this environment belongs to
	id := "se-12345678" // string | The ID of the service environment

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.ServiceEnvironmentApiAPI.ServiceEnvironmentApiDescribeServiceEnvironment(context.Background(), serviceId, id).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `ServiceEnvironmentApiAPI.ServiceEnvironmentApiDescribeServiceEnvironment``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `ServiceEnvironmentApiDescribeServiceEnvironment`: DescribeServiceEnvironmentResult
	fmt.Fprintf(os.Stdout, "Response from `ServiceEnvironmentApiAPI.ServiceEnvironmentApiDescribeServiceEnvironment`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**serviceId** | **string** | The ID of the service this environment belongs to | 
**id** | **string** | The ID of the service environment | 

### Other Parameters

Other parameters are passed through a pointer to a apiServiceEnvironmentApiDescribeServiceEnvironmentRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------



### Return type

[**DescribeServiceEnvironmentResult**](DescribeServiceEnvironmentResult.md)

### Authorization

[api_key_header_Authorization](../README.md#api_key_header_Authorization)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json, application/vnd.goa.error

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## ServiceEnvironmentApiListServiceEnvironment

> ListServiceEnvironmentsResult ServiceEnvironmentApiListServiceEnvironment(ctx, serviceId).Execute()

ListServiceEnvironment service-environment-api

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
	serviceId := "s-12345678" // string | The ID of the service this environment belongs to

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.ServiceEnvironmentApiAPI.ServiceEnvironmentApiListServiceEnvironment(context.Background(), serviceId).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `ServiceEnvironmentApiAPI.ServiceEnvironmentApiListServiceEnvironment``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `ServiceEnvironmentApiListServiceEnvironment`: ListServiceEnvironmentsResult
	fmt.Fprintf(os.Stdout, "Response from `ServiceEnvironmentApiAPI.ServiceEnvironmentApiListServiceEnvironment`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**serviceId** | **string** | The ID of the service this environment belongs to | 

### Other Parameters

Other parameters are passed through a pointer to a apiServiceEnvironmentApiListServiceEnvironmentRequest struct via the builder pattern


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


## ServiceEnvironmentApiPromoteServiceEnvironment

> ServiceEnvironmentApiPromoteServiceEnvironment(ctx, serviceId, id).PromoteServiceEnvironmentRequest2(promoteServiceEnvironmentRequest2).Execute()

PromoteServiceEnvironment service-environment-api

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
	serviceId := "s-12345678" // string | The ID of the service this environment belongs to
	id := "se-12345678" // string | The ID of the service environment
	promoteServiceEnvironmentRequest2 := *openapiclient.NewPromoteServiceEnvironmentRequest2() // PromoteServiceEnvironmentRequest2 | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	r, err := apiClient.ServiceEnvironmentApiAPI.ServiceEnvironmentApiPromoteServiceEnvironment(context.Background(), serviceId, id).PromoteServiceEnvironmentRequest2(promoteServiceEnvironmentRequest2).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `ServiceEnvironmentApiAPI.ServiceEnvironmentApiPromoteServiceEnvironment``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**serviceId** | **string** | The ID of the service this environment belongs to | 
**id** | **string** | The ID of the service environment | 

### Other Parameters

Other parameters are passed through a pointer to a apiServiceEnvironmentApiPromoteServiceEnvironmentRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


 **promoteServiceEnvironmentRequest2** | [**PromoteServiceEnvironmentRequest2**](PromoteServiceEnvironmentRequest2.md) |  | 

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


## ServiceEnvironmentApiPromoteServiceEnvironmentStatus

> []EnvironmentPromotionStatus ServiceEnvironmentApiPromoteServiceEnvironmentStatus(ctx, serviceId, id).Execute()

PromoteServiceEnvironmentStatus service-environment-api

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
	serviceId := "s-12345678" // string | The ID of the service this environment belongs to
	id := "se-12345678" // string | The ID of the service environment

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.ServiceEnvironmentApiAPI.ServiceEnvironmentApiPromoteServiceEnvironmentStatus(context.Background(), serviceId, id).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `ServiceEnvironmentApiAPI.ServiceEnvironmentApiPromoteServiceEnvironmentStatus``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `ServiceEnvironmentApiPromoteServiceEnvironmentStatus`: []EnvironmentPromotionStatus
	fmt.Fprintf(os.Stdout, "Response from `ServiceEnvironmentApiAPI.ServiceEnvironmentApiPromoteServiceEnvironmentStatus`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**serviceId** | **string** | The ID of the service this environment belongs to | 
**id** | **string** | The ID of the service environment | 

### Other Parameters

Other parameters are passed through a pointer to a apiServiceEnvironmentApiPromoteServiceEnvironmentStatusRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------



### Return type

[**[]EnvironmentPromotionStatus**](EnvironmentPromotionStatus.md)

### Authorization

[api_key_header_Authorization](../README.md#api_key_header_Authorization)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json, application/vnd.goa.error

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## ServiceEnvironmentApiUpdateServiceEnvironment

> ServiceEnvironmentApiUpdateServiceEnvironment(ctx, serviceId, id).UpdateServiceEnvironmentRequest2(updateServiceEnvironmentRequest2).Execute()

UpdateServiceEnvironment service-environment-api

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
	serviceId := "s-12345678" // string | The ID of the service this environment belongs to
	id := "se-12345678" // string | The ID of the service environment
	updateServiceEnvironmentRequest2 := *openapiclient.NewUpdateServiceEnvironmentRequest2() // UpdateServiceEnvironmentRequest2 | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	r, err := apiClient.ServiceEnvironmentApiAPI.ServiceEnvironmentApiUpdateServiceEnvironment(context.Background(), serviceId, id).UpdateServiceEnvironmentRequest2(updateServiceEnvironmentRequest2).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `ServiceEnvironmentApiAPI.ServiceEnvironmentApiUpdateServiceEnvironment``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**serviceId** | **string** | The ID of the service this environment belongs to | 
**id** | **string** | The ID of the service environment | 

### Other Parameters

Other parameters are passed through a pointer to a apiServiceEnvironmentApiUpdateServiceEnvironmentRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


 **updateServiceEnvironmentRequest2** | [**UpdateServiceEnvironmentRequest2**](UpdateServiceEnvironmentRequest2.md) |  | 

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

