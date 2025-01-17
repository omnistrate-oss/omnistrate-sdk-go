# SearchInventoryRequest

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Query** | **string** | The search query. | 
**Token** | **string** | JWT token used to perform authorization | 

## Methods

### NewSearchInventoryRequest

`func NewSearchInventoryRequest(query string, token string, ) *SearchInventoryRequest`

NewSearchInventoryRequest instantiates a new SearchInventoryRequest object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewSearchInventoryRequestWithDefaults

`func NewSearchInventoryRequestWithDefaults() *SearchInventoryRequest`

NewSearchInventoryRequestWithDefaults instantiates a new SearchInventoryRequest object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetQuery

`func (o *SearchInventoryRequest) GetQuery() string`

GetQuery returns the Query field if non-nil, zero value otherwise.

### GetQueryOk

`func (o *SearchInventoryRequest) GetQueryOk() (*string, bool)`

GetQueryOk returns a tuple with the Query field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetQuery

`func (o *SearchInventoryRequest) SetQuery(v string)`

SetQuery sets Query field to given value.


### GetToken

`func (o *SearchInventoryRequest) GetToken() string`

GetToken returns the Token field if non-nil, zero value otherwise.

### GetTokenOk

`func (o *SearchInventoryRequest) GetTokenOk() (*string, bool)`

GetTokenOk returns a tuple with the Token field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetToken

`func (o *SearchInventoryRequest) SetToken(v string)`

SetToken sets Token field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


