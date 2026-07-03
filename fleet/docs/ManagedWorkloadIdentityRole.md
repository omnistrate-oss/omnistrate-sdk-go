# ManagedWorkloadIdentityRole

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Name** | **string** | The cloud-provider role name. | 
**Type** | Pointer to **string** | The cloud-provider role type. Default is used for GCP IAM roles and Azure RBAC roles. Entra is used for Azure Entra roles. | [optional] 

## Methods

### NewManagedWorkloadIdentityRole

`func NewManagedWorkloadIdentityRole(name string, ) *ManagedWorkloadIdentityRole`

NewManagedWorkloadIdentityRole instantiates a new ManagedWorkloadIdentityRole object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewManagedWorkloadIdentityRoleWithDefaults

`func NewManagedWorkloadIdentityRoleWithDefaults() *ManagedWorkloadIdentityRole`

NewManagedWorkloadIdentityRoleWithDefaults instantiates a new ManagedWorkloadIdentityRole object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetName

`func (o *ManagedWorkloadIdentityRole) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *ManagedWorkloadIdentityRole) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *ManagedWorkloadIdentityRole) SetName(v string)`

SetName sets Name field to given value.


### GetType

`func (o *ManagedWorkloadIdentityRole) GetType() string`

GetType returns the Type field if non-nil, zero value otherwise.

### GetTypeOk

`func (o *ManagedWorkloadIdentityRole) GetTypeOk() (*string, bool)`

GetTypeOk returns a tuple with the Type field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetType

`func (o *ManagedWorkloadIdentityRole) SetType(v string)`

SetType sets Type field to given value.

### HasType

`func (o *ManagedWorkloadIdentityRole) HasType() bool`

HasType returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


