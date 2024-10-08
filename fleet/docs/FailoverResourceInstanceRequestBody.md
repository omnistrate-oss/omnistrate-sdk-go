# FailoverResourceInstanceRequestBody

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**FailedReplicaAction** | Pointer to **string** | The failed replica action | [optional] 
**FailedReplicaID** | **string** | The failed replica ID | 

## Methods

### NewFailoverResourceInstanceRequestBody

`func NewFailoverResourceInstanceRequestBody(failedReplicaID string, ) *FailoverResourceInstanceRequestBody`

NewFailoverResourceInstanceRequestBody instantiates a new FailoverResourceInstanceRequestBody object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewFailoverResourceInstanceRequestBodyWithDefaults

`func NewFailoverResourceInstanceRequestBodyWithDefaults() *FailoverResourceInstanceRequestBody`

NewFailoverResourceInstanceRequestBodyWithDefaults instantiates a new FailoverResourceInstanceRequestBody object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetFailedReplicaAction

`func (o *FailoverResourceInstanceRequestBody) GetFailedReplicaAction() string`

GetFailedReplicaAction returns the FailedReplicaAction field if non-nil, zero value otherwise.

### GetFailedReplicaActionOk

`func (o *FailoverResourceInstanceRequestBody) GetFailedReplicaActionOk() (*string, bool)`

GetFailedReplicaActionOk returns a tuple with the FailedReplicaAction field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFailedReplicaAction

`func (o *FailoverResourceInstanceRequestBody) SetFailedReplicaAction(v string)`

SetFailedReplicaAction sets FailedReplicaAction field to given value.

### HasFailedReplicaAction

`func (o *FailoverResourceInstanceRequestBody) HasFailedReplicaAction() bool`

HasFailedReplicaAction returns a boolean if a field has been set.

### GetFailedReplicaID

`func (o *FailoverResourceInstanceRequestBody) GetFailedReplicaID() string`

GetFailedReplicaID returns the FailedReplicaID field if non-nil, zero value otherwise.

### GetFailedReplicaIDOk

`func (o *FailoverResourceInstanceRequestBody) GetFailedReplicaIDOk() (*string, bool)`

GetFailedReplicaIDOk returns a tuple with the FailedReplicaID field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFailedReplicaID

`func (o *FailoverResourceInstanceRequestBody) SetFailedReplicaID(v string)`

SetFailedReplicaID sets FailedReplicaID field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


