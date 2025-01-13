# \NotificationsApiAPI

All URIs are relative to *https://api.omnistrate.cloud*

Method | HTTP request | Description
------------- | ------------- | -------------
[**NotificationsApiCreateNotificationChannel**](NotificationsApiAPI.md#NotificationsApiCreateNotificationChannel) | **Post** /2022-09-01-00/fleet/notification-channel | CreateNotificationChannel notifications-api
[**NotificationsApiDeleteNotificationChannel**](NotificationsApiAPI.md#NotificationsApiDeleteNotificationChannel) | **Delete** /2022-09-01-00/fleet/notification-channel/{id} | DeleteNotificationChannel notifications-api
[**NotificationsApiDescribeNotificationChannel**](NotificationsApiAPI.md#NotificationsApiDescribeNotificationChannel) | **Get** /2022-09-01-00/fleet/notification-channel/{id} | DescribeNotificationChannel notifications-api
[**NotificationsApiListNotificationChannels**](NotificationsApiAPI.md#NotificationsApiListNotificationChannels) | **Get** /2022-09-01-00/fleet/notification-channel | ListNotificationChannels notifications-api
[**NotificationsApiUpdateNotificationChannel**](NotificationsApiAPI.md#NotificationsApiUpdateNotificationChannel) | **Patch** /2022-09-01-00/fleet/notification-channel/{id} | UpdateNotificationChannel notifications-api



## NotificationsApiCreateNotificationChannel

> string NotificationsApiCreateNotificationChannel(ctx).CreateNotificationChannelRequestBody(createNotificationChannelRequestBody).Execute()

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
	createNotificationChannelRequestBody := *openapiclient.NewCreateNotificationChannelRequestBody("My Channel", *openapiclient.NewChannelSubscription([]string{"Ut ea saepe voluptatibus velit a."}, []string{"Qui eum itaque occaecati corrupti ut reiciendis."}, []string{"Temporibus mollitia quo omnis illo."}, []string{"Voluptate est modi."}, []string{"Sint enim amet repellat suscipit et iusto."})) // CreateNotificationChannelRequestBody | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.NotificationsApiAPI.NotificationsApiCreateNotificationChannel(context.Background()).CreateNotificationChannelRequestBody(createNotificationChannelRequestBody).Execute()
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
 **createNotificationChannelRequestBody** | [**CreateNotificationChannelRequestBody**](CreateNotificationChannelRequestBody.md) |  | 

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


## NotificationsApiUpdateNotificationChannel

> NotificationsApiUpdateNotificationChannel(ctx, id).UpdateNotificationChannelRequestBody(updateNotificationChannelRequestBody).Execute()

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
	updateNotificationChannelRequestBody := *openapiclient.NewUpdateNotificationChannelRequestBody() // UpdateNotificationChannelRequestBody | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	r, err := apiClient.NotificationsApiAPI.NotificationsApiUpdateNotificationChannel(context.Background(), id).UpdateNotificationChannelRequestBody(updateNotificationChannelRequestBody).Execute()
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

 **updateNotificationChannelRequestBody** | [**UpdateNotificationChannelRequestBody**](UpdateNotificationChannelRequestBody.md) |  | 

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

