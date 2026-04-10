# FleetUpdateAccountConfigResourceInstanceRequest2

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**NebiusBindings** | Pointer to [**[]FleetUpdateAccountConfigNebiusBindingInput**](FleetUpdateAccountConfigNebiusBindingInput.md) | Full replacement set of Nebius bindings for this account config instance | [optional] 
**SetConnection** | Pointer to **bool** | set account config instance connection | [optional] 

## Methods

### NewFleetUpdateAccountConfigResourceInstanceRequest2

`func NewFleetUpdateAccountConfigResourceInstanceRequest2() *FleetUpdateAccountConfigResourceInstanceRequest2`

NewFleetUpdateAccountConfigResourceInstanceRequest2 instantiates a new FleetUpdateAccountConfigResourceInstanceRequest2 object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewFleetUpdateAccountConfigResourceInstanceRequest2WithDefaults

`func NewFleetUpdateAccountConfigResourceInstanceRequest2WithDefaults() *FleetUpdateAccountConfigResourceInstanceRequest2`

NewFleetUpdateAccountConfigResourceInstanceRequest2WithDefaults instantiates a new FleetUpdateAccountConfigResourceInstanceRequest2 object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetNebiusBindings

`func (o *FleetUpdateAccountConfigResourceInstanceRequest2) GetNebiusBindings() []FleetUpdateAccountConfigNebiusBindingInput`

GetNebiusBindings returns the NebiusBindings field if non-nil, zero value otherwise.

### GetNebiusBindingsOk

`func (o *FleetUpdateAccountConfigResourceInstanceRequest2) GetNebiusBindingsOk() (*[]FleetUpdateAccountConfigNebiusBindingInput, bool)`

GetNebiusBindingsOk returns a tuple with the NebiusBindings field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNebiusBindings

`func (o *FleetUpdateAccountConfigResourceInstanceRequest2) SetNebiusBindings(v []FleetUpdateAccountConfigNebiusBindingInput)`

SetNebiusBindings sets NebiusBindings field to given value.

### HasNebiusBindings

`func (o *FleetUpdateAccountConfigResourceInstanceRequest2) HasNebiusBindings() bool`

HasNebiusBindings returns a boolean if a field has been set.

### GetSetConnection

`func (o *FleetUpdateAccountConfigResourceInstanceRequest2) GetSetConnection() bool`

GetSetConnection returns the SetConnection field if non-nil, zero value otherwise.

### GetSetConnectionOk

`func (o *FleetUpdateAccountConfigResourceInstanceRequest2) GetSetConnectionOk() (*bool, bool)`

GetSetConnectionOk returns a tuple with the SetConnection field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSetConnection

`func (o *FleetUpdateAccountConfigResourceInstanceRequest2) SetSetConnection(v bool)`

SetSetConnection sets SetConnection field to given value.

### HasSetConnection

`func (o *FleetUpdateAccountConfigResourceInstanceRequest2) HasSetConnection() bool`

HasSetConnection returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


