# ListDependentResourcesResult

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Id** | Pointer to **string** | The ID of the resource | [optional] 
**Ids** | **[]string** | List of dependent resource IDs | 
**NextPageToken** | Pointer to **string** | Token to use for the next page | [optional] 
**ServiceId** | Pointer to **string** | The service ID that this API bundle belongs to | [optional] 

## Methods

### NewListDependentResourcesResult

`func NewListDependentResourcesResult(ids []string, ) *ListDependentResourcesResult`

NewListDependentResourcesResult instantiates a new ListDependentResourcesResult object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewListDependentResourcesResultWithDefaults

`func NewListDependentResourcesResultWithDefaults() *ListDependentResourcesResult`

NewListDependentResourcesResultWithDefaults instantiates a new ListDependentResourcesResult object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetId

`func (o *ListDependentResourcesResult) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *ListDependentResourcesResult) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *ListDependentResourcesResult) SetId(v string)`

SetId sets Id field to given value.

### HasId

`func (o *ListDependentResourcesResult) HasId() bool`

HasId returns a boolean if a field has been set.

### GetIds

`func (o *ListDependentResourcesResult) GetIds() []string`

GetIds returns the Ids field if non-nil, zero value otherwise.

### GetIdsOk

`func (o *ListDependentResourcesResult) GetIdsOk() (*[]string, bool)`

GetIdsOk returns a tuple with the Ids field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIds

`func (o *ListDependentResourcesResult) SetIds(v []string)`

SetIds sets Ids field to given value.


### GetNextPageToken

`func (o *ListDependentResourcesResult) GetNextPageToken() string`

GetNextPageToken returns the NextPageToken field if non-nil, zero value otherwise.

### GetNextPageTokenOk

`func (o *ListDependentResourcesResult) GetNextPageTokenOk() (*string, bool)`

GetNextPageTokenOk returns a tuple with the NextPageToken field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNextPageToken

`func (o *ListDependentResourcesResult) SetNextPageToken(v string)`

SetNextPageToken sets NextPageToken field to given value.

### HasNextPageToken

`func (o *ListDependentResourcesResult) HasNextPageToken() bool`

HasNextPageToken returns a boolean if a field has been set.

### GetServiceId

`func (o *ListDependentResourcesResult) GetServiceId() string`

GetServiceId returns the ServiceId field if non-nil, zero value otherwise.

### GetServiceIdOk

`func (o *ListDependentResourcesResult) GetServiceIdOk() (*string, bool)`

GetServiceIdOk returns a tuple with the ServiceId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetServiceId

`func (o *ListDependentResourcesResult) SetServiceId(v string)`

SetServiceId sets ServiceId field to given value.

### HasServiceId

`func (o *ListDependentResourcesResult) HasServiceId() bool`

HasServiceId returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


