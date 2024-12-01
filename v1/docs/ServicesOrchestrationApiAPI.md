# \ServicesOrchestrationApiAPI

All URIs are relative to *https://api.omnistrate.cloud*

Method | HTTP request | Description
------------- | ------------- | -------------
[**ServicesOrchestrationApiCreateServicesOrchestration**](ServicesOrchestrationApiAPI.md#ServicesOrchestrationApiCreateServicesOrchestration) | **Post** /2022-09-01-00/services-orchestration | CreateServicesOrchestration services-orchestration-api
[**ServicesOrchestrationApiDeleteServicesOrchestration**](ServicesOrchestrationApiAPI.md#ServicesOrchestrationApiDeleteServicesOrchestration) | **Delete** /2022-09-01-00/services-orchestration/{id} | DeleteServicesOrchestration services-orchestration-api
[**ServicesOrchestrationApiDescribeServicesOrchestration**](ServicesOrchestrationApiAPI.md#ServicesOrchestrationApiDescribeServicesOrchestration) | **Get** /2022-09-01-00/services-orchestration/{id} | DescribeServicesOrchestration services-orchestration-api
[**ServicesOrchestrationApiListServicesOrchestration**](ServicesOrchestrationApiAPI.md#ServicesOrchestrationApiListServicesOrchestration) | **Get** /2022-09-01-00/services-orchestration | ListServicesOrchestration services-orchestration-api
[**ServicesOrchestrationApiModifyServicesOrchestration**](ServicesOrchestrationApiAPI.md#ServicesOrchestrationApiModifyServicesOrchestration) | **Patch** /2022-09-01-00/services-orchestration/{id} | ModifyServicesOrchestration services-orchestration-api



## ServicesOrchestrationApiCreateServicesOrchestration

> CreateServicesOrchestrationResponseBody ServicesOrchestrationApiCreateServicesOrchestration(ctx).CreateServicesOrchestrationRequestBody(createServicesOrchestrationRequestBody).Execute()

CreateServicesOrchestration services-orchestration-api

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
	createServicesOrchestrationRequestBody := *openapiclient.NewCreateServicesOrchestrationRequestBody("Et vero pariatur sed et veritatis possimus.") // CreateServicesOrchestrationRequestBody | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.ServicesOrchestrationApiAPI.ServicesOrchestrationApiCreateServicesOrchestration(context.Background()).CreateServicesOrchestrationRequestBody(createServicesOrchestrationRequestBody).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `ServicesOrchestrationApiAPI.ServicesOrchestrationApiCreateServicesOrchestration``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `ServicesOrchestrationApiCreateServicesOrchestration`: CreateServicesOrchestrationResponseBody
	fmt.Fprintf(os.Stdout, "Response from `ServicesOrchestrationApiAPI.ServicesOrchestrationApiCreateServicesOrchestration`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiServicesOrchestrationApiCreateServicesOrchestrationRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **createServicesOrchestrationRequestBody** | [**CreateServicesOrchestrationRequestBody**](CreateServicesOrchestrationRequestBody.md) |  | 

### Return type

[**CreateServicesOrchestrationResponseBody**](CreateServicesOrchestrationResponseBody.md)

### Authorization

[api_key_header_Authorization](../README.md#api_key_header_Authorization)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json, application/vnd.goa.error

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## ServicesOrchestrationApiDeleteServicesOrchestration

> ServicesOrchestrationApiDeleteServicesOrchestration(ctx, id).Execute()

DeleteServicesOrchestration services-orchestration-api

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
	id := "so-12345678" // string | The ID of the services orchestration

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	r, err := apiClient.ServicesOrchestrationApiAPI.ServicesOrchestrationApiDeleteServicesOrchestration(context.Background(), id).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `ServicesOrchestrationApiAPI.ServicesOrchestrationApiDeleteServicesOrchestration``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **string** | The ID of the services orchestration | 

### Other Parameters

Other parameters are passed through a pointer to a apiServicesOrchestrationApiDeleteServicesOrchestrationRequest struct via the builder pattern


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


## ServicesOrchestrationApiDescribeServicesOrchestration

> DescribeServicesOrchestrationResult ServicesOrchestrationApiDescribeServicesOrchestration(ctx, id).Execute()

DescribeServicesOrchestration services-orchestration-api

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
	id := "so-12345678" // string | The ID of the services orchestration

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.ServicesOrchestrationApiAPI.ServicesOrchestrationApiDescribeServicesOrchestration(context.Background(), id).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `ServicesOrchestrationApiAPI.ServicesOrchestrationApiDescribeServicesOrchestration``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `ServicesOrchestrationApiDescribeServicesOrchestration`: DescribeServicesOrchestrationResult
	fmt.Fprintf(os.Stdout, "Response from `ServicesOrchestrationApiAPI.ServicesOrchestrationApiDescribeServicesOrchestration`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **string** | The ID of the services orchestration | 

### Other Parameters

Other parameters are passed through a pointer to a apiServicesOrchestrationApiDescribeServicesOrchestrationRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

[**DescribeServicesOrchestrationResult**](DescribeServicesOrchestrationResult.md)

### Authorization

[api_key_header_Authorization](../README.md#api_key_header_Authorization)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json, application/vnd.goa.error

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## ServicesOrchestrationApiListServicesOrchestration

> []DescribeServicesOrchestrationResult ServicesOrchestrationApiListServicesOrchestration(ctx).Execute()

ListServicesOrchestration services-orchestration-api

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
	resp, r, err := apiClient.ServicesOrchestrationApiAPI.ServicesOrchestrationApiListServicesOrchestration(context.Background()).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `ServicesOrchestrationApiAPI.ServicesOrchestrationApiListServicesOrchestration``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `ServicesOrchestrationApiListServicesOrchestration`: []DescribeServicesOrchestrationResult
	fmt.Fprintf(os.Stdout, "Response from `ServicesOrchestrationApiAPI.ServicesOrchestrationApiListServicesOrchestration`: %v\n", resp)
}
```

### Path Parameters

This endpoint does not need any parameter.

### Other Parameters

Other parameters are passed through a pointer to a apiServicesOrchestrationApiListServicesOrchestrationRequest struct via the builder pattern


### Return type

[**[]DescribeServicesOrchestrationResult**](DescribeServicesOrchestrationResult.md)

### Authorization

[api_key_header_Authorization](../README.md#api_key_header_Authorization)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json, application/vnd.goa.error

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## ServicesOrchestrationApiModifyServicesOrchestration

> CreateServicesOrchestrationResponseBody ServicesOrchestrationApiModifyServicesOrchestration(ctx, id).ModifyServicesOrchestrationRequestBody(modifyServicesOrchestrationRequestBody).Execute()

ModifyServicesOrchestration services-orchestration-api

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
	id := "so-12345678" // string | The ID of the services orchestration
	modifyServicesOrchestrationRequestBody := *openapiclient.NewModifyServicesOrchestrationRequestBody("Voluptatibus quod fuga sapiente doloremque deleniti.") // ModifyServicesOrchestrationRequestBody | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.ServicesOrchestrationApiAPI.ServicesOrchestrationApiModifyServicesOrchestration(context.Background(), id).ModifyServicesOrchestrationRequestBody(modifyServicesOrchestrationRequestBody).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `ServicesOrchestrationApiAPI.ServicesOrchestrationApiModifyServicesOrchestration``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `ServicesOrchestrationApiModifyServicesOrchestration`: CreateServicesOrchestrationResponseBody
	fmt.Fprintf(os.Stdout, "Response from `ServicesOrchestrationApiAPI.ServicesOrchestrationApiModifyServicesOrchestration`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **string** | The ID of the services orchestration | 

### Other Parameters

Other parameters are passed through a pointer to a apiServicesOrchestrationApiModifyServicesOrchestrationRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **modifyServicesOrchestrationRequestBody** | [**ModifyServicesOrchestrationRequestBody**](ModifyServicesOrchestrationRequestBody.md) |  | 

### Return type

[**CreateServicesOrchestrationResponseBody**](CreateServicesOrchestrationResponseBody.md)

### Authorization

[api_key_header_Authorization](../README.md#api_key_header_Authorization)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json, application/vnd.goa.error

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

