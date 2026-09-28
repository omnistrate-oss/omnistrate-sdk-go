# InfrastructureAppliedObject

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**ApiVersion** | Pointer to **string** |  | [optional] 
**AppliedBy** | Pointer to [**InfrastructureAppliedBy**](InfrastructureAppliedBy.md) |  | [optional] 
**Availability** | Pointer to [**InfrastructureAvailability**](InfrastructureAvailability.md) |  | [optional] 
**Conditions** | Pointer to [**[]InfrastructureCondition**](InfrastructureCondition.md) | At most three, messages capped | [optional] 
**ConditionsTruncation** | Pointer to [**CheckpointTruncation**](CheckpointTruncation.md) |  | [optional] 
**CustomResource** | Pointer to **bool** | The kind is served by a CustomResourceDefinition | [optional] 
**Generation** | Pointer to **int64** |  | [optional] 
**ObservedGeneration** | Pointer to **int64** |  | [optional] 
**RecordedUid** | Pointer to **string** | UID when the step applied it | [optional] 
**Ref** | Pointer to [**InfrastructureObjectReference**](InfrastructureObjectReference.md) |  | [optional] 

## Methods

### NewInfrastructureAppliedObject

`func NewInfrastructureAppliedObject() *InfrastructureAppliedObject`

NewInfrastructureAppliedObject instantiates a new InfrastructureAppliedObject object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewInfrastructureAppliedObjectWithDefaults

`func NewInfrastructureAppliedObjectWithDefaults() *InfrastructureAppliedObject`

NewInfrastructureAppliedObjectWithDefaults instantiates a new InfrastructureAppliedObject object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetApiVersion

`func (o *InfrastructureAppliedObject) GetApiVersion() string`

GetApiVersion returns the ApiVersion field if non-nil, zero value otherwise.

### GetApiVersionOk

`func (o *InfrastructureAppliedObject) GetApiVersionOk() (*string, bool)`

GetApiVersionOk returns a tuple with the ApiVersion field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetApiVersion

`func (o *InfrastructureAppliedObject) SetApiVersion(v string)`

SetApiVersion sets ApiVersion field to given value.

### HasApiVersion

`func (o *InfrastructureAppliedObject) HasApiVersion() bool`

HasApiVersion returns a boolean if a field has been set.

### GetAppliedBy

`func (o *InfrastructureAppliedObject) GetAppliedBy() InfrastructureAppliedBy`

GetAppliedBy returns the AppliedBy field if non-nil, zero value otherwise.

### GetAppliedByOk

`func (o *InfrastructureAppliedObject) GetAppliedByOk() (*InfrastructureAppliedBy, bool)`

GetAppliedByOk returns a tuple with the AppliedBy field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAppliedBy

`func (o *InfrastructureAppliedObject) SetAppliedBy(v InfrastructureAppliedBy)`

SetAppliedBy sets AppliedBy field to given value.

### HasAppliedBy

`func (o *InfrastructureAppliedObject) HasAppliedBy() bool`

HasAppliedBy returns a boolean if a field has been set.

### GetAvailability

`func (o *InfrastructureAppliedObject) GetAvailability() InfrastructureAvailability`

GetAvailability returns the Availability field if non-nil, zero value otherwise.

### GetAvailabilityOk

`func (o *InfrastructureAppliedObject) GetAvailabilityOk() (*InfrastructureAvailability, bool)`

GetAvailabilityOk returns a tuple with the Availability field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAvailability

`func (o *InfrastructureAppliedObject) SetAvailability(v InfrastructureAvailability)`

SetAvailability sets Availability field to given value.

### HasAvailability

`func (o *InfrastructureAppliedObject) HasAvailability() bool`

HasAvailability returns a boolean if a field has been set.

### GetConditions

`func (o *InfrastructureAppliedObject) GetConditions() []InfrastructureCondition`

GetConditions returns the Conditions field if non-nil, zero value otherwise.

### GetConditionsOk

`func (o *InfrastructureAppliedObject) GetConditionsOk() (*[]InfrastructureCondition, bool)`

GetConditionsOk returns a tuple with the Conditions field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetConditions

`func (o *InfrastructureAppliedObject) SetConditions(v []InfrastructureCondition)`

SetConditions sets Conditions field to given value.

