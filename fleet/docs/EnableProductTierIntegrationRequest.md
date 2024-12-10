# EnableProductTierIntegrationRequest

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Configuration** | Pointer to **map[string]interface{}** | The configuration parameters for the integration provider. | [optional] 
**Id** | **string** | ID of a Product Tier | 
**IntegrationProviderName** | **string** | The provider offering the integration for the product tier feature. | 
**IntegrationType** | **string** | ProductTierFeatureType is to enable / disable features per product tier | 
**ServiceId** | **string** | ID of a Service | 
**Token** | **string** | JWT token used to perform authorization | 

## Methods

### NewEnableProductTierIntegrationRequest

`func NewEnableProductTierIntegrationRequest(id string, integrationProviderName string, integrationType string, serviceId string, token string, ) *EnableProductTierIntegrationRequest`

NewEnableProductTierIntegrationRequest instantiates a new EnableProductTierIntegrationRequest object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewEnableProductTierIntegrationRequestWithDefaults

`func NewEnableProductTierIntegrationRequestWithDefaults() *EnableProductTierIntegrationRequest`

NewEnableProductTierIntegrationRequestWithDefaults instantiates a new EnableProductTierIntegrationRequest object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetConfiguration

`func (o *EnableProductTierIntegrationRequest) GetConfiguration() map[string]interface{}`

GetConfiguration returns the Configuration field if non-nil, zero value otherwise.

### GetConfigurationOk

`func (o *EnableProductTierIntegrationRequest) GetConfigurationOk() (*map[string]interface{}, bool)`

GetConfigurationOk returns a tuple with the Configuration field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetConfiguration

`func (o *EnableProductTierIntegrationRequest) SetConfiguration(v map[string]interface{})`

SetConfiguration sets Configuration field to given value.

### HasConfiguration

`func (o *EnableProductTierIntegrationRequest) HasConfiguration() bool`

HasConfiguration returns a boolean if a field has been set.

### GetId

`func (o *EnableProductTierIntegrationRequest) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *EnableProductTierIntegrationRequest) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *EnableProductTierIntegrationRequest) SetId(v string)`

SetId sets Id field to given value.


### GetIntegrationProviderName

`func (o *EnableProductTierIntegrationRequest) GetIntegrationProviderName() string`

GetIntegrationProviderName returns the IntegrationProviderName field if non-nil, zero value otherwise.

### GetIntegrationProviderNameOk

`func (o *EnableProductTierIntegrationRequest) GetIntegrationProviderNameOk() (*string, bool)`

GetIntegrationProviderNameOk returns a tuple with the IntegrationProviderName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIntegrationProviderName

`func (o *EnableProductTierIntegrationRequest) SetIntegrationProviderName(v string)`

SetIntegrationProviderName sets IntegrationProviderName field to given value.


### GetIntegrationType

`func (o *EnableProductTierIntegrationRequest) GetIntegrationType() string`

GetIntegrationType returns the IntegrationType field if non-nil, zero value otherwise.

### GetIntegrationTypeOk

`func (o *EnableProductTierIntegrationRequest) GetIntegrationTypeOk() (*string, bool)`

GetIntegrationTypeOk returns a tuple with the IntegrationType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIntegrationType

`func (o *EnableProductTierIntegrationRequest) SetIntegrationType(v string)`

SetIntegrationType sets IntegrationType field to given value.


### GetServiceId

`func (o *EnableProductTierIntegrationRequest) GetServiceId() string`

GetServiceId returns the ServiceId field if non-nil, zero value otherwise.

### GetServiceIdOk

`func (o *EnableProductTierIntegrationRequest) GetServiceIdOk() (*string, bool)`

GetServiceIdOk returns a tuple with the ServiceId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetServiceId

`func (o *EnableProductTierIntegrationRequest) SetServiceId(v string)`

SetServiceId sets ServiceId field to given value.


### GetToken

`func (o *EnableProductTierIntegrationRequest) GetToken() string`

GetToken returns the Token field if non-nil, zero value otherwise.

### GetTokenOk

`func (o *EnableProductTierIntegrationRequest) GetTokenOk() (*string, bool)`

GetTokenOk returns a tuple with the Token field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetToken

`func (o *EnableProductTierIntegrationRequest) SetToken(v string)`

SetToken sets Token field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


