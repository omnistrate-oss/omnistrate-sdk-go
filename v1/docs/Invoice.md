# Invoice

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Currency** | Pointer to **string** | The currency of the charge | [optional] 
**DueDate** | Pointer to **string** | Due date in ISO 8601 format | [optional] 
**InvoiceDate** | Pointer to **string** | Date of the invoice | [optional] 
**InvoiceId** | Pointer to **string** | ID of the invoice | [optional] 
**InvoiceNumber** | Pointer to **string** | The invoice number | [optional] 
**InvoicePdf** | Pointer to **string** | URL for the PDF of the invoice | [optional] 
**InvoiceStatus** | Pointer to **string** | This describes the status of the invoice and is set by the payment provider | [optional] 
**InvoiceUrl** | Pointer to **string** | URL for this invoice | [optional] 
**PaymentTerms** | Pointer to **string** | Payment terms of the invoice | [optional] 
**TaxAmount** | Pointer to **float64** | Amount of tax, if any | [optional] 
**TotalAmount** | Pointer to **float64** | Amount of Invoice | [optional] 
**TotalAmountWithoutTax** | Pointer to **float64** | Amount of Invoice without tax | [optional] 

## Methods

### NewInvoice

`func NewInvoice() *Invoice`

NewInvoice instantiates a new Invoice object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewInvoiceWithDefaults

`func NewInvoiceWithDefaults() *Invoice`

NewInvoiceWithDefaults instantiates a new Invoice object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetCurrency

`func (o *Invoice) GetCurrency() string`

GetCurrency returns the Currency field if non-nil, zero value otherwise.

### GetCurrencyOk

`func (o *Invoice) GetCurrencyOk() (*string, bool)`

GetCurrencyOk returns a tuple with the Currency field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCurrency

`func (o *Invoice) SetCurrency(v string)`

SetCurrency sets Currency field to given value.

### HasCurrency

`func (o *Invoice) HasCurrency() bool`

HasCurrency returns a boolean if a field has been set.

### GetDueDate

`func (o *Invoice) GetDueDate() string`

GetDueDate returns the DueDate field if non-nil, zero value otherwise.

### GetDueDateOk

`func (o *Invoice) GetDueDateOk() (*string, bool)`

GetDueDateOk returns a tuple with the DueDate field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDueDate

`func (o *Invoice) SetDueDate(v string)`

SetDueDate sets DueDate field to given value.

### HasDueDate

`func (o *Invoice) HasDueDate() bool`

HasDueDate returns a boolean if a field has been set.

### GetInvoiceDate

`func (o *Invoice) GetInvoiceDate() string`

GetInvoiceDate returns the InvoiceDate field if non-nil, zero value otherwise.

### GetInvoiceDateOk

`func (o *Invoice) GetInvoiceDateOk() (*string, bool)`

GetInvoiceDateOk returns a tuple with the InvoiceDate field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetInvoiceDate

`func (o *Invoice) SetInvoiceDate(v string)`

SetInvoiceDate sets InvoiceDate field to given value.

### HasInvoiceDate

`func (o *Invoice) HasInvoiceDate() bool`

HasInvoiceDate returns a boolean if a field has been set.

### GetInvoiceId

`func (o *Invoice) GetInvoiceId() string`

GetInvoiceId returns the InvoiceId field if non-nil, zero value otherwise.

### GetInvoiceIdOk

`func (o *Invoice) GetInvoiceIdOk() (*string, bool)`

GetInvoiceIdOk returns a tuple with the InvoiceId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetInvoiceId

`func (o *Invoice) SetInvoiceId(v string)`

SetInvoiceId sets InvoiceId field to given value.

### HasInvoiceId

`func (o *Invoice) HasInvoiceId() bool`

HasInvoiceId returns a boolean if a field has been set.

### GetInvoiceNumber

`func (o *Invoice) GetInvoiceNumber() string`

GetInvoiceNumber returns the InvoiceNumber field if non-nil, zero value otherwise.

### GetInvoiceNumberOk

`func (o *Invoice) GetInvoiceNumberOk() (*string, bool)`

GetInvoiceNumberOk returns a tuple with the InvoiceNumber field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetInvoiceNumber

`func (o *Invoice) SetInvoiceNumber(v string)`

SetInvoiceNumber sets InvoiceNumber field to given value.

### HasInvoiceNumber

`func (o *Invoice) HasInvoiceNumber() bool`

HasInvoiceNumber returns a boolean if a field has been set.

### GetInvoicePdf

`func (o *Invoice) GetInvoicePdf() string`

