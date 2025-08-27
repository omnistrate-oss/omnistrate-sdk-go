# ExpressionEvaluatorResult

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Error** | Pointer to **string** | Error message if the evaluation failed | [optional] 
**Result** | Pointer to **interface{}** | The evaluated result of the expression | [optional] 

## Methods

### NewExpressionEvaluatorResult

`func NewExpressionEvaluatorResult() *ExpressionEvaluatorResult`

NewExpressionEvaluatorResult instantiates a new ExpressionEvaluatorResult object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewExpressionEvaluatorResultWithDefaults

`func NewExpressionEvaluatorResultWithDefaults() *ExpressionEvaluatorResult`

NewExpressionEvaluatorResultWithDefaults instantiates a new ExpressionEvaluatorResult object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetError

`func (o *ExpressionEvaluatorResult) GetError() string`

GetError returns the Error field if non-nil, zero value otherwise.

### GetErrorOk

`func (o *ExpressionEvaluatorResult) GetErrorOk() (*string, bool)`

GetErrorOk returns a tuple with the Error field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetError

`func (o *ExpressionEvaluatorResult) SetError(v string)`

SetError sets Error field to given value.

### HasError

`func (o *ExpressionEvaluatorResult) HasError() bool`

HasError returns a boolean if a field has been set.

### GetResult

`func (o *ExpressionEvaluatorResult) GetResult() interface{}`

GetResult returns the Result field if non-nil, zero value otherwise.

### GetResultOk

`func (o *ExpressionEvaluatorResult) GetResultOk() (*interface{}, bool)`

GetResultOk returns a tuple with the Result field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetResult

`func (o *ExpressionEvaluatorResult) SetResult(v interface{})`

SetResult sets Result field to given value.

### HasResult

`func (o *ExpressionEvaluatorResult) HasResult() bool`

HasResult returns a boolean if a field has been set.

### SetResultNil

`func (o *ExpressionEvaluatorResult) SetResultNil(b bool)`

 SetResultNil sets the value for Result to be an explicit nil

### UnsetResult
`func (o *ExpressionEvaluatorResult) UnsetResult()`

UnsetResult ensures that no value is present for Result, not even an explicit nil

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


