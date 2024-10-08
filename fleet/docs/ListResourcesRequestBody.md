# ListResourcesRequestBody

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**NextPageToken** | Pointer to **string** | Token to use for the next request | [optional] 
**PageSize** | Pointer to **int64** | The number of resources to return per page. | [optional] 

## Methods

### NewListResourcesRequestBody

`func NewListResourcesRequestBody() *ListResourcesRequestBody`

NewListResourcesRequestBody instantiates a new ListResourcesRequestBody object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewListResourcesRequestBodyWithDefaults

`func NewListResourcesRequestBodyWithDefaults() *ListResourcesRequestBody`

NewListResourcesRequestBodyWithDefaults instantiates a new ListResourcesRequestBody object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetNextPageToken

`func (o *ListResourcesRequestBody) GetNextPageToken() string`

GetNextPageToken returns the NextPageToken field if non-nil, zero value otherwise.

### GetNextPageTokenOk

`func (o *ListResourcesRequestBody) GetNextPageTokenOk() (*string, bool)`

GetNextPageTokenOk returns a tuple with the NextPageToken field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNextPageToken

`func (o *ListResourcesRequestBody) SetNextPageToken(v string)`

SetNextPageToken sets NextPageToken field to given value.

### HasNextPageToken

`func (o *ListResourcesRequestBody) HasNextPageToken() bool`

HasNextPageToken returns a boolean if a field has been set.

### GetPageSize

`func (o *ListResourcesRequestBody) GetPageSize() int64`

GetPageSize returns the PageSize field if non-nil, zero value otherwise.

### GetPageSizeOk

`func (o *ListResourcesRequestBody) GetPageSizeOk() (*int64, bool)`

GetPageSizeOk returns a tuple with the PageSize field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPageSize

`func (o *ListResourcesRequestBody) SetPageSize(v int64)`

SetPageSize sets PageSize field to given value.

### HasPageSize

`func (o *ListResourcesRequestBody) HasPageSize() bool`

HasPageSize returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


