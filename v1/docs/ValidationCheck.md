# ValidationCheck

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Name** | **string** | The validation check category. Every category is reported on every response; categories that the submitted specification does not exercise are reported explicitly as NOT_APPLICABLE rather than omitted. | 
**Status** | **string** | The outcome of a single validation check category. PASSED means the category completed with no error. FAILED means the category produced at least one error diagnostic. INCOMPLETE means the category could not be completed, including when it was blocked by failed syntax or normalization. NOT_APPLICABLE means the submitted specification does not exercise the category at all. | 

## Methods

### NewValidationCheck

`func NewValidationCheck(name string, status string, ) *ValidationCheck`

NewValidationCheck instantiates a new ValidationCheck object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewValidationCheckWithDefaults

`func NewValidationCheckWithDefaults() *ValidationCheck`

NewValidationCheckWithDefaults instantiates a new ValidationCheck object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetName

`func (o *ValidationCheck) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *ValidationCheck) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *ValidationCheck) SetName(v string)`

SetName sets Name field to given value.


### GetStatus

`func (o *ValidationCheck) GetStatus() string`

GetStatus returns the Status field if non-nil, zero value otherwise.

### GetStatusOk

`func (o *ValidationCheck) GetStatusOk() (*string, bool)`

GetStatusOk returns a tuple with the Status field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStatus

`func (o *ValidationCheck) SetStatus(v string)`

SetStatus sets Status field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


