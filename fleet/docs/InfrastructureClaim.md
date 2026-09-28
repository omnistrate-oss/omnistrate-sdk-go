# InfrastructureClaim

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**AccessModes** | Pointer to **[]string** |  | [optional] 
**AccessModesTruncation** | Pointer to [**CheckpointTruncation**](CheckpointTruncation.md) |  | [optional] 
**BindingMode** | Pointer to **string** |  | [optional] 
**Capacity** | Pointer to **string** |  | [optional] 
**Controller** | Pointer to [**InfrastructureObjectReference**](InfrastructureObjectReference.md) |  | [optional] 
**Phase** | Pointer to **string** |  | [optional] 
**Provisioner** | Pointer to **string** |  | [optional] 
**Ref** | Pointer to [**InfrastructureObjectReference**](InfrastructureObjectReference.md) |  | [optional] 
**RequestedCapacity** | Pointer to **string** |  | [optional] 
**StorageClass** | Pointer to **string** |  | [optional] 
**StorageClassAvailability** | Pointer to [**InfrastructureAvailability**](InfrastructureAvailability.md) |  | [optional] 
**Volume** | Pointer to [**InfrastructureVolume**](InfrastructureVolume.md) |  | [optional] 

## Methods

### NewInfrastructureClaim

`func NewInfrastructureClaim() *InfrastructureClaim`

NewInfrastructureClaim instantiates a new InfrastructureClaim object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewInfrastructureClaimWithDefaults

`func NewInfrastructureClaimWithDefaults() *InfrastructureClaim`

NewInfrastructureClaimWithDefaults instantiates a new InfrastructureClaim object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetAccessModes

`func (o *InfrastructureClaim) GetAccessModes() []string`

GetAccessModes returns the AccessModes field if non-nil, zero value otherwise.

### GetAccessModesOk

`func (o *InfrastructureClaim) GetAccessModesOk() (*[]string, bool)`

GetAccessModesOk returns a tuple with the AccessModes field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAccessModes

`func (o *InfrastructureClaim) SetAccessModes(v []string)`

SetAccessModes sets AccessModes field to given value.

### HasAccessModes

`func (o *InfrastructureClaim) HasAccessModes() bool`

HasAccessModes returns a boolean if a field has been set.

### GetAccessModesTruncation

`func (o *InfrastructureClaim) GetAccessModesTruncation() CheckpointTruncation`

GetAccessModesTruncation returns the AccessModesTruncation field if non-nil, zero value otherwise.

### GetAccessModesTruncationOk

`func (o *InfrastructureClaim) GetAccessModesTruncationOk() (*CheckpointTruncation, bool)`

GetAccessModesTruncationOk returns a tuple with the AccessModesTruncation field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAccessModesTruncation

`func (o *InfrastructureClaim) SetAccessModesTruncation(v CheckpointTruncation)`

SetAccessModesTruncation sets AccessModesTruncation field to given value.

### HasAccessModesTruncation

`func (o *InfrastructureClaim) HasAccessModesTruncation() bool`

HasAccessModesTruncation returns a boolean if a field has been set.

### GetBindingMode

`func (o *InfrastructureClaim) GetBindingMode() string`

GetBindingMode returns the BindingMode field if non-nil, zero value otherwise.

### GetBindingModeOk

`func (o *InfrastructureClaim) GetBindingModeOk() (*string, bool)`

GetBindingModeOk returns a tuple with the BindingMode field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBindingMode

`func (o *InfrastructureClaim) SetBindingMode(v string)`

SetBindingMode sets BindingMode field to given value.

### HasBindingMode

`func (o *InfrastructureClaim) HasBindingMode() bool`

HasBindingMode returns a boolean if a field has been set.

### GetCapacity

`func (o *InfrastructureClaim) GetCapacity() string`

GetCapacity returns the Capacity field if non-nil, zero value otherwise.

### GetCapacityOk

`func (o *InfrastructureClaim) GetCapacityOk() (*string, bool)`

GetCapacityOk returns a tuple with the Capacity field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCapacity

`func (o *InfrastructureClaim) SetCapacity(v string)`

SetCapacity sets Capacity field to given value.

### HasCapacity

`func (o *InfrastructureClaim) HasCapacity() bool`

HasCapacity returns a boolean if a field has been set.

### GetController

`func (o *InfrastructureClaim) GetController() InfrastructureObjectReference`

GetController returns the Controller field if non-nil, zero value otherwise.

### GetControllerOk

`func (o *InfrastructureClaim) GetControllerOk() (*InfrastructureObjectReference, bool)`

GetControllerOk returns a tuple with the Controller field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetController

`func (o *InfrastructureClaim) SetController(v InfrastructureObjectReference)`

