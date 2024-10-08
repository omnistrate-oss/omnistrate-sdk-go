# \ServicePlanApiAPI

All URIs are relative to *https://api.omnistrate.cloud*

Method | HTTP request | Description
------------- | ------------- | -------------
[**ServicePlanApiGetServicePlan**](ServicePlanApiAPI.md#ServicePlanApiGetServicePlan) | **Get** /2022-09-01-00/service/{serviceId}/product-tier/{productTierId}/service-plan | GetServicePlan service-plan-api
[**ServicePlanApiListServicePlans**](ServicePlanApiAPI.md#ServicePlanApiListServicePlans) | **Get** /2022-09-01-00/service/{serviceId}/environment/{serviceEnvironmentId}/service-plan | ListServicePlans service-plan-api



## ServicePlanApiGetServicePlan

> GetServicePlanResult ServicePlanApiGetServicePlan(ctx, serviceId, productTierId).Execute()

GetServicePlan service-plan-api

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
	serviceId := "s-12345678" // string | Service ID
	productTierId := "pt-12345678" // string | Product tier ID

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.ServicePlanApiAPI.ServicePlanApiGetServicePlan(context.Background(), serviceId, productTierId).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `ServicePlanApiAPI.ServicePlanApiGetServicePlan``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `ServicePlanApiGetServicePlan`: GetServicePlanResult
	fmt.Fprintf(os.Stdout, "Response from `ServicePlanApiAPI.ServicePlanApiGetServicePlan`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**serviceId** | **string** | Service ID | 
**productTierId** | **string** | Product tier ID | 

### Other Parameters

Other parameters are passed through a pointer to a apiServicePlanApiGetServicePlanRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------



### Return type

[**GetServicePlanResult**](GetServicePlanResult.md)

### Authorization

[api_key_header_Authorization](../README.md#api_key_header_Authorization)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json, application/vnd.goa.error

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## ServicePlanApiListServicePlans

> ListServicePlansResult ServicePlanApiListServicePlans(ctx, serviceId, serviceEnvironmentId).Execute()

ListServicePlans service-plan-api

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
	serviceId := "s-12345678" // string | Service ID
	serviceEnvironmentId := "se-12345678" // string | Service environment ID

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.ServicePlanApiAPI.ServicePlanApiListServicePlans(context.Background(), serviceId, serviceEnvironmentId).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `ServicePlanApiAPI.ServicePlanApiListServicePlans``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `ServicePlanApiListServicePlans`: ListServicePlansResult
	fmt.Fprintf(os.Stdout, "Response from `ServicePlanApiAPI.ServicePlanApiListServicePlans`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**serviceId** | **string** | Service ID | 
**serviceEnvironmentId** | **string** | Service environment ID | 

### Other Parameters

Other parameters are passed through a pointer to a apiServicePlanApiListServicePlansRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------



### Return type

[**ListServicePlansResult**](ListServicePlansResult.md)

### Authorization

[api_key_header_Authorization](../README.md#api_key_header_Authorization)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json, application/vnd.goa.error

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

