# ListActionHooksResult

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Hooks** | [**[]RegisterActionHookRequestBody**](RegisterActionHookRequestBody.md) | The list of action hooks | 

## Methods

### NewListActionHooksResult

`func NewListActionHooksResult(hooks []RegisterActionHookRequestBody, ) *ListActionHooksResult`

NewListActionHooksResult instantiates a new ListActionHooksResult object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewListActionHooksResultWithDefaults

`func NewListActionHooksResultWithDefaults() *ListActionHooksResult`

NewListActionHooksResultWithDefaults instantiates a new ListActionHooksResult object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetHooks

`func (o *ListActionHooksResult) GetHooks() []RegisterActionHookRequestBody`

GetHooks returns the Hooks field if non-nil, zero value otherwise.

### GetHooksOk

`func (o *ListActionHooksResult) GetHooksOk() (*[]RegisterActionHookRequestBody, bool)`

GetHooksOk returns a tuple with the Hooks field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetHooks

`func (o *ListActionHooksResult) SetHooks(v []RegisterActionHookRequestBody)`

SetHooks sets Hooks field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


