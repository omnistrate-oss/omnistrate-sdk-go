# RegisterActionHookRequest2

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**CommandTemplate** | **string** | The Base64 encoded command template to execute | 
**CustomCommand** | Pointer to **[]string** | The custom command to execute the hook with | [optional] 
**CustomImage** | Pointer to **string** | The custom image to execute the hook | [optional] 
**Scope** | **string** | The scope of the hook | 
**Type** | **string** | The type of hook to execute | 

## Methods

### NewRegisterActionHookRequest2

`func NewRegisterActionHookRequest2(commandTemplate string, scope string, type_ string, ) *RegisterActionHookRequest2`

NewRegisterActionHookRequest2 instantiates a new RegisterActionHookRequest2 object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewRegisterActionHookRequest2WithDefaults

`func NewRegisterActionHookRequest2WithDefaults() *RegisterActionHookRequest2`

NewRegisterActionHookRequest2WithDefaults instantiates a new RegisterActionHookRequest2 object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetCommandTemplate

`func (o *RegisterActionHookRequest2) GetCommandTemplate() string`

GetCommandTemplate returns the CommandTemplate field if non-nil, zero value otherwise.

### GetCommandTemplateOk

`func (o *RegisterActionHookRequest2) GetCommandTemplateOk() (*string, bool)`

GetCommandTemplateOk returns a tuple with the CommandTemplate field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCommandTemplate

`func (o *RegisterActionHookRequest2) SetCommandTemplate(v string)`

SetCommandTemplate sets CommandTemplate field to given value.


### GetCustomCommand

`func (o *RegisterActionHookRequest2) GetCustomCommand() []string`

GetCustomCommand returns the CustomCommand field if non-nil, zero value otherwise.

### GetCustomCommandOk

`func (o *RegisterActionHookRequest2) GetCustomCommandOk() (*[]string, bool)`

GetCustomCommandOk returns a tuple with the CustomCommand field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCustomCommand

`func (o *RegisterActionHookRequest2) SetCustomCommand(v []string)`

SetCustomCommand sets CustomCommand field to given value.

### HasCustomCommand

`func (o *RegisterActionHookRequest2) HasCustomCommand() bool`

HasCustomCommand returns a boolean if a field has been set.

### GetCustomImage

`func (o *RegisterActionHookRequest2) GetCustomImage() string`

GetCustomImage returns the CustomImage field if non-nil, zero value otherwise.

### GetCustomImageOk

`func (o *RegisterActionHookRequest2) GetCustomImageOk() (*string, bool)`

GetCustomImageOk returns a tuple with the CustomImage field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCustomImage

`func (o *RegisterActionHookRequest2) SetCustomImage(v string)`

SetCustomImage sets CustomImage field to given value.

### HasCustomImage

`func (o *RegisterActionHookRequest2) HasCustomImage() bool`

HasCustomImage returns a boolean if a field has been set.

### GetScope

`func (o *RegisterActionHookRequest2) GetScope() string`

GetScope returns the Scope field if non-nil, zero value otherwise.

### GetScopeOk

`func (o *RegisterActionHookRequest2) GetScopeOk() (*string, bool)`

GetScopeOk returns a tuple with the Scope field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetScope

`func (o *RegisterActionHookRequest2) SetScope(v string)`

SetScope sets Scope field to given value.


### GetType

`func (o *RegisterActionHookRequest2) GetType() string`

GetType returns the Type field if non-nil, zero value otherwise.

### GetTypeOk

`func (o *RegisterActionHookRequest2) GetTypeOk() (*string, bool)`

GetTypeOk returns a tuple with the Type field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetType

`func (o *RegisterActionHookRequest2) SetType(v string)`

SetType sets Type field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


