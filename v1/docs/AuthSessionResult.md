# AuthSessionResult

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Exp** | **int64** | JWT expiration as a Unix timestamp in seconds | 
**OrganizationID** | Pointer to **string** | The verified organization scope, when present in the JWT | [optional] 
**ServiceProviderID** | Pointer to **string** | The verified service provider scope, when present in the JWT | [optional] 
**UserID** | **string** | The authenticated user ID | 

## Methods

### NewAuthSessionResult

`func NewAuthSessionResult(exp int64, userID string, ) *AuthSessionResult`

NewAuthSessionResult instantiates a new AuthSessionResult object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewAuthSessionResultWithDefaults

`func NewAuthSessionResultWithDefaults() *AuthSessionResult`

NewAuthSessionResultWithDefaults instantiates a new AuthSessionResult object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetExp

`func (o *AuthSessionResult) GetExp() int64`

GetExp returns the Exp field if non-nil, zero value otherwise.

### GetExpOk

`func (o *AuthSessionResult) GetExpOk() (*int64, bool)`

GetExpOk returns a tuple with the Exp field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetExp

`func (o *AuthSessionResult) SetExp(v int64)`

SetExp sets Exp field to given value.


### GetOrganizationID

`func (o *AuthSessionResult) GetOrganizationID() string`

GetOrganizationID returns the OrganizationID field if non-nil, zero value otherwise.

### GetOrganizationIDOk

`func (o *AuthSessionResult) GetOrganizationIDOk() (*string, bool)`

GetOrganizationIDOk returns a tuple with the OrganizationID field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOrganizationID

`func (o *AuthSessionResult) SetOrganizationID(v string)`

SetOrganizationID sets OrganizationID field to given value.

### HasOrganizationID

`func (o *AuthSessionResult) HasOrganizationID() bool`

HasOrganizationID returns a boolean if a field has been set.

### GetServiceProviderID

`func (o *AuthSessionResult) GetServiceProviderID() string`

GetServiceProviderID returns the ServiceProviderID field if non-nil, zero value otherwise.

### GetServiceProviderIDOk

`func (o *AuthSessionResult) GetServiceProviderIDOk() (*string, bool)`

GetServiceProviderIDOk returns a tuple with the ServiceProviderID field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetServiceProviderID

`func (o *AuthSessionResult) SetServiceProviderID(v string)`

SetServiceProviderID sets ServiceProviderID field to given value.

### HasServiceProviderID

`func (o *AuthSessionResult) HasServiceProviderID() bool`

HasServiceProviderID returns a boolean if a field has been set.

### GetUserID

`func (o *AuthSessionResult) GetUserID() string`

GetUserID returns the UserID field if non-nil, zero value otherwise.

### GetUserIDOk

`func (o *AuthSessionResult) GetUserIDOk() (*string, bool)`

GetUserIDOk returns a tuple with the UserID field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUserID

`func (o *AuthSessionResult) SetUserID(v string)`

SetUserID sets UserID field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


