# InfrastructureNode

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Availability** | Pointer to [**InfrastructureAvailability**](InfrastructureAvailability.md) |  | [optional] 
**Capacity** | Pointer to [**InfrastructureResources**](InfrastructureResources.md) |  | [optional] 
**CreatedAt** | Pointer to **string** |  | [optional] 
**Gpu** | Pointer to [**InfrastructureGPU**](InfrastructureGPU.md) |  | [optional] 
**InstancePodCount** | Pointer to **int64** |  | [optional] 
**InstanceType** | Pointer to **string** |  | [optional] 
**InternalIP** | Pointer to **string** |  | [optional] 
**KubeletVersion** | Pointer to **string** |  | [optional] 
**Labels** | Pointer to [**[]InfrastructureKeyValue**](InfrastructureKeyValue.md) |  | [optional] 
**LabelsTruncation** | Pointer to [**CheckpointTruncation**](CheckpointTruncation.md) |  | [optional] 
**PodNames** | Pointer to **[]string** |  | [optional] 
**PodsTruncation** | Pointer to [**CheckpointTruncation**](CheckpointTruncation.md) |  | [optional] 
**Pool** | Pointer to **string** |  | [optional] 
**ProviderID** | Pointer to **string** |  | [optional] 
**PurchaseModel** | Pointer to **string** |  | [optional] 
**Ready** | Pointer to **bool** |  | [optional] 
**Ref** | Pointer to [**InfrastructureObjectReference**](InfrastructureObjectReference.md) |  | [optional] 
**Region** | Pointer to **string** |  | [optional] 
**Roles** | Pointer to **[]string** |  | [optional] 
**RolesTruncation** | Pointer to [**CheckpointTruncation**](CheckpointTruncation.md) |  | [optional] 
**Taints** | Pointer to [**[]InfrastructureSchedulingRule**](InfrastructureSchedulingRule.md) |  | [optional] 
**TaintsTruncation** | Pointer to [**CheckpointTruncation**](CheckpointTruncation.md) |  | [optional] 
**Zone** | Pointer to **string** |  | [optional] 

## Methods

### NewInfrastructureNode

`func NewInfrastructureNode() *InfrastructureNode`

NewInfrastructureNode instantiates a new InfrastructureNode object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewInfrastructureNodeWithDefaults

`func NewInfrastructureNodeWithDefaults() *InfrastructureNode`

NewInfrastructureNodeWithDefaults instantiates a new InfrastructureNode object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetAvailability

`func (o *InfrastructureNode) GetAvailability() InfrastructureAvailability`

GetAvailability returns the Availability field if non-nil, zero value otherwise.

### GetAvailabilityOk

`func (o *InfrastructureNode) GetAvailabilityOk() (*InfrastructureAvailability, bool)`

GetAvailabilityOk returns a tuple with the Availability field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAvailability

`func (o *InfrastructureNode) SetAvailability(v InfrastructureAvailability)`

SetAvailability sets Availability field to given value.

### HasAvailability

`func (o *InfrastructureNode) HasAvailability() bool`

HasAvailability returns a boolean if a field has been set.

### GetCapacity

`func (o *InfrastructureNode) GetCapacity() InfrastructureResources`

GetCapacity returns the Capacity field if non-nil, zero value otherwise.

### GetCapacityOk

`func (o *InfrastructureNode) GetCapacityOk() (*InfrastructureResources, bool)`

GetCapacityOk returns a tuple with the Capacity field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCapacity

`func (o *InfrastructureNode) SetCapacity(v InfrastructureResources)`

SetCapacity sets Capacity field to given value.

### HasCapacity

`func (o *InfrastructureNode) HasCapacity() bool`

HasCapacity returns a boolean if a field has been set.

### GetCreatedAt

`func (o *InfrastructureNode) GetCreatedAt() string`

GetCreatedAt returns the CreatedAt field if non-nil, zero value otherwise.

### GetCreatedAtOk

`func (o *InfrastructureNode) GetCreatedAtOk() (*string, bool)`

GetCreatedAtOk returns a tuple with the CreatedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCreatedAt

`func (o *InfrastructureNode) SetCreatedAt(v string)`

SetCreatedAt sets CreatedAt field to given value.

### HasCreatedAt

