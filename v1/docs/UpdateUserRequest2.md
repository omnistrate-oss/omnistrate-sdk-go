# UpdateUserRequest2

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Address** | Pointer to [**Address**](Address.md) |  | [optional] 
**Attributes** | Pointer to **map[string]string** | Additional attributes of the user. | [optional] 
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

## Methods

### NewUpdateUserRequest2

`func NewUpdateUserRequest2() *UpdateUserRequest2`

NewUpdateUserRequest2 instantiates a new UpdateUserRequest2 object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewUpdateUserRequest2WithDefaults

`func NewUpdateUserRequest2WithDefaults() *UpdateUserRequest2`

NewUpdateUserRequest2WithDefaults instantiates a new UpdateUserRequest2 object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetAddress

`func (o *UpdateUserRequest2) GetAddress() Address`

GetAddress returns the Address field if non-nil, zero value otherwise.

### GetAddressOk

`func (o *UpdateUserRequest2) GetAddressOk() (*Address, bool)`

GetAddressOk returns a tuple with the Address field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAddress

`func (o *UpdateUserRequest2) SetAddress(v Address)`

SetAddress sets Address field to given value.

### HasAddress

`func (o *UpdateUserRequest2) HasAddress() bool`

HasAddress returns a boolean if a field has been set.

### GetAttributes

`func (o *UpdateUserRequest2) GetAttributes() map[string]string`

GetAttributes returns the Attributes field if non-nil, zero value otherwise.

### GetAttributesOk

`func (o *UpdateUserRequest2) GetAttributesOk() (*map[string]string, bool)`

GetAttributesOk returns a tuple with the Attributes field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAttributes

`func (o *UpdateUserRequest2) SetAttributes(v map[string]string)`

SetAttributes sets Attributes field to given value.

### HasAttributes

`func (o *UpdateUserRequest2) HasAttributes() bool`

HasAttributes returns a boolean if a field has been set.

### GetName

`func (o *UpdateUserRequest2) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *UpdateUserRequest2) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *UpdateUserRequest2) SetName(v string)`

SetName sets Name field to given value.

### HasName

`func (o *UpdateUserRequest2) HasName() bool`

HasName returns a boolean if a field has been set.

### GetOrgCookiePolicy

`func (o *UpdateUserRequest2) GetOrgCookiePolicy() string`

GetOrgCookiePolicy returns the OrgCookiePolicy field if non-nil, zero value otherwise.

### GetOrgCookiePolicyOk

`func (o *UpdateUserRequest2) GetOrgCookiePolicyOk() (*string, bool)`

GetOrgCookiePolicyOk returns a tuple with the OrgCookiePolicy field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOrgCookiePolicy

`func (o *UpdateUserRequest2) SetOrgCookiePolicy(v string)`

SetOrgCookiePolicy sets OrgCookiePolicy field to given value.

### HasOrgCookiePolicy

`func (o *UpdateUserRequest2) HasOrgCookiePolicy() bool`

HasOrgCookiePolicy returns a boolean if a field has been set.

### GetOrgDescription

`func (o *UpdateUserRequest2) GetOrgDescription() string`

GetOrgDescription returns the OrgDescription field if non-nil, zero value otherwise.

### GetOrgDescriptionOk

`func (o *UpdateUserRequest2) GetOrgDescriptionOk() (*string, bool)`

GetOrgDescriptionOk returns a tuple with the OrgDescription field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOrgDescription

`func (o *UpdateUserRequest2) SetOrgDescription(v string)`

SetOrgDescription sets OrgDescription field to given value.

### HasOrgDescription

`func (o *UpdateUserRequest2) HasOrgDescription() bool`

HasOrgDescription returns a boolean if a field has been set.

### GetOrgFavIconURL

`func (o *UpdateUserRequest2) GetOrgFavIconURL() string`

GetOrgFavIconURL returns the OrgFavIconURL field if non-nil, zero value otherwise.

### GetOrgFavIconURLOk

`func (o *UpdateUserRequest2) GetOrgFavIconURLOk() (*string, bool)`

GetOrgFavIconURLOk returns a tuple with the OrgFavIconURL field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOrgFavIconURL

`func (o *UpdateUserRequest2) SetOrgFavIconURL(v string)`

SetOrgFavIconURL sets OrgFavIconURL field to given value.

### HasOrgFavIconURL

`func (o *UpdateUserRequest2) HasOrgFavIconURL() bool`

HasOrgFavIconURL returns a boolean if a field has been set.

