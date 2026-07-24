# ManagedWorkloadIdentityBinding

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**ServiceAccount** | [**ManagedWorkloadIdentityServiceAccount**](ManagedWorkloadIdentityServiceAccount.md) |  | 

## Methods

### NewManagedWorkloadIdentityBinding

`func NewManagedWorkloadIdentityBinding(serviceAccount ManagedWorkloadIdentityServiceAccount, ) *ManagedWorkloadIdentityBinding`

NewManagedWorkloadIdentityBinding instantiates a new ManagedWorkloadIdentityBinding object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewManagedWorkloadIdentityBindingWithDefaults

`func NewManagedWorkloadIdentityBindingWithDefaults() *ManagedWorkloadIdentityBinding`

NewManagedWorkloadIdentityBindingWithDefaults instantiates a new ManagedWorkloadIdentityBinding object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetServiceAccount

`func (o *ManagedWorkloadIdentityBinding) GetServiceAccount() ManagedWorkloadIdentityServiceAccount`

GetServiceAccount returns the ServiceAccount field if non-nil, zero value otherwise.

### GetServiceAccountOk

`func (o *ManagedWorkloadIdentityBinding) GetServiceAccountOk() (*ManagedWorkloadIdentityServiceAccount, bool)`

GetServiceAccountOk returns a tuple with the ServiceAccount field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetServiceAccount

`func (o *ManagedWorkloadIdentityBinding) SetServiceAccount(v ManagedWorkloadIdentityServiceAccount)`

SetServiceAccount sets ServiceAccount field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


