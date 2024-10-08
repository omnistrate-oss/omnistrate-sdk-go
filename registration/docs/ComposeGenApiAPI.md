# \ComposeGenApiAPI

All URIs are relative to *https://api.omnistrate.cloud*

Method | HTTP request | Description
------------- | ------------- | -------------
[**ComposeGenApiCheckIfContainerImageAccessible**](ComposeGenApiAPI.md#ComposeGenApiCheckIfContainerImageAccessible) | **Get** /2022-09-01-00/compose-gen/image | CheckIfContainerImageAccessible compose-gen-api
[**ComposeGenApiGenerateComposeSpecFromContainerImage**](ComposeGenApiAPI.md#ComposeGenApiGenerateComposeSpecFromContainerImage) | **Post** /2022-09-01-00/compose-gen/image | GenerateComposeSpecFromContainerImage compose-gen-api



## ComposeGenApiCheckIfContainerImageAccessible

> CheckIfContainerImageAccessibleResult ComposeGenApiCheckIfContainerImageAccessible(ctx).ImageRegistry(imageRegistry).Image(image).Username(username).Password(password).Execute()

CheckIfContainerImageAccessible compose-gen-api

### Example

```go
package main

import (
	"context"
	"fmt"
	"os"
	openapiclient "github.com/omnistrate/omnistrate-sdk-go/registration"
)

func main() {
	imageRegistry := "docker.io" // string | Registry where the image is stored
	image := "library/mysql:5.7" // string | Name of the image along with the tag. Include the repository name if the image is not from the official repository
	username := "username" // string | Username to access the image registry (optional)
	password := "password" // string | Password to access the image registry (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.ComposeGenApiAPI.ComposeGenApiCheckIfContainerImageAccessible(context.Background()).ImageRegistry(imageRegistry).Image(image).Username(username).Password(password).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `ComposeGenApiAPI.ComposeGenApiCheckIfContainerImageAccessible``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `ComposeGenApiCheckIfContainerImageAccessible`: CheckIfContainerImageAccessibleResult
	fmt.Fprintf(os.Stdout, "Response from `ComposeGenApiAPI.ComposeGenApiCheckIfContainerImageAccessible`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiComposeGenApiCheckIfContainerImageAccessibleRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **imageRegistry** | **string** | Registry where the image is stored | 
 **image** | **string** | Name of the image along with the tag. Include the repository name if the image is not from the official repository | 
 **username** | **string** | Username to access the image registry | 
 **password** | **string** | Password to access the image registry | 

### Return type

[**CheckIfContainerImageAccessibleResult**](CheckIfContainerImageAccessibleResult.md)

### Authorization

[api_key_header_Authorization](../README.md#api_key_header_Authorization)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json, application/vnd.goa.error

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## ComposeGenApiGenerateComposeSpecFromContainerImage

> GenerateComposeSpecFromContainerImageResult ComposeGenApiGenerateComposeSpecFromContainerImage(ctx).GenerateComposeSpecFromContainerImageRequestBody(generateComposeSpecFromContainerImageRequestBody).Execute()

GenerateComposeSpecFromContainerImage compose-gen-api

### Example

```go
package main

import (
	"context"
	"fmt"
	"os"
	openapiclient "github.com/omnistrate/omnistrate-sdk-go/registration"
)

func main() {
	generateComposeSpecFromContainerImageRequestBody := *openapiclient.NewGenerateComposeSpecFromContainerImageRequestBody("library/mysql:5.7", "docker.io") // GenerateComposeSpecFromContainerImageRequestBody | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.ComposeGenApiAPI.ComposeGenApiGenerateComposeSpecFromContainerImage(context.Background()).GenerateComposeSpecFromContainerImageRequestBody(generateComposeSpecFromContainerImageRequestBody).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `ComposeGenApiAPI.ComposeGenApiGenerateComposeSpecFromContainerImage``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `ComposeGenApiGenerateComposeSpecFromContainerImage`: GenerateComposeSpecFromContainerImageResult
	fmt.Fprintf(os.Stdout, "Response from `ComposeGenApiAPI.ComposeGenApiGenerateComposeSpecFromContainerImage`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiComposeGenApiGenerateComposeSpecFromContainerImageRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **generateComposeSpecFromContainerImageRequestBody** | [**GenerateComposeSpecFromContainerImageRequestBody**](GenerateComposeSpecFromContainerImageRequestBody.md) |  | 

### Return type

[**GenerateComposeSpecFromContainerImageResult**](GenerateComposeSpecFromContainerImageResult.md)

### Authorization

[api_key_header_Authorization](../README.md#api_key_header_Authorization)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json, application/vnd.goa.error

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

