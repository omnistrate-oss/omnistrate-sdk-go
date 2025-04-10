# DescribeUserResult

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Address** | Pointer to [**Address**](Address.md) |  | [optional] 
**CreatedAt** | Pointer to **string** | The user creation time | [optional] 
**Email** | Pointer to **string** | The email of the user | [optional] 
**Enabled** | Pointer to **bool** | Is the user enabled. | [optional] 
**Id** | **string** | The User ID | 
**LastModifiedAt** | Pointer to **string** | The user update time | [optional] 
**Name** | Pointer to **string** | The name of the user | [optional] 
**OrgCookiePolicy** | Pointer to **string** | The cookie policy for the org that this user owns | [optional] 
**OrgDescription** | Pointer to **string** | The description of the org that this user owns | [optional] 
**OrgFavIconURL** | Pointer to **string** | The favicon of the org that this user owns | [optional] 
**OrgId** | Pointer to **string** | The ID of the org that this user owns | [optional] 
**OrgLogoURL** | Pointer to **string** | The logo of the org that this user owns | [optional] 
**OrgName** | Pointer to **string** | The org name that this user owns | [optional] 
**OrgPrivacyPolicy** | Pointer to **string** | The privacy policy for the org that this user owns | [optional] 
**OrgSupportEmail** | Pointer to **string** | The support email of the org that this user owns | [optional] 
**OrgTermsOfUse** | Pointer to **string** | The terms of use for the org that this user owns | [optional] 
**OrgURL** | Pointer to **string** | The url of the org that this user owns | [optional] 
**PlanName** | Pointer to **string** | This parameter is used to select the appropriate Product Plan | [optional] 
**RoleType** | Pointer to **string** | Type of the role | [optional] 
**Status** | Pointer to **string** | The status of the user. | [optional] 

## Methods

### NewDescribeUserResult

`func NewDescribeUserResult(id string, ) *DescribeUserResult`

NewDescribeUserResult instantiates a new DescribeUserResult object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewDescribeUserResultWithDefaults

`func NewDescribeUserResultWithDefaults() *DescribeUserResult`

NewDescribeUserResultWithDefaults instantiates a new DescribeUserResult object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetAddress

`func (o *DescribeUserResult) GetAddress() Address`

GetAddress returns the Address field if non-nil, zero value otherwise.

### GetAddressOk

`func (o *DescribeUserResult) GetAddressOk() (*Address, bool)`

GetAddressOk returns a tuple with the Address field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAddress

`func (o *DescribeUserResult) SetAddress(v Address)`

SetAddress sets Address field to given value.

### HasAddress

`func (o *DescribeUserResult) HasAddress() bool`

HasAddress returns a boolean if a field has been set.

### GetCreatedAt

`func (o *DescribeUserResult) GetCreatedAt() string`

GetCreatedAt returns the CreatedAt field if non-nil, zero value otherwise.

### GetCreatedAtOk

`func (o *DescribeUserResult) GetCreatedAtOk() (*string, bool)`

GetCreatedAtOk returns a tuple with the CreatedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCreatedAt

`func (o *DescribeUserResult) SetCreatedAt(v string)`

SetCreatedAt sets CreatedAt field to given value.

### HasCreatedAt

`func (o *DescribeUserResult) HasCreatedAt() bool`

HasCreatedAt returns a boolean if a field has been set.

### GetEmail

`func (o *DescribeUserResult) GetEmail() string`

GetEmail returns the Email field if non-nil, zero value otherwise.

### GetEmailOk

`func (o *DescribeUserResult) GetEmailOk() (*string, bool)`

GetEmailOk returns a tuple with the Email field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEmail

`func (o *DescribeUserResult) SetEmail(v string)`

SetEmail sets Email field to given value.

### HasEmail

`func (o *DescribeUserResult) HasEmail() bool`

HasEmail returns a boolean if a field has been set.

### GetEnabled

`func (o *DescribeUserResult) GetEnabled() bool`

GetEnabled returns the Enabled field if non-nil, zero value otherwise.

### GetEnabledOk

`func (o *DescribeUserResult) GetEnabledOk() (*bool, bool)`

GetEnabledOk returns a tuple with the Enabled field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEnabled

`func (o *DescribeUserResult) SetEnabled(v bool)`

SetEnabled sets Enabled field to given value.

### HasEnabled

`func (o *DescribeUserResult) HasEnabled() bool`

HasEnabled returns a boolean if a field has been set.

### GetId

`func (o *DescribeUserResult) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *DescribeUserResult) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *DescribeUserResult) SetId(v string)`

SetId sets Id field to given value.


### GetLastModifiedAt

`func (o *DescribeUserResult) GetLastModifiedAt() string`

GetLastModifiedAt returns the LastModifiedAt field if non-nil, zero value otherwise.

