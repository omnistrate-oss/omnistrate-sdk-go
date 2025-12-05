# FleetInvoice

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**BillIssuedToUserEmail** | Pointer to **string** | Email of the user that this invoice is issued to | [optional] 
**BillIssuedToUserID** | Pointer to **string** | ID of the user that this invoice is issued to | [optional] 
**BillIssuedToUserName** | Pointer to **string** | Name of the user that this invoice is issued to | [optional] 
**Currency** | Pointer to **string** | Currency of the invoice | [optional] 
**CustomerId** | Pointer to **string** | ID of an Org | [optional] 
**CustomerName** | Pointer to **string** | Organization name of the customer | [optional] 
**DueDate** | Pointer to **string** | Due date in ISO 8601 format | [optional] 
**Id** | Pointer to **string** | ID of the invoice | [optional] 
**InvoiceDate** | Pointer to **string** | Date of the invoice in ISO 8601 format | [optional] 
**InvoicePdf** | Pointer to **string** | URL for the PDF of the invoice | [optional] 
**InvoiceUrl** | Pointer to **string** | Stripe URL for this invoice | [optional] 
**PaymentTerms** | Pointer to **string** | Payment terms of the invoice | [optional] 
**Status** | Pointer to **string** | This describes the status of the invoice and is set by the payment provider | [optional] 
**TaxAmount** | Pointer to **float64** | Amount of tax, if any | [optional] 
**TotalAmount** | Pointer to **float64** | Amount of Invoice | [optional] 
**TotalAmountWithoutTax** | Pointer to **float64** | Amount of Invoice without tax | [optional] 

## Methods

### NewFleetInvoice

`func NewFleetInvoice() *FleetInvoice`

NewFleetInvoice instantiates a new FleetInvoice object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewFleetInvoiceWithDefaults

`func NewFleetInvoiceWithDefaults() *FleetInvoice`

NewFleetInvoiceWithDefaults instantiates a new FleetInvoice object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetBillIssuedToUserEmail

`func (o *FleetInvoice) GetBillIssuedToUserEmail() string`

GetBillIssuedToUserEmail returns the BillIssuedToUserEmail field if non-nil, zero value otherwise.

### GetBillIssuedToUserEmailOk

`func (o *FleetInvoice) GetBillIssuedToUserEmailOk() (*string, bool)`

GetBillIssuedToUserEmailOk returns a tuple with the BillIssuedToUserEmail field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBillIssuedToUserEmail

`func (o *FleetInvoice) SetBillIssuedToUserEmail(v string)`

SetBillIssuedToUserEmail sets BillIssuedToUserEmail field to given value.

### HasBillIssuedToUserEmail

`func (o *FleetInvoice) HasBillIssuedToUserEmail() bool`

HasBillIssuedToUserEmail returns a boolean if a field has been set.

### GetBillIssuedToUserID

`func (o *FleetInvoice) GetBillIssuedToUserID() string`

GetBillIssuedToUserID returns the BillIssuedToUserID field if non-nil, zero value otherwise.

### GetBillIssuedToUserIDOk

`func (o *FleetInvoice) GetBillIssuedToUserIDOk() (*string, bool)`

GetBillIssuedToUserIDOk returns a tuple with the BillIssuedToUserID field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBillIssuedToUserID

`func (o *FleetInvoice) SetBillIssuedToUserID(v string)`

SetBillIssuedToUserID sets BillIssuedToUserID field to given value.

### HasBillIssuedToUserID

`func (o *FleetInvoice) HasBillIssuedToUserID() bool`

HasBillIssuedToUserID returns a boolean if a field has been set.

### GetBillIssuedToUserName

`func (o *FleetInvoice) GetBillIssuedToUserName() string`

GetBillIssuedToUserName returns the BillIssuedToUserName field if non-nil, zero value otherwise.

### GetBillIssuedToUserNameOk

`func (o *FleetInvoice) GetBillIssuedToUserNameOk() (*string, bool)`

GetBillIssuedToUserNameOk returns a tuple with the BillIssuedToUserName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBillIssuedToUserName

`func (o *FleetInvoice) SetBillIssuedToUserName(v string)`

