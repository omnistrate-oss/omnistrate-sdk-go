# \InvoiceApiAPI

All URIs are relative to *https://api.omnistrate.cloud*

Method | HTTP request | Description
------------- | ------------- | -------------
[**InvoiceApiApproveInvoice**](InvoiceApiAPI.md#InvoiceApiApproveInvoice) | **Post** /2022-09-01-00/fleet/invoice/{id}/approve | ApproveInvoice invoice-api
[**InvoiceApiDeleteInvoice**](InvoiceApiAPI.md#InvoiceApiDeleteInvoice) | **Delete** /2022-09-01-00/fleet/invoice/{id} | DeleteInvoice invoice-api
[**InvoiceApiDescribeInvoice**](InvoiceApiAPI.md#InvoiceApiDescribeInvoice) | **Get** /2022-09-01-00/fleet/invoices/{id} | DescribeInvoice invoice-api
[**InvoiceApiListInvoices**](InvoiceApiAPI.md#InvoiceApiListInvoices) | **Get** /2022-09-01-00/fleet/invoices | ListInvoices invoice-api
[**InvoiceApiResendInvoice**](InvoiceApiAPI.md#InvoiceApiResendInvoice) | **Post** /2022-09-01-00/fleet/invoice/{id}/resend | ResendInvoice invoice-api
[**InvoiceApiVoidInvoice**](InvoiceApiAPI.md#InvoiceApiVoidInvoice) | **Post** /2022-09-01-00/fleet/invoice/{id}/void | VoidInvoice invoice-api



## InvoiceApiApproveInvoice

> InvoiceApiApproveInvoice(ctx, id).Execute()

ApproveInvoice invoice-api

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
	id := "invoice-1234568" // string | ID of the invoice

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	r, err := apiClient.InvoiceApiAPI.InvoiceApiApproveInvoice(context.Background(), id).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `InvoiceApiAPI.InvoiceApiApproveInvoice``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **string** | ID of the invoice | 

### Other Parameters

Other parameters are passed through a pointer to a apiInvoiceApiApproveInvoiceRequest struct via the builder pattern


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


## InvoiceApiDeleteInvoice

> InvoiceApiDeleteInvoice(ctx, id).Execute()

DeleteInvoice invoice-api

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
	id := "invoice-1234568" // string | ID of the invoice

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	r, err := apiClient.InvoiceApiAPI.InvoiceApiDeleteInvoice(context.Background(), id).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `InvoiceApiAPI.InvoiceApiDeleteInvoice``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **string** | ID of the invoice | 

### Other Parameters

Other parameters are passed through a pointer to a apiInvoiceApiDeleteInvoiceRequest struct via the builder pattern


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


## InvoiceApiDescribeInvoice

> FleetInvoice InvoiceApiDescribeInvoice(ctx, id).Execute()

DescribeInvoice invoice-api

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
	id := "invoice-1234568" // string | ID of the invoice

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.InvoiceApiAPI.InvoiceApiDescribeInvoice(context.Background(), id).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `InvoiceApiAPI.InvoiceApiDescribeInvoice``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `InvoiceApiDescribeInvoice`: FleetInvoice
	fmt.Fprintf(os.Stdout, "Response from `InvoiceApiAPI.InvoiceApiDescribeInvoice`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **string** | ID of the invoice | 

### Other Parameters

Other parameters are passed through a pointer to a apiInvoiceApiDescribeInvoiceRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

[**FleetInvoice**](FleetInvoice.md)

### Authorization

[api_key_header_Authorization](../README.md#api_key_header_Authorization)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json, application/vnd.goa.error

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## InvoiceApiListInvoices

> FleetListInvoicesResult InvoiceApiListInvoices(ctx).StartDate(startDate).EndDate(endDate).Status(status).CustomerId(customerId).BillingProvider(billingProvider).Execute()

ListInvoices invoice-api

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
	startDate := time.Now() // time.Time | Start date for filtering invoices (optional)
	endDate := time.Now() // time.Time | End date for filtering invoices (optional)
	status := "open" // string | Filter by invoice status (optional)
	customerId := "customer-1234568" // string | Filter by customer ID (optional)
	billingProvider := "STRIPE" // string | Filter by billing provider (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.InvoiceApiAPI.InvoiceApiListInvoices(context.Background()).StartDate(startDate).EndDate(endDate).Status(status).CustomerId(customerId).BillingProvider(billingProvider).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `InvoiceApiAPI.InvoiceApiListInvoices``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `InvoiceApiListInvoices`: FleetListInvoicesResult
	fmt.Fprintf(os.Stdout, "Response from `InvoiceApiAPI.InvoiceApiListInvoices`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiInvoiceApiListInvoicesRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **startDate** | **time.Time** | Start date for filtering invoices | 
 **endDate** | **time.Time** | End date for filtering invoices | 
 **status** | **string** | Filter by invoice status | 
 **customerId** | **string** | Filter by customer ID | 
 **billingProvider** | **string** | Filter by billing provider | 

### Return type

[**FleetListInvoicesResult**](FleetListInvoicesResult.md)

### Authorization

[api_key_header_Authorization](../README.md#api_key_header_Authorization)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json, application/vnd.goa.error

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## InvoiceApiResendInvoice

> InvoiceApiResendInvoice(ctx, id).Execute()

ResendInvoice invoice-api

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
	id := "invoice-1234568" // string | ID of the invoice

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	r, err := apiClient.InvoiceApiAPI.InvoiceApiResendInvoice(context.Background(), id).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `InvoiceApiAPI.InvoiceApiResendInvoice``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **string** | ID of the invoice | 

### Other Parameters

Other parameters are passed through a pointer to a apiInvoiceApiResendInvoiceRequest struct via the builder pattern


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


## InvoiceApiVoidInvoice

> InvoiceApiVoidInvoice(ctx, id).Execute()

VoidInvoice invoice-api

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
	id := "invoice-1234568" // string | ID of the invoice

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	r, err := apiClient.InvoiceApiAPI.InvoiceApiVoidInvoice(context.Background(), id).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `InvoiceApiAPI.InvoiceApiVoidInvoice``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **string** | ID of the invoice | 

### Other Parameters

Other parameters are passed through a pointer to a apiInvoiceApiVoidInvoiceRequest struct via the builder pattern


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

