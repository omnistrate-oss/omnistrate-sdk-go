# \FleetWorkflowsApiAPI

All URIs are relative to *https://api.omnistrate.cloud*

Method | HTTP request | Description
------------- | ------------- | -------------
[**FleetWorkflowsApiDescribeServiceWorkflow**](FleetWorkflowsApiAPI.md#FleetWorkflowsApiDescribeServiceWorkflow) | **Get** /2022-09-01-00/fleet/service/{serviceId}/environment/{environmentId}/service-workflows/{id} | DescribeServiceWorkflow fleet-workflows-api
[**FleetWorkflowsApiDescribeServiceWorkflowSummary**](FleetWorkflowsApiAPI.md#FleetWorkflowsApiDescribeServiceWorkflowSummary) | **Get** /2022-09-01-00/fleet/service/{serviceId}/environment/{environmentId}/service-workflows-summary | DescribeServiceWorkflowSummary fleet-workflows-api
[**FleetWorkflowsApiGetWorkflowEvents**](FleetWorkflowsApiAPI.md#FleetWorkflowsApiGetWorkflowEvents) | **Get** /2022-09-01-00/fleet/service/{serviceId}/environment/{environmentId}/service-workflows/{id}/events | GetWorkflowEvents fleet-workflows-api
[**FleetWorkflowsApiListServiceWorkflows**](FleetWorkflowsApiAPI.md#FleetWorkflowsApiListServiceWorkflows) | **Get** /2022-09-01-00/fleet/service/{serviceId}/environment/{environmentId}/service-workflows | ListServiceWorkflows fleet-workflows-api
[**FleetWorkflowsApiTerminateServiceWorkflow**](FleetWorkflowsApiAPI.md#FleetWorkflowsApiTerminateServiceWorkflow) | **Delete** /2022-09-01-00/fleet/service/{serviceId}/environment/{environmentId}/service-workflows/{id} | TerminateServiceWorkflow fleet-workflows-api
[**FleetWorkflowsApiUpdateServiceWorkflow**](FleetWorkflowsApiAPI.md#FleetWorkflowsApiUpdateServiceWorkflow) | **Patch** /2022-09-01-00/fleet/service/{serviceId}/environment/{environmentId}/service-workflows/{id} | UpdateServiceWorkflow fleet-workflows-api



## FleetWorkflowsApiDescribeServiceWorkflow

> DescribeServiceWorkflowResult FleetWorkflowsApiDescribeServiceWorkflow(ctx, serviceId, environmentId, id).Execute()

DescribeServiceWorkflow fleet-workflows-api

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
	id := "submit-create-instance-plan-instance-50h74sj46" // string | ID of the ServiceWorkflow

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.FleetWorkflowsApiAPI.FleetWorkflowsApiDescribeServiceWorkflow(context.Background(), serviceId, environmentId, id).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `FleetWorkflowsApiAPI.FleetWorkflowsApiDescribeServiceWorkflow``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `FleetWorkflowsApiDescribeServiceWorkflow`: DescribeServiceWorkflowResult
	fmt.Fprintf(os.Stdout, "Response from `FleetWorkflowsApiAPI.FleetWorkflowsApiDescribeServiceWorkflow`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**serviceId** | **string** | The service ID this workflow belongs to. | 
**environmentId** | **string** | The service environment ID this workflow belongs to. | 
**id** | **string** | ID of the ServiceWorkflow | 

### Other Parameters

Other parameters are passed through a pointer to a apiFleetWorkflowsApiDescribeServiceWorkflowRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------




### Return type

[**DescribeServiceWorkflowResult**](DescribeServiceWorkflowResult.md)

### Authorization

[api_key_header_Authorization](../README.md#api_key_header_Authorization)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json, application/vnd.goa.error

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## FleetWorkflowsApiDescribeServiceWorkflowSummary

> DescribeServiceWorkflowSummaryResult FleetWorkflowsApiDescribeServiceWorkflowSummary(ctx, serviceId, environmentId).Execute()

DescribeServiceWorkflowSummary fleet-workflows-api

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
	resp, r, err := apiClient.FleetWorkflowsApiAPI.FleetWorkflowsApiDescribeServiceWorkflowSummary(context.Background(), serviceId, environmentId).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `FleetWorkflowsApiAPI.FleetWorkflowsApiDescribeServiceWorkflowSummary``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `FleetWorkflowsApiDescribeServiceWorkflowSummary`: DescribeServiceWorkflowSummaryResult
	fmt.Fprintf(os.Stdout, "Response from `FleetWorkflowsApiAPI.FleetWorkflowsApiDescribeServiceWorkflowSummary`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**serviceId** | **string** | The service ID this workflow belongs to. | 
**environmentId** | **string** | The service environment ID this workflow belongs to. | 

### Other Parameters

Other parameters are passed through a pointer to a apiFleetWorkflowsApiDescribeServiceWorkflowSummaryRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------



### Return type

[**DescribeServiceWorkflowSummaryResult**](DescribeServiceWorkflowSummaryResult.md)

### Authorization

[api_key_header_Authorization](../README.md#api_key_header_Authorization)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json, application/vnd.goa.error

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## FleetWorkflowsApiGetWorkflowEvents

> GetWorkflowEventsResult FleetWorkflowsApiGetWorkflowEvents(ctx, serviceId, environmentId, id).Execute()

GetWorkflowEvents fleet-workflows-api

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
	id := "submit-create-instance-plan-instance-50h74sj46" // string | ID of the ServiceWorkflow

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.FleetWorkflowsApiAPI.FleetWorkflowsApiGetWorkflowEvents(context.Background(), serviceId, environmentId, id).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `FleetWorkflowsApiAPI.FleetWorkflowsApiGetWorkflowEvents``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `FleetWorkflowsApiGetWorkflowEvents`: GetWorkflowEventsResult
	fmt.Fprintf(os.Stdout, "Response from `FleetWorkflowsApiAPI.FleetWorkflowsApiGetWorkflowEvents`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**serviceId** | **string** | The service ID this workflow belongs to. | 
**environmentId** | **string** | The service environment ID this workflow belongs to. | 
**id** | **string** | ID of the ServiceWorkflow | 

### Other Parameters

Other parameters are passed through a pointer to a apiFleetWorkflowsApiGetWorkflowEventsRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------




### Return type

[**GetWorkflowEventsResult**](GetWorkflowEventsResult.md)

### Authorization

[api_key_header_Authorization](../README.md#api_key_header_Authorization)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json, application/vnd.goa.error

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## FleetWorkflowsApiListServiceWorkflows

> ListServiceWorkflowsResult FleetWorkflowsApiListServiceWorkflows(ctx, serviceId, environmentId).InstanceId(instanceId).NextPageToken(nextPageToken).PageSize(pageSize).Execute()

ListServiceWorkflows fleet-workflows-api

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
	instanceId := "instance-12345678" // string | The instance ID of the workflow (optional)
	nextPageToken := "token" // string | The next token to use for pagination (optional)
	pageSize := int64(10) // int64 | The number of resources to return per page (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.FleetWorkflowsApiAPI.FleetWorkflowsApiListServiceWorkflows(context.Background(), serviceId, environmentId).InstanceId(instanceId).NextPageToken(nextPageToken).PageSize(pageSize).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `FleetWorkflowsApiAPI.FleetWorkflowsApiListServiceWorkflows``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `FleetWorkflowsApiListServiceWorkflows`: ListServiceWorkflowsResult
	fmt.Fprintf(os.Stdout, "Response from `FleetWorkflowsApiAPI.FleetWorkflowsApiListServiceWorkflows`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**serviceId** | **string** | The service ID this workflow belongs to. | 
**environmentId** | **string** | The service environment ID this workflow belongs to. | 

### Other Parameters

Other parameters are passed through a pointer to a apiFleetWorkflowsApiListServiceWorkflowsRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


 **instanceId** | **string** | The instance ID of the workflow | 
 **nextPageToken** | **string** | The next token to use for pagination | 
 **pageSize** | **int64** | The number of resources to return per page | 

### Return type

[**ListServiceWorkflowsResult**](ListServiceWorkflowsResult.md)

### Authorization

[api_key_header_Authorization](../README.md#api_key_header_Authorization)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json, application/vnd.goa.error

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## FleetWorkflowsApiTerminateServiceWorkflow

> ServiceWorkflow FleetWorkflowsApiTerminateServiceWorkflow(ctx, serviceId, environmentId, id).Execute()

TerminateServiceWorkflow fleet-workflows-api

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
	id := "submit-create-instance-plan-instance-50h74sj46" // string | ID of the ServiceWorkflow

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.FleetWorkflowsApiAPI.FleetWorkflowsApiTerminateServiceWorkflow(context.Background(), serviceId, environmentId, id).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `FleetWorkflowsApiAPI.FleetWorkflowsApiTerminateServiceWorkflow``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `FleetWorkflowsApiTerminateServiceWorkflow`: ServiceWorkflow
	fmt.Fprintf(os.Stdout, "Response from `FleetWorkflowsApiAPI.FleetWorkflowsApiTerminateServiceWorkflow`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**serviceId** | **string** | The service ID this workflow belongs to. | 
**environmentId** | **string** | The service environment ID this workflow belongs to. | 
**id** | **string** | ID of the ServiceWorkflow | 

### Other Parameters

Other parameters are passed through a pointer to a apiFleetWorkflowsApiTerminateServiceWorkflowRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------




### Return type

[**ServiceWorkflow**](ServiceWorkflow.md)

### Authorization

[api_key_header_Authorization](../README.md#api_key_header_Authorization)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json, application/vnd.goa.error

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## FleetWorkflowsApiUpdateServiceWorkflow

> ServiceWorkflow FleetWorkflowsApiUpdateServiceWorkflow(ctx, serviceId, environmentId, id).UpdateServiceWorkflowRequestBody(updateServiceWorkflowRequestBody).Execute()

UpdateServiceWorkflow fleet-workflows-api

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
	id := "submit-create-instance-plan-instance-50h74sj46" // string | ID of the ServiceWorkflow
	updateServiceWorkflowRequestBody := *openapiclient.NewUpdateServiceWorkflowRequestBody("Pause") // UpdateServiceWorkflowRequestBody | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.FleetWorkflowsApiAPI.FleetWorkflowsApiUpdateServiceWorkflow(context.Background(), serviceId, environmentId, id).UpdateServiceWorkflowRequestBody(updateServiceWorkflowRequestBody).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `FleetWorkflowsApiAPI.FleetWorkflowsApiUpdateServiceWorkflow``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `FleetWorkflowsApiUpdateServiceWorkflow`: ServiceWorkflow
	fmt.Fprintf(os.Stdout, "Response from `FleetWorkflowsApiAPI.FleetWorkflowsApiUpdateServiceWorkflow`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**serviceId** | **string** | The service ID this workflow belongs to. | 
**environmentId** | **string** | The service environment ID this workflow belongs to. | 
**id** | **string** | ID of the ServiceWorkflow | 

### Other Parameters

Other parameters are passed through a pointer to a apiFleetWorkflowsApiUpdateServiceWorkflowRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------



 **updateServiceWorkflowRequestBody** | [**UpdateServiceWorkflowRequestBody**](UpdateServiceWorkflowRequestBody.md) |  | 

### Return type

[**ServiceWorkflow**](ServiceWorkflow.md)

### Authorization

[api_key_header_Authorization](../README.md#api_key_header_Authorization)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json, application/vnd.goa.error

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

