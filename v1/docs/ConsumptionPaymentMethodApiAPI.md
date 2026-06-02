# \ConsumptionPaymentMethodApiAPI

All URIs are relative to *https://api.omnistrate.cloud*

Method | HTTP request | Description
------------- | ------------- | -------------
[**ConsumptionPaymentMethodApiCreateConsumptionSetupIntent**](ConsumptionPaymentMethodApiAPI.md#ConsumptionPaymentMethodApiCreateConsumptionSetupIntent) | **Post** /2022-09-01-00/resource-instance/billing/stripe/payment-methods/setup-intent | CreateConsumptionSetupIntent consumption-payment-method-api
[**ConsumptionPaymentMethodApiGetConsumptionStripeConfig**](ConsumptionPaymentMethodApiAPI.md#ConsumptionPaymentMethodApiGetConsumptionStripeConfig) | **Get** /2022-09-01-00/resource-instance/billing/stripe/config | GetConsumptionStripeConfig consumption-payment-method-api
[**ConsumptionPaymentMethodApiListConsumptionPaymentMethods**](ConsumptionPaymentMethodApiAPI.md#ConsumptionPaymentMethodApiListConsumptionPaymentMethods) | **Get** /2022-09-01-00/resource-instance/billing/stripe/payment-methods | ListConsumptionPaymentMethods consumption-payment-method-api
[**ConsumptionPaymentMethodApiRemoveConsumptionPaymentMethod**](ConsumptionPaymentMethodApiAPI.md#ConsumptionPaymentMethodApiRemoveConsumptionPaymentMethod) | **Delete** /2022-09-01-00/resource-instance/billing/stripe/payment-methods/{id} | RemoveConsumptionPaymentMethod consumption-payment-method-api
[**ConsumptionPaymentMethodApiSetDefaultConsumptionPaymentMethod**](ConsumptionPaymentMethodApiAPI.md#ConsumptionPaymentMethodApiSetDefaultConsumptionPaymentMethod) | **Post** /2022-09-01-00/resource-instance/billing/stripe/payment-methods/{id}/default | SetDefaultConsumptionPaymentMethod consumption-payment-method-api



## ConsumptionPaymentMethodApiCreateConsumptionSetupIntent

> CreateSetupIntentResult ConsumptionPaymentMethodApiCreateConsumptionSetupIntent(ctx).Execute()

CreateConsumptionSetupIntent consumption-payment-method-api



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
	resp, r, err := apiClient.ConsumptionPaymentMethodApiAPI.ConsumptionPaymentMethodApiCreateConsumptionSetupIntent(context.Background()).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `ConsumptionPaymentMethodApiAPI.ConsumptionPaymentMethodApiCreateConsumptionSetupIntent``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `ConsumptionPaymentMethodApiCreateConsumptionSetupIntent`: CreateSetupIntentResult
	fmt.Fprintf(os.Stdout, "Response from `ConsumptionPaymentMethodApiAPI.ConsumptionPaymentMethodApiCreateConsumptionSetupIntent`: %v\n", resp)
}
```

### Path Parameters

This endpoint does not need any parameter.

### Other Parameters

Other parameters are passed through a pointer to a apiConsumptionPaymentMethodApiCreateConsumptionSetupIntentRequest struct via the builder pattern


### Return type

[**CreateSetupIntentResult**](CreateSetupIntentResult.md)

### Authorization

[api_key_header_Authorization](../README.md#api_key_header_Authorization)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json, application/vnd.goa.error

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## ConsumptionPaymentMethodApiGetConsumptionStripeConfig

> StripeConfigResult ConsumptionPaymentMethodApiGetConsumptionStripeConfig(ctx).Execute()

GetConsumptionStripeConfig consumption-payment-method-api



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
	resp, r, err := apiClient.ConsumptionPaymentMethodApiAPI.ConsumptionPaymentMethodApiGetConsumptionStripeConfig(context.Background()).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `ConsumptionPaymentMethodApiAPI.ConsumptionPaymentMethodApiGetConsumptionStripeConfig``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `ConsumptionPaymentMethodApiGetConsumptionStripeConfig`: StripeConfigResult
	fmt.Fprintf(os.Stdout, "Response from `ConsumptionPaymentMethodApiAPI.ConsumptionPaymentMethodApiGetConsumptionStripeConfig`: %v\n", resp)
}
```

### Path Parameters

This endpoint does not need any parameter.

### Other Parameters

Other parameters are passed through a pointer to a apiConsumptionPaymentMethodApiGetConsumptionStripeConfigRequest struct via the builder pattern


### Return type

[**StripeConfigResult**](StripeConfigResult.md)

### Authorization

[api_key_header_Authorization](../README.md#api_key_header_Authorization)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json, application/vnd.goa.error

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## ConsumptionPaymentMethodApiListConsumptionPaymentMethods

> ListPaymentMethodsResult ConsumptionPaymentMethodApiListConsumptionPaymentMethods(ctx).Execute()

ListConsumptionPaymentMethods consumption-payment-method-api



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
	resp, r, err := apiClient.ConsumptionPaymentMethodApiAPI.ConsumptionPaymentMethodApiListConsumptionPaymentMethods(context.Background()).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `ConsumptionPaymentMethodApiAPI.ConsumptionPaymentMethodApiListConsumptionPaymentMethods``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `ConsumptionPaymentMethodApiListConsumptionPaymentMethods`: ListPaymentMethodsResult
	fmt.Fprintf(os.Stdout, "Response from `ConsumptionPaymentMethodApiAPI.ConsumptionPaymentMethodApiListConsumptionPaymentMethods`: %v\n", resp)
}
```

### Path Parameters

This endpoint does not need any parameter.

### Other Parameters

Other parameters are passed through a pointer to a apiConsumptionPaymentMethodApiListConsumptionPaymentMethodsRequest struct via the builder pattern


### Return type

[**ListPaymentMethodsResult**](ListPaymentMethodsResult.md)

### Authorization

[api_key_header_Authorization](../README.md#api_key_header_Authorization)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json, application/vnd.goa.error

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## ConsumptionPaymentMethodApiRemoveConsumptionPaymentMethod

> ConsumptionPaymentMethodApiRemoveConsumptionPaymentMethod(ctx, id).Execute()

RemoveConsumptionPaymentMethod consumption-payment-method-api



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
	id := "pm_1NVChw2eZvKYlo2CwkOranTh" // string | Payment method ID to remove

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	r, err := apiClient.ConsumptionPaymentMethodApiAPI.ConsumptionPaymentMethodApiRemoveConsumptionPaymentMethod(context.Background(), id).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `ConsumptionPaymentMethodApiAPI.ConsumptionPaymentMethodApiRemoveConsumptionPaymentMethod``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **string** | Payment method ID to remove | 

### Other Parameters

Other parameters are passed through a pointer to a apiConsumptionPaymentMethodApiRemoveConsumptionPaymentMethodRequest struct via the builder pattern


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


## ConsumptionPaymentMethodApiSetDefaultConsumptionPaymentMethod

> ConsumptionPaymentMethodApiSetDefaultConsumptionPaymentMethod(ctx, id).Execute()

SetDefaultConsumptionPaymentMethod consumption-payment-method-api



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
	id := "pm_1NVChw2eZvKYlo2CwkOranTh" // string | Payment method ID to set as default

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	r, err := apiClient.ConsumptionPaymentMethodApiAPI.ConsumptionPaymentMethodApiSetDefaultConsumptionPaymentMethod(context.Background(), id).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `ConsumptionPaymentMethodApiAPI.ConsumptionPaymentMethodApiSetDefaultConsumptionPaymentMethod``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **string** | Payment method ID to set as default | 

### Other Parameters

Other parameters are passed through a pointer to a apiConsumptionPaymentMethodApiSetDefaultConsumptionPaymentMethodRequest struct via the builder pattern


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

