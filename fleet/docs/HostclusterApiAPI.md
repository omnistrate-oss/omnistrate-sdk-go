# \HostclusterApiAPI

All URIs are relative to *https://api.omnistrate.cloud*

Method | HTTP request | Description
------------- | ------------- | -------------
[**HostclusterApiListHostClusters**](HostclusterApiAPI.md#HostclusterApiListHostClusters) | **Get** /2022-09-01-00/fleet/host-clusters | ListHostClusters hostcluster-api



## HostclusterApiListHostClusters

> ListHostClustersResult HostclusterApiListHostClusters(ctx).AccountConfigId(accountConfigId).RegionId(regionId).IncludeProvisionerClusters(includeProvisionerClusters).Execute()

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

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.HostclusterApiAPI.HostclusterApiListHostClusters(context.Background()).AccountConfigId(accountConfigId).RegionId(regionId).IncludeProvisionerClusters(includeProvisionerClusters).Execute()
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