SetBillIssuedToUserName sets BillIssuedToUserName field to given value.

### HasBillIssuedToUserName

`func (o *FleetInvoice) HasBillIssuedToUserName() bool`

HasBillIssuedToUserName returns a boolean if a field has been set.

### GetCurrency

`func (o *FleetInvoice) GetCurrency() string`

GetCurrency returns the Currency field if non-nil, zero value otherwise.

### GetCurrencyOk

`func (o *FleetInvoice) GetCurrencyOk() (*string, bool)`

GetCurrencyOk returns a tuple with the Currency field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCurrency

`func (o *FleetInvoice) SetCurrency(v string)`

SetCurrency sets Currency field to given value.

### HasCurrency

`func (o *FleetInvoice) HasCurrency() bool`

HasCurrency returns a boolean if a field has been set.

### GetCustomerId

`func (o *FleetInvoice) GetCustomerId() string`

GetCustomerId returns the CustomerId field if non-nil, zero value otherwise.

### GetCustomerIdOk

`func (o *FleetInvoice) GetCustomerIdOk() (*string, bool)`

GetCustomerIdOk returns a tuple with the CustomerId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCustomerId

`func (o *FleetInvoice) SetCustomerId(v string)`

SetCustomerId sets CustomerId field to given value.

### HasCustomerId

`func (o *FleetInvoice) HasCustomerId() bool`

HasCustomerId returns a boolean if a field has been set.

### GetCustomerName

`func (o *FleetInvoice) GetCustomerName() string`

GetCustomerName returns the CustomerName field if non-nil, zero value otherwise.

### GetCustomerNameOk

`func (o *FleetInvoice) GetCustomerNameOk() (*string, bool)`

GetCustomerNameOk returns a tuple with the CustomerName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCustomerName

`func (o *FleetInvoice) SetCustomerName(v string)`

SetCustomerName sets CustomerName field to given value.

### HasCustomerName

`func (o *FleetInvoice) HasCustomerName() bool`

HasCustomerName returns a boolean if a field has been set.

### GetDueDate

`func (o *FleetInvoice) GetDueDate() string`

GetDueDate returns the DueDate field if non-nil, zero value otherwise.

### GetDueDateOk

`func (o *FleetInvoice) GetDueDateOk() (*string, bool)`

GetDueDateOk returns a tuple with the DueDate field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDueDate

`func (o *FleetInvoice) SetDueDate(v string)`

SetDueDate sets DueDate field to given value.

### HasDueDate

`func (o *FleetInvoice) HasDueDate() bool`

HasDueDate returns a boolean if a field has been set.

### GetId

`func (o *FleetInvoice) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *FleetInvoice) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *FleetInvoice) SetId(v string)`

SetId sets Id field to given value.

### HasId

`func (o *FleetInvoice) HasId() bool`

HasId returns a boolean if a field has been set.

### GetInvoiceDate

`func (o *FleetInvoice) GetInvoiceDate() string`

GetInvoiceDate returns the InvoiceDate field if non-nil, zero value otherwise.

### GetInvoiceDateOk

`func (o *FleetInvoice) GetInvoiceDateOk() (*string, bool)`

GetInvoiceDateOk returns a tuple with the InvoiceDate field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetInvoiceDate

`func (o *FleetInvoice) SetInvoiceDate(v string)`

SetInvoiceDate sets InvoiceDate field to given value.

### HasInvoiceDate

`func (o *FleetInvoice) HasInvoiceDate() bool`

HasInvoiceDate returns a boolean if a field has been set.

### GetInvoicePdf

`func (o *FleetInvoice) GetInvoicePdf() string`

GetInvoicePdf returns the InvoicePdf field if non-nil, zero value otherwise.

### GetInvoicePdfOk

`func (o *FleetInvoice) GetInvoicePdfOk() (*string, bool)`

GetInvoicePdfOk returns a tuple with the InvoicePdf field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetInvoicePdf

`func (o *FleetInvoice) SetInvoicePdf(v string)`

SetInvoicePdf sets InvoicePdf field to given value.

### HasInvoicePdf

`func (o *FleetInvoice) HasInvoicePdf() bool`

