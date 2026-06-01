# ResourceInstanceSupportedOperation

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**ApiParameters** | Pointer to [**[]InputParameterEntity**](InputParameterEntity.md) | The API parameters users provide when invoking this operation. | [optional] 
**Configuration** | Pointer to **map[string]interface{}** | Operation-specific configuration that does not require user input. For backup operations this can include backupRetentionInDays, backupPeriodInHours, snapshotBeforeDeletion, and other backup-related settings. | [optional] 
**Description** | Pointer to **string** | A brief description of the operation. | [optional] 
**Id** | Pointer to **string** | ID of a Custom Workflow | [optional] 
**Name** | **string** | The display name of the operation. | 
**Scope** | Pointer to **[]string** | The user context scopes that can invoke this operation. | [optional] 
**Source** | **string** | The implementation source backing a supported operation. | 
**Verb** | **string** | The verb used to invoke this operation. | 

## Methods

### NewResourceInstanceSupportedOperation

`func NewResourceInstanceSupportedOperation(name string, source string, verb string, ) *ResourceInstanceSupportedOperation`

NewResourceInstanceSupportedOperation instantiates a new ResourceInstanceSupportedOperation object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewResourceInstanceSupportedOperationWithDefaults

`func NewResourceInstanceSupportedOperationWithDefaults() *ResourceInstanceSupportedOperation`

NewResourceInstanceSupportedOperationWithDefaults instantiates a new ResourceInstanceSupportedOperation object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetApiParameters

`func (o *ResourceInstanceSupportedOperation) GetApiParameters() []InputParameterEntity`

GetApiParameters returns the ApiParameters field if non-nil, zero value otherwise.

### GetApiParametersOk

`func (o *ResourceInstanceSupportedOperation) GetApiParametersOk() (*[]InputParameterEntity, bool)`

GetApiParametersOk returns a tuple with the ApiParameters field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetApiParameters

`func (o *ResourceInstanceSupportedOperation) SetApiParameters(v []InputParameterEntity)`

SetApiParameters sets ApiParameters field to given value.

### HasApiParameters

`func (o *ResourceInstanceSupportedOperation) HasApiParameters() bool`

HasApiParameters returns a boolean if a field has been set.

### GetConfiguration

`func (o *ResourceInstanceSupportedOperation) GetConfiguration() map[string]interface{}`

GetConfiguration returns the Configuration field if non-nil, zero value otherwise.

### GetConfigurationOk

`func (o *ResourceInstanceSupportedOperation) GetConfigurationOk() (*map[string]interface{}, bool)`

GetConfigurationOk returns a tuple with the Configuration field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetConfiguration

`func (o *ResourceInstanceSupportedOperation) SetConfiguration(v map[string]interface{})`

SetConfiguration sets Configuration field to given value.

### HasConfiguration

`func (o *ResourceInstanceSupportedOperation) HasConfiguration() bool`

HasConfiguration returns a boolean if a field has been set.

### GetDescription

`func (o *ResourceInstanceSupportedOperation) GetDescription() string`

GetDescription returns the Description field if non-nil, zero value otherwise.

### GetDescriptionOk

`func (o *ResourceInstanceSupportedOperation) GetDescriptionOk() (*string, bool)`

GetDescriptionOk returns a tuple with the Description field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDescription

`func (o *ResourceInstanceSupportedOperation) SetDescription(v string)`

SetDescription sets Description field to given value.

### HasDescription

`func (o *ResourceInstanceSupportedOperation) HasDescription() bool`

HasDescription returns a boolean if a field has been set.

### GetId

`func (o *ResourceInstanceSupportedOperation) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *ResourceInstanceSupportedOperation) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *ResourceInstanceSupportedOperation) SetId(v string)`

SetId sets Id field to given value.

### HasId

`func (o *ResourceInstanceSupportedOperation) HasId() bool`

HasId returns a boolean if a field has been set.

### GetName

`func (o *ResourceInstanceSupportedOperation) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *ResourceInstanceSupportedOperation) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *ResourceInstanceSupportedOperation) SetName(v string)`

SetName sets Name field to given value.


### GetScope

`func (o *ResourceInstanceSupportedOperation) GetScope() []string`

GetScope returns the Scope field if non-nil, zero value otherwise.

### GetScopeOk

`func (o *ResourceInstanceSupportedOperation) GetScopeOk() (*[]string, bool)`

GetScopeOk returns a tuple with the Scope field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetScope

`func (o *ResourceInstanceSupportedOperation) SetScope(v []string)`

SetScope sets Scope field to given value.

### HasScope

`func (o *ResourceInstanceSupportedOperation) HasScope() bool`

HasScope returns a boolean if a field has been set.

### GetSource

`func (o *ResourceInstanceSupportedOperation) GetSource() string`

GetSource returns the Source field if non-nil, zero value otherwise.

### GetSourceOk

`func (o *ResourceInstanceSupportedOperation) GetSourceOk() (*string, bool)`

GetSourceOk returns a tuple with the Source field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSource

`func (o *ResourceInstanceSupportedOperation) SetSource(v string)`

SetSource sets Source field to given value.


### GetVerb

`func (o *ResourceInstanceSupportedOperation) GetVerb() string`

GetVerb returns the Verb field if non-nil, zero value otherwise.

### GetVerbOk

`func (o *ResourceInstanceSupportedOperation) GetVerbOk() (*string, bool)`

GetVerbOk returns a tuple with the Verb field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVerb

`func (o *ResourceInstanceSupportedOperation) SetVerb(v string)`

SetVerb sets Verb field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


