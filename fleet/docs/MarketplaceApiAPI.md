# \MarketplaceApiAPI

All URIs are relative to *https://api.omnistrate.cloud*

Method | HTTP request | Description
------------- | ------------- | -------------
[**MarketplaceApiConfirmMarketplaceFulfillment**](MarketplaceApiAPI.md#MarketplaceApiConfirmMarketplaceFulfillment) | **Post** /2022-09-01-00/fleet/marketplace/contract/{id}/fulfillment/confirm | ConfirmMarketplaceFulfillment marketplace-api
[**MarketplaceApiConnectMarketplaceChannel**](MarketplaceApiAPI.md#MarketplaceApiConnectMarketplaceChannel) | **Post** /2022-09-01-00/fleet/marketplace/channel | ConnectMarketplaceChannel marketplace-api
[**MarketplaceApiDenyMarketplaceFulfillment**](MarketplaceApiAPI.md#MarketplaceApiDenyMarketplaceFulfillment) | **Post** /2022-09-01-00/fleet/marketplace/contract/{id}/fulfillment/deny | DenyMarketplaceFulfillment marketplace-api
[**MarketplaceApiDescribeMarketplaceChannel**](MarketplaceApiAPI.md#MarketplaceApiDescribeMarketplaceChannel) | **Get** /2022-09-01-00/fleet/marketplace/channel/{channel} | DescribeMarketplaceChannel marketplace-api
[**MarketplaceApiDescribeMarketplaceContract**](MarketplaceApiAPI.md#MarketplaceApiDescribeMarketplaceContract) | **Get** /2022-09-01-00/fleet/marketplace/contract/{id} | DescribeMarketplaceContract marketplace-api
[**MarketplaceApiDescribeMarketplaceFulfillment**](MarketplaceApiAPI.md#MarketplaceApiDescribeMarketplaceFulfillment) | **Get** /2022-09-01-00/fleet/marketplace/contract/{id}/fulfillment | DescribeMarketplaceFulfillment marketplace-api
[**MarketplaceApiDescribeSandboxRun**](MarketplaceApiAPI.md#MarketplaceApiDescribeSandboxRun) | **Get** /2022-09-01-00/fleet/marketplace/sandbox/run | DescribeSandboxRun marketplace-api
[**MarketplaceApiDisconnectMarketplaceChannel**](MarketplaceApiAPI.md#MarketplaceApiDisconnectMarketplaceChannel) | **Delete** /2022-09-01-00/fleet/marketplace/channel/{channel} | DisconnectMarketplaceChannel marketplace-api
[**MarketplaceApiListMarketplaceChannelListings**](MarketplaceApiAPI.md#MarketplaceApiListMarketplaceChannelListings) | **Get** /2022-09-01-00/fleet/marketplace/channel/{channel}/listing | ListMarketplaceChannelListings marketplace-api
[**MarketplaceApiListMarketplaceChannels**](MarketplaceApiAPI.md#MarketplaceApiListMarketplaceChannels) | **Get** /2022-09-01-00/fleet/marketplace/channel | ListMarketplaceChannels marketplace-api
[**MarketplaceApiListMarketplaceContracts**](MarketplaceApiAPI.md#MarketplaceApiListMarketplaceContracts) | **Get** /2022-09-01-00/fleet/marketplace/contract | ListMarketplaceContracts marketplace-api
[**MarketplaceApiListMarketplaceDeliveries**](MarketplaceApiAPI.md#MarketplaceApiListMarketplaceDeliveries) | **Get** /2022-09-01-00/fleet/marketplace/delivery | ListMarketplaceDeliveries marketplace-api
[**MarketplaceApiMarketplaceLanding**](MarketplaceApiAPI.md#MarketplaceApiMarketplaceLanding) | **Get** /2022-09-01-00/fleet/marketplace/land/{channel}/{serviceProviderOrgId} | MarketplaceLanding marketplace-api
[**MarketplaceApiRedeemHandoff**](MarketplaceApiAPI.md#MarketplaceApiRedeemHandoff) | **Post** /2022-09-01-00/fleet/marketplace/handoff/redeem | RedeemHandoff marketplace-api
[**MarketplaceApiRedeliverMarketplaceDelivery**](MarketplaceApiAPI.md#MarketplaceApiRedeliverMarketplaceDelivery) | **Post** /2022-09-01-00/fleet/marketplace/delivery/{deliveryId}/redeliver | RedeliverMarketplaceDelivery marketplace-api
[**MarketplaceApiRetryMarketplaceFulfillment**](MarketplaceApiAPI.md#MarketplaceApiRetryMarketplaceFulfillment) | **Post** /2022-09-01-00/fleet/marketplace/contract/{id}/fulfillment/retry | RetryMarketplaceFulfillment marketplace-api
[**MarketplaceApiRotateSandboxSecret**](MarketplaceApiAPI.md#MarketplaceApiRotateSandboxSecret) | **Post** /2022-09-01-00/fleet/marketplace/sandbox/secret/rotate | RotateSandboxSecret marketplace-api
[**MarketplaceApiSandboxCheckout**](MarketplaceApiAPI.md#MarketplaceApiSandboxCheckout) | **Get** /2022-09-01-00/fleet/marketplace/sandbox/checkout/{externalRef} | SandboxCheckout marketplace-api
[**MarketplaceApiSendSandboxEvent**](MarketplaceApiAPI.md#MarketplaceApiSendSandboxEvent) | **Post** /2022-09-01-00/fleet/marketplace/sandbox/event | SendSandboxEvent marketplace-api
[**MarketplaceApiSimulateContractEvent**](MarketplaceApiAPI.md#MarketplaceApiSimulateContractEvent) | **Post** /2022-09-01-00/fleet/marketplace/sandbox/simulate | SimulateContractEvent marketplace-api
[**MarketplaceApiStartSandboxRun**](MarketplaceApiAPI.md#MarketplaceApiStartSandboxRun) | **Post** /2022-09-01-00/fleet/marketplace/sandbox/run | StartSandboxRun marketplace-api
[**MarketplaceApiUpdateMarketplaceChannel**](MarketplaceApiAPI.md#MarketplaceApiUpdateMarketplaceChannel) | **Patch** /2022-09-01-00/fleet/marketplace/channel/{channel} | UpdateMarketplaceChannel marketplace-api
[**MarketplaceApiValidateMarketplaceChannel**](MarketplaceApiAPI.md#MarketplaceApiValidateMarketplaceChannel) | **Post** /2022-09-01-00/fleet/marketplace/channel/{channel}/validate | ValidateMarketplaceChannel marketplace-api



## MarketplaceApiConfirmMarketplaceFulfillment

> MarketplaceFulfillment MarketplaceApiConfirmMarketplaceFulfillment(ctx, id).ConfirmMarketplaceFulfillmentRequest2(confirmMarketplaceFulfillmentRequest2).Execute()

ConfirmMarketplaceFulfillment marketplace-api

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
	id := "mkc-4t8w2qbnz1lp" // string | The Omnistrate contract identifier, as carried by marketplaceContractId on the contract.discovered event and returned by the handoff redeem
	confirmMarketplaceFulfillmentRequest2 := *openapiclient.NewConfirmMarketplaceFulfillmentRequest2() // ConfirmMarketplaceFulfillmentRequest2 | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.MarketplaceApiAPI.MarketplaceApiConfirmMarketplaceFulfillment(context.Background(), id).ConfirmMarketplaceFulfillmentRequest2(confirmMarketplaceFulfillmentRequest2).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `MarketplaceApiAPI.MarketplaceApiConfirmMarketplaceFulfillment``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `MarketplaceApiConfirmMarketplaceFulfillment`: MarketplaceFulfillment
	fmt.Fprintf(os.Stdout, "Response from `MarketplaceApiAPI.MarketplaceApiConfirmMarketplaceFulfillment`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **string** | The Omnistrate contract identifier, as carried by marketplaceContractId on the contract.discovered event and returned by the handoff redeem | 

### Other Parameters

Other parameters are passed through a pointer to a apiMarketplaceApiConfirmMarketplaceFulfillmentRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **confirmMarketplaceFulfillmentRequest2** | [**ConfirmMarketplaceFulfillmentRequest2**](ConfirmMarketplaceFulfillmentRequest2.md) |  | 

### Return type

[**MarketplaceFulfillment**](MarketplaceFulfillment.md)

### Authorization

[api_key_header_Authorization](../README.md#api_key_header_Authorization)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json, application/vnd.goa.error

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## MarketplaceApiConnectMarketplaceChannel

> MarketplaceChannelConfig MarketplaceApiConnectMarketplaceChannel(ctx).ConnectMarketplaceChannelRequest2(connectMarketplaceChannelRequest2).Execute()

ConnectMarketplaceChannel marketplace-api

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
	connectMarketplaceChannelRequest2 := *openapiclient.NewConnectMarketplaceChannelRequest2("SUGER|SANDBOX") // ConnectMarketplaceChannelRequest2 | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.MarketplaceApiAPI.MarketplaceApiConnectMarketplaceChannel(context.Background()).ConnectMarketplaceChannelRequest2(connectMarketplaceChannelRequest2).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `MarketplaceApiAPI.MarketplaceApiConnectMarketplaceChannel``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `MarketplaceApiConnectMarketplaceChannel`: MarketplaceChannelConfig
	fmt.Fprintf(os.Stdout, "Response from `MarketplaceApiAPI.MarketplaceApiConnectMarketplaceChannel`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiMarketplaceApiConnectMarketplaceChannelRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **connectMarketplaceChannelRequest2** | [**ConnectMarketplaceChannelRequest2**](ConnectMarketplaceChannelRequest2.md) |  | 

### Return type

[**MarketplaceChannelConfig**](MarketplaceChannelConfig.md)

### Authorization

[api_key_header_Authorization](../README.md#api_key_header_Authorization)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json, application/vnd.goa.error

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## MarketplaceApiDenyMarketplaceFulfillment

> MarketplaceFulfillment MarketplaceApiDenyMarketplaceFulfillment(ctx, id).DenyMarketplaceFulfillmentRequest2(denyMarketplaceFulfillmentRequest2).Execute()

DenyMarketplaceFulfillment marketplace-api

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
	id := "mkc-4t8w2qbnz1lp" // string | The Omnistrate contract identifier
	denyMarketplaceFulfillmentRequest2 := *openapiclient.NewDenyMarketplaceFulfillmentRequest2() // DenyMarketplaceFulfillmentRequest2 | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.MarketplaceApiAPI.MarketplaceApiDenyMarketplaceFulfillment(context.Background(), id).DenyMarketplaceFulfillmentRequest2(denyMarketplaceFulfillmentRequest2).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `MarketplaceApiAPI.MarketplaceApiDenyMarketplaceFulfillment``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `MarketplaceApiDenyMarketplaceFulfillment`: MarketplaceFulfillment
	fmt.Fprintf(os.Stdout, "Response from `MarketplaceApiAPI.MarketplaceApiDenyMarketplaceFulfillment`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **string** | The Omnistrate contract identifier | 

### Other Parameters

Other parameters are passed through a pointer to a apiMarketplaceApiDenyMarketplaceFulfillmentRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **denyMarketplaceFulfillmentRequest2** | [**DenyMarketplaceFulfillmentRequest2**](DenyMarketplaceFulfillmentRequest2.md) |  | 

### Return type

[**MarketplaceFulfillment**](MarketplaceFulfillment.md)

### Authorization

[api_key_header_Authorization](../README.md#api_key_header_Authorization)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json, application/vnd.goa.error

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## MarketplaceApiDescribeMarketplaceChannel

> MarketplaceChannelConfig MarketplaceApiDescribeMarketplaceChannel(ctx, channel).Execute()

DescribeMarketplaceChannel marketplace-api

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
	channel := "SUGER|SANDBOX" // string | Which channel to describe

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.MarketplaceApiAPI.MarketplaceApiDescribeMarketplaceChannel(context.Background(), channel).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `MarketplaceApiAPI.MarketplaceApiDescribeMarketplaceChannel``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `MarketplaceApiDescribeMarketplaceChannel`: MarketplaceChannelConfig
	fmt.Fprintf(os.Stdout, "Response from `MarketplaceApiAPI.MarketplaceApiDescribeMarketplaceChannel`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**channel** | **string** | Which channel to describe | 

### Other Parameters

Other parameters are passed through a pointer to a apiMarketplaceApiDescribeMarketplaceChannelRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

[**MarketplaceChannelConfig**](MarketplaceChannelConfig.md)

### Authorization

[api_key_header_Authorization](../README.md#api_key_header_Authorization)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json, application/vnd.goa.error

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## MarketplaceApiDescribeMarketplaceContract

> MarketplaceContract MarketplaceApiDescribeMarketplaceContract(ctx, id).Channel(channel).ExternalRef(externalRef).Execute()

DescribeMarketplaceContract marketplace-api

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
	id := "mkc-4t8w2qbnz1lp" // string | The Omnistrate contract identifier
	channel := "SUGER|SANDBOX" // string | With externalRef, resolves a contract from the marketplace identifier alone, which is what an ISV has to hand when the buyer contacts them (optional)
	externalRef := "Veniam suscipit provident cumque." // string |  (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.MarketplaceApiAPI.MarketplaceApiDescribeMarketplaceContract(context.Background(), id).Channel(channel).ExternalRef(externalRef).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `MarketplaceApiAPI.MarketplaceApiDescribeMarketplaceContract``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `MarketplaceApiDescribeMarketplaceContract`: MarketplaceContract
	fmt.Fprintf(os.Stdout, "Response from `MarketplaceApiAPI.MarketplaceApiDescribeMarketplaceContract`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **string** | The Omnistrate contract identifier | 

### Other Parameters

Other parameters are passed through a pointer to a apiMarketplaceApiDescribeMarketplaceContractRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **channel** | **string** | With externalRef, resolves a contract from the marketplace identifier alone, which is what an ISV has to hand when the buyer contacts them | 
 **externalRef** | **string** |  | 

### Return type

[**MarketplaceContract**](MarketplaceContract.md)

### Authorization

[api_key_header_Authorization](../README.md#api_key_header_Authorization)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json, application/vnd.goa.error

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## MarketplaceApiDescribeMarketplaceFulfillment

> MarketplaceFulfillment MarketplaceApiDescribeMarketplaceFulfillment(ctx, id).Execute()

DescribeMarketplaceFulfillment marketplace-api

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
	id := "mkc-4t8w2qbnz1lp" // string | The Omnistrate contract identifier

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.MarketplaceApiAPI.MarketplaceApiDescribeMarketplaceFulfillment(context.Background(), id).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `MarketplaceApiAPI.MarketplaceApiDescribeMarketplaceFulfillment``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `MarketplaceApiDescribeMarketplaceFulfillment`: MarketplaceFulfillment
	fmt.Fprintf(os.Stdout, "Response from `MarketplaceApiAPI.MarketplaceApiDescribeMarketplaceFulfillment`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **string** | The Omnistrate contract identifier | 

### Other Parameters

Other parameters are passed through a pointer to a apiMarketplaceApiDescribeMarketplaceFulfillmentRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

[**MarketplaceFulfillment**](MarketplaceFulfillment.md)

### Authorization

[api_key_header_Authorization](../README.md#api_key_header_Authorization)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json, application/vnd.goa.error

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## MarketplaceApiDescribeSandboxRun

> SandboxRun MarketplaceApiDescribeSandboxRun(ctx).Execute()

DescribeSandboxRun marketplace-api

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
	resp, r, err := apiClient.MarketplaceApiAPI.MarketplaceApiDescribeSandboxRun(context.Background()).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `MarketplaceApiAPI.MarketplaceApiDescribeSandboxRun``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `MarketplaceApiDescribeSandboxRun`: SandboxRun
	fmt.Fprintf(os.Stdout, "Response from `MarketplaceApiAPI.MarketplaceApiDescribeSandboxRun`: %v\n", resp)
}
```

### Path Parameters

This endpoint does not need any parameter.

### Other Parameters

Other parameters are passed through a pointer to a apiMarketplaceApiDescribeSandboxRunRequest struct via the builder pattern


### Return type

[**SandboxRun**](SandboxRun.md)

### Authorization

[api_key_header_Authorization](../README.md#api_key_header_Authorization)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json, application/vnd.goa.error

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## MarketplaceApiDisconnectMarketplaceChannel

> MarketplaceApiDisconnectMarketplaceChannel(ctx, channel).Execute()

DisconnectMarketplaceChannel marketplace-api

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
	channel := "SUGER|SANDBOX" // string | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	r, err := apiClient.MarketplaceApiAPI.MarketplaceApiDisconnectMarketplaceChannel(context.Background(), channel).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `MarketplaceApiAPI.MarketplaceApiDisconnectMarketplaceChannel``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**channel** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiMarketplaceApiDisconnectMarketplaceChannelRequest struct via the builder pattern


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


## MarketplaceApiListMarketplaceChannelListings

> ListMarketplaceChannelListingsResult MarketplaceApiListMarketplaceChannelListings(ctx, channel).Execute()

ListMarketplaceChannelListings marketplace-api

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
	channel := "SUGER|SANDBOX" // string | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.MarketplaceApiAPI.MarketplaceApiListMarketplaceChannelListings(context.Background(), channel).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `MarketplaceApiAPI.MarketplaceApiListMarketplaceChannelListings``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `MarketplaceApiListMarketplaceChannelListings`: ListMarketplaceChannelListingsResult
	fmt.Fprintf(os.Stdout, "Response from `MarketplaceApiAPI.MarketplaceApiListMarketplaceChannelListings`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**channel** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiMarketplaceApiListMarketplaceChannelListingsRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

[**ListMarketplaceChannelListingsResult**](ListMarketplaceChannelListingsResult.md)

### Authorization

[api_key_header_Authorization](../README.md#api_key_header_Authorization)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json, application/vnd.goa.error

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## MarketplaceApiListMarketplaceChannels

> ListMarketplaceChannelsResult MarketplaceApiListMarketplaceChannels(ctx).Execute()

ListMarketplaceChannels marketplace-api

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
	resp, r, err := apiClient.MarketplaceApiAPI.MarketplaceApiListMarketplaceChannels(context.Background()).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `MarketplaceApiAPI.MarketplaceApiListMarketplaceChannels``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `MarketplaceApiListMarketplaceChannels`: ListMarketplaceChannelsResult
	fmt.Fprintf(os.Stdout, "Response from `MarketplaceApiAPI.MarketplaceApiListMarketplaceChannels`: %v\n", resp)
}
```

### Path Parameters

This endpoint does not need any parameter.

### Other Parameters

Other parameters are passed through a pointer to a apiMarketplaceApiListMarketplaceChannelsRequest struct via the builder pattern


### Return type

[**ListMarketplaceChannelsResult**](ListMarketplaceChannelsResult.md)

### Authorization

[api_key_header_Authorization](../README.md#api_key_header_Authorization)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json, application/vnd.goa.error

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## MarketplaceApiListMarketplaceContracts

> ListMarketplaceContractsResult MarketplaceApiListMarketplaceContracts(ctx).Channel(channel).FulfillmentState(fulfillmentState).Quadrant(quadrant).NeedsAttention(needsAttention).IncludeSimulated(includeSimulated).Execute()

ListMarketplaceContracts marketplace-api

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
	channel := "SUGER|SANDBOX" // string | Filter to one channel (optional)
	fulfillmentState := "DISCOVERED|IDENTIFIED|AWAITING_ISV|READY|SUSPENDED|DEPROVISIONING|CLOSED" // string | Filter to one fulfillment state. This is how an operator lists the fulfillments that are stuck, rather than through a separate endpoint (optional)
	quadrant := "HEALTHY|ONBOARDING|PAYING_NOT_PROVISIONED|SERVING_NOT_BILLED|CLOSED" // string | Filter to one reconciliation quadrant (optional)
	needsAttention := true // bool | Only contracts past the handoff SLA while billable. The operator's working queue: every one of these is a customer paying for something they cannot use (optional)
	includeSimulated := false // bool | Sandbox contracts are excluded unless this is set, so a revenue view cannot pick them up by forgetting a filter (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.MarketplaceApiAPI.MarketplaceApiListMarketplaceContracts(context.Background()).Channel(channel).FulfillmentState(fulfillmentState).Quadrant(quadrant).NeedsAttention(needsAttention).IncludeSimulated(includeSimulated).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `MarketplaceApiAPI.MarketplaceApiListMarketplaceContracts``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `MarketplaceApiListMarketplaceContracts`: ListMarketplaceContractsResult
	fmt.Fprintf(os.Stdout, "Response from `MarketplaceApiAPI.MarketplaceApiListMarketplaceContracts`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiMarketplaceApiListMarketplaceContractsRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **channel** | **string** | Filter to one channel | 
 **fulfillmentState** | **string** | Filter to one fulfillment state. This is how an operator lists the fulfillments that are stuck, rather than through a separate endpoint | 
 **quadrant** | **string** | Filter to one reconciliation quadrant | 
 **needsAttention** | **bool** | Only contracts past the handoff SLA while billable. The operator&#39;s working queue: every one of these is a customer paying for something they cannot use | 
 **includeSimulated** | **bool** | Sandbox contracts are excluded unless this is set, so a revenue view cannot pick them up by forgetting a filter | 

### Return type

[**ListMarketplaceContractsResult**](ListMarketplaceContractsResult.md)

### Authorization

[api_key_header_Authorization](../README.md#api_key_header_Authorization)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json, application/vnd.goa.error

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## MarketplaceApiListMarketplaceDeliveries

> ListMarketplaceDeliveriesResult MarketplaceApiListMarketplaceDeliveries(ctx).Channel(channel).MarketplaceContractId(marketplaceContractId).Direction(direction).Status(status).EventType(eventType).FailuresOnly(failuresOnly).Since(since).Execute()

ListMarketplaceDeliveries marketplace-api

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
	channel := "SUGER|SANDBOX" // string | Filter to one channel (optional)
	marketplaceContractId := "Laborum est impedit temporibus." // string | Filter to one contract, which is how a contract detail view scopes its own trail (optional)
	direction := "OUTBOUND|INBOUND" // string |  (optional)
	status := "PENDING|DELIVERED|FAILED|BLOCKED" // string |  (optional)
	eventType := "contract.discovered|entitlement.updated|contract.suspended|contract.cancelled|fulfillment.failed" // string |  (optional)
	failuresOnly := false // bool | A shortcut for the only filter combination anybody types twice (optional)
	since := time.Now() // time.Time | Deliveries sent or received at or after this instant (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.MarketplaceApiAPI.MarketplaceApiListMarketplaceDeliveries(context.Background()).Channel(channel).MarketplaceContractId(marketplaceContractId).Direction(direction).Status(status).EventType(eventType).FailuresOnly(failuresOnly).Since(since).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `MarketplaceApiAPI.MarketplaceApiListMarketplaceDeliveries``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `MarketplaceApiListMarketplaceDeliveries`: ListMarketplaceDeliveriesResult
	fmt.Fprintf(os.Stdout, "Response from `MarketplaceApiAPI.MarketplaceApiListMarketplaceDeliveries`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiMarketplaceApiListMarketplaceDeliveriesRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **channel** | **string** | Filter to one channel | 
 **marketplaceContractId** | **string** | Filter to one contract, which is how a contract detail view scopes its own trail | 
 **direction** | **string** |  | 
 **status** | **string** |  | 
 **eventType** | **string** |  | 
 **failuresOnly** | **bool** | A shortcut for the only filter combination anybody types twice | 
 **since** | **time.Time** | Deliveries sent or received at or after this instant | 

### Return type

[**ListMarketplaceDeliveriesResult**](ListMarketplaceDeliveriesResult.md)

### Authorization

[api_key_header_Authorization](../README.md#api_key_header_Authorization)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json, application/vnd.goa.error

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## MarketplaceApiMarketplaceLanding

> MarketplaceApiMarketplaceLanding(ctx, channel, serviceProviderOrgId).SugerEntitlementId(sugerEntitlementId).Partner(partner).OfferType(offerType).Execute()

MarketplaceLanding marketplace-api

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
	channel := "SUGER|SANDBOX" // string | Which channel the buyer arrived from. Selects the channel implementation that knows how to read this arrival's token back
	serviceProviderOrgId := "org-8Hn2Kq4Vd1" // string | Which ISV organization's listing was purchased. A SELECTOR for whose stored channel credentials perform the server-side readback, not a credential: it grants nothing, and substituting another organization's id resolves the token against an account where it does not exist
	sugerEntitlementId := "ent_01J9Q7VZ3K8MTRQ2X4W6H0N5PD" // string | Suger's entitlement identifier, appended by Suger's signup redirect. A pointer to be read back, never believed (optional)
	partner := "aws|azure|gcp" // string | Which cloud the purchase originated on, as Suger reports it. Recorded for audit; the authoritative value comes from the readback (optional)
	offerType := "Aut quaerat omnis in totam voluptatem." // string | Suger's offer type for the purchase. Recorded for audit only (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	r, err := apiClient.MarketplaceApiAPI.MarketplaceApiMarketplaceLanding(context.Background(), channel, serviceProviderOrgId).SugerEntitlementId(sugerEntitlementId).Partner(partner).OfferType(offerType).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `MarketplaceApiAPI.MarketplaceApiMarketplaceLanding``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**channel** | **string** | Which channel the buyer arrived from. Selects the channel implementation that knows how to read this arrival&#39;s token back | 
**serviceProviderOrgId** | **string** | Which ISV organization&#39;s listing was purchased. A SELECTOR for whose stored channel credentials perform the server-side readback, not a credential: it grants nothing, and substituting another organization&#39;s id resolves the token against an account where it does not exist | 

### Other Parameters

Other parameters are passed through a pointer to a apiMarketplaceApiMarketplaceLandingRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


 **sugerEntitlementId** | **string** | Suger&#39;s entitlement identifier, appended by Suger&#39;s signup redirect. A pointer to be read back, never believed | 
 **partner** | **string** | Which cloud the purchase originated on, as Suger reports it. Recorded for audit; the authoritative value comes from the readback | 
 **offerType** | **string** | Suger&#39;s offer type for the purchase. Recorded for audit only | 

### Return type

 (empty response body)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/vnd.goa.error

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## MarketplaceApiRedeemHandoff

> RedeemHandoffResult MarketplaceApiRedeemHandoff(ctx).RedeemHandoffRequest2(redeemHandoffRequest2).Execute()

RedeemHandoff marketplace-api

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
	redeemHandoffRequest2 := *openapiclient.NewRedeemHandoffRequest2("hoff_01J9XQ3M7ZK4V8N2PY6TBW5RCD") // RedeemHandoffRequest2 | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.MarketplaceApiAPI.MarketplaceApiRedeemHandoff(context.Background()).RedeemHandoffRequest2(redeemHandoffRequest2).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `MarketplaceApiAPI.MarketplaceApiRedeemHandoff``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `MarketplaceApiRedeemHandoff`: RedeemHandoffResult
	fmt.Fprintf(os.Stdout, "Response from `MarketplaceApiAPI.MarketplaceApiRedeemHandoff`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiMarketplaceApiRedeemHandoffRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **redeemHandoffRequest2** | [**RedeemHandoffRequest2**](RedeemHandoffRequest2.md) |  | 

### Return type

[**RedeemHandoffResult**](RedeemHandoffResult.md)

### Authorization

[api_key_header_Authorization](../README.md#api_key_header_Authorization)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json, application/vnd.goa.error

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## MarketplaceApiRedeliverMarketplaceDelivery

> MarketplaceDelivery MarketplaceApiRedeliverMarketplaceDelivery(ctx, deliveryId).RedeliverMarketplaceDeliveryRequest2(redeliverMarketplaceDeliveryRequest2).Execute()

RedeliverMarketplaceDelivery marketplace-api

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
	deliveryId := "Saepe nobis aperiam sint." // string | Which delivery to send again. Must be an outbound one: there is nothing to redeliver about a call the ISV made. The redelivery carries the SAME eventId, so a receiver that already processed this event correctly treats it as a duplicate and does nothing
	redeliverMarketplaceDeliveryRequest2 := *openapiclient.NewRedeliverMarketplaceDeliveryRequest2() // RedeliverMarketplaceDeliveryRequest2 | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.MarketplaceApiAPI.MarketplaceApiRedeliverMarketplaceDelivery(context.Background(), deliveryId).RedeliverMarketplaceDeliveryRequest2(redeliverMarketplaceDeliveryRequest2).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `MarketplaceApiAPI.MarketplaceApiRedeliverMarketplaceDelivery``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `MarketplaceApiRedeliverMarketplaceDelivery`: MarketplaceDelivery
	fmt.Fprintf(os.Stdout, "Response from `MarketplaceApiAPI.MarketplaceApiRedeliverMarketplaceDelivery`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**deliveryId** | **string** | Which delivery to send again. Must be an outbound one: there is nothing to redeliver about a call the ISV made. The redelivery carries the SAME eventId, so a receiver that already processed this event correctly treats it as a duplicate and does nothing | 

### Other Parameters

Other parameters are passed through a pointer to a apiMarketplaceApiRedeliverMarketplaceDeliveryRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **redeliverMarketplaceDeliveryRequest2** | [**RedeliverMarketplaceDeliveryRequest2**](RedeliverMarketplaceDeliveryRequest2.md) |  | 

### Return type

[**MarketplaceDelivery**](MarketplaceDelivery.md)

### Authorization

[api_key_header_Authorization](../README.md#api_key_header_Authorization)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json, application/vnd.goa.error

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## MarketplaceApiRetryMarketplaceFulfillment

> MarketplaceFulfillment MarketplaceApiRetryMarketplaceFulfillment(ctx, id).RetryMarketplaceFulfillmentRequest2(retryMarketplaceFulfillmentRequest2).Execute()

RetryMarketplaceFulfillment marketplace-api

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
	id := "mkc-4t8w2qbnz1lp" // string | The Omnistrate contract identifier
	retryMarketplaceFulfillmentRequest2 := *openapiclient.NewRetryMarketplaceFulfillmentRequest2() // RetryMarketplaceFulfillmentRequest2 | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.MarketplaceApiAPI.MarketplaceApiRetryMarketplaceFulfillment(context.Background(), id).RetryMarketplaceFulfillmentRequest2(retryMarketplaceFulfillmentRequest2).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `MarketplaceApiAPI.MarketplaceApiRetryMarketplaceFulfillment``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `MarketplaceApiRetryMarketplaceFulfillment`: MarketplaceFulfillment
	fmt.Fprintf(os.Stdout, "Response from `MarketplaceApiAPI.MarketplaceApiRetryMarketplaceFulfillment`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **string** | The Omnistrate contract identifier | 

### Other Parameters

Other parameters are passed through a pointer to a apiMarketplaceApiRetryMarketplaceFulfillmentRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **retryMarketplaceFulfillmentRequest2** | [**RetryMarketplaceFulfillmentRequest2**](RetryMarketplaceFulfillmentRequest2.md) |  | 

### Return type

[**MarketplaceFulfillment**](MarketplaceFulfillment.md)

### Authorization

[api_key_header_Authorization](../README.md#api_key_header_Authorization)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json, application/vnd.goa.error

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## MarketplaceApiRotateSandboxSecret

> RotateSandboxSecretResult MarketplaceApiRotateSandboxSecret(ctx).Execute()

RotateSandboxSecret marketplace-api

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
	resp, r, err := apiClient.MarketplaceApiAPI.MarketplaceApiRotateSandboxSecret(context.Background()).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `MarketplaceApiAPI.MarketplaceApiRotateSandboxSecret``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `MarketplaceApiRotateSandboxSecret`: RotateSandboxSecretResult
	fmt.Fprintf(os.Stdout, "Response from `MarketplaceApiAPI.MarketplaceApiRotateSandboxSecret`: %v\n", resp)
}
```

### Path Parameters

This endpoint does not need any parameter.

### Other Parameters

Other parameters are passed through a pointer to a apiMarketplaceApiRotateSandboxSecretRequest struct via the builder pattern


### Return type

[**RotateSandboxSecretResult**](RotateSandboxSecretResult.md)

### Authorization

[api_key_header_Authorization](../README.md#api_key_header_Authorization)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json, application/vnd.goa.error

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## MarketplaceApiSandboxCheckout

> MarketplaceApiSandboxCheckout(ctx, externalRef).Execute()

SandboxCheckout marketplace-api

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
	externalRef := "sbx-ent-4f2a9c1b" // string | The armed purchase's reference on the simulated channel, which is what the purchase control returned. Resolved against the CALLER'S sandbox: a reference belonging to another organization is simply not found

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	r, err := apiClient.MarketplaceApiAPI.MarketplaceApiSandboxCheckout(context.Background(), externalRef).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `MarketplaceApiAPI.MarketplaceApiSandboxCheckout``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**externalRef** | **string** | The armed purchase&#39;s reference on the simulated channel, which is what the purchase control returned. Resolved against the CALLER&#39;S sandbox: a reference belonging to another organization is simply not found | 

### Other Parameters

Other parameters are passed through a pointer to a apiMarketplaceApiSandboxCheckoutRequest struct via the builder pattern


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


## MarketplaceApiSendSandboxEvent

> SandboxRun MarketplaceApiSendSandboxEvent(ctx).SendSandboxEventRequest2(sendSandboxEventRequest2).Execute()

SendSandboxEvent marketplace-api

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
	sendSandboxEventRequest2 := *openapiclient.NewSendSandboxEventRequest2("contract.discovered|entitlement.updated|contract.suspended|contract.cancelled|fulfillment.failed") // SendSandboxEventRequest2 | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.MarketplaceApiAPI.MarketplaceApiSendSandboxEvent(context.Background()).SendSandboxEventRequest2(sendSandboxEventRequest2).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `MarketplaceApiAPI.MarketplaceApiSendSandboxEvent``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `MarketplaceApiSendSandboxEvent`: SandboxRun
	fmt.Fprintf(os.Stdout, "Response from `MarketplaceApiAPI.MarketplaceApiSendSandboxEvent`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiMarketplaceApiSendSandboxEventRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **sendSandboxEventRequest2** | [**SendSandboxEventRequest2**](SendSandboxEventRequest2.md) |  | 

### Return type

[**SandboxRun**](SandboxRun.md)

### Authorization

[api_key_header_Authorization](../README.md#api_key_header_Authorization)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json, application/vnd.goa.error

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## MarketplaceApiSimulateContractEvent

> SimulateContractEventResult MarketplaceApiSimulateContractEvent(ctx).SimulateContractEventRequest2(simulateContractEventRequest2).Execute()

SimulateContractEvent marketplace-api

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
	simulateContractEventRequest2 := *openapiclient.NewSimulateContractEventRequest2("purchase|plan_change|quantity_change|suspend|reinstate|cancel|release_usage_gate|stall_handoff") // SimulateContractEventRequest2 | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.MarketplaceApiAPI.MarketplaceApiSimulateContractEvent(context.Background()).SimulateContractEventRequest2(simulateContractEventRequest2).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `MarketplaceApiAPI.MarketplaceApiSimulateContractEvent``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `MarketplaceApiSimulateContractEvent`: SimulateContractEventResult
	fmt.Fprintf(os.Stdout, "Response from `MarketplaceApiAPI.MarketplaceApiSimulateContractEvent`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiMarketplaceApiSimulateContractEventRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **simulateContractEventRequest2** | [**SimulateContractEventRequest2**](SimulateContractEventRequest2.md) |  | 

### Return type

[**SimulateContractEventResult**](SimulateContractEventResult.md)

### Authorization

[api_key_header_Authorization](../README.md#api_key_header_Authorization)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json, application/vnd.goa.error

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## MarketplaceApiStartSandboxRun

> SandboxRun MarketplaceApiStartSandboxRun(ctx).StartSandboxRunRequest2(startSandboxRunRequest2).Execute()

StartSandboxRun marketplace-api

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
	startSandboxRunRequest2 := *openapiclient.NewStartSandboxRunRequest2("https://hooks.acme.example.com/omnistrate") // StartSandboxRunRequest2 | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.MarketplaceApiAPI.MarketplaceApiStartSandboxRun(context.Background()).StartSandboxRunRequest2(startSandboxRunRequest2).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `MarketplaceApiAPI.MarketplaceApiStartSandboxRun``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `MarketplaceApiStartSandboxRun`: SandboxRun
	fmt.Fprintf(os.Stdout, "Response from `MarketplaceApiAPI.MarketplaceApiStartSandboxRun`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiMarketplaceApiStartSandboxRunRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **startSandboxRunRequest2** | [**StartSandboxRunRequest2**](StartSandboxRunRequest2.md) |  | 

### Return type

[**SandboxRun**](SandboxRun.md)

### Authorization

[api_key_header_Authorization](../README.md#api_key_header_Authorization)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json, application/vnd.goa.error

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## MarketplaceApiUpdateMarketplaceChannel

> MarketplaceChannelConfig MarketplaceApiUpdateMarketplaceChannel(ctx, channel).UpdateMarketplaceChannelRequest2(updateMarketplaceChannelRequest2).Execute()

UpdateMarketplaceChannel marketplace-api

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
	channel := "SUGER|SANDBOX" // string | 
	updateMarketplaceChannelRequest2 := *openapiclient.NewUpdateMarketplaceChannelRequest2() // UpdateMarketplaceChannelRequest2 | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.MarketplaceApiAPI.MarketplaceApiUpdateMarketplaceChannel(context.Background(), channel).UpdateMarketplaceChannelRequest2(updateMarketplaceChannelRequest2).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `MarketplaceApiAPI.MarketplaceApiUpdateMarketplaceChannel``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `MarketplaceApiUpdateMarketplaceChannel`: MarketplaceChannelConfig
	fmt.Fprintf(os.Stdout, "Response from `MarketplaceApiAPI.MarketplaceApiUpdateMarketplaceChannel`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**channel** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiMarketplaceApiUpdateMarketplaceChannelRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **updateMarketplaceChannelRequest2** | [**UpdateMarketplaceChannelRequest2**](UpdateMarketplaceChannelRequest2.md) |  | 

### Return type

[**MarketplaceChannelConfig**](MarketplaceChannelConfig.md)

### Authorization

[api_key_header_Authorization](../README.md#api_key_header_Authorization)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json, application/vnd.goa.error

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## MarketplaceApiValidateMarketplaceChannel

> MarketplaceChannelValidation MarketplaceApiValidateMarketplaceChannel(ctx, channel).Execute()

ValidateMarketplaceChannel marketplace-api

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
	channel := "SUGER|SANDBOX" // string | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.MarketplaceApiAPI.MarketplaceApiValidateMarketplaceChannel(context.Background(), channel).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `MarketplaceApiAPI.MarketplaceApiValidateMarketplaceChannel``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `MarketplaceApiValidateMarketplaceChannel`: MarketplaceChannelValidation
	fmt.Fprintf(os.Stdout, "Response from `MarketplaceApiAPI.MarketplaceApiValidateMarketplaceChannel`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**channel** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiMarketplaceApiValidateMarketplaceChannelRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

[**MarketplaceChannelValidation**](MarketplaceChannelValidation.md)

### Authorization

[api_key_header_Authorization](../README.md#api_key_header_Authorization)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json, application/vnd.goa.error

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

