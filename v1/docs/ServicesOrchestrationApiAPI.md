# \ServicesOrchestrationApiAPI

All URIs are relative to *https://api.omnistrate.cloud*

Method | HTTP request | Description
------------- | ------------- | -------------
[**ServicesOrchestrationApiCreateServicesOrchestration**](ServicesOrchestrationApiAPI.md#ServicesOrchestrationApiCreateServicesOrchestration) | **Post** /2022-09-01-00/resource-instance/services-orchestration | CreateServicesOrchestration services-orchestration-api
[**ServicesOrchestrationApiDeleteServicesOrchestration**](ServicesOrchestrationApiAPI.md#ServicesOrchestrationApiDeleteServicesOrchestration) | **Delete** /2022-09-01-00/resource-instance/services-orchestration/{id} | DeleteServicesOrchestration services-orchestration-api
[**ServicesOrchestrationApiDescribeServicesOrchestration**](ServicesOrchestrationApiAPI.md#ServicesOrchestrationApiDescribeServicesOrchestration) | **Get** /2022-09-01-00/resource-instance/services-orchestration/{id} | DescribeServicesOrchestration services-orchestration-api
[**ServicesOrchestrationApiListServicesOrchestrations**](ServicesOrchestrationApiAPI.md#ServicesOrchestrationApiListServicesOrchestrations) | **Get** /2022-09-01-00/resource-instance/services-orchestration | ListServicesOrchestrations services-orchestration-api
[**ServicesOrchestrationApiModifyServicesOrchestration**](ServicesOrchestrationApiAPI.md#ServicesOrchestrationApiModifyServicesOrchestration) | **Patch** /2022-09-01-00/resource-instance/services-orchestration/{id} | ModifyServicesOrchestration services-orchestration-api



## ServicesOrchestrationApiCreateServicesOrchestration

> CreateServicesOrchestrationResponseBody ServicesOrchestrationApiCreateServicesOrchestration(ctx).CreateServicesOrchestrationRequest2(createServicesOrchestrationRequest2).Execute()

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
	createServicesOrchestrationRequest2 := *openapiclient.NewCreateServicesOrchestrationRequest2("Eum blanditiis animi.") // CreateServicesOrchestrationRequest2 | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.ServicesOrchestrationApiAPI.ServicesOrchestrationApiCreateServicesOrchestration(context.Background()).CreateServicesOrchestrationRequest2(createServicesOrchestrationRequest2).Execute()
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
 **createServicesOrchestrationRequest2** | [**CreateServicesOrchestrationRequest2**](CreateServicesOrchestrationRequest2.md) |  | 

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


## ServicesOrchestrationApiListServicesOrchestrations

> []DescribeServicesOrchestrationResult ServicesOrchestrationApiListServicesOrchestrations(ctx).Execute()

ListServicesOrchestrations services-orchestration-api

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
	resp, r, err := apiClient.ServicesOrchestrationApiAPI.ServicesOrchestrationApiListServicesOrchestrations(context.Background()).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `ServicesOrchestrationApiAPI.ServicesOrchestrationApiListServicesOrchestrations``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `ServicesOrchestrationApiListServicesOrchestrations`: []DescribeServicesOrchestrationResult
	fmt.Fprintf(os.Stdout, "Response from `ServicesOrchestrationApiAPI.ServicesOrchestrationApiListServicesOrchestrations`: %v\n", resp)
}
```

### Path Parameters

This endpoint does not need any parameter.

### Other Parameters

Other parameters are passed through a pointer to a apiServicesOrchestrationApiListServicesOrchestrationsRequest struct via the builder pattern


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

> ServicesOrchestrationApiModifyServicesOrchestration(ctx, id).ModifyServicesOrchestrationRequest2(modifyServicesOrchestrationRequest2).Execute()

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
	modifyServicesOrchestrationRequest2 := *openapiclient.NewModifyServicesOrchestrationRequest2("Cum voluptas commodi.") // ModifyServicesOrchestrationRequest2 | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	r, err := apiClient.ServicesOrchestrationApiAPI.ServicesOrchestrationApiModifyServicesOrchestration(context.Background(), id).ModifyServicesOrchestrationRequest2(modifyServicesOrchestrationRequest2).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `ServicesOrchestrationApiAPI.ServicesOrchestrationApiModifyServicesOrchestration``: %v\n", err)
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

Other parameters are passed through a pointer to a apiServicesOrchestrationApiModifyServicesOrchestrationRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **modifyServicesOrchestrationRequest2** | [**ModifyServicesOrchestrationRequest2**](ModifyServicesOrchestrationRequest2.md) |  | 

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

