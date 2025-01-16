# InstanceHealthSummary

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**InstanceID** | Pointer to **string** | ID of a Resource Instance | [optional] 
**LifeCycleStatus** | Pointer to **string** | The status of an operation | [optional] 
**ResourcesHealth** | Pointer to [**map[string]ResourceHealthSummary**](ResourceHealthSummary.md) | The health summary of the resources by resource ID | [optional] 
**Status** | Pointer to **string** | The heath status of a resource | [optional] 

## Methods

### NewInstanceHealthSummary

`func NewInstanceHealthSummary() *InstanceHealthSummary`

NewInstanceHealthSummary instantiates a new InstanceHealthSummary object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewInstanceHealthSummaryWithDefaults

`func NewInstanceHealthSummaryWithDefaults() *InstanceHealthSummary`

NewInstanceHealthSummaryWithDefaults instantiates a new InstanceHealthSummary object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetInstanceID

`func (o *InstanceHealthSummary) GetInstanceID() string`

GetInstanceID returns the InstanceID field if non-nil, zero value otherwise.

### GetInstanceIDOk

`func (o *InstanceHealthSummary) GetInstanceIDOk() (*string, bool)`

GetInstanceIDOk returns a tuple with the InstanceID field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetInstanceID

`func (o *InstanceHealthSummary) SetInstanceID(v string)`

SetInstanceID sets InstanceID field to given value.

### HasInstanceID

`func (o *InstanceHealthSummary) HasInstanceID() bool`

HasInstanceID returns a boolean if a field has been set.

### GetLifeCycleStatus

`func (o *InstanceHealthSummary) GetLifeCycleStatus() string`

GetLifeCycleStatus returns the LifeCycleStatus field if non-nil, zero value otherwise.

### GetLifeCycleStatusOk

`func (o *InstanceHealthSummary) GetLifeCycleStatusOk() (*string, bool)`

GetLifeCycleStatusOk returns a tuple with the LifeCycleStatus field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLifeCycleStatus

`func (o *InstanceHealthSummary) SetLifeCycleStatus(v string)`

SetLifeCycleStatus sets LifeCycleStatus field to given value.

### HasLifeCycleStatus

`func (o *InstanceHealthSummary) HasLifeCycleStatus() bool`

HasLifeCycleStatus returns a boolean if a field has been set.

### GetResourcesHealth

`func (o *InstanceHealthSummary) GetResourcesHealth() map[string]ResourceHealthSummary`

GetResourcesHealth returns the ResourcesHealth field if non-nil, zero value otherwise.

### GetResourcesHealthOk

`func (o *InstanceHealthSummary) GetResourcesHealthOk() (*map[string]ResourceHealthSummary, bool)`

GetResourcesHealthOk returns a tuple with the ResourcesHealth field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetResourcesHealth

`func (o *InstanceHealthSummary) SetResourcesHealth(v map[string]ResourceHealthSummary)`

SetResourcesHealth sets ResourcesHealth field to given value.

### HasResourcesHealth

`func (o *InstanceHealthSummary) HasResourcesHealth() bool`

HasResourcesHealth returns a boolean if a field has been set.

### GetStatus

`func (o *InstanceHealthSummary) GetStatus() string`

GetStatus returns the Status field if non-nil, zero value otherwise.

### GetStatusOk

`func (o *InstanceHealthSummary) GetStatusOk() (*string, bool)`

GetStatusOk returns a tuple with the Status field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStatus

`func (o *InstanceHealthSummary) SetStatus(v string)`

SetStatus sets Status field to given value.

### HasStatus

`func (o *InstanceHealthSummary) HasStatus() bool`

HasStatus returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


