# JITAccessEventRequester

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**UserEmail** | Pointer to **string** | The requester email captured when the event was written. | [optional] 
**UserId** | Pointer to **string** | ID of a User | [optional] 
**UserName** | Pointer to **string** | The requester name captured when the event was written. | [optional] 

## Methods

### NewJITAccessEventRequester

`func NewJITAccessEventRequester() *JITAccessEventRequester`

NewJITAccessEventRequester instantiates a new JITAccessEventRequester object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewJITAccessEventRequesterWithDefaults

`func NewJITAccessEventRequesterWithDefaults() *JITAccessEventRequester`

NewJITAccessEventRequesterWithDefaults instantiates a new JITAccessEventRequester object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetUserEmail

`func (o *JITAccessEventRequester) GetUserEmail() string`

GetUserEmail returns the UserEmail field if non-nil, zero value otherwise.

### GetUserEmailOk

`func (o *JITAccessEventRequester) GetUserEmailOk() (*string, bool)`

GetUserEmailOk returns a tuple with the UserEmail field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUserEmail

`func (o *JITAccessEventRequester) SetUserEmail(v string)`

SetUserEmail sets UserEmail field to given value.

### HasUserEmail

`func (o *JITAccessEventRequester) HasUserEmail() bool`

HasUserEmail returns a boolean if a field has been set.

### GetUserId

`func (o *JITAccessEventRequester) GetUserId() string`

GetUserId returns the UserId field if non-nil, zero value otherwise.

### GetUserIdOk

`func (o *JITAccessEventRequester) GetUserIdOk() (*string, bool)`

GetUserIdOk returns a tuple with the UserId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUserId

`func (o *JITAccessEventRequester) SetUserId(v string)`

SetUserId sets UserId field to given value.

### HasUserId

`func (o *JITAccessEventRequester) HasUserId() bool`

HasUserId returns a boolean if a field has been set.

### GetUserName

`func (o *JITAccessEventRequester) GetUserName() string`

GetUserName returns the UserName field if non-nil, zero value otherwise.

### GetUserNameOk

`func (o *JITAccessEventRequester) GetUserNameOk() (*string, bool)`

GetUserNameOk returns a tuple with the UserName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUserName

`func (o *JITAccessEventRequester) SetUserName(v string)`

SetUserName sets UserName field to given value.

### HasUserName

`func (o *JITAccessEventRequester) HasUserName() bool`

HasUserName returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


