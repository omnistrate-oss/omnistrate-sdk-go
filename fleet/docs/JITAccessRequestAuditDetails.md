# JITAccessRequestAuditDetails

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**ApprovalMethod** | Pointer to **string** | The approval method, MANUAL or AUTOMATIC. Omitted before approval. | [optional] 
**CreatedAt** | Pointer to **time.Time** | The time the access request was created. | [optional] 
**DeploymentCellId** | Pointer to **string** | ID of a Host Cluster | [optional] 
**ExpiredAt** | Pointer to **time.Time** | The effective expiration time, populated only for expiration events. | [optional] 
**ExpiresAt** | Pointer to **time.Time** | The immutable grant expiry. Omitted before approval. | [optional] 
**GrantStartedAt** | Pointer to **time.Time** | The original grant start time. Omitted before approval. | [optional] 
**GrantedDurationSeconds** | Pointer to **int64** | The original approved duration in seconds. Omitted before approval. | [optional] 
**Id** | Pointer to **string** | Server-generated ID of a provider cluster access request. | [optional] 
**RequestedRole** | Pointer to **string** | Requested Kubernetes access role. Known values: cluster-reader, cluster-admin. The server validates authorization; role is not part of requester-plus-cell deduplication. Clients must preserve unknown future values. | [optional] 
**Requester** | Pointer to [**JITAccessEventRequester**](JITAccessEventRequester.md) |  | [optional] 
**RevokedAt** | Pointer to **time.Time** | The revocation time, populated only for revocation events. | [optional] 

## Methods

### NewJITAccessRequestAuditDetails

`func NewJITAccessRequestAuditDetails() *JITAccessRequestAuditDetails`

NewJITAccessRequestAuditDetails instantiates a new JITAccessRequestAuditDetails object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewJITAccessRequestAuditDetailsWithDefaults

`func NewJITAccessRequestAuditDetailsWithDefaults() *JITAccessRequestAuditDetails`

NewJITAccessRequestAuditDetailsWithDefaults instantiates a new JITAccessRequestAuditDetails object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetApprovalMethod

`func (o *JITAccessRequestAuditDetails) GetApprovalMethod() string`

GetApprovalMethod returns the ApprovalMethod field if non-nil, zero value otherwise.

### GetApprovalMethodOk

`func (o *JITAccessRequestAuditDetails) GetApprovalMethodOk() (*string, bool)`

GetApprovalMethodOk returns a tuple with the ApprovalMethod field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetApprovalMethod

`func (o *JITAccessRequestAuditDetails) SetApprovalMethod(v string)`

SetApprovalMethod sets ApprovalMethod field to given value.

### HasApprovalMethod

`func (o *JITAccessRequestAuditDetails) HasApprovalMethod() bool`

HasApprovalMethod returns a boolean if a field has been set.

### GetCreatedAt

`func (o *JITAccessRequestAuditDetails) GetCreatedAt() time.Time`

GetCreatedAt returns the CreatedAt field if non-nil, zero value otherwise.

### GetCreatedAtOk

`func (o *JITAccessRequestAuditDetails) GetCreatedAtOk() (*time.Time, bool)`

GetCreatedAtOk returns a tuple with the CreatedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCreatedAt

`func (o *JITAccessRequestAuditDetails) SetCreatedAt(v time.Time)`

SetCreatedAt sets CreatedAt field to given value.

### HasCreatedAt

`func (o *JITAccessRequestAuditDetails) HasCreatedAt() bool`

HasCreatedAt returns a boolean if a field has been set.

### GetDeploymentCellId

`func (o *JITAccessRequestAuditDetails) GetDeploymentCellId() string`

GetDeploymentCellId returns the DeploymentCellId field if non-nil, zero value otherwise.

### GetDeploymentCellIdOk

`func (o *JITAccessRequestAuditDetails) GetDeploymentCellIdOk() (*string, bool)`

GetDeploymentCellIdOk returns a tuple with the DeploymentCellId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDeploymentCellId

`func (o *JITAccessRequestAuditDetails) SetDeploymentCellId(v string)`

SetDeploymentCellId sets DeploymentCellId field to given value.

### HasDeploymentCellId

`func (o *JITAccessRequestAuditDetails) HasDeploymentCellId() bool`

HasDeploymentCellId returns a boolean if a field has been set.

### GetExpiredAt

`func (o *JITAccessRequestAuditDetails) GetExpiredAt() time.Time`

GetExpiredAt returns the ExpiredAt field if non-nil, zero value otherwise.

### GetExpiredAtOk

`func (o *JITAccessRequestAuditDetails) GetExpiredAtOk() (*time.Time, bool)`

GetExpiredAtOk returns a tuple with the ExpiredAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetExpiredAt

`func (o *JITAccessRequestAuditDetails) SetExpiredAt(v time.Time)`

SetExpiredAt sets ExpiredAt field to given value.

### HasExpiredAt

`func (o *JITAccessRequestAuditDetails) HasExpiredAt() bool`

