# BuildServiceFromComposeSpecResult

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**ProductTierID** | **string** | ID of a Product Tier | 
**ServiceEnvironmentID** | **string** | ID of a Service Environment | 
**ServiceID** | **string** | ID of a Service | 
**UndefinedResources** | Pointer to **map[string]string** | Resources that appear in the service plan but were not defined in the spec. These resources were not processed during the build. | [optional] 

## Methods

### NewBuildServiceFromComposeSpecResult

`func NewBuildServiceFromComposeSpecResult(productTierID string, serviceEnvironmentID string, serviceID string, ) *BuildServiceFromComposeSpecResult`

NewBuildServiceFromComposeSpecResult instantiates a new BuildServiceFromComposeSpecResult object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewBuildServiceFromComposeSpecResultWithDefaults

`func NewBuildServiceFromComposeSpecResultWithDefaults() *BuildServiceFromComposeSpecResult`

NewBuildServiceFromComposeSpecResultWithDefaults instantiates a new BuildServiceFromComposeSpecResult object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetProductTierID

`func (o *BuildServiceFromComposeSpecResult) GetProductTierID() string`

GetProductTierID returns the ProductTierID field if non-nil, zero value otherwise.

### GetProductTierIDOk

`func (o *BuildServiceFromComposeSpecResult) GetProductTierIDOk() (*string, bool)`

GetProductTierIDOk returns a tuple with the ProductTierID field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetProductTierID

`func (o *BuildServiceFromComposeSpecResult) SetProductTierID(v string)`

SetProductTierID sets ProductTierID field to given value.


### GetServiceEnvironmentID

`func (o *BuildServiceFromComposeSpecResult) GetServiceEnvironmentID() string`

GetServiceEnvironmentID returns the ServiceEnvironmentID field if non-nil, zero value otherwise.

### GetServiceEnvironmentIDOk

`func (o *BuildServiceFromComposeSpecResult) GetServiceEnvironmentIDOk() (*string, bool)`

GetServiceEnvironmentIDOk returns a tuple with the ServiceEnvironmentID field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetServiceEnvironmentID

`func (o *BuildServiceFromComposeSpecResult) SetServiceEnvironmentID(v string)`

SetServiceEnvironmentID sets ServiceEnvironmentID field to given value.


### GetServiceID

`func (o *BuildServiceFromComposeSpecResult) GetServiceID() string`

GetServiceID returns the ServiceID field if non-nil, zero value otherwise.

### GetServiceIDOk

`func (o *BuildServiceFromComposeSpecResult) GetServiceIDOk() (*string, bool)`

GetServiceIDOk returns a tuple with the ServiceID field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetServiceID

`func (o *BuildServiceFromComposeSpecResult) SetServiceID(v string)`

SetServiceID sets ServiceID field to given value.


### GetUndefinedResources

`func (o *BuildServiceFromComposeSpecResult) GetUndefinedResources() map[string]string`

GetUndefinedResources returns the UndefinedResources field if non-nil, zero value otherwise.

### GetUndefinedResourcesOk

`func (o *BuildServiceFromComposeSpecResult) GetUndefinedResourcesOk() (*map[string]string, bool)`

GetUndefinedResourcesOk returns a tuple with the UndefinedResources field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUndefinedResources

`func (o *BuildServiceFromComposeSpecResult) SetUndefinedResources(v map[string]string)`

SetUndefinedResources sets UndefinedResources field to given value.

### HasUndefinedResources

`func (o *BuildServiceFromComposeSpecResult) HasUndefinedResources() bool`

HasUndefinedResources returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


