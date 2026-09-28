# InfrastructureNetworkPolicy

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**PolicyTypes** | Pointer to **[]string** |  | [optional] 
**Ref** | Pointer to [**InfrastructureObjectReference**](InfrastructureObjectReference.md) |  | [optional] 
**Selector** | Pointer to **string** |  | [optional] 

## Methods

### NewInfrastructureNetworkPolicy

`func NewInfrastructureNetworkPolicy() *InfrastructureNetworkPolicy`

NewInfrastructureNetworkPolicy instantiates a new InfrastructureNetworkPolicy object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewInfrastructureNetworkPolicyWithDefaults

`func NewInfrastructureNetworkPolicyWithDefaults() *InfrastructureNetworkPolicy`

NewInfrastructureNetworkPolicyWithDefaults instantiates a new InfrastructureNetworkPolicy object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetPolicyTypes

`func (o *InfrastructureNetworkPolicy) GetPolicyTypes() []string`

GetPolicyTypes returns the PolicyTypes field if non-nil, zero value otherwise.

### GetPolicyTypesOk

`func (o *InfrastructureNetworkPolicy) GetPolicyTypesOk() (*[]string, bool)`

GetPolicyTypesOk returns a tuple with the PolicyTypes field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPolicyTypes

`func (o *InfrastructureNetworkPolicy) SetPolicyTypes(v []string)`

SetPolicyTypes sets PolicyTypes field to given value.

### HasPolicyTypes

`func (o *InfrastructureNetworkPolicy) HasPolicyTypes() bool`

HasPolicyTypes returns a boolean if a field has been set.

### GetRef

`func (o *InfrastructureNetworkPolicy) GetRef() InfrastructureObjectReference`

GetRef returns the Ref field if non-nil, zero value otherwise.

### GetRefOk

`func (o *InfrastructureNetworkPolicy) GetRefOk() (*InfrastructureObjectReference, bool)`

GetRefOk returns a tuple with the Ref field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRef

`func (o *InfrastructureNetworkPolicy) SetRef(v InfrastructureObjectReference)`

SetRef sets Ref field to given value.

### HasRef

`func (o *InfrastructureNetworkPolicy) HasRef() bool`

HasRef returns a boolean if a field has been set.

### GetSelector

`func (o *InfrastructureNetworkPolicy) GetSelector() string`

GetSelector returns the Selector field if non-nil, zero value otherwise.

### GetSelectorOk

`func (o *InfrastructureNetworkPolicy) GetSelectorOk() (*string, bool)`

GetSelectorOk returns a tuple with the Selector field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSelector

`func (o *InfrastructureNetworkPolicy) SetSelector(v string)`

SetSelector sets Selector field to given value.

### HasSelector

`func (o *InfrastructureNetworkPolicy) HasSelector() bool`

HasSelector returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


