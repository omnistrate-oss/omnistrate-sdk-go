# EnableBillingProviderRequest

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**BalanceDueLink** | Pointer to **string** | The URL to the balance due page | [optional] 
**BillingProviderType** | **string** | The billing provider type | 
**LogoURL** | Pointer to **string** | The URL of the logo for the billing provider. Only present when the billing provider is BRING_YOUR_OWN | [optional] 
**Name** | Pointer to **string** | A custom name for the billing provider. Only present when the billing provider is BRING_YOUR_OWN | [optional] 
**Token** | **string** | JWT token used to perform authorization | 

## Methods

### NewEnableBillingProviderRequest

`func NewEnableBillingProviderRequest(billingProviderType string, token string, ) *EnableBillingProviderRequest`

NewEnableBillingProviderRequest instantiates a new EnableBillingProviderRequest object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewEnableBillingProviderRequestWithDefaults

`func NewEnableBillingProviderRequestWithDefaults() *EnableBillingProviderRequest`

NewEnableBillingProviderRequestWithDefaults instantiates a new EnableBillingProviderRequest object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetBalanceDueLink

`func (o *EnableBillingProviderRequest) GetBalanceDueLink() string`

GetBalanceDueLink returns the BalanceDueLink field if non-nil, zero value otherwise.

### GetBalanceDueLinkOk

`func (o *EnableBillingProviderRequest) GetBalanceDueLinkOk() (*string, bool)`

GetBalanceDueLinkOk returns a tuple with the BalanceDueLink field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBalanceDueLink

`func (o *EnableBillingProviderRequest) SetBalanceDueLink(v string)`

SetBalanceDueLink sets BalanceDueLink field to given value.

### HasBalanceDueLink

`func (o *EnableBillingProviderRequest) HasBalanceDueLink() bool`

HasBalanceDueLink returns a boolean if a field has been set.

### GetBillingProviderType

`func (o *EnableBillingProviderRequest) GetBillingProviderType() string`

GetBillingProviderType returns the BillingProviderType field if non-nil, zero value otherwise.

### GetBillingProviderTypeOk

`func (o *EnableBillingProviderRequest) GetBillingProviderTypeOk() (*string, bool)`

GetBillingProviderTypeOk returns a tuple with the BillingProviderType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBillingProviderType

`func (o *EnableBillingProviderRequest) SetBillingProviderType(v string)`

SetBillingProviderType sets BillingProviderType field to given value.


### GetLogoURL

`func (o *EnableBillingProviderRequest) GetLogoURL() string`

GetLogoURL returns the LogoURL field if non-nil, zero value otherwise.

### GetLogoURLOk

`func (o *EnableBillingProviderRequest) GetLogoURLOk() (*string, bool)`

GetLogoURLOk returns a tuple with the LogoURL field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLogoURL

`func (o *EnableBillingProviderRequest) SetLogoURL(v string)`

SetLogoURL sets LogoURL field to given value.

### HasLogoURL

`func (o *EnableBillingProviderRequest) HasLogoURL() bool`

HasLogoURL returns a boolean if a field has been set.

### GetName

`func (o *EnableBillingProviderRequest) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *EnableBillingProviderRequest) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *EnableBillingProviderRequest) SetName(v string)`

SetName sets Name field to given value.

### HasName

`func (o *EnableBillingProviderRequest) HasName() bool`

HasName returns a boolean if a field has been set.

### GetToken

`func (o *EnableBillingProviderRequest) GetToken() string`

GetToken returns the Token field if non-nil, zero value otherwise.

### GetTokenOk

`func (o *EnableBillingProviderRequest) GetTokenOk() (*string, bool)`

GetTokenOk returns a tuple with the Token field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetToken

`func (o *EnableBillingProviderRequest) SetToken(v string)`

SetToken sets Token field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


