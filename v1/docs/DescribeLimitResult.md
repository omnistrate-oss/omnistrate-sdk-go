# DescribeLimitResult

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Description** | **string** | A brief description of the limit | 
**Family** | **string** | The limit family | 
**Key** | **string** | Unique key to identify the limit | 
**Modifiable** | **bool** | Whether the limit can be modified | 
**Name** | **string** | Name of the limit | 
**Value** | **int64** | Value of the limit being enforced | 

## Methods

### NewDescribeLimitResult

`func NewDescribeLimitResult(description string, family string, key string, modifiable bool, name string, value int64, ) *DescribeLimitResult`

NewDescribeLimitResult instantiates a new DescribeLimitResult object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewDescribeLimitResultWithDefaults

`func NewDescribeLimitResultWithDefaults() *DescribeLimitResult`

NewDescribeLimitResultWithDefaults instantiates a new DescribeLimitResult object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetDescription

`func (o *DescribeLimitResult) GetDescription() string`

GetDescription returns the Description field if non-nil, zero value otherwise.

### GetDescriptionOk

`func (o *DescribeLimitResult) GetDescriptionOk() (*string, bool)`

GetDescriptionOk returns a tuple with the Description field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDescription

`func (o *DescribeLimitResult) SetDescription(v string)`

SetDescription sets Description field to given value.


### GetFamily

`func (o *DescribeLimitResult) GetFamily() string`

GetFamily returns the Family field if non-nil, zero value otherwise.

### GetFamilyOk

`func (o *DescribeLimitResult) GetFamilyOk() (*string, bool)`

GetFamilyOk returns a tuple with the Family field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFamily

`func (o *DescribeLimitResult) SetFamily(v string)`

SetFamily sets Family field to given value.


### GetKey

`func (o *DescribeLimitResult) GetKey() string`

GetKey returns the Key field if non-nil, zero value otherwise.

### GetKeyOk

`func (o *DescribeLimitResult) GetKeyOk() (*string, bool)`

GetKeyOk returns a tuple with the Key field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetKey

`func (o *DescribeLimitResult) SetKey(v string)`

SetKey sets Key field to given value.


### GetModifiable

`func (o *DescribeLimitResult) GetModifiable() bool`

GetModifiable returns the Modifiable field if non-nil, zero value otherwise.

### GetModifiableOk

`func (o *DescribeLimitResult) GetModifiableOk() (*bool, bool)`

GetModifiableOk returns a tuple with the Modifiable field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetModifiable

`func (o *DescribeLimitResult) SetModifiable(v bool)`

SetModifiable sets Modifiable field to given value.


### GetName

`func (o *DescribeLimitResult) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *DescribeLimitResult) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *DescribeLimitResult) SetName(v string)`

SetName sets Name field to given value.


### GetValue

`func (o *DescribeLimitResult) GetValue() int64`

GetValue returns the Value field if non-nil, zero value otherwise.

### GetValueOk

`func (o *DescribeLimitResult) GetValueOk() (*int64, bool)`

GetValueOk returns a tuple with the Value field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetValue

`func (o *DescribeLimitResult) SetValue(v int64)`

SetValue sets Value field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


