# InfrastructurePod

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Conditions** | Pointer to [**[]InfrastructureCondition**](InfrastructureCondition.md) |  | [optional] 
**ConditionsTruncation** | Pointer to [**CheckpointTruncation**](CheckpointTruncation.md) |  | [optional] 
**Containers** | Pointer to [**[]InfrastructureContainer**](InfrastructureContainer.md) |  | [optional] 
**ContainersTruncation** | Pointer to [**CheckpointTruncation**](CheckpointTruncation.md) |  | [optional] 
**Controller** | Pointer to [**InfrastructureObjectReference**](InfrastructureObjectReference.md) |  | [optional] 
**CreatedAt** | Pointer to **string** |  | [optional] 
**DeletedAt** | Pointer to **string** |  | [optional] 
**HostPaths** | Pointer to **[]string** |  | [optional] 
**HostPathsTruncation** | Pointer to [**CheckpointTruncation**](CheckpointTruncation.md) |  | [optional] 
**NodeName** | Pointer to **string** |  | [optional] 
**Owner** | Pointer to [**InfrastructureObjectReference**](InfrastructureObjectReference.md) |  | [optional] 
**Phase** | Pointer to **string** |  | [optional] 
**PodAntiAffinity** | Pointer to **bool** |  | [optional] 
**PodIP** | Pointer to **string** |  | [optional] 
**QosClass** | Pointer to **string** |  | [optional] 
**Ready** | Pointer to **bool** |  | [optional] 
**ReadyContainers** | Pointer to **int64** |  | [optional] 
**Reason** | Pointer to **string** |  | [optional] 
**Ref** | Pointer to [**InfrastructureObjectReference**](InfrastructureObjectReference.md) |  | [optional] 
**References** | Pointer to [**[]InfrastructureObjectReference**](InfrastructureObjectReference.md) |  | [optional] 
**ReferencesTruncation** | Pointer to [**CheckpointTruncation**](CheckpointTruncation.md) |  | [optional] 
**Restarts** | Pointer to **int64** |  | [optional] 
**Tolerations** | Pointer to [**[]InfrastructureSchedulingRule**](InfrastructureSchedulingRule.md) |  | [optional] 
**TolerationsTruncation** | Pointer to [**CheckpointTruncation**](CheckpointTruncation.md) |  | [optional] 
**TotalContainers** | Pointer to **int64** |  | [optional] 

## Methods

### NewInfrastructurePod

`func NewInfrastructurePod() *InfrastructurePod`

NewInfrastructurePod instantiates a new InfrastructurePod object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewInfrastructurePodWithDefaults

`func NewInfrastructurePodWithDefaults() *InfrastructurePod`

NewInfrastructurePodWithDefaults instantiates a new InfrastructurePod object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetConditions

`func (o *InfrastructurePod) GetConditions() []InfrastructureCondition`

GetConditions returns the Conditions field if non-nil, zero value otherwise.

### GetConditionsOk

`func (o *InfrastructurePod) GetConditionsOk() (*[]InfrastructureCondition, bool)`

GetConditionsOk returns a tuple with the Conditions field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetConditions

`func (o *InfrastructurePod) SetConditions(v []InfrastructureCondition)`

SetConditions sets Conditions field to given value.

### HasConditions

`func (o *InfrastructurePod) HasConditions() bool`

HasConditions returns a boolean if a field has been set.

### GetConditionsTruncation

`func (o *InfrastructurePod) GetConditionsTruncation() CheckpointTruncation`

GetConditionsTruncation returns the ConditionsTruncation field if non-nil, zero value otherwise.

### GetConditionsTruncationOk

`func (o *InfrastructurePod) GetConditionsTruncationOk() (*CheckpointTruncation, bool)`

GetConditionsTruncationOk returns a tuple with the ConditionsTruncation field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetConditionsTruncation

`func (o *InfrastructurePod) SetConditionsTruncation(v CheckpointTruncation)`

SetConditionsTruncation sets ConditionsTruncation field to given value.

### HasConditionsTruncation

`func (o *InfrastructurePod) HasConditionsTruncation() bool`

HasConditionsTruncation returns a boolean if a field has been set.

### GetContainers

`func (o *InfrastructurePod) GetContainers() []InfrastructureContainer`

GetContainers returns the Containers field if non-nil, zero value otherwise.

### GetContainersOk

`func (o *InfrastructurePod) GetContainersOk() (*[]InfrastructureContainer, bool)`

GetContainersOk returns a tuple with the Containers field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetContainers

`func (o *InfrastructurePod) SetContainers(v []InfrastructureContainer)`

SetContainers sets Containers field to given value.

### HasContainers

`func (o *InfrastructurePod) HasContainers() bool`

