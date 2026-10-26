# AccountGovernanceControls

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Connected** | Pointer to **bool** | Whether dataplane connectivity to the control plane is requested. Defaults to true for requested settings. Disconnect overrides the effective debug and infrastructure permissions without discarding those preferences. | [optional] 
**DebugAccessEnabled** | Pointer to **bool** | Whether Kubernetes debug access is requested. Defaults to false for AWS requested settings; omitted for other cloud providers unless explicitly configured. | [optional] 
**InfrastructureMutationEnabled** | Pointer to **bool** | Whether the dataplane agent may change cloud infrastructure. Defaults to true for AWS requested settings; omitted for other cloud providers unless explicitly configured. | [optional] 

## Methods

### NewAccountGovernanceControls

`func NewAccountGovernanceControls() *AccountGovernanceControls`

NewAccountGovernanceControls instantiates a new AccountGovernanceControls object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewAccountGovernanceControlsWithDefaults

`func NewAccountGovernanceControlsWithDefaults() *AccountGovernanceControls`

NewAccountGovernanceControlsWithDefaults instantiates a new AccountGovernanceControls object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetConnected

`func (o *AccountGovernanceControls) GetConnected() bool`

GetConnected returns the Connected field if non-nil, zero value otherwise.

### GetConnectedOk

`func (o *AccountGovernanceControls) GetConnectedOk() (*bool, bool)`

GetConnectedOk returns a tuple with the Connected field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetConnected

`func (o *AccountGovernanceControls) SetConnected(v bool)`

SetConnected sets Connected field to given value.

### HasConnected

`func (o *AccountGovernanceControls) HasConnected() bool`

HasConnected returns a boolean if a field has been set.

### GetDebugAccessEnabled

`func (o *AccountGovernanceControls) GetDebugAccessEnabled() bool`

GetDebugAccessEnabled returns the DebugAccessEnabled field if non-nil, zero value otherwise.

### GetDebugAccessEnabledOk

`func (o *AccountGovernanceControls) GetDebugAccessEnabledOk() (*bool, bool)`

GetDebugAccessEnabledOk returns a tuple with the DebugAccessEnabled field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDebugAccessEnabled

`func (o *AccountGovernanceControls) SetDebugAccessEnabled(v bool)`

SetDebugAccessEnabled sets DebugAccessEnabled field to given value.

### HasDebugAccessEnabled

`func (o *AccountGovernanceControls) HasDebugAccessEnabled() bool`

HasDebugAccessEnabled returns a boolean if a field has been set.

### GetInfrastructureMutationEnabled

`func (o *AccountGovernanceControls) GetInfrastructureMutationEnabled() bool`

GetInfrastructureMutationEnabled returns the InfrastructureMutationEnabled field if non-nil, zero value otherwise.

### GetInfrastructureMutationEnabledOk

`func (o *AccountGovernanceControls) GetInfrastructureMutationEnabledOk() (*bool, bool)`

GetInfrastructureMutationEnabledOk returns a tuple with the InfrastructureMutationEnabled field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetInfrastructureMutationEnabled

`func (o *AccountGovernanceControls) SetInfrastructureMutationEnabled(v bool)`

SetInfrastructureMutationEnabled sets InfrastructureMutationEnabled field to given value.

### HasInfrastructureMutationEnabled

`func (o *AccountGovernanceControls) HasInfrastructureMutationEnabled() bool`

HasInfrastructureMutationEnabled returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


