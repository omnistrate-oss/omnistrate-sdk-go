# ListResourcesRequest2

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**EnvironmentType** | Pointer to **string** | Filter resources by environment type. | [optional] 
**NextPageToken** | Pointer to **string** | Token to use for the next request | [optional] 
**OrgId** | Pointer to **string** | Filter resources by organization ID. | [optional] 
**PageSize** | Pointer to **int64** | The number of resources to return per page. | [optional] 
**UserId** | Pointer to **string** | Filter resources by user ID. | [optional] 

## Methods

### NewListResourcesRequest2

`func NewListResourcesRequest2() *ListResourcesRequest2`

NewListResourcesRequest2 instantiates a new ListResourcesRequest2 object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewListResourcesRequest2WithDefaults

`func NewListResourcesRequest2WithDefaults() *ListResourcesRequest2`

NewListResourcesRequest2WithDefaults instantiates a new ListResourcesRequest2 object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetEnvironmentType

`func (o *ListResourcesRequest2) GetEnvironmentType() string`

GetEnvironmentType returns the EnvironmentType field if non-nil, zero value otherwise.

### GetEnvironmentTypeOk

`func (o *ListResourcesRequest2) GetEnvironmentTypeOk() (*string, bool)`

GetEnvironmentTypeOk returns a tuple with the EnvironmentType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEnvironmentType

`func (o *ListResourcesRequest2) SetEnvironmentType(v string)`

SetEnvironmentType sets EnvironmentType field to given value.

### HasEnvironmentType

`func (o *ListResourcesRequest2) HasEnvironmentType() bool`

HasEnvironmentType returns a boolean if a field has been set.

### GetNextPageToken

`func (o *ListResourcesRequest2) GetNextPageToken() string`

GetNextPageToken returns the NextPageToken field if non-nil, zero value otherwise.

### GetNextPageTokenOk

`func (o *ListResourcesRequest2) GetNextPageTokenOk() (*string, bool)`

GetNextPageTokenOk returns a tuple with the NextPageToken field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNextPageToken

`func (o *ListResourcesRequest2) SetNextPageToken(v string)`

SetNextPageToken sets NextPageToken field to given value.

### HasNextPageToken

`func (o *ListResourcesRequest2) HasNextPageToken() bool`

HasNextPageToken returns a boolean if a field has been set.

### GetOrgId

`func (o *ListResourcesRequest2) GetOrgId() string`

GetOrgId returns the OrgId field if non-nil, zero value otherwise.

### GetOrgIdOk

`func (o *ListResourcesRequest2) GetOrgIdOk() (*string, bool)`

GetOrgIdOk returns a tuple with the OrgId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOrgId

`func (o *ListResourcesRequest2) SetOrgId(v string)`

SetOrgId sets OrgId field to given value.

### HasOrgId

`func (o *ListResourcesRequest2) HasOrgId() bool`

HasOrgId returns a boolean if a field has been set.

### GetPageSize

`func (o *ListResourcesRequest2) GetPageSize() int64`

GetPageSize returns the PageSize field if non-nil, zero value otherwise.

### GetPageSizeOk

`func (o *ListResourcesRequest2) GetPageSizeOk() (*int64, bool)`

GetPageSizeOk returns a tuple with the PageSize field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPageSize

`func (o *ListResourcesRequest2) SetPageSize(v int64)`

SetPageSize sets PageSize field to given value.

### HasPageSize

`func (o *ListResourcesRequest2) HasPageSize() bool`

HasPageSize returns a boolean if a field has been set.

### GetUserId

`func (o *ListResourcesRequest2) GetUserId() string`

GetUserId returns the UserId field if non-nil, zero value otherwise.

### GetUserIdOk

`func (o *ListResourcesRequest2) GetUserIdOk() (*string, bool)`

GetUserIdOk returns a tuple with the UserId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUserId

`func (o *ListResourcesRequest2) SetUserId(v string)`

SetUserId sets UserId field to given value.

### HasUserId

`func (o *ListResourcesRequest2) HasUserId() bool`

HasUserId returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


