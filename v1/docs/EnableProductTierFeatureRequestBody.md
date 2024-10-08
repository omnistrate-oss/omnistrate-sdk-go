# EnableProductTierFeatureRequestBody

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Configuration** | Pointer to **map[string]interface{}** | The configuration parameters of the product tier feature | [optional] 
**Feature** | **string** | Feature to enable | 
**Scope** | Pointer to **string** | Feature scope | [optional] 

## Methods

### NewEnableProductTierFeatureRequestBody

`func NewEnableProductTierFeatureRequestBody(feature string, ) *EnableProductTierFeatureRequestBody`

NewEnableProductTierFeatureRequestBody instantiates a new EnableProductTierFeatureRequestBody object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewEnableProductTierFeatureRequestBodyWithDefaults

`func NewEnableProductTierFeatureRequestBodyWithDefaults() *EnableProductTierFeatureRequestBody`

NewEnableProductTierFeatureRequestBodyWithDefaults instantiates a new EnableProductTierFeatureRequestBody object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetConfiguration

`func (o *EnableProductTierFeatureRequestBody) GetConfiguration() map[string]interface{}`

GetConfiguration returns the Configuration field if non-nil, zero value otherwise.

### GetConfigurationOk

`func (o *EnableProductTierFeatureRequestBody) GetConfigurationOk() (*map[string]interface{}, bool)`

GetConfigurationOk returns a tuple with the Configuration field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetConfiguration

`func (o *EnableProductTierFeatureRequestBody) SetConfiguration(v map[string]interface{})`

SetConfiguration sets Configuration field to given value.

### HasConfiguration

`func (o *EnableProductTierFeatureRequestBody) HasConfiguration() bool`

HasConfiguration returns a boolean if a field has been set.

### GetFeature

`func (o *EnableProductTierFeatureRequestBody) GetFeature() string`

GetFeature returns the Feature field if non-nil, zero value otherwise.

### GetFeatureOk

`func (o *EnableProductTierFeatureRequestBody) GetFeatureOk() (*string, bool)`

GetFeatureOk returns a tuple with the Feature field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFeature

`func (o *EnableProductTierFeatureRequestBody) SetFeature(v string)`

SetFeature sets Feature field to given value.


### GetScope

`func (o *EnableProductTierFeatureRequestBody) GetScope() string`

GetScope returns the Scope field if non-nil, zero value otherwise.

### GetScopeOk

`func (o *EnableProductTierFeatureRequestBody) GetScopeOk() (*string, bool)`

GetScopeOk returns a tuple with the Scope field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetScope

`func (o *EnableProductTierFeatureRequestBody) SetScope(v string)`

SetScope sets Scope field to given value.

### HasScope

`func (o *EnableProductTierFeatureRequestBody) HasScope() bool`

HasScope returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


