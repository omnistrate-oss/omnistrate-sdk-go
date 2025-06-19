# ListServicePlansRequest

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**NextPageToken** | Pointer to **string** | The next token to use for pagination | [optional] 
**PageSize** | Pointer to **int64** | The number of resources to return per page | [optional] 
**ServiceEnvironmentId** | **string** | ID of a Service Environment | 
**ServiceId** | **string** | ID of a Service | 
**SkipHasPendingChangesCheck** | Pointer to **bool** | Whether to skip the check for pending changes in the service plans | [optional] 
**Token** | **string** | JWT token used to perform authorization | 

## Methods

### NewListServicePlansRequest

`func NewListServicePlansRequest(serviceEnvironmentId string, serviceId string, token string, ) *ListServicePlansRequest`

NewListServicePlansRequest instantiates a new ListServicePlansRequest object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewListServicePlansRequestWithDefaults

`func NewListServicePlansRequestWithDefaults() *ListServicePlansRequest`

NewListServicePlansRequestWithDefaults instantiates a new ListServicePlansRequest object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetNextPageToken

`func (o *ListServicePlansRequest) GetNextPageToken() string`

GetNextPageToken returns the NextPageToken field if non-nil, zero value otherwise.

### GetNextPageTokenOk

`func (o *ListServicePlansRequest) GetNextPageTokenOk() (*string, bool)`

GetNextPageTokenOk returns a tuple with the NextPageToken field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNextPageToken

`func (o *ListServicePlansRequest) SetNextPageToken(v string)`

SetNextPageToken sets NextPageToken field to given value.

### HasNextPageToken

`func (o *ListServicePlansRequest) HasNextPageToken() bool`

HasNextPageToken returns a boolean if a field has been set.

### GetPageSize

`func (o *ListServicePlansRequest) GetPageSize() int64`

GetPageSize returns the PageSize field if non-nil, zero value otherwise.

### GetPageSizeOk

`func (o *ListServicePlansRequest) GetPageSizeOk() (*int64, bool)`

GetPageSizeOk returns a tuple with the PageSize field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPageSize

`func (o *ListServicePlansRequest) SetPageSize(v int64)`

SetPageSize sets PageSize field to given value.

### HasPageSize

`func (o *ListServicePlansRequest) HasPageSize() bool`

HasPageSize returns a boolean if a field has been set.

### GetServiceEnvironmentId

`func (o *ListServicePlansRequest) GetServiceEnvironmentId() string`

GetServiceEnvironmentId returns the ServiceEnvironmentId field if non-nil, zero value otherwise.

### GetServiceEnvironmentIdOk

`func (o *ListServicePlansRequest) GetServiceEnvironmentIdOk() (*string, bool)`

GetServiceEnvironmentIdOk returns a tuple with the ServiceEnvironmentId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetServiceEnvironmentId

`func (o *ListServicePlansRequest) SetServiceEnvironmentId(v string)`

SetServiceEnvironmentId sets ServiceEnvironmentId field to given value.


### GetServiceId

`func (o *ListServicePlansRequest) GetServiceId() string`

GetServiceId returns the ServiceId field if non-nil, zero value otherwise.

### GetServiceIdOk

`func (o *ListServicePlansRequest) GetServiceIdOk() (*string, bool)`

GetServiceIdOk returns a tuple with the ServiceId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetServiceId

`func (o *ListServicePlansRequest) SetServiceId(v string)`

SetServiceId sets ServiceId field to given value.


### GetSkipHasPendingChangesCheck

`func (o *ListServicePlansRequest) GetSkipHasPendingChangesCheck() bool`

GetSkipHasPendingChangesCheck returns the SkipHasPendingChangesCheck field if non-nil, zero value otherwise.

### GetSkipHasPendingChangesCheckOk

`func (o *ListServicePlansRequest) GetSkipHasPendingChangesCheckOk() (*bool, bool)`

GetSkipHasPendingChangesCheckOk returns a tuple with the SkipHasPendingChangesCheck field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSkipHasPendingChangesCheck

`func (o *ListServicePlansRequest) SetSkipHasPendingChangesCheck(v bool)`

SetSkipHasPendingChangesCheck sets SkipHasPendingChangesCheck field to given value.

### HasSkipHasPendingChangesCheck

`func (o *ListServicePlansRequest) HasSkipHasPendingChangesCheck() bool`

HasSkipHasPendingChangesCheck returns a boolean if a field has been set.

### GetToken

`func (o *ListServicePlansRequest) GetToken() string`

GetToken returns the Token field if non-nil, zero value otherwise.

### GetTokenOk

`func (o *ListServicePlansRequest) GetTokenOk() (*string, bool)`

GetTokenOk returns a tuple with the Token field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetToken

`func (o *ListServicePlansRequest) SetToken(v string)`

SetToken sets Token field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


