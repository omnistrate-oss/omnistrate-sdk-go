# ListLimitRequest

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Family** | **string** | The limit family | 
**Token** | **string** | JWT token used to perform authorization | 

## Methods

### NewListLimitRequest

`func NewListLimitRequest(family string, token string, ) *ListLimitRequest`

NewListLimitRequest instantiates a new ListLimitRequest object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewListLimitRequestWithDefaults

`func NewListLimitRequestWithDefaults() *ListLimitRequest`

NewListLimitRequestWithDefaults instantiates a new ListLimitRequest object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetFamily

`func (o *ListLimitRequest) GetFamily() string`

GetFamily returns the Family field if non-nil, zero value otherwise.

### GetFamilyOk

`func (o *ListLimitRequest) GetFamilyOk() (*string, bool)`

GetFamilyOk returns a tuple with the Family field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFamily

`func (o *ListLimitRequest) SetFamily(v string)`

SetFamily sets Family field to given value.


### GetToken

`func (o *ListLimitRequest) GetToken() string`

GetToken returns the Token field if non-nil, zero value otherwise.

### GetTokenOk

`func (o *ListLimitRequest) GetTokenOk() (*string, bool)`

GetTokenOk returns a tuple with the Token field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetToken

`func (o *ListLimitRequest) SetToken(v string)`

SetToken sets Token field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


