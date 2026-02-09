# \CustomNetworkApiAPI

All URIs are relative to *https://api.omnistrate.cloud*

Method | HTTP request | Description
------------- | ------------- | -------------
[**CustomNetworkApiCreateCustomNetwork**](CustomNetworkApiAPI.md#CustomNetworkApiCreateCustomNetwork) | **Post** /2022-09-01-00/resource-instance/custom-network | CreateCustomNetwork custom-network-api
[**CustomNetworkApiDeleteCustomNetwork**](CustomNetworkApiAPI.md#CustomNetworkApiDeleteCustomNetwork) | **Delete** /2022-09-01-00/resource-instance/custom-network/{id} | DeleteCustomNetwork custom-network-api
[**CustomNetworkApiDescribeCustomNetwork**](CustomNetworkApiAPI.md#CustomNetworkApiDescribeCustomNetwork) | **Get** /2022-09-01-00/resource-instance/custom-network/{id} | DescribeCustomNetwork custom-network-api
[**CustomNetworkApiListCustomNetworks**](CustomNetworkApiAPI.md#CustomNetworkApiListCustomNetworks) | **Get** /2022-09-01-00/resource-instance/custom-network | ListCustomNetworks custom-network-api
[**CustomNetworkApiUpdateCustomNetwork**](CustomNetworkApiAPI.md#CustomNetworkApiUpdateCustomNetwork) | **Patch** /2022-09-01-00/resource-instance/custom-network/{id} | UpdateCustomNetwork custom-network-api



## CustomNetworkApiCreateCustomNetwork

> CustomNetwork CustomNetworkApiCreateCustomNetwork(ctx).CreateCustomNetworkRequest2(createCustomNetworkRequest2).Execute()

CreateCustomNetwork custom-network-api

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
	createCustomNetworkRequest2 := *openapiclient.NewCreateCustomNetworkRequest2("aws|azure|gcp|all", "us-east-1") // CreateCustomNetworkRequest2 | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.CustomNetworkApiAPI.CustomNetworkApiCreateCustomNetwork(context.Background()).CreateCustomNetworkRequest2(createCustomNetworkRequest2).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `CustomNetworkApiAPI.CustomNetworkApiCreateCustomNetwork``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `CustomNetworkApiCreateCustomNetwork`: CustomNetwork
	fmt.Fprintf(os.Stdout, "Response from `CustomNetworkApiAPI.CustomNetworkApiCreateCustomNetwork`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiCustomNetworkApiCreateCustomNetworkRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **createCustomNetworkRequest2** | [**CreateCustomNetworkRequest2**](CreateCustomNetworkRequest2.md) |  | 

### Return type

[**CustomNetwork**](CustomNetwork.md)

### Authorization

[api_key_header_Authorization](../README.md#api_key_header_Authorization)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json, application/vnd.goa.error

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## CustomNetworkApiDeleteCustomNetwork

> CustomNetworkApiDeleteCustomNetwork(ctx, id).Execute()

DeleteCustomNetwork custom-network-api

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
	id := "n-12345678" // string | ID of a custom network

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	r, err := apiClient.CustomNetworkApiAPI.CustomNetworkApiDeleteCustomNetwork(context.Background(), id).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `CustomNetworkApiAPI.CustomNetworkApiDeleteCustomNetwork``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **string** | ID of a custom network | 

### Other Parameters

Other parameters are passed through a pointer to a apiCustomNetworkApiDeleteCustomNetworkRequest struct via the builder pattern


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


## CustomNetworkApiDescribeCustomNetwork

> CustomNetwork CustomNetworkApiDescribeCustomNetwork(ctx, id).Execute()

DescribeCustomNetwork custom-network-api

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
	id := "n-12345678" // string | ID of a custom network

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.CustomNetworkApiAPI.CustomNetworkApiDescribeCustomNetwork(context.Background(), id).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `CustomNetworkApiAPI.CustomNetworkApiDescribeCustomNetwork``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `CustomNetworkApiDescribeCustomNetwork`: CustomNetwork
	fmt.Fprintf(os.Stdout, "Response from `CustomNetworkApiAPI.CustomNetworkApiDescribeCustomNetwork`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **string** | ID of a custom network | 

### Other Parameters

Other parameters are passed through a pointer to a apiCustomNetworkApiDescribeCustomNetworkRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

[**CustomNetwork**](CustomNetwork.md)

### Authorization

[api_key_header_Authorization](../README.md#api_key_header_Authorization)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json, application/vnd.goa.error

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## CustomNetworkApiListCustomNetworks

> ListCustomNetworksResult CustomNetworkApiListCustomNetworks(ctx).CloudProviderName(cloudProviderName).CloudProviderRegion(cloudProviderRegion).CustomNetworksOnly(customNetworksOnly).SubscriptionId(subscriptionId).Execute()

ListCustomNetworks custom-network-api

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
	cloudProviderName := "aws" // string | The name of the cloud provider to filter custom networks by (optional)
	cloudProviderRegion := "us-east-1" // string | The region of the cloud provider to filter custom networks by (optional)
	customNetworksOnly := false // bool | Flag indicating whether to return only custom networks, or to include default and imported networks as well (optional)
	subscriptionId := "sub-12345678" // string | List available custom networks for the given subscription ID (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.CustomNetworkApiAPI.CustomNetworkApiListCustomNetworks(context.Background()).CloudProviderName(cloudProviderName).CloudProviderRegion(cloudProviderRegion).CustomNetworksOnly(customNetworksOnly).SubscriptionId(subscriptionId).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `CustomNetworkApiAPI.CustomNetworkApiListCustomNetworks``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `CustomNetworkApiListCustomNetworks`: ListCustomNetworksResult
	fmt.Fprintf(os.Stdout, "Response from `CustomNetworkApiAPI.CustomNetworkApiListCustomNetworks`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiCustomNetworkApiListCustomNetworksRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **cloudProviderName** | **string** | The name of the cloud provider to filter custom networks by | 
 **cloudProviderRegion** | **string** | The region of the cloud provider to filter custom networks by | 
 **customNetworksOnly** | **bool** | Flag indicating whether to return only custom networks, or to include default and imported networks as well | 
 **subscriptionId** | **string** | List available custom networks for the given subscription ID | 

### Return type

[**ListCustomNetworksResult**](ListCustomNetworksResult.md)

### Authorization

[api_key_header_Authorization](../README.md#api_key_header_Authorization)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json, application/vnd.goa.error

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## CustomNetworkApiUpdateCustomNetwork

> CustomNetwork CustomNetworkApiUpdateCustomNetwork(ctx, id).UpdateCustomNetworkRequest2(updateCustomNetworkRequest2).Execute()

UpdateCustomNetwork custom-network-api

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
	id := "n-12345678" // string | ID of a custom network
	updateCustomNetworkRequest2 := *openapiclient.NewUpdateCustomNetworkRequest2() // UpdateCustomNetworkRequest2 | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.CustomNetworkApiAPI.CustomNetworkApiUpdateCustomNetwork(context.Background(), id).UpdateCustomNetworkRequest2(updateCustomNetworkRequest2).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `CustomNetworkApiAPI.CustomNetworkApiUpdateCustomNetwork``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `CustomNetworkApiUpdateCustomNetwork`: CustomNetwork
	fmt.Fprintf(os.Stdout, "Response from `CustomNetworkApiAPI.CustomNetworkApiUpdateCustomNetwork`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **string** | ID of a custom network | 

### Other Parameters

Other parameters are passed through a pointer to a apiCustomNetworkApiUpdateCustomNetworkRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **updateCustomNetworkRequest2** | [**UpdateCustomNetworkRequest2**](UpdateCustomNetworkRequest2.md) |  | 

### Return type

[**CustomNetwork**](CustomNetwork.md)

### Authorization

[api_key_header_Authorization](../README.md#api_key_header_Authorization)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json, application/vnd.goa.error

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

