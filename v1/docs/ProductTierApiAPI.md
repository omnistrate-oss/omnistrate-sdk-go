# \ProductTierApiAPI

All URIs are relative to *https://api.omnistrate.cloud*

Method | HTTP request | Description
------------- | ------------- | -------------
[**ProductTierApiCopyProductTier**](ProductTierApiAPI.md#ProductTierApiCopyProductTier) | **Post** /2022-09-01-00/service/{serviceId}/product-tier/copy/{sourceId} | CopyProductTier product-tier-api
[**ProductTierApiCreateProductTier**](ProductTierApiAPI.md#ProductTierApiCreateProductTier) | **Post** /2022-09-01-00/service/{serviceId}/product-tier | CreateProductTier product-tier-api
[**ProductTierApiDeleteProductTier**](ProductTierApiAPI.md#ProductTierApiDeleteProductTier) | **Delete** /2022-09-01-00/service/{serviceId}/product-tier/{id} | DeleteProductTier product-tier-api
[**ProductTierApiDescribeProductTier**](ProductTierApiAPI.md#ProductTierApiDescribeProductTier) | **Get** /2022-09-01-00/service/{serviceId}/product-tier/{id} | DescribeProductTier product-tier-api
[**ProductTierApiDisableProductTierFeature**](ProductTierApiAPI.md#ProductTierApiDisableProductTierFeature) | **Delete** /2022-09-01-00/service/{serviceId}/product-tier/{id}/feature | DisableProductTierFeature product-tier-api
[**ProductTierApiEnableProductTierFeature**](ProductTierApiAPI.md#ProductTierApiEnableProductTierFeature) | **Put** /2022-09-01-00/service/{serviceId}/product-tier/{id}/feature | EnableProductTierFeature product-tier-api
[**ProductTierApiListProductTier**](ProductTierApiAPI.md#ProductTierApiListProductTier) | **Get** /2022-09-01-00/service/{serviceId}/model/{serviceModelId}/product-tier | ListProductTier product-tier-api
[**ProductTierApiUpdateProductTier**](ProductTierApiAPI.md#ProductTierApiUpdateProductTier) | **Patch** /2022-09-01-00/service/{serviceId}/product-tier/{id} | UpdateProductTier product-tier-api



## ProductTierApiCopyProductTier

> string ProductTierApiCopyProductTier(ctx, serviceId, sourceId).CopyProductTierRequest2(copyProductTierRequest2).Execute()

CopyProductTier product-tier-api

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
	sourceId := "pt-12345678" // string | The source product tier ID
	copyProductTierRequest2 := *openapiclient.NewCopyProductTierRequest2("A premium product tier", "Premium", "Amet temporibus.") // CopyProductTierRequest2 | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.ProductTierApiAPI.ProductTierApiCopyProductTier(context.Background(), serviceId, sourceId).CopyProductTierRequest2(copyProductTierRequest2).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `ProductTierApiAPI.ProductTierApiCopyProductTier``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `ProductTierApiCopyProductTier`: string
	fmt.Fprintf(os.Stdout, "Response from `ProductTierApiAPI.ProductTierApiCopyProductTier`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**serviceId** | **string** | Service ID | 
**sourceId** | **string** | The source product tier ID | 

### Other Parameters

Other parameters are passed through a pointer to a apiProductTierApiCopyProductTierRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


 **copyProductTierRequest2** | [**CopyProductTierRequest2**](CopyProductTierRequest2.md) |  | 

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


## ProductTierApiCreateProductTier

> string ProductTierApiCreateProductTier(ctx, serviceId).CreateProductTierRequest2(createProductTierRequest2).Execute()

CreateProductTier product-tier-api

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
	createProductTierRequest2 := *openapiclient.NewCreateProductTierRequest2("A premium product tier", "Premium", "A premium plan", "Autem ut sapiente dolor perspiciatis sit pariatur.", "OMNISTRATE_DEDICATED_TENANCY|OMNISTRATE_MULTI_TENANCY|CUSTOM_TENANCY") // CreateProductTierRequest2 | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.ProductTierApiAPI.ProductTierApiCreateProductTier(context.Background(), serviceId).CreateProductTierRequest2(createProductTierRequest2).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `ProductTierApiAPI.ProductTierApiCreateProductTier``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `ProductTierApiCreateProductTier`: string
	fmt.Fprintf(os.Stdout, "Response from `ProductTierApiAPI.ProductTierApiCreateProductTier`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**serviceId** | **string** | Service ID | 

### Other Parameters

Other parameters are passed through a pointer to a apiProductTierApiCreateProductTierRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **createProductTierRequest2** | [**CreateProductTierRequest2**](CreateProductTierRequest2.md) |  | 

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


## ProductTierApiDeleteProductTier

> ProductTierApiDeleteProductTier(ctx, serviceId, id).Execute()

DeleteProductTier product-tier-api

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
	id := "pt-12345678" // string | Product tier ID

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	r, err := apiClient.ProductTierApiAPI.ProductTierApiDeleteProductTier(context.Background(), serviceId, id).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `ProductTierApiAPI.ProductTierApiDeleteProductTier``: %v\n", err)
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

Other parameters are passed through a pointer to a apiProductTierApiDeleteProductTierRequest struct via the builder pattern


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


## ProductTierApiDescribeProductTier

> DescribeProductTierResult ProductTierApiDescribeProductTier(ctx, serviceId, id).Version(version).Execute()

DescribeProductTier product-tier-api

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
	id := "pt-12345678" // string | Product tier ID
	version := "3.0" // string | The version number for the specific version set (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.ProductTierApiAPI.ProductTierApiDescribeProductTier(context.Background(), serviceId, id).Version(version).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `ProductTierApiAPI.ProductTierApiDescribeProductTier``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `ProductTierApiDescribeProductTier`: DescribeProductTierResult
	fmt.Fprintf(os.Stdout, "Response from `ProductTierApiAPI.ProductTierApiDescribeProductTier`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**serviceId** | **string** | Service ID | 
**id** | **string** | Product tier ID | 

### Other Parameters

Other parameters are passed through a pointer to a apiProductTierApiDescribeProductTierRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


 **version** | **string** | The version number for the specific version set | 

### Return type

[**DescribeProductTierResult**](DescribeProductTierResult.md)

### Authorization

[api_key_header_Authorization](../README.md#api_key_header_Authorization)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json, application/vnd.goa.error

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## ProductTierApiDisableProductTierFeature

> ProductTierApiDisableProductTierFeature(ctx, serviceId, id).DisableProductTierFeatureRequest2(disableProductTierFeatureRequest2).Execute()

DisableProductTierFeature product-tier-api

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
	id := "pt-12345678" // string | Product tier ID
	disableProductTierFeatureRequest2 := *openapiclient.NewDisableProductTierFeatureRequest2("LOGS|METRICS|CLOUD_INSURANCE|MARKETPLACE|OPERATIONAL_STATUS|COMPLIANCE|APPLICATION_SECURITY") // DisableProductTierFeatureRequest2 | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	r, err := apiClient.ProductTierApiAPI.ProductTierApiDisableProductTierFeature(context.Background(), serviceId, id).DisableProductTierFeatureRequest2(disableProductTierFeatureRequest2).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `ProductTierApiAPI.ProductTierApiDisableProductTierFeature``: %v\n", err)
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

Other parameters are passed through a pointer to a apiProductTierApiDisableProductTierFeatureRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


 **disableProductTierFeatureRequest2** | [**DisableProductTierFeatureRequest2**](DisableProductTierFeatureRequest2.md) |  | 

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


## ProductTierApiEnableProductTierFeature

> ProductTierApiEnableProductTierFeature(ctx, serviceId, id).EnableProductTierFeatureRequest2(enableProductTierFeatureRequest2).Execute()

EnableProductTierFeature product-tier-api

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
	id := "pt-12345678" // string | Product tier ID
	enableProductTierFeatureRequest2 := *openapiclient.NewEnableProductTierFeatureRequest2("LOGS|METRICS|CLOUD_INSURANCE|MARKETPLACE|OPERATIONAL_STATUS|COMPLIANCE|APPLICATION_SECURITY") // EnableProductTierFeatureRequest2 | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	r, err := apiClient.ProductTierApiAPI.ProductTierApiEnableProductTierFeature(context.Background(), serviceId, id).EnableProductTierFeatureRequest2(enableProductTierFeatureRequest2).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `ProductTierApiAPI.ProductTierApiEnableProductTierFeature``: %v\n", err)
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

Other parameters are passed through a pointer to a apiProductTierApiEnableProductTierFeatureRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


 **enableProductTierFeatureRequest2** | [**EnableProductTierFeatureRequest2**](EnableProductTierFeatureRequest2.md) |  | 

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


## ProductTierApiListProductTier

> ListProductTiersResult ProductTierApiListProductTier(ctx, serviceId, serviceModelId).Execute()

ListProductTier product-tier-api

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
	serviceModelId := "sm-12345678" // string | Service model ID

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.ProductTierApiAPI.ProductTierApiListProductTier(context.Background(), serviceId, serviceModelId).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `ProductTierApiAPI.ProductTierApiListProductTier``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `ProductTierApiListProductTier`: ListProductTiersResult
	fmt.Fprintf(os.Stdout, "Response from `ProductTierApiAPI.ProductTierApiListProductTier`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**serviceId** | **string** | Service ID | 
**serviceModelId** | **string** | Service model ID | 

### Other Parameters

Other parameters are passed through a pointer to a apiProductTierApiListProductTierRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------



### Return type

[**ListProductTiersResult**](ListProductTiersResult.md)

### Authorization

[api_key_header_Authorization](../README.md#api_key_header_Authorization)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json, application/vnd.goa.error

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## ProductTierApiUpdateProductTier

> ProductTierApiUpdateProductTier(ctx, serviceId, id).UpdateProductTierRequest2(updateProductTierRequest2).Execute()

UpdateProductTier product-tier-api

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
	id := "pt-12345678" // string | Product tier ID
	updateProductTierRequest2 := *openapiclient.NewUpdateProductTierRequest2() // UpdateProductTierRequest2 | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	r, err := apiClient.ProductTierApiAPI.ProductTierApiUpdateProductTier(context.Background(), serviceId, id).UpdateProductTierRequest2(updateProductTierRequest2).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `ProductTierApiAPI.ProductTierApiUpdateProductTier``: %v\n", err)
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

Other parameters are passed through a pointer to a apiProductTierApiUpdateProductTierRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


 **updateProductTierRequest2** | [**UpdateProductTierRequest2**](UpdateProductTierRequest2.md) |  | 

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

