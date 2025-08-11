# \NotificationsApiAPI

All URIs are relative to *https://api.omnistrate.cloud*

Method | HTTP request | Description
------------- | ------------- | -------------
[**NotificationsApiCreateNotificationChannel**](NotificationsApiAPI.md#NotificationsApiCreateNotificationChannel) | **Post** /2022-09-01-00/fleet/notification-channel | CreateNotificationChannel notifications-api
[**NotificationsApiDeleteNotificationChannel**](NotificationsApiAPI.md#NotificationsApiDeleteNotificationChannel) | **Delete** /2022-09-01-00/fleet/notification-channel/{id} | DeleteNotificationChannel notifications-api
[**NotificationsApiDescribeNotificationChannel**](NotificationsApiAPI.md#NotificationsApiDescribeNotificationChannel) | **Get** /2022-09-01-00/fleet/notification-channel/{id} | DescribeNotificationChannel notifications-api
[**NotificationsApiListNotificationChannels**](NotificationsApiAPI.md#NotificationsApiListNotificationChannels) | **Get** /2022-09-01-00/fleet/notification-channel | ListNotificationChannels notifications-api
[**NotificationsApiNotificationChannelEventHistory**](NotificationsApiAPI.md#NotificationsApiNotificationChannelEventHistory) | **Get** /2022-09-01-00/fleet/notification-channel/{id}/event-history | NotificationChannelEventHistory notifications-api
[**NotificationsApiReplayEvent**](NotificationsApiAPI.md#NotificationsApiReplayEvent) | **Post** /2022-09-01-00/fleet/notification-event/{id}/replay | ReplayEvent notifications-api
[**NotificationsApiUpdateNotificationChannel**](NotificationsApiAPI.md#NotificationsApiUpdateNotificationChannel) | **Patch** /2022-09-01-00/fleet/notification-channel/{id} | UpdateNotificationChannel notifications-api



## NotificationsApiCreateNotificationChannel

> string NotificationsApiCreateNotificationChannel(ctx).CreateNotificationChannelRequest2(createNotificationChannelRequest2).Execute()

CreateNotificationChannel notifications-api

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
	createNotificationChannelRequest2 := *openapiclient.NewCreateNotificationChannelRequest2("My Channel", *openapiclient.NewChannelSubscription([]string{"Alarm|Notification"}, []string{"PROD|PRIVATE|CANARY|STAGING|QA|DEV|GLOBAL"}, []string{"InstanceEvent|ServiceEvent|UserEvent|IdentityProviderEvent"}, []string{"Critical|High|Medium|Low"}, []string{"UnhealthyInstance|FailedDeployment|ScaleOut|UserSignUp|UserSubscription"})) // CreateNotificationChannelRequest2 | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.NotificationsApiAPI.NotificationsApiCreateNotificationChannel(context.Background()).CreateNotificationChannelRequest2(createNotificationChannelRequest2).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `NotificationsApiAPI.NotificationsApiCreateNotificationChannel``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `NotificationsApiCreateNotificationChannel`: string
	fmt.Fprintf(os.Stdout, "Response from `NotificationsApiAPI.NotificationsApiCreateNotificationChannel`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiNotificationsApiCreateNotificationChannelRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **createNotificationChannelRequest2** | [**CreateNotificationChannelRequest2**](CreateNotificationChannelRequest2.md) |  | 

### Return type

**string**

### Authorization

[api_key_header_Authorization](../README.md#api_key_header_Authorization)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json, application/vnd.goa.error

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## NotificationsApiDeleteNotificationChannel

> NotificationsApiDeleteNotificationChannel(ctx, id).Execute()

DeleteNotificationChannel notifications-api

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
	id := "channel-12345678" // string | Unique identifier of the channel

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	r, err := apiClient.NotificationsApiAPI.NotificationsApiDeleteNotificationChannel(context.Background(), id).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `NotificationsApiAPI.NotificationsApiDeleteNotificationChannel``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **string** | Unique identifier of the channel | 

### Other Parameters

Other parameters are passed through a pointer to a apiNotificationsApiDeleteNotificationChannelRequest struct via the builder pattern


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


## NotificationsApiDescribeNotificationChannel

> Channel NotificationsApiDescribeNotificationChannel(ctx, id).Execute()

DescribeNotificationChannel notifications-api

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
	id := "channel-12345678" // string | Unique identifier of the channel

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.NotificationsApiAPI.NotificationsApiDescribeNotificationChannel(context.Background(), id).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `NotificationsApiAPI.NotificationsApiDescribeNotificationChannel``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `NotificationsApiDescribeNotificationChannel`: Channel
	fmt.Fprintf(os.Stdout, "Response from `NotificationsApiAPI.NotificationsApiDescribeNotificationChannel`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **string** | Unique identifier of the channel | 

### Other Parameters

Other parameters are passed through a pointer to a apiNotificationsApiDescribeNotificationChannelRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

[**Channel**](Channel.md)

### Authorization

[api_key_header_Authorization](../README.md#api_key_header_Authorization)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json, application/vnd.goa.error

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## NotificationsApiListNotificationChannels

> ListNotificationChannelsResult NotificationsApiListNotificationChannels(ctx).Execute()

ListNotificationChannels notifications-api

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

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.NotificationsApiAPI.NotificationsApiListNotificationChannels(context.Background()).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `NotificationsApiAPI.NotificationsApiListNotificationChannels``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `NotificationsApiListNotificationChannels`: ListNotificationChannelsResult
	fmt.Fprintf(os.Stdout, "Response from `NotificationsApiAPI.NotificationsApiListNotificationChannels`: %v\n", resp)
}
```

### Path Parameters

This endpoint does not need any parameter.

### Other Parameters

Other parameters are passed through a pointer to a apiNotificationsApiListNotificationChannelsRequest struct via the builder pattern


### Return type

[**ListNotificationChannelsResult**](ListNotificationChannelsResult.md)

### Authorization

[api_key_header_Authorization](../README.md#api_key_header_Authorization)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json, application/vnd.goa.error

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## NotificationsApiNotificationChannelEventHistory

> ChannelEventHistoryResult NotificationsApiNotificationChannelEventHistory(ctx, id).StartTime(startTime).EndTime(endTime).Execute()

NotificationChannelEventHistory notifications-api

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
	id := "channel-12345678" // string | Unique identifier of the channel
	startTime := time.Now() // time.Time | The start time of the range (optional)
	endTime := time.Now() // time.Time | The end time of the range (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.NotificationsApiAPI.NotificationsApiNotificationChannelEventHistory(context.Background(), id).StartTime(startTime).EndTime(endTime).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `NotificationsApiAPI.NotificationsApiNotificationChannelEventHistory``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `NotificationsApiNotificationChannelEventHistory`: ChannelEventHistoryResult
	fmt.Fprintf(os.Stdout, "Response from `NotificationsApiAPI.NotificationsApiNotificationChannelEventHistory`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **string** | Unique identifier of the channel | 

### Other Parameters

Other parameters are passed through a pointer to a apiNotificationsApiNotificationChannelEventHistoryRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **startTime** | **time.Time** | The start time of the range | 
 **endTime** | **time.Time** | The end time of the range | 

### Return type

[**ChannelEventHistoryResult**](ChannelEventHistoryResult.md)

### Authorization

[api_key_header_Authorization](../README.md#api_key_header_Authorization)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json, application/vnd.goa.error

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## NotificationsApiReplayEvent

> NotificationsApiReplayEvent(ctx, id).Execute()

ReplayEvent notifications-api

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
	id := "event-12345678" // string | Unique identifier of the event to replay

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	r, err := apiClient.NotificationsApiAPI.NotificationsApiReplayEvent(context.Background(), id).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `NotificationsApiAPI.NotificationsApiReplayEvent``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **string** | Unique identifier of the event to replay | 

### Other Parameters

Other parameters are passed through a pointer to a apiNotificationsApiReplayEventRequest struct via the builder pattern


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


## NotificationsApiUpdateNotificationChannel

> NotificationsApiUpdateNotificationChannel(ctx, id).UpdateNotificationChannelRequest2(updateNotificationChannelRequest2).Execute()

UpdateNotificationChannel notifications-api

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
	id := "channel-12345678" // string | Unique identifier of the channel
	updateNotificationChannelRequest2 := *openapiclient.NewUpdateNotificationChannelRequest2() // UpdateNotificationChannelRequest2 | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	r, err := apiClient.NotificationsApiAPI.NotificationsApiUpdateNotificationChannel(context.Background(), id).UpdateNotificationChannelRequest2(updateNotificationChannelRequest2).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `NotificationsApiAPI.NotificationsApiUpdateNotificationChannel``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **string** | Unique identifier of the channel | 

### Other Parameters

Other parameters are passed through a pointer to a apiNotificationsApiUpdateNotificationChannelRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **updateNotificationChannelRequest2** | [**UpdateNotificationChannelRequest2**](UpdateNotificationChannelRequest2.md) |  | 

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

