# EnableBillingProviderRequest2

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**BalanceDueLink** | Pointer to **string** | The URL to the balance due page | [optional] 
**LogoURL** | Pointer to **string** | The URL of the logo for the billing provider. Only present when the billing provider is BRING_YOUR_OWN | [optional] 
**Name** | Pointer to **string** | A custom name for the billing provider. Only present when the billing provider is BRING_YOUR_OWN | [optional] 

## Methods

### NewEnableBillingProviderRequest2

`func NewEnableBillingProviderRequest2() *EnableBillingProviderRequest2`

NewEnableBillingProviderRequest2 instantiates a new EnableBillingProviderRequest2 object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewEnableBillingProviderRequest2WithDefaults

`func NewEnableBillingProviderRequest2WithDefaults() *EnableBillingProviderRequest2`

NewEnableBillingProviderRequest2WithDefaults instantiates a new EnableBillingProviderRequest2 object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetBalanceDueLink

`func (o *EnableBillingProviderRequest2) GetBalanceDueLink() string`

GetBalanceDueLink returns the BalanceDueLink field if non-nil, zero value otherwise.

### GetBalanceDueLinkOk

`func (o *EnableBillingProviderRequest2) GetBalanceDueLinkOk() (*string, bool)`

GetBalanceDueLinkOk returns a tuple with the BalanceDueLink field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBalanceDueLink

`func (o *EnableBillingProviderRequest2) SetBalanceDueLink(v string)`

SetBalanceDueLink sets BalanceDueLink field to given value.

### HasBalanceDueLink

`func (o *EnableBillingProviderRequest2) HasBalanceDueLink() bool`

HasBalanceDueLink returns a boolean if a field has been set.

### GetLogoURL

`func (o *EnableBillingProviderRequest2) GetLogoURL() string`

GetLogoURL returns the LogoURL field if non-nil, zero value otherwise.

### GetLogoURLOk

`func (o *EnableBillingProviderRequest2) GetLogoURLOk() (*string, bool)`

GetLogoURLOk returns a tuple with the LogoURL field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLogoURL

`func (o *EnableBillingProviderRequest2) SetLogoURL(v string)`

SetLogoURL sets LogoURL field to given value.

### HasLogoURL

`func (o *EnableBillingProviderRequest2) HasLogoURL() bool`

HasLogoURL returns a boolean if a field has been set.

### GetName

`func (o *EnableBillingProviderRequest2) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *EnableBillingProviderRequest2) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *EnableBillingProviderRequest2) SetName(v string)`

SetName sets Name field to given value.

### HasName

`func (o *EnableBillingProviderRequest2) HasName() bool`

HasName returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


