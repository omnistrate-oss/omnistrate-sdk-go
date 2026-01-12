# \InfraConfigApiAPI

All URIs are relative to *https://api.omnistrate.cloud*

Method | HTTP request | Description
------------- | ------------- | -------------
[**InfraConfigApiCreateInfraConfig**](InfraConfigApiAPI.md#InfraConfigApiCreateInfraConfig) | **Post** /2022-09-01-00/service/{serviceId}/infra-config | CreateInfraConfig infra-config-api
[**InfraConfigApiDeleteInfraConfig**](InfraConfigApiAPI.md#InfraConfigApiDeleteInfraConfig) | **Delete** /2022-09-01-00/service/{serviceId}/infra-config/{id} | DeleteInfraConfig infra-config-api
[**InfraConfigApiDescribeInfraConfig**](InfraConfigApiAPI.md#InfraConfigApiDescribeInfraConfig) | **Get** /2022-09-01-00/service/{serviceId}/infra-config/{id} | DescribeInfraConfig infra-config-api
[**InfraConfigApiDetachComputeConfig**](InfraConfigApiAPI.md#InfraConfigApiDetachComputeConfig) | **Delete** /2022-09-01-00/service/{serviceId}/infra-config/{id}/compute-config | DetachComputeConfig infra-config-api
[**InfraConfigApiDetachNetworkConfig**](InfraConfigApiAPI.md#InfraConfigApiDetachNetworkConfig) | **Delete** /2022-09-01-00/service/{serviceId}/infra-config/{id}/network-config | DetachNetworkConfig infra-config-api
[**InfraConfigApiDetachStorageConfig**](InfraConfigApiAPI.md#InfraConfigApiDetachStorageConfig) | **Delete** /2022-09-01-00/service/{serviceId}/infra-config/{id}/storage-config | DetachStorageConfig infra-config-api
[**InfraConfigApiListAssociatedResources**](InfraConfigApiAPI.md#InfraConfigApiListAssociatedResources) | **Get** /2022-09-01-00/service/{serviceId}/infra-config/{id}/associated-resources | ListAssociatedResources infra-config-api
[**InfraConfigApiListInfraConfig**](InfraConfigApiAPI.md#InfraConfigApiListInfraConfig) | **Get** /2022-09-01-00/service/{serviceId}/serviceenvironment/{serviceEnvironmentId}/infra-config | ListInfraConfig infra-config-api
[**InfraConfigApiReleaseInfraConfig**](InfraConfigApiAPI.md#InfraConfigApiReleaseInfraConfig) | **Post** /2022-09-01-00/service/{serviceId}/infra-config/{id}/release | ReleaseInfraConfig infra-config-api
[**InfraConfigApiRolloutFleetInfra**](InfraConfigApiAPI.md#InfraConfigApiRolloutFleetInfra) | **Post** /2022-09-01-00/service/{serviceId}/infra-config/{id}/rollout | RolloutFleetInfra infra-config-api
[**InfraConfigApiRolloutFleetInfraStatus**](InfraConfigApiAPI.md#InfraConfigApiRolloutFleetInfraStatus) | **Get** /2022-09-01-00/service/{serviceId}/infra-config/{id}/rollout/status | RolloutFleetInfraStatus infra-config-api
[**InfraConfigApiUpdateInfraConfig**](InfraConfigApiAPI.md#InfraConfigApiUpdateInfraConfig) | **Patch** /2022-09-01-00/service/{serviceId}/infra-config/{id} | UpdateInfraConfig infra-config-api



## InfraConfigApiCreateInfraConfig

> string InfraConfigApiCreateInfraConfig(ctx, serviceId).CreateInfraConfigRequest2(createInfraConfigRequest2).Execute()

CreateInfraConfig infra-config-api

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
	serviceId := "s-12345678" // string | The service ID
	createInfraConfigRequest2 := *openapiclient.NewCreateInfraConfigRequest2("Infra config used for the base tier MySQL service", "MySQL Writer Infra Config", "se-123456") // CreateInfraConfigRequest2 | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.InfraConfigApiAPI.InfraConfigApiCreateInfraConfig(context.Background(), serviceId).CreateInfraConfigRequest2(createInfraConfigRequest2).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `InfraConfigApiAPI.InfraConfigApiCreateInfraConfig``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `InfraConfigApiCreateInfraConfig`: string
	fmt.Fprintf(os.Stdout, "Response from `InfraConfigApiAPI.InfraConfigApiCreateInfraConfig`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**serviceId** | **string** | The service ID | 

### Other Parameters

Other parameters are passed through a pointer to a apiInfraConfigApiCreateInfraConfigRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **createInfraConfigRequest2** | [**CreateInfraConfigRequest2**](CreateInfraConfigRequest2.md) |  | 

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


## InfraConfigApiDeleteInfraConfig

> InfraConfigApiDeleteInfraConfig(ctx, serviceId, id).Execute()

DeleteInfraConfig infra-config-api

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
	serviceId := "s-12345678" // string | The service ID
	id := "ic-12345678" // string | Infra Config ID to operate on

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	r, err := apiClient.InfraConfigApiAPI.InfraConfigApiDeleteInfraConfig(context.Background(), serviceId, id).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `InfraConfigApiAPI.InfraConfigApiDeleteInfraConfig``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**serviceId** | **string** | The service ID | 
**id** | **string** | Infra Config ID to operate on | 

### Other Parameters

Other parameters are passed through a pointer to a apiInfraConfigApiDeleteInfraConfigRequest struct via the builder pattern


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


## InfraConfigApiDescribeInfraConfig

> DescribeInfraConfigResult InfraConfigApiDescribeInfraConfig(ctx, serviceId, id).ProductTierVersion(productTierVersion).ProductTierId(productTierId).Execute()

DescribeInfraConfig infra-config-api

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
	serviceId := "s-12345678" // string | The service ID
	id := "ic-12345678" // string | Infra Config ID to operate on
	productTierVersion := "Voluptatem et corrupti dolores accusantium tempora." // string | Product tier version of the infra config to describe. If not specified, the latest version is described. (optional)
	productTierId := "Beatae beatae." // string | ProductTierId of the infra config to describe. Needs to specified in combination with the product tier version (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.InfraConfigApiAPI.InfraConfigApiDescribeInfraConfig(context.Background(), serviceId, id).ProductTierVersion(productTierVersion).ProductTierId(productTierId).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `InfraConfigApiAPI.InfraConfigApiDescribeInfraConfig``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `InfraConfigApiDescribeInfraConfig`: DescribeInfraConfigResult
	fmt.Fprintf(os.Stdout, "Response from `InfraConfigApiAPI.InfraConfigApiDescribeInfraConfig`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**serviceId** | **string** | The service ID | 
**id** | **string** | Infra Config ID to operate on | 

### Other Parameters

Other parameters are passed through a pointer to a apiInfraConfigApiDescribeInfraConfigRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


 **productTierVersion** | **string** | Product tier version of the infra config to describe. If not specified, the latest version is described. | 
 **productTierId** | **string** | ProductTierId of the infra config to describe. Needs to specified in combination with the product tier version | 

### Return type

[**DescribeInfraConfigResult**](DescribeInfraConfigResult.md)

### Authorization

[api_key_header_Authorization](../README.md#api_key_header_Authorization)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json, application/vnd.goa.error

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## InfraConfigApiDetachComputeConfig

> InfraConfigApiDetachComputeConfig(ctx, serviceId, id).Execute()

DetachComputeConfig infra-config-api

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
	serviceId := "s-12345678" // string | The service ID
	id := "ic-12345678" // string | Infra Config ID to operate on

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	r, err := apiClient.InfraConfigApiAPI.InfraConfigApiDetachComputeConfig(context.Background(), serviceId, id).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `InfraConfigApiAPI.InfraConfigApiDetachComputeConfig``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**serviceId** | **string** | The service ID | 
**id** | **string** | Infra Config ID to operate on | 

### Other Parameters

Other parameters are passed through a pointer to a apiInfraConfigApiDetachComputeConfigRequest struct via the builder pattern


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


## InfraConfigApiDetachNetworkConfig

> InfraConfigApiDetachNetworkConfig(ctx, serviceId, id).Execute()

DetachNetworkConfig infra-config-api

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
	serviceId := "s-12345678" // string | The service ID
	id := "ic-12345678" // string | Infra Config ID to operate on

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	r, err := apiClient.InfraConfigApiAPI.InfraConfigApiDetachNetworkConfig(context.Background(), serviceId, id).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `InfraConfigApiAPI.InfraConfigApiDetachNetworkConfig``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**serviceId** | **string** | The service ID | 
**id** | **string** | Infra Config ID to operate on | 

### Other Parameters

Other parameters are passed through a pointer to a apiInfraConfigApiDetachNetworkConfigRequest struct via the builder pattern


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


## InfraConfigApiDetachStorageConfig

> InfraConfigApiDetachStorageConfig(ctx, serviceId, id).Execute()

DetachStorageConfig infra-config-api

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
	serviceId := "s-12345678" // string | The service ID
	id := "ic-12345678" // string | Infra Config ID to operate on

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	r, err := apiClient.InfraConfigApiAPI.InfraConfigApiDetachStorageConfig(context.Background(), serviceId, id).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `InfraConfigApiAPI.InfraConfigApiDetachStorageConfig``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**serviceId** | **string** | The service ID | 
**id** | **string** | Infra Config ID to operate on | 

### Other Parameters

Other parameters are passed through a pointer to a apiInfraConfigApiDetachStorageConfigRequest struct via the builder pattern


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


## InfraConfigApiListAssociatedResources

> ListAssociatedResourcesResult InfraConfigApiListAssociatedResources(ctx, serviceId, id).Execute()

ListAssociatedResources infra-config-api

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
	serviceId := "s-12345678" // string | The service ID
	id := "ic-12345678" // string | Infra Config ID to operate on

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.InfraConfigApiAPI.InfraConfigApiListAssociatedResources(context.Background(), serviceId, id).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `InfraConfigApiAPI.InfraConfigApiListAssociatedResources``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `InfraConfigApiListAssociatedResources`: ListAssociatedResourcesResult
	fmt.Fprintf(os.Stdout, "Response from `InfraConfigApiAPI.InfraConfigApiListAssociatedResources`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**serviceId** | **string** | The service ID | 
**id** | **string** | Infra Config ID to operate on | 

### Other Parameters

Other parameters are passed through a pointer to a apiInfraConfigApiListAssociatedResourcesRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------



### Return type

[**ListAssociatedResourcesResult**](ListAssociatedResourcesResult.md)

### Authorization

[api_key_header_Authorization](../README.md#api_key_header_Authorization)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json, application/vnd.goa.error

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## InfraConfigApiListInfraConfig

> ListInfraConfigResult InfraConfigApiListInfraConfig(ctx, serviceId, serviceEnvironmentId).Managed(managed).Execute()

ListInfraConfig infra-config-api

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
	serviceId := "s-12345678" // string | The service ID
	serviceEnvironmentId := "se-12345678" // string | The service environment ID
	managed := false // bool | Is infra config managed by omnistrate (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.InfraConfigApiAPI.InfraConfigApiListInfraConfig(context.Background(), serviceId, serviceEnvironmentId).Managed(managed).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `InfraConfigApiAPI.InfraConfigApiListInfraConfig``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `InfraConfigApiListInfraConfig`: ListInfraConfigResult
	fmt.Fprintf(os.Stdout, "Response from `InfraConfigApiAPI.InfraConfigApiListInfraConfig`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**serviceId** | **string** | The service ID | 
**serviceEnvironmentId** | **string** | The service environment ID | 

### Other Parameters

Other parameters are passed through a pointer to a apiInfraConfigApiListInfraConfigRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


 **managed** | **bool** | Is infra config managed by omnistrate | 

### Return type

[**ListInfraConfigResult**](ListInfraConfigResult.md)

### Authorization

[api_key_header_Authorization](../README.md#api_key_header_Authorization)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json, application/vnd.goa.error

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## InfraConfigApiReleaseInfraConfig

> InfraConfigApiReleaseInfraConfig(ctx, serviceId, id).ReleaseInfraConfigRequest2(releaseInfraConfigRequest2).Execute()

ReleaseInfraConfig infra-config-api

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
	serviceId := "s-12345678" // string | The service ID
	id := "ic-12345678" // string | Infra Config ID to operate on
	releaseInfraConfigRequest2 := *openapiclient.NewReleaseInfraConfigRequest2() // ReleaseInfraConfigRequest2 | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	r, err := apiClient.InfraConfigApiAPI.InfraConfigApiReleaseInfraConfig(context.Background(), serviceId, id).ReleaseInfraConfigRequest2(releaseInfraConfigRequest2).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `InfraConfigApiAPI.InfraConfigApiReleaseInfraConfig``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**serviceId** | **string** | The service ID | 
**id** | **string** | Infra Config ID to operate on | 

### Other Parameters

Other parameters are passed through a pointer to a apiInfraConfigApiReleaseInfraConfigRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


 **releaseInfraConfigRequest2** | [**ReleaseInfraConfigRequest2**](ReleaseInfraConfigRequest2.md) |  | 

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


## InfraConfigApiRolloutFleetInfra

> InfraConfigApiRolloutFleetInfra(ctx, serviceId, id).Execute()

RolloutFleetInfra infra-config-api

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
	serviceId := "s-12345678" // string | The service ID
	id := "ic-12345678" // string | Infra Config ID to operate on

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	r, err := apiClient.InfraConfigApiAPI.InfraConfigApiRolloutFleetInfra(context.Background(), serviceId, id).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `InfraConfigApiAPI.InfraConfigApiRolloutFleetInfra``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**serviceId** | **string** | The service ID | 
**id** | **string** | Infra Config ID to operate on | 

### Other Parameters

Other parameters are passed through a pointer to a apiInfraConfigApiRolloutFleetInfraRequest struct via the builder pattern


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


## InfraConfigApiRolloutFleetInfraStatus

> RolloutFleetInfraStatusResult InfraConfigApiRolloutFleetInfraStatus(ctx, serviceId, id).Execute()

RolloutFleetInfraStatus infra-config-api

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
	serviceId := "s-12345678" // string | The service ID
	id := "ic-12345678" // string | Infra Config ID to operate on

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.InfraConfigApiAPI.InfraConfigApiRolloutFleetInfraStatus(context.Background(), serviceId, id).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `InfraConfigApiAPI.InfraConfigApiRolloutFleetInfraStatus``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `InfraConfigApiRolloutFleetInfraStatus`: RolloutFleetInfraStatusResult
	fmt.Fprintf(os.Stdout, "Response from `InfraConfigApiAPI.InfraConfigApiRolloutFleetInfraStatus`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**serviceId** | **string** | The service ID | 
**id** | **string** | Infra Config ID to operate on | 

### Other Parameters

Other parameters are passed through a pointer to a apiInfraConfigApiRolloutFleetInfraStatusRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------



### Return type

[**RolloutFleetInfraStatusResult**](RolloutFleetInfraStatusResult.md)

### Authorization

[api_key_header_Authorization](../README.md#api_key_header_Authorization)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json, application/vnd.goa.error

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## InfraConfigApiUpdateInfraConfig

> InfraConfigApiUpdateInfraConfig(ctx, serviceId, id).UpdateInfraConfigRequest2(updateInfraConfigRequest2).Execute()

UpdateInfraConfig infra-config-api

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
	serviceId := "s-12345678" // string | The service ID
	id := "ic-12345678" // string | Infra Config ID to operate on
	updateInfraConfigRequest2 := *openapiclient.NewUpdateInfraConfigRequest2() // UpdateInfraConfigRequest2 | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	r, err := apiClient.InfraConfigApiAPI.InfraConfigApiUpdateInfraConfig(context.Background(), serviceId, id).UpdateInfraConfigRequest2(updateInfraConfigRequest2).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `InfraConfigApiAPI.InfraConfigApiUpdateInfraConfig``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**serviceId** | **string** | The service ID | 
**id** | **string** | Infra Config ID to operate on | 

### Other Parameters

Other parameters are passed through a pointer to a apiInfraConfigApiUpdateInfraConfigRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


 **updateInfraConfigRequest2** | [**UpdateInfraConfigRequest2**](UpdateInfraConfigRequest2.md) |  | 

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

