# RolloutFleetInfraStatusRequest

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Id** | **string** | ID of an Infra Config | 
**ServiceId** | **string** | ID of a Service | 
**Token** | **string** | JWT token used to perform authorization | 

## Methods

### NewRolloutFleetInfraStatusRequest

`func NewRolloutFleetInfraStatusRequest(id string, serviceId string, token string, ) *RolloutFleetInfraStatusRequest`

NewRolloutFleetInfraStatusRequest instantiates a new RolloutFleetInfraStatusRequest object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewRolloutFleetInfraStatusRequestWithDefaults

`func NewRolloutFleetInfraStatusRequestWithDefaults() *RolloutFleetInfraStatusRequest`

NewRolloutFleetInfraStatusRequestWithDefaults instantiates a new RolloutFleetInfraStatusRequest object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetId

`func (o *RolloutFleetInfraStatusRequest) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *RolloutFleetInfraStatusRequest) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *RolloutFleetInfraStatusRequest) SetId(v string)`

SetId sets Id field to given value.


### GetServiceId

`func (o *RolloutFleetInfraStatusRequest) GetServiceId() string`

GetServiceId returns the ServiceId field if non-nil, zero value otherwise.

### GetServiceIdOk

`func (o *RolloutFleetInfraStatusRequest) GetServiceIdOk() (*string, bool)`

GetServiceIdOk returns a tuple with the ServiceId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetServiceId

`func (o *RolloutFleetInfraStatusRequest) SetServiceId(v string)`

SetServiceId sets ServiceId field to given value.


### GetToken

`func (o *RolloutFleetInfraStatusRequest) GetToken() string`

GetToken returns the Token field if non-nil, zero value otherwise.

### GetTokenOk

`func (o *RolloutFleetInfraStatusRequest) GetTokenOk() (*string, bool)`

GetTokenOk returns a tuple with the Token field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetToken

`func (o *RolloutFleetInfraStatusRequest) SetToken(v string)`

SetToken sets Token field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


