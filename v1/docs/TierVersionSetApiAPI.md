# \TierVersionSetApiAPI

All URIs are relative to *https://api.omnistrate.cloud*

Method | HTTP request | Description
------------- | ------------- | -------------
[**TierVersionSetApiCreateTierVersionSet**](TierVersionSetApiAPI.md#TierVersionSetApiCreateTierVersionSet) | **Post** /2022-09-01-00/service/{serviceId}/productTier/{productTierId}/version-set | CreateTierVersionSet tier-version-set-api
[**TierVersionSetApiCustomerListTierVersionSets**](TierVersionSetApiAPI.md#TierVersionSetApiCustomerListTierVersionSets) | **Get** /2022-09-01-00/service/{serviceId}/productTier/{productTierId}/customer-version-set | CustomerListTierVersionSets tier-version-set-api
[**TierVersionSetApiDeprecateTierVersionSet**](TierVersionSetApiAPI.md#TierVersionSetApiDeprecateTierVersionSet) | **Patch** /2022-09-01-00/service/{serviceId}/productTier/{productTierId}/version-set/{version}/deprecate | DeprecateTierVersionSet tier-version-set-api
[**TierVersionSetApiDescribeTierVersionSet**](TierVersionSetApiAPI.md#TierVersionSetApiDescribeTierVersionSet) | **Get** /2022-09-01-00/service/{serviceId}/productTier/{productTierId}/version-set/{version} | DescribeTierVersionSet tier-version-set-api
[**TierVersionSetApiDescribeTierVersionSetMetadata**](TierVersionSetApiAPI.md#TierVersionSetApiDescribeTierVersionSetMetadata) | **Get** /2022-09-01-00/service/{serviceId}/productTier/{productTierId}/version-set/{version}/metadata | DescribeTierVersionSetMetadata tier-version-set-api
[**TierVersionSetApiDiffTierVersionSets**](TierVersionSetApiAPI.md#TierVersionSetApiDiffTierVersionSets) | **Get** /2022-09-01-00/service/{serviceId}/productTier/{productTierId}/version-set/{version}/diff/{anotherVersion} | DiffTierVersionSets tier-version-set-api
[**TierVersionSetApiGetTierVersionSetSpec**](TierVersionSetApiAPI.md#TierVersionSetApiGetTierVersionSetSpec) | **Get** /2022-09-01-00/service/{serviceId}/productTier/{productTierId}/version-set/spec | GetTierVersionSetSpec tier-version-set-api
[**TierVersionSetApiListTierVersionSetSpecs**](TierVersionSetApiAPI.md#TierVersionSetApiListTierVersionSetSpecs) | **Get** /2022-09-01-00/service/{serviceId}/productTier/{productTierId}/version-set/specs | ListTierVersionSetSpecs tier-version-set-api
[**TierVersionSetApiListTierVersionSets**](TierVersionSetApiAPI.md#TierVersionSetApiListTierVersionSets) | **Get** /2022-09-01-00/service/{serviceId}/productTier/{productTierId}/version-set | ListTierVersionSets tier-version-set-api
[**TierVersionSetApiPromoteTierVersionSet**](TierVersionSetApiAPI.md#TierVersionSetApiPromoteTierVersionSet) | **Patch** /2022-09-01-00/service/{serviceId}/productTier/{productTierId}/version-set/{version}/promote | PromoteTierVersionSet tier-version-set-api
[**TierVersionSetApiReleaseTierVersionSet**](TierVersionSetApiAPI.md#TierVersionSetApiReleaseTierVersionSet) | **Patch** /2022-09-01-00/service/{serviceId}/productTier/{productTierId}/version-set/{version}/release | ReleaseTierVersionSet tier-version-set-api
[**TierVersionSetApiUpdateTierVersionSet**](TierVersionSetApiAPI.md#TierVersionSetApiUpdateTierVersionSet) | **Patch** /2022-09-01-00/service/{serviceId}/productTier/{productTierId}/version-set/{version} | UpdateTierVersionSet tier-version-set-api
[**TierVersionSetApiUpdateTierVersionSetMetadata**](TierVersionSetApiAPI.md#TierVersionSetApiUpdateTierVersionSetMetadata) | **Patch** /2022-09-01-00/service/{serviceId}/productTier/{productTierId}/version-set/{version}/metadata | UpdateTierVersionSetMetadata tier-version-set-api



## TierVersionSetApiCreateTierVersionSet

> CreateTierVersionSetResult TierVersionSetApiCreateTierVersionSet(ctx, serviceId, productTierId).CreateTierVersionSetRequest2(createTierVersionSetRequest2).Execute()

CreateTierVersionSet tier-version-set-api

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
	serviceId := "s-12345678" // string | ID of the Service
	productTierId := "Beatae beatae." // string | The product tier ID that this version set belongs to.
	createTierVersionSetRequest2 := *openapiclient.NewCreateTierVersionSetRequest2("Major|Incremental|UserDefined") // CreateTierVersionSetRequest2 | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.TierVersionSetApiAPI.TierVersionSetApiCreateTierVersionSet(context.Background(), serviceId, productTierId).CreateTierVersionSetRequest2(createTierVersionSetRequest2).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `TierVersionSetApiAPI.TierVersionSetApiCreateTierVersionSet``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `TierVersionSetApiCreateTierVersionSet`: CreateTierVersionSetResult
	fmt.Fprintf(os.Stdout, "Response from `TierVersionSetApiAPI.TierVersionSetApiCreateTierVersionSet`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**serviceId** | **string** | ID of the Service | 
**productTierId** | **string** | The product tier ID that this version set belongs to. | 

### Other Parameters

Other parameters are passed through a pointer to a apiTierVersionSetApiCreateTierVersionSetRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


 **createTierVersionSetRequest2** | [**CreateTierVersionSetRequest2**](CreateTierVersionSetRequest2.md) |  | 

### Return type

[**CreateTierVersionSetResult**](CreateTierVersionSetResult.md)

### Authorization

[api_key_header_Authorization](../README.md#api_key_header_Authorization)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json, application/vnd.goa.error

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## TierVersionSetApiCustomerListTierVersionSets

> CustomerListTierVersionSetsResult TierVersionSetApiCustomerListTierVersionSets(ctx, serviceId, productTierId).NextPageToken(nextPageToken).PageSize(pageSize).Execute()

CustomerListTierVersionSets tier-version-set-api

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
	serviceId := "s-12345678" // string | ID of the Service
	productTierId := "Beatae beatae." // string | The product tier ID that this version set belongs to.
	nextPageToken := "token" // string |  (optional)
	pageSize := int64(10) // int64 |  (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.TierVersionSetApiAPI.TierVersionSetApiCustomerListTierVersionSets(context.Background(), serviceId, productTierId).NextPageToken(nextPageToken).PageSize(pageSize).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `TierVersionSetApiAPI.TierVersionSetApiCustomerListTierVersionSets``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `TierVersionSetApiCustomerListTierVersionSets`: CustomerListTierVersionSetsResult
	fmt.Fprintf(os.Stdout, "Response from `TierVersionSetApiAPI.TierVersionSetApiCustomerListTierVersionSets`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**serviceId** | **string** | ID of the Service | 
**productTierId** | **string** | The product tier ID that this version set belongs to. | 

### Other Parameters

Other parameters are passed through a pointer to a apiTierVersionSetApiCustomerListTierVersionSetsRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


 **nextPageToken** | **string** |  | 
 **pageSize** | **int64** |  | 

### Return type

[**CustomerListTierVersionSetsResult**](CustomerListTierVersionSetsResult.md)

### Authorization

[api_key_header_Authorization](../README.md#api_key_header_Authorization)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json, application/vnd.goa.error

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## TierVersionSetApiDeprecateTierVersionSet

> TierVersionSet TierVersionSetApiDeprecateTierVersionSet(ctx, serviceId, productTierId, version).Execute()

DeprecateTierVersionSet tier-version-set-api

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
	serviceId := "s-12345678" // string | ID of the Service
	productTierId := "Beatae beatae." // string | The product tier ID that this version set belongs to.
	version := "3.0" // string | The version number for the specific version set.

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.TierVersionSetApiAPI.TierVersionSetApiDeprecateTierVersionSet(context.Background(), serviceId, productTierId, version).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `TierVersionSetApiAPI.TierVersionSetApiDeprecateTierVersionSet``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `TierVersionSetApiDeprecateTierVersionSet`: TierVersionSet
	fmt.Fprintf(os.Stdout, "Response from `TierVersionSetApiAPI.TierVersionSetApiDeprecateTierVersionSet`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**serviceId** | **string** | ID of the Service | 
**productTierId** | **string** | The product tier ID that this version set belongs to. | 
**version** | **string** | The version number for the specific version set. | 

### Other Parameters

Other parameters are passed through a pointer to a apiTierVersionSetApiDeprecateTierVersionSetRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------




### Return type

[**TierVersionSet**](TierVersionSet.md)

### Authorization

[api_key_header_Authorization](../README.md#api_key_header_Authorization)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json, application/vnd.goa.error

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## TierVersionSetApiDescribeTierVersionSet

> TierVersionSet TierVersionSetApiDescribeTierVersionSet(ctx, serviceId, productTierId, version).Execute()

DescribeTierVersionSet tier-version-set-api

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
	serviceId := "s-12345678" // string | ID of the Service
	productTierId := "Beatae beatae." // string | The product tier ID that this version set belongs to.
	version := "3.0" // string | The version number for the specific version set.

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.TierVersionSetApiAPI.TierVersionSetApiDescribeTierVersionSet(context.Background(), serviceId, productTierId, version).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `TierVersionSetApiAPI.TierVersionSetApiDescribeTierVersionSet``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `TierVersionSetApiDescribeTierVersionSet`: TierVersionSet
	fmt.Fprintf(os.Stdout, "Response from `TierVersionSetApiAPI.TierVersionSetApiDescribeTierVersionSet`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**serviceId** | **string** | ID of the Service | 
**productTierId** | **string** | The product tier ID that this version set belongs to. | 
**version** | **string** | The version number for the specific version set. | 

### Other Parameters

Other parameters are passed through a pointer to a apiTierVersionSetApiDescribeTierVersionSetRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------




### Return type

[**TierVersionSet**](TierVersionSet.md)

### Authorization

[api_key_header_Authorization](../README.md#api_key_header_Authorization)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json, application/vnd.goa.error

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## TierVersionSetApiDescribeTierVersionSetMetadata

> TierVersionSetMetadata TierVersionSetApiDescribeTierVersionSetMetadata(ctx, serviceId, productTierId, version).Execute()

DescribeTierVersionSetMetadata tier-version-set-api

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
	serviceId := "s-12345678" // string | ID of the Service
	productTierId := "Beatae beatae." // string | The product tier ID that this version set belongs to.
	version := "3.0" // string | The version number for the specific version set.

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.TierVersionSetApiAPI.TierVersionSetApiDescribeTierVersionSetMetadata(context.Background(), serviceId, productTierId, version).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `TierVersionSetApiAPI.TierVersionSetApiDescribeTierVersionSetMetadata``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `TierVersionSetApiDescribeTierVersionSetMetadata`: TierVersionSetMetadata
	fmt.Fprintf(os.Stdout, "Response from `TierVersionSetApiAPI.TierVersionSetApiDescribeTierVersionSetMetadata`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**serviceId** | **string** | ID of the Service | 
**productTierId** | **string** | The product tier ID that this version set belongs to. | 
**version** | **string** | The version number for the specific version set. | 

### Other Parameters

Other parameters are passed through a pointer to a apiTierVersionSetApiDescribeTierVersionSetMetadataRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------




### Return type

[**TierVersionSetMetadata**](TierVersionSetMetadata.md)

### Authorization

[api_key_header_Authorization](../README.md#api_key_header_Authorization)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json, application/vnd.goa.error

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## TierVersionSetApiDiffTierVersionSets

> DiffTierVersionSetsResult TierVersionSetApiDiffTierVersionSets(ctx, serviceId, productTierId, version, anotherVersion).Execute()

DiffTierVersionSets tier-version-set-api

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
	serviceId := "s-12345678" // string | ID of the Service
	productTierId := "Beatae beatae." // string | The product tier ID that this version set belongs to.
	version := "3.0" // string | The version number for the version set.
	anotherVersion := "3.1" // string | The target version to compare against.

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.TierVersionSetApiAPI.TierVersionSetApiDiffTierVersionSets(context.Background(), serviceId, productTierId, version, anotherVersion).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `TierVersionSetApiAPI.TierVersionSetApiDiffTierVersionSets``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `TierVersionSetApiDiffTierVersionSets`: DiffTierVersionSetsResult
	fmt.Fprintf(os.Stdout, "Response from `TierVersionSetApiAPI.TierVersionSetApiDiffTierVersionSets`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**serviceId** | **string** | ID of the Service | 
**productTierId** | **string** | The product tier ID that this version set belongs to. | 
**version** | **string** | The version number for the version set. | 
**anotherVersion** | **string** | The target version to compare against. | 

### Other Parameters

Other parameters are passed through a pointer to a apiTierVersionSetApiDiffTierVersionSetsRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------





### Return type

[**DiffTierVersionSetsResult**](DiffTierVersionSetsResult.md)

### Authorization

[api_key_header_Authorization](../README.md#api_key_header_Authorization)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json, application/vnd.goa.error

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## TierVersionSetApiGetTierVersionSetSpec

> GetTierVersionSetSpecResult TierVersionSetApiGetTierVersionSetSpec(ctx, serviceId, productTierId).Version(version).Execute()

GetTierVersionSetSpec tier-version-set-api

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
	serviceId := "s-12345678" // string | ID of the Service
	productTierId := "Beatae beatae." // string | The product tier ID that this version set belongs to.
	version := "3.0" // string | The version number for the specific version set. If not specified, returns the latest preferred or active version. (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.TierVersionSetApiAPI.TierVersionSetApiGetTierVersionSetSpec(context.Background(), serviceId, productTierId).Version(version).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `TierVersionSetApiAPI.TierVersionSetApiGetTierVersionSetSpec``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `TierVersionSetApiGetTierVersionSetSpec`: GetTierVersionSetSpecResult
	fmt.Fprintf(os.Stdout, "Response from `TierVersionSetApiAPI.TierVersionSetApiGetTierVersionSetSpec`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**serviceId** | **string** | ID of the Service | 
**productTierId** | **string** | The product tier ID that this version set belongs to. | 

### Other Parameters

Other parameters are passed through a pointer to a apiTierVersionSetApiGetTierVersionSetSpecRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


 **version** | **string** | The version number for the specific version set. If not specified, returns the latest preferred or active version. | 

### Return type

[**GetTierVersionSetSpecResult**](GetTierVersionSetSpecResult.md)

### Authorization

[api_key_header_Authorization](../README.md#api_key_header_Authorization)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json, application/vnd.goa.error

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## TierVersionSetApiListTierVersionSetSpecs

> ListTierVersionSetSpecsResult TierVersionSetApiListTierVersionSetSpecs(ctx, serviceId, productTierId).Execute()

ListTierVersionSetSpecs tier-version-set-api

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
	serviceId := "s-12345678" // string | ID of the Service
	productTierId := "Beatae beatae." // string | The product tier ID that this version set belongs to.

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.TierVersionSetApiAPI.TierVersionSetApiListTierVersionSetSpecs(context.Background(), serviceId, productTierId).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `TierVersionSetApiAPI.TierVersionSetApiListTierVersionSetSpecs``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `TierVersionSetApiListTierVersionSetSpecs`: ListTierVersionSetSpecsResult
	fmt.Fprintf(os.Stdout, "Response from `TierVersionSetApiAPI.TierVersionSetApiListTierVersionSetSpecs`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**serviceId** | **string** | ID of the Service | 
**productTierId** | **string** | The product tier ID that this version set belongs to. | 

### Other Parameters

Other parameters are passed through a pointer to a apiTierVersionSetApiListTierVersionSetSpecsRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------



### Return type

[**ListTierVersionSetSpecsResult**](ListTierVersionSetSpecsResult.md)

### Authorization

[api_key_header_Authorization](../README.md#api_key_header_Authorization)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json, application/vnd.goa.error

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## TierVersionSetApiListTierVersionSets

> ListTierVersionSetsResult TierVersionSetApiListTierVersionSets(ctx, serviceId, productTierId).LatestMajorVersionOnly(latestMajorVersionOnly).LatestIncrementalVersionForMajorVersion(latestIncrementalVersionForMajorVersion).NextPageToken(nextPageToken).PageSize(pageSize).Execute()

ListTierVersionSets tier-version-set-api

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
	serviceId := "s-12345678" // string | ID of the Service
	productTierId := "Beatae beatae." // string | The product tier ID that this version set belongs to.
	latestMajorVersionOnly := true // bool | If true, the latest major version is returned. The parameter needs to be specified in isolation. (optional)
	latestIncrementalVersionForMajorVersion := "3.0" // string | Returns the latest incremental version for the given major version. The paramenter needs to be specified in isolation. (optional)
	nextPageToken := "token" // string |  (optional)
	pageSize := int64(10) // int64 |  (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.TierVersionSetApiAPI.TierVersionSetApiListTierVersionSets(context.Background(), serviceId, productTierId).LatestMajorVersionOnly(latestMajorVersionOnly).LatestIncrementalVersionForMajorVersion(latestIncrementalVersionForMajorVersion).NextPageToken(nextPageToken).PageSize(pageSize).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `TierVersionSetApiAPI.TierVersionSetApiListTierVersionSets``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `TierVersionSetApiListTierVersionSets`: ListTierVersionSetsResult
	fmt.Fprintf(os.Stdout, "Response from `TierVersionSetApiAPI.TierVersionSetApiListTierVersionSets`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**serviceId** | **string** | ID of the Service | 
**productTierId** | **string** | The product tier ID that this version set belongs to. | 

### Other Parameters

Other parameters are passed through a pointer to a apiTierVersionSetApiListTierVersionSetsRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


 **latestMajorVersionOnly** | **bool** | If true, the latest major version is returned. The parameter needs to be specified in isolation. | 
 **latestIncrementalVersionForMajorVersion** | **string** | Returns the latest incremental version for the given major version. The paramenter needs to be specified in isolation. | 
 **nextPageToken** | **string** |  | 
 **pageSize** | **int64** |  | 

### Return type

[**ListTierVersionSetsResult**](ListTierVersionSetsResult.md)

### Authorization

[api_key_header_Authorization](../README.md#api_key_header_Authorization)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json, application/vnd.goa.error

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## TierVersionSetApiPromoteTierVersionSet

> TierVersionSet TierVersionSetApiPromoteTierVersionSet(ctx, serviceId, productTierId, version).Execute()

PromoteTierVersionSet tier-version-set-api

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
	serviceId := "s-12345678" // string | ID of the Service
	productTierId := "Beatae beatae." // string | The product tier ID that this version set belongs to.
	version := "3.0" // string | The version number for the specific version set.

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.TierVersionSetApiAPI.TierVersionSetApiPromoteTierVersionSet(context.Background(), serviceId, productTierId, version).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `TierVersionSetApiAPI.TierVersionSetApiPromoteTierVersionSet``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `TierVersionSetApiPromoteTierVersionSet`: TierVersionSet
	fmt.Fprintf(os.Stdout, "Response from `TierVersionSetApiAPI.TierVersionSetApiPromoteTierVersionSet`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**serviceId** | **string** | ID of the Service | 
**productTierId** | **string** | The product tier ID that this version set belongs to. | 
**version** | **string** | The version number for the specific version set. | 

### Other Parameters

Other parameters are passed through a pointer to a apiTierVersionSetApiPromoteTierVersionSetRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------




### Return type

[**TierVersionSet**](TierVersionSet.md)

### Authorization

[api_key_header_Authorization](../README.md#api_key_header_Authorization)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json, application/vnd.goa.error

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## TierVersionSetApiReleaseTierVersionSet

> TierVersionSet TierVersionSetApiReleaseTierVersionSet(ctx, serviceId, productTierId, version).ReleaseTierVersionSetRequest2(releaseTierVersionSetRequest2).Execute()

ReleaseTierVersionSet tier-version-set-api

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
	serviceId := "s-12345678" // string | ID of the Service
	productTierId := "Beatae beatae." // string | The product tier ID that this version set belongs to.
	version := "3.0" // string | The version number for the specific version set.
	releaseTierVersionSetRequest2 := *openapiclient.NewReleaseTierVersionSetRequest2() // ReleaseTierVersionSetRequest2 | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.TierVersionSetApiAPI.TierVersionSetApiReleaseTierVersionSet(context.Background(), serviceId, productTierId, version).ReleaseTierVersionSetRequest2(releaseTierVersionSetRequest2).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `TierVersionSetApiAPI.TierVersionSetApiReleaseTierVersionSet``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `TierVersionSetApiReleaseTierVersionSet`: TierVersionSet
	fmt.Fprintf(os.Stdout, "Response from `TierVersionSetApiAPI.TierVersionSetApiReleaseTierVersionSet`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**serviceId** | **string** | ID of the Service | 
**productTierId** | **string** | The product tier ID that this version set belongs to. | 
**version** | **string** | The version number for the specific version set. | 

### Other Parameters

Other parameters are passed through a pointer to a apiTierVersionSetApiReleaseTierVersionSetRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------



 **releaseTierVersionSetRequest2** | [**ReleaseTierVersionSetRequest2**](ReleaseTierVersionSetRequest2.md) |  | 

### Return type

[**TierVersionSet**](TierVersionSet.md)

### Authorization

[api_key_header_Authorization](../README.md#api_key_header_Authorization)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json, application/vnd.goa.error

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## TierVersionSetApiUpdateTierVersionSet

> TierVersionSet TierVersionSetApiUpdateTierVersionSet(ctx, serviceId, productTierId, version).UpdateTierVersionSetRequest2(updateTierVersionSetRequest2).Execute()

UpdateTierVersionSet tier-version-set-api

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
	serviceId := "s-12345678" // string | ID of the Service
	productTierId := "Beatae beatae." // string | The product tier ID that this version set belongs to.
	version := "3.0" // string | The version number for the specific version set.
	updateTierVersionSetRequest2 := *openapiclient.NewUpdateTierVersionSetRequest2("Updated version set name") // UpdateTierVersionSetRequest2 | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.TierVersionSetApiAPI.TierVersionSetApiUpdateTierVersionSet(context.Background(), serviceId, productTierId, version).UpdateTierVersionSetRequest2(updateTierVersionSetRequest2).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `TierVersionSetApiAPI.TierVersionSetApiUpdateTierVersionSet``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `TierVersionSetApiUpdateTierVersionSet`: TierVersionSet
	fmt.Fprintf(os.Stdout, "Response from `TierVersionSetApiAPI.TierVersionSetApiUpdateTierVersionSet`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**serviceId** | **string** | ID of the Service | 
**productTierId** | **string** | The product tier ID that this version set belongs to. | 
**version** | **string** | The version number for the specific version set. | 

### Other Parameters

Other parameters are passed through a pointer to a apiTierVersionSetApiUpdateTierVersionSetRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------



 **updateTierVersionSetRequest2** | [**UpdateTierVersionSetRequest2**](UpdateTierVersionSetRequest2.md) |  | 

### Return type

[**TierVersionSet**](TierVersionSet.md)

### Authorization

[api_key_header_Authorization](../README.md#api_key_header_Authorization)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json, application/vnd.goa.error

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## TierVersionSetApiUpdateTierVersionSetMetadata

> TierVersionSetMetadata TierVersionSetApiUpdateTierVersionSetMetadata(ctx, serviceId, productTierId, version).UpdateTierVersionSetMetadataRequest2(updateTierVersionSetMetadataRequest2).Execute()

UpdateTierVersionSetMetadata tier-version-set-api

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
	serviceId := "s-12345678" // string | ID of the Service
	productTierId := "Beatae beatae." // string | The product tier ID that this version set belongs to.
	version := "3.0" // string | The version number for the specific version set.
	updateTierVersionSetMetadataRequest2 := *openapiclient.NewUpdateTierVersionSetMetadataRequest2() // UpdateTierVersionSetMetadataRequest2 | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.TierVersionSetApiAPI.TierVersionSetApiUpdateTierVersionSetMetadata(context.Background(), serviceId, productTierId, version).UpdateTierVersionSetMetadataRequest2(updateTierVersionSetMetadataRequest2).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `TierVersionSetApiAPI.TierVersionSetApiUpdateTierVersionSetMetadata``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `TierVersionSetApiUpdateTierVersionSetMetadata`: TierVersionSetMetadata
	fmt.Fprintf(os.Stdout, "Response from `TierVersionSetApiAPI.TierVersionSetApiUpdateTierVersionSetMetadata`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**serviceId** | **string** | ID of the Service | 
**productTierId** | **string** | The product tier ID that this version set belongs to. | 
**version** | **string** | The version number for the specific version set. | 

### Other Parameters

Other parameters are passed through a pointer to a apiTierVersionSetApiUpdateTierVersionSetMetadataRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------



 **updateTierVersionSetMetadataRequest2** | [**UpdateTierVersionSetMetadataRequest2**](UpdateTierVersionSetMetadataRequest2.md) |  | 

### Return type

[**TierVersionSetMetadata**](TierVersionSetMetadata.md)

### Authorization

[api_key_header_Authorization](../README.md#api_key_header_Authorization)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json, application/vnd.goa.error

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