SetController sets Controller field to given value.

### HasController

`func (o *InfrastructureClaim) HasController() bool`

HasController returns a boolean if a field has been set.

### GetPhase

`func (o *InfrastructureClaim) GetPhase() string`

GetPhase returns the Phase field if non-nil, zero value otherwise.

### GetPhaseOk

`func (o *InfrastructureClaim) GetPhaseOk() (*string, bool)`

GetPhaseOk returns a tuple with the Phase field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPhase

`func (o *InfrastructureClaim) SetPhase(v string)`

SetPhase sets Phase field to given value.

### HasPhase

`func (o *InfrastructureClaim) HasPhase() bool`

HasPhase returns a boolean if a field has been set.

### GetProvisioner

`func (o *InfrastructureClaim) GetProvisioner() string`

GetProvisioner returns the Provisioner field if non-nil, zero value otherwise.

### GetProvisionerOk

`func (o *InfrastructureClaim) GetProvisionerOk() (*string, bool)`

GetProvisionerOk returns a tuple with the Provisioner field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetProvisioner

`func (o *InfrastructureClaim) SetProvisioner(v string)`

SetProvisioner sets Provisioner field to given value.

### HasProvisioner

`func (o *InfrastructureClaim) HasProvisioner() bool`

HasProvisioner returns a boolean if a field has been set.

### GetRef

`func (o *InfrastructureClaim) GetRef() InfrastructureObjectReference`

GetRef returns the Ref field if non-nil, zero value otherwise.

### GetRefOk

`func (o *InfrastructureClaim) GetRefOk() (*InfrastructureObjectReference, bool)`

GetRefOk returns a tuple with the Ref field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRef

`func (o *InfrastructureClaim) SetRef(v InfrastructureObjectReference)`

SetRef sets Ref field to given value.

### HasRef

`func (o *InfrastructureClaim) HasRef() bool`

HasRef returns a boolean if a field has been set.

### GetRequestedCapacity

`func (o *InfrastructureClaim) GetRequestedCapacity() string`

GetRequestedCapacity returns the RequestedCapacity field if non-nil, zero value otherwise.

### GetRequestedCapacityOk

`func (o *InfrastructureClaim) GetRequestedCapacityOk() (*string, bool)`

GetRequestedCapacityOk returns a tuple with the RequestedCapacity field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRequestedCapacity

`func (o *InfrastructureClaim) SetRequestedCapacity(v string)`

SetRequestedCapacity sets RequestedCapacity field to given value.

### HasRequestedCapacity

`func (o *InfrastructureClaim) HasRequestedCapacity() bool`

HasRequestedCapacity returns a boolean if a field has been set.

### GetStorageClass

`func (o *InfrastructureClaim) GetStorageClass() string`

GetStorageClass returns the StorageClass field if non-nil, zero value otherwise.

### GetStorageClassOk

`func (o *InfrastructureClaim) GetStorageClassOk() (*string, bool)`

GetStorageClassOk returns a tuple with the StorageClass field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStorageClass

`func (o *InfrastructureClaim) SetStorageClass(v string)`

SetStorageClass sets StorageClass field to given value.

### HasStorageClass

`func (o *InfrastructureClaim) HasStorageClass() bool`

HasStorageClass returns a boolean if a field has been set.

### GetStorageClassAvailability

`func (o *InfrastructureClaim) GetStorageClassAvailability() InfrastructureAvailability`

GetStorageClassAvailability returns the StorageClassAvailability field if non-nil, zero value otherwise.

### GetStorageClassAvailabilityOk

`func (o *InfrastructureClaim) GetStorageClassAvailabilityOk() (*InfrastructureAvailability, bool)`

GetStorageClassAvailabilityOk returns a tuple with the StorageClassAvailability field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStorageClassAvailability

`func (o *InfrastructureClaim) SetStorageClassAvailability(v InfrastructureAvailability)`

SetStorageClassAvailability sets StorageClassAvailability field to given value.

### HasStorageClassAvailability

`func (o *InfrastructureClaim) HasStorageClassAvailability() bool`

HasStorageClassAvailability returns a boolean if a field has been set.

### GetVolume

`func (o *InfrastructureClaim) GetVolume() InfrastructureVolume`

GetVolume returns the Volume field if non-nil, zero value otherwise.

### GetVolumeOk

`func (o *InfrastructureClaim) GetVolumeOk() (*InfrastructureVolume, bool)`

GetVolumeOk returns a tuple with the Volume field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVolume

`func (o *InfrastructureClaim) SetVolume(v InfrastructureVolume)`

SetVolume sets Volume field to given value.

### HasVolume

`func (o *InfrastructureClaim) HasVolume() bool`

HasVolume returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


