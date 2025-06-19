# BillingProvider

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**BalanceDueLink** | Pointer to **string** | The URL to the balance due page | [optional] 
**ConnectClientID** | Pointer to **string** | The Stripe Connect Client ID associated with the billing provider. Only present when the billing provider is STRIPE | [optional] 
**LogoURL** | Pointer to **string** | The URL of the logo for the billing provider. Only present when the billing provider is BRING_YOUR_OWN | [optional] 
**Name** | Pointer to **string** | A custom name for the billing provider. Only present when the billing provider is BRING_YOUR_OWN | [optional] 
**Status** | **string** | The status of the billing provider | 
**StripeCustomerID** | Pointer to **string** | The Stripe Customer ID associated with the billing provider. Only present when the billing provider is STRIPE | [optional] 
**Type** | **string** | The billing provider type | 

## Methods

### NewBillingProvider

`func NewBillingProvider(status string, type_ string, ) *BillingProvider`

NewBillingProvider instantiates a new BillingProvider object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewBillingProviderWithDefaults

`func NewBillingProviderWithDefaults() *BillingProvider`

NewBillingProviderWithDefaults instantiates a new BillingProvider object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetBalanceDueLink

`func (o *BillingProvider) GetBalanceDueLink() string`

GetBalanceDueLink returns the BalanceDueLink field if non-nil, zero value otherwise.

### GetBalanceDueLinkOk

`func (o *BillingProvider) GetBalanceDueLinkOk() (*string, bool)`

GetBalanceDueLinkOk returns a tuple with the BalanceDueLink field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBalanceDueLink

`func (o *BillingProvider) SetBalanceDueLink(v string)`

SetBalanceDueLink sets BalanceDueLink field to given value.

### HasBalanceDueLink

`func (o *BillingProvider) HasBalanceDueLink() bool`

HasBalanceDueLink returns a boolean if a field has been set.

### GetConnectClientID

`func (o *BillingProvider) GetConnectClientID() string`

GetConnectClientID returns the ConnectClientID field if non-nil, zero value otherwise.

### GetConnectClientIDOk

`func (o *BillingProvider) GetConnectClientIDOk() (*string, bool)`

GetConnectClientIDOk returns a tuple with the ConnectClientID field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetConnectClientID

`func (o *BillingProvider) SetConnectClientID(v string)`

SetConnectClientID sets ConnectClientID field to given value.

### HasConnectClientID

`func (o *BillingProvider) HasConnectClientID() bool`

HasConnectClientID returns a boolean if a field has been set.

### GetLogoURL

`func (o *BillingProvider) GetLogoURL() string`

GetLogoURL returns the LogoURL field if non-nil, zero value otherwise.

### GetLogoURLOk

`func (o *BillingProvider) GetLogoURLOk() (*string, bool)`

GetLogoURLOk returns a tuple with the LogoURL field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLogoURL

`func (o *BillingProvider) SetLogoURL(v string)`

SetLogoURL sets LogoURL field to given value.

### HasLogoURL

`func (o *BillingProvider) HasLogoURL() bool`

HasLogoURL returns a boolean if a field has been set.

### GetName

`func (o *BillingProvider) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *BillingProvider) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *BillingProvider) SetName(v string)`

SetName sets Name field to given value.

### HasName

`func (o *BillingProvider) HasName() bool`

HasName returns a boolean if a field has been set.

### GetStatus

`func (o *BillingProvider) GetStatus() string`

GetStatus returns the Status field if non-nil, zero value otherwise.

### GetStatusOk

`func (o *BillingProvider) GetStatusOk() (*string, bool)`

GetStatusOk returns a tuple with the Status field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStatus

`func (o *BillingProvider) SetStatus(v string)`

SetStatus sets Status field to given value.


### GetStripeCustomerID

`func (o *BillingProvider) GetStripeCustomerID() string`

GetStripeCustomerID returns the StripeCustomerID field if non-nil, zero value otherwise.

### GetStripeCustomerIDOk

`func (o *BillingProvider) GetStripeCustomerIDOk() (*string, bool)`

GetStripeCustomerIDOk returns a tuple with the StripeCustomerID field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStripeCustomerID

`func (o *BillingProvider) SetStripeCustomerID(v string)`

SetStripeCustomerID sets StripeCustomerID field to given value.

### HasStripeCustomerID

`func (o *BillingProvider) HasStripeCustomerID() bool`

HasStripeCustomerID returns a boolean if a field has been set.

### GetType

`func (o *BillingProvider) GetType() string`

GetType returns the Type field if non-nil, zero value otherwise.

### GetTypeOk

`func (o *BillingProvider) GetTypeOk() (*string, bool)`

GetTypeOk returns a tuple with the Type field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetType

`func (o *BillingProvider) SetType(v string)`

SetType sets Type field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


