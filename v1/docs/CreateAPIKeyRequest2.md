# CreateAPIKeyRequest2

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Description** | Pointer to **string** | Optional free-text description. No character of the plaintext key is ever stored or returned beyond the prefix label. | [optional] 
**ExpiresAt** | Pointer to **time.Time** | Optional ISO 8601 expiry. Omit for non-expiring keys. | [optional] 
**Name** | **string** | Operator-facing display name for the key. Unique within the org; conflicts return 409. | 
**RoleType** | **string** | Authority granted to the key. Must be assignable by the caller per the role priority cap. The &#39;root&#39; role is never assignable to an API key. | 

## Methods

### NewCreateAPIKeyRequest2

`func NewCreateAPIKeyRequest2(name string, roleType string, ) *CreateAPIKeyRequest2`

NewCreateAPIKeyRequest2 instantiates a new CreateAPIKeyRequest2 object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewCreateAPIKeyRequest2WithDefaults

`func NewCreateAPIKeyRequest2WithDefaults() *CreateAPIKeyRequest2`

NewCreateAPIKeyRequest2WithDefaults instantiates a new CreateAPIKeyRequest2 object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetDescription

`func (o *CreateAPIKeyRequest2) GetDescription() string`

GetDescription returns the Description field if non-nil, zero value otherwise.

### GetDescriptionOk

`func (o *CreateAPIKeyRequest2) GetDescriptionOk() (*string, bool)`

GetDescriptionOk returns a tuple with the Description field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDescription

`func (o *CreateAPIKeyRequest2) SetDescription(v string)`

SetDescription sets Description field to given value.

### HasDescription

`func (o *CreateAPIKeyRequest2) HasDescription() bool`

HasDescription returns a boolean if a field has been set.

### GetExpiresAt

`func (o *CreateAPIKeyRequest2) GetExpiresAt() time.Time`

GetExpiresAt returns the ExpiresAt field if non-nil, zero value otherwise.

### GetExpiresAtOk

`func (o *CreateAPIKeyRequest2) GetExpiresAtOk() (*time.Time, bool)`

GetExpiresAtOk returns a tuple with the ExpiresAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetExpiresAt

`func (o *CreateAPIKeyRequest2) SetExpiresAt(v time.Time)`

SetExpiresAt sets ExpiresAt field to given value.

### HasExpiresAt

`func (o *CreateAPIKeyRequest2) HasExpiresAt() bool`

HasExpiresAt returns a boolean if a field has been set.

### GetName

`func (o *CreateAPIKeyRequest2) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *CreateAPIKeyRequest2) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *CreateAPIKeyRequest2) SetName(v string)`

SetName sets Name field to given value.


### GetRoleType

`func (o *CreateAPIKeyRequest2) GetRoleType() string`

GetRoleType returns the RoleType field if non-nil, zero value otherwise.

### GetRoleTypeOk

`func (o *CreateAPIKeyRequest2) GetRoleTypeOk() (*string, bool)`

GetRoleTypeOk returns a tuple with the RoleType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRoleType

`func (o *CreateAPIKeyRequest2) SetRoleType(v string)`

SetRoleType sets RoleType field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


