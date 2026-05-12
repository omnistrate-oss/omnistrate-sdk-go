# DescribeCustomWorkflowResult

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**ApiParameters** | Pointer to [**[]InputParameterEntity**](InputParameterEntity.md) | The API parameters that users provide when invoking this workflow | [optional] 
**Configuration** | Pointer to **map[string]interface{}** | Workflow-specific configuration as a flexible key-value map. For backup workflows: backupRetentionInDays (int), backupPeriodInHours (int), snapshotBeforeDeletion (bool). | [optional] 
**Description** | Pointer to **string** | A brief description of the workflow | [optional] 
**Id** | **string** | ID of a Custom Workflow | 
**IsSystemWorkflow** | **bool** | Whether this workflow was defined from a reserved systemWorkflows entry. | 
**Name** | **string** | The display name of the workflow | 
**ProductTierId** | **string** | ID of a Product Tier | 
**Scope** | Pointer to **[]string** | The user context scopes that can invoke this workflow | [optional] 
**ServiceId** | **string** | ID of a Service | 
**Spec** | **interface{}** | The workflow definition stored with the service plan. This is persisted as JSON and mirrors the workflow spec model in commons. | 
**Verb** | **string** | The provider-defined verb associated with a custom workflow. Any custom action value is accepted; reserved system workflow verbs are returned with isSystemWorkflow&#x3D;true. | 

## Methods

### NewDescribeCustomWorkflowResult

`func NewDescribeCustomWorkflowResult(id string, isSystemWorkflow bool, name string, productTierId string, serviceId string, spec interface{}, verb string, ) *DescribeCustomWorkflowResult`

NewDescribeCustomWorkflowResult instantiates a new DescribeCustomWorkflowResult object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewDescribeCustomWorkflowResultWithDefaults

`func NewDescribeCustomWorkflowResultWithDefaults() *DescribeCustomWorkflowResult`

NewDescribeCustomWorkflowResultWithDefaults instantiates a new DescribeCustomWorkflowResult object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetApiParameters

`func (o *DescribeCustomWorkflowResult) GetApiParameters() []InputParameterEntity`

GetApiParameters returns the ApiParameters field if non-nil, zero value otherwise.

### GetApiParametersOk

`func (o *DescribeCustomWorkflowResult) GetApiParametersOk() (*[]InputParameterEntity, bool)`

GetApiParametersOk returns a tuple with the ApiParameters field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetApiParameters

`func (o *DescribeCustomWorkflowResult) SetApiParameters(v []InputParameterEntity)`

SetApiParameters sets ApiParameters field to given value.

### HasApiParameters

`func (o *DescribeCustomWorkflowResult) HasApiParameters() bool`

HasApiParameters returns a boolean if a field has been set.

### GetConfiguration

`func (o *DescribeCustomWorkflowResult) GetConfiguration() map[string]interface{}`

GetConfiguration returns the Configuration field if non-nil, zero value otherwise.

### GetConfigurationOk

`func (o *DescribeCustomWorkflowResult) GetConfigurationOk() (*map[string]interface{}, bool)`

GetConfigurationOk returns a tuple with the Configuration field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetConfiguration

`func (o *DescribeCustomWorkflowResult) SetConfiguration(v map[string]interface{})`

SetConfiguration sets Configuration field to given value.

### HasConfiguration

`func (o *DescribeCustomWorkflowResult) HasConfiguration() bool`

HasConfiguration returns a boolean if a field has been set.

### GetDescription

`func (o *DescribeCustomWorkflowResult) GetDescription() string`

GetDescription returns the Description field if non-nil, zero value otherwise.

### GetDescriptionOk

`func (o *DescribeCustomWorkflowResult) GetDescriptionOk() (*string, bool)`

GetDescriptionOk returns a tuple with the Description field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDescription

`func (o *DescribeCustomWorkflowResult) SetDescription(v string)`

SetDescription sets Description field to given value.

### HasDescription

`func (o *DescribeCustomWorkflowResult) HasDescription() bool`

