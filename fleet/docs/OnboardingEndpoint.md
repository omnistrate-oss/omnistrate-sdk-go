# OnboardingEndpoint

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Host** | Pointer to **string** | The host of the endpoint. | [optional] 
**NetworkingType** | Pointer to **string** | The networking type of the endpoint (PUBLIC or PRIVATE). | [optional] 
**Ports** | Pointer to **[]int64** | The ports of the endpoint. | [optional] 
**Primary** | Pointer to **bool** | Whether this is the primary endpoint. | [optional] 

## Methods

### NewOnboardingEndpoint

`func NewOnboardingEndpoint() *OnboardingEndpoint`

NewOnboardingEndpoint instantiates a new OnboardingEndpoint object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewOnboardingEndpointWithDefaults

`func NewOnboardingEndpointWithDefaults() *OnboardingEndpoint`

NewOnboardingEndpointWithDefaults instantiates a new OnboardingEndpoint object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetHost

`func (o *OnboardingEndpoint) GetHost() string`

GetHost returns the Host field if non-nil, zero value otherwise.

### GetHostOk

`func (o *OnboardingEndpoint) GetHostOk() (*string, bool)`

GetHostOk returns a tuple with the Host field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetHost

`func (o *OnboardingEndpoint) SetHost(v string)`

SetHost sets Host field to given value.

### HasHost

`func (o *OnboardingEndpoint) HasHost() bool`

HasHost returns a boolean if a field has been set.

### GetNetworkingType

`func (o *OnboardingEndpoint) GetNetworkingType() string`

GetNetworkingType returns the NetworkingType field if non-nil, zero value otherwise.

### GetNetworkingTypeOk

`func (o *OnboardingEndpoint) GetNetworkingTypeOk() (*string, bool)`

GetNetworkingTypeOk returns a tuple with the NetworkingType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNetworkingType

`func (o *OnboardingEndpoint) SetNetworkingType(v string)`

SetNetworkingType sets NetworkingType field to given value.

### HasNetworkingType

`func (o *OnboardingEndpoint) HasNetworkingType() bool`

HasNetworkingType returns a boolean if a field has been set.

### GetPorts

`func (o *OnboardingEndpoint) GetPorts() []int64`

GetPorts returns the Ports field if non-nil, zero value otherwise.

### GetPortsOk

`func (o *OnboardingEndpoint) GetPortsOk() (*[]int64, bool)`

GetPortsOk returns a tuple with the Ports field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPorts

`func (o *OnboardingEndpoint) SetPorts(v []int64)`

SetPorts sets Ports field to given value.

### HasPorts

`func (o *OnboardingEndpoint) HasPorts() bool`

HasPorts returns a boolean if a field has been set.

### GetPrimary

`func (o *OnboardingEndpoint) GetPrimary() bool`

GetPrimary returns the Primary field if non-nil, zero value otherwise.

### GetPrimaryOk

`func (o *OnboardingEndpoint) GetPrimaryOk() (*bool, bool)`

GetPrimaryOk returns a tuple with the Primary field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPrimary

`func (o *OnboardingEndpoint) SetPrimary(v bool)`

SetPrimary sets Primary field to given value.

### HasPrimary

`func (o *OnboardingEndpoint) HasPrimary() bool`

HasPrimary returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


