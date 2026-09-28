# InfrastructureVolume

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Availability** | Pointer to [**InfrastructureAvailability**](InfrastructureAvailability.md) |  | [optional] 
**CsiDriver** | Pointer to **string** |  | [optional] 
**FsType** | Pointer to **string** |  | [optional] 
**ReclaimPolicy** | Pointer to **string** |  | [optional] 
**Ref** | Pointer to [**InfrastructureObjectReference**](InfrastructureObjectReference.md) |  | [optional] 
**VolumeHandle** | Pointer to **string** |  | [optional] 
**VolumeMode** | Pointer to **string** |  | [optional] 
**Zones** | Pointer to **[]string** |  | [optional] 
**ZonesTruncation** | Pointer to [**CheckpointTruncation**](CheckpointTruncation.md) |  | [optional] 

## Methods

### NewInfrastructureVolume

`func NewInfrastructureVolume() *InfrastructureVolume`

NewInfrastructureVolume instantiates a new InfrastructureVolume object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewInfrastructureVolumeWithDefaults

`func NewInfrastructureVolumeWithDefaults() *InfrastructureVolume`

NewInfrastructureVolumeWithDefaults instantiates a new InfrastructureVolume object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetAvailability

`func (o *InfrastructureVolume) GetAvailability() InfrastructureAvailability`

GetAvailability returns the Availability field if non-nil, zero value otherwise.

### GetAvailabilityOk

`func (o *InfrastructureVolume) GetAvailabilityOk() (*InfrastructureAvailability, bool)`

GetAvailabilityOk returns a tuple with the Availability field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAvailability

`func (o *InfrastructureVolume) SetAvailability(v InfrastructureAvailability)`

SetAvailability sets Availability field to given value.

### HasAvailability

`func (o *InfrastructureVolume) HasAvailability() bool`

HasAvailability returns a boolean if a field has been set.

### GetCsiDriver

`func (o *InfrastructureVolume) GetCsiDriver() string`

GetCsiDriver returns the CsiDriver field if non-nil, zero value otherwise.

### GetCsiDriverOk

`func (o *InfrastructureVolume) GetCsiDriverOk() (*string, bool)`

GetCsiDriverOk returns a tuple with the CsiDriver field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCsiDriver

`func (o *InfrastructureVolume) SetCsiDriver(v string)`

SetCsiDriver sets CsiDriver field to given value.

### HasCsiDriver

`func (o *InfrastructureVolume) HasCsiDriver() bool`

HasCsiDriver returns a boolean if a field has been set.

### GetFsType

`func (o *InfrastructureVolume) GetFsType() string`

GetFsType returns the FsType field if non-nil, zero value otherwise.

### GetFsTypeOk

`func (o *InfrastructureVolume) GetFsTypeOk() (*string, bool)`

GetFsTypeOk returns a tuple with the FsType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFsType

`func (o *InfrastructureVolume) SetFsType(v string)`

SetFsType sets FsType field to given value.

### HasFsType

`func (o *InfrastructureVolume) HasFsType() bool`

HasFsType returns a boolean if a field has been set.

### GetReclaimPolicy

`func (o *InfrastructureVolume) GetReclaimPolicy() string`

GetReclaimPolicy returns the ReclaimPolicy field if non-nil, zero value otherwise.

### GetReclaimPolicyOk

`func (o *InfrastructureVolume) GetReclaimPolicyOk() (*string, bool)`

GetReclaimPolicyOk returns a tuple with the ReclaimPolicy field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetReclaimPolicy

`func (o *InfrastructureVolume) SetReclaimPolicy(v string)`

SetReclaimPolicy sets ReclaimPolicy field to given value.

### HasReclaimPolicy

`func (o *InfrastructureVolume) HasReclaimPolicy() bool`

HasReclaimPolicy returns a boolean if a field has been set.

### GetRef

`func (o *InfrastructureVolume) GetRef() InfrastructureObjectReference`

GetRef returns the Ref field if non-nil, zero value otherwise.

### GetRefOk

`func (o *InfrastructureVolume) GetRefOk() (*InfrastructureObjectReference, bool)`

GetRefOk returns a tuple with the Ref field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRef

`func (o *InfrastructureVolume) SetRef(v InfrastructureObjectReference)`

SetRef sets Ref field to given value.

### HasRef

`func (o *InfrastructureVolume) HasRef() bool`

HasRef returns a boolean if a field has been set.

### GetVolumeHandle

`func (o *InfrastructureVolume) GetVolumeHandle() string`

GetVolumeHandle returns the VolumeHandle field if non-nil, zero value otherwise.

### GetVolumeHandleOk

`func (o *InfrastructureVolume) GetVolumeHandleOk() (*string, bool)`

GetVolumeHandleOk returns a tuple with the VolumeHandle field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVolumeHandle

`func (o *InfrastructureVolume) SetVolumeHandle(v string)`

SetVolumeHandle sets VolumeHandle field to given value.

### HasVolumeHandle

`func (o *InfrastructureVolume) HasVolumeHandle() bool`

HasVolumeHandle returns a boolean if a field has been set.

### GetVolumeMode

`func (o *InfrastructureVolume) GetVolumeMode() string`

GetVolumeMode returns the VolumeMode field if non-nil, zero value otherwise.

### GetVolumeModeOk

`func (o *InfrastructureVolume) GetVolumeModeOk() (*string, bool)`

GetVolumeModeOk returns a tuple with the VolumeMode field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVolumeMode

`func (o *InfrastructureVolume) SetVolumeMode(v string)`

SetVolumeMode sets VolumeMode field to given value.

### HasVolumeMode

`func (o *InfrastructureVolume) HasVolumeMode() bool`

HasVolumeMode returns a boolean if a field has been set.

### GetZones

`func (o *InfrastructureVolume) GetZones() []string`

GetZones returns the Zones field if non-nil, zero value otherwise.

### GetZonesOk

`func (o *InfrastructureVolume) GetZonesOk() (*[]string, bool)`

GetZonesOk returns a tuple with the Zones field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetZones

`func (o *InfrastructureVolume) SetZones(v []string)`

SetZones sets Zones field to given value.

### HasZones

`func (o *InfrastructureVolume) HasZones() bool`

HasZones returns a boolean if a field has been set.

### GetZonesTruncation

`func (o *InfrastructureVolume) GetZonesTruncation() CheckpointTruncation`

GetZonesTruncation returns the ZonesTruncation field if non-nil, zero value otherwise.

### GetZonesTruncationOk

`func (o *InfrastructureVolume) GetZonesTruncationOk() (*CheckpointTruncation, bool)`

GetZonesTruncationOk returns a tuple with the ZonesTruncation field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetZonesTruncation

`func (o *InfrastructureVolume) SetZonesTruncation(v CheckpointTruncation)`

SetZonesTruncation sets ZonesTruncation field to given value.

### HasZonesTruncation

`func (o *InfrastructureVolume) HasZonesTruncation() bool`

HasZonesTruncation returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


