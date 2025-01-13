# IntegrationStatus

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Url** | Pointer to **string** | Additional URL relevant for integration (optional and integration specific) | [optional] 
**Configuration** | Pointer to **map[string]interface{}** | The configuration of the integration | [optional] 
**HealthStatus** | Pointer to **string** | The health status of the integration | [optional] 
**IntegrationType** | Pointer to **string** | The name of the integration or feature | [optional] 
**Message** | Pointer to **string** | Additional details regarding integration health | [optional] 
**Scope** | Pointer to **string** | Scope of the feature/integration | [optional] 

## Methods

### NewIntegrationStatus

`func NewIntegrationStatus() *IntegrationStatus`

NewIntegrationStatus instantiates a new IntegrationStatus object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewIntegrationStatusWithDefaults

`func NewIntegrationStatusWithDefaults() *IntegrationStatus`

NewIntegrationStatusWithDefaults instantiates a new IntegrationStatus object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetUrl

`func (o *IntegrationStatus) GetUrl() string`

GetUrl returns the Url field if non-nil, zero value otherwise.

### GetUrlOk

`func (o *IntegrationStatus) GetUrlOk() (*string, bool)`

GetUrlOk returns a tuple with the Url field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUrl

`func (o *IntegrationStatus) SetUrl(v string)`

SetUrl sets Url field to given value.

### HasUrl

`func (o *IntegrationStatus) HasUrl() bool`

HasUrl returns a boolean if a field has been set.

### GetConfiguration

`func (o *IntegrationStatus) GetConfiguration() map[string]interface{}`

GetConfiguration returns the Configuration field if non-nil, zero value otherwise.

### GetConfigurationOk

`func (o *IntegrationStatus) GetConfigurationOk() (*map[string]interface{}, bool)`

GetConfigurationOk returns a tuple with the Configuration field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetConfiguration

`func (o *IntegrationStatus) SetConfiguration(v map[string]interface{})`

SetConfiguration sets Configuration field to given value.

### HasConfiguration

`func (o *IntegrationStatus) HasConfiguration() bool`

HasConfiguration returns a boolean if a field has been set.

### GetHealthStatus

`func (o *IntegrationStatus) GetHealthStatus() string`

GetHealthStatus returns the HealthStatus field if non-nil, zero value otherwise.

### GetHealthStatusOk

`func (o *IntegrationStatus) GetHealthStatusOk() (*string, bool)`

GetHealthStatusOk returns a tuple with the HealthStatus field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetHealthStatus

`func (o *IntegrationStatus) SetHealthStatus(v string)`

SetHealthStatus sets HealthStatus field to given value.

### HasHealthStatus

`func (o *IntegrationStatus) HasHealthStatus() bool`

HasHealthStatus returns a boolean if a field has been set.

### GetIntegrationType

`func (o *IntegrationStatus) GetIntegrationType() string`

GetIntegrationType returns the IntegrationType field if non-nil, zero value otherwise.

### GetIntegrationTypeOk

`func (o *IntegrationStatus) GetIntegrationTypeOk() (*string, bool)`

GetIntegrationTypeOk returns a tuple with the IntegrationType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIntegrationType

`func (o *IntegrationStatus) SetIntegrationType(v string)`

SetIntegrationType sets IntegrationType field to given value.

### HasIntegrationType

`func (o *IntegrationStatus) HasIntegrationType() bool`

HasIntegrationType returns a boolean if a field has been set.

### GetMessage

`func (o *IntegrationStatus) GetMessage() string`

GetMessage returns the Message field if non-nil, zero value otherwise.

### GetMessageOk

`func (o *IntegrationStatus) GetMessageOk() (*string, bool)`

GetMessageOk returns a tuple with the Message field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMessage

`func (o *IntegrationStatus) SetMessage(v string)`

SetMessage sets Message field to given value.

### HasMessage

`func (o *IntegrationStatus) HasMessage() bool`

HasMessage returns a boolean if a field has been set.

### GetScope

`func (o *IntegrationStatus) GetScope() string`

GetScope returns the Scope field if non-nil, zero value otherwise.

### GetScopeOk

`func (o *IntegrationStatus) GetScopeOk() (*string, bool)`

GetScopeOk returns a tuple with the Scope field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetScope

`func (o *IntegrationStatus) SetScope(v string)`

SetScope sets Scope field to given value.

### HasScope

`func (o *IntegrationStatus) HasScope() bool`

HasScope returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


