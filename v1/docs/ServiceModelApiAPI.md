# \ServiceModelApiAPI

All URIs are relative to *https://api.omnistrate.cloud*

Method | HTTP request | Description
------------- | ------------- | -------------
[**ServiceModelApiAddAccountConfigToServiceModel**](ServiceModelApiAPI.md#ServiceModelApiAddAccountConfigToServiceModel) | **Post** /2022-09-01-00/service/{serviceId}/model/{id}/account-config | AddAccountConfigToServiceModel service-model-api
[**ServiceModelApiCopyServiceModel**](ServiceModelApiAPI.md#ServiceModelApiCopyServiceModel) | **Post** /2022-09-01-00/service/{serviceId}/model/copy/{sourceId} | CopyServiceModel service-model-api
[**ServiceModelApiCreateServiceModel**](ServiceModelApiAPI.md#ServiceModelApiCreateServiceModel) | **Post** /2022-09-01-00/service/{serviceId}/model | CreateServiceModel service-model-api
[**ServiceModelApiDeleteServiceModel**](ServiceModelApiAPI.md#ServiceModelApiDeleteServiceModel) | **Delete** /2022-09-01-00/service/{serviceId}/model/{id} | DeleteServiceModel service-model-api
[**ServiceModelApiDescribeServiceModel**](ServiceModelApiAPI.md#ServiceModelApiDescribeServiceModel) | **Get** /2022-09-01-00/service/{serviceId}/model/{id} | DescribeServiceModel service-model-api
[**ServiceModelApiDisableServiceModelFeature**](ServiceModelApiAPI.md#ServiceModelApiDisableServiceModelFeature) | **Delete** /2022-09-01-00/service/{serviceId}/model/{id}/feature | DisableServiceModelFeature service-model-api
[**ServiceModelApiEnableServiceModelFeature**](ServiceModelApiAPI.md#ServiceModelApiEnableServiceModelFeature) | **Put** /2022-09-01-00/service/{serviceId}/model/{id}/feature | EnableServiceModelFeature service-model-api
[**ServiceModelApiListServiceModel**](ServiceModelApiAPI.md#ServiceModelApiListServiceModel) | **Get** /2022-09-01-00/service/{serviceId}/serviceapi/{serviceApiId}/model | ListServiceModel service-model-api
[**ServiceModelApiReleaseServiceModelStatus**](ServiceModelApiAPI.md#ServiceModelApiReleaseServiceModelStatus) | **Get** /2022-09-01-00/service/{serviceId}/model/{id}/release | ReleaseServiceModelStatus service-model-api
[**ServiceModelApiRemoveAccountConfigFromServiceModel**](ServiceModelApiAPI.md#ServiceModelApiRemoveAccountConfigFromServiceModel) | **Delete** /2022-09-01-00/service/{serviceId}/model/{id}/account-config | RemoveAccountConfigFromServiceModel service-model-api
[**ServiceModelApiSetActiveAccountConfig**](ServiceModelApiAPI.md#ServiceModelApiSetActiveAccountConfig) | **Post** /2022-09-01-00/service/{serviceId}/model/{id}/account-config/{accountConfigId}/active | SetActiveAccountConfig service-model-api
[**ServiceModelApiUpdateServiceModel**](ServiceModelApiAPI.md#ServiceModelApiUpdateServiceModel) | **Patch** /2022-09-01-00/service/{serviceId}/model/{id} | UpdateServiceModel service-model-api



## ServiceModelApiAddAccountConfigToServiceModel

> ServiceModelApiAddAccountConfigToServiceModel(ctx, serviceId, id).AddAccountConfigToServiceModelRequest2(addAccountConfigToServiceModelRequest2).Execute()

AddAccountConfigToServiceModel service-model-api

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
	serviceId := "s-12345678" // string | The service ID this model belongs to
	id := "sm-12345678" // string | The service model ID
	addAccountConfigToServiceModelRequest2 := *openapiclient.NewAddAccountConfigToServiceModelRequest2("ac-12345678") // AddAccountConfigToServiceModelRequest2 | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	r, err := apiClient.ServiceModelApiAPI.ServiceModelApiAddAccountConfigToServiceModel(context.Background(), serviceId, id).AddAccountConfigToServiceModelRequest2(addAccountConfigToServiceModelRequest2).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `ServiceModelApiAPI.ServiceModelApiAddAccountConfigToServiceModel``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**serviceId** | **string** | The service ID this model belongs to | 
**id** | **string** | The service model ID | 

### Other Parameters

Other parameters are passed through a pointer to a apiServiceModelApiAddAccountConfigToServiceModelRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


 **addAccountConfigToServiceModelRequest2** | [**AddAccountConfigToServiceModelRequest2**](AddAccountConfigToServiceModelRequest2.md) |  | 

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


## ServiceModelApiCopyServiceModel

> string ServiceModelApiCopyServiceModel(ctx, serviceId, sourceId).CopyServiceModelRequest2(copyServiceModelRequest2).Execute()

CopyServiceModel service-model-api

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
	serviceId := "s-12345678" // string | The service ID this model belongs to
	sourceId := "sm-12345678" // string | The source service model ID
	copyServiceModelRequest2 := *openapiclient.NewCopyServiceModelRequest2("A MySQL Hosted SaaS specializing in multi-writer clusters for high availability", "MySQL multi-writer service hosted model", "CUSTOMER_HOSTED") // CopyServiceModelRequest2 | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.ServiceModelApiAPI.ServiceModelApiCopyServiceModel(context.Background(), serviceId, sourceId).CopyServiceModelRequest2(copyServiceModelRequest2).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `ServiceModelApiAPI.ServiceModelApiCopyServiceModel``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `ServiceModelApiCopyServiceModel`: string
	fmt.Fprintf(os.Stdout, "Response from `ServiceModelApiAPI.ServiceModelApiCopyServiceModel`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**serviceId** | **string** | The service ID this model belongs to | 
**sourceId** | **string** | The source service model ID | 

### Other Parameters

Other parameters are passed through a pointer to a apiServiceModelApiCopyServiceModelRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


 **copyServiceModelRequest2** | [**CopyServiceModelRequest2**](CopyServiceModelRequest2.md) |  | 

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


## ServiceModelApiCreateServiceModel

> string ServiceModelApiCreateServiceModel(ctx, serviceId).CreateServiceModelRequest2(createServiceModelRequest2).Execute()

CreateServiceModel service-model-api

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
	serviceId := "s-12345678" // string | The service this model is for
	createServiceModelRequest2 := *openapiclient.NewCreateServiceModelRequest2("A MySQL Hosted SaaS specializing in multi-writer clusters for high availability", "CUSTOMER_HOSTED", "MySQL multi-writer service hosted model", "sa-12345678") // CreateServiceModelRequest2 | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.ServiceModelApiAPI.ServiceModelApiCreateServiceModel(context.Background(), serviceId).CreateServiceModelRequest2(createServiceModelRequest2).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `ServiceModelApiAPI.ServiceModelApiCreateServiceModel``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `ServiceModelApiCreateServiceModel`: string
	fmt.Fprintf(os.Stdout, "Response from `ServiceModelApiAPI.ServiceModelApiCreateServiceModel`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**serviceId** | **string** | The service this model is for | 

### Other Parameters

Other parameters are passed through a pointer to a apiServiceModelApiCreateServiceModelRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **createServiceModelRequest2** | [**CreateServiceModelRequest2**](CreateServiceModelRequest2.md) |  | 

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


## ServiceModelApiDeleteServiceModel

> ServiceModelApiDeleteServiceModel(ctx, serviceId, id).Execute()

DeleteServiceModel service-model-api

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
	serviceId := "s-12345678" // string | The service ID this model belongs to
	id := "sm-12345678" // string | The service model ID

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	r, err := apiClient.ServiceModelApiAPI.ServiceModelApiDeleteServiceModel(context.Background(), serviceId, id).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `ServiceModelApiAPI.ServiceModelApiDeleteServiceModel``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**serviceId** | **string** | The service ID this model belongs to | 
**id** | **string** | The service model ID | 

### Other Parameters

Other parameters are passed through a pointer to a apiServiceModelApiDeleteServiceModelRequest struct via the builder pattern


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


## ServiceModelApiDescribeServiceModel

> DescribeServiceModelResult ServiceModelApiDescribeServiceModel(ctx, serviceId, id).Execute()

DescribeServiceModel service-model-api

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
	serviceId := "s-12345678" // string | The service ID this model belongs to
	id := "sm-12345678" // string | The service model ID

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.ServiceModelApiAPI.ServiceModelApiDescribeServiceModel(context.Background(), serviceId, id).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `ServiceModelApiAPI.ServiceModelApiDescribeServiceModel``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `ServiceModelApiDescribeServiceModel`: DescribeServiceModelResult
	fmt.Fprintf(os.Stdout, "Response from `ServiceModelApiAPI.ServiceModelApiDescribeServiceModel`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**serviceId** | **string** | The service ID this model belongs to | 
**id** | **string** | The service model ID | 

### Other Parameters

Other parameters are passed through a pointer to a apiServiceModelApiDescribeServiceModelRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------



### Return type

[**DescribeServiceModelResult**](DescribeServiceModelResult.md)

### Authorization

[api_key_header_Authorization](../README.md#api_key_header_Authorization)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json, application/vnd.goa.error

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## ServiceModelApiDisableServiceModelFeature

> ServiceModelApiDisableServiceModelFeature(ctx, serviceId, id).DisableServiceModelFeatureRequest2(disableServiceModelFeatureRequest2).Execute()

DisableServiceModelFeature service-model-api

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
	serviceId := "s-12345678" // string | The service ID this model belongs to
	id := "sm-12345678" // string | The service model ID
	disableServiceModelFeatureRequest2 := *openapiclient.NewDisableServiceModelFeatureRequest2("Possimus vitae eveniet est eum.") // DisableServiceModelFeatureRequest2 | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	r, err := apiClient.ServiceModelApiAPI.ServiceModelApiDisableServiceModelFeature(context.Background(), serviceId, id).DisableServiceModelFeatureRequest2(disableServiceModelFeatureRequest2).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `ServiceModelApiAPI.ServiceModelApiDisableServiceModelFeature``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**serviceId** | **string** | The service ID this model belongs to | 
**id** | **string** | The service model ID | 

### Other Parameters

Other parameters are passed through a pointer to a apiServiceModelApiDisableServiceModelFeatureRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


 **disableServiceModelFeatureRequest2** | [**DisableServiceModelFeatureRequest2**](DisableServiceModelFeatureRequest2.md) |  | 

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


## ServiceModelApiEnableServiceModelFeature

> ServiceModelApiEnableServiceModelFeature(ctx, serviceId, id).EnableServiceModelFeatureRequest2(enableServiceModelFeatureRequest2).Execute()

EnableServiceModelFeature service-model-api

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
	serviceId := "s-12345678" // string | The service ID this model belongs to
	id := "sm-12345678" // string | The service model ID
	enableServiceModelFeatureRequest2 := *openapiclient.NewEnableServiceModelFeatureRequest2(map[string]interface{}{"key": interface{}(123)}, "Quis doloribus consequuntur.") // EnableServiceModelFeatureRequest2 | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	r, err := apiClient.ServiceModelApiAPI.ServiceModelApiEnableServiceModelFeature(context.Background(), serviceId, id).EnableServiceModelFeatureRequest2(enableServiceModelFeatureRequest2).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `ServiceModelApiAPI.ServiceModelApiEnableServiceModelFeature``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**serviceId** | **string** | The service ID this model belongs to | 
**id** | **string** | The service model ID | 

### Other Parameters

Other parameters are passed through a pointer to a apiServiceModelApiEnableServiceModelFeatureRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


 **enableServiceModelFeatureRequest2** | [**EnableServiceModelFeatureRequest2**](EnableServiceModelFeatureRequest2.md) |  | 

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


## ServiceModelApiListServiceModel

> ListServiceModelsResult ServiceModelApiListServiceModel(ctx, serviceId, serviceApiId).Execute()

ListServiceModel service-model-api

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
	serviceId := "s-12345678" // string | The service ID
	serviceApiId := "sa-12345678" // string | The service API ID

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.ServiceModelApiAPI.ServiceModelApiListServiceModel(context.Background(), serviceId, serviceApiId).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `ServiceModelApiAPI.ServiceModelApiListServiceModel``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `ServiceModelApiListServiceModel`: ListServiceModelsResult
	fmt.Fprintf(os.Stdout, "Response from `ServiceModelApiAPI.ServiceModelApiListServiceModel`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**serviceId** | **string** | The service ID | 
**serviceApiId** | **string** | The service API ID | 

### Other Parameters

Other parameters are passed through a pointer to a apiServiceModelApiListServiceModelRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------



### Return type

[**ListServiceModelsResult**](ListServiceModelsResult.md)

### Authorization

[api_key_header_Authorization](../README.md#api_key_header_Authorization)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json, application/vnd.goa.error

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## ServiceModelApiReleaseServiceModelStatus

> ReleaseServiceModelResult ServiceModelApiReleaseServiceModelStatus(ctx, serviceId, id).Execute()

ReleaseServiceModelStatus service-model-api

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
	serviceId := "s-12345678" // string | The service ID this model belongs to
	id := "sm-12345678" // string | The service model ID

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.ServiceModelApiAPI.ServiceModelApiReleaseServiceModelStatus(context.Background(), serviceId, id).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `ServiceModelApiAPI.ServiceModelApiReleaseServiceModelStatus``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `ServiceModelApiReleaseServiceModelStatus`: ReleaseServiceModelResult
	fmt.Fprintf(os.Stdout, "Response from `ServiceModelApiAPI.ServiceModelApiReleaseServiceModelStatus`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**serviceId** | **string** | The service ID this model belongs to | 
**id** | **string** | The service model ID | 

### Other Parameters

Other parameters are passed through a pointer to a apiServiceModelApiReleaseServiceModelStatusRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------



### Return type

[**ReleaseServiceModelResult**](ReleaseServiceModelResult.md)

### Authorization

[api_key_header_Authorization](../README.md#api_key_header_Authorization)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json, application/vnd.goa.error

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## ServiceModelApiRemoveAccountConfigFromServiceModel

> ServiceModelApiRemoveAccountConfigFromServiceModel(ctx, serviceId, id).RemoveAccountConfigFromServiceModelRequest2(removeAccountConfigFromServiceModelRequest2).Execute()

RemoveAccountConfigFromServiceModel service-model-api

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
	serviceId := "s-12345678" // string | The service ID this model belongs to
	id := "sm-12345678" // string | The service model ID
	removeAccountConfigFromServiceModelRequest2 := *openapiclient.NewRemoveAccountConfigFromServiceModelRequest2("ac-12345678") // RemoveAccountConfigFromServiceModelRequest2 | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	r, err := apiClient.ServiceModelApiAPI.ServiceModelApiRemoveAccountConfigFromServiceModel(context.Background(), serviceId, id).RemoveAccountConfigFromServiceModelRequest2(removeAccountConfigFromServiceModelRequest2).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `ServiceModelApiAPI.ServiceModelApiRemoveAccountConfigFromServiceModel``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**serviceId** | **string** | The service ID this model belongs to | 
**id** | **string** | The service model ID | 

### Other Parameters

Other parameters are passed through a pointer to a apiServiceModelApiRemoveAccountConfigFromServiceModelRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


 **removeAccountConfigFromServiceModelRequest2** | [**RemoveAccountConfigFromServiceModelRequest2**](RemoveAccountConfigFromServiceModelRequest2.md) |  | 

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


## ServiceModelApiSetActiveAccountConfig

> ServiceModelApiSetActiveAccountConfig(ctx, serviceId, id, accountConfigId).Execute()

SetActiveAccountConfig service-model-api

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
	serviceId := "s-12345678" // string | The service ID this model belongs to
	id := "sm-12345678" // string | The service model ID
	accountConfigId := "ac-12345678" // string | The infrastructure account configuration ID

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	r, err := apiClient.ServiceModelApiAPI.ServiceModelApiSetActiveAccountConfig(context.Background(), serviceId, id, accountConfigId).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `ServiceModelApiAPI.ServiceModelApiSetActiveAccountConfig``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**serviceId** | **string** | The service ID this model belongs to | 
**id** | **string** | The service model ID | 
**accountConfigId** | **string** | The infrastructure account configuration ID | 

### Other Parameters

Other parameters are passed through a pointer to a apiServiceModelApiSetActiveAccountConfigRequest struct via the builder pattern


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


## ServiceModelApiUpdateServiceModel

> ServiceModelApiUpdateServiceModel(ctx, serviceId, id).UpdateServiceModelRequest2(updateServiceModelRequest2).Execute()

UpdateServiceModel service-model-api

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
	serviceId := "s-12345678" // string | The service ID this model belongs to
	id := "sm-12345678" // string | The service model ID
	updateServiceModelRequest2 := *openapiclient.NewUpdateServiceModelRequest2() // UpdateServiceModelRequest2 | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	r, err := apiClient.ServiceModelApiAPI.ServiceModelApiUpdateServiceModel(context.Background(), serviceId, id).UpdateServiceModelRequest2(updateServiceModelRequest2).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `ServiceModelApiAPI.ServiceModelApiUpdateServiceModel``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**serviceId** | **string** | The service ID this model belongs to | 
**id** | **string** | The service model ID | 

### Other Parameters

Other parameters are passed through a pointer to a apiServiceModelApiUpdateServiceModelRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


 **updateServiceModelRequest2** | [**UpdateServiceModelRequest2**](UpdateServiceModelRequest2.md) |  | 

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

