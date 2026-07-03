# ManagedWorkloadIdentityPermissions

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Permissions** | Pointer to **map[string][]string** | Cloud-provider permission statements assigned to the managed workload identity. Use this for OCI workload identity policies. | [optional] 
**Policies** | Pointer to **map[string]string** | Cloud-provider policy documents for the managed workload identity. Use this for AWS IAM policy documents. | [optional] 
**Roles** | Pointer to [**map[string][]ManagedWorkloadIdentityRole**](array.md) | Cloud-provider roles assigned to the managed workload identity. Use this for GCP IAM roles and Azure RBAC/Entra roles. | [optional] 

## Methods

### NewManagedWorkloadIdentityPermissions

`func NewManagedWorkloadIdentityPermissions() *ManagedWorkloadIdentityPermissions`

NewManagedWorkloadIdentityPermissions instantiates a new ManagedWorkloadIdentityPermissions object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewManagedWorkloadIdentityPermissionsWithDefaults

`func NewManagedWorkloadIdentityPermissionsWithDefaults() *ManagedWorkloadIdentityPermissions`

NewManagedWorkloadIdentityPermissionsWithDefaults instantiates a new ManagedWorkloadIdentityPermissions object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetPermissions

`func (o *ManagedWorkloadIdentityPermissions) GetPermissions() map[string][]string`

GetPermissions returns the Permissions field if non-nil, zero value otherwise.

### GetPermissionsOk

`func (o *ManagedWorkloadIdentityPermissions) GetPermissionsOk() (*map[string][]string, bool)`

GetPermissionsOk returns a tuple with the Permissions field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPermissions

`func (o *ManagedWorkloadIdentityPermissions) SetPermissions(v map[string][]string)`

SetPermissions sets Permissions field to given value.

### HasPermissions

`func (o *ManagedWorkloadIdentityPermissions) HasPermissions() bool`

HasPermissions returns a boolean if a field has been set.

### GetPolicies

`func (o *ManagedWorkloadIdentityPermissions) GetPolicies() map[string]string`

GetPolicies returns the Policies field if non-nil, zero value otherwise.

### GetPoliciesOk

`func (o *ManagedWorkloadIdentityPermissions) GetPoliciesOk() (*map[string]string, bool)`

GetPoliciesOk returns a tuple with the Policies field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPolicies

`func (o *ManagedWorkloadIdentityPermissions) SetPolicies(v map[string]string)`

SetPolicies sets Policies field to given value.

### HasPolicies

`func (o *ManagedWorkloadIdentityPermissions) HasPolicies() bool`

HasPolicies returns a boolean if a field has been set.

### GetRoles

`func (o *ManagedWorkloadIdentityPermissions) GetRoles() map[string][]ManagedWorkloadIdentityRole`

GetRoles returns the Roles field if non-nil, zero value otherwise.

### GetRolesOk

`func (o *ManagedWorkloadIdentityPermissions) GetRolesOk() (*map[string][]ManagedWorkloadIdentityRole, bool)`

GetRolesOk returns a tuple with the Roles field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRoles

`func (o *ManagedWorkloadIdentityPermissions) SetRoles(v map[string][]ManagedWorkloadIdentityRole)`

SetRoles sets Roles field to given value.

### HasRoles

`func (o *ManagedWorkloadIdentityPermissions) HasRoles() bool`

HasRoles returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


