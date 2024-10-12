# \UsersApiAPI

All URIs are relative to *https://api.omnistrate.cloud*

Method | HTTP request | Description
------------- | ------------- | -------------
[**UsersApiCustomerInviteUser**](UsersApiAPI.md#UsersApiCustomerInviteUser) | **Post** /2022-09-01-00/customer-invite-user | CustomerInviteUser users-api
[**UsersApiCustomerLoginWithIdentityProvider**](UsersApiAPI.md#UsersApiCustomerLoginWithIdentityProvider) | **Post** /2022-09-01-00/customer-login-with-identity-provider | CustomerLoginWithIdentityProvider users-api
[**UsersApiCustomerResetPassword**](UsersApiAPI.md#UsersApiCustomerResetPassword) | **Post** /2022-09-01-00/customer-reset-password | CustomerResetPassword users-api
[**UsersApiCustomerSignin**](UsersApiAPI.md#UsersApiCustomerSignin) | **Post** /2022-09-01-00/customer-user-signin | CustomerSignin users-api
[**UsersApiCustomerSignup**](UsersApiAPI.md#UsersApiCustomerSignup) | **Post** /2022-09-01-00/customer-user-signup | CustomerSignup users-api
[**UsersApiDeleteUser**](UsersApiAPI.md#UsersApiDeleteUser) | **Delete** /2022-09-01-00/user | DeleteUser users-api
[**UsersApiDescribeUser**](UsersApiAPI.md#UsersApiDescribeUser) | **Get** /2022-09-01-00/user | DescribeUser users-api
[**UsersApiDescribeUsersByOrg**](UsersApiAPI.md#UsersApiDescribeUsersByOrg) | **Get** /2022-09-01-00/org-users | DescribeUsersByOrg users-api
[**UsersApiInviteUser**](UsersApiAPI.md#UsersApiInviteUser) | **Post** /2022-09-01-00/invite-user | InviteUser users-api
[**UsersApiLogout**](UsersApiAPI.md#UsersApiLogout) | **Post** /2022-09-01-00/logout | Logout users-api
[**UsersApiRevokeUserRole**](UsersApiAPI.md#UsersApiRevokeUserRole) | **Delete** /2022-09-01-00/revoke-user-role | RevokeUserRole users-api
[**UsersApiUpdatePassword**](UsersApiAPI.md#UsersApiUpdatePassword) | **Post** /2022-09-01-00/update-password | UpdatePassword users-api
[**UsersApiUpdateUser**](UsersApiAPI.md#UsersApiUpdateUser) | **Patch** /2022-09-01-00/user/{id} | UpdateUser users-api



## UsersApiCustomerInviteUser

> UsersApiCustomerInviteUser(ctx).ResetPasswordRequestBody(resetPasswordRequestBody).Execute()

CustomerInviteUser users-api

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
	resetPasswordRequestBody := *openapiclient.NewResetPasswordRequestBody("abc@gmail.com") // ResetPasswordRequestBody | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	r, err := apiClient.UsersApiAPI.UsersApiCustomerInviteUser(context.Background()).ResetPasswordRequestBody(resetPasswordRequestBody).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `UsersApiAPI.UsersApiCustomerInviteUser``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiUsersApiCustomerInviteUserRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **resetPasswordRequestBody** | [**ResetPasswordRequestBody**](ResetPasswordRequestBody.md) |  | 

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


## UsersApiCustomerLoginWithIdentityProvider

> SigninResult UsersApiCustomerLoginWithIdentityProvider(ctx).CustomerLoginWithIdentityProviderRequestBody(customerLoginWithIdentityProviderRequestBody).Execute()

CustomerLoginWithIdentityProvider users-api

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
	customerLoginWithIdentityProviderRequestBody := *openapiclient.NewCustomerLoginWithIdentityProviderRequestBody("4/P7q7W91a-oMsCeLvIaQm6bTrgtp7&", "Google") // CustomerLoginWithIdentityProviderRequestBody | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.UsersApiAPI.UsersApiCustomerLoginWithIdentityProvider(context.Background()).CustomerLoginWithIdentityProviderRequestBody(customerLoginWithIdentityProviderRequestBody).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `UsersApiAPI.UsersApiCustomerLoginWithIdentityProvider``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `UsersApiCustomerLoginWithIdentityProvider`: SigninResult
	fmt.Fprintf(os.Stdout, "Response from `UsersApiAPI.UsersApiCustomerLoginWithIdentityProvider`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiUsersApiCustomerLoginWithIdentityProviderRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **customerLoginWithIdentityProviderRequestBody** | [**CustomerLoginWithIdentityProviderRequestBody**](CustomerLoginWithIdentityProviderRequestBody.md) |  | 

### Return type

[**SigninResult**](SigninResult.md)

### Authorization

[api_key_header_Authorization](../README.md#api_key_header_Authorization)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json, application/vnd.goa.error

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## UsersApiCustomerResetPassword

> UsersApiCustomerResetPassword(ctx).ResetPasswordRequestBody(resetPasswordRequestBody).Execute()

CustomerResetPassword users-api

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
	resetPasswordRequestBody := *openapiclient.NewResetPasswordRequestBody("abc@gmail.com") // ResetPasswordRequestBody | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	r, err := apiClient.UsersApiAPI.UsersApiCustomerResetPassword(context.Background()).ResetPasswordRequestBody(resetPasswordRequestBody).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `UsersApiAPI.UsersApiCustomerResetPassword``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiUsersApiCustomerResetPasswordRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **resetPasswordRequestBody** | [**ResetPasswordRequestBody**](ResetPasswordRequestBody.md) |  | 

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


## UsersApiCustomerSignin

> SigninResult UsersApiCustomerSignin(ctx).CustomerSigninRequestBody(customerSigninRequestBody).Execute()

CustomerSignin users-api

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
	customerSigninRequestBody := *openapiclient.NewCustomerSigninRequestBody("abc@gmail.com") // CustomerSigninRequestBody | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.UsersApiAPI.UsersApiCustomerSignin(context.Background()).CustomerSigninRequestBody(customerSigninRequestBody).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `UsersApiAPI.UsersApiCustomerSignin``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `UsersApiCustomerSignin`: SigninResult
	fmt.Fprintf(os.Stdout, "Response from `UsersApiAPI.UsersApiCustomerSignin`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiUsersApiCustomerSigninRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **customerSigninRequestBody** | [**CustomerSigninRequestBody**](CustomerSigninRequestBody.md) |  | 

### Return type

[**SigninResult**](SigninResult.md)

### Authorization

[api_key_header_Authorization](../README.md#api_key_header_Authorization)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json, application/vnd.goa.error

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## UsersApiCustomerSignup

> UsersApiCustomerSignup(ctx).SignupRequestBody(signupRequestBody).Execute()

CustomerSignup users-api

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
	signupRequestBody := *openapiclient.NewSignupRequestBody("abc@gmail.com", "John Doe", "password") // SignupRequestBody | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	r, err := apiClient.UsersApiAPI.UsersApiCustomerSignup(context.Background()).SignupRequestBody(signupRequestBody).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `UsersApiAPI.UsersApiCustomerSignup``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiUsersApiCustomerSignupRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **signupRequestBody** | [**SignupRequestBody**](SignupRequestBody.md) |  | 

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


## UsersApiDeleteUser

> UsersApiDeleteUser(ctx).Execute()

DeleteUser users-api

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
	r, err := apiClient.UsersApiAPI.UsersApiDeleteUser(context.Background()).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `UsersApiAPI.UsersApiDeleteUser``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
}
```

### Path Parameters

This endpoint does not need any parameter.

### Other Parameters

Other parameters are passed through a pointer to a apiUsersApiDeleteUserRequest struct via the builder pattern


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


## UsersApiDescribeUser

> DescribeUserResult UsersApiDescribeUser(ctx).Execute()

DescribeUser users-api

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
	resp, r, err := apiClient.UsersApiAPI.UsersApiDescribeUser(context.Background()).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `UsersApiAPI.UsersApiDescribeUser``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `UsersApiDescribeUser`: DescribeUserResult
	fmt.Fprintf(os.Stdout, "Response from `UsersApiAPI.UsersApiDescribeUser`: %v\n", resp)
}
```

### Path Parameters

This endpoint does not need any parameter.

### Other Parameters

Other parameters are passed through a pointer to a apiUsersApiDescribeUserRequest struct via the builder pattern


### Return type

[**DescribeUserResult**](DescribeUserResult.md)

### Authorization

[api_key_header_Authorization](../README.md#api_key_header_Authorization)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json, application/vnd.goa.error

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## UsersApiDescribeUsersByOrg

> DescribeUsersByOrgResult UsersApiDescribeUsersByOrg(ctx).Execute()

DescribeUsersByOrg users-api

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
	resp, r, err := apiClient.UsersApiAPI.UsersApiDescribeUsersByOrg(context.Background()).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `UsersApiAPI.UsersApiDescribeUsersByOrg``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `UsersApiDescribeUsersByOrg`: DescribeUsersByOrgResult
	fmt.Fprintf(os.Stdout, "Response from `UsersApiAPI.UsersApiDescribeUsersByOrg`: %v\n", resp)
}
```

### Path Parameters

This endpoint does not need any parameter.

### Other Parameters

Other parameters are passed through a pointer to a apiUsersApiDescribeUsersByOrgRequest struct via the builder pattern


### Return type

[**DescribeUsersByOrgResult**](DescribeUsersByOrgResult.md)

### Authorization

[api_key_header_Authorization](../README.md#api_key_header_Authorization)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json, application/vnd.goa.error

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## UsersApiInviteUser

> UsersApiInviteUser(ctx).InviteUserRequestBody(inviteUserRequestBody).Execute()

InviteUser users-api

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
	inviteUserRequestBody := *openapiclient.NewInviteUserRequestBody("abc@gmail.com", "reader") // InviteUserRequestBody | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	r, err := apiClient.UsersApiAPI.UsersApiInviteUser(context.Background()).InviteUserRequestBody(inviteUserRequestBody).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `UsersApiAPI.UsersApiInviteUser``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiUsersApiInviteUserRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **inviteUserRequestBody** | [**InviteUserRequestBody**](InviteUserRequestBody.md) |  | 

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


## UsersApiLogout

> UsersApiLogout(ctx).Execute()

Logout users-api

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
	r, err := apiClient.UsersApiAPI.UsersApiLogout(context.Background()).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `UsersApiAPI.UsersApiLogout``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
}
```

### Path Parameters

This endpoint does not need any parameter.

### Other Parameters

Other parameters are passed through a pointer to a apiUsersApiLogoutRequest struct via the builder pattern


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


## UsersApiRevokeUserRole

> UsersApiRevokeUserRole(ctx).InviteUserRequestBody(inviteUserRequestBody).Execute()

RevokeUserRole users-api

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
	inviteUserRequestBody := *openapiclient.NewInviteUserRequestBody("abc@gmail.com", "reader") // InviteUserRequestBody | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	r, err := apiClient.UsersApiAPI.UsersApiRevokeUserRole(context.Background()).InviteUserRequestBody(inviteUserRequestBody).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `UsersApiAPI.UsersApiRevokeUserRole``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiUsersApiRevokeUserRoleRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **inviteUserRequestBody** | [**InviteUserRequestBody**](InviteUserRequestBody.md) |  | 

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


## UsersApiUpdatePassword

> UsersApiUpdatePassword(ctx).UpdatePasswordRequestBody(updatePasswordRequestBody).Execute()

UpdatePassword users-api

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
	updatePasswordRequestBody := *openapiclient.NewUpdatePasswordRequestBody("password") // UpdatePasswordRequestBody | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	r, err := apiClient.UsersApiAPI.UsersApiUpdatePassword(context.Background()).UpdatePasswordRequestBody(updatePasswordRequestBody).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `UsersApiAPI.UsersApiUpdatePassword``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiUsersApiUpdatePasswordRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **updatePasswordRequestBody** | [**UpdatePasswordRequestBody**](UpdatePasswordRequestBody.md) |  | 

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


## UsersApiUpdateUser

> UsersApiUpdateUser(ctx, id).UpdateUserRequestBody(updateUserRequestBody).Execute()

UpdateUser users-api

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
	id := "user-abcd1234" // string | The User ID
	updateUserRequestBody := *openapiclient.NewUpdateUserRequestBody() // UpdateUserRequestBody | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	r, err := apiClient.UsersApiAPI.UsersApiUpdateUser(context.Background(), id).UpdateUserRequestBody(updateUserRequestBody).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `UsersApiAPI.UsersApiUpdateUser``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **string** | The User ID | 

### Other Parameters

Other parameters are passed through a pointer to a apiUsersApiUpdateUserRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **updateUserRequestBody** | [**UpdateUserRequestBody**](UpdateUserRequestBody.md) |  | 

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

