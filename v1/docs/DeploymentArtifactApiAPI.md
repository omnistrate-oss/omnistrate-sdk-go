# \DeploymentArtifactApiAPI

All URIs are relative to *https://api.omnistrate.cloud*

Method | HTTP request | Description
------------- | ------------- | -------------
[**DeploymentArtifactApiUploadDeploymentArtifact**](DeploymentArtifactApiAPI.md#DeploymentArtifactApiUploadDeploymentArtifact) | **Post** /2022-09-01-00/deployment-artifact | UploadDeploymentArtifact deployment-artifact-api



## DeploymentArtifactApiUploadDeploymentArtifact

> string DeploymentArtifactApiUploadDeploymentArtifact(ctx).UploadDeploymentArtifactRequest2(uploadDeploymentArtifactRequest2).Execute()

UploadDeploymentArtifact deployment-artifact-api

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
	uploadDeploymentArtifactRequest2 := *openapiclient.NewUploadDeploymentArtifactRequest2([]*os.File{"TODO"}, "/path/to", "standard", "my-service") // UploadDeploymentArtifactRequest2 | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.DeploymentArtifactApiAPI.DeploymentArtifactApiUploadDeploymentArtifact(context.Background()).UploadDeploymentArtifactRequest2(uploadDeploymentArtifactRequest2).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `DeploymentArtifactApiAPI.DeploymentArtifactApiUploadDeploymentArtifact``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `DeploymentArtifactApiUploadDeploymentArtifact`: string
	fmt.Fprintf(os.Stdout, "Response from `DeploymentArtifactApiAPI.DeploymentArtifactApiUploadDeploymentArtifact`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiDeploymentArtifactApiUploadDeploymentArtifactRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **uploadDeploymentArtifactRequest2** | [**UploadDeploymentArtifactRequest2**](UploadDeploymentArtifactRequest2.md) |  | 

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

