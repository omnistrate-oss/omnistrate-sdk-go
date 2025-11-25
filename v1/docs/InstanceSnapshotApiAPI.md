# \InstanceSnapshotApiAPI

All URIs are relative to *https://api.omnistrate.cloud*

Method | HTTP request | Description
------------- | ------------- | -------------
[**InstanceSnapshotApiDeleteResourceInstanceSnapshot**](InstanceSnapshotApiAPI.md#InstanceSnapshotApiDeleteResourceInstanceSnapshot) | **Delete** /2022-09-01-00/resource-instance/snapshot/{snapshotId} | DeleteResourceInstanceSnapshot instance-snapshot-api



## InstanceSnapshotApiDeleteResourceInstanceSnapshot

> InstanceSnapshotApiDeleteResourceInstanceSnapshot(ctx, snapshotId).SubscriptionId(subscriptionId).Execute()

DeleteResourceInstanceSnapshot instance-snapshot-api

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
	snapshotId := "instance-ss-12345678" // string | The instance snapshot ID
	subscriptionId := "sub-abcd1234" // string | Subscription Id (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	r, err := apiClient.InstanceSnapshotApiAPI.InstanceSnapshotApiDeleteResourceInstanceSnapshot(context.Background(), snapshotId).SubscriptionId(subscriptionId).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `InstanceSnapshotApiAPI.InstanceSnapshotApiDeleteResourceInstanceSnapshot``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**snapshotId** | **string** | The instance snapshot ID | 

### Other Parameters

Other parameters are passed through a pointer to a apiInstanceSnapshotApiDeleteResourceInstanceSnapshotRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **subscriptionId** | **string** | Subscription Id | 

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

