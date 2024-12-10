# ProductTierIntegration

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Configuration** | **map[string]interface{}** | The configuration parameters for the integration provider. | 
**IntegrationFailureDetails** | Pointer to **string** | The details of the integration failure | [optional] 
**IntegrationProviderName** | **string** | The provider offering the integration for the product tier feature. | 
**IntegrationStatus** | **string** | The status of the integration | 
**IntegrationType** | **string** | ProductTierFeatureType is to enable / disable features per product tier | 

## Methods

### NewProductTierIntegration

`func NewProductTierIntegration(configuration map[string]interface{}, integrationProviderName string, integrationStatus string, integrationType string, ) *ProductTierIntegration`

NewProductTierIntegration instantiates a new ProductTierIntegration object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewProductTierIntegrationWithDefaults

`func NewProductTierIntegrationWithDefaults() *ProductTierIntegration`

NewProductTierIntegrationWithDefaults instantiates a new ProductTierIntegration object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetConfiguration

`func (o *ProductTierIntegration) GetConfiguration() map[string]interface{}`

GetConfiguration returns the Configuration field if non-nil, zero value otherwise.

### GetConfigurationOk

`func (o *ProductTierIntegration) GetConfigurationOk() (*map[string]interface{}, bool)`

GetConfigurationOk returns a tuple with the Configuration field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetConfiguration

`func (o *ProductTierIntegration) SetConfiguration(v map[string]interface{})`

SetConfiguration sets Configuration field to given value.


### GetIntegrationFailureDetails

`func (o *ProductTierIntegration) GetIntegrationFailureDetails() string`

GetIntegrationFailureDetails returns the IntegrationFailureDetails field if non-nil, zero value otherwise.

### GetIntegrationFailureDetailsOk

`func (o *ProductTierIntegration) GetIntegrationFailureDetailsOk() (*string, bool)`

GetIntegrationFailureDetailsOk returns a tuple with the IntegrationFailureDetails field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIntegrationFailureDetails

`func (o *ProductTierIntegration) SetIntegrationFailureDetails(v string)`

SetIntegrationFailureDetails sets IntegrationFailureDetails field to given value.

### HasIntegrationFailureDetails

`func (o *ProductTierIntegration) HasIntegrationFailureDetails() bool`

HasIntegrationFailureDetails returns a boolean if a field has been set.

### GetIntegrationProviderName

`func (o *ProductTierIntegration) GetIntegrationProviderName() string`

GetIntegrationProviderName returns the IntegrationProviderName field if non-nil, zero value otherwise.

### GetIntegrationProviderNameOk

`func (o *ProductTierIntegration) GetIntegrationProviderNameOk() (*string, bool)`

GetIntegrationProviderNameOk returns a tuple with the IntegrationProviderName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIntegrationProviderName

`func (o *ProductTierIntegration) SetIntegrationProviderName(v string)`

SetIntegrationProviderName sets IntegrationProviderName field to given value.


### GetIntegrationStatus

`func (o *ProductTierIntegration) GetIntegrationStatus() string`

GetIntegrationStatus returns the IntegrationStatus field if non-nil, zero value otherwise.

### GetIntegrationStatusOk

`func (o *ProductTierIntegration) GetIntegrationStatusOk() (*string, bool)`

GetIntegrationStatusOk returns a tuple with the IntegrationStatus field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIntegrationStatus

`func (o *ProductTierIntegration) SetIntegrationStatus(v string)`

SetIntegrationStatus sets IntegrationStatus field to given value.


### GetIntegrationType

`func (o *ProductTierIntegration) GetIntegrationType() string`

GetIntegrationType returns the IntegrationType field if non-nil, zero value otherwise.

### GetIntegrationTypeOk

`func (o *ProductTierIntegration) GetIntegrationTypeOk() (*string, bool)`

GetIntegrationTypeOk returns a tuple with the IntegrationType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIntegrationType

`func (o *ProductTierIntegration) SetIntegrationType(v string)`

SetIntegrationType sets IntegrationType field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


