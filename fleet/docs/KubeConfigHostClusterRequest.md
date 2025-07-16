# KubeConfigHostClusterRequest

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Id** | **string** | ID of a Host Cluster | 
**Role** | Pointer to **string** | Role of the service account to use in the kubeconfig | [optional] [default to "cluster-reader"]
**Token** | **string** | JWT token used to perform authorization | 

## Methods

### NewKubeConfigHostClusterRequest

`func NewKubeConfigHostClusterRequest(id string, token string, ) *KubeConfigHostClusterRequest`

NewKubeConfigHostClusterRequest instantiates a new KubeConfigHostClusterRequest object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewKubeConfigHostClusterRequestWithDefaults

`func NewKubeConfigHostClusterRequestWithDefaults() *KubeConfigHostClusterRequest`

NewKubeConfigHostClusterRequestWithDefaults instantiates a new KubeConfigHostClusterRequest object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetId

`func (o *KubeConfigHostClusterRequest) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *KubeConfigHostClusterRequest) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *KubeConfigHostClusterRequest) SetId(v string)`

SetId sets Id field to given value.


### GetRole

`func (o *KubeConfigHostClusterRequest) GetRole() string`

GetRole returns the Role field if non-nil, zero value otherwise.

### GetRoleOk

`func (o *KubeConfigHostClusterRequest) GetRoleOk() (*string, bool)`

GetRoleOk returns a tuple with the Role field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRole

`func (o *KubeConfigHostClusterRequest) SetRole(v string)`

SetRole sets Role field to given value.

### HasRole

`func (o *KubeConfigHostClusterRequest) HasRole() bool`

HasRole returns a boolean if a field has been set.

### GetToken

`func (o *KubeConfigHostClusterRequest) GetToken() string`

GetToken returns the Token field if non-nil, zero value otherwise.

### GetTokenOk

`func (o *KubeConfigHostClusterRequest) GetTokenOk() (*string, bool)`

GetTokenOk returns a tuple with the Token field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetToken

`func (o *KubeConfigHostClusterRequest) SetToken(v string)`

SetToken sets Token field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


