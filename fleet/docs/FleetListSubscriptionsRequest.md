# FleetListSubscriptionsRequest

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**EnvironmentId** | **string** | ID of a Service Environment | 
**ExcludePricing** | Pointer to **bool** | Whether to exclude billing details (active pricing, scheduled pricing, audit logs, billing provider) from the response | [optional] 
**ExcludeStats** | Pointer to **bool** | Whether to exclude subscription statistics (user count, instance count, first usage time) from the response | [optional] 
**IncludeInactive** | Pointer to **bool** | Whether to include inactive (suspended, cancelled, terminated) subscriptions | [optional] 
**NextPageToken** | Pointer to **string** | The next token to use for pagination | [optional] 
**PageSize** | Pointer to **int64** | The number of resources to return per page | [optional] 
**ProductTierId** | Pointer to **string** | ID of a Product Tier | [optional] 
**ServiceId** | **string** | ID of a Service | 
**Token** | **string** | JWT token used to perform authorization | 

## Methods

### NewFleetListSubscriptionsRequest

`func NewFleetListSubscriptionsRequest(environmentId string, serviceId string, token string, ) *FleetListSubscriptionsRequest`

NewFleetListSubscriptionsRequest instantiates a new FleetListSubscriptionsRequest object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewFleetListSubscriptionsRequestWithDefaults

`func NewFleetListSubscriptionsRequestWithDefaults() *FleetListSubscriptionsRequest`

NewFleetListSubscriptionsRequestWithDefaults instantiates a new FleetListSubscriptionsRequest object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetEnvironmentId

`func (o *FleetListSubscriptionsRequest) GetEnvironmentId() string`

GetEnvironmentId returns the EnvironmentId field if non-nil, zero value otherwise.

### GetEnvironmentIdOk

`func (o *FleetListSubscriptionsRequest) GetEnvironmentIdOk() (*string, bool)`

GetEnvironmentIdOk returns a tuple with the EnvironmentId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEnvironmentId

`func (o *FleetListSubscriptionsRequest) SetEnvironmentId(v string)`

SetEnvironmentId sets EnvironmentId field to given value.


### GetExcludePricing

`func (o *FleetListSubscriptionsRequest) GetExcludePricing() bool`

GetExcludePricing returns the ExcludePricing field if non-nil, zero value otherwise.

### GetExcludePricingOk

`func (o *FleetListSubscriptionsRequest) GetExcludePricingOk() (*bool, bool)`

GetExcludePricingOk returns a tuple with the ExcludePricing field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetExcludePricing

`func (o *FleetListSubscriptionsRequest) SetExcludePricing(v bool)`

SetExcludePricing sets ExcludePricing field to given value.

### HasExcludePricing

`func (o *FleetListSubscriptionsRequest) HasExcludePricing() bool`

HasExcludePricing returns a boolean if a field has been set.

### GetExcludeStats

`func (o *FleetListSubscriptionsRequest) GetExcludeStats() bool`

GetExcludeStats returns the ExcludeStats field if non-nil, zero value otherwise.

### GetExcludeStatsOk

`func (o *FleetListSubscriptionsRequest) GetExcludeStatsOk() (*bool, bool)`

GetExcludeStatsOk returns a tuple with the ExcludeStats field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetExcludeStats

`func (o *FleetListSubscriptionsRequest) SetExcludeStats(v bool)`

SetExcludeStats sets ExcludeStats field to given value.

### HasExcludeStats

`func (o *FleetListSubscriptionsRequest) HasExcludeStats() bool`

HasExcludeStats returns a boolean if a field has been set.

### GetIncludeInactive

`func (o *FleetListSubscriptionsRequest) GetIncludeInactive() bool`

GetIncludeInactive returns the IncludeInactive field if non-nil, zero value otherwise.

### GetIncludeInactiveOk

`func (o *FleetListSubscriptionsRequest) GetIncludeInactiveOk() (*bool, bool)`

GetIncludeInactiveOk returns a tuple with the IncludeInactive field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIncludeInactive

`func (o *FleetListSubscriptionsRequest) SetIncludeInactive(v bool)`

SetIncludeInactive sets IncludeInactive field to given value.

### HasIncludeInactive

`func (o *FleetListSubscriptionsRequest) HasIncludeInactive() bool`

HasIncludeInactive returns a boolean if a field has been set.

### GetNextPageToken

`func (o *FleetListSubscriptionsRequest) GetNextPageToken() string`

GetNextPageToken returns the NextPageToken field if non-nil, zero value otherwise.

### GetNextPageTokenOk

`func (o *FleetListSubscriptionsRequest) GetNextPageTokenOk() (*string, bool)`

GetNextPageTokenOk returns a tuple with the NextPageToken field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNextPageToken

`func (o *FleetListSubscriptionsRequest) SetNextPageToken(v string)`

SetNextPageToken sets NextPageToken field to given value.

### HasNextPageToken

`func (o *FleetListSubscriptionsRequest) HasNextPageToken() bool`

HasNextPageToken returns a boolean if a field has been set.

### GetPageSize

`func (o *FleetListSubscriptionsRequest) GetPageSize() int64`

GetPageSize returns the PageSize field if non-nil, zero value otherwise.

### GetPageSizeOk

`func (o *FleetListSubscriptionsRequest) GetPageSizeOk() (*int64, bool)`

GetPageSizeOk returns a tuple with the PageSize field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPageSize

`func (o *FleetListSubscriptionsRequest) SetPageSize(v int64)`

SetPageSize sets PageSize field to given value.

### HasPageSize

`func (o *FleetListSubscriptionsRequest) HasPageSize() bool`

HasPageSize returns a boolean if a field has been set.

### GetProductTierId

`func (o *FleetListSubscriptionsRequest) GetProductTierId() string`

GetProductTierId returns the ProductTierId field if non-nil, zero value otherwise.

### GetProductTierIdOk

`func (o *FleetListSubscriptionsRequest) GetProductTierIdOk() (*string, bool)`

GetProductTierIdOk returns a tuple with the ProductTierId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetProductTierId

`func (o *FleetListSubscriptionsRequest) SetProductTierId(v string)`

SetProductTierId sets ProductTierId field to given value.

### HasProductTierId

`func (o *FleetListSubscriptionsRequest) HasProductTierId() bool`

HasProductTierId returns a boolean if a field has been set.

### GetServiceId

`func (o *FleetListSubscriptionsRequest) GetServiceId() string`

GetServiceId returns the ServiceId field if non-nil, zero value otherwise.

### GetServiceIdOk

`func (o *FleetListSubscriptionsRequest) GetServiceIdOk() (*string, bool)`

GetServiceIdOk returns a tuple with the ServiceId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetServiceId

`func (o *FleetListSubscriptionsRequest) SetServiceId(v string)`

SetServiceId sets ServiceId field to given value.


### GetToken

`func (o *FleetListSubscriptionsRequest) GetToken() string`

GetToken returns the Token field if non-nil, zero value otherwise.

### GetTokenOk

`func (o *FleetListSubscriptionsRequest) GetTokenOk() (*string, bool)`

GetTokenOk returns a tuple with the Token field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetToken

`func (o *FleetListSubscriptionsRequest) SetToken(v string)`

SetToken sets Token field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


