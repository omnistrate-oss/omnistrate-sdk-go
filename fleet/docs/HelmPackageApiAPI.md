# \HelmPackageApiAPI

All URIs are relative to *https://api.omnistrate.cloud*

Method | HTTP request | Description
------------- | ------------- | -------------
[**HelmPackageApiListHelmPackageInstallations**](HelmPackageApiAPI.md#HelmPackageApiListHelmPackageInstallations) | **Get** /2022-09-01-00/fleet/helm-package-installations | ListHelmPackageInstallations helm-package-api



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

