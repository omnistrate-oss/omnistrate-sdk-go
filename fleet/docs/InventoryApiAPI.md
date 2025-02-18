# \InventoryApiAPI

All URIs are relative to *https://api.omnistrate.cloud*

Method | HTTP request | Description
------------- | ------------- | -------------
[**InventoryApiAddCapacityToResourceInstance**](InventoryApiAPI.md#InventoryApiAddCapacityToResourceInstance) | **Post** /2022-09-01-00/fleet/service/{serviceId}/environment/{environmentId}/instance/{instanceId}/add-capacity | AddCapacityToResourceInstance inventory-api
[**InventoryApiAddCustomDNSToResourceInstance**](InventoryApiAPI.md#InventoryApiAddCustomDNSToResourceInstance) | **Post** /2022-09-01-00/fleet/service/{serviceId}/environment/{environmentId}/{resourceKey}/instance/{instanceId}/custom-dns | AddCustomDNSToResourceInstance inventory-api
[**InventoryApiApproveSubscriptionRequest**](InventoryApiAPI.md#InventoryApiApproveSubscriptionRequest) | **Put** /2022-09-01-00/fleet/service/{serviceId}/environment/{environmentId}/subscription/request/{id} | ApproveSubscriptionRequest inventory-api
[**InventoryApiCancelUpgradePath**](InventoryApiAPI.md#InventoryApiCancelUpgradePath) | **Post** /2022-09-01-00/fleet/service/{serviceId}/productTier/{productTierId}/upgrade-path/{upgradePathId}/cancel | CancelUpgradePath inventory-api
[**InventoryApiCreateConsumptionUser**](InventoryApiAPI.md#InventoryApiCreateConsumptionUser) | **Post** /2022-09-01-00/fleet/user | CreateConsumptionUser inventory-api
[**InventoryApiCreateProxyResourceInstance**](InventoryApiAPI.md#InventoryApiCreateProxyResourceInstance) | **Post** /2022-09-01-00/fleet/proxy-resource-instance/{serviceProviderId}/{serviceKey}/{serviceAPIVersion}/{serviceEnvironmentKey}/{serviceModelKey}/{productTierKey}/{sourceResourceKey} | CreateProxyResourceInstance inventory-api
[**InventoryApiCreateResourceInstance**](InventoryApiAPI.md#InventoryApiCreateResourceInstance) | **Post** /2022-09-01-00/fleet/resource-instance/{serviceProviderId}/{serviceKey}/{serviceAPIVersion}/{serviceEnvironmentKey}/{serviceModelKey}/{productTierKey}/{resourceKey} | CreateResourceInstance inventory-api
[**InventoryApiCreateResourceInstanceSnapshot**](InventoryApiAPI.md#InventoryApiCreateResourceInstanceSnapshot) | **Post** /2022-09-01-00/fleet/service/{serviceId}/environment/{environmentId}/instance/{instanceId}/snapshot | CreateResourceInstanceSnapshot inventory-api
[**InventoryApiCreateServicesOrchestration**](InventoryApiAPI.md#InventoryApiCreateServicesOrchestration) | **Post** /2022-09-01-00/fleet/services-orchestration | CreateServicesOrchestration inventory-api
[**InventoryApiCreateUpgradePath**](InventoryApiAPI.md#InventoryApiCreateUpgradePath) | **Post** /2022-09-01-00/fleet/service/{serviceId}/productTier/{productTierId}/upgrade-path | CreateUpgradePath inventory-api
[**InventoryApiDebugResourceInstance**](InventoryApiAPI.md#InventoryApiDebugResourceInstance) | **Get** /2022-09-01-00/fleet/service/{serviceId}/environment/{environmentId}/instance/{instanceId}/debug | DebugResourceInstance inventory-api
[**InventoryApiDeleteProxyResourceInstance**](InventoryApiAPI.md#InventoryApiDeleteProxyResourceInstance) | **Delete** /2022-09-01-00/fleet/proxy-resource-instance/{serviceProviderId}/{serviceKey}/{serviceAPIVersion}/{serviceEnvironmentKey}/{serviceModelKey}/{productTierKey}/{id} | DeleteProxyResourceInstance inventory-api
[**InventoryApiDeleteResourceInstance**](InventoryApiAPI.md#InventoryApiDeleteResourceInstance) | **Delete** /2022-09-01-00/fleet/service/{serviceId}/environment/{environmentId}/instance/{instanceId} | DeleteResourceInstance inventory-api
[**InventoryApiDeleteResourceInstanceSnapshot**](InventoryApiAPI.md#InventoryApiDeleteResourceInstanceSnapshot) | **Delete** /2022-09-01-00/fleet/service/{serviceId}/environment/{environmentId}/instance/{instanceId}/snapshot/{snapshotId} | DeleteResourceInstanceSnapshot inventory-api
[**InventoryApiDeleteServicesOrchestration**](InventoryApiAPI.md#InventoryApiDeleteServicesOrchestration) | **Delete** /2022-09-01-00/fleet/services-orchestration/{id} | DeleteServicesOrchestration inventory-api
[**InventoryApiDeleteUser**](InventoryApiAPI.md#InventoryApiDeleteUser) | **Delete** /2022-09-01-00/fleet/user/{userId} | DeleteUser inventory-api
[**InventoryApiDenySubscriptionRequest**](InventoryApiAPI.md#InventoryApiDenySubscriptionRequest) | **Delete** /2022-09-01-00/fleet/service/{serviceId}/environment/{environmentId}/subscription/request/{id} | DenySubscriptionRequest inventory-api
[**InventoryApiDescribeHostCluster**](InventoryApiAPI.md#InventoryApiDescribeHostCluster) | **Get** /2022-09-01-00/fleet/service/{serviceId}/environment/{environmentId}/host-cluster/{id} | DescribeHostCluster inventory-api
[**InventoryApiDescribeInstanceEvent**](InventoryApiAPI.md#InventoryApiDescribeInstanceEvent) | **Get** /2022-09-01-00/fleet/service/{serviceId}/environment/{environmentId}/instance/{instanceId}/event/{id} | DescribeInstanceEvent inventory-api
[**InventoryApiDescribeInventorySummary**](InventoryApiAPI.md#InventoryApiDescribeInventorySummary) | **Get** /2022-09-01-00/fleet/service/{serviceId}/environment/{environmentId}/service-inventory-summary | DescribeInventorySummary inventory-api
[**InventoryApiDescribeOrgUser**](InventoryApiAPI.md#InventoryApiDescribeOrgUser) | **Get** /2022-09-01-00/fleet/user/{userId} | DescribeOrgUser inventory-api
[**InventoryApiDescribeOrganization**](InventoryApiAPI.md#InventoryApiDescribeOrganization) | **Get** /2022-09-01-00/fleet/service/{serviceId}/environment/{environmentId}/organization/{organizationId} | DescribeOrganization inventory-api
[**InventoryApiDescribeResource**](InventoryApiAPI.md#InventoryApiDescribeResource) | **Get** /2022-09-01-00/fleet/service/{serviceId}/environment/{environmentId}/model/{serviceModelId}/productTier/{productTierId}/resource/{resourceId} | DescribeResource inventory-api
[**InventoryApiDescribeResourceInstance**](InventoryApiAPI.md#InventoryApiDescribeResourceInstance) | **Get** /2022-09-01-00/fleet/service/{serviceId}/environment/{environmentId}/instance/{instanceId} | DescribeResourceInstance inventory-api
[**InventoryApiDescribeResourceInstanceSnapshotFromTime**](InventoryApiAPI.md#InventoryApiDescribeResourceInstanceSnapshotFromTime) | **Get** /2022-09-01-00/fleet/service/{serviceId}/environment/{environmentId}/instance/{instanceId}/restore | DescribeResourceInstanceSnapshotFromTime inventory-api
[**InventoryApiDescribeServiceOffering**](InventoryApiAPI.md#InventoryApiDescribeServiceOffering) | **Get** /2022-09-01-00/fleet/service-offering/{serviceId} | DescribeServiceOffering inventory-api
[**InventoryApiDescribeServiceOfferingResource**](InventoryApiAPI.md#InventoryApiDescribeServiceOfferingResource) | **Get** /2022-09-01-00/fleet/service-offering/{serviceId}/resource/{resourceId}/instance/{instanceId} | DescribeServiceOfferingResource inventory-api
[**InventoryApiDescribeServicesOrchestration**](InventoryApiAPI.md#InventoryApiDescribeServicesOrchestration) | **Get** /2022-09-01-00/fleet/services-orchestration/{id} | DescribeServicesOrchestration inventory-api
[**InventoryApiDescribeSubscription**](InventoryApiAPI.md#InventoryApiDescribeSubscription) | **Get** /2022-09-01-00/fleet/service/{serviceId}/environment/{environmentId}/subscription/{id} | DescribeSubscription inventory-api
[**InventoryApiDescribeSubscriptionRequest**](InventoryApiAPI.md#InventoryApiDescribeSubscriptionRequest) | **Get** /2022-09-01-00/fleet/service/{serviceId}/environment/{environmentId}/subscription/request/{id} | DescribeSubscriptionRequest inventory-api
[**InventoryApiDescribeUpgradePath**](InventoryApiAPI.md#InventoryApiDescribeUpgradePath) | **Get** /2022-09-01-00/fleet/service/{serviceId}/productTier/{productTierId}/upgrade-path/{upgradePathId} | DescribeUpgradePath inventory-api
[**InventoryApiDescribeUser**](InventoryApiAPI.md#InventoryApiDescribeUser) | **Get** /2022-09-01-00/fleet/service/{serviceId}/environment/{environmentId}/user/{userId} | DescribeUser inventory-api
[**InventoryApiEnableResourceInstanceManualOverride**](InventoryApiAPI.md#InventoryApiEnableResourceInstanceManualOverride) | **Post** /2022-09-01-00/fleet/service/{serviceId}/environment/{environmentId}/instance/{instanceId}/override | EnableResourceInstanceManualOverride inventory-api
[**InventoryApiFailoverResourceInstance**](InventoryApiAPI.md#InventoryApiFailoverResourceInstance) | **Post** /2022-09-01-00/fleet/service/{serviceId}/environment/{environmentId}/instance/{instanceId}/failover | FailoverResourceInstance inventory-api
[**InventoryApiGenerateTokenForHostClusterDashboard**](InventoryApiAPI.md#InventoryApiGenerateTokenForHostClusterDashboard) | **Post** /2022-09-01-00/fleet/service/{serviceId}/environment/{environmentId}/host-cluster/{id}/dashboard/token | GenerateTokenForHostClusterDashboard inventory-api
[**InventoryApiListActiveOrganizations**](InventoryApiAPI.md#InventoryApiListActiveOrganizations) | **Get** /2022-09-01-00/fleet/service/{serviceId}/environment/{environmentId}/organizations | ListActiveOrganizations inventory-api
[**InventoryApiListAllUsers**](InventoryApiAPI.md#InventoryApiListAllUsers) | **Get** /2022-09-01-00/fleet/users | ListAllUsers inventory-api
[**InventoryApiListDependentComponents**](InventoryApiAPI.md#InventoryApiListDependentComponents) | **Get** /2022-09-01-00/fleet/service/{serviceId}/environment/{environmentId}/instance/{instanceId}/dependent-components | ListDependentComponents inventory-api
[**InventoryApiListEligibleInstancesPerUpgrade**](InventoryApiAPI.md#InventoryApiListEligibleInstancesPerUpgrade) | **Get** /2022-09-01-00/fleet/service/{serviceId}/productTier/{productTierId}/upgrade-path/{upgradePathId}/eligible-instances | ListEligibleInstancesPerUpgrade inventory-api
[**InventoryApiListHostClusters**](InventoryApiAPI.md#InventoryApiListHostClusters) | **Get** /2022-09-01-00/fleet/service/{serviceId}/environment/{environmentId}/host-clusters | ListHostClusters inventory-api
[**InventoryApiListInstanceEvents**](InventoryApiAPI.md#InventoryApiListInstanceEvents) | **Get** /2022-09-01-00/fleet/service/{serviceId}/environment/{environmentId}/instance/{instanceId}/events | ListInstanceEvents inventory-api
[**InventoryApiListLinkedInstances**](InventoryApiAPI.md#InventoryApiListLinkedInstances) | **Get** /2022-09-01-00/fleet/service/{serviceId}/environment/{environmentId}/instance/{instanceId}/linked-instances | ListLinkedInstances inventory-api
[**InventoryApiListResourceInstanceSnapshots**](InventoryApiAPI.md#InventoryApiListResourceInstanceSnapshots) | **Get** /2022-09-01-00/fleet/service/{serviceId}/environment/{environmentId}/instance/{instanceId}/snapshot | ListResourceInstanceSnapshots inventory-api
[**InventoryApiListResourceInstances**](InventoryApiAPI.md#InventoryApiListResourceInstances) | **Get** /2022-09-01-00/fleet/service/{serviceId}/environment/{environmentId}/instances/ | ListResourceInstances inventory-api
[**InventoryApiListResources**](InventoryApiAPI.md#InventoryApiListResources) | **Get** /2022-09-01-00/fleet/service/{serviceId}/environment/{environmentId}/model/{serviceModelId}/productTier/{productTierId}/resources | ListResources inventory-api
[**InventoryApiListServiceOfferings**](InventoryApiAPI.md#InventoryApiListServiceOfferings) | **Get** /2022-09-01-00/fleet/service-offering | ListServiceOfferings inventory-api
[**InventoryApiListServicesOrchestrations**](InventoryApiAPI.md#InventoryApiListServicesOrchestrations) | **Get** /2022-09-01-00/fleet/services-orchestration | ListServicesOrchestrations inventory-api
[**InventoryApiListSubscription**](InventoryApiAPI.md#InventoryApiListSubscription) | **Get** /2022-09-01-00/fleet/service/{serviceId}/environment/{environmentId}/subscription | ListSubscription inventory-api
[**InventoryApiListSubscriptionRequests**](InventoryApiAPI.md#InventoryApiListSubscriptionRequests) | **Get** /2022-09-01-00/fleet/service/{serviceId}/environment/{environmentId}/subscription/request | ListSubscriptionRequests inventory-api
[**InventoryApiListUpgradePaths**](InventoryApiAPI.md#InventoryApiListUpgradePaths) | **Get** /2022-09-01-00/fleet/service/{serviceId}/productTier/{productTierId}/upgrade-paths | ListUpgradePaths inventory-api
[**InventoryApiListUsers**](InventoryApiAPI.md#InventoryApiListUsers) | **Get** /2022-09-01-00/fleet/service/{serviceId}/environment/{environmentId}/users | ListUsers inventory-api
[**InventoryApiModifyServicesOrchestration**](InventoryApiAPI.md#InventoryApiModifyServicesOrchestration) | **Patch** /2022-09-01-00/fleet/services-orchestration/{id} | ModifyServicesOrchestration inventory-api
[**InventoryApiRemoveCapacityFromResourceInstance**](InventoryApiAPI.md#InventoryApiRemoveCapacityFromResourceInstance) | **Post** /2022-09-01-00/fleet/service/{serviceId}/environment/{environmentId}/instance/{instanceId}/remove-capacity | RemoveCapacityFromResourceInstance inventory-api
[**InventoryApiRemoveCustomDNSFromResourceInstance**](InventoryApiAPI.md#InventoryApiRemoveCustomDNSFromResourceInstance) | **Delete** /2022-09-01-00/fleet/service/{serviceId}/environment/{environmentId}/{resourceKey}/instance/{instanceId}/custom-dns | RemoveCustomDNSFromResourceInstance inventory-api
[**InventoryApiRestartResourceInstance**](InventoryApiAPI.md#InventoryApiRestartResourceInstance) | **Post** /2022-09-01-00/fleet/service/{serviceId}/environment/{environmentId}/instance/{instanceId}/restart | RestartResourceInstance inventory-api
[**InventoryApiRestoreResourceInstance**](InventoryApiAPI.md#InventoryApiRestoreResourceInstance) | **Post** /2022-09-01-00/fleet/service/{serviceId}/environment/{environmentId}/instance/{instanceId}/restore | RestoreResourceInstance inventory-api
[**InventoryApiRestoreResourceInstanceFromSnapshot**](InventoryApiAPI.md#InventoryApiRestoreResourceInstanceFromSnapshot) | **Post** /2022-09-01-00/fleet/service/{serviceId}/environment/{environmentId}/snapshot/{snapshotId}/restore | RestoreResourceInstanceFromSnapshot inventory-api
[**InventoryApiResumeSubscription**](InventoryApiAPI.md#InventoryApiResumeSubscription) | **Put** /2022-09-01-00/fleet/service/{serviceId}/environment/{environmentId}/subscription/{id}/resume | ResumeSubscription inventory-api
[**InventoryApiSearchInventory**](InventoryApiAPI.md#InventoryApiSearchInventory) | **Post** /2022-09-01-00/fleet/search-inventory | SearchInventory inventory-api
[**InventoryApiSearchServiceInventory**](InventoryApiAPI.md#InventoryApiSearchServiceInventory) | **Post** /2022-09-01-00/fleet/service/{serviceId}/environment/{environmentId}/search-inventory | SearchServiceInventory inventory-api
[**InventoryApiStartResourceInstance**](InventoryApiAPI.md#InventoryApiStartResourceInstance) | **Post** /2022-09-01-00/fleet/service/{serviceId}/environment/{environmentId}/instance/{instanceId}/start | StartResourceInstance inventory-api
[**InventoryApiStopResourceInstance**](InventoryApiAPI.md#InventoryApiStopResourceInstance) | **Post** /2022-09-01-00/fleet/service/{serviceId}/environment/{environmentId}/instance/{instanceId}/stop | StopResourceInstance inventory-api
[**InventoryApiSuspendSubscription**](InventoryApiAPI.md#InventoryApiSuspendSubscription) | **Put** /2022-09-01-00/fleet/service/{serviceId}/environment/{environmentId}/subscription/{id}/suspend | SuspendSubscription inventory-api
[**InventoryApiSuspendUser**](InventoryApiAPI.md#InventoryApiSuspendUser) | **Put** /2022-09-01-00/fleet/user/{userId}/suspend | SuspendUser inventory-api
[**InventoryApiTerminateSubscription**](InventoryApiAPI.md#InventoryApiTerminateSubscription) | **Delete** /2022-09-01-00/fleet/service/{serviceId}/environment/{environmentId}/subscription/{id} | TerminateSubscription inventory-api
[**InventoryApiUnsuspendUser**](InventoryApiAPI.md#InventoryApiUnsuspendUser) | **Put** /2022-09-01-00/fleet/user/{userId}/unsuspend | UnsuspendUser inventory-api
[**InventoryApiUpdateResourceInstance**](InventoryApiAPI.md#InventoryApiUpdateResourceInstance) | **Patch** /2022-09-01-00/fleet/service/{serviceId}/environment/{environmentId}/instance/{instanceId} | UpdateResourceInstance inventory-api



## InventoryApiAddCapacityToResourceInstance

> InventoryApiAddCapacityToResourceInstance(ctx, serviceId, environmentId, instanceId).FleetAddCapacityToResourceInstanceRequest2(fleetAddCapacityToResourceInstanceRequest2).Execute()

AddCapacityToResourceInstance inventory-api

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
	serviceId := "s-12345678" // string | The service ID this workflow belongs to.
	environmentId := "se-12345678" // string | The service environment ID this workflow belongs to.
	instanceId := "instance-12345678" // string | The resource instance ID.
	fleetAddCapacityToResourceInstanceRequest2 := *openapiclient.NewFleetAddCapacityToResourceInstanceRequest2(int64(3), "r-12345678") // FleetAddCapacityToResourceInstanceRequest2 | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	r, err := apiClient.InventoryApiAPI.InventoryApiAddCapacityToResourceInstance(context.Background(), serviceId, environmentId, instanceId).FleetAddCapacityToResourceInstanceRequest2(fleetAddCapacityToResourceInstanceRequest2).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `InventoryApiAPI.InventoryApiAddCapacityToResourceInstance``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**serviceId** | **string** | The service ID this workflow belongs to. | 
**environmentId** | **string** | The service environment ID this workflow belongs to. | 
**instanceId** | **string** | The resource instance ID. | 

### Other Parameters

Other parameters are passed through a pointer to a apiInventoryApiAddCapacityToResourceInstanceRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------



 **fleetAddCapacityToResourceInstanceRequest2** | [**FleetAddCapacityToResourceInstanceRequest2**](FleetAddCapacityToResourceInstanceRequest2.md) |  | 

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


## InventoryApiAddCustomDNSToResourceInstance

> InventoryApiAddCustomDNSToResourceInstance(ctx, serviceId, environmentId, resourceKey, instanceId).FleetAddCustomDNSToResourceInstanceRequest2(fleetAddCustomDNSToResourceInstanceRequest2).Execute()

AddCustomDNSToResourceInstance inventory-api

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
	serviceId := "s-12345678" // string | The service ID this workflow belongs to.
	environmentId := "se-12345678" // string | The service environment ID this workflow belongs to.
	resourceKey := "http-service" // string | The resource key
	instanceId := "instance-12345678" // string | The resource instance ID.
	fleetAddCustomDNSToResourceInstanceRequest2 := *openapiclient.NewFleetAddCustomDNSToResourceInstanceRequest2("my-custom-dns.com") // FleetAddCustomDNSToResourceInstanceRequest2 | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	r, err := apiClient.InventoryApiAPI.InventoryApiAddCustomDNSToResourceInstance(context.Background(), serviceId, environmentId, resourceKey, instanceId).FleetAddCustomDNSToResourceInstanceRequest2(fleetAddCustomDNSToResourceInstanceRequest2).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `InventoryApiAPI.InventoryApiAddCustomDNSToResourceInstance``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**serviceId** | **string** | The service ID this workflow belongs to. | 
**environmentId** | **string** | The service environment ID this workflow belongs to. | 
**resourceKey** | **string** | The resource key | 
**instanceId** | **string** | The resource instance ID. | 

### Other Parameters

Other parameters are passed through a pointer to a apiInventoryApiAddCustomDNSToResourceInstanceRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------




 **fleetAddCustomDNSToResourceInstanceRequest2** | [**FleetAddCustomDNSToResourceInstanceRequest2**](FleetAddCustomDNSToResourceInstanceRequest2.md) |  | 

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


## InventoryApiApproveSubscriptionRequest

> InventoryApiApproveSubscriptionRequest(ctx, serviceId, environmentId, id).Execute()

ApproveSubscriptionRequest inventory-api

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
	serviceId := "s-12345678" // string | The service ID this workflow belongs to.
	environmentId := "se-12345678" // string | The service environment ID this workflow belongs to.
	id := "subr-12345678" // string | The subscription ID

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	r, err := apiClient.InventoryApiAPI.InventoryApiApproveSubscriptionRequest(context.Background(), serviceId, environmentId, id).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `InventoryApiAPI.InventoryApiApproveSubscriptionRequest``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**serviceId** | **string** | The service ID this workflow belongs to. | 
**environmentId** | **string** | The service environment ID this workflow belongs to. | 
**id** | **string** | The subscription ID | 

### Other Parameters

Other parameters are passed through a pointer to a apiInventoryApiApproveSubscriptionRequestRequest struct via the builder pattern


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


## InventoryApiCancelUpgradePath

> UpgradePath InventoryApiCancelUpgradePath(ctx, serviceId, productTierId, upgradePathId).Execute()

CancelUpgradePath inventory-api

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
	serviceId := "s-12345678" // string | The service ID this workflow belongs to.
	productTierId := "pt-12345678" // string | The product tier ID that this upgrade path belongs to
	upgradePathId := "up-12345678" // string | The upgrade path ID

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.InventoryApiAPI.InventoryApiCancelUpgradePath(context.Background(), serviceId, productTierId, upgradePathId).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `InventoryApiAPI.InventoryApiCancelUpgradePath``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `InventoryApiCancelUpgradePath`: UpgradePath
	fmt.Fprintf(os.Stdout, "Response from `InventoryApiAPI.InventoryApiCancelUpgradePath`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**serviceId** | **string** | The service ID this workflow belongs to. | 
**productTierId** | **string** | The product tier ID that this upgrade path belongs to | 
**upgradePathId** | **string** | The upgrade path ID | 

### Other Parameters

Other parameters are passed through a pointer to a apiInventoryApiCancelUpgradePathRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------




### Return type

[**UpgradePath**](UpgradePath.md)

### Authorization

[api_key_header_Authorization](../README.md#api_key_header_Authorization)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json, application/vnd.goa.error

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## InventoryApiCreateConsumptionUser

> string InventoryApiCreateConsumptionUser(ctx).FleetCreateConsumptionUserRequest2(fleetCreateConsumptionUserRequest2).Execute()

CreateConsumptionUser inventory-api

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
	fleetCreateConsumptionUserRequest2 := *openapiclient.NewFleetCreateConsumptionUserRequest2("abc@gmail.com", "mywebsite", "John Doe", "password") // FleetCreateConsumptionUserRequest2 | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.InventoryApiAPI.InventoryApiCreateConsumptionUser(context.Background()).FleetCreateConsumptionUserRequest2(fleetCreateConsumptionUserRequest2).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `InventoryApiAPI.InventoryApiCreateConsumptionUser``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `InventoryApiCreateConsumptionUser`: string
	fmt.Fprintf(os.Stdout, "Response from `InventoryApiAPI.InventoryApiCreateConsumptionUser`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiInventoryApiCreateConsumptionUserRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **fleetCreateConsumptionUserRequest2** | [**FleetCreateConsumptionUserRequest2**](FleetCreateConsumptionUserRequest2.md) |  | 

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


## InventoryApiCreateProxyResourceInstance

> CreateResourceInstanceResponseBody InventoryApiCreateProxyResourceInstance(ctx, serviceProviderId, serviceKey, serviceAPIVersion, serviceEnvironmentKey, serviceModelKey, productTierKey, sourceResourceKey).CreateProxyResourceInstanceRequest2(createProxyResourceInstanceRequest2).Execute()

CreateProxyResourceInstance inventory-api

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
	serviceProviderId := "omnistrate" // string | The service provider ID
	serviceKey := "service-orchestration" // string | The service name
	serviceAPIVersion := "v1" // string | The service API version
	serviceEnvironmentKey := "dev" // string | The service environment name
	serviceModelKey := "hosted" // string | The service model name
	productTierKey := "premium" // string | The product tier name
	sourceResourceKey := "mysql" // string | The serverless resource key
	createProxyResourceInstanceRequest2 := *openapiclient.NewCreateProxyResourceInstanceRequest2() // CreateProxyResourceInstanceRequest2 | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.InventoryApiAPI.InventoryApiCreateProxyResourceInstance(context.Background(), serviceProviderId, serviceKey, serviceAPIVersion, serviceEnvironmentKey, serviceModelKey, productTierKey, sourceResourceKey).CreateProxyResourceInstanceRequest2(createProxyResourceInstanceRequest2).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `InventoryApiAPI.InventoryApiCreateProxyResourceInstance``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `InventoryApiCreateProxyResourceInstance`: CreateResourceInstanceResponseBody
	fmt.Fprintf(os.Stdout, "Response from `InventoryApiAPI.InventoryApiCreateProxyResourceInstance`: %v\n", resp)
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
**sourceResourceKey** | **string** | The serverless resource key | 

### Other Parameters

Other parameters are passed through a pointer to a apiInventoryApiCreateProxyResourceInstanceRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------







 **createProxyResourceInstanceRequest2** | [**CreateProxyResourceInstanceRequest2**](CreateProxyResourceInstanceRequest2.md) |  | 

### Return type

[**CreateResourceInstanceResponseBody**](CreateResourceInstanceResponseBody.md)

### Authorization

[api_key_header_Authorization](../README.md#api_key_header_Authorization)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json, application/vnd.goa.error

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## InventoryApiCreateResourceInstance

> CreateResourceInstanceResponseBody InventoryApiCreateResourceInstance(ctx, serviceProviderId, serviceKey, serviceAPIVersion, serviceEnvironmentKey, serviceModelKey, productTierKey, resourceKey).FleetCreateResourceInstanceRequest2(fleetCreateResourceInstanceRequest2).Execute()

CreateResourceInstance inventory-api

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
	serviceProviderId := "omnistrate" // string | The service provider ID
	serviceKey := "service-orchestration" // string | The service name
	serviceAPIVersion := "v1" // string | The service API version
	serviceEnvironmentKey := "dev" // string | The service environment name
	serviceModelKey := "hosted" // string | The service model name
	productTierKey := "premium" // string | The product tier name
	resourceKey := "mysql" // string | The resource key
	fleetCreateResourceInstanceRequest2 := *openapiclient.NewFleetCreateResourceInstanceRequest2() // FleetCreateResourceInstanceRequest2 | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.InventoryApiAPI.InventoryApiCreateResourceInstance(context.Background(), serviceProviderId, serviceKey, serviceAPIVersion, serviceEnvironmentKey, serviceModelKey, productTierKey, resourceKey).FleetCreateResourceInstanceRequest2(fleetCreateResourceInstanceRequest2).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `InventoryApiAPI.InventoryApiCreateResourceInstance``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `InventoryApiCreateResourceInstance`: CreateResourceInstanceResponseBody
	fmt.Fprintf(os.Stdout, "Response from `InventoryApiAPI.InventoryApiCreateResourceInstance`: %v\n", resp)
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

Other parameters are passed through a pointer to a apiInventoryApiCreateResourceInstanceRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------







 **fleetCreateResourceInstanceRequest2** | [**FleetCreateResourceInstanceRequest2**](FleetCreateResourceInstanceRequest2.md) |  | 

### Return type

[**CreateResourceInstanceResponseBody**](CreateResourceInstanceResponseBody.md)

### Authorization

[api_key_header_Authorization](../README.md#api_key_header_Authorization)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json, application/vnd.goa.error

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## InventoryApiCreateResourceInstanceSnapshot

> FleetCreateInstanceSnapshotResult InventoryApiCreateResourceInstanceSnapshot(ctx, serviceId, environmentId, instanceId).Execute()

CreateResourceInstanceSnapshot inventory-api

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
	serviceId := "s-12345678" // string | The service ID this workflow belongs to.
	environmentId := "se-12345678" // string | The service environment ID this workflow belongs to.
	instanceId := "instance-12345678" // string | The resource instance ID.

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.InventoryApiAPI.InventoryApiCreateResourceInstanceSnapshot(context.Background(), serviceId, environmentId, instanceId).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `InventoryApiAPI.InventoryApiCreateResourceInstanceSnapshot``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `InventoryApiCreateResourceInstanceSnapshot`: FleetCreateInstanceSnapshotResult
	fmt.Fprintf(os.Stdout, "Response from `InventoryApiAPI.InventoryApiCreateResourceInstanceSnapshot`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**serviceId** | **string** | The service ID this workflow belongs to. | 
**environmentId** | **string** | The service environment ID this workflow belongs to. | 
**instanceId** | **string** | The resource instance ID. | 

### Other Parameters

Other parameters are passed through a pointer to a apiInventoryApiCreateResourceInstanceSnapshotRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------




### Return type

[**FleetCreateInstanceSnapshotResult**](FleetCreateInstanceSnapshotResult.md)

### Authorization

[api_key_header_Authorization](../README.md#api_key_header_Authorization)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json, application/vnd.goa.error

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## InventoryApiCreateServicesOrchestration

> CreateResourceInstanceResponseBody InventoryApiCreateServicesOrchestration(ctx).FleetCreateServicesOrchestrationRequest2(fleetCreateServicesOrchestrationRequest2).Execute()

CreateServicesOrchestration inventory-api

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
	fleetCreateServicesOrchestrationRequest2 := *openapiclient.NewFleetCreateServicesOrchestrationRequest2("Nobis cupiditate dicta.") // FleetCreateServicesOrchestrationRequest2 | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.InventoryApiAPI.InventoryApiCreateServicesOrchestration(context.Background()).FleetCreateServicesOrchestrationRequest2(fleetCreateServicesOrchestrationRequest2).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `InventoryApiAPI.InventoryApiCreateServicesOrchestration``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `InventoryApiCreateServicesOrchestration`: CreateResourceInstanceResponseBody
	fmt.Fprintf(os.Stdout, "Response from `InventoryApiAPI.InventoryApiCreateServicesOrchestration`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiInventoryApiCreateServicesOrchestrationRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **fleetCreateServicesOrchestrationRequest2** | [**FleetCreateServicesOrchestrationRequest2**](FleetCreateServicesOrchestrationRequest2.md) |  | 

### Return type

[**CreateResourceInstanceResponseBody**](CreateResourceInstanceResponseBody.md)

### Authorization

[api_key_header_Authorization](../README.md#api_key_header_Authorization)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json, application/vnd.goa.error

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## InventoryApiCreateUpgradePath

> UpgradePath InventoryApiCreateUpgradePath(ctx, serviceId, productTierId).CreateUpgradePathRequest2(createUpgradePathRequest2).Execute()

CreateUpgradePath inventory-api

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
	serviceId := "s-12345678" // string | The service ID this workflow belongs to.
	productTierId := "pt-12345678" // string | The product tier ID that this upgrade path belongs to
	createUpgradePathRequest2 := *openapiclient.NewCreateUpgradePathRequest2("1.0", "2.0", map[string][]string{"key": []string{"Occaecati sit."}}) // CreateUpgradePathRequest2 | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.InventoryApiAPI.InventoryApiCreateUpgradePath(context.Background(), serviceId, productTierId).CreateUpgradePathRequest2(createUpgradePathRequest2).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `InventoryApiAPI.InventoryApiCreateUpgradePath``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `InventoryApiCreateUpgradePath`: UpgradePath
	fmt.Fprintf(os.Stdout, "Response from `InventoryApiAPI.InventoryApiCreateUpgradePath`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**serviceId** | **string** | The service ID this workflow belongs to. | 
**productTierId** | **string** | The product tier ID that this upgrade path belongs to | 

### Other Parameters

Other parameters are passed through a pointer to a apiInventoryApiCreateUpgradePathRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


 **createUpgradePathRequest2** | [**CreateUpgradePathRequest2**](CreateUpgradePathRequest2.md) |  | 

### Return type

[**UpgradePath**](UpgradePath.md)

### Authorization

[api_key_header_Authorization](../README.md#api_key_header_Authorization)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json, application/vnd.goa.error

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## InventoryApiDebugResourceInstance

> DebugResourceInstanceResult InventoryApiDebugResourceInstance(ctx, serviceId, environmentId, instanceId).Execute()

DebugResourceInstance inventory-api

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
	serviceId := "s-12345678" // string | The service ID this workflow belongs to.
	environmentId := "se-12345678" // string | The service environment ID this workflow belongs to.
	instanceId := "instance-12345678" // string | The resource instance ID.

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.InventoryApiAPI.InventoryApiDebugResourceInstance(context.Background(), serviceId, environmentId, instanceId).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `InventoryApiAPI.InventoryApiDebugResourceInstance``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `InventoryApiDebugResourceInstance`: DebugResourceInstanceResult
	fmt.Fprintf(os.Stdout, "Response from `InventoryApiAPI.InventoryApiDebugResourceInstance`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**serviceId** | **string** | The service ID this workflow belongs to. | 
**environmentId** | **string** | The service environment ID this workflow belongs to. | 
**instanceId** | **string** | The resource instance ID. | 

### Other Parameters

Other parameters are passed through a pointer to a apiInventoryApiDebugResourceInstanceRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------




### Return type

[**DebugResourceInstanceResult**](DebugResourceInstanceResult.md)

### Authorization

[api_key_header_Authorization](../README.md#api_key_header_Authorization)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json, application/vnd.goa.error

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## InventoryApiDeleteProxyResourceInstance

> InventoryApiDeleteProxyResourceInstance(ctx, serviceProviderId, serviceKey, serviceAPIVersion, serviceEnvironmentKey, serviceModelKey, productTierKey, id).Execute()

DeleteProxyResourceInstance inventory-api

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
	serviceProviderId := "omnistrate" // string | The service provider ID
	serviceKey := "service-orchestration" // string | The service name
	serviceAPIVersion := "v1" // string | The service API version
	serviceEnvironmentKey := "dev" // string | The service environment name
	serviceModelKey := "hosted" // string | The service model name
	productTierKey := "premium" // string | The product tier name
	id := "instance-abcd1234" // string | The instance ID

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	r, err := apiClient.InventoryApiAPI.InventoryApiDeleteProxyResourceInstance(context.Background(), serviceProviderId, serviceKey, serviceAPIVersion, serviceEnvironmentKey, serviceModelKey, productTierKey, id).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `InventoryApiAPI.InventoryApiDeleteProxyResourceInstance``: %v\n", err)
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
**id** | **string** | The instance ID | 

### Other Parameters

Other parameters are passed through a pointer to a apiInventoryApiDeleteProxyResourceInstanceRequest struct via the builder pattern


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


## InventoryApiDeleteResourceInstance

> InventoryApiDeleteResourceInstance(ctx, serviceId, environmentId, instanceId).FleetDeleteResourceInstanceRequest2(fleetDeleteResourceInstanceRequest2).Execute()

DeleteResourceInstance inventory-api

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
	serviceId := "s-12345678" // string | The service ID this workflow belongs to.
	environmentId := "se-12345678" // string | The service environment ID this workflow belongs to.
	instanceId := "instance-12345678" // string | The resource instance ID.
	fleetDeleteResourceInstanceRequest2 := *openapiclient.NewFleetDeleteResourceInstanceRequest2("r-12345678") // FleetDeleteResourceInstanceRequest2 | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	r, err := apiClient.InventoryApiAPI.InventoryApiDeleteResourceInstance(context.Background(), serviceId, environmentId, instanceId).FleetDeleteResourceInstanceRequest2(fleetDeleteResourceInstanceRequest2).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `InventoryApiAPI.InventoryApiDeleteResourceInstance``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**serviceId** | **string** | The service ID this workflow belongs to. | 
**environmentId** | **string** | The service environment ID this workflow belongs to. | 
**instanceId** | **string** | The resource instance ID. | 

### Other Parameters

Other parameters are passed through a pointer to a apiInventoryApiDeleteResourceInstanceRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------



 **fleetDeleteResourceInstanceRequest2** | [**FleetDeleteResourceInstanceRequest2**](FleetDeleteResourceInstanceRequest2.md) |  | 

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


## InventoryApiDeleteResourceInstanceSnapshot

> InventoryApiDeleteResourceInstanceSnapshot(ctx, serviceId, environmentId, instanceId, snapshotId).Execute()

DeleteResourceInstanceSnapshot inventory-api

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
	serviceId := "s-12345678" // string | The service ID this workflow belongs to.
	environmentId := "se-12345678" // string | The service environment ID this workflow belongs to.
	instanceId := "instance-12345678" // string | The resource instance ID.
	snapshotId := "instance-ss-12345678" // string | The instance snapshot ID

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	r, err := apiClient.InventoryApiAPI.InventoryApiDeleteResourceInstanceSnapshot(context.Background(), serviceId, environmentId, instanceId, snapshotId).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `InventoryApiAPI.InventoryApiDeleteResourceInstanceSnapshot``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**serviceId** | **string** | The service ID this workflow belongs to. | 
**environmentId** | **string** | The service environment ID this workflow belongs to. | 
**instanceId** | **string** | The resource instance ID. | 
**snapshotId** | **string** | The instance snapshot ID | 

### Other Parameters

Other parameters are passed through a pointer to a apiInventoryApiDeleteResourceInstanceSnapshotRequest struct via the builder pattern


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


## InventoryApiDeleteServicesOrchestration

> InventoryApiDeleteServicesOrchestration(ctx, id).Execute()

DeleteServicesOrchestration inventory-api

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
	id := "so-12345678" // string | The ID of the services orchestration

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	r, err := apiClient.InventoryApiAPI.InventoryApiDeleteServicesOrchestration(context.Background(), id).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `InventoryApiAPI.InventoryApiDeleteServicesOrchestration``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **string** | The ID of the services orchestration | 

### Other Parameters

Other parameters are passed through a pointer to a apiInventoryApiDeleteServicesOrchestrationRequest struct via the builder pattern


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


## InventoryApiDeleteUser

> InventoryApiDeleteUser(ctx, userId).Execute()

DeleteUser inventory-api

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
	userId := "user-12345678" // string | The user ID

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	r, err := apiClient.InventoryApiAPI.InventoryApiDeleteUser(context.Background(), userId).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `InventoryApiAPI.InventoryApiDeleteUser``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**userId** | **string** | The user ID | 

### Other Parameters

Other parameters are passed through a pointer to a apiInventoryApiDeleteUserRequest struct via the builder pattern


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


## InventoryApiDenySubscriptionRequest

> InventoryApiDenySubscriptionRequest(ctx, serviceId, environmentId, id).Execute()

DenySubscriptionRequest inventory-api

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
	serviceId := "s-12345678" // string | The service ID this workflow belongs to.
	environmentId := "se-12345678" // string | The service environment ID this workflow belongs to.
	id := "subr-12345678" // string | The subscription ID

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	r, err := apiClient.InventoryApiAPI.InventoryApiDenySubscriptionRequest(context.Background(), serviceId, environmentId, id).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `InventoryApiAPI.InventoryApiDenySubscriptionRequest``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**serviceId** | **string** | The service ID this workflow belongs to. | 
**environmentId** | **string** | The service environment ID this workflow belongs to. | 
**id** | **string** | The subscription ID | 

### Other Parameters

Other parameters are passed through a pointer to a apiInventoryApiDenySubscriptionRequestRequest struct via the builder pattern


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


## InventoryApiDescribeHostCluster

> FleetDescribeHostClusterResult InventoryApiDescribeHostCluster(ctx, serviceId, environmentId, id).Execute()

DescribeHostCluster inventory-api

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
	serviceId := "s-12345678" // string | The service ID this workflow belongs to.
	environmentId := "se-12345678" // string | The service environment ID this workflow belongs to.
	id := "hc-12345678" // string | The host cluster ID

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.InventoryApiAPI.InventoryApiDescribeHostCluster(context.Background(), serviceId, environmentId, id).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `InventoryApiAPI.InventoryApiDescribeHostCluster``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `InventoryApiDescribeHostCluster`: FleetDescribeHostClusterResult
	fmt.Fprintf(os.Stdout, "Response from `InventoryApiAPI.InventoryApiDescribeHostCluster`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**serviceId** | **string** | The service ID this workflow belongs to. | 
**environmentId** | **string** | The service environment ID this workflow belongs to. | 
**id** | **string** | The host cluster ID | 

### Other Parameters

Other parameters are passed through a pointer to a apiInventoryApiDescribeHostClusterRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------




### Return type

[**FleetDescribeHostClusterResult**](FleetDescribeHostClusterResult.md)

### Authorization

[api_key_header_Authorization](../README.md#api_key_header_Authorization)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json, application/vnd.goa.error

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## InventoryApiDescribeInstanceEvent

> FleetDescribeEventResult InventoryApiDescribeInstanceEvent(ctx, serviceId, environmentId, instanceId, id).Execute()

DescribeInstanceEvent inventory-api

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
	serviceId := "s-12345678" // string | The service ID this workflow belongs to.
	environmentId := "se-12345678" // string | The service environment ID this workflow belongs to.
	instanceId := "instance-12345678" // string | The resource instance ID.
	id := "event-12345678" // string | The ID of the event

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.InventoryApiAPI.InventoryApiDescribeInstanceEvent(context.Background(), serviceId, environmentId, instanceId, id).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `InventoryApiAPI.InventoryApiDescribeInstanceEvent``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `InventoryApiDescribeInstanceEvent`: FleetDescribeEventResult
	fmt.Fprintf(os.Stdout, "Response from `InventoryApiAPI.InventoryApiDescribeInstanceEvent`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**serviceId** | **string** | The service ID this workflow belongs to. | 
**environmentId** | **string** | The service environment ID this workflow belongs to. | 
**instanceId** | **string** | The resource instance ID. | 
**id** | **string** | The ID of the event | 

### Other Parameters

Other parameters are passed through a pointer to a apiInventoryApiDescribeInstanceEventRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------





### Return type

[**FleetDescribeEventResult**](FleetDescribeEventResult.md)

### Authorization

[api_key_header_Authorization](../README.md#api_key_header_Authorization)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json, application/vnd.goa.error

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## InventoryApiDescribeInventorySummary

> DescribeInventorySummaryResult InventoryApiDescribeInventorySummary(ctx, serviceId, environmentId).Execute()

DescribeInventorySummary inventory-api

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
	serviceId := "s-12345678" // string | The service ID this workflow belongs to.
	environmentId := "se-12345678" // string | The service environment ID this workflow belongs to.

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.InventoryApiAPI.InventoryApiDescribeInventorySummary(context.Background(), serviceId, environmentId).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `InventoryApiAPI.InventoryApiDescribeInventorySummary``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `InventoryApiDescribeInventorySummary`: DescribeInventorySummaryResult
	fmt.Fprintf(os.Stdout, "Response from `InventoryApiAPI.InventoryApiDescribeInventorySummary`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**serviceId** | **string** | The service ID this workflow belongs to. | 
**environmentId** | **string** | The service environment ID this workflow belongs to. | 

### Other Parameters

Other parameters are passed through a pointer to a apiInventoryApiDescribeInventorySummaryRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------



### Return type

[**DescribeInventorySummaryResult**](DescribeInventorySummaryResult.md)

### Authorization

[api_key_header_Authorization](../README.md#api_key_header_Authorization)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json, application/vnd.goa.error

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## InventoryApiDescribeOrgUser

> FleetDescribeUserResult InventoryApiDescribeOrgUser(ctx, userId).Execute()

DescribeOrgUser inventory-api

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
	userId := "user-12345678" // string | The user ID

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.InventoryApiAPI.InventoryApiDescribeOrgUser(context.Background(), userId).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `InventoryApiAPI.InventoryApiDescribeOrgUser``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `InventoryApiDescribeOrgUser`: FleetDescribeUserResult
	fmt.Fprintf(os.Stdout, "Response from `InventoryApiAPI.InventoryApiDescribeOrgUser`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**userId** | **string** | The user ID | 

### Other Parameters

Other parameters are passed through a pointer to a apiInventoryApiDescribeOrgUserRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

[**FleetDescribeUserResult**](FleetDescribeUserResult.md)

### Authorization

[api_key_header_Authorization](../README.md#api_key_header_Authorization)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json, application/vnd.goa.error

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## InventoryApiDescribeOrganization

> Organization InventoryApiDescribeOrganization(ctx, serviceId, environmentId, organizationId).Execute()

DescribeOrganization inventory-api

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
	serviceId := "s-12345678" // string | The service ID this workflow belongs to.
	environmentId := "se-12345678" // string | The service environment ID this workflow belongs to.
	organizationId := "o-12345678" // string | The organization ID.

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.InventoryApiAPI.InventoryApiDescribeOrganization(context.Background(), serviceId, environmentId, organizationId).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `InventoryApiAPI.InventoryApiDescribeOrganization``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `InventoryApiDescribeOrganization`: Organization
	fmt.Fprintf(os.Stdout, "Response from `InventoryApiAPI.InventoryApiDescribeOrganization`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**serviceId** | **string** | The service ID this workflow belongs to. | 
**environmentId** | **string** | The service environment ID this workflow belongs to. | 
**organizationId** | **string** | The organization ID. | 

### Other Parameters

Other parameters are passed through a pointer to a apiInventoryApiDescribeOrganizationRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------




### Return type

[**Organization**](Organization.md)

### Authorization

[api_key_header_Authorization](../README.md#api_key_header_Authorization)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json, application/vnd.goa.error

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## InventoryApiDescribeResource

> Resource InventoryApiDescribeResource(ctx, serviceId, environmentId, serviceModelId, productTierId, resourceId).Execute()

DescribeResource inventory-api

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
	serviceId := "s-12345678" // string | The service ID this workflow belongs to.
	environmentId := "se-12345678" // string | The service environment ID this workflow belongs to.
	serviceModelId := "sm-12345678" // string | The service model ID.
	productTierId := "pt-12345678" // string | The product tier ID.
	resourceId := "r-12345678" // string | The resource ID.

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.InventoryApiAPI.InventoryApiDescribeResource(context.Background(), serviceId, environmentId, serviceModelId, productTierId, resourceId).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `InventoryApiAPI.InventoryApiDescribeResource``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `InventoryApiDescribeResource`: Resource
	fmt.Fprintf(os.Stdout, "Response from `InventoryApiAPI.InventoryApiDescribeResource`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**serviceId** | **string** | The service ID this workflow belongs to. | 
**environmentId** | **string** | The service environment ID this workflow belongs to. | 
**serviceModelId** | **string** | The service model ID. | 
**productTierId** | **string** | The product tier ID. | 
**resourceId** | **string** | The resource ID. | 

### Other Parameters

Other parameters are passed through a pointer to a apiInventoryApiDescribeResourceRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------






### Return type

[**Resource**](Resource.md)

### Authorization

[api_key_header_Authorization](../README.md#api_key_header_Authorization)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json, application/vnd.goa.error

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## InventoryApiDescribeResourceInstance

> ResourceInstance InventoryApiDescribeResourceInstance(ctx, serviceId, environmentId, instanceId).Execute()

DescribeResourceInstance inventory-api

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
	serviceId := "s-12345678" // string | The service ID this workflow belongs to.
	environmentId := "se-12345678" // string | The service environment ID this workflow belongs to.
	instanceId := "instance-12345678" // string | The resource instance ID.

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.InventoryApiAPI.InventoryApiDescribeResourceInstance(context.Background(), serviceId, environmentId, instanceId).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `InventoryApiAPI.InventoryApiDescribeResourceInstance``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `InventoryApiDescribeResourceInstance`: ResourceInstance
	fmt.Fprintf(os.Stdout, "Response from `InventoryApiAPI.InventoryApiDescribeResourceInstance`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**serviceId** | **string** | The service ID this workflow belongs to. | 
**environmentId** | **string** | The service environment ID this workflow belongs to. | 
**instanceId** | **string** | The resource instance ID. | 

### Other Parameters

Other parameters are passed through a pointer to a apiInventoryApiDescribeResourceInstanceRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------




### Return type

[**ResourceInstance**](ResourceInstance.md)

### Authorization

[api_key_header_Authorization](../README.md#api_key_header_Authorization)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json, application/vnd.goa.error

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## InventoryApiDescribeResourceInstanceSnapshotFromTime

> FleetDescribeInstanceSnapshotFromTimeResult InventoryApiDescribeResourceInstanceSnapshotFromTime(ctx, serviceId, environmentId, instanceId).FleetDescribeInstanceSnapshotFromTimeRequest2(fleetDescribeInstanceSnapshotFromTimeRequest2).Execute()

DescribeResourceInstanceSnapshotFromTime inventory-api

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
	serviceId := "s-12345678" // string | The service ID this workflow belongs to.
	environmentId := "se-12345678" // string | The service environment ID this workflow belongs to.
	instanceId := "instance-12345678" // string | The resource instance ID.
	fleetDescribeInstanceSnapshotFromTimeRequest2 := *openapiclient.NewFleetDescribeInstanceSnapshotFromTimeRequest2("2021-09-01T00:00:00Z") // FleetDescribeInstanceSnapshotFromTimeRequest2 | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.InventoryApiAPI.InventoryApiDescribeResourceInstanceSnapshotFromTime(context.Background(), serviceId, environmentId, instanceId).FleetDescribeInstanceSnapshotFromTimeRequest2(fleetDescribeInstanceSnapshotFromTimeRequest2).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `InventoryApiAPI.InventoryApiDescribeResourceInstanceSnapshotFromTime``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `InventoryApiDescribeResourceInstanceSnapshotFromTime`: FleetDescribeInstanceSnapshotFromTimeResult
	fmt.Fprintf(os.Stdout, "Response from `InventoryApiAPI.InventoryApiDescribeResourceInstanceSnapshotFromTime`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**serviceId** | **string** | The service ID this workflow belongs to. | 
**environmentId** | **string** | The service environment ID this workflow belongs to. | 
**instanceId** | **string** | The resource instance ID. | 

### Other Parameters

Other parameters are passed through a pointer to a apiInventoryApiDescribeResourceInstanceSnapshotFromTimeRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------



 **fleetDescribeInstanceSnapshotFromTimeRequest2** | [**FleetDescribeInstanceSnapshotFromTimeRequest2**](FleetDescribeInstanceSnapshotFromTimeRequest2.md) |  | 

### Return type

[**FleetDescribeInstanceSnapshotFromTimeResult**](FleetDescribeInstanceSnapshotFromTimeResult.md)

### Authorization

[api_key_header_Authorization](../README.md#api_key_header_Authorization)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json, application/vnd.goa.error

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## InventoryApiDescribeServiceOffering

> InventoryDescribeServiceOfferingResult InventoryApiDescribeServiceOffering(ctx, serviceId).Visibility(visibility).ProductTierId(productTierId).ProductTierVersion(productTierVersion).Execute()

DescribeServiceOffering inventory-api

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
	serviceId := "s-12345678" // string | The service ID
	visibility := "PRIVATE" // string | The visibility of service offering (optional)
	productTierId := "pt-12345678" // string | The product tier Id (optional)
	productTierVersion := "1.0.0" // string | The product tier version (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.InventoryApiAPI.InventoryApiDescribeServiceOffering(context.Background(), serviceId).Visibility(visibility).ProductTierId(productTierId).ProductTierVersion(productTierVersion).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `InventoryApiAPI.InventoryApiDescribeServiceOffering``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `InventoryApiDescribeServiceOffering`: InventoryDescribeServiceOfferingResult
	fmt.Fprintf(os.Stdout, "Response from `InventoryApiAPI.InventoryApiDescribeServiceOffering`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**serviceId** | **string** | The service ID | 

### Other Parameters

Other parameters are passed through a pointer to a apiInventoryApiDescribeServiceOfferingRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **visibility** | **string** | The visibility of service offering | 
 **productTierId** | **string** | The product tier Id | 
 **productTierVersion** | **string** | The product tier version | 

### Return type

[**InventoryDescribeServiceOfferingResult**](InventoryDescribeServiceOfferingResult.md)

### Authorization

[api_key_header_Authorization](../README.md#api_key_header_Authorization)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json, application/vnd.goa.error

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## InventoryApiDescribeServiceOfferingResource

> InventoryDescribeServiceOfferingResourceResult InventoryApiDescribeServiceOfferingResource(ctx, serviceId, resourceId, instanceId).ProductTierId(productTierId).ProductTierVersion(productTierVersion).Execute()

DescribeServiceOfferingResource inventory-api

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
	serviceId := "s-12345678" // string | The service ID
	resourceId := "r-12345678" // string | The resource ID
	instanceId := "instance-12345678" // string | The instance ID (default to "none")
	productTierId := "pt-12345678" // string | The product tier Id (optional)
	productTierVersion := "1.0.0" // string | The product tier version (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.InventoryApiAPI.InventoryApiDescribeServiceOfferingResource(context.Background(), serviceId, resourceId, instanceId).ProductTierId(productTierId).ProductTierVersion(productTierVersion).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `InventoryApiAPI.InventoryApiDescribeServiceOfferingResource``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `InventoryApiDescribeServiceOfferingResource`: InventoryDescribeServiceOfferingResourceResult
	fmt.Fprintf(os.Stdout, "Response from `InventoryApiAPI.InventoryApiDescribeServiceOfferingResource`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**serviceId** | **string** | The service ID | 
**resourceId** | **string** | The resource ID | 
**instanceId** | **string** | The instance ID | [default to &quot;none&quot;]

### Other Parameters

Other parameters are passed through a pointer to a apiInventoryApiDescribeServiceOfferingResourceRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------



 **productTierId** | **string** | The product tier Id | 
 **productTierVersion** | **string** | The product tier version | 

### Return type

[**InventoryDescribeServiceOfferingResourceResult**](InventoryDescribeServiceOfferingResourceResult.md)

### Authorization

[api_key_header_Authorization](../README.md#api_key_header_Authorization)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json, application/vnd.goa.error

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## InventoryApiDescribeServicesOrchestration

> FleetDescribeServicesOrchestrationResult InventoryApiDescribeServicesOrchestration(ctx, id).Execute()

DescribeServicesOrchestration inventory-api

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
	id := "so-12345678" // string | The ID of the services orchestration

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.InventoryApiAPI.InventoryApiDescribeServicesOrchestration(context.Background(), id).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `InventoryApiAPI.InventoryApiDescribeServicesOrchestration``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `InventoryApiDescribeServicesOrchestration`: FleetDescribeServicesOrchestrationResult
	fmt.Fprintf(os.Stdout, "Response from `InventoryApiAPI.InventoryApiDescribeServicesOrchestration`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **string** | The ID of the services orchestration | 

### Other Parameters

Other parameters are passed through a pointer to a apiInventoryApiDescribeServicesOrchestrationRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

[**FleetDescribeServicesOrchestrationResult**](FleetDescribeServicesOrchestrationResult.md)

### Authorization

[api_key_header_Authorization](../README.md#api_key_header_Authorization)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json, application/vnd.goa.error

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## InventoryApiDescribeSubscription

> FleetDescribeSubscriptionResult InventoryApiDescribeSubscription(ctx, serviceId, environmentId, id).Execute()

DescribeSubscription inventory-api

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
	serviceId := "s-12345678" // string | The service ID this workflow belongs to.
	environmentId := "se-12345678" // string | The service environment ID this workflow belongs to.
	id := "sub-12345678" // string | The subscription ID

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.InventoryApiAPI.InventoryApiDescribeSubscription(context.Background(), serviceId, environmentId, id).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `InventoryApiAPI.InventoryApiDescribeSubscription``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `InventoryApiDescribeSubscription`: FleetDescribeSubscriptionResult
	fmt.Fprintf(os.Stdout, "Response from `InventoryApiAPI.InventoryApiDescribeSubscription`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**serviceId** | **string** | The service ID this workflow belongs to. | 
**environmentId** | **string** | The service environment ID this workflow belongs to. | 
**id** | **string** | The subscription ID | 

### Other Parameters

Other parameters are passed through a pointer to a apiInventoryApiDescribeSubscriptionRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------




### Return type

[**FleetDescribeSubscriptionResult**](FleetDescribeSubscriptionResult.md)

### Authorization

[api_key_header_Authorization](../README.md#api_key_header_Authorization)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json, application/vnd.goa.error

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## InventoryApiDescribeSubscriptionRequest

> DescribeSubscriptionRequestResult InventoryApiDescribeSubscriptionRequest(ctx, serviceId, environmentId, id).Execute()

DescribeSubscriptionRequest inventory-api

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
	serviceId := "s-12345678" // string | The service ID this workflow belongs to.
	environmentId := "se-12345678" // string | The service environment ID this workflow belongs to.
	id := "subr-12345678" // string | The subscription ID

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.InventoryApiAPI.InventoryApiDescribeSubscriptionRequest(context.Background(), serviceId, environmentId, id).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `InventoryApiAPI.InventoryApiDescribeSubscriptionRequest``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `InventoryApiDescribeSubscriptionRequest`: DescribeSubscriptionRequestResult
	fmt.Fprintf(os.Stdout, "Response from `InventoryApiAPI.InventoryApiDescribeSubscriptionRequest`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**serviceId** | **string** | The service ID this workflow belongs to. | 
**environmentId** | **string** | The service environment ID this workflow belongs to. | 
**id** | **string** | The subscription ID | 

### Other Parameters

Other parameters are passed through a pointer to a apiInventoryApiDescribeSubscriptionRequestRequest struct via the builder pattern


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


## InventoryApiDescribeUpgradePath

> UpgradePath InventoryApiDescribeUpgradePath(ctx, serviceId, productTierId, upgradePathId).Execute()

DescribeUpgradePath inventory-api

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
	serviceId := "s-12345678" // string | The service ID this workflow belongs to.
	productTierId := "pt-12345678" // string | The product tier ID that this upgrade path belongs to
	upgradePathId := "up-12345678" // string | The upgrade path ID

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.InventoryApiAPI.InventoryApiDescribeUpgradePath(context.Background(), serviceId, productTierId, upgradePathId).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `InventoryApiAPI.InventoryApiDescribeUpgradePath``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `InventoryApiDescribeUpgradePath`: UpgradePath
	fmt.Fprintf(os.Stdout, "Response from `InventoryApiAPI.InventoryApiDescribeUpgradePath`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**serviceId** | **string** | The service ID this workflow belongs to. | 
**productTierId** | **string** | The product tier ID that this upgrade path belongs to | 
**upgradePathId** | **string** | The upgrade path ID | 

### Other Parameters

Other parameters are passed through a pointer to a apiInventoryApiDescribeUpgradePathRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------




### Return type

[**UpgradePath**](UpgradePath.md)

### Authorization

[api_key_header_Authorization](../README.md#api_key_header_Authorization)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json, application/vnd.goa.error

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## InventoryApiDescribeUser

> FleetDescribeUserResult InventoryApiDescribeUser(ctx, serviceId, environmentId, userId).Execute()

DescribeUser inventory-api

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
	serviceId := "s-12345678" // string | The service ID this workflow belongs to.
	environmentId := "se-12345678" // string | The service environment ID this workflow belongs to.
	userId := "user-12345678" // string | The user ID

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.InventoryApiAPI.InventoryApiDescribeUser(context.Background(), serviceId, environmentId, userId).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `InventoryApiAPI.InventoryApiDescribeUser``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `InventoryApiDescribeUser`: FleetDescribeUserResult
	fmt.Fprintf(os.Stdout, "Response from `InventoryApiAPI.InventoryApiDescribeUser`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**serviceId** | **string** | The service ID this workflow belongs to. | 
**environmentId** | **string** | The service environment ID this workflow belongs to. | 
**userId** | **string** | The user ID | 

### Other Parameters

Other parameters are passed through a pointer to a apiInventoryApiDescribeUserRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------




### Return type

[**FleetDescribeUserResult**](FleetDescribeUserResult.md)

### Authorization

[api_key_header_Authorization](../README.md#api_key_header_Authorization)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json, application/vnd.goa.error

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## InventoryApiEnableResourceInstanceManualOverride

> InventoryApiEnableResourceInstanceManualOverride(ctx, serviceId, environmentId, instanceId).Execute()

EnableResourceInstanceManualOverride inventory-api

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
	serviceId := "s-12345678" // string | The service ID this workflow belongs to.
	environmentId := "se-12345678" // string | The service environment ID this workflow belongs to.
	instanceId := "instance-12345678" // string | The resource instance ID.

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	r, err := apiClient.InventoryApiAPI.InventoryApiEnableResourceInstanceManualOverride(context.Background(), serviceId, environmentId, instanceId).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `InventoryApiAPI.InventoryApiEnableResourceInstanceManualOverride``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**serviceId** | **string** | The service ID this workflow belongs to. | 
**environmentId** | **string** | The service environment ID this workflow belongs to. | 
**instanceId** | **string** | The resource instance ID. | 

### Other Parameters

Other parameters are passed through a pointer to a apiInventoryApiEnableResourceInstanceManualOverrideRequest struct via the builder pattern


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


## InventoryApiFailoverResourceInstance

> InventoryApiFailoverResourceInstance(ctx, serviceId, environmentId, instanceId).FleetFailoverResourceInstanceRequest2(fleetFailoverResourceInstanceRequest2).Execute()

FailoverResourceInstance inventory-api

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
	serviceId := "s-12345678" // string | The service ID this workflow belongs to.
	environmentId := "se-12345678" // string | The service environment ID this workflow belongs to.
	instanceId := "instance-12345678" // string | The resource instance ID.
	fleetFailoverResourceInstanceRequest2 := *openapiclient.NewFleetFailoverResourceInstanceRequest2("db-0") // FleetFailoverResourceInstanceRequest2 | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	r, err := apiClient.InventoryApiAPI.InventoryApiFailoverResourceInstance(context.Background(), serviceId, environmentId, instanceId).FleetFailoverResourceInstanceRequest2(fleetFailoverResourceInstanceRequest2).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `InventoryApiAPI.InventoryApiFailoverResourceInstance``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**serviceId** | **string** | The service ID this workflow belongs to. | 
**environmentId** | **string** | The service environment ID this workflow belongs to. | 
**instanceId** | **string** | The resource instance ID. | 

### Other Parameters

Other parameters are passed through a pointer to a apiInventoryApiFailoverResourceInstanceRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------



 **fleetFailoverResourceInstanceRequest2** | [**FleetFailoverResourceInstanceRequest2**](FleetFailoverResourceInstanceRequest2.md) |  | 

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


## InventoryApiGenerateTokenForHostClusterDashboard

> FleetGenerateTokenForHostClusterDashboardResult InventoryApiGenerateTokenForHostClusterDashboard(ctx, serviceId, environmentId, id).Execute()

GenerateTokenForHostClusterDashboard inventory-api

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
	serviceId := "s-12345678" // string | The service ID this workflow belongs to.
	environmentId := "se-12345678" // string | The service environment ID this workflow belongs to.
	id := "hc-12345678" // string | The host cluster ID

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.InventoryApiAPI.InventoryApiGenerateTokenForHostClusterDashboard(context.Background(), serviceId, environmentId, id).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `InventoryApiAPI.InventoryApiGenerateTokenForHostClusterDashboard``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `InventoryApiGenerateTokenForHostClusterDashboard`: FleetGenerateTokenForHostClusterDashboardResult
	fmt.Fprintf(os.Stdout, "Response from `InventoryApiAPI.InventoryApiGenerateTokenForHostClusterDashboard`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**serviceId** | **string** | The service ID this workflow belongs to. | 
**environmentId** | **string** | The service environment ID this workflow belongs to. | 
**id** | **string** | The host cluster ID | 

### Other Parameters

Other parameters are passed through a pointer to a apiInventoryApiGenerateTokenForHostClusterDashboardRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------




### Return type

[**FleetGenerateTokenForHostClusterDashboardResult**](FleetGenerateTokenForHostClusterDashboardResult.md)

### Authorization

[api_key_header_Authorization](../README.md#api_key_header_Authorization)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json, application/vnd.goa.error

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## InventoryApiListActiveOrganizations

> ListOrganizationsResult InventoryApiListActiveOrganizations(ctx, serviceId, environmentId).Execute()

ListActiveOrganizations inventory-api

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
	serviceId := "s-12345678" // string | The service ID this workflow belongs to.
	environmentId := "se-12345678" // string | The service environment ID this workflow belongs to.

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.InventoryApiAPI.InventoryApiListActiveOrganizations(context.Background(), serviceId, environmentId).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `InventoryApiAPI.InventoryApiListActiveOrganizations``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `InventoryApiListActiveOrganizations`: ListOrganizationsResult
	fmt.Fprintf(os.Stdout, "Response from `InventoryApiAPI.InventoryApiListActiveOrganizations`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**serviceId** | **string** | The service ID this workflow belongs to. | 
**environmentId** | **string** | The service environment ID this workflow belongs to. | 

### Other Parameters

Other parameters are passed through a pointer to a apiInventoryApiListActiveOrganizationsRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------



### Return type

[**ListOrganizationsResult**](ListOrganizationsResult.md)

### Authorization

[api_key_header_Authorization](../README.md#api_key_header_Authorization)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json, application/vnd.goa.error

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## InventoryApiListAllUsers

> FleetListAllUsersResult InventoryApiListAllUsers(ctx).ServiceId(serviceId).Execute()

ListAllUsers inventory-api

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
	serviceId := "s-12345678" // string | The service ID of the users (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.InventoryApiAPI.InventoryApiListAllUsers(context.Background()).ServiceId(serviceId).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `InventoryApiAPI.InventoryApiListAllUsers``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `InventoryApiListAllUsers`: FleetListAllUsersResult
	fmt.Fprintf(os.Stdout, "Response from `InventoryApiAPI.InventoryApiListAllUsers`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiInventoryApiListAllUsersRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **serviceId** | **string** | The service ID of the users | 

### Return type

[**FleetListAllUsersResult**](FleetListAllUsersResult.md)

### Authorization

[api_key_header_Authorization](../README.md#api_key_header_Authorization)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json, application/vnd.goa.error

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## InventoryApiListDependentComponents

> FleetListDependentComponentsResult InventoryApiListDependentComponents(ctx, serviceId, environmentId, instanceId).Execute()

ListDependentComponents inventory-api

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
	serviceId := "s-12345678" // string | The service ID this workflow belongs to.
	environmentId := "se-12345678" // string | The service environment ID this workflow belongs to.
	instanceId := "instance-12345678" // string | The resource instance ID.

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.InventoryApiAPI.InventoryApiListDependentComponents(context.Background(), serviceId, environmentId, instanceId).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `InventoryApiAPI.InventoryApiListDependentComponents``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `InventoryApiListDependentComponents`: FleetListDependentComponentsResult
	fmt.Fprintf(os.Stdout, "Response from `InventoryApiAPI.InventoryApiListDependentComponents`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**serviceId** | **string** | The service ID this workflow belongs to. | 
**environmentId** | **string** | The service environment ID this workflow belongs to. | 
**instanceId** | **string** | The resource instance ID. | 

### Other Parameters

Other parameters are passed through a pointer to a apiInventoryApiListDependentComponentsRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------




### Return type

[**FleetListDependentComponentsResult**](FleetListDependentComponentsResult.md)

### Authorization

[api_key_header_Authorization](../README.md#api_key_header_Authorization)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json, application/vnd.goa.error

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## InventoryApiListEligibleInstancesPerUpgrade

> ListEligibleInstancesPerUpgradeResult InventoryApiListEligibleInstancesPerUpgrade(ctx, serviceId, productTierId, upgradePathId).NextPageToken(nextPageToken).Execute()

ListEligibleInstancesPerUpgrade inventory-api

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
	serviceId := "s-12345678" // string | The service ID this workflow belongs to.
	productTierId := "pt-12345678" // string | The product tier ID that this upgrade path belongs to
	upgradePathId := "up-12345678" // string | The upgrade path ID
	nextPageToken := "next-token" // string | The next page token to use for pagination (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.InventoryApiAPI.InventoryApiListEligibleInstancesPerUpgrade(context.Background(), serviceId, productTierId, upgradePathId).NextPageToken(nextPageToken).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `InventoryApiAPI.InventoryApiListEligibleInstancesPerUpgrade``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `InventoryApiListEligibleInstancesPerUpgrade`: ListEligibleInstancesPerUpgradeResult
	fmt.Fprintf(os.Stdout, "Response from `InventoryApiAPI.InventoryApiListEligibleInstancesPerUpgrade`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**serviceId** | **string** | The service ID this workflow belongs to. | 
**productTierId** | **string** | The product tier ID that this upgrade path belongs to | 
**upgradePathId** | **string** | The upgrade path ID | 

### Other Parameters

Other parameters are passed through a pointer to a apiInventoryApiListEligibleInstancesPerUpgradeRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------



 **nextPageToken** | **string** | The next page token to use for pagination | 

### Return type

[**ListEligibleInstancesPerUpgradeResult**](ListEligibleInstancesPerUpgradeResult.md)

### Authorization

[api_key_header_Authorization](../README.md#api_key_header_Authorization)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json, application/vnd.goa.error

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## InventoryApiListHostClusters

> FleetListHostClustersResult InventoryApiListHostClusters(ctx, serviceId, environmentId).Execute()

ListHostClusters inventory-api

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
	serviceId := "s-12345678" // string | The service ID this workflow belongs to.
	environmentId := "se-12345678" // string | The service environment ID this workflow belongs to.

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.InventoryApiAPI.InventoryApiListHostClusters(context.Background(), serviceId, environmentId).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `InventoryApiAPI.InventoryApiListHostClusters``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `InventoryApiListHostClusters`: FleetListHostClustersResult
	fmt.Fprintf(os.Stdout, "Response from `InventoryApiAPI.InventoryApiListHostClusters`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**serviceId** | **string** | The service ID this workflow belongs to. | 
**environmentId** | **string** | The service environment ID this workflow belongs to. | 

### Other Parameters

Other parameters are passed through a pointer to a apiInventoryApiListHostClustersRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------



### Return type

[**FleetListHostClustersResult**](FleetListHostClustersResult.md)

### Authorization

[api_key_header_Authorization](../README.md#api_key_header_Authorization)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json, application/vnd.goa.error

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## InventoryApiListInstanceEvents

> FleetListEventsResult InventoryApiListInstanceEvents(ctx, serviceId, environmentId, instanceId).StartTime(startTime).EndTime(endTime).Execute()

ListInstanceEvents inventory-api

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
	serviceId := "s-12345678" // string | The service ID this workflow belongs to.
	environmentId := "se-12345678" // string | The service environment ID this workflow belongs to.
	instanceId := "instance-12345678" // string | The resource instance ID.
	startTime := "2023-01-10T00:00:00Z" // string | Filter events that occurred after this time (optional)
	endTime := "2023-01-10T00:00:00Z" // string | Filter events that occurred before this time (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.InventoryApiAPI.InventoryApiListInstanceEvents(context.Background(), serviceId, environmentId, instanceId).StartTime(startTime).EndTime(endTime).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `InventoryApiAPI.InventoryApiListInstanceEvents``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `InventoryApiListInstanceEvents`: FleetListEventsResult
	fmt.Fprintf(os.Stdout, "Response from `InventoryApiAPI.InventoryApiListInstanceEvents`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**serviceId** | **string** | The service ID this workflow belongs to. | 
**environmentId** | **string** | The service environment ID this workflow belongs to. | 
**instanceId** | **string** | The resource instance ID. | 

### Other Parameters

Other parameters are passed through a pointer to a apiInventoryApiListInstanceEventsRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------



 **startTime** | **string** | Filter events that occurred after this time | 
 **endTime** | **string** | Filter events that occurred before this time | 

### Return type

[**FleetListEventsResult**](FleetListEventsResult.md)

### Authorization

[api_key_header_Authorization](../README.md#api_key_header_Authorization)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json, application/vnd.goa.error

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## InventoryApiListLinkedInstances

> FleetListLinkedInstancesResult InventoryApiListLinkedInstances(ctx, serviceId, environmentId, instanceId).Execute()

ListLinkedInstances inventory-api

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
	serviceId := "s-12345678" // string | The service ID this workflow belongs to.
	environmentId := "se-12345678" // string | The service environment ID this workflow belongs to.
	instanceId := "instance-12345678" // string | The resource instance ID.

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.InventoryApiAPI.InventoryApiListLinkedInstances(context.Background(), serviceId, environmentId, instanceId).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `InventoryApiAPI.InventoryApiListLinkedInstances``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `InventoryApiListLinkedInstances`: FleetListLinkedInstancesResult
	fmt.Fprintf(os.Stdout, "Response from `InventoryApiAPI.InventoryApiListLinkedInstances`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**serviceId** | **string** | The service ID this workflow belongs to. | 
**environmentId** | **string** | The service environment ID this workflow belongs to. | 
**instanceId** | **string** | The resource instance ID. | 

### Other Parameters

Other parameters are passed through a pointer to a apiInventoryApiListLinkedInstancesRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------




### Return type

[**FleetListLinkedInstancesResult**](FleetListLinkedInstancesResult.md)

### Authorization

[api_key_header_Authorization](../README.md#api_key_header_Authorization)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json, application/vnd.goa.error

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## InventoryApiListResourceInstanceSnapshots

> FleetListInstanceSnapshotResult InventoryApiListResourceInstanceSnapshots(ctx, serviceId, environmentId, instanceId).Execute()

ListResourceInstanceSnapshots inventory-api

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
	serviceId := "s-12345678" // string | The service ID this workflow belongs to.
	environmentId := "se-12345678" // string | The service environment ID this workflow belongs to.
	instanceId := "instance-12345678" // string | The resource instance ID.

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.InventoryApiAPI.InventoryApiListResourceInstanceSnapshots(context.Background(), serviceId, environmentId, instanceId).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `InventoryApiAPI.InventoryApiListResourceInstanceSnapshots``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `InventoryApiListResourceInstanceSnapshots`: FleetListInstanceSnapshotResult
	fmt.Fprintf(os.Stdout, "Response from `InventoryApiAPI.InventoryApiListResourceInstanceSnapshots`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**serviceId** | **string** | The service ID this workflow belongs to. | 
**environmentId** | **string** | The service environment ID this workflow belongs to. | 
**instanceId** | **string** | The resource instance ID. | 

### Other Parameters

Other parameters are passed through a pointer to a apiInventoryApiListResourceInstanceSnapshotsRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------




### Return type

[**FleetListInstanceSnapshotResult**](FleetListInstanceSnapshotResult.md)

### Authorization

[api_key_header_Authorization](../README.md#api_key_header_Authorization)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json, application/vnd.goa.error

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## InventoryApiListResourceInstances

> ListFleetResourceInstancesResultInternal InventoryApiListResourceInstances(ctx, serviceId, environmentId).ProductTierVersion(productTierVersion).ProductTierId(productTierId).SubscriptionId(subscriptionId).Execute()

ListResourceInstances inventory-api

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
	serviceId := "s-12345678" // string | The service ID this workflow belongs to.
	environmentId := "se-12345678" // string | The service environment ID this workflow belongs to.
	productTierVersion := "Ab nemo maiores voluptates." // string | Product tier version of the instance to describe. If not specified, the latest version is described. (optional)
	productTierId := "Et et dolorem harum." // string | Product tier id of the instance to describe. Needs to specified in combination with the product tier version (optional)
	subscriptionId := "Laboriosam dolorem dignissimos quasi ex totam expedita." // string | Subscription id of the instance to describe. (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.InventoryApiAPI.InventoryApiListResourceInstances(context.Background(), serviceId, environmentId).ProductTierVersion(productTierVersion).ProductTierId(productTierId).SubscriptionId(subscriptionId).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `InventoryApiAPI.InventoryApiListResourceInstances``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `InventoryApiListResourceInstances`: ListFleetResourceInstancesResultInternal
	fmt.Fprintf(os.Stdout, "Response from `InventoryApiAPI.InventoryApiListResourceInstances`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**serviceId** | **string** | The service ID this workflow belongs to. | 
**environmentId** | **string** | The service environment ID this workflow belongs to. | 

### Other Parameters

Other parameters are passed through a pointer to a apiInventoryApiListResourceInstancesRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


 **productTierVersion** | **string** | Product tier version of the instance to describe. If not specified, the latest version is described. | 
 **productTierId** | **string** | Product tier id of the instance to describe. Needs to specified in combination with the product tier version | 
 **subscriptionId** | **string** | Subscription id of the instance to describe. | 

### Return type

[**ListFleetResourceInstancesResultInternal**](ListFleetResourceInstancesResultInternal.md)

### Authorization

[api_key_header_Authorization](../README.md#api_key_header_Authorization)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json, application/vnd.goa.error

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## InventoryApiListResources

> ListResourcesResult InventoryApiListResources(ctx, serviceId, environmentId, serviceModelId, productTierId).ListResourcesRequest2(listResourcesRequest2).Execute()

ListResources inventory-api

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
	serviceId := "s-12345678" // string | The service ID this workflow belongs to.
	environmentId := "se-12345678" // string | The service environment ID this workflow belongs to.
	serviceModelId := "sm-12345678" // string | The service model ID.
	productTierId := "pt-12345678" // string | The product tier ID.
	listResourcesRequest2 := *openapiclient.NewListResourcesRequest2() // ListResourcesRequest2 | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.InventoryApiAPI.InventoryApiListResources(context.Background(), serviceId, environmentId, serviceModelId, productTierId).ListResourcesRequest2(listResourcesRequest2).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `InventoryApiAPI.InventoryApiListResources``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `InventoryApiListResources`: ListResourcesResult
	fmt.Fprintf(os.Stdout, "Response from `InventoryApiAPI.InventoryApiListResources`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**serviceId** | **string** | The service ID this workflow belongs to. | 
**environmentId** | **string** | The service environment ID this workflow belongs to. | 
**serviceModelId** | **string** | The service model ID. | 
**productTierId** | **string** | The product tier ID. | 

### Other Parameters

Other parameters are passed through a pointer to a apiInventoryApiListResourcesRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------




 **listResourcesRequest2** | [**ListResourcesRequest2**](ListResourcesRequest2.md) |  | 

### Return type

[**ListResourcesResult**](ListResourcesResult.md)

### Authorization

[api_key_header_Authorization](../README.md#api_key_header_Authorization)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json, application/vnd.goa.error

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## InventoryApiListServiceOfferings

> InventoryListServiceOfferingsResult InventoryApiListServiceOfferings(ctx).OrgId(orgId).Visibility(visibility).Execute()

ListServiceOfferings inventory-api

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
	orgId := "org-12345678" // string | Org Id (optional)
	visibility := "PRIVATE" // string | The visibility of service offering (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.InventoryApiAPI.InventoryApiListServiceOfferings(context.Background()).OrgId(orgId).Visibility(visibility).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `InventoryApiAPI.InventoryApiListServiceOfferings``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `InventoryApiListServiceOfferings`: InventoryListServiceOfferingsResult
	fmt.Fprintf(os.Stdout, "Response from `InventoryApiAPI.InventoryApiListServiceOfferings`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiInventoryApiListServiceOfferingsRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **orgId** | **string** | Org Id | 
 **visibility** | **string** | The visibility of service offering | 

### Return type

[**InventoryListServiceOfferingsResult**](InventoryListServiceOfferingsResult.md)

### Authorization

[api_key_header_Authorization](../README.md#api_key_header_Authorization)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json, application/vnd.goa.error

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## InventoryApiListServicesOrchestrations

> []FleetDescribeServicesOrchestrationResult InventoryApiListServicesOrchestrations(ctx).EnvironmentType(environmentType).Execute()

ListServicesOrchestrations inventory-api

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
	environmentType := "PROD|PRIVATE|CANARY|STAGING|QA|DEV" // string | The environment type

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.InventoryApiAPI.InventoryApiListServicesOrchestrations(context.Background()).EnvironmentType(environmentType).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `InventoryApiAPI.InventoryApiListServicesOrchestrations``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `InventoryApiListServicesOrchestrations`: []FleetDescribeServicesOrchestrationResult
	fmt.Fprintf(os.Stdout, "Response from `InventoryApiAPI.InventoryApiListServicesOrchestrations`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiInventoryApiListServicesOrchestrationsRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **environmentType** | **string** | The environment type | 

### Return type

[**[]FleetDescribeServicesOrchestrationResult**](FleetDescribeServicesOrchestrationResult.md)

### Authorization

[api_key_header_Authorization](../README.md#api_key_header_Authorization)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json, application/vnd.goa.error

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## InventoryApiListSubscription

> FleetListSubscriptionsResult InventoryApiListSubscription(ctx, serviceId, environmentId).Execute()

ListSubscription inventory-api

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
	serviceId := "s-12345678" // string | The service ID this workflow belongs to.
	environmentId := "se-12345678" // string | The service environment ID this workflow belongs to.

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.InventoryApiAPI.InventoryApiListSubscription(context.Background(), serviceId, environmentId).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `InventoryApiAPI.InventoryApiListSubscription``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `InventoryApiListSubscription`: FleetListSubscriptionsResult
	fmt.Fprintf(os.Stdout, "Response from `InventoryApiAPI.InventoryApiListSubscription`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**serviceId** | **string** | The service ID this workflow belongs to. | 
**environmentId** | **string** | The service environment ID this workflow belongs to. | 

### Other Parameters

Other parameters are passed through a pointer to a apiInventoryApiListSubscriptionRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------



### Return type

[**FleetListSubscriptionsResult**](FleetListSubscriptionsResult.md)

### Authorization

[api_key_header_Authorization](../README.md#api_key_header_Authorization)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json, application/vnd.goa.error

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## InventoryApiListSubscriptionRequests

> ListSubscriptionRequestsResult InventoryApiListSubscriptionRequests(ctx, serviceId, environmentId).Status(status).Execute()

ListSubscriptionRequests inventory-api

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
	serviceId := "s-12345678" // string | The service ID this workflow belongs to.
	environmentId := "se-12345678" // string | The service environment ID this workflow belongs to.
	status := "PENDING" // string | The status of the subscription request to filter by (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.InventoryApiAPI.InventoryApiListSubscriptionRequests(context.Background(), serviceId, environmentId).Status(status).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `InventoryApiAPI.InventoryApiListSubscriptionRequests``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `InventoryApiListSubscriptionRequests`: ListSubscriptionRequestsResult
	fmt.Fprintf(os.Stdout, "Response from `InventoryApiAPI.InventoryApiListSubscriptionRequests`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**serviceId** | **string** | The service ID this workflow belongs to. | 
**environmentId** | **string** | The service environment ID this workflow belongs to. | 

### Other Parameters

Other parameters are passed through a pointer to a apiInventoryApiListSubscriptionRequestsRequest struct via the builder pattern


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


## InventoryApiListUpgradePaths

> ListUpgradePathsResult InventoryApiListUpgradePaths(ctx, serviceId, productTierId).SourceProductTierVersion(sourceProductTierVersion).TargetProductTierVersion(targetProductTierVersion).Status(status).Type_(type_).NextPageToken(nextPageToken).Execute()

ListUpgradePaths inventory-api

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
	serviceId := "s-12345678" // string | The service ID this workflow belongs to.
	productTierId := "pt-12345678" // string | The product tier ID that this upgrade path belongs to
	sourceProductTierVersion := "1.0" // string | The source product tier version to list upgrade paths for (optional)
	targetProductTierVersion := "2.0" // string | The target product tier version to list upgrade paths for (optional)
	status := "COMPLETE" // string | The status of the upgrade path to filter by (optional)
	type_ := "Major" // string | The type of the upgrade path to filter by (optional)
	nextPageToken := "next-token" // string | The next page token to use for pagination (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.InventoryApiAPI.InventoryApiListUpgradePaths(context.Background(), serviceId, productTierId).SourceProductTierVersion(sourceProductTierVersion).TargetProductTierVersion(targetProductTierVersion).Status(status).Type_(type_).NextPageToken(nextPageToken).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `InventoryApiAPI.InventoryApiListUpgradePaths``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `InventoryApiListUpgradePaths`: ListUpgradePathsResult
	fmt.Fprintf(os.Stdout, "Response from `InventoryApiAPI.InventoryApiListUpgradePaths`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**serviceId** | **string** | The service ID this workflow belongs to. | 
**productTierId** | **string** | The product tier ID that this upgrade path belongs to | 

### Other Parameters

Other parameters are passed through a pointer to a apiInventoryApiListUpgradePathsRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


 **sourceProductTierVersion** | **string** | The source product tier version to list upgrade paths for | 
 **targetProductTierVersion** | **string** | The target product tier version to list upgrade paths for | 
 **status** | **string** | The status of the upgrade path to filter by | 
 **type_** | **string** | The type of the upgrade path to filter by | 
 **nextPageToken** | **string** | The next page token to use for pagination | 

### Return type

[**ListUpgradePathsResult**](ListUpgradePathsResult.md)

### Authorization

[api_key_header_Authorization](../README.md#api_key_header_Authorization)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json, application/vnd.goa.error

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## InventoryApiListUsers

> FleetListUsersResult InventoryApiListUsers(ctx, serviceId, environmentId).SubscriptionId(subscriptionId).Execute()

ListUsers inventory-api

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
	serviceId := "s-12345678" // string | The service ID this workflow belongs to.
	environmentId := "se-12345678" // string | The service environment ID this workflow belongs to.
	subscriptionId := "sub-12345678" // string | The subscription ID of the user (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.InventoryApiAPI.InventoryApiListUsers(context.Background(), serviceId, environmentId).SubscriptionId(subscriptionId).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `InventoryApiAPI.InventoryApiListUsers``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `InventoryApiListUsers`: FleetListUsersResult
	fmt.Fprintf(os.Stdout, "Response from `InventoryApiAPI.InventoryApiListUsers`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**serviceId** | **string** | The service ID this workflow belongs to. | 
**environmentId** | **string** | The service environment ID this workflow belongs to. | 

### Other Parameters

Other parameters are passed through a pointer to a apiInventoryApiListUsersRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


 **subscriptionId** | **string** | The subscription ID of the user | 

### Return type

[**FleetListUsersResult**](FleetListUsersResult.md)

### Authorization

[api_key_header_Authorization](../README.md#api_key_header_Authorization)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json, application/vnd.goa.error

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## InventoryApiModifyServicesOrchestration

> InventoryApiModifyServicesOrchestration(ctx, id).FleetModifyServicesOrchestrationRequest2(fleetModifyServicesOrchestrationRequest2).Execute()

ModifyServicesOrchestration inventory-api

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
	id := "so-12345678" // string | The ID of the services orchestration
	fleetModifyServicesOrchestrationRequest2 := *openapiclient.NewFleetModifyServicesOrchestrationRequest2("Voluptatem dolorem.") // FleetModifyServicesOrchestrationRequest2 | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	r, err := apiClient.InventoryApiAPI.InventoryApiModifyServicesOrchestration(context.Background(), id).FleetModifyServicesOrchestrationRequest2(fleetModifyServicesOrchestrationRequest2).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `InventoryApiAPI.InventoryApiModifyServicesOrchestration``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **string** | The ID of the services orchestration | 

### Other Parameters

Other parameters are passed through a pointer to a apiInventoryApiModifyServicesOrchestrationRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **fleetModifyServicesOrchestrationRequest2** | [**FleetModifyServicesOrchestrationRequest2**](FleetModifyServicesOrchestrationRequest2.md) |  | 

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


## InventoryApiRemoveCapacityFromResourceInstance

> InventoryApiRemoveCapacityFromResourceInstance(ctx, serviceId, environmentId, instanceId).FleetRemoveCapacityFromResourceInstanceRequest2(fleetRemoveCapacityFromResourceInstanceRequest2).Execute()

RemoveCapacityFromResourceInstance inventory-api

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
	serviceId := "s-12345678" // string | The service ID this workflow belongs to.
	environmentId := "se-12345678" // string | The service environment ID this workflow belongs to.
	instanceId := "instance-12345678" // string | The resource instance ID.
	fleetRemoveCapacityFromResourceInstanceRequest2 := *openapiclient.NewFleetRemoveCapacityFromResourceInstanceRequest2(int64(3), "r-12345678") // FleetRemoveCapacityFromResourceInstanceRequest2 | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	r, err := apiClient.InventoryApiAPI.InventoryApiRemoveCapacityFromResourceInstance(context.Background(), serviceId, environmentId, instanceId).FleetRemoveCapacityFromResourceInstanceRequest2(fleetRemoveCapacityFromResourceInstanceRequest2).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `InventoryApiAPI.InventoryApiRemoveCapacityFromResourceInstance``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**serviceId** | **string** | The service ID this workflow belongs to. | 
**environmentId** | **string** | The service environment ID this workflow belongs to. | 
**instanceId** | **string** | The resource instance ID. | 

### Other Parameters

Other parameters are passed through a pointer to a apiInventoryApiRemoveCapacityFromResourceInstanceRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------



 **fleetRemoveCapacityFromResourceInstanceRequest2** | [**FleetRemoveCapacityFromResourceInstanceRequest2**](FleetRemoveCapacityFromResourceInstanceRequest2.md) |  | 

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


## InventoryApiRemoveCustomDNSFromResourceInstance

> InventoryApiRemoveCustomDNSFromResourceInstance(ctx, serviceId, environmentId, resourceKey, instanceId).Execute()

RemoveCustomDNSFromResourceInstance inventory-api

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
	serviceId := "s-12345678" // string | The service ID this workflow belongs to.
	environmentId := "se-12345678" // string | The service environment ID this workflow belongs to.
	resourceKey := "http-service" // string | The resource key
	instanceId := "instance-12345678" // string | The resource instance ID.

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	r, err := apiClient.InventoryApiAPI.InventoryApiRemoveCustomDNSFromResourceInstance(context.Background(), serviceId, environmentId, resourceKey, instanceId).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `InventoryApiAPI.InventoryApiRemoveCustomDNSFromResourceInstance``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**serviceId** | **string** | The service ID this workflow belongs to. | 
**environmentId** | **string** | The service environment ID this workflow belongs to. | 
**resourceKey** | **string** | The resource key | 
**instanceId** | **string** | The resource instance ID. | 

### Other Parameters

Other parameters are passed through a pointer to a apiInventoryApiRemoveCustomDNSFromResourceInstanceRequest struct via the builder pattern


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


## InventoryApiRestartResourceInstance

> InventoryApiRestartResourceInstance(ctx, serviceId, environmentId, instanceId).FleetRestartResourceInstanceRequest2(fleetRestartResourceInstanceRequest2).Execute()

RestartResourceInstance inventory-api

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
	serviceId := "s-12345678" // string | The service ID this workflow belongs to.
	environmentId := "se-12345678" // string | The service environment ID this workflow belongs to.
	instanceId := "instance-12345678" // string | The resource instance ID.
	fleetRestartResourceInstanceRequest2 := *openapiclient.NewFleetRestartResourceInstanceRequest2("r-12345678") // FleetRestartResourceInstanceRequest2 | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	r, err := apiClient.InventoryApiAPI.InventoryApiRestartResourceInstance(context.Background(), serviceId, environmentId, instanceId).FleetRestartResourceInstanceRequest2(fleetRestartResourceInstanceRequest2).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `InventoryApiAPI.InventoryApiRestartResourceInstance``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**serviceId** | **string** | The service ID this workflow belongs to. | 
**environmentId** | **string** | The service environment ID this workflow belongs to. | 
**instanceId** | **string** | The resource instance ID. | 

### Other Parameters

Other parameters are passed through a pointer to a apiInventoryApiRestartResourceInstanceRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------



 **fleetRestartResourceInstanceRequest2** | [**FleetRestartResourceInstanceRequest2**](FleetRestartResourceInstanceRequest2.md) |  | 

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


## InventoryApiRestoreResourceInstance

> CreateResourceInstanceResponseBody InventoryApiRestoreResourceInstance(ctx, serviceId, environmentId, instanceId).FleetRestoreResourceInstanceRequest2(fleetRestoreResourceInstanceRequest2).Execute()

RestoreResourceInstance inventory-api

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
	serviceId := "s-12345678" // string | The service ID this workflow belongs to.
	environmentId := "se-12345678" // string | The service environment ID this workflow belongs to.
	instanceId := "instance-12345678" // string | The resource instance ID.
	fleetRestoreResourceInstanceRequest2 := *openapiclient.NewFleetRestoreResourceInstanceRequest2("2021-09-01T00:00:00Z") // FleetRestoreResourceInstanceRequest2 | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.InventoryApiAPI.InventoryApiRestoreResourceInstance(context.Background(), serviceId, environmentId, instanceId).FleetRestoreResourceInstanceRequest2(fleetRestoreResourceInstanceRequest2).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `InventoryApiAPI.InventoryApiRestoreResourceInstance``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `InventoryApiRestoreResourceInstance`: CreateResourceInstanceResponseBody
	fmt.Fprintf(os.Stdout, "Response from `InventoryApiAPI.InventoryApiRestoreResourceInstance`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**serviceId** | **string** | The service ID this workflow belongs to. | 
**environmentId** | **string** | The service environment ID this workflow belongs to. | 
**instanceId** | **string** | The resource instance ID. | 

### Other Parameters

Other parameters are passed through a pointer to a apiInventoryApiRestoreResourceInstanceRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------



 **fleetRestoreResourceInstanceRequest2** | [**FleetRestoreResourceInstanceRequest2**](FleetRestoreResourceInstanceRequest2.md) |  | 

### Return type

[**CreateResourceInstanceResponseBody**](CreateResourceInstanceResponseBody.md)

### Authorization

[api_key_header_Authorization](../README.md#api_key_header_Authorization)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json, application/vnd.goa.error

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## InventoryApiRestoreResourceInstanceFromSnapshot

> CreateResourceInstanceResponseBody InventoryApiRestoreResourceInstanceFromSnapshot(ctx, serviceId, environmentId, snapshotId).FleetRestoreResourceInstanceFromSnapshotRequest2(fleetRestoreResourceInstanceFromSnapshotRequest2).Execute()

RestoreResourceInstanceFromSnapshot inventory-api

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
	serviceId := "s-12345678" // string | The service ID this workflow belongs to.
	environmentId := "se-12345678" // string | The service environment ID this workflow belongs to.
	snapshotId := "instance-ss-12345678" // string | The snapshot ID
	fleetRestoreResourceInstanceFromSnapshotRequest2 := *openapiclient.NewFleetRestoreResourceInstanceFromSnapshotRequest2() // FleetRestoreResourceInstanceFromSnapshotRequest2 | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.InventoryApiAPI.InventoryApiRestoreResourceInstanceFromSnapshot(context.Background(), serviceId, environmentId, snapshotId).FleetRestoreResourceInstanceFromSnapshotRequest2(fleetRestoreResourceInstanceFromSnapshotRequest2).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `InventoryApiAPI.InventoryApiRestoreResourceInstanceFromSnapshot``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `InventoryApiRestoreResourceInstanceFromSnapshot`: CreateResourceInstanceResponseBody
	fmt.Fprintf(os.Stdout, "Response from `InventoryApiAPI.InventoryApiRestoreResourceInstanceFromSnapshot`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**serviceId** | **string** | The service ID this workflow belongs to. | 
**environmentId** | **string** | The service environment ID this workflow belongs to. | 
**snapshotId** | **string** | The snapshot ID | 

### Other Parameters

Other parameters are passed through a pointer to a apiInventoryApiRestoreResourceInstanceFromSnapshotRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------



 **fleetRestoreResourceInstanceFromSnapshotRequest2** | [**FleetRestoreResourceInstanceFromSnapshotRequest2**](FleetRestoreResourceInstanceFromSnapshotRequest2.md) |  | 

### Return type

[**CreateResourceInstanceResponseBody**](CreateResourceInstanceResponseBody.md)

### Authorization

[api_key_header_Authorization](../README.md#api_key_header_Authorization)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json, application/vnd.goa.error

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## InventoryApiResumeSubscription

> InventoryApiResumeSubscription(ctx, serviceId, environmentId, id).Execute()

ResumeSubscription inventory-api

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
	serviceId := "s-12345678" // string | The service ID this workflow belongs to.
	environmentId := "se-12345678" // string | The service environment ID this workflow belongs to.
	id := "sub-12345678" // string | The subscription ID

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	r, err := apiClient.InventoryApiAPI.InventoryApiResumeSubscription(context.Background(), serviceId, environmentId, id).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `InventoryApiAPI.InventoryApiResumeSubscription``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**serviceId** | **string** | The service ID this workflow belongs to. | 
**environmentId** | **string** | The service environment ID this workflow belongs to. | 
**id** | **string** | The subscription ID | 

### Other Parameters

Other parameters are passed through a pointer to a apiInventoryApiResumeSubscriptionRequest struct via the builder pattern


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


## InventoryApiSearchInventory

> SearchInventoryResult InventoryApiSearchInventory(ctx).SearchInventoryRequest2(searchInventoryRequest2).Execute()

SearchInventory inventory-api

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
	searchInventoryRequest2 := *openapiclient.NewSearchInventoryRequest2("foo") // SearchInventoryRequest2 | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.InventoryApiAPI.InventoryApiSearchInventory(context.Background()).SearchInventoryRequest2(searchInventoryRequest2).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `InventoryApiAPI.InventoryApiSearchInventory``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `InventoryApiSearchInventory`: SearchInventoryResult
	fmt.Fprintf(os.Stdout, "Response from `InventoryApiAPI.InventoryApiSearchInventory`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiInventoryApiSearchInventoryRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **searchInventoryRequest2** | [**SearchInventoryRequest2**](SearchInventoryRequest2.md) |  | 

### Return type

[**SearchInventoryResult**](SearchInventoryResult.md)

### Authorization

[api_key_header_Authorization](../README.md#api_key_header_Authorization)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json, application/vnd.goa.error

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## InventoryApiSearchServiceInventory

> SearchServiceInventoryResult InventoryApiSearchServiceInventory(ctx, serviceId, environmentId).SearchServiceInventoryRequest2(searchServiceInventoryRequest2).Execute()

SearchServiceInventory inventory-api

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
	serviceId := "s-12345678" // string | The service ID this workflow belongs to.
	environmentId := "se-12345678" // string | The service environment ID this workflow belongs to.
	searchServiceInventoryRequest2 := *openapiclient.NewSearchServiceInventoryRequest2("foo") // SearchServiceInventoryRequest2 | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.InventoryApiAPI.InventoryApiSearchServiceInventory(context.Background(), serviceId, environmentId).SearchServiceInventoryRequest2(searchServiceInventoryRequest2).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `InventoryApiAPI.InventoryApiSearchServiceInventory``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `InventoryApiSearchServiceInventory`: SearchServiceInventoryResult
	fmt.Fprintf(os.Stdout, "Response from `InventoryApiAPI.InventoryApiSearchServiceInventory`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**serviceId** | **string** | The service ID this workflow belongs to. | 
**environmentId** | **string** | The service environment ID this workflow belongs to. | 

### Other Parameters

Other parameters are passed through a pointer to a apiInventoryApiSearchServiceInventoryRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


 **searchServiceInventoryRequest2** | [**SearchServiceInventoryRequest2**](SearchServiceInventoryRequest2.md) |  | 

### Return type

[**SearchServiceInventoryResult**](SearchServiceInventoryResult.md)

### Authorization

[api_key_header_Authorization](../README.md#api_key_header_Authorization)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json, application/vnd.goa.error

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## InventoryApiStartResourceInstance

> InventoryApiStartResourceInstance(ctx, serviceId, environmentId, instanceId).FleetStartResourceInstanceRequest2(fleetStartResourceInstanceRequest2).Execute()

StartResourceInstance inventory-api

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
	serviceId := "s-12345678" // string | The service ID this workflow belongs to.
	environmentId := "se-12345678" // string | The service environment ID this workflow belongs to.
	instanceId := "instance-12345678" // string | The resource instance ID.
	fleetStartResourceInstanceRequest2 := *openapiclient.NewFleetStartResourceInstanceRequest2("r-12345678") // FleetStartResourceInstanceRequest2 | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	r, err := apiClient.InventoryApiAPI.InventoryApiStartResourceInstance(context.Background(), serviceId, environmentId, instanceId).FleetStartResourceInstanceRequest2(fleetStartResourceInstanceRequest2).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `InventoryApiAPI.InventoryApiStartResourceInstance``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**serviceId** | **string** | The service ID this workflow belongs to. | 
**environmentId** | **string** | The service environment ID this workflow belongs to. | 
**instanceId** | **string** | The resource instance ID. | 

### Other Parameters

Other parameters are passed through a pointer to a apiInventoryApiStartResourceInstanceRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------



 **fleetStartResourceInstanceRequest2** | [**FleetStartResourceInstanceRequest2**](FleetStartResourceInstanceRequest2.md) |  | 

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


## InventoryApiStopResourceInstance

> InventoryApiStopResourceInstance(ctx, serviceId, environmentId, instanceId).FleetStopResourceInstanceRequest2(fleetStopResourceInstanceRequest2).Execute()

StopResourceInstance inventory-api

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
	serviceId := "s-12345678" // string | The service ID this workflow belongs to.
	environmentId := "se-12345678" // string | The service environment ID this workflow belongs to.
	instanceId := "instance-12345678" // string | The resource instance ID.
	fleetStopResourceInstanceRequest2 := *openapiclient.NewFleetStopResourceInstanceRequest2("r-12345678") // FleetStopResourceInstanceRequest2 | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	r, err := apiClient.InventoryApiAPI.InventoryApiStopResourceInstance(context.Background(), serviceId, environmentId, instanceId).FleetStopResourceInstanceRequest2(fleetStopResourceInstanceRequest2).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `InventoryApiAPI.InventoryApiStopResourceInstance``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**serviceId** | **string** | The service ID this workflow belongs to. | 
**environmentId** | **string** | The service environment ID this workflow belongs to. | 
**instanceId** | **string** | The resource instance ID. | 

### Other Parameters

Other parameters are passed through a pointer to a apiInventoryApiStopResourceInstanceRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------



 **fleetStopResourceInstanceRequest2** | [**FleetStopResourceInstanceRequest2**](FleetStopResourceInstanceRequest2.md) |  | 

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


## InventoryApiSuspendSubscription

> InventoryApiSuspendSubscription(ctx, serviceId, environmentId, id).Execute()

SuspendSubscription inventory-api

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
	serviceId := "s-12345678" // string | The service ID this workflow belongs to.
	environmentId := "se-12345678" // string | The service environment ID this workflow belongs to.
	id := "sub-12345678" // string | The subscription ID

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	r, err := apiClient.InventoryApiAPI.InventoryApiSuspendSubscription(context.Background(), serviceId, environmentId, id).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `InventoryApiAPI.InventoryApiSuspendSubscription``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**serviceId** | **string** | The service ID this workflow belongs to. | 
**environmentId** | **string** | The service environment ID this workflow belongs to. | 
**id** | **string** | The subscription ID | 

### Other Parameters

Other parameters are passed through a pointer to a apiInventoryApiSuspendSubscriptionRequest struct via the builder pattern


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


## InventoryApiSuspendUser

> InventoryApiSuspendUser(ctx, userId).Execute()

SuspendUser inventory-api

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
	userId := "user-12345678" // string | The user ID

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	r, err := apiClient.InventoryApiAPI.InventoryApiSuspendUser(context.Background(), userId).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `InventoryApiAPI.InventoryApiSuspendUser``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**userId** | **string** | The user ID | 

### Other Parameters

Other parameters are passed through a pointer to a apiInventoryApiSuspendUserRequest struct via the builder pattern


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


## InventoryApiTerminateSubscription

> InventoryApiTerminateSubscription(ctx, serviceId, environmentId, id).Execute()

TerminateSubscription inventory-api

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
	serviceId := "s-12345678" // string | The service ID this workflow belongs to.
	environmentId := "se-12345678" // string | The service environment ID this workflow belongs to.
	id := "sub-12345678" // string | The subscription ID

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	r, err := apiClient.InventoryApiAPI.InventoryApiTerminateSubscription(context.Background(), serviceId, environmentId, id).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `InventoryApiAPI.InventoryApiTerminateSubscription``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**serviceId** | **string** | The service ID this workflow belongs to. | 
**environmentId** | **string** | The service environment ID this workflow belongs to. | 
**id** | **string** | The subscription ID | 

### Other Parameters

Other parameters are passed through a pointer to a apiInventoryApiTerminateSubscriptionRequest struct via the builder pattern


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


## InventoryApiUnsuspendUser

> InventoryApiUnsuspendUser(ctx, userId).Execute()

UnsuspendUser inventory-api

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
	userId := "user-12345678" // string | The user ID

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	r, err := apiClient.InventoryApiAPI.InventoryApiUnsuspendUser(context.Background(), userId).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `InventoryApiAPI.InventoryApiUnsuspendUser``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**userId** | **string** | The user ID | 

### Other Parameters

Other parameters are passed through a pointer to a apiInventoryApiUnsuspendUserRequest struct via the builder pattern


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


## InventoryApiUpdateResourceInstance

> InventoryApiUpdateResourceInstance(ctx, serviceId, environmentId, instanceId).FleetUpdateResourceInstanceRequest2(fleetUpdateResourceInstanceRequest2).Execute()

UpdateResourceInstance inventory-api

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
	serviceId := "s-12345678" // string | The service ID this workflow belongs to.
	environmentId := "se-12345678" // string | The service environment ID this workflow belongs to.
	instanceId := "instance-12345678" // string | The resource instance ID.
	fleetUpdateResourceInstanceRequest2 := *openapiclient.NewFleetUpdateResourceInstanceRequest2("r-12345678") // FleetUpdateResourceInstanceRequest2 | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	r, err := apiClient.InventoryApiAPI.InventoryApiUpdateResourceInstance(context.Background(), serviceId, environmentId, instanceId).FleetUpdateResourceInstanceRequest2(fleetUpdateResourceInstanceRequest2).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `InventoryApiAPI.InventoryApiUpdateResourceInstance``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**serviceId** | **string** | The service ID this workflow belongs to. | 
**environmentId** | **string** | The service environment ID this workflow belongs to. | 
**instanceId** | **string** | The resource instance ID. | 

### Other Parameters

Other parameters are passed through a pointer to a apiInventoryApiUpdateResourceInstanceRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------



 **fleetUpdateResourceInstanceRequest2** | [**FleetUpdateResourceInstanceRequest2**](FleetUpdateResourceInstanceRequest2.md) |  | 

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

