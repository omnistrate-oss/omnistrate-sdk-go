# UpdateAccountConfigRequest2

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**AllowNewCloudNativeNetworkCreation** | Pointer to **bool** | Whether Omnistrate is allowed to create new CloudNativeNetworks in this account when no registered cloud native network is selected at deployment time | [optional] 
**Description** | Pointer to **string** | The updated description for the account | [optional] 
**Name** | Pointer to **string** | The updated name of the account | [optional] 
**NebiusBindings** | Pointer to [**[]NebiusAccountBindingInput**](NebiusAccountBindingInput.md) | Full replacement set of Nebius project/service-account bindings for an existing tenant-scoped Nebius account configuration | [optional] 

## Methods

### NewUpdateAccountConfigRequest2

`func NewUpdateAccountConfigRequest2() *UpdateAccountConfigRequest2`

NewUpdateAccountConfigRequest2 instantiates a new UpdateAccountConfigRequest2 object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewUpdateAccountConfigRequest2WithDefaults

`func NewUpdateAccountConfigRequest2WithDefaults() *UpdateAccountConfigRequest2`

NewUpdateAccountConfigRequest2WithDefaults instantiates a new UpdateAccountConfigRequest2 object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetAllowNewCloudNativeNetworkCreation

`func (o *UpdateAccountConfigRequest2) GetAllowNewCloudNativeNetworkCreation() bool`

GetAllowNewCloudNativeNetworkCreation returns the AllowNewCloudNativeNetworkCreation field if non-nil, zero value otherwise.

### GetAllowNewCloudNativeNetworkCreationOk

`func (o *UpdateAccountConfigRequest2) GetAllowNewCloudNativeNetworkCreationOk() (*bool, bool)`

GetAllowNewCloudNativeNetworkCreationOk returns a tuple with the AllowNewCloudNativeNetworkCreation field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAllowNewCloudNativeNetworkCreation

`func (o *UpdateAccountConfigRequest2) SetAllowNewCloudNativeNetworkCreation(v bool)`

SetAllowNewCloudNativeNetworkCreation sets AllowNewCloudNativeNetworkCreation field to given value.

### HasAllowNewCloudNativeNetworkCreation

`func (o *UpdateAccountConfigRequest2) HasAllowNewCloudNativeNetworkCreation() bool`

HasAllowNewCloudNativeNetworkCreation returns a boolean if a field has been set.

### GetDescription

`func (o *UpdateAccountConfigRequest2) GetDescription() string`

GetDescription returns the Description field if non-nil, zero value otherwise.

### GetDescriptionOk

`func (o *UpdateAccountConfigRequest2) GetDescriptionOk() (*string, bool)`

GetDescriptionOk returns a tuple with the Description field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDescription

`func (o *UpdateAccountConfigRequest2) SetDescription(v string)`

SetDescription sets Description field to given value.

### HasDescription

`func (o *UpdateAccountConfigRequest2) HasDescription() bool`

HasDescription returns a boolean if a field has been set.

### GetName

`func (o *UpdateAccountConfigRequest2) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *UpdateAccountConfigRequest2) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *UpdateAccountConfigRequest2) SetName(v string)`

SetName sets Name field to given value.

### HasName

`func (o *UpdateAccountConfigRequest2) HasName() bool`

HasName returns a boolean if a field has been set.

### GetNebiusBindings

`func (o *UpdateAccountConfigRequest2) GetNebiusBindings() []NebiusAccountBindingInput`

GetNebiusBindings returns the NebiusBindings field if non-nil, zero value otherwise.

### GetNebiusBindingsOk

`func (o *UpdateAccountConfigRequest2) GetNebiusBindingsOk() (*[]NebiusAccountBindingInput, bool)`

GetNebiusBindingsOk returns a tuple with the NebiusBindings field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNebiusBindings

`func (o *UpdateAccountConfigRequest2) SetNebiusBindings(v []NebiusAccountBindingInput)`

SetNebiusBindings sets NebiusBindings field to given value.

### HasNebiusBindings

`func (o *UpdateAccountConfigRequest2) HasNebiusBindings() bool`

HasNebiusBindings returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


