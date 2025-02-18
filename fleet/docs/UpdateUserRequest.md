# UpdateUserRequest

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Address** | Pointer to [**Address**](Address.md) |  | [optional] 
**Id** | **string** | The User ID | 
**Name** | Pointer to **string** | The name of the user | [optional] 
**OrgCookiePolicy** | Pointer to **string** | The cookie policy for the org that this user owns in an HTML format | [optional] 
**OrgDescription** | Pointer to **string** | The description of the org that this user owns | [optional] 
**OrgFavIconURL** | Pointer to **string** | The favicon of the org that this user owns | [optional] 
**OrgLogoURL** | Pointer to **string** | The logo of the org that this user owns | [optional] 
**OrgName** | Pointer to **string** | The org name that this user owns | [optional] 
**OrgPrivacyPolicy** | Pointer to **string** | The privacy policy for the org that this user owns in an HTML format | [optional] 
**OrgSupportEmail** | Pointer to **string** | The support email of the org that this user owns | [optional] 
**OrgTermsOfUse** | Pointer to **string** | The terms of use for the org that this user owns in an HTML format | [optional] 
**OrgURL** | Pointer to **string** | The url of the org that this user owns | [optional] 
**Token** | **string** | JWT token used to perform authorization | 

## Methods

### NewUpdateUserRequest

`func NewUpdateUserRequest(id string, token string, ) *UpdateUserRequest`

NewUpdateUserRequest instantiates a new UpdateUserRequest object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewUpdateUserRequestWithDefaults

`func NewUpdateUserRequestWithDefaults() *UpdateUserRequest`

NewUpdateUserRequestWithDefaults instantiates a new UpdateUserRequest object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetAddress

`func (o *UpdateUserRequest) GetAddress() Address`

GetAddress returns the Address field if non-nil, zero value otherwise.

### GetAddressOk

`func (o *UpdateUserRequest) GetAddressOk() (*Address, bool)`

GetAddressOk returns a tuple with the Address field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAddress

`func (o *UpdateUserRequest) SetAddress(v Address)`

SetAddress sets Address field to given value.

### HasAddress

`func (o *UpdateUserRequest) HasAddress() bool`

HasAddress returns a boolean if a field has been set.

### GetId

`func (o *UpdateUserRequest) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *UpdateUserRequest) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *UpdateUserRequest) SetId(v string)`

SetId sets Id field to given value.


### GetName

`func (o *UpdateUserRequest) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *UpdateUserRequest) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *UpdateUserRequest) SetName(v string)`

SetName sets Name field to given value.

### HasName

`func (o *UpdateUserRequest) HasName() bool`

HasName returns a boolean if a field has been set.

### GetOrgCookiePolicy

`func (o *UpdateUserRequest) GetOrgCookiePolicy() string`

GetOrgCookiePolicy returns the OrgCookiePolicy field if non-nil, zero value otherwise.

### GetOrgCookiePolicyOk

`func (o *UpdateUserRequest) GetOrgCookiePolicyOk() (*string, bool)`

GetOrgCookiePolicyOk returns a tuple with the OrgCookiePolicy field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOrgCookiePolicy

`func (o *UpdateUserRequest) SetOrgCookiePolicy(v string)`

SetOrgCookiePolicy sets OrgCookiePolicy field to given value.

### HasOrgCookiePolicy

`func (o *UpdateUserRequest) HasOrgCookiePolicy() bool`

HasOrgCookiePolicy returns a boolean if a field has been set.

### GetOrgDescription

`func (o *UpdateUserRequest) GetOrgDescription() string`

GetOrgDescription returns the OrgDescription field if non-nil, zero value otherwise.

### GetOrgDescriptionOk

`func (o *UpdateUserRequest) GetOrgDescriptionOk() (*string, bool)`

GetOrgDescriptionOk returns a tuple with the OrgDescription field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOrgDescription

`func (o *UpdateUserRequest) SetOrgDescription(v string)`

SetOrgDescription sets OrgDescription field to given value.

### HasOrgDescription

`func (o *UpdateUserRequest) HasOrgDescription() bool`

HasOrgDescription returns a boolean if a field has been set.

### GetOrgFavIconURL

`func (o *UpdateUserRequest) GetOrgFavIconURL() string`

GetOrgFavIconURL returns the OrgFavIconURL field if non-nil, zero value otherwise.

### GetOrgFavIconURLOk

`func (o *UpdateUserRequest) GetOrgFavIconURLOk() (*string, bool)`

GetOrgFavIconURLOk returns a tuple with the OrgFavIconURL field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOrgFavIconURL

`func (o *UpdateUserRequest) SetOrgFavIconURL(v string)`

SetOrgFavIconURL sets OrgFavIconURL field to given value.

### HasOrgFavIconURL

`func (o *UpdateUserRequest) HasOrgFavIconURL() bool`

