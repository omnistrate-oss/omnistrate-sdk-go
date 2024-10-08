# EnableResourceCapabilityRequestBody

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Capability** | **string** | The type of capability of a resource | 
**Configuration** | Pointer to **map[string]interface{}** | The configuration parameters of a capability of a resource | [optional] 

## Methods

### NewEnableResourceCapabilityRequestBody

`func NewEnableResourceCapabilityRequestBody(capability string, ) *EnableResourceCapabilityRequestBody`

NewEnableResourceCapabilityRequestBody instantiates a new EnableResourceCapabilityRequestBody object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewEnableResourceCapabilityRequestBodyWithDefaults

`func NewEnableResourceCapabilityRequestBodyWithDefaults() *EnableResourceCapabilityRequestBody`

NewEnableResourceCapabilityRequestBodyWithDefaults instantiates a new EnableResourceCapabilityRequestBody object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetCapability

`func (o *EnableResourceCapabilityRequestBody) GetCapability() string`

GetCapability returns the Capability field if non-nil, zero value otherwise.

### GetCapabilityOk

`func (o *EnableResourceCapabilityRequestBody) GetCapabilityOk() (*string, bool)`

GetCapabilityOk returns a tuple with the Capability field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCapability

`func (o *EnableResourceCapabilityRequestBody) SetCapability(v string)`

SetCapability sets Capability field to given value.


### GetConfiguration

`func (o *EnableResourceCapabilityRequestBody) GetConfiguration() map[string]interface{}`

GetConfiguration returns the Configuration field if non-nil, zero value otherwise.

### GetConfigurationOk

`func (o *EnableResourceCapabilityRequestBody) GetConfigurationOk() (*map[string]interface{}, bool)`

GetConfigurationOk returns a tuple with the Configuration field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetConfiguration

`func (o *EnableResourceCapabilityRequestBody) SetConfiguration(v map[string]interface{})`

SetConfiguration sets Configuration field to given value.

### HasConfiguration

`func (o *EnableResourceCapabilityRequestBody) HasConfiguration() bool`

HasConfiguration returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


