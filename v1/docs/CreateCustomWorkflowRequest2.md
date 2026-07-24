# CreateCustomWorkflowRequest2

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**ApiParameters** | Pointer to [**[]InputParameterEntity**](InputParameterEntity.md) | The API parameters that users provide when invoking this workflow | [optional] 
**Configuration** | Pointer to **map[string]interface{}** | Workflow-specific configuration as a flexible key-value map. For backup workflows: backupRetentionInDays (int), backupPeriodInHours (int), snapshotBeforeDeletion (bool). New properties can be added without API or schema changes. | [optional] 
**Description** | Pointer to **string** | A brief description of the workflow | [optional] 
**Name** | **string** | The display name of the workflow | 
**OutputParameterSpecs** | Pointer to [**[]WorkflowOutputParameterSpec**](WorkflowOutputParameterSpec.md) | Typed declarations of the output parameters this workflow produces. | [optional] 
**Scope** | Pointer to **[]string** | The user context scopes that can invoke this workflow | [optional] 
**Spec** | **interface{}** | The workflow definition persisted for the service plan | 
**Verb** | **string** | The workflow verb | 

## Methods

### NewCreateCustomWorkflowRequest2

`func NewCreateCustomWorkflowRequest2(name string, spec interface{}, verb string, ) *CreateCustomWorkflowRequest2`

NewCreateCustomWorkflowRequest2 instantiates a new CreateCustomWorkflowRequest2 object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewCreateCustomWorkflowRequest2WithDefaults

`func NewCreateCustomWorkflowRequest2WithDefaults() *CreateCustomWorkflowRequest2`

NewCreateCustomWorkflowRequest2WithDefaults instantiates a new CreateCustomWorkflowRequest2 object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetApiParameters

`func (o *CreateCustomWorkflowRequest2) GetApiParameters() []InputParameterEntity`

GetApiParameters returns the ApiParameters field if non-nil, zero value otherwise.

### GetApiParametersOk

`func (o *CreateCustomWorkflowRequest2) GetApiParametersOk() (*[]InputParameterEntity, bool)`

GetApiParametersOk returns a tuple with the ApiParameters field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetApiParameters

`func (o *CreateCustomWorkflowRequest2) SetApiParameters(v []InputParameterEntity)`

SetApiParameters sets ApiParameters field to given value.

### HasApiParameters

`func (o *CreateCustomWorkflowRequest2) HasApiParameters() bool`

HasApiParameters returns a boolean if a field has been set.

### GetConfiguration

`func (o *CreateCustomWorkflowRequest2) GetConfiguration() map[string]interface{}`

GetConfiguration returns the Configuration field if non-nil, zero value otherwise.

### GetConfigurationOk

`func (o *CreateCustomWorkflowRequest2) GetConfigurationOk() (*map[string]interface{}, bool)`

GetConfigurationOk returns a tuple with the Configuration field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetConfiguration

`func (o *CreateCustomWorkflowRequest2) SetConfiguration(v map[string]interface{})`

SetConfiguration sets Configuration field to given value.

### HasConfiguration

`func (o *CreateCustomWorkflowRequest2) HasConfiguration() bool`

HasConfiguration returns a boolean if a field has been set.

### GetDescription

`func (o *CreateCustomWorkflowRequest2) GetDescription() string`

GetDescription returns the Description field if non-nil, zero value otherwise.

### GetDescriptionOk

`func (o *CreateCustomWorkflowRequest2) GetDescriptionOk() (*string, bool)`

GetDescriptionOk returns a tuple with the Description field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDescription

`func (o *CreateCustomWorkflowRequest2) SetDescription(v string)`

SetDescription sets Description field to given value.

### HasDescription

`func (o *CreateCustomWorkflowRequest2) HasDescription() bool`

HasDescription returns a boolean if a field has been set.

### GetName

`func (o *CreateCustomWorkflowRequest2) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *CreateCustomWorkflowRequest2) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *CreateCustomWorkflowRequest2) SetName(v string)`

SetName sets Name field to given value.


### GetOutputParameterSpecs

`func (o *CreateCustomWorkflowRequest2) GetOutputParameterSpecs() []WorkflowOutputParameterSpec`

GetOutputParameterSpecs returns the OutputParameterSpecs field if non-nil, zero value otherwise.

### GetOutputParameterSpecsOk

`func (o *CreateCustomWorkflowRequest2) GetOutputParameterSpecsOk() (*[]WorkflowOutputParameterSpec, bool)`

GetOutputParameterSpecsOk returns a tuple with the OutputParameterSpecs field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOutputParameterSpecs

`func (o *CreateCustomWorkflowRequest2) SetOutputParameterSpecs(v []WorkflowOutputParameterSpec)`

SetOutputParameterSpecs sets OutputParameterSpecs field to given value.

### HasOutputParameterSpecs

`func (o *CreateCustomWorkflowRequest2) HasOutputParameterSpecs() bool`

HasOutputParameterSpecs returns a boolean if a field has been set.

### GetScope

`func (o *CreateCustomWorkflowRequest2) GetScope() []string`

GetScope returns the Scope field if non-nil, zero value otherwise.

### GetScopeOk

`func (o *CreateCustomWorkflowRequest2) GetScopeOk() (*[]string, bool)`

GetScopeOk returns a tuple with the Scope field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetScope

`func (o *CreateCustomWorkflowRequest2) SetScope(v []string)`

SetScope sets Scope field to given value.

### HasScope

`func (o *CreateCustomWorkflowRequest2) HasScope() bool`

HasScope returns a boolean if a field has been set.

### GetSpec

`func (o *CreateCustomWorkflowRequest2) GetSpec() interface{}`

GetSpec returns the Spec field if non-nil, zero value otherwise.

### GetSpecOk

`func (o *CreateCustomWorkflowRequest2) GetSpecOk() (*interface{}, bool)`

GetSpecOk returns a tuple with the Spec field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSpec

`func (o *CreateCustomWorkflowRequest2) SetSpec(v interface{})`

SetSpec sets Spec field to given value.


### SetSpecNil

`func (o *CreateCustomWorkflowRequest2) SetSpecNil(b bool)`

 SetSpecNil sets the value for Spec to be an explicit nil

### UnsetSpec
`func (o *CreateCustomWorkflowRequest2) UnsetSpec()`

UnsetSpec ensures that no value is present for Spec, not even an explicit nil
### GetVerb

`func (o *CreateCustomWorkflowRequest2) GetVerb() string`

GetVerb returns the Verb field if non-nil, zero value otherwise.

### GetVerbOk

`func (o *CreateCustomWorkflowRequest2) GetVerbOk() (*string, bool)`

GetVerbOk returns a tuple with the Verb field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVerb

`func (o *CreateCustomWorkflowRequest2) SetVerb(v string)`

SetVerb sets Verb field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


