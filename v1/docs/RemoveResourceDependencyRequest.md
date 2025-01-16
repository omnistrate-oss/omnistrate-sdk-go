# RemoveResourceDependencyRequest

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Id** | **string** | ID of a resource | 
**ResourceDependencyId** | **string** | ID of a resource | 
**ServiceId** | **string** | ID of a Service | 
**Token** | **string** | JWT token used to perform authorization | 

## Methods

### NewRemoveResourceDependencyRequest

`func NewRemoveResourceDependencyRequest(id string, resourceDependencyId string, serviceId string, token string, ) *RemoveResourceDependencyRequest`

NewRemoveResourceDependencyRequest instantiates a new RemoveResourceDependencyRequest object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewRemoveResourceDependencyRequestWithDefaults

`func NewRemoveResourceDependencyRequestWithDefaults() *RemoveResourceDependencyRequest`

NewRemoveResourceDependencyRequestWithDefaults instantiates a new RemoveResourceDependencyRequest object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetId

`func (o *RemoveResourceDependencyRequest) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *RemoveResourceDependencyRequest) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *RemoveResourceDependencyRequest) SetId(v string)`

SetId sets Id field to given value.


### GetResourceDependencyId

`func (o *RemoveResourceDependencyRequest) GetResourceDependencyId() string`

GetResourceDependencyId returns the ResourceDependencyId field if non-nil, zero value otherwise.

### GetResourceDependencyIdOk

`func (o *RemoveResourceDependencyRequest) GetResourceDependencyIdOk() (*string, bool)`

GetResourceDependencyIdOk returns a tuple with the ResourceDependencyId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetResourceDependencyId

`func (o *RemoveResourceDependencyRequest) SetResourceDependencyId(v string)`

SetResourceDependencyId sets ResourceDependencyId field to given value.


### GetServiceId

`func (o *RemoveResourceDependencyRequest) GetServiceId() string`

GetServiceId returns the ServiceId field if non-nil, zero value otherwise.

### GetServiceIdOk

`func (o *RemoveResourceDependencyRequest) GetServiceIdOk() (*string, bool)`

GetServiceIdOk returns a tuple with the ServiceId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetServiceId

`func (o *RemoveResourceDependencyRequest) SetServiceId(v string)`

SetServiceId sets ServiceId field to given value.


### GetToken

`func (o *RemoveResourceDependencyRequest) GetToken() string`

GetToken returns the Token field if non-nil, zero value otherwise.

### GetTokenOk

`func (o *RemoveResourceDependencyRequest) GetTokenOk() (*string, bool)`

GetTokenOk returns a tuple with the Token field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetToken

`func (o *RemoveResourceDependencyRequest) SetToken(v string)`

SetToken sets Token field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


