# APIKeyMetadata

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**CreatedAt** | **time.Time** | Creation timestamp. | 
**CreatedByUserId** | **string** | ID of a User | 
**Description** | Pointer to **string** | Optional free-text description. Editable via Update metadata. | [optional] 
**ExpiresAt** | Pointer to **time.Time** | Optional expiry. Null for non-expiring keys. | [optional] 
**Id** | **string** | ID of an API Key | 
**LastUsedAt** | Pointer to **time.Time** | Timestamp of the most recent successful authentication. Null when the key has never been used. System-managed; updated monotonically (never moves backwards). | [optional] 
**Name** | **string** | Operator-facing display name. Unique within the org. Editable via Update metadata. | 
**Prefix** | **string** | The fixed prefix label embedded in the plaintext at creation time. The only characters of the plaintext that are ever stored or returned after creation. Immutable. | 
**RevokedAt** | Pointer to **time.Time** | Revocation timestamp. Null when the key is not revoked. System-managed; set by the Revoke endpoint. | [optional] 
**RevokedByUserId** | Pointer to **string** | ID of a User | [optional] 
**RoleType** | **string** | Type of the role | 
**Status** | **string** | Derived lifecycle status of the API key, computed from revokedAt and expiresAt at read time. | 

## Methods

### NewAPIKeyMetadata

`func NewAPIKeyMetadata(createdAt time.Time, createdByUserId string, id string, name string, prefix string, roleType string, status string, ) *APIKeyMetadata`

NewAPIKeyMetadata instantiates a new APIKeyMetadata object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewAPIKeyMetadataWithDefaults

`func NewAPIKeyMetadataWithDefaults() *APIKeyMetadata`

NewAPIKeyMetadataWithDefaults instantiates a new APIKeyMetadata object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetCreatedAt

`func (o *APIKeyMetadata) GetCreatedAt() time.Time`

GetCreatedAt returns the CreatedAt field if non-nil, zero value otherwise.

### GetCreatedAtOk

`func (o *APIKeyMetadata) GetCreatedAtOk() (*time.Time, bool)`

GetCreatedAtOk returns a tuple with the CreatedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCreatedAt

`func (o *APIKeyMetadata) SetCreatedAt(v time.Time)`

SetCreatedAt sets CreatedAt field to given value.


### GetCreatedByUserId

`func (o *APIKeyMetadata) GetCreatedByUserId() string`

GetCreatedByUserId returns the CreatedByUserId field if non-nil, zero value otherwise.

### GetCreatedByUserIdOk

`func (o *APIKeyMetadata) GetCreatedByUserIdOk() (*string, bool)`

GetCreatedByUserIdOk returns a tuple with the CreatedByUserId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCreatedByUserId

`func (o *APIKeyMetadata) SetCreatedByUserId(v string)`

SetCreatedByUserId sets CreatedByUserId field to given value.


### GetDescription

`func (o *APIKeyMetadata) GetDescription() string`

GetDescription returns the Description field if non-nil, zero value otherwise.

### GetDescriptionOk

`func (o *APIKeyMetadata) GetDescriptionOk() (*string, bool)`

GetDescriptionOk returns a tuple with the Description field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDescription

`func (o *APIKeyMetadata) SetDescription(v string)`

SetDescription sets Description field to given value.

### HasDescription

`func (o *APIKeyMetadata) HasDescription() bool`

HasDescription returns a boolean if a field has been set.

### GetExpiresAt

`func (o *APIKeyMetadata) GetExpiresAt() time.Time`

GetExpiresAt returns the ExpiresAt field if non-nil, zero value otherwise.

### GetExpiresAtOk

`func (o *APIKeyMetadata) GetExpiresAtOk() (*time.Time, bool)`

GetExpiresAtOk returns a tuple with the ExpiresAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetExpiresAt

`func (o *APIKeyMetadata) SetExpiresAt(v time.Time)`

SetExpiresAt sets ExpiresAt field to given value.

