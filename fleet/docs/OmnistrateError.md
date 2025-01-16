# OmnistrateError

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Code** | **string** | Application-specific error code | 
**Message** | **string** | More context on the error that occurred | 
**Name** | **string** | Name of the error | 

## Methods

### NewOmnistrateError

`func NewOmnistrateError(code string, message string, name string, ) *OmnistrateError`

NewOmnistrateError instantiates a new OmnistrateError object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewOmnistrateErrorWithDefaults

`func NewOmnistrateErrorWithDefaults() *OmnistrateError`

NewOmnistrateErrorWithDefaults instantiates a new OmnistrateError object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetCode

`func (o *OmnistrateError) GetCode() string`

GetCode returns the Code field if non-nil, zero value otherwise.

### GetCodeOk

`func (o *OmnistrateError) GetCodeOk() (*string, bool)`

GetCodeOk returns a tuple with the Code field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCode

`func (o *OmnistrateError) SetCode(v string)`

SetCode sets Code field to given value.


### GetMessage

`func (o *OmnistrateError) GetMessage() string`

GetMessage returns the Message field if non-nil, zero value otherwise.

### GetMessageOk

`func (o *OmnistrateError) GetMessageOk() (*string, bool)`

GetMessageOk returns a tuple with the Message field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMessage

`func (o *OmnistrateError) SetMessage(v string)`

SetMessage sets Message field to given value.


### GetName

`func (o *OmnistrateError) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *OmnistrateError) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *OmnistrateError) SetName(v string)`

SetName sets Name field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