HasOrgFavIconURL returns a boolean if a field has been set.

### GetOrgLogoURL

`func (o *UpdateUserRequest) GetOrgLogoURL() string`

GetOrgLogoURL returns the OrgLogoURL field if non-nil, zero value otherwise.

### GetOrgLogoURLOk

`func (o *UpdateUserRequest) GetOrgLogoURLOk() (*string, bool)`

GetOrgLogoURLOk returns a tuple with the OrgLogoURL field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOrgLogoURL

`func (o *UpdateUserRequest) SetOrgLogoURL(v string)`

SetOrgLogoURL sets OrgLogoURL field to given value.

### HasOrgLogoURL

`func (o *UpdateUserRequest) HasOrgLogoURL() bool`

HasOrgLogoURL returns a boolean if a field has been set.

### GetOrgName

`func (o *UpdateUserRequest) GetOrgName() string`

GetOrgName returns the OrgName field if non-nil, zero value otherwise.

### GetOrgNameOk

`func (o *UpdateUserRequest) GetOrgNameOk() (*string, bool)`

GetOrgNameOk returns a tuple with the OrgName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOrgName

`func (o *UpdateUserRequest) SetOrgName(v string)`

SetOrgName sets OrgName field to given value.

### HasOrgName

`func (o *UpdateUserRequest) HasOrgName() bool`

HasOrgName returns a boolean if a field has been set.

### GetOrgPrivacyPolicy

`func (o *UpdateUserRequest) GetOrgPrivacyPolicy() string`

GetOrgPrivacyPolicy returns the OrgPrivacyPolicy field if non-nil, zero value otherwise.

### GetOrgPrivacyPolicyOk

`func (o *UpdateUserRequest) GetOrgPrivacyPolicyOk() (*string, bool)`

GetOrgPrivacyPolicyOk returns a tuple with the OrgPrivacyPolicy field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOrgPrivacyPolicy

`func (o *UpdateUserRequest) SetOrgPrivacyPolicy(v string)`

SetOrgPrivacyPolicy sets OrgPrivacyPolicy field to given value.

### HasOrgPrivacyPolicy

`func (o *UpdateUserRequest) HasOrgPrivacyPolicy() bool`

HasOrgPrivacyPolicy returns a boolean if a field has been set.

### GetOrgSupportEmail

`func (o *UpdateUserRequest) GetOrgSupportEmail() string`

GetOrgSupportEmail returns the OrgSupportEmail field if non-nil, zero value otherwise.

### GetOrgSupportEmailOk

`func (o *UpdateUserRequest) GetOrgSupportEmailOk() (*string, bool)`

GetOrgSupportEmailOk returns a tuple with the OrgSupportEmail field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOrgSupportEmail

`func (o *UpdateUserRequest) SetOrgSupportEmail(v string)`

SetOrgSupportEmail sets OrgSupportEmail field to given value.

### HasOrgSupportEmail

`func (o *UpdateUserRequest) HasOrgSupportEmail() bool`

HasOrgSupportEmail returns a boolean if a field has been set.

### GetOrgTermsOfUse

`func (o *UpdateUserRequest) GetOrgTermsOfUse() string`

GetOrgTermsOfUse returns the OrgTermsOfUse field if non-nil, zero value otherwise.

### GetOrgTermsOfUseOk

`func (o *UpdateUserRequest) GetOrgTermsOfUseOk() (*string, bool)`

GetOrgTermsOfUseOk returns a tuple with the OrgTermsOfUse field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOrgTermsOfUse

`func (o *UpdateUserRequest) SetOrgTermsOfUse(v string)`

SetOrgTermsOfUse sets OrgTermsOfUse field to given value.

### HasOrgTermsOfUse

`func (o *UpdateUserRequest) HasOrgTermsOfUse() bool`

HasOrgTermsOfUse returns a boolean if a field has been set.

### GetOrgURL

`func (o *UpdateUserRequest) GetOrgURL() string`

GetOrgURL returns the OrgURL field if non-nil, zero value otherwise.

### GetOrgURLOk

`func (o *UpdateUserRequest) GetOrgURLOk() (*string, bool)`

GetOrgURLOk returns a tuple with the OrgURL field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOrgURL

`func (o *UpdateUserRequest) SetOrgURL(v string)`

SetOrgURL sets OrgURL field to given value.

### HasOrgURL

`func (o *UpdateUserRequest) HasOrgURL() bool`

HasOrgURL returns a boolean if a field has been set.

### GetToken

`func (o *UpdateUserRequest) GetToken() string`

GetToken returns the Token field if non-nil, zero value otherwise.

### GetTokenOk

`func (o *UpdateUserRequest) GetTokenOk() (*string, bool)`

GetTokenOk returns a tuple with the Token field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetToken

`func (o *UpdateUserRequest) SetToken(v string)`

SetToken sets Token field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


