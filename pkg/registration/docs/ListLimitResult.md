# ListLimitResult

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Keys** | **[]string** | List of limit keys | 
**NextPageToken** | Pointer to **string** | Token to use for the next page | [optional] 

## Methods

### NewListLimitResult

`func NewListLimitResult(keys []string, ) *ListLimitResult`

NewListLimitResult instantiates a new ListLimitResult object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewListLimitResultWithDefaults

`func NewListLimitResultWithDefaults() *ListLimitResult`

NewListLimitResultWithDefaults instantiates a new ListLimitResult object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetKeys

`func (o *ListLimitResult) GetKeys() []string`

GetKeys returns the Keys field if non-nil, zero value otherwise.

### GetKeysOk

`func (o *ListLimitResult) GetKeysOk() (*[]string, bool)`

GetKeysOk returns a tuple with the Keys field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetKeys

`func (o *ListLimitResult) SetKeys(v []string)`

SetKeys sets Keys field to given value.


### GetNextPageToken

`func (o *ListLimitResult) GetNextPageToken() string`

GetNextPageToken returns the NextPageToken field if non-nil, zero value otherwise.

### GetNextPageTokenOk

`func (o *ListLimitResult) GetNextPageTokenOk() (*string, bool)`

GetNextPageTokenOk returns a tuple with the NextPageToken field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNextPageToken

`func (o *ListLimitResult) SetNextPageToken(v string)`

SetNextPageToken sets NextPageToken field to given value.

### HasNextPageToken

`func (o *ListLimitResult) HasNextPageToken() bool`

HasNextPageToken returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


