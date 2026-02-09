# CreateInstanceSnapshotRequest2

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**InstanceId** | **string** | The instance ID | 
**TargetRegion** | Pointer to **string** | The target region to create the snapshot in. If not specified, use the same region as the instance | [optional] 

## Methods

### NewCreateInstanceSnapshotRequest2

`func NewCreateInstanceSnapshotRequest2(instanceId string, ) *CreateInstanceSnapshotRequest2`

NewCreateInstanceSnapshotRequest2 instantiates a new CreateInstanceSnapshotRequest2 object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewCreateInstanceSnapshotRequest2WithDefaults

`func NewCreateInstanceSnapshotRequest2WithDefaults() *CreateInstanceSnapshotRequest2`

NewCreateInstanceSnapshotRequest2WithDefaults instantiates a new CreateInstanceSnapshotRequest2 object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetInstanceId

`func (o *CreateInstanceSnapshotRequest2) GetInstanceId() string`

GetInstanceId returns the InstanceId field if non-nil, zero value otherwise.

### GetInstanceIdOk

`func (o *CreateInstanceSnapshotRequest2) GetInstanceIdOk() (*string, bool)`

GetInstanceIdOk returns a tuple with the InstanceId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetInstanceId

`func (o *CreateInstanceSnapshotRequest2) SetInstanceId(v string)`

SetInstanceId sets InstanceId field to given value.


### GetTargetRegion

`func (o *CreateInstanceSnapshotRequest2) GetTargetRegion() string`

GetTargetRegion returns the TargetRegion field if non-nil, zero value otherwise.

### GetTargetRegionOk

`func (o *CreateInstanceSnapshotRequest2) GetTargetRegionOk() (*string, bool)`

GetTargetRegionOk returns a tuple with the TargetRegion field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTargetRegion

`func (o *CreateInstanceSnapshotRequest2) SetTargetRegion(v string)`

SetTargetRegion sets TargetRegion field to given value.

### HasTargetRegion

`func (o *CreateInstanceSnapshotRequest2) HasTargetRegion() bool`

HasTargetRegion returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


