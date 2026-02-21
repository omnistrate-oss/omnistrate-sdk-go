# \InstanceSnapshotApiAPI

All URIs are relative to *https://api.omnistrate.cloud*

Method | HTTP request | Description
------------- | ------------- | -------------
[**InstanceSnapshotApiCopyInstanceSnapshot**](InstanceSnapshotApiAPI.md#InstanceSnapshotApiCopyInstanceSnapshot) | **Post** /2022-09-01-00/resource-instance/snapshot/{sourceSnapshotId} | CopyInstanceSnapshot instance-snapshot-api
[**InstanceSnapshotApiCreateInstanceSnapshot**](InstanceSnapshotApiAPI.md#InstanceSnapshotApiCreateInstanceSnapshot) | **Post** /2022-09-01-00/resource-instance/snapshot | CreateInstanceSnapshot instance-snapshot-api
[**InstanceSnapshotApiDeleteInstanceSnapshot**](InstanceSnapshotApiAPI.md#InstanceSnapshotApiDeleteInstanceSnapshot) | **Delete** /2022-09-01-00/resource-instance/snapshot/{id} | DeleteInstanceSnapshot instance-snapshot-api
[**InstanceSnapshotApiDescribeInstanceSnapshot**](InstanceSnapshotApiAPI.md#InstanceSnapshotApiDescribeInstanceSnapshot) | **Get** /2022-09-01-00/resource-instance/snapshot/{id} | DescribeInstanceSnapshot instance-snapshot-api
[**InstanceSnapshotApiListAllInstanceSnapshots**](InstanceSnapshotApiAPI.md#InstanceSnapshotApiListAllInstanceSnapshots) | **Get** /2022-09-01-00/resource-instance/snapshot | ListAllInstanceSnapshots instance-snapshot-api
[**InstanceSnapshotApiRestoreResourceInstanceFromSnapshot**](InstanceSnapshotApiAPI.md#InstanceSnapshotApiRestoreResourceInstanceFromSnapshot) | **Post** /2022-09-01-00/resource-instance/snapshot/{snapshotId}/restore | RestoreResourceInstanceFromSnapshot instance-snapshot-api



## InstanceSnapshotApiCopyInstanceSnapshot

> CreateServicesOrchestrationResponseBody InstanceSnapshotApiCopyInstanceSnapshot(ctx, sourceSnapshotId).CopyInstanceSnapshotRequest2(copyInstanceSnapshotRequest2).SubscriptionId(subscriptionId).Execute()

CopyInstanceSnapshot instance-snapshot-api

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
	sourceSnapshotId := "instance-ss-abcd1234" // string | The source snapshot ID
	copyInstanceSnapshotRequest2 := *openapiclient.NewCopyInstanceSnapshotRequest2("us-west-2") // CopyInstanceSnapshotRequest2 | 
	subscriptionId := "sub-abcd1234" // string | The subscription ID (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.InstanceSnapshotApiAPI.InstanceSnapshotApiCopyInstanceSnapshot(context.Background(), sourceSnapshotId).CopyInstanceSnapshotRequest2(copyInstanceSnapshotRequest2).SubscriptionId(subscriptionId).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `InstanceSnapshotApiAPI.InstanceSnapshotApiCopyInstanceSnapshot``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `InstanceSnapshotApiCopyInstanceSnapshot`: CreateServicesOrchestrationResponseBody
	fmt.Fprintf(os.Stdout, "Response from `InstanceSnapshotApiAPI.InstanceSnapshotApiCopyInstanceSnapshot`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**sourceSnapshotId** | **string** | The source snapshot ID | 

### Other Parameters

Other parameters are passed through a pointer to a apiInstanceSnapshotApiCopyInstanceSnapshotRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **copyInstanceSnapshotRequest2** | [**CopyInstanceSnapshotRequest2**](CopyInstanceSnapshotRequest2.md) |  | 
 **subscriptionId** | **string** | The subscription ID | 

### Return type

[**CreateServicesOrchestrationResponseBody**](CreateServicesOrchestrationResponseBody.md)

### Authorization

[api_key_header_Authorization](../README.md#api_key_header_Authorization)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json, application/vnd.goa.error

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## InstanceSnapshotApiCreateInstanceSnapshot

> CreateServicesOrchestrationResponseBody InstanceSnapshotApiCreateInstanceSnapshot(ctx).CreateInstanceSnapshotRequest2(createInstanceSnapshotRequest2).SubscriptionId(subscriptionId).Execute()

CreateInstanceSnapshot instance-snapshot-api

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
	createInstanceSnapshotRequest2 := *openapiclient.NewCreateInstanceSnapshotRequest2("instance-abcd1234") // CreateInstanceSnapshotRequest2 | 
	subscriptionId := "sub-abcd1234" // string | The subscription ID (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.InstanceSnapshotApiAPI.InstanceSnapshotApiCreateInstanceSnapshot(context.Background()).CreateInstanceSnapshotRequest2(createInstanceSnapshotRequest2).SubscriptionId(subscriptionId).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `InstanceSnapshotApiAPI.InstanceSnapshotApiCreateInstanceSnapshot``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `InstanceSnapshotApiCreateInstanceSnapshot`: CreateServicesOrchestrationResponseBody
	fmt.Fprintf(os.Stdout, "Response from `InstanceSnapshotApiAPI.InstanceSnapshotApiCreateInstanceSnapshot`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiInstanceSnapshotApiCreateInstanceSnapshotRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **createInstanceSnapshotRequest2** | [**CreateInstanceSnapshotRequest2**](CreateInstanceSnapshotRequest2.md) |  | 
 **subscriptionId** | **string** | The subscription ID | 

### Return type

[**CreateServicesOrchestrationResponseBody**](CreateServicesOrchestrationResponseBody.md)

### Authorization

[api_key_header_Authorization](../README.md#api_key_header_Authorization)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json, application/vnd.goa.error

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## InstanceSnapshotApiDeleteInstanceSnapshot

> InstanceSnapshotApiDeleteInstanceSnapshot(ctx, id).SubscriptionId(subscriptionId).Execute()

DeleteInstanceSnapshot instance-snapshot-api

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
	id := "instance-ss-12345678" // string | The instance snapshot ID
	subscriptionId := "sub-abcd1234" // string | The subscription ID (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	r, err := apiClient.InstanceSnapshotApiAPI.InstanceSnapshotApiDeleteInstanceSnapshot(context.Background(), id).SubscriptionId(subscriptionId).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `InstanceSnapshotApiAPI.InstanceSnapshotApiDeleteInstanceSnapshot``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **string** | The instance snapshot ID | 

### Other Parameters

Other parameters are passed through a pointer to a apiInstanceSnapshotApiDeleteInstanceSnapshotRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **subscriptionId** | **string** | The subscription ID | 

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


## InstanceSnapshotApiDescribeInstanceSnapshot

> DescribeInstanceSnapshotResult InstanceSnapshotApiDescribeInstanceSnapshot(ctx, id).SubscriptionId(subscriptionId).Execute()

DescribeInstanceSnapshot instance-snapshot-api

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
	id := "instance-ss-12345678" // string | The instance snapshot ID
	subscriptionId := "sub-abcd1234" // string | The subscription ID (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.InstanceSnapshotApiAPI.InstanceSnapshotApiDescribeInstanceSnapshot(context.Background(), id).SubscriptionId(subscriptionId).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `InstanceSnapshotApiAPI.InstanceSnapshotApiDescribeInstanceSnapshot``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `InstanceSnapshotApiDescribeInstanceSnapshot`: DescribeInstanceSnapshotResult
	fmt.Fprintf(os.Stdout, "Response from `InstanceSnapshotApiAPI.InstanceSnapshotApiDescribeInstanceSnapshot`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **string** | The instance snapshot ID | 

### Other Parameters

Other parameters are passed through a pointer to a apiInstanceSnapshotApiDescribeInstanceSnapshotRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **subscriptionId** | **string** | The subscription ID | 

### Return type

[**DescribeInstanceSnapshotResult**](DescribeInstanceSnapshotResult.md)

### Authorization

[api_key_header_Authorization](../README.md#api_key_header_Authorization)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json, application/vnd.goa.error

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## InstanceSnapshotApiListAllInstanceSnapshots

> ListInstanceSnapshotsResult InstanceSnapshotApiListAllInstanceSnapshots(ctx).EnvironmentType(environmentType).SnapshotType(snapshotType).Execute()

ListAllInstanceSnapshots instance-snapshot-api

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
	environmentType := "DEV" // string | The environment type to filter by (optional)
	snapshotType := "AutomatedSnapshot|ManualSnapshot" // string | The type of snapshot to list. Valid values are: 'ManualSnapshot' and 'AutomatedSnapshot' (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.InstanceSnapshotApiAPI.InstanceSnapshotApiListAllInstanceSnapshots(context.Background()).EnvironmentType(environmentType).SnapshotType(snapshotType).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `InstanceSnapshotApiAPI.InstanceSnapshotApiListAllInstanceSnapshots``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `InstanceSnapshotApiListAllInstanceSnapshots`: ListInstanceSnapshotsResult
	fmt.Fprintf(os.Stdout, "Response from `InstanceSnapshotApiAPI.InstanceSnapshotApiListAllInstanceSnapshots`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiInstanceSnapshotApiListAllInstanceSnapshotsRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **environmentType** | **string** | The environment type to filter by | 
 **snapshotType** | **string** | The type of snapshot to list. Valid values are: &#39;ManualSnapshot&#39; and &#39;AutomatedSnapshot&#39; | 

### Return type

[**ListInstanceSnapshotsResult**](ListInstanceSnapshotsResult.md)

### Authorization

[api_key_header_Authorization](../README.md#api_key_header_Authorization)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json, application/vnd.goa.error

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## InstanceSnapshotApiRestoreResourceInstanceFromSnapshot

> RestoreResourceInstanceFromSnapshotResponseBody InstanceSnapshotApiRestoreResourceInstanceFromSnapshot(ctx, snapshotId).RestoreFromInstanceSnapshotRequest2(restoreFromInstanceSnapshotRequest2).SubscriptionId(subscriptionId).Execute()

RestoreResourceInstanceFromSnapshot instance-snapshot-api

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
	snapshotId := "instance-ss-abcd1234" // string | The snapshot ID
	restoreFromInstanceSnapshotRequest2 := *openapiclient.NewRestoreFromInstanceSnapshotRequest2() // RestoreFromInstanceSnapshotRequest2 | 
	subscriptionId := "sub-abcd1234" // string | The subscription ID (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.InstanceSnapshotApiAPI.InstanceSnapshotApiRestoreResourceInstanceFromSnapshot(context.Background(), snapshotId).RestoreFromInstanceSnapshotRequest2(restoreFromInstanceSnapshotRequest2).SubscriptionId(subscriptionId).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `InstanceSnapshotApiAPI.InstanceSnapshotApiRestoreResourceInstanceFromSnapshot``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `InstanceSnapshotApiRestoreResourceInstanceFromSnapshot`: RestoreResourceInstanceFromSnapshotResponseBody
	fmt.Fprintf(os.Stdout, "Response from `InstanceSnapshotApiAPI.InstanceSnapshotApiRestoreResourceInstanceFromSnapshot`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**snapshotId** | **string** | The snapshot ID | 

### Other Parameters

Other parameters are passed through a pointer to a apiInstanceSnapshotApiRestoreResourceInstanceFromSnapshotRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **restoreFromInstanceSnapshotRequest2** | [**RestoreFromInstanceSnapshotRequest2**](RestoreFromInstanceSnapshotRequest2.md) |  | 
 **subscriptionId** | **string** | The subscription ID | 

### Return type

[**RestoreResourceInstanceFromSnapshotResponseBody**](RestoreResourceInstanceFromSnapshotResponseBody.md)

### Authorization

[api_key_header_Authorization](../README.md#api_key_header_Authorization)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json, application/vnd.goa.error

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

