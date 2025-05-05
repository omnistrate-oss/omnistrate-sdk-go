# \InputParameterApiAPI

All URIs are relative to *https://api.omnistrate.cloud*

Method | HTTP request | Description
------------- | ------------- | -------------
[**InputParameterApiCreateInputParameter**](InputParameterApiAPI.md#InputParameterApiCreateInputParameter) | **Post** /2022-09-01-00/service/{serviceId}/input-parameter | CreateInputParameter input-parameter-api
[**InputParameterApiDeleteInputParameter**](InputParameterApiAPI.md#InputParameterApiDeleteInputParameter) | **Delete** /2022-09-01-00/service/{serviceId}/input-parameter/{id} | DeleteInputParameter input-parameter-api
[**InputParameterApiDescribeInputParameter**](InputParameterApiAPI.md#InputParameterApiDescribeInputParameter) | **Get** /2022-09-01-00/service/{serviceId}/input-parameter/{id} | DescribeInputParameter input-parameter-api
[**InputParameterApiListInputParameter**](InputParameterApiAPI.md#InputParameterApiListInputParameter) | **Get** /2022-09-01-00/service/{serviceId}/resource/{resourceId}/input-parameter | ListInputParameter input-parameter-api
[**InputParameterApiUpdateInputParameter**](InputParameterApiAPI.md#InputParameterApiUpdateInputParameter) | **Patch** /2022-09-01-00/service/{serviceId}/input-parameter/{id} | UpdateInputParameter input-parameter-api



## InputParameterApiCreateInputParameter

> string InputParameterApiCreateInputParameter(ctx, serviceId).CreateInputParameterRequest2(createInputParameterRequest2).Execute()

CreateInputParameter input-parameter-api

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
	createInputParameterRequest2 := *openapiclient.NewCreateInputParameterRequest2("Quas minima non voluptas.", "9", false, "3h8", true, "r-12345678", "Sapiente quaerat illo hic suscipit asperiores.") // CreateInputParameterRequest2 | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.InputParameterApiAPI.InputParameterApiCreateInputParameter(context.Background(), serviceId).CreateInputParameterRequest2(createInputParameterRequest2).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `InputParameterApiAPI.InputParameterApiCreateInputParameter``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `InputParameterApiCreateInputParameter`: string
	fmt.Fprintf(os.Stdout, "Response from `InputParameterApiAPI.InputParameterApiCreateInputParameter`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**serviceId** | **string** | The ID of the service that this output parameter belongs to | 

### Other Parameters

Other parameters are passed through a pointer to a apiInputParameterApiCreateInputParameterRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **createInputParameterRequest2** | [**CreateInputParameterRequest2**](CreateInputParameterRequest2.md) |  | 

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


## InputParameterApiDeleteInputParameter

> InputParameterApiDeleteInputParameter(ctx, serviceId, id).Execute()

DeleteInputParameter input-parameter-api

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
	id := "var-12345678" // string | ID of the input parameter

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	r, err := apiClient.InputParameterApiAPI.InputParameterApiDeleteInputParameter(context.Background(), serviceId, id).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `InputParameterApiAPI.InputParameterApiDeleteInputParameter``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**serviceId** | **string** | The ID of the service that this output parameter belongs to | 
**id** | **string** | ID of the input parameter | 

### Other Parameters

Other parameters are passed through a pointer to a apiInputParameterApiDeleteInputParameterRequest struct via the builder pattern


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


## InputParameterApiDescribeInputParameter

> DescribeInputParameterResult InputParameterApiDescribeInputParameter(ctx, serviceId, id).ProductTierVersion(productTierVersion).ProductTierId(productTierId).Execute()

DescribeInputParameter input-parameter-api

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
	id := "var-12345678" // string | ID of the input parameter
	productTierVersion := "Doloribus qui ea modi." // string | Product tier version of the instance to describe. If not specified, the latest version is described. (optional)
	productTierId := "Beatae beatae." // string | Product tier id of the instance to describe. Needs to specified in combination with the product tier version (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.InputParameterApiAPI.InputParameterApiDescribeInputParameter(context.Background(), serviceId, id).ProductTierVersion(productTierVersion).ProductTierId(productTierId).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `InputParameterApiAPI.InputParameterApiDescribeInputParameter``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `InputParameterApiDescribeInputParameter`: DescribeInputParameterResult
	fmt.Fprintf(os.Stdout, "Response from `InputParameterApiAPI.InputParameterApiDescribeInputParameter`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**serviceId** | **string** | The ID of the service that this output parameter belongs to | 
**id** | **string** | ID of the input parameter | 

### Other Parameters

Other parameters are passed through a pointer to a apiInputParameterApiDescribeInputParameterRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


 **productTierVersion** | **string** | Product tier version of the instance to describe. If not specified, the latest version is described. | 
 **productTierId** | **string** | Product tier id of the instance to describe. Needs to specified in combination with the product tier version | 

### Return type

[**DescribeInputParameterResult**](DescribeInputParameterResult.md)

### Authorization

[api_key_header_Authorization](../README.md#api_key_header_Authorization)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json, application/vnd.goa.error

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## InputParameterApiListInputParameter

> ListInputParametersResult InputParameterApiListInputParameter(ctx, serviceId, resourceId).ProductTierVersion(productTierVersion).ProductTierId(productTierId).Execute()

ListInputParameter input-parameter-api

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
	resourceId := "r-12345678" // string | The ID of the resource that this input parameter belongs to
	productTierVersion := "Placeat cum itaque aut in voluptate eaque." // string | Product tier version of the instance to describe. If not specified, the latest version is described. (optional)
	productTierId := "Beatae beatae." // string | Product tier id of the instance to describe. Needs to specified in combination with the product tier version (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.InputParameterApiAPI.InputParameterApiListInputParameter(context.Background(), serviceId, resourceId).ProductTierVersion(productTierVersion).ProductTierId(productTierId).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `InputParameterApiAPI.InputParameterApiListInputParameter``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `InputParameterApiListInputParameter`: ListInputParametersResult
	fmt.Fprintf(os.Stdout, "Response from `InputParameterApiAPI.InputParameterApiListInputParameter`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**serviceId** | **string** | The ID of the service that this output parameter belongs to | 
**resourceId** | **string** | The ID of the resource that this input parameter belongs to | 

### Other Parameters

Other parameters are passed through a pointer to a apiInputParameterApiListInputParameterRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


 **productTierVersion** | **string** | Product tier version of the instance to describe. If not specified, the latest version is described. | 
 **productTierId** | **string** | Product tier id of the instance to describe. Needs to specified in combination with the product tier version | 

### Return type

[**ListInputParametersResult**](ListInputParametersResult.md)

### Authorization

[api_key_header_Authorization](../README.md#api_key_header_Authorization)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json, application/vnd.goa.error

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## InputParameterApiUpdateInputParameter

> InputParameterApiUpdateInputParameter(ctx, serviceId, id).UpdateInputParameterRequest2(updateInputParameterRequest2).Execute()

UpdateInputParameter input-parameter-api

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
	id := "var-12345678" // string | ID of the input parameter
	updateInputParameterRequest2 := *openapiclient.NewUpdateInputParameterRequest2() // UpdateInputParameterRequest2 | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	r, err := apiClient.InputParameterApiAPI.InputParameterApiUpdateInputParameter(context.Background(), serviceId, id).UpdateInputParameterRequest2(updateInputParameterRequest2).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `InputParameterApiAPI.InputParameterApiUpdateInputParameter``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**serviceId** | **string** | The ID of the service that this output parameter belongs to | 
**id** | **string** | ID of the input parameter | 

### Other Parameters

Other parameters are passed through a pointer to a apiInputParameterApiUpdateInputParameterRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


 **updateInputParameterRequest2** | [**UpdateInputParameterRequest2**](UpdateInputParameterRequest2.md) |  | 

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

