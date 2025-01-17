# ApproveSubscriptionRequestRequest

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**EnvironmentId** | **string** | ID of a Service Environment | 
**Id** | **string** | ID of a Subscription Request | 
**ServiceId** | **string** | ID of a Service | 
**Token** | **string** | JWT token used to perform authorization | 

## Methods

### NewApproveSubscriptionRequestRequest

`func NewApproveSubscriptionRequestRequest(environmentId string, id string, serviceId string, token string, ) *ApproveSubscriptionRequestRequest`

NewApproveSubscriptionRequestRequest instantiates a new ApproveSubscriptionRequestRequest object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewApproveSubscriptionRequestRequestWithDefaults

`func NewApproveSubscriptionRequestRequestWithDefaults() *ApproveSubscriptionRequestRequest`

NewApproveSubscriptionRequestRequestWithDefaults instantiates a new ApproveSubscriptionRequestRequest object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetEnvironmentId

`func (o *ApproveSubscriptionRequestRequest) GetEnvironmentId() string`

GetEnvironmentId returns the EnvironmentId field if non-nil, zero value otherwise.

### GetEnvironmentIdOk

`func (o *ApproveSubscriptionRequestRequest) GetEnvironmentIdOk() (*string, bool)`

GetEnvironmentIdOk returns a tuple with the EnvironmentId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEnvironmentId

`func (o *ApproveSubscriptionRequestRequest) SetEnvironmentId(v string)`

SetEnvironmentId sets EnvironmentId field to given value.


### GetId

`func (o *ApproveSubscriptionRequestRequest) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *ApproveSubscriptionRequestRequest) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *ApproveSubscriptionRequestRequest) SetId(v string)`

SetId sets Id field to given value.


### GetServiceId

`func (o *ApproveSubscriptionRequestRequest) GetServiceId() string`

GetServiceId returns the ServiceId field if non-nil, zero value otherwise.

### GetServiceIdOk

`func (o *ApproveSubscriptionRequestRequest) GetServiceIdOk() (*string, bool)`

GetServiceIdOk returns a tuple with the ServiceId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetServiceId

`func (o *ApproveSubscriptionRequestRequest) SetServiceId(v string)`

SetServiceId sets ServiceId field to given value.


### GetToken

`func (o *ApproveSubscriptionRequestRequest) GetToken() string`

GetToken returns the Token field if non-nil, zero value otherwise.

### GetTokenOk

`func (o *ApproveSubscriptionRequestRequest) GetTokenOk() (*string, bool)`

GetTokenOk returns a tuple with the Token field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetToken

`func (o *ApproveSubscriptionRequestRequest) SetToken(v string)`

SetToken sets Token field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


