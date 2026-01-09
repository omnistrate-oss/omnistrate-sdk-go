# FleetGetConsolidatedUsageResult

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**ConsolidatedUsage** | Pointer to [**[]UsagePerSubscription**](UsagePerSubscription.md) | The consolidated usage for the requested service context | [optional] 

## Methods

### NewFleetGetConsolidatedUsageResult

`func NewFleetGetConsolidatedUsageResult() *FleetGetConsolidatedUsageResult`

NewFleetGetConsolidatedUsageResult instantiates a new FleetGetConsolidatedUsageResult object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewFleetGetConsolidatedUsageResultWithDefaults

`func NewFleetGetConsolidatedUsageResultWithDefaults() *FleetGetConsolidatedUsageResult`

NewFleetGetConsolidatedUsageResultWithDefaults instantiates a new FleetGetConsolidatedUsageResult object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetConsolidatedUsage

`func (o *FleetGetConsolidatedUsageResult) GetConsolidatedUsage() []UsagePerSubscription`

GetConsolidatedUsage returns the ConsolidatedUsage field if non-nil, zero value otherwise.

### GetConsolidatedUsageOk

`func (o *FleetGetConsolidatedUsageResult) GetConsolidatedUsageOk() (*[]UsagePerSubscription, bool)`

GetConsolidatedUsageOk returns a tuple with the ConsolidatedUsage field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetConsolidatedUsage

`func (o *FleetGetConsolidatedUsageResult) SetConsolidatedUsage(v []UsagePerSubscription)`

SetConsolidatedUsage sets ConsolidatedUsage field to given value.

### HasConsolidatedUsage

`func (o *FleetGetConsolidatedUsageResult) HasConsolidatedUsage() bool`

HasConsolidatedUsage returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


