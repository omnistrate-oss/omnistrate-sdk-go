# UpdateCustomWorkflowRequest

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**ApiParameters** | Pointer to [**[]InputParameterEntity**](InputParameterEntity.md) | The API parameters that users provide when invoking this workflow | [optional] 
**Configuration** | Pointer to **map[string]interface{}** | Workflow-specific configuration as a flexible key-value map. For backup workflows: backupRetentionInDays (int), backupPeriodInHours (int), snapshotBeforeDeletion (bool). New properties can be added without API or schema changes. | [optional] 
**Description** | Pointer to **string** | A brief description of the workflow | [optional] 
**Id** | **string** | ID of a Custom Workflow | 
**Name** | Pointer to **string** | The display name of the workflow | [optional] 
**OutputParameterSpecs** | Pointer to [**[]WorkflowOutputParameterSpec**](WorkflowOutputParameterSpec.md) | Typed declarations of the output parameters this workflow produces. | [optional] 
**Scope** | Pointer to **[]string** | The user context scopes that can invoke this workflow | [optional] 
**ServiceId** | **string** | ID of a Service | 
**Spec** | Pointer to **interface{}** | The workflow definition stored with the service plan. This is persisted as JSON and mirrors the workflow spec model in commons. | [optional] 
**Token** | **string** | JWT token used to perform authorization | 
**Verb** | Pointer to **string** | The provider-defined verb associated with a custom workflow. Any custom action value is accepted; reserved system workflow verbs are returned with isSystemWorkflow&#x3D;true. | [optional] 

## Methods

### NewUpdateCustomWorkflowRequest

`func NewUpdateCustomWorkflowRequest(id string, serviceId string, token string, ) *UpdateCustomWorkflowRequest`

NewUpdateCustomWorkflowRequest instantiates a new UpdateCustomWorkflowRequest object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewUpdateCustomWorkflowRequestWithDefaults

`func NewUpdateCustomWorkflowRequestWithDefaults() *UpdateCustomWorkflowRequest`

NewUpdateCustomWorkflowRequestWithDefaults instantiates a new UpdateCustomWorkflowRequest object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetApiParameters

`func (o *UpdateCustomWorkflowRequest) GetApiParameters() []InputParameterEntity`

GetApiParameters returns the ApiParameters field if non-nil, zero value otherwise.

### GetApiParametersOk

`func (o *UpdateCustomWorkflowRequest) GetApiParametersOk() (*[]InputParameterEntity, bool)`

GetApiParametersOk returns a tuple with the ApiParameters field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetApiParameters

`func (o *UpdateCustomWorkflowRequest) SetApiParameters(v []InputParameterEntity)`

SetApiParameters sets ApiParameters field to given value.

### HasApiParameters

`func (o *UpdateCustomWorkflowRequest) HasApiParameters() bool`

HasApiParameters returns a boolean if a field has been set.

### GetConfiguration

`func (o *UpdateCustomWorkflowRequest) GetConfiguration() map[string]interface{}`

GetConfiguration returns the Configuration field if non-nil, zero value otherwise.

### GetConfigurationOk

`func (o *UpdateCustomWorkflowRequest) GetConfigurationOk() (*map[string]interface{}, bool)`

GetConfigurationOk returns a tuple with the Configuration field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetConfiguration

`func (o *UpdateCustomWorkflowRequest) SetConfiguration(v map[string]interface{})`

SetConfiguration sets Configuration field to given value.

### HasConfiguration

`func (o *UpdateCustomWorkflowRequest) HasConfiguration() bool`

HasConfiguration returns a boolean if a field has been set.

### GetDescription

`func (o *UpdateCustomWorkflowRequest) GetDescription() string`

GetDescription returns the Description field if non-nil, zero value otherwise.

### GetDescriptionOk

`func (o *UpdateCustomWorkflowRequest) GetDescriptionOk() (*string, bool)`

GetDescriptionOk returns a tuple with the Description field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDescription

`func (o *UpdateCustomWorkflowRequest) SetDescription(v string)`

SetDescription sets Description field to given value.

### HasDescription

`func (o *UpdateCustomWorkflowRequest) HasDescription() bool`

HasDescription returns a boolean if a field has been set.

### GetId

`func (o *UpdateCustomWorkflowRequest) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *UpdateCustomWorkflowRequest) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *UpdateCustomWorkflowRequest) SetId(v string)`

SetId sets Id field to given value.


