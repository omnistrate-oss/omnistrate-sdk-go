# \HelmPackageApiAPI

All URIs are relative to *https://api.omnistrate.cloud*

Method | HTTP request | Description
------------- | ------------- | -------------
[**HelmPackageApiDeleteHelmPackage**](HelmPackageApiAPI.md#HelmPackageApiDeleteHelmPackage) | **Delete** /2022-09-01-00/fleet/helm-package/{chartName}/{chartVersion} | DeleteHelmPackage helm-package-api
[**HelmPackageApiDescribeHelmPackage**](HelmPackageApiAPI.md#HelmPackageApiDescribeHelmPackage) | **Get** /2022-09-01-00/fleet/helm-package/{chartName}/{chartVersion} | DescribeHelmPackage helm-package-api
[**HelmPackageApiListHelmPackageInstallations**](HelmPackageApiAPI.md#HelmPackageApiListHelmPackageInstallations) | **Get** /2022-09-01-00/fleet/helm-package-installations | ListHelmPackageInstallations helm-package-api
[**HelmPackageApiListHelmPackages**](HelmPackageApiAPI.md#HelmPackageApiListHelmPackages) | **Get** /2022-09-01-00/fleet/helm-package | ListHelmPackages helm-package-api
[**HelmPackageApiSaveHelmPackage**](HelmPackageApiAPI.md#HelmPackageApiSaveHelmPackage) | **Post** /2022-09-01-00/fleet/helm-package | SaveHelmPackage helm-package-api



## HelmPackageApiDeleteHelmPackage

> HelmPackageApiDeleteHelmPackage(ctx, chartName, chartVersion).Execute()

DeleteHelmPackage helm-package-api

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
	chartName := "my-chart" // string | The chart name of the Helm package
	chartVersion := "1.0.0" // string | The chart version of the Helm package

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	r, err := apiClient.HelmPackageApiAPI.HelmPackageApiDeleteHelmPackage(context.Background(), chartName, chartVersion).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `HelmPackageApiAPI.HelmPackageApiDeleteHelmPackage``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**chartName** | **string** | The chart name of the Helm package | 
**chartVersion** | **string** | The chart version of the Helm package | 

### Other Parameters

Other parameters are passed through a pointer to a apiHelmPackageApiDeleteHelmPackageRequest struct via the builder pattern


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


## HelmPackageApiDescribeHelmPackage

> HelmPackage HelmPackageApiDescribeHelmPackage(ctx, chartName, chartVersion).Execute()

DescribeHelmPackage helm-package-api

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
	chartName := "my-chart" // string | The chart name of the Helm package
	chartVersion := "1.0.0" // string | The chart version of the Helm package

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.HelmPackageApiAPI.HelmPackageApiDescribeHelmPackage(context.Background(), chartName, chartVersion).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `HelmPackageApiAPI.HelmPackageApiDescribeHelmPackage``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `HelmPackageApiDescribeHelmPackage`: HelmPackage
	fmt.Fprintf(os.Stdout, "Response from `HelmPackageApiAPI.HelmPackageApiDescribeHelmPackage`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**chartName** | **string** | The chart name of the Helm package | 
**chartVersion** | **string** | The chart version of the Helm package | 

### Other Parameters

Other parameters are passed through a pointer to a apiHelmPackageApiDescribeHelmPackageRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------



### Return type

[**HelmPackage**](HelmPackage.md)

### Authorization

[api_key_header_Authorization](../README.md#api_key_header_Authorization)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json, application/vnd.goa.error

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## HelmPackageApiListHelmPackageInstallations

> ListHelmPackageInstallationsResult HelmPackageApiListHelmPackageInstallations(ctx).HostClusterID(hostClusterID).Execute()

ListHelmPackageInstallations helm-package-api

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
	hostClusterID := "hc-12345678" // string | The host cluster ID to list helm packages for (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.HelmPackageApiAPI.HelmPackageApiListHelmPackageInstallations(context.Background()).HostClusterID(hostClusterID).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `HelmPackageApiAPI.HelmPackageApiListHelmPackageInstallations``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `HelmPackageApiListHelmPackageInstallations`: ListHelmPackageInstallationsResult
	fmt.Fprintf(os.Stdout, "Response from `HelmPackageApiAPI.HelmPackageApiListHelmPackageInstallations`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiHelmPackageApiListHelmPackageInstallationsRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **hostClusterID** | **string** | The host cluster ID to list helm packages for | 

### Return type

[**ListHelmPackageInstallationsResult**](ListHelmPackageInstallationsResult.md)

### Authorization

[api_key_header_Authorization](../README.md#api_key_header_Authorization)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json, application/vnd.goa.error

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## HelmPackageApiListHelmPackages

> ListHelmPackagesResult HelmPackageApiListHelmPackages(ctx).Execute()

ListHelmPackages helm-package-api

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

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.HelmPackageApiAPI.HelmPackageApiListHelmPackages(context.Background()).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `HelmPackageApiAPI.HelmPackageApiListHelmPackages``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `HelmPackageApiListHelmPackages`: ListHelmPackagesResult
	fmt.Fprintf(os.Stdout, "Response from `HelmPackageApiAPI.HelmPackageApiListHelmPackages`: %v\n", resp)
}
```

### Path Parameters

This endpoint does not need any parameter.

### Other Parameters

Other parameters are passed through a pointer to a apiHelmPackageApiListHelmPackagesRequest struct via the builder pattern


### Return type

[**ListHelmPackagesResult**](ListHelmPackagesResult.md)

### Authorization

[api_key_header_Authorization](../README.md#api_key_header_Authorization)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json, application/vnd.goa.error

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## HelmPackageApiSaveHelmPackage

> HelmPackage HelmPackageApiSaveHelmPackage(ctx).SaveHelmPackageRequestBody(saveHelmPackageRequestBody).Execute()

SaveHelmPackage helm-package-api

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
	saveHelmPackageRequestBody := *openapiclient.NewSaveHelmPackageRequestBody(*openapiclient.NewHelmPackage("my-chart", "1.0.0", "default", "https://charts.bitnami.com/bitnami")) // SaveHelmPackageRequestBody | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.HelmPackageApiAPI.HelmPackageApiSaveHelmPackage(context.Background()).SaveHelmPackageRequestBody(saveHelmPackageRequestBody).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `HelmPackageApiAPI.HelmPackageApiSaveHelmPackage``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `HelmPackageApiSaveHelmPackage`: HelmPackage
	fmt.Fprintf(os.Stdout, "Response from `HelmPackageApiAPI.HelmPackageApiSaveHelmPackage`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiHelmPackageApiSaveHelmPackageRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **saveHelmPackageRequestBody** | [**SaveHelmPackageRequestBody**](SaveHelmPackageRequestBody.md) |  | 

### Return type

[**HelmPackage**](HelmPackage.md)

### Authorization

[api_key_header_Authorization](../README.md#api_key_header_Authorization)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json, application/vnd.goa.error

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

