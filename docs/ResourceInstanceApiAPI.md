# \ResourceInstanceApiAPI

All URIs are relative to *https://api.omnistrate.cloud*

Method | HTTP request | Description
------------- | ------------- | -------------
[**ResourceInstanceApiAddCapacityToResourceInstance**](ResourceInstanceApiAPI.md#ResourceInstanceApiAddCapacityToResourceInstance) | **Post** /2022-09-01-00/resource-instance/{serviceProviderId}/{serviceKey}/{serviceAPIVersion}/{serviceEnvironmentKey}/{serviceModelKey}/{productTierKey}/{resourceKey}/{id}/add-capacity | AddCapacityToResourceInstance resource-instance-api
[**ResourceInstanceApiAddCustomDNSToResourceInstance**](ResourceInstanceApiAPI.md#ResourceInstanceApiAddCustomDNSToResourceInstance) | **Post** /2022-09-01-00/resource-instance/{serviceProviderId}/{serviceKey}/{serviceAPIVersion}/{serviceEnvironmentKey}/{serviceModelKey}/{productTierKey}/{resourceKey}/{id}/custom-dns | AddCustomDNSToResourceInstance resource-instance-api
[**ResourceInstanceApiCreateResourceInstance**](ResourceInstanceApiAPI.md#ResourceInstanceApiCreateResourceInstance) | **Post** /2022-09-01-00/resource-instance/{serviceProviderId}/{serviceKey}/{serviceAPIVersion}/{serviceEnvironmentKey}/{serviceModelKey}/{productTierKey}/{resourceKey} | CreateResourceInstance resource-instance-api
[**ResourceInstanceApiDeleteResourceInstance**](ResourceInstanceApiAPI.md#ResourceInstanceApiDeleteResourceInstance) | **Delete** /2022-09-01-00/resource-instance/{serviceProviderId}/{serviceKey}/{serviceAPIVersion}/{serviceEnvironmentKey}/{serviceModelKey}/{productTierKey}/{resourceKey}/{id} | DeleteResourceInstance resource-instance-api
[**ResourceInstanceApiDescribeResourceInstance**](ResourceInstanceApiAPI.md#ResourceInstanceApiDescribeResourceInstance) | **Get** /2022-09-01-00/resource-instance/{serviceProviderId}/{serviceKey}/{serviceAPIVersion}/{serviceEnvironmentKey}/{serviceModelKey}/{productTierKey}/{resourceKey}/{id} | DescribeResourceInstance resource-instance-api
[**ResourceInstanceApiFailoverResourceInstance**](ResourceInstanceApiAPI.md#ResourceInstanceApiFailoverResourceInstance) | **Post** /2022-09-01-00/resource-instance/{serviceProviderId}/{serviceKey}/{serviceAPIVersion}/{serviceEnvironmentKey}/{serviceModelKey}/{productTierKey}/{resourceKey}/{id}/failover | FailoverResourceInstance resource-instance-api
[**ResourceInstanceApiListResourceInstances**](ResourceInstanceApiAPI.md#ResourceInstanceApiListResourceInstances) | **Get** /2022-09-01-00/resource-instance/{serviceProviderId}/{serviceKey}/{serviceAPIVersion}/{serviceEnvironmentKey}/{serviceModelKey}/{productTierKey}/{resourceKey} | ListResourceInstances resource-instance-api
[**ResourceInstanceApiRemoveCapacityFromResourceInstance**](ResourceInstanceApiAPI.md#ResourceInstanceApiRemoveCapacityFromResourceInstance) | **Post** /2022-09-01-00/resource-instance/{serviceProviderId}/{serviceKey}/{serviceAPIVersion}/{serviceEnvironmentKey}/{serviceModelKey}/{productTierKey}/{resourceKey}/{id}/remove-capacity | RemoveCapacityFromResourceInstance resource-instance-api
[**ResourceInstanceApiRemoveCustomDNSFromResourceInstance**](ResourceInstanceApiAPI.md#ResourceInstanceApiRemoveCustomDNSFromResourceInstance) | **Delete** /2022-09-01-00/resource-instance/{serviceProviderId}/{serviceKey}/{serviceAPIVersion}/{serviceEnvironmentKey}/{serviceModelKey}/{productTierKey}/{resourceKey}/{id}/custom-dns | RemoveCustomDNSFromResourceInstance resource-instance-api
[**ResourceInstanceApiResourceInstanceProvisionerSetupKit**](ResourceInstanceApiAPI.md#ResourceInstanceApiResourceInstanceProvisionerSetupKit) | **Get** /2022-09-01-00/resource-instance/{serviceProviderId}/{serviceKey}/{serviceAPIVersion}/{serviceEnvironmentKey}/{serviceModelKey}/setup-kit | ResourceInstanceProvisionerSetupKit resource-instance-api
[**ResourceInstanceApiRestartResourceInstance**](ResourceInstanceApiAPI.md#ResourceInstanceApiRestartResourceInstance) | **Post** /2022-09-01-00/resource-instance/{serviceProviderId}/{serviceKey}/{serviceAPIVersion}/{serviceEnvironmentKey}/{serviceModelKey}/{productTierKey}/{resourceKey}/{id}/restart | RestartResourceInstance resource-instance-api
[**ResourceInstanceApiRestoreResourceInstance**](ResourceInstanceApiAPI.md#ResourceInstanceApiRestoreResourceInstance) | **Post** /2022-09-01-00/resource-instance/{serviceProviderId}/{serviceKey}/{serviceAPIVersion}/{serviceEnvironmentKey}/{serviceModelKey}/{productTierKey}/{resourceKey}/{id}/restore | RestoreResourceInstance resource-instance-api
[**ResourceInstanceApiStartResourceInstance**](ResourceInstanceApiAPI.md#ResourceInstanceApiStartResourceInstance) | **Post** /2022-09-01-00/resource-instance/{serviceProviderId}/{serviceKey}/{serviceAPIVersion}/{serviceEnvironmentKey}/{serviceModelKey}/{productTierKey}/{resourceKey}/{id}/start | StartResourceInstance resource-instance-api
[**ResourceInstanceApiStopResourceInstance**](ResourceInstanceApiAPI.md#ResourceInstanceApiStopResourceInstance) | **Post** /2022-09-01-00/resource-instance/{serviceProviderId}/{serviceKey}/{serviceAPIVersion}/{serviceEnvironmentKey}/{serviceModelKey}/{productTierKey}/{resourceKey}/{id}/stop | StopResourceInstance resource-instance-api
[**ResourceInstanceApiUpdateResourceInstance**](ResourceInstanceApiAPI.md#ResourceInstanceApiUpdateResourceInstance) | **Patch** /2022-09-01-00/resource-instance/{serviceProviderId}/{serviceKey}/{serviceAPIVersion}/{serviceEnvironmentKey}/{serviceModelKey}/{productTierKey}/{resourceKey}/{id} | UpdateResourceInstance resource-instance-api



## ResourceInstanceApiAddCapacityToResourceInstance

> ResourceInstanceApiAddCapacityToResourceInstance(ctx, serviceProviderId, serviceKey, serviceAPIVersion, serviceEnvironmentKey, serviceModelKey, productTierKey, resourceKey, id).AddCapacityToResourceInstanceRequestBody(addCapacityToResourceInstanceRequestBody).SubscriptionId(subscriptionId).Execute()

AddCapacityToResourceInstance resource-instance-api

### Example

```go
package main

import (
	"context"
	"fmt"
	"os"
	openapiclient "github.com/omnistrate/omnistrate-sdk-go"
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
	addCapacityToResourceInstanceRequestBody := *openapiclient.NewAddCapacityToResourceInstanceRequestBody(int64(3)) // AddCapacityToResourceInstanceRequestBody | 
	subscriptionId := "sub-abcd1234" // string | Subscription Id (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	r, err := apiClient.ResourceInstanceApiAPI.ResourceInstanceApiAddCapacityToResourceInstance(context.Background(), serviceProviderId, serviceKey, serviceAPIVersion, serviceEnvironmentKey, serviceModelKey, productTierKey, resourceKey, id).AddCapacityToResourceInstanceRequestBody(addCapacityToResourceInstanceRequestBody).SubscriptionId(subscriptionId).Execute()
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








 **addCapacityToResourceInstanceRequestBody** | [**AddCapacityToResourceInstanceRequestBody**](AddCapacityToResourceInstanceRequestBody.md) |  | 
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

> ResourceInstanceApiAddCustomDNSToResourceInstance(ctx, serviceProviderId, serviceKey, serviceAPIVersion, serviceEnvironmentKey, serviceModelKey, productTierKey, resourceKey, id).AddCustomDNSToResourceInstanceRequestBody(addCustomDNSToResourceInstanceRequestBody).SubscriptionId(subscriptionId).Execute()

AddCustomDNSToResourceInstance resource-instance-api

### Example

```go
package main

import (
	"context"
	"fmt"
	"os"
	openapiclient "github.com/omnistrate/omnistrate-sdk-go"
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
	addCustomDNSToResourceInstanceRequestBody := *openapiclient.NewAddCustomDNSToResourceInstanceRequestBody("my-custom-dns.com") // AddCustomDNSToResourceInstanceRequestBody | 
	subscriptionId := "sub-abcd1234" // string | Subscription Id (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	r, err := apiClient.ResourceInstanceApiAPI.ResourceInstanceApiAddCustomDNSToResourceInstance(context.Background(), serviceProviderId, serviceKey, serviceAPIVersion, serviceEnvironmentKey, serviceModelKey, productTierKey, resourceKey, id).AddCustomDNSToResourceInstanceRequestBody(addCustomDNSToResourceInstanceRequestBody).SubscriptionId(subscriptionId).Execute()
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








 **addCustomDNSToResourceInstanceRequestBody** | [**AddCustomDNSToResourceInstanceRequestBody**](AddCustomDNSToResourceInstanceRequestBody.md) |  | 
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


## ResourceInstanceApiCreateResourceInstance

> CreateResourceInstanceResponseBody ResourceInstanceApiCreateResourceInstance(ctx, serviceProviderId, serviceKey, serviceAPIVersion, serviceEnvironmentKey, serviceModelKey, productTierKey, resourceKey).CreateResourceInstanceRequestBody(createResourceInstanceRequestBody).SubscriptionId(subscriptionId).Execute()

CreateResourceInstance resource-instance-api

### Example

```go
package main

import (
	"context"
	"fmt"
	"os"
	openapiclient "github.com/omnistrate/omnistrate-sdk-go"
)

func main() {
	serviceProviderId := "omnistrate" // string | The service provider ID
	serviceKey := "service-orchestration" // string | The service name
	serviceAPIVersion := "v1" // string | The service API version
	serviceEnvironmentKey := "dev" // string | The service environment name
	serviceModelKey := "hosted" // string | The service model name
	productTierKey := "premium" // string | The product tier name
	resourceKey := "mysql" // string | The resource key
	createResourceInstanceRequestBody := *openapiclient.NewCreateResourceInstanceRequestBody() // CreateResourceInstanceRequestBody | 
	subscriptionId := "sub-abcd1234" // string | Subscription Id (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.ResourceInstanceApiAPI.ResourceInstanceApiCreateResourceInstance(context.Background(), serviceProviderId, serviceKey, serviceAPIVersion, serviceEnvironmentKey, serviceModelKey, productTierKey, resourceKey).CreateResourceInstanceRequestBody(createResourceInstanceRequestBody).SubscriptionId(subscriptionId).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `ResourceInstanceApiAPI.ResourceInstanceApiCreateResourceInstance``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `ResourceInstanceApiCreateResourceInstance`: CreateResourceInstanceResponseBody
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







 **createResourceInstanceRequestBody** | [**CreateResourceInstanceRequestBody**](CreateResourceInstanceRequestBody.md) |  | 
 **subscriptionId** | **string** | Subscription Id | 

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
	openapiclient "github.com/omnistrate/omnistrate-sdk-go"
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
	openapiclient "github.com/omnistrate/omnistrate-sdk-go"
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

> ResourceInstanceApiFailoverResourceInstance(ctx, serviceProviderId, serviceKey, serviceAPIVersion, serviceEnvironmentKey, serviceModelKey, productTierKey, resourceKey, id).FailoverResourceInstanceRequestBody(failoverResourceInstanceRequestBody).SubscriptionId(subscriptionId).Execute()

FailoverResourceInstance resource-instance-api

### Example

```go
package main

import (
	"context"
	"fmt"
	"os"
	openapiclient "github.com/omnistrate/omnistrate-sdk-go"
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
	failoverResourceInstanceRequestBody := *openapiclient.NewFailoverResourceInstanceRequestBody("db-0") // FailoverResourceInstanceRequestBody | 
	subscriptionId := "sub-abcd1234" // string | Subscription Id (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	r, err := apiClient.ResourceInstanceApiAPI.ResourceInstanceApiFailoverResourceInstance(context.Background(), serviceProviderId, serviceKey, serviceAPIVersion, serviceEnvironmentKey, serviceModelKey, productTierKey, resourceKey, id).FailoverResourceInstanceRequestBody(failoverResourceInstanceRequestBody).SubscriptionId(subscriptionId).Execute()
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








 **failoverResourceInstanceRequestBody** | [**FailoverResourceInstanceRequestBody**](FailoverResourceInstanceRequestBody.md) |  | 
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


## ResourceInstanceApiListResourceInstances

> ListServiceEnvironmentsResult ResourceInstanceApiListResourceInstances(ctx, serviceProviderId, serviceKey, serviceAPIVersion, serviceEnvironmentKey, serviceModelKey, productTierKey, resourceKey).SubscriptionId(subscriptionId).Execute()

ListResourceInstances resource-instance-api

### Example

```go
package main

import (
	"context"
	"fmt"
	"os"
	openapiclient "github.com/omnistrate/omnistrate-sdk-go"
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
	// response from `ResourceInstanceApiListResourceInstances`: ListServiceEnvironmentsResult
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

[**ListServiceEnvironmentsResult**](ListServiceEnvironmentsResult.md)

### Authorization

[api_key_header_Authorization](../README.md#api_key_header_Authorization)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json, application/vnd.goa.error

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## ResourceInstanceApiRemoveCapacityFromResourceInstance

> ResourceInstanceApiRemoveCapacityFromResourceInstance(ctx, serviceProviderId, serviceKey, serviceAPIVersion, serviceEnvironmentKey, serviceModelKey, productTierKey, resourceKey, id).RemoveCapacityFromResourceInstanceRequestBody(removeCapacityFromResourceInstanceRequestBody).SubscriptionId(subscriptionId).Execute()

RemoveCapacityFromResourceInstance resource-instance-api

### Example

```go
package main

import (
	"context"
	"fmt"
	"os"
	openapiclient "github.com/omnistrate/omnistrate-sdk-go"
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
	removeCapacityFromResourceInstanceRequestBody := *openapiclient.NewRemoveCapacityFromResourceInstanceRequestBody(int64(3)) // RemoveCapacityFromResourceInstanceRequestBody | 
	subscriptionId := "sub-abcd1234" // string | Subscription Id (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	r, err := apiClient.ResourceInstanceApiAPI.ResourceInstanceApiRemoveCapacityFromResourceInstance(context.Background(), serviceProviderId, serviceKey, serviceAPIVersion, serviceEnvironmentKey, serviceModelKey, productTierKey, resourceKey, id).RemoveCapacityFromResourceInstanceRequestBody(removeCapacityFromResourceInstanceRequestBody).SubscriptionId(subscriptionId).Execute()
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








 **removeCapacityFromResourceInstanceRequestBody** | [**RemoveCapacityFromResourceInstanceRequestBody**](RemoveCapacityFromResourceInstanceRequestBody.md) |  | 
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
	openapiclient "github.com/omnistrate/omnistrate-sdk-go"
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
	openapiclient "github.com/omnistrate/omnistrate-sdk-go"
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
	openapiclient "github.com/omnistrate/omnistrate-sdk-go"
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

> CreateResourceInstanceResponseBody ResourceInstanceApiRestoreResourceInstance(ctx, serviceProviderId, serviceKey, serviceAPIVersion, serviceEnvironmentKey, serviceModelKey, productTierKey, resourceKey, id).RestoreResourceInstanceRequestBody(restoreResourceInstanceRequestBody).SubscriptionId(subscriptionId).Execute()

RestoreResourceInstance resource-instance-api

### Example

```go
package main

import (
	"context"
	"fmt"
	"os"
	openapiclient "github.com/omnistrate/omnistrate-sdk-go"
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
	restoreResourceInstanceRequestBody := *openapiclient.NewRestoreResourceInstanceRequestBody("2021-09-01T00:00:00Z") // RestoreResourceInstanceRequestBody | 
	subscriptionId := "sub-abcd1234" // string | Subscription Id (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.ResourceInstanceApiAPI.ResourceInstanceApiRestoreResourceInstance(context.Background(), serviceProviderId, serviceKey, serviceAPIVersion, serviceEnvironmentKey, serviceModelKey, productTierKey, resourceKey, id).RestoreResourceInstanceRequestBody(restoreResourceInstanceRequestBody).SubscriptionId(subscriptionId).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `ResourceInstanceApiAPI.ResourceInstanceApiRestoreResourceInstance``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `ResourceInstanceApiRestoreResourceInstance`: CreateResourceInstanceResponseBody
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








 **restoreResourceInstanceRequestBody** | [**RestoreResourceInstanceRequestBody**](RestoreResourceInstanceRequestBody.md) |  | 
 **subscriptionId** | **string** | Subscription Id | 

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
	openapiclient "github.com/omnistrate/omnistrate-sdk-go"
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
	openapiclient "github.com/omnistrate/omnistrate-sdk-go"
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


## ResourceInstanceApiUpdateResourceInstance

> ResourceInstanceApiUpdateResourceInstance(ctx, serviceProviderId, serviceKey, serviceAPIVersion, serviceEnvironmentKey, serviceModelKey, productTierKey, resourceKey, id).UpdateResourceInstanceRequestBody(updateResourceInstanceRequestBody).SubscriptionId(subscriptionId).Execute()

UpdateResourceInstance resource-instance-api

### Example

```go
package main

import (
	"context"
	"fmt"
	"os"
	openapiclient "github.com/omnistrate/omnistrate-sdk-go"
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
	updateResourceInstanceRequestBody := *openapiclient.NewUpdateResourceInstanceRequestBody() // UpdateResourceInstanceRequestBody | 
	subscriptionId := "sub-abcd1234" // string | Subscription Id (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	r, err := apiClient.ResourceInstanceApiAPI.ResourceInstanceApiUpdateResourceInstance(context.Background(), serviceProviderId, serviceKey, serviceAPIVersion, serviceEnvironmentKey, serviceModelKey, productTierKey, resourceKey, id).UpdateResourceInstanceRequestBody(updateResourceInstanceRequestBody).SubscriptionId(subscriptionId).Execute()
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








 **updateResourceInstanceRequestBody** | [**UpdateResourceInstanceRequestBody**](UpdateResourceInstanceRequestBody.md) |  | 
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