### GetName

`func (o *UpdateCustomWorkflowRequest) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *UpdateCustomWorkflowRequest) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *UpdateCustomWorkflowRequest) SetName(v string)`

SetName sets Name field to given value.

### HasName

`func (o *UpdateCustomWorkflowRequest) HasName() bool`

HasName returns a boolean if a field has been set.

### GetOutputParameterSpecs

`func (o *UpdateCustomWorkflowRequest) GetOutputParameterSpecs() []WorkflowOutputParameterSpec`

GetOutputParameterSpecs returns the OutputParameterSpecs field if non-nil, zero value otherwise.

### GetOutputParameterSpecsOk

`func (o *UpdateCustomWorkflowRequest) GetOutputParameterSpecsOk() (*[]WorkflowOutputParameterSpec, bool)`

GetOutputParameterSpecsOk returns a tuple with the OutputParameterSpecs field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOutputParameterSpecs

`func (o *UpdateCustomWorkflowRequest) SetOutputParameterSpecs(v []WorkflowOutputParameterSpec)`

SetOutputParameterSpecs sets OutputParameterSpecs field to given value.

### HasOutputParameterSpecs

`func (o *UpdateCustomWorkflowRequest) HasOutputParameterSpecs() bool`

HasOutputParameterSpecs returns a boolean if a field has been set.

### GetScope

`func (o *UpdateCustomWorkflowRequest) GetScope() []string`

GetScope returns the Scope field if non-nil, zero value otherwise.

### GetScopeOk

`func (o *UpdateCustomWorkflowRequest) GetScopeOk() (*[]string, bool)`

GetScopeOk returns a tuple with the Scope field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetScope

`func (o *UpdateCustomWorkflowRequest) SetScope(v []string)`

SetScope sets Scope field to given value.

### HasScope

`func (o *UpdateCustomWorkflowRequest) HasScope() bool`

HasScope returns a boolean if a field has been set.

### GetServiceId

`func (o *UpdateCustomWorkflowRequest) GetServiceId() string`

GetServiceId returns the ServiceId field if non-nil, zero value otherwise.

### GetServiceIdOk

`func (o *UpdateCustomWorkflowRequest) GetServiceIdOk() (*string, bool)`

GetServiceIdOk returns a tuple with the ServiceId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetServiceId

`func (o *UpdateCustomWorkflowRequest) SetServiceId(v string)`

SetServiceId sets ServiceId field to given value.


### GetSpec

`func (o *UpdateCustomWorkflowRequest) GetSpec() interface{}`

GetSpec returns the Spec field if non-nil, zero value otherwise.

### GetSpecOk

`func (o *UpdateCustomWorkflowRequest) GetSpecOk() (*interface{}, bool)`

GetSpecOk returns a tuple with the Spec field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSpec

`func (o *UpdateCustomWorkflowRequest) SetSpec(v interface{})`

SetSpec sets Spec field to given value.

### HasSpec

`func (o *UpdateCustomWorkflowRequest) HasSpec() bool`

HasSpec returns a boolean if a field has been set.

### SetSpecNil

`func (o *UpdateCustomWorkflowRequest) SetSpecNil(b bool)`

 SetSpecNil sets the value for Spec to be an explicit nil

### UnsetSpec
`func (o *UpdateCustomWorkflowRequest) UnsetSpec()`

UnsetSpec ensures that no value is present for Spec, not even an explicit nil
### GetToken

`func (o *UpdateCustomWorkflowRequest) GetToken() string`

GetToken returns the Token field if non-nil, zero value otherwise.

### GetTokenOk

`func (o *UpdateCustomWorkflowRequest) GetTokenOk() (*string, bool)`

GetTokenOk returns a tuple with the Token field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetToken

`func (o *UpdateCustomWorkflowRequest) SetToken(v string)`

SetToken sets Token field to given value.


### GetVerb

`func (o *UpdateCustomWorkflowRequest) GetVerb() string`

GetVerb returns the Verb field if non-nil, zero value otherwise.

### GetVerbOk

`func (o *UpdateCustomWorkflowRequest) GetVerbOk() (*string, bool)`

GetVerbOk returns a tuple with the Verb field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVerb

`func (o *UpdateCustomWorkflowRequest) SetVerb(v string)`

SetVerb sets Verb field to given value.

### HasVerb

`func (o *UpdateCustomWorkflowRequest) HasVerb() bool`

HasVerb returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


