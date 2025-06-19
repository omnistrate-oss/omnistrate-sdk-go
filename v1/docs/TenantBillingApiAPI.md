# \TenantBillingApiAPI

All URIs are relative to *https://api.omnistrate.cloud*

Method | HTTP request | Description
------------- | ------------- | -------------
[**TenantBillingApiDisableBillingProvider**](TenantBillingApiAPI.md#TenantBillingApiDisableBillingProvider) | **Delete** /2022-09-01-00/tenant-billing/disable/{billingProviderType} | DisableBillingProvider tenant-billing-api
[**TenantBillingApiDisableTenantBilling**](TenantBillingApiAPI.md#TenantBillingApiDisableTenantBilling) | **Delete** /2022-09-01-00/tenant-billing/disable | DisableTenantBilling tenant-billing-api
[**TenantBillingApiEnableBillingProvider**](TenantBillingApiAPI.md#TenantBillingApiEnableBillingProvider) | **Post** /2022-09-01-00/tenant-billing/enable/{billingProviderType} | EnableBillingProvider tenant-billing-api
[**TenantBillingApiEnableTenantBilling**](TenantBillingApiAPI.md#TenantBillingApiEnableTenantBilling) | **Post** /2022-09-01-00/tenant-billing/enable | EnableTenantBilling tenant-billing-api
[**TenantBillingApiGetTenantBillingStatus**](TenantBillingApiAPI.md#TenantBillingApiGetTenantBillingStatus) | **Get** /2022-09-01-00/tenant-billing | GetTenantBillingStatus tenant-billing-api



## TenantBillingApiDisableBillingProvider

> TenantBillingApiDisableBillingProvider(ctx, billingProviderType).Execute()

DisableBillingProvider tenant-billing-api

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
	billingProviderType := "STRIPE" // string | The type of billing provider to enable

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	r, err := apiClient.TenantBillingApiAPI.TenantBillingApiDisableBillingProvider(context.Background(), billingProviderType).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `TenantBillingApiAPI.TenantBillingApiDisableBillingProvider``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**billingProviderType** | **string** | The type of billing provider to enable | 

### Other Parameters

Other parameters are passed through a pointer to a apiTenantBillingApiDisableBillingProviderRequest struct via the builder pattern


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


## TenantBillingApiDisableTenantBilling

> TenantBillingApiDisableTenantBilling(ctx).Execute()

DisableTenantBilling tenant-billing-api

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
	r, err := apiClient.TenantBillingApiAPI.TenantBillingApiDisableTenantBilling(context.Background()).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `TenantBillingApiAPI.TenantBillingApiDisableTenantBilling``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
}
```

### Path Parameters

This endpoint does not need any parameter.

### Other Parameters

Other parameters are passed through a pointer to a apiTenantBillingApiDisableTenantBillingRequest struct via the builder pattern


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


## TenantBillingApiEnableBillingProvider

> TenantBillingApiEnableBillingProvider(ctx, billingProviderType).EnableBillingProviderRequest2(enableBillingProviderRequest2).Execute()

EnableBillingProvider tenant-billing-api

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
	billingProviderType := "STRIPE" // string | The type of billing provider to enable
	enableBillingProviderRequest2 := *openapiclient.NewEnableBillingProviderRequest2() // EnableBillingProviderRequest2 | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	r, err := apiClient.TenantBillingApiAPI.TenantBillingApiEnableBillingProvider(context.Background(), billingProviderType).EnableBillingProviderRequest2(enableBillingProviderRequest2).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `TenantBillingApiAPI.TenantBillingApiEnableBillingProvider``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**billingProviderType** | **string** | The type of billing provider to enable | 

### Other Parameters

Other parameters are passed through a pointer to a apiTenantBillingApiEnableBillingProviderRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **enableBillingProviderRequest2** | [**EnableBillingProviderRequest2**](EnableBillingProviderRequest2.md) |  | 

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


## TenantBillingApiEnableTenantBilling

> TenantBillingApiEnableTenantBilling(ctx).Execute()

EnableTenantBilling tenant-billing-api

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
	r, err := apiClient.TenantBillingApiAPI.TenantBillingApiEnableTenantBilling(context.Background()).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `TenantBillingApiAPI.TenantBillingApiEnableTenantBilling``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
}
```

### Path Parameters

This endpoint does not need any parameter.

### Other Parameters

Other parameters are passed through a pointer to a apiTenantBillingApiEnableTenantBillingRequest struct via the builder pattern


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


## TenantBillingApiGetTenantBillingStatus

> GetTenantBillingStatusResult TenantBillingApiGetTenantBillingStatus(ctx).Execute()

GetTenantBillingStatus tenant-billing-api

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
	resp, r, err := apiClient.TenantBillingApiAPI.TenantBillingApiGetTenantBillingStatus(context.Background()).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `TenantBillingApiAPI.TenantBillingApiGetTenantBillingStatus``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `TenantBillingApiGetTenantBillingStatus`: GetTenantBillingStatusResult
	fmt.Fprintf(os.Stdout, "Response from `TenantBillingApiAPI.TenantBillingApiGetTenantBillingStatus`: %v\n", resp)
}
```

### Path Parameters

This endpoint does not need any parameter.

### Other Parameters

Other parameters are passed through a pointer to a apiTenantBillingApiGetTenantBillingStatusRequest struct via the builder pattern


### Return type

[**GetTenantBillingStatusResult**](GetTenantBillingStatusResult.md)

### Authorization

[api_key_header_Authorization](../README.md#api_key_header_Authorization)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json, application/vnd.goa.error

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

