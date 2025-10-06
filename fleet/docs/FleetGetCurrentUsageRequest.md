# FleetGetCurrentUsageRequest

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**EndDate** | Pointer to **time.Time** | End time of the window in RFC 3339 format. If omitted, the filter is open-ended at the start. | [optional] 
**StartDate** | Pointer to **time.Time** | Start time of the window in RFC 3339 format. If omitted, the filter is open-ended at the start. | [optional] 
**SubscriptionIDs** | Pointer to **[]string** | Optionally filter by subscription IDs | [optional] 
**Token** | **string** | JWT token used to perform authorization | 

## Methods

### NewFleetGetCurrentUsageRequest

`func NewFleetGetCurrentUsageRequest(token string, ) *FleetGetCurrentUsageRequest`

NewFleetGetCurrentUsageRequest instantiates a new FleetGetCurrentUsageRequest object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewFleetGetCurrentUsageRequestWithDefaults

`func NewFleetGetCurrentUsageRequestWithDefaults() *FleetGetCurrentUsageRequest`

NewFleetGetCurrentUsageRequestWithDefaults instantiates a new FleetGetCurrentUsageRequest object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetEndDate

`func (o *FleetGetCurrentUsageRequest) GetEndDate() time.Time`

GetEndDate returns the EndDate field if non-nil, zero value otherwise.

### GetEndDateOk

`func (o *FleetGetCurrentUsageRequest) GetEndDateOk() (*time.Time, bool)`

GetEndDateOk returns a tuple with the EndDate field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEndDate

`func (o *FleetGetCurrentUsageRequest) SetEndDate(v time.Time)`

SetEndDate sets EndDate field to given value.

### HasEndDate

`func (o *FleetGetCurrentUsageRequest) HasEndDate() bool`

HasEndDate returns a boolean if a field has been set.

### GetStartDate

`func (o *FleetGetCurrentUsageRequest) GetStartDate() time.Time`

GetStartDate returns the StartDate field if non-nil, zero value otherwise.

### GetStartDateOk

`func (o *FleetGetCurrentUsageRequest) GetStartDateOk() (*time.Time, bool)`

GetStartDateOk returns a tuple with the StartDate field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStartDate

`func (o *FleetGetCurrentUsageRequest) SetStartDate(v time.Time)`

SetStartDate sets StartDate field to given value.

### HasStartDate

`func (o *FleetGetCurrentUsageRequest) HasStartDate() bool`

HasStartDate returns a boolean if a field has been set.

### GetSubscriptionIDs

`func (o *FleetGetCurrentUsageRequest) GetSubscriptionIDs() []string`

GetSubscriptionIDs returns the SubscriptionIDs field if non-nil, zero value otherwise.

### GetSubscriptionIDsOk

`func (o *FleetGetCurrentUsageRequest) GetSubscriptionIDsOk() (*[]string, bool)`

GetSubscriptionIDsOk returns a tuple with the SubscriptionIDs field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSubscriptionIDs

`func (o *FleetGetCurrentUsageRequest) SetSubscriptionIDs(v []string)`

SetSubscriptionIDs sets SubscriptionIDs field to given value.

### HasSubscriptionIDs

`func (o *FleetGetCurrentUsageRequest) HasSubscriptionIDs() bool`

HasSubscriptionIDs returns a boolean if a field has been set.

### GetToken

`func (o *FleetGetCurrentUsageRequest) GetToken() string`

GetToken returns the Token field if non-nil, zero value otherwise.

### GetTokenOk

`func (o *FleetGetCurrentUsageRequest) GetTokenOk() (*string, bool)`

GetTokenOk returns a tuple with the Token field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetToken

`func (o *FleetGetCurrentUsageRequest) SetToken(v string)`

SetToken sets Token field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


