# ListInvoicesRequest

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**BillingProvider** | Pointer to **string** | The billing provider type | [optional] 
**Token** | **string** | JWT token used to perform authorization | 

## Methods

### NewListInvoicesRequest

`func NewListInvoicesRequest(token string, ) *ListInvoicesRequest`

NewListInvoicesRequest instantiates a new ListInvoicesRequest object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewListInvoicesRequestWithDefaults

`func NewListInvoicesRequestWithDefaults() *ListInvoicesRequest`

NewListInvoicesRequestWithDefaults instantiates a new ListInvoicesRequest object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetBillingProvider

`func (o *ListInvoicesRequest) GetBillingProvider() string`

GetBillingProvider returns the BillingProvider field if non-nil, zero value otherwise.

### GetBillingProviderOk

`func (o *ListInvoicesRequest) GetBillingProviderOk() (*string, bool)`

GetBillingProviderOk returns a tuple with the BillingProvider field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBillingProvider

`func (o *ListInvoicesRequest) SetBillingProvider(v string)`

SetBillingProvider sets BillingProvider field to given value.

### HasBillingProvider

`func (o *ListInvoicesRequest) HasBillingProvider() bool`

HasBillingProvider returns a boolean if a field has been set.

### GetToken

`func (o *ListInvoicesRequest) GetToken() string`

GetToken returns the Token field if non-nil, zero value otherwise.

### GetTokenOk

`func (o *ListInvoicesRequest) GetTokenOk() (*string, bool)`

GetTokenOk returns a tuple with the Token field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetToken

`func (o *ListInvoicesRequest) SetToken(v string)`

SetToken sets Token field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


