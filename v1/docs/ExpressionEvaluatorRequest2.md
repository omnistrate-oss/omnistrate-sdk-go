# ExpressionEvaluatorRequest2

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**DeploymentCellID** | Pointer to **string** | Either this or the instanceID must be provided to evaluate the expression(s) | [optional] 
**Expression** | Pointer to **string** | Expression containing system and api parameter variable references | [optional] 
**ExpressionMap** | Pointer to **map[string]interface{}** | If provided, the &#x60;expression&#x60; field is ignored and all expressions in the map are evaluated. The map keys are preserved in the result. | [optional] 
**InstanceID** | Pointer to **string** | Either this or the deploymentCellID must be provided to evaluate the expression(s) | [optional] 
**ProductTierID** | Pointer to **string** | Mandatory if the instanceID is not provided | [optional] 
**ResourceKey** | **string** | The resource key to use for evaluating resource parameters | 
**ServiceID** | **string** | The service ID to use for evaluating service parameters | 

## Methods

### NewExpressionEvaluatorRequest2

`func NewExpressionEvaluatorRequest2(resourceKey string, serviceID string, ) *ExpressionEvaluatorRequest2`

NewExpressionEvaluatorRequest2 instantiates a new ExpressionEvaluatorRequest2 object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewExpressionEvaluatorRequest2WithDefaults

`func NewExpressionEvaluatorRequest2WithDefaults() *ExpressionEvaluatorRequest2`

NewExpressionEvaluatorRequest2WithDefaults instantiates a new ExpressionEvaluatorRequest2 object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetDeploymentCellID

`func (o *ExpressionEvaluatorRequest2) GetDeploymentCellID() string`

GetDeploymentCellID returns the DeploymentCellID field if non-nil, zero value otherwise.

### GetDeploymentCellIDOk

`func (o *ExpressionEvaluatorRequest2) GetDeploymentCellIDOk() (*string, bool)`

GetDeploymentCellIDOk returns a tuple with the DeploymentCellID field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDeploymentCellID

`func (o *ExpressionEvaluatorRequest2) SetDeploymentCellID(v string)`

SetDeploymentCellID sets DeploymentCellID field to given value.

### HasDeploymentCellID

`func (o *ExpressionEvaluatorRequest2) HasDeploymentCellID() bool`

HasDeploymentCellID returns a boolean if a field has been set.

### GetExpression

`func (o *ExpressionEvaluatorRequest2) GetExpression() string`

GetExpression returns the Expression field if non-nil, zero value otherwise.

### GetExpressionOk

`func (o *ExpressionEvaluatorRequest2) GetExpressionOk() (*string, bool)`

GetExpressionOk returns a tuple with the Expression field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetExpression

`func (o *ExpressionEvaluatorRequest2) SetExpression(v string)`

SetExpression sets Expression field to given value.

### HasExpression

`func (o *ExpressionEvaluatorRequest2) HasExpression() bool`

HasExpression returns a boolean if a field has been set.

### GetExpressionMap

`func (o *ExpressionEvaluatorRequest2) GetExpressionMap() map[string]interface{}`

GetExpressionMap returns the ExpressionMap field if non-nil, zero value otherwise.

### GetExpressionMapOk

`func (o *ExpressionEvaluatorRequest2) GetExpressionMapOk() (*map[string]interface{}, bool)`

GetExpressionMapOk returns a tuple with the ExpressionMap field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetExpressionMap

`func (o *ExpressionEvaluatorRequest2) SetExpressionMap(v map[string]interface{})`

SetExpressionMap sets ExpressionMap field to given value.

### HasExpressionMap

`func (o *ExpressionEvaluatorRequest2) HasExpressionMap() bool`

HasExpressionMap returns a boolean if a field has been set.

### GetInstanceID

`func (o *ExpressionEvaluatorRequest2) GetInstanceID() string`

GetInstanceID returns the InstanceID field if non-nil, zero value otherwise.

### GetInstanceIDOk

`func (o *ExpressionEvaluatorRequest2) GetInstanceIDOk() (*string, bool)`

GetInstanceIDOk returns a tuple with the InstanceID field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetInstanceID

`func (o *ExpressionEvaluatorRequest2) SetInstanceID(v string)`

SetInstanceID sets InstanceID field to given value.

### HasInstanceID

`func (o *ExpressionEvaluatorRequest2) HasInstanceID() bool`

HasInstanceID returns a boolean if a field has been set.

### GetProductTierID

`func (o *ExpressionEvaluatorRequest2) GetProductTierID() string`

GetProductTierID returns the ProductTierID field if non-nil, zero value otherwise.

### GetProductTierIDOk

`func (o *ExpressionEvaluatorRequest2) GetProductTierIDOk() (*string, bool)`

GetProductTierIDOk returns a tuple with the ProductTierID field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetProductTierID

`func (o *ExpressionEvaluatorRequest2) SetProductTierID(v string)`

SetProductTierID sets ProductTierID field to given value.

### HasProductTierID

`func (o *ExpressionEvaluatorRequest2) HasProductTierID() bool`

HasProductTierID returns a boolean if a field has been set.

### GetResourceKey

`func (o *ExpressionEvaluatorRequest2) GetResourceKey() string`

GetResourceKey returns the ResourceKey field if non-nil, zero value otherwise.

### GetResourceKeyOk

`func (o *ExpressionEvaluatorRequest2) GetResourceKeyOk() (*string, bool)`

GetResourceKeyOk returns a tuple with the ResourceKey field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetResourceKey

`func (o *ExpressionEvaluatorRequest2) SetResourceKey(v string)`

SetResourceKey sets ResourceKey field to given value.


### GetServiceID

`func (o *ExpressionEvaluatorRequest2) GetServiceID() string`

GetServiceID returns the ServiceID field if non-nil, zero value otherwise.

### GetServiceIDOk

`func (o *ExpressionEvaluatorRequest2) GetServiceIDOk() (*string, bool)`

GetServiceIDOk returns a tuple with the ServiceID field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetServiceID

`func (o *ExpressionEvaluatorRequest2) SetServiceID(v string)`

SetServiceID sets ServiceID field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


