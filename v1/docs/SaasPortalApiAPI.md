# \SaasPortalApiAPI

All URIs are relative to *https://api.omnistrate.cloud*

Method | HTTP request | Description
------------- | ------------- | -------------
[**SaasPortalApiCreateSaaSPortalCustomDomain**](SaasPortalApiAPI.md#SaasPortalApiCreateSaaSPortalCustomDomain) | **Post** /2022-09-01-00/saas-portal-custom-domain | CreateSaaSPortalCustomDomain saas-portal-api
[**SaasPortalApiDeleteSaaSPortalCustomDomain**](SaasPortalApiAPI.md#SaasPortalApiDeleteSaaSPortalCustomDomain) | **Delete** /2022-09-01-00/saas-portal-custom-domain/{environmentType} | DeleteSaaSPortalCustomDomain saas-portal-api
[**SaasPortalApiListSaaSPortalCustomDomains**](SaasPortalApiAPI.md#SaasPortalApiListSaaSPortalCustomDomains) | **Get** /2022-09-01-00/saas-portal-custom-domain | ListSaaSPortalCustomDomains saas-portal-api
[**SaasPortalApiListSaaSPortals**](SaasPortalApiAPI.md#SaasPortalApiListSaaSPortals) | **Get** /2022-09-01-00/saas-portal | ListSaaSPortals saas-portal-api
[**SaasPortalApiUpdateSaaSPortal**](SaasPortalApiAPI.md#SaasPortalApiUpdateSaaSPortal) | **Patch** /2022-09-01-00/saas-portal/{environmentType} | UpdateSaaSPortal saas-portal-api
[**SaasPortalApiUpdateSaaSPortalCustomDomain**](SaasPortalApiAPI.md#SaasPortalApiUpdateSaaSPortalCustomDomain) | **Patch** /2022-09-01-00/saas-portal-custom-domain/{environmentType} | UpdateSaaSPortalCustomDomain saas-portal-api



## SaasPortalApiCreateSaaSPortalCustomDomain

> SaasPortalApiCreateSaaSPortalCustomDomain(ctx).CreateSaaSPortalCustomDomainRequestBody(createSaaSPortalCustomDomainRequestBody).Execute()

CreateSaaSPortalCustomDomain saas-portal-api

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
	createSaaSPortalCustomDomainRequestBody := *openapiclient.NewCreateSaaSPortalCustomDomainRequestBody("mycustomdomain.com", "My custom domain description", "DEV", "MyCustomDomain") // CreateSaaSPortalCustomDomainRequestBody | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	r, err := apiClient.SaasPortalApiAPI.SaasPortalApiCreateSaaSPortalCustomDomain(context.Background()).CreateSaaSPortalCustomDomainRequestBody(createSaaSPortalCustomDomainRequestBody).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `SaasPortalApiAPI.SaasPortalApiCreateSaaSPortalCustomDomain``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiSaasPortalApiCreateSaaSPortalCustomDomainRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **createSaaSPortalCustomDomainRequestBody** | [**CreateSaaSPortalCustomDomainRequestBody**](CreateSaaSPortalCustomDomainRequestBody.md) |  | 

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


## SaasPortalApiDeleteSaaSPortalCustomDomain

> SaasPortalApiDeleteSaaSPortalCustomDomain(ctx, environmentType).Execute()

DeleteSaaSPortalCustomDomain saas-portal-api

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
	environmentType := "DEV" // string | The environment type for the custom domain

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	r, err := apiClient.SaasPortalApiAPI.SaasPortalApiDeleteSaaSPortalCustomDomain(context.Background(), environmentType).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `SaasPortalApiAPI.SaasPortalApiDeleteSaaSPortalCustomDomain``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**environmentType** | **string** | The environment type for the custom domain | 

### Other Parameters

Other parameters are passed through a pointer to a apiSaasPortalApiDeleteSaaSPortalCustomDomainRequest struct via the builder pattern


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


## SaasPortalApiListSaaSPortalCustomDomains

> ListSaaSPortalCustomDomainsResult SaasPortalApiListSaaSPortalCustomDomains(ctx).Execute()

ListSaaSPortalCustomDomains saas-portal-api

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
	resp, r, err := apiClient.SaasPortalApiAPI.SaasPortalApiListSaaSPortalCustomDomains(context.Background()).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `SaasPortalApiAPI.SaasPortalApiListSaaSPortalCustomDomains``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `SaasPortalApiListSaaSPortalCustomDomains`: ListSaaSPortalCustomDomainsResult
	fmt.Fprintf(os.Stdout, "Response from `SaasPortalApiAPI.SaasPortalApiListSaaSPortalCustomDomains`: %v\n", resp)
}
```

### Path Parameters

This endpoint does not need any parameter.

### Other Parameters

Other parameters are passed through a pointer to a apiSaasPortalApiListSaaSPortalCustomDomainsRequest struct via the builder pattern


### Return type

[**ListSaaSPortalCustomDomainsResult**](ListSaaSPortalCustomDomainsResult.md)

### Authorization

[api_key_header_Authorization](../README.md#api_key_header_Authorization)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json, application/vnd.goa.error

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## SaasPortalApiListSaaSPortals

> ListSaaSPortalsResult SaasPortalApiListSaaSPortals(ctx).Execute()

ListSaaSPortals saas-portal-api

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
	resp, r, err := apiClient.SaasPortalApiAPI.SaasPortalApiListSaaSPortals(context.Background()).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `SaasPortalApiAPI.SaasPortalApiListSaaSPortals``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `SaasPortalApiListSaaSPortals`: ListSaaSPortalsResult
	fmt.Fprintf(os.Stdout, "Response from `SaasPortalApiAPI.SaasPortalApiListSaaSPortals`: %v\n", resp)
}
```

### Path Parameters

This endpoint does not need any parameter.

### Other Parameters

Other parameters are passed through a pointer to a apiSaasPortalApiListSaaSPortalsRequest struct via the builder pattern


### Return type

[**ListSaaSPortalsResult**](ListSaaSPortalsResult.md)

### Authorization

[api_key_header_Authorization](../README.md#api_key_header_Authorization)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json, application/vnd.goa.error

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## SaasPortalApiUpdateSaaSPortal

> SaasPortalApiUpdateSaaSPortal(ctx, environmentType).UpdateSaaSPortalRequestBody(updateSaaSPortalRequestBody).Execute()

UpdateSaaSPortal saas-portal-api

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
	environmentType := "DEV" // string | The environment type for the saas portal custom domain to update
	updateSaaSPortalRequestBody := *openapiclient.NewUpdateSaaSPortalRequestBody() // UpdateSaaSPortalRequestBody | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	r, err := apiClient.SaasPortalApiAPI.SaasPortalApiUpdateSaaSPortal(context.Background(), environmentType).UpdateSaaSPortalRequestBody(updateSaaSPortalRequestBody).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `SaasPortalApiAPI.SaasPortalApiUpdateSaaSPortal``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**environmentType** | **string** | The environment type for the saas portal custom domain to update | 

### Other Parameters

Other parameters are passed through a pointer to a apiSaasPortalApiUpdateSaaSPortalRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **updateSaaSPortalRequestBody** | [**UpdateSaaSPortalRequestBody**](UpdateSaaSPortalRequestBody.md) |  | 

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


## SaasPortalApiUpdateSaaSPortalCustomDomain

> SaasPortalApiUpdateSaaSPortalCustomDomain(ctx, environmentType).UpdateSaaSPortalCustomDomainRequestBody(updateSaaSPortalCustomDomainRequestBody).Execute()

UpdateSaaSPortalCustomDomain saas-portal-api

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
	environmentType := "DEV" // string | The environment type for the saas portal custom domain to update
	updateSaaSPortalCustomDomainRequestBody := *openapiclient.NewUpdateSaaSPortalCustomDomainRequestBody() // UpdateSaaSPortalCustomDomainRequestBody | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	r, err := apiClient.SaasPortalApiAPI.SaasPortalApiUpdateSaaSPortalCustomDomain(context.Background(), environmentType).UpdateSaaSPortalCustomDomainRequestBody(updateSaaSPortalCustomDomainRequestBody).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `SaasPortalApiAPI.SaasPortalApiUpdateSaaSPortalCustomDomain``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**environmentType** | **string** | The environment type for the saas portal custom domain to update | 

### Other Parameters

Other parameters are passed through a pointer to a apiSaasPortalApiUpdateSaaSPortalCustomDomainRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **updateSaaSPortalCustomDomainRequestBody** | [**UpdateSaaSPortalCustomDomainRequestBody**](UpdateSaaSPortalCustomDomainRequestBody.md) |  | 

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

