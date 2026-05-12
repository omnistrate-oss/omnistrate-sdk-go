# UpdateCustomWorkflowRequest2

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**ApiParameters** | Pointer to [**[]InputParameterEntity**](InputParameterEntity.md) | The API parameters that users provide when invoking this workflow | [optional] 
**Configuration** | Pointer to **map[string]interface{}** | Workflow-specific configuration as a flexible key-value map. For backup workflows: backupRetentionInDays (int), backupPeriodInHours (int), snapshotBeforeDeletion (bool). New properties can be added without API or schema changes. | [optional] 
**Description** | Pointer to **string** | A brief description of the workflow | [optional] 
**Name** | Pointer to **string** | The display name of the workflow | [optional] 
**Scope** | Pointer to **[]string** | The user context scopes that can invoke this workflow | [optional] 
**Spec** | Pointer to **interface{}** | The workflow definition persisted for the service plan | [optional] 
**Verb** | Pointer to **string** | The workflow verb | [optional] 

## Methods

### NewUpdateCustomWorkflowRequest2

`func NewUpdateCustomWorkflowRequest2() *UpdateCustomWorkflowRequest2`

NewUpdateCustomWorkflowRequest2 instantiates a new UpdateCustomWorkflowRequest2 object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewUpdateCustomWorkflowRequest2WithDefaults

`func NewUpdateCustomWorkflowRequest2WithDefaults() *UpdateCustomWorkflowRequest2`

NewUpdateCustomWorkflowRequest2WithDefaults instantiates a new UpdateCustomWorkflowRequest2 object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetApiParameters

`func (o *UpdateCustomWorkflowRequest2) GetApiParameters() []InputParameterEntity`

GetApiParameters returns the ApiParameters field if non-nil, zero value otherwise.

### GetApiParametersOk

`func (o *UpdateCustomWorkflowRequest2) GetApiParametersOk() (*[]InputParameterEntity, bool)`

GetApiParametersOk returns a tuple with the ApiParameters field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetApiParameters

`func (o *UpdateCustomWorkflowRequest2) SetApiParameters(v []InputParameterEntity)`

SetApiParameters sets ApiParameters field to given value.

### HasApiParameters

`func (o *UpdateCustomWorkflowRequest2) HasApiParameters() bool`

HasApiParameters returns a boolean if a field has been set.

### GetConfiguration

`func (o *UpdateCustomWorkflowRequest2) GetConfiguration() map[string]interface{}`

GetConfiguration returns the Configuration field if non-nil, zero value otherwise.

### GetConfigurationOk

`func (o *UpdateCustomWorkflowRequest2) GetConfigurationOk() (*map[string]interface{}, bool)`

GetConfigurationOk returns a tuple with the Configuration field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetConfiguration

`func (o *UpdateCustomWorkflowRequest2) SetConfiguration(v map[string]interface{})`

SetConfiguration sets Configuration field to given value.

### HasConfiguration

`func (o *UpdateCustomWorkflowRequest2) HasConfiguration() bool`

HasConfiguration returns a boolean if a field has been set.

### GetDescription

`func (o *UpdateCustomWorkflowRequest2) GetDescription() string`

GetDescription returns the Description field if non-nil, zero value otherwise.

### GetDescriptionOk

`func (o *UpdateCustomWorkflowRequest2) GetDescriptionOk() (*string, bool)`

GetDescriptionOk returns a tuple with the Description field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDescription

`func (o *UpdateCustomWorkflowRequest2) SetDescription(v string)`

SetDescription sets Description field to given value.

### HasDescription

`func (o *UpdateCustomWorkflowRequest2) HasDescription() bool`

HasDescription returns a boolean if a field has been set.

### GetName

`func (o *UpdateCustomWorkflowRequest2) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *UpdateCustomWorkflowRequest2) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *UpdateCustomWorkflowRequest2) SetName(v string)`

SetName sets Name field to given value.

### HasName

`func (o *UpdateCustomWorkflowRequest2) HasName() bool`

HasName returns a boolean if a field has been set.

### GetScope

`func (o *UpdateCustomWorkflowRequest2) GetScope() []string`

GetScope returns the Scope field if non-nil, zero value otherwise.

### GetScopeOk

`func (o *UpdateCustomWorkflowRequest2) GetScopeOk() (*[]string, bool)`

GetScopeOk returns a tuple with the Scope field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetScope

`func (o *UpdateCustomWorkflowRequest2) SetScope(v []string)`

SetScope sets Scope field to given value.

### HasScope

`func (o *UpdateCustomWorkflowRequest2) HasScope() bool`

HasScope returns a boolean if a field has been set.

### GetSpec

`func (o *UpdateCustomWorkflowRequest2) GetSpec() interface{}`

GetSpec returns the Spec field if non-nil, zero value otherwise.

### GetSpecOk

`func (o *UpdateCustomWorkflowRequest2) GetSpecOk() (*interface{}, bool)`

GetSpecOk returns a tuple with the Spec field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSpec

`func (o *UpdateCustomWorkflowRequest2) SetSpec(v interface{})`

SetSpec sets Spec field to given value.

### HasSpec

`func (o *UpdateCustomWorkflowRequest2) HasSpec() bool`

HasSpec returns a boolean if a field has been set.

### SetSpecNil

`func (o *UpdateCustomWorkflowRequest2) SetSpecNil(b bool)`

 SetSpecNil sets the value for Spec to be an explicit nil

### UnsetSpec
`func (o *UpdateCustomWorkflowRequest2) UnsetSpec()`

UnsetSpec ensures that no value is present for Spec, not even an explicit nil
### GetVerb

`func (o *UpdateCustomWorkflowRequest2) GetVerb() string`

GetVerb returns the Verb field if non-nil, zero value otherwise.

### GetVerbOk

`func (o *UpdateCustomWorkflowRequest2) GetVerbOk() (*string, bool)`

GetVerbOk returns a tuple with the Verb field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVerb

`func (o *UpdateCustomWorkflowRequest2) SetVerb(v string)`

SetVerb sets Verb field to given value.

### HasVerb

`func (o *UpdateCustomWorkflowRequest2) HasVerb() bool`

HasVerb returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


