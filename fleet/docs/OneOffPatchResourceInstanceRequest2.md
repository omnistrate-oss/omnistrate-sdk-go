# OneOffPatchResourceInstanceRequest2

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**ResourceId** | **string** | The resource ID. | 
**ResourceOverrideConfiguration** | Pointer to [**map[string]ResourceOneOffPatchConfigurationOverride**](ResourceOneOffPatchConfigurationOverride.md) | The resource override configuration for one-off patching. | [optional] 
**TargetTierVersion** | Pointer to **string** | The target resource version. | [optional] 

## Methods

### NewOneOffPatchResourceInstanceRequest2

`func NewOneOffPatchResourceInstanceRequest2(resourceId string, ) *OneOffPatchResourceInstanceRequest2`

NewOneOffPatchResourceInstanceRequest2 instantiates a new OneOffPatchResourceInstanceRequest2 object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewOneOffPatchResourceInstanceRequest2WithDefaults

`func NewOneOffPatchResourceInstanceRequest2WithDefaults() *OneOffPatchResourceInstanceRequest2`

NewOneOffPatchResourceInstanceRequest2WithDefaults instantiates a new OneOffPatchResourceInstanceRequest2 object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetResourceId

`func (o *OneOffPatchResourceInstanceRequest2) GetResourceId() string`

GetResourceId returns the ResourceId field if non-nil, zero value otherwise.

### GetResourceIdOk

`func (o *OneOffPatchResourceInstanceRequest2) GetResourceIdOk() (*string, bool)`

GetResourceIdOk returns a tuple with the ResourceId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetResourceId

`func (o *OneOffPatchResourceInstanceRequest2) SetResourceId(v string)`

SetResourceId sets ResourceId field to given value.


### GetResourceOverrideConfiguration

`func (o *OneOffPatchResourceInstanceRequest2) GetResourceOverrideConfiguration() map[string]ResourceOneOffPatchConfigurationOverride`

GetResourceOverrideConfiguration returns the ResourceOverrideConfiguration field if non-nil, zero value otherwise.

### GetResourceOverrideConfigurationOk

`func (o *OneOffPatchResourceInstanceRequest2) GetResourceOverrideConfigurationOk() (*map[string]ResourceOneOffPatchConfigurationOverride, bool)`

GetResourceOverrideConfigurationOk returns a tuple with the ResourceOverrideConfiguration field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetResourceOverrideConfiguration

`func (o *OneOffPatchResourceInstanceRequest2) SetResourceOverrideConfiguration(v map[string]ResourceOneOffPatchConfigurationOverride)`

SetResourceOverrideConfiguration sets ResourceOverrideConfiguration field to given value.

### HasResourceOverrideConfiguration

`func (o *OneOffPatchResourceInstanceRequest2) HasResourceOverrideConfiguration() bool`

HasResourceOverrideConfiguration returns a boolean if a field has been set.

### GetTargetTierVersion

`func (o *OneOffPatchResourceInstanceRequest2) GetTargetTierVersion() string`

GetTargetTierVersion returns the TargetTierVersion field if non-nil, zero value otherwise.

### GetTargetTierVersionOk

`func (o *OneOffPatchResourceInstanceRequest2) GetTargetTierVersionOk() (*string, bool)`

GetTargetTierVersionOk returns a tuple with the TargetTierVersion field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTargetTierVersion

`func (o *OneOffPatchResourceInstanceRequest2) SetTargetTierVersion(v string)`

SetTargetTierVersion sets TargetTierVersion field to given value.

### HasTargetTierVersion

`func (o *OneOffPatchResourceInstanceRequest2) HasTargetTierVersion() bool`

HasTargetTierVersion returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


