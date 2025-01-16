# UpdateSaaSPortalCustomDomainRequest

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**CustomDomain** | Pointer to **string** | The custom domain | [optional] 
**Description** | Pointer to **string** | The custom domain description | [optional] 
**EnvironmentType** | **string** | The type of service environment | 
**Name** | Pointer to **string** | The custom domain name | [optional] 
**Token** | **string** | JWT token used to perform authorization | 

## Methods

### NewUpdateSaaSPortalCustomDomainRequest

`func NewUpdateSaaSPortalCustomDomainRequest(environmentType string, token string, ) *UpdateSaaSPortalCustomDomainRequest`

NewUpdateSaaSPortalCustomDomainRequest instantiates a new UpdateSaaSPortalCustomDomainRequest object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewUpdateSaaSPortalCustomDomainRequestWithDefaults

`func NewUpdateSaaSPortalCustomDomainRequestWithDefaults() *UpdateSaaSPortalCustomDomainRequest`

NewUpdateSaaSPortalCustomDomainRequestWithDefaults instantiates a new UpdateSaaSPortalCustomDomainRequest object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetCustomDomain

`func (o *UpdateSaaSPortalCustomDomainRequest) GetCustomDomain() string`

GetCustomDomain returns the CustomDomain field if non-nil, zero value otherwise.

### GetCustomDomainOk

`func (o *UpdateSaaSPortalCustomDomainRequest) GetCustomDomainOk() (*string, bool)`

GetCustomDomainOk returns a tuple with the CustomDomain field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCustomDomain

`func (o *UpdateSaaSPortalCustomDomainRequest) SetCustomDomain(v string)`

SetCustomDomain sets CustomDomain field to given value.

### HasCustomDomain

`func (o *UpdateSaaSPortalCustomDomainRequest) HasCustomDomain() bool`

HasCustomDomain returns a boolean if a field has been set.

### GetDescription

`func (o *UpdateSaaSPortalCustomDomainRequest) GetDescription() string`

GetDescription returns the Description field if non-nil, zero value otherwise.

### GetDescriptionOk

`func (o *UpdateSaaSPortalCustomDomainRequest) GetDescriptionOk() (*string, bool)`

GetDescriptionOk returns a tuple with the Description field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDescription

`func (o *UpdateSaaSPortalCustomDomainRequest) SetDescription(v string)`

SetDescription sets Description field to given value.

### HasDescription

`func (o *UpdateSaaSPortalCustomDomainRequest) HasDescription() bool`

HasDescription returns a boolean if a field has been set.

### GetEnvironmentType

`func (o *UpdateSaaSPortalCustomDomainRequest) GetEnvironmentType() string`

GetEnvironmentType returns the EnvironmentType field if non-nil, zero value otherwise.

### GetEnvironmentTypeOk

`func (o *UpdateSaaSPortalCustomDomainRequest) GetEnvironmentTypeOk() (*string, bool)`

GetEnvironmentTypeOk returns a tuple with the EnvironmentType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEnvironmentType

`func (o *UpdateSaaSPortalCustomDomainRequest) SetEnvironmentType(v string)`

SetEnvironmentType sets EnvironmentType field to given value.


### GetName

`func (o *UpdateSaaSPortalCustomDomainRequest) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *UpdateSaaSPortalCustomDomainRequest) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *UpdateSaaSPortalCustomDomainRequest) SetName(v string)`

SetName sets Name field to given value.

### HasName

`func (o *UpdateSaaSPortalCustomDomainRequest) HasName() bool`

HasName returns a boolean if a field has been set.

### GetToken

`func (o *UpdateSaaSPortalCustomDomainRequest) GetToken() string`

GetToken returns the Token field if non-nil, zero value otherwise.

### GetTokenOk

`func (o *UpdateSaaSPortalCustomDomainRequest) GetTokenOk() (*string, bool)`

GetTokenOk returns a tuple with the Token field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetToken

`func (o *UpdateSaaSPortalCustomDomainRequest) SetToken(v string)`

SetToken sets Token field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


