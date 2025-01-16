# EnableResourceCapabilityRequest

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Capability** | **string** | The type of capability of a resource | 
**Configuration** | Pointer to **map[string]interface{}** | The configuration parameters of a capability of a resource | [optional] 
**Id** | **string** | ID of a resource | 
**ServiceId** | **string** | ID of a Service | 
**Token** | **string** | JWT token used to perform authorization | 

## Methods

### NewEnableResourceCapabilityRequest

`func NewEnableResourceCapabilityRequest(capability string, id string, serviceId string, token string, ) *EnableResourceCapabilityRequest`

NewEnableResourceCapabilityRequest instantiates a new EnableResourceCapabilityRequest object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewEnableResourceCapabilityRequestWithDefaults

`func NewEnableResourceCapabilityRequestWithDefaults() *EnableResourceCapabilityRequest`

NewEnableResourceCapabilityRequestWithDefaults instantiates a new EnableResourceCapabilityRequest object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetCapability

`func (o *EnableResourceCapabilityRequest) GetCapability() string`

GetCapability returns the Capability field if non-nil, zero value otherwise.

### GetCapabilityOk

`func (o *EnableResourceCapabilityRequest) GetCapabilityOk() (*string, bool)`

GetCapabilityOk returns a tuple with the Capability field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCapability

`func (o *EnableResourceCapabilityRequest) SetCapability(v string)`

SetCapability sets Capability field to given value.


### GetConfiguration

`func (o *EnableResourceCapabilityRequest) GetConfiguration() map[string]interface{}`

GetConfiguration returns the Configuration field if non-nil, zero value otherwise.

### GetConfigurationOk

`func (o *EnableResourceCapabilityRequest) GetConfigurationOk() (*map[string]interface{}, bool)`

GetConfigurationOk returns a tuple with the Configuration field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetConfiguration

`func (o *EnableResourceCapabilityRequest) SetConfiguration(v map[string]interface{})`

SetConfiguration sets Configuration field to given value.

### HasConfiguration

`func (o *EnableResourceCapabilityRequest) HasConfiguration() bool`

HasConfiguration returns a boolean if a field has been set.

### GetId

`func (o *EnableResourceCapabilityRequest) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *EnableResourceCapabilityRequest) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *EnableResourceCapabilityRequest) SetId(v string)`

SetId sets Id field to given value.


### GetServiceId

`func (o *EnableResourceCapabilityRequest) GetServiceId() string`

GetServiceId returns the ServiceId field if non-nil, zero value otherwise.

### GetServiceIdOk

`func (o *EnableResourceCapabilityRequest) GetServiceIdOk() (*string, bool)`

GetServiceIdOk returns a tuple with the ServiceId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetServiceId

`func (o *EnableResourceCapabilityRequest) SetServiceId(v string)`

SetServiceId sets ServiceId field to given value.


### GetToken

`func (o *EnableResourceCapabilityRequest) GetToken() string`

GetToken returns the Token field if non-nil, zero value otherwise.

### GetTokenOk

`func (o *EnableResourceCapabilityRequest) GetTokenOk() (*string, bool)`

GetTokenOk returns a tuple with the Token field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetToken

`func (o *EnableResourceCapabilityRequest) SetToken(v string)`

SetToken sets Token field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


