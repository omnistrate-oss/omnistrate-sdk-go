# \SigninApiAPI

All URIs are relative to *https://api.omnistrate.cloud*

Method | HTTP request | Description
------------- | ------------- | -------------
[**SigninApiLoginWithIdentityProvider**](SigninApiAPI.md#SigninApiLoginWithIdentityProvider) | **Post** /2022-09-01-00/login-with-identity-provider | LoginWithIdentityProvider signin-api
[**SigninApiSignin**](SigninApiAPI.md#SigninApiSignin) | **Post** /2022-09-01-00/signin | Signin signin-api



## SigninApiLoginWithIdentityProvider

> SigninResult SigninApiLoginWithIdentityProvider(ctx).LoginWithIdentityProviderRequestBody(loginWithIdentityProviderRequestBody).Execute()

LoginWithIdentityProvider signin-api

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
	loginWithIdentityProviderRequestBody := *openapiclient.NewLoginWithIdentityProviderRequestBody("Google") // LoginWithIdentityProviderRequestBody | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.SigninApiAPI.SigninApiLoginWithIdentityProvider(context.Background()).LoginWithIdentityProviderRequestBody(loginWithIdentityProviderRequestBody).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `SigninApiAPI.SigninApiLoginWithIdentityProvider``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `SigninApiLoginWithIdentityProvider`: SigninResult
	fmt.Fprintf(os.Stdout, "Response from `SigninApiAPI.SigninApiLoginWithIdentityProvider`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiSigninApiLoginWithIdentityProviderRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **loginWithIdentityProviderRequestBody** | [**LoginWithIdentityProviderRequestBody**](LoginWithIdentityProviderRequestBody.md) |  | 

### Return type

[**SigninResult**](SigninResult.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json, application/vnd.goa.error

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## SigninApiSignin

> SigninResult SigninApiSignin(ctx).SigninRequestBody(signinRequestBody).Execute()

Signin signin-api

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
	signinRequestBody := *openapiclient.NewSigninRequestBody("abc@gmail.com") // SigninRequestBody | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.SigninApiAPI.SigninApiSignin(context.Background()).SigninRequestBody(signinRequestBody).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `SigninApiAPI.SigninApiSignin``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `SigninApiSignin`: SigninResult
	fmt.Fprintf(os.Stdout, "Response from `SigninApiAPI.SigninApiSignin`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiSigninApiSigninRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **signinRequestBody** | [**SigninRequestBody**](SigninRequestBody.md) |  | 

### Return type

[**SigninResult**](SigninResult.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json, application/vnd.goa.error

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

