# CreateAccessRequestRequest

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**DeploymentCellId** | **string** | ID of a Host Cluster | 
**Reason** | Pointer to **string** | A nonblank access reason is required in MANUAL. Missing configuration uses automatic approval when a request is explicitly created; callers that never create a request retain legacy access. | [optional] 
**Role** | **string** | Requested Kubernetes access role. Known values: cluster-reader, cluster-admin. The server validates authorization; role is not part of requester-plus-cell deduplication. Clients must preserve unknown future values. | 
**Token** | **string** | JWT token used to perform authorization | 

## Methods

### NewCreateAccessRequestRequest

`func NewCreateAccessRequestRequest(deploymentCellId string, role string, token string, ) *CreateAccessRequestRequest`

NewCreateAccessRequestRequest instantiates a new CreateAccessRequestRequest object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewCreateAccessRequestRequestWithDefaults

`func NewCreateAccessRequestRequestWithDefaults() *CreateAccessRequestRequest`

NewCreateAccessRequestRequestWithDefaults instantiates a new CreateAccessRequestRequest object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetDeploymentCellId

`func (o *CreateAccessRequestRequest) GetDeploymentCellId() string`

GetDeploymentCellId returns the DeploymentCellId field if non-nil, zero value otherwise.

### GetDeploymentCellIdOk

`func (o *CreateAccessRequestRequest) GetDeploymentCellIdOk() (*string, bool)`

GetDeploymentCellIdOk returns a tuple with the DeploymentCellId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDeploymentCellId

`func (o *CreateAccessRequestRequest) SetDeploymentCellId(v string)`

SetDeploymentCellId sets DeploymentCellId field to given value.


### GetReason

`func (o *CreateAccessRequestRequest) GetReason() string`

GetReason returns the Reason field if non-nil, zero value otherwise.

### GetReasonOk

`func (o *CreateAccessRequestRequest) GetReasonOk() (*string, bool)`

GetReasonOk returns a tuple with the Reason field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetReason

`func (o *CreateAccessRequestRequest) SetReason(v string)`

SetReason sets Reason field to given value.

### HasReason

`func (o *CreateAccessRequestRequest) HasReason() bool`

HasReason returns a boolean if a field has been set.

### GetRole

`func (o *CreateAccessRequestRequest) GetRole() string`

GetRole returns the Role field if non-nil, zero value otherwise.

### GetRoleOk

`func (o *CreateAccessRequestRequest) GetRoleOk() (*string, bool)`

GetRoleOk returns a tuple with the Role field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRole

`func (o *CreateAccessRequestRequest) SetRole(v string)`

SetRole sets Role field to given value.


### GetToken

`func (o *CreateAccessRequestRequest) GetToken() string`

GetToken returns the Token field if non-nil, zero value otherwise.

### GetTokenOk

`func (o *CreateAccessRequestRequest) GetTokenOk() (*string, bool)`

GetTokenOk returns a tuple with the Token field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetToken

`func (o *CreateAccessRequestRequest) SetToken(v string)`

SetToken sets Token field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


