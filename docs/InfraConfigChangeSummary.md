# InfraConfigChangeSummary

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**ChangeState** | Pointer to **string** | The pending change state for the infra configuration | [optional] 
**InfraConfigId** | Pointer to **string** | The ID of the infrastructure configuration that this resource refers to | [optional] 
**Name** | Pointer to **string** | The name of the infra config | [optional] 

## Methods

### NewInfraConfigChangeSummary

`func NewInfraConfigChangeSummary() *InfraConfigChangeSummary`

NewInfraConfigChangeSummary instantiates a new InfraConfigChangeSummary object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewInfraConfigChangeSummaryWithDefaults

`func NewInfraConfigChangeSummaryWithDefaults() *InfraConfigChangeSummary`

NewInfraConfigChangeSummaryWithDefaults instantiates a new InfraConfigChangeSummary object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetChangeState

`func (o *InfraConfigChangeSummary) GetChangeState() string`

GetChangeState returns the ChangeState field if non-nil, zero value otherwise.

### GetChangeStateOk

`func (o *InfraConfigChangeSummary) GetChangeStateOk() (*string, bool)`

GetChangeStateOk returns a tuple with the ChangeState field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetChangeState

`func (o *InfraConfigChangeSummary) SetChangeState(v string)`

SetChangeState sets ChangeState field to given value.

### HasChangeState

`func (o *InfraConfigChangeSummary) HasChangeState() bool`

HasChangeState returns a boolean if a field has been set.

### GetInfraConfigId

`func (o *InfraConfigChangeSummary) GetInfraConfigId() string`

GetInfraConfigId returns the InfraConfigId field if non-nil, zero value otherwise.

### GetInfraConfigIdOk

`func (o *InfraConfigChangeSummary) GetInfraConfigIdOk() (*string, bool)`

GetInfraConfigIdOk returns a tuple with the InfraConfigId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetInfraConfigId

`func (o *InfraConfigChangeSummary) SetInfraConfigId(v string)`

SetInfraConfigId sets InfraConfigId field to given value.

### HasInfraConfigId

`func (o *InfraConfigChangeSummary) HasInfraConfigId() bool`

HasInfraConfigId returns a boolean if a field has been set.

### GetName

`func (o *InfraConfigChangeSummary) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *InfraConfigChangeSummary) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *InfraConfigChangeSummary) SetName(v string)`

SetName sets Name field to given value.

### HasName

`func (o *InfraConfigChangeSummary) HasName() bool`

HasName returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


