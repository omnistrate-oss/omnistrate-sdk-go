# UpdateUserRequestBody

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Address** | Pointer to [**Address**](Address.md) |  | [optional] 
**Name** | Pointer to **string** | The name of the user | [optional] 
**OrgDescription** | Pointer to **string** | The description of the org that this user owns | [optional] 
**OrgFavIconURL** | Pointer to **string** | The favicon of the org that this user owns | [optional] 
**OrgLogoURL** | Pointer to **string** | The logo of the org that this user owns | [optional] 
**OrgName** | Pointer to **string** | The org name that this user owns | [optional] 
**OrgPrivacyPolicy** | Pointer to **string** | The privacy policy for the org that this user owns in an HTML format | [optional] 
**OrgSupportEmail** | Pointer to **string** | The support email of the org that this user owns | [optional] 
**OrgTermsOfUse** | Pointer to **string** | The terms of use for the org that this user owns in an HTML format | [optional] 
**OrgURL** | Pointer to **string** | The url of the org that this user owns | [optional] 

## Methods

### NewUpdateUserRequestBody

`func NewUpdateUserRequestBody() *UpdateUserRequestBody`

NewUpdateUserRequestBody instantiates a new UpdateUserRequestBody object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewUpdateUserRequestBodyWithDefaults

`func NewUpdateUserRequestBodyWithDefaults() *UpdateUserRequestBody`

NewUpdateUserRequestBodyWithDefaults instantiates a new UpdateUserRequestBody object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetAddress

`func (o *UpdateUserRequestBody) GetAddress() Address`

GetAddress returns the Address field if non-nil, zero value otherwise.

### GetAddressOk

`func (o *UpdateUserRequestBody) GetAddressOk() (*Address, bool)`

GetAddressOk returns a tuple with the Address field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAddress

`func (o *UpdateUserRequestBody) SetAddress(v Address)`

SetAddress sets Address field to given value.

### HasAddress

`func (o *UpdateUserRequestBody) HasAddress() bool`

HasAddress returns a boolean if a field has been set.

### GetName

`func (o *UpdateUserRequestBody) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *UpdateUserRequestBody) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *UpdateUserRequestBody) SetName(v string)`

SetName sets Name field to given value.

### HasName

`func (o *UpdateUserRequestBody) HasName() bool`

HasName returns a boolean if a field has been set.

### GetOrgDescription

`func (o *UpdateUserRequestBody) GetOrgDescription() string`

GetOrgDescription returns the OrgDescription field if non-nil, zero value otherwise.

### GetOrgDescriptionOk

`func (o *UpdateUserRequestBody) GetOrgDescriptionOk() (*string, bool)`

GetOrgDescriptionOk returns a tuple with the OrgDescription field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOrgDescription

`func (o *UpdateUserRequestBody) SetOrgDescription(v string)`

SetOrgDescription sets OrgDescription field to given value.

### HasOrgDescription

`func (o *UpdateUserRequestBody) HasOrgDescription() bool`

HasOrgDescription returns a boolean if a field has been set.

### GetOrgFavIconURL

`func (o *UpdateUserRequestBody) GetOrgFavIconURL() string`

GetOrgFavIconURL returns the OrgFavIconURL field if non-nil, zero value otherwise.

### GetOrgFavIconURLOk

`func (o *UpdateUserRequestBody) GetOrgFavIconURLOk() (*string, bool)`

GetOrgFavIconURLOk returns a tuple with the OrgFavIconURL field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOrgFavIconURL

`func (o *UpdateUserRequestBody) SetOrgFavIconURL(v string)`

SetOrgFavIconURL sets OrgFavIconURL field to given value.

### HasOrgFavIconURL

`func (o *UpdateUserRequestBody) HasOrgFavIconURL() bool`

HasOrgFavIconURL returns a boolean if a field has been set.

### GetOrgLogoURL

`func (o *UpdateUserRequestBody) GetOrgLogoURL() string`

GetOrgLogoURL returns the OrgLogoURL field if non-nil, zero value otherwise.

### GetOrgLogoURLOk

`func (o *UpdateUserRequestBody) GetOrgLogoURLOk() (*string, bool)`

