# DescribePipelineRequest

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Id** | Pointer to **string** | ID of a Pipeline | [optional] 
**ServiceId** | Pointer to **string** | ID of a Service | [optional] 
**Token** | **string** | JWT token used to perform authorization | 

## Methods

### NewDescribePipelineRequest

`func NewDescribePipelineRequest(token string, ) *DescribePipelineRequest`

NewDescribePipelineRequest instantiates a new DescribePipelineRequest object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewDescribePipelineRequestWithDefaults

`func NewDescribePipelineRequestWithDefaults() *DescribePipelineRequest`

NewDescribePipelineRequestWithDefaults instantiates a new DescribePipelineRequest object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetId

`func (o *DescribePipelineRequest) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *DescribePipelineRequest) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *DescribePipelineRequest) SetId(v string)`

SetId sets Id field to given value.

### HasId

`func (o *DescribePipelineRequest) HasId() bool`

HasId returns a boolean if a field has been set.

### GetServiceId

`func (o *DescribePipelineRequest) GetServiceId() string`

GetServiceId returns the ServiceId field if non-nil, zero value otherwise.

### GetServiceIdOk

`func (o *DescribePipelineRequest) GetServiceIdOk() (*string, bool)`

GetServiceIdOk returns a tuple with the ServiceId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetServiceId

`func (o *DescribePipelineRequest) SetServiceId(v string)`

SetServiceId sets ServiceId field to given value.

### HasServiceId

`func (o *DescribePipelineRequest) HasServiceId() bool`

HasServiceId returns a boolean if a field has been set.

### GetToken

`func (o *DescribePipelineRequest) GetToken() string`

GetToken returns the Token field if non-nil, zero value otherwise.

### GetTokenOk

`func (o *DescribePipelineRequest) GetTokenOk() (*string, bool)`

GetTokenOk returns a tuple with the Token field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetToken

`func (o *DescribePipelineRequest) SetToken(v string)`

SetToken sets Token field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


