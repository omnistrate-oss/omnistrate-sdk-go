# \ResourceApiAPI

All URIs are relative to *https://api.omnistrate.cloud*

Method | HTTP request | Description
------------- | ------------- | -------------
[**ResourceApiAddFile**](ResourceApiAPI.md#ResourceApiAddFile) | **Post** /2022-09-01-00/service/{serviceId}/resource/{id}/file | AddFile resource-api
[**ResourceApiAddResourceDependency**](ResourceApiAPI.md#ResourceApiAddResourceDependency) | **Post** /2022-09-01-00/service/{serviceId}/resource/{id}/resource-dependency/{resourceDependencyId} | AddResourceDependency resource-api
[**ResourceApiCreateResource**](ResourceApiAPI.md#ResourceApiCreateResource) | **Post** /2022-09-01-00/service/{serviceId}/resource | CreateResource resource-api
[**ResourceApiDeleteResource**](ResourceApiAPI.md#ResourceApiDeleteResource) | **Delete** /2022-09-01-00/service/{serviceId}/resource/{id} | DeleteResource resource-api
[**ResourceApiDeprecateResource**](ResourceApiAPI.md#ResourceApiDeprecateResource) | **Patch** /2022-09-01-00/service/{serviceId}/resource/{id}/deprecate | DeprecateResource resource-api
[**ResourceApiDeregisterActionHook**](ResourceApiAPI.md#ResourceApiDeregisterActionHook) | **Delete** /2022-09-01-00/service/{serviceId}/resource/{id}/action-hook | DeregisterActionHook resource-api
[**ResourceApiDeregisterResourceMetricsConfig**](ResourceApiAPI.md#ResourceApiDeregisterResourceMetricsConfig) | **Delete** /2022-09-01-00/service/{serviceId}/resource/{id}/metrics | DeregisterResourceMetricsConfig resource-api
[**ResourceApiDescribeResource**](ResourceApiAPI.md#ResourceApiDescribeResource) | **Get** /2022-09-01-00/service/{serviceId}/resource/{id} | DescribeResource resource-api
[**ResourceApiDescribeResourceMetricsConfig**](ResourceApiAPI.md#ResourceApiDescribeResourceMetricsConfig) | **Get** /2022-09-01-00/service/{serviceId}/resource/{id}/metrics | DescribeResourceMetricsConfig resource-api
[**ResourceApiDisableResourceCapability**](ResourceApiAPI.md#ResourceApiDisableResourceCapability) | **Delete** /2022-09-01-00/service/{serviceId}/resource/{id}/capability | DisableResourceCapability resource-api
[**ResourceApiEnableResourceCapability**](ResourceApiAPI.md#ResourceApiEnableResourceCapability) | **Put** /2022-09-01-00/service/{serviceId}/resource/{id}/capability | EnableResourceCapability resource-api
[**ResourceApiGetFile**](ResourceApiAPI.md#ResourceApiGetFile) | **Get** /2022-09-01-00/service/{serviceId}/resource/{id}/file/{fileId} | GetFile resource-api
[**ResourceApiListActionHooks**](ResourceApiAPI.md#ResourceApiListActionHooks) | **Get** /2022-09-01-00/service/{serviceId}/resource/{id}/action-hook | ListActionHooks resource-api
[**ResourceApiListDependentResource**](ResourceApiAPI.md#ResourceApiListDependentResource) | **Get** /2022-09-01-00/service/{serviceId}/resource/{id}/dependent-resource | ListDependentResource resource-api
[**ResourceApiListFiles**](ResourceApiAPI.md#ResourceApiListFiles) | **Get** /2022-09-01-00/service/{serviceId}/resource/{id}/file | ListFiles resource-api
[**ResourceApiListResourceCapabilities**](ResourceApiAPI.md#ResourceApiListResourceCapabilities) | **Get** /2022-09-01-00/service/{serviceId}/resource/{id}/capability | ListResourceCapabilities resource-api
[**ResourceApiListResources**](ResourceApiAPI.md#ResourceApiListResources) | **Get** /2022-09-01-00/service/{serviceId}/producttier/{productTierId}/resource | ListResources resource-api
[**ResourceApiRegisterActionHook**](ResourceApiAPI.md#ResourceApiRegisterActionHook) | **Post** /2022-09-01-00/service/{serviceId}/resource/{id}/action-hook | RegisterActionHook resource-api
[**ResourceApiRegisterResourceMetricsConfig**](ResourceApiAPI.md#ResourceApiRegisterResourceMetricsConfig) | **Post** /2022-09-01-00/service/{serviceId}/resource/{id}/metrics | RegisterResourceMetricsConfig resource-api
[**ResourceApiRemoveFile**](ResourceApiAPI.md#ResourceApiRemoveFile) | **Delete** /2022-09-01-00/service/{serviceId}/resource/{id}/file/{fileId} | RemoveFile resource-api
[**ResourceApiRemoveResourceDependency**](ResourceApiAPI.md#ResourceApiRemoveResourceDependency) | **Delete** /2022-09-01-00/service/{serviceId}/resource/{id}/resource-dependency/{resourceDependencyId} | RemoveResourceDependency resource-api
[**ResourceApiSetEnvironmentVariables**](ResourceApiAPI.md#ResourceApiSetEnvironmentVariables) | **Patch** /2022-09-01-00/service/{serviceId}/resource/{id}/environment-variables | SetEnvironmentVariables resource-api
[**ResourceApiUndeprecateResource**](ResourceApiAPI.md#ResourceApiUndeprecateResource) | **Patch** /2022-09-01-00/service/{serviceId}/resource/{id}/undeprecate | UndeprecateResource resource-api
[**ResourceApiUnsetEnvironmentVariables**](ResourceApiAPI.md#ResourceApiUnsetEnvironmentVariables) | **Delete** /2022-09-01-00/service/{serviceId}/resource/{id}/environment-variables | UnsetEnvironmentVariables resource-api
[**ResourceApiUpdateFileContent**](ResourceApiAPI.md#ResourceApiUpdateFileContent) | **Put** /2022-09-01-00/service/{serviceId}/resource/{id}/file/{fileId}/content | UpdateFileContent resource-api
[**ResourceApiUpdateFileMetadata**](ResourceApiAPI.md#ResourceApiUpdateFileMetadata) | **Patch** /2022-09-01-00/service/{serviceId}/resource/{id}/file/{fileId}/metadata | UpdateFileMetadata resource-api
[**ResourceApiUpdateResource**](ResourceApiAPI.md#ResourceApiUpdateResource) | **Patch** /2022-09-01-00/service/{serviceId}/resource/{id} | UpdateResource resource-api



## ResourceApiAddFile

> string ResourceApiAddFile(ctx, serviceId, id).Name(name).Description(description).FileType(fileType).MountPath(mountPath).ContentType(contentType).Execute()

AddFile resource-api

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
	serviceId := "s-12345678" // string | The service ID that this API bundle belongs to
	id := "r-12345678" // string | The ID of the resource
	name := "file.txt" // string | The name of the file
	description := "A file to store the configuration for the resource" // string | The description of the file
	fileType := "Config" // string | The type of the file
	mountPath := "/etc/config" // string | The mount path of the file
	contentType := "multipart/form-data; boundary=boundary" // string | Content-Type header, must define value for multipart boundary. (optional) (default to "multipart/form-data; boundary=boundary")

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.ResourceApiAPI.ResourceApiAddFile(context.Background(), serviceId, id).Name(name).Description(description).FileType(fileType).MountPath(mountPath).ContentType(contentType).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `ResourceApiAPI.ResourceApiAddFile``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `ResourceApiAddFile`: string
	fmt.Fprintf(os.Stdout, "Response from `ResourceApiAPI.ResourceApiAddFile`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**serviceId** | **string** | The service ID that this API bundle belongs to | 
**id** | **string** | The ID of the resource | 

### Other Parameters

Other parameters are passed through a pointer to a apiResourceApiAddFileRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


 **name** | **string** | The name of the file | 
 **description** | **string** | The description of the file | 
 **fileType** | **string** | The type of the file | 
 **mountPath** | **string** | The mount path of the file | 
 **contentType** | **string** | Content-Type header, must define value for multipart boundary. | [default to &quot;multipart/form-data; boundary&#x3D;boundary&quot;]

### Return type

**string**

### Authorization

[api_key_header_Authorization](../README.md#api_key_header_Authorization)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json, application/vnd.goa.error

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## ResourceApiAddResourceDependency

> ResourceApiAddResourceDependency(ctx, serviceId, id, resourceDependencyId).AddResourceDependencyRequest2(addResourceDependencyRequest2).Execute()

AddResourceDependency resource-api

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
	serviceId := "s-12345678" // string | The service ID that this API bundle belongs to
	id := "r-12345678" // string | The ID of the resource
	resourceDependencyId := "r-12345678" // string | The ID of the resource to be added as a dependency
	addResourceDependencyRequest2 := *openapiclient.NewAddResourceDependencyRequest2() // AddResourceDependencyRequest2 | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	r, err := apiClient.ResourceApiAPI.ResourceApiAddResourceDependency(context.Background(), serviceId, id, resourceDependencyId).AddResourceDependencyRequest2(addResourceDependencyRequest2).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `ResourceApiAPI.ResourceApiAddResourceDependency``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**serviceId** | **string** | The service ID that this API bundle belongs to | 
**id** | **string** | The ID of the resource | 
**resourceDependencyId** | **string** | The ID of the resource to be added as a dependency | 

### Other Parameters

Other parameters are passed through a pointer to a apiResourceApiAddResourceDependencyRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------



 **addResourceDependencyRequest2** | [**AddResourceDependencyRequest2**](AddResourceDependencyRequest2.md) |  | 

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


## ResourceApiCreateResource

> string ResourceApiCreateResource(ctx, serviceId).CreateResourceRequest2(createResourceRequest2).Execute()

CreateResource resource-api

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
	serviceId := "s-12345678" // string | The service ID that this API bundle belongs to
	createResourceRequest2 := *openapiclient.NewCreateResourceRequest2("A resource to manage a hosted public SaaS offering of a multi-writer MySQL service", "Galera", "Quas doloribus ad est illo.") // CreateResourceRequest2 | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.ResourceApiAPI.ResourceApiCreateResource(context.Background(), serviceId).CreateResourceRequest2(createResourceRequest2).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `ResourceApiAPI.ResourceApiCreateResource``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `ResourceApiCreateResource`: string
	fmt.Fprintf(os.Stdout, "Response from `ResourceApiAPI.ResourceApiCreateResource`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**serviceId** | **string** | The service ID that this API bundle belongs to | 

### Other Parameters

Other parameters are passed through a pointer to a apiResourceApiCreateResourceRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **createResourceRequest2** | [**CreateResourceRequest2**](CreateResourceRequest2.md) |  | 

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


## ResourceApiDeleteResource

> ResourceApiDeleteResource(ctx, serviceId, id).Execute()

DeleteResource resource-api

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
	serviceId := "s-12345678" // string | The service ID that this API bundle belongs to
	id := "r-12345678" // string | The ID of the resource

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	r, err := apiClient.ResourceApiAPI.ResourceApiDeleteResource(context.Background(), serviceId, id).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `ResourceApiAPI.ResourceApiDeleteResource``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**serviceId** | **string** | The service ID that this API bundle belongs to | 
**id** | **string** | The ID of the resource | 

### Other Parameters

Other parameters are passed through a pointer to a apiResourceApiDeleteResourceRequest struct via the builder pattern


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


## ResourceApiDeprecateResource

> ResourceApiDeprecateResource(ctx, serviceId, id).Execute()

DeprecateResource resource-api

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
	serviceId := "s-12345678" // string | The service ID that this API bundle belongs to
	id := "r-12345678" // string | The ID of the resource

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	r, err := apiClient.ResourceApiAPI.ResourceApiDeprecateResource(context.Background(), serviceId, id).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `ResourceApiAPI.ResourceApiDeprecateResource``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**serviceId** | **string** | The service ID that this API bundle belongs to | 
**id** | **string** | The ID of the resource | 

### Other Parameters

Other parameters are passed through a pointer to a apiResourceApiDeprecateResourceRequest struct via the builder pattern


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


## ResourceApiDeregisterActionHook

> ResourceApiDeregisterActionHook(ctx, serviceId, id).DeregisterActionHookRequest2(deregisterActionHookRequest2).Execute()

DeregisterActionHook resource-api

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
	serviceId := "s-12345678" // string | The ID of the service to which the hook belongs
	id := "r-12345678" // string | The ID of the resource to which the hook belongs
	deregisterActionHookRequest2 := *openapiclient.NewDeregisterActionHookRequest2("CLUSTER|NODE", "INIT|ADD|REMOVE|PROMOTE|DEMOTE|HEALTH_CHECK|READINESS_CHECK|STARTUP_CHECK|PRE_START|POST_START|PRE_UPGRADE|POST_UPGRADE|PRE_STOP|POST_STOP") // DeregisterActionHookRequest2 | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	r, err := apiClient.ResourceApiAPI.ResourceApiDeregisterActionHook(context.Background(), serviceId, id).DeregisterActionHookRequest2(deregisterActionHookRequest2).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `ResourceApiAPI.ResourceApiDeregisterActionHook``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**serviceId** | **string** | The ID of the service to which the hook belongs | 
**id** | **string** | The ID of the resource to which the hook belongs | 

### Other Parameters

Other parameters are passed through a pointer to a apiResourceApiDeregisterActionHookRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


 **deregisterActionHookRequest2** | [**DeregisterActionHookRequest2**](DeregisterActionHookRequest2.md) |  | 

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


## ResourceApiDeregisterResourceMetricsConfig

> ResourceApiDeregisterResourceMetricsConfig(ctx, serviceId, id).Execute()

DeregisterResourceMetricsConfig resource-api

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
	serviceId := "s-12345678" // string | The service ID that this API bundle belongs to
	id := "r-12345678" // string | The ID of the resource

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	r, err := apiClient.ResourceApiAPI.ResourceApiDeregisterResourceMetricsConfig(context.Background(), serviceId, id).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `ResourceApiAPI.ResourceApiDeregisterResourceMetricsConfig``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**serviceId** | **string** | The service ID that this API bundle belongs to | 
**id** | **string** | The ID of the resource | 

### Other Parameters

Other parameters are passed through a pointer to a apiResourceApiDeregisterResourceMetricsConfigRequest struct via the builder pattern


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


## ResourceApiDescribeResource

> DescribeResourceResult ResourceApiDescribeResource(ctx, serviceId, id).ProductTierVersion(productTierVersion).ProductTierId(productTierId).Execute()

DescribeResource resource-api

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
	serviceId := "s-12345678" // string | The service ID that this API bundle belongs to
	id := "r-12345678" // string | The ID of the resource
	productTierVersion := "Harum quia quia a." // string | Product tier version of the resource to describe. If not specified, the latest version is described. (optional)
	productTierId := "Beatae beatae." // string | ProductTierId of the resource to describe. Needs to specified in combination with the product tier version (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.ResourceApiAPI.ResourceApiDescribeResource(context.Background(), serviceId, id).ProductTierVersion(productTierVersion).ProductTierId(productTierId).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `ResourceApiAPI.ResourceApiDescribeResource``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `ResourceApiDescribeResource`: DescribeResourceResult
	fmt.Fprintf(os.Stdout, "Response from `ResourceApiAPI.ResourceApiDescribeResource`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**serviceId** | **string** | The service ID that this API bundle belongs to | 
**id** | **string** | The ID of the resource | 

### Other Parameters

Other parameters are passed through a pointer to a apiResourceApiDescribeResourceRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


 **productTierVersion** | **string** | Product tier version of the resource to describe. If not specified, the latest version is described. | 
 **productTierId** | **string** | ProductTierId of the resource to describe. Needs to specified in combination with the product tier version | 

### Return type

[**DescribeResourceResult**](DescribeResourceResult.md)

### Authorization

[api_key_header_Authorization](../README.md#api_key_header_Authorization)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json, application/vnd.goa.error

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## ResourceApiDescribeResourceMetricsConfig

> Describeresourcemetricsconfigresult ResourceApiDescribeResourceMetricsConfig(ctx, serviceId, id).ProductTierVersion(productTierVersion).ProductTierId(productTierId).Execute()

DescribeResourceMetricsConfig resource-api

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
	serviceId := "s-12345678" // string | The service ID that this API bundle belongs to
	id := "r-12345678" // string | The ID of the resource
	productTierVersion := "Pariatur sit ut vitae officiis." // string | Product tier version of the resource to describe. If not specified, the latest version is described. (optional)
	productTierId := "Beatae beatae." // string | ProductTierId of the resource to describe. Needs to specified in combination with the product tier version (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.ResourceApiAPI.ResourceApiDescribeResourceMetricsConfig(context.Background(), serviceId, id).ProductTierVersion(productTierVersion).ProductTierId(productTierId).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `ResourceApiAPI.ResourceApiDescribeResourceMetricsConfig``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `ResourceApiDescribeResourceMetricsConfig`: Describeresourcemetricsconfigresult
	fmt.Fprintf(os.Stdout, "Response from `ResourceApiAPI.ResourceApiDescribeResourceMetricsConfig`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**serviceId** | **string** | The service ID that this API bundle belongs to | 
**id** | **string** | The ID of the resource | 

### Other Parameters

Other parameters are passed through a pointer to a apiResourceApiDescribeResourceMetricsConfigRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


 **productTierVersion** | **string** | Product tier version of the resource to describe. If not specified, the latest version is described. | 
 **productTierId** | **string** | ProductTierId of the resource to describe. Needs to specified in combination with the product tier version | 

### Return type

[**Describeresourcemetricsconfigresult**](Describeresourcemetricsconfigresult.md)

### Authorization

[api_key_header_Authorization](../README.md#api_key_header_Authorization)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json, application/vnd.goa.error

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## ResourceApiDisableResourceCapability

> ResourceApiDisableResourceCapability(ctx, serviceId, id).DisableResourceCapabilityRequest2(disableResourceCapabilityRequest2).Execute()

DisableResourceCapability resource-api

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
	serviceId := "s-12345678" // string | The service ID that this API bundle belongs to
	id := "r-12345678" // string | The ID of the resource
	disableResourceCapabilityRequest2 := *openapiclient.NewDisableResourceCapabilityRequest2("SERVERLESS|SERVICE_ACCOUNT_POLICIES|PROCESS_CORE_DUMP|CUSTOM_DNS|SIDECARS") // DisableResourceCapabilityRequest2 | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	r, err := apiClient.ResourceApiAPI.ResourceApiDisableResourceCapability(context.Background(), serviceId, id).DisableResourceCapabilityRequest2(disableResourceCapabilityRequest2).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `ResourceApiAPI.ResourceApiDisableResourceCapability``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**serviceId** | **string** | The service ID that this API bundle belongs to | 
**id** | **string** | The ID of the resource | 

### Other Parameters

Other parameters are passed through a pointer to a apiResourceApiDisableResourceCapabilityRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


 **disableResourceCapabilityRequest2** | [**DisableResourceCapabilityRequest2**](DisableResourceCapabilityRequest2.md) |  | 

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


## ResourceApiEnableResourceCapability

> ResourceApiEnableResourceCapability(ctx, serviceId, id).EnableResourceCapabilityRequest2(enableResourceCapabilityRequest2).Execute()

EnableResourceCapability resource-api

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
	serviceId := "s-12345678" // string | The service ID that this API bundle belongs to
	id := "r-12345678" // string | The ID of the resource
	enableResourceCapabilityRequest2 := *openapiclient.NewEnableResourceCapabilityRequest2("SERVERLESS|SERVICE_ACCOUNT_POLICIES|PROCESS_CORE_DUMP|CUSTOM_DNS|SIDECARS") // EnableResourceCapabilityRequest2 | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	r, err := apiClient.ResourceApiAPI.ResourceApiEnableResourceCapability(context.Background(), serviceId, id).EnableResourceCapabilityRequest2(enableResourceCapabilityRequest2).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `ResourceApiAPI.ResourceApiEnableResourceCapability``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**serviceId** | **string** | The service ID that this API bundle belongs to | 
**id** | **string** | The ID of the resource | 

### Other Parameters

Other parameters are passed through a pointer to a apiResourceApiEnableResourceCapabilityRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


 **enableResourceCapabilityRequest2** | [**EnableResourceCapabilityRequest2**](EnableResourceCapabilityRequest2.md) |  | 

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


## ResourceApiGetFile

> *os.File ResourceApiGetFile(ctx, serviceId, id, fileId).ProductTierId(productTierId).ProductTierVersion(productTierVersion).Execute()

GetFile resource-api

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
	serviceId := "s-12345678" // string | The service ID that this API bundle belongs to
	id := "r-12345678" // string | The ID of the resource
	fileId := "f-12345678" // string | The ID of the file
	productTierId := "pt-12345678" // string | Product Tier ID of the config files to list (optional)
	productTierVersion := "1.0" // string | Product Tier version of the config files to list. If missing, last version is used (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.ResourceApiAPI.ResourceApiGetFile(context.Background(), serviceId, id, fileId).ProductTierId(productTierId).ProductTierVersion(productTierVersion).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `ResourceApiAPI.ResourceApiGetFile``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `ResourceApiGetFile`: *os.File
	fmt.Fprintf(os.Stdout, "Response from `ResourceApiAPI.ResourceApiGetFile`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**serviceId** | **string** | The service ID that this API bundle belongs to | 
**id** | **string** | The ID of the resource | 
**fileId** | **string** | The ID of the file | 

### Other Parameters

Other parameters are passed through a pointer to a apiResourceApiGetFileRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------



 **productTierId** | **string** | Product Tier ID of the config files to list | 
 **productTierVersion** | **string** | Product Tier version of the config files to list. If missing, last version is used | 

### Return type

[***os.File**](*os.File.md)

### Authorization

[api_key_header_Authorization](../README.md#api_key_header_Authorization)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json, application/vnd.goa.error

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## ResourceApiListActionHooks

> ListActionHooksResult ResourceApiListActionHooks(ctx, serviceId, id).Execute()

ListActionHooks resource-api

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
	serviceId := "s-12345678" // string | The ID of the service to which the hook belongs
	id := "r-12345678" // string | The ID of the resource to which the hook belongs

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.ResourceApiAPI.ResourceApiListActionHooks(context.Background(), serviceId, id).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `ResourceApiAPI.ResourceApiListActionHooks``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `ResourceApiListActionHooks`: ListActionHooksResult
	fmt.Fprintf(os.Stdout, "Response from `ResourceApiAPI.ResourceApiListActionHooks`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**serviceId** | **string** | The ID of the service to which the hook belongs | 
**id** | **string** | The ID of the resource to which the hook belongs | 

### Other Parameters

Other parameters are passed through a pointer to a apiResourceApiListActionHooksRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------



### Return type

[**ListActionHooksResult**](ListActionHooksResult.md)

### Authorization

[api_key_header_Authorization](../README.md#api_key_header_Authorization)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json, application/vnd.goa.error

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## ResourceApiListDependentResource

> ListDependentResourcesResult ResourceApiListDependentResource(ctx, serviceId, id).Execute()

ListDependentResource resource-api

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
	serviceId := "s-12345678" // string | The service ID that this API bundle belongs to
	id := "r-12345678" // string | The ID of the resource

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.ResourceApiAPI.ResourceApiListDependentResource(context.Background(), serviceId, id).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `ResourceApiAPI.ResourceApiListDependentResource``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `ResourceApiListDependentResource`: ListDependentResourcesResult
	fmt.Fprintf(os.Stdout, "Response from `ResourceApiAPI.ResourceApiListDependentResource`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**serviceId** | **string** | The service ID that this API bundle belongs to | 
**id** | **string** | The ID of the resource | 

### Other Parameters

Other parameters are passed through a pointer to a apiResourceApiListDependentResourceRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------



### Return type

[**ListDependentResourcesResult**](ListDependentResourcesResult.md)

### Authorization

[api_key_header_Authorization](../README.md#api_key_header_Authorization)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json, application/vnd.goa.error

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## ResourceApiListFiles

> ListFilesResult ResourceApiListFiles(ctx, serviceId, id).ProductTierId(productTierId).ProductTierVersion(productTierVersion).Execute()

ListFiles resource-api

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
	serviceId := "s-12345678" // string | The service ID that this API bundle belongs to
	id := "r-12345678" // string | The ID of the resource
	productTierId := "pt-12345678" // string | ProductTierId of the config file to describe (optional)
	productTierVersion := "1.0" // string | Product tier version of the config file to describe. If missing, last version is described (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.ResourceApiAPI.ResourceApiListFiles(context.Background(), serviceId, id).ProductTierId(productTierId).ProductTierVersion(productTierVersion).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `ResourceApiAPI.ResourceApiListFiles``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `ResourceApiListFiles`: ListFilesResult
	fmt.Fprintf(os.Stdout, "Response from `ResourceApiAPI.ResourceApiListFiles`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**serviceId** | **string** | The service ID that this API bundle belongs to | 
**id** | **string** | The ID of the resource | 

### Other Parameters

Other parameters are passed through a pointer to a apiResourceApiListFilesRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


 **productTierId** | **string** | ProductTierId of the config file to describe | 
 **productTierVersion** | **string** | Product tier version of the config file to describe. If missing, last version is described | 

### Return type

[**ListFilesResult**](ListFilesResult.md)

### Authorization

[api_key_header_Authorization](../README.md#api_key_header_Authorization)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json, application/vnd.goa.error

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## ResourceApiListResourceCapabilities

> ListResourceCapabilitiesResponse ResourceApiListResourceCapabilities(ctx, serviceId, id).ProductTierVersion(productTierVersion).ProductTierId(productTierId).Execute()

ListResourceCapabilities resource-api

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
	serviceId := "s-12345678" // string | The service ID that this API bundle belongs to
	id := "r-12345678" // string | The ID of the resource
	productTierVersion := "Similique et fuga aut." // string | Product tier version of the instance to describe. If not specified, the latest version is described. (optional)
	productTierId := "Beatae beatae." // string | Product tier id of the instance to describe. Needs to specified in combination with the product tier version (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.ResourceApiAPI.ResourceApiListResourceCapabilities(context.Background(), serviceId, id).ProductTierVersion(productTierVersion).ProductTierId(productTierId).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `ResourceApiAPI.ResourceApiListResourceCapabilities``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `ResourceApiListResourceCapabilities`: ListResourceCapabilitiesResponse
	fmt.Fprintf(os.Stdout, "Response from `ResourceApiAPI.ResourceApiListResourceCapabilities`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**serviceId** | **string** | The service ID that this API bundle belongs to | 
**id** | **string** | The ID of the resource | 

### Other Parameters

Other parameters are passed through a pointer to a apiResourceApiListResourceCapabilitiesRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


 **productTierVersion** | **string** | Product tier version of the instance to describe. If not specified, the latest version is described. | 
 **productTierId** | **string** | Product tier id of the instance to describe. Needs to specified in combination with the product tier version | 

### Return type

[**ListResourceCapabilitiesResponse**](ListResourceCapabilitiesResponse.md)

### Authorization

[api_key_header_Authorization](../README.md#api_key_header_Authorization)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json, application/vnd.goa.error

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## ResourceApiListResources

> ListResourcesResult ResourceApiListResources(ctx, serviceId, productTierId).Managed(managed).ProductTierVersion(productTierVersion).Execute()

ListResources resource-api

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
	serviceId := "si-12345678" // string | The ID of the service to list resources for
	productTierId := "pt-12345678" // string | The product tier ID
	managed := false // bool | Is resource managed by omnistrate (optional)
	productTierVersion := "Consequatur illo ut." // string | Product tier version of the instance to describe. If not specified, the latest version is described. (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.ResourceApiAPI.ResourceApiListResources(context.Background(), serviceId, productTierId).Managed(managed).ProductTierVersion(productTierVersion).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `ResourceApiAPI.ResourceApiListResources``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `ResourceApiListResources`: ListResourcesResult
	fmt.Fprintf(os.Stdout, "Response from `ResourceApiAPI.ResourceApiListResources`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**serviceId** | **string** | The ID of the service to list resources for | 
**productTierId** | **string** | The product tier ID | 

### Other Parameters

Other parameters are passed through a pointer to a apiResourceApiListResourcesRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


 **managed** | **bool** | Is resource managed by omnistrate | 
 **productTierVersion** | **string** | Product tier version of the instance to describe. If not specified, the latest version is described. | 

### Return type

[**ListResourcesResult**](ListResourcesResult.md)

### Authorization

[api_key_header_Authorization](../README.md#api_key_header_Authorization)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json, application/vnd.goa.error

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## ResourceApiRegisterActionHook

> ResourceApiRegisterActionHook(ctx, serviceId, id).RegisterActionHookRequest2(registerActionHookRequest2).Execute()

RegisterActionHook resource-api

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
	serviceId := "s-12345678" // string | The ID of the service to which the hook belongs
	id := "r-12345678" // string | The ID of the resource to which the hook belongs
	registerActionHookRequest2 := *openapiclient.NewRegisterActionHookRequest2("ZWNobyAiaGVsbG8gd29ybGQi", "CLUSTER|NODE", "INIT|ADD|REMOVE|PROMOTE|DEMOTE|HEALTH_CHECK|READINESS_CHECK|STARTUP_CHECK|PRE_START|POST_START|PRE_UPGRADE|POST_UPGRADE|PRE_STOP|POST_STOP") // RegisterActionHookRequest2 | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	r, err := apiClient.ResourceApiAPI.ResourceApiRegisterActionHook(context.Background(), serviceId, id).RegisterActionHookRequest2(registerActionHookRequest2).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `ResourceApiAPI.ResourceApiRegisterActionHook``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**serviceId** | **string** | The ID of the service to which the hook belongs | 
**id** | **string** | The ID of the resource to which the hook belongs | 

### Other Parameters

Other parameters are passed through a pointer to a apiResourceApiRegisterActionHookRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


 **registerActionHookRequest2** | [**RegisterActionHookRequest2**](RegisterActionHookRequest2.md) |  | 

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


## ResourceApiRegisterResourceMetricsConfig

> ResourceApiRegisterResourceMetricsConfig(ctx, serviceId, id).RegisterResourceMetricsConfigRequest2(registerResourceMetricsConfigRequest2).Execute()

RegisterResourceMetricsConfig resource-api

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
	serviceId := "s-12345678" // string | The service ID that this API bundle belongs to
	id := "r-12345678" // string | The ID of the resource
	registerResourceMetricsConfigRequest2 := *openapiclient.NewRegisterResourceMetricsConfigRequest2("http://localhost:9187/metrics") // RegisterResourceMetricsConfigRequest2 | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	r, err := apiClient.ResourceApiAPI.ResourceApiRegisterResourceMetricsConfig(context.Background(), serviceId, id).RegisterResourceMetricsConfigRequest2(registerResourceMetricsConfigRequest2).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `ResourceApiAPI.ResourceApiRegisterResourceMetricsConfig``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**serviceId** | **string** | The service ID that this API bundle belongs to | 
**id** | **string** | The ID of the resource | 

### Other Parameters

Other parameters are passed through a pointer to a apiResourceApiRegisterResourceMetricsConfigRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


 **registerResourceMetricsConfigRequest2** | [**RegisterResourceMetricsConfigRequest2**](RegisterResourceMetricsConfigRequest2.md) |  | 

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


## ResourceApiRemoveFile

> ResourceApiRemoveFile(ctx, serviceId, id, fileId).Execute()

RemoveFile resource-api

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
	serviceId := "s-12345678" // string | The service ID that this API bundle belongs to
	id := "r-12345678" // string | The ID of the resource
	fileId := "file-12345678" // string | The ID of the file

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	r, err := apiClient.ResourceApiAPI.ResourceApiRemoveFile(context.Background(), serviceId, id, fileId).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `ResourceApiAPI.ResourceApiRemoveFile``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**serviceId** | **string** | The service ID that this API bundle belongs to | 
**id** | **string** | The ID of the resource | 
**fileId** | **string** | The ID of the file | 

### Other Parameters

Other parameters are passed through a pointer to a apiResourceApiRemoveFileRequest struct via the builder pattern


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


## ResourceApiRemoveResourceDependency

> ResourceApiRemoveResourceDependency(ctx, serviceId, id, resourceDependencyId).Execute()

RemoveResourceDependency resource-api

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
	serviceId := "s-12345678" // string | The service ID that this API bundle belongs to
	id := "r-12345678" // string | The ID of the resource
	resourceDependencyId := "r-12345678" // string | The ID of the resource dependency to remove

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	r, err := apiClient.ResourceApiAPI.ResourceApiRemoveResourceDependency(context.Background(), serviceId, id, resourceDependencyId).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `ResourceApiAPI.ResourceApiRemoveResourceDependency``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**serviceId** | **string** | The service ID that this API bundle belongs to | 
**id** | **string** | The ID of the resource | 
**resourceDependencyId** | **string** | The ID of the resource dependency to remove | 

### Other Parameters

Other parameters are passed through a pointer to a apiResourceApiRemoveResourceDependencyRequest struct via the builder pattern


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


## ResourceApiSetEnvironmentVariables

> ResourceApiSetEnvironmentVariables(ctx, serviceId, id).SetEnvironmentVariablesRequest2(setEnvironmentVariablesRequest2).Execute()

SetEnvironmentVariables resource-api

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
	serviceId := "s-12345678" // string | The service ID that this API bundle belongs to
	id := "r-12345678" // string | The ID of the resource
	setEnvironmentVariablesRequest2 := *openapiclient.NewSetEnvironmentVariablesRequest2() // SetEnvironmentVariablesRequest2 | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	r, err := apiClient.ResourceApiAPI.ResourceApiSetEnvironmentVariables(context.Background(), serviceId, id).SetEnvironmentVariablesRequest2(setEnvironmentVariablesRequest2).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `ResourceApiAPI.ResourceApiSetEnvironmentVariables``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**serviceId** | **string** | The service ID that this API bundle belongs to | 
**id** | **string** | The ID of the resource | 

### Other Parameters

Other parameters are passed through a pointer to a apiResourceApiSetEnvironmentVariablesRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


 **setEnvironmentVariablesRequest2** | [**SetEnvironmentVariablesRequest2**](SetEnvironmentVariablesRequest2.md) |  | 

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


## ResourceApiUndeprecateResource

> ResourceApiUndeprecateResource(ctx, serviceId, id).Execute()

UndeprecateResource resource-api

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
	serviceId := "s-12345678" // string | The service ID that this API bundle belongs to
	id := "r-12345678" // string | The ID of the resource

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	r, err := apiClient.ResourceApiAPI.ResourceApiUndeprecateResource(context.Background(), serviceId, id).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `ResourceApiAPI.ResourceApiUndeprecateResource``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**serviceId** | **string** | The service ID that this API bundle belongs to | 
**id** | **string** | The ID of the resource | 

### Other Parameters

Other parameters are passed through a pointer to a apiResourceApiUndeprecateResourceRequest struct via the builder pattern


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


## ResourceApiUnsetEnvironmentVariables

> ResourceApiUnsetEnvironmentVariables(ctx, serviceId, id).UnsetEnvironmentVariablesRequest2(unsetEnvironmentVariablesRequest2).Execute()

UnsetEnvironmentVariables resource-api

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
	serviceId := "s-12345678" // string | The service ID that this API bundle belongs to
	id := "r-12345678" // string | The ID of the resource
	unsetEnvironmentVariablesRequest2 := *openapiclient.NewUnsetEnvironmentVariablesRequest2() // UnsetEnvironmentVariablesRequest2 | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	r, err := apiClient.ResourceApiAPI.ResourceApiUnsetEnvironmentVariables(context.Background(), serviceId, id).UnsetEnvironmentVariablesRequest2(unsetEnvironmentVariablesRequest2).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `ResourceApiAPI.ResourceApiUnsetEnvironmentVariables``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**serviceId** | **string** | The service ID that this API bundle belongs to | 
**id** | **string** | The ID of the resource | 

### Other Parameters

Other parameters are passed through a pointer to a apiResourceApiUnsetEnvironmentVariablesRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


 **unsetEnvironmentVariablesRequest2** | [**UnsetEnvironmentVariablesRequest2**](UnsetEnvironmentVariablesRequest2.md) |  | 

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


## ResourceApiUpdateFileContent

> ResourceApiUpdateFileContent(ctx, serviceId, id, fileId).ContentType(contentType).Execute()

UpdateFileContent resource-api

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
	serviceId := "s-12345678" // string | The service ID that this API bundle belongs to
	id := "r-12345678" // string | The ID of the resource
	fileId := "f-12345678" // string | The ID of the file
	contentType := "multipart/form-data; boundary=boundary" // string | Content-Type header, must define value for multipart boundary. (optional) (default to "multipart/form-data; boundary=boundary")

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	r, err := apiClient.ResourceApiAPI.ResourceApiUpdateFileContent(context.Background(), serviceId, id, fileId).ContentType(contentType).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `ResourceApiAPI.ResourceApiUpdateFileContent``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**serviceId** | **string** | The service ID that this API bundle belongs to | 
**id** | **string** | The ID of the resource | 
**fileId** | **string** | The ID of the file | 

### Other Parameters

Other parameters are passed through a pointer to a apiResourceApiUpdateFileContentRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------



 **contentType** | **string** | Content-Type header, must define value for multipart boundary. | [default to &quot;multipart/form-data; boundary&#x3D;boundary&quot;]

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


## ResourceApiUpdateFileMetadata

> ResourceApiUpdateFileMetadata(ctx, serviceId, id, fileId).UpdateFileMetadataRequest2(updateFileMetadataRequest2).Execute()

UpdateFileMetadata resource-api

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
	serviceId := "s-12345678" // string | The service ID that this API bundle belongs to
	id := "r-12345678" // string | The ID of the resource
	fileId := "f-12345678" // string | The ID of the file
	updateFileMetadataRequest2 := *openapiclient.NewUpdateFileMetadataRequest2() // UpdateFileMetadataRequest2 | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	r, err := apiClient.ResourceApiAPI.ResourceApiUpdateFileMetadata(context.Background(), serviceId, id, fileId).UpdateFileMetadataRequest2(updateFileMetadataRequest2).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `ResourceApiAPI.ResourceApiUpdateFileMetadata``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**serviceId** | **string** | The service ID that this API bundle belongs to | 
**id** | **string** | The ID of the resource | 
**fileId** | **string** | The ID of the file | 

### Other Parameters

Other parameters are passed through a pointer to a apiResourceApiUpdateFileMetadataRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------



 **updateFileMetadataRequest2** | [**UpdateFileMetadataRequest2**](UpdateFileMetadataRequest2.md) |  | 

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


## ResourceApiUpdateResource

> ResourceApiUpdateResource(ctx, serviceId, id).UpdateResourceRequest2(updateResourceRequest2).Execute()

UpdateResource resource-api

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
	serviceId := "s-12345678" // string | The service ID that this API bundle belongs to
	id := "r-12345678" // string | The ID of the resource
	updateResourceRequest2 := *openapiclient.NewUpdateResourceRequest2() // UpdateResourceRequest2 | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	r, err := apiClient.ResourceApiAPI.ResourceApiUpdateResource(context.Background(), serviceId, id).UpdateResourceRequest2(updateResourceRequest2).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `ResourceApiAPI.ResourceApiUpdateResource``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**serviceId** | **string** | The service ID that this API bundle belongs to | 
**id** | **string** | The ID of the resource | 

### Other Parameters

Other parameters are passed through a pointer to a apiResourceApiUpdateResourceRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


 **updateResourceRequest2** | [**UpdateResourceRequest2**](UpdateResourceRequest2.md) |  | 

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

