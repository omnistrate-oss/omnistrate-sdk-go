# \InvoiceApiAPI

All URIs are relative to *https://api.omnistrate.cloud*

Method | HTTP request | Description
------------- | ------------- | -------------
[**InvoiceApiListInvoices**](InvoiceApiAPI.md#InvoiceApiListInvoices) | **Get** /2022-09-01-00/invoice | ListInvoices invoice-api



## InvoiceApiListInvoices

> ListInvoicesResult InvoiceApiListInvoices(ctx).Execute()

ListInvoices invoice-api

### Example

```go
package main

import (
	"context"
	"fmt"
	"os"
	openapiclient "github.com/omnistrate/omnistrate-sdk-go"
)

func main() {

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.InvoiceApiAPI.InvoiceApiListInvoices(context.Background()).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `InvoiceApiAPI.InvoiceApiListInvoices``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `InvoiceApiListInvoices`: ListInvoicesResult
	fmt.Fprintf(os.Stdout, "Response from `InvoiceApiAPI.InvoiceApiListInvoices`: %v\n", resp)
}
```

### Path Parameters

This endpoint does not need any parameter.

### Other Parameters

Other parameters are passed through a pointer to a apiInvoiceApiListInvoicesRequest struct via the builder pattern


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