GetInvoicePdf returns the InvoicePdf field if non-nil, zero value otherwise.

### GetInvoicePdfOk

`func (o *Invoice) GetInvoicePdfOk() (*string, bool)`

GetInvoicePdfOk returns a tuple with the InvoicePdf field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetInvoicePdf

`func (o *Invoice) SetInvoicePdf(v string)`

SetInvoicePdf sets InvoicePdf field to given value.

### HasInvoicePdf

`func (o *Invoice) HasInvoicePdf() bool`

HasInvoicePdf returns a boolean if a field has been set.

### GetInvoiceStatus

`func (o *Invoice) GetInvoiceStatus() string`

GetInvoiceStatus returns the InvoiceStatus field if non-nil, zero value otherwise.

### GetInvoiceStatusOk

`func (o *Invoice) GetInvoiceStatusOk() (*string, bool)`

GetInvoiceStatusOk returns a tuple with the InvoiceStatus field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetInvoiceStatus

`func (o *Invoice) SetInvoiceStatus(v string)`

SetInvoiceStatus sets InvoiceStatus field to given value.

### HasInvoiceStatus

`func (o *Invoice) HasInvoiceStatus() bool`

HasInvoiceStatus returns a boolean if a field has been set.

### GetInvoiceUrl

`func (o *Invoice) GetInvoiceUrl() string`

GetInvoiceUrl returns the InvoiceUrl field if non-nil, zero value otherwise.

### GetInvoiceUrlOk

`func (o *Invoice) GetInvoiceUrlOk() (*string, bool)`

GetInvoiceUrlOk returns a tuple with the InvoiceUrl field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetInvoiceUrl

`func (o *Invoice) SetInvoiceUrl(v string)`

SetInvoiceUrl sets InvoiceUrl field to given value.

### HasInvoiceUrl

`func (o *Invoice) HasInvoiceUrl() bool`

HasInvoiceUrl returns a boolean if a field has been set.

### GetPaymentTerms

`func (o *Invoice) GetPaymentTerms() string`

GetPaymentTerms returns the PaymentTerms field if non-nil, zero value otherwise.

### GetPaymentTermsOk

`func (o *Invoice) GetPaymentTermsOk() (*string, bool)`

GetPaymentTermsOk returns a tuple with the PaymentTerms field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPaymentTerms

`func (o *Invoice) SetPaymentTerms(v string)`

SetPaymentTerms sets PaymentTerms field to given value.

### HasPaymentTerms

`func (o *Invoice) HasPaymentTerms() bool`

HasPaymentTerms returns a boolean if a field has been set.

### GetTaxAmount

`func (o *Invoice) GetTaxAmount() float64`

GetTaxAmount returns the TaxAmount field if non-nil, zero value otherwise.

### GetTaxAmountOk

`func (o *Invoice) GetTaxAmountOk() (*float64, bool)`

GetTaxAmountOk returns a tuple with the TaxAmount field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTaxAmount

`func (o *Invoice) SetTaxAmount(v float64)`

SetTaxAmount sets TaxAmount field to given value.

### HasTaxAmount

`func (o *Invoice) HasTaxAmount() bool`

HasTaxAmount returns a boolean if a field has been set.

### GetTotalAmount

`func (o *Invoice) GetTotalAmount() float64`

GetTotalAmount returns the TotalAmount field if non-nil, zero value otherwise.

### GetTotalAmountOk

`func (o *Invoice) GetTotalAmountOk() (*float64, bool)`

GetTotalAmountOk returns a tuple with the TotalAmount field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTotalAmount

`func (o *Invoice) SetTotalAmount(v float64)`

SetTotalAmount sets TotalAmount field to given value.

### HasTotalAmount

`func (o *Invoice) HasTotalAmount() bool`

HasTotalAmount returns a boolean if a field has been set.

### GetTotalAmountWithoutTax

`func (o *Invoice) GetTotalAmountWithoutTax() float64`

GetTotalAmountWithoutTax returns the TotalAmountWithoutTax field if non-nil, zero value otherwise.

### GetTotalAmountWithoutTaxOk

`func (o *Invoice) GetTotalAmountWithoutTaxOk() (*float64, bool)`

GetTotalAmountWithoutTaxOk returns a tuple with the TotalAmountWithoutTax field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTotalAmountWithoutTax

`func (o *Invoice) SetTotalAmountWithoutTax(v float64)`

SetTotalAmountWithoutTax sets TotalAmountWithoutTax field to given value.

### HasTotalAmountWithoutTax

`func (o *Invoice) HasTotalAmountWithoutTax() bool`

HasTotalAmountWithoutTax returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


