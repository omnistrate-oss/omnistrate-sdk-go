# BuildServiceFromServicePlanSpecResult

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**IsNewServicePlanVersionCreated** | Pointer to **bool** | Indicates if a new service plan version was created | [optional] 
**ProductTierID** | **string** | ID of a Product Tier | 
**ServiceEnvironmentID** | **string** | ID of a Service Environment | 
**ServiceID** | **string** | ID of a Service | 
**UndefinedResources** | Pointer to **map[string]string** | Resources that appear in the service plan but were not defined in the spec. These resources were not processed during the build. | [optional] 

## Methods

### NewBuildServiceFromServicePlanSpecResult

`func NewBuildServiceFromServicePlanSpecResult(productTierID string, serviceEnvironmentID string, serviceID string, ) *BuildServiceFromServicePlanSpecResult`

NewBuildServiceFromServicePlanSpecResult instantiates a new BuildServiceFromServicePlanSpecResult object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewBuildServiceFromServicePlanSpecResultWithDefaults

`func NewBuildServiceFromServicePlanSpecResultWithDefaults() *BuildServiceFromServicePlanSpecResult`

NewBuildServiceFromServicePlanSpecResultWithDefaults instantiates a new BuildServiceFromServicePlanSpecResult object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetIsNewServicePlanVersionCreated

`func (o *BuildServiceFromServicePlanSpecResult) GetIsNewServicePlanVersionCreated() bool`

GetIsNewServicePlanVersionCreated returns the IsNewServicePlanVersionCreated field if non-nil, zero value otherwise.

### GetIsNewServicePlanVersionCreatedOk

`func (o *BuildServiceFromServicePlanSpecResult) GetIsNewServicePlanVersionCreatedOk() (*bool, bool)`

GetIsNewServicePlanVersionCreatedOk returns a tuple with the IsNewServicePlanVersionCreated field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIsNewServicePlanVersionCreated

`func (o *BuildServiceFromServicePlanSpecResult) SetIsNewServicePlanVersionCreated(v bool)`

SetIsNewServicePlanVersionCreated sets IsNewServicePlanVersionCreated field to given value.

### HasIsNewServicePlanVersionCreated

`func (o *BuildServiceFromServicePlanSpecResult) HasIsNewServicePlanVersionCreated() bool`

HasIsNewServicePlanVersionCreated returns a boolean if a field has been set.

### GetProductTierID

`func (o *BuildServiceFromServicePlanSpecResult) GetProductTierID() string`

GetProductTierID returns the ProductTierID field if non-nil, zero value otherwise.

### GetProductTierIDOk

`func (o *BuildServiceFromServicePlanSpecResult) GetProductTierIDOk() (*string, bool)`

GetProductTierIDOk returns a tuple with the ProductTierID field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetProductTierID

`func (o *BuildServiceFromServicePlanSpecResult) SetProductTierID(v string)`

SetProductTierID sets ProductTierID field to given value.


### GetServiceEnvironmentID

`func (o *BuildServiceFromServicePlanSpecResult) GetServiceEnvironmentID() string`

GetServiceEnvironmentID returns the ServiceEnvironmentID field if non-nil, zero value otherwise.

### GetServiceEnvironmentIDOk

`func (o *BuildServiceFromServicePlanSpecResult) GetServiceEnvironmentIDOk() (*string, bool)`

GetServiceEnvironmentIDOk returns a tuple with the ServiceEnvironmentID field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetServiceEnvironmentID

`func (o *BuildServiceFromServicePlanSpecResult) SetServiceEnvironmentID(v string)`

SetServiceEnvironmentID sets ServiceEnvironmentID field to given value.


### GetServiceID

`func (o *BuildServiceFromServicePlanSpecResult) GetServiceID() string`

GetServiceID returns the ServiceID field if non-nil, zero value otherwise.

### GetServiceIDOk

`func (o *BuildServiceFromServicePlanSpecResult) GetServiceIDOk() (*string, bool)`

GetServiceIDOk returns a tuple with the ServiceID field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetServiceID

`func (o *BuildServiceFromServicePlanSpecResult) SetServiceID(v string)`

SetServiceID sets ServiceID field to given value.


### GetUndefinedResources

`func (o *BuildServiceFromServicePlanSpecResult) GetUndefinedResources() map[string]string`

GetUndefinedResources returns the UndefinedResources field if non-nil, zero value otherwise.

### GetUndefinedResourcesOk

`func (o *BuildServiceFromServicePlanSpecResult) GetUndefinedResourcesOk() (*map[string]string, bool)`

GetUndefinedResourcesOk returns a tuple with the UndefinedResources field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUndefinedResources

`func (o *BuildServiceFromServicePlanSpecResult) SetUndefinedResources(v map[string]string)`

SetUndefinedResources sets UndefinedResources field to given value.

### HasUndefinedResources

`func (o *BuildServiceFromServicePlanSpecResult) HasUndefinedResources() bool`

HasUndefinedResources returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


