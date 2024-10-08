# UpdateSaaSPortalRequestBody

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**CustomDomain** | Pointer to **string** | The custom domain for the SaaS portal | [optional] 
**EmailConfig** | Pointer to [**SaaSPortalEmailConfig**](SaaSPortalEmailConfig.md) |  | [optional] 
**GoogleAnalyticsTagID** | Pointer to **string** | The Google Analytics tag ID for the SaaS portal | [optional] 
**ImageConfig** | Pointer to [**SaaSPortalImageConfig**](SaaSPortalImageConfig.md) |  | [optional] 

## Methods

### NewUpdateSaaSPortalRequestBody

`func NewUpdateSaaSPortalRequestBody() *UpdateSaaSPortalRequestBody`

NewUpdateSaaSPortalRequestBody instantiates a new UpdateSaaSPortalRequestBody object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewUpdateSaaSPortalRequestBodyWithDefaults

`func NewUpdateSaaSPortalRequestBodyWithDefaults() *UpdateSaaSPortalRequestBody`

NewUpdateSaaSPortalRequestBodyWithDefaults instantiates a new UpdateSaaSPortalRequestBody object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetCustomDomain

`func (o *UpdateSaaSPortalRequestBody) GetCustomDomain() string`

GetCustomDomain returns the CustomDomain field if non-nil, zero value otherwise.

### GetCustomDomainOk

`func (o *UpdateSaaSPortalRequestBody) GetCustomDomainOk() (*string, bool)`

GetCustomDomainOk returns a tuple with the CustomDomain field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCustomDomain

`func (o *UpdateSaaSPortalRequestBody) SetCustomDomain(v string)`

SetCustomDomain sets CustomDomain field to given value.

### HasCustomDomain

`func (o *UpdateSaaSPortalRequestBody) HasCustomDomain() bool`

HasCustomDomain returns a boolean if a field has been set.

### GetEmailConfig

`func (o *UpdateSaaSPortalRequestBody) GetEmailConfig() SaaSPortalEmailConfig`

GetEmailConfig returns the EmailConfig field if non-nil, zero value otherwise.

### GetEmailConfigOk

`func (o *UpdateSaaSPortalRequestBody) GetEmailConfigOk() (*SaaSPortalEmailConfig, bool)`

GetEmailConfigOk returns a tuple with the EmailConfig field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEmailConfig

`func (o *UpdateSaaSPortalRequestBody) SetEmailConfig(v SaaSPortalEmailConfig)`

SetEmailConfig sets EmailConfig field to given value.

### HasEmailConfig

`func (o *UpdateSaaSPortalRequestBody) HasEmailConfig() bool`

HasEmailConfig returns a boolean if a field has been set.

### GetGoogleAnalyticsTagID

`func (o *UpdateSaaSPortalRequestBody) GetGoogleAnalyticsTagID() string`

GetGoogleAnalyticsTagID returns the GoogleAnalyticsTagID field if non-nil, zero value otherwise.

### GetGoogleAnalyticsTagIDOk

`func (o *UpdateSaaSPortalRequestBody) GetGoogleAnalyticsTagIDOk() (*string, bool)`

GetGoogleAnalyticsTagIDOk returns a tuple with the GoogleAnalyticsTagID field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetGoogleAnalyticsTagID

`func (o *UpdateSaaSPortalRequestBody) SetGoogleAnalyticsTagID(v string)`

SetGoogleAnalyticsTagID sets GoogleAnalyticsTagID field to given value.

### HasGoogleAnalyticsTagID

`func (o *UpdateSaaSPortalRequestBody) HasGoogleAnalyticsTagID() bool`

HasGoogleAnalyticsTagID returns a boolean if a field has been set.

### GetImageConfig

`func (o *UpdateSaaSPortalRequestBody) GetImageConfig() SaaSPortalImageConfig`

GetImageConfig returns the ImageConfig field if non-nil, zero value otherwise.

### GetImageConfigOk

`func (o *UpdateSaaSPortalRequestBody) GetImageConfigOk() (*SaaSPortalImageConfig, bool)`

GetImageConfigOk returns a tuple with the ImageConfig field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetImageConfig

`func (o *UpdateSaaSPortalRequestBody) SetImageConfig(v SaaSPortalImageConfig)`

SetImageConfig sets ImageConfig field to given value.

### HasImageConfig

`func (o *UpdateSaaSPortalRequestBody) HasImageConfig() bool`

HasImageConfig returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


