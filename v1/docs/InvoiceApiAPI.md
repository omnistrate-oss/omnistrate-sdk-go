# \InvoiceApiAPI

All URIs are relative to *https://api.omnistrate.cloud*

Method | HTTP request | Description
------------- | ------------- | -------------
[**InvoiceApiListInvoices**](InvoiceApiAPI.md#InvoiceApiListInvoices) | **Get** /2022-09-01-00/invoice | ListInvoices invoice-api



## InvoiceApiListInvoices

> ListInvoicesResult InvoiceApiListInvoices(ctx).BillingProvider(billingProvider).Execute()

ListInvoices invoice-api

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
	billingProvider := "STRIPE" // string | Billing provider. If specified, list invoices for the specified billing provider. (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.InvoiceApiAPI.InvoiceApiListInvoices(context.Background()).BillingProvider(billingProvider).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `InvoiceApiAPI.InvoiceApiListInvoices``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `InvoiceApiListInvoices`: ListInvoicesResult
	fmt.Fprintf(os.Stdout, "Response from `InvoiceApiAPI.InvoiceApiListInvoices`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiInvoiceApiListInvoicesRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **billingProvider** | **string** | Billing provider. If specified, list invoices for the specified billing provider. | 

### Return type

[**ListInvoicesResult**](ListInvoicesResult.md)

### Authorization

[api_key_header_Authorization](../README.md#api_key_header_Authorization)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json, application/vnd.goa.error

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

