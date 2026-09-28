# InfrastructurePort

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Name** | Pointer to **string** |  | [optional] 
**NodePort** | Pointer to **int64** |  | [optional] 
**Port** | Pointer to **int64** |  | [optional] 
**Protocol** | Pointer to **string** |  | [optional] 
**TargetPort** | Pointer to **string** |  | [optional] 

## Methods

### NewInfrastructurePort

`func NewInfrastructurePort() *InfrastructurePort`

NewInfrastructurePort instantiates a new InfrastructurePort object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewInfrastructurePortWithDefaults

`func NewInfrastructurePortWithDefaults() *InfrastructurePort`

NewInfrastructurePortWithDefaults instantiates a new InfrastructurePort object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetName

`func (o *InfrastructurePort) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *InfrastructurePort) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *InfrastructurePort) SetName(v string)`

SetName sets Name field to given value.

### HasName

`func (o *InfrastructurePort) HasName() bool`

HasName returns a boolean if a field has been set.

### GetNodePort

`func (o *InfrastructurePort) GetNodePort() int64`

GetNodePort returns the NodePort field if non-nil, zero value otherwise.

### GetNodePortOk

`func (o *InfrastructurePort) GetNodePortOk() (*int64, bool)`

GetNodePortOk returns a tuple with the NodePort field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNodePort

`func (o *InfrastructurePort) SetNodePort(v int64)`

SetNodePort sets NodePort field to given value.

### HasNodePort

`func (o *InfrastructurePort) HasNodePort() bool`

HasNodePort returns a boolean if a field has been set.

### GetPort

`func (o *InfrastructurePort) GetPort() int64`

GetPort returns the Port field if non-nil, zero value otherwise.

### GetPortOk

`func (o *InfrastructurePort) GetPortOk() (*int64, bool)`

GetPortOk returns a tuple with the Port field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPort

`func (o *InfrastructurePort) SetPort(v int64)`

SetPort sets Port field to given value.

### HasPort

`func (o *InfrastructurePort) HasPort() bool`

HasPort returns a boolean if a field has been set.

### GetProtocol

`func (o *InfrastructurePort) GetProtocol() string`

GetProtocol returns the Protocol field if non-nil, zero value otherwise.

### GetProtocolOk

`func (o *InfrastructurePort) GetProtocolOk() (*string, bool)`

GetProtocolOk returns a tuple with the Protocol field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetProtocol

`func (o *InfrastructurePort) SetProtocol(v string)`

SetProtocol sets Protocol field to given value.

### HasProtocol

`func (o *InfrastructurePort) HasProtocol() bool`

HasProtocol returns a boolean if a field has been set.

### GetTargetPort

`func (o *InfrastructurePort) GetTargetPort() string`

GetTargetPort returns the TargetPort field if non-nil, zero value otherwise.

### GetTargetPortOk

`func (o *InfrastructurePort) GetTargetPortOk() (*string, bool)`

GetTargetPortOk returns a tuple with the TargetPort field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTargetPort

`func (o *InfrastructurePort) SetTargetPort(v string)`

SetTargetPort sets TargetPort field to given value.

### HasTargetPort

`func (o *InfrastructurePort) HasTargetPort() bool`

HasTargetPort returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


