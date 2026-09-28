# \ManagedArtifactApiAPI

All URIs are relative to *https://api.omnistrate.cloud*

Method | HTTP request | Description
------------- | ------------- | -------------
[**ManagedArtifactApiDescribeManagedArtifactRelease**](ManagedArtifactApiAPI.md#ManagedArtifactApiDescribeManagedArtifactRelease) | **Get** /2022-09-01-00/managed-artifact/releases/{bundleVersion} | DescribeManagedArtifactRelease managed-artifact-api
[**ManagedArtifactApiDescribeManagedArtifactReleasePolicy**](ManagedArtifactApiAPI.md#ManagedArtifactApiDescribeManagedArtifactReleasePolicy) | **Get** /2022-09-01-00/managed-artifact/release-policy/{environmentType}/{cloudProvider} | DescribeManagedArtifactReleasePolicy managed-artifact-api
[**ManagedArtifactApiDescribeManagedArtifactSync**](ManagedArtifactApiAPI.md#ManagedArtifactApiDescribeManagedArtifactSync) | **Get** /2022-09-01-00/managed-artifact/syncs/{id} | DescribeManagedArtifactSync managed-artifact-api
[**ManagedArtifactApiListManagedArtifactReleases**](ManagedArtifactApiAPI.md#ManagedArtifactApiListManagedArtifactReleases) | **Get** /2022-09-01-00/managed-artifact/releases | ListManagedArtifactReleases managed-artifact-api
[**ManagedArtifactApiListManagedArtifactSyncs**](ManagedArtifactApiAPI.md#ManagedArtifactApiListManagedArtifactSyncs) | **Get** /2022-09-01-00/managed-artifact/syncs | ListManagedArtifactSyncs managed-artifact-api
[**ManagedArtifactApiUpdateManagedArtifactReleasePolicy**](ManagedArtifactApiAPI.md#ManagedArtifactApiUpdateManagedArtifactReleasePolicy) | **Put** /2022-09-01-00/managed-artifact/release-policy/{environmentType}/{cloudProvider} | UpdateManagedArtifactReleasePolicy managed-artifact-api



## ManagedArtifactApiDescribeManagedArtifactRelease

> ManagedArtifactRelease ManagedArtifactApiDescribeManagedArtifactRelease(ctx, bundleVersion).Execute()

DescribeManagedArtifactRelease managed-artifact-api



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
	bundleVersion := "r0000020" // string | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.ManagedArtifactApiAPI.ManagedArtifactApiDescribeManagedArtifactRelease(context.Background(), bundleVersion).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `ManagedArtifactApiAPI.ManagedArtifactApiDescribeManagedArtifactRelease``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `ManagedArtifactApiDescribeManagedArtifactRelease`: ManagedArtifactRelease
	fmt.Fprintf(os.Stdout, "Response from `ManagedArtifactApiAPI.ManagedArtifactApiDescribeManagedArtifactRelease`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**bundleVersion** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiManagedArtifactApiDescribeManagedArtifactReleaseRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

[**ManagedArtifactRelease**](ManagedArtifactRelease.md)

### Authorization

[api_key_header_Authorization](../README.md#api_key_header_Authorization)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json, application/vnd.goa.error

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## ManagedArtifactApiDescribeManagedArtifactReleasePolicy

> ManagedArtifactReleasePolicy ManagedArtifactApiDescribeManagedArtifactReleasePolicy(ctx, environmentType, cloudProvider).Execute()

DescribeManagedArtifactReleasePolicy managed-artifact-api



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
	environmentType := "PROD|PRIVATE|CANARY|STAGING|QA|DEV|GLOBAL" // string | 
	cloudProvider := "aws|azure|gcp|nebius|oci|byoc-onprem|all" // string | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.ManagedArtifactApiAPI.ManagedArtifactApiDescribeManagedArtifactReleasePolicy(context.Background(), environmentType, cloudProvider).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `ManagedArtifactApiAPI.ManagedArtifactApiDescribeManagedArtifactReleasePolicy``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `ManagedArtifactApiDescribeManagedArtifactReleasePolicy`: ManagedArtifactReleasePolicy
	fmt.Fprintf(os.Stdout, "Response from `ManagedArtifactApiAPI.ManagedArtifactApiDescribeManagedArtifactReleasePolicy`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**environmentType** | **string** |  | 
**cloudProvider** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiManagedArtifactApiDescribeManagedArtifactReleasePolicyRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------



### Return type

[**ManagedArtifactReleasePolicy**](ManagedArtifactReleasePolicy.md)

### Authorization

[api_key_header_Authorization](../README.md#api_key_header_Authorization)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json, application/vnd.goa.error

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## ManagedArtifactApiDescribeManagedArtifactSync

> ManagedArtifactSync ManagedArtifactApiDescribeManagedArtifactSync(ctx, id).Execute()

DescribeManagedArtifactSync managed-artifact-api



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
	id := "spabs-M" // string | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.ManagedArtifactApiAPI.ManagedArtifactApiDescribeManagedArtifactSync(context.Background(), id).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `ManagedArtifactApiAPI.ManagedArtifactApiDescribeManagedArtifactSync``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `ManagedArtifactApiDescribeManagedArtifactSync`: ManagedArtifactSync
	fmt.Fprintf(os.Stdout, "Response from `ManagedArtifactApiAPI.ManagedArtifactApiDescribeManagedArtifactSync`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiManagedArtifactApiDescribeManagedArtifactSyncRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

[**ManagedArtifactSync**](ManagedArtifactSync.md)

### Authorization

[api_key_header_Authorization](../README.md#api_key_header_Authorization)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json, application/vnd.goa.error

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## ManagedArtifactApiListManagedArtifactReleases

> ListManagedArtifactReleasesResult ManagedArtifactApiListManagedArtifactReleases(ctx).BundleVersion(bundleVersion).ReleasedAfter(releasedAfter).ReleasedBefore(releasedBefore).Limit(limit).NextPageToken(nextPageToken).Execute()

ListManagedArtifactReleases managed-artifact-api



### Example

```go
package main

import (
	"context"
	"fmt"
	"os"
    "time"
	openapiclient "github.com/omnistrate-oss/omnistrate-sdk-go/v1"
)

func main() {
	bundleVersion := "r12258919" // string |  (optional)
	releasedAfter := time.Now() // time.Time |  (optional)
	releasedBefore := time.Now() // time.Time |  (optional)
	limit := int64(65) // int64 |  (optional) (default to 20)
	nextPageToken := "qa5" // string | Opaque token returned by the previous list response. (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.ManagedArtifactApiAPI.ManagedArtifactApiListManagedArtifactReleases(context.Background()).BundleVersion(bundleVersion).ReleasedAfter(releasedAfter).ReleasedBefore(releasedBefore).Limit(limit).NextPageToken(nextPageToken).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `ManagedArtifactApiAPI.ManagedArtifactApiListManagedArtifactReleases``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `ManagedArtifactApiListManagedArtifactReleases`: ListManagedArtifactReleasesResult
	fmt.Fprintf(os.Stdout, "Response from `ManagedArtifactApiAPI.ManagedArtifactApiListManagedArtifactReleases`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiManagedArtifactApiListManagedArtifactReleasesRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **bundleVersion** | **string** |  | 
 **releasedAfter** | **time.Time** |  | 
 **releasedBefore** | **time.Time** |  | 
 **limit** | **int64** |  | [default to 20]
 **nextPageToken** | **string** | Opaque token returned by the previous list response. | 

### Return type

[**ListManagedArtifactReleasesResult**](ListManagedArtifactReleasesResult.md)

### Authorization

[api_key_header_Authorization](../README.md#api_key_header_Authorization)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json, application/vnd.goa.error

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## ManagedArtifactApiListManagedArtifactSyncs

> ListManagedArtifactSyncsResult ManagedArtifactApiListManagedArtifactSyncs(ctx).BundleVersion(bundleVersion).Status(status).TargetId(targetId).UpdatedAfter(updatedAfter).UpdatedBefore(updatedBefore).Limit(limit).NextPageToken(nextPageToken).Execute()

ListManagedArtifactSyncs managed-artifact-api



### Example

```go
package main

import (
	"context"
	"fmt"
	"os"
    "time"
	openapiclient "github.com/omnistrate-oss/omnistrate-sdk-go/v1"
)

func main() {
	bundleVersion := "r1467398" // string |  (optional)
	status := "READY" // string |  (optional)
	targetId := "q" // string |  (optional)
	updatedAfter := time.Now() // time.Time |  (optional)
	updatedBefore := time.Now() // time.Time |  (optional)
	limit := int64(56) // int64 |  (optional) (default to 20)
	nextPageToken := "blt" // string | Opaque token returned by the previous list response. (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.ManagedArtifactApiAPI.ManagedArtifactApiListManagedArtifactSyncs(context.Background()).BundleVersion(bundleVersion).Status(status).TargetId(targetId).UpdatedAfter(updatedAfter).UpdatedBefore(updatedBefore).Limit(limit).NextPageToken(nextPageToken).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `ManagedArtifactApiAPI.ManagedArtifactApiListManagedArtifactSyncs``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `ManagedArtifactApiListManagedArtifactSyncs`: ListManagedArtifactSyncsResult
	fmt.Fprintf(os.Stdout, "Response from `ManagedArtifactApiAPI.ManagedArtifactApiListManagedArtifactSyncs`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiManagedArtifactApiListManagedArtifactSyncsRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **bundleVersion** | **string** |  | 
 **status** | **string** |  | 
 **targetId** | **string** |  | 
 **updatedAfter** | **time.Time** |  | 
 **updatedBefore** | **time.Time** |  | 
 **limit** | **int64** |  | [default to 20]
 **nextPageToken** | **string** | Opaque token returned by the previous list response. | 

### Return type

[**ListManagedArtifactSyncsResult**](ListManagedArtifactSyncsResult.md)

### Authorization

[api_key_header_Authorization](../README.md#api_key_header_Authorization)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json, application/vnd.goa.error

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## ManagedArtifactApiUpdateManagedArtifactReleasePolicy

> ManagedArtifactReleasePolicy ManagedArtifactApiUpdateManagedArtifactReleasePolicy(ctx, environmentType, cloudProvider).UpdateManagedArtifactReleasePolicyRequest2(updateManagedArtifactReleasePolicyRequest2).Execute()

UpdateManagedArtifactReleasePolicy managed-artifact-api



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
	environmentType := "PROD|PRIVATE|CANARY|STAGING|QA|DEV|GLOBAL" // string | 
	cloudProvider := "aws|azure|gcp|nebius|oci|byoc-onprem|all" // string | 
	updateManagedArtifactReleasePolicyRequest2 := *openapiclient.NewUpdateManagedArtifactReleasePolicyRequest2(false) // UpdateManagedArtifactReleasePolicyRequest2 | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.ManagedArtifactApiAPI.ManagedArtifactApiUpdateManagedArtifactReleasePolicy(context.Background(), environmentType, cloudProvider).UpdateManagedArtifactReleasePolicyRequest2(updateManagedArtifactReleasePolicyRequest2).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `ManagedArtifactApiAPI.ManagedArtifactApiUpdateManagedArtifactReleasePolicy``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `ManagedArtifactApiUpdateManagedArtifactReleasePolicy`: ManagedArtifactReleasePolicy
	fmt.Fprintf(os.Stdout, "Response from `ManagedArtifactApiAPI.ManagedArtifactApiUpdateManagedArtifactReleasePolicy`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**environmentType** | **string** |  | 
**cloudProvider** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiManagedArtifactApiUpdateManagedArtifactReleasePolicyRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


 **updateManagedArtifactReleasePolicyRequest2** | [**UpdateManagedArtifactReleasePolicyRequest2**](UpdateManagedArtifactReleasePolicyRequest2.md) |  | 

### Return type

[**ManagedArtifactReleasePolicy**](ManagedArtifactReleasePolicy.md)

### Authorization

[api_key_header_Authorization](../README.md#api_key_header_Authorization)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json, application/vnd.goa.error

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

