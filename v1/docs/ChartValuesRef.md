# ChartValuesRef

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Scope** | Pointer to **map[string]string** | Scope condition for applying these chart values (can be used with either values or valuesFile). Key is the parameter name which supports template expressions like &#39;{{ $sys.cloudProviderName }}&#39;, value is the expected value.  | [optional] 
**Values** | Pointer to **interface{}** | Inline chart values as a map (mutually exclusive with valuesFile) | [optional] 
**ValuesFile** | Pointer to [**ValuesFile**](ValuesFile.md) |  | [optional] 

## Methods

### NewChartValuesRef

`func NewChartValuesRef() *ChartValuesRef`

NewChartValuesRef instantiates a new ChartValuesRef object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewChartValuesRefWithDefaults

`func NewChartValuesRefWithDefaults() *ChartValuesRef`

NewChartValuesRefWithDefaults instantiates a new ChartValuesRef object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetScope

`func (o *ChartValuesRef) GetScope() map[string]string`

GetScope returns the Scope field if non-nil, zero value otherwise.

### GetScopeOk

`func (o *ChartValuesRef) GetScopeOk() (*map[string]string, bool)`

GetScopeOk returns a tuple with the Scope field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetScope

`func (o *ChartValuesRef) SetScope(v map[string]string)`

SetScope sets Scope field to given value.

### HasScope

`func (o *ChartValuesRef) HasScope() bool`

HasScope returns a boolean if a field has been set.

### GetValues

`func (o *ChartValuesRef) GetValues() interface{}`

GetValues returns the Values field if non-nil, zero value otherwise.

### GetValuesOk

`func (o *ChartValuesRef) GetValuesOk() (*interface{}, bool)`

GetValuesOk returns a tuple with the Values field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetValues

`func (o *ChartValuesRef) SetValues(v interface{})`

SetValues sets Values field to given value.

### HasValues

`func (o *ChartValuesRef) HasValues() bool`

HasValues returns a boolean if a field has been set.

### SetValuesNil

`func (o *ChartValuesRef) SetValuesNil(b bool)`

 SetValuesNil sets the value for Values to be an explicit nil

### UnsetValues
`func (o *ChartValuesRef) UnsetValues()`

UnsetValues ensures that no value is present for Values, not even an explicit nil
### GetValuesFile

`func (o *ChartValuesRef) GetValuesFile() ValuesFile`

GetValuesFile returns the ValuesFile field if non-nil, zero value otherwise.

### GetValuesFileOk

`func (o *ChartValuesRef) GetValuesFileOk() (*ValuesFile, bool)`

GetValuesFileOk returns a tuple with the ValuesFile field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetValuesFile

`func (o *ChartValuesRef) SetValuesFile(v ValuesFile)`

SetValuesFile sets ValuesFile field to given value.

### HasValuesFile

`func (o *ChartValuesRef) HasValuesFile() bool`

HasValuesFile returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


