# AcceleratorConfiguration

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Count** | **int64** | Number of accelerators to attach | 
**Type** | **string** | Type of accelerator (GPU) | 

## Methods

### NewAcceleratorConfiguration

`func NewAcceleratorConfiguration(count int64, type_ string, ) *AcceleratorConfiguration`

NewAcceleratorConfiguration instantiates a new AcceleratorConfiguration object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewAcceleratorConfigurationWithDefaults

`func NewAcceleratorConfigurationWithDefaults() *AcceleratorConfiguration`

NewAcceleratorConfigurationWithDefaults instantiates a new AcceleratorConfiguration object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetCount

`func (o *AcceleratorConfiguration) GetCount() int64`

GetCount returns the Count field if non-nil, zero value otherwise.

### GetCountOk

`func (o *AcceleratorConfiguration) GetCountOk() (*int64, bool)`

GetCountOk returns a tuple with the Count field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCount

`func (o *AcceleratorConfiguration) SetCount(v int64)`

SetCount sets Count field to given value.


### GetType

`func (o *AcceleratorConfiguration) GetType() string`

GetType returns the Type field if non-nil, zero value otherwise.

### GetTypeOk

`func (o *AcceleratorConfiguration) GetTypeOk() (*string, bool)`

GetTypeOk returns a tuple with the Type field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetType

`func (o *AcceleratorConfiguration) SetType(v string)`

SetType sets Type field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


