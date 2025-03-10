# FleetUpdateAccountConfigResourceInstanceRequest

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Disconnect** | Pointer to **bool** | Disconnect account config instance or not | [optional] 
**InstanceId** | **string** | ID of a Resource Instance | 
**ServiceId** | **string** | ID of a Service | 
**Token** | **string** | JWT token used to perform authorization | 

## Methods

### NewFleetUpdateAccountConfigResourceInstanceRequest

`func NewFleetUpdateAccountConfigResourceInstanceRequest(instanceId string, serviceId string, token string, ) *FleetUpdateAccountConfigResourceInstanceRequest`

NewFleetUpdateAccountConfigResourceInstanceRequest instantiates a new FleetUpdateAccountConfigResourceInstanceRequest object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewFleetUpdateAccountConfigResourceInstanceRequestWithDefaults

`func NewFleetUpdateAccountConfigResourceInstanceRequestWithDefaults() *FleetUpdateAccountConfigResourceInstanceRequest`

NewFleetUpdateAccountConfigResourceInstanceRequestWithDefaults instantiates a new FleetUpdateAccountConfigResourceInstanceRequest object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetDisconnect

`func (o *FleetUpdateAccountConfigResourceInstanceRequest) GetDisconnect() bool`

GetDisconnect returns the Disconnect field if non-nil, zero value otherwise.

### GetDisconnectOk

`func (o *FleetUpdateAccountConfigResourceInstanceRequest) GetDisconnectOk() (*bool, bool)`

GetDisconnectOk returns a tuple with the Disconnect field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDisconnect

`func (o *FleetUpdateAccountConfigResourceInstanceRequest) SetDisconnect(v bool)`

SetDisconnect sets Disconnect field to given value.

### HasDisconnect

`func (o *FleetUpdateAccountConfigResourceInstanceRequest) HasDisconnect() bool`

HasDisconnect returns a boolean if a field has been set.

### GetInstanceId

`func (o *FleetUpdateAccountConfigResourceInstanceRequest) GetInstanceId() string`

GetInstanceId returns the InstanceId field if non-nil, zero value otherwise.

### GetInstanceIdOk

`func (o *FleetUpdateAccountConfigResourceInstanceRequest) GetInstanceIdOk() (*string, bool)`

GetInstanceIdOk returns a tuple with the InstanceId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetInstanceId

`func (o *FleetUpdateAccountConfigResourceInstanceRequest) SetInstanceId(v string)`

SetInstanceId sets InstanceId field to given value.


### GetServiceId

`func (o *FleetUpdateAccountConfigResourceInstanceRequest) GetServiceId() string`

GetServiceId returns the ServiceId field if non-nil, zero value otherwise.

### GetServiceIdOk

`func (o *FleetUpdateAccountConfigResourceInstanceRequest) GetServiceIdOk() (*string, bool)`

GetServiceIdOk returns a tuple with the ServiceId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetServiceId

`func (o *FleetUpdateAccountConfigResourceInstanceRequest) SetServiceId(v string)`

SetServiceId sets ServiceId field to given value.


### GetToken

`func (o *FleetUpdateAccountConfigResourceInstanceRequest) GetToken() string`

GetToken returns the Token field if non-nil, zero value otherwise.

### GetTokenOk

`func (o *FleetUpdateAccountConfigResourceInstanceRequest) GetTokenOk() (*string, bool)`

GetTokenOk returns a tuple with the Token field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetToken

`func (o *FleetUpdateAccountConfigResourceInstanceRequest) SetToken(v string)`

SetToken sets Token field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


