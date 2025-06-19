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
	openapiclient "github.com/omnistrate-oss/omnistrate-sdk-go/v1"
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

> ListServicePlansResult ServicePlanApiListServicePlans(ctx, serviceId, serviceEnvironmentId).NextPageToken(nextPageToken).PageSize(pageSize).SkipHasPendingChangesCheck(skipHasPendingChangesCheck).Execute()

ListServicePlans service-plan-api

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
	serviceId := "s-12345678" // string | Service ID
	serviceEnvironmentId := "se-12345678" // string | Service environment ID
	nextPageToken := "token" // string |  (optional)
	pageSize := int64(10) // int64 |  (optional)
	skipHasPendingChangesCheck := false // bool | Skip the check for pending changes in the service plans (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.ServicePlanApiAPI.ServicePlanApiListServicePlans(context.Background(), serviceId, serviceEnvironmentId).NextPageToken(nextPageToken).PageSize(pageSize).SkipHasPendingChangesCheck(skipHasPendingChangesCheck).Execute()
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


 **nextPageToken** | **string** |  | 
 **pageSize** | **int64** |  | 
 **skipHasPendingChangesCheck** | **bool** | Skip the check for pending changes in the service plans | 

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

