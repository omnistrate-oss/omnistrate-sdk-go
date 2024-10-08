# SearchServiceInventoryResult

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Results** | [**[]SearchRecord**](SearchRecord.md) | The search results | 
**ServiceId** | **string** | The service ID this workflow belongs to. | 

## Methods

### NewSearchServiceInventoryResult

`func NewSearchServiceInventoryResult(results []SearchRecord, serviceId string, ) *SearchServiceInventoryResult`

NewSearchServiceInventoryResult instantiates a new SearchServiceInventoryResult object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewSearchServiceInventoryResultWithDefaults

`func NewSearchServiceInventoryResultWithDefaults() *SearchServiceInventoryResult`

NewSearchServiceInventoryResultWithDefaults instantiates a new SearchServiceInventoryResult object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetResults

`func (o *SearchServiceInventoryResult) GetResults() []SearchRecord`

GetResults returns the Results field if non-nil, zero value otherwise.

### GetResultsOk

`func (o *SearchServiceInventoryResult) GetResultsOk() (*[]SearchRecord, bool)`

GetResultsOk returns a tuple with the Results field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetResults

`func (o *SearchServiceInventoryResult) SetResults(v []SearchRecord)`

SetResults sets Results field to given value.


### GetServiceId

`func (o *SearchServiceInventoryResult) GetServiceId() string`

GetServiceId returns the ServiceId field if non-nil, zero value otherwise.

### GetServiceIdOk

`func (o *SearchServiceInventoryResult) GetServiceIdOk() (*string, bool)`

GetServiceIdOk returns a tuple with the ServiceId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetServiceId

`func (o *SearchServiceInventoryResult) SetServiceId(v string)`

SetServiceId sets ServiceId field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


