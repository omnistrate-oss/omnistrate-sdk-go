# ReportHealthResult

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Health** | Pointer to **string** | Service Health Status | [optional] 
**Id** | Pointer to **string** | ID of a Service | [optional] 
**Report** | Pointer to [**map[string]EnvironmentHealthReport**](EnvironmentHealthReport.md) | The health of each environment under this service | [optional] 

## Methods

### NewReportHealthResult

`func NewReportHealthResult() *ReportHealthResult`

NewReportHealthResult instantiates a new ReportHealthResult object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewReportHealthResultWithDefaults

`func NewReportHealthResultWithDefaults() *ReportHealthResult`

NewReportHealthResultWithDefaults instantiates a new ReportHealthResult object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetHealth

`func (o *ReportHealthResult) GetHealth() string`

GetHealth returns the Health field if non-nil, zero value otherwise.

### GetHealthOk

`func (o *ReportHealthResult) GetHealthOk() (*string, bool)`

GetHealthOk returns a tuple with the Health field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetHealth

`func (o *ReportHealthResult) SetHealth(v string)`

SetHealth sets Health field to given value.

### HasHealth

`func (o *ReportHealthResult) HasHealth() bool`

HasHealth returns a boolean if a field has been set.

### GetId

`func (o *ReportHealthResult) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *ReportHealthResult) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *ReportHealthResult) SetId(v string)`

SetId sets Id field to given value.

### HasId

`func (o *ReportHealthResult) HasId() bool`

HasId returns a boolean if a field has been set.

### GetReport

`func (o *ReportHealthResult) GetReport() map[string]EnvironmentHealthReport`

GetReport returns the Report field if non-nil, zero value otherwise.

### GetReportOk

`func (o *ReportHealthResult) GetReportOk() (*map[string]EnvironmentHealthReport, bool)`

GetReportOk returns a tuple with the Report field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetReport

`func (o *ReportHealthResult) SetReport(v map[string]EnvironmentHealthReport)`

SetReport sets Report field to given value.

### HasReport

`func (o *ReportHealthResult) HasReport() bool`

HasReport returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


