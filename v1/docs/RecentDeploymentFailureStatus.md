# RecentDeploymentFailureStatus

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**FailedAt** | Pointer to **string** | The time at which the deployment failed | [optional] 
**FailureReason** | Pointer to **string** | The reason for the deployment failure | [optional] 

## Methods

### NewRecentDeploymentFailureStatus

`func NewRecentDeploymentFailureStatus() *RecentDeploymentFailureStatus`

NewRecentDeploymentFailureStatus instantiates a new RecentDeploymentFailureStatus object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewRecentDeploymentFailureStatusWithDefaults

`func NewRecentDeploymentFailureStatusWithDefaults() *RecentDeploymentFailureStatus`

NewRecentDeploymentFailureStatusWithDefaults instantiates a new RecentDeploymentFailureStatus object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetFailedAt

`func (o *RecentDeploymentFailureStatus) GetFailedAt() string`

GetFailedAt returns the FailedAt field if non-nil, zero value otherwise.

### GetFailedAtOk

`func (o *RecentDeploymentFailureStatus) GetFailedAtOk() (*string, bool)`

GetFailedAtOk returns a tuple with the FailedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFailedAt

`func (o *RecentDeploymentFailureStatus) SetFailedAt(v string)`

SetFailedAt sets FailedAt field to given value.

### HasFailedAt

`func (o *RecentDeploymentFailureStatus) HasFailedAt() bool`

HasFailedAt returns a boolean if a field has been set.

### GetFailureReason

`func (o *RecentDeploymentFailureStatus) GetFailureReason() string`

GetFailureReason returns the FailureReason field if non-nil, zero value otherwise.

### GetFailureReasonOk

`func (o *RecentDeploymentFailureStatus) GetFailureReasonOk() (*string, bool)`

GetFailureReasonOk returns a tuple with the FailureReason field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFailureReason

`func (o *RecentDeploymentFailureStatus) SetFailureReason(v string)`

SetFailureReason sets FailureReason field to given value.

### HasFailureReason

`func (o *RecentDeploymentFailureStatus) HasFailureReason() bool`

HasFailureReason returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


