# \IntegrationsApiAPI

All URIs are relative to *https://api.omnistrate.cloud*

Method | HTTP request | Description
------------- | ------------- | -------------
[**IntegrationsApiDisableProductTierIntegration**](IntegrationsApiAPI.md#IntegrationsApiDisableProductTierIntegration) | **Delete** /2022-09-01-00/fleet/service/{serviceId}/product-tier/{id}/integration | DisableProductTierIntegration integrations-api
[**IntegrationsApiEnableProductTierIntegration**](IntegrationsApiAPI.md#IntegrationsApiEnableProductTierIntegration) | **Put** /2022-09-01-00/fleet/service/{serviceId}/product-tier/{id}/integration | EnableProductTierIntegration integrations-api
[**IntegrationsApiListProductTierIntegrations**](IntegrationsApiAPI.md#IntegrationsApiListProductTierIntegrations) | **Get** /2022-09-01-00/fleet/service/{serviceId}/product-tier/{id}/integration | ListProductTierIntegrations integrations-api



## IntegrationsApiDisableProductTierIntegration

> IntegrationsApiDisableProductTierIntegration(ctx, serviceId, id).DisableProductTierIntegrationRequestBody(disableProductTierIntegrationRequestBody).Execute()

DisableProductTierIntegration integrations-api

### Example

```go
package main

import (
	"context"
	"fmt"
	"os"
	openapiclient "github.com/omnistrate/omnistrate-sdk-go/fleet"
)

func main() {
	serviceId := "s-12345678" // string | Service ID
	id := "pt-12345678" // string | Product tier ID
	disableProductTierIntegrationRequestBody := *openapiclient.NewDisableProductTierIntegrationRequestBody("PAGERDUTY", "ALERTS") // DisableProductTierIntegrationRequestBody | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	r, err := apiClient.IntegrationsApiAPI.IntegrationsApiDisableProductTierIntegration(context.Background(), serviceId, id).DisableProductTierIntegrationRequestBody(disableProductTierIntegrationRequestBody).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `IntegrationsApiAPI.IntegrationsApiDisableProductTierIntegration``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**serviceId** | **string** | Service ID | 
**id** | **string** | Product tier ID | 

### Other Parameters

Other parameters are passed through a pointer to a apiIntegrationsApiDisableProductTierIntegrationRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


 **disableProductTierIntegrationRequestBody** | [**DisableProductTierIntegrationRequestBody**](DisableProductTierIntegrationRequestBody.md) |  | 

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


## IntegrationsApiEnableProductTierIntegration

> IntegrationsApiEnableProductTierIntegration(ctx, serviceId, id).EnableProductTierIntegrationRequestBody(enableProductTierIntegrationRequestBody).Execute()

EnableProductTierIntegration integrations-api

### Example

```go
package main

import (
	"context"
	"fmt"
	"os"
	openapiclient "github.com/omnistrate/omnistrate-sdk-go/fleet"
)

func main() {
	serviceId := "s-12345678" // string | Service ID
	id := "pt-12345678" // string | Product tier ID
	enableProductTierIntegrationRequestBody := *openapiclient.NewEnableProductTierIntegrationRequestBody("PAGERDUTY", "ALERTS") // EnableProductTierIntegrationRequestBody | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	r, err := apiClient.IntegrationsApiAPI.IntegrationsApiEnableProductTierIntegration(context.Background(), serviceId, id).EnableProductTierIntegrationRequestBody(enableProductTierIntegrationRequestBody).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `IntegrationsApiAPI.IntegrationsApiEnableProductTierIntegration``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**serviceId** | **string** | Service ID | 
**id** | **string** | Product tier ID | 

### Other Parameters

Other parameters are passed through a pointer to a apiIntegrationsApiEnableProductTierIntegrationRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


 **enableProductTierIntegrationRequestBody** | [**EnableProductTierIntegrationRequestBody**](EnableProductTierIntegrationRequestBody.md) |  | 

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


## IntegrationsApiListProductTierIntegrations

> ListProductTierIntegrationsResult IntegrationsApiListProductTierIntegrations(ctx, serviceId, id).Execute()

ListProductTierIntegrations integrations-api

### Example

```go
package main

import (
	"context"
	"fmt"
	"os"
	openapiclient "github.com/omnistrate/omnistrate-sdk-go/fleet"
)

func main() {
	serviceId := "s-12345678" // string | Service ID
	id := "pt-12345678" // string | Product tier ID

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.IntegrationsApiAPI.IntegrationsApiListProductTierIntegrations(context.Background(), serviceId, id).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `IntegrationsApiAPI.IntegrationsApiListProductTierIntegrations``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `IntegrationsApiListProductTierIntegrations`: ListProductTierIntegrationsResult
	fmt.Fprintf(os.Stdout, "Response from `IntegrationsApiAPI.IntegrationsApiListProductTierIntegrations`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**serviceId** | **string** | Service ID | 
**id** | **string** | Product tier ID | 

### Other Parameters

Other parameters are passed through a pointer to a apiIntegrationsApiListProductTierIntegrationsRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------



### Return type

[**ListProductTierIntegrationsResult**](ListProductTierIntegrationsResult.md)

### Authorization

[api_key_header_Authorization](../README.md#api_key_header_Authorization)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json, application/vnd.goa.error

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

