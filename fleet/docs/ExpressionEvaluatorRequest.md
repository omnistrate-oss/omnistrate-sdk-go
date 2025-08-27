# ExpressionEvaluatorRequest

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**DeploymentCellID** | Pointer to **string** | ID of a Host Cluster | [optional] 
**Expression** | Pointer to **string** | Expression containing system and api parameter variable references | [optional] 
**ExpressionMap** | Pointer to **map[string]interface{}** | If provided, the &#x60;expression&#x60; field is ignored and all expressions in the map are evaluated. The map keys are preserved in the result. | [optional] 
**InstanceID** | Pointer to **string** | ID of a Resource Instance | [optional] 
**ProductTierID** | Pointer to **string** | ID of a Product Tier | [optional] 
**ResourceKey** | **string** | The resource key to use for evaluating resource parameters | 
**ServiceID** | **string** | ID of a Service | 
**Token** | **string** | JWT token used to perform authorization | 

## Methods

### NewExpressionEvaluatorRequest

`func NewExpressionEvaluatorRequest(resourceKey string, serviceID string, token string, ) *ExpressionEvaluatorRequest`

NewExpressionEvaluatorRequest instantiates a new ExpressionEvaluatorRequest object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewExpressionEvaluatorRequestWithDefaults

`func NewExpressionEvaluatorRequestWithDefaults() *ExpressionEvaluatorRequest`

NewExpressionEvaluatorRequestWithDefaults instantiates a new ExpressionEvaluatorRequest object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetDeploymentCellID

`func (o *ExpressionEvaluatorRequest) GetDeploymentCellID() string`

GetDeploymentCellID returns the DeploymentCellID field if non-nil, zero value otherwise.

### GetDeploymentCellIDOk

`func (o *ExpressionEvaluatorRequest) GetDeploymentCellIDOk() (*string, bool)`

GetDeploymentCellIDOk returns a tuple with the DeploymentCellID field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDeploymentCellID

`func (o *ExpressionEvaluatorRequest) SetDeploymentCellID(v string)`

SetDeploymentCellID sets DeploymentCellID field to given value.

### HasDeploymentCellID

`func (o *ExpressionEvaluatorRequest) HasDeploymentCellID() bool`

HasDeploymentCellID returns a boolean if a field has been set.

### GetExpression

`func (o *ExpressionEvaluatorRequest) GetExpression() string`

GetExpression returns the Expression field if non-nil, zero value otherwise.

### GetExpressionOk

`func (o *ExpressionEvaluatorRequest) GetExpressionOk() (*string, bool)`

GetExpressionOk returns a tuple with the Expression field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetExpression

`func (o *ExpressionEvaluatorRequest) SetExpression(v string)`

SetExpression sets Expression field to given value.

### HasExpression

`func (o *ExpressionEvaluatorRequest) HasExpression() bool`

HasExpression returns a boolean if a field has been set.

### GetExpressionMap

`func (o *ExpressionEvaluatorRequest) GetExpressionMap() map[string]interface{}`

GetExpressionMap returns the ExpressionMap field if non-nil, zero value otherwise.

### GetExpressionMapOk

`func (o *ExpressionEvaluatorRequest) GetExpressionMapOk() (*map[string]interface{}, bool)`

GetExpressionMapOk returns a tuple with the ExpressionMap field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetExpressionMap

`func (o *ExpressionEvaluatorRequest) SetExpressionMap(v map[string]interface{})`

SetExpressionMap sets ExpressionMap field to given value.

### HasExpressionMap

`func (o *ExpressionEvaluatorRequest) HasExpressionMap() bool`

HasExpressionMap returns a boolean if a field has been set.

### GetInstanceID

`func (o *ExpressionEvaluatorRequest) GetInstanceID() string`

GetInstanceID returns the InstanceID field if non-nil, zero value otherwise.

### GetInstanceIDOk

`func (o *ExpressionEvaluatorRequest) GetInstanceIDOk() (*string, bool)`

GetInstanceIDOk returns a tuple with the InstanceID field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetInstanceID

`func (o *ExpressionEvaluatorRequest) SetInstanceID(v string)`

SetInstanceID sets InstanceID field to given value.

### HasInstanceID

`func (o *ExpressionEvaluatorRequest) HasInstanceID() bool`

HasInstanceID returns a boolean if a field has been set.

### GetProductTierID

`func (o *ExpressionEvaluatorRequest) GetProductTierID() string`

GetProductTierID returns the ProductTierID field if non-nil, zero value otherwise.

### GetProductTierIDOk

`func (o *ExpressionEvaluatorRequest) GetProductTierIDOk() (*string, bool)`

GetProductTierIDOk returns a tuple with the ProductTierID field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetProductTierID

`func (o *ExpressionEvaluatorRequest) SetProductTierID(v string)`

SetProductTierID sets ProductTierID field to given value.

### HasProductTierID

`func (o *ExpressionEvaluatorRequest) HasProductTierID() bool`

HasProductTierID returns a boolean if a field has been set.

### GetResourceKey

`func (o *ExpressionEvaluatorRequest) GetResourceKey() string`

GetResourceKey returns the ResourceKey field if non-nil, zero value otherwise.

### GetResourceKeyOk

`func (o *ExpressionEvaluatorRequest) GetResourceKeyOk() (*string, bool)`

GetResourceKeyOk returns a tuple with the ResourceKey field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetResourceKey

`func (o *ExpressionEvaluatorRequest) SetResourceKey(v string)`

SetResourceKey sets ResourceKey field to given value.


### GetServiceID

`func (o *ExpressionEvaluatorRequest) GetServiceID() string`

GetServiceID returns the ServiceID field if non-nil, zero value otherwise.

### GetServiceIDOk

`func (o *ExpressionEvaluatorRequest) GetServiceIDOk() (*string, bool)`

GetServiceIDOk returns a tuple with the ServiceID field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetServiceID

`func (o *ExpressionEvaluatorRequest) SetServiceID(v string)`

SetServiceID sets ServiceID field to given value.


### GetToken

`func (o *ExpressionEvaluatorRequest) GetToken() string`

GetToken returns the Token field if non-nil, zero value otherwise.

### GetTokenOk

`func (o *ExpressionEvaluatorRequest) GetTokenOk() (*string, bool)`

GetTokenOk returns a tuple with the Token field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetToken

`func (o *ExpressionEvaluatorRequest) SetToken(v string)`

SetToken sets Token field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


