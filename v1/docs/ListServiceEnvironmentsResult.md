# ListServiceEnvironmentsResult

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Ids** | **[]string** | The list of service environment IDs | 
**NextPageToken** | Pointer to **string** | The next token to use to retrieve the next page of results | [optional] 

## Methods

### NewListServiceEnvironmentsResult

`func NewListServiceEnvironmentsResult(ids []string, ) *ListServiceEnvironmentsResult`

NewListServiceEnvironmentsResult instantiates a new ListServiceEnvironmentsResult object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewListServiceEnvironmentsResultWithDefaults

`func NewListServiceEnvironmentsResultWithDefaults() *ListServiceEnvironmentsResult`

NewListServiceEnvironmentsResultWithDefaults instantiates a new ListServiceEnvironmentsResult object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetIds

`func (o *ListServiceEnvironmentsResult) GetIds() []string`

GetIds returns the Ids field if non-nil, zero value otherwise.

### GetIdsOk

`func (o *ListServiceEnvironmentsResult) GetIdsOk() (*[]string, bool)`

GetIdsOk returns a tuple with the Ids field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIds

`func (o *ListServiceEnvironmentsResult) SetIds(v []string)`

SetIds sets Ids field to given value.


### GetNextPageToken

`func (o *ListServiceEnvironmentsResult) GetNextPageToken() string`

GetNextPageToken returns the NextPageToken field if non-nil, zero value otherwise.

### GetNextPageTokenOk

`func (o *ListServiceEnvironmentsResult) GetNextPageTokenOk() (*string, bool)`

GetNextPageTokenOk returns a tuple with the NextPageToken field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNextPageToken

`func (o *ListServiceEnvironmentsResult) SetNextPageToken(v string)`

SetNextPageToken sets NextPageToken field to given value.

### HasNextPageToken

`func (o *ListServiceEnvironmentsResult) HasNextPageToken() bool`

HasNextPageToken returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


