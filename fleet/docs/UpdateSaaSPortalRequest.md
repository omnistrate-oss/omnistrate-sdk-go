# UpdateSaaSPortalRequest

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**CustomDomain** | Pointer to **string** | The custom domain for the SaaS portal | [optional] 
**EmailConfig** | Pointer to [**SaaSPortalEmailConfig**](SaaSPortalEmailConfig.md) |  | [optional] 
**EnvironmentType** | **string** | The type of service environment | 
**GoogleAnalyticsTagID** | Pointer to **string** | The Google Analytics tag ID for the SaaS portal | [optional] 
**ImageConfig** | Pointer to [**SaaSPortalImageConfig**](SaaSPortalImageConfig.md) |  | [optional] 
**Token** | **string** | JWT token used to perform authorization | 

## Methods

### NewUpdateSaaSPortalRequest

`func NewUpdateSaaSPortalRequest(environmentType string, token string, ) *UpdateSaaSPortalRequest`

NewUpdateSaaSPortalRequest instantiates a new UpdateSaaSPortalRequest object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewUpdateSaaSPortalRequestWithDefaults

`func NewUpdateSaaSPortalRequestWithDefaults() *UpdateSaaSPortalRequest`

NewUpdateSaaSPortalRequestWithDefaults instantiates a new UpdateSaaSPortalRequest object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetCustomDomain

`func (o *UpdateSaaSPortalRequest) GetCustomDomain() string`

GetCustomDomain returns the CustomDomain field if non-nil, zero value otherwise.

### GetCustomDomainOk

`func (o *UpdateSaaSPortalRequest) GetCustomDomainOk() (*string, bool)`

GetCustomDomainOk returns a tuple with the CustomDomain field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCustomDomain

`func (o *UpdateSaaSPortalRequest) SetCustomDomain(v string)`

SetCustomDomain sets CustomDomain field to given value.

### HasCustomDomain

`func (o *UpdateSaaSPortalRequest) HasCustomDomain() bool`

HasCustomDomain returns a boolean if a field has been set.

### GetEmailConfig

`func (o *UpdateSaaSPortalRequest) GetEmailConfig() SaaSPortalEmailConfig`

GetEmailConfig returns the EmailConfig field if non-nil, zero value otherwise.

### GetEmailConfigOk

`func (o *UpdateSaaSPortalRequest) GetEmailConfigOk() (*SaaSPortalEmailConfig, bool)`

GetEmailConfigOk returns a tuple with the EmailConfig field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEmailConfig

`func (o *UpdateSaaSPortalRequest) SetEmailConfig(v SaaSPortalEmailConfig)`

SetEmailConfig sets EmailConfig field to given value.

### HasEmailConfig

`func (o *UpdateSaaSPortalRequest) HasEmailConfig() bool`

HasEmailConfig returns a boolean if a field has been set.

### GetEnvironmentType

`func (o *UpdateSaaSPortalRequest) GetEnvironmentType() string`

GetEnvironmentType returns the EnvironmentType field if non-nil, zero value otherwise.

### GetEnvironmentTypeOk

`func (o *UpdateSaaSPortalRequest) GetEnvironmentTypeOk() (*string, bool)`

GetEnvironmentTypeOk returns a tuple with the EnvironmentType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEnvironmentType

`func (o *UpdateSaaSPortalRequest) SetEnvironmentType(v string)`

SetEnvironmentType sets EnvironmentType field to given value.


### GetGoogleAnalyticsTagID

`func (o *UpdateSaaSPortalRequest) GetGoogleAnalyticsTagID() string`

GetGoogleAnalyticsTagID returns the GoogleAnalyticsTagID field if non-nil, zero value otherwise.

### GetGoogleAnalyticsTagIDOk

`func (o *UpdateSaaSPortalRequest) GetGoogleAnalyticsTagIDOk() (*string, bool)`

GetGoogleAnalyticsTagIDOk returns a tuple with the GoogleAnalyticsTagID field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetGoogleAnalyticsTagID

`func (o *UpdateSaaSPortalRequest) SetGoogleAnalyticsTagID(v string)`

SetGoogleAnalyticsTagID sets GoogleAnalyticsTagID field to given value.

### HasGoogleAnalyticsTagID

`func (o *UpdateSaaSPortalRequest) HasGoogleAnalyticsTagID() bool`

HasGoogleAnalyticsTagID returns a boolean if a field has been set.

### GetImageConfig

`func (o *UpdateSaaSPortalRequest) GetImageConfig() SaaSPortalImageConfig`

GetImageConfig returns the ImageConfig field if non-nil, zero value otherwise.

### GetImageConfigOk

`func (o *UpdateSaaSPortalRequest) GetImageConfigOk() (*SaaSPortalImageConfig, bool)`

GetImageConfigOk returns a tuple with the ImageConfig field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetImageConfig

`func (o *UpdateSaaSPortalRequest) SetImageConfig(v SaaSPortalImageConfig)`

SetImageConfig sets ImageConfig field to given value.

### HasImageConfig

`func (o *UpdateSaaSPortalRequest) HasImageConfig() bool`

HasImageConfig returns a boolean if a field has been set.

### GetToken

`func (o *UpdateSaaSPortalRequest) GetToken() string`

GetToken returns the Token field if non-nil, zero value otherwise.

### GetTokenOk

`func (o *UpdateSaaSPortalRequest) GetTokenOk() (*string, bool)`

GetTokenOk returns a tuple with the Token field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetToken

`func (o *UpdateSaaSPortalRequest) SetToken(v string)`

SetToken sets Token field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


