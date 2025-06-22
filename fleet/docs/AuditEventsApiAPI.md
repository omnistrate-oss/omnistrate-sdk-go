# \AuditEventsApiAPI

All URIs are relative to *https://api.omnistrate.cloud*

Method | HTTP request | Description
------------- | ------------- | -------------
[**AuditEventsApiAuditEvents**](AuditEventsApiAPI.md#AuditEventsApiAuditEvents) | **Get** /2022-09-01-00/fleet/audit-events | AuditEvents audit-events-api



## AuditEventsApiAuditEvents

> FleetAuditEventsResult AuditEventsApiAuditEvents(ctx).NextPageToken(nextPageToken).PageSize(pageSize).ServiceID(serviceID).EnvironmentType(environmentType).EventSourceTypes(eventSourceTypes).InstanceID(instanceID).ProductTierID(productTierID).StartDate(startDate).EndDate(endDate).Execute()

AuditEvents audit-events-api

### Example

```go
package main

import (
	"context"
	"fmt"
	"os"
    "time"
	openapiclient "github.com/omnistrate-oss/omnistrate-sdk-go/fleet"
)

func main() {
	nextPageToken := "token" // string |  (optional)
	pageSize := int64(10) // int64 |  (optional)
	serviceID := "s-123456" // string | The service ID to list events for (optional)
	environmentType := "PROD|PRIVATE|CANARY|STAGING|QA|DEV" // string |  (optional)
	eventSourceTypes := []string{"Sint est in."} // []string | The event types to filter by (optional)
	instanceID := "instance-12345678" // string | The instance ID to list events for (optional)
	productTierID := "Eum officiis et." // string |  (optional)
	startDate := time.Now() // time.Time | Start date of the events (optional)
	endDate := time.Now() // time.Time | End date of the events (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.AuditEventsApiAPI.AuditEventsApiAuditEvents(context.Background()).NextPageToken(nextPageToken).PageSize(pageSize).ServiceID(serviceID).EnvironmentType(environmentType).EventSourceTypes(eventSourceTypes).InstanceID(instanceID).ProductTierID(productTierID).StartDate(startDate).EndDate(endDate).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `AuditEventsApiAPI.AuditEventsApiAuditEvents``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `AuditEventsApiAuditEvents`: FleetAuditEventsResult
	fmt.Fprintf(os.Stdout, "Response from `AuditEventsApiAPI.AuditEventsApiAuditEvents`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiAuditEventsApiAuditEventsRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **nextPageToken** | **string** |  | 
 **pageSize** | **int64** |  | 
 **serviceID** | **string** | The service ID to list events for | 
 **environmentType** | **string** |  | 
 **eventSourceTypes** | **[]string** | The event types to filter by | 
 **instanceID** | **string** | The instance ID to list events for | 
 **productTierID** | **string** |  | 
 **startDate** | **time.Time** | Start date of the events | 
 **endDate** | **time.Time** | End date of the events | 

### Return type

[**FleetAuditEventsResult**](FleetAuditEventsResult.md)

### Authorization

[api_key_header_Authorization](../README.md#api_key_header_Authorization)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json, application/vnd.goa.error

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

