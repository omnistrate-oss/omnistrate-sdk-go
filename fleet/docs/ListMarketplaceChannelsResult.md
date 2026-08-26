# ListMarketplaceChannelsResult

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Channels** | [**[]MarketplaceChannelConfig**](MarketplaceChannelConfig.md) | Includes channels with no configuration yet, reported NOT_CONNECTED. A list of only the configured ones could not answer \&quot;what could I sell through\&quot;, which is the question an ISV opens this page with | 

## Methods

### NewListMarketplaceChannelsResult

`func NewListMarketplaceChannelsResult(channels []MarketplaceChannelConfig, ) *ListMarketplaceChannelsResult`

NewListMarketplaceChannelsResult instantiates a new ListMarketplaceChannelsResult object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewListMarketplaceChannelsResultWithDefaults

`func NewListMarketplaceChannelsResultWithDefaults() *ListMarketplaceChannelsResult`

NewListMarketplaceChannelsResultWithDefaults instantiates a new ListMarketplaceChannelsResult object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetChannels

`func (o *ListMarketplaceChannelsResult) GetChannels() []MarketplaceChannelConfig`

GetChannels returns the Channels field if non-nil, zero value otherwise.

### GetChannelsOk

`func (o *ListMarketplaceChannelsResult) GetChannelsOk() (*[]MarketplaceChannelConfig, bool)`

GetChannelsOk returns a tuple with the Channels field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetChannels

`func (o *ListMarketplaceChannelsResult) SetChannels(v []MarketplaceChannelConfig)`

SetChannels sets Channels field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