### GetLastModifiedAtOk

`func (o *DescribeUserResult) GetLastModifiedAtOk() (*string, bool)`

GetLastModifiedAtOk returns a tuple with the LastModifiedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLastModifiedAt

`func (o *DescribeUserResult) SetLastModifiedAt(v string)`

SetLastModifiedAt sets LastModifiedAt field to given value.

### HasLastModifiedAt

`func (o *DescribeUserResult) HasLastModifiedAt() bool`

HasLastModifiedAt returns a boolean if a field has been set.

### GetName

`func (o *DescribeUserResult) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *DescribeUserResult) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *DescribeUserResult) SetName(v string)`

SetName sets Name field to given value.

### HasName

`func (o *DescribeUserResult) HasName() bool`

HasName returns a boolean if a field has been set.

### GetOrgCookiePolicy

`func (o *DescribeUserResult) GetOrgCookiePolicy() string`

GetOrgCookiePolicy returns the OrgCookiePolicy field if non-nil, zero value otherwise.

### GetOrgCookiePolicyOk

`func (o *DescribeUserResult) GetOrgCookiePolicyOk() (*string, bool)`

GetOrgCookiePolicyOk returns a tuple with the OrgCookiePolicy field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOrgCookiePolicy

`func (o *DescribeUserResult) SetOrgCookiePolicy(v string)`

SetOrgCookiePolicy sets OrgCookiePolicy field to given value.

### HasOrgCookiePolicy

`func (o *DescribeUserResult) HasOrgCookiePolicy() bool`

HasOrgCookiePolicy returns a boolean if a field has been set.

### GetOrgDescription

`func (o *DescribeUserResult) GetOrgDescription() string`

GetOrgDescription returns the OrgDescription field if non-nil, zero value otherwise.

### GetOrgDescriptionOk

`func (o *DescribeUserResult) GetOrgDescriptionOk() (*string, bool)`

GetOrgDescriptionOk returns a tuple with the OrgDescription field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOrgDescription

`func (o *DescribeUserResult) SetOrgDescription(v string)`

SetOrgDescription sets OrgDescription field to given value.

### HasOrgDescription

`func (o *DescribeUserResult) HasOrgDescription() bool`

HasOrgDescription returns a boolean if a field has been set.

### GetOrgFavIconURL

`func (o *DescribeUserResult) GetOrgFavIconURL() string`

GetOrgFavIconURL returns the OrgFavIconURL field if non-nil, zero value otherwise.

### GetOrgFavIconURLOk

`func (o *DescribeUserResult) GetOrgFavIconURLOk() (*string, bool)`

GetOrgFavIconURLOk returns a tuple with the OrgFavIconURL field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOrgFavIconURL

`func (o *DescribeUserResult) SetOrgFavIconURL(v string)`

SetOrgFavIconURL sets OrgFavIconURL field to given value.

### HasOrgFavIconURL

`func (o *DescribeUserResult) HasOrgFavIconURL() bool`

HasOrgFavIconURL returns a boolean if a field has been set.

### GetOrgId

`func (o *DescribeUserResult) GetOrgId() string`

GetOrgId returns the OrgId field if non-nil, zero value otherwise.

### GetOrgIdOk

`func (o *DescribeUserResult) GetOrgIdOk() (*string, bool)`

GetOrgIdOk returns a tuple with the OrgId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOrgId

`func (o *DescribeUserResult) SetOrgId(v string)`

SetOrgId sets OrgId field to given value.

### HasOrgId

`func (o *DescribeUserResult) HasOrgId() bool`

HasOrgId returns a boolean if a field has been set.

### GetOrgLogoURL

`func (o *DescribeUserResult) GetOrgLogoURL() string`

GetOrgLogoURL returns the OrgLogoURL field if non-nil, zero value otherwise.

### GetOrgLogoURLOk

`func (o *DescribeUserResult) GetOrgLogoURLOk() (*string, bool)`

GetOrgLogoURLOk returns a tuple with the OrgLogoURL field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOrgLogoURL

`func (o *DescribeUserResult) SetOrgLogoURL(v string)`

SetOrgLogoURL sets OrgLogoURL field to given value.

### HasOrgLogoURL

`func (o *DescribeUserResult) HasOrgLogoURL() bool`

HasOrgLogoURL returns a boolean if a field has been set.

### GetOrgName

`func (o *DescribeUserResult) GetOrgName() string`

GetOrgName returns the OrgName field if non-nil, zero value otherwise.

### GetOrgNameOk

`func (o *DescribeUserResult) GetOrgNameOk() (*string, bool)`

GetOrgNameOk returns a tuple with the OrgName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOrgName

`func (o *DescribeUserResult) SetOrgName(v string)`

SetOrgName sets OrgName field to given value.

### HasOrgName

