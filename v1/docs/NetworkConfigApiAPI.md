# \NetworkConfigApiAPI

All URIs are relative to *https://api.omnistrate.cloud*

Method | HTTP request | Description
------------- | ------------- | -------------
[**NetworkConfigApiCreateNetworkConfig**](NetworkConfigApiAPI.md#NetworkConfigApiCreateNetworkConfig) | **Post** /2022-09-01-00/service/{serviceId}/network-config | CreateNetworkConfig network-config-api
[**NetworkConfigApiDeleteNetworkConfig**](NetworkConfigApiAPI.md#NetworkConfigApiDeleteNetworkConfig) | **Delete** /2022-09-01-00/service/{serviceId}/network-config/{id} | DeleteNetworkConfig network-config-api
[**NetworkConfigApiDescribeNetworkConfig**](NetworkConfigApiAPI.md#NetworkConfigApiDescribeNetworkConfig) | **Get** /2022-09-01-00/service/{serviceId}/network-config/{id} | DescribeNetworkConfig network-config-api
[**NetworkConfigApiListNetworkConfig**](NetworkConfigApiAPI.md#NetworkConfigApiListNetworkConfig) | **Get** /2022-09-01-00/service/{serviceId}/network-config | ListNetworkConfig network-config-api
[**NetworkConfigApiUpdateNetworkConfig**](NetworkConfigApiAPI.md#NetworkConfigApiUpdateNetworkConfig) | **Patch** /2022-09-01-00/service/{serviceId}/network-config/{id} | UpdateNetworkConfig network-config-api



## NetworkConfigApiCreateNetworkConfig

> string NetworkConfigApiCreateNetworkConfig(ctx, serviceId).CreateNetworkConfigRequest2(createNetworkConfigRequest2).Execute()

CreateNetworkConfig network-config-api

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
	serviceId := "s-12345678" // string | The ID of the service
	createNetworkConfigRequest2 := *openapiclient.NewCreateNetworkConfigRequest2("A multi-zone HA network config", true, "Multi-zone") // CreateNetworkConfigRequest2 | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.NetworkConfigApiAPI.NetworkConfigApiCreateNetworkConfig(context.Background(), serviceId).CreateNetworkConfigRequest2(createNetworkConfigRequest2).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `NetworkConfigApiAPI.NetworkConfigApiCreateNetworkConfig``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `NetworkConfigApiCreateNetworkConfig`: string
	fmt.Fprintf(os.Stdout, "Response from `NetworkConfigApiAPI.NetworkConfigApiCreateNetworkConfig`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**serviceId** | **string** | The ID of the service | 

### Other Parameters

Other parameters are passed through a pointer to a apiNetworkConfigApiCreateNetworkConfigRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **createNetworkConfigRequest2** | [**CreateNetworkConfigRequest2**](CreateNetworkConfigRequest2.md) |  | 

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


## NetworkConfigApiDeleteNetworkConfig

> NetworkConfigApiDeleteNetworkConfig(ctx, serviceId, id).Execute()

DeleteNetworkConfig network-config-api

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
	serviceId := "s-12345678" // string | The ID of the service
	id := "nc-12345678" // string | The ID of the network config

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	r, err := apiClient.NetworkConfigApiAPI.NetworkConfigApiDeleteNetworkConfig(context.Background(), serviceId, id).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `NetworkConfigApiAPI.NetworkConfigApiDeleteNetworkConfig``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**serviceId** | **string** | The ID of the service | 
**id** | **string** | The ID of the network config | 

### Other Parameters

Other parameters are passed through a pointer to a apiNetworkConfigApiDeleteNetworkConfigRequest struct via the builder pattern


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


## NetworkConfigApiDescribeNetworkConfig

> DescribeNetworkConfigResult NetworkConfigApiDescribeNetworkConfig(ctx, serviceId, id).ProductTierVersion(productTierVersion).ProductTierId(productTierId).Execute()

DescribeNetworkConfig network-config-api

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
	serviceId := "s-12345678" // string | The ID of the service
	id := "nc-12345678" // string | The ID of the network config
	productTierVersion := "Est facere id odit quidem nam." // string | Product tier version of the network config to describe. If not specified, the latest version is described. (optional)
	productTierId := "Beatae beatae." // string | ProductTierId of the network config to describe. Needs to specified in combination with the product tier version (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.NetworkConfigApiAPI.NetworkConfigApiDescribeNetworkConfig(context.Background(), serviceId, id).ProductTierVersion(productTierVersion).ProductTierId(productTierId).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `NetworkConfigApiAPI.NetworkConfigApiDescribeNetworkConfig``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `NetworkConfigApiDescribeNetworkConfig`: DescribeNetworkConfigResult
	fmt.Fprintf(os.Stdout, "Response from `NetworkConfigApiAPI.NetworkConfigApiDescribeNetworkConfig`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**serviceId** | **string** | The ID of the service | 
**id** | **string** | The ID of the network config | 

### Other Parameters

Other parameters are passed through a pointer to a apiNetworkConfigApiDescribeNetworkConfigRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


 **productTierVersion** | **string** | Product tier version of the network config to describe. If not specified, the latest version is described. | 
 **productTierId** | **string** | ProductTierId of the network config to describe. Needs to specified in combination with the product tier version | 

### Return type

[**DescribeNetworkConfigResult**](DescribeNetworkConfigResult.md)

### Authorization

[api_key_header_Authorization](../README.md#api_key_header_Authorization)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json, application/vnd.goa.error

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## NetworkConfigApiListNetworkConfig

> ListNetworkConfigsResult NetworkConfigApiListNetworkConfig(ctx, serviceId).Managed(managed).Execute()

ListNetworkConfig network-config-api

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
	serviceId := "s-12345678" // string | The service ID to list network configs for
	managed := false // bool | Is network config managed by omnistrate (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.NetworkConfigApiAPI.NetworkConfigApiListNetworkConfig(context.Background(), serviceId).Managed(managed).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `NetworkConfigApiAPI.NetworkConfigApiListNetworkConfig``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `NetworkConfigApiListNetworkConfig`: ListNetworkConfigsResult
	fmt.Fprintf(os.Stdout, "Response from `NetworkConfigApiAPI.NetworkConfigApiListNetworkConfig`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**serviceId** | **string** | The service ID to list network configs for | 

### Other Parameters

Other parameters are passed through a pointer to a apiNetworkConfigApiListNetworkConfigRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **managed** | **bool** | Is network config managed by omnistrate | 

### Return type

[**ListNetworkConfigsResult**](ListNetworkConfigsResult.md)

### Authorization

[api_key_header_Authorization](../README.md#api_key_header_Authorization)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json, application/vnd.goa.error

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## NetworkConfigApiUpdateNetworkConfig

> NetworkConfigApiUpdateNetworkConfig(ctx, serviceId, id).UpdateNetworkConfigRequest2(updateNetworkConfigRequest2).Execute()

UpdateNetworkConfig network-config-api

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
	serviceId := "s-12345678" // string | The ID of the service
	id := "nc-12345678" // string | The ID of the network config
	updateNetworkConfigRequest2 := *openapiclient.NewUpdateNetworkConfigRequest2() // UpdateNetworkConfigRequest2 | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	r, err := apiClient.NetworkConfigApiAPI.NetworkConfigApiUpdateNetworkConfig(context.Background(), serviceId, id).UpdateNetworkConfigRequest2(updateNetworkConfigRequest2).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `NetworkConfigApiAPI.NetworkConfigApiUpdateNetworkConfig``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**serviceId** | **string** | The ID of the service | 
**id** | **string** | The ID of the network config | 

### Other Parameters

Other parameters are passed through a pointer to a apiNetworkConfigApiUpdateNetworkConfigRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


 **updateNetworkConfigRequest2** | [**UpdateNetworkConfigRequest2**](UpdateNetworkConfigRequest2.md) |  | 

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

