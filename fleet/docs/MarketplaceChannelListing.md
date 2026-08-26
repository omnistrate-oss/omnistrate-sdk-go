# MarketplaceChannelListing

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Description** | Pointer to **string** |  | [optional] 
**MappedProductTierId** | Pointer to **string** |  | [optional] 
**MappedServiceId** | Pointer to **string** |  | [optional] 
**MarketplaceProductId** | Pointer to **string** | The channel&#39;s product identifier, which is what an ISV recognises in its console | [optional] 
**Name** | **string** | What the listing is called on the marketplace, product and plan together | 
**PlanRef** | **string** | The channel&#39;s own identifier, exactly as it arrives on a contract. The mapping key | 
**Status** | Pointer to **string** | The channel&#39;s own word for whether the listing is live. Displayed, never interpreted | [optional] 

## Methods

### NewMarketplaceChannelListing

`func NewMarketplaceChannelListing(name string, planRef string, ) *MarketplaceChannelListing`

NewMarketplaceChannelListing instantiates a new MarketplaceChannelListing object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewMarketplaceChannelListingWithDefaults

`func NewMarketplaceChannelListingWithDefaults() *MarketplaceChannelListing`

NewMarketplaceChannelListingWithDefaults instantiates a new MarketplaceChannelListing object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetDescription

`func (o *MarketplaceChannelListing) GetDescription() string`

GetDescription returns the Description field if non-nil, zero value otherwise.

### GetDescriptionOk

`func (o *MarketplaceChannelListing) GetDescriptionOk() (*string, bool)`

GetDescriptionOk returns a tuple with the Description field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDescription

`func (o *MarketplaceChannelListing) SetDescription(v string)`

SetDescription sets Description field to given value.

### HasDescription

`func (o *MarketplaceChannelListing) HasDescription() bool`

HasDescription returns a boolean if a field has been set.

### GetMappedProductTierId

`func (o *MarketplaceChannelListing) GetMappedProductTierId() string`

GetMappedProductTierId returns the MappedProductTierId field if non-nil, zero value otherwise.

### GetMappedProductTierIdOk

`func (o *MarketplaceChannelListing) GetMappedProductTierIdOk() (*string, bool)`

GetMappedProductTierIdOk returns a tuple with the MappedProductTierId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMappedProductTierId

`func (o *MarketplaceChannelListing) SetMappedProductTierId(v string)`

SetMappedProductTierId sets MappedProductTierId field to given value.

### HasMappedProductTierId

`func (o *MarketplaceChannelListing) HasMappedProductTierId() bool`

HasMappedProductTierId returns a boolean if a field has been set.

### GetMappedServiceId

`func (o *MarketplaceChannelListing) GetMappedServiceId() string`

GetMappedServiceId returns the MappedServiceId field if non-nil, zero value otherwise.

### GetMappedServiceIdOk

`func (o *MarketplaceChannelListing) GetMappedServiceIdOk() (*string, bool)`

GetMappedServiceIdOk returns a tuple with the MappedServiceId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMappedServiceId

`func (o *MarketplaceChannelListing) SetMappedServiceId(v string)`

SetMappedServiceId sets MappedServiceId field to given value.

### HasMappedServiceId

`func (o *MarketplaceChannelListing) HasMappedServiceId() bool`

HasMappedServiceId returns a boolean if a field has been set.

### GetMarketplaceProductId

`func (o *MarketplaceChannelListing) GetMarketplaceProductId() string`

GetMarketplaceProductId returns the MarketplaceProductId field if non-nil, zero value otherwise.

### GetMarketplaceProductIdOk

`func (o *MarketplaceChannelListing) GetMarketplaceProductIdOk() (*string, bool)`

GetMarketplaceProductIdOk returns a tuple with the MarketplaceProductId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMarketplaceProductId

`func (o *MarketplaceChannelListing) SetMarketplaceProductId(v string)`

SetMarketplaceProductId sets MarketplaceProductId field to given value.

### HasMarketplaceProductId

`func (o *MarketplaceChannelListing) HasMarketplaceProductId() bool`

HasMarketplaceProductId returns a boolean if a field has been set.

### GetName

`func (o *MarketplaceChannelListing) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *MarketplaceChannelListing) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *MarketplaceChannelListing) SetName(v string)`

SetName sets Name field to given value.


### GetPlanRef

`func (o *MarketplaceChannelListing) GetPlanRef() string`

GetPlanRef returns the PlanRef field if non-nil, zero value otherwise.

### GetPlanRefOk

`func (o *MarketplaceChannelListing) GetPlanRefOk() (*string, bool)`

GetPlanRefOk returns a tuple with the PlanRef field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPlanRef

`func (o *MarketplaceChannelListing) SetPlanRef(v string)`

SetPlanRef sets PlanRef field to given value.


### GetStatus

`func (o *MarketplaceChannelListing) GetStatus() string`

GetStatus returns the Status field if non-nil, zero value otherwise.

### GetStatusOk

`func (o *MarketplaceChannelListing) GetStatusOk() (*string, bool)`

GetStatusOk returns a tuple with the Status field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStatus

`func (o *MarketplaceChannelListing) SetStatus(v string)`

SetStatus sets Status field to given value.

### HasStatus

`func (o *MarketplaceChannelListing) HasStatus() bool`

HasStatus returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


