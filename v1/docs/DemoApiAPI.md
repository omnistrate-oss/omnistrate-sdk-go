# \DemoApiAPI

All URIs are relative to *https://api.omnistrate.cloud*

Method | HTTP request | Description
------------- | ------------- | -------------
[**DemoApiDemo**](DemoApiAPI.md#DemoApiDemo) | **Post** /2022-09-01-00/demo | Demo demo-api



## DemoApiDemo

> DemoApiDemo(ctx).DemoRequestBody(demoRequestBody).Execute()

Demo demo-api

### Example

```go
package main

import (
	"context"
	"fmt"
	"os"
	openapiclient "github.com/omnistrate/omnistrate-sdk-go/v1"
)

func main() {
	demoRequestBody := *openapiclient.NewDemoRequestBody("ABC", "abc@gmail.com", "John Doe") // DemoRequestBody | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	r, err := apiClient.DemoApiAPI.DemoApiDemo(context.Background()).DemoRequestBody(demoRequestBody).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `DemoApiAPI.DemoApiDemo``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiDemoApiDemoRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **demoRequestBody** | [**DemoRequestBody**](DemoRequestBody.md) |  | 

### Return type

 (empty response body)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/vnd.goa.error

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

