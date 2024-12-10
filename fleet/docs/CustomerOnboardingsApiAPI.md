# \CustomerOnboardingsApiAPI

All URIs are relative to *https://api.omnistrate.cloud*

Method | HTTP request | Description
------------- | ------------- | -------------
[**CustomerOnboardingsApiCreateCustomerOnboarding**](CustomerOnboardingsApiAPI.md#CustomerOnboardingsApiCreateCustomerOnboarding) | **Post** /2022-09-01-00/fleet/customer-onboarding | CreateCustomerOnboarding customer-onboardings-api
[**CustomerOnboardingsApiDeleteCustomerOnboarding**](CustomerOnboardingsApiAPI.md#CustomerOnboardingsApiDeleteCustomerOnboarding) | **Delete** /2022-09-01-00/fleet/customer-onboarding/{id} | DeleteCustomerOnboarding customer-onboardings-api
[**CustomerOnboardingsApiDescribeCustomerOnboarding**](CustomerOnboardingsApiAPI.md#CustomerOnboardingsApiDescribeCustomerOnboarding) | **Get** /2022-09-01-00/fleet/customer-onboarding/{id} | DescribeCustomerOnboarding customer-onboardings-api
[**CustomerOnboardingsApiListCustomerOnboardingStages**](CustomerOnboardingsApiAPI.md#CustomerOnboardingsApiListCustomerOnboardingStages) | **Get** /2022-09-01-00/fleet/customer-onboarding-stages | ListCustomerOnboardingStages customer-onboardings-api
[**CustomerOnboardingsApiListCustomerOnboardings**](CustomerOnboardingsApiAPI.md#CustomerOnboardingsApiListCustomerOnboardings) | **Get** /2022-09-01-00/fleet/customer-onboarding | ListCustomerOnboardings customer-onboardings-api
[**CustomerOnboardingsApiUpdateCustomerOnboarding**](CustomerOnboardingsApiAPI.md#CustomerOnboardingsApiUpdateCustomerOnboarding) | **Patch** /2022-09-01-00/fleet/customer-onboarding/{id} | UpdateCustomerOnboarding customer-onboardings-api



## CustomerOnboardingsApiCreateCustomerOnboarding

> string CustomerOnboardingsApiCreateCustomerOnboarding(ctx).CreateCustomerOnboardingRequest2(createCustomerOnboardingRequest2).Execute()

CreateCustomerOnboarding customer-onboardings-api

### Example

```go
package main

import (
	"context"
	"fmt"
	"os"
	openapiclient "github.com/omnistrate-oss/omnistrate-sdk-go/fleet"
)

func main() {
	createCustomerOnboardingRequest2 := *openapiclient.NewCreateCustomerOnboardingRequest2() // CreateCustomerOnboardingRequest2 | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.CustomerOnboardingsApiAPI.CustomerOnboardingsApiCreateCustomerOnboarding(context.Background()).CreateCustomerOnboardingRequest2(createCustomerOnboardingRequest2).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `CustomerOnboardingsApiAPI.CustomerOnboardingsApiCreateCustomerOnboarding``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `CustomerOnboardingsApiCreateCustomerOnboarding`: string
	fmt.Fprintf(os.Stdout, "Response from `CustomerOnboardingsApiAPI.CustomerOnboardingsApiCreateCustomerOnboarding`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiCustomerOnboardingsApiCreateCustomerOnboardingRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **createCustomerOnboardingRequest2** | [**CreateCustomerOnboardingRequest2**](CreateCustomerOnboardingRequest2.md) |  | 

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


## CustomerOnboardingsApiDeleteCustomerOnboarding

> CustomerOnboardingsApiDeleteCustomerOnboarding(ctx, id).Execute()

DeleteCustomerOnboarding customer-onboardings-api

### Example

```go
package main

import (
	"context"
	"fmt"
	"os"
	openapiclient "github.com/omnistrate-oss/omnistrate-sdk-go/fleet"
)

func main() {
	id := "onboarding-1234567890" // string | The ID of the onboarding.

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	r, err := apiClient.CustomerOnboardingsApiAPI.CustomerOnboardingsApiDeleteCustomerOnboarding(context.Background(), id).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `CustomerOnboardingsApiAPI.CustomerOnboardingsApiDeleteCustomerOnboarding``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **string** | The ID of the onboarding. | 

### Other Parameters

Other parameters are passed through a pointer to a apiCustomerOnboardingsApiDeleteCustomerOnboardingRequest struct via the builder pattern


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


## CustomerOnboardingsApiDescribeCustomerOnboarding

> CustomerOnboarding CustomerOnboardingsApiDescribeCustomerOnboarding(ctx, id).Execute()

DescribeCustomerOnboarding customer-onboardings-api

### Example

```go
package main

import (
	"context"
	"fmt"
	"os"
	openapiclient "github.com/omnistrate-oss/omnistrate-sdk-go/fleet"
)

func main() {
	id := "onboarding-1234567890" // string | The ID of the onboarding.

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.CustomerOnboardingsApiAPI.CustomerOnboardingsApiDescribeCustomerOnboarding(context.Background(), id).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `CustomerOnboardingsApiAPI.CustomerOnboardingsApiDescribeCustomerOnboarding``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `CustomerOnboardingsApiDescribeCustomerOnboarding`: CustomerOnboarding
	fmt.Fprintf(os.Stdout, "Response from `CustomerOnboardingsApiAPI.CustomerOnboardingsApiDescribeCustomerOnboarding`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **string** | The ID of the onboarding. | 

### Other Parameters

Other parameters are passed through a pointer to a apiCustomerOnboardingsApiDescribeCustomerOnboardingRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

[**CustomerOnboarding**](CustomerOnboarding.md)

### Authorization

[api_key_header_Authorization](../README.md#api_key_header_Authorization)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json, application/vnd.goa.error

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## CustomerOnboardingsApiListCustomerOnboardingStages

> ListCustomerOnboardingStagesResult CustomerOnboardingsApiListCustomerOnboardingStages(ctx).Execute()

ListCustomerOnboardingStages customer-onboardings-api

### Example

```go
package main

import (
	"context"
	"fmt"
	"os"
	openapiclient "github.com/omnistrate-oss/omnistrate-sdk-go/fleet"
)

func main() {

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.CustomerOnboardingsApiAPI.CustomerOnboardingsApiListCustomerOnboardingStages(context.Background()).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `CustomerOnboardingsApiAPI.CustomerOnboardingsApiListCustomerOnboardingStages``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `CustomerOnboardingsApiListCustomerOnboardingStages`: ListCustomerOnboardingStagesResult
	fmt.Fprintf(os.Stdout, "Response from `CustomerOnboardingsApiAPI.CustomerOnboardingsApiListCustomerOnboardingStages`: %v\n", resp)
}
```

### Path Parameters

This endpoint does not need any parameter.

### Other Parameters

Other parameters are passed through a pointer to a apiCustomerOnboardingsApiListCustomerOnboardingStagesRequest struct via the builder pattern


### Return type

[**ListCustomerOnboardingStagesResult**](ListCustomerOnboardingStagesResult.md)

### Authorization

[api_key_header_Authorization](../README.md#api_key_header_Authorization)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json, application/vnd.goa.error

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## CustomerOnboardingsApiListCustomerOnboardings

> ListCustomerOnboardingResult CustomerOnboardingsApiListCustomerOnboardings(ctx).PendingOnly(pendingOnly).Execute()

ListCustomerOnboardings customer-onboardings-api

### Example

```go
package main

import (
	"context"
	"fmt"
	"os"
	openapiclient "github.com/omnistrate-oss/omnistrate-sdk-go/fleet"
)

func main() {
	pendingOnly := true // bool | Whether to return only pending onboardings. (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.CustomerOnboardingsApiAPI.CustomerOnboardingsApiListCustomerOnboardings(context.Background()).PendingOnly(pendingOnly).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `CustomerOnboardingsApiAPI.CustomerOnboardingsApiListCustomerOnboardings``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `CustomerOnboardingsApiListCustomerOnboardings`: ListCustomerOnboardingResult
	fmt.Fprintf(os.Stdout, "Response from `CustomerOnboardingsApiAPI.CustomerOnboardingsApiListCustomerOnboardings`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiCustomerOnboardingsApiListCustomerOnboardingsRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **pendingOnly** | **bool** | Whether to return only pending onboardings. | 

### Return type

[**ListCustomerOnboardingResult**](ListCustomerOnboardingResult.md)

### Authorization

[api_key_header_Authorization](../README.md#api_key_header_Authorization)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json, application/vnd.goa.error

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## CustomerOnboardingsApiUpdateCustomerOnboarding

> CustomerOnboardingsApiUpdateCustomerOnboarding(ctx, id).UpdateCustomerOnboardingRequest2(updateCustomerOnboardingRequest2).Execute()

UpdateCustomerOnboarding customer-onboardings-api

### Example

```go
package main

import (
	"context"
	"fmt"
	"os"
	openapiclient "github.com/omnistrate-oss/omnistrate-sdk-go/fleet"
)

func main() {
	id := "onboarding-1234567890" // string | The ID of the onboarding.
	updateCustomerOnboardingRequest2 := *openapiclient.NewUpdateCustomerOnboardingRequest2() // UpdateCustomerOnboardingRequest2 | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	r, err := apiClient.CustomerOnboardingsApiAPI.CustomerOnboardingsApiUpdateCustomerOnboarding(context.Background(), id).UpdateCustomerOnboardingRequest2(updateCustomerOnboardingRequest2).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `CustomerOnboardingsApiAPI.CustomerOnboardingsApiUpdateCustomerOnboarding``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **string** | The ID of the onboarding. | 

### Other Parameters

Other parameters are passed through a pointer to a apiCustomerOnboardingsApiUpdateCustomerOnboardingRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **updateCustomerOnboardingRequest2** | [**UpdateCustomerOnboardingRequest2**](UpdateCustomerOnboardingRequest2.md) |  | 

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

