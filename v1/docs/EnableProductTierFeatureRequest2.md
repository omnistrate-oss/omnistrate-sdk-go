# EnableProductTierFeatureRequest2

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Configuration** | Pointer to **map[string]interface{}** | The configuration parameters of the product tier feature | [optional] 
**Feature** | **string** | Feature to enable | 
**Scope** | Pointer to **string** | Feature scope | [optional] 

## Methods

### NewEnableProductTierFeatureRequest2

`func NewEnableProductTierFeatureRequest2(feature string, ) *EnableProductTierFeatureRequest2`

NewEnableProductTierFeatureRequest2 instantiates a new EnableProductTierFeatureRequest2 object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewEnableProductTierFeatureRequest2WithDefaults

`func NewEnableProductTierFeatureRequest2WithDefaults() *EnableProductTierFeatureRequest2`

NewEnableProductTierFeatureRequest2WithDefaults instantiates a new EnableProductTierFeatureRequest2 object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetConfiguration

`func (o *EnableProductTierFeatureRequest2) GetConfiguration() map[string]interface{}`

GetConfiguration returns the Configuration field if non-nil, zero value otherwise.

### GetConfigurationOk

`func (o *EnableProductTierFeatureRequest2) GetConfigurationOk() (*map[string]interface{}, bool)`

GetConfigurationOk returns a tuple with the Configuration field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetConfiguration

`func (o *EnableProductTierFeatureRequest2) SetConfiguration(v map[string]interface{})`

SetConfiguration sets Configuration field to given value.

### HasConfiguration

`func (o *EnableProductTierFeatureRequest2) HasConfiguration() bool`

HasConfiguration returns a boolean if a field has been set.

### GetFeature

`func (o *EnableProductTierFeatureRequest2) GetFeature() string`

GetFeature returns the Feature field if non-nil, zero value otherwise.

### GetFeatureOk

`func (o *EnableProductTierFeatureRequest2) GetFeatureOk() (*string, bool)`

GetFeatureOk returns a tuple with the Feature field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFeature

`func (o *EnableProductTierFeatureRequest2) SetFeature(v string)`

SetFeature sets Feature field to given value.


### GetScope

`func (o *EnableProductTierFeatureRequest2) GetScope() string`

GetScope returns the Scope field if non-nil, zero value otherwise.

### GetScopeOk

`func (o *EnableProductTierFeatureRequest2) GetScopeOk() (*string, bool)`

GetScopeOk returns a tuple with the Scope field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetScope

`func (o *EnableProductTierFeatureRequest2) SetScope(v string)`

SetScope sets Scope field to given value.

### HasScope

`func (o *EnableProductTierFeatureRequest2) HasScope() bool`

HasScope returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


