# EnableProductTierIntegrationRequest2

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Configuration** | Pointer to **map[string]interface{}** | The configuration parameters for the integration provider. | [optional] 
**IntegrationProviderName** | **string** | Name of the product tier integration provider. | 
**IntegrationType** | **string** | Type of the product tier integration. | 

## Methods

### NewEnableProductTierIntegrationRequest2

`func NewEnableProductTierIntegrationRequest2(integrationProviderName string, integrationType string, ) *EnableProductTierIntegrationRequest2`

NewEnableProductTierIntegrationRequest2 instantiates a new EnableProductTierIntegrationRequest2 object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewEnableProductTierIntegrationRequest2WithDefaults

`func NewEnableProductTierIntegrationRequest2WithDefaults() *EnableProductTierIntegrationRequest2`

NewEnableProductTierIntegrationRequest2WithDefaults instantiates a new EnableProductTierIntegrationRequest2 object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetConfiguration

`func (o *EnableProductTierIntegrationRequest2) GetConfiguration() map[string]interface{}`

GetConfiguration returns the Configuration field if non-nil, zero value otherwise.

### GetConfigurationOk

`func (o *EnableProductTierIntegrationRequest2) GetConfigurationOk() (*map[string]interface{}, bool)`

GetConfigurationOk returns a tuple with the Configuration field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetConfiguration

`func (o *EnableProductTierIntegrationRequest2) SetConfiguration(v map[string]interface{})`

SetConfiguration sets Configuration field to given value.

### HasConfiguration

`func (o *EnableProductTierIntegrationRequest2) HasConfiguration() bool`

HasConfiguration returns a boolean if a field has been set.

### GetIntegrationProviderName

`func (o *EnableProductTierIntegrationRequest2) GetIntegrationProviderName() string`

GetIntegrationProviderName returns the IntegrationProviderName field if non-nil, zero value otherwise.

### GetIntegrationProviderNameOk

`func (o *EnableProductTierIntegrationRequest2) GetIntegrationProviderNameOk() (*string, bool)`

GetIntegrationProviderNameOk returns a tuple with the IntegrationProviderName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIntegrationProviderName

`func (o *EnableProductTierIntegrationRequest2) SetIntegrationProviderName(v string)`

SetIntegrationProviderName sets IntegrationProviderName field to given value.


### GetIntegrationType

`func (o *EnableProductTierIntegrationRequest2) GetIntegrationType() string`

GetIntegrationType returns the IntegrationType field if non-nil, zero value otherwise.

### GetIntegrationTypeOk

`func (o *EnableProductTierIntegrationRequest2) GetIntegrationTypeOk() (*string, bool)`

GetIntegrationTypeOk returns a tuple with the IntegrationType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIntegrationType

`func (o *EnableProductTierIntegrationRequest2) SetIntegrationType(v string)`

SetIntegrationType sets IntegrationType field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


