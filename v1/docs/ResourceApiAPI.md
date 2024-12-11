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

> ResourceApiAddResourceDependency(ctx, serviceId, id, resourceDependencyId).AddResourceDependencyRequestBody(addResourceDependencyRequestBody).Execute()

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
	addResourceDependencyRequestBody := *openapiclient.NewAddResourceDependencyRequestBody() // AddResourceDependencyRequestBody | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	r, err := apiClient.ResourceApiAPI.ResourceApiAddResourceDependency(context.Background(), serviceId, id, resourceDependencyId).AddResourceDependencyRequestBody(addResourceDependencyRequestBody).Execute()
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



 **addResourceDependencyRequestBody** | [**AddResourceDependencyRequestBody**](AddResourceDependencyRequestBody.md) |  | 

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

> string ResourceApiCreateResource(ctx, serviceId).CreateResourceRequestBody(createResourceRequestBody).Execute()

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
	createResourceRequestBody := *openapiclient.NewCreateResourceRequestBody("A resource to manage a hosted public SaaS offering of a multi-writer MySQL service", "Galera", "pt-12345678") // CreateResourceRequestBody | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.ResourceApiAPI.ResourceApiCreateResource(context.Background(), serviceId).CreateResourceRequestBody(createResourceRequestBody).Execute()
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

 **createResourceRequestBody** | [**CreateResourceRequestBody**](CreateResourceRequestBody.md) |  | 

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

> ResourceApiDeregisterActionHook(ctx, serviceId, id).DeregisterActionHookRequestBody(deregisterActionHookRequestBody).Execute()

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
	deregisterActionHookRequestBody := *openapiclient.NewDeregisterActionHookRequestBody("CLUSTER", "ADD") // DeregisterActionHookRequestBody | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	r, err := apiClient.ResourceApiAPI.ResourceApiDeregisterActionHook(context.Background(), serviceId, id).DeregisterActionHookRequestBody(deregisterActionHookRequestBody).Execute()
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


 **deregisterActionHookRequestBody** | [**DeregisterActionHookRequestBody**](DeregisterActionHookRequestBody.md) |  | 

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
	productTierVersion := "Quaerat harum totam omnis." // string | Product tier version of the resource to describe. If not specified, the latest version is described. (optional)
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
	productTierVersion := "Et quo enim." // string | Product tier version of the resource to describe. If not specified, the latest version is described. (optional)
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

> ResourceApiDisableResourceCapability(ctx, serviceId, id).DisableResourceCapabilityRequestBody(disableResourceCapabilityRequestBody).Execute()

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
	disableResourceCapabilityRequestBody := *openapiclient.NewDisableResourceCapabilityRequestBody("SERVERLESS") // DisableResourceCapabilityRequestBody | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	r, err := apiClient.ResourceApiAPI.ResourceApiDisableResourceCapability(context.Background(), serviceId, id).DisableResourceCapabilityRequestBody(disableResourceCapabilityRequestBody).Execute()
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


 **disableResourceCapabilityRequestBody** | [**DisableResourceCapabilityRequestBody**](DisableResourceCapabilityRequestBody.md) |  | 

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

> ResourceApiEnableResourceCapability(ctx, serviceId, id).EnableResourceCapabilityRequestBody(enableResourceCapabilityRequestBody).Execute()

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
	enableResourceCapabilityRequestBody := *openapiclient.NewEnableResourceCapabilityRequestBody("SERVERLESS") // EnableResourceCapabilityRequestBody | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	r, err := apiClient.ResourceApiAPI.ResourceApiEnableResourceCapability(context.Background(), serviceId, id).EnableResourceCapabilityRequestBody(enableResourceCapabilityRequestBody).Execute()
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


 **enableResourceCapabilityRequestBody** | [**EnableResourceCapabilityRequestBody**](EnableResourceCapabilityRequestBody.md) |  | 

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
	productTierVersion := "Deserunt ut et ut illum." // string | Product tier version of the instance to describe. If not specified, the latest version is described. (optional)
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
	productTierVersion := "Consectetur error at quo." // string | Product tier version of the instance to describe. If not specified, the latest version is described. (optional)

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

