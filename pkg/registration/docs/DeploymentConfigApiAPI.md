# \DeploymentConfigApiAPI

All URIs are relative to *https://api.omnistrate.cloud*

Method | HTTP request | Description
------------- | ------------- | -------------
[**DeploymentConfigApiCreateDeploymentConfig**](DeploymentConfigApiAPI.md#DeploymentConfigApiCreateDeploymentConfig) | **Post** /2022-09-01-00/deployment-config | CreateDeploymentConfig deployment-config-api
[**DeploymentConfigApiDeleteDeploymentConfig**](DeploymentConfigApiAPI.md#DeploymentConfigApiDeleteDeploymentConfig) | **Delete** /2022-09-01-00/deployment-config/{id} | DeleteDeploymentConfig deployment-config-api
[**DeploymentConfigApiDescribeDeploymentConfig**](DeploymentConfigApiAPI.md#DeploymentConfigApiDescribeDeploymentConfig) | **Get** /2022-09-01-00/deployment-config/{id} | DescribeDeploymentConfig deployment-config-api
[**DeploymentConfigApiListDeploymentConfigs**](DeploymentConfigApiAPI.md#DeploymentConfigApiListDeploymentConfigs) | **Get** /2022-09-01-00/deployment-config | ListDeploymentConfigs deployment-config-api
[**DeploymentConfigApiUpdateDeploymentConfig**](DeploymentConfigApiAPI.md#DeploymentConfigApiUpdateDeploymentConfig) | **Patch** /2022-09-01-00/deployment-config/{id} | UpdateDeploymentConfig deployment-config-api



## DeploymentConfigApiCreateDeploymentConfig

> string DeploymentConfigApiCreateDeploymentConfig(ctx).CreateDeploymentConfigRequestBody(createDeploymentConfigRequestBody).Execute()

CreateDeploymentConfig deployment-config-api

### Example

```go
package main

import (
	"context"
	"fmt"
	"os"
	openapiclient "github.com/omnistrate/omnistrate-sdk-go/registration"
)

func main() {
	createDeploymentConfigRequestBody := *openapiclient.NewCreateDeploymentConfigRequestBody("A production deployment config", "Production") // CreateDeploymentConfigRequestBody | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.DeploymentConfigApiAPI.DeploymentConfigApiCreateDeploymentConfig(context.Background()).CreateDeploymentConfigRequestBody(createDeploymentConfigRequestBody).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `DeploymentConfigApiAPI.DeploymentConfigApiCreateDeploymentConfig``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `DeploymentConfigApiCreateDeploymentConfig`: string
	fmt.Fprintf(os.Stdout, "Response from `DeploymentConfigApiAPI.DeploymentConfigApiCreateDeploymentConfig`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiDeploymentConfigApiCreateDeploymentConfigRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **createDeploymentConfigRequestBody** | [**CreateDeploymentConfigRequestBody**](CreateDeploymentConfigRequestBody.md) |  | 

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


## DeploymentConfigApiDeleteDeploymentConfig

> DeploymentConfigApiDeleteDeploymentConfig(ctx, id).Execute()

DeleteDeploymentConfig deployment-config-api

### Example

```go
package main

import (
	"context"
	"fmt"
	"os"
	openapiclient "github.com/omnistrate/omnistrate-sdk-go/registration"
)

func main() {
	id := "dc-12345678" // string | The deployment configuration ID

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	r, err := apiClient.DeploymentConfigApiAPI.DeploymentConfigApiDeleteDeploymentConfig(context.Background(), id).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `DeploymentConfigApiAPI.DeploymentConfigApiDeleteDeploymentConfig``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **string** | The deployment configuration ID | 

### Other Parameters

Other parameters are passed through a pointer to a apiDeploymentConfigApiDeleteDeploymentConfigRequest struct via the builder pattern


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


## DeploymentConfigApiDescribeDeploymentConfig

> DescribeDeploymentConfigResult DeploymentConfigApiDescribeDeploymentConfig(ctx, id).Execute()

DescribeDeploymentConfig deployment-config-api

### Example

```go
package main

import (
	"context"
	"fmt"
	"os"
	openapiclient "github.com/omnistrate/omnistrate-sdk-go/registration"
)

func main() {
	id := "default" // string | The deployment configuration ID

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.DeploymentConfigApiAPI.DeploymentConfigApiDescribeDeploymentConfig(context.Background(), id).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `DeploymentConfigApiAPI.DeploymentConfigApiDescribeDeploymentConfig``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `DeploymentConfigApiDescribeDeploymentConfig`: DescribeDeploymentConfigResult
	fmt.Fprintf(os.Stdout, "Response from `DeploymentConfigApiAPI.DeploymentConfigApiDescribeDeploymentConfig`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **string** | The deployment configuration ID | 

### Other Parameters

Other parameters are passed through a pointer to a apiDeploymentConfigApiDescribeDeploymentConfigRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

[**DescribeDeploymentConfigResult**](DescribeDeploymentConfigResult.md)

### Authorization

[api_key_header_Authorization](../README.md#api_key_header_Authorization)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json, application/vnd.goa.error

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## DeploymentConfigApiListDeploymentConfigs

> ListServiceEnvironmentsResult DeploymentConfigApiListDeploymentConfigs(ctx).Execute()

ListDeploymentConfigs deployment-config-api

### Example

```go
package main

import (
	"context"
	"fmt"
	"os"
	openapiclient "github.com/omnistrate/omnistrate-sdk-go/registration"
)

func main() {

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.DeploymentConfigApiAPI.DeploymentConfigApiListDeploymentConfigs(context.Background()).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `DeploymentConfigApiAPI.DeploymentConfigApiListDeploymentConfigs``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `DeploymentConfigApiListDeploymentConfigs`: ListServiceEnvironmentsResult
	fmt.Fprintf(os.Stdout, "Response from `DeploymentConfigApiAPI.DeploymentConfigApiListDeploymentConfigs`: %v\n", resp)
}
```

### Path Parameters

This endpoint does not need any parameter.

### Other Parameters

Other parameters are passed through a pointer to a apiDeploymentConfigApiListDeploymentConfigsRequest struct via the builder pattern


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


## DeploymentConfigApiUpdateDeploymentConfig

> DeploymentConfigApiUpdateDeploymentConfig(ctx, id).UpdateDeploymentConfigRequestBody(updateDeploymentConfigRequestBody).Execute()

UpdateDeploymentConfig deployment-config-api

### Example

```go
package main

import (
	"context"
	"fmt"
	"os"
	openapiclient "github.com/omnistrate/omnistrate-sdk-go/registration"
)

func main() {
	updateDeploymentConfigRequestBody := *openapiclient.NewUpdateDeploymentConfigRequestBody() // UpdateDeploymentConfigRequestBody | 
	id := "dc-12345678" // string | The deployment configuration ID

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	r, err := apiClient.DeploymentConfigApiAPI.DeploymentConfigApiUpdateDeploymentConfig(context.Background(), id).UpdateDeploymentConfigRequestBody(updateDeploymentConfigRequestBody).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `DeploymentConfigApiAPI.DeploymentConfigApiUpdateDeploymentConfig``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **string** | The deployment configuration ID | 

### Other Parameters

Other parameters are passed through a pointer to a apiDeploymentConfigApiUpdateDeploymentConfigRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **updateDeploymentConfigRequestBody** | [**UpdateDeploymentConfigRequestBody**](UpdateDeploymentConfigRequestBody.md) |  | 


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

