# \CustomDomainApiAPI

All URIs are relative to *https://api.omnistrate.cloud*

Method | HTTP request | Description
------------- | ------------- | -------------
[**CustomDomainApiCreateCustomDomain**](CustomDomainApiAPI.md#CustomDomainApiCreateCustomDomain) | **Post** /2022-09-01-00/customdomain | CreateCustomDomain custom-domain-api
[**CustomDomainApiCustomDomainIdentityID**](CustomDomainApiAPI.md#CustomDomainApiCustomDomainIdentityID) | **Get** /2022-09-01-00/customdomain/identityid | CustomDomainIdentityID custom-domain-api
[**CustomDomainApiDeleteCustomDomain**](CustomDomainApiAPI.md#CustomDomainApiDeleteCustomDomain) | **Delete** /2022-09-01-00/customdomain/{id} | DeleteCustomDomain custom-domain-api
[**CustomDomainApiDescribeCustomDomain**](CustomDomainApiAPI.md#CustomDomainApiDescribeCustomDomain) | **Get** /2022-09-01-00/customdomain/{id} | DescribeCustomDomain custom-domain-api
[**CustomDomainApiListCustomDomain**](CustomDomainApiAPI.md#CustomDomainApiListCustomDomain) | **Get** /2022-09-01-00/customdomain | ListCustomDomain custom-domain-api
[**CustomDomainApiVerifyCustomDomain**](CustomDomainApiAPI.md#CustomDomainApiVerifyCustomDomain) | **Post** /2022-09-01-00/customdomain/verify/{id} | VerifyCustomDomain custom-domain-api



## CustomDomainApiCreateCustomDomain

> string CustomDomainApiCreateCustomDomain(ctx).CreateCustomDomainRequestBody(createCustomDomainRequestBody).Execute()

CreateCustomDomain custom-domain-api

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
	createCustomDomainRequestBody := *openapiclient.NewCreateCustomDomainRequestBody("mydomain.dev", "Description of the domain", "Dev domain", *openapiclient.NewRoute53Configuration("123456789012")) // CreateCustomDomainRequestBody | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.CustomDomainApiAPI.CustomDomainApiCreateCustomDomain(context.Background()).CreateCustomDomainRequestBody(createCustomDomainRequestBody).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `CustomDomainApiAPI.CustomDomainApiCreateCustomDomain``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `CustomDomainApiCreateCustomDomain`: string
	fmt.Fprintf(os.Stdout, "Response from `CustomDomainApiAPI.CustomDomainApiCreateCustomDomain`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiCustomDomainApiCreateCustomDomainRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **createCustomDomainRequestBody** | [**CreateCustomDomainRequestBody**](CreateCustomDomainRequestBody.md) |  | 

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


## CustomDomainApiCustomDomainIdentityID

> AccountConfigIdentityIDResult CustomDomainApiCustomDomainIdentityID(ctx).Execute()

CustomDomainIdentityID custom-domain-api



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
	resp, r, err := apiClient.CustomDomainApiAPI.CustomDomainApiCustomDomainIdentityID(context.Background()).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `CustomDomainApiAPI.CustomDomainApiCustomDomainIdentityID``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `CustomDomainApiCustomDomainIdentityID`: AccountConfigIdentityIDResult
	fmt.Fprintf(os.Stdout, "Response from `CustomDomainApiAPI.CustomDomainApiCustomDomainIdentityID`: %v\n", resp)
}
```

### Path Parameters

This endpoint does not need any parameter.

### Other Parameters

Other parameters are passed through a pointer to a apiCustomDomainApiCustomDomainIdentityIDRequest struct via the builder pattern


### Return type

[**AccountConfigIdentityIDResult**](AccountConfigIdentityIDResult.md)

### Authorization

[api_key_header_Authorization](../README.md#api_key_header_Authorization)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json, application/vnd.goa.error

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## CustomDomainApiDeleteCustomDomain

> CustomDomainApiDeleteCustomDomain(ctx, id).Execute()

DeleteCustomDomain custom-domain-api

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
	id := "cd-12345678" // string | custom domain ID to operate on

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	r, err := apiClient.CustomDomainApiAPI.CustomDomainApiDeleteCustomDomain(context.Background(), id).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `CustomDomainApiAPI.CustomDomainApiDeleteCustomDomain``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **string** | custom domain ID to operate on | 

### Other Parameters

Other parameters are passed through a pointer to a apiCustomDomainApiDeleteCustomDomainRequest struct via the builder pattern


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


## CustomDomainApiDescribeCustomDomain

> DescribeCustomDomainResult CustomDomainApiDescribeCustomDomain(ctx, id).Execute()

DescribeCustomDomain custom-domain-api

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
	id := "cd-12345678" // string | custom domain ID to operate on

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.CustomDomainApiAPI.CustomDomainApiDescribeCustomDomain(context.Background(), id).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `CustomDomainApiAPI.CustomDomainApiDescribeCustomDomain``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `CustomDomainApiDescribeCustomDomain`: DescribeCustomDomainResult
	fmt.Fprintf(os.Stdout, "Response from `CustomDomainApiAPI.CustomDomainApiDescribeCustomDomain`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **string** | custom domain ID to operate on | 

### Other Parameters

Other parameters are passed through a pointer to a apiCustomDomainApiDescribeCustomDomainRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

[**DescribeCustomDomainResult**](DescribeCustomDomainResult.md)

### Authorization

[api_key_header_Authorization](../README.md#api_key_header_Authorization)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json, application/vnd.goa.error

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## CustomDomainApiListCustomDomain

> ListCustomDomainResult CustomDomainApiListCustomDomain(ctx).Execute()

ListCustomDomain custom-domain-api

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
	resp, r, err := apiClient.CustomDomainApiAPI.CustomDomainApiListCustomDomain(context.Background()).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `CustomDomainApiAPI.CustomDomainApiListCustomDomain``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `CustomDomainApiListCustomDomain`: ListCustomDomainResult
	fmt.Fprintf(os.Stdout, "Response from `CustomDomainApiAPI.CustomDomainApiListCustomDomain`: %v\n", resp)
}
```

### Path Parameters

This endpoint does not need any parameter.

### Other Parameters

Other parameters are passed through a pointer to a apiCustomDomainApiListCustomDomainRequest struct via the builder pattern


### Return type

[**ListCustomDomainResult**](ListCustomDomainResult.md)

### Authorization

[api_key_header_Authorization](../README.md#api_key_header_Authorization)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json, application/vnd.goa.error

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## CustomDomainApiVerifyCustomDomain

> CustomDomainApiVerifyCustomDomain(ctx, id).Execute()

VerifyCustomDomain custom-domain-api

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
	id := "cd-12345678" // string | custom domain ID to operate on

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	r, err := apiClient.CustomDomainApiAPI.CustomDomainApiVerifyCustomDomain(context.Background(), id).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `CustomDomainApiAPI.CustomDomainApiVerifyCustomDomain``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **string** | custom domain ID to operate on | 

### Other Parameters

Other parameters are passed through a pointer to a apiCustomDomainApiVerifyCustomDomainRequest struct via the builder pattern


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