`func (o *DescribeUserResult) HasOrgName() bool`

HasOrgName returns a boolean if a field has been set.

### GetOrgPrivacyPolicy

`func (o *DescribeUserResult) GetOrgPrivacyPolicy() string`

GetOrgPrivacyPolicy returns the OrgPrivacyPolicy field if non-nil, zero value otherwise.

### GetOrgPrivacyPolicyOk

`func (o *DescribeUserResult) GetOrgPrivacyPolicyOk() (*string, bool)`

GetOrgPrivacyPolicyOk returns a tuple with the OrgPrivacyPolicy field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOrgPrivacyPolicy

`func (o *DescribeUserResult) SetOrgPrivacyPolicy(v string)`

SetOrgPrivacyPolicy sets OrgPrivacyPolicy field to given value.

### HasOrgPrivacyPolicy

`func (o *DescribeUserResult) HasOrgPrivacyPolicy() bool`

HasOrgPrivacyPolicy returns a boolean if a field has been set.

### GetOrgSupportEmail

`func (o *DescribeUserResult) GetOrgSupportEmail() string`

GetOrgSupportEmail returns the OrgSupportEmail field if non-nil, zero value otherwise.

### GetOrgSupportEmailOk

`func (o *DescribeUserResult) GetOrgSupportEmailOk() (*string, bool)`

GetOrgSupportEmailOk returns a tuple with the OrgSupportEmail field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOrgSupportEmail

`func (o *DescribeUserResult) SetOrgSupportEmail(v string)`

SetOrgSupportEmail sets OrgSupportEmail field to given value.

### HasOrgSupportEmail

`func (o *DescribeUserResult) HasOrgSupportEmail() bool`

HasOrgSupportEmail returns a boolean if a field has been set.

### GetOrgTermsOfUse

`func (o *DescribeUserResult) GetOrgTermsOfUse() string`

GetOrgTermsOfUse returns the OrgTermsOfUse field if non-nil, zero value otherwise.

### GetOrgTermsOfUseOk

`func (o *DescribeUserResult) GetOrgTermsOfUseOk() (*string, bool)`

GetOrgTermsOfUseOk returns a tuple with the OrgTermsOfUse field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOrgTermsOfUse

`func (o *DescribeUserResult) SetOrgTermsOfUse(v string)`

SetOrgTermsOfUse sets OrgTermsOfUse field to given value.

### HasOrgTermsOfUse

`func (o *DescribeUserResult) HasOrgTermsOfUse() bool`

HasOrgTermsOfUse returns a boolean if a field has been set.

### GetOrgURL

`func (o *DescribeUserResult) GetOrgURL() string`

GetOrgURL returns the OrgURL field if non-nil, zero value otherwise.

### GetOrgURLOk

`func (o *DescribeUserResult) GetOrgURLOk() (*string, bool)`

GetOrgURLOk returns a tuple with the OrgURL field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOrgURL

`func (o *DescribeUserResult) SetOrgURL(v string)`

SetOrgURL sets OrgURL field to given value.

### HasOrgURL

`func (o *DescribeUserResult) HasOrgURL() bool`

HasOrgURL returns a boolean if a field has been set.

### GetPlanName

`func (o *DescribeUserResult) GetPlanName() string`

GetPlanName returns the PlanName field if non-nil, zero value otherwise.

### GetPlanNameOk

`func (o *DescribeUserResult) GetPlanNameOk() (*string, bool)`

GetPlanNameOk returns a tuple with the PlanName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPlanName

`func (o *DescribeUserResult) SetPlanName(v string)`

SetPlanName sets PlanName field to given value.

### HasPlanName

`func (o *DescribeUserResult) HasPlanName() bool`

HasPlanName returns a boolean if a field has been set.

### GetRoleType

`func (o *DescribeUserResult) GetRoleType() string`

GetRoleType returns the RoleType field if non-nil, zero value otherwise.

### GetRoleTypeOk

`func (o *DescribeUserResult) GetRoleTypeOk() (*string, bool)`

GetRoleTypeOk returns a tuple with the RoleType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRoleType

`func (o *DescribeUserResult) SetRoleType(v string)`

SetRoleType sets RoleType field to given value.

### HasRoleType

`func (o *DescribeUserResult) HasRoleType() bool`

HasRoleType returns a boolean if a field has been set.

### GetStatus

`func (o *DescribeUserResult) GetStatus() string`

GetStatus returns the Status field if non-nil, zero value otherwise.

### GetStatusOk

`func (o *DescribeUserResult) GetStatusOk() (*string, bool)`

GetStatusOk returns a tuple with the Status field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStatus

`func (o *DescribeUserResult) SetStatus(v string)`

SetStatus sets Status field to given value.

### HasStatus

`func (o *DescribeUserResult) HasStatus() bool`

HasStatus returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


