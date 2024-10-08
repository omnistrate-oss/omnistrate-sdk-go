# \ConsumptionUserApiAPI

All URIs are relative to *https://api.omnistrate.cloud*

Method | HTTP request | Description
------------- | ------------- | -------------
[**ConsumptionUserApiDescribeUserBillingDetails**](ConsumptionUserApiAPI.md#ConsumptionUserApiDescribeUserBillingDetails) | **Get** /2022-09-01-00/resource-instance/user/{id}/billing-details | DescribeUserBillingDetails consumption-user-api
[**ConsumptionUserApiDescribeUsersBySubscription**](ConsumptionUserApiAPI.md#ConsumptionUserApiDescribeUsersBySubscription) | **Get** /2022-09-01-00/resource-instance/subscription/{subscriptionId}/subscription-users | DescribeUsersBySubscription consumption-user-api
[**ConsumptionUserApiInviteUser**](ConsumptionUserApiAPI.md#ConsumptionUserApiInviteUser) | **Post** /2022-09-01-00/resource-instance/subscription/{subscriptionId}/invite-user | InviteUser consumption-user-api
[**ConsumptionUserApiRevokeUserRole**](ConsumptionUserApiAPI.md#ConsumptionUserApiRevokeUserRole) | **Delete** /2022-09-01-00/resource-instance/subscription/{subscriptionId}/revoke-user-role | RevokeUserRole consumption-user-api
[**ConsumptionUserApiSignin**](ConsumptionUserApiAPI.md#ConsumptionUserApiSignin) | **Post** /2022-09-01-00/resource-instance/user/signin | Signin consumption-user-api



## ConsumptionUserApiDescribeUserBillingDetails

> DescribeUserBillingDetailsResult ConsumptionUserApiDescribeUserBillingDetails(ctx, id).Execute()

DescribeUserBillingDetails consumption-user-api

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
	resp, r, err := apiClient.ConsumptionUserApiAPI.ConsumptionUserApiDescribeUserBillingDetails(context.Background(), id).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `ConsumptionUserApiAPI.ConsumptionUserApiDescribeUserBillingDetails``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `ConsumptionUserApiDescribeUserBillingDetails`: DescribeUserBillingDetailsResult
	fmt.Fprintf(os.Stdout, "Response from `ConsumptionUserApiAPI.ConsumptionUserApiDescribeUserBillingDetails`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiConsumptionUserApiDescribeUserBillingDetailsRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

[**DescribeUserBillingDetailsResult**](DescribeUserBillingDetailsResult.md)

### Authorization

[api_key_header_Authorization](../README.md#api_key_header_Authorization)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json, application/vnd.goa.error

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## ConsumptionUserApiDescribeUsersBySubscription

> DescribeUsersBySubscriptionResult ConsumptionUserApiDescribeUsersBySubscription(ctx, subscriptionId).Execute()

DescribeUsersBySubscription consumption-user-api

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
	subscriptionId := "sub-abcd1234" // string | The subscription ID

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.ConsumptionUserApiAPI.ConsumptionUserApiDescribeUsersBySubscription(context.Background(), subscriptionId).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `ConsumptionUserApiAPI.ConsumptionUserApiDescribeUsersBySubscription``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `ConsumptionUserApiDescribeUsersBySubscription`: DescribeUsersBySubscriptionResult
	fmt.Fprintf(os.Stdout, "Response from `ConsumptionUserApiAPI.ConsumptionUserApiDescribeUsersBySubscription`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**subscriptionId** | **string** | The subscription ID | 

### Other Parameters

Other parameters are passed through a pointer to a apiConsumptionUserApiDescribeUsersBySubscriptionRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

[**DescribeUsersBySubscriptionResult**](DescribeUsersBySubscriptionResult.md)

### Authorization

[api_key_header_Authorization](../README.md#api_key_header_Authorization)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json, application/vnd.goa.error

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## ConsumptionUserApiInviteUser

> ConsumptionUserApiInviteUser(ctx, subscriptionId).InviteUserRequestBody(inviteUserRequestBody).Execute()

InviteUser consumption-user-api

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
	subscriptionId := "sub-abcd1234" // string | The subscription ID
	inviteUserRequestBody := *openapiclient.NewInviteUserRequestBody("abc@gmail.com", "reader") // InviteUserRequestBody | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	r, err := apiClient.ConsumptionUserApiAPI.ConsumptionUserApiInviteUser(context.Background(), subscriptionId).InviteUserRequestBody(inviteUserRequestBody).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `ConsumptionUserApiAPI.ConsumptionUserApiInviteUser``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**subscriptionId** | **string** | The subscription ID | 

### Other Parameters

Other parameters are passed through a pointer to a apiConsumptionUserApiInviteUserRequest struct via the builder pattern


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


## ConsumptionUserApiRevokeUserRole

> ConsumptionUserApiRevokeUserRole(ctx, subscriptionId).InviteUserRequestBody(inviteUserRequestBody).Execute()

RevokeUserRole consumption-user-api

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
	subscriptionId := "sub-abcd1234" // string | The subscription ID
	inviteUserRequestBody := *openapiclient.NewInviteUserRequestBody("abc@gmail.com", "reader") // InviteUserRequestBody | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	r, err := apiClient.ConsumptionUserApiAPI.ConsumptionUserApiRevokeUserRole(context.Background(), subscriptionId).InviteUserRequestBody(inviteUserRequestBody).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `ConsumptionUserApiAPI.ConsumptionUserApiRevokeUserRole``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**subscriptionId** | **string** | The subscription ID | 

### Other Parameters

Other parameters are passed through a pointer to a apiConsumptionUserApiRevokeUserRoleRequest struct via the builder pattern


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


## ConsumptionUserApiSignin

> ConsumptionServiceAuthResult ConsumptionUserApiSignin(ctx).Execute()

Signin consumption-user-api

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

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.ConsumptionUserApiAPI.ConsumptionUserApiSignin(context.Background()).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `ConsumptionUserApiAPI.ConsumptionUserApiSignin``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `ConsumptionUserApiSignin`: ConsumptionServiceAuthResult
	fmt.Fprintf(os.Stdout, "Response from `ConsumptionUserApiAPI.ConsumptionUserApiSignin`: %v\n", resp)
}
```

### Path Parameters

This endpoint does not need any parameter.

### Other Parameters

Other parameters are passed through a pointer to a apiConsumptionUserApiSigninRequest struct via the builder pattern


### Return type

[**ConsumptionServiceAuthResult**](ConsumptionServiceAuthResult.md)

### Authorization

[basic_header_Authorization](../README.md#basic_header_Authorization)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json, application/vnd.goa.error

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

