# \ProductTierApiAPI

All URIs are relative to *https://api.omnistrate.cloud*

Method | HTTP request | Description
------------- | ------------- | -------------
[**ProductTierApiCopyProductTier**](ProductTierApiAPI.md#ProductTierApiCopyProductTier) | **Post** /2022-09-01-00/service/{serviceId}/product-tier/copy/{sourceId} | CopyProductTier product-tier-api
[**ProductTierApiCreateProductTier**](ProductTierApiAPI.md#ProductTierApiCreateProductTier) | **Post** /2022-09-01-00/service/{serviceId}/product-tier | CreateProductTier product-tier-api
[**ProductTierApiCreateProductTierBillingPlan**](ProductTierApiAPI.md#ProductTierApiCreateProductTierBillingPlan) | **Post** /2022-09-01-00/service/{serviceId}/product-tier/{id}/billing-plan | CreateProductTierBillingPlan product-tier-api
[**ProductTierApiDeleteProductTier**](ProductTierApiAPI.md#ProductTierApiDeleteProductTier) | **Delete** /2022-09-01-00/service/{serviceId}/product-tier/{id} | DeleteProductTier product-tier-api
[**ProductTierApiDeleteProductTierBillingPlan**](ProductTierApiAPI.md#ProductTierApiDeleteProductTierBillingPlan) | **Delete** /2022-09-01-00/service/{serviceId}/product-tier/{productTierId}/billing-plan/{id} | DeleteProductTierBillingPlan product-tier-api
[**ProductTierApiDescribeProductTier**](ProductTierApiAPI.md#ProductTierApiDescribeProductTier) | **Get** /2022-09-01-00/service/{serviceId}/product-tier/{id} | DescribeProductTier product-tier-api
[**ProductTierApiDescribeProductTierBillingPlan**](ProductTierApiAPI.md#ProductTierApiDescribeProductTierBillingPlan) | **Get** /2022-09-01-00/service/{serviceId}/product-tier/{productTierId}/billing-plan/{id} | DescribeProductTierBillingPlan product-tier-api
[**ProductTierApiDisableProductTierFeature**](ProductTierApiAPI.md#ProductTierApiDisableProductTierFeature) | **Delete** /2022-09-01-00/service/{serviceId}/product-tier/{id}/feature | DisableProductTierFeature product-tier-api
[**ProductTierApiEnableProductTierFeature**](ProductTierApiAPI.md#ProductTierApiEnableProductTierFeature) | **Put** /2022-09-01-00/service/{serviceId}/product-tier/{id}/feature | EnableProductTierFeature product-tier-api
[**ProductTierApiListProductTier**](ProductTierApiAPI.md#ProductTierApiListProductTier) | **Get** /2022-09-01-00/service/{serviceId}/model/{serviceModelId}/product-tier | ListProductTier product-tier-api
[**ProductTierApiListProductTierBillingPlan**](ProductTierApiAPI.md#ProductTierApiListProductTierBillingPlan) | **Get** /2022-09-01-00/service/{serviceId}/product-tier/{id}/billing-plan | ListProductTierBillingPlan product-tier-api
[**ProductTierApiUpdateProductTier**](ProductTierApiAPI.md#ProductTierApiUpdateProductTier) | **Patch** /2022-09-01-00/service/{serviceId}/product-tier/{id} | UpdateProductTier product-tier-api
[**ProductTierApiUpdateProductTierBillingPlan**](ProductTierApiAPI.md#ProductTierApiUpdateProductTierBillingPlan) | **Patch** /2022-09-01-00/service/{serviceId}/product-tier/{productTierId}/billing-plan/{id} | UpdateProductTierBillingPlan product-tier-api



## ProductTierApiCopyProductTier

> string ProductTierApiCopyProductTier(ctx, serviceId, sourceId).CopyProductTierRequestBody(copyProductTierRequestBody).Execute()

CopyProductTier product-tier-api

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
	serviceId := "s-12345678" // string | Service ID
	sourceId := "pt-12345678" // string | The source product tier ID
	copyProductTierRequestBody := *openapiclient.NewCopyProductTierRequestBody("A premium product tier", "Premium", "sm-12345678") // CopyProductTierRequestBody | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.ProductTierApiAPI.ProductTierApiCopyProductTier(context.Background(), serviceId, sourceId).CopyProductTierRequestBody(copyProductTierRequestBody).Execute()
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


 **copyProductTierRequestBody** | [**CopyProductTierRequestBody**](CopyProductTierRequestBody.md) |  | 

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

> string ProductTierApiCreateProductTier(ctx, serviceId).CreateProductTierRequestBody(createProductTierRequestBody).Execute()

CreateProductTier product-tier-api

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
	serviceId := "s-12345678" // string | Service ID
	createProductTierRequestBody := *openapiclient.NewCreateProductTierRequestBody("A premium product tier", "Premium", "A premium plan", "sm-12345678", "OMNISTRATE_DEDICATED_TENANCY") // CreateProductTierRequestBody | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.ProductTierApiAPI.ProductTierApiCreateProductTier(context.Background(), serviceId).CreateProductTierRequestBody(createProductTierRequestBody).Execute()
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

 **createProductTierRequestBody** | [**CreateProductTierRequestBody**](CreateProductTierRequestBody.md) |  | 

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


## ProductTierApiCreateProductTierBillingPlan

> BillingPlan ProductTierApiCreateProductTierBillingPlan(ctx, serviceId, id).CreateProductTierBillingPlanRequestBody(createProductTierBillingPlanRequestBody).Execute()

CreateProductTierBillingPlan product-tier-api

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
	serviceId := "s-12345678" // string | Service ID
	id := "pt-12345678" // string | Product tier ID
	createProductTierBillingPlanRequestBody := *openapiclient.NewCreateProductTierBillingPlanRequestBody(true, int64(5), "STARTER", interface{}({"cpuCoreHours":"0.001","memoryGiBHours":"0.0001","storageGiBHours":"0.0001"})) // CreateProductTierBillingPlanRequestBody | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.ProductTierApiAPI.ProductTierApiCreateProductTierBillingPlan(context.Background(), serviceId, id).CreateProductTierBillingPlanRequestBody(createProductTierBillingPlanRequestBody).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `ProductTierApiAPI.ProductTierApiCreateProductTierBillingPlan``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `ProductTierApiCreateProductTierBillingPlan`: BillingPlan
	fmt.Fprintf(os.Stdout, "Response from `ProductTierApiAPI.ProductTierApiCreateProductTierBillingPlan`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**serviceId** | **string** | Service ID | 
**id** | **string** | Product tier ID | 

### Other Parameters

Other parameters are passed through a pointer to a apiProductTierApiCreateProductTierBillingPlanRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


 **createProductTierBillingPlanRequestBody** | [**CreateProductTierBillingPlanRequestBody**](CreateProductTierBillingPlanRequestBody.md) |  | 

### Return type

[**BillingPlan**](BillingPlan.md)

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
	openapiclient "github.com/omnistrate/omnistrate-sdk-go/v1"
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


## ProductTierApiDeleteProductTierBillingPlan

> ProductTierApiDeleteProductTierBillingPlan(ctx, serviceId, productTierId, id).Execute()

DeleteProductTierBillingPlan product-tier-api

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
	serviceId := "s-12345678" // string | Service ID
	productTierId := "pt-12345678" // string | Product tier ID
	id := "bp-12345678" // string | Product tier billing plan ID

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	r, err := apiClient.ProductTierApiAPI.ProductTierApiDeleteProductTierBillingPlan(context.Background(), serviceId, productTierId, id).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `ProductTierApiAPI.ProductTierApiDeleteProductTierBillingPlan``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**serviceId** | **string** | Service ID | 
**productTierId** | **string** | Product tier ID | 
**id** | **string** | Product tier billing plan ID | 

### Other Parameters

Other parameters are passed through a pointer to a apiProductTierApiDeleteProductTierBillingPlanRequest struct via the builder pattern


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
	openapiclient "github.com/omnistrate/omnistrate-sdk-go/v1"
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


## ProductTierApiDescribeProductTierBillingPlan

> BillingPlan ProductTierApiDescribeProductTierBillingPlan(ctx, serviceId, productTierId, id).Execute()

DescribeProductTierBillingPlan product-tier-api

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
	serviceId := "s-12345678" // string | Service ID
	productTierId := "pt-12345678" // string | Product tier ID
	id := "bp-12345678" // string | Product tier billing plan ID

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.ProductTierApiAPI.ProductTierApiDescribeProductTierBillingPlan(context.Background(), serviceId, productTierId, id).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `ProductTierApiAPI.ProductTierApiDescribeProductTierBillingPlan``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `ProductTierApiDescribeProductTierBillingPlan`: BillingPlan
	fmt.Fprintf(os.Stdout, "Response from `ProductTierApiAPI.ProductTierApiDescribeProductTierBillingPlan`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**serviceId** | **string** | Service ID | 
**productTierId** | **string** | Product tier ID | 
**id** | **string** | Product tier billing plan ID | 

### Other Parameters

Other parameters are passed through a pointer to a apiProductTierApiDescribeProductTierBillingPlanRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------




### Return type

[**BillingPlan**](BillingPlan.md)

### Authorization

[api_key_header_Authorization](../README.md#api_key_header_Authorization)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json, application/vnd.goa.error

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## ProductTierApiDisableProductTierFeature

> ProductTierApiDisableProductTierFeature(ctx, serviceId, id).DisableProductTierFeatureRequestBody(disableProductTierFeatureRequestBody).Execute()

DisableProductTierFeature product-tier-api

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
	serviceId := "s-12345678" // string | Service ID
	id := "pt-12345678" // string | Product tier ID
	disableProductTierFeatureRequestBody := *openapiclient.NewDisableProductTierFeatureRequestBody("BILLING") // DisableProductTierFeatureRequestBody | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	r, err := apiClient.ProductTierApiAPI.ProductTierApiDisableProductTierFeature(context.Background(), serviceId, id).DisableProductTierFeatureRequestBody(disableProductTierFeatureRequestBody).Execute()
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


 **disableProductTierFeatureRequestBody** | [**DisableProductTierFeatureRequestBody**](DisableProductTierFeatureRequestBody.md) |  | 

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

> ProductTierApiEnableProductTierFeature(ctx, serviceId, id).EnableProductTierFeatureRequestBody(enableProductTierFeatureRequestBody).Execute()

EnableProductTierFeature product-tier-api

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
	serviceId := "s-12345678" // string | Service ID
	id := "pt-12345678" // string | Product tier ID
	enableProductTierFeatureRequestBody := *openapiclient.NewEnableProductTierFeatureRequestBody("BILLING") // EnableProductTierFeatureRequestBody | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	r, err := apiClient.ProductTierApiAPI.ProductTierApiEnableProductTierFeature(context.Background(), serviceId, id).EnableProductTierFeatureRequestBody(enableProductTierFeatureRequestBody).Execute()
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


 **enableProductTierFeatureRequestBody** | [**EnableProductTierFeatureRequestBody**](EnableProductTierFeatureRequestBody.md) |  | 

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

> ListServiceEnvironmentsResult ProductTierApiListProductTier(ctx, serviceId, serviceModelId).Execute()

ListProductTier product-tier-api

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
	serviceId := "s-12345678" // string | Service ID
	serviceModelId := "sm-12345678" // string | Service model ID

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.ProductTierApiAPI.ProductTierApiListProductTier(context.Background(), serviceId, serviceModelId).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `ProductTierApiAPI.ProductTierApiListProductTier``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `ProductTierApiListProductTier`: ListServiceEnvironmentsResult
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

[**ListServiceEnvironmentsResult**](ListServiceEnvironmentsResult.md)

### Authorization

[api_key_header_Authorization](../README.md#api_key_header_Authorization)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json, application/vnd.goa.error

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## ProductTierApiListProductTierBillingPlan

> ListProductTierBillingPlanResult ProductTierApiListProductTierBillingPlan(ctx, serviceId, id).Execute()

ListProductTierBillingPlan product-tier-api

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
	serviceId := "s-12345678" // string | Service ID
	id := "pt-12345678" // string | Product tier ID

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.ProductTierApiAPI.ProductTierApiListProductTierBillingPlan(context.Background(), serviceId, id).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `ProductTierApiAPI.ProductTierApiListProductTierBillingPlan``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `ProductTierApiListProductTierBillingPlan`: ListProductTierBillingPlanResult
	fmt.Fprintf(os.Stdout, "Response from `ProductTierApiAPI.ProductTierApiListProductTierBillingPlan`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**serviceId** | **string** | Service ID | 
**id** | **string** | Product tier ID | 

### Other Parameters

Other parameters are passed through a pointer to a apiProductTierApiListProductTierBillingPlanRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------



### Return type

[**ListProductTierBillingPlanResult**](ListProductTierBillingPlanResult.md)

### Authorization

[api_key_header_Authorization](../README.md#api_key_header_Authorization)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json, application/vnd.goa.error

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## ProductTierApiUpdateProductTier

> ProductTierApiUpdateProductTier(ctx, serviceId, id).UpdateProductTierRequestBody(updateProductTierRequestBody).Execute()

UpdateProductTier product-tier-api

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
	serviceId := "s-12345678" // string | Service ID
	id := "pt-12345678" // string | Product tier ID
	updateProductTierRequestBody := *openapiclient.NewUpdateProductTierRequestBody() // UpdateProductTierRequestBody | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	r, err := apiClient.ProductTierApiAPI.ProductTierApiUpdateProductTier(context.Background(), serviceId, id).UpdateProductTierRequestBody(updateProductTierRequestBody).Execute()
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


 **updateProductTierRequestBody** | [**UpdateProductTierRequestBody**](UpdateProductTierRequestBody.md) |  | 

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


## ProductTierApiUpdateProductTierBillingPlan

> BillingPlan ProductTierApiUpdateProductTierBillingPlan(ctx, serviceId, productTierId, id).UpdateProductTierBillingPlanRequestBody(updateProductTierBillingPlanRequestBody).Execute()

UpdateProductTierBillingPlan product-tier-api

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
	serviceId := "s-12345678" // string | Service ID
	productTierId := "pt-12345678" // string | Product tier ID
	id := "bp-12345678" // string | Product tier billing plan ID
	updateProductTierBillingPlanRequestBody := *openapiclient.NewUpdateProductTierBillingPlanRequestBody() // UpdateProductTierBillingPlanRequestBody | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.ProductTierApiAPI.ProductTierApiUpdateProductTierBillingPlan(context.Background(), serviceId, productTierId, id).UpdateProductTierBillingPlanRequestBody(updateProductTierBillingPlanRequestBody).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `ProductTierApiAPI.ProductTierApiUpdateProductTierBillingPlan``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `ProductTierApiUpdateProductTierBillingPlan`: BillingPlan
	fmt.Fprintf(os.Stdout, "Response from `ProductTierApiAPI.ProductTierApiUpdateProductTierBillingPlan`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**serviceId** | **string** | Service ID | 
**productTierId** | **string** | Product tier ID | 
**id** | **string** | Product tier billing plan ID | 

### Other Parameters

Other parameters are passed through a pointer to a apiProductTierApiUpdateProductTierBillingPlanRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------



 **updateProductTierBillingPlanRequestBody** | [**UpdateProductTierBillingPlanRequestBody**](UpdateProductTierBillingPlanRequestBody.md) |  | 

### Return type

[**BillingPlan**](BillingPlan.md)

### Authorization

[api_key_header_Authorization](../README.md#api_key_header_Authorization)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json, application/vnd.goa.error

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

