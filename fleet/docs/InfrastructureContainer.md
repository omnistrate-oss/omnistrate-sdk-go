# InfrastructureContainer

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Kind** | Pointer to **string** |  | [optional] 
**Limits** | Pointer to [**InfrastructureResources**](InfrastructureResources.md) |  | [optional] 
**Mounts** | Pointer to [**[]InfrastructureMount**](InfrastructureMount.md) |  | [optional] 
**MountsTruncation** | Pointer to [**CheckpointTruncation**](CheckpointTruncation.md) |  | [optional] 
**Name** | Pointer to **string** |  | [optional] 
**Requests** | Pointer to [**InfrastructureResources**](InfrastructureResources.md) |  | [optional] 

## Methods

### NewInfrastructureContainer

`func NewInfrastructureContainer() *InfrastructureContainer`

NewInfrastructureContainer instantiates a new InfrastructureContainer object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewInfrastructureContainerWithDefaults

`func NewInfrastructureContainerWithDefaults() *InfrastructureContainer`

NewInfrastructureContainerWithDefaults instantiates a new InfrastructureContainer object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetKind

`func (o *InfrastructureContainer) GetKind() string`

GetKind returns the Kind field if non-nil, zero value otherwise.

### GetKindOk

`func (o *InfrastructureContainer) GetKindOk() (*string, bool)`

GetKindOk returns a tuple with the Kind field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetKind

`func (o *InfrastructureContainer) SetKind(v string)`

SetKind sets Kind field to given value.

### HasKind

`func (o *InfrastructureContainer) HasKind() bool`

HasKind returns a boolean if a field has been set.

### GetLimits

`func (o *InfrastructureContainer) GetLimits() InfrastructureResources`

GetLimits returns the Limits field if non-nil, zero value otherwise.

### GetLimitsOk

`func (o *InfrastructureContainer) GetLimitsOk() (*InfrastructureResources, bool)`

GetLimitsOk returns a tuple with the Limits field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLimits

`func (o *InfrastructureContainer) SetLimits(v InfrastructureResources)`

SetLimits sets Limits field to given value.

### HasLimits

`func (o *InfrastructureContainer) HasLimits() bool`

HasLimits returns a boolean if a field has been set.

### GetMounts

`func (o *InfrastructureContainer) GetMounts() []InfrastructureMount`

GetMounts returns the Mounts field if non-nil, zero value otherwise.

### GetMountsOk

`func (o *InfrastructureContainer) GetMountsOk() (*[]InfrastructureMount, bool)`

GetMountsOk returns a tuple with the Mounts field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMounts

`func (o *InfrastructureContainer) SetMounts(v []InfrastructureMount)`

SetMounts sets Mounts field to given value.

### HasMounts

`func (o *InfrastructureContainer) HasMounts() bool`

HasMounts returns a boolean if a field has been set.

### GetMountsTruncation

`func (o *InfrastructureContainer) GetMountsTruncation() CheckpointTruncation`

GetMountsTruncation returns the MountsTruncation field if non-nil, zero value otherwise.

### GetMountsTruncationOk

`func (o *InfrastructureContainer) GetMountsTruncationOk() (*CheckpointTruncation, bool)`

GetMountsTruncationOk returns a tuple with the MountsTruncation field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMountsTruncation

`func (o *InfrastructureContainer) SetMountsTruncation(v CheckpointTruncation)`

SetMountsTruncation sets MountsTruncation field to given value.

### HasMountsTruncation

`func (o *InfrastructureContainer) HasMountsTruncation() bool`

HasMountsTruncation returns a boolean if a field has been set.

### GetName

`func (o *InfrastructureContainer) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *InfrastructureContainer) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *InfrastructureContainer) SetName(v string)`

SetName sets Name field to given value.

### HasName

`func (o *InfrastructureContainer) HasName() bool`

HasName returns a boolean if a field has been set.

### GetRequests

`func (o *InfrastructureContainer) GetRequests() InfrastructureResources`

GetRequests returns the Requests field if non-nil, zero value otherwise.

### GetRequestsOk

`func (o *InfrastructureContainer) GetRequestsOk() (*InfrastructureResources, bool)`

GetRequestsOk returns a tuple with the Requests field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRequests

`func (o *InfrastructureContainer) SetRequests(v InfrastructureResources)`

SetRequests sets Requests field to given value.

### HasRequests

`func (o *InfrastructureContainer) HasRequests() bool`

HasRequests returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


