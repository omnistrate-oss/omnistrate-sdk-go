# \SubscriptionRequestApiAPI

All URIs are relative to *https://api.omnistrate.cloud*

Method | HTTP request | Description
------------- | ------------- | -------------
[**SubscriptionRequestApiCancelSubscriptionRequest**](SubscriptionRequestApiAPI.md#SubscriptionRequestApiCancelSubscriptionRequest) | **Delete** /2022-09-01-00/subscription/request/{id} | CancelSubscriptionRequest subscription-request-api
[**SubscriptionRequestApiCreateSubscriptionRequest**](SubscriptionRequestApiAPI.md#SubscriptionRequestApiCreateSubscriptionRequest) | **Post** /2022-09-01-00/subscription/request | CreateSubscriptionRequest subscription-request-api
[**SubscriptionRequestApiDescribeSubscriptionRequest**](SubscriptionRequestApiAPI.md#SubscriptionRequestApiDescribeSubscriptionRequest) | **Get** /2022-09-01-00/subscription/request/{id} | DescribeSubscriptionRequest subscription-request-api
[**SubscriptionRequestApiListSubscriptionRequests**](SubscriptionRequestApiAPI.md#SubscriptionRequestApiListSubscriptionRequests) | **Get** /2022-09-01-00/subscription/request | ListSubscriptionRequests subscription-request-api



## SubscriptionRequestApiCancelSubscriptionRequest

> SubscriptionRequestApiCancelSubscriptionRequest(ctx, id).Execute()

CancelSubscriptionRequest subscription-request-api

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
	id := "subr-12345678" // string | The subscription ID

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	r, err := apiClient.SubscriptionRequestApiAPI.SubscriptionRequestApiCancelSubscriptionRequest(context.Background(), id).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `SubscriptionRequestApiAPI.SubscriptionRequestApiCancelSubscriptionRequest``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **string** | The subscription ID | 

### Other Parameters

Other parameters are passed through a pointer to a apiSubscriptionRequestApiCancelSubscriptionRequestRequest struct via the builder pattern


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


## SubscriptionRequestApiCreateSubscriptionRequest

> string SubscriptionRequestApiCreateSubscriptionRequest(ctx).CreateSubscriptionRequestRequest2(createSubscriptionRequestRequest2).Execute()

CreateSubscriptionRequest subscription-request-api

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
	createSubscriptionRequestRequest2 := *openapiclient.NewCreateSubscriptionRequestRequest2("Qui praesentium sed.", "s-123456") // CreateSubscriptionRequestRequest2 | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.SubscriptionRequestApiAPI.SubscriptionRequestApiCreateSubscriptionRequest(context.Background()).CreateSubscriptionRequestRequest2(createSubscriptionRequestRequest2).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `SubscriptionRequestApiAPI.SubscriptionRequestApiCreateSubscriptionRequest``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `SubscriptionRequestApiCreateSubscriptionRequest`: string
	fmt.Fprintf(os.Stdout, "Response from `SubscriptionRequestApiAPI.SubscriptionRequestApiCreateSubscriptionRequest`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiSubscriptionRequestApiCreateSubscriptionRequestRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **createSubscriptionRequestRequest2** | [**CreateSubscriptionRequestRequest2**](CreateSubscriptionRequestRequest2.md) |  | 

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


## SubscriptionRequestApiDescribeSubscriptionRequest

> DescribeSubscriptionRequestResult SubscriptionRequestApiDescribeSubscriptionRequest(ctx, id).Execute()

DescribeSubscriptionRequest subscription-request-api

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
	id := "subr-12345678" // string | The subscription ID

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.SubscriptionRequestApiAPI.SubscriptionRequestApiDescribeSubscriptionRequest(context.Background(), id).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `SubscriptionRequestApiAPI.SubscriptionRequestApiDescribeSubscriptionRequest``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `SubscriptionRequestApiDescribeSubscriptionRequest`: DescribeSubscriptionRequestResult
	fmt.Fprintf(os.Stdout, "Response from `SubscriptionRequestApiAPI.SubscriptionRequestApiDescribeSubscriptionRequest`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **string** | The subscription ID | 

### Other Parameters

Other parameters are passed through a pointer to a apiSubscriptionRequestApiDescribeSubscriptionRequestRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

[**DescribeSubscriptionRequestResult**](DescribeSubscriptionRequestResult.md)

### Authorization

[api_key_header_Authorization](../README.md#api_key_header_Authorization)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json, application/vnd.goa.error

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## SubscriptionRequestApiListSubscriptionRequests

> ListSubscriptionRequestsResult SubscriptionRequestApiListSubscriptionRequests(ctx).Status(status).Execute()

ListSubscriptionRequests subscription-request-api

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
	status := "PENDING" // string | The status of the subscription request to filter by (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.SubscriptionRequestApiAPI.SubscriptionRequestApiListSubscriptionRequests(context.Background()).Status(status).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `SubscriptionRequestApiAPI.SubscriptionRequestApiListSubscriptionRequests``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `SubscriptionRequestApiListSubscriptionRequests`: ListSubscriptionRequestsResult
	fmt.Fprintf(os.Stdout, "Response from `SubscriptionRequestApiAPI.SubscriptionRequestApiListSubscriptionRequests`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiSubscriptionRequestApiListSubscriptionRequestsRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **status** | **string** | The status of the subscription request to filter by | 

### Return type

[**ListSubscriptionRequestsResult**](ListSubscriptionRequestsResult.md)

### Authorization

[api_key_header_Authorization](../README.md#api_key_header_Authorization)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json, application/vnd.goa.error

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

