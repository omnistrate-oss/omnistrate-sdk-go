# ListMarketplaceChannelListingsResult

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**DiscoverySupported** | **bool** | Whether this channel can enumerate its own listings | 
**Listings** | [**[]MarketplaceChannelListing**](MarketplaceChannelListing.md) |  | 

## Methods

### NewListMarketplaceChannelListingsResult

`func NewListMarketplaceChannelListingsResult(discoverySupported bool, listings []MarketplaceChannelListing, ) *ListMarketplaceChannelListingsResult`

NewListMarketplaceChannelListingsResult instantiates a new ListMarketplaceChannelListingsResult object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewListMarketplaceChannelListingsResultWithDefaults

`func NewListMarketplaceChannelListingsResultWithDefaults() *ListMarketplaceChannelListingsResult`

NewListMarketplaceChannelListingsResultWithDefaults instantiates a new ListMarketplaceChannelListingsResult object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetDiscoverySupported

`func (o *ListMarketplaceChannelListingsResult) GetDiscoverySupported() bool`

GetDiscoverySupported returns the DiscoverySupported field if non-nil, zero value otherwise.

### GetDiscoverySupportedOk

`func (o *ListMarketplaceChannelListingsResult) GetDiscoverySupportedOk() (*bool, bool)`

GetDiscoverySupportedOk returns a tuple with the DiscoverySupported field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDiscoverySupported

`func (o *ListMarketplaceChannelListingsResult) SetDiscoverySupported(v bool)`

SetDiscoverySupported sets DiscoverySupported field to given value.


### GetListings

`func (o *ListMarketplaceChannelListingsResult) GetListings() []MarketplaceChannelListing`

GetListings returns the Listings field if non-nil, zero value otherwise.

### GetListingsOk

`func (o *ListMarketplaceChannelListingsResult) GetListingsOk() (*[]MarketplaceChannelListing, bool)`

GetListingsOk returns a tuple with the Listings field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetListings

`func (o *ListMarketplaceChannelListingsResult) SetListings(v []MarketplaceChannelListing)`

SetListings sets Listings field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


