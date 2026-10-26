# \AuthSessionApiAPI

All URIs are relative to *https://api.omnistrate.cloud*

Method | HTTP request | Description
------------- | ------------- | -------------
[**AuthSessionApiDescribeAuthSession**](AuthSessionApiAPI.md#AuthSessionApiDescribeAuthSession) | **Get** /2022-09-01-00/auth/session | DescribeAuthSession auth-session-api



## AuthSessionApiDescribeAuthSession

> AuthSessionResult AuthSessionApiDescribeAuthSession(ctx).Execute()

DescribeAuthSession auth-session-api



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
	resp, r, err := apiClient.AuthSessionApiAPI.AuthSessionApiDescribeAuthSession(context.Background()).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `AuthSessionApiAPI.AuthSessionApiDescribeAuthSession``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `AuthSessionApiDescribeAuthSession`: AuthSessionResult
	fmt.Fprintf(os.Stdout, "Response from `AuthSessionApiAPI.AuthSessionApiDescribeAuthSession`: %v\n", resp)
}
```

### Path Parameters

This endpoint does not need any parameter.

### Other Parameters

Other parameters are passed through a pointer to a apiAuthSessionApiDescribeAuthSessionRequest struct via the builder pattern


### Return type

[**AuthSessionResult**](AuthSessionResult.md)

### Authorization

[api_key_header_Authorization](../README.md#api_key_header_Authorization)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json, application/vnd.goa.error

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

