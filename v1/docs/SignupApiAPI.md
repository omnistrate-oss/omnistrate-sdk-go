# \SignupApiAPI

All URIs are relative to *https://api.omnistrate.cloud*

Method | HTTP request | Description
------------- | ------------- | -------------
[**SignupApiChangePassword**](SignupApiAPI.md#SignupApiChangePassword) | **Post** /2022-09-01-00/change-password | ChangePassword signup-api
[**SignupApiRegenerateToken**](SignupApiAPI.md#SignupApiRegenerateToken) | **Post** /2022-09-01-00/users/{id}/regenerate-token | RegenerateToken signup-api
[**SignupApiResetPassword**](SignupApiAPI.md#SignupApiResetPassword) | **Post** /2022-09-01-00/reset-password | ResetPassword signup-api
[**SignupApiSignup**](SignupApiAPI.md#SignupApiSignup) | **Post** /2022-09-01-00/signup | Signup signup-api
[**SignupApiValidateToken**](SignupApiAPI.md#SignupApiValidateToken) | **Post** /2022-09-01-00/validate-token | ValidateToken signup-api



## SignupApiChangePassword

> SignupApiChangePassword(ctx).ChangePasswordRequestBody(changePasswordRequestBody).Execute()

ChangePassword signup-api

### Example

```go
package main

import (
	"context"
	"fmt"
	"os"
	openapiclient "github.com/omnistrate/omnistrate-sdk-go/v1"
)

func main() {
	changePasswordRequestBody := *openapiclient.NewChangePasswordRequestBody("abc@gmail.com", "password", "fc7b8dea-a50b-4c9a-8864-fc3d845a2be6") // ChangePasswordRequestBody | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	r, err := apiClient.SignupApiAPI.SignupApiChangePassword(context.Background()).ChangePasswordRequestBody(changePasswordRequestBody).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `SignupApiAPI.SignupApiChangePassword``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiSignupApiChangePasswordRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **changePasswordRequestBody** | [**ChangePasswordRequestBody**](ChangePasswordRequestBody.md) |  | 

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


## SignupApiRegenerateToken

> SignupApiRegenerateToken(ctx, id).Execute()

RegenerateToken signup-api

### Example

```go
package main

import (
	"context"
	"fmt"
	"os"
	openapiclient "github.com/omnistrate/omnistrate-sdk-go/v1"
)

func main() {
	id := "user-abcd1234" // string | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	r, err := apiClient.SignupApiAPI.SignupApiRegenerateToken(context.Background(), id).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `SignupApiAPI.SignupApiRegenerateToken``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiSignupApiRegenerateTokenRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

 (empty response body)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/vnd.goa.error

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## SignupApiResetPassword

> SignupApiResetPassword(ctx).ResetPasswordRequestBody(resetPasswordRequestBody).Execute()

ResetPassword signup-api

### Example

```go
package main

import (
	"context"
	"fmt"
	"os"
	openapiclient "github.com/omnistrate/omnistrate-sdk-go/v1"
)

func main() {
	resetPasswordRequestBody := *openapiclient.NewResetPasswordRequestBody("abc@gmail.com") // ResetPasswordRequestBody | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	r, err := apiClient.SignupApiAPI.SignupApiResetPassword(context.Background()).ResetPasswordRequestBody(resetPasswordRequestBody).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `SignupApiAPI.SignupApiResetPassword``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiSignupApiResetPasswordRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **resetPasswordRequestBody** | [**ResetPasswordRequestBody**](ResetPasswordRequestBody.md) |  | 

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


## SignupApiSignup

> SignupApiSignup(ctx).SignupRequestBody(signupRequestBody).Execute()

Signup signup-api

### Example

```go
package main

import (
	"context"
	"fmt"
	"os"
	openapiclient "github.com/omnistrate/omnistrate-sdk-go/v1"
)

func main() {
	signupRequestBody := *openapiclient.NewSignupRequestBody("abc@gmail.com", "John Doe", "password") // SignupRequestBody | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	r, err := apiClient.SignupApiAPI.SignupApiSignup(context.Background()).SignupRequestBody(signupRequestBody).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `SignupApiAPI.SignupApiSignup``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiSignupApiSignupRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **signupRequestBody** | [**SignupRequestBody**](SignupRequestBody.md) |  | 

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


## SignupApiValidateToken

> SignupApiValidateToken(ctx).ValidateTokenRequestBody(validateTokenRequestBody).Execute()

ValidateToken signup-api

### Example

```go
package main

import (
	"context"
	"fmt"
	"os"
	openapiclient "github.com/omnistrate/omnistrate-sdk-go/v1"
)

func main() {
	validateTokenRequestBody := *openapiclient.NewValidateTokenRequestBody("abc@gmail.com", "fc7b8dea-a50b-4c9a-8864-fc3d845a2be6") // ValidateTokenRequestBody | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	r, err := apiClient.SignupApiAPI.SignupApiValidateToken(context.Background()).ValidateTokenRequestBody(validateTokenRequestBody).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `SignupApiAPI.SignupApiValidateToken``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiSignupApiValidateTokenRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **validateTokenRequestBody** | [**ValidateTokenRequestBody**](ValidateTokenRequestBody.md) |  | 

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

