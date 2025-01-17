# DeletePipelineRequest

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Id** | Pointer to **string** | ID of a Pipeline | [optional] 
**ServiceId** | Pointer to **string** | ID of a Service | [optional] 
**Token** | **string** | JWT token used to perform authorization | 

## Methods

### NewDeletePipelineRequest

`func NewDeletePipelineRequest(token string, ) *DeletePipelineRequest`

NewDeletePipelineRequest instantiates a new DeletePipelineRequest object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewDeletePipelineRequestWithDefaults

`func NewDeletePipelineRequestWithDefaults() *DeletePipelineRequest`

NewDeletePipelineRequestWithDefaults instantiates a new DeletePipelineRequest object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetId

`func (o *DeletePipelineRequest) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *DeletePipelineRequest) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *DeletePipelineRequest) SetId(v string)`

SetId sets Id field to given value.

### HasId

`func (o *DeletePipelineRequest) HasId() bool`

HasId returns a boolean if a field has been set.

### GetServiceId

`func (o *DeletePipelineRequest) GetServiceId() string`

GetServiceId returns the ServiceId field if non-nil, zero value otherwise.

### GetServiceIdOk

`func (o *DeletePipelineRequest) GetServiceIdOk() (*string, bool)`

GetServiceIdOk returns a tuple with the ServiceId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetServiceId

`func (o *DeletePipelineRequest) SetServiceId(v string)`

SetServiceId sets ServiceId field to given value.

### HasServiceId

`func (o *DeletePipelineRequest) HasServiceId() bool`

HasServiceId returns a boolean if a field has been set.

### GetToken

`func (o *DeletePipelineRequest) GetToken() string`

GetToken returns the Token field if non-nil, zero value otherwise.

### GetTokenOk

`func (o *DeletePipelineRequest) GetTokenOk() (*string, bool)`

GetTokenOk returns a tuple with the Token field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetToken

`func (o *DeletePipelineRequest) SetToken(v string)`

SetToken sets Token field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


