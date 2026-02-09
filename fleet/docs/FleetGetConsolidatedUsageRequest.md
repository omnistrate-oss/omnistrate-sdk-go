# FleetGetConsolidatedUsageRequest

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**EndDate** | Pointer to **time.Time** | End time of the window in RFC 3339 format. If omitted, the filter is open-ended at the start. | [optional] 
**EnvironmentId** | **string** | ID of a Service Environment | 
**ProductTierId** | Pointer to **string** | ID of a Product Tier | [optional] 
**ServiceId** | **string** | ID of a Service | 
**StartDate** | Pointer to **time.Time** | Start time of the window in RFC 3339 format. If omitted, the filter is open-ended at the start. | [optional] 
**Token** | **string** | JWT token used to perform authorization | 

## Methods

### NewFleetGetConsolidatedUsageRequest

`func NewFleetGetConsolidatedUsageRequest(environmentId string, serviceId string, token string, ) *FleetGetConsolidatedUsageRequest`

NewFleetGetConsolidatedUsageRequest instantiates a new FleetGetConsolidatedUsageRequest object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewFleetGetConsolidatedUsageRequestWithDefaults

`func NewFleetGetConsolidatedUsageRequestWithDefaults() *FleetGetConsolidatedUsageRequest`

NewFleetGetConsolidatedUsageRequestWithDefaults instantiates a new FleetGetConsolidatedUsageRequest object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetEndDate

`func (o *FleetGetConsolidatedUsageRequest) GetEndDate() time.Time`

GetEndDate returns the EndDate field if non-nil, zero value otherwise.

### GetEndDateOk

`func (o *FleetGetConsolidatedUsageRequest) GetEndDateOk() (*time.Time, bool)`

GetEndDateOk returns a tuple with the EndDate field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEndDate

`func (o *FleetGetConsolidatedUsageRequest) SetEndDate(v time.Time)`

SetEndDate sets EndDate field to given value.

### HasEndDate

`func (o *FleetGetConsolidatedUsageRequest) HasEndDate() bool`

HasEndDate returns a boolean if a field has been set.

### GetEnvironmentId

`func (o *FleetGetConsolidatedUsageRequest) GetEnvironmentId() string`

GetEnvironmentId returns the EnvironmentId field if non-nil, zero value otherwise.

### GetEnvironmentIdOk

`func (o *FleetGetConsolidatedUsageRequest) GetEnvironmentIdOk() (*string, bool)`

GetEnvironmentIdOk returns a tuple with the EnvironmentId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEnvironmentId

`func (o *FleetGetConsolidatedUsageRequest) SetEnvironmentId(v string)`

SetEnvironmentId sets EnvironmentId field to given value.


### GetProductTierId

`func (o *FleetGetConsolidatedUsageRequest) GetProductTierId() string`

GetProductTierId returns the ProductTierId field if non-nil, zero value otherwise.

### GetProductTierIdOk

`func (o *FleetGetConsolidatedUsageRequest) GetProductTierIdOk() (*string, bool)`

GetProductTierIdOk returns a tuple with the ProductTierId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetProductTierId

`func (o *FleetGetConsolidatedUsageRequest) SetProductTierId(v string)`

SetProductTierId sets ProductTierId field to given value.

### HasProductTierId

`func (o *FleetGetConsolidatedUsageRequest) HasProductTierId() bool`

HasProductTierId returns a boolean if a field has been set.

### GetServiceId

`func (o *FleetGetConsolidatedUsageRequest) GetServiceId() string`

GetServiceId returns the ServiceId field if non-nil, zero value otherwise.

### GetServiceIdOk

`func (o *FleetGetConsolidatedUsageRequest) GetServiceIdOk() (*string, bool)`

GetServiceIdOk returns a tuple with the ServiceId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetServiceId

`func (o *FleetGetConsolidatedUsageRequest) SetServiceId(v string)`

SetServiceId sets ServiceId field to given value.


### GetStartDate

`func (o *FleetGetConsolidatedUsageRequest) GetStartDate() time.Time`

GetStartDate returns the StartDate field if non-nil, zero value otherwise.

### GetStartDateOk

`func (o *FleetGetConsolidatedUsageRequest) GetStartDateOk() (*time.Time, bool)`

GetStartDateOk returns a tuple with the StartDate field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStartDate

`func (o *FleetGetConsolidatedUsageRequest) SetStartDate(v time.Time)`

SetStartDate sets StartDate field to given value.

### HasStartDate

`func (o *FleetGetConsolidatedUsageRequest) HasStartDate() bool`

HasStartDate returns a boolean if a field has been set.

### GetToken

`func (o *FleetGetConsolidatedUsageRequest) GetToken() string`

GetToken returns the Token field if non-nil, zero value otherwise.

### GetTokenOk

`func (o *FleetGetConsolidatedUsageRequest) GetTokenOk() (*string, bool)`

GetTokenOk returns a tuple with the Token field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetToken

`func (o *FleetGetConsolidatedUsageRequest) SetToken(v string)`

SetToken sets Token field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


