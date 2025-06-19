# DisableBillingProviderRequest

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**BillingProviderType** | **string** | The billing provider type | 
**Token** | **string** | JWT token used to perform authorization | 

## Methods

### NewDisableBillingProviderRequest

`func NewDisableBillingProviderRequest(billingProviderType string, token string, ) *DisableBillingProviderRequest`

NewDisableBillingProviderRequest instantiates a new DisableBillingProviderRequest object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewDisableBillingProviderRequestWithDefaults

`func NewDisableBillingProviderRequestWithDefaults() *DisableBillingProviderRequest`

NewDisableBillingProviderRequestWithDefaults instantiates a new DisableBillingProviderRequest object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetBillingProviderType

`func (o *DisableBillingProviderRequest) GetBillingProviderType() string`

GetBillingProviderType returns the BillingProviderType field if non-nil, zero value otherwise.

### GetBillingProviderTypeOk

`func (o *DisableBillingProviderRequest) GetBillingProviderTypeOk() (*string, bool)`

GetBillingProviderTypeOk returns a tuple with the BillingProviderType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBillingProviderType

`func (o *DisableBillingProviderRequest) SetBillingProviderType(v string)`

SetBillingProviderType sets BillingProviderType field to given value.


### GetToken

`func (o *DisableBillingProviderRequest) GetToken() string`

GetToken returns the Token field if non-nil, zero value otherwise.

### GetTokenOk

`func (o *DisableBillingProviderRequest) GetTokenOk() (*string, bool)`

GetTokenOk returns a tuple with the Token field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetToken

`func (o *DisableBillingProviderRequest) SetToken(v string)`

SetToken sets Token field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


