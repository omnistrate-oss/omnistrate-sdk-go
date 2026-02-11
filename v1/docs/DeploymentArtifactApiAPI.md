# \DeploymentArtifactApiAPI

All URIs are relative to *https://api.omnistrate.cloud*

Method | HTTP request | Description
------------- | ------------- | -------------
[**DeploymentArtifactApiDescribeDeploymentArtifact**](DeploymentArtifactApiAPI.md#DeploymentArtifactApiDescribeDeploymentArtifact) | **Get** /2022-09-01-00/deployment-artifact/{id} | DescribeDeploymentArtifact deployment-artifact-api
[**DeploymentArtifactApiUploadDeploymentArtifact**](DeploymentArtifactApiAPI.md#DeploymentArtifactApiUploadDeploymentArtifact) | **Post** /2022-09-01-00/deployment-artifact | UploadDeploymentArtifact deployment-artifact-api



## DeploymentArtifactApiDescribeDeploymentArtifact

> DescribeDeploymentArtifactResult DeploymentArtifactApiDescribeDeploymentArtifact(ctx, id).Execute()

DescribeDeploymentArtifact deployment-artifact-api

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
	id := "da-12345678" // string | The ID of the deployment artifact

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.DeploymentArtifactApiAPI.DeploymentArtifactApiDescribeDeploymentArtifact(context.Background(), id).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `DeploymentArtifactApiAPI.DeploymentArtifactApiDescribeDeploymentArtifact``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `DeploymentArtifactApiDescribeDeploymentArtifact`: DescribeDeploymentArtifactResult
	fmt.Fprintf(os.Stdout, "Response from `DeploymentArtifactApiAPI.DeploymentArtifactApiDescribeDeploymentArtifact`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **string** | The ID of the deployment artifact | 

### Other Parameters

Other parameters are passed through a pointer to a apiDeploymentArtifactApiDescribeDeploymentArtifactRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

[**DescribeDeploymentArtifactResult**](DescribeDeploymentArtifactResult.md)

### Authorization

[api_key_header_Authorization](../README.md#api_key_header_Authorization)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json, application/vnd.goa.error

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


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
	uploadDeploymentArtifactRequest2 := *openapiclient.NewUploadDeploymentArtifactRequest2("ac-1234567890", "/path/to", "Harum quia quia a.", "PROD|PRIVATE|CANARY|STAGING|QA|DEV|GLOBAL", "standard", "my-service") // UploadDeploymentArtifactRequest2 | 

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

