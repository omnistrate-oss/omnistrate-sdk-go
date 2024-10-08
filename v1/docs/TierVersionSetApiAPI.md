# \TierVersionSetApiAPI

All URIs are relative to *https://api.omnistrate.cloud*

Method | HTTP request | Description
------------- | ------------- | -------------
[**TierVersionSetApiCreateTierVersionSet**](TierVersionSetApiAPI.md#TierVersionSetApiCreateTierVersionSet) | **Post** /2022-09-01-00/service/{serviceId}/productTier/{productTierId}/version-set | CreateTierVersionSet tier-version-set-api
[**TierVersionSetApiDeprecateTierVersionSet**](TierVersionSetApiAPI.md#TierVersionSetApiDeprecateTierVersionSet) | **Patch** /2022-09-01-00/service/{serviceId}/productTier/{productTierId}/version-set/{version}/deprecate | DeprecateTierVersionSet tier-version-set-api
[**TierVersionSetApiDescribeTierVersionSet**](TierVersionSetApiAPI.md#TierVersionSetApiDescribeTierVersionSet) | **Get** /2022-09-01-00/service/{serviceId}/productTier/{productTierId}/version-set/{version} | DescribeTierVersionSet tier-version-set-api
[**TierVersionSetApiDiffTierVersionSets**](TierVersionSetApiAPI.md#TierVersionSetApiDiffTierVersionSets) | **Get** /2022-09-01-00/service/{serviceId}/productTier/{productTierId}/version-set/{version}/diff/{anotherVersion} | DiffTierVersionSets tier-version-set-api
[**TierVersionSetApiListTierVersionSets**](TierVersionSetApiAPI.md#TierVersionSetApiListTierVersionSets) | **Get** /2022-09-01-00/service/{serviceId}/productTier/{productTierId}/version-set | ListTierVersionSets tier-version-set-api
[**TierVersionSetApiPromoteTierVersionSet**](TierVersionSetApiAPI.md#TierVersionSetApiPromoteTierVersionSet) | **Patch** /2022-09-01-00/service/{serviceId}/productTier/{productTierId}/version-set/{version}/promote | PromoteTierVersionSet tier-version-set-api
[**TierVersionSetApiReleaseTierVersionSet**](TierVersionSetApiAPI.md#TierVersionSetApiReleaseTierVersionSet) | **Patch** /2022-09-01-00/service/{serviceId}/productTier/{productTierId}/version-set/{version}/release | ReleaseTierVersionSet tier-version-set-api



## TierVersionSetApiCreateTierVersionSet

> CreateTierVersionSetResult TierVersionSetApiCreateTierVersionSet(ctx, serviceId, productTierId).CreateTierVersionSetRequestBody(createTierVersionSetRequestBody).Execute()

CreateTierVersionSet tier-version-set-api

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
	serviceId := "s-12345678" // string | ID of the Service
	productTierId := "Beatae beatae." // string | The product tier ID that this version set belongs to.
	createTierVersionSetRequestBody := *openapiclient.NewCreateTierVersionSetRequestBody("Major|Incremental|UserDefined") // CreateTierVersionSetRequestBody | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.TierVersionSetApiAPI.TierVersionSetApiCreateTierVersionSet(context.Background(), serviceId, productTierId).CreateTierVersionSetRequestBody(createTierVersionSetRequestBody).Execute()
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


 **createTierVersionSetRequestBody** | [**CreateTierVersionSetRequestBody**](CreateTierVersionSetRequestBody.md) |  | 

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
	openapiclient "github.com/omnistrate/omnistrate-sdk-go/v1"
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
	openapiclient "github.com/omnistrate/omnistrate-sdk-go/v1"
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
	openapiclient "github.com/omnistrate/omnistrate-sdk-go/v1"
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


## TierVersionSetApiListTierVersionSets

> ListTierVersionSetsResult TierVersionSetApiListTierVersionSets(ctx, serviceId, productTierId).LatestMajorVersionOnly(latestMajorVersionOnly).LatestIncrementalVersionForMajorVersion(latestIncrementalVersionForMajorVersion).Execute()

ListTierVersionSets tier-version-set-api

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
	serviceId := "s-12345678" // string | ID of the Service
	productTierId := "Beatae beatae." // string | The product tier ID that this version set belongs to.
	latestMajorVersionOnly := true // bool | If true, the latest major version is returned. The parameter needs to be specified in isolation. (optional)
	latestIncrementalVersionForMajorVersion := "3.0" // string | Returns the latest incremental version for the given major version. The paramenter needs to be specified in isolation. (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.TierVersionSetApiAPI.TierVersionSetApiListTierVersionSets(context.Background(), serviceId, productTierId).LatestMajorVersionOnly(latestMajorVersionOnly).LatestIncrementalVersionForMajorVersion(latestIncrementalVersionForMajorVersion).Execute()
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
	openapiclient "github.com/omnistrate/omnistrate-sdk-go/v1"
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

> TierVersionSet TierVersionSetApiReleaseTierVersionSet(ctx, serviceId, productTierId, version).ReleaseTierVersionSetRequestBody(releaseTierVersionSetRequestBody).Execute()

ReleaseTierVersionSet tier-version-set-api

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
	serviceId := "s-12345678" // string | ID of the Service
	productTierId := "Beatae beatae." // string | The product tier ID that this version set belongs to.
	version := "3.0" // string | The version number for the specific version set.
	releaseTierVersionSetRequestBody := *openapiclient.NewReleaseTierVersionSetRequestBody() // ReleaseTierVersionSetRequestBody | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.TierVersionSetApiAPI.TierVersionSetApiReleaseTierVersionSet(context.Background(), serviceId, productTierId, version).ReleaseTierVersionSetRequestBody(releaseTierVersionSetRequestBody).Execute()
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



 **releaseTierVersionSetRequestBody** | [**ReleaseTierVersionSetRequestBody**](ReleaseTierVersionSetRequestBody.md) |  | 

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

