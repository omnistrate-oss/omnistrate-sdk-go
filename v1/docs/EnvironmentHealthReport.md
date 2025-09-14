# EnvironmentHealthReport

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Models** | Pointer to [**map[string]VUnitHealthReport**](VUnitHealthReport.md) | Health report for each model in the environment | [optional] 

## Methods

### NewEnvironmentHealthReport

`func NewEnvironmentHealthReport() *EnvironmentHealthReport`

NewEnvironmentHealthReport instantiates a new EnvironmentHealthReport object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewEnvironmentHealthReportWithDefaults

`func NewEnvironmentHealthReportWithDefaults() *EnvironmentHealthReport`

NewEnvironmentHealthReportWithDefaults instantiates a new EnvironmentHealthReport object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetModels

`func (o *EnvironmentHealthReport) GetModels() map[string]VUnitHealthReport`

GetModels returns the Models field if non-nil, zero value otherwise.

### GetModelsOk

`func (o *EnvironmentHealthReport) GetModelsOk() (*map[string]VUnitHealthReport, bool)`

GetModelsOk returns a tuple with the Models field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetModels

`func (o *EnvironmentHealthReport) SetModels(v map[string]VUnitHealthReport)`

SetModels sets Models field to given value.

### HasModels

`func (o *EnvironmentHealthReport) HasModels() bool`

HasModels returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


