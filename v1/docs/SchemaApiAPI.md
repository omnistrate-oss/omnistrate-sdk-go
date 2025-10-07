# \SchemaApiAPI

All URIs are relative to *https://api.omnistrate.cloud*

Method | HTTP request | Description
------------- | ------------- | -------------
[**SchemaApiGetJSONSchema**](SchemaApiAPI.md#SchemaApiGetJSONSchema) | **Get** /2022-09-01-00/json-schema | GetJSONSchema schema-api



## SchemaApiGetJSONSchema

> interface{} SchemaApiGetJSONSchema(ctx).Type_(type_).Execute()

GetJSONSchema schema-api

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
	type_ := "x-omnistrate-service-plan|x-omnistrate-integrations|x-omnistrate-mode-internal|x-omnistrate-proxy-type|x-omnistrate-actionhooks|x-omnistrate-api-params|x-omnistrate-capabilities|x-omnistrate-compute|x-omnistrate-job-config|x-omnistrate-storage|x-omnistrate-image-registry-attributes|x-omnistrate-load-balancer|compose|service-plan|deployment-cell-amenities|system-parameters|services-orchestration-create-dsl|service-orchestration-modify-dsl" // string | The type of the schema to retrieve

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.SchemaApiAPI.SchemaApiGetJSONSchema(context.Background()).Type_(type_).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `SchemaApiAPI.SchemaApiGetJSONSchema``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `SchemaApiGetJSONSchema`: interface{}
	fmt.Fprintf(os.Stdout, "Response from `SchemaApiAPI.SchemaApiGetJSONSchema`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiSchemaApiGetJSONSchemaRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **type_** | **string** | The type of the schema to retrieve | 

### Return type

**interface{}**

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json, application/vnd.goa.error

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

