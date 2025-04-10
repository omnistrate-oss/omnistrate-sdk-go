# ListAllOrganizationsRequest

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**HasInvoice** | Pointer to **bool** | Filter for organizations with invoices | [optional] 
**HasInvoiceFromDate** | Pointer to **time.Time** | Start time of the window in RFC 3339 format. If omitted, the filter is open-ended at the start. | [optional] 
**HasInvoiceToDate** | Pointer to **time.Time** | End time of the window in RFC 3339 format. If omitted, the filter is open-ended at the start. | [optional] 
**HasInvoiceWithStatus** | Pointer to **string** | This describes the status of the invoice and is set by the payment provider | [optional] 
**Token** | **string** | JWT token used to perform authorization | 

## Methods

### NewListAllOrganizationsRequest

`func NewListAllOrganizationsRequest(token string, ) *ListAllOrganizationsRequest`

NewListAllOrganizationsRequest instantiates a new ListAllOrganizationsRequest object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewListAllOrganizationsRequestWithDefaults

`func NewListAllOrganizationsRequestWithDefaults() *ListAllOrganizationsRequest`

NewListAllOrganizationsRequestWithDefaults instantiates a new ListAllOrganizationsRequest object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetHasInvoice

`func (o *ListAllOrganizationsRequest) GetHasInvoice() bool`

GetHasInvoice returns the HasInvoice field if non-nil, zero value otherwise.

### GetHasInvoiceOk

`func (o *ListAllOrganizationsRequest) GetHasInvoiceOk() (*bool, bool)`

GetHasInvoiceOk returns a tuple with the HasInvoice field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetHasInvoice

`func (o *ListAllOrganizationsRequest) SetHasInvoice(v bool)`

SetHasInvoice sets HasInvoice field to given value.

### HasHasInvoice

`func (o *ListAllOrganizationsRequest) HasHasInvoice() bool`

HasHasInvoice returns a boolean if a field has been set.

### GetHasInvoiceFromDate

`func (o *ListAllOrganizationsRequest) GetHasInvoiceFromDate() time.Time`

GetHasInvoiceFromDate returns the HasInvoiceFromDate field if non-nil, zero value otherwise.

### GetHasInvoiceFromDateOk

`func (o *ListAllOrganizationsRequest) GetHasInvoiceFromDateOk() (*time.Time, bool)`

GetHasInvoiceFromDateOk returns a tuple with the HasInvoiceFromDate field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetHasInvoiceFromDate

`func (o *ListAllOrganizationsRequest) SetHasInvoiceFromDate(v time.Time)`

SetHasInvoiceFromDate sets HasInvoiceFromDate field to given value.

### HasHasInvoiceFromDate

`func (o *ListAllOrganizationsRequest) HasHasInvoiceFromDate() bool`

HasHasInvoiceFromDate returns a boolean if a field has been set.

### GetHasInvoiceToDate

`func (o *ListAllOrganizationsRequest) GetHasInvoiceToDate() time.Time`

GetHasInvoiceToDate returns the HasInvoiceToDate field if non-nil, zero value otherwise.

### GetHasInvoiceToDateOk

`func (o *ListAllOrganizationsRequest) GetHasInvoiceToDateOk() (*time.Time, bool)`

GetHasInvoiceToDateOk returns a tuple with the HasInvoiceToDate field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetHasInvoiceToDate

`func (o *ListAllOrganizationsRequest) SetHasInvoiceToDate(v time.Time)`

SetHasInvoiceToDate sets HasInvoiceToDate field to given value.

### HasHasInvoiceToDate

`func (o *ListAllOrganizationsRequest) HasHasInvoiceToDate() bool`

HasHasInvoiceToDate returns a boolean if a field has been set.

### GetHasInvoiceWithStatus

`func (o *ListAllOrganizationsRequest) GetHasInvoiceWithStatus() string`

GetHasInvoiceWithStatus returns the HasInvoiceWithStatus field if non-nil, zero value otherwise.

### GetHasInvoiceWithStatusOk

`func (o *ListAllOrganizationsRequest) GetHasInvoiceWithStatusOk() (*string, bool)`

GetHasInvoiceWithStatusOk returns a tuple with the HasInvoiceWithStatus field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetHasInvoiceWithStatus

`func (o *ListAllOrganizationsRequest) SetHasInvoiceWithStatus(v string)`

SetHasInvoiceWithStatus sets HasInvoiceWithStatus field to given value.

### HasHasInvoiceWithStatus

`func (o *ListAllOrganizationsRequest) HasHasInvoiceWithStatus() bool`

HasHasInvoiceWithStatus returns a boolean if a field has been set.

### GetToken

`func (o *ListAllOrganizationsRequest) GetToken() string`

GetToken returns the Token field if non-nil, zero value otherwise.

### GetTokenOk

`func (o *ListAllOrganizationsRequest) GetTokenOk() (*string, bool)`

GetTokenOk returns a tuple with the Token field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetToken

`func (o *ListAllOrganizationsRequest) SetToken(v string)`

SetToken sets Token field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


