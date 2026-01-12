# \ResourceInstanceApiAPI

All URIs are relative to *https://api.omnistrate.cloud*

Method | HTTP request | Description
------------- | ------------- | -------------
[**ResourceInstanceApiAddCapacityToResourceInstance**](ResourceInstanceApiAPI.md#ResourceInstanceApiAddCapacityToResourceInstance) | **Post** /2022-09-01-00/resource-instance/{serviceProviderId}/{serviceKey}/{serviceAPIVersion}/{serviceEnvironmentKey}/{serviceModelKey}/{productTierKey}/{resourceKey}/{id}/add-capacity | AddCapacityToResourceInstance resource-instance-api
[**ResourceInstanceApiAddCustomDNSToResourceInstance**](ResourceInstanceApiAPI.md#ResourceInstanceApiAddCustomDNSToResourceInstance) | **Post** /2022-09-01-00/resource-instance/{serviceProviderId}/{serviceKey}/{serviceAPIVersion}/{serviceEnvironmentKey}/{serviceModelKey}/{productTierKey}/{resourceKey}/{id}/custom-dns | AddCustomDNSToResourceInstance resource-instance-api
[**ResourceInstanceApiCopyResourceInstanceSnapshot**](ResourceInstanceApiAPI.md#ResourceInstanceApiCopyResourceInstanceSnapshot) | **Post** /2022-09-01-00/resource-instance/{serviceProviderId}/{serviceKey}/{serviceAPIVersion}/{serviceEnvironmentKey}/{serviceModelKey}/{productTierKey}/{resourceKey}/{id}/copy-snapshot | CopyResourceInstanceSnapshot resource-instance-api
[**ResourceInstanceApiCreateResourceInstance**](ResourceInstanceApiAPI.md#ResourceInstanceApiCreateResourceInstance) | **Post** /2022-09-01-00/resource-instance/{serviceProviderId}/{serviceKey}/{serviceAPIVersion}/{serviceEnvironmentKey}/{serviceModelKey}/{productTierKey}/{resourceKey} | CreateResourceInstance resource-instance-api
[**ResourceInstanceApiDeleteResourceInstance**](ResourceInstanceApiAPI.md#ResourceInstanceApiDeleteResourceInstance) | **Delete** /2022-09-01-00/resource-instance/{serviceProviderId}/{serviceKey}/{serviceAPIVersion}/{serviceEnvironmentKey}/{serviceModelKey}/{productTierKey}/{resourceKey}/{id} | DeleteResourceInstance resource-instance-api
[**ResourceInstanceApiDescribeResourceInstance**](ResourceInstanceApiAPI.md#ResourceInstanceApiDescribeResourceInstance) | **Get** /2022-09-01-00/resource-instance/{serviceProviderId}/{serviceKey}/{serviceAPIVersion}/{serviceEnvironmentKey}/{serviceModelKey}/{productTierKey}/{resourceKey}/{id} | DescribeResourceInstance resource-instance-api
[**ResourceInstanceApiFailoverResourceInstance**](ResourceInstanceApiAPI.md#ResourceInstanceApiFailoverResourceInstance) | **Post** /2022-09-01-00/resource-instance/{serviceProviderId}/{serviceKey}/{serviceAPIVersion}/{serviceEnvironmentKey}/{serviceModelKey}/{productTierKey}/{resourceKey}/{id}/failover | FailoverResourceInstance resource-instance-api
[**ResourceInstanceApiGenerateTokenForDeploymentCellDashboard**](ResourceInstanceApiAPI.md#ResourceInstanceApiGenerateTokenForDeploymentCellDashboard) | **Post** /2022-09-01-00/resource-instance/{id}/deployment-cell-dashboard/token | GenerateTokenForDeploymentCellDashboard resource-instance-api
[**ResourceInstanceApiListAllResourceInstances**](ResourceInstanceApiAPI.md#ResourceInstanceApiListAllResourceInstances) | **Get** /2022-09-01-00/resource-instance | ListAllResourceInstances resource-instance-api
[**ResourceInstanceApiListResourceInstanceSnapshots**](ResourceInstanceApiAPI.md#ResourceInstanceApiListResourceInstanceSnapshots) | **Get** /2022-09-01-00/resource-instance/{serviceProviderId}/{serviceKey}/{serviceAPIVersion}/{serviceEnvironmentKey}/{serviceModelKey}/{productTierKey}/{resourceKey}/{id}/snapshot | ListResourceInstanceSnapshots resource-instance-api
[**ResourceInstanceApiListResourceInstances**](ResourceInstanceApiAPI.md#ResourceInstanceApiListResourceInstances) | **Get** /2022-09-01-00/resource-instance/{serviceProviderId}/{serviceKey}/{serviceAPIVersion}/{serviceEnvironmentKey}/{serviceModelKey}/{productTierKey}/{resourceKey} | ListResourceInstances resource-instance-api
[**ResourceInstanceApiRemoveCapacityFromResourceInstance**](ResourceInstanceApiAPI.md#ResourceInstanceApiRemoveCapacityFromResourceInstance) | **Post** /2022-09-01-00/resource-instance/{serviceProviderId}/{serviceKey}/{serviceAPIVersion}/{serviceEnvironmentKey}/{serviceModelKey}/{productTierKey}/{resourceKey}/{id}/remove-capacity | RemoveCapacityFromResourceInstance resource-instance-api
[**ResourceInstanceApiRemoveCustomDNSFromResourceInstance**](ResourceInstanceApiAPI.md#ResourceInstanceApiRemoveCustomDNSFromResourceInstance) | **Delete** /2022-09-01-00/resource-instance/{serviceProviderId}/{serviceKey}/{serviceAPIVersion}/{serviceEnvironmentKey}/{serviceModelKey}/{productTierKey}/{resourceKey}/{id}/custom-dns | RemoveCustomDNSFromResourceInstance resource-instance-api
[**ResourceInstanceApiResourceInstanceProvisionerSetupKit**](ResourceInstanceApiAPI.md#ResourceInstanceApiResourceInstanceProvisionerSetupKit) | **Get** /2022-09-01-00/resource-instance/{serviceProviderId}/{serviceKey}/{serviceAPIVersion}/{serviceEnvironmentKey}/{serviceModelKey}/setup-kit | ResourceInstanceProvisionerSetupKit resource-instance-api
[**ResourceInstanceApiRestartResourceInstance**](ResourceInstanceApiAPI.md#ResourceInstanceApiRestartResourceInstance) | **Post** /2022-09-01-00/resource-instance/{serviceProviderId}/{serviceKey}/{serviceAPIVersion}/{serviceEnvironmentKey}/{serviceModelKey}/{productTierKey}/{resourceKey}/{id}/restart | RestartResourceInstance resource-instance-api
[**ResourceInstanceApiRestoreResourceInstance**](ResourceInstanceApiAPI.md#ResourceInstanceApiRestoreResourceInstance) | **Post** /2022-09-01-00/resource-instance/{serviceProviderId}/{serviceKey}/{serviceAPIVersion}/{serviceEnvironmentKey}/{serviceModelKey}/{productTierKey}/{resourceKey}/{id}/restore | RestoreResourceInstance resource-instance-api
[**ResourceInstanceApiRestoreResourceInstanceFromSnapshot**](ResourceInstanceApiAPI.md#ResourceInstanceApiRestoreResourceInstanceFromSnapshot) | **Post** /2022-09-01-00/resource-instance/{serviceProviderId}/{serviceKey}/{serviceAPIVersion}/{serviceEnvironmentKey}/{serviceModelKey}/{productTierKey}/{resourceKey}/snapshot/{snapshotId}/restore | RestoreResourceInstanceFromSnapshot resource-instance-api
[**ResourceInstanceApiStartResourceInstance**](ResourceInstanceApiAPI.md#ResourceInstanceApiStartResourceInstance) | **Post** /2022-09-01-00/resource-instance/{serviceProviderId}/{serviceKey}/{serviceAPIVersion}/{serviceEnvironmentKey}/{serviceModelKey}/{productTierKey}/{resourceKey}/{id}/start | StartResourceInstance resource-instance-api
[**ResourceInstanceApiStopResourceInstance**](ResourceInstanceApiAPI.md#ResourceInstanceApiStopResourceInstance) | **Post** /2022-09-01-00/resource-instance/{serviceProviderId}/{serviceKey}/{serviceAPIVersion}/{serviceEnvironmentKey}/{serviceModelKey}/{productTierKey}/{resourceKey}/{id}/stop | StopResourceInstance resource-instance-api
[**ResourceInstanceApiUpdateAccountConfigResourceInstance**](ResourceInstanceApiAPI.md#ResourceInstanceApiUpdateAccountConfigResourceInstance) | **Post** /2022-09-01-00/resource-instance/account-config/{id} | UpdateAccountConfigResourceInstance resource-instance-api
[**ResourceInstanceApiUpdateResourceInstance**](ResourceInstanceApiAPI.md#ResourceInstanceApiUpdateResourceInstance) | **Patch** /2022-09-01-00/resource-instance/{serviceProviderId}/{serviceKey}/{serviceAPIVersion}/{serviceEnvironmentKey}/{serviceModelKey}/{productTierKey}/{resourceKey}/{id} | UpdateResourceInstance resource-instance-api
[**ResourceInstanceApiUpgradeResourceInstanceVersion**](ResourceInstanceApiAPI.md#ResourceInstanceApiUpgradeResourceInstanceVersion) | **Post** /2022-09-01-00/resource-instance/{id}/version-upgrade | UpgradeResourceInstanceVersion resource-instance-api



## ResourceInstanceApiAddCapacityToResourceInstance

> ResourceInstanceApiAddCapacityToResourceInstance(ctx, serviceProviderId, serviceKey, serviceAPIVersion, serviceEnvironmentKey, serviceModelKey, productTierKey, resourceKey, id).AddCapacityToResourceInstanceRequest2(addCapacityToResourceInstanceRequest2).SubscriptionId(subscriptionId).Execute()

AddCapacityToResourceInstance resource-instance-api

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
	serviceProviderId := "omnistrate" // string | The service provider ID
	serviceKey := "service-orchestration" // string | The service name
	serviceAPIVersion := "v1" // string | The service API version
	serviceEnvironmentKey := "dev" // string | The service environment name
	serviceModelKey := "hosted" // string | The service model name
	productTierKey := "premium" // string | The product tier name
	resourceKey := "mysql" // string | The resource key
	id := "instance-abcd1234" // string | The instance ID
	addCapacityToResourceInstanceRequest2 := *openapiclient.NewAddCapacityToResourceInstanceRequest2(int64(3)) // AddCapacityToResourceInstanceRequest2 | 
	subscriptionId := "sub-abcd1234" // string | Subscription Id (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	r, err := apiClient.ResourceInstanceApiAPI.ResourceInstanceApiAddCapacityToResourceInstance(context.Background(), serviceProviderId, serviceKey, serviceAPIVersion, serviceEnvironmentKey, serviceModelKey, productTierKey, resourceKey, id).AddCapacityToResourceInstanceRequest2(addCapacityToResourceInstanceRequest2).SubscriptionId(subscriptionId).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `ResourceInstanceApiAPI.ResourceInstanceApiAddCapacityToResourceInstance``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**serviceProviderId** | **string** | The service provider ID | 
**serviceKey** | **string** | The service name | 
**serviceAPIVersion** | **string** | The service API version | 
**serviceEnvironmentKey** | **string** | The service environment name | 
**serviceModelKey** | **string** | The service model name | 
**productTierKey** | **string** | The product tier name | 
**resourceKey** | **string** | The resource key | 
**id** | **string** | The instance ID | 

### Other Parameters

Other parameters are passed through a pointer to a apiResourceInstanceApiAddCapacityToResourceInstanceRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------








 **addCapacityToResourceInstanceRequest2** | [**AddCapacityToResourceInstanceRequest2**](AddCapacityToResourceInstanceRequest2.md) |  | 
 **subscriptionId** | **string** | Subscription Id | 

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


## ResourceInstanceApiAddCustomDNSToResourceInstance

> ResourceInstanceApiAddCustomDNSToResourceInstance(ctx, serviceProviderId, serviceKey, serviceAPIVersion, serviceEnvironmentKey, serviceModelKey, productTierKey, resourceKey, id).AddCustomDNSToResourceInstanceRequest2(addCustomDNSToResourceInstanceRequest2).SubscriptionId(subscriptionId).Execute()

AddCustomDNSToResourceInstance resource-instance-api

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
	serviceProviderId := "omnistrate" // string | The service provider ID
	serviceKey := "service-orchestration" // string | The service name
	serviceAPIVersion := "v1" // string | The service API version
	serviceEnvironmentKey := "dev" // string | The service environment name
	serviceModelKey := "hosted" // string | The service model name
	productTierKey := "premium" // string | The product tier name
	resourceKey := "http-service" // string | The resource key
	id := "instance-abcd1234" // string | The instance ID
	addCustomDNSToResourceInstanceRequest2 := *openapiclient.NewAddCustomDNSToResourceInstanceRequest2("my-custom-dns.com") // AddCustomDNSToResourceInstanceRequest2 | 
	subscriptionId := "sub-abcd1234" // string | Subscription Id (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	r, err := apiClient.ResourceInstanceApiAPI.ResourceInstanceApiAddCustomDNSToResourceInstance(context.Background(), serviceProviderId, serviceKey, serviceAPIVersion, serviceEnvironmentKey, serviceModelKey, productTierKey, resourceKey, id).AddCustomDNSToResourceInstanceRequest2(addCustomDNSToResourceInstanceRequest2).SubscriptionId(subscriptionId).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `ResourceInstanceApiAPI.ResourceInstanceApiAddCustomDNSToResourceInstance``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**serviceProviderId** | **string** | The service provider ID | 
**serviceKey** | **string** | The service name | 
**serviceAPIVersion** | **string** | The service API version | 
**serviceEnvironmentKey** | **string** | The service environment name | 
**serviceModelKey** | **string** | The service model name | 
**productTierKey** | **string** | The product tier name | 
**resourceKey** | **string** | The resource key | 
**id** | **string** | The instance ID | 

### Other Parameters

Other parameters are passed through a pointer to a apiResourceInstanceApiAddCustomDNSToResourceInstanceRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------








 **addCustomDNSToResourceInstanceRequest2** | [**AddCustomDNSToResourceInstanceRequest2**](AddCustomDNSToResourceInstanceRequest2.md) |  | 
 **subscriptionId** | **string** | Subscription Id | 

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


## ResourceInstanceApiCopyResourceInstanceSnapshot

> CopyResourceInstanceSnapshotResponseBody ResourceInstanceApiCopyResourceInstanceSnapshot(ctx, serviceProviderId, serviceKey, serviceAPIVersion, serviceEnvironmentKey, serviceModelKey, productTierKey, resourceKey, id).CopyResourceInstanceSnapshotRequest2(copyResourceInstanceSnapshotRequest2).SubscriptionId(subscriptionId).Execute()

CopyResourceInstanceSnapshot resource-instance-api

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
	serviceProviderId := "omnistrate" // string | The service provider ID
	serviceKey := "service-orchestration" // string | The service name
	serviceAPIVersion := "v1" // string | The service API version
	serviceEnvironmentKey := "dev" // string | The service environment name
	serviceModelKey := "hosted" // string | The service model name
	productTierKey := "premium" // string | The product tier name
	resourceKey := "mysql" // string | The resource key
	id := "instance-abcd1234" // string | The instance ID
	copyResourceInstanceSnapshotRequest2 := *openapiclient.NewCopyResourceInstanceSnapshotRequest2("us-west-2") // CopyResourceInstanceSnapshotRequest2 | 
	subscriptionId := "sub-abcd1234" // string | Subscription Id (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.ResourceInstanceApiAPI.ResourceInstanceApiCopyResourceInstanceSnapshot(context.Background(), serviceProviderId, serviceKey, serviceAPIVersion, serviceEnvironmentKey, serviceModelKey, productTierKey, resourceKey, id).CopyResourceInstanceSnapshotRequest2(copyResourceInstanceSnapshotRequest2).SubscriptionId(subscriptionId).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `ResourceInstanceApiAPI.ResourceInstanceApiCopyResourceInstanceSnapshot``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `ResourceInstanceApiCopyResourceInstanceSnapshot`: CopyResourceInstanceSnapshotResponseBody
	fmt.Fprintf(os.Stdout, "Response from `ResourceInstanceApiAPI.ResourceInstanceApiCopyResourceInstanceSnapshot`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**serviceProviderId** | **string** | The service provider ID | 
**serviceKey** | **string** | The service name | 
**serviceAPIVersion** | **string** | The service API version | 
**serviceEnvironmentKey** | **string** | The service environment name | 
**serviceModelKey** | **string** | The service model name | 
**productTierKey** | **string** | The product tier name | 
**resourceKey** | **string** | The resource key | 
**id** | **string** | The instance ID | 

### Other Parameters

Other parameters are passed through a pointer to a apiResourceInstanceApiCopyResourceInstanceSnapshotRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------








 **copyResourceInstanceSnapshotRequest2** | [**CopyResourceInstanceSnapshotRequest2**](CopyResourceInstanceSnapshotRequest2.md) |  | 
 **subscriptionId** | **string** | Subscription Id | 

### Return type

[**CopyResourceInstanceSnapshotResponseBody**](CopyResourceInstanceSnapshotResponseBody.md)

### Authorization

[api_key_header_Authorization](../README.md#api_key_header_Authorization)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json, application/vnd.goa.error

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## ResourceInstanceApiCreateResourceInstance

> CreateServicesOrchestrationResponseBody ResourceInstanceApiCreateResourceInstance(ctx, serviceProviderId, serviceKey, serviceAPIVersion, serviceEnvironmentKey, serviceModelKey, productTierKey, resourceKey).CreateResourceInstanceRequest2(createResourceInstanceRequest2).SubscriptionId(subscriptionId).Execute()

CreateResourceInstance resource-instance-api

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
	serviceProviderId := "omnistrate" // string | The service provider ID
	serviceKey := "service-orchestration" // string | The service name
	serviceAPIVersion := "v1" // string | The service API version
	serviceEnvironmentKey := "dev" // string | The service environment name
	serviceModelKey := "hosted" // string | The service model name
	productTierKey := "premium" // string | The product tier name
	resourceKey := "mysql" // string | The resource key
	createResourceInstanceRequest2 := *openapiclient.NewCreateResourceInstanceRequest2() // CreateResourceInstanceRequest2 | 
	subscriptionId := "sub-abcd1234" // string | Subscription Id (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.ResourceInstanceApiAPI.ResourceInstanceApiCreateResourceInstance(context.Background(), serviceProviderId, serviceKey, serviceAPIVersion, serviceEnvironmentKey, serviceModelKey, productTierKey, resourceKey).CreateResourceInstanceRequest2(createResourceInstanceRequest2).SubscriptionId(subscriptionId).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `ResourceInstanceApiAPI.ResourceInstanceApiCreateResourceInstance``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `ResourceInstanceApiCreateResourceInstance`: CreateServicesOrchestrationResponseBody
	fmt.Fprintf(os.Stdout, "Response from `ResourceInstanceApiAPI.ResourceInstanceApiCreateResourceInstance`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**serviceProviderId** | **string** | The service provider ID | 
**serviceKey** | **string** | The service name | 
**serviceAPIVersion** | **string** | The service API version | 
**serviceEnvironmentKey** | **string** | The service environment name | 
**serviceModelKey** | **string** | The service model name | 
**productTierKey** | **string** | The product tier name | 
**resourceKey** | **string** | The resource key | 

### Other Parameters

Other parameters are passed through a pointer to a apiResourceInstanceApiCreateResourceInstanceRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------







 **createResourceInstanceRequest2** | [**CreateResourceInstanceRequest2**](CreateResourceInstanceRequest2.md) |  | 
 **subscriptionId** | **string** | Subscription Id | 

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


## ResourceInstanceApiDeleteResourceInstance

> ResourceInstanceApiDeleteResourceInstance(ctx, serviceProviderId, serviceKey, serviceAPIVersion, serviceEnvironmentKey, serviceModelKey, productTierKey, resourceKey, id).SubscriptionId(subscriptionId).Execute()

DeleteResourceInstance resource-instance-api

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
	serviceProviderId := "omnistrate" // string | The service provider ID
	serviceKey := "service-orchestration" // string | The service name
	serviceAPIVersion := "v1" // string | The service API version
	serviceEnvironmentKey := "dev" // string | The service environment name
	serviceModelKey := "hosted" // string | The service model name
	productTierKey := "premium" // string | The product tier name
	resourceKey := "mysql" // string | The resource key
	id := "instance-abcd1234" // string | The instance ID
	subscriptionId := "sub-abcd1234" // string | Subscription Id (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	r, err := apiClient.ResourceInstanceApiAPI.ResourceInstanceApiDeleteResourceInstance(context.Background(), serviceProviderId, serviceKey, serviceAPIVersion, serviceEnvironmentKey, serviceModelKey, productTierKey, resourceKey, id).SubscriptionId(subscriptionId).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `ResourceInstanceApiAPI.ResourceInstanceApiDeleteResourceInstance``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**serviceProviderId** | **string** | The service provider ID | 
**serviceKey** | **string** | The service name | 
**serviceAPIVersion** | **string** | The service API version | 
**serviceEnvironmentKey** | **string** | The service environment name | 
**serviceModelKey** | **string** | The service model name | 
**productTierKey** | **string** | The product tier name | 
**resourceKey** | **string** | The resource key | 
**id** | **string** | The instance ID | 

### Other Parameters

Other parameters are passed through a pointer to a apiResourceInstanceApiDeleteResourceInstanceRequest struct via the builder pattern


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


## ResourceInstanceApiDescribeResourceInstance

> DescribeResourceInstanceResult ResourceInstanceApiDescribeResourceInstance(ctx, serviceProviderId, serviceKey, serviceAPIVersion, serviceEnvironmentKey, serviceModelKey, productTierKey, resourceKey, id).SubscriptionId(subscriptionId).Execute()

DescribeResourceInstance resource-instance-api

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
	serviceProviderId := "omnistrate" // string | The service provider ID
	serviceKey := "service-orchestration" // string | The service name
	serviceAPIVersion := "v1" // string | The service API version
	serviceEnvironmentKey := "dev" // string | The service environment name
	serviceModelKey := "hosted" // string | The service model name
	productTierKey := "premium" // string | The product tier name
	resourceKey := "mysql" // string | The resource key
	id := "instance-abcd1234" // string | The instance ID
	subscriptionId := "sub-abcd1234" // string | Subscription Id (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.ResourceInstanceApiAPI.ResourceInstanceApiDescribeResourceInstance(context.Background(), serviceProviderId, serviceKey, serviceAPIVersion, serviceEnvironmentKey, serviceModelKey, productTierKey, resourceKey, id).SubscriptionId(subscriptionId).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `ResourceInstanceApiAPI.ResourceInstanceApiDescribeResourceInstance``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `ResourceInstanceApiDescribeResourceInstance`: DescribeResourceInstanceResult
	fmt.Fprintf(os.Stdout, "Response from `ResourceInstanceApiAPI.ResourceInstanceApiDescribeResourceInstance`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**serviceProviderId** | **string** | The service provider ID | 
**serviceKey** | **string** | The service name | 
**serviceAPIVersion** | **string** | The service API version | 
**serviceEnvironmentKey** | **string** | The service environment name | 
**serviceModelKey** | **string** | The service model name | 
**productTierKey** | **string** | The product tier name | 
**resourceKey** | **string** | The resource key | 
**id** | **string** | The instance ID | 

### Other Parameters

Other parameters are passed through a pointer to a apiResourceInstanceApiDescribeResourceInstanceRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------








 **subscriptionId** | **string** | Subscription Id | 

### Return type

[**DescribeResourceInstanceResult**](DescribeResourceInstanceResult.md)

### Authorization

[api_key_header_Authorization](../README.md#api_key_header_Authorization)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json, application/vnd.goa.error

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## ResourceInstanceApiFailoverResourceInstance

> ResourceInstanceApiFailoverResourceInstance(ctx, serviceProviderId, serviceKey, serviceAPIVersion, serviceEnvironmentKey, serviceModelKey, productTierKey, resourceKey, id).FailoverResourceInstanceRequest2(failoverResourceInstanceRequest2).SubscriptionId(subscriptionId).Execute()

FailoverResourceInstance resource-instance-api

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
	serviceProviderId := "omnistrate" // string | The service provider ID
	serviceKey := "service-orchestration" // string | The service name
	serviceAPIVersion := "v1" // string | The service API version
	serviceEnvironmentKey := "dev" // string | The service environment name
	serviceModelKey := "hosted" // string | The service model name
	productTierKey := "premium" // string | The product tier name
	resourceKey := "mysql" // string | The resource key
	id := "instance-abcd1234" // string | The instance ID
	failoverResourceInstanceRequest2 := *openapiclient.NewFailoverResourceInstanceRequest2("db-0") // FailoverResourceInstanceRequest2 | 
	subscriptionId := "sub-abcd1234" // string | Subscription Id (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	r, err := apiClient.ResourceInstanceApiAPI.ResourceInstanceApiFailoverResourceInstance(context.Background(), serviceProviderId, serviceKey, serviceAPIVersion, serviceEnvironmentKey, serviceModelKey, productTierKey, resourceKey, id).FailoverResourceInstanceRequest2(failoverResourceInstanceRequest2).SubscriptionId(subscriptionId).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `ResourceInstanceApiAPI.ResourceInstanceApiFailoverResourceInstance``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**serviceProviderId** | **string** | The service provider ID | 
**serviceKey** | **string** | The service name | 
**serviceAPIVersion** | **string** | The service API version | 
**serviceEnvironmentKey** | **string** | The service environment name | 
**serviceModelKey** | **string** | The service model name | 
**productTierKey** | **string** | The product tier name | 
**resourceKey** | **string** | The resource key | 
**id** | **string** | The instance ID | 

### Other Parameters

Other parameters are passed through a pointer to a apiResourceInstanceApiFailoverResourceInstanceRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------








 **failoverResourceInstanceRequest2** | [**FailoverResourceInstanceRequest2**](FailoverResourceInstanceRequest2.md) |  | 
 **subscriptionId** | **string** | Subscription Id | 

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


## ResourceInstanceApiGenerateTokenForDeploymentCellDashboard

> GenerateTokenForDeploymentCellDashboardResult ResourceInstanceApiGenerateTokenForDeploymentCellDashboard(ctx, id).SubscriptionId(subscriptionId).Execute()

GenerateTokenForDeploymentCellDashboard resource-instance-api

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
	subscriptionId := "sub-abcd1234" // string | Subscription Id
	id := "instance-abcd1234" // string | The instance ID whose deployment cell dashboard token is to be generated

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.ResourceInstanceApiAPI.ResourceInstanceApiGenerateTokenForDeploymentCellDashboard(context.Background(), id).SubscriptionId(subscriptionId).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `ResourceInstanceApiAPI.ResourceInstanceApiGenerateTokenForDeploymentCellDashboard``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `ResourceInstanceApiGenerateTokenForDeploymentCellDashboard`: GenerateTokenForDeploymentCellDashboardResult
	fmt.Fprintf(os.Stdout, "Response from `ResourceInstanceApiAPI.ResourceInstanceApiGenerateTokenForDeploymentCellDashboard`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **string** | The instance ID whose deployment cell dashboard token is to be generated | 

### Other Parameters

Other parameters are passed through a pointer to a apiResourceInstanceApiGenerateTokenForDeploymentCellDashboardRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **subscriptionId** | **string** | Subscription Id | 


### Return type

[**GenerateTokenForDeploymentCellDashboardResult**](GenerateTokenForDeploymentCellDashboardResult.md)

### Authorization

[api_key_header_Authorization](../README.md#api_key_header_Authorization)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json, application/vnd.goa.error

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## ResourceInstanceApiListAllResourceInstances

> ListAllResourceInstancesResult ResourceInstanceApiListAllResourceInstances(ctx).EnvironmentType(environmentType).Execute()

ListAllResourceInstances resource-instance-api

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

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.ResourceInstanceApiAPI.ResourceInstanceApiListAllResourceInstances(context.Background()).EnvironmentType(environmentType).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `ResourceInstanceApiAPI.ResourceInstanceApiListAllResourceInstances``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `ResourceInstanceApiListAllResourceInstances`: ListAllResourceInstancesResult
	fmt.Fprintf(os.Stdout, "Response from `ResourceInstanceApiAPI.ResourceInstanceApiListAllResourceInstances`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiResourceInstanceApiListAllResourceInstancesRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **environmentType** | **string** | The environment type to filter by | 

### Return type

[**ListAllResourceInstancesResult**](ListAllResourceInstancesResult.md)

### Authorization

[api_key_header_Authorization](../README.md#api_key_header_Authorization)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json, application/vnd.goa.error

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## ResourceInstanceApiListResourceInstanceSnapshots

> ListResourceInstanceSnapshotsResult ResourceInstanceApiListResourceInstanceSnapshots(ctx, serviceProviderId, serviceKey, serviceAPIVersion, serviceEnvironmentKey, serviceModelKey, productTierKey, resourceKey, id).SubscriptionId(subscriptionId).Execute()

ListResourceInstanceSnapshots resource-instance-api

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
	serviceProviderId := "omnistrate" // string | The service provider ID
	serviceKey := "service-orchestration" // string | The service key
	serviceAPIVersion := "v1" // string | The service API version
	serviceEnvironmentKey := "dev" // string | The service environment name
	serviceModelKey := "hosted" // string | The service model name
	productTierKey := "premium" // string | The product tier name
	resourceKey := "mysql" // string | The resource key
	id := "instance-abcd1234" // string | The instance ID
	subscriptionId := "sub-abcd1234" // string | Subscription Id (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.ResourceInstanceApiAPI.ResourceInstanceApiListResourceInstanceSnapshots(context.Background(), serviceProviderId, serviceKey, serviceAPIVersion, serviceEnvironmentKey, serviceModelKey, productTierKey, resourceKey, id).SubscriptionId(subscriptionId).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `ResourceInstanceApiAPI.ResourceInstanceApiListResourceInstanceSnapshots``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `ResourceInstanceApiListResourceInstanceSnapshots`: ListResourceInstanceSnapshotsResult
	fmt.Fprintf(os.Stdout, "Response from `ResourceInstanceApiAPI.ResourceInstanceApiListResourceInstanceSnapshots`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**serviceProviderId** | **string** | The service provider ID | 
**serviceKey** | **string** | The service key | 
**serviceAPIVersion** | **string** | The service API version | 
**serviceEnvironmentKey** | **string** | The service environment name | 
**serviceModelKey** | **string** | The service model name | 
**productTierKey** | **string** | The product tier name | 
**resourceKey** | **string** | The resource key | 
**id** | **string** | The instance ID | 

### Other Parameters

Other parameters are passed through a pointer to a apiResourceInstanceApiListResourceInstanceSnapshotsRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------








 **subscriptionId** | **string** | Subscription Id | 

### Return type

[**ListResourceInstanceSnapshotsResult**](ListResourceInstanceSnapshotsResult.md)

### Authorization

[api_key_header_Authorization](../README.md#api_key_header_Authorization)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json, application/vnd.goa.error

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## ResourceInstanceApiListResourceInstances

> ListResourceInstancesResult ResourceInstanceApiListResourceInstances(ctx, serviceProviderId, serviceKey, serviceAPIVersion, serviceEnvironmentKey, serviceModelKey, productTierKey, resourceKey).SubscriptionId(subscriptionId).Execute()

ListResourceInstances resource-instance-api

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
	serviceProviderId := "omnistrate" // string | The service provider ID
	serviceKey := "service-orchestration" // string | The service key
	serviceAPIVersion := "v1" // string | The service API version
	serviceEnvironmentKey := "dev" // string | The service environment name
	serviceModelKey := "hosted" // string | The service model name
	productTierKey := "premium" // string | The product tier name
	resourceKey := "mysql" // string | The resource key
	subscriptionId := "sub-abcd1234" // string | Subscription Id (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.ResourceInstanceApiAPI.ResourceInstanceApiListResourceInstances(context.Background(), serviceProviderId, serviceKey, serviceAPIVersion, serviceEnvironmentKey, serviceModelKey, productTierKey, resourceKey).SubscriptionId(subscriptionId).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `ResourceInstanceApiAPI.ResourceInstanceApiListResourceInstances``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `ResourceInstanceApiListResourceInstances`: ListResourceInstancesResult
	fmt.Fprintf(os.Stdout, "Response from `ResourceInstanceApiAPI.ResourceInstanceApiListResourceInstances`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**serviceProviderId** | **string** | The service provider ID | 
**serviceKey** | **string** | The service key | 
**serviceAPIVersion** | **string** | The service API version | 
**serviceEnvironmentKey** | **string** | The service environment name | 
**serviceModelKey** | **string** | The service model name | 
**productTierKey** | **string** | The product tier name | 
**resourceKey** | **string** | The resource key | 

### Other Parameters

Other parameters are passed through a pointer to a apiResourceInstanceApiListResourceInstancesRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------







 **subscriptionId** | **string** | Subscription Id | 

### Return type

[**ListResourceInstancesResult**](ListResourceInstancesResult.md)

### Authorization

[api_key_header_Authorization](../README.md#api_key_header_Authorization)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json, application/vnd.goa.error

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## ResourceInstanceApiRemoveCapacityFromResourceInstance

> ResourceInstanceApiRemoveCapacityFromResourceInstance(ctx, serviceProviderId, serviceKey, serviceAPIVersion, serviceEnvironmentKey, serviceModelKey, productTierKey, resourceKey, id).RemoveCapacityFromResourceInstanceRequest2(removeCapacityFromResourceInstanceRequest2).SubscriptionId(subscriptionId).Execute()

RemoveCapacityFromResourceInstance resource-instance-api

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
	serviceProviderId := "omnistrate" // string | The service provider ID
	serviceKey := "service-orchestration" // string | The service name
	serviceAPIVersion := "v1" // string | The service API version
	serviceEnvironmentKey := "dev" // string | The service environment name
	serviceModelKey := "hosted" // string | The service model name
	productTierKey := "premium" // string | The product tier name
	resourceKey := "mysql" // string | The resource key
	id := "instance-abcd1234" // string | The instance ID
	removeCapacityFromResourceInstanceRequest2 := *openapiclient.NewRemoveCapacityFromResourceInstanceRequest2(int64(3)) // RemoveCapacityFromResourceInstanceRequest2 | 
	subscriptionId := "sub-abcd1234" // string | Subscription Id (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	r, err := apiClient.ResourceInstanceApiAPI.ResourceInstanceApiRemoveCapacityFromResourceInstance(context.Background(), serviceProviderId, serviceKey, serviceAPIVersion, serviceEnvironmentKey, serviceModelKey, productTierKey, resourceKey, id).RemoveCapacityFromResourceInstanceRequest2(removeCapacityFromResourceInstanceRequest2).SubscriptionId(subscriptionId).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `ResourceInstanceApiAPI.ResourceInstanceApiRemoveCapacityFromResourceInstance``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**serviceProviderId** | **string** | The service provider ID | 
**serviceKey** | **string** | The service name | 
**serviceAPIVersion** | **string** | The service API version | 
**serviceEnvironmentKey** | **string** | The service environment name | 
**serviceModelKey** | **string** | The service model name | 
**productTierKey** | **string** | The product tier name | 
**resourceKey** | **string** | The resource key | 
**id** | **string** | The instance ID | 

### Other Parameters

Other parameters are passed through a pointer to a apiResourceInstanceApiRemoveCapacityFromResourceInstanceRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------








 **removeCapacityFromResourceInstanceRequest2** | [**RemoveCapacityFromResourceInstanceRequest2**](RemoveCapacityFromResourceInstanceRequest2.md) |  | 
 **subscriptionId** | **string** | Subscription Id | 

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


## ResourceInstanceApiRemoveCustomDNSFromResourceInstance

> ResourceInstanceApiRemoveCustomDNSFromResourceInstance(ctx, serviceProviderId, serviceKey, serviceAPIVersion, serviceEnvironmentKey, serviceModelKey, productTierKey, resourceKey, id).SubscriptionId(subscriptionId).Execute()

RemoveCustomDNSFromResourceInstance resource-instance-api

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
	serviceProviderId := "omnistrate" // string | The service provider ID
	serviceKey := "service-orchestration" // string | The service name
	serviceAPIVersion := "v1" // string | The service API version
	serviceEnvironmentKey := "dev" // string | The service environment name
	serviceModelKey := "hosted" // string | The service model name
	productTierKey := "premium" // string | The product tier name
	resourceKey := "http-service" // string | The resource key
	id := "instance-abcd1234" // string | The instance ID
	subscriptionId := "sub-abcd1234" // string | Subscription Id (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	r, err := apiClient.ResourceInstanceApiAPI.ResourceInstanceApiRemoveCustomDNSFromResourceInstance(context.Background(), serviceProviderId, serviceKey, serviceAPIVersion, serviceEnvironmentKey, serviceModelKey, productTierKey, resourceKey, id).SubscriptionId(subscriptionId).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `ResourceInstanceApiAPI.ResourceInstanceApiRemoveCustomDNSFromResourceInstance``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**serviceProviderId** | **string** | The service provider ID | 
**serviceKey** | **string** | The service name | 
**serviceAPIVersion** | **string** | The service API version | 
**serviceEnvironmentKey** | **string** | The service environment name | 
**serviceModelKey** | **string** | The service model name | 
**productTierKey** | **string** | The product tier name | 
**resourceKey** | **string** | The resource key | 
**id** | **string** | The instance ID | 

### Other Parameters

Other parameters are passed through a pointer to a apiResourceInstanceApiRemoveCustomDNSFromResourceInstanceRequest struct via the builder pattern


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


## ResourceInstanceApiResourceInstanceProvisionerSetupKit

> ResourceInstanceApiResourceInstanceProvisionerSetupKit(ctx, serviceProviderId, serviceKey, serviceAPIVersion, serviceEnvironmentKey, serviceModelKey).SubscriptionId(subscriptionId).Execute()

ResourceInstanceProvisionerSetupKit resource-instance-api

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
	serviceProviderId := "omnistrate" // string | The service provider ID
	serviceKey := "service-orchestration" // string | The service name
	serviceAPIVersion := "v1" // string | The service API version
	serviceEnvironmentKey := "dev" // string | The service environment name
	serviceModelKey := "hosted" // string | The service model name
	subscriptionId := "sub-abcd1234" // string | Subscription Id (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	r, err := apiClient.ResourceInstanceApiAPI.ResourceInstanceApiResourceInstanceProvisionerSetupKit(context.Background(), serviceProviderId, serviceKey, serviceAPIVersion, serviceEnvironmentKey, serviceModelKey).SubscriptionId(subscriptionId).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `ResourceInstanceApiAPI.ResourceInstanceApiResourceInstanceProvisionerSetupKit``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**serviceProviderId** | **string** | The service provider ID | 
**serviceKey** | **string** | The service name | 
**serviceAPIVersion** | **string** | The service API version | 
**serviceEnvironmentKey** | **string** | The service environment name | 
**serviceModelKey** | **string** | The service model name | 

### Other Parameters

Other parameters are passed through a pointer to a apiResourceInstanceApiResourceInstanceProvisionerSetupKitRequest struct via the builder pattern


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


## ResourceInstanceApiRestartResourceInstance

> ResourceInstanceApiRestartResourceInstance(ctx, serviceProviderId, serviceKey, serviceAPIVersion, serviceEnvironmentKey, serviceModelKey, productTierKey, resourceKey, id).SubscriptionId(subscriptionId).Execute()

RestartResourceInstance resource-instance-api

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
	serviceProviderId := "omnistrate" // string | The service provider ID
	serviceKey := "service-orchestration" // string | The service name
	serviceAPIVersion := "v1" // string | The service API version
	serviceEnvironmentKey := "dev" // string | The service environment name
	serviceModelKey := "hosted" // string | The service model name
	productTierKey := "premium" // string | The product tier name
	resourceKey := "mysql" // string | The resource key
	id := "instance-abcd1234" // string | The instance ID
	subscriptionId := "sub-abcd1234" // string | Subscription Id (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	r, err := apiClient.ResourceInstanceApiAPI.ResourceInstanceApiRestartResourceInstance(context.Background(), serviceProviderId, serviceKey, serviceAPIVersion, serviceEnvironmentKey, serviceModelKey, productTierKey, resourceKey, id).SubscriptionId(subscriptionId).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `ResourceInstanceApiAPI.ResourceInstanceApiRestartResourceInstance``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**serviceProviderId** | **string** | The service provider ID | 
**serviceKey** | **string** | The service name | 
**serviceAPIVersion** | **string** | The service API version | 
**serviceEnvironmentKey** | **string** | The service environment name | 
**serviceModelKey** | **string** | The service model name | 
**productTierKey** | **string** | The product tier name | 
**resourceKey** | **string** | The resource key | 
**id** | **string** | The instance ID | 

### Other Parameters

Other parameters are passed through a pointer to a apiResourceInstanceApiRestartResourceInstanceRequest struct via the builder pattern


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


## ResourceInstanceApiRestoreResourceInstance

> CreateServicesOrchestrationResponseBody ResourceInstanceApiRestoreResourceInstance(ctx, serviceProviderId, serviceKey, serviceAPIVersion, serviceEnvironmentKey, serviceModelKey, productTierKey, resourceKey, id).RestoreResourceInstanceRequest2(restoreResourceInstanceRequest2).SubscriptionId(subscriptionId).Execute()

RestoreResourceInstance resource-instance-api

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
	serviceProviderId := "omnistrate" // string | The service provider ID
	serviceKey := "service-orchestration" // string | The service name
	serviceAPIVersion := "v1" // string | The service API version
	serviceEnvironmentKey := "dev" // string | The service environment name
	serviceModelKey := "hosted" // string | The service model name
	productTierKey := "premium" // string | The product tier name
	resourceKey := "mysql" // string | The resource key
	id := "instance-abcd1234" // string | The instance ID
	restoreResourceInstanceRequest2 := *openapiclient.NewRestoreResourceInstanceRequest2("2021-09-01T00:00:00Z") // RestoreResourceInstanceRequest2 | 
	subscriptionId := "sub-abcd1234" // string | Subscription Id (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.ResourceInstanceApiAPI.ResourceInstanceApiRestoreResourceInstance(context.Background(), serviceProviderId, serviceKey, serviceAPIVersion, serviceEnvironmentKey, serviceModelKey, productTierKey, resourceKey, id).RestoreResourceInstanceRequest2(restoreResourceInstanceRequest2).SubscriptionId(subscriptionId).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `ResourceInstanceApiAPI.ResourceInstanceApiRestoreResourceInstance``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `ResourceInstanceApiRestoreResourceInstance`: CreateServicesOrchestrationResponseBody
	fmt.Fprintf(os.Stdout, "Response from `ResourceInstanceApiAPI.ResourceInstanceApiRestoreResourceInstance`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**serviceProviderId** | **string** | The service provider ID | 
**serviceKey** | **string** | The service name | 
**serviceAPIVersion** | **string** | The service API version | 
**serviceEnvironmentKey** | **string** | The service environment name | 
**serviceModelKey** | **string** | The service model name | 
**productTierKey** | **string** | The product tier name | 
**resourceKey** | **string** | The resource key | 
**id** | **string** | The instance ID | 

### Other Parameters

Other parameters are passed through a pointer to a apiResourceInstanceApiRestoreResourceInstanceRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------








 **restoreResourceInstanceRequest2** | [**RestoreResourceInstanceRequest2**](RestoreResourceInstanceRequest2.md) |  | 
 **subscriptionId** | **string** | Subscription Id | 

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


## ResourceInstanceApiRestoreResourceInstanceFromSnapshot

> CreateServicesOrchestrationResponseBody ResourceInstanceApiRestoreResourceInstanceFromSnapshot(ctx, serviceProviderId, serviceKey, serviceAPIVersion, serviceEnvironmentKey, serviceModelKey, productTierKey, resourceKey, snapshotId).RestoreResourceInstanceFromSnapshotRequest2(restoreResourceInstanceFromSnapshotRequest2).SubscriptionId(subscriptionId).Execute()

RestoreResourceInstanceFromSnapshot resource-instance-api

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
	serviceProviderId := "omnistrate" // string | The service provider ID
	serviceKey := "service-orchestration" // string | The service name
	serviceAPIVersion := "v1" // string | The service API version
	serviceEnvironmentKey := "dev" // string | The service environment name
	serviceModelKey := "hosted" // string | The service model name
	productTierKey := "premium" // string | The product tier name
	resourceKey := "mysql" // string | The resource key
	snapshotId := "instance-ss-abcd1234" // string | The snapshot ID
	restoreResourceInstanceFromSnapshotRequest2 := *openapiclient.NewRestoreResourceInstanceFromSnapshotRequest2() // RestoreResourceInstanceFromSnapshotRequest2 | 
	subscriptionId := "sub-abcd1234" // string | Subscription Id (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.ResourceInstanceApiAPI.ResourceInstanceApiRestoreResourceInstanceFromSnapshot(context.Background(), serviceProviderId, serviceKey, serviceAPIVersion, serviceEnvironmentKey, serviceModelKey, productTierKey, resourceKey, snapshotId).RestoreResourceInstanceFromSnapshotRequest2(restoreResourceInstanceFromSnapshotRequest2).SubscriptionId(subscriptionId).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `ResourceInstanceApiAPI.ResourceInstanceApiRestoreResourceInstanceFromSnapshot``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `ResourceInstanceApiRestoreResourceInstanceFromSnapshot`: CreateServicesOrchestrationResponseBody
	fmt.Fprintf(os.Stdout, "Response from `ResourceInstanceApiAPI.ResourceInstanceApiRestoreResourceInstanceFromSnapshot`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**serviceProviderId** | **string** | The service provider ID | 
**serviceKey** | **string** | The service name | 
**serviceAPIVersion** | **string** | The service API version | 
**serviceEnvironmentKey** | **string** | The service environment name | 
**serviceModelKey** | **string** | The service model name | 
**productTierKey** | **string** | The product tier name | 
**resourceKey** | **string** | The resource key | 
**snapshotId** | **string** | The snapshot ID | 

### Other Parameters

Other parameters are passed through a pointer to a apiResourceInstanceApiRestoreResourceInstanceFromSnapshotRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------








 **restoreResourceInstanceFromSnapshotRequest2** | [**RestoreResourceInstanceFromSnapshotRequest2**](RestoreResourceInstanceFromSnapshotRequest2.md) |  | 
 **subscriptionId** | **string** | Subscription Id | 

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


## ResourceInstanceApiStartResourceInstance

> ResourceInstanceApiStartResourceInstance(ctx, serviceProviderId, serviceKey, serviceAPIVersion, serviceEnvironmentKey, serviceModelKey, productTierKey, resourceKey, id).SubscriptionId(subscriptionId).Execute()

StartResourceInstance resource-instance-api

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
	serviceProviderId := "omnistrate" // string | The service provider ID
	serviceKey := "service-orchestration" // string | The service name
	serviceAPIVersion := "v1" // string | The service API version
	serviceEnvironmentKey := "dev" // string | The service environment name
	serviceModelKey := "hosted" // string | The service model name
	productTierKey := "premium" // string | The product tier name
	resourceKey := "mysql" // string | The resource key
	id := "instance-abcd1234" // string | The instance ID
	subscriptionId := "sub-abcd1234" // string | Subscription Id (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	r, err := apiClient.ResourceInstanceApiAPI.ResourceInstanceApiStartResourceInstance(context.Background(), serviceProviderId, serviceKey, serviceAPIVersion, serviceEnvironmentKey, serviceModelKey, productTierKey, resourceKey, id).SubscriptionId(subscriptionId).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `ResourceInstanceApiAPI.ResourceInstanceApiStartResourceInstance``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**serviceProviderId** | **string** | The service provider ID | 
**serviceKey** | **string** | The service name | 
**serviceAPIVersion** | **string** | The service API version | 
**serviceEnvironmentKey** | **string** | The service environment name | 
**serviceModelKey** | **string** | The service model name | 
**productTierKey** | **string** | The product tier name | 
**resourceKey** | **string** | The resource key | 
**id** | **string** | The instance ID | 

### Other Parameters

Other parameters are passed through a pointer to a apiResourceInstanceApiStartResourceInstanceRequest struct via the builder pattern


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


## ResourceInstanceApiStopResourceInstance

> ResourceInstanceApiStopResourceInstance(ctx, serviceProviderId, serviceKey, serviceAPIVersion, serviceEnvironmentKey, serviceModelKey, productTierKey, resourceKey, id).SubscriptionId(subscriptionId).Execute()

StopResourceInstance resource-instance-api

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
	serviceProviderId := "omnistrate" // string | The service provider ID
	serviceKey := "service-orchestration" // string | The service name
	serviceAPIVersion := "v1" // string | The service API version
	serviceEnvironmentKey := "dev" // string | The service environment name
	serviceModelKey := "hosted" // string | The service model name
	productTierKey := "premium" // string | The product tier name
	resourceKey := "mysql" // string | The resource key
	id := "instance-abcd1234" // string | The instance ID
	subscriptionId := "sub-abcd1234" // string | Subscription Id (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	r, err := apiClient.ResourceInstanceApiAPI.ResourceInstanceApiStopResourceInstance(context.Background(), serviceProviderId, serviceKey, serviceAPIVersion, serviceEnvironmentKey, serviceModelKey, productTierKey, resourceKey, id).SubscriptionId(subscriptionId).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `ResourceInstanceApiAPI.ResourceInstanceApiStopResourceInstance``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**serviceProviderId** | **string** | The service provider ID | 
**serviceKey** | **string** | The service name | 
**serviceAPIVersion** | **string** | The service API version | 
**serviceEnvironmentKey** | **string** | The service environment name | 
**serviceModelKey** | **string** | The service model name | 
**productTierKey** | **string** | The product tier name | 
**resourceKey** | **string** | The resource key | 
**id** | **string** | The instance ID | 

### Other Parameters

Other parameters are passed through a pointer to a apiResourceInstanceApiStopResourceInstanceRequest struct via the builder pattern


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


## ResourceInstanceApiUpdateAccountConfigResourceInstance

> ResourceInstanceApiUpdateAccountConfigResourceInstance(ctx, id).UpdateAccountConfigResourceInstanceRequest2(updateAccountConfigResourceInstanceRequest2).Execute()

UpdateAccountConfigResourceInstance resource-instance-api

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
	id := "instance-abcd1234" // string | The instance ID
	updateAccountConfigResourceInstanceRequest2 := *openapiclient.NewUpdateAccountConfigResourceInstanceRequest2("service-1234", "sub-abcd1234") // UpdateAccountConfigResourceInstanceRequest2 | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	r, err := apiClient.ResourceInstanceApiAPI.ResourceInstanceApiUpdateAccountConfigResourceInstance(context.Background(), id).UpdateAccountConfigResourceInstanceRequest2(updateAccountConfigResourceInstanceRequest2).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `ResourceInstanceApiAPI.ResourceInstanceApiUpdateAccountConfigResourceInstance``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **string** | The instance ID | 

### Other Parameters

Other parameters are passed through a pointer to a apiResourceInstanceApiUpdateAccountConfigResourceInstanceRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **updateAccountConfigResourceInstanceRequest2** | [**UpdateAccountConfigResourceInstanceRequest2**](UpdateAccountConfigResourceInstanceRequest2.md) |  | 

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


## ResourceInstanceApiUpdateResourceInstance

> ResourceInstanceApiUpdateResourceInstance(ctx, serviceProviderId, serviceKey, serviceAPIVersion, serviceEnvironmentKey, serviceModelKey, productTierKey, resourceKey, id).UpdateResourceInstanceRequest2(updateResourceInstanceRequest2).SubscriptionId(subscriptionId).Execute()

UpdateResourceInstance resource-instance-api

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
	serviceProviderId := "omnistrate" // string | The service provider ID
	serviceKey := "service-orchestration" // string | The service name
	serviceAPIVersion := "v1" // string | The service API version
	serviceEnvironmentKey := "dev" // string | The service environment name
	serviceModelKey := "hosted" // string | The service model name
	productTierKey := "premium" // string | The product tier name
	resourceKey := "mysql" // string | The resource key
	id := "instance-abcd1234" // string | The instance ID
	updateResourceInstanceRequest2 := *openapiclient.NewUpdateResourceInstanceRequest2() // UpdateResourceInstanceRequest2 | 
	subscriptionId := "sub-abcd1234" // string | Subscription Id (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	r, err := apiClient.ResourceInstanceApiAPI.ResourceInstanceApiUpdateResourceInstance(context.Background(), serviceProviderId, serviceKey, serviceAPIVersion, serviceEnvironmentKey, serviceModelKey, productTierKey, resourceKey, id).UpdateResourceInstanceRequest2(updateResourceInstanceRequest2).SubscriptionId(subscriptionId).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `ResourceInstanceApiAPI.ResourceInstanceApiUpdateResourceInstance``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**serviceProviderId** | **string** | The service provider ID | 
**serviceKey** | **string** | The service name | 
**serviceAPIVersion** | **string** | The service API version | 
**serviceEnvironmentKey** | **string** | The service environment name | 
**serviceModelKey** | **string** | The service model name | 
**productTierKey** | **string** | The product tier name | 
**resourceKey** | **string** | The resource key | 
**id** | **string** | The instance ID | 

### Other Parameters

Other parameters are passed through a pointer to a apiResourceInstanceApiUpdateResourceInstanceRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------








 **updateResourceInstanceRequest2** | [**UpdateResourceInstanceRequest2**](UpdateResourceInstanceRequest2.md) |  | 
 **subscriptionId** | **string** | Subscription Id | 

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


## ResourceInstanceApiUpgradeResourceInstanceVersion

> ResourceInstanceApiUpgradeResourceInstanceVersion(ctx, id).UpgradeResourceInstanceVersionRequest2(upgradeResourceInstanceVersionRequest2).SubscriptionId(subscriptionId).Execute()

UpgradeResourceInstanceVersion resource-instance-api

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
	id := "instance-abcd1234" // string | The instance ID
	upgradeResourceInstanceVersionRequest2 := *openapiclient.NewUpgradeResourceInstanceVersionRequest2("premium", "mysql", "v1", "dev", "service-orchestration", "hosted", "Ab alias.") // UpgradeResourceInstanceVersionRequest2 | 
	subscriptionId := "sub-abcd1234" // string | Subscription Id (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	r, err := apiClient.ResourceInstanceApiAPI.ResourceInstanceApiUpgradeResourceInstanceVersion(context.Background(), id).UpgradeResourceInstanceVersionRequest2(upgradeResourceInstanceVersionRequest2).SubscriptionId(subscriptionId).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `ResourceInstanceApiAPI.ResourceInstanceApiUpgradeResourceInstanceVersion``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **string** | The instance ID | 

### Other Parameters

Other parameters are passed through a pointer to a apiResourceInstanceApiUpgradeResourceInstanceVersionRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **upgradeResourceInstanceVersionRequest2** | [**UpgradeResourceInstanceVersionRequest2**](UpgradeResourceInstanceVersionRequest2.md) |  | 
 **subscriptionId** | **string** | Subscription Id | 

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

