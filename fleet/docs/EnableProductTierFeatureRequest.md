# EnableProductTierFeatureRequest

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Configuration** | Pointer to **map[string]interface{}** | The configuration parameters of the product tier feature | [optional] 
**Feature** | **string** | ProductTierFeatureType is to enable / disable features per product tier | 
**Id** | **string** | ID of a Product Tier | 
**Scope** | Pointer to **string** | ProductTierFeatureScope defines scope of the feature within product tier | [optional] 
**ServiceId** | **string** | ID of a Service | 
**Token** | **string** | JWT token used to perform authorization | 

## Methods

### NewEnableProductTierFeatureRequest

`func NewEnableProductTierFeatureRequest(feature string, id string, serviceId string, token string, ) *EnableProductTierFeatureRequest`

NewEnableProductTierFeatureRequest instantiates a new EnableProductTierFeatureRequest object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewEnableProductTierFeatureRequestWithDefaults

`func NewEnableProductTierFeatureRequestWithDefaults() *EnableProductTierFeatureRequest`

NewEnableProductTierFeatureRequestWithDefaults instantiates a new EnableProductTierFeatureRequest object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetConfiguration

`func (o *EnableProductTierFeatureRequest) GetConfiguration() map[string]interface{}`

GetConfiguration returns the Configuration field if non-nil, zero value otherwise.

### GetConfigurationOk

`func (o *EnableProductTierFeatureRequest) GetConfigurationOk() (*map[string]interface{}, bool)`

GetConfigurationOk returns a tuple with the Configuration field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetConfiguration

`func (o *EnableProductTierFeatureRequest) SetConfiguration(v map[string]interface{})`

SetConfiguration sets Configuration field to given value.

### HasConfiguration

`func (o *EnableProductTierFeatureRequest) HasConfiguration() bool`

HasConfiguration returns a boolean if a field has been set.

### GetFeature

`func (o *EnableProductTierFeatureRequest) GetFeature() string`

GetFeature returns the Feature field if non-nil, zero value otherwise.

### GetFeatureOk

`func (o *EnableProductTierFeatureRequest) GetFeatureOk() (*string, bool)`

GetFeatureOk returns a tuple with the Feature field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFeature

`func (o *EnableProductTierFeatureRequest) SetFeature(v string)`

SetFeature sets Feature field to given value.


### GetId

`func (o *EnableProductTierFeatureRequest) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *EnableProductTierFeatureRequest) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *EnableProductTierFeatureRequest) SetId(v string)`

SetId sets Id field to given value.


### GetScope

`func (o *EnableProductTierFeatureRequest) GetScope() string`

GetScope returns the Scope field if non-nil, zero value otherwise.

### GetScopeOk

`func (o *EnableProductTierFeatureRequest) GetScopeOk() (*string, bool)`

GetScopeOk returns a tuple with the Scope field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetScope

`func (o *EnableProductTierFeatureRequest) SetScope(v string)`

SetScope sets Scope field to given value.

### HasScope

`func (o *EnableProductTierFeatureRequest) HasScope() bool`

HasScope returns a boolean if a field has been set.

### GetServiceId

`func (o *EnableProductTierFeatureRequest) GetServiceId() string`

GetServiceId returns the ServiceId field if non-nil, zero value otherwise.

### GetServiceIdOk

`func (o *EnableProductTierFeatureRequest) GetServiceIdOk() (*string, bool)`

GetServiceIdOk returns a tuple with the ServiceId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetServiceId

`func (o *EnableProductTierFeatureRequest) SetServiceId(v string)`

SetServiceId sets ServiceId field to given value.


### GetToken

`func (o *EnableProductTierFeatureRequest) GetToken() string`

GetToken returns the Token field if non-nil, zero value otherwise.

### GetTokenOk

`func (o *EnableProductTierFeatureRequest) GetTokenOk() (*string, bool)`

GetTokenOk returns a tuple with the Token field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetToken

`func (o *EnableProductTierFeatureRequest) SetToken(v string)`

SetToken sets Token field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


