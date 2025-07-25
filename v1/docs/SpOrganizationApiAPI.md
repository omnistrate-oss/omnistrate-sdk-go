# \SpOrganizationApiAPI

All URIs are relative to *https://api.omnistrate.cloud*

Method | HTTP request | Description
------------- | ------------- | -------------
[**SpOrganizationApiDescribeServiceProviderOrganization**](SpOrganizationApiAPI.md#SpOrganizationApiDescribeServiceProviderOrganization) | **Get** /2022-09-01-00/sp-organization | DescribeServiceProviderOrganization sp-organization-api
[**SpOrganizationApiModifyServiceProviderOrganization**](SpOrganizationApiAPI.md#SpOrganizationApiModifyServiceProviderOrganization) | **Patch** /2022-09-01-00/sp-organization | ModifyServiceProviderOrganization sp-organization-api



## SpOrganizationApiDescribeServiceProviderOrganization

> DescribeServiceProviderOrganizationResult SpOrganizationApiDescribeServiceProviderOrganization(ctx).Execute()

DescribeServiceProviderOrganization sp-organization-api

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

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.SpOrganizationApiAPI.SpOrganizationApiDescribeServiceProviderOrganization(context.Background()).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `SpOrganizationApiAPI.SpOrganizationApiDescribeServiceProviderOrganization``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `SpOrganizationApiDescribeServiceProviderOrganization`: DescribeServiceProviderOrganizationResult
	fmt.Fprintf(os.Stdout, "Response from `SpOrganizationApiAPI.SpOrganizationApiDescribeServiceProviderOrganization`: %v\n", resp)
}
```

### Path Parameters

This endpoint does not need any parameter.

### Other Parameters

Other parameters are passed through a pointer to a apiSpOrganizationApiDescribeServiceProviderOrganizationRequest struct via the builder pattern


### Return type

[**DescribeServiceProviderOrganizationResult**](DescribeServiceProviderOrganizationResult.md)

### Authorization

[api_key_header_Authorization](../README.md#api_key_header_Authorization)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json, application/vnd.goa.error

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## SpOrganizationApiModifyServiceProviderOrganization

> SpOrganizationApiModifyServiceProviderOrganization(ctx).ModifyServiceProviderOrganizationRequest2(modifyServiceProviderOrganizationRequest2).Execute()

ModifyServiceProviderOrganization sp-organization-api

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
	modifyServiceProviderOrganizationRequest2 := *openapiclient.NewModifyServiceProviderOrganizationRequest2() // ModifyServiceProviderOrganizationRequest2 | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	r, err := apiClient.SpOrganizationApiAPI.SpOrganizationApiModifyServiceProviderOrganization(context.Background()).ModifyServiceProviderOrganizationRequest2(modifyServiceProviderOrganizationRequest2).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `SpOrganizationApiAPI.SpOrganizationApiModifyServiceProviderOrganization``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiSpOrganizationApiModifyServiceProviderOrganizationRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **modifyServiceProviderOrganizationRequest2** | [**ModifyServiceProviderOrganizationRequest2**](ModifyServiceProviderOrganizationRequest2.md) |  | 

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