> ResourceApiRegisterActionHook(ctx, serviceId, id).RegisterActionHookRequestBody(registerActionHookRequestBody).Execute()

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
	registerActionHookRequestBody := *openapiclient.NewRegisterActionHookRequestBody("ZWNobyAiaGVsbG8gd29ybGQi", "CLUSTER", "ADD") // RegisterActionHookRequestBody | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	r, err := apiClient.ResourceApiAPI.ResourceApiRegisterActionHook(context.Background(), serviceId, id).RegisterActionHookRequestBody(registerActionHookRequestBody).Execute()
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


 **registerActionHookRequestBody** | [**RegisterActionHookRequestBody**](RegisterActionHookRequestBody.md) |  | 

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

> ResourceApiRegisterResourceMetricsConfig(ctx, serviceId, id).RegisterResourceMetricsConfigRequestBody(registerResourceMetricsConfigRequestBody).Execute()

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
	registerResourceMetricsConfigRequestBody := *openapiclient.NewRegisterResourceMetricsConfigRequestBody("http://localhost:9187/metrics") // RegisterResourceMetricsConfigRequestBody | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	r, err := apiClient.ResourceApiAPI.ResourceApiRegisterResourceMetricsConfig(context.Background(), serviceId, id).RegisterResourceMetricsConfigRequestBody(registerResourceMetricsConfigRequestBody).Execute()
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


 **registerResourceMetricsConfigRequestBody** | [**RegisterResourceMetricsConfigRequestBody**](RegisterResourceMetricsConfigRequestBody.md) |  | 

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

> ResourceApiSetEnvironmentVariables(ctx, serviceId, id).SetEnvironmentVariablesRequestBody(setEnvironmentVariablesRequestBody).Execute()

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
	setEnvironmentVariablesRequestBody := *openapiclient.NewSetEnvironmentVariablesRequestBody() // SetEnvironmentVariablesRequestBody | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	r, err := apiClient.ResourceApiAPI.ResourceApiSetEnvironmentVariables(context.Background(), serviceId, id).SetEnvironmentVariablesRequestBody(setEnvironmentVariablesRequestBody).Execute()
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


 **setEnvironmentVariablesRequestBody** | [**SetEnvironmentVariablesRequestBody**](SetEnvironmentVariablesRequestBody.md) |  | 

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


## ResourceApiUnsetEnvironmentVariables

> ResourceApiUnsetEnvironmentVariables(ctx, serviceId, id).SetEnvironmentVariablesRequestBody(setEnvironmentVariablesRequestBody).Execute()

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
	setEnvironmentVariablesRequestBody := *openapiclient.NewSetEnvironmentVariablesRequestBody() // SetEnvironmentVariablesRequestBody | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	r, err := apiClient.ResourceApiAPI.ResourceApiUnsetEnvironmentVariables(context.Background(), serviceId, id).SetEnvironmentVariablesRequestBody(setEnvironmentVariablesRequestBody).Execute()
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


 **setEnvironmentVariablesRequestBody** | [**SetEnvironmentVariablesRequestBody**](SetEnvironmentVariablesRequestBody.md) |  | 

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

> ResourceApiUpdateFileMetadata(ctx, serviceId, id, fileId).UpdateFileMetadataRequestBody(updateFileMetadataRequestBody).Execute()

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
	updateFileMetadataRequestBody := *openapiclient.NewUpdateFileMetadataRequestBody() // UpdateFileMetadataRequestBody | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	r, err := apiClient.ResourceApiAPI.ResourceApiUpdateFileMetadata(context.Background(), serviceId, id, fileId).UpdateFileMetadataRequestBody(updateFileMetadataRequestBody).Execute()
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



 **updateFileMetadataRequestBody** | [**UpdateFileMetadataRequestBody**](UpdateFileMetadataRequestBody.md) |  | 

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

> ResourceApiUpdateResource(ctx, serviceId, id).UpdateResourceRequestBody(updateResourceRequestBody).Execute()

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
	updateResourceRequestBody := *openapiclient.NewUpdateResourceRequestBody() // UpdateResourceRequestBody | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	r, err := apiClient.ResourceApiAPI.ResourceApiUpdateResource(context.Background(), serviceId, id).UpdateResourceRequestBody(updateResourceRequestBody).Execute()
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


 **updateResourceRequestBody** | [**UpdateResourceRequestBody**](UpdateResourceRequestBody.md) |  | 

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

