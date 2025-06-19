# RenderIdentityProvidersRequest

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**EnvironmentType** | Pointer to **string** | The type of environment to filter Identity Providers by | [optional] 
**LoginHint** | Pointer to **string** | The login hint to use for the Identity Provider | [optional] 
**RedirectUrl** | Pointer to **string** | The URL to redirect to after login | [optional] 
**Token** | **string** | JWT token used to perform authorization | 

## Methods

### NewRenderIdentityProvidersRequest

`func NewRenderIdentityProvidersRequest(token string, ) *RenderIdentityProvidersRequest`

NewRenderIdentityProvidersRequest instantiates a new RenderIdentityProvidersRequest object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewRenderIdentityProvidersRequestWithDefaults

`func NewRenderIdentityProvidersRequestWithDefaults() *RenderIdentityProvidersRequest`

NewRenderIdentityProvidersRequestWithDefaults instantiates a new RenderIdentityProvidersRequest object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetEnvironmentType

`func (o *RenderIdentityProvidersRequest) GetEnvironmentType() string`

GetEnvironmentType returns the EnvironmentType field if non-nil, zero value otherwise.

### GetEnvironmentTypeOk

`func (o *RenderIdentityProvidersRequest) GetEnvironmentTypeOk() (*string, bool)`

GetEnvironmentTypeOk returns a tuple with the EnvironmentType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEnvironmentType

`func (o *RenderIdentityProvidersRequest) SetEnvironmentType(v string)`

SetEnvironmentType sets EnvironmentType field to given value.

### HasEnvironmentType

`func (o *RenderIdentityProvidersRequest) HasEnvironmentType() bool`

HasEnvironmentType returns a boolean if a field has been set.

### GetLoginHint

`func (o *RenderIdentityProvidersRequest) GetLoginHint() string`

GetLoginHint returns the LoginHint field if non-nil, zero value otherwise.

### GetLoginHintOk

`func (o *RenderIdentityProvidersRequest) GetLoginHintOk() (*string, bool)`

GetLoginHintOk returns a tuple with the LoginHint field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLoginHint

`func (o *RenderIdentityProvidersRequest) SetLoginHint(v string)`

SetLoginHint sets LoginHint field to given value.

### HasLoginHint

`func (o *RenderIdentityProvidersRequest) HasLoginHint() bool`

HasLoginHint returns a boolean if a field has been set.

### GetRedirectUrl

`func (o *RenderIdentityProvidersRequest) GetRedirectUrl() string`

GetRedirectUrl returns the RedirectUrl field if non-nil, zero value otherwise.

### GetRedirectUrlOk

`func (o *RenderIdentityProvidersRequest) GetRedirectUrlOk() (*string, bool)`

GetRedirectUrlOk returns a tuple with the RedirectUrl field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRedirectUrl

`func (o *RenderIdentityProvidersRequest) SetRedirectUrl(v string)`

SetRedirectUrl sets RedirectUrl field to given value.

### HasRedirectUrl

`func (o *RenderIdentityProvidersRequest) HasRedirectUrl() bool`

HasRedirectUrl returns a boolean if a field has been set.

### GetToken

`func (o *RenderIdentityProvidersRequest) GetToken() string`

GetToken returns the Token field if non-nil, zero value otherwise.

### GetTokenOk

`func (o *RenderIdentityProvidersRequest) GetTokenOk() (*string, bool)`

GetTokenOk returns a tuple with the Token field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetToken

`func (o *RenderIdentityProvidersRequest) SetToken(v string)`

SetToken sets Token field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


