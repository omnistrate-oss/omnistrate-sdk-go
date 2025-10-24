# \HostclusterApiAPI

All URIs are relative to *https://api.omnistrate.cloud*

Method | HTTP request | Description
------------- | ------------- | -------------
[**HostclusterApiAdoptHostCluster**](HostclusterApiAPI.md#HostclusterApiAdoptHostCluster) | **Post** /2022-09-01-00/fleet/host-cluster/adopt | AdoptHostCluster hostcluster-api
[**HostclusterApiApplyPendingChangesToHostCluster**](HostclusterApiAPI.md#HostclusterApiApplyPendingChangesToHostCluster) | **Post** /2022-09-01-00/fleet/host-cluster/{id}/apply-pending-changes | ApplyPendingChangesToHostCluster hostcluster-api
[**HostclusterApiDebugHostCluster**](HostclusterApiAPI.md#HostclusterApiDebugHostCluster) | **Get** /2022-09-01-00/fleet/host-cluster/{id}/debug | DebugHostCluster hostcluster-api
[**HostclusterApiDeleteHostCluster**](HostclusterApiAPI.md#HostclusterApiDeleteHostCluster) | **Delete** /2022-09-01-00/fleet/host-cluster/{id} | DeleteHostCluster hostcluster-api
[**HostclusterApiDescribeHostCluster**](HostclusterApiAPI.md#HostclusterApiDescribeHostCluster) | **Get** /2022-09-01-00/fleet/host-cluster/{id} | DescribeHostCluster hostcluster-api
[**HostclusterApiDescribeHostClusterEntity**](HostclusterApiAPI.md#HostclusterApiDescribeHostClusterEntity) | **Get** /2022-09-01-00/fleet/host-cluster/{hostClusterID}/entityType/{entityType}/entityID/{entityID} | DescribeHostClusterEntity hostcluster-api
[**HostclusterApiGenerateTokenForHostClusterDashboard**](HostclusterApiAPI.md#HostclusterApiGenerateTokenForHostClusterDashboard) | **Post** /2022-09-01-00/fleet/host-cluster/{id}/dashboard/token | GenerateTokenForHostClusterDashboard hostcluster-api
[**HostclusterApiKubeConfigHostCluster**](HostclusterApiAPI.md#HostclusterApiKubeConfigHostCluster) | **Get** /2022-09-01-00/fleet/host-cluster/{id}/kubeconfig | KubeConfigHostCluster hostcluster-api
[**HostclusterApiListHostClusterEntities**](HostclusterApiAPI.md#HostclusterApiListHostClusterEntities) | **Get** /2022-09-01-00/fleet/host-cluster/{hostClusterID}/entityType/{entityType} | ListHostClusterEntities hostcluster-api
[**HostclusterApiListHostClusters**](HostclusterApiAPI.md#HostclusterApiListHostClusters) | **Get** /2022-09-01-00/fleet/host-clusters | ListHostClusters hostcluster-api
[**HostclusterApiSetNodePoolProperty**](HostclusterApiAPI.md#HostclusterApiSetNodePoolProperty) | **Patch** /2022-09-01-00/fleet/host-cluster/{hostClusterID}/node-pool/{nodePoolName} | SetNodePoolProperty hostcluster-api
[**HostclusterApiUpdateHostCluster**](HostclusterApiAPI.md#HostclusterApiUpdateHostCluster) | **Patch** /2022-09-01-00/fleet/host-cluster/{id} | UpdateHostCluster hostcluster-api



## HostclusterApiAdoptHostCluster

> AdoptHostClusterResult HostclusterApiAdoptHostCluster(ctx).AdoptHostClusterRequest2(adoptHostClusterRequest2).Execute()

AdoptHostCluster hostcluster-api



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
	adoptHostClusterRequest2 := *openapiclient.NewAdoptHostClusterRequest2("aws|azure|gcp|all", "My Adopted Host Cluster", "Libero molestias beatae cupiditate aliquam itaque.", "us-east-1") // AdoptHostClusterRequest2 | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.HostclusterApiAPI.HostclusterApiAdoptHostCluster(context.Background()).AdoptHostClusterRequest2(adoptHostClusterRequest2).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `HostclusterApiAPI.HostclusterApiAdoptHostCluster``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `HostclusterApiAdoptHostCluster`: AdoptHostClusterResult
	fmt.Fprintf(os.Stdout, "Response from `HostclusterApiAPI.HostclusterApiAdoptHostCluster`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiHostclusterApiAdoptHostClusterRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **adoptHostClusterRequest2** | [**AdoptHostClusterRequest2**](AdoptHostClusterRequest2.md) |  | 

### Return type

[**AdoptHostClusterResult**](AdoptHostClusterResult.md)

### Authorization

[api_key_header_Authorization](../README.md#api_key_header_Authorization)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json, application/vnd.goa.error

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## HostclusterApiApplyPendingChangesToHostCluster

> HostclusterApiApplyPendingChangesToHostCluster(ctx, id).Execute()

ApplyPendingChangesToHostCluster hostcluster-api

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
	id := "hc-12345678" // string | ID of the host cluster to apply pending changes to

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	r, err := apiClient.HostclusterApiAPI.HostclusterApiApplyPendingChangesToHostCluster(context.Background(), id).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `HostclusterApiAPI.HostclusterApiApplyPendingChangesToHostCluster``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **string** | ID of the host cluster to apply pending changes to | 

### Other Parameters

Other parameters are passed through a pointer to a apiHostclusterApiApplyPendingChangesToHostClusterRequest struct via the builder pattern


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


## HostclusterApiDebugHostCluster

> DebugHostClusterResult HostclusterApiDebugHostCluster(ctx, id).IncludeAmenitiesInstallationLogs(includeAmenitiesInstallationLogs).Execute()

DebugHostCluster hostcluster-api



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
	id := "hc-12345678" // string | ID of the host cluster to debug
	includeAmenitiesInstallationLogs := true // bool | Include the installation logs of the amenities in the host cluster (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.HostclusterApiAPI.HostclusterApiDebugHostCluster(context.Background(), id).IncludeAmenitiesInstallationLogs(includeAmenitiesInstallationLogs).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `HostclusterApiAPI.HostclusterApiDebugHostCluster``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `HostclusterApiDebugHostCluster`: DebugHostClusterResult
	fmt.Fprintf(os.Stdout, "Response from `HostclusterApiAPI.HostclusterApiDebugHostCluster`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **string** | ID of the host cluster to debug | 

### Other Parameters

Other parameters are passed through a pointer to a apiHostclusterApiDebugHostClusterRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **includeAmenitiesInstallationLogs** | **bool** | Include the installation logs of the amenities in the host cluster | 

### Return type

[**DebugHostClusterResult**](DebugHostClusterResult.md)

### Authorization

[api_key_header_Authorization](../README.md#api_key_header_Authorization)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json, application/vnd.goa.error

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## HostclusterApiDeleteHostCluster

> HostclusterApiDeleteHostCluster(ctx, id).Execute()

DeleteHostCluster hostcluster-api



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
	id := "hc-12345678" // string | ID of the host cluster to delete

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	r, err := apiClient.HostclusterApiAPI.HostclusterApiDeleteHostCluster(context.Background(), id).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `HostclusterApiAPI.HostclusterApiDeleteHostCluster``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **string** | ID of the host cluster to delete | 

### Other Parameters

Other parameters are passed through a pointer to a apiHostclusterApiDeleteHostClusterRequest struct via the builder pattern


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


## HostclusterApiDescribeHostCluster

> HostCluster HostclusterApiDescribeHostCluster(ctx, id).Execute()

DescribeHostCluster hostcluster-api

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
	id := "hc-12345678" // string | ID of the host cluster to describe

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.HostclusterApiAPI.HostclusterApiDescribeHostCluster(context.Background(), id).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `HostclusterApiAPI.HostclusterApiDescribeHostCluster``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `HostclusterApiDescribeHostCluster`: HostCluster
	fmt.Fprintf(os.Stdout, "Response from `HostclusterApiAPI.HostclusterApiDescribeHostCluster`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **string** | ID of the host cluster to describe | 

### Other Parameters

Other parameters are passed through a pointer to a apiHostclusterApiDescribeHostClusterRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

[**HostCluster**](HostCluster.md)

### Authorization

[api_key_header_Authorization](../README.md#api_key_header_Authorization)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json, application/vnd.goa.error

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## HostclusterApiDescribeHostClusterEntity

> Entity HostclusterApiDescribeHostClusterEntity(ctx, hostClusterID, entityType, entityID).Execute()

DescribeHostClusterEntity hostcluster-api



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
	hostClusterID := "hc-12345abcd" // string | ID of the host cluster to which the entity belongs
	entityType := "NAMESPACE" // string | Type of the entity (e.g., NAMESPACE, SERVICE, POD, etc.)
	entityID := "namespace-12345" // string | Unique identifier of the entity to describe

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.HostclusterApiAPI.HostclusterApiDescribeHostClusterEntity(context.Background(), hostClusterID, entityType, entityID).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `HostclusterApiAPI.HostclusterApiDescribeHostClusterEntity``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `HostclusterApiDescribeHostClusterEntity`: Entity
	fmt.Fprintf(os.Stdout, "Response from `HostclusterApiAPI.HostclusterApiDescribeHostClusterEntity`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**hostClusterID** | **string** | ID of the host cluster to which the entity belongs | 
**entityType** | **string** | Type of the entity (e.g., NAMESPACE, SERVICE, POD, etc.) | 
**entityID** | **string** | Unique identifier of the entity to describe | 

### Other Parameters

Other parameters are passed through a pointer to a apiHostclusterApiDescribeHostClusterEntityRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------




### Return type

[**Entity**](Entity.md)

### Authorization

[api_key_header_Authorization](../README.md#api_key_header_Authorization)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json, application/vnd.goa.error

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## HostclusterApiGenerateTokenForHostClusterDashboard

> GenerateTokenForHostClusterDashboardResult HostclusterApiGenerateTokenForHostClusterDashboard(ctx, id).Execute()

GenerateTokenForHostClusterDashboard hostcluster-api

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
	id := "hc-12345678" // string | ID of the host cluster to delete

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.HostclusterApiAPI.HostclusterApiGenerateTokenForHostClusterDashboard(context.Background(), id).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `HostclusterApiAPI.HostclusterApiGenerateTokenForHostClusterDashboard``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `HostclusterApiGenerateTokenForHostClusterDashboard`: GenerateTokenForHostClusterDashboardResult
	fmt.Fprintf(os.Stdout, "Response from `HostclusterApiAPI.HostclusterApiGenerateTokenForHostClusterDashboard`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **string** | ID of the host cluster to delete | 

### Other Parameters

Other parameters are passed through a pointer to a apiHostclusterApiGenerateTokenForHostClusterDashboardRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

[**GenerateTokenForHostClusterDashboardResult**](GenerateTokenForHostClusterDashboardResult.md)

### Authorization

[api_key_header_Authorization](../README.md#api_key_header_Authorization)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json, application/vnd.goa.error

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## HostclusterApiKubeConfigHostCluster

> KubeConfigHostClusterResult HostclusterApiKubeConfigHostCluster(ctx, id).Role(role).Execute()

KubeConfigHostCluster hostcluster-api



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
	id := "hc-12345678" // string | ID of the host cluster to get the kubeconfig for
	role := "cluster-admin|cluster-reader" // string | The role of the service account to use for the kubeconfig (optional) (default to "cluster-reader")

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.HostclusterApiAPI.HostclusterApiKubeConfigHostCluster(context.Background(), id).Role(role).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `HostclusterApiAPI.HostclusterApiKubeConfigHostCluster``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `HostclusterApiKubeConfigHostCluster`: KubeConfigHostClusterResult
	fmt.Fprintf(os.Stdout, "Response from `HostclusterApiAPI.HostclusterApiKubeConfigHostCluster`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **string** | ID of the host cluster to get the kubeconfig for | 

### Other Parameters

Other parameters are passed through a pointer to a apiHostclusterApiKubeConfigHostClusterRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **role** | **string** | The role of the service account to use for the kubeconfig | [default to &quot;cluster-reader&quot;]

### Return type

[**KubeConfigHostClusterResult**](KubeConfigHostClusterResult.md)

### Authorization

[api_key_header_Authorization](../README.md#api_key_header_Authorization)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json, application/vnd.goa.error

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## HostclusterApiListHostClusterEntities

> ListEntitiesResult HostclusterApiListHostClusterEntities(ctx, hostClusterID, entityType).Execute()

ListHostClusterEntities hostcluster-api



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
	hostClusterID := "hc-12345abcd" // string | ID of the host cluster whose entities are to be listed
	entityType := "NAMESPACE" // string | Type of entities to list (e.g., NAMESPACE, SERVICE, POD, etc.)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.HostclusterApiAPI.HostclusterApiListHostClusterEntities(context.Background(), hostClusterID, entityType).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `HostclusterApiAPI.HostclusterApiListHostClusterEntities``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `HostclusterApiListHostClusterEntities`: ListEntitiesResult
	fmt.Fprintf(os.Stdout, "Response from `HostclusterApiAPI.HostclusterApiListHostClusterEntities`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**hostClusterID** | **string** | ID of the host cluster whose entities are to be listed | 
**entityType** | **string** | Type of entities to list (e.g., NAMESPACE, SERVICE, POD, etc.) | 

### Other Parameters

Other parameters are passed through a pointer to a apiHostclusterApiListHostClusterEntitiesRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------



### Return type

[**ListEntitiesResult**](ListEntitiesResult.md)

### Authorization

[api_key_header_Authorization](../README.md#api_key_header_Authorization)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json, application/vnd.goa.error

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## HostclusterApiListHostClusters

> ListHostClustersResult HostclusterApiListHostClusters(ctx).AccountConfigId(accountConfigId).RegionId(regionId).IncludeProvisionerClusters(includeProvisionerClusters).CustomerEmail(customerEmail).Execute()

ListHostClusters hostcluster-api

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
	accountConfigId := "ac-12345678" // string | The account config ID of the host cluster (optional)
	regionId := "region-12345678" // string | The region ID of the host cluster (optional)
	includeProvisionerClusters := true // bool | Whether to include provisioner clusters in the response (optional)
	customerEmail := "admin@example.com" // string | The email of the customer to filter host clusters by (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.HostclusterApiAPI.HostclusterApiListHostClusters(context.Background()).AccountConfigId(accountConfigId).RegionId(regionId).IncludeProvisionerClusters(includeProvisionerClusters).CustomerEmail(customerEmail).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `HostclusterApiAPI.HostclusterApiListHostClusters``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `HostclusterApiListHostClusters`: ListHostClustersResult
	fmt.Fprintf(os.Stdout, "Response from `HostclusterApiAPI.HostclusterApiListHostClusters`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiHostclusterApiListHostClustersRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **accountConfigId** | **string** | The account config ID of the host cluster | 
 **regionId** | **string** | The region ID of the host cluster | 
 **includeProvisionerClusters** | **bool** | Whether to include provisioner clusters in the response | 
 **customerEmail** | **string** | The email of the customer to filter host clusters by | 

### Return type

[**ListHostClustersResult**](ListHostClustersResult.md)

### Authorization

[api_key_header_Authorization](../README.md#api_key_header_Authorization)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json, application/vnd.goa.error

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## HostclusterApiSetNodePoolProperty

> HostclusterApiSetNodePoolProperty(ctx, hostClusterID, nodePoolName).SetNodePoolPropertyRequest2(setNodePoolPropertyRequest2).Execute()

SetNodePoolProperty hostcluster-api



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
	hostClusterID := "hc-12345abcd" // string | ID of the host cluster to which the entity belongs
	nodePoolName := "nodepool-1" // string | Unique identifier of the node pool to update
	setNodePoolPropertyRequest2 := *openapiclient.NewSetNodePoolPropertyRequest2(int64(10)) // SetNodePoolPropertyRequest2 | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	r, err := apiClient.HostclusterApiAPI.HostclusterApiSetNodePoolProperty(context.Background(), hostClusterID, nodePoolName).SetNodePoolPropertyRequest2(setNodePoolPropertyRequest2).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `HostclusterApiAPI.HostclusterApiSetNodePoolProperty``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**hostClusterID** | **string** | ID of the host cluster to which the entity belongs | 
**nodePoolName** | **string** | Unique identifier of the node pool to update | 

### Other Parameters

Other parameters are passed through a pointer to a apiHostclusterApiSetNodePoolPropertyRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


 **setNodePoolPropertyRequest2** | [**SetNodePoolPropertyRequest2**](SetNodePoolPropertyRequest2.md) |  | 

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


## HostclusterApiUpdateHostCluster

> HostclusterApiUpdateHostCluster(ctx, id).UpdateHostClusterRequest2(updateHostClusterRequest2).Execute()

UpdateHostCluster hostcluster-api

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
	id := "hc-12345678" // string | ID of the host cluster to update
	updateHostClusterRequest2 := *openapiclient.NewUpdateHostClusterRequest2() // UpdateHostClusterRequest2 | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	r, err := apiClient.HostclusterApiAPI.HostclusterApiUpdateHostCluster(context.Background(), id).UpdateHostClusterRequest2(updateHostClusterRequest2).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `HostclusterApiAPI.HostclusterApiUpdateHostCluster``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **string** | ID of the host cluster to update | 

### Other Parameters

Other parameters are passed through a pointer to a apiHostclusterApiUpdateHostClusterRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **updateHostClusterRequest2** | [**UpdateHostClusterRequest2**](UpdateHostClusterRequest2.md) |  | 

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

