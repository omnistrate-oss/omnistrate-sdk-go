# \PlanApiAPI

All URIs are relative to *https://api.omnistrate.cloud*

Method | HTTP request | Description
------------- | ------------- | -------------
[**PlanApiChangePlan**](PlanApiAPI.md#PlanApiChangePlan) | **Post** /2022-09-01-00/plan | ChangePlan plan-api
[**PlanApiDescribePlan**](PlanApiAPI.md#PlanApiDescribePlan) | **Get** /2022-09-01-00/plan | DescribePlan plan-api



## PlanApiChangePlan

> PlanApiChangePlan(ctx).ChangePlanRequest2(changePlanRequest2).Execute()

ChangePlan plan-api

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
	changePlanRequest2 := *openapiclient.NewChangePlanRequest2("STARTER_NO_COMMIT|STARTER|GROWTH|ENTERPRISE") // ChangePlanRequest2 | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	r, err := apiClient.PlanApiAPI.PlanApiChangePlan(context.Background()).ChangePlanRequest2(changePlanRequest2).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `PlanApiAPI.PlanApiChangePlan``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiPlanApiChangePlanRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **changePlanRequest2** | [**ChangePlanRequest2**](ChangePlanRequest2.md) |  | 

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


## PlanApiDescribePlan

> DescribePlanResult PlanApiDescribePlan(ctx).Execute()

DescribePlan plan-api

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
	resp, r, err := apiClient.PlanApiAPI.PlanApiDescribePlan(context.Background()).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `PlanApiAPI.PlanApiDescribePlan``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `PlanApiDescribePlan`: DescribePlanResult
	fmt.Fprintf(os.Stdout, "Response from `PlanApiAPI.PlanApiDescribePlan`: %v\n", resp)
}
```

### Path Parameters

This endpoint does not need any parameter.

### Other Parameters

Other parameters are passed through a pointer to a apiPlanApiDescribePlanRequest struct via the builder pattern


### Return type

[**DescribePlanResult**](DescribePlanResult.md)

### Authorization

[api_key_header_Authorization](../README.md#api_key_header_Authorization)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json, application/vnd.goa.error

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

