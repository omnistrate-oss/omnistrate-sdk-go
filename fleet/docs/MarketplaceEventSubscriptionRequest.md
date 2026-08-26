# MarketplaceEventSubscriptionRequest

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**EnvironmentId** | **string** |  | 
**Id** | **string** | The subscription request to approve or deny. This is the id the ISV passes to the approve endpoint, and approving it is the only mandatory call in the integration | 
**ProductTierId** | **string** |  | 
**ServiceId** | **string** |  | 
**Status** | **string** | Status of the subscription request | 

## Methods

### NewMarketplaceEventSubscriptionRequest

`func NewMarketplaceEventSubscriptionRequest(environmentId string, id string, productTierId string, serviceId string, status string, ) *MarketplaceEventSubscriptionRequest`

NewMarketplaceEventSubscriptionRequest instantiates a new MarketplaceEventSubscriptionRequest object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewMarketplaceEventSubscriptionRequestWithDefaults

`func NewMarketplaceEventSubscriptionRequestWithDefaults() *MarketplaceEventSubscriptionRequest`

NewMarketplaceEventSubscriptionRequestWithDefaults instantiates a new MarketplaceEventSubscriptionRequest object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetEnvironmentId

`func (o *MarketplaceEventSubscriptionRequest) GetEnvironmentId() string`

GetEnvironmentId returns the EnvironmentId field if non-nil, zero value otherwise.

### GetEnvironmentIdOk

`func (o *MarketplaceEventSubscriptionRequest) GetEnvironmentIdOk() (*string, bool)`

GetEnvironmentIdOk returns a tuple with the EnvironmentId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEnvironmentId

`func (o *MarketplaceEventSubscriptionRequest) SetEnvironmentId(v string)`

SetEnvironmentId sets EnvironmentId field to given value.


### GetId

`func (o *MarketplaceEventSubscriptionRequest) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *MarketplaceEventSubscriptionRequest) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *MarketplaceEventSubscriptionRequest) SetId(v string)`

SetId sets Id field to given value.


### GetProductTierId

`func (o *MarketplaceEventSubscriptionRequest) GetProductTierId() string`

GetProductTierId returns the ProductTierId field if non-nil, zero value otherwise.

### GetProductTierIdOk

`func (o *MarketplaceEventSubscriptionRequest) GetProductTierIdOk() (*string, bool)`

GetProductTierIdOk returns a tuple with the ProductTierId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetProductTierId

`func (o *MarketplaceEventSubscriptionRequest) SetProductTierId(v string)`

SetProductTierId sets ProductTierId field to given value.


### GetServiceId

`func (o *MarketplaceEventSubscriptionRequest) GetServiceId() string`

GetServiceId returns the ServiceId field if non-nil, zero value otherwise.

### GetServiceIdOk

`func (o *MarketplaceEventSubscriptionRequest) GetServiceIdOk() (*string, bool)`

GetServiceIdOk returns a tuple with the ServiceId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetServiceId

`func (o *MarketplaceEventSubscriptionRequest) SetServiceId(v string)`

SetServiceId sets ServiceId field to given value.


### GetStatus

`func (o *MarketplaceEventSubscriptionRequest) GetStatus() string`

GetStatus returns the Status field if non-nil, zero value otherwise.

### GetStatusOk

`func (o *MarketplaceEventSubscriptionRequest) GetStatusOk() (*string, bool)`

GetStatusOk returns a tuple with the Status field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStatus

`func (o *MarketplaceEventSubscriptionRequest) SetStatus(v string)`

SetStatus sets Status field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