GetOrgLogoURLOk returns a tuple with the OrgLogoURL field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOrgLogoURL

`func (o *UpdateUserRequestBody) SetOrgLogoURL(v string)`

SetOrgLogoURL sets OrgLogoURL field to given value.

### HasOrgLogoURL

`func (o *UpdateUserRequestBody) HasOrgLogoURL() bool`

HasOrgLogoURL returns a boolean if a field has been set.

### GetOrgName

`func (o *UpdateUserRequestBody) GetOrgName() string`

GetOrgName returns the OrgName field if non-nil, zero value otherwise.

### GetOrgNameOk

`func (o *UpdateUserRequestBody) GetOrgNameOk() (*string, bool)`

GetOrgNameOk returns a tuple with the OrgName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOrgName

`func (o *UpdateUserRequestBody) SetOrgName(v string)`

SetOrgName sets OrgName field to given value.

### HasOrgName

`func (o *UpdateUserRequestBody) HasOrgName() bool`

HasOrgName returns a boolean if a field has been set.

### GetOrgPrivacyPolicy

`func (o *UpdateUserRequestBody) GetOrgPrivacyPolicy() string`

GetOrgPrivacyPolicy returns the OrgPrivacyPolicy field if non-nil, zero value otherwise.

### GetOrgPrivacyPolicyOk

`func (o *UpdateUserRequestBody) GetOrgPrivacyPolicyOk() (*string, bool)`

GetOrgPrivacyPolicyOk returns a tuple with the OrgPrivacyPolicy field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOrgPrivacyPolicy

`func (o *UpdateUserRequestBody) SetOrgPrivacyPolicy(v string)`

SetOrgPrivacyPolicy sets OrgPrivacyPolicy field to given value.

### HasOrgPrivacyPolicy

`func (o *UpdateUserRequestBody) HasOrgPrivacyPolicy() bool`

HasOrgPrivacyPolicy returns a boolean if a field has been set.

### GetOrgSupportEmail

`func (o *UpdateUserRequestBody) GetOrgSupportEmail() string`

GetOrgSupportEmail returns the OrgSupportEmail field if non-nil, zero value otherwise.

### GetOrgSupportEmailOk

`func (o *UpdateUserRequestBody) GetOrgSupportEmailOk() (*string, bool)`

GetOrgSupportEmailOk returns a tuple with the OrgSupportEmail field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOrgSupportEmail

`func (o *UpdateUserRequestBody) SetOrgSupportEmail(v string)`

SetOrgSupportEmail sets OrgSupportEmail field to given value.

### HasOrgSupportEmail

`func (o *UpdateUserRequestBody) HasOrgSupportEmail() bool`

HasOrgSupportEmail returns a boolean if a field has been set.

### GetOrgTermsOfUse

`func (o *UpdateUserRequestBody) GetOrgTermsOfUse() string`

GetOrgTermsOfUse returns the OrgTermsOfUse field if non-nil, zero value otherwise.

### GetOrgTermsOfUseOk

`func (o *UpdateUserRequestBody) GetOrgTermsOfUseOk() (*string, bool)`

GetOrgTermsOfUseOk returns a tuple with the OrgTermsOfUse field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOrgTermsOfUse

`func (o *UpdateUserRequestBody) SetOrgTermsOfUse(v string)`

SetOrgTermsOfUse sets OrgTermsOfUse field to given value.

### HasOrgTermsOfUse

`func (o *UpdateUserRequestBody) HasOrgTermsOfUse() bool`

HasOrgTermsOfUse returns a boolean if a field has been set.

### GetOrgURL

`func (o *UpdateUserRequestBody) GetOrgURL() string`

GetOrgURL returns the OrgURL field if non-nil, zero value otherwise.

### GetOrgURLOk

`func (o *UpdateUserRequestBody) GetOrgURLOk() (*string, bool)`

GetOrgURLOk returns a tuple with the OrgURL field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOrgURL

`func (o *UpdateUserRequestBody) SetOrgURL(v string)`

SetOrgURL sets OrgURL field to given value.

### HasOrgURL

`func (o *UpdateUserRequestBody) HasOrgURL() bool`

HasOrgURL returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


