# GetUsageResult

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**EndTime** | Pointer to **string** | End timestamp of usage | [optional] 
**PlanName** | Pointer to **string** | This parameter is used to select the appropriate Product Plan | [optional] 
**StartTime** | Pointer to **string** | Start timestamp of usage | [optional] 
**Usage** | Pointer to [**[]UsagePerDimension**](UsagePerDimension.md) | Usage for the current plan | [optional] 

## Methods

### NewGetUsageResult

`func NewGetUsageResult() *GetUsageResult`

NewGetUsageResult instantiates a new GetUsageResult object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewGetUsageResultWithDefaults

`func NewGetUsageResultWithDefaults() *GetUsageResult`

NewGetUsageResultWithDefaults instantiates a new GetUsageResult object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetEndTime

`func (o *GetUsageResult) GetEndTime() string`

GetEndTime returns the EndTime field if non-nil, zero value otherwise.

### GetEndTimeOk

`func (o *GetUsageResult) GetEndTimeOk() (*string, bool)`

GetEndTimeOk returns a tuple with the EndTime field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEndTime

`func (o *GetUsageResult) SetEndTime(v string)`

SetEndTime sets EndTime field to given value.

### HasEndTime

`func (o *GetUsageResult) HasEndTime() bool`

HasEndTime returns a boolean if a field has been set.

### GetPlanName

`func (o *GetUsageResult) GetPlanName() string`

GetPlanName returns the PlanName field if non-nil, zero value otherwise.

### GetPlanNameOk

`func (o *GetUsageResult) GetPlanNameOk() (*string, bool)`

GetPlanNameOk returns a tuple with the PlanName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPlanName

`func (o *GetUsageResult) SetPlanName(v string)`

SetPlanName sets PlanName field to given value.

### HasPlanName

`func (o *GetUsageResult) HasPlanName() bool`

HasPlanName returns a boolean if a field has been set.

### GetStartTime

`func (o *GetUsageResult) GetStartTime() string`

GetStartTime returns the StartTime field if non-nil, zero value otherwise.

### GetStartTimeOk

`func (o *GetUsageResult) GetStartTimeOk() (*string, bool)`

GetStartTimeOk returns a tuple with the StartTime field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStartTime

`func (o *GetUsageResult) SetStartTime(v string)`

SetStartTime sets StartTime field to given value.

### HasStartTime

`func (o *GetUsageResult) HasStartTime() bool`

HasStartTime returns a boolean if a field has been set.

### GetUsage

`func (o *GetUsageResult) GetUsage() []UsagePerDimension`

GetUsage returns the Usage field if non-nil, zero value otherwise.

### GetUsageOk

`func (o *GetUsageResult) GetUsageOk() (*[]UsagePerDimension, bool)`

GetUsageOk returns a tuple with the Usage field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUsage

`func (o *GetUsageResult) SetUsage(v []UsagePerDimension)`

SetUsage sets Usage field to given value.

### HasUsage

`func (o *GetUsageResult) HasUsage() bool`

HasUsage returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


