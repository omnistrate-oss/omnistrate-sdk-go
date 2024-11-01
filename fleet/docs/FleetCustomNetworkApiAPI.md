# \FleetCustomNetworkApiAPI

All URIs are relative to *https://api.omnistrate.cloud*

Method | HTTP request | Description
------------- | ------------- | -------------
[**FleetCustomNetworkApiCreateCustomNetwork**](FleetCustomNetworkApiAPI.md#FleetCustomNetworkApiCreateCustomNetwork) | **Post** /2022-09-01-00/fleet/custom-network | CreateCustomNetwork fleet-custom-network-api
[**FleetCustomNetworkApiDeleteCustomNetwork**](FleetCustomNetworkApiAPI.md#FleetCustomNetworkApiDeleteCustomNetwork) | **Delete** /2022-09-01-00/fleet/custom-network/{id} | DeleteCustomNetwork fleet-custom-network-api
[**FleetCustomNetworkApiDescribeCustomNetwork**](FleetCustomNetworkApiAPI.md#FleetCustomNetworkApiDescribeCustomNetwork) | **Get** /2022-09-01-00/fleet/custom-network/{id} | DescribeCustomNetwork fleet-custom-network-api
[**FleetCustomNetworkApiListCustomNetworks**](FleetCustomNetworkApiAPI.md#FleetCustomNetworkApiListCustomNetworks) | **Get** /2022-09-01-00/fleet/custom-network | ListCustomNetworks fleet-custom-network-api



## FleetCustomNetworkApiCreateCustomNetwork

> FleetCustomNetwork FleetCustomNetworkApiCreateCustomNetwork(ctx).CreateCustomNetworkRequestBody(createCustomNetworkRequestBody).Execute()

CreateCustomNetwork fleet-custom-network-api

### Example

```go
package main

import (
	"context"
	"fmt"
	"os"
	openapiclient "github.com/omnistrate-oss/omnistrate-sdk-go/fleet"
)

func main() {
	createCustomNetworkRequestBody := *openapiclient.NewCreateCustomNetworkRequestBody("10.0.0.0/16", "aws", "us-east-1") // CreateCustomNetworkRequestBody | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.FleetCustomNetworkApiAPI.FleetCustomNetworkApiCreateCustomNetwork(context.Background()).CreateCustomNetworkRequestBody(createCustomNetworkRequestBody).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `FleetCustomNetworkApiAPI.FleetCustomNetworkApiCreateCustomNetwork``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `FleetCustomNetworkApiCreateCustomNetwork`: FleetCustomNetwork
	fmt.Fprintf(os.Stdout, "Response from `FleetCustomNetworkApiAPI.FleetCustomNetworkApiCreateCustomNetwork`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiFleetCustomNetworkApiCreateCustomNetworkRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **createCustomNetworkRequestBody** | [**CreateCustomNetworkRequestBody**](CreateCustomNetworkRequestBody.md) |  | 

### Return type

[**FleetCustomNetwork**](FleetCustomNetwork.md)

### Authorization

[api_key_header_Authorization](../README.md#api_key_header_Authorization)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json, application/vnd.goa.error

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## FleetCustomNetworkApiDeleteCustomNetwork

> FleetCustomNetworkApiDeleteCustomNetwork(ctx, id).Execute()

DeleteCustomNetwork fleet-custom-network-api

### Example

```go
package main

import (
	"context"
	"fmt"
	"os"
	openapiclient "github.com/omnistrate-oss/omnistrate-sdk-go/fleet"
)

func main() {
	id := "n-12345678" // string | ID of a custom network

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	r, err := apiClient.FleetCustomNetworkApiAPI.FleetCustomNetworkApiDeleteCustomNetwork(context.Background(), id).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `FleetCustomNetworkApiAPI.FleetCustomNetworkApiDeleteCustomNetwork``: %v\n", err)
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

Other parameters are passed through a pointer to a apiFleetCustomNetworkApiDeleteCustomNetworkRequest struct via the builder pattern


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


## FleetCustomNetworkApiDescribeCustomNetwork

> FleetCustomNetwork FleetCustomNetworkApiDescribeCustomNetwork(ctx, id).Execute()

DescribeCustomNetwork fleet-custom-network-api

### Example

```go
package main

import (
	"context"
	"fmt"
	"os"
	openapiclient "github.com/omnistrate-oss/omnistrate-sdk-go/fleet"
)

func main() {
	id := "n-12345678" // string | ID of a custom network

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.FleetCustomNetworkApiAPI.FleetCustomNetworkApiDescribeCustomNetwork(context.Background(), id).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `FleetCustomNetworkApiAPI.FleetCustomNetworkApiDescribeCustomNetwork``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `FleetCustomNetworkApiDescribeCustomNetwork`: FleetCustomNetwork
	fmt.Fprintf(os.Stdout, "Response from `FleetCustomNetworkApiAPI.FleetCustomNetworkApiDescribeCustomNetwork`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **string** | ID of a custom network | 

### Other Parameters

Other parameters are passed through a pointer to a apiFleetCustomNetworkApiDescribeCustomNetworkRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

[**FleetCustomNetwork**](FleetCustomNetwork.md)

### Authorization

[api_key_header_Authorization](../README.md#api_key_header_Authorization)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json, application/vnd.goa.error

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## FleetCustomNetworkApiListCustomNetworks

> FleetListCustomNetworksResult FleetCustomNetworkApiListCustomNetworks(ctx).CloudProviderName(cloudProviderName).CloudProviderRegion(cloudProviderRegion).CustomNetworksOnly(customNetworksOnly).Execute()

ListCustomNetworks fleet-custom-network-api

### Example

```go
package main

import (
	"context"
	"fmt"
	"os"
	openapiclient "github.com/omnistrate-oss/omnistrate-sdk-go/fleet"
)

func main() {
	cloudProviderName := "aws" // string | The name of the cloud provider that custom network should be created in (optional)
	cloudProviderRegion := "us-east-1" // string | The region of the cloud provider that the network should be created in (optional)
	customNetworksOnly := false // bool | Flag indicating whether to return only custom networks, or to include default and imported networks as well (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.FleetCustomNetworkApiAPI.FleetCustomNetworkApiListCustomNetworks(context.Background()).CloudProviderName(cloudProviderName).CloudProviderRegion(cloudProviderRegion).CustomNetworksOnly(customNetworksOnly).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `FleetCustomNetworkApiAPI.FleetCustomNetworkApiListCustomNetworks``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `FleetCustomNetworkApiListCustomNetworks`: FleetListCustomNetworksResult
	fmt.Fprintf(os.Stdout, "Response from `FleetCustomNetworkApiAPI.FleetCustomNetworkApiListCustomNetworks`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiFleetCustomNetworkApiListCustomNetworksRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **cloudProviderName** | **string** | The name of the cloud provider that custom network should be created in | 
 **cloudProviderRegion** | **string** | The region of the cloud provider that the network should be created in | 
 **customNetworksOnly** | **bool** | Flag indicating whether to return only custom networks, or to include default and imported networks as well | 

### Return type

[**FleetListCustomNetworksResult**](FleetListCustomNetworksResult.md)

### Authorization

[api_key_header_Authorization](../README.md#api_key_header_Authorization)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json, application/vnd.goa.error

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