HasContainers returns a boolean if a field has been set.

### GetContainersTruncation

`func (o *InfrastructurePod) GetContainersTruncation() CheckpointTruncation`

GetContainersTruncation returns the ContainersTruncation field if non-nil, zero value otherwise.

### GetContainersTruncationOk

`func (o *InfrastructurePod) GetContainersTruncationOk() (*CheckpointTruncation, bool)`

GetContainersTruncationOk returns a tuple with the ContainersTruncation field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetContainersTruncation

`func (o *InfrastructurePod) SetContainersTruncation(v CheckpointTruncation)`

SetContainersTruncation sets ContainersTruncation field to given value.

### HasContainersTruncation

`func (o *InfrastructurePod) HasContainersTruncation() bool`

HasContainersTruncation returns a boolean if a field has been set.

### GetController

`func (o *InfrastructurePod) GetController() InfrastructureObjectReference`

GetController returns the Controller field if non-nil, zero value otherwise.

### GetControllerOk

`func (o *InfrastructurePod) GetControllerOk() (*InfrastructureObjectReference, bool)`

GetControllerOk returns a tuple with the Controller field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetController

`func (o *InfrastructurePod) SetController(v InfrastructureObjectReference)`

SetController sets Controller field to given value.

### HasController

`func (o *InfrastructurePod) HasController() bool`

HasController returns a boolean if a field has been set.

### GetCreatedAt

`func (o *InfrastructurePod) GetCreatedAt() string`

GetCreatedAt returns the CreatedAt field if non-nil, zero value otherwise.

### GetCreatedAtOk

`func (o *InfrastructurePod) GetCreatedAtOk() (*string, bool)`

GetCreatedAtOk returns a tuple with the CreatedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCreatedAt

`func (o *InfrastructurePod) SetCreatedAt(v string)`

SetCreatedAt sets CreatedAt field to given value.

### HasCreatedAt

`func (o *InfrastructurePod) HasCreatedAt() bool`

HasCreatedAt returns a boolean if a field has been set.

### GetDeletedAt

`func (o *InfrastructurePod) GetDeletedAt() string`

GetDeletedAt returns the DeletedAt field if non-nil, zero value otherwise.

### GetDeletedAtOk

`func (o *InfrastructurePod) GetDeletedAtOk() (*string, bool)`

GetDeletedAtOk returns a tuple with the DeletedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDeletedAt

`func (o *InfrastructurePod) SetDeletedAt(v string)`

SetDeletedAt sets DeletedAt field to given value.

### HasDeletedAt

`func (o *InfrastructurePod) HasDeletedAt() bool`

HasDeletedAt returns a boolean if a field has been set.

### GetHostPaths

`func (o *InfrastructurePod) GetHostPaths() []string`

GetHostPaths returns the HostPaths field if non-nil, zero value otherwise.

### GetHostPathsOk

`func (o *InfrastructurePod) GetHostPathsOk() (*[]string, bool)`

GetHostPathsOk returns a tuple with the HostPaths field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetHostPaths

`func (o *InfrastructurePod) SetHostPaths(v []string)`

SetHostPaths sets HostPaths field to given value.

### HasHostPaths

`func (o *InfrastructurePod) HasHostPaths() bool`

HasHostPaths returns a boolean if a field has been set.

### GetHostPathsTruncation

`func (o *InfrastructurePod) GetHostPathsTruncation() CheckpointTruncation`

GetHostPathsTruncation returns the HostPathsTruncation field if non-nil, zero value otherwise.

### GetHostPathsTruncationOk

`func (o *InfrastructurePod) GetHostPathsTruncationOk() (*CheckpointTruncation, bool)`

GetHostPathsTruncationOk returns a tuple with the HostPathsTruncation field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetHostPathsTruncation

`func (o *InfrastructurePod) SetHostPathsTruncation(v CheckpointTruncation)`

SetHostPathsTruncation sets HostPathsTruncation field to given value.

### HasHostPathsTruncation

`func (o *InfrastructurePod) HasHostPathsTruncation() bool`

HasHostPathsTruncation returns a boolean if a field has been set.

### GetNodeName

`func (o *InfrastructurePod) GetNodeName() string`

GetNodeName returns the NodeName field if non-nil, zero value otherwise.

### GetNodeNameOk

`func (o *InfrastructurePod) GetNodeNameOk() (*string, bool)`

GetNodeNameOk returns a tuple with the NodeName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNodeName

`func (o *InfrastructurePod) SetNodeName(v string)`

SetNodeName sets NodeName field to given value.

### HasNodeName

`func (o *InfrastructurePod) HasNodeName() bool`

HasNodeName returns a boolean if a field has been set.

### GetOwner

