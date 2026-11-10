# CreateAccessRequestRequest2

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**DeploymentCellId** | **string** | The target deployment cell ID. Provider organization and environment are resolved and authorized by the service. | 
**Reason** | Pointer to **string** | A nonblank access reason is required in MANUAL. Missing configuration uses automatic approval when a request is explicitly created; callers that never create a request retain legacy access. | [optional] 
**Role** | **string** | The requested cluster-reader or cluster-admin role. Validated by the server and not part of the requester-plus-cell deduplication key. | 

## Methods

### NewCreateAccessRequestRequest2

`func NewCreateAccessRequestRequest2(deploymentCellId string, role string, ) *CreateAccessRequestRequest2`

NewCreateAccessRequestRequest2 instantiates a new CreateAccessRequestRequest2 object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewCreateAccessRequestRequest2WithDefaults

`func NewCreateAccessRequestRequest2WithDefaults() *CreateAccessRequestRequest2`

NewCreateAccessRequestRequest2WithDefaults instantiates a new CreateAccessRequestRequest2 object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetDeploymentCellId

`func (o *CreateAccessRequestRequest2) GetDeploymentCellId() string`

GetDeploymentCellId returns the DeploymentCellId field if non-nil, zero value otherwise.

### GetDeploymentCellIdOk

`func (o *CreateAccessRequestRequest2) GetDeploymentCellIdOk() (*string, bool)`

GetDeploymentCellIdOk returns a tuple with the DeploymentCellId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDeploymentCellId

`func (o *CreateAccessRequestRequest2) SetDeploymentCellId(v string)`

SetDeploymentCellId sets DeploymentCellId field to given value.


### GetReason

`func (o *CreateAccessRequestRequest2) GetReason() string`

GetReason returns the Reason field if non-nil, zero value otherwise.

### GetReasonOk

`func (o *CreateAccessRequestRequest2) GetReasonOk() (*string, bool)`

GetReasonOk returns a tuple with the Reason field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetReason

`func (o *CreateAccessRequestRequest2) SetReason(v string)`

SetReason sets Reason field to given value.

### HasReason

`func (o *CreateAccessRequestRequest2) HasReason() bool`

HasReason returns a boolean if a field has been set.

### GetRole

`func (o *CreateAccessRequestRequest2) GetRole() string`

GetRole returns the Role field if non-nil, zero value otherwise.

### GetRoleOk

`func (o *CreateAccessRequestRequest2) GetRoleOk() (*string, bool)`

GetRoleOk returns a tuple with the Role field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRole

`func (o *CreateAccessRequestRequest2) SetRole(v string)`

SetRole sets Role field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


