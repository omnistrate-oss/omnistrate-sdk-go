# \UsersApiAPI

All URIs are relative to *https://api.omnistrate.cloud*

Method | HTTP request | Description
------------- | ------------- | -------------
[**UsersApiCustomerDeleteUser**](UsersApiAPI.md#UsersApiCustomerDeleteUser) | **Delete** /2022-09-01-00/customer-delete-user | CustomerDeleteUser users-api
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



## UsersApiCustomerDeleteUser

> UsersApiCustomerDeleteUser(ctx).Execute()

CustomerDeleteUser users-api

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
	r, err := apiClient.UsersApiAPI.UsersApiCustomerDeleteUser(context.Background()).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `UsersApiAPI.UsersApiCustomerDeleteUser``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
}
```

### Path Parameters

This endpoint does not need any parameter.

### Other Parameters

Other parameters are passed through a pointer to a apiUsersApiCustomerDeleteUserRequest struct via the builder pattern


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


## UsersApiCustomerInviteUser

> UsersApiCustomerInviteUser(ctx).CustomerInviteUserRequest2(customerInviteUserRequest2).Execute()

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
	customerInviteUserRequest2 := *openapiclient.NewCustomerInviteUserRequest2("abc@example.com") // CustomerInviteUserRequest2 | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	r, err := apiClient.UsersApiAPI.UsersApiCustomerInviteUser(context.Background()).CustomerInviteUserRequest2(customerInviteUserRequest2).Execute()
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
 **customerInviteUserRequest2** | [**CustomerInviteUserRequest2**](CustomerInviteUserRequest2.md) |  | 

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

> CustomerLoginWithIdentityProviderResult UsersApiCustomerLoginWithIdentityProvider(ctx).CustomerLoginWithIdentityProviderRequest2(customerLoginWithIdentityProviderRequest2).Execute()

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
	customerLoginWithIdentityProviderRequest2 := *openapiclient.NewCustomerLoginWithIdentityProviderRequest2("4/P7q7W91a-oMsCeLvIaQm6bTrgtp7&") // CustomerLoginWithIdentityProviderRequest2 | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.UsersApiAPI.UsersApiCustomerLoginWithIdentityProvider(context.Background()).CustomerLoginWithIdentityProviderRequest2(customerLoginWithIdentityProviderRequest2).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `UsersApiAPI.UsersApiCustomerLoginWithIdentityProvider``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `UsersApiCustomerLoginWithIdentityProvider`: CustomerLoginWithIdentityProviderResult
	fmt.Fprintf(os.Stdout, "Response from `UsersApiAPI.UsersApiCustomerLoginWithIdentityProvider`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiUsersApiCustomerLoginWithIdentityProviderRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **customerLoginWithIdentityProviderRequest2** | [**CustomerLoginWithIdentityProviderRequest2**](CustomerLoginWithIdentityProviderRequest2.md) |  | 

### Return type

[**CustomerLoginWithIdentityProviderResult**](CustomerLoginWithIdentityProviderResult.md)

### Authorization

[api_key_header_Authorization](../README.md#api_key_header_Authorization)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json, application/vnd.goa.error

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## UsersApiCustomerResetPassword

> UsersApiCustomerResetPassword(ctx).CustomerResetPasswordRequest2(customerResetPasswordRequest2).Execute()

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
	customerResetPasswordRequest2 := *openapiclient.NewCustomerResetPasswordRequest2("abc@example.com") // CustomerResetPasswordRequest2 | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	r, err := apiClient.UsersApiAPI.UsersApiCustomerResetPassword(context.Background()).CustomerResetPasswordRequest2(customerResetPasswordRequest2).Execute()
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
 **customerResetPasswordRequest2** | [**CustomerResetPasswordRequest2**](CustomerResetPasswordRequest2.md) |  | 

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

> CustomerSigninResult UsersApiCustomerSignin(ctx).CustomerSigninRequest2(customerSigninRequest2).Execute()

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
	customerSigninRequest2 := *openapiclient.NewCustomerSigninRequest2("abc@gmail.com") // CustomerSigninRequest2 | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.UsersApiAPI.UsersApiCustomerSignin(context.Background()).CustomerSigninRequest2(customerSigninRequest2).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `UsersApiAPI.UsersApiCustomerSignin``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `UsersApiCustomerSignin`: CustomerSigninResult
	fmt.Fprintf(os.Stdout, "Response from `UsersApiAPI.UsersApiCustomerSignin`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiUsersApiCustomerSigninRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **customerSigninRequest2** | [**CustomerSigninRequest2**](CustomerSigninRequest2.md) |  | 

### Return type

[**CustomerSigninResult**](CustomerSigninResult.md)

### Authorization

[api_key_header_Authorization](../README.md#api_key_header_Authorization)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json, application/vnd.goa.error

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## UsersApiCustomerSignup

> UsersApiCustomerSignup(ctx).CustomerSignupRequest2(customerSignupRequest2).Execute()

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
	customerSignupRequest2 := *openapiclient.NewCustomerSignupRequest2("abc@gmail.com", "John Doe", "password") // CustomerSignupRequest2 | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	r, err := apiClient.UsersApiAPI.UsersApiCustomerSignup(context.Background()).CustomerSignupRequest2(customerSignupRequest2).Execute()
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
 **customerSignupRequest2** | [**CustomerSignupRequest2**](CustomerSignupRequest2.md) |  | 

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

> UsersApiInviteUser(ctx).InviteUserRequest2(inviteUserRequest2).Execute()

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
	inviteUserRequest2 := *openapiclient.NewInviteUserRequest2("abc@gmail.com", "root|editor|reader|service_editor|service_reader|admin|service_operator") // InviteUserRequest2 | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	r, err := apiClient.UsersApiAPI.UsersApiInviteUser(context.Background()).InviteUserRequest2(inviteUserRequest2).Execute()
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
 **inviteUserRequest2** | [**InviteUserRequest2**](InviteUserRequest2.md) |  | 

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

> UsersApiRevokeUserRole(ctx).RevokeUserRoleRequest2(revokeUserRoleRequest2).Execute()

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
	revokeUserRoleRequest2 := *openapiclient.NewRevokeUserRoleRequest2("abc@gmail.com", "root|editor|reader|service_editor|service_reader|admin|service_operator") // RevokeUserRoleRequest2 | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	r, err := apiClient.UsersApiAPI.UsersApiRevokeUserRole(context.Background()).RevokeUserRoleRequest2(revokeUserRoleRequest2).Execute()
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
 **revokeUserRoleRequest2** | [**RevokeUserRoleRequest2**](RevokeUserRoleRequest2.md) |  | 

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

> UsersApiUpdatePassword(ctx).UpdatePasswordRequest2(updatePasswordRequest2).Execute()

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
	updatePasswordRequest2 := *openapiclient.NewUpdatePasswordRequest2("password") // UpdatePasswordRequest2 | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	r, err := apiClient.UsersApiAPI.UsersApiUpdatePassword(context.Background()).UpdatePasswordRequest2(updatePasswordRequest2).Execute()
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
 **updatePasswordRequest2** | [**UpdatePasswordRequest2**](UpdatePasswordRequest2.md) |  | 

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

> UsersApiUpdateUser(ctx, id).UpdateUserRequest2(updateUserRequest2).Execute()

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
	updateUserRequest2 := *openapiclient.NewUpdateUserRequest2() // UpdateUserRequest2 | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	r, err := apiClient.UsersApiAPI.UsersApiUpdateUser(context.Background(), id).UpdateUserRequest2(updateUserRequest2).Execute()
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

 **updateUserRequest2** | [**UpdateUserRequest2**](UpdateUserRequest2.md) |  | 

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

