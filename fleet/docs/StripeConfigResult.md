# StripeConfigResult

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**PublishableKey** | **string** | Stripe publishable key for mounting Elements | 
**StripeAccountId** | Pointer to **string** | Connected Stripe account ID (SP-tier only) | [optional] 

## Methods

### NewStripeConfigResult

`func NewStripeConfigResult(publishableKey string, ) *StripeConfigResult`

NewStripeConfigResult instantiates a new StripeConfigResult object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewStripeConfigResultWithDefaults

`func NewStripeConfigResultWithDefaults() *StripeConfigResult`

NewStripeConfigResultWithDefaults instantiates a new StripeConfigResult object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetPublishableKey

`func (o *StripeConfigResult) GetPublishableKey() string`

GetPublishableKey returns the PublishableKey field if non-nil, zero value otherwise.

### GetPublishableKeyOk

`func (o *StripeConfigResult) GetPublishableKeyOk() (*string, bool)`

GetPublishableKeyOk returns a tuple with the PublishableKey field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPublishableKey

`func (o *StripeConfigResult) SetPublishableKey(v string)`

SetPublishableKey sets PublishableKey field to given value.


### GetStripeAccountId

`func (o *StripeConfigResult) GetStripeAccountId() string`

GetStripeAccountId returns the StripeAccountId field if non-nil, zero value otherwise.

### GetStripeAccountIdOk

`func (o *StripeConfigResult) GetStripeAccountIdOk() (*string, bool)`

GetStripeAccountIdOk returns a tuple with the StripeAccountId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStripeAccountId

`func (o *StripeConfigResult) SetStripeAccountId(v string)`

SetStripeAccountId sets StripeAccountId field to given value.

### HasStripeAccountId

`func (o *StripeConfigResult) HasStripeAccountId() bool`

HasStripeAccountId returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


