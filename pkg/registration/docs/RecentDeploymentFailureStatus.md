# RecentDeploymentFailureStatus

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**FailedAt** | **string** | The time at which the deployment failed | 
**FailureReason** | **string** | The reason for the deployment failure | 

## Methods

### NewRecentDeploymentFailureStatus

`func NewRecentDeploymentFailureStatus(failedAt string, failureReason string, ) *RecentDeploymentFailureStatus`

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



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


