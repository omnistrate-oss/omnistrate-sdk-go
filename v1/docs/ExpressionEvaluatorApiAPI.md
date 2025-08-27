# \ExpressionEvaluatorApiAPI

All URIs are relative to *https://api.omnistrate.cloud*

Method | HTTP request | Description
------------- | ------------- | -------------
[**ExpressionEvaluatorApiExpressionEvaluator**](ExpressionEvaluatorApiAPI.md#ExpressionEvaluatorApiExpressionEvaluator) | **Post** /2022-09-01-00/resource-instance/expression-evaluator | ExpressionEvaluator expression-evaluator-api



## ExpressionEvaluatorApiExpressionEvaluator

> ExpressionEvaluatorResult ExpressionEvaluatorApiExpressionEvaluator(ctx).ExpressionEvaluatorRequest2(expressionEvaluatorRequest2).Execute()

ExpressionEvaluator expression-evaluator-api

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
	expressionEvaluatorRequest2 := *openapiclient.NewExpressionEvaluatorRequest2("mysql", "s-12345678") // ExpressionEvaluatorRequest2 | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.ExpressionEvaluatorApiAPI.ExpressionEvaluatorApiExpressionEvaluator(context.Background()).ExpressionEvaluatorRequest2(expressionEvaluatorRequest2).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `ExpressionEvaluatorApiAPI.ExpressionEvaluatorApiExpressionEvaluator``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `ExpressionEvaluatorApiExpressionEvaluator`: ExpressionEvaluatorResult
	fmt.Fprintf(os.Stdout, "Response from `ExpressionEvaluatorApiAPI.ExpressionEvaluatorApiExpressionEvaluator`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiExpressionEvaluatorApiExpressionEvaluatorRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **expressionEvaluatorRequest2** | [**ExpressionEvaluatorRequest2**](ExpressionEvaluatorRequest2.md) |  | 

### Return type

[**ExpressionEvaluatorResult**](ExpressionEvaluatorResult.md)

### Authorization

[api_key_header_Authorization](../README.md#api_key_header_Authorization)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json, application/vnd.goa.error

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

