# DeleteLimitRequest

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Family** | **string** | The limit family | 
**Key** | **string** | Unique key to identify the limit | 
**Token** | **string** | JWT token used to perform authorization | 

## Methods

### NewDeleteLimitRequest

`func NewDeleteLimitRequest(family string, key string, token string, ) *DeleteLimitRequest`

NewDeleteLimitRequest instantiates a new DeleteLimitRequest object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewDeleteLimitRequestWithDefaults

`func NewDeleteLimitRequestWithDefaults() *DeleteLimitRequest`

NewDeleteLimitRequestWithDefaults instantiates a new DeleteLimitRequest object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetFamily

`func (o *DeleteLimitRequest) GetFamily() string`

GetFamily returns the Family field if non-nil, zero value otherwise.

### GetFamilyOk

`func (o *DeleteLimitRequest) GetFamilyOk() (*string, bool)`

GetFamilyOk returns a tuple with the Family field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFamily

`func (o *DeleteLimitRequest) SetFamily(v string)`

SetFamily sets Family field to given value.


### GetKey

`func (o *DeleteLimitRequest) GetKey() string`

GetKey returns the Key field if non-nil, zero value otherwise.

### GetKeyOk

`func (o *DeleteLimitRequest) GetKeyOk() (*string, bool)`

GetKeyOk returns a tuple with the Key field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetKey

`func (o *DeleteLimitRequest) SetKey(v string)`

SetKey sets Key field to given value.


### GetToken

`func (o *DeleteLimitRequest) GetToken() string`

GetToken returns the Token field if non-nil, zero value otherwise.

### GetTokenOk

`func (o *DeleteLimitRequest) GetTokenOk() (*string, bool)`

GetTokenOk returns a tuple with the Token field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetToken

`func (o *DeleteLimitRequest) SetToken(v string)`

SetToken sets Token field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


