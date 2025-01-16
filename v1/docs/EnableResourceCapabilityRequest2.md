# EnableResourceCapabilityRequest2

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Capability** | **string** | The type of capability of a resource | 
**Configuration** | Pointer to **map[string]interface{}** | The configuration parameters of a capability of a resource | [optional] 

## Methods

### NewEnableResourceCapabilityRequest2

`func NewEnableResourceCapabilityRequest2(capability string, ) *EnableResourceCapabilityRequest2`

NewEnableResourceCapabilityRequest2 instantiates a new EnableResourceCapabilityRequest2 object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewEnableResourceCapabilityRequest2WithDefaults

`func NewEnableResourceCapabilityRequest2WithDefaults() *EnableResourceCapabilityRequest2`

NewEnableResourceCapabilityRequest2WithDefaults instantiates a new EnableResourceCapabilityRequest2 object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetCapability

`func (o *EnableResourceCapabilityRequest2) GetCapability() string`

GetCapability returns the Capability field if non-nil, zero value otherwise.

### GetCapabilityOk

`func (o *EnableResourceCapabilityRequest2) GetCapabilityOk() (*string, bool)`

GetCapabilityOk returns a tuple with the Capability field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCapability

`func (o *EnableResourceCapabilityRequest2) SetCapability(v string)`

SetCapability sets Capability field to given value.


### GetConfiguration

`func (o *EnableResourceCapabilityRequest2) GetConfiguration() map[string]interface{}`

GetConfiguration returns the Configuration field if non-nil, zero value otherwise.

### GetConfigurationOk

`func (o *EnableResourceCapabilityRequest2) GetConfigurationOk() (*map[string]interface{}, bool)`

GetConfigurationOk returns a tuple with the Configuration field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetConfiguration

`func (o *EnableResourceCapabilityRequest2) SetConfiguration(v map[string]interface{})`

SetConfiguration sets Configuration field to given value.

### HasConfiguration

`func (o *EnableResourceCapabilityRequest2) HasConfiguration() bool`

HasConfiguration returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