`func (o *InfrastructureNode) HasCreatedAt() bool`

HasCreatedAt returns a boolean if a field has been set.

### GetGpu

`func (o *InfrastructureNode) GetGpu() InfrastructureGPU`

GetGpu returns the Gpu field if non-nil, zero value otherwise.

### GetGpuOk

`func (o *InfrastructureNode) GetGpuOk() (*InfrastructureGPU, bool)`

GetGpuOk returns a tuple with the Gpu field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetGpu

`func (o *InfrastructureNode) SetGpu(v InfrastructureGPU)`

SetGpu sets Gpu field to given value.

### HasGpu

`func (o *InfrastructureNode) HasGpu() bool`

HasGpu returns a boolean if a field has been set.

### GetInstancePodCount

`func (o *InfrastructureNode) GetInstancePodCount() int64`

GetInstancePodCount returns the InstancePodCount field if non-nil, zero value otherwise.

### GetInstancePodCountOk

`func (o *InfrastructureNode) GetInstancePodCountOk() (*int64, bool)`

GetInstancePodCountOk returns a tuple with the InstancePodCount field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetInstancePodCount

`func (o *InfrastructureNode) SetInstancePodCount(v int64)`

SetInstancePodCount sets InstancePodCount field to given value.

### HasInstancePodCount

`func (o *InfrastructureNode) HasInstancePodCount() bool`

HasInstancePodCount returns a boolean if a field has been set.

### GetInstanceType

`func (o *InfrastructureNode) GetInstanceType() string`

GetInstanceType returns the InstanceType field if non-nil, zero value otherwise.

### GetInstanceTypeOk

`func (o *InfrastructureNode) GetInstanceTypeOk() (*string, bool)`

GetInstanceTypeOk returns a tuple with the InstanceType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetInstanceType

`func (o *InfrastructureNode) SetInstanceType(v string)`

SetInstanceType sets InstanceType field to given value.

### HasInstanceType

`func (o *InfrastructureNode) HasInstanceType() bool`

HasInstanceType returns a boolean if a field has been set.

### GetInternalIP

`func (o *InfrastructureNode) GetInternalIP() string`

GetInternalIP returns the InternalIP field if non-nil, zero value otherwise.

### GetInternalIPOk

`func (o *InfrastructureNode) GetInternalIPOk() (*string, bool)`

GetInternalIPOk returns a tuple with the InternalIP field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetInternalIP

`func (o *InfrastructureNode) SetInternalIP(v string)`

SetInternalIP sets InternalIP field to given value.

### HasInternalIP

`func (o *InfrastructureNode) HasInternalIP() bool`

HasInternalIP returns a boolean if a field has been set.

### GetKubeletVersion

`func (o *InfrastructureNode) GetKubeletVersion() string`

GetKubeletVersion returns the KubeletVersion field if non-nil, zero value otherwise.

### GetKubeletVersionOk

`func (o *InfrastructureNode) GetKubeletVersionOk() (*string, bool)`

GetKubeletVersionOk returns a tuple with the KubeletVersion field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetKubeletVersion

`func (o *InfrastructureNode) SetKubeletVersion(v string)`

SetKubeletVersion sets KubeletVersion field to given value.

### HasKubeletVersion

`func (o *InfrastructureNode) HasKubeletVersion() bool`

HasKubeletVersion returns a boolean if a field has been set.

### GetLabels

`func (o *InfrastructureNode) GetLabels() []InfrastructureKeyValue`

GetLabels returns the Labels field if non-nil, zero value otherwise.

### GetLabelsOk

`func (o *InfrastructureNode) GetLabelsOk() (*[]InfrastructureKeyValue, bool)`

GetLabelsOk returns a tuple with the Labels field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLabels

`func (o *InfrastructureNode) SetLabels(v []InfrastructureKeyValue)`

SetLabels sets Labels field to given value.

### HasLabels

`func (o *InfrastructureNode) HasLabels() bool`

HasLabels returns a boolean if a field has been set.

### GetLabelsTruncation

`func (o *InfrastructureNode) GetLabelsTruncation() CheckpointTruncation`

GetLabelsTruncation returns the LabelsTruncation field if non-nil, zero value otherwise.

### GetLabelsTruncationOk

