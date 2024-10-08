# \ContactusApiAPI

All URIs are relative to *https://api.omnistrate.cloud*

Method | HTTP request | Description
------------- | ------------- | -------------
[**ContactusApiContactus**](ContactusApiAPI.md#ContactusApiContactus) | **Post** /2022-09-01-00/contactus | Contactus contactus-api



## ContactusApiContactus

> ContactusApiContactus(ctx).ContactusRequestBody(contactusRequestBody).Execute()

Contactus contactus-api

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
	contactusRequestBody := *openapiclient.NewContactusRequestBody("ABC", "abc@gmail.com", "this is a test", "John Doe") // ContactusRequestBody | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	r, err := apiClient.ContactusApiAPI.ContactusApiContactus(context.Background()).ContactusRequestBody(contactusRequestBody).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `ContactusApiAPI.ContactusApiContactus``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiContactusApiContactusRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **contactusRequestBody** | [**ContactusRequestBody**](ContactusRequestBody.md) |  | 

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

