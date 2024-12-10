# ListInvoicesResult

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Invoices** | Pointer to [**[]Invoice**](Invoice.md) | List of Invoices | [optional] 

## Methods

### NewListInvoicesResult

`func NewListInvoicesResult() *ListInvoicesResult`

NewListInvoicesResult instantiates a new ListInvoicesResult object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewListInvoicesResultWithDefaults

`func NewListInvoicesResultWithDefaults() *ListInvoicesResult`

NewListInvoicesResultWithDefaults instantiates a new ListInvoicesResult object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetInvoices

`func (o *ListInvoicesResult) GetInvoices() []Invoice`

GetInvoices returns the Invoices field if non-nil, zero value otherwise.

### GetInvoicesOk

`func (o *ListInvoicesResult) GetInvoicesOk() (*[]Invoice, bool)`

GetInvoicesOk returns a tuple with the Invoices field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetInvoices

`func (o *ListInvoicesResult) SetInvoices(v []Invoice)`

SetInvoices sets Invoices field to given value.

### HasInvoices

`func (o *ListInvoicesResult) HasInvoices() bool`

HasInvoices returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


