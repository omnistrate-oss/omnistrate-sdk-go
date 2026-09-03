# MarketplaceChannelWebhookResult

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Accepted** | **bool** | True once the delivery has been verified and recorded. Not a statement that anything changed: most deliveries resolve to a re-read that finds the state already correct | 

## Methods

### NewMarketplaceChannelWebhookResult

`func NewMarketplaceChannelWebhookResult(accepted bool, ) *MarketplaceChannelWebhookResult`

NewMarketplaceChannelWebhookResult instantiates a new MarketplaceChannelWebhookResult object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewMarketplaceChannelWebhookResultWithDefaults

`func NewMarketplaceChannelWebhookResultWithDefaults() *MarketplaceChannelWebhookResult`

NewMarketplaceChannelWebhookResultWithDefaults instantiates a new MarketplaceChannelWebhookResult object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetAccepted

`func (o *MarketplaceChannelWebhookResult) GetAccepted() bool`

GetAccepted returns the Accepted field if non-nil, zero value otherwise.

### GetAcceptedOk

`func (o *MarketplaceChannelWebhookResult) GetAcceptedOk() (*bool, bool)`

GetAcceptedOk returns a tuple with the Accepted field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAccepted

`func (o *MarketplaceChannelWebhookResult) SetAccepted(v bool)`

SetAccepted sets Accepted field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


