# \SigninApiAPI

All URIs are relative to *https://api.omnistrate.cloud*

Method | HTTP request | Description
------------- | ------------- | -------------
[**SigninApiLoginWithIdentityProvider**](SigninApiAPI.md#SigninApiLoginWithIdentityProvider) | **Post** /2022-09-01-00/login-with-identity-provider | LoginWithIdentityProvider signin-api
[**SigninApiRefreshToken**](SigninApiAPI.md#SigninApiRefreshToken) | **Post** /2022-09-01-00/refresh-token | RefreshToken signin-api
[**SigninApiRevokeToken**](SigninApiAPI.md#SigninApiRevokeToken) | **Post** /2022-09-01-00/revoke-token | RevokeToken signin-api
[**SigninApiSignin**](SigninApiAPI.md#SigninApiSignin) | **Post** /2022-09-01-00/signin | Signin signin-api



## SigninApiLoginWithIdentityProvider

> LoginWithIdentityProviderResult SigninApiLoginWithIdentityProvider(ctx).LoginWithIdentityProviderRequest(loginWithIdentityProviderRequest).Execute()

LoginWithIdentityProvider signin-api

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
	loginWithIdentityProviderRequest := *openapiclient.NewLoginWithIdentityProviderRequest("Google|GitHub|Microsoft Entra|Amazon Cognito|Okta|Auth0|Keycloak|OIDC") // LoginWithIdentityProviderRequest | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.SigninApiAPI.SigninApiLoginWithIdentityProvider(context.Background()).LoginWithIdentityProviderRequest(loginWithIdentityProviderRequest).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `SigninApiAPI.SigninApiLoginWithIdentityProvider``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `SigninApiLoginWithIdentityProvider`: LoginWithIdentityProviderResult
	fmt.Fprintf(os.Stdout, "Response from `SigninApiAPI.SigninApiLoginWithIdentityProvider`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiSigninApiLoginWithIdentityProviderRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **loginWithIdentityProviderRequest** | [**LoginWithIdentityProviderRequest**](LoginWithIdentityProviderRequest.md) |  | 

### Return type

[**LoginWithIdentityProviderResult**](LoginWithIdentityProviderResult.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json, application/vnd.goa.error

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## SigninApiRefreshToken

> RefreshTokenResult SigninApiRefreshToken(ctx).RefreshTokenRequest(refreshTokenRequest).Execute()

RefreshToken signin-api



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
	refreshTokenRequest := *openapiclient.NewRefreshTokenRequest() // RefreshTokenRequest | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.SigninApiAPI.SigninApiRefreshToken(context.Background()).RefreshTokenRequest(refreshTokenRequest).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `SigninApiAPI.SigninApiRefreshToken``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `SigninApiRefreshToken`: RefreshTokenResult
	fmt.Fprintf(os.Stdout, "Response from `SigninApiAPI.SigninApiRefreshToken`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiSigninApiRefreshTokenRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **refreshTokenRequest** | [**RefreshTokenRequest**](RefreshTokenRequest.md) |  | 

### Return type

[**RefreshTokenResult**](RefreshTokenResult.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json, application/vnd.goa.error

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## SigninApiRevokeToken

> SigninApiRevokeToken(ctx).RevokeTokenRequest(revokeTokenRequest).Execute()

RevokeToken signin-api



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
	revokeTokenRequest := *openapiclient.NewRevokeTokenRequest() // RevokeTokenRequest | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	r, err := apiClient.SigninApiAPI.SigninApiRevokeToken(context.Background()).RevokeTokenRequest(revokeTokenRequest).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `SigninApiAPI.SigninApiRevokeToken``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiSigninApiRevokeTokenRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **revokeTokenRequest** | [**RevokeTokenRequest**](RevokeTokenRequest.md) |  | 

### Return type

 (empty response body)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/vnd.goa.error

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## SigninApiSignin

> SigninResult SigninApiSignin(ctx).SigninRequest(signinRequest).Execute()

Signin signin-api

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
	signinRequest := *openapiclient.NewSigninRequest("abc@gmail.com") // SigninRequest | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.SigninApiAPI.SigninApiSignin(context.Background()).SigninRequest(signinRequest).Execute()
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
 **signinRequest** | [**SigninRequest**](SigninRequest.md) |  | 

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

