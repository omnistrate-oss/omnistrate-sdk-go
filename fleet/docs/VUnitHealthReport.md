# VUnitHealthReport

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Vunits** | Pointer to **map[string]interface{}** | The health of each vunit under this service environment | [optional] 

## Methods

### NewVUnitHealthReport

`func NewVUnitHealthReport() *VUnitHealthReport`

NewVUnitHealthReport instantiates a new VUnitHealthReport object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewVUnitHealthReportWithDefaults

`func NewVUnitHealthReportWithDefaults() *VUnitHealthReport`

NewVUnitHealthReportWithDefaults instantiates a new VUnitHealthReport object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetVunits

`func (o *VUnitHealthReport) GetVunits() map[string]interface{}`

GetVunits returns the Vunits field if non-nil, zero value otherwise.

### GetVunitsOk

`func (o *VUnitHealthReport) GetVunitsOk() (*map[string]interface{}, bool)`

GetVunitsOk returns a tuple with the Vunits field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVunits

`func (o *VUnitHealthReport) SetVunits(v map[string]interface{})`

SetVunits sets Vunits field to given value.

### HasVunits

`func (o *VUnitHealthReport) HasVunits() bool`

HasVunits returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


