# \IdentityProviderApiAPI

All URIs are relative to *https://api.omnistrate.cloud*

Method | HTTP request | Description
------------- | ------------- | -------------
[**IdentityProviderApiCreateIdentityProvider**](IdentityProviderApiAPI.md#IdentityProviderApiCreateIdentityProvider) | **Post** /2022-09-01-00/identity-provider | CreateIdentityProvider identity-provider-api
[**IdentityProviderApiDeleteIdentityProvider**](IdentityProviderApiAPI.md#IdentityProviderApiDeleteIdentityProvider) | **Delete** /2022-09-01-00/identity-provider/{id} | DeleteIdentityProvider identity-provider-api
[**IdentityProviderApiDescribeIdentityProvider**](IdentityProviderApiAPI.md#IdentityProviderApiDescribeIdentityProvider) | **Get** /2022-09-01-00/identity-provider/{id} | DescribeIdentityProvider identity-provider-api
[**IdentityProviderApiListIdentityProviders**](IdentityProviderApiAPI.md#IdentityProviderApiListIdentityProviders) | **Get** /2022-09-01-00/identity-provider | ListIdentityProviders identity-provider-api
[**IdentityProviderApiUpdateIdentityProvider**](IdentityProviderApiAPI.md#IdentityProviderApiUpdateIdentityProvider) | **Patch** /2022-09-01-00/identity-provider/{id} | UpdateIdentityProvider identity-provider-api
[**IdentityProviderApiVerifyIdentityProvider**](IdentityProviderApiAPI.md#IdentityProviderApiVerifyIdentityProvider) | **Post** /2022-09-01-00/identity-provider/{id} | VerifyIdentityProvider identity-provider-api



## IdentityProviderApiCreateIdentityProvider

> string IdentityProviderApiCreateIdentityProvider(ctx).CreateIdentityProviderRequestBody(createIdentityProviderRequestBody).Execute()

CreateIdentityProvider identity-provider-api

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
	createIdentityProviderRequestBody := *openapiclient.NewCreateIdentityProviderRequestBody("205376496935-vtfpdnseqmjhsynlh0bsufl38k0test.apps.googleusercontent.com", "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", "Google") // CreateIdentityProviderRequestBody | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.IdentityProviderApiAPI.IdentityProviderApiCreateIdentityProvider(context.Background()).CreateIdentityProviderRequestBody(createIdentityProviderRequestBody).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `IdentityProviderApiAPI.IdentityProviderApiCreateIdentityProvider``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `IdentityProviderApiCreateIdentityProvider`: string
	fmt.Fprintf(os.Stdout, "Response from `IdentityProviderApiAPI.IdentityProviderApiCreateIdentityProvider`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiIdentityProviderApiCreateIdentityProviderRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **createIdentityProviderRequestBody** | [**CreateIdentityProviderRequestBody**](CreateIdentityProviderRequestBody.md) |  | 

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


## IdentityProviderApiDeleteIdentityProvider

> IdentityProviderApiDeleteIdentityProvider(ctx, id).Execute()

DeleteIdentityProvider identity-provider-api

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
	id := "idp-12345678" // string | The Identity Provider ID

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	r, err := apiClient.IdentityProviderApiAPI.IdentityProviderApiDeleteIdentityProvider(context.Background(), id).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `IdentityProviderApiAPI.IdentityProviderApiDeleteIdentityProvider``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **string** | The Identity Provider ID | 

### Other Parameters

Other parameters are passed through a pointer to a apiIdentityProviderApiDeleteIdentityProviderRequest struct via the builder pattern


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


## IdentityProviderApiDescribeIdentityProvider

> DescribeIdentityProviderResult IdentityProviderApiDescribeIdentityProvider(ctx, id).Execute()

DescribeIdentityProvider identity-provider-api

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
	id := "idp-12345678" // string | The Identity Provider ID

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.IdentityProviderApiAPI.IdentityProviderApiDescribeIdentityProvider(context.Background(), id).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `IdentityProviderApiAPI.IdentityProviderApiDescribeIdentityProvider``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `IdentityProviderApiDescribeIdentityProvider`: DescribeIdentityProviderResult
	fmt.Fprintf(os.Stdout, "Response from `IdentityProviderApiAPI.IdentityProviderApiDescribeIdentityProvider`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **string** | The Identity Provider ID | 

### Other Parameters

Other parameters are passed through a pointer to a apiIdentityProviderApiDescribeIdentityProviderRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

[**DescribeIdentityProviderResult**](DescribeIdentityProviderResult.md)

### Authorization

[api_key_header_Authorization](../README.md#api_key_header_Authorization)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json, application/vnd.goa.error

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## IdentityProviderApiListIdentityProviders

> ListIdentityProvidersResult IdentityProviderApiListIdentityProviders(ctx).Execute()

ListIdentityProviders identity-provider-api

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
	resp, r, err := apiClient.IdentityProviderApiAPI.IdentityProviderApiListIdentityProviders(context.Background()).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `IdentityProviderApiAPI.IdentityProviderApiListIdentityProviders``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `IdentityProviderApiListIdentityProviders`: ListIdentityProvidersResult
	fmt.Fprintf(os.Stdout, "Response from `IdentityProviderApiAPI.IdentityProviderApiListIdentityProviders`: %v\n", resp)
}
```

### Path Parameters

This endpoint does not need any parameter.

### Other Parameters

Other parameters are passed through a pointer to a apiIdentityProviderApiListIdentityProvidersRequest struct via the builder pattern


### Return type

[**ListIdentityProvidersResult**](ListIdentityProvidersResult.md)

### Authorization

[api_key_header_Authorization](../README.md#api_key_header_Authorization)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json, application/vnd.goa.error

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## IdentityProviderApiUpdateIdentityProvider

> IdentityProviderApiUpdateIdentityProvider(ctx, id).UpdateIdentityProviderRequestBody(updateIdentityProviderRequestBody).Execute()

UpdateIdentityProvider identity-provider-api

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
	id := "idp-12345678" // string | The Identity Provider ID
	updateIdentityProviderRequestBody := *openapiclient.NewUpdateIdentityProviderRequestBody() // UpdateIdentityProviderRequestBody | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	r, err := apiClient.IdentityProviderApiAPI.IdentityProviderApiUpdateIdentityProvider(context.Background(), id).UpdateIdentityProviderRequestBody(updateIdentityProviderRequestBody).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `IdentityProviderApiAPI.IdentityProviderApiUpdateIdentityProvider``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **string** | The Identity Provider ID | 

### Other Parameters

Other parameters are passed through a pointer to a apiIdentityProviderApiUpdateIdentityProviderRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **updateIdentityProviderRequestBody** | [**UpdateIdentityProviderRequestBody**](UpdateIdentityProviderRequestBody.md) |  | 

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


## IdentityProviderApiVerifyIdentityProvider

> DescribeIdentityProviderResult IdentityProviderApiVerifyIdentityProvider(ctx, id).Execute()

VerifyIdentityProvider identity-provider-api

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
	id := "idp-12345678" // string | The Identity Provider ID

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.IdentityProviderApiAPI.IdentityProviderApiVerifyIdentityProvider(context.Background(), id).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `IdentityProviderApiAPI.IdentityProviderApiVerifyIdentityProvider``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `IdentityProviderApiVerifyIdentityProvider`: DescribeIdentityProviderResult
	fmt.Fprintf(os.Stdout, "Response from `IdentityProviderApiAPI.IdentityProviderApiVerifyIdentityProvider`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **string** | The Identity Provider ID | 

### Other Parameters

Other parameters are passed through a pointer to a apiIdentityProviderApiVerifyIdentityProviderRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

[**DescribeIdentityProviderResult**](DescribeIdentityProviderResult.md)

### Authorization

[api_key_header_Authorization](../README.md#api_key_header_Authorization)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json, application/vnd.goa.error

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

