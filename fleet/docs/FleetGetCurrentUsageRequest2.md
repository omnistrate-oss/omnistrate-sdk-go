# FleetGetCurrentUsageRequest2

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**EndDate** | Pointer to **time.Time** |  | [optional] 
**StartDate** | Pointer to **time.Time** |  | [optional] 
**SubscriptionIDs** | Pointer to **[]string** | Optionally filter by subscription IDs | [optional] 

## Methods

### NewFleetGetCurrentUsageRequest2

`func NewFleetGetCurrentUsageRequest2() *FleetGetCurrentUsageRequest2`

NewFleetGetCurrentUsageRequest2 instantiates a new FleetGetCurrentUsageRequest2 object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewFleetGetCurrentUsageRequest2WithDefaults

`func NewFleetGetCurrentUsageRequest2WithDefaults() *FleetGetCurrentUsageRequest2`

NewFleetGetCurrentUsageRequest2WithDefaults instantiates a new FleetGetCurrentUsageRequest2 object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetEndDate

`func (o *FleetGetCurrentUsageRequest2) GetEndDate() time.Time`

GetEndDate returns the EndDate field if non-nil, zero value otherwise.

### GetEndDateOk

`func (o *FleetGetCurrentUsageRequest2) GetEndDateOk() (*time.Time, bool)`

GetEndDateOk returns a tuple with the EndDate field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEndDate

`func (o *FleetGetCurrentUsageRequest2) SetEndDate(v time.Time)`

SetEndDate sets EndDate field to given value.

### HasEndDate

`func (o *FleetGetCurrentUsageRequest2) HasEndDate() bool`

HasEndDate returns a boolean if a field has been set.

### GetStartDate

`func (o *FleetGetCurrentUsageRequest2) GetStartDate() time.Time`

GetStartDate returns the StartDate field if non-nil, zero value otherwise.

### GetStartDateOk

`func (o *FleetGetCurrentUsageRequest2) GetStartDateOk() (*time.Time, bool)`

GetStartDateOk returns a tuple with the StartDate field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStartDate

`func (o *FleetGetCurrentUsageRequest2) SetStartDate(v time.Time)`

SetStartDate sets StartDate field to given value.

### HasStartDate

`func (o *FleetGetCurrentUsageRequest2) HasStartDate() bool`

HasStartDate returns a boolean if a field has been set.

### GetSubscriptionIDs

`func (o *FleetGetCurrentUsageRequest2) GetSubscriptionIDs() []string`

GetSubscriptionIDs returns the SubscriptionIDs field if non-nil, zero value otherwise.

### GetSubscriptionIDsOk

`func (o *FleetGetCurrentUsageRequest2) GetSubscriptionIDsOk() (*[]string, bool)`

GetSubscriptionIDsOk returns a tuple with the SubscriptionIDs field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSubscriptionIDs

`func (o *FleetGetCurrentUsageRequest2) SetSubscriptionIDs(v []string)`

SetSubscriptionIDs sets SubscriptionIDs field to given value.

### HasSubscriptionIDs

`func (o *FleetGetCurrentUsageRequest2) HasSubscriptionIDs() bool`

HasSubscriptionIDs returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


