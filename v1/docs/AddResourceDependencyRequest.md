# AddResourceDependencyRequest

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Id** | **string** | ID of a resource | 
**ParameterMap** | Pointer to **map[string]string** | A map of the source parameter to the dependency resource parameter | [optional] 
**ResourceDependencyId** | **string** | ID of a resource | 
**ServiceId** | **string** | ID of a Service | 
**Token** | **string** | JWT token used to perform authorization | 

## Methods

### NewAddResourceDependencyRequest

`func NewAddResourceDependencyRequest(id string, resourceDependencyId string, serviceId string, token string, ) *AddResourceDependencyRequest`

NewAddResourceDependencyRequest instantiates a new AddResourceDependencyRequest object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewAddResourceDependencyRequestWithDefaults

`func NewAddResourceDependencyRequestWithDefaults() *AddResourceDependencyRequest`

NewAddResourceDependencyRequestWithDefaults instantiates a new AddResourceDependencyRequest object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetId

`func (o *AddResourceDependencyRequest) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *AddResourceDependencyRequest) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *AddResourceDependencyRequest) SetId(v string)`

SetId sets Id field to given value.


### GetParameterMap

`func (o *AddResourceDependencyRequest) GetParameterMap() map[string]string`

GetParameterMap returns the ParameterMap field if non-nil, zero value otherwise.

### GetParameterMapOk

`func (o *AddResourceDependencyRequest) GetParameterMapOk() (*map[string]string, bool)`

GetParameterMapOk returns a tuple with the ParameterMap field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetParameterMap

`func (o *AddResourceDependencyRequest) SetParameterMap(v map[string]string)`

SetParameterMap sets ParameterMap field to given value.

### HasParameterMap

`func (o *AddResourceDependencyRequest) HasParameterMap() bool`

HasParameterMap returns a boolean if a field has been set.

### GetResourceDependencyId

`func (o *AddResourceDependencyRequest) GetResourceDependencyId() string`

GetResourceDependencyId returns the ResourceDependencyId field if non-nil, zero value otherwise.

### GetResourceDependencyIdOk

`func (o *AddResourceDependencyRequest) GetResourceDependencyIdOk() (*string, bool)`

GetResourceDependencyIdOk returns a tuple with the ResourceDependencyId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetResourceDependencyId

`func (o *AddResourceDependencyRequest) SetResourceDependencyId(v string)`

SetResourceDependencyId sets ResourceDependencyId field to given value.


### GetServiceId

`func (o *AddResourceDependencyRequest) GetServiceId() string`

GetServiceId returns the ServiceId field if non-nil, zero value otherwise.

### GetServiceIdOk

`func (o *AddResourceDependencyRequest) GetServiceIdOk() (*string, bool)`

GetServiceIdOk returns a tuple with the ServiceId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetServiceId

`func (o *AddResourceDependencyRequest) SetServiceId(v string)`

SetServiceId sets ServiceId field to given value.


### GetToken

`func (o *AddResourceDependencyRequest) GetToken() string`

GetToken returns the Token field if non-nil, zero value otherwise.

### GetTokenOk

`func (o *AddResourceDependencyRequest) GetTokenOk() (*string, bool)`

GetTokenOk returns a tuple with the Token field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetToken

`func (o *AddResourceDependencyRequest) SetToken(v string)`

SetToken sets Token field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


