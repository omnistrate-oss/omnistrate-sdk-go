# Invoice

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Currency** | Pointer to **string** | The currency of the charge | [optional] 
**InvoiceDate** | Pointer to **string** | Date of the invoice | [optional] 
**InvoiceId** | Pointer to **string** | ID of the invoice | [optional] 
**InvoicePaymentTerm** | Pointer to **string** | Invoice Payment Term | [optional] 
**InvoiceStatus** | Pointer to **string** | This describes the status of the invoice and is set by the payment provider | [optional] 
**InvoiceUrl** | Pointer to **string** | URL for this invoice | [optional] 
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

### GetInvoicePaymentTerm

`func (o *Invoice) GetInvoicePaymentTerm() string`

GetInvoicePaymentTerm returns the InvoicePaymentTerm field if non-nil, zero value otherwise.

### GetInvoicePaymentTermOk

`func (o *Invoice) GetInvoicePaymentTermOk() (*string, bool)`

GetInvoicePaymentTermOk returns a tuple with the InvoicePaymentTerm field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetInvoicePaymentTerm

`func (o *Invoice) SetInvoicePaymentTerm(v string)`

SetInvoicePaymentTerm sets InvoicePaymentTerm field to given value.

### HasInvoicePaymentTerm

`func (o *Invoice) HasInvoicePaymentTerm() bool`

HasInvoicePaymentTerm returns a boolean if a field has been set.

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


