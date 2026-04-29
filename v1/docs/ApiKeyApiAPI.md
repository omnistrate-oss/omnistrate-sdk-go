# \ApiKeyApiAPI

All URIs are relative to *https://api.omnistrate.cloud*

Method | HTTP request | Description
------------- | ------------- | -------------
[**ApiKeyApiCreateAPIKey**](ApiKeyApiAPI.md#ApiKeyApiCreateAPIKey) | **Post** /2022-09-01-00/api-key | CreateAPIKey api-key-api
[**ApiKeyApiDeleteAPIKey**](ApiKeyApiAPI.md#ApiKeyApiDeleteAPIKey) | **Delete** /2022-09-01-00/api-key/{id} | DeleteAPIKey api-key-api
[**ApiKeyApiDescribeAPIKey**](ApiKeyApiAPI.md#ApiKeyApiDescribeAPIKey) | **Get** /2022-09-01-00/api-key/{id} | DescribeAPIKey api-key-api
[**ApiKeyApiListAPIKeys**](ApiKeyApiAPI.md#ApiKeyApiListAPIKeys) | **Get** /2022-09-01-00/api-key | ListAPIKeys api-key-api
[**ApiKeyApiRevokeAPIKey**](ApiKeyApiAPI.md#ApiKeyApiRevokeAPIKey) | **Post** /2022-09-01-00/api-key/{id}:revoke | RevokeAPIKey api-key-api
[**ApiKeyApiUpdateAPIKeyMetadata**](ApiKeyApiAPI.md#ApiKeyApiUpdateAPIKeyMetadata) | **Patch** /2022-09-01-00/api-key/{id} | UpdateAPIKeyMetadata api-key-api



## ApiKeyApiCreateAPIKey

> CreateAPIKeyResult ApiKeyApiCreateAPIKey(ctx).CreateAPIKeyRequest2(createAPIKeyRequest2).Execute()

CreateAPIKey api-key-api

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
	createAPIKeyRequest2 := *openapiclient.NewCreateAPIKeyRequest2("ci-pipeline-prod", "root|editor|reader|service_editor|service_reader|admin|service_operator") // CreateAPIKeyRequest2 | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.ApiKeyApiAPI.ApiKeyApiCreateAPIKey(context.Background()).CreateAPIKeyRequest2(createAPIKeyRequest2).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `ApiKeyApiAPI.ApiKeyApiCreateAPIKey``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `ApiKeyApiCreateAPIKey`: CreateAPIKeyResult
	fmt.Fprintf(os.Stdout, "Response from `ApiKeyApiAPI.ApiKeyApiCreateAPIKey`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiApiKeyApiCreateAPIKeyRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **createAPIKeyRequest2** | [**CreateAPIKeyRequest2**](CreateAPIKeyRequest2.md) |  | 

### Return type

[**CreateAPIKeyResult**](CreateAPIKeyResult.md)

### Authorization

[api_key_header_Authorization](../README.md#api_key_header_Authorization)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json, application/vnd.goa.error

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## ApiKeyApiDeleteAPIKey

> RevokeAPIKeyResult ApiKeyApiDeleteAPIKey(ctx, id).Execute()

DeleteAPIKey api-key-api

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
	id := "apikey-12345678" // string | The API key ID.

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.ApiKeyApiAPI.ApiKeyApiDeleteAPIKey(context.Background(), id).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `ApiKeyApiAPI.ApiKeyApiDeleteAPIKey``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `ApiKeyApiDeleteAPIKey`: RevokeAPIKeyResult
	fmt.Fprintf(os.Stdout, "Response from `ApiKeyApiAPI.ApiKeyApiDeleteAPIKey`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **string** | The API key ID. | 

### Other Parameters

Other parameters are passed through a pointer to a apiApiKeyApiDeleteAPIKeyRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

[**RevokeAPIKeyResult**](RevokeAPIKeyResult.md)

### Authorization

[api_key_header_Authorization](../README.md#api_key_header_Authorization)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json, application/vnd.goa.error

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## ApiKeyApiDescribeAPIKey

> DescribeAPIKeyResult ApiKeyApiDescribeAPIKey(ctx, id).Execute()

DescribeAPIKey api-key-api

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
	id := "apikey-12345678" // string | The API key ID.

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.ApiKeyApiAPI.ApiKeyApiDescribeAPIKey(context.Background(), id).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `ApiKeyApiAPI.ApiKeyApiDescribeAPIKey``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `ApiKeyApiDescribeAPIKey`: DescribeAPIKeyResult
	fmt.Fprintf(os.Stdout, "Response from `ApiKeyApiAPI.ApiKeyApiDescribeAPIKey`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **string** | The API key ID. | 

### Other Parameters

Other parameters are passed through a pointer to a apiApiKeyApiDescribeAPIKeyRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

[**DescribeAPIKeyResult**](DescribeAPIKeyResult.md)

### Authorization

[api_key_header_Authorization](../README.md#api_key_header_Authorization)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json, application/vnd.goa.error

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## ApiKeyApiListAPIKeys

> ListAPIKeysResult ApiKeyApiListAPIKeys(ctx).Execute()

ListAPIKeys api-key-api

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
	resp, r, err := apiClient.ApiKeyApiAPI.ApiKeyApiListAPIKeys(context.Background()).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `ApiKeyApiAPI.ApiKeyApiListAPIKeys``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `ApiKeyApiListAPIKeys`: ListAPIKeysResult
	fmt.Fprintf(os.Stdout, "Response from `ApiKeyApiAPI.ApiKeyApiListAPIKeys`: %v\n", resp)
}
```

### Path Parameters

This endpoint does not need any parameter.

### Other Parameters

Other parameters are passed through a pointer to a apiApiKeyApiListAPIKeysRequest struct via the builder pattern


### Return type

[**ListAPIKeysResult**](ListAPIKeysResult.md)

### Authorization

[api_key_header_Authorization](../README.md#api_key_header_Authorization)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json, application/vnd.goa.error

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## ApiKeyApiRevokeAPIKey

> RevokeAPIKeyResult ApiKeyApiRevokeAPIKey(ctx, id).Execute()

RevokeAPIKey api-key-api

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
	id := "apikey-12345678" // string | The API key ID.

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.ApiKeyApiAPI.ApiKeyApiRevokeAPIKey(context.Background(), id).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `ApiKeyApiAPI.ApiKeyApiRevokeAPIKey``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `ApiKeyApiRevokeAPIKey`: RevokeAPIKeyResult
	fmt.Fprintf(os.Stdout, "Response from `ApiKeyApiAPI.ApiKeyApiRevokeAPIKey`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **string** | The API key ID. | 

### Other Parameters

Other parameters are passed through a pointer to a apiApiKeyApiRevokeAPIKeyRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

[**RevokeAPIKeyResult**](RevokeAPIKeyResult.md)

### Authorization

[api_key_header_Authorization](../README.md#api_key_header_Authorization)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json, application/vnd.goa.error

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## ApiKeyApiUpdateAPIKeyMetadata

> UpdateAPIKeyMetadataResult ApiKeyApiUpdateAPIKeyMetadata(ctx, id).UpdateAPIKeyMetadataRequest2(updateAPIKeyMetadataRequest2).Execute()

UpdateAPIKeyMetadata api-key-api

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
	id := "apikey-12345678" // string | The API key ID.
	updateAPIKeyMetadataRequest2 := *openapiclient.NewUpdateAPIKeyMetadataRequest2() // UpdateAPIKeyMetadataRequest2 | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.ApiKeyApiAPI.ApiKeyApiUpdateAPIKeyMetadata(context.Background(), id).UpdateAPIKeyMetadataRequest2(updateAPIKeyMetadataRequest2).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `ApiKeyApiAPI.ApiKeyApiUpdateAPIKeyMetadata``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `ApiKeyApiUpdateAPIKeyMetadata`: UpdateAPIKeyMetadataResult
	fmt.Fprintf(os.Stdout, "Response from `ApiKeyApiAPI.ApiKeyApiUpdateAPIKeyMetadata`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **string** | The API key ID. | 

### Other Parameters

Other parameters are passed through a pointer to a apiApiKeyApiUpdateAPIKeyMetadataRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **updateAPIKeyMetadataRequest2** | [**UpdateAPIKeyMetadataRequest2**](UpdateAPIKeyMetadataRequest2.md) |  | 

### Return type

[**UpdateAPIKeyMetadataResult**](UpdateAPIKeyMetadataResult.md)

### Authorization

[api_key_header_Authorization](../README.md#api_key_header_Authorization)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json, application/vnd.goa.error

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

