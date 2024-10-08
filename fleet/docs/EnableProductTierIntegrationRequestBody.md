# EnableProductTierIntegrationRequestBody

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Configuration** | Pointer to **map[string]interface{}** | The configuration parameters for the integration provider. | [optional] 
**IntegrationProviderName** | **string** | Name of the product tier integration provider. | 
**IntegrationType** | **string** | Type of the product tier integration. | 

## Methods

### NewEnableProductTierIntegrationRequestBody

`func NewEnableProductTierIntegrationRequestBody(integrationProviderName string, integrationType string, ) *EnableProductTierIntegrationRequestBody`

NewEnableProductTierIntegrationRequestBody instantiates a new EnableProductTierIntegrationRequestBody object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewEnableProductTierIntegrationRequestBodyWithDefaults

`func NewEnableProductTierIntegrationRequestBodyWithDefaults() *EnableProductTierIntegrationRequestBody`

NewEnableProductTierIntegrationRequestBodyWithDefaults instantiates a new EnableProductTierIntegrationRequestBody object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetConfiguration

`func (o *EnableProductTierIntegrationRequestBody) GetConfiguration() map[string]interface{}`

GetConfiguration returns the Configuration field if non-nil, zero value otherwise.

### GetConfigurationOk

`func (o *EnableProductTierIntegrationRequestBody) GetConfigurationOk() (*map[string]interface{}, bool)`

GetConfigurationOk returns a tuple with the Configuration field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetConfiguration

`func (o *EnableProductTierIntegrationRequestBody) SetConfiguration(v map[string]interface{})`

SetConfiguration sets Configuration field to given value.

### HasConfiguration

`func (o *EnableProductTierIntegrationRequestBody) HasConfiguration() bool`

HasConfiguration returns a boolean if a field has been set.

### GetIntegrationProviderName

`func (o *EnableProductTierIntegrationRequestBody) GetIntegrationProviderName() string`

GetIntegrationProviderName returns the IntegrationProviderName field if non-nil, zero value otherwise.

### GetIntegrationProviderNameOk

`func (o *EnableProductTierIntegrationRequestBody) GetIntegrationProviderNameOk() (*string, bool)`

GetIntegrationProviderNameOk returns a tuple with the IntegrationProviderName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIntegrationProviderName

`func (o *EnableProductTierIntegrationRequestBody) SetIntegrationProviderName(v string)`

SetIntegrationProviderName sets IntegrationProviderName field to given value.


### GetIntegrationType

`func (o *EnableProductTierIntegrationRequestBody) GetIntegrationType() string`

GetIntegrationType returns the IntegrationType field if non-nil, zero value otherwise.

### GetIntegrationTypeOk

`func (o *EnableProductTierIntegrationRequestBody) GetIntegrationTypeOk() (*string, bool)`

GetIntegrationTypeOk returns a tuple with the IntegrationType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIntegrationType

`func (o *EnableProductTierIntegrationRequestBody) SetIntegrationType(v string)`

SetIntegrationType sets IntegrationType field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


