# DisableResourceCapabilityRequest

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Capability** | **string** | The type of capability of a resource | 
**Id** | **string** | ID of a resource | 
**ServiceId** | **string** | ID of a Service | 
**Token** | **string** | JWT token used to perform authorization | 

## Methods

### NewDisableResourceCapabilityRequest

`func NewDisableResourceCapabilityRequest(capability string, id string, serviceId string, token string, ) *DisableResourceCapabilityRequest`

NewDisableResourceCapabilityRequest instantiates a new DisableResourceCapabilityRequest object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewDisableResourceCapabilityRequestWithDefaults

`func NewDisableResourceCapabilityRequestWithDefaults() *DisableResourceCapabilityRequest`

NewDisableResourceCapabilityRequestWithDefaults instantiates a new DisableResourceCapabilityRequest object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetCapability

`func (o *DisableResourceCapabilityRequest) GetCapability() string`

GetCapability returns the Capability field if non-nil, zero value otherwise.

### GetCapabilityOk

`func (o *DisableResourceCapabilityRequest) GetCapabilityOk() (*string, bool)`

GetCapabilityOk returns a tuple with the Capability field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCapability

`func (o *DisableResourceCapabilityRequest) SetCapability(v string)`

SetCapability sets Capability field to given value.


### GetId

`func (o *DisableResourceCapabilityRequest) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *DisableResourceCapabilityRequest) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *DisableResourceCapabilityRequest) SetId(v string)`

SetId sets Id field to given value.


### GetServiceId

`func (o *DisableResourceCapabilityRequest) GetServiceId() string`

GetServiceId returns the ServiceId field if non-nil, zero value otherwise.

### GetServiceIdOk

`func (o *DisableResourceCapabilityRequest) GetServiceIdOk() (*string, bool)`

GetServiceIdOk returns a tuple with the ServiceId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetServiceId

`func (o *DisableResourceCapabilityRequest) SetServiceId(v string)`

SetServiceId sets ServiceId field to given value.


### GetToken

`func (o *DisableResourceCapabilityRequest) GetToken() string`

GetToken returns the Token field if non-nil, zero value otherwise.

### GetTokenOk

`func (o *DisableResourceCapabilityRequest) GetTokenOk() (*string, bool)`

GetTokenOk returns a tuple with the Token field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetToken

`func (o *DisableResourceCapabilityRequest) SetToken(v string)`

SetToken sets Token field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


