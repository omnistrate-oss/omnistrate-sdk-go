# MarketplacePlanMapping

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**BillingProvider** | Pointer to **string** | The billing provider type: STRIPE, BRING_YOUR_OWN, AWS_MARKETPLACE, GCP_MARKETPLACE or AZURE_MARKETPLACE. Deprecated: MARKETPLACE is retained for legacy channel billing bindings and does not identify one cloud marketplace | [optional] 
**ProductTierId** | **string** |  | 
**ServiceId** | **string** |  | 

## Methods

### NewMarketplacePlanMapping

`func NewMarketplacePlanMapping(productTierId string, serviceId string, ) *MarketplacePlanMapping`

NewMarketplacePlanMapping instantiates a new MarketplacePlanMapping object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewMarketplacePlanMappingWithDefaults

`func NewMarketplacePlanMappingWithDefaults() *MarketplacePlanMapping`

NewMarketplacePlanMappingWithDefaults instantiates a new MarketplacePlanMapping object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetBillingProvider

`func (o *MarketplacePlanMapping) GetBillingProvider() string`

GetBillingProvider returns the BillingProvider field if non-nil, zero value otherwise.

### GetBillingProviderOk

`func (o *MarketplacePlanMapping) GetBillingProviderOk() (*string, bool)`

GetBillingProviderOk returns a tuple with the BillingProvider field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBillingProvider

`func (o *MarketplacePlanMapping) SetBillingProvider(v string)`

SetBillingProvider sets BillingProvider field to given value.

### HasBillingProvider

`func (o *MarketplacePlanMapping) HasBillingProvider() bool`

HasBillingProvider returns a boolean if a field has been set.

### GetProductTierId

`func (o *MarketplacePlanMapping) GetProductTierId() string`

GetProductTierId returns the ProductTierId field if non-nil, zero value otherwise.

### GetProductTierIdOk

`func (o *MarketplacePlanMapping) GetProductTierIdOk() (*string, bool)`

GetProductTierIdOk returns a tuple with the ProductTierId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetProductTierId

`func (o *MarketplacePlanMapping) SetProductTierId(v string)`

SetProductTierId sets ProductTierId field to given value.


### GetServiceId

`func (o *MarketplacePlanMapping) GetServiceId() string`

GetServiceId returns the ServiceId field if non-nil, zero value otherwise.

### GetServiceIdOk

`func (o *MarketplacePlanMapping) GetServiceIdOk() (*string, bool)`

GetServiceIdOk returns a tuple with the ServiceId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetServiceId

`func (o *MarketplacePlanMapping) SetServiceId(v string)`

SetServiceId sets ServiceId field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