### HasExpiresAt

`func (o *APIKeyMetadata) HasExpiresAt() bool`

HasExpiresAt returns a boolean if a field has been set.

### GetId

`func (o *APIKeyMetadata) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *APIKeyMetadata) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *APIKeyMetadata) SetId(v string)`

SetId sets Id field to given value.


### GetLastUsedAt

`func (o *APIKeyMetadata) GetLastUsedAt() time.Time`

GetLastUsedAt returns the LastUsedAt field if non-nil, zero value otherwise.

### GetLastUsedAtOk

`func (o *APIKeyMetadata) GetLastUsedAtOk() (*time.Time, bool)`

GetLastUsedAtOk returns a tuple with the LastUsedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLastUsedAt

`func (o *APIKeyMetadata) SetLastUsedAt(v time.Time)`

SetLastUsedAt sets LastUsedAt field to given value.

### HasLastUsedAt

`func (o *APIKeyMetadata) HasLastUsedAt() bool`

HasLastUsedAt returns a boolean if a field has been set.

### GetName

`func (o *APIKeyMetadata) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *APIKeyMetadata) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *APIKeyMetadata) SetName(v string)`

SetName sets Name field to given value.


### GetPrefix

`func (o *APIKeyMetadata) GetPrefix() string`

GetPrefix returns the Prefix field if non-nil, zero value otherwise.

### GetPrefixOk

`func (o *APIKeyMetadata) GetPrefixOk() (*string, bool)`

GetPrefixOk returns a tuple with the Prefix field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPrefix

`func (o *APIKeyMetadata) SetPrefix(v string)`

SetPrefix sets Prefix field to given value.


### GetRevokedAt

`func (o *APIKeyMetadata) GetRevokedAt() time.Time`

GetRevokedAt returns the RevokedAt field if non-nil, zero value otherwise.

### GetRevokedAtOk

`func (o *APIKeyMetadata) GetRevokedAtOk() (*time.Time, bool)`

GetRevokedAtOk returns a tuple with the RevokedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRevokedAt

`func (o *APIKeyMetadata) SetRevokedAt(v time.Time)`

SetRevokedAt sets RevokedAt field to given value.

### HasRevokedAt

`func (o *APIKeyMetadata) HasRevokedAt() bool`

HasRevokedAt returns a boolean if a field has been set.

### GetRevokedByUserId

`func (o *APIKeyMetadata) GetRevokedByUserId() string`

GetRevokedByUserId returns the RevokedByUserId field if non-nil, zero value otherwise.

### GetRevokedByUserIdOk

`func (o *APIKeyMetadata) GetRevokedByUserIdOk() (*string, bool)`

GetRevokedByUserIdOk returns a tuple with the RevokedByUserId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRevokedByUserId

`func (o *APIKeyMetadata) SetRevokedByUserId(v string)`

SetRevokedByUserId sets RevokedByUserId field to given value.

### HasRevokedByUserId

`func (o *APIKeyMetadata) HasRevokedByUserId() bool`

HasRevokedByUserId returns a boolean if a field has been set.

### GetRoleType

`func (o *APIKeyMetadata) GetRoleType() string`

GetRoleType returns the RoleType field if non-nil, zero value otherwise.

### GetRoleTypeOk

`func (o *APIKeyMetadata) GetRoleTypeOk() (*string, bool)`

GetRoleTypeOk returns a tuple with the RoleType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRoleType

`func (o *APIKeyMetadata) SetRoleType(v string)`

SetRoleType sets RoleType field to given value.


### GetStatus

`func (o *APIKeyMetadata) GetStatus() string`

GetStatus returns the Status field if non-nil, zero value otherwise.

### GetStatusOk

`func (o *APIKeyMetadata) GetStatusOk() (*string, bool)`

GetStatusOk returns a tuple with the Status field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStatus

`func (o *APIKeyMetadata) SetStatus(v string)`

SetStatus sets Status field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


