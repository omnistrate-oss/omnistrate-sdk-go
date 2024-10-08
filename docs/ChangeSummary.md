# ChangeSummary

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Changes** | [**[]Change**](Change.md) | List of individual changes | 
**Status** | **string** | The product tier feature changes for the resource. | 

## Methods

### NewChangeSummary

`func NewChangeSummary(changes []Change, status string, ) *ChangeSummary`

NewChangeSummary instantiates a new ChangeSummary object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewChangeSummaryWithDefaults

`func NewChangeSummaryWithDefaults() *ChangeSummary`

NewChangeSummaryWithDefaults instantiates a new ChangeSummary object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetChanges

`func (o *ChangeSummary) GetChanges() []Change`

GetChanges returns the Changes field if non-nil, zero value otherwise.

### GetChangesOk

`func (o *ChangeSummary) GetChangesOk() (*[]Change, bool)`

GetChangesOk returns a tuple with the Changes field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetChanges

`func (o *ChangeSummary) SetChanges(v []Change)`

SetChanges sets Changes field to given value.


### GetStatus

`func (o *ChangeSummary) GetStatus() string`

GetStatus returns the Status field if non-nil, zero value otherwise.

### GetStatusOk

`func (o *ChangeSummary) GetStatusOk() (*string, bool)`

GetStatusOk returns a tuple with the Status field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStatus

`func (o *ChangeSummary) SetStatus(v string)`

SetStatus sets Status field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


