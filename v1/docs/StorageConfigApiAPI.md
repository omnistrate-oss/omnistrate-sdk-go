# \StorageConfigApiAPI

All URIs are relative to *https://api.omnistrate.cloud*

Method | HTTP request | Description
------------- | ------------- | -------------
[**StorageConfigApiAddStorageVolumeConfig**](StorageConfigApiAPI.md#StorageConfigApiAddStorageVolumeConfig) | **Put** /2022-09-01-00/service/{serviceId}/storage-config/{id}/volume/{storageVolumeConfigId} | AddStorageVolumeConfig storage-config-api
[**StorageConfigApiCreateStorageConfig**](StorageConfigApiAPI.md#StorageConfigApiCreateStorageConfig) | **Post** /2022-09-01-00/service/{serviceId}/storage-config | CreateStorageConfig storage-config-api
[**StorageConfigApiDeleteStorageConfig**](StorageConfigApiAPI.md#StorageConfigApiDeleteStorageConfig) | **Delete** /2022-09-01-00/service/{serviceId}/storage-config/{id} | DeleteStorageConfig storage-config-api
[**StorageConfigApiDescribeStorageConfig**](StorageConfigApiAPI.md#StorageConfigApiDescribeStorageConfig) | **Get** /2022-09-01-00/service/{serviceId}/storage-config/{id} | DescribeStorageConfig storage-config-api
[**StorageConfigApiListStorageConfig**](StorageConfigApiAPI.md#StorageConfigApiListStorageConfig) | **Get** /2022-09-01-00/service/{serviceId}/storage-config | ListStorageConfig storage-config-api
[**StorageConfigApiRemoveStorageVolumeConfig**](StorageConfigApiAPI.md#StorageConfigApiRemoveStorageVolumeConfig) | **Delete** /2022-09-01-00/service/{serviceId}/storage-config/{id}/volume/{storageVolumeConfigId} | RemoveStorageVolumeConfig storage-config-api
[**StorageConfigApiUpdateStorageConfig**](StorageConfigApiAPI.md#StorageConfigApiUpdateStorageConfig) | **Patch** /2022-09-01-00/service/{serviceId}/storage-config/{id} | UpdateStorageConfig storage-config-api



## StorageConfigApiAddStorageVolumeConfig

> StorageConfigApiAddStorageVolumeConfig(ctx, serviceId, id, storageVolumeConfigId).AddStorageVolumeConfigRequest2(addStorageVolumeConfigRequest2).Execute()

AddStorageVolumeConfig storage-config-api

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
	id := "sc-12345678" // string | The storage config ID
	storageVolumeConfigId := "svc-12345678" // string | The storage volume config ID
	addStorageVolumeConfigRequest2 := *openapiclient.NewAddStorageVolumeConfigRequest2() // AddStorageVolumeConfigRequest2 | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	r, err := apiClient.StorageConfigApiAPI.StorageConfigApiAddStorageVolumeConfig(context.Background(), serviceId, id, storageVolumeConfigId).AddStorageVolumeConfigRequest2(addStorageVolumeConfigRequest2).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `StorageConfigApiAPI.StorageConfigApiAddStorageVolumeConfig``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**serviceId** | **string** | The service ID | 
**id** | **string** | The storage config ID | 
**storageVolumeConfigId** | **string** | The storage volume config ID | 

### Other Parameters

Other parameters are passed through a pointer to a apiStorageConfigApiAddStorageVolumeConfigRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------



 **addStorageVolumeConfigRequest2** | [**AddStorageVolumeConfigRequest2**](AddStorageVolumeConfigRequest2.md) |  | 

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


## StorageConfigApiCreateStorageConfig

> string StorageConfigApiCreateStorageConfig(ctx, serviceId).CreateStorageConfigRequest2(createStorageConfigRequest2).Execute()

CreateStorageConfig storage-config-api

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
	serviceId := "s-12345678" // string | The service to which this storage config belongs
	createStorageConfigRequest2 := *openapiclient.NewCreateStorageConfigRequest2("my-storage-config-description", "my-storage-config") // CreateStorageConfigRequest2 | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.StorageConfigApiAPI.StorageConfigApiCreateStorageConfig(context.Background(), serviceId).CreateStorageConfigRequest2(createStorageConfigRequest2).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `StorageConfigApiAPI.StorageConfigApiCreateStorageConfig``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `StorageConfigApiCreateStorageConfig`: string
	fmt.Fprintf(os.Stdout, "Response from `StorageConfigApiAPI.StorageConfigApiCreateStorageConfig`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**serviceId** | **string** | The service to which this storage config belongs | 

### Other Parameters

Other parameters are passed through a pointer to a apiStorageConfigApiCreateStorageConfigRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **createStorageConfigRequest2** | [**CreateStorageConfigRequest2**](CreateStorageConfigRequest2.md) |  | 

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


## StorageConfigApiDeleteStorageConfig

> StorageConfigApiDeleteStorageConfig(ctx, serviceId, id).Execute()

DeleteStorageConfig storage-config-api

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
	id := "sc-12345678" // string | The storage config ID

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	r, err := apiClient.StorageConfigApiAPI.StorageConfigApiDeleteStorageConfig(context.Background(), serviceId, id).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `StorageConfigApiAPI.StorageConfigApiDeleteStorageConfig``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**serviceId** | **string** | The service ID | 
**id** | **string** | The storage config ID | 

### Other Parameters

Other parameters are passed through a pointer to a apiStorageConfigApiDeleteStorageConfigRequest struct via the builder pattern


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


## StorageConfigApiDescribeStorageConfig

> DescribeStorageConfigResult StorageConfigApiDescribeStorageConfig(ctx, serviceId, id).ProductTierVersion(productTierVersion).ProductTierId(productTierId).Execute()

DescribeStorageConfig storage-config-api

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
	id := "sc-12345678" // string | The storage config ID
	productTierVersion := "Eum blanditiis animi." // string | Product tier version of the storage config to describe. If not specified, the latest version is described. (optional)
	productTierId := "Beatae beatae." // string | ProductTierId of the storage config to describe. Needs to specified in combination with the product tier version (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.StorageConfigApiAPI.StorageConfigApiDescribeStorageConfig(context.Background(), serviceId, id).ProductTierVersion(productTierVersion).ProductTierId(productTierId).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `StorageConfigApiAPI.StorageConfigApiDescribeStorageConfig``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `StorageConfigApiDescribeStorageConfig`: DescribeStorageConfigResult
	fmt.Fprintf(os.Stdout, "Response from `StorageConfigApiAPI.StorageConfigApiDescribeStorageConfig`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**serviceId** | **string** | The service ID | 
**id** | **string** | The storage config ID | 

### Other Parameters

Other parameters are passed through a pointer to a apiStorageConfigApiDescribeStorageConfigRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


 **productTierVersion** | **string** | Product tier version of the storage config to describe. If not specified, the latest version is described. | 
 **productTierId** | **string** | ProductTierId of the storage config to describe. Needs to specified in combination with the product tier version | 

### Return type

[**DescribeStorageConfigResult**](DescribeStorageConfigResult.md)

### Authorization

[api_key_header_Authorization](../README.md#api_key_header_Authorization)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json, application/vnd.goa.error

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## StorageConfigApiListStorageConfig

> ListStorageConfigsResult StorageConfigApiListStorageConfig(ctx, serviceId).Managed(managed).Execute()

ListStorageConfig storage-config-api

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
	serviceId := "s-12345678" // string | The service id to filter by
	managed := false // bool | Is storage config managed by omnistrate (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.StorageConfigApiAPI.StorageConfigApiListStorageConfig(context.Background(), serviceId).Managed(managed).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `StorageConfigApiAPI.StorageConfigApiListStorageConfig``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `StorageConfigApiListStorageConfig`: ListStorageConfigsResult
	fmt.Fprintf(os.Stdout, "Response from `StorageConfigApiAPI.StorageConfigApiListStorageConfig`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**serviceId** | **string** | The service id to filter by | 

### Other Parameters

Other parameters are passed through a pointer to a apiStorageConfigApiListStorageConfigRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **managed** | **bool** | Is storage config managed by omnistrate | 

### Return type

[**ListStorageConfigsResult**](ListStorageConfigsResult.md)

### Authorization

[api_key_header_Authorization](../README.md#api_key_header_Authorization)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json, application/vnd.goa.error

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## StorageConfigApiRemoveStorageVolumeConfig

> StorageConfigApiRemoveStorageVolumeConfig(ctx, serviceId, id, storageVolumeConfigId).RemoveStorageVolumeConfigRequest2(removeStorageVolumeConfigRequest2).Execute()

RemoveStorageVolumeConfig storage-config-api

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
	id := "sc-12345678" // string | The storage config ID
	storageVolumeConfigId := "svc-12345678" // string | The storage volume config ID
	removeStorageVolumeConfigRequest2 := *openapiclient.NewRemoveStorageVolumeConfigRequest2() // RemoveStorageVolumeConfigRequest2 | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	r, err := apiClient.StorageConfigApiAPI.StorageConfigApiRemoveStorageVolumeConfig(context.Background(), serviceId, id, storageVolumeConfigId).RemoveStorageVolumeConfigRequest2(removeStorageVolumeConfigRequest2).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `StorageConfigApiAPI.StorageConfigApiRemoveStorageVolumeConfig``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**serviceId** | **string** | The service ID | 
**id** | **string** | The storage config ID | 
**storageVolumeConfigId** | **string** | The storage volume config ID | 

### Other Parameters

Other parameters are passed through a pointer to a apiStorageConfigApiRemoveStorageVolumeConfigRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------



 **removeStorageVolumeConfigRequest2** | [**RemoveStorageVolumeConfigRequest2**](RemoveStorageVolumeConfigRequest2.md) |  | 

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


## StorageConfigApiUpdateStorageConfig

> StorageConfigApiUpdateStorageConfig(ctx, serviceId, id).UpdateStorageConfigRequest2(updateStorageConfigRequest2).Execute()

UpdateStorageConfig storage-config-api

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
	id := "sc-12345678" // string | The storage config ID
	updateStorageConfigRequest2 := *openapiclient.NewUpdateStorageConfigRequest2() // UpdateStorageConfigRequest2 | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	r, err := apiClient.StorageConfigApiAPI.StorageConfigApiUpdateStorageConfig(context.Background(), serviceId, id).UpdateStorageConfigRequest2(updateStorageConfigRequest2).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `StorageConfigApiAPI.StorageConfigApiUpdateStorageConfig``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**serviceId** | **string** | The service ID | 
**id** | **string** | The storage config ID | 

### Other Parameters

Other parameters are passed through a pointer to a apiStorageConfigApiUpdateStorageConfigRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


 **updateStorageConfigRequest2** | [**UpdateStorageConfigRequest2**](UpdateStorageConfigRequest2.md) |  | 

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

