# DescribeAccessRequestRequest

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Id** | **string** | Server-generated ID of a provider cluster access request. | 
**Token** | **string** | JWT token used to perform authorization | 

## Methods

### NewDescribeAccessRequestRequest

`func NewDescribeAccessRequestRequest(id string, token string, ) *DescribeAccessRequestRequest`

NewDescribeAccessRequestRequest instantiates a new DescribeAccessRequestRequest object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewDescribeAccessRequestRequestWithDefaults

`func NewDescribeAccessRequestRequestWithDefaults() *DescribeAccessRequestRequest`

NewDescribeAccessRequestRequestWithDefaults instantiates a new DescribeAccessRequestRequest object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetId

`func (o *DescribeAccessRequestRequest) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *DescribeAccessRequestRequest) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *DescribeAccessRequestRequest) SetId(v string)`

SetId sets Id field to given value.


### GetToken

`func (o *DescribeAccessRequestRequest) GetToken() string`

GetToken returns the Token field if non-nil, zero value otherwise.

### GetTokenOk

`func (o *DescribeAccessRequestRequest) GetTokenOk() (*string, bool)`

GetTokenOk returns a tuple with the Token field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetToken

`func (o *DescribeAccessRequestRequest) SetToken(v string)`

SetToken sets Token field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


