# \EventsApiAPI

All URIs are relative to *https://api.omnistrate.cloud*

Method | HTTP request | Description
------------- | ------------- | -------------
[**EventsApiAcknowledgeEvent**](EventsApiAPI.md#EventsApiAcknowledgeEvent) | **Delete** /2022-09-01-00/fleet/events/{id} | AcknowledgeEvent events-api
[**EventsApiListEvents**](EventsApiAPI.md#EventsApiListEvents) | **Get** /2022-09-01-00/fleet/events | ListEvents events-api



## EventsApiAcknowledgeEvent

> EventsApiAcknowledgeEvent(ctx, id).Execute()

AcknowledgeEvent events-api

### Example

```go
package main

import (
	"context"
	"fmt"
	"os"
	openapiclient "github.com/omnistrate-oss/omnistrate-sdk-go/fleet"
)

func main() {
	id := "e-123456" // string | The ID of the event

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	r, err := apiClient.EventsApiAPI.EventsApiAcknowledgeEvent(context.Background(), id).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `EventsApiAPI.EventsApiAcknowledgeEvent``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **string** | The ID of the event | 

### Other Parameters

Other parameters are passed through a pointer to a apiEventsApiAcknowledgeEventRequest struct via the builder pattern


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


## EventsApiListEvents

> ListEndCustomerEventsResult EventsApiListEvents(ctx).EnvironmentType(environmentType).Execute()

ListEvents events-api

### Example

```go
package main

import (
	"context"
	"fmt"
	"os"
	openapiclient "github.com/omnistrate-oss/omnistrate-sdk-go/fleet"
)

func main() {
	environmentType := "DEV" // string | The environment type to filter by (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.EventsApiAPI.EventsApiListEvents(context.Background()).EnvironmentType(environmentType).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `EventsApiAPI.EventsApiListEvents``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `EventsApiListEvents`: ListEndCustomerEventsResult
	fmt.Fprintf(os.Stdout, "Response from `EventsApiAPI.EventsApiListEvents`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiEventsApiListEventsRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **environmentType** | **string** | The environment type to filter by | 

### Return type

[**ListEndCustomerEventsResult**](ListEndCustomerEventsResult.md)

### Authorization

[api_key_header_Authorization](../README.md#api_key_header_Authorization)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json, application/vnd.goa.error

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

