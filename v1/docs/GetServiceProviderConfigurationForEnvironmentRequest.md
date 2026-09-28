# GetServiceProviderConfigurationForEnvironmentRequest

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**EnvironmentType** | **string** | The type of the environment whose service provider configuration should be returned | 
**Token** | **string** | JWT token used to perform authorization | 

## Methods

### NewGetServiceProviderConfigurationForEnvironmentRequest

`func NewGetServiceProviderConfigurationForEnvironmentRequest(environmentType string, token string, ) *GetServiceProviderConfigurationForEnvironmentRequest`

NewGetServiceProviderConfigurationForEnvironmentRequest instantiates a new GetServiceProviderConfigurationForEnvironmentRequest object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewGetServiceProviderConfigurationForEnvironmentRequestWithDefaults

`func NewGetServiceProviderConfigurationForEnvironmentRequestWithDefaults() *GetServiceProviderConfigurationForEnvironmentRequest`

NewGetServiceProviderConfigurationForEnvironmentRequestWithDefaults instantiates a new GetServiceProviderConfigurationForEnvironmentRequest object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetEnvironmentType

`func (o *GetServiceProviderConfigurationForEnvironmentRequest) GetEnvironmentType() string`

GetEnvironmentType returns the EnvironmentType field if non-nil, zero value otherwise.

### GetEnvironmentTypeOk

`func (o *GetServiceProviderConfigurationForEnvironmentRequest) GetEnvironmentTypeOk() (*string, bool)`

GetEnvironmentTypeOk returns a tuple with the EnvironmentType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEnvironmentType

`func (o *GetServiceProviderConfigurationForEnvironmentRequest) SetEnvironmentType(v string)`

SetEnvironmentType sets EnvironmentType field to given value.


### GetToken

`func (o *GetServiceProviderConfigurationForEnvironmentRequest) GetToken() string`

GetToken returns the Token field if non-nil, zero value otherwise.

### GetTokenOk

`func (o *GetServiceProviderConfigurationForEnvironmentRequest) GetTokenOk() (*string, bool)`

GetTokenOk returns a tuple with the Token field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetToken

`func (o *GetServiceProviderConfigurationForEnvironmentRequest) SetToken(v string)`

SetToken sets Token field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


