# CreateCustomWorkflowRequest

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**ApiParameters** | Pointer to [**[]InputParameterEntity**](InputParameterEntity.md) | The API parameters that users provide when invoking this workflow | [optional] 
**Configuration** | Pointer to **map[string]interface{}** | Workflow-specific configuration as a flexible key-value map. For backup workflows: backupRetentionInDays (int), backupPeriodInHours (int), snapshotBeforeDeletion (bool). New properties can be added without API or schema changes. | [optional] 
**Description** | Pointer to **string** | A brief description of the workflow | [optional] 
**Name** | **string** | The display name of the workflow | 
**ProductTierId** | **string** | ID of a Product Tier | 
**Scope** | Pointer to **[]string** | The user context scopes that can invoke this workflow | [optional] 
**ServiceId** | **string** | ID of a Service | 
**Spec** | **interface{}** | The workflow definition stored with the service plan. This is persisted as JSON and mirrors the workflow spec model in commons. | 
**Token** | **string** | JWT token used to perform authorization | 
**Verb** | **string** | The provider-defined verb associated with a custom workflow. Any custom action value is accepted; reserved system workflow verbs are returned with isSystemWorkflow&#x3D;true. | 

## Methods

### NewCreateCustomWorkflowRequest

`func NewCreateCustomWorkflowRequest(name string, productTierId string, serviceId string, spec interface{}, token string, verb string, ) *CreateCustomWorkflowRequest`

NewCreateCustomWorkflowRequest instantiates a new CreateCustomWorkflowRequest object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewCreateCustomWorkflowRequestWithDefaults

`func NewCreateCustomWorkflowRequestWithDefaults() *CreateCustomWorkflowRequest`

NewCreateCustomWorkflowRequestWithDefaults instantiates a new CreateCustomWorkflowRequest object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetApiParameters

`func (o *CreateCustomWorkflowRequest) GetApiParameters() []InputParameterEntity`

GetApiParameters returns the ApiParameters field if non-nil, zero value otherwise.

### GetApiParametersOk

`func (o *CreateCustomWorkflowRequest) GetApiParametersOk() (*[]InputParameterEntity, bool)`

GetApiParametersOk returns a tuple with the ApiParameters field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetApiParameters

`func (o *CreateCustomWorkflowRequest) SetApiParameters(v []InputParameterEntity)`

SetApiParameters sets ApiParameters field to given value.

### HasApiParameters

`func (o *CreateCustomWorkflowRequest) HasApiParameters() bool`

HasApiParameters returns a boolean if a field has been set.

### GetConfiguration

`func (o *CreateCustomWorkflowRequest) GetConfiguration() map[string]interface{}`

GetConfiguration returns the Configuration field if non-nil, zero value otherwise.

### GetConfigurationOk

`func (o *CreateCustomWorkflowRequest) GetConfigurationOk() (*map[string]interface{}, bool)`

GetConfigurationOk returns a tuple with the Configuration field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetConfiguration

`func (o *CreateCustomWorkflowRequest) SetConfiguration(v map[string]interface{})`

SetConfiguration sets Configuration field to given value.

### HasConfiguration

`func (o *CreateCustomWorkflowRequest) HasConfiguration() bool`

HasConfiguration returns a boolean if a field has been set.

### GetDescription

`func (o *CreateCustomWorkflowRequest) GetDescription() string`

GetDescription returns the Description field if non-nil, zero value otherwise.

### GetDescriptionOk

`func (o *CreateCustomWorkflowRequest) GetDescriptionOk() (*string, bool)`

GetDescriptionOk returns a tuple with the Description field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDescription

`func (o *CreateCustomWorkflowRequest) SetDescription(v string)`

SetDescription sets Description field to given value.

### HasDescription

`func (o *CreateCustomWorkflowRequest) HasDescription() bool`

HasDescription returns a boolean if a field has been set.

### GetName

`func (o *CreateCustomWorkflowRequest) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *CreateCustomWorkflowRequest) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *CreateCustomWorkflowRequest) SetName(v string)`

SetName sets Name field to given value.


### GetProductTierId

`func (o *CreateCustomWorkflowRequest) GetProductTierId() string`

GetProductTierId returns the ProductTierId field if non-nil, zero value otherwise.

### GetProductTierIdOk

`func (o *CreateCustomWorkflowRequest) GetProductTierIdOk() (*string, bool)`

GetProductTierIdOk returns a tuple with the ProductTierId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetProductTierId

`func (o *CreateCustomWorkflowRequest) SetProductTierId(v string)`

SetProductTierId sets ProductTierId field to given value.


### GetScope

`func (o *CreateCustomWorkflowRequest) GetScope() []string`

GetScope returns the Scope field if non-nil, zero value otherwise.

### GetScopeOk

`func (o *CreateCustomWorkflowRequest) GetScopeOk() (*[]string, bool)`

GetScopeOk returns a tuple with the Scope field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetScope

`func (o *CreateCustomWorkflowRequest) SetScope(v []string)`

SetScope sets Scope field to given value.

### HasScope

`func (o *CreateCustomWorkflowRequest) HasScope() bool`

HasScope returns a boolean if a field has been set.

### GetServiceId

`func (o *CreateCustomWorkflowRequest) GetServiceId() string`

GetServiceId returns the ServiceId field if non-nil, zero value otherwise.

### GetServiceIdOk

`func (o *CreateCustomWorkflowRequest) GetServiceIdOk() (*string, bool)`

GetServiceIdOk returns a tuple with the ServiceId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetServiceId

`func (o *CreateCustomWorkflowRequest) SetServiceId(v string)`

SetServiceId sets ServiceId field to given value.


### GetSpec

`func (o *CreateCustomWorkflowRequest) GetSpec() interface{}`

GetSpec returns the Spec field if non-nil, zero value otherwise.

### GetSpecOk

`func (o *CreateCustomWorkflowRequest) GetSpecOk() (*interface{}, bool)`

GetSpecOk returns a tuple with the Spec field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSpec

`func (o *CreateCustomWorkflowRequest) SetSpec(v interface{})`

SetSpec sets Spec field to given value.


### SetSpecNil

`func (o *CreateCustomWorkflowRequest) SetSpecNil(b bool)`

 SetSpecNil sets the value for Spec to be an explicit nil

### UnsetSpec
`func (o *CreateCustomWorkflowRequest) UnsetSpec()`

UnsetSpec ensures that no value is present for Spec, not even an explicit nil
### GetToken

`func (o *CreateCustomWorkflowRequest) GetToken() string`

GetToken returns the Token field if non-nil, zero value otherwise.

### GetTokenOk

`func (o *CreateCustomWorkflowRequest) GetTokenOk() (*string, bool)`

GetTokenOk returns a tuple with the Token field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetToken

`func (o *CreateCustomWorkflowRequest) SetToken(v string)`

SetToken sets Token field to given value.


### GetVerb

`func (o *CreateCustomWorkflowRequest) GetVerb() string`

GetVerb returns the Verb field if non-nil, zero value otherwise.

### GetVerbOk

`func (o *CreateCustomWorkflowRequest) GetVerbOk() (*string, bool)`

GetVerbOk returns a tuple with the Verb field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVerb

`func (o *CreateCustomWorkflowRequest) SetVerb(v string)`

SetVerb sets Verb field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