`func (o *InfrastructurePod) GetOwner() InfrastructureObjectReference`

GetOwner returns the Owner field if non-nil, zero value otherwise.

### GetOwnerOk

`func (o *InfrastructurePod) GetOwnerOk() (*InfrastructureObjectReference, bool)`

GetOwnerOk returns a tuple with the Owner field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOwner

`func (o *InfrastructurePod) SetOwner(v InfrastructureObjectReference)`

SetOwner sets Owner field to given value.

### HasOwner

`func (o *InfrastructurePod) HasOwner() bool`

HasOwner returns a boolean if a field has been set.

### GetPhase

`func (o *InfrastructurePod) GetPhase() string`

GetPhase returns the Phase field if non-nil, zero value otherwise.

### GetPhaseOk

`func (o *InfrastructurePod) GetPhaseOk() (*string, bool)`

GetPhaseOk returns a tuple with the Phase field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPhase

`func (o *InfrastructurePod) SetPhase(v string)`

SetPhase sets Phase field to given value.

### HasPhase

`func (o *InfrastructurePod) HasPhase() bool`

HasPhase returns a boolean if a field has been set.

### GetPodAntiAffinity

`func (o *InfrastructurePod) GetPodAntiAffinity() bool`

GetPodAntiAffinity returns the PodAntiAffinity field if non-nil, zero value otherwise.

### GetPodAntiAffinityOk

`func (o *InfrastructurePod) GetPodAntiAffinityOk() (*bool, bool)`

GetPodAntiAffinityOk returns a tuple with the PodAntiAffinity field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPodAntiAffinity

`func (o *InfrastructurePod) SetPodAntiAffinity(v bool)`

SetPodAntiAffinity sets PodAntiAffinity field to given value.

### HasPodAntiAffinity

`func (o *InfrastructurePod) HasPodAntiAffinity() bool`

HasPodAntiAffinity returns a boolean if a field has been set.

### GetPodIP

`func (o *InfrastructurePod) GetPodIP() string`

GetPodIP returns the PodIP field if non-nil, zero value otherwise.

### GetPodIPOk

`func (o *InfrastructurePod) GetPodIPOk() (*string, bool)`

GetPodIPOk returns a tuple with the PodIP field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPodIP

`func (o *InfrastructurePod) SetPodIP(v string)`

SetPodIP sets PodIP field to given value.

### HasPodIP

`func (o *InfrastructurePod) HasPodIP() bool`

HasPodIP returns a boolean if a field has been set.

### GetQosClass

`func (o *InfrastructurePod) GetQosClass() string`

GetQosClass returns the QosClass field if non-nil, zero value otherwise.

### GetQosClassOk

`func (o *InfrastructurePod) GetQosClassOk() (*string, bool)`

GetQosClassOk returns a tuple with the QosClass field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetQosClass

`func (o *InfrastructurePod) SetQosClass(v string)`

SetQosClass sets QosClass field to given value.

### HasQosClass

`func (o *InfrastructurePod) HasQosClass() bool`

HasQosClass returns a boolean if a field has been set.

### GetReady

`func (o *InfrastructurePod) GetReady() bool`

GetReady returns the Ready field if non-nil, zero value otherwise.

### GetReadyOk

`func (o *InfrastructurePod) GetReadyOk() (*bool, bool)`

GetReadyOk returns a tuple with the Ready field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetReady

`func (o *InfrastructurePod) SetReady(v bool)`

SetReady sets Ready field to given value.

### HasReady

`func (o *InfrastructurePod) HasReady() bool`

HasReady returns a boolean if a field has been set.

### GetReadyContainers

`func (o *InfrastructurePod) GetReadyContainers() int64`

GetReadyContainers returns the ReadyContainers field if non-nil, zero value otherwise.

### GetReadyContainersOk

`func (o *InfrastructurePod) GetReadyContainersOk() (*int64, bool)`

GetReadyContainersOk returns a tuple with the ReadyContainers field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetReadyContainers

`func (o *InfrastructurePod) SetReadyContainers(v int64)`

SetReadyContainers sets ReadyContainers field to given value.

### HasReadyContainers

`func (o *InfrastructurePod) HasReadyContainers() bool`

HasReadyContainers returns a boolean if a field has been set.

### GetReason

`func (o *InfrastructurePod) GetReason() string`

GetReason returns the Reason field if non-nil, zero value otherwise.

### GetReasonOk

`func (o *InfrastructurePod) GetReasonOk() (*string, bool)`

GetReasonOk returns a tuple with the Reason field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetReason

`func (o *InfrastructurePod) SetReason(v string)`

SetReason sets Reason field to given value.

### HasReason

`func (o *InfrastructurePod) HasReason() bool`

HasReason returns a boolean if a field has been set.