### GetOrgLogoURL

`func (o *UpdateUserRequest2) GetOrgLogoURL() string`

GetOrgLogoURL returns the OrgLogoURL field if non-nil, zero value otherwise.

### GetOrgLogoURLOk

`func (o *UpdateUserRequest2) GetOrgLogoURLOk() (*string, bool)`

GetOrgLogoURLOk returns a tuple with the OrgLogoURL field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOrgLogoURL

`func (o *UpdateUserRequest2) SetOrgLogoURL(v string)`

SetOrgLogoURL sets OrgLogoURL field to given value.

### HasOrgLogoURL

`func (o *UpdateUserRequest2) HasOrgLogoURL() bool`

HasOrgLogoURL returns a boolean if a field has been set.

### GetOrgName

`func (o *UpdateUserRequest2) GetOrgName() string`

GetOrgName returns the OrgName field if non-nil, zero value otherwise.

### GetOrgNameOk

`func (o *UpdateUserRequest2) GetOrgNameOk() (*string, bool)`

GetOrgNameOk returns a tuple with the OrgName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOrgName

`func (o *UpdateUserRequest2) SetOrgName(v string)`

SetOrgName sets OrgName field to given value.

### HasOrgName

`func (o *UpdateUserRequest2) HasOrgName() bool`

HasOrgName returns a boolean if a field has been set.

### GetOrgPrivacyPolicy

`func (o *UpdateUserRequest2) GetOrgPrivacyPolicy() string`

GetOrgPrivacyPolicy returns the OrgPrivacyPolicy field if non-nil, zero value otherwise.

### GetOrgPrivacyPolicyOk

`func (o *UpdateUserRequest2) GetOrgPrivacyPolicyOk() (*string, bool)`

GetOrgPrivacyPolicyOk returns a tuple with the OrgPrivacyPolicy field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOrgPrivacyPolicy

`func (o *UpdateUserRequest2) SetOrgPrivacyPolicy(v string)`

SetOrgPrivacyPolicy sets OrgPrivacyPolicy field to given value.

### HasOrgPrivacyPolicy

`func (o *UpdateUserRequest2) HasOrgPrivacyPolicy() bool`

HasOrgPrivacyPolicy returns a boolean if a field has been set.

### GetOrgSupportEmail

`func (o *UpdateUserRequest2) GetOrgSupportEmail() string`

GetOrgSupportEmail returns the OrgSupportEmail field if non-nil, zero value otherwise.

### GetOrgSupportEmailOk

`func (o *UpdateUserRequest2) GetOrgSupportEmailOk() (*string, bool)`

GetOrgSupportEmailOk returns a tuple with the OrgSupportEmail field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOrgSupportEmail

`func (o *UpdateUserRequest2) SetOrgSupportEmail(v string)`

SetOrgSupportEmail sets OrgSupportEmail field to given value.

### HasOrgSupportEmail

`func (o *UpdateUserRequest2) HasOrgSupportEmail() bool`

HasOrgSupportEmail returns a boolean if a field has been set.

### GetOrgTermsOfUse

`func (o *UpdateUserRequest2) GetOrgTermsOfUse() string`

GetOrgTermsOfUse returns the OrgTermsOfUse field if non-nil, zero value otherwise.

### GetOrgTermsOfUseOk

`func (o *UpdateUserRequest2) GetOrgTermsOfUseOk() (*string, bool)`

GetOrgTermsOfUseOk returns a tuple with the OrgTermsOfUse field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOrgTermsOfUse

`func (o *UpdateUserRequest2) SetOrgTermsOfUse(v string)`

SetOrgTermsOfUse sets OrgTermsOfUse field to given value.

### HasOrgTermsOfUse

`func (o *UpdateUserRequest2) HasOrgTermsOfUse() bool`

HasOrgTermsOfUse returns a boolean if a field has been set.

### GetOrgURL

`func (o *UpdateUserRequest2) GetOrgURL() string`

GetOrgURL returns the OrgURL field if non-nil, zero value otherwise.

### GetOrgURLOk

`func (o *UpdateUserRequest2) GetOrgURLOk() (*string, bool)`

GetOrgURLOk returns a tuple with the OrgURL field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOrgURL

`func (o *UpdateUserRequest2) SetOrgURL(v string)`

SetOrgURL sets OrgURL field to given value.

### HasOrgURL

`func (o *UpdateUserRequest2) HasOrgURL() bool`

HasOrgURL returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


