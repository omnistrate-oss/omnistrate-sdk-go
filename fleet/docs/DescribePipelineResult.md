# DescribePipelineResult

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Description** | Pointer to **string** | Description of the pipeline | [optional] 
**Id** | Pointer to **string** | ID of a Pipeline | [optional] 
**Name** | Pointer to **string** | Name of the pipeline | [optional] 
**ServiceEnvironmentSequence** | Pointer to **[]string** | Sequence of service environments to be deployed in the pipeline | [optional] 
**ServiceId** | Pointer to **string** | ID of a Service | [optional] 

## Methods

### NewDescribePipelineResult

`func NewDescribePipelineResult() *DescribePipelineResult`

NewDescribePipelineResult instantiates a new DescribePipelineResult object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewDescribePipelineResultWithDefaults

`func NewDescribePipelineResultWithDefaults() *DescribePipelineResult`

NewDescribePipelineResultWithDefaults instantiates a new DescribePipelineResult object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetDescription

`func (o *DescribePipelineResult) GetDescription() string`

GetDescription returns the Description field if non-nil, zero value otherwise.

### GetDescriptionOk

`func (o *DescribePipelineResult) GetDescriptionOk() (*string, bool)`

GetDescriptionOk returns a tuple with the Description field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDescription

`func (o *DescribePipelineResult) SetDescription(v string)`

SetDescription sets Description field to given value.

### HasDescription

`func (o *DescribePipelineResult) HasDescription() bool`

HasDescription returns a boolean if a field has been set.

### GetId

`func (o *DescribePipelineResult) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *DescribePipelineResult) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *DescribePipelineResult) SetId(v string)`

SetId sets Id field to given value.

### HasId

`func (o *DescribePipelineResult) HasId() bool`

HasId returns a boolean if a field has been set.

### GetName

`func (o *DescribePipelineResult) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *DescribePipelineResult) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *DescribePipelineResult) SetName(v string)`

SetName sets Name field to given value.

### HasName

`func (o *DescribePipelineResult) HasName() bool`

HasName returns a boolean if a field has been set.

### GetServiceEnvironmentSequence

`func (o *DescribePipelineResult) GetServiceEnvironmentSequence() []string`

GetServiceEnvironmentSequence returns the ServiceEnvironmentSequence field if non-nil, zero value otherwise.

### GetServiceEnvironmentSequenceOk

`func (o *DescribePipelineResult) GetServiceEnvironmentSequenceOk() (*[]string, bool)`

GetServiceEnvironmentSequenceOk returns a tuple with the ServiceEnvironmentSequence field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetServiceEnvironmentSequence

`func (o *DescribePipelineResult) SetServiceEnvironmentSequence(v []string)`

SetServiceEnvironmentSequence sets ServiceEnvironmentSequence field to given value.

### HasServiceEnvironmentSequence

`func (o *DescribePipelineResult) HasServiceEnvironmentSequence() bool`

HasServiceEnvironmentSequence returns a boolean if a field has been set.

### GetServiceId

`func (o *DescribePipelineResult) GetServiceId() string`

GetServiceId returns the ServiceId field if non-nil, zero value otherwise.

### GetServiceIdOk

`func (o *DescribePipelineResult) GetServiceIdOk() (*string, bool)`

GetServiceIdOk returns a tuple with the ServiceId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetServiceId

`func (o *DescribePipelineResult) SetServiceId(v string)`

SetServiceId sets ServiceId field to given value.

### HasServiceId

`func (o *DescribePipelineResult) HasServiceId() bool`

HasServiceId returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


