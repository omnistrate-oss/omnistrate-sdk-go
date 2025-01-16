# RegisterActionHookRequest

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**CommandTemplate** | **string** | The Base64 encoded command template to execute | 
**CustomCommand** | Pointer to **[]string** | The custom command to execute the hook with | [optional] 
**CustomImage** | Pointer to **string** | The custom image to execute the hook | [optional] 
**Id** | **string** | ID of a resource | 
**Scope** | **string** | The scope of the hook | 
**ServiceId** | **string** | ID of a Service | 
**Token** | **string** | JWT token used to perform authorization | 
**Type** | **string** | The type of hook to execute | 

## Methods

### NewRegisterActionHookRequest

`func NewRegisterActionHookRequest(commandTemplate string, id string, scope string, serviceId string, token string, type_ string, ) *RegisterActionHookRequest`

NewRegisterActionHookRequest instantiates a new RegisterActionHookRequest object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewRegisterActionHookRequestWithDefaults

`func NewRegisterActionHookRequestWithDefaults() *RegisterActionHookRequest`

NewRegisterActionHookRequestWithDefaults instantiates a new RegisterActionHookRequest object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetCommandTemplate

`func (o *RegisterActionHookRequest) GetCommandTemplate() string`

GetCommandTemplate returns the CommandTemplate field if non-nil, zero value otherwise.

### GetCommandTemplateOk

`func (o *RegisterActionHookRequest) GetCommandTemplateOk() (*string, bool)`

GetCommandTemplateOk returns a tuple with the CommandTemplate field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCommandTemplate

`func (o *RegisterActionHookRequest) SetCommandTemplate(v string)`

SetCommandTemplate sets CommandTemplate field to given value.


### GetCustomCommand

`func (o *RegisterActionHookRequest) GetCustomCommand() []string`

GetCustomCommand returns the CustomCommand field if non-nil, zero value otherwise.

### GetCustomCommandOk

`func (o *RegisterActionHookRequest) GetCustomCommandOk() (*[]string, bool)`

GetCustomCommandOk returns a tuple with the CustomCommand field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCustomCommand

`func (o *RegisterActionHookRequest) SetCustomCommand(v []string)`

SetCustomCommand sets CustomCommand field to given value.

### HasCustomCommand

`func (o *RegisterActionHookRequest) HasCustomCommand() bool`

HasCustomCommand returns a boolean if a field has been set.

### GetCustomImage

`func (o *RegisterActionHookRequest) GetCustomImage() string`

GetCustomImage returns the CustomImage field if non-nil, zero value otherwise.

### GetCustomImageOk

`func (o *RegisterActionHookRequest) GetCustomImageOk() (*string, bool)`

GetCustomImageOk returns a tuple with the CustomImage field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCustomImage

`func (o *RegisterActionHookRequest) SetCustomImage(v string)`

SetCustomImage sets CustomImage field to given value.

### HasCustomImage

`func (o *RegisterActionHookRequest) HasCustomImage() bool`

HasCustomImage returns a boolean if a field has been set.

### GetId

`func (o *RegisterActionHookRequest) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *RegisterActionHookRequest) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *RegisterActionHookRequest) SetId(v string)`

SetId sets Id field to given value.


### GetScope

`func (o *RegisterActionHookRequest) GetScope() string`

GetScope returns the Scope field if non-nil, zero value otherwise.

### GetScopeOk

`func (o *RegisterActionHookRequest) GetScopeOk() (*string, bool)`

GetScopeOk returns a tuple with the Scope field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetScope

`func (o *RegisterActionHookRequest) SetScope(v string)`

SetScope sets Scope field to given value.


### GetServiceId

`func (o *RegisterActionHookRequest) GetServiceId() string`

GetServiceId returns the ServiceId field if non-nil, zero value otherwise.

### GetServiceIdOk

`func (o *RegisterActionHookRequest) GetServiceIdOk() (*string, bool)`

GetServiceIdOk returns a tuple with the ServiceId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetServiceId

`func (o *RegisterActionHookRequest) SetServiceId(v string)`

SetServiceId sets ServiceId field to given value.


### GetToken

`func (o *RegisterActionHookRequest) GetToken() string`

GetToken returns the Token field if non-nil, zero value otherwise.

### GetTokenOk

`func (o *RegisterActionHookRequest) GetTokenOk() (*string, bool)`

GetTokenOk returns a tuple with the Token field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetToken

`func (o *RegisterActionHookRequest) SetToken(v string)`

SetToken sets Token field to given value.


### GetType

`func (o *RegisterActionHookRequest) GetType() string`

GetType returns the Type field if non-nil, zero value otherwise.

### GetTypeOk

`func (o *RegisterActionHookRequest) GetTypeOk() (*string, bool)`

GetTypeOk returns a tuple with the Type field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetType

`func (o *RegisterActionHookRequest) SetType(v string)`

SetType sets Type field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


