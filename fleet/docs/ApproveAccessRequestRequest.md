# ApproveAccessRequestRequest

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**DurationMinutes** | Pointer to **int64** | MANUAL approval duration. Omission uses the finite backend default within 1–720; no client default, unlimited choice, active extension, or credential renewal. | [optional] 
**Id** | **string** | Server-generated ID of a provider cluster access request. | 
**Token** | **string** | JWT token used to perform authorization | 

## Methods

### NewApproveAccessRequestRequest

`func NewApproveAccessRequestRequest(id string, token string, ) *ApproveAccessRequestRequest`

NewApproveAccessRequestRequest instantiates a new ApproveAccessRequestRequest object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewApproveAccessRequestRequestWithDefaults

`func NewApproveAccessRequestRequestWithDefaults() *ApproveAccessRequestRequest`

NewApproveAccessRequestRequestWithDefaults instantiates a new ApproveAccessRequestRequest object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetDurationMinutes

`func (o *ApproveAccessRequestRequest) GetDurationMinutes() int64`

GetDurationMinutes returns the DurationMinutes field if non-nil, zero value otherwise.

### GetDurationMinutesOk

`func (o *ApproveAccessRequestRequest) GetDurationMinutesOk() (*int64, bool)`

GetDurationMinutesOk returns a tuple with the DurationMinutes field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDurationMinutes

`func (o *ApproveAccessRequestRequest) SetDurationMinutes(v int64)`

SetDurationMinutes sets DurationMinutes field to given value.

### HasDurationMinutes

`func (o *ApproveAccessRequestRequest) HasDurationMinutes() bool`

HasDurationMinutes returns a boolean if a field has been set.

### GetId

`func (o *ApproveAccessRequestRequest) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *ApproveAccessRequestRequest) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *ApproveAccessRequestRequest) SetId(v string)`

SetId sets Id field to given value.


### GetToken

`func (o *ApproveAccessRequestRequest) GetToken() string`

GetToken returns the Token field if non-nil, zero value otherwise.

### GetTokenOk

`func (o *ApproveAccessRequestRequest) GetTokenOk() (*string, bool)`

GetTokenOk returns a tuple with the Token field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetToken

`func (o *ApproveAccessRequestRequest) SetToken(v string)`

SetToken sets Token field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


