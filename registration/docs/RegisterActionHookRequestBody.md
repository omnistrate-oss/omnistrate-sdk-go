# RegisterActionHookRequestBody

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**CommandTemplate** | **string** | The Base64 encoded command template to execute | 
**Scope** | **string** | The scope of the hook | 
**Type** | **string** | The type of hook to execute | 

## Methods

### NewRegisterActionHookRequestBody

`func NewRegisterActionHookRequestBody(commandTemplate string, scope string, type_ string, ) *RegisterActionHookRequestBody`

NewRegisterActionHookRequestBody instantiates a new RegisterActionHookRequestBody object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewRegisterActionHookRequestBodyWithDefaults

`func NewRegisterActionHookRequestBodyWithDefaults() *RegisterActionHookRequestBody`

NewRegisterActionHookRequestBodyWithDefaults instantiates a new RegisterActionHookRequestBody object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetCommandTemplate

`func (o *RegisterActionHookRequestBody) GetCommandTemplate() string`

GetCommandTemplate returns the CommandTemplate field if non-nil, zero value otherwise.

### GetCommandTemplateOk

`func (o *RegisterActionHookRequestBody) GetCommandTemplateOk() (*string, bool)`

GetCommandTemplateOk returns a tuple with the CommandTemplate field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCommandTemplate

`func (o *RegisterActionHookRequestBody) SetCommandTemplate(v string)`

SetCommandTemplate sets CommandTemplate field to given value.


### GetScope

`func (o *RegisterActionHookRequestBody) GetScope() string`

GetScope returns the Scope field if non-nil, zero value otherwise.

### GetScopeOk

`func (o *RegisterActionHookRequestBody) GetScopeOk() (*string, bool)`

GetScopeOk returns a tuple with the Scope field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetScope

`func (o *RegisterActionHookRequestBody) SetScope(v string)`

SetScope sets Scope field to given value.


### GetType

`func (o *RegisterActionHookRequestBody) GetType() string`

GetType returns the Type field if non-nil, zero value otherwise.

### GetTypeOk

`func (o *RegisterActionHookRequestBody) GetTypeOk() (*string, bool)`

GetTypeOk returns a tuple with the Type field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetType

`func (o *RegisterActionHookRequestBody) SetType(v string)`

SetType sets Type field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


