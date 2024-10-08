# ListOutputParametersResult

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Ids** | **[]string** | The IDs of the output parameters | 
**NextPageToken** | Pointer to **string** | The token to use for the next page of results | [optional] 
**OutputParameters** | Pointer to [**[]DescribeOutputParameterResult**](DescribeOutputParameterResult.md) | The output parameters | [optional] 

## Methods

### NewListOutputParametersResult

`func NewListOutputParametersResult(ids []string, ) *ListOutputParametersResult`

NewListOutputParametersResult instantiates a new ListOutputParametersResult object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewListOutputParametersResultWithDefaults

`func NewListOutputParametersResultWithDefaults() *ListOutputParametersResult`

NewListOutputParametersResultWithDefaults instantiates a new ListOutputParametersResult object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetIds

`func (o *ListOutputParametersResult) GetIds() []string`

GetIds returns the Ids field if non-nil, zero value otherwise.

### GetIdsOk

`func (o *ListOutputParametersResult) GetIdsOk() (*[]string, bool)`

GetIdsOk returns a tuple with the Ids field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIds

`func (o *ListOutputParametersResult) SetIds(v []string)`

SetIds sets Ids field to given value.


### GetNextPageToken

`func (o *ListOutputParametersResult) GetNextPageToken() string`

GetNextPageToken returns the NextPageToken field if non-nil, zero value otherwise.

### GetNextPageTokenOk

`func (o *ListOutputParametersResult) GetNextPageTokenOk() (*string, bool)`

GetNextPageTokenOk returns a tuple with the NextPageToken field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNextPageToken

`func (o *ListOutputParametersResult) SetNextPageToken(v string)`

SetNextPageToken sets NextPageToken field to given value.

### HasNextPageToken

`func (o *ListOutputParametersResult) HasNextPageToken() bool`

HasNextPageToken returns a boolean if a field has been set.

### GetOutputParameters

`func (o *ListOutputParametersResult) GetOutputParameters() []DescribeOutputParameterResult`

GetOutputParameters returns the OutputParameters field if non-nil, zero value otherwise.

### GetOutputParametersOk

`func (o *ListOutputParametersResult) GetOutputParametersOk() (*[]DescribeOutputParameterResult, bool)`

GetOutputParametersOk returns a tuple with the OutputParameters field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOutputParameters

`func (o *ListOutputParametersResult) SetOutputParameters(v []DescribeOutputParameterResult)`

SetOutputParameters sets OutputParameters field to given value.

### HasOutputParameters

`func (o *ListOutputParametersResult) HasOutputParameters() bool`

HasOutputParameters returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


