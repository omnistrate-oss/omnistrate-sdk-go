# DisableProductTierFeatureRequest

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Feature** | **string** | ProductTierFeatureType is to enable / disable features per product tier | 
**Id** | **string** | ID of a Product Tier | 
**Scope** | Pointer to **string** | ProductTierFeatureScope defines scope of the feature within product tier | [optional] 
**ServiceId** | **string** | ID of a Service | 
**Token** | **string** | JWT token used to perform authorization | 

## Methods

### NewDisableProductTierFeatureRequest

`func NewDisableProductTierFeatureRequest(feature string, id string, serviceId string, token string, ) *DisableProductTierFeatureRequest`

NewDisableProductTierFeatureRequest instantiates a new DisableProductTierFeatureRequest object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewDisableProductTierFeatureRequestWithDefaults

`func NewDisableProductTierFeatureRequestWithDefaults() *DisableProductTierFeatureRequest`

NewDisableProductTierFeatureRequestWithDefaults instantiates a new DisableProductTierFeatureRequest object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetFeature

`func (o *DisableProductTierFeatureRequest) GetFeature() string`

GetFeature returns the Feature field if non-nil, zero value otherwise.

### GetFeatureOk

`func (o *DisableProductTierFeatureRequest) GetFeatureOk() (*string, bool)`

GetFeatureOk returns a tuple with the Feature field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFeature

`func (o *DisableProductTierFeatureRequest) SetFeature(v string)`

SetFeature sets Feature field to given value.


### GetId

`func (o *DisableProductTierFeatureRequest) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *DisableProductTierFeatureRequest) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *DisableProductTierFeatureRequest) SetId(v string)`

SetId sets Id field to given value.


### GetScope

`func (o *DisableProductTierFeatureRequest) GetScope() string`

GetScope returns the Scope field if non-nil, zero value otherwise.

### GetScopeOk

`func (o *DisableProductTierFeatureRequest) GetScopeOk() (*string, bool)`

GetScopeOk returns a tuple with the Scope field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetScope

`func (o *DisableProductTierFeatureRequest) SetScope(v string)`

SetScope sets Scope field to given value.

### HasScope

`func (o *DisableProductTierFeatureRequest) HasScope() bool`

HasScope returns a boolean if a field has been set.

### GetServiceId

`func (o *DisableProductTierFeatureRequest) GetServiceId() string`

GetServiceId returns the ServiceId field if non-nil, zero value otherwise.

### GetServiceIdOk

`func (o *DisableProductTierFeatureRequest) GetServiceIdOk() (*string, bool)`

GetServiceIdOk returns a tuple with the ServiceId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetServiceId

`func (o *DisableProductTierFeatureRequest) SetServiceId(v string)`

SetServiceId sets ServiceId field to given value.


### GetToken

`func (o *DisableProductTierFeatureRequest) GetToken() string`

GetToken returns the Token field if non-nil, zero value otherwise.

### GetTokenOk

`func (o *DisableProductTierFeatureRequest) GetTokenOk() (*string, bool)`

GetTokenOk returns a tuple with the Token field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetToken

`func (o *DisableProductTierFeatureRequest) SetToken(v string)`

SetToken sets Token field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


