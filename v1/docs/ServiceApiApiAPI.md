# \ServiceApiApiAPI

All URIs are relative to *https://api.omnistrate.cloud*

Method | HTTP request | Description
------------- | ------------- | -------------
[**ServiceApiApiCreateServiceAPI**](ServiceApiApiAPI.md#ServiceApiApiCreateServiceAPI) | **Post** /2022-09-01-00/service/{serviceId}/service-api | CreateServiceAPI service-api-api
[**ServiceApiApiDeleteServiceAPI**](ServiceApiApiAPI.md#ServiceApiApiDeleteServiceAPI) | **Delete** /2022-09-01-00/service/{serviceId}/service-api/{id} | DeleteServiceAPI service-api-api
[**ServiceApiApiDeprecateServiceAPI**](ServiceApiApiAPI.md#ServiceApiApiDeprecateServiceAPI) | **Post** /2022-09-01-00/service/{serviceId}/service-api/{id}/deprecate | DeprecateServiceAPI service-api-api
[**ServiceApiApiDescribePendingChanges**](ServiceApiApiAPI.md#ServiceApiApiDescribePendingChanges) | **Get** /2022-09-01-00/service/{serviceId}/service-api/{id}/all-pending-changes | DescribePendingChanges service-api-api
[**ServiceApiApiDescribeServiceAPI**](ServiceApiApiAPI.md#ServiceApiApiDescribeServiceAPI) | **Get** /2022-09-01-00/service/{serviceId}/service-api/{id} | DescribeServiceAPI service-api-api
[**ServiceApiApiDiscardPendingChanges**](ServiceApiApiAPI.md#ServiceApiApiDiscardPendingChanges) | **Delete** /2022-09-01-00/service/{serviceId}/service-api/{id}/all-pending-changes | DiscardPendingChanges service-api-api
[**ServiceApiApiListServiceAPI**](ServiceApiApiAPI.md#ServiceApiApiListServiceAPI) | **Get** /2022-09-01-00/service/{serviceId}/serviceenvironment/{serviceEnvironmentId}/service-api | ListServiceAPI service-api-api
[**ServiceApiApiReleaseServiceAPI**](ServiceApiApiAPI.md#ServiceApiApiReleaseServiceAPI) | **Post** /2022-09-01-00/service/{serviceId}/service-api/{id}/release | ReleaseServiceAPI service-api-api
[**ServiceApiApiUpdateServiceAPI**](ServiceApiApiAPI.md#ServiceApiApiUpdateServiceAPI) | **Patch** /2022-09-01-00/service/{serviceId}/service-api/{id} | UpdateServiceAPI service-api-api



## ServiceApiApiCreateServiceAPI

> string ServiceApiApiCreateServiceAPI(ctx, serviceId).CreateServiceAPIRequest2(createServiceAPIRequest2).Execute()

CreateServiceAPI service-api-api

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
	serviceId := "s-12345678" // string | The service ID that this API bundle belongs to
	createServiceAPIRequest2 := *openapiclient.NewCreateServiceAPIRequest2("A MySQL SaaS API specializing in multi-writer multi-tenant clusters for high availability", "se-123456") // CreateServiceAPIRequest2 | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.ServiceApiApiAPI.ServiceApiApiCreateServiceAPI(context.Background(), serviceId).CreateServiceAPIRequest2(createServiceAPIRequest2).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `ServiceApiApiAPI.ServiceApiApiCreateServiceAPI``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `ServiceApiApiCreateServiceAPI`: string
	fmt.Fprintf(os.Stdout, "Response from `ServiceApiApiAPI.ServiceApiApiCreateServiceAPI`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**serviceId** | **string** | The service ID that this API bundle belongs to | 

### Other Parameters

Other parameters are passed through a pointer to a apiServiceApiApiCreateServiceAPIRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **createServiceAPIRequest2** | [**CreateServiceAPIRequest2**](CreateServiceAPIRequest2.md) |  | 

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


## ServiceApiApiDeleteServiceAPI

> ServiceApiApiDeleteServiceAPI(ctx, serviceId, id).Execute()

DeleteServiceAPI service-api-api

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
	serviceId := "s-12345678" // string | The service ID that this API bundle belongs to
	id := "sa-12345678" // string | The service API ID

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	r, err := apiClient.ServiceApiApiAPI.ServiceApiApiDeleteServiceAPI(context.Background(), serviceId, id).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `ServiceApiApiAPI.ServiceApiApiDeleteServiceAPI``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**serviceId** | **string** | The service ID that this API bundle belongs to | 
**id** | **string** | The service API ID | 

### Other Parameters

Other parameters are passed through a pointer to a apiServiceApiApiDeleteServiceAPIRequest struct via the builder pattern


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


## ServiceApiApiDeprecateServiceAPI

> ServiceApiApiDeprecateServiceAPI(ctx, serviceId, id).Execute()

DeprecateServiceAPI service-api-api

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
	serviceId := "s-12345678" // string | The service ID that this API bundle belongs to
	id := "sa-12345678" // string | The service API ID

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	r, err := apiClient.ServiceApiApiAPI.ServiceApiApiDeprecateServiceAPI(context.Background(), serviceId, id).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `ServiceApiApiAPI.ServiceApiApiDeprecateServiceAPI``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**serviceId** | **string** | The service ID that this API bundle belongs to | 
**id** | **string** | The service API ID | 

### Other Parameters

Other parameters are passed through a pointer to a apiServiceApiApiDeprecateServiceAPIRequest struct via the builder pattern


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


## ServiceApiApiDescribePendingChanges

> DescribePendingChangesResult ServiceApiApiDescribePendingChanges(ctx, serviceId, id).ProductTierId(productTierId).Execute()

DescribePendingChanges service-api-api

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
	serviceId := "s-12345678" // string | The service ID that this API bundle belongs to
	id := "sa-12345678" // string | The service API ID
	productTierId := "pt-12345678" // string | ProductTierID of the resource to describe pending changes forNeeds to specified in combination with the product tier version (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.ServiceApiApiAPI.ServiceApiApiDescribePendingChanges(context.Background(), serviceId, id).ProductTierId(productTierId).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `ServiceApiApiAPI.ServiceApiApiDescribePendingChanges``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `ServiceApiApiDescribePendingChanges`: DescribePendingChangesResult
	fmt.Fprintf(os.Stdout, "Response from `ServiceApiApiAPI.ServiceApiApiDescribePendingChanges`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**serviceId** | **string** | The service ID that this API bundle belongs to | 
**id** | **string** | The service API ID | 

### Other Parameters

Other parameters are passed through a pointer to a apiServiceApiApiDescribePendingChangesRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


 **productTierId** | **string** | ProductTierID of the resource to describe pending changes forNeeds to specified in combination with the product tier version | 

### Return type

[**DescribePendingChangesResult**](DescribePendingChangesResult.md)

### Authorization

[api_key_header_Authorization](../README.md#api_key_header_Authorization)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json, application/vnd.goa.error

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## ServiceApiApiDescribeServiceAPI

> DescribeServiceAPIResult ServiceApiApiDescribeServiceAPI(ctx, serviceId, id).Execute()

DescribeServiceAPI service-api-api

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
	serviceId := "s-12345678" // string | The service ID that this API bundle belongs to
	id := "sa-12345678" // string | The service API ID

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.ServiceApiApiAPI.ServiceApiApiDescribeServiceAPI(context.Background(), serviceId, id).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `ServiceApiApiAPI.ServiceApiApiDescribeServiceAPI``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `ServiceApiApiDescribeServiceAPI`: DescribeServiceAPIResult
	fmt.Fprintf(os.Stdout, "Response from `ServiceApiApiAPI.ServiceApiApiDescribeServiceAPI`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**serviceId** | **string** | The service ID that this API bundle belongs to | 
**id** | **string** | The service API ID | 

### Other Parameters

Other parameters are passed through a pointer to a apiServiceApiApiDescribeServiceAPIRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------



### Return type

[**DescribeServiceAPIResult**](DescribeServiceAPIResult.md)

### Authorization

[api_key_header_Authorization](../README.md#api_key_header_Authorization)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json, application/vnd.goa.error

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## ServiceApiApiDiscardPendingChanges

> ServiceApiApiDiscardPendingChanges(ctx, serviceId, id).ProductTierId(productTierId).Execute()

DiscardPendingChanges service-api-api

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
	serviceId := "s-12345678" // string | The service ID that this API bundle belongs to
	id := "sa-12345678" // string | The service API ID
	productTierId := "pt-12345678" // string | ProductTierID of the resource to describe pending changes forNeeds to specified in combination with the product tier version (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	r, err := apiClient.ServiceApiApiAPI.ServiceApiApiDiscardPendingChanges(context.Background(), serviceId, id).ProductTierId(productTierId).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `ServiceApiApiAPI.ServiceApiApiDiscardPendingChanges``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**serviceId** | **string** | The service ID that this API bundle belongs to | 
**id** | **string** | The service API ID | 

### Other Parameters

Other parameters are passed through a pointer to a apiServiceApiApiDiscardPendingChangesRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


 **productTierId** | **string** | ProductTierID of the resource to describe pending changes forNeeds to specified in combination with the product tier version | 

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


## ServiceApiApiListServiceAPI

> ListServiceAPIsResult ServiceApiApiListServiceAPI(ctx, serviceId, serviceEnvironmentId).Execute()

ListServiceAPI service-api-api

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
	serviceId := "s-12345678" // string | The service ID that this API bundle belongs to
	serviceEnvironmentId := "se-12345678" // string | The service environment ID

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.ServiceApiApiAPI.ServiceApiApiListServiceAPI(context.Background(), serviceId, serviceEnvironmentId).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `ServiceApiApiAPI.ServiceApiApiListServiceAPI``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `ServiceApiApiListServiceAPI`: ListServiceAPIsResult
	fmt.Fprintf(os.Stdout, "Response from `ServiceApiApiAPI.ServiceApiApiListServiceAPI`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**serviceId** | **string** | The service ID that this API bundle belongs to | 
**serviceEnvironmentId** | **string** | The service environment ID | 

### Other Parameters

Other parameters are passed through a pointer to a apiServiceApiApiListServiceAPIRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------



### Return type

[**ListServiceAPIsResult**](ListServiceAPIsResult.md)

### Authorization

[api_key_header_Authorization](../README.md#api_key_header_Authorization)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json, application/vnd.goa.error

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## ServiceApiApiReleaseServiceAPI

> ServiceApiApiReleaseServiceAPI(ctx, serviceId, id).ReleaseServiceAPIRequest2(releaseServiceAPIRequest2).Execute()

ReleaseServiceAPI service-api-api

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
	serviceId := "s-12345678" // string | The service ID that this API bundle belongs to
	id := "sa-12345678" // string | The service API ID
	releaseServiceAPIRequest2 := *openapiclient.NewReleaseServiceAPIRequest2() // ReleaseServiceAPIRequest2 | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	r, err := apiClient.ServiceApiApiAPI.ServiceApiApiReleaseServiceAPI(context.Background(), serviceId, id).ReleaseServiceAPIRequest2(releaseServiceAPIRequest2).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `ServiceApiApiAPI.ServiceApiApiReleaseServiceAPI``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**serviceId** | **string** | The service ID that this API bundle belongs to | 
**id** | **string** | The service API ID | 

### Other Parameters

Other parameters are passed through a pointer to a apiServiceApiApiReleaseServiceAPIRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


 **releaseServiceAPIRequest2** | [**ReleaseServiceAPIRequest2**](ReleaseServiceAPIRequest2.md) |  | 

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


## ServiceApiApiUpdateServiceAPI

> ServiceApiApiUpdateServiceAPI(ctx, serviceId, id).UpdateServiceAPIRequest2(updateServiceAPIRequest2).Execute()

UpdateServiceAPI service-api-api

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
	serviceId := "s-12345678" // string | The service ID that this API bundle belongs to
	id := "sa-12345678" // string | The service API ID
	updateServiceAPIRequest2 := *openapiclient.NewUpdateServiceAPIRequest2() // UpdateServiceAPIRequest2 | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	r, err := apiClient.ServiceApiApiAPI.ServiceApiApiUpdateServiceAPI(context.Background(), serviceId, id).UpdateServiceAPIRequest2(updateServiceAPIRequest2).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `ServiceApiApiAPI.ServiceApiApiUpdateServiceAPI``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**serviceId** | **string** | The service ID that this API bundle belongs to | 
**id** | **string** | The service API ID | 

### Other Parameters

Other parameters are passed through a pointer to a apiServiceApiApiUpdateServiceAPIRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


 **updateServiceAPIRequest2** | [**UpdateServiceAPIRequest2**](UpdateServiceAPIRequest2.md) |  | 

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

