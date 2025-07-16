# \StripeConnectApiAPI

All URIs are relative to *https://api.omnistrate.cloud*

Method | HTTP request | Description
------------- | ------------- | -------------
[**StripeConnectApiCompleteOAuthConnection**](StripeConnectApiAPI.md#StripeConnectApiCompleteOAuthConnection) | **Post** /2022-09-01-00/stripe-setup | CompleteOAuthConnection stripe-connect-api
[**StripeConnectApiGetStripeAuthorizeUrl**](StripeConnectApiAPI.md#StripeConnectApiGetStripeAuthorizeUrl) | **Get** /2022-09-01-00/stripe-setup | GetStripeAuthorizeUrl stripe-connect-api



## StripeConnectApiCompleteOAuthConnection

> CompleteOAuthConnectionResult StripeConnectApiCompleteOAuthConnection(ctx).CompleteOAuthConnectionRequest2(completeOAuthConnectionRequest2).Execute()

CompleteOAuthConnection stripe-connect-api

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
	completeOAuthConnectionRequest2 := *openapiclient.NewCompleteOAuthConnectionRequest2() // CompleteOAuthConnectionRequest2 | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.StripeConnectApiAPI.StripeConnectApiCompleteOAuthConnection(context.Background()).CompleteOAuthConnectionRequest2(completeOAuthConnectionRequest2).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `StripeConnectApiAPI.StripeConnectApiCompleteOAuthConnection``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `StripeConnectApiCompleteOAuthConnection`: CompleteOAuthConnectionResult
	fmt.Fprintf(os.Stdout, "Response from `StripeConnectApiAPI.StripeConnectApiCompleteOAuthConnection`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiStripeConnectApiCompleteOAuthConnectionRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **completeOAuthConnectionRequest2** | [**CompleteOAuthConnectionRequest2**](CompleteOAuthConnectionRequest2.md) |  | 

### Return type

[**CompleteOAuthConnectionResult**](CompleteOAuthConnectionResult.md)

### Authorization

[api_key_header_Authorization](../README.md#api_key_header_Authorization)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json, application/vnd.goa.error

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## StripeConnectApiGetStripeAuthorizeUrl

> StripeAuthorizeURLResult StripeConnectApiGetStripeAuthorizeUrl(ctx).Execute()

GetStripeAuthorizeUrl stripe-connect-api

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
	resp, r, err := apiClient.StripeConnectApiAPI.StripeConnectApiGetStripeAuthorizeUrl(context.Background()).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `StripeConnectApiAPI.StripeConnectApiGetStripeAuthorizeUrl``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `StripeConnectApiGetStripeAuthorizeUrl`: StripeAuthorizeURLResult
	fmt.Fprintf(os.Stdout, "Response from `StripeConnectApiAPI.StripeConnectApiGetStripeAuthorizeUrl`: %v\n", resp)
}
```

### Path Parameters

This endpoint does not need any parameter.

### Other Parameters

Other parameters are passed through a pointer to a apiStripeConnectApiGetStripeAuthorizeUrlRequest struct via the builder pattern


### Return type

[**StripeAuthorizeURLResult**](StripeAuthorizeURLResult.md)

### Authorization

[api_key_header_Authorization](../README.md#api_key_header_Authorization)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json, application/vnd.goa.error

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

