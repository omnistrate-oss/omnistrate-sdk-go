# NamedPortSpec

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Port** | Pointer to **int64** | Single port number (mutually exclusive with portRange) | [optional] 
**PortsRange** | Pointer to [**PortsRange**](PortsRange.md) |  | [optional] 

## Methods

### NewNamedPortSpec

`func NewNamedPortSpec() *NamedPortSpec`

NewNamedPortSpec instantiates a new NamedPortSpec object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewNamedPortSpecWithDefaults

`func NewNamedPortSpecWithDefaults() *NamedPortSpec`

NewNamedPortSpecWithDefaults instantiates a new NamedPortSpec object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetPort

`func (o *NamedPortSpec) GetPort() int64`

GetPort returns the Port field if non-nil, zero value otherwise.

### GetPortOk

`func (o *NamedPortSpec) GetPortOk() (*int64, bool)`

GetPortOk returns a tuple with the Port field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPort

`func (o *NamedPortSpec) SetPort(v int64)`

SetPort sets Port field to given value.

### HasPort

`func (o *NamedPortSpec) HasPort() bool`

HasPort returns a boolean if a field has been set.

### GetPortsRange

`func (o *NamedPortSpec) GetPortsRange() PortsRange`

GetPortsRange returns the PortsRange field if non-nil, zero value otherwise.

### GetPortsRangeOk

`func (o *NamedPortSpec) GetPortsRangeOk() (*PortsRange, bool)`

GetPortsRangeOk returns a tuple with the PortsRange field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPortsRange

`func (o *NamedPortSpec) SetPortsRange(v PortsRange)`

SetPortsRange sets PortsRange field to given value.

### HasPortsRange

`func (o *NamedPortSpec) HasPortsRange() bool`

HasPortsRange returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


