# GetConsumptionUsageResult

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**EndTime** | Pointer to **string** | End timestamp of usage | [optional] 
**StartTime** | Pointer to **string** | Start timestamp of usage | [optional] 
**Usage** | Pointer to [**[]UsagePerDimension**](UsagePerDimension.md) | Usage for the current plan | [optional] 

## Methods

### NewGetConsumptionUsageResult

`func NewGetConsumptionUsageResult() *GetConsumptionUsageResult`

NewGetConsumptionUsageResult instantiates a new GetConsumptionUsageResult object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewGetConsumptionUsageResultWithDefaults

`func NewGetConsumptionUsageResultWithDefaults() *GetConsumptionUsageResult`

NewGetConsumptionUsageResultWithDefaults instantiates a new GetConsumptionUsageResult object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetEndTime

`func (o *GetConsumptionUsageResult) GetEndTime() string`

GetEndTime returns the EndTime field if non-nil, zero value otherwise.

### GetEndTimeOk

`func (o *GetConsumptionUsageResult) GetEndTimeOk() (*string, bool)`

GetEndTimeOk returns a tuple with the EndTime field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEndTime

`func (o *GetConsumptionUsageResult) SetEndTime(v string)`

SetEndTime sets EndTime field to given value.

### HasEndTime

`func (o *GetConsumptionUsageResult) HasEndTime() bool`

HasEndTime returns a boolean if a field has been set.

### GetStartTime

`func (o *GetConsumptionUsageResult) GetStartTime() string`

GetStartTime returns the StartTime field if non-nil, zero value otherwise.

### GetStartTimeOk

`func (o *GetConsumptionUsageResult) GetStartTimeOk() (*string, bool)`

GetStartTimeOk returns a tuple with the StartTime field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStartTime

`func (o *GetConsumptionUsageResult) SetStartTime(v string)`

SetStartTime sets StartTime field to given value.

### HasStartTime

`func (o *GetConsumptionUsageResult) HasStartTime() bool`

HasStartTime returns a boolean if a field has been set.

### GetUsage

`func (o *GetConsumptionUsageResult) GetUsage() []UsagePerDimension`

GetUsage returns the Usage field if non-nil, zero value otherwise.

### GetUsageOk

`func (o *GetConsumptionUsageResult) GetUsageOk() (*[]UsagePerDimension, bool)`

GetUsageOk returns a tuple with the Usage field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUsage

`func (o *GetConsumptionUsageResult) SetUsage(v []UsagePerDimension)`

SetUsage sets Usage field to given value.

### HasUsage

`func (o *GetConsumptionUsageResult) HasUsage() bool`

HasUsage returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