HasDescription returns a boolean if a field has been set.

### GetId

`func (o *DescribeCustomWorkflowResult) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *DescribeCustomWorkflowResult) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *DescribeCustomWorkflowResult) SetId(v string)`

SetId sets Id field to given value.


### GetIsSystemWorkflow

`func (o *DescribeCustomWorkflowResult) GetIsSystemWorkflow() bool`

GetIsSystemWorkflow returns the IsSystemWorkflow field if non-nil, zero value otherwise.

### GetIsSystemWorkflowOk

`func (o *DescribeCustomWorkflowResult) GetIsSystemWorkflowOk() (*bool, bool)`

GetIsSystemWorkflowOk returns a tuple with the IsSystemWorkflow field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIsSystemWorkflow

`func (o *DescribeCustomWorkflowResult) SetIsSystemWorkflow(v bool)`

SetIsSystemWorkflow sets IsSystemWorkflow field to given value.


### GetName

`func (o *DescribeCustomWorkflowResult) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *DescribeCustomWorkflowResult) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *DescribeCustomWorkflowResult) SetName(v string)`

SetName sets Name field to given value.


### GetProductTierId

`func (o *DescribeCustomWorkflowResult) GetProductTierId() string`

GetProductTierId returns the ProductTierId field if non-nil, zero value otherwise.

### GetProductTierIdOk

`func (o *DescribeCustomWorkflowResult) GetProductTierIdOk() (*string, bool)`

GetProductTierIdOk returns a tuple with the ProductTierId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetProductTierId

`func (o *DescribeCustomWorkflowResult) SetProductTierId(v string)`

SetProductTierId sets ProductTierId field to given value.


### GetScope

`func (o *DescribeCustomWorkflowResult) GetScope() []string`

GetScope returns the Scope field if non-nil, zero value otherwise.

### GetScopeOk

`func (o *DescribeCustomWorkflowResult) GetScopeOk() (*[]string, bool)`

GetScopeOk returns a tuple with the Scope field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetScope

`func (o *DescribeCustomWorkflowResult) SetScope(v []string)`

SetScope sets Scope field to given value.

### HasScope

`func (o *DescribeCustomWorkflowResult) HasScope() bool`

HasScope returns a boolean if a field has been set.

### GetServiceId

`func (o *DescribeCustomWorkflowResult) GetServiceId() string`

GetServiceId returns the ServiceId field if non-nil, zero value otherwise.

### GetServiceIdOk

`func (o *DescribeCustomWorkflowResult) GetServiceIdOk() (*string, bool)`

GetServiceIdOk returns a tuple with the ServiceId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetServiceId

`func (o *DescribeCustomWorkflowResult) SetServiceId(v string)`

SetServiceId sets ServiceId field to given value.


### GetSpec

`func (o *DescribeCustomWorkflowResult) GetSpec() interface{}`

GetSpec returns the Spec field if non-nil, zero value otherwise.

### GetSpecOk

`func (o *DescribeCustomWorkflowResult) GetSpecOk() (*interface{}, bool)`

GetSpecOk returns a tuple with the Spec field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSpec

`func (o *DescribeCustomWorkflowResult) SetSpec(v interface{})`

SetSpec sets Spec field to given value.


### SetSpecNil

`func (o *DescribeCustomWorkflowResult) SetSpecNil(b bool)`

 SetSpecNil sets the value for Spec to be an explicit nil

### UnsetSpec
`func (o *DescribeCustomWorkflowResult) UnsetSpec()`

UnsetSpec ensures that no value is present for Spec, not even an explicit nil
### GetVerb

`func (o *DescribeCustomWorkflowResult) GetVerb() string`

GetVerb returns the Verb field if non-nil, zero value otherwise.

### GetVerbOk

`func (o *DescribeCustomWorkflowResult) GetVerbOk() (*string, bool)`

GetVerbOk returns a tuple with the Verb field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVerb

`func (o *DescribeCustomWorkflowResult) SetVerb(v string)`

SetVerb sets Verb field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


