# GetTenantBillingStatusResult

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**BillingProviders** | Pointer to [**[]BillingProvider**](BillingProvider.md) | List of billing providers | [optional] 
**Enabled** | Pointer to **bool** | Whether tenant billing is enabled for the organization | [optional] 

## Methods

### NewGetTenantBillingStatusResult

`func NewGetTenantBillingStatusResult() *GetTenantBillingStatusResult`

NewGetTenantBillingStatusResult instantiates a new GetTenantBillingStatusResult object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewGetTenantBillingStatusResultWithDefaults

`func NewGetTenantBillingStatusResultWithDefaults() *GetTenantBillingStatusResult`

NewGetTenantBillingStatusResultWithDefaults instantiates a new GetTenantBillingStatusResult object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetBillingProviders

`func (o *GetTenantBillingStatusResult) GetBillingProviders() []BillingProvider`

GetBillingProviders returns the BillingProviders field if non-nil, zero value otherwise.

### GetBillingProvidersOk

`func (o *GetTenantBillingStatusResult) GetBillingProvidersOk() (*[]BillingProvider, bool)`

GetBillingProvidersOk returns a tuple with the BillingProviders field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBillingProviders

`func (o *GetTenantBillingStatusResult) SetBillingProviders(v []BillingProvider)`

SetBillingProviders sets BillingProviders field to given value.

### HasBillingProviders

`func (o *GetTenantBillingStatusResult) HasBillingProviders() bool`

HasBillingProviders returns a boolean if a field has been set.

### GetEnabled

`func (o *GetTenantBillingStatusResult) GetEnabled() bool`

GetEnabled returns the Enabled field if non-nil, zero value otherwise.

### GetEnabledOk

`func (o *GetTenantBillingStatusResult) GetEnabledOk() (*bool, bool)`

GetEnabledOk returns a tuple with the Enabled field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEnabled

`func (o *GetTenantBillingStatusResult) SetEnabled(v bool)`

SetEnabled sets Enabled field to given value.

### HasEnabled

`func (o *GetTenantBillingStatusResult) HasEnabled() bool`

HasEnabled returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


