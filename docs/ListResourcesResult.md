# ListResourcesResult

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Ids** | **[]string** | List of resource IDs | 
**NextPageToken** | Pointer to **string** | Token to use for the next page | [optional] 
**Resources** | Pointer to [**[]DescribeResourceResult**](DescribeResourceResult.md) | List of resources | [optional] 

## Methods

### NewListResourcesResult

`func NewListResourcesResult(ids []string, ) *ListResourcesResult`

NewListResourcesResult instantiates a new ListResourcesResult object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewListResourcesResultWithDefaults

`func NewListResourcesResultWithDefaults() *ListResourcesResult`

NewListResourcesResultWithDefaults instantiates a new ListResourcesResult object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetIds

`func (o *ListResourcesResult) GetIds() []string`

GetIds returns the Ids field if non-nil, zero value otherwise.

### GetIdsOk

`func (o *ListResourcesResult) GetIdsOk() (*[]string, bool)`

GetIdsOk returns a tuple with the Ids field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIds

`func (o *ListResourcesResult) SetIds(v []string)`

SetIds sets Ids field to given value.


### GetNextPageToken

`func (o *ListResourcesResult) GetNextPageToken() string`

GetNextPageToken returns the NextPageToken field if non-nil, zero value otherwise.

### GetNextPageTokenOk

`func (o *ListResourcesResult) GetNextPageTokenOk() (*string, bool)`

GetNextPageTokenOk returns a tuple with the NextPageToken field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNextPageToken

`func (o *ListResourcesResult) SetNextPageToken(v string)`

SetNextPageToken sets NextPageToken field to given value.

### HasNextPageToken

`func (o *ListResourcesResult) HasNextPageToken() bool`

HasNextPageToken returns a boolean if a field has been set.

### GetResources

`func (o *ListResourcesResult) GetResources() []DescribeResourceResult`

GetResources returns the Resources field if non-nil, zero value otherwise.

### GetResourcesOk

`func (o *ListResourcesResult) GetResourcesOk() (*[]DescribeResourceResult, bool)`

GetResourcesOk returns a tuple with the Resources field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetResources

`func (o *ListResourcesResult) SetResources(v []DescribeResourceResult)`

SetResources sets Resources field to given value.

### HasResources

`func (o *ListResourcesResult) HasResources() bool`

HasResources returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


