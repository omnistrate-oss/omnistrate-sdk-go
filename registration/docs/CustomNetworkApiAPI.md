# \CustomNetworkApiAPI

All URIs are relative to *https://api.omnistrate.cloud*

Method | HTTP request | Description
------------- | ------------- | -------------
[**CustomNetworkApiCreateCustomNetwork**](CustomNetworkApiAPI.md#CustomNetworkApiCreateCustomNetwork) | **Post** /2022-09-01-00/resource-instance/custom-network | CreateCustomNetwork custom-network-api
[**CustomNetworkApiDeleteCustomNetwork**](CustomNetworkApiAPI.md#CustomNetworkApiDeleteCustomNetwork) | **Delete** /2022-09-01-00/resource-instance/custom-network/{id} | DeleteCustomNetwork custom-network-api
[**CustomNetworkApiDescribeCustomNetwork**](CustomNetworkApiAPI.md#CustomNetworkApiDescribeCustomNetwork) | **Get** /2022-09-01-00/resource-instance/custom-network/{id} | DescribeCustomNetwork custom-network-api
[**CustomNetworkApiListCustomNetworks**](CustomNetworkApiAPI.md#CustomNetworkApiListCustomNetworks) | **Get** /2022-09-01-00/resource-instance/custom-network | ListCustomNetworks custom-network-api



## CustomNetworkApiCreateCustomNetwork

> CustomNetwork CustomNetworkApiCreateCustomNetwork(ctx).CreateCustomNetworkRequestBody(createCustomNetworkRequestBody).Execute()

CreateCustomNetwork custom-network-api

### Example

```go
package main

import (
	"context"
	"fmt"
	"os"
	openapiclient "github.com/omnistrate/omnistrate-sdk-go/registration"
)

func main() {
	createCustomNetworkRequestBody := *openapiclient.NewCreateCustomNetworkRequestBody("aws", "us-east-1") // CreateCustomNetworkRequestBody | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.CustomNetworkApiAPI.CustomNetworkApiCreateCustomNetwork(context.Background()).CreateCustomNetworkRequestBody(createCustomNetworkRequestBody).Execute()
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
 **createCustomNetworkRequestBody** | [**CreateCustomNetworkRequestBody**](CreateCustomNetworkRequestBody.md) |  | 

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
	openapiclient "github.com/omnistrate/omnistrate-sdk-go/registration"
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
	openapiclient "github.com/omnistrate/omnistrate-sdk-go/registration"
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

> ListCustomNetworksResult CustomNetworkApiListCustomNetworks(ctx).CloudProviderName(cloudProviderName).CloudProviderRegion(cloudProviderRegion).Execute()

ListCustomNetworks custom-network-api

### Example

```go
package main

import (
	"context"
	"fmt"
	"os"
	openapiclient "github.com/omnistrate/omnistrate-sdk-go/registration"
)

func main() {
	cloudProviderName := "aws" // string | The name of the cloud provider that custom network should be created in (optional)
	cloudProviderRegion := "us-east-1" // string | The region of the cloud provider that the network should be created in (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.CustomNetworkApiAPI.CustomNetworkApiListCustomNetworks(context.Background()).CloudProviderName(cloudProviderName).CloudProviderRegion(cloudProviderRegion).Execute()
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
 **cloudProviderName** | **string** | The name of the cloud provider that custom network should be created in | 
 **cloudProviderRegion** | **string** | The region of the cloud provider that the network should be created in | 

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

