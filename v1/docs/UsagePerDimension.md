# UsagePerDimension

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Dimension** | Pointer to **string** | Dimension of usage | [optional] 
**Total** | Pointer to **float64** | Total amount of usage during the period | [optional] 

## Methods

### NewUsagePerDimension

`func NewUsagePerDimension() *UsagePerDimension`

NewUsagePerDimension instantiates a new UsagePerDimension object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewUsagePerDimensionWithDefaults

`func NewUsagePerDimensionWithDefaults() *UsagePerDimension`

NewUsagePerDimensionWithDefaults instantiates a new UsagePerDimension object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetDimension

`func (o *UsagePerDimension) GetDimension() string`

GetDimension returns the Dimension field if non-nil, zero value otherwise.

### GetDimensionOk

`func (o *UsagePerDimension) GetDimensionOk() (*string, bool)`

GetDimensionOk returns a tuple with the Dimension field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDimension

`func (o *UsagePerDimension) SetDimension(v string)`

SetDimension sets Dimension field to given value.

### HasDimension

`func (o *UsagePerDimension) HasDimension() bool`

HasDimension returns a boolean if a field has been set.

### GetTotal

`func (o *UsagePerDimension) GetTotal() float64`

GetTotal returns the Total field if non-nil, zero value otherwise.

### GetTotalOk

`func (o *UsagePerDimension) GetTotalOk() (*float64, bool)`

GetTotalOk returns a tuple with the Total field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTotal

`func (o *UsagePerDimension) SetTotal(v float64)`

SetTotal sets Total field to given value.

### HasTotal

`func (o *UsagePerDimension) HasTotal() bool`

HasTotal returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)

