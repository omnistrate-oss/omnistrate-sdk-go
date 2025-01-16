# ListUpgradePathResult

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Ids** | Pointer to **[]string** | List of upgrade path IDs | [optional] 
**NextPageToken** | Pointer to **string** | Token to use to get the next page of results | [optional] 

## Methods

### NewListUpgradePathResult

`func NewListUpgradePathResult() *ListUpgradePathResult`

NewListUpgradePathResult instantiates a new ListUpgradePathResult object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewListUpgradePathResultWithDefaults

`func NewListUpgradePathResultWithDefaults() *ListUpgradePathResult`

NewListUpgradePathResultWithDefaults instantiates a new ListUpgradePathResult object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetIds

`func (o *ListUpgradePathResult) GetIds() []string`

GetIds returns the Ids field if non-nil, zero value otherwise.

### GetIdsOk

`func (o *ListUpgradePathResult) GetIdsOk() (*[]string, bool)`

GetIdsOk returns a tuple with the Ids field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIds

`func (o *ListUpgradePathResult) SetIds(v []string)`

SetIds sets Ids field to given value.

### HasIds

`func (o *ListUpgradePathResult) HasIds() bool`

HasIds returns a boolean if a field has been set.

### GetNextPageToken

`func (o *ListUpgradePathResult) GetNextPageToken() string`

GetNextPageToken returns the NextPageToken field if non-nil, zero value otherwise.

### GetNextPageTokenOk

`func (o *ListUpgradePathResult) GetNextPageTokenOk() (*string, bool)`

GetNextPageTokenOk returns a tuple with the NextPageToken field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNextPageToken

`func (o *ListUpgradePathResult) SetNextPageToken(v string)`

SetNextPageToken sets NextPageToken field to given value.

### HasNextPageToken

`func (o *ListUpgradePathResult) HasNextPageToken() bool`

HasNextPageToken returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