HasInvoicePdf returns a boolean if a field has been set.

### GetInvoiceUrl

`func (o *FleetInvoice) GetInvoiceUrl() string`

GetInvoiceUrl returns the InvoiceUrl field if non-nil, zero value otherwise.

### GetInvoiceUrlOk

`func (o *FleetInvoice) GetInvoiceUrlOk() (*string, bool)`

GetInvoiceUrlOk returns a tuple with the InvoiceUrl field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetInvoiceUrl

`func (o *FleetInvoice) SetInvoiceUrl(v string)`

SetInvoiceUrl sets InvoiceUrl field to given value.

### HasInvoiceUrl

`func (o *FleetInvoice) HasInvoiceUrl() bool`

HasInvoiceUrl returns a boolean if a field has been set.

### GetPaymentTerms

`func (o *FleetInvoice) GetPaymentTerms() string`

GetPaymentTerms returns the PaymentTerms field if non-nil, zero value otherwise.

### GetPaymentTermsOk

`func (o *FleetInvoice) GetPaymentTermsOk() (*string, bool)`

GetPaymentTermsOk returns a tuple with the PaymentTerms field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPaymentTerms

`func (o *FleetInvoice) SetPaymentTerms(v string)`

SetPaymentTerms sets PaymentTerms field to given value.

### HasPaymentTerms

`func (o *FleetInvoice) HasPaymentTerms() bool`

HasPaymentTerms returns a boolean if a field has been set.

### GetStatus

`func (o *FleetInvoice) GetStatus() string`

GetStatus returns the Status field if non-nil, zero value otherwise.

### GetStatusOk

`func (o *FleetInvoice) GetStatusOk() (*string, bool)`

GetStatusOk returns a tuple with the Status field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStatus

`func (o *FleetInvoice) SetStatus(v string)`

SetStatus sets Status field to given value.

### HasStatus

`func (o *FleetInvoice) HasStatus() bool`

HasStatus returns a boolean if a field has been set.

### GetTaxAmount

`func (o *FleetInvoice) GetTaxAmount() float64`

GetTaxAmount returns the TaxAmount field if non-nil, zero value otherwise.

### GetTaxAmountOk

`func (o *FleetInvoice) GetTaxAmountOk() (*float64, bool)`

GetTaxAmountOk returns a tuple with the TaxAmount field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTaxAmount

`func (o *FleetInvoice) SetTaxAmount(v float64)`

SetTaxAmount sets TaxAmount field to given value.

### HasTaxAmount

`func (o *FleetInvoice) HasTaxAmount() bool`

HasTaxAmount returns a boolean if a field has been set.

### GetTotalAmount

`func (o *FleetInvoice) GetTotalAmount() float64`

GetTotalAmount returns the TotalAmount field if non-nil, zero value otherwise.

### GetTotalAmountOk

`func (o *FleetInvoice) GetTotalAmountOk() (*float64, bool)`

GetTotalAmountOk returns a tuple with the TotalAmount field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTotalAmount

`func (o *FleetInvoice) SetTotalAmount(v float64)`

SetTotalAmount sets TotalAmount field to given value.

### HasTotalAmount

`func (o *FleetInvoice) HasTotalAmount() bool`

HasTotalAmount returns a boolean if a field has been set.

### GetTotalAmountWithoutTax

`func (o *FleetInvoice) GetTotalAmountWithoutTax() float64`

GetTotalAmountWithoutTax returns the TotalAmountWithoutTax field if non-nil, zero value otherwise.

### GetTotalAmountWithoutTaxOk

`func (o *FleetInvoice) GetTotalAmountWithoutTaxOk() (*float64, bool)`

GetTotalAmountWithoutTaxOk returns a tuple with the TotalAmountWithoutTax field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTotalAmountWithoutTax

`func (o *FleetInvoice) SetTotalAmountWithoutTax(v float64)`

SetTotalAmountWithoutTax sets TotalAmountWithoutTax field to given value.

### HasTotalAmountWithoutTax

`func (o *FleetInvoice) HasTotalAmountWithoutTax() bool`

HasTotalAmountWithoutTax returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


