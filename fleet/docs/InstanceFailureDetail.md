# InstanceFailureDetail

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Reason** | **string** | The reason the instance was not eligible for upgrade. | 

## Methods

### NewInstanceFailureDetail

`func NewInstanceFailureDetail(reason string, ) *InstanceFailureDetail`

NewInstanceFailureDetail instantiates a new InstanceFailureDetail object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewInstanceFailureDetailWithDefaults

`func NewInstanceFailureDetailWithDefaults() *InstanceFailureDetail`

NewInstanceFailureDetailWithDefaults instantiates a new InstanceFailureDetail object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetReason

`func (o *InstanceFailureDetail) GetReason() string`

GetReason returns the Reason field if non-nil, zero value otherwise.

### GetReasonOk

`func (o *InstanceFailureDetail) GetReasonOk() (*string, bool)`

GetReasonOk returns a tuple with the Reason field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetReason

`func (o *InstanceFailureDetail) SetReason(v string)`

SetReason sets Reason field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


