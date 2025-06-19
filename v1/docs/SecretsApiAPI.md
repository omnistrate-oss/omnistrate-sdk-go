# \SecretsApiAPI

All URIs are relative to *https://api.omnistrate.cloud*

Method | HTTP request | Description
------------- | ------------- | -------------
[**SecretsApiDeleteSecret**](SecretsApiAPI.md#SecretsApiDeleteSecret) | **Delete** /2022-09-01-00/environmentType/{environmentType}/secret/{name} | DeleteSecret secrets-api
[**SecretsApiGetSecret**](SecretsApiAPI.md#SecretsApiGetSecret) | **Get** /2022-09-01-00/environmentType/{environmentType}/secret/{name} | GetSecret secrets-api
[**SecretsApiListSecrets**](SecretsApiAPI.md#SecretsApiListSecrets) | **Get** /2022-09-01-00/environmentType/{environmentType}/secrets | ListSecrets secrets-api
[**SecretsApiSetSecret**](SecretsApiAPI.md#SecretsApiSetSecret) | **Put** /2022-09-01-00/environmentType/{environmentType}/secret/{name} | SetSecret secrets-api



## SecretsApiDeleteSecret

> SecretsApiDeleteSecret(ctx, environmentType, name).Execute()

DeleteSecret secrets-api

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
	environmentType := "DEV" // string | The environment type for the secret
	name := "my-secret" // string | Name of the secret

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	r, err := apiClient.SecretsApiAPI.SecretsApiDeleteSecret(context.Background(), environmentType, name).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `SecretsApiAPI.SecretsApiDeleteSecret``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**environmentType** | **string** | The environment type for the secret | 
**name** | **string** | Name of the secret | 

### Other Parameters

Other parameters are passed through a pointer to a apiSecretsApiDeleteSecretRequest struct via the builder pattern


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


## SecretsApiGetSecret

> GetSecretResult SecretsApiGetSecret(ctx, environmentType, name).Execute()

GetSecret secrets-api

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
	environmentType := "DEV" // string | The environment type for the secret
	name := "my-secret" // string | Name of the secret

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.SecretsApiAPI.SecretsApiGetSecret(context.Background(), environmentType, name).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `SecretsApiAPI.SecretsApiGetSecret``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `SecretsApiGetSecret`: GetSecretResult
	fmt.Fprintf(os.Stdout, "Response from `SecretsApiAPI.SecretsApiGetSecret`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**environmentType** | **string** | The environment type for the secret | 
**name** | **string** | Name of the secret | 

### Other Parameters

Other parameters are passed through a pointer to a apiSecretsApiGetSecretRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------



### Return type

[**GetSecretResult**](GetSecretResult.md)

### Authorization

[api_key_header_Authorization](../README.md#api_key_header_Authorization)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json, application/vnd.goa.error

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## SecretsApiListSecrets

> ListSecretsResult SecretsApiListSecrets(ctx, environmentType).Execute()

ListSecrets secrets-api

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
	environmentType := "DEV" // string | The environment type for the secret

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.SecretsApiAPI.SecretsApiListSecrets(context.Background(), environmentType).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `SecretsApiAPI.SecretsApiListSecrets``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `SecretsApiListSecrets`: ListSecretsResult
	fmt.Fprintf(os.Stdout, "Response from `SecretsApiAPI.SecretsApiListSecrets`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**environmentType** | **string** | The environment type for the secret | 

### Other Parameters

Other parameters are passed through a pointer to a apiSecretsApiListSecretsRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

[**ListSecretsResult**](ListSecretsResult.md)

### Authorization

[api_key_header_Authorization](../README.md#api_key_header_Authorization)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json, application/vnd.goa.error

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## SecretsApiSetSecret

> SecretsApiSetSecret(ctx, environmentType, name).SetSecretRequest2(setSecretRequest2).Execute()

SetSecret secrets-api

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
	environmentType := "DEV" // string | The environment type for the secret
	name := "my-secret" // string | Name of the secret
	setSecretRequest2 := *openapiclient.NewSetSecretRequest2("super-secret-value") // SetSecretRequest2 | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	r, err := apiClient.SecretsApiAPI.SecretsApiSetSecret(context.Background(), environmentType, name).SetSecretRequest2(setSecretRequest2).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `SecretsApiAPI.SecretsApiSetSecret``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**environmentType** | **string** | The environment type for the secret | 
**name** | **string** | Name of the secret | 

### Other Parameters

Other parameters are passed through a pointer to a apiSecretsApiSetSecretRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


 **setSecretRequest2** | [**SetSecretRequest2**](SetSecretRequest2.md) |  | 

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

