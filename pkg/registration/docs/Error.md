# Error

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Fault** | **bool** | Is the error a server-side fault? | 
**Id** | **string** | ID is a unique identifier for this particular occurrence of the problem. | 
**Message** | **string** | Message is a human-readable explanation specific to this occurrence of the problem. | 
**Name** | **string** | Name is the name of this class of errors. | 
**Temporary** | **bool** | Is the error temporary? | 
**Timeout** | **bool** | Is the error a timeout? | 

## Methods

### NewError

`func NewError(fault bool, id string, message string, name string, temporary bool, timeout bool, ) *Error`

NewError instantiates a new Error object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewErrorWithDefaults

`func NewErrorWithDefaults() *Error`

NewErrorWithDefaults instantiates a new Error object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetFault

`func (o *Error) GetFault() bool`

GetFault returns the Fault field if non-nil, zero value otherwise.

### GetFaultOk

`func (o *Error) GetFaultOk() (*bool, bool)`

GetFaultOk returns a tuple with the Fault field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFault

`func (o *Error) SetFault(v bool)`

SetFault sets Fault field to given value.


### GetId

`func (o *Error) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *Error) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *Error) SetId(v string)`

SetId sets Id field to given value.


### GetMessage

`func (o *Error) GetMessage() string`

GetMessage returns the Message field if non-nil, zero value otherwise.

### GetMessageOk

`func (o *Error) GetMessageOk() (*string, bool)`

GetMessageOk returns a tuple with the Message field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMessage

`func (o *Error) SetMessage(v string)`

SetMessage sets Message field to given value.


### GetName

`func (o *Error) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *Error) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *Error) SetName(v string)`

SetName sets Name field to given value.


### GetTemporary

`func (o *Error) GetTemporary() bool`

GetTemporary returns the Temporary field if non-nil, zero value otherwise.

### GetTemporaryOk

`func (o *Error) GetTemporaryOk() (*bool, bool)`

GetTemporaryOk returns a tuple with the Temporary field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTemporary

`func (o *Error) SetTemporary(v bool)`

SetTemporary sets Temporary field to given value.


### GetTimeout

`func (o *Error) GetTimeout() bool`

GetTimeout returns the Timeout field if non-nil, zero value otherwise.

### GetTimeoutOk

`func (o *Error) GetTimeoutOk() (*bool, bool)`

GetTimeoutOk returns a tuple with the Timeout field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTimeout

`func (o *Error) SetTimeout(v bool)`

SetTimeout sets Timeout field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


