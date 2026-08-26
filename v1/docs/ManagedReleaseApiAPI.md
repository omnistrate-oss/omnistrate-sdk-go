# \ManagedReleaseApiAPI

All URIs are relative to *https://api.omnistrate.cloud*

Method | HTTP request | Description
------------- | ------------- | -------------
[**ManagedReleaseApiDescribeManagedReleaseRevision**](ManagedReleaseApiAPI.md#ManagedReleaseApiDescribeManagedReleaseRevision) | **Get** /2022-09-01-00/internal/v1/managed-release-revisions/{bundleVersion} | DescribeManagedReleaseRevision managed-release-api
[**ManagedReleaseApiDescribeManagedReleaseSync**](ManagedReleaseApiAPI.md#ManagedReleaseApiDescribeManagedReleaseSync) | **Get** /2022-09-01-00/internal/v1/organizations/{organizationId}/provisioner-targets/{provisionerTargetId}/managed-release-syncs/{bundleRevisionId} | DescribeManagedReleaseSync managed-release-api
[**ManagedReleaseApiListManagedReleaseRevisions**](ManagedReleaseApiAPI.md#ManagedReleaseApiListManagedReleaseRevisions) | **Get** /2022-09-01-00/internal/v1/managed-release-revisions | ListManagedReleaseRevisions managed-release-api
[**ManagedReleaseApiListManagedReleaseSyncs**](ManagedReleaseApiAPI.md#ManagedReleaseApiListManagedReleaseSyncs) | **Get** /2022-09-01-00/internal/v1/organizations/{organizationId}/provisioner-targets/{provisionerTargetId}/managed-release-syncs | ListManagedReleaseSyncs managed-release-api
[**ManagedReleaseApiPublishManagedRelease**](ManagedReleaseApiAPI.md#ManagedReleaseApiPublishManagedRelease) | **Post** /2022-09-01-00/internal/v1/managed-release-revisions:publish-managed-release | PublishManagedRelease managed-release-api



## ManagedReleaseApiDescribeManagedReleaseRevision

> ManagedReleaseRevision3 ManagedReleaseApiDescribeManagedReleaseRevision(ctx, bundleVersion).IncludeArtifacts(includeArtifacts).Execute()

DescribeManagedReleaseRevision managed-release-api

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
	bundleVersion := "e6" // string | 
	includeArtifacts := false // bool |  (optional) (default to true)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.ManagedReleaseApiAPI.ManagedReleaseApiDescribeManagedReleaseRevision(context.Background(), bundleVersion).IncludeArtifacts(includeArtifacts).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `ManagedReleaseApiAPI.ManagedReleaseApiDescribeManagedReleaseRevision``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `ManagedReleaseApiDescribeManagedReleaseRevision`: ManagedReleaseRevision3
	fmt.Fprintf(os.Stdout, "Response from `ManagedReleaseApiAPI.ManagedReleaseApiDescribeManagedReleaseRevision`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**bundleVersion** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiManagedReleaseApiDescribeManagedReleaseRevisionRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **includeArtifacts** | **bool** |  | [default to true]

### Return type

[**ManagedReleaseRevision3**](ManagedReleaseRevision3.md)

### Authorization

[api_key_header_Authorization](../README.md#api_key_header_Authorization)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json, application/vnd.goa.error

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## ManagedReleaseApiDescribeManagedReleaseSync

> ManagedReleaseSync ManagedReleaseApiDescribeManagedReleaseSync(ctx, organizationId, provisionerTargetId, bundleRevisionId).IncludeArtifactResults(includeArtifactResults).Execute()

DescribeManagedReleaseSync managed-release-api

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
	organizationId := "k4" // string | 
	provisionerTargetId := "u" // string | 
	bundleRevisionId := "oabr-vA" // string | 
	includeArtifactResults := false // bool |  (optional) (default to true)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.ManagedReleaseApiAPI.ManagedReleaseApiDescribeManagedReleaseSync(context.Background(), organizationId, provisionerTargetId, bundleRevisionId).IncludeArtifactResults(includeArtifactResults).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `ManagedReleaseApiAPI.ManagedReleaseApiDescribeManagedReleaseSync``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `ManagedReleaseApiDescribeManagedReleaseSync`: ManagedReleaseSync
	fmt.Fprintf(os.Stdout, "Response from `ManagedReleaseApiAPI.ManagedReleaseApiDescribeManagedReleaseSync`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**organizationId** | **string** |  | 
**provisionerTargetId** | **string** |  | 
**bundleRevisionId** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiManagedReleaseApiDescribeManagedReleaseSyncRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------



 **includeArtifactResults** | **bool** |  | [default to true]

### Return type

[**ManagedReleaseSync**](ManagedReleaseSync.md)

### Authorization

[api_key_header_Authorization](../README.md#api_key_header_Authorization)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json, application/vnd.goa.error

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## ManagedReleaseApiListManagedReleaseRevisions

> ListManagedReleaseRevisionsResult2 ManagedReleaseApiListManagedReleaseRevisions(ctx).Status(status).ReleaseComponent(releaseComponent).Limit(limit).IncludeArtifacts(includeArtifacts).Execute()

ListManagedReleaseRevisions managed-release-api

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
	status := "AVAILABLE" // string |  (optional)
	releaseComponent := "bootstrap-service" // string |  (optional)
	limit := int64(61) // int64 |  (optional) (default to 20)
	includeArtifacts := false // bool |  (optional) (default to false)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.ManagedReleaseApiAPI.ManagedReleaseApiListManagedReleaseRevisions(context.Background()).Status(status).ReleaseComponent(releaseComponent).Limit(limit).IncludeArtifacts(includeArtifacts).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `ManagedReleaseApiAPI.ManagedReleaseApiListManagedReleaseRevisions``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `ManagedReleaseApiListManagedReleaseRevisions`: ListManagedReleaseRevisionsResult2
	fmt.Fprintf(os.Stdout, "Response from `ManagedReleaseApiAPI.ManagedReleaseApiListManagedReleaseRevisions`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiManagedReleaseApiListManagedReleaseRevisionsRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **status** | **string** |  | 
 **releaseComponent** | **string** |  | 
 **limit** | **int64** |  | [default to 20]
 **includeArtifacts** | **bool** |  | [default to false]

### Return type

[**ListManagedReleaseRevisionsResult2**](ListManagedReleaseRevisionsResult2.md)

### Authorization

[api_key_header_Authorization](../README.md#api_key_header_Authorization)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json, application/vnd.goa.error

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## ManagedReleaseApiListManagedReleaseSyncs

> ListManagedReleaseSyncsResult ManagedReleaseApiListManagedReleaseSyncs(ctx, organizationId, provisionerTargetId).Status(status).IncludeArtifactResults(includeArtifactResults).Execute()

ListManagedReleaseSyncs managed-release-api

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
	organizationId := "sea" // string | 
	provisionerTargetId := "wy6" // string | 
	status := "AVAILABLE" // string |  (optional)
	includeArtifactResults := true // bool |  (optional) (default to false)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.ManagedReleaseApiAPI.ManagedReleaseApiListManagedReleaseSyncs(context.Background(), organizationId, provisionerTargetId).Status(status).IncludeArtifactResults(includeArtifactResults).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `ManagedReleaseApiAPI.ManagedReleaseApiListManagedReleaseSyncs``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `ManagedReleaseApiListManagedReleaseSyncs`: ListManagedReleaseSyncsResult
	fmt.Fprintf(os.Stdout, "Response from `ManagedReleaseApiAPI.ManagedReleaseApiListManagedReleaseSyncs`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**organizationId** | **string** |  | 
**provisionerTargetId** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiManagedReleaseApiListManagedReleaseSyncsRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


 **status** | **string** |  | 
 **includeArtifactResults** | **bool** |  | [default to false]

### Return type

[**ListManagedReleaseSyncsResult**](ListManagedReleaseSyncsResult.md)

### Authorization

[api_key_header_Authorization](../README.md#api_key_header_Authorization)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json, application/vnd.goa.error

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## ManagedReleaseApiPublishManagedRelease

> PublishManagedReleaseResult ManagedReleaseApiPublishManagedRelease(ctx).PublishManagedReleaseRequest2(publishManagedReleaseRequest2).Execute()

PublishManagedRelease managed-release-api

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
	publishManagedReleaseRequest2 := *openapiclient.NewPublishManagedReleaseRequest2([]openapiclient.ManagedReleaseArtifact2{*openapiclient.NewManagedReleaseArtifact2("container-image/dataplane-agent", []string{"bootstrap"}, "bootstrap-service", "2au", "p8x", "container-image")}, "bootstrap-service", "g4", "v6", "was") // PublishManagedReleaseRequest2 | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.ManagedReleaseApiAPI.ManagedReleaseApiPublishManagedRelease(context.Background()).PublishManagedReleaseRequest2(publishManagedReleaseRequest2).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `ManagedReleaseApiAPI.ManagedReleaseApiPublishManagedRelease``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `ManagedReleaseApiPublishManagedRelease`: PublishManagedReleaseResult
	fmt.Fprintf(os.Stdout, "Response from `ManagedReleaseApiAPI.ManagedReleaseApiPublishManagedRelease`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiManagedReleaseApiPublishManagedReleaseRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **publishManagedReleaseRequest2** | [**PublishManagedReleaseRequest2**](PublishManagedReleaseRequest2.md) |  | 

### Return type

[**PublishManagedReleaseResult**](PublishManagedReleaseResult.md)

### Authorization

[api_key_header_Authorization](../README.md#api_key_header_Authorization)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json, application/vnd.goa.error

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