`func (o *InfrastructureNode) GetLabelsTruncationOk() (*CheckpointTruncation, bool)`

GetLabelsTruncationOk returns a tuple with the LabelsTruncation field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLabelsTruncation

`func (o *InfrastructureNode) SetLabelsTruncation(v CheckpointTruncation)`

SetLabelsTruncation sets LabelsTruncation field to given value.

### HasLabelsTruncation

`func (o *InfrastructureNode) HasLabelsTruncation() bool`

HasLabelsTruncation returns a boolean if a field has been set.

### GetPodNames

`func (o *InfrastructureNode) GetPodNames() []string`

GetPodNames returns the PodNames field if non-nil, zero value otherwise.

### GetPodNamesOk

`func (o *InfrastructureNode) GetPodNamesOk() (*[]string, bool)`

GetPodNamesOk returns a tuple with the PodNames field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPodNames

`func (o *InfrastructureNode) SetPodNames(v []string)`

SetPodNames sets PodNames field to given value.

### HasPodNames

`func (o *InfrastructureNode) HasPodNames() bool`

HasPodNames returns a boolean if a field has been set.

### GetPodsTruncation

`func (o *InfrastructureNode) GetPodsTruncation() CheckpointTruncation`

GetPodsTruncation returns the PodsTruncation field if non-nil, zero value otherwise.

### GetPodsTruncationOk

`func (o *InfrastructureNode) GetPodsTruncationOk() (*CheckpointTruncation, bool)`

GetPodsTruncationOk returns a tuple with the PodsTruncation field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPodsTruncation

`func (o *InfrastructureNode) SetPodsTruncation(v CheckpointTruncation)`

SetPodsTruncation sets PodsTruncation field to given value.

### HasPodsTruncation

`func (o *InfrastructureNode) HasPodsTruncation() bool`

HasPodsTruncation returns a boolean if a field has been set.

### GetPool

`func (o *InfrastructureNode) GetPool() string`

GetPool returns the Pool field if non-nil, zero value otherwise.

### GetPoolOk

`func (o *InfrastructureNode) GetPoolOk() (*string, bool)`

GetPoolOk returns a tuple with the Pool field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPool

`func (o *InfrastructureNode) SetPool(v string)`

SetPool sets Pool field to given value.

### HasPool

`func (o *InfrastructureNode) HasPool() bool`

HasPool returns a boolean if a field has been set.

### GetProviderID

`func (o *InfrastructureNode) GetProviderID() string`

GetProviderID returns the ProviderID field if non-nil, zero value otherwise.

### GetProviderIDOk

`func (o *InfrastructureNode) GetProviderIDOk() (*string, bool)`

GetProviderIDOk returns a tuple with the ProviderID field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetProviderID

`func (o *InfrastructureNode) SetProviderID(v string)`

SetProviderID sets ProviderID field to given value.

### HasProviderID

`func (o *InfrastructureNode) HasProviderID() bool`

HasProviderID returns a boolean if a field has been set.

### GetPurchaseModel

`func (o *InfrastructureNode) GetPurchaseModel() string`

GetPurchaseModel returns the PurchaseModel field if non-nil, zero value otherwise.

### GetPurchaseModelOk

`func (o *InfrastructureNode) GetPurchaseModelOk() (*string, bool)`

GetPurchaseModelOk returns a tuple with the PurchaseModel field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPurchaseModel

`func (o *InfrastructureNode) SetPurchaseModel(v string)`

SetPurchaseModel sets PurchaseModel field to given value.

### HasPurchaseModel

`func (o *InfrastructureNode) HasPurchaseModel() bool`

HasPurchaseModel returns a boolean if a field has been set.

### GetReady

`func (o *InfrastructureNode) GetReady() bool`

GetReady returns the Ready field if non-nil, zero value otherwise.

### GetReadyOk

`func (o *InfrastructureNode) GetReadyOk() (*bool, bool)`

GetReadyOk returns a tuple with the Ready field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetReady

`func (o *InfrastructureNode) SetReady(v bool)`

SetReady sets Ready field to given value.

### HasReady

`func (o *InfrastructureNode) HasReady() bool`

HasReady returns a boolean if a field has been set.

### GetRef

`func (o *InfrastructureNode) GetRef() InfrastructureObjectReference`

