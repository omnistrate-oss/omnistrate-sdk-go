# ListAccountConfigResult

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**AccountConfigs** | Pointer to [**[]DescribeAccountConfigResult**](DescribeAccountConfigResult.md) | The list of account configs | [optional] 
**Ids** | Pointer to **[]string** |  | [optional] 
**NextPageToken** | Pointer to **string** | Token to use for the next page | [optional] 

## Methods

### NewListAccountConfigResult

`func NewListAccountConfigResult() *ListAccountConfigResult`

NewListAccountConfigResult instantiates a new ListAccountConfigResult object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewListAccountConfigResultWithDefaults

`func NewListAccountConfigResultWithDefaults() *ListAccountConfigResult`

NewListAccountConfigResultWithDefaults instantiates a new ListAccountConfigResult object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetAccountConfigs

`func (o *ListAccountConfigResult) GetAccountConfigs() []DescribeAccountConfigResult`

GetAccountConfigs returns the AccountConfigs field if non-nil, zero value otherwise.

### GetAccountConfigsOk

`func (o *ListAccountConfigResult) GetAccountConfigsOk() (*[]DescribeAccountConfigResult, bool)`

GetAccountConfigsOk returns a tuple with the AccountConfigs field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAccountConfigs

`func (o *ListAccountConfigResult) SetAccountConfigs(v []DescribeAccountConfigResult)`

SetAccountConfigs sets AccountConfigs field to given value.

### HasAccountConfigs

`func (o *ListAccountConfigResult) HasAccountConfigs() bool`

HasAccountConfigs returns a boolean if a field has been set.

### GetIds

`func (o *ListAccountConfigResult) GetIds() []string`

GetIds returns the Ids field if non-nil, zero value otherwise.

### GetIdsOk

`func (o *ListAccountConfigResult) GetIdsOk() (*[]string, bool)`

GetIdsOk returns a tuple with the Ids field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIds

`func (o *ListAccountConfigResult) SetIds(v []string)`

SetIds sets Ids field to given value.

### HasIds

`func (o *ListAccountConfigResult) HasIds() bool`

HasIds returns a boolean if a field has been set.

### GetNextPageToken

`func (o *ListAccountConfigResult) GetNextPageToken() string`

GetNextPageToken returns the NextPageToken field if non-nil, zero value otherwise.

### GetNextPageTokenOk

`func (o *ListAccountConfigResult) GetNextPageTokenOk() (*string, bool)`

GetNextPageTokenOk returns a tuple with the NextPageToken field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNextPageToken

`func (o *ListAccountConfigResult) SetNextPageToken(v string)`

SetNextPageToken sets NextPageToken field to given value.

### HasNextPageToken

`func (o *ListAccountConfigResult) HasNextPageToken() bool`

HasNextPageToken returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


