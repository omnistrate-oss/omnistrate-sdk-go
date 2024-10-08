# \HostclusterApiAPI

All URIs are relative to *https://api.omnistrate.cloud*

Method | HTTP request | Description
------------- | ------------- | -------------
[**HostclusterApiListHostClusters**](HostclusterApiAPI.md#HostclusterApiListHostClusters) | **Get** /2022-09-01-00/fleet/host-clusters | ListHostClusters hostcluster-api



## HostclusterApiListHostClusters

> ListHostClustersResult HostclusterApiListHostClusters(ctx).ListHostClustersRequestBody(listHostClustersRequestBody).Execute()

ListHostClusters hostcluster-api

### Example

```go
package main

import (
	"context"
	"fmt"
	"os"
	openapiclient "github.com/GIT_USER_ID/GIT_REPO_ID"
)

func main() {
	listHostClustersRequestBody := *openapiclient.NewListHostClustersRequestBody() // ListHostClustersRequestBody | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.HostclusterApiAPI.HostclusterApiListHostClusters(context.Background()).ListHostClustersRequestBody(listHostClustersRequestBody).Execute()
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
 **listHostClustersRequestBody** | [**ListHostClustersRequestBody**](ListHostClustersRequestBody.md) |  | 

### Return type

[**ListHostClustersResult**](ListHostClustersResult.md)

### Authorization

[api_key_header_Authorization](../README.md#api_key_header_Authorization)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json, application/vnd.goa.error

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

