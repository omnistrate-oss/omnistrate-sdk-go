# ListStorageConfigsRequest

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Managed** | Pointer to **bool** | Is storage config managed by omnistrate | [optional] 
**ServiceId** | **string** | ID of a Service | 
**Token** | **string** | JWT token used to perform authorization | 

## Methods

### NewListStorageConfigsRequest

`func NewListStorageConfigsRequest(serviceId string, token string, ) *ListStorageConfigsRequest`

NewListStorageConfigsRequest instantiates a new ListStorageConfigsRequest object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewListStorageConfigsRequestWithDefaults

`func NewListStorageConfigsRequestWithDefaults() *ListStorageConfigsRequest`

NewListStorageConfigsRequestWithDefaults instantiates a new ListStorageConfigsRequest object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetManaged

`func (o *ListStorageConfigsRequest) GetManaged() bool`

GetManaged returns the Managed field if non-nil, zero value otherwise.

### GetManagedOk

`func (o *ListStorageConfigsRequest) GetManagedOk() (*bool, bool)`

GetManagedOk returns a tuple with the Managed field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetManaged

`func (o *ListStorageConfigsRequest) SetManaged(v bool)`

SetManaged sets Managed field to given value.

### HasManaged

`func (o *ListStorageConfigsRequest) HasManaged() bool`

HasManaged returns a boolean if a field has been set.

### GetServiceId

`func (o *ListStorageConfigsRequest) GetServiceId() string`

GetServiceId returns the ServiceId field if non-nil, zero value otherwise.

### GetServiceIdOk

`func (o *ListStorageConfigsRequest) GetServiceIdOk() (*string, bool)`

GetServiceIdOk returns a tuple with the ServiceId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetServiceId

`func (o *ListStorageConfigsRequest) SetServiceId(v string)`

SetServiceId sets ServiceId field to given value.


### GetToken

`func (o *ListStorageConfigsRequest) GetToken() string`

GetToken returns the Token field if non-nil, zero value otherwise.

### GetTokenOk

`func (o *ListStorageConfigsRequest) GetTokenOk() (*string, bool)`

GetTokenOk returns a tuple with the Token field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetToken

`func (o *ListStorageConfigsRequest) SetToken(v string)`

SetToken sets Token field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