### GetRef

`func (o *InfrastructurePod) GetRef() InfrastructureObjectReference`

GetRef returns the Ref field if non-nil, zero value otherwise.

### GetRefOk

`func (o *InfrastructurePod) GetRefOk() (*InfrastructureObjectReference, bool)`

GetRefOk returns a tuple with the Ref field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRef

`func (o *InfrastructurePod) SetRef(v InfrastructureObjectReference)`

SetRef sets Ref field to given value.

### HasRef

`func (o *InfrastructurePod) HasRef() bool`

HasRef returns a boolean if a field has been set.

### GetReferences

`func (o *InfrastructurePod) GetReferences() []InfrastructureObjectReference`

GetReferences returns the References field if non-nil, zero value otherwise.

### GetReferencesOk

`func (o *InfrastructurePod) GetReferencesOk() (*[]InfrastructureObjectReference, bool)`

GetReferencesOk returns a tuple with the References field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetReferences

`func (o *InfrastructurePod) SetReferences(v []InfrastructureObjectReference)`

SetReferences sets References field to given value.

### HasReferences

`func (o *InfrastructurePod) HasReferences() bool`

HasReferences returns a boolean if a field has been set.

### GetReferencesTruncation

`func (o *InfrastructurePod) GetReferencesTruncation() CheckpointTruncation`

GetReferencesTruncation returns the ReferencesTruncation field if non-nil, zero value otherwise.

### GetReferencesTruncationOk

`func (o *InfrastructurePod) GetReferencesTruncationOk() (*CheckpointTruncation, bool)`

GetReferencesTruncationOk returns a tuple with the ReferencesTruncation field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetReferencesTruncation

`func (o *InfrastructurePod) SetReferencesTruncation(v CheckpointTruncation)`

SetReferencesTruncation sets ReferencesTruncation field to given value.

### HasReferencesTruncation

`func (o *InfrastructurePod) HasReferencesTruncation() bool`

HasReferencesTruncation returns a boolean if a field has been set.

### GetRestarts

`func (o *InfrastructurePod) GetRestarts() int64`

GetRestarts returns the Restarts field if non-nil, zero value otherwise.

### GetRestartsOk

`func (o *InfrastructurePod) GetRestartsOk() (*int64, bool)`

GetRestartsOk returns a tuple with the Restarts field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRestarts

`func (o *InfrastructurePod) SetRestarts(v int64)`

SetRestarts sets Restarts field to given value.

### HasRestarts

`func (o *InfrastructurePod) HasRestarts() bool`

HasRestarts returns a boolean if a field has been set.

### GetTolerations

`func (o *InfrastructurePod) GetTolerations() []InfrastructureSchedulingRule`

GetTolerations returns the Tolerations field if non-nil, zero value otherwise.

### GetTolerationsOk

`func (o *InfrastructurePod) GetTolerationsOk() (*[]InfrastructureSchedulingRule, bool)`

GetTolerationsOk returns a tuple with the Tolerations field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTolerations

`func (o *InfrastructurePod) SetTolerations(v []InfrastructureSchedulingRule)`

SetTolerations sets Tolerations field to given value.

### HasTolerations

`func (o *InfrastructurePod) HasTolerations() bool`

HasTolerations returns a boolean if a field has been set.

### GetTolerationsTruncation

`func (o *InfrastructurePod) GetTolerationsTruncation() CheckpointTruncation`

GetTolerationsTruncation returns the TolerationsTruncation field if non-nil, zero value otherwise.

### GetTolerationsTruncationOk

`func (o *InfrastructurePod) GetTolerationsTruncationOk() (*CheckpointTruncation, bool)`

GetTolerationsTruncationOk returns a tuple with the TolerationsTruncation field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTolerationsTruncation

`func (o *InfrastructurePod) SetTolerationsTruncation(v CheckpointTruncation)`

SetTolerationsTruncation sets TolerationsTruncation field to given value.

### HasTolerationsTruncation

`func (o *InfrastructurePod) HasTolerationsTruncation() bool`

HasTolerationsTruncation returns a boolean if a field has been set.

### GetTotalContainers

`func (o *InfrastructurePod) GetTotalContainers() int64`

GetTotalContainers returns the TotalContainers field if non-nil, zero value otherwise.

### GetTotalContainersOk

`func (o *InfrastructurePod) GetTotalContainersOk() (*int64, bool)`

GetTotalContainersOk returns a tuple with the TotalContainers field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTotalContainers

`func (o *InfrastructurePod) SetTotalContainers(v int64)`

SetTotalContainers sets TotalContainers field to given value.

### HasTotalContainers

`func (o *InfrastructurePod) HasTotalContainers() bool`

HasTotalContainers returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


