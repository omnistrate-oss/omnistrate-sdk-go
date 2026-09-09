# MarketplaceEventReceiver

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**EventType** | **string** | The wire event name this routes, for example contract.cancelled. The published name rather than an internal one, so it does not move when Omnistrate refactors | 
**ReceiverUrl** | **string** | Where deliveries of this event go. https, a host, no query, no fragment, and not a private, loopback, link-local or metadata address. Required: an entry naming no URL would be an entry that says nothing, and removing the entry is how you stop delivering | 

## Methods

### NewMarketplaceEventReceiver

`func NewMarketplaceEventReceiver(eventType string, receiverUrl string, ) *MarketplaceEventReceiver`

NewMarketplaceEventReceiver instantiates a new MarketplaceEventReceiver object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewMarketplaceEventReceiverWithDefaults

`func NewMarketplaceEventReceiverWithDefaults() *MarketplaceEventReceiver`

NewMarketplaceEventReceiverWithDefaults instantiates a new MarketplaceEventReceiver object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetEventType

`func (o *MarketplaceEventReceiver) GetEventType() string`

GetEventType returns the EventType field if non-nil, zero value otherwise.

### GetEventTypeOk

`func (o *MarketplaceEventReceiver) GetEventTypeOk() (*string, bool)`

GetEventTypeOk returns a tuple with the EventType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEventType

`func (o *MarketplaceEventReceiver) SetEventType(v string)`

SetEventType sets EventType field to given value.


### GetReceiverUrl

`func (o *MarketplaceEventReceiver) GetReceiverUrl() string`

GetReceiverUrl returns the ReceiverUrl field if non-nil, zero value otherwise.

### GetReceiverUrlOk

`func (o *MarketplaceEventReceiver) GetReceiverUrlOk() (*string, bool)`

GetReceiverUrlOk returns a tuple with the ReceiverUrl field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetReceiverUrl

`func (o *MarketplaceEventReceiver) SetReceiverUrl(v string)`

SetReceiverUrl sets ReceiverUrl field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


