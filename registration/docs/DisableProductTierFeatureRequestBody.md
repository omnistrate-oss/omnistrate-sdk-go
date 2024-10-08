# DisableProductTierFeatureRequestBody

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Feature** | **string** | Feature to disable | 
**Scope** | Pointer to **string** | Feature scope | [optional] 

## Methods

### NewDisableProductTierFeatureRequestBody

`func NewDisableProductTierFeatureRequestBody(feature string, ) *DisableProductTierFeatureRequestBody`

NewDisableProductTierFeatureRequestBody instantiates a new DisableProductTierFeatureRequestBody object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewDisableProductTierFeatureRequestBodyWithDefaults

`func NewDisableProductTierFeatureRequestBodyWithDefaults() *DisableProductTierFeatureRequestBody`

NewDisableProductTierFeatureRequestBodyWithDefaults instantiates a new DisableProductTierFeatureRequestBody object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetFeature

`func (o *DisableProductTierFeatureRequestBody) GetFeature() string`

GetFeature returns the Feature field if non-nil, zero value otherwise.

### GetFeatureOk

`func (o *DisableProductTierFeatureRequestBody) GetFeatureOk() (*string, bool)`

GetFeatureOk returns a tuple with the Feature field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFeature

`func (o *DisableProductTierFeatureRequestBody) SetFeature(v string)`

SetFeature sets Feature field to given value.


### GetScope

`func (o *DisableProductTierFeatureRequestBody) GetScope() string`

GetScope returns the Scope field if non-nil, zero value otherwise.

### GetScopeOk

`func (o *DisableProductTierFeatureRequestBody) GetScopeOk() (*string, bool)`

GetScopeOk returns a tuple with the Scope field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetScope

`func (o *DisableProductTierFeatureRequestBody) SetScope(v string)`

SetScope sets Scope field to given value.

### HasScope

`func (o *DisableProductTierFeatureRequestBody) HasScope() bool`

HasScope returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


