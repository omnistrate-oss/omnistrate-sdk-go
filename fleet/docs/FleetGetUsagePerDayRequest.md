# FleetGetUsagePerDayRequest

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**EndDate** | Pointer to **time.Time** | End time of the window in RFC 3339 format. If omitted, the filter is open-ended at the start. | [optional] 
**StartDate** | Pointer to **time.Time** | Start time of the window in RFC 3339 format. If omitted, the filter is open-ended at the start. | [optional] 
**SubscriptionIDs** | Pointer to **[]string** | Optionally filter by subscription IDs | [optional] 
**Token** | **string** | JWT token used to perform authorization | 

## Methods

### NewFleetGetUsagePerDayRequest

`func NewFleetGetUsagePerDayRequest(token string, ) *FleetGetUsagePerDayRequest`

NewFleetGetUsagePerDayRequest instantiates a new FleetGetUsagePerDayRequest object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewFleetGetUsagePerDayRequestWithDefaults

`func NewFleetGetUsagePerDayRequestWithDefaults() *FleetGetUsagePerDayRequest`

NewFleetGetUsagePerDayRequestWithDefaults instantiates a new FleetGetUsagePerDayRequest object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetEndDate

`func (o *FleetGetUsagePerDayRequest) GetEndDate() time.Time`

GetEndDate returns the EndDate field if non-nil, zero value otherwise.

### GetEndDateOk

`func (o *FleetGetUsagePerDayRequest) GetEndDateOk() (*time.Time, bool)`

GetEndDateOk returns a tuple with the EndDate field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEndDate

`func (o *FleetGetUsagePerDayRequest) SetEndDate(v time.Time)`

SetEndDate sets EndDate field to given value.

### HasEndDate

`func (o *FleetGetUsagePerDayRequest) HasEndDate() bool`

HasEndDate returns a boolean if a field has been set.

### GetStartDate

`func (o *FleetGetUsagePerDayRequest) GetStartDate() time.Time`

GetStartDate returns the StartDate field if non-nil, zero value otherwise.

### GetStartDateOk

`func (o *FleetGetUsagePerDayRequest) GetStartDateOk() (*time.Time, bool)`

GetStartDateOk returns a tuple with the StartDate field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStartDate

`func (o *FleetGetUsagePerDayRequest) SetStartDate(v time.Time)`

SetStartDate sets StartDate field to given value.

### HasStartDate

`func (o *FleetGetUsagePerDayRequest) HasStartDate() bool`

HasStartDate returns a boolean if a field has been set.

### GetSubscriptionIDs

`func (o *FleetGetUsagePerDayRequest) GetSubscriptionIDs() []string`

GetSubscriptionIDs returns the SubscriptionIDs field if non-nil, zero value otherwise.

### GetSubscriptionIDsOk

`func (o *FleetGetUsagePerDayRequest) GetSubscriptionIDsOk() (*[]string, bool)`

GetSubscriptionIDsOk returns a tuple with the SubscriptionIDs field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSubscriptionIDs

`func (o *FleetGetUsagePerDayRequest) SetSubscriptionIDs(v []string)`

SetSubscriptionIDs sets SubscriptionIDs field to given value.

### HasSubscriptionIDs

`func (o *FleetGetUsagePerDayRequest) HasSubscriptionIDs() bool`

HasSubscriptionIDs returns a boolean if a field has been set.

### GetToken

`func (o *FleetGetUsagePerDayRequest) GetToken() string`

GetToken returns the Token field if non-nil, zero value otherwise.

### GetTokenOk

`func (o *FleetGetUsagePerDayRequest) GetTokenOk() (*string, bool)`

GetTokenOk returns a tuple with the Token field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetToken

`func (o *FleetGetUsagePerDayRequest) SetToken(v string)`

SetToken sets Token field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