GetRef returns the Ref field if non-nil, zero value otherwise.

### GetRefOk

`func (o *InfrastructureNode) GetRefOk() (*InfrastructureObjectReference, bool)`

GetRefOk returns a tuple with the Ref field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRef

`func (o *InfrastructureNode) SetRef(v InfrastructureObjectReference)`

SetRef sets Ref field to given value.

### HasRef

`func (o *InfrastructureNode) HasRef() bool`

HasRef returns a boolean if a field has been set.

### GetRegion

`func (o *InfrastructureNode) GetRegion() string`

GetRegion returns the Region field if non-nil, zero value otherwise.

### GetRegionOk

`func (o *InfrastructureNode) GetRegionOk() (*string, bool)`

GetRegionOk returns a tuple with the Region field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRegion

`func (o *InfrastructureNode) SetRegion(v string)`

SetRegion sets Region field to given value.

### HasRegion

`func (o *InfrastructureNode) HasRegion() bool`

HasRegion returns a boolean if a field has been set.

### GetRoles

`func (o *InfrastructureNode) GetRoles() []string`

GetRoles returns the Roles field if non-nil, zero value otherwise.

### GetRolesOk

`func (o *InfrastructureNode) GetRolesOk() (*[]string, bool)`

GetRolesOk returns a tuple with the Roles field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRoles

`func (o *InfrastructureNode) SetRoles(v []string)`

SetRoles sets Roles field to given value.

### HasRoles

`func (o *InfrastructureNode) HasRoles() bool`

HasRoles returns a boolean if a field has been set.

### GetRolesTruncation

`func (o *InfrastructureNode) GetRolesTruncation() CheckpointTruncation`

GetRolesTruncation returns the RolesTruncation field if non-nil, zero value otherwise.

### GetRolesTruncationOk

`func (o *InfrastructureNode) GetRolesTruncationOk() (*CheckpointTruncation, bool)`

GetRolesTruncationOk returns a tuple with the RolesTruncation field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRolesTruncation

`func (o *InfrastructureNode) SetRolesTruncation(v CheckpointTruncation)`

SetRolesTruncation sets RolesTruncation field to given value.

### HasRolesTruncation

`func (o *InfrastructureNode) HasRolesTruncation() bool`

HasRolesTruncation returns a boolean if a field has been set.

### GetTaints

`func (o *InfrastructureNode) GetTaints() []InfrastructureSchedulingRule`

GetTaints returns the Taints field if non-nil, zero value otherwise.

### GetTaintsOk

`func (o *InfrastructureNode) GetTaintsOk() (*[]InfrastructureSchedulingRule, bool)`

GetTaintsOk returns a tuple with the Taints field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTaints

`func (o *InfrastructureNode) SetTaints(v []InfrastructureSchedulingRule)`

SetTaints sets Taints field to given value.

### HasTaints

`func (o *InfrastructureNode) HasTaints() bool`

HasTaints returns a boolean if a field has been set.

### GetTaintsTruncation

`func (o *InfrastructureNode) GetTaintsTruncation() CheckpointTruncation`

GetTaintsTruncation returns the TaintsTruncation field if non-nil, zero value otherwise.

### GetTaintsTruncationOk

`func (o *InfrastructureNode) GetTaintsTruncationOk() (*CheckpointTruncation, bool)`

GetTaintsTruncationOk returns a tuple with the TaintsTruncation field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTaintsTruncation

`func (o *InfrastructureNode) SetTaintsTruncation(v CheckpointTruncation)`

SetTaintsTruncation sets TaintsTruncation field to given value.

### HasTaintsTruncation

`func (o *InfrastructureNode) HasTaintsTruncation() bool`

HasTaintsTruncation returns a boolean if a field has been set.

### GetZone

`func (o *InfrastructureNode) GetZone() string`

GetZone returns the Zone field if non-nil, zero value otherwise.

### GetZoneOk

`func (o *InfrastructureNode) GetZoneOk() (*string, bool)`

GetZoneOk returns a tuple with the Zone field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetZone

`func (o *InfrastructureNode) SetZone(v string)`

SetZone sets Zone field to given value.

### HasZone

`func (o *InfrastructureNode) HasZone() bool`

HasZone returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


