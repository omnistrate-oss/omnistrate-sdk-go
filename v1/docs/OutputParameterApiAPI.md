# \OutputParameterApiAPI

All URIs are relative to *https://api.omnistrate.cloud*

Method | HTTP request | Description
------------- | ------------- | -------------
[**OutputParameterApiCreateOutputParameter**](OutputParameterApiAPI.md#OutputParameterApiCreateOutputParameter) | **Post** /2022-09-01-00/service/{serviceId}/output-parameter | CreateOutputParameter output-parameter-api
[**OutputParameterApiDeleteOutputParameter**](OutputParameterApiAPI.md#OutputParameterApiDeleteOutputParameter) | **Delete** /2022-09-01-00/service/{serviceId}/output-parameter/{id} | DeleteOutputParameter output-parameter-api
[**OutputParameterApiDescribeOutputParameter**](OutputParameterApiAPI.md#OutputParameterApiDescribeOutputParameter) | **Get** /2022-09-01-00/service/{serviceId}/output-parameter/{id} | DescribeOutputParameter output-parameter-api
[**OutputParameterApiListOutputParameter**](OutputParameterApiAPI.md#OutputParameterApiListOutputParameter) | **Get** /2022-09-01-00/service/{serviceId}/resource/{resourceId}/output-parameter | ListOutputParameter output-parameter-api
[**OutputParameterApiUpdateOutputParameter**](OutputParameterApiAPI.md#OutputParameterApiUpdateOutputParameter) | **Patch** /2022-09-01-00/service/{serviceId}/output-parameter/{id} | UpdateOutputParameter output-parameter-api



## OutputParameterApiCreateOutputParameter

> string OutputParameterApiCreateOutputParameter(ctx, serviceId).CreateOutputParameterRequestBody(createOutputParameterRequestBody).Execute()

CreateOutputParameter output-parameter-api

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
	serviceId := "s-12345678" // string | The ID of the service that this output parameter belongs to
	createOutputParameterRequestBody := *openapiclient.NewCreateOutputParameterRequestBody("Username of the user created in the target system", "username", "Username", "r-12345678") // CreateOutputParameterRequestBody | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.OutputParameterApiAPI.OutputParameterApiCreateOutputParameter(context.Background(), serviceId).CreateOutputParameterRequestBody(createOutputParameterRequestBody).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `OutputParameterApiAPI.OutputParameterApiCreateOutputParameter``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `OutputParameterApiCreateOutputParameter`: string
	fmt.Fprintf(os.Stdout, "Response from `OutputParameterApiAPI.OutputParameterApiCreateOutputParameter`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**serviceId** | **string** | The ID of the service that this output parameter belongs to | 

### Other Parameters

Other parameters are passed through a pointer to a apiOutputParameterApiCreateOutputParameterRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **createOutputParameterRequestBody** | [**CreateOutputParameterRequestBody**](CreateOutputParameterRequestBody.md) |  | 

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


## OutputParameterApiDeleteOutputParameter

> OutputParameterApiDeleteOutputParameter(ctx, serviceId, id).Execute()

DeleteOutputParameter output-parameter-api

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
	serviceId := "s-12345678" // string | The ID of the service that this output parameter belongs to
	id := "op-12345678" // string | The ID of the output parameter

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	r, err := apiClient.OutputParameterApiAPI.OutputParameterApiDeleteOutputParameter(context.Background(), serviceId, id).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `OutputParameterApiAPI.OutputParameterApiDeleteOutputParameter``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**serviceId** | **string** | The ID of the service that this output parameter belongs to | 
**id** | **string** | The ID of the output parameter | 

### Other Parameters

Other parameters are passed through a pointer to a apiOutputParameterApiDeleteOutputParameterRequest struct via the builder pattern


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


## OutputParameterApiDescribeOutputParameter

> DescribeOutputParameterResult OutputParameterApiDescribeOutputParameter(ctx, serviceId, id).ProductTierVersion(productTierVersion).ProductTierId(productTierId).Execute()

DescribeOutputParameter output-parameter-api

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
	serviceId := "s-12345678" // string | The ID of the service that this output parameter belongs to
	id := "op-12345678" // string | The ID of the output parameter
	productTierVersion := "Voluptas voluptas impedit." // string | Product tier version of the resource to describe. If not specified, the latest version is described. (optional)
	productTierId := "Beatae beatae." // string | ProductTierId of the resource to describe. Needs to specified in combination with the product tier version (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.OutputParameterApiAPI.OutputParameterApiDescribeOutputParameter(context.Background(), serviceId, id).ProductTierVersion(productTierVersion).ProductTierId(productTierId).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `OutputParameterApiAPI.OutputParameterApiDescribeOutputParameter``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `OutputParameterApiDescribeOutputParameter`: DescribeOutputParameterResult
	fmt.Fprintf(os.Stdout, "Response from `OutputParameterApiAPI.OutputParameterApiDescribeOutputParameter`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**serviceId** | **string** | The ID of the service that this output parameter belongs to | 
**id** | **string** | The ID of the output parameter | 

### Other Parameters

Other parameters are passed through a pointer to a apiOutputParameterApiDescribeOutputParameterRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


 **productTierVersion** | **string** | Product tier version of the resource to describe. If not specified, the latest version is described. | 
 **productTierId** | **string** | ProductTierId of the resource to describe. Needs to specified in combination with the product tier version | 

### Return type

[**DescribeOutputParameterResult**](DescribeOutputParameterResult.md)

### Authorization

[api_key_header_Authorization](../README.md#api_key_header_Authorization)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json, application/vnd.goa.error

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## OutputParameterApiListOutputParameter

> ListOutputParametersResult OutputParameterApiListOutputParameter(ctx, serviceId, resourceId).ProductTierVersion(productTierVersion).ProductTierId(productTierId).Execute()

ListOutputParameter output-parameter-api

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
	serviceId := "s-12345678" // string | The ID of the service that this output parameter belongs to
	resourceId := "r-12345678" // string | The ID of the resource that this output parameter belongs to
	productTierVersion := "Corporis dolorum." // string | Product tier version of the resource to describe. If not specified, the latest version is described. (optional)
	productTierId := "Beatae beatae." // string | ProductTierId of the resource to describe. Needs to specified in combination with the product tier version (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.OutputParameterApiAPI.OutputParameterApiListOutputParameter(context.Background(), serviceId, resourceId).ProductTierVersion(productTierVersion).ProductTierId(productTierId).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `OutputParameterApiAPI.OutputParameterApiListOutputParameter``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `OutputParameterApiListOutputParameter`: ListOutputParametersResult
	fmt.Fprintf(os.Stdout, "Response from `OutputParameterApiAPI.OutputParameterApiListOutputParameter`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**serviceId** | **string** | The ID of the service that this output parameter belongs to | 
**resourceId** | **string** | The ID of the resource that this output parameter belongs to | 

### Other Parameters

Other parameters are passed through a pointer to a apiOutputParameterApiListOutputParameterRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


 **productTierVersion** | **string** | Product tier version of the resource to describe. If not specified, the latest version is described. | 
 **productTierId** | **string** | ProductTierId of the resource to describe. Needs to specified in combination with the product tier version | 

### Return type

[**ListOutputParametersResult**](ListOutputParametersResult.md)

### Authorization

[api_key_header_Authorization](../README.md#api_key_header_Authorization)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json, application/vnd.goa.error

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## OutputParameterApiUpdateOutputParameter

> OutputParameterApiUpdateOutputParameter(ctx, serviceId, id).UpdateOutputParameterRequestBody(updateOutputParameterRequestBody).Execute()

UpdateOutputParameter output-parameter-api

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
	serviceId := "s-12345678" // string | The ID of the service that this output parameter belongs to
	id := "op-12345678" // string | The ID of the output parameter
	updateOutputParameterRequestBody := *openapiclient.NewUpdateOutputParameterRequestBody() // UpdateOutputParameterRequestBody | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	r, err := apiClient.OutputParameterApiAPI.OutputParameterApiUpdateOutputParameter(context.Background(), serviceId, id).UpdateOutputParameterRequestBody(updateOutputParameterRequestBody).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `OutputParameterApiAPI.OutputParameterApiUpdateOutputParameter``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**serviceId** | **string** | The ID of the service that this output parameter belongs to | 
**id** | **string** | The ID of the output parameter | 

### Other Parameters

Other parameters are passed through a pointer to a apiOutputParameterApiUpdateOutputParameterRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


 **updateOutputParameterRequestBody** | [**UpdateOutputParameterRequestBody**](UpdateOutputParameterRequestBody.md) |  | 

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

