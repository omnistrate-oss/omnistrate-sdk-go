# \StorageVolumeConfigApiAPI

All URIs are relative to *https://api.omnistrate.cloud*

Method | HTTP request | Description
------------- | ------------- | -------------
[**StorageVolumeConfigApiCreateStorageVolumeConfig**](StorageVolumeConfigApiAPI.md#StorageVolumeConfigApiCreateStorageVolumeConfig) | **Post** /2022-09-01-00/service/{serviceId}/storage-volume-config | CreateStorageVolumeConfig storage-volume-config-api
[**StorageVolumeConfigApiDeleteStorageVolumeConfig**](StorageVolumeConfigApiAPI.md#StorageVolumeConfigApiDeleteStorageVolumeConfig) | **Delete** /2022-09-01-00/service/{serviceId}/storage-volume-config/{id} | DeleteStorageVolumeConfig storage-volume-config-api
[**StorageVolumeConfigApiDescribeStorageVolumeConfig**](StorageVolumeConfigApiAPI.md#StorageVolumeConfigApiDescribeStorageVolumeConfig) | **Get** /2022-09-01-00/service/{serviceId}/storage-volume-config/{id} | DescribeStorageVolumeConfig storage-volume-config-api
[**StorageVolumeConfigApiListStorageVolumeConfig**](StorageVolumeConfigApiAPI.md#StorageVolumeConfigApiListStorageVolumeConfig) | **Get** /2022-09-01-00/service/{serviceId}/storage-volume-config | ListStorageVolumeConfig storage-volume-config-api
[**StorageVolumeConfigApiUpdateInstanceStorageVolumeConfig**](StorageVolumeConfigApiAPI.md#StorageVolumeConfigApiUpdateInstanceStorageVolumeConfig) | **Patch** /2022-09-01-00/service/{serviceId}/storage-volume-config/{id}/instance | UpdateInstanceStorageVolumeConfig storage-volume-config-api
[**StorageVolumeConfigApiUpdateStorageVolumeConfig**](StorageVolumeConfigApiAPI.md#StorageVolumeConfigApiUpdateStorageVolumeConfig) | **Patch** /2022-09-01-00/service/{serviceId}/storage-volume-config/{id} | UpdateStorageVolumeConfig storage-volume-config-api
[**StorageVolumeConfigApiUpdateStorageVolumeSizeConfig**](StorageVolumeConfigApiAPI.md#StorageVolumeConfigApiUpdateStorageVolumeSizeConfig) | **Patch** /2022-09-01-00/service/{serviceId}/storage-volume-config/{id}/size | UpdateStorageVolumeSizeConfig storage-volume-config-api



## StorageVolumeConfigApiCreateStorageVolumeConfig

> string StorageVolumeConfigApiCreateStorageVolumeConfig(ctx, serviceId).CreateStorageVolumeConfigRequest2(createStorageVolumeConfigRequest2).Execute()

CreateStorageVolumeConfig storage-volume-config-api

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
	createStorageVolumeConfigRequest2 := *openapiclient.NewCreateStorageVolumeConfigRequest2("A storage volume set to store the MySQL data directory", "MySQL Data Volume") // CreateStorageVolumeConfigRequest2 | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.StorageVolumeConfigApiAPI.StorageVolumeConfigApiCreateStorageVolumeConfig(context.Background(), serviceId).CreateStorageVolumeConfigRequest2(createStorageVolumeConfigRequest2).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `StorageVolumeConfigApiAPI.StorageVolumeConfigApiCreateStorageVolumeConfig``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `StorageVolumeConfigApiCreateStorageVolumeConfig`: string
	fmt.Fprintf(os.Stdout, "Response from `StorageVolumeConfigApiAPI.StorageVolumeConfigApiCreateStorageVolumeConfig`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**serviceId** | **string** | The service ID | 

### Other Parameters

Other parameters are passed through a pointer to a apiStorageVolumeConfigApiCreateStorageVolumeConfigRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **createStorageVolumeConfigRequest2** | [**CreateStorageVolumeConfigRequest2**](CreateStorageVolumeConfigRequest2.md) |  | 

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


## StorageVolumeConfigApiDeleteStorageVolumeConfig

> StorageVolumeConfigApiDeleteStorageVolumeConfig(ctx, serviceId, id).Execute()

DeleteStorageVolumeConfig storage-volume-config-api

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
	id := "svc-12345678" // string | The storage volume config ID

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	r, err := apiClient.StorageVolumeConfigApiAPI.StorageVolumeConfigApiDeleteStorageVolumeConfig(context.Background(), serviceId, id).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `StorageVolumeConfigApiAPI.StorageVolumeConfigApiDeleteStorageVolumeConfig``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**serviceId** | **string** | The service ID | 
**id** | **string** | The storage volume config ID | 

### Other Parameters

Other parameters are passed through a pointer to a apiStorageVolumeConfigApiDeleteStorageVolumeConfigRequest struct via the builder pattern


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


## StorageVolumeConfigApiDescribeStorageVolumeConfig

> DescribeStorageVolumeConfigResult StorageVolumeConfigApiDescribeStorageVolumeConfig(ctx, serviceId, id).ProductTierVersion(productTierVersion).ProductTierId(productTierId).Execute()

DescribeStorageVolumeConfig storage-volume-config-api

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
	id := "svc-12345678" // string | The storage volume config ID
	productTierVersion := "Facilis rerum necessitatibus voluptas." // string | Product tier version of the network config to describe. If not specified, the latest version is described. (optional)
	productTierId := "Beatae beatae." // string | ProductTierId of the network config to describe. Needs to specified in combination with the product tier version (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.StorageVolumeConfigApiAPI.StorageVolumeConfigApiDescribeStorageVolumeConfig(context.Background(), serviceId, id).ProductTierVersion(productTierVersion).ProductTierId(productTierId).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `StorageVolumeConfigApiAPI.StorageVolumeConfigApiDescribeStorageVolumeConfig``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `StorageVolumeConfigApiDescribeStorageVolumeConfig`: DescribeStorageVolumeConfigResult
	fmt.Fprintf(os.Stdout, "Response from `StorageVolumeConfigApiAPI.StorageVolumeConfigApiDescribeStorageVolumeConfig`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**serviceId** | **string** | The service ID | 
**id** | **string** | The storage volume config ID | 

### Other Parameters

Other parameters are passed through a pointer to a apiStorageVolumeConfigApiDescribeStorageVolumeConfigRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


 **productTierVersion** | **string** | Product tier version of the network config to describe. If not specified, the latest version is described. | 
 **productTierId** | **string** | ProductTierId of the network config to describe. Needs to specified in combination with the product tier version | 

### Return type

[**DescribeStorageVolumeConfigResult**](DescribeStorageVolumeConfigResult.md)

### Authorization

[api_key_header_Authorization](../README.md#api_key_header_Authorization)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json, application/vnd.goa.error

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## StorageVolumeConfigApiListStorageVolumeConfig

> ListStorageVolumeConfigsResult StorageVolumeConfigApiListStorageVolumeConfig(ctx, serviceId).Managed(managed).Execute()

ListStorageVolumeConfig storage-volume-config-api

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
	serviceId := "s-12345678" // string | The service to list storage volume configs for
	managed := false // bool | Is storage volume config managed by omnistrate (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.StorageVolumeConfigApiAPI.StorageVolumeConfigApiListStorageVolumeConfig(context.Background(), serviceId).Managed(managed).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `StorageVolumeConfigApiAPI.StorageVolumeConfigApiListStorageVolumeConfig``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `StorageVolumeConfigApiListStorageVolumeConfig`: ListStorageVolumeConfigsResult
	fmt.Fprintf(os.Stdout, "Response from `StorageVolumeConfigApiAPI.StorageVolumeConfigApiListStorageVolumeConfig`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**serviceId** | **string** | The service to list storage volume configs for | 

### Other Parameters

Other parameters are passed through a pointer to a apiStorageVolumeConfigApiListStorageVolumeConfigRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **managed** | **bool** | Is storage volume config managed by omnistrate | 

### Return type

[**ListStorageVolumeConfigsResult**](ListStorageVolumeConfigsResult.md)

### Authorization

[api_key_header_Authorization](../README.md#api_key_header_Authorization)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json, application/vnd.goa.error

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## StorageVolumeConfigApiUpdateInstanceStorageVolumeConfig

> StorageVolumeConfigApiUpdateInstanceStorageVolumeConfig(ctx, serviceId, id).UpdateInstanceStorageVolumeConfigRequest2(updateInstanceStorageVolumeConfigRequest2).Execute()

UpdateInstanceStorageVolumeConfig storage-volume-config-api

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
	id := "svc-12345678" // string | The storage volume config ID
	updateInstanceStorageVolumeConfigRequest2 := *openapiclient.NewUpdateInstanceStorageVolumeConfigRequest2() // UpdateInstanceStorageVolumeConfigRequest2 | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	r, err := apiClient.StorageVolumeConfigApiAPI.StorageVolumeConfigApiUpdateInstanceStorageVolumeConfig(context.Background(), serviceId, id).UpdateInstanceStorageVolumeConfigRequest2(updateInstanceStorageVolumeConfigRequest2).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `StorageVolumeConfigApiAPI.StorageVolumeConfigApiUpdateInstanceStorageVolumeConfig``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**serviceId** | **string** | The service ID | 
**id** | **string** | The storage volume config ID | 

### Other Parameters

Other parameters are passed through a pointer to a apiStorageVolumeConfigApiUpdateInstanceStorageVolumeConfigRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


 **updateInstanceStorageVolumeConfigRequest2** | [**UpdateInstanceStorageVolumeConfigRequest2**](UpdateInstanceStorageVolumeConfigRequest2.md) |  | 

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


## StorageVolumeConfigApiUpdateStorageVolumeConfig

> StorageVolumeConfigApiUpdateStorageVolumeConfig(ctx, serviceId, id).UpdateStorageVolumeConfigRequest2(updateStorageVolumeConfigRequest2).Execute()

UpdateStorageVolumeConfig storage-volume-config-api

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
	id := "svc-12345678" // string | The storage volume config ID
	updateStorageVolumeConfigRequest2 := *openapiclient.NewUpdateStorageVolumeConfigRequest2() // UpdateStorageVolumeConfigRequest2 | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	r, err := apiClient.StorageVolumeConfigApiAPI.StorageVolumeConfigApiUpdateStorageVolumeConfig(context.Background(), serviceId, id).UpdateStorageVolumeConfigRequest2(updateStorageVolumeConfigRequest2).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `StorageVolumeConfigApiAPI.StorageVolumeConfigApiUpdateStorageVolumeConfig``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**serviceId** | **string** | The service ID | 
**id** | **string** | The storage volume config ID | 

### Other Parameters

Other parameters are passed through a pointer to a apiStorageVolumeConfigApiUpdateStorageVolumeConfigRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


 **updateStorageVolumeConfigRequest2** | [**UpdateStorageVolumeConfigRequest2**](UpdateStorageVolumeConfigRequest2.md) |  | 

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


## StorageVolumeConfigApiUpdateStorageVolumeSizeConfig

> StorageVolumeConfigApiUpdateStorageVolumeSizeConfig(ctx, serviceId, id).UpdateStorageVolumeSizeConfigRequest2(updateStorageVolumeSizeConfigRequest2).Execute()

UpdateStorageVolumeSizeConfig storage-volume-config-api

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
	id := "svc-12345678" // string | The storage volume config ID
	updateStorageVolumeSizeConfigRequest2 := *openapiclient.NewUpdateStorageVolumeSizeConfigRequest2("$var.storage_size") // UpdateStorageVolumeSizeConfigRequest2 | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	r, err := apiClient.StorageVolumeConfigApiAPI.StorageVolumeConfigApiUpdateStorageVolumeSizeConfig(context.Background(), serviceId, id).UpdateStorageVolumeSizeConfigRequest2(updateStorageVolumeSizeConfigRequest2).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `StorageVolumeConfigApiAPI.StorageVolumeConfigApiUpdateStorageVolumeSizeConfig``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**serviceId** | **string** | The service ID | 
**id** | **string** | The storage volume config ID | 

### Other Parameters

Other parameters are passed through a pointer to a apiStorageVolumeConfigApiUpdateStorageVolumeSizeConfigRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


 **updateStorageVolumeSizeConfigRequest2** | [**UpdateStorageVolumeSizeConfigRequest2**](UpdateStorageVolumeSizeConfigRequest2.md) |  | 

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

