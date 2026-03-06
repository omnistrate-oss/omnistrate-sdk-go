# EnvironmentPromotionStatus

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Status** | **string** | The status of the promotion | 
**TargetEnvironmentID** | **string** | ID of a Service Environment | 

## Methods

### NewEnvironmentPromotionStatus

`func NewEnvironmentPromotionStatus(status string, targetEnvironmentID string, ) *EnvironmentPromotionStatus`

NewEnvironmentPromotionStatus instantiates a new EnvironmentPromotionStatus object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewEnvironmentPromotionStatusWithDefaults

`func NewEnvironmentPromotionStatusWithDefaults() *EnvironmentPromotionStatus`

NewEnvironmentPromotionStatusWithDefaults instantiates a new EnvironmentPromotionStatus object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetStatus

`func (o *EnvironmentPromotionStatus) GetStatus() string`

GetStatus returns the Status field if non-nil, zero value otherwise.

### GetStatusOk

`func (o *EnvironmentPromotionStatus) GetStatusOk() (*string, bool)`

GetStatusOk returns a tuple with the Status field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStatus

`func (o *EnvironmentPromotionStatus) SetStatus(v string)`

SetStatus sets Status field to given value.


### GetTargetEnvironmentID

`func (o *EnvironmentPromotionStatus) GetTargetEnvironmentID() string`

GetTargetEnvironmentID returns the TargetEnvironmentID field if non-nil, zero value otherwise.

### GetTargetEnvironmentIDOk

`func (o *EnvironmentPromotionStatus) GetTargetEnvironmentIDOk() (*string, bool)`

GetTargetEnvironmentIDOk returns a tuple with the TargetEnvironmentID field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTargetEnvironmentID

`func (o *EnvironmentPromotionStatus) SetTargetEnvironmentID(v string)`

SetTargetEnvironmentID sets TargetEnvironmentID field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


