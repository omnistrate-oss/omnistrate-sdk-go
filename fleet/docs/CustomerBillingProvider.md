# CustomerBillingProvider

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**BalanceDueLink** | Pointer to **string** | The URL to the balance due page. Only present when the payment channel is BRING_YOUR_OWN | [optional] 
**LogoURL** | Pointer to **string** | The URL of the logo for the payment channel. Only present when the payment channel is BRING_YOUR_OWN | [optional] 
**Name** | Pointer to **string** | A custom name for the payment channel | [optional] 
**PaymentConfigurable** | Pointer to **bool** | Whether the payment channel payment information can be configured. | [optional] 
**PaymentConfigured** | Pointer to **bool** | Whether the customer has configured their payment information. | [optional] 
**PaymentInfoPortalURL** | Pointer to **string** | The URL from the billing provide to redirect users to so they can enter their payment information.  Only present when first adding payment information | [optional] 
**Type** | **string** | The billing provider type | 

## Methods

### NewCustomerBillingProvider

`func NewCustomerBillingProvider(type_ string, ) *CustomerBillingProvider`

NewCustomerBillingProvider instantiates a new CustomerBillingProvider object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewCustomerBillingProviderWithDefaults

`func NewCustomerBillingProviderWithDefaults() *CustomerBillingProvider`

NewCustomerBillingProviderWithDefaults instantiates a new CustomerBillingProvider object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetBalanceDueLink

`func (o *CustomerBillingProvider) GetBalanceDueLink() string`

GetBalanceDueLink returns the BalanceDueLink field if non-nil, zero value otherwise.

### GetBalanceDueLinkOk

`func (o *CustomerBillingProvider) GetBalanceDueLinkOk() (*string, bool)`

GetBalanceDueLinkOk returns a tuple with the BalanceDueLink field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBalanceDueLink

`func (o *CustomerBillingProvider) SetBalanceDueLink(v string)`

SetBalanceDueLink sets BalanceDueLink field to given value.

### HasBalanceDueLink

`func (o *CustomerBillingProvider) HasBalanceDueLink() bool`

HasBalanceDueLink returns a boolean if a field has been set.

### GetLogoURL

`func (o *CustomerBillingProvider) GetLogoURL() string`

GetLogoURL returns the LogoURL field if non-nil, zero value otherwise.

### GetLogoURLOk

`func (o *CustomerBillingProvider) GetLogoURLOk() (*string, bool)`

GetLogoURLOk returns a tuple with the LogoURL field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLogoURL

`func (o *CustomerBillingProvider) SetLogoURL(v string)`

SetLogoURL sets LogoURL field to given value.

### HasLogoURL

`func (o *CustomerBillingProvider) HasLogoURL() bool`

HasLogoURL returns a boolean if a field has been set.

### GetName

`func (o *CustomerBillingProvider) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *CustomerBillingProvider) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *CustomerBillingProvider) SetName(v string)`

SetName sets Name field to given value.

### HasName

`func (o *CustomerBillingProvider) HasName() bool`

HasName returns a boolean if a field has been set.

### GetPaymentConfigurable

`func (o *CustomerBillingProvider) GetPaymentConfigurable() bool`

GetPaymentConfigurable returns the PaymentConfigurable field if non-nil, zero value otherwise.

### GetPaymentConfigurableOk

`func (o *CustomerBillingProvider) GetPaymentConfigurableOk() (*bool, bool)`

GetPaymentConfigurableOk returns a tuple with the PaymentConfigurable field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPaymentConfigurable

`func (o *CustomerBillingProvider) SetPaymentConfigurable(v bool)`

SetPaymentConfigurable sets PaymentConfigurable field to given value.

### HasPaymentConfigurable

`func (o *CustomerBillingProvider) HasPaymentConfigurable() bool`

HasPaymentConfigurable returns a boolean if a field has been set.

### GetPaymentConfigured

`func (o *CustomerBillingProvider) GetPaymentConfigured() bool`

GetPaymentConfigured returns the PaymentConfigured field if non-nil, zero value otherwise.

### GetPaymentConfiguredOk

`func (o *CustomerBillingProvider) GetPaymentConfiguredOk() (*bool, bool)`

GetPaymentConfiguredOk returns a tuple with the PaymentConfigured field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPaymentConfigured

`func (o *CustomerBillingProvider) SetPaymentConfigured(v bool)`

SetPaymentConfigured sets PaymentConfigured field to given value.

### HasPaymentConfigured

`func (o *CustomerBillingProvider) HasPaymentConfigured() bool`

HasPaymentConfigured returns a boolean if a field has been set.

### GetPaymentInfoPortalURL

`func (o *CustomerBillingProvider) GetPaymentInfoPortalURL() string`

GetPaymentInfoPortalURL returns the PaymentInfoPortalURL field if non-nil, zero value otherwise.

### GetPaymentInfoPortalURLOk

`func (o *CustomerBillingProvider) GetPaymentInfoPortalURLOk() (*string, bool)`

GetPaymentInfoPortalURLOk returns a tuple with the PaymentInfoPortalURL field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPaymentInfoPortalURL

`func (o *CustomerBillingProvider) SetPaymentInfoPortalURL(v string)`

SetPaymentInfoPortalURL sets PaymentInfoPortalURL field to given value.

### HasPaymentInfoPortalURL

`func (o *CustomerBillingProvider) HasPaymentInfoPortalURL() bool`

HasPaymentInfoPortalURL returns a boolean if a field has been set.

### GetType

`func (o *CustomerBillingProvider) GetType() string`

GetType returns the Type field if non-nil, zero value otherwise.

### GetTypeOk

`func (o *CustomerBillingProvider) GetTypeOk() (*string, bool)`

GetTypeOk returns a tuple with the Type field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetType

`func (o *CustomerBillingProvider) SetType(v string)`

SetType sets Type field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


