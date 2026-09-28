# InfrastructureExcludedObject

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**ApiVersion** | Pointer to **string** |  | [optional] 
**AppliedBy** | Pointer to [**InfrastructureAppliedBy**](InfrastructureAppliedBy.md) |  | [optional] 
**Kind** | Pointer to **string** |  | [optional] 
**Name** | Pointer to **string** |  | [optional] 
**Namespace** | Pointer to **string** |  | [optional] 
**Reason** | Pointer to **string** | outside-namespace or cluster-scoped | [optional] 

## Methods

### NewInfrastructureExcludedObject

`func NewInfrastructureExcludedObject() *InfrastructureExcludedObject`

NewInfrastructureExcludedObject instantiates a new InfrastructureExcludedObject object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewInfrastructureExcludedObjectWithDefaults

`func NewInfrastructureExcludedObjectWithDefaults() *InfrastructureExcludedObject`

NewInfrastructureExcludedObjectWithDefaults instantiates a new InfrastructureExcludedObject object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetApiVersion

`func (o *InfrastructureExcludedObject) GetApiVersion() string`

GetApiVersion returns the ApiVersion field if non-nil, zero value otherwise.

### GetApiVersionOk

`func (o *InfrastructureExcludedObject) GetApiVersionOk() (*string, bool)`

GetApiVersionOk returns a tuple with the ApiVersion field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetApiVersion

`func (o *InfrastructureExcludedObject) SetApiVersion(v string)`

SetApiVersion sets ApiVersion field to given value.

### HasApiVersion

`func (o *InfrastructureExcludedObject) HasApiVersion() bool`

HasApiVersion returns a boolean if a field has been set.

### GetAppliedBy

`func (o *InfrastructureExcludedObject) GetAppliedBy() InfrastructureAppliedBy`

GetAppliedBy returns the AppliedBy field if non-nil, zero value otherwise.

### GetAppliedByOk

`func (o *InfrastructureExcludedObject) GetAppliedByOk() (*InfrastructureAppliedBy, bool)`

GetAppliedByOk returns a tuple with the AppliedBy field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAppliedBy

`func (o *InfrastructureExcludedObject) SetAppliedBy(v InfrastructureAppliedBy)`

SetAppliedBy sets AppliedBy field to given value.

### HasAppliedBy

`func (o *InfrastructureExcludedObject) HasAppliedBy() bool`

HasAppliedBy returns a boolean if a field has been set.

### GetKind

`func (o *InfrastructureExcludedObject) GetKind() string`

GetKind returns the Kind field if non-nil, zero value otherwise.

### GetKindOk

`func (o *InfrastructureExcludedObject) GetKindOk() (*string, bool)`

GetKindOk returns a tuple with the Kind field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetKind

`func (o *InfrastructureExcludedObject) SetKind(v string)`

SetKind sets Kind field to given value.

### HasKind

`func (o *InfrastructureExcludedObject) HasKind() bool`

HasKind returns a boolean if a field has been set.

### GetName

`func (o *InfrastructureExcludedObject) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *InfrastructureExcludedObject) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *InfrastructureExcludedObject) SetName(v string)`

SetName sets Name field to given value.

### HasName

`func (o *InfrastructureExcludedObject) HasName() bool`

HasName returns a boolean if a field has been set.

### GetNamespace

`func (o *InfrastructureExcludedObject) GetNamespace() string`

GetNamespace returns the Namespace field if non-nil, zero value otherwise.

### GetNamespaceOk

`func (o *InfrastructureExcludedObject) GetNamespaceOk() (*string, bool)`

GetNamespaceOk returns a tuple with the Namespace field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNamespace

`func (o *InfrastructureExcludedObject) SetNamespace(v string)`

SetNamespace sets Namespace field to given value.

### HasNamespace

`func (o *InfrastructureExcludedObject) HasNamespace() bool`

HasNamespace returns a boolean if a field has been set.

### GetReason

`func (o *InfrastructureExcludedObject) GetReason() string`

GetReason returns the Reason field if non-nil, zero value otherwise.

### GetReasonOk

`func (o *InfrastructureExcludedObject) GetReasonOk() (*string, bool)`

GetReasonOk returns a tuple with the Reason field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetReason

`func (o *InfrastructureExcludedObject) SetReason(v string)`

SetReason sets Reason field to given value.

### HasReason

`func (o *InfrastructureExcludedObject) HasReason() bool`

HasReason returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


