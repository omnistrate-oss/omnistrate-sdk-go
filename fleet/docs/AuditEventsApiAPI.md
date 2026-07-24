# \AuditEventsApiAPI

All URIs are relative to *https://api.omnistrate.cloud*

Method | HTTP request | Description
------------- | ------------- | -------------
[**AuditEventsApiAuditEvents**](AuditEventsApiAPI.md#AuditEventsApiAuditEvents) | **Get** /2022-09-01-00/fleet/audit-events | AuditEvents audit-events-api



## AuditEventsApiAuditEvents

> FleetAuditEventsResult AuditEventsApiAuditEvents(ctx).NextPageToken(nextPageToken).PageSize(pageSize).ServiceID(serviceID).EnvironmentType(environmentType).EventSourceTypes(eventSourceTypes).InstanceID(instanceID).ProductTierID(productTierID).StartDate(startDate).EndDate(endDate).ExcludeWorkflowFailures(excludeWorkflowFailures).UserId(userId).OrgId(orgId).SubscriptionId(subscriptionId).IpAddress(ipAddress).BillingProvider(billingProvider).ExternalPayerId(externalPayerId).Execute()

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
	environmentType := "PROD|PRIVATE|CANARY|STAGING|QA|DEV|GLOBAL" // string |  (optional)
	eventSourceTypes := []string{"Eos qui molestias."} // []string | The event types to filter by (optional)
	instanceID := "instance-12345678" // string | The instance ID to list events for (optional)
	productTierID := "Est corporis quia." // string |  (optional)
	startDate := time.Now() // time.Time | Start date of the events (optional)
	endDate := time.Now() // time.Time | End date of the events (optional)
	excludeWorkflowFailures := true // bool | Whether to exclude workflow failure details from the response. (optional)
	userId := "Fuga et natus." // string | The user ID to filter events by (optional)
	orgId := "Sunt sed quis." // string | The organization ID to filter events by (optional)
	subscriptionId := "sub-abcd1234" // string | The subscription ID to filter events by (optional)
	ipAddress := "203.0.113.42" // string | The IP address to filter events by (exact match) (optional)
	billingProvider := "STRIPE" // string | The billing provider on the instance's subscription. Empty when no subscription is linked. (optional)
	externalPayerId := "cus_NXyz12345abcdef" // string | The raw external payer ID on the instance's subscription. (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.AuditEventsApiAPI.AuditEventsApiAuditEvents(context.Background()).NextPageToken(nextPageToken).PageSize(pageSize).ServiceID(serviceID).EnvironmentType(environmentType).EventSourceTypes(eventSourceTypes).InstanceID(instanceID).ProductTierID(productTierID).StartDate(startDate).EndDate(endDate).ExcludeWorkflowFailures(excludeWorkflowFailures).UserId(userId).OrgId(orgId).SubscriptionId(subscriptionId).IpAddress(ipAddress).BillingProvider(billingProvider).ExternalPayerId(externalPayerId).Execute()
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
 **excludeWorkflowFailures** | **bool** | Whether to exclude workflow failure details from the response. | 
 **userId** | **string** | The user ID to filter events by | 
 **orgId** | **string** | The organization ID to filter events by | 
 **subscriptionId** | **string** | The subscription ID to filter events by | 
 **ipAddress** | **string** | The IP address to filter events by (exact match) | 
 **billingProvider** | **string** | The billing provider on the instance&#39;s subscription. Empty when no subscription is linked. | 
 **externalPayerId** | **string** | The raw external payer ID on the instance&#39;s subscription. | 

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

