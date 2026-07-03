# ManagedWorkloadIdentityServiceAccount

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Name** | **string** | The Kubernetes service account name. | 
**Namespace** | **string** | The Kubernetes namespace containing the service account. | 

## Methods

### NewManagedWorkloadIdentityServiceAccount

`func NewManagedWorkloadIdentityServiceAccount(name string, namespace string, ) *ManagedWorkloadIdentityServiceAccount`

NewManagedWorkloadIdentityServiceAccount instantiates a new ManagedWorkloadIdentityServiceAccount object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewManagedWorkloadIdentityServiceAccountWithDefaults

`func NewManagedWorkloadIdentityServiceAccountWithDefaults() *ManagedWorkloadIdentityServiceAccount`

NewManagedWorkloadIdentityServiceAccountWithDefaults instantiates a new ManagedWorkloadIdentityServiceAccount object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetName

`func (o *ManagedWorkloadIdentityServiceAccount) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *ManagedWorkloadIdentityServiceAccount) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *ManagedWorkloadIdentityServiceAccount) SetName(v string)`

SetName sets Name field to given value.


### GetNamespace

`func (o *ManagedWorkloadIdentityServiceAccount) GetNamespace() string`

GetNamespace returns the Namespace field if non-nil, zero value otherwise.

### GetNamespaceOk

`func (o *ManagedWorkloadIdentityServiceAccount) GetNamespaceOk() (*string, bool)`

GetNamespaceOk returns a tuple with the Namespace field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNamespace

`func (o *ManagedWorkloadIdentityServiceAccount) SetNamespace(v string)`

SetNamespace sets Namespace field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


