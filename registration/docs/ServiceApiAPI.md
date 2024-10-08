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

> BuildServiceFromServicePlanSpecResult ServiceApiBuildServiceFromComposeSpec(ctx).BuildServiceFromComposeSpecRequestBody(buildServiceFromComposeSpecRequestBody).Execute()

BuildServiceFromComposeSpec service-api

### Example

```go
package main

import (
	"context"
	"fmt"
	"os"
	openapiclient "github.com/omnistrate/omnistrate-sdk-go"
)

func main() {
	buildServiceFromComposeSpecRequestBody := *openapiclient.NewBuildServiceFromComposeSpecRequestBody("Rerum rem et saepe et voluptatem sit.", "MySQL multi-writer service") // BuildServiceFromComposeSpecRequestBody | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.ServiceApiAPI.ServiceApiBuildServiceFromComposeSpec(context.Background()).BuildServiceFromComposeSpecRequestBody(buildServiceFromComposeSpecRequestBody).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `ServiceApiAPI.ServiceApiBuildServiceFromComposeSpec``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `ServiceApiBuildServiceFromComposeSpec`: BuildServiceFromServicePlanSpecResult
	fmt.Fprintf(os.Stdout, "Response from `ServiceApiAPI.ServiceApiBuildServiceFromComposeSpec`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiServiceApiBuildServiceFromComposeSpecRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **buildServiceFromComposeSpecRequestBody** | [**BuildServiceFromComposeSpecRequestBody**](BuildServiceFromComposeSpecRequestBody.md) |  | 

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


## ServiceApiBuildServiceFromServicePlanSpec

> BuildServiceFromServicePlanSpecResult ServiceApiBuildServiceFromServicePlanSpec(ctx).BuildServiceFromServicePlanSpecRequestBody(buildServiceFromServicePlanSpecRequestBody).Execute()

BuildServiceFromServicePlanSpec service-api

### Example

```go
package main

import (
	"context"
	"fmt"
	"os"
	openapiclient "github.com/omnistrate/omnistrate-sdk-go"
)

func main() {
	buildServiceFromServicePlanSpecRequestBody := *openapiclient.NewBuildServiceFromServicePlanSpecRequestBody("Et sint temporibus et in ut.", "MySQL multi-writer service") // BuildServiceFromServicePlanSpecRequestBody | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.ServiceApiAPI.ServiceApiBuildServiceFromServicePlanSpec(context.Background()).BuildServiceFromServicePlanSpecRequestBody(buildServiceFromServicePlanSpecRequestBody).Execute()
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
 **buildServiceFromServicePlanSpecRequestBody** | [**BuildServiceFromServicePlanSpecRequestBody**](BuildServiceFromServicePlanSpecRequestBody.md) |  | 

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

> string ServiceApiCreateService(ctx).CreateServiceRequestBody(createServiceRequestBody).Execute()

CreateService service-api

### Example

```go
package main

import (
	"context"
	"fmt"
	"os"
	openapiclient "github.com/omnistrate/omnistrate-sdk-go"
)

func main() {
	createServiceRequestBody := *openapiclient.NewCreateServiceRequestBody("A MySQL SaaS specializing in multi-writer clusters for high availability", "MySQL multi-writer service") // CreateServiceRequestBody | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.ServiceApiAPI.ServiceApiCreateService(context.Background()).CreateServiceRequestBody(createServiceRequestBody).Execute()
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
 **createServiceRequestBody** | [**CreateServiceRequestBody**](CreateServiceRequestBody.md) |  | 

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

> string ServiceApiCreateServiceFromComposeSpec(ctx).CreateServiceFromComposeSpecRequestBody(createServiceFromComposeSpecRequestBody).Execute()

CreateServiceFromComposeSpec service-api

### Example

```go
package main

import (
	"context"
	"fmt"
	"os"
	openapiclient "github.com/omnistrate/omnistrate-sdk-go"
)

func main() {
	createServiceFromComposeSpecRequestBody := *openapiclient.NewCreateServiceFromComposeSpecRequestBody("A MySQL SaaS specializing in multi-writer clusters for high availability", "Ipsam vero minima et soluta eius.", "text/plain", "mysql.yaml", "MySQL multi-writer service") // CreateServiceFromComposeSpecRequestBody | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.ServiceApiAPI.ServiceApiCreateServiceFromComposeSpec(context.Background()).CreateServiceFromComposeSpecRequestBody(createServiceFromComposeSpecRequestBody).Execute()
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
 **createServiceFromComposeSpecRequestBody** | [**CreateServiceFromComposeSpecRequestBody**](CreateServiceFromComposeSpecRequestBody.md) |  | 

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
	openapiclient "github.com/omnistrate/omnistrate-sdk-go"
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
	openapiclient "github.com/omnistrate/omnistrate-sdk-go"
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
	openapiclient "github.com/omnistrate/omnistrate-sdk-go"
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
	openapiclient "github.com/omnistrate/omnistrate-sdk-go"
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

> ServiceApiUpdateService(ctx, id).UpdateServiceRequestBody(updateServiceRequestBody).Execute()

UpdateService service-api

### Example

```go
package main

import (
	"context"
	"fmt"
	"os"
	openapiclient "github.com/omnistrate/omnistrate-sdk-go"
)

func main() {
	id := "s-12345678" // string | The service ID to operate on
	updateServiceRequestBody := *openapiclient.NewUpdateServiceRequestBody() // UpdateServiceRequestBody | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	r, err := apiClient.ServiceApiAPI.ServiceApiUpdateService(context.Background(), id).UpdateServiceRequestBody(updateServiceRequestBody).Execute()
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

 **updateServiceRequestBody** | [**UpdateServiceRequestBody**](UpdateServiceRequestBody.md) |  | 

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

