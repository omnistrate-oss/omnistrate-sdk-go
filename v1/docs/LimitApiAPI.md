# \LimitApiAPI

All URIs are relative to *https://api.omnistrate.cloud*

Method | HTTP request | Description
------------- | ------------- | -------------
[**LimitApiDeleteLimit**](LimitApiAPI.md#LimitApiDeleteLimit) | **Delete** /2022-09-01-00/limit/family/{family}/key/{key} | DeleteLimit limit-api
[**LimitApiDescribeLimit**](LimitApiAPI.md#LimitApiDescribeLimit) | **Get** /2022-09-01-00/limit/family/{family}/key/{key} | DescribeLimit limit-api
[**LimitApiListLimit**](LimitApiAPI.md#LimitApiListLimit) | **Get** /2022-09-01-00/limit | ListLimit limit-api
[**LimitApiUpdateLimit**](LimitApiAPI.md#LimitApiUpdateLimit) | **Patch** /2022-09-01-00/limit/family/{family}/key/{key} | UpdateLimit limit-api



## LimitApiDeleteLimit

> LimitApiDeleteLimit(ctx, family, key).Execute()

DeleteLimit limit-api

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
	family := "COMPUTE_INFRA" // string | Limit family
	key := "MAX_VIRTUAL_CORES_PER_RESOURCE" // string | Unique key to identify the limit

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	r, err := apiClient.LimitApiAPI.LimitApiDeleteLimit(context.Background(), family, key).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `LimitApiAPI.LimitApiDeleteLimit``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**family** | **string** | Limit family | 
**key** | **string** | Unique key to identify the limit | 

### Other Parameters

Other parameters are passed through a pointer to a apiLimitApiDeleteLimitRequest struct via the builder pattern


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


## LimitApiDescribeLimit

> DescribeLimitResult LimitApiDescribeLimit(ctx, family, key).Execute()

DescribeLimit limit-api

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
	family := "COMPUTE_INFRA" // string | Limit family
	key := "MAX_VIRTUAL_CORES_PER_RESOURCE" // string | Unique key to identify the limit

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.LimitApiAPI.LimitApiDescribeLimit(context.Background(), family, key).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `LimitApiAPI.LimitApiDescribeLimit``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `LimitApiDescribeLimit`: DescribeLimitResult
	fmt.Fprintf(os.Stdout, "Response from `LimitApiAPI.LimitApiDescribeLimit`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**family** | **string** | Limit family | 
**key** | **string** | Unique key to identify the limit | 

### Other Parameters

Other parameters are passed through a pointer to a apiLimitApiDescribeLimitRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------



### Return type

[**DescribeLimitResult**](DescribeLimitResult.md)

### Authorization

[api_key_header_Authorization](../README.md#api_key_header_Authorization)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json, application/vnd.goa.error

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## LimitApiListLimit

> ListLimitResult LimitApiListLimit(ctx).ListLimitRequest2(listLimitRequest2).Execute()

ListLimit limit-api

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
	listLimitRequest2 := *openapiclient.NewListLimitRequest2("COMPUTE_INFRA") // ListLimitRequest2 | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.LimitApiAPI.LimitApiListLimit(context.Background()).ListLimitRequest2(listLimitRequest2).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `LimitApiAPI.LimitApiListLimit``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `LimitApiListLimit`: ListLimitResult
	fmt.Fprintf(os.Stdout, "Response from `LimitApiAPI.LimitApiListLimit`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiLimitApiListLimitRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **listLimitRequest2** | [**ListLimitRequest2**](ListLimitRequest2.md) |  | 

### Return type

[**ListLimitResult**](ListLimitResult.md)

### Authorization

[api_key_header_Authorization](../README.md#api_key_header_Authorization)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json, application/vnd.goa.error

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## LimitApiUpdateLimit

> LimitApiUpdateLimit(ctx, family, key).UpdateLimitRequest2(updateLimitRequest2).Execute()

UpdateLimit limit-api

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
	family := "COMPUTE_INFRA" // string | Limit family
	key := "MAX_VIRTUAL_CORES_PER_RESOURCE" // string | Unique key to identify the limit
	updateLimitRequest2 := *openapiclient.NewUpdateLimitRequest2(int64(4)) // UpdateLimitRequest2 | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	r, err := apiClient.LimitApiAPI.LimitApiUpdateLimit(context.Background(), family, key).UpdateLimitRequest2(updateLimitRequest2).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `LimitApiAPI.LimitApiUpdateLimit``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**family** | **string** | Limit family | 
**key** | **string** | Unique key to identify the limit | 

### Other Parameters

Other parameters are passed through a pointer to a apiLimitApiUpdateLimitRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


 **updateLimitRequest2** | [**UpdateLimitRequest2**](UpdateLimitRequest2.md) |  | 

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

