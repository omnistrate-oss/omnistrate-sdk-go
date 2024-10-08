# ListServiceResult

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Ids** | **[]string** | List of service IDs | 
**NextPageToken** | Pointer to **string** | Token to retrieve the next page of results | [optional] 
**Services** | Pointer to [**[]DescribeServiceResult**](DescribeServiceResult.md) | List of services | [optional] 

## Methods

### NewListServiceResult

`func NewListServiceResult(ids []string, ) *ListServiceResult`

NewListServiceResult instantiates a new ListServiceResult object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewListServiceResultWithDefaults

`func NewListServiceResultWithDefaults() *ListServiceResult`

NewListServiceResultWithDefaults instantiates a new ListServiceResult object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetIds

`func (o *ListServiceResult) GetIds() []string`

GetIds returns the Ids field if non-nil, zero value otherwise.

### GetIdsOk

`func (o *ListServiceResult) GetIdsOk() (*[]string, bool)`

GetIdsOk returns a tuple with the Ids field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIds

`func (o *ListServiceResult) SetIds(v []string)`

SetIds sets Ids field to given value.


### GetNextPageToken

`func (o *ListServiceResult) GetNextPageToken() string`

GetNextPageToken returns the NextPageToken field if non-nil, zero value otherwise.

### GetNextPageTokenOk

`func (o *ListServiceResult) GetNextPageTokenOk() (*string, bool)`

GetNextPageTokenOk returns a tuple with the NextPageToken field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNextPageToken

`func (o *ListServiceResult) SetNextPageToken(v string)`

SetNextPageToken sets NextPageToken field to given value.

### HasNextPageToken

`func (o *ListServiceResult) HasNextPageToken() bool`

HasNextPageToken returns a boolean if a field has been set.

### GetServices

`func (o *ListServiceResult) GetServices() []DescribeServiceResult`

GetServices returns the Services field if non-nil, zero value otherwise.

### GetServicesOk

`func (o *ListServiceResult) GetServicesOk() (*[]DescribeServiceResult, bool)`

GetServicesOk returns a tuple with the Services field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetServices

`func (o *ListServiceResult) SetServices(v []DescribeServiceResult)`

SetServices sets Services field to given value.

### HasServices

`func (o *ListServiceResult) HasServices() bool`

HasServices returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


