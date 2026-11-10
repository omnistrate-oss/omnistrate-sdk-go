# \AccessRequestsApiAPI

All URIs are relative to *https://api.omnistrate.cloud*

Method | HTTP request | Description
------------- | ------------- | -------------
[**AccessRequestsApiApproveAccessRequest**](AccessRequestsApiAPI.md#AccessRequestsApiApproveAccessRequest) | **Post** /2022-09-01-00/fleet/access-requests/{id}/approve | ApproveAccessRequest access-requests-api
[**AccessRequestsApiBatchRevokeAccessRequests**](AccessRequestsApiAPI.md#AccessRequestsApiBatchRevokeAccessRequests) | **Post** /2022-09-01-00/fleet/access-requests/batch-revoke | BatchRevokeAccessRequests access-requests-api
[**AccessRequestsApiCreateAccessRequest**](AccessRequestsApiAPI.md#AccessRequestsApiCreateAccessRequest) | **Post** /2022-09-01-00/fleet/access-requests | CreateAccessRequest access-requests-api
[**AccessRequestsApiDenyAccessRequest**](AccessRequestsApiAPI.md#AccessRequestsApiDenyAccessRequest) | **Post** /2022-09-01-00/fleet/access-requests/{id}/deny | DenyAccessRequest access-requests-api
[**AccessRequestsApiDescribeAccessRequest**](AccessRequestsApiAPI.md#AccessRequestsApiDescribeAccessRequest) | **Get** /2022-09-01-00/fleet/access-requests/{id} | DescribeAccessRequest access-requests-api
[**AccessRequestsApiListAccessRequests**](AccessRequestsApiAPI.md#AccessRequestsApiListAccessRequests) | **Get** /2022-09-01-00/fleet/access-requests | ListAccessRequests access-requests-api
[**AccessRequestsApiRevokeAccessRequest**](AccessRequestsApiAPI.md#AccessRequestsApiRevokeAccessRequest) | **Post** /2022-09-01-00/fleet/access-requests/{id}/revoke | RevokeAccessRequest access-requests-api



## AccessRequestsApiApproveAccessRequest

> AccessRequest AccessRequestsApiApproveAccessRequest(ctx, id).ApproveAccessRequestRequest2(approveAccessRequestRequest2).Execute()

ApproveAccessRequest access-requests-api



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
	id := "ar-12345678" // string | The ID of the pending access request to approve.
	approveAccessRequestRequest2 := *openapiclient.NewApproveAccessRequestRequest2() // ApproveAccessRequestRequest2 | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.AccessRequestsApiAPI.AccessRequestsApiApproveAccessRequest(context.Background(), id).ApproveAccessRequestRequest2(approveAccessRequestRequest2).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `AccessRequestsApiAPI.AccessRequestsApiApproveAccessRequest``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `AccessRequestsApiApproveAccessRequest`: AccessRequest
	fmt.Fprintf(os.Stdout, "Response from `AccessRequestsApiAPI.AccessRequestsApiApproveAccessRequest`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **string** | The ID of the pending access request to approve. | 

### Other Parameters

Other parameters are passed through a pointer to a apiAccessRequestsApiApproveAccessRequestRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **approveAccessRequestRequest2** | [**ApproveAccessRequestRequest2**](ApproveAccessRequestRequest2.md) |  | 

### Return type

[**AccessRequest**](AccessRequest.md)

### Authorization

[api_key_header_Authorization](../README.md#api_key_header_Authorization)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json, application/vnd.goa.error

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## AccessRequestsApiBatchRevokeAccessRequests

> BatchRevokeAccessRequestsResult AccessRequestsApiBatchRevokeAccessRequests(ctx).BatchRevokeAccessRequestsRequest2(batchRevokeAccessRequestsRequest2).Execute()

BatchRevokeAccessRequests access-requests-api



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
	batchRevokeAccessRequestsRequest2 := *openapiclient.NewBatchRevokeAccessRequestsRequest2([]string{"ar-12345678"}) // BatchRevokeAccessRequestsRequest2 | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.AccessRequestsApiAPI.AccessRequestsApiBatchRevokeAccessRequests(context.Background()).BatchRevokeAccessRequestsRequest2(batchRevokeAccessRequestsRequest2).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `AccessRequestsApiAPI.AccessRequestsApiBatchRevokeAccessRequests``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `AccessRequestsApiBatchRevokeAccessRequests`: BatchRevokeAccessRequestsResult
	fmt.Fprintf(os.Stdout, "Response from `AccessRequestsApiAPI.AccessRequestsApiBatchRevokeAccessRequests`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiAccessRequestsApiBatchRevokeAccessRequestsRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **batchRevokeAccessRequestsRequest2** | [**BatchRevokeAccessRequestsRequest2**](BatchRevokeAccessRequestsRequest2.md) |  | 

### Return type

[**BatchRevokeAccessRequestsResult**](BatchRevokeAccessRequestsResult.md)

### Authorization

[api_key_header_Authorization](../README.md#api_key_header_Authorization)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json, application/vnd.goa.error

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## AccessRequestsApiCreateAccessRequest

> AccessRequest AccessRequestsApiCreateAccessRequest(ctx).CreateAccessRequestRequest2(createAccessRequestRequest2).Execute()

CreateAccessRequest access-requests-api



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
	createAccessRequestRequest2 := *openapiclient.NewCreateAccessRequestRequest2("Sint similique est.", "cluster-reader|cluster-admin") // CreateAccessRequestRequest2 | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.AccessRequestsApiAPI.AccessRequestsApiCreateAccessRequest(context.Background()).CreateAccessRequestRequest2(createAccessRequestRequest2).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `AccessRequestsApiAPI.AccessRequestsApiCreateAccessRequest``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `AccessRequestsApiCreateAccessRequest`: AccessRequest
	fmt.Fprintf(os.Stdout, "Response from `AccessRequestsApiAPI.AccessRequestsApiCreateAccessRequest`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiAccessRequestsApiCreateAccessRequestRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **createAccessRequestRequest2** | [**CreateAccessRequestRequest2**](CreateAccessRequestRequest2.md) |  | 

### Return type

[**AccessRequest**](AccessRequest.md)

### Authorization

[api_key_header_Authorization](../README.md#api_key_header_Authorization)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json, application/vnd.goa.error

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## AccessRequestsApiDenyAccessRequest

> AccessRequest AccessRequestsApiDenyAccessRequest(ctx, id).Execute()

DenyAccessRequest access-requests-api



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
	id := "ar-12345678" // string | The ID of the pending access request to deny.

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.AccessRequestsApiAPI.AccessRequestsApiDenyAccessRequest(context.Background(), id).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `AccessRequestsApiAPI.AccessRequestsApiDenyAccessRequest``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `AccessRequestsApiDenyAccessRequest`: AccessRequest
	fmt.Fprintf(os.Stdout, "Response from `AccessRequestsApiAPI.AccessRequestsApiDenyAccessRequest`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **string** | The ID of the pending access request to deny. | 

### Other Parameters

Other parameters are passed through a pointer to a apiAccessRequestsApiDenyAccessRequestRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

[**AccessRequest**](AccessRequest.md)

### Authorization

[api_key_header_Authorization](../README.md#api_key_header_Authorization)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json, application/vnd.goa.error

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## AccessRequestsApiDescribeAccessRequest

> AccessRequest AccessRequestsApiDescribeAccessRequest(ctx, id).Execute()

DescribeAccessRequest access-requests-api



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
	id := "ar-12345678" // string | The ID of the access request to describe.

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.AccessRequestsApiAPI.AccessRequestsApiDescribeAccessRequest(context.Background(), id).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `AccessRequestsApiAPI.AccessRequestsApiDescribeAccessRequest``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `AccessRequestsApiDescribeAccessRequest`: AccessRequest
	fmt.Fprintf(os.Stdout, "Response from `AccessRequestsApiAPI.AccessRequestsApiDescribeAccessRequest`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **string** | The ID of the access request to describe. | 

### Other Parameters

Other parameters are passed through a pointer to a apiAccessRequestsApiDescribeAccessRequestRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

[**AccessRequest**](AccessRequest.md)

### Authorization

[api_key_header_Authorization](../README.md#api_key_header_Authorization)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json, application/vnd.goa.error

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## AccessRequestsApiListAccessRequests

> ListAccessRequestsResult AccessRequestsApiListAccessRequests(ctx).Statuses(statuses).IncludeDeleted(includeDeleted).Limit(limit).Execute()

ListAccessRequests access-requests-api



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
	statuses := []string{"PENDING|APPROVING|APPROVED|REVOKING|DENIED|REVOKED|EXPIRED"} // []string | Optional lifecycle statuses to match. Omission or an empty list returns all statuses. (optional)
	includeDeleted := false // bool | Whether to include all retained soft-deleted requests instead of only terminal requests from the previous 24 hours; omission defaults to false. (optional)
	limit := int64(1000) // int64 | Maximum requests returned with includeDeleted=true; omission defaults to 1000 and the allowed range is 1-5000. (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.AccessRequestsApiAPI.AccessRequestsApiListAccessRequests(context.Background()).Statuses(statuses).IncludeDeleted(includeDeleted).Limit(limit).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `AccessRequestsApiAPI.AccessRequestsApiListAccessRequests``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `AccessRequestsApiListAccessRequests`: ListAccessRequestsResult
	fmt.Fprintf(os.Stdout, "Response from `AccessRequestsApiAPI.AccessRequestsApiListAccessRequests`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiAccessRequestsApiListAccessRequestsRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **statuses** | **[]string** | Optional lifecycle statuses to match. Omission or an empty list returns all statuses. | 
 **includeDeleted** | **bool** | Whether to include all retained soft-deleted requests instead of only terminal requests from the previous 24 hours; omission defaults to false. | 
 **limit** | **int64** | Maximum requests returned with includeDeleted&#x3D;true; omission defaults to 1000 and the allowed range is 1-5000. | 

### Return type

[**ListAccessRequestsResult**](ListAccessRequestsResult.md)

### Authorization

[api_key_header_Authorization](../README.md#api_key_header_Authorization)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json, application/vnd.goa.error

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## AccessRequestsApiRevokeAccessRequest

> AccessRequest AccessRequestsApiRevokeAccessRequest(ctx, id).Execute()

RevokeAccessRequest access-requests-api



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
	id := "ar-12345678" // string | The ID of the approved access request to revoke.

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.AccessRequestsApiAPI.AccessRequestsApiRevokeAccessRequest(context.Background(), id).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `AccessRequestsApiAPI.AccessRequestsApiRevokeAccessRequest``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `AccessRequestsApiRevokeAccessRequest`: AccessRequest
	fmt.Fprintf(os.Stdout, "Response from `AccessRequestsApiAPI.AccessRequestsApiRevokeAccessRequest`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **string** | The ID of the approved access request to revoke. | 

### Other Parameters

Other parameters are passed through a pointer to a apiAccessRequestsApiRevokeAccessRequestRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

[**AccessRequest**](AccessRequest.md)

### Authorization

[api_key_header_Authorization](../README.md#api_key_header_Authorization)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json, application/vnd.goa.error

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

