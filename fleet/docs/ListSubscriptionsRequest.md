# ListSubscriptionsRequest

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**EnvironmentType** | Pointer to **string** | The type of service environment | [optional] 
**IncludeInactive** | Pointer to **bool** | Whether to include inactive (suspended, cancelled, terminated) subscriptions | [optional] 
**ServiceId** | Pointer to **string** | ID of a Service | [optional] 
**Token** | **string** | JWT token used to perform authorization | 

## Methods

### NewListSubscriptionsRequest

`func NewListSubscriptionsRequest(token string, ) *ListSubscriptionsRequest`

NewListSubscriptionsRequest instantiates a new ListSubscriptionsRequest object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewListSubscriptionsRequestWithDefaults

`func NewListSubscriptionsRequestWithDefaults() *ListSubscriptionsRequest`

NewListSubscriptionsRequestWithDefaults instantiates a new ListSubscriptionsRequest object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetEnvironmentType

`func (o *ListSubscriptionsRequest) GetEnvironmentType() string`

GetEnvironmentType returns the EnvironmentType field if non-nil, zero value otherwise.

### GetEnvironmentTypeOk

`func (o *ListSubscriptionsRequest) GetEnvironmentTypeOk() (*string, bool)`

GetEnvironmentTypeOk returns a tuple with the EnvironmentType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEnvironmentType

`func (o *ListSubscriptionsRequest) SetEnvironmentType(v string)`

SetEnvironmentType sets EnvironmentType field to given value.

### HasEnvironmentType

`func (o *ListSubscriptionsRequest) HasEnvironmentType() bool`

HasEnvironmentType returns a boolean if a field has been set.

### GetIncludeInactive

`func (o *ListSubscriptionsRequest) GetIncludeInactive() bool`

GetIncludeInactive returns the IncludeInactive field if non-nil, zero value otherwise.

### GetIncludeInactiveOk

`func (o *ListSubscriptionsRequest) GetIncludeInactiveOk() (*bool, bool)`

GetIncludeInactiveOk returns a tuple with the IncludeInactive field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIncludeInactive

`func (o *ListSubscriptionsRequest) SetIncludeInactive(v bool)`

SetIncludeInactive sets IncludeInactive field to given value.

### HasIncludeInactive

`func (o *ListSubscriptionsRequest) HasIncludeInactive() bool`

HasIncludeInactive returns a boolean if a field has been set.

### GetServiceId

`func (o *ListSubscriptionsRequest) GetServiceId() string`

GetServiceId returns the ServiceId field if non-nil, zero value otherwise.

### GetServiceIdOk

`func (o *ListSubscriptionsRequest) GetServiceIdOk() (*string, bool)`

GetServiceIdOk returns a tuple with the ServiceId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetServiceId

`func (o *ListSubscriptionsRequest) SetServiceId(v string)`

SetServiceId sets ServiceId field to given value.

### HasServiceId

`func (o *ListSubscriptionsRequest) HasServiceId() bool`

HasServiceId returns a boolean if a field has been set.

### GetToken

`func (o *ListSubscriptionsRequest) GetToken() string`

GetToken returns the Token field if non-nil, zero value otherwise.

### GetTokenOk

`func (o *ListSubscriptionsRequest) GetTokenOk() (*string, bool)`

GetTokenOk returns a tuple with the Token field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetToken

`func (o *ListSubscriptionsRequest) SetToken(v string)`

SetToken sets Token field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


