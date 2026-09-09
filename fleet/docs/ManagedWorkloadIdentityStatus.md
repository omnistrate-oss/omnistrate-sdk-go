# ManagedWorkloadIdentityStatus

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**CloudIdentifier** | Pointer to **string** | The cloud-specific identifier created for this identity (Azure client ID, GCP service account email, or OCI policy OCID) | [optional] 
**Name** | **string** | The managed workload identity name | 
**Status** | **string** | The current status of this managed workload identity | 

## Methods

### NewManagedWorkloadIdentityStatus

`func NewManagedWorkloadIdentityStatus(name string, status string, ) *ManagedWorkloadIdentityStatus`

NewManagedWorkloadIdentityStatus instantiates a new ManagedWorkloadIdentityStatus object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewManagedWorkloadIdentityStatusWithDefaults

`func NewManagedWorkloadIdentityStatusWithDefaults() *ManagedWorkloadIdentityStatus`

NewManagedWorkloadIdentityStatusWithDefaults instantiates a new ManagedWorkloadIdentityStatus object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetCloudIdentifier

`func (o *ManagedWorkloadIdentityStatus) GetCloudIdentifier() string`

GetCloudIdentifier returns the CloudIdentifier field if non-nil, zero value otherwise.

### GetCloudIdentifierOk

`func (o *ManagedWorkloadIdentityStatus) GetCloudIdentifierOk() (*string, bool)`

GetCloudIdentifierOk returns a tuple with the CloudIdentifier field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCloudIdentifier

`func (o *ManagedWorkloadIdentityStatus) SetCloudIdentifier(v string)`

SetCloudIdentifier sets CloudIdentifier field to given value.

### HasCloudIdentifier

`func (o *ManagedWorkloadIdentityStatus) HasCloudIdentifier() bool`

HasCloudIdentifier returns a boolean if a field has been set.

### GetName

`func (o *ManagedWorkloadIdentityStatus) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *ManagedWorkloadIdentityStatus) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *ManagedWorkloadIdentityStatus) SetName(v string)`

SetName sets Name field to given value.


### GetStatus

`func (o *ManagedWorkloadIdentityStatus) GetStatus() string`

GetStatus returns the Status field if non-nil, zero value otherwise.

### GetStatusOk

`func (o *ManagedWorkloadIdentityStatus) GetStatusOk() (*string, bool)`

GetStatusOk returns a tuple with the Status field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStatus

`func (o *ManagedWorkloadIdentityStatus) SetStatus(v string)`

SetStatus sets Status field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


