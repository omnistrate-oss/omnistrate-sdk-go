# ManagedWorkloadIdentity

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Bindings** | Pointer to [**[]ManagedWorkloadIdentityBinding**](ManagedWorkloadIdentityBinding.md) | The workload identity bindings for this managed workload identity. API validation accepts service account bindings for OCI only. | [optional] 
**Description** | Pointer to **string** | A description of the managed workload identity. | [optional] 
**Identifier** | **string** | The identifier of the managed workload identity. | 
**Permissions** | Pointer to [**ManagedWorkloadIdentityPermissions**](ManagedWorkloadIdentityPermissions.md) |  | [optional] 

## Methods

### NewManagedWorkloadIdentity

`func NewManagedWorkloadIdentity(identifier string, ) *ManagedWorkloadIdentity`

NewManagedWorkloadIdentity instantiates a new ManagedWorkloadIdentity object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewManagedWorkloadIdentityWithDefaults

`func NewManagedWorkloadIdentityWithDefaults() *ManagedWorkloadIdentity`

NewManagedWorkloadIdentityWithDefaults instantiates a new ManagedWorkloadIdentity object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetBindings

`func (o *ManagedWorkloadIdentity) GetBindings() []ManagedWorkloadIdentityBinding`

GetBindings returns the Bindings field if non-nil, zero value otherwise.

### GetBindingsOk

`func (o *ManagedWorkloadIdentity) GetBindingsOk() (*[]ManagedWorkloadIdentityBinding, bool)`

GetBindingsOk returns a tuple with the Bindings field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBindings

`func (o *ManagedWorkloadIdentity) SetBindings(v []ManagedWorkloadIdentityBinding)`

SetBindings sets Bindings field to given value.

### HasBindings

`func (o *ManagedWorkloadIdentity) HasBindings() bool`

HasBindings returns a boolean if a field has been set.

### GetDescription

`func (o *ManagedWorkloadIdentity) GetDescription() string`

GetDescription returns the Description field if non-nil, zero value otherwise.

### GetDescriptionOk

`func (o *ManagedWorkloadIdentity) GetDescriptionOk() (*string, bool)`

GetDescriptionOk returns a tuple with the Description field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDescription

`func (o *ManagedWorkloadIdentity) SetDescription(v string)`

SetDescription sets Description field to given value.

### HasDescription

`func (o *ManagedWorkloadIdentity) HasDescription() bool`

HasDescription returns a boolean if a field has been set.

### GetIdentifier

`func (o *ManagedWorkloadIdentity) GetIdentifier() string`

GetIdentifier returns the Identifier field if non-nil, zero value otherwise.

### GetIdentifierOk

`func (o *ManagedWorkloadIdentity) GetIdentifierOk() (*string, bool)`

GetIdentifierOk returns a tuple with the Identifier field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIdentifier

`func (o *ManagedWorkloadIdentity) SetIdentifier(v string)`

SetIdentifier sets Identifier field to given value.


### GetPermissions

`func (o *ManagedWorkloadIdentity) GetPermissions() ManagedWorkloadIdentityPermissions`

GetPermissions returns the Permissions field if non-nil, zero value otherwise.

### GetPermissionsOk

`func (o *ManagedWorkloadIdentity) GetPermissionsOk() (*ManagedWorkloadIdentityPermissions, bool)`

GetPermissionsOk returns a tuple with the Permissions field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPermissions

`func (o *ManagedWorkloadIdentity) SetPermissions(v ManagedWorkloadIdentityPermissions)`

SetPermissions sets Permissions field to given value.

### HasPermissions

`func (o *ManagedWorkloadIdentity) HasPermissions() bool`

HasPermissions returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


