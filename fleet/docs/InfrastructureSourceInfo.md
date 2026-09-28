# InfrastructureSourceInfo

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**DeclaredObjects** | Pointer to **int64** | Kustomize: object documents the stored manifest declares | [optional] 
**Excluded** | Pointer to [**[]InfrastructureExcludedObject**](InfrastructureExcludedObject.md) | Objects in scope that were not read | [optional] 
**ExcludedTruncation** | Pointer to [**CheckpointTruncation**](CheckpointTruncation.md) |  | [optional] 
**LatestExecution** | Pointer to [**InfrastructureExecutionSummary**](InfrastructureExecutionSummary.md) |  | [optional] 
**Name** | Pointer to **string** | The resource&#39;s identity within its type: release, resource key, alias or kustomization name | [optional] 
**ScopeBasis** | Pointer to **string** | manifest, recorded-objects or workflow-applied | [optional] 
**Type** | Pointer to **string** | helm, compose, operator or kustomize | [optional] 

## Methods

### NewInfrastructureSourceInfo

`func NewInfrastructureSourceInfo() *InfrastructureSourceInfo`

NewInfrastructureSourceInfo instantiates a new InfrastructureSourceInfo object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewInfrastructureSourceInfoWithDefaults

`func NewInfrastructureSourceInfoWithDefaults() *InfrastructureSourceInfo`

NewInfrastructureSourceInfoWithDefaults instantiates a new InfrastructureSourceInfo object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetDeclaredObjects

`func (o *InfrastructureSourceInfo) GetDeclaredObjects() int64`

GetDeclaredObjects returns the DeclaredObjects field if non-nil, zero value otherwise.

### GetDeclaredObjectsOk

`func (o *InfrastructureSourceInfo) GetDeclaredObjectsOk() (*int64, bool)`

GetDeclaredObjectsOk returns a tuple with the DeclaredObjects field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDeclaredObjects

`func (o *InfrastructureSourceInfo) SetDeclaredObjects(v int64)`

SetDeclaredObjects sets DeclaredObjects field to given value.

### HasDeclaredObjects

`func (o *InfrastructureSourceInfo) HasDeclaredObjects() bool`

HasDeclaredObjects returns a boolean if a field has been set.

### GetExcluded

`func (o *InfrastructureSourceInfo) GetExcluded() []InfrastructureExcludedObject`

GetExcluded returns the Excluded field if non-nil, zero value otherwise.

### GetExcludedOk

`func (o *InfrastructureSourceInfo) GetExcludedOk() (*[]InfrastructureExcludedObject, bool)`

GetExcludedOk returns a tuple with the Excluded field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetExcluded

`func (o *InfrastructureSourceInfo) SetExcluded(v []InfrastructureExcludedObject)`

SetExcluded sets Excluded field to given value.

### HasExcluded

`func (o *InfrastructureSourceInfo) HasExcluded() bool`

HasExcluded returns a boolean if a field has been set.

### GetExcludedTruncation

`func (o *InfrastructureSourceInfo) GetExcludedTruncation() CheckpointTruncation`

GetExcludedTruncation returns the ExcludedTruncation field if non-nil, zero value otherwise.

### GetExcludedTruncationOk

`func (o *InfrastructureSourceInfo) GetExcludedTruncationOk() (*CheckpointTruncation, bool)`

GetExcludedTruncationOk returns a tuple with the ExcludedTruncation field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetExcludedTruncation

`func (o *InfrastructureSourceInfo) SetExcludedTruncation(v CheckpointTruncation)`

SetExcludedTruncation sets ExcludedTruncation field to given value.

### HasExcludedTruncation

`func (o *InfrastructureSourceInfo) HasExcludedTruncation() bool`

HasExcludedTruncation returns a boolean if a field has been set.

### GetLatestExecution

`func (o *InfrastructureSourceInfo) GetLatestExecution() InfrastructureExecutionSummary`

GetLatestExecution returns the LatestExecution field if non-nil, zero value otherwise.

### GetLatestExecutionOk

`func (o *InfrastructureSourceInfo) GetLatestExecutionOk() (*InfrastructureExecutionSummary, bool)`

GetLatestExecutionOk returns a tuple with the LatestExecution field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLatestExecution

`func (o *InfrastructureSourceInfo) SetLatestExecution(v InfrastructureExecutionSummary)`

SetLatestExecution sets LatestExecution field to given value.

### HasLatestExecution

`func (o *InfrastructureSourceInfo) HasLatestExecution() bool`

HasLatestExecution returns a boolean if a field has been set.

### GetName

`func (o *InfrastructureSourceInfo) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *InfrastructureSourceInfo) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *InfrastructureSourceInfo) SetName(v string)`

SetName sets Name field to given value.

### HasName

`func (o *InfrastructureSourceInfo) HasName() bool`

HasName returns a boolean if a field has been set.

### GetScopeBasis

`func (o *InfrastructureSourceInfo) GetScopeBasis() string`

GetScopeBasis returns the ScopeBasis field if non-nil, zero value otherwise.

### GetScopeBasisOk

`func (o *InfrastructureSourceInfo) GetScopeBasisOk() (*string, bool)`

GetScopeBasisOk returns a tuple with the ScopeBasis field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetScopeBasis

`func (o *InfrastructureSourceInfo) SetScopeBasis(v string)`

SetScopeBasis sets ScopeBasis field to given value.

### HasScopeBasis

`func (o *InfrastructureSourceInfo) HasScopeBasis() bool`

HasScopeBasis returns a boolean if a field has been set.

### GetType

`func (o *InfrastructureSourceInfo) GetType() string`

GetType returns the Type field if non-nil, zero value otherwise.

### GetTypeOk

`func (o *InfrastructureSourceInfo) GetTypeOk() (*string, bool)`

GetTypeOk returns a tuple with the Type field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetType

`func (o *InfrastructureSourceInfo) SetType(v string)`

SetType sets Type field to given value.

### HasType

`func (o *InfrastructureSourceInfo) HasType() bool`

HasType returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


