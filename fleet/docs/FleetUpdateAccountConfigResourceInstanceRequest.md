# FleetUpdateAccountConfigResourceInstanceRequest

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**InstanceId** | **string** | ID of a Resource Instance | 
**ServiceId** | **string** | ID of a Service | 
**SetConnection** | Pointer to **bool** | set account config instance connection | [optional] 
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


### GetSetConnection

`func (o *FleetUpdateAccountConfigResourceInstanceRequest) GetSetConnection() bool`

GetSetConnection returns the SetConnection field if non-nil, zero value otherwise.

### GetSetConnectionOk

`func (o *FleetUpdateAccountConfigResourceInstanceRequest) GetSetConnectionOk() (*bool, bool)`

GetSetConnectionOk returns a tuple with the SetConnection field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSetConnection

`func (o *FleetUpdateAccountConfigResourceInstanceRequest) SetSetConnection(v bool)`

SetSetConnection sets SetConnection field to given value.

### HasSetConnection

`func (o *FleetUpdateAccountConfigResourceInstanceRequest) HasSetConnection() bool`

HasSetConnection returns a boolean if a field has been set.

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


