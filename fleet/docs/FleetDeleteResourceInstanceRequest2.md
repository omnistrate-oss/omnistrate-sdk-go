# FleetDeleteResourceInstanceRequest2

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**ResourceId** | **string** | The resource ID. | 
**SkipFinalSnapshot** | Pointer to **bool** | If true, skip the automatic final snapshot before deletion even if snapshotBeforeDeletion is enabled for the resource. | [optional] [default to false]

## Methods

### NewFleetDeleteResourceInstanceRequest2

`func NewFleetDeleteResourceInstanceRequest2(resourceId string, ) *FleetDeleteResourceInstanceRequest2`

NewFleetDeleteResourceInstanceRequest2 instantiates a new FleetDeleteResourceInstanceRequest2 object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewFleetDeleteResourceInstanceRequest2WithDefaults

`func NewFleetDeleteResourceInstanceRequest2WithDefaults() *FleetDeleteResourceInstanceRequest2`

NewFleetDeleteResourceInstanceRequest2WithDefaults instantiates a new FleetDeleteResourceInstanceRequest2 object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetResourceId

`func (o *FleetDeleteResourceInstanceRequest2) GetResourceId() string`

GetResourceId returns the ResourceId field if non-nil, zero value otherwise.

### GetResourceIdOk

`func (o *FleetDeleteResourceInstanceRequest2) GetResourceIdOk() (*string, bool)`

GetResourceIdOk returns a tuple with the ResourceId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetResourceId

`func (o *FleetDeleteResourceInstanceRequest2) SetResourceId(v string)`

SetResourceId sets ResourceId field to given value.


### GetSkipFinalSnapshot

`func (o *FleetDeleteResourceInstanceRequest2) GetSkipFinalSnapshot() bool`

GetSkipFinalSnapshot returns the SkipFinalSnapshot field if non-nil, zero value otherwise.

### GetSkipFinalSnapshotOk

`func (o *FleetDeleteResourceInstanceRequest2) GetSkipFinalSnapshotOk() (*bool, bool)`

GetSkipFinalSnapshotOk returns a tuple with the SkipFinalSnapshot field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSkipFinalSnapshot

`func (o *FleetDeleteResourceInstanceRequest2) SetSkipFinalSnapshot(v bool)`

SetSkipFinalSnapshot sets SkipFinalSnapshot field to given value.

### HasSkipFinalSnapshot

`func (o *FleetDeleteResourceInstanceRequest2) HasSkipFinalSnapshot() bool`

HasSkipFinalSnapshot returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


