# ListVUnitsResult

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Ids** | Pointer to **map[string][]string** | List of VUnit IDs per Cloud provider | [optional] 
**NextPageToken** | Pointer to **string** | Next page token | [optional] 

## Methods

### NewListVUnitsResult

`func NewListVUnitsResult() *ListVUnitsResult`

NewListVUnitsResult instantiates a new ListVUnitsResult object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewListVUnitsResultWithDefaults

`func NewListVUnitsResultWithDefaults() *ListVUnitsResult`

NewListVUnitsResultWithDefaults instantiates a new ListVUnitsResult object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetIds

`func (o *ListVUnitsResult) GetIds() map[string][]string`

GetIds returns the Ids field if non-nil, zero value otherwise.

### GetIdsOk

`func (o *ListVUnitsResult) GetIdsOk() (*map[string][]string, bool)`

GetIdsOk returns a tuple with the Ids field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIds

`func (o *ListVUnitsResult) SetIds(v map[string][]string)`

SetIds sets Ids field to given value.

### HasIds

`func (o *ListVUnitsResult) HasIds() bool`

HasIds returns a boolean if a field has been set.

### GetNextPageToken

`func (o *ListVUnitsResult) GetNextPageToken() string`

GetNextPageToken returns the NextPageToken field if non-nil, zero value otherwise.

### GetNextPageTokenOk

`func (o *ListVUnitsResult) GetNextPageTokenOk() (*string, bool)`

GetNextPageTokenOk returns a tuple with the NextPageToken field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNextPageToken

`func (o *ListVUnitsResult) SetNextPageToken(v string)`

SetNextPageToken sets NextPageToken field to given value.

### HasNextPageToken

`func (o *ListVUnitsResult) HasNextPageToken() bool`

HasNextPageToken returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


