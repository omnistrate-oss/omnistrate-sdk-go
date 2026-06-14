# \CustomWorkflowApiAPI

All URIs are relative to *https://api.omnistrate.cloud*

Method | HTTP request | Description
------------- | ------------- | -------------
[**CustomWorkflowApiCreateCustomWorkflow**](CustomWorkflowApiAPI.md#CustomWorkflowApiCreateCustomWorkflow) | **Post** /2022-09-01-00/service/{serviceId}/product-tier/{productTierId}/custom-workflow | CreateCustomWorkflow custom-workflow-api
[**CustomWorkflowApiDeleteCustomWorkflow**](CustomWorkflowApiAPI.md#CustomWorkflowApiDeleteCustomWorkflow) | **Delete** /2022-09-01-00/service/{serviceId}/custom-workflow/{id} | DeleteCustomWorkflow custom-workflow-api
[**CustomWorkflowApiDescribeCustomWorkflow**](CustomWorkflowApiAPI.md#CustomWorkflowApiDescribeCustomWorkflow) | **Get** /2022-09-01-00/service/{serviceId}/custom-workflow/{id} | DescribeCustomWorkflow custom-workflow-api
[**CustomWorkflowApiListCustomWorkflows**](CustomWorkflowApiAPI.md#CustomWorkflowApiListCustomWorkflows) | **Get** /2022-09-01-00/service/{serviceId}/product-tier/{productTierId}/custom-workflow | ListCustomWorkflows custom-workflow-api
[**CustomWorkflowApiUpdateCustomWorkflow**](CustomWorkflowApiAPI.md#CustomWorkflowApiUpdateCustomWorkflow) | **Patch** /2022-09-01-00/service/{serviceId}/custom-workflow/{id} | UpdateCustomWorkflow custom-workflow-api



## CustomWorkflowApiCreateCustomWorkflow

> string CustomWorkflowApiCreateCustomWorkflow(ctx, serviceId, productTierId).CreateCustomWorkflowRequest2(createCustomWorkflowRequest2).Execute()

CreateCustomWorkflow custom-workflow-api

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
	serviceId := "s-12345678" // string | The ID of the service that this workflow belongs to
	productTierId := "pt-12345678" // string | The product tier ID that this workflow belongs to
	createCustomWorkflowRequest2 := *openapiclient.NewCreateCustomWorkflowRequest2("Backup now", interface{}({"entrypoint":"backup","templates":[{"dag":{"tasks":[{"name":"run-backup","template":"backup-cluster"}]},"name":"backup"}]}), "SWITCHOVER") // CreateCustomWorkflowRequest2 | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.CustomWorkflowApiAPI.CustomWorkflowApiCreateCustomWorkflow(context.Background(), serviceId, productTierId).CreateCustomWorkflowRequest2(createCustomWorkflowRequest2).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `CustomWorkflowApiAPI.CustomWorkflowApiCreateCustomWorkflow``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `CustomWorkflowApiCreateCustomWorkflow`: string
	fmt.Fprintf(os.Stdout, "Response from `CustomWorkflowApiAPI.CustomWorkflowApiCreateCustomWorkflow`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**serviceId** | **string** | The ID of the service that this workflow belongs to | 
**productTierId** | **string** | The product tier ID that this workflow belongs to | 

### Other Parameters

Other parameters are passed through a pointer to a apiCustomWorkflowApiCreateCustomWorkflowRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


 **createCustomWorkflowRequest2** | [**CreateCustomWorkflowRequest2**](CreateCustomWorkflowRequest2.md) |  | 

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


## CustomWorkflowApiDeleteCustomWorkflow

> CustomWorkflowApiDeleteCustomWorkflow(ctx, serviceId, id).Execute()

DeleteCustomWorkflow custom-workflow-api

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
	serviceId := "s-12345678" // string | The ID of the service that this workflow belongs to
	id := "cwt-12345678" // string | The custom workflow ID

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	r, err := apiClient.CustomWorkflowApiAPI.CustomWorkflowApiDeleteCustomWorkflow(context.Background(), serviceId, id).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `CustomWorkflowApiAPI.CustomWorkflowApiDeleteCustomWorkflow``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**serviceId** | **string** | The ID of the service that this workflow belongs to | 
**id** | **string** | The custom workflow ID | 

### Other Parameters

Other parameters are passed through a pointer to a apiCustomWorkflowApiDeleteCustomWorkflowRequest struct via the builder pattern


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


## CustomWorkflowApiDescribeCustomWorkflow

> DescribeCustomWorkflowResult CustomWorkflowApiDescribeCustomWorkflow(ctx, serviceId, id).ProductTierVersion(productTierVersion).ProductTierId(productTierId).Execute()

DescribeCustomWorkflow custom-workflow-api

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
	serviceId := "s-12345678" // string | The ID of the service that this workflow belongs to
	id := "cwt-12345678" // string | The custom workflow ID
	productTierVersion := "Id provident ea repellendus non nihil quam." // string | Product tier version of the workflow to describe. If not specified, the latest version is described. (optional)
	productTierId := "pt-12345678" // string | Product tier ID of the workflow to describe. Needs to be specified in combination with the product tier version (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.CustomWorkflowApiAPI.CustomWorkflowApiDescribeCustomWorkflow(context.Background(), serviceId, id).ProductTierVersion(productTierVersion).ProductTierId(productTierId).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `CustomWorkflowApiAPI.CustomWorkflowApiDescribeCustomWorkflow``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `CustomWorkflowApiDescribeCustomWorkflow`: DescribeCustomWorkflowResult
	fmt.Fprintf(os.Stdout, "Response from `CustomWorkflowApiAPI.CustomWorkflowApiDescribeCustomWorkflow`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**serviceId** | **string** | The ID of the service that this workflow belongs to | 
**id** | **string** | The custom workflow ID | 

### Other Parameters

Other parameters are passed through a pointer to a apiCustomWorkflowApiDescribeCustomWorkflowRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


 **productTierVersion** | **string** | Product tier version of the workflow to describe. If not specified, the latest version is described. | 
 **productTierId** | **string** | Product tier ID of the workflow to describe. Needs to be specified in combination with the product tier version | 

### Return type

[**DescribeCustomWorkflowResult**](DescribeCustomWorkflowResult.md)

### Authorization

[api_key_header_Authorization](../README.md#api_key_header_Authorization)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json, application/vnd.goa.error

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## CustomWorkflowApiListCustomWorkflows

> ListCustomWorkflowsResult CustomWorkflowApiListCustomWorkflows(ctx, serviceId, productTierId).ProductTierVersion(productTierVersion).Execute()

ListCustomWorkflows custom-workflow-api

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
	serviceId := "s-12345678" // string | The ID of the service that the workflows belong to
	productTierId := "pt-12345678" // string | The product tier ID
	productTierVersion := "Ullam dolores molestias." // string | Product tier version of the workflows to describe. If not specified, the latest version is described. (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.CustomWorkflowApiAPI.CustomWorkflowApiListCustomWorkflows(context.Background(), serviceId, productTierId).ProductTierVersion(productTierVersion).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `CustomWorkflowApiAPI.CustomWorkflowApiListCustomWorkflows``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `CustomWorkflowApiListCustomWorkflows`: ListCustomWorkflowsResult
	fmt.Fprintf(os.Stdout, "Response from `CustomWorkflowApiAPI.CustomWorkflowApiListCustomWorkflows`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**serviceId** | **string** | The ID of the service that the workflows belong to | 
**productTierId** | **string** | The product tier ID | 

### Other Parameters

Other parameters are passed through a pointer to a apiCustomWorkflowApiListCustomWorkflowsRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


 **productTierVersion** | **string** | Product tier version of the workflows to describe. If not specified, the latest version is described. | 

### Return type

[**ListCustomWorkflowsResult**](ListCustomWorkflowsResult.md)

### Authorization

[api_key_header_Authorization](../README.md#api_key_header_Authorization)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json, application/vnd.goa.error

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## CustomWorkflowApiUpdateCustomWorkflow

> CustomWorkflowApiUpdateCustomWorkflow(ctx, serviceId, id).UpdateCustomWorkflowRequest2(updateCustomWorkflowRequest2).Execute()

UpdateCustomWorkflow custom-workflow-api

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
	serviceId := "s-12345678" // string | The ID of the service that this workflow belongs to
	id := "cwt-12345678" // string | The custom workflow ID
	updateCustomWorkflowRequest2 := *openapiclient.NewUpdateCustomWorkflowRequest2() // UpdateCustomWorkflowRequest2 | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	r, err := apiClient.CustomWorkflowApiAPI.CustomWorkflowApiUpdateCustomWorkflow(context.Background(), serviceId, id).UpdateCustomWorkflowRequest2(updateCustomWorkflowRequest2).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `CustomWorkflowApiAPI.CustomWorkflowApiUpdateCustomWorkflow``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**serviceId** | **string** | The ID of the service that this workflow belongs to | 
**id** | **string** | The custom workflow ID | 

### Other Parameters

Other parameters are passed through a pointer to a apiCustomWorkflowApiUpdateCustomWorkflowRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


 **updateCustomWorkflowRequest2** | [**UpdateCustomWorkflowRequest2**](UpdateCustomWorkflowRequest2.md) |  | 

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