HasExpiredAt returns a boolean if a field has been set.

### GetExpiresAt

`func (o *JITAccessRequestAuditDetails) GetExpiresAt() time.Time`

GetExpiresAt returns the ExpiresAt field if non-nil, zero value otherwise.

### GetExpiresAtOk

`func (o *JITAccessRequestAuditDetails) GetExpiresAtOk() (*time.Time, bool)`

GetExpiresAtOk returns a tuple with the ExpiresAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetExpiresAt

`func (o *JITAccessRequestAuditDetails) SetExpiresAt(v time.Time)`

SetExpiresAt sets ExpiresAt field to given value.

### HasExpiresAt

`func (o *JITAccessRequestAuditDetails) HasExpiresAt() bool`

HasExpiresAt returns a boolean if a field has been set.

### GetGrantStartedAt

`func (o *JITAccessRequestAuditDetails) GetGrantStartedAt() time.Time`

GetGrantStartedAt returns the GrantStartedAt field if non-nil, zero value otherwise.

### GetGrantStartedAtOk

`func (o *JITAccessRequestAuditDetails) GetGrantStartedAtOk() (*time.Time, bool)`

GetGrantStartedAtOk returns a tuple with the GrantStartedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetGrantStartedAt

`func (o *JITAccessRequestAuditDetails) SetGrantStartedAt(v time.Time)`

SetGrantStartedAt sets GrantStartedAt field to given value.

### HasGrantStartedAt

`func (o *JITAccessRequestAuditDetails) HasGrantStartedAt() bool`

HasGrantStartedAt returns a boolean if a field has been set.

### GetGrantedDurationSeconds

`func (o *JITAccessRequestAuditDetails) GetGrantedDurationSeconds() int64`

GetGrantedDurationSeconds returns the GrantedDurationSeconds field if non-nil, zero value otherwise.

### GetGrantedDurationSecondsOk

`func (o *JITAccessRequestAuditDetails) GetGrantedDurationSecondsOk() (*int64, bool)`

GetGrantedDurationSecondsOk returns a tuple with the GrantedDurationSeconds field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetGrantedDurationSeconds

`func (o *JITAccessRequestAuditDetails) SetGrantedDurationSeconds(v int64)`

SetGrantedDurationSeconds sets GrantedDurationSeconds field to given value.

### HasGrantedDurationSeconds

`func (o *JITAccessRequestAuditDetails) HasGrantedDurationSeconds() bool`

HasGrantedDurationSeconds returns a boolean if a field has been set.

### GetId

`func (o *JITAccessRequestAuditDetails) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *JITAccessRequestAuditDetails) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *JITAccessRequestAuditDetails) SetId(v string)`

SetId sets Id field to given value.

### HasId

`func (o *JITAccessRequestAuditDetails) HasId() bool`

HasId returns a boolean if a field has been set.

### GetRequestedRole

`func (o *JITAccessRequestAuditDetails) GetRequestedRole() string`

GetRequestedRole returns the RequestedRole field if non-nil, zero value otherwise.

### GetRequestedRoleOk

`func (o *JITAccessRequestAuditDetails) GetRequestedRoleOk() (*string, bool)`

GetRequestedRoleOk returns a tuple with the RequestedRole field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRequestedRole

`func (o *JITAccessRequestAuditDetails) SetRequestedRole(v string)`

SetRequestedRole sets RequestedRole field to given value.

### HasRequestedRole

`func (o *JITAccessRequestAuditDetails) HasRequestedRole() bool`

HasRequestedRole returns a boolean if a field has been set.

### GetRequester

`func (o *JITAccessRequestAuditDetails) GetRequester() JITAccessEventRequester`

GetRequester returns the Requester field if non-nil, zero value otherwise.

### GetRequesterOk

`func (o *JITAccessRequestAuditDetails) GetRequesterOk() (*JITAccessEventRequester, bool)`

GetRequesterOk returns a tuple with the Requester field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRequester

`func (o *JITAccessRequestAuditDetails) SetRequester(v JITAccessEventRequester)`

SetRequester sets Requester field to given value.

### HasRequester

`func (o *JITAccessRequestAuditDetails) HasRequester() bool`

HasRequester returns a boolean if a field has been set.

### GetRevokedAt

`func (o *JITAccessRequestAuditDetails) GetRevokedAt() time.Time`

GetRevokedAt returns the RevokedAt field if non-nil, zero value otherwise.

### GetRevokedAtOk

`func (o *JITAccessRequestAuditDetails) GetRevokedAtOk() (*time.Time, bool)`

GetRevokedAtOk returns a tuple with the RevokedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRevokedAt

`func (o *JITAccessRequestAuditDetails) SetRevokedAt(v time.Time)`

SetRevokedAt sets RevokedAt field to given value.

### HasRevokedAt

`func (o *JITAccessRequestAuditDetails) HasRevokedAt() bool`

HasRevokedAt returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