### HasConditions

`func (o *InfrastructureAppliedObject) HasConditions() bool`

HasConditions returns a boolean if a field has been set.

### GetConditionsTruncation

`func (o *InfrastructureAppliedObject) GetConditionsTruncation() CheckpointTruncation`

GetConditionsTruncation returns the ConditionsTruncation field if non-nil, zero value otherwise.

### GetConditionsTruncationOk

`func (o *InfrastructureAppliedObject) GetConditionsTruncationOk() (*CheckpointTruncation, bool)`

GetConditionsTruncationOk returns a tuple with the ConditionsTruncation field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetConditionsTruncation

`func (o *InfrastructureAppliedObject) SetConditionsTruncation(v CheckpointTruncation)`

SetConditionsTruncation sets ConditionsTruncation field to given value.

### HasConditionsTruncation

`func (o *InfrastructureAppliedObject) HasConditionsTruncation() bool`

HasConditionsTruncation returns a boolean if a field has been set.

### GetCustomResource

`func (o *InfrastructureAppliedObject) GetCustomResource() bool`

GetCustomResource returns the CustomResource field if non-nil, zero value otherwise.

### GetCustomResourceOk

`func (o *InfrastructureAppliedObject) GetCustomResourceOk() (*bool, bool)`

GetCustomResourceOk returns a tuple with the CustomResource field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCustomResource

`func (o *InfrastructureAppliedObject) SetCustomResource(v bool)`

SetCustomResource sets CustomResource field to given value.

### HasCustomResource

`func (o *InfrastructureAppliedObject) HasCustomResource() bool`

HasCustomResource returns a boolean if a field has been set.

### GetGeneration

`func (o *InfrastructureAppliedObject) GetGeneration() int64`

GetGeneration returns the Generation field if non-nil, zero value otherwise.

### GetGenerationOk

`func (o *InfrastructureAppliedObject) GetGenerationOk() (*int64, bool)`

GetGenerationOk returns a tuple with the Generation field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetGeneration

`func (o *InfrastructureAppliedObject) SetGeneration(v int64)`

SetGeneration sets Generation field to given value.

### HasGeneration

`func (o *InfrastructureAppliedObject) HasGeneration() bool`

HasGeneration returns a boolean if a field has been set.

### GetObservedGeneration

`func (o *InfrastructureAppliedObject) GetObservedGeneration() int64`

GetObservedGeneration returns the ObservedGeneration field if non-nil, zero value otherwise.

### GetObservedGenerationOk

`func (o *InfrastructureAppliedObject) GetObservedGenerationOk() (*int64, bool)`

GetObservedGenerationOk returns a tuple with the ObservedGeneration field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetObservedGeneration

`func (o *InfrastructureAppliedObject) SetObservedGeneration(v int64)`

SetObservedGeneration sets ObservedGeneration field to given value.

### HasObservedGeneration

`func (o *InfrastructureAppliedObject) HasObservedGeneration() bool`

HasObservedGeneration returns a boolean if a field has been set.

### GetRecordedUid

`func (o *InfrastructureAppliedObject) GetRecordedUid() string`

GetRecordedUid returns the RecordedUid field if non-nil, zero value otherwise.

### GetRecordedUidOk

`func (o *InfrastructureAppliedObject) GetRecordedUidOk() (*string, bool)`

GetRecordedUidOk returns a tuple with the RecordedUid field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRecordedUid

`func (o *InfrastructureAppliedObject) SetRecordedUid(v string)`

SetRecordedUid sets RecordedUid field to given value.

### HasRecordedUid

`func (o *InfrastructureAppliedObject) HasRecordedUid() bool`

HasRecordedUid returns a boolean if a field has been set.

### GetRef

`func (o *InfrastructureAppliedObject) GetRef() InfrastructureObjectReference`

GetRef returns the Ref field if non-nil, zero value otherwise.

### GetRefOk

`func (o *InfrastructureAppliedObject) GetRefOk() (*InfrastructureObjectReference, bool)`

GetRefOk returns a tuple with the Ref field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRef

`func (o *InfrastructureAppliedObject) SetRef(v InfrastructureObjectReference)`

SetRef sets Ref field to given value.

### HasRef

`func (o *InfrastructureAppliedObject) HasRef() bool`

HasRef returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


