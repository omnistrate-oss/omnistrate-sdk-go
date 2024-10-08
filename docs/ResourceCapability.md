# ResourceCapability

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Capability** | **string** | The type of capability of a resource | 
**Configuration** | **map[string]interface{}** | The configuration parameters of a capability of a resource | 

## Methods

### NewResourceCapability

`func NewResourceCapability(capability string, configuration map[string]interface{}, ) *ResourceCapability`

NewResourceCapability instantiates a new ResourceCapability object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewResourceCapabilityWithDefaults

`func NewResourceCapabilityWithDefaults() *ResourceCapability`

NewResourceCapabilityWithDefaults instantiates a new ResourceCapability object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetCapability

`func (o *ResourceCapability) GetCapability() string`

GetCapability returns the Capability field if non-nil, zero value otherwise.

### GetCapabilityOk

`func (o *ResourceCapability) GetCapabilityOk() (*string, bool)`

GetCapabilityOk returns a tuple with the Capability field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCapability

`func (o *ResourceCapability) SetCapability(v string)`

SetCapability sets Capability field to given value.


### GetConfiguration

`func (o *ResourceCapability) GetConfiguration() map[string]interface{}`

GetConfiguration returns the Configuration field if non-nil, zero value otherwise.

### GetConfigurationOk

`func (o *ResourceCapability) GetConfigurationOk() (*map[string]interface{}, bool)`

GetConfigurationOk returns a tuple with the Configuration field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetConfiguration

`func (o *ResourceCapability) SetConfiguration(v map[string]interface{})`

SetConfiguration sets Configuration field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


