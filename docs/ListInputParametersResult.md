# ListInputParametersResult

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Ids** | **[]string** | List of input parameter IDs | 
**InputParameters** | Pointer to [**[]DescribeInputParameterResult**](DescribeInputParameterResult.md) | The input parameters | [optional] 
**NextPageToken** | Pointer to **string** | Token to retrieve the next page of results | [optional] 

## Methods

### NewListInputParametersResult

`func NewListInputParametersResult(ids []string, ) *ListInputParametersResult`

NewListInputParametersResult instantiates a new ListInputParametersResult object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewListInputParametersResultWithDefaults

`func NewListInputParametersResultWithDefaults() *ListInputParametersResult`

NewListInputParametersResultWithDefaults instantiates a new ListInputParametersResult object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetIds

`func (o *ListInputParametersResult) GetIds() []string`

GetIds returns the Ids field if non-nil, zero value otherwise.

### GetIdsOk

`func (o *ListInputParametersResult) GetIdsOk() (*[]string, bool)`

GetIdsOk returns a tuple with the Ids field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIds

`func (o *ListInputParametersResult) SetIds(v []string)`

SetIds sets Ids field to given value.


### GetInputParameters

`func (o *ListInputParametersResult) GetInputParameters() []DescribeInputParameterResult`

GetInputParameters returns the InputParameters field if non-nil, zero value otherwise.

### GetInputParametersOk

`func (o *ListInputParametersResult) GetInputParametersOk() (*[]DescribeInputParameterResult, bool)`

GetInputParametersOk returns a tuple with the InputParameters field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetInputParameters

`func (o *ListInputParametersResult) SetInputParameters(v []DescribeInputParameterResult)`

SetInputParameters sets InputParameters field to given value.

### HasInputParameters

`func (o *ListInputParametersResult) HasInputParameters() bool`

HasInputParameters returns a boolean if a field has been set.

### GetNextPageToken

`func (o *ListInputParametersResult) GetNextPageToken() string`

GetNextPageToken returns the NextPageToken field if non-nil, zero value otherwise.

### GetNextPageTokenOk

`func (o *ListInputParametersResult) GetNextPageTokenOk() (*string, bool)`

GetNextPageTokenOk returns a tuple with the NextPageToken field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNextPageToken

`func (o *ListInputParametersResult) SetNextPageToken(v string)`

SetNextPageToken sets NextPageToken field to given value.

### HasNextPageToken

`func (o *ListInputParametersResult) HasNextPageToken() bool`

HasNextPageToken returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


