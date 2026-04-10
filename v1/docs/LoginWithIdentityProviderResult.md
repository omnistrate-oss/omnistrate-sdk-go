# LoginWithIdentityProviderResult

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**JwtToken** | **string** | The jwt token | 
**RefreshToken** | Pointer to **string** | Opaque refresh token for obtaining new JWT tokens | [optional] 

## Methods

### NewLoginWithIdentityProviderResult

`func NewLoginWithIdentityProviderResult(jwtToken string, ) *LoginWithIdentityProviderResult`

NewLoginWithIdentityProviderResult instantiates a new LoginWithIdentityProviderResult object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewLoginWithIdentityProviderResultWithDefaults

`func NewLoginWithIdentityProviderResultWithDefaults() *LoginWithIdentityProviderResult`

NewLoginWithIdentityProviderResultWithDefaults instantiates a new LoginWithIdentityProviderResult object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetJwtToken

`func (o *LoginWithIdentityProviderResult) GetJwtToken() string`

GetJwtToken returns the JwtToken field if non-nil, zero value otherwise.

### GetJwtTokenOk

`func (o *LoginWithIdentityProviderResult) GetJwtTokenOk() (*string, bool)`

GetJwtTokenOk returns a tuple with the JwtToken field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetJwtToken

`func (o *LoginWithIdentityProviderResult) SetJwtToken(v string)`

SetJwtToken sets JwtToken field to given value.


### GetRefreshToken

`func (o *LoginWithIdentityProviderResult) GetRefreshToken() string`

GetRefreshToken returns the RefreshToken field if non-nil, zero value otherwise.

### GetRefreshTokenOk

`func (o *LoginWithIdentityProviderResult) GetRefreshTokenOk() (*string, bool)`

GetRefreshTokenOk returns a tuple with the RefreshToken field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRefreshToken

`func (o *LoginWithIdentityProviderResult) SetRefreshToken(v string)`

SetRefreshToken sets RefreshToken field to given value.

### HasRefreshToken

`func (o *LoginWithIdentityProviderResult) HasRefreshToken() bool`

HasRefreshToken returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


