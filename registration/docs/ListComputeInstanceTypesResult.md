# ListComputeInstanceTypesResult

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**NextPageToken** | Pointer to **string** | The next page token | [optional] 
**Types** | **[]string** | The list of compute instance type IDs | 

## Methods

### NewListComputeInstanceTypesResult

`func NewListComputeInstanceTypesResult(types []string, ) *ListComputeInstanceTypesResult`

NewListComputeInstanceTypesResult instantiates a new ListComputeInstanceTypesResult object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewListComputeInstanceTypesResultWithDefaults

`func NewListComputeInstanceTypesResultWithDefaults() *ListComputeInstanceTypesResult`

NewListComputeInstanceTypesResultWithDefaults instantiates a new ListComputeInstanceTypesResult object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetNextPageToken

`func (o *ListComputeInstanceTypesResult) GetNextPageToken() string`

GetNextPageToken returns the NextPageToken field if non-nil, zero value otherwise.

### GetNextPageTokenOk

`func (o *ListComputeInstanceTypesResult) GetNextPageTokenOk() (*string, bool)`

GetNextPageTokenOk returns a tuple with the NextPageToken field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNextPageToken

`func (o *ListComputeInstanceTypesResult) SetNextPageToken(v string)`

SetNextPageToken sets NextPageToken field to given value.

### HasNextPageToken

`func (o *ListComputeInstanceTypesResult) HasNextPageToken() bool`

HasNextPageToken returns a boolean if a field has been set.

### GetTypes

`func (o *ListComputeInstanceTypesResult) GetTypes() []string`

GetTypes returns the Types field if non-nil, zero value otherwise.

### GetTypesOk

`func (o *ListComputeInstanceTypesResult) GetTypesOk() (*[]string, bool)`

GetTypesOk returns a tuple with the Types field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTypes

`func (o *ListComputeInstanceTypesResult) SetTypes(v []string)`

SetTypes sets Types field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


