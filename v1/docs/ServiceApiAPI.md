# \ServiceApiAPI

All URIs are relative to *https://api.omnistrate.cloud*

Method | HTTP request | Description
------------- | ------------- | -------------
[**ServiceApiBuildServiceFromComposeSpec**](ServiceApiAPI.md#ServiceApiBuildServiceFromComposeSpec) | **Put** /2022-09-01-00/service/composespec | BuildServiceFromComposeSpec service-api
[**ServiceApiBuildServiceFromServicePlanSpec**](ServiceApiAPI.md#ServiceApiBuildServiceFromServicePlanSpec) | **Put** /2022-09-01-00/service/serviceplanspec | BuildServiceFromServicePlanSpec service-api
[**ServiceApiCreateService**](ServiceApiAPI.md#ServiceApiCreateService) | **Post** /2022-09-01-00/service | CreateService service-api
[**ServiceApiCreateServiceFromComposeSpec**](ServiceApiAPI.md#ServiceApiCreateServiceFromComposeSpec) | **Post** /2022-09-01-00/service/composespec | CreateServiceFromComposeSpec service-api
[**ServiceApiDeleteService**](ServiceApiAPI.md#ServiceApiDeleteService) | **Delete** /2022-09-01-00/service/{id} | DeleteService service-api
[**ServiceApiDescribeService**](ServiceApiAPI.md#ServiceApiDescribeService) | **Get** /2022-09-01-00/service/{id} | DescribeService service-api
[**ServiceApiListService**](ServiceApiAPI.md#ServiceApiListService) | **Get** /2022-09-01-00/service | ListService service-api
[**ServiceApiServiceHealth**](ServiceApiAPI.md#ServiceApiServiceHealth) | **Get** /2022-09-01-00/service/{id}/health | ServiceHealth service-api
[**ServiceApiUpdateService**](ServiceApiAPI.md#ServiceApiUpdateService) | **Patch** /2022-09-01-00/service/{id} | UpdateService service-api



## ServiceApiBuildServiceFromComposeSpec

> BuildServiceFromComposeSpecResult ServiceApiBuildServiceFromComposeSpec(ctx).BuildServiceFromComposeSpecRequest2(buildServiceFromComposeSpecRequest2).Execute()

BuildServiceFromComposeSpec service-api

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
	buildServiceFromComposeSpecRequest2 := *openapiclient.NewBuildServiceFromComposeSpecRequest2("Et qui debitis.", "MySQL multi-writer service") // BuildServiceFromComposeSpecRequest2 | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.ServiceApiAPI.ServiceApiBuildServiceFromComposeSpec(context.Background()).BuildServiceFromComposeSpecRequest2(buildServiceFromComposeSpecRequest2).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `ServiceApiAPI.ServiceApiBuildServiceFromComposeSpec``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `ServiceApiBuildServiceFromComposeSpec`: BuildServiceFromComposeSpecResult
	fmt.Fprintf(os.Stdout, "Response from `ServiceApiAPI.ServiceApiBuildServiceFromComposeSpec`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiServiceApiBuildServiceFromComposeSpecRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **buildServiceFromComposeSpecRequest2** | [**BuildServiceFromComposeSpecRequest2**](BuildServiceFromComposeSpecRequest2.md) |  | 

### Return type

[**BuildServiceFromComposeSpecResult**](BuildServiceFromComposeSpecResult.md)

### Authorization

[api_key_header_Authorization](../README.md#api_key_header_Authorization)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json, application/vnd.goa.error

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## ServiceApiBuildServiceFromServicePlanSpec

> BuildServiceFromServicePlanSpecResult ServiceApiBuildServiceFromServicePlanSpec(ctx).BuildServiceFromServicePlanSpecRequest2(buildServiceFromServicePlanSpecRequest2).Execute()

BuildServiceFromServicePlanSpec service-api

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
	buildServiceFromServicePlanSpecRequest2 := *openapiclient.NewBuildServiceFromServicePlanSpecRequest2("Ut voluptas sit dignissimos quia doloribus nesciunt.", "MySQL multi-writer service") // BuildServiceFromServicePlanSpecRequest2 | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.ServiceApiAPI.ServiceApiBuildServiceFromServicePlanSpec(context.Background()).BuildServiceFromServicePlanSpecRequest2(buildServiceFromServicePlanSpecRequest2).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `ServiceApiAPI.ServiceApiBuildServiceFromServicePlanSpec``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `ServiceApiBuildServiceFromServicePlanSpec`: BuildServiceFromServicePlanSpecResult
	fmt.Fprintf(os.Stdout, "Response from `ServiceApiAPI.ServiceApiBuildServiceFromServicePlanSpec`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiServiceApiBuildServiceFromServicePlanSpecRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **buildServiceFromServicePlanSpecRequest2** | [**BuildServiceFromServicePlanSpecRequest2**](BuildServiceFromServicePlanSpecRequest2.md) |  | 

### Return type

[**BuildServiceFromServicePlanSpecResult**](BuildServiceFromServicePlanSpecResult.md)

### Authorization

[api_key_header_Authorization](../README.md#api_key_header_Authorization)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json, application/vnd.goa.error

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## ServiceApiCreateService

> string ServiceApiCreateService(ctx).CreateServiceRequest2(createServiceRequest2).Execute()

CreateService service-api

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
	createServiceRequest2 := *openapiclient.NewCreateServiceRequest2("A MySQL SaaS specializing in multi-writer clusters for high availability", "MySQL multi-writer service") // CreateServiceRequest2 | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.ServiceApiAPI.ServiceApiCreateService(context.Background()).CreateServiceRequest2(createServiceRequest2).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `ServiceApiAPI.ServiceApiCreateService``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `ServiceApiCreateService`: string
	fmt.Fprintf(os.Stdout, "Response from `ServiceApiAPI.ServiceApiCreateService`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiServiceApiCreateServiceRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **createServiceRequest2** | [**CreateServiceRequest2**](CreateServiceRequest2.md) |  | 

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


## ServiceApiCreateServiceFromComposeSpec

> string ServiceApiCreateServiceFromComposeSpec(ctx).CreateServiceFromComposeSpecRequest2(createServiceFromComposeSpecRequest2).Execute()

CreateServiceFromComposeSpec service-api

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
	createServiceFromComposeSpecRequest2 := *openapiclient.NewCreateServiceFromComposeSpecRequest2("A MySQL SaaS specializing in multi-writer clusters for high availability", "Laboriosam sed quaerat.", "text/plain", "mysql.yaml", "MySQL multi-writer service") // CreateServiceFromComposeSpecRequest2 | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.ServiceApiAPI.ServiceApiCreateServiceFromComposeSpec(context.Background()).CreateServiceFromComposeSpecRequest2(createServiceFromComposeSpecRequest2).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `ServiceApiAPI.ServiceApiCreateServiceFromComposeSpec``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `ServiceApiCreateServiceFromComposeSpec`: string
	fmt.Fprintf(os.Stdout, "Response from `ServiceApiAPI.ServiceApiCreateServiceFromComposeSpec`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiServiceApiCreateServiceFromComposeSpecRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **createServiceFromComposeSpecRequest2** | [**CreateServiceFromComposeSpecRequest2**](CreateServiceFromComposeSpecRequest2.md) |  | 

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


## ServiceApiDeleteService

> ServiceApiDeleteService(ctx, id).Execute()

DeleteService service-api

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
	id := "s-12345678" // string | The service ID to operate on

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	r, err := apiClient.ServiceApiAPI.ServiceApiDeleteService(context.Background(), id).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `ServiceApiAPI.ServiceApiDeleteService``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **string** | The service ID to operate on | 

### Other Parameters

Other parameters are passed through a pointer to a apiServiceApiDeleteServiceRequest struct via the builder pattern


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


## ServiceApiDescribeService

> DescribeServiceResult ServiceApiDescribeService(ctx, id).Execute()

DescribeService service-api

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
	id := "s-12345678" // string | The service ID to operate on

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.ServiceApiAPI.ServiceApiDescribeService(context.Background(), id).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `ServiceApiAPI.ServiceApiDescribeService``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `ServiceApiDescribeService`: DescribeServiceResult
	fmt.Fprintf(os.Stdout, "Response from `ServiceApiAPI.ServiceApiDescribeService`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **string** | The service ID to operate on | 

### Other Parameters

Other parameters are passed through a pointer to a apiServiceApiDescribeServiceRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

[**DescribeServiceResult**](DescribeServiceResult.md)

### Authorization

[api_key_header_Authorization](../README.md#api_key_header_Authorization)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json, application/vnd.goa.error

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## ServiceApiListService

> ListServiceResult ServiceApiListService(ctx).Execute()

ListService service-api

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
	resp, r, err := apiClient.ServiceApiAPI.ServiceApiListService(context.Background()).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `ServiceApiAPI.ServiceApiListService``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `ServiceApiListService`: ListServiceResult
	fmt.Fprintf(os.Stdout, "Response from `ServiceApiAPI.ServiceApiListService`: %v\n", resp)
}
```

### Path Parameters

This endpoint does not need any parameter.

### Other Parameters

Other parameters are passed through a pointer to a apiServiceApiListServiceRequest struct via the builder pattern


### Return type

[**ListServiceResult**](ListServiceResult.md)

### Authorization

[api_key_header_Authorization](../README.md#api_key_header_Authorization)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json, application/vnd.goa.error

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## ServiceApiServiceHealth

> ReportHealthResult ServiceApiServiceHealth(ctx, id).Execute()

ServiceHealth service-api

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
	id := "s-12345678" // string | The ID of the service

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.ServiceApiAPI.ServiceApiServiceHealth(context.Background(), id).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `ServiceApiAPI.ServiceApiServiceHealth``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `ServiceApiServiceHealth`: ReportHealthResult
	fmt.Fprintf(os.Stdout, "Response from `ServiceApiAPI.ServiceApiServiceHealth`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **string** | The ID of the service | 

### Other Parameters

Other parameters are passed through a pointer to a apiServiceApiServiceHealthRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

[**ReportHealthResult**](ReportHealthResult.md)

### Authorization

[api_key_header_Authorization](../README.md#api_key_header_Authorization)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json, application/vnd.goa.error

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## ServiceApiUpdateService

> ServiceApiUpdateService(ctx, id).UpdateServiceRequest2(updateServiceRequest2).Execute()

UpdateService service-api

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
	id := "s-12345678" // string | The service ID to operate on
	updateServiceRequest2 := *openapiclient.NewUpdateServiceRequest2() // UpdateServiceRequest2 | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	r, err := apiClient.ServiceApiAPI.ServiceApiUpdateService(context.Background(), id).UpdateServiceRequest2(updateServiceRequest2).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `ServiceApiAPI.ServiceApiUpdateService``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **string** | The service ID to operate on | 

### Other Parameters

Other parameters are passed through a pointer to a apiServiceApiUpdateServiceRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **updateServiceRequest2** | [**UpdateServiceRequest2**](UpdateServiceRequest2.md) |  | 

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

