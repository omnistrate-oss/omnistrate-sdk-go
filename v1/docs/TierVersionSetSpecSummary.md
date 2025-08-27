# TierVersionSetSpecSummary

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**CreatedAt** | Pointer to **time.Time** | When the version was created | [optional] 
**HasSpecFile** | **bool** | Whether this version has a stored YAML spec file | 
**ProductTierId** | **string** | ID of a Product Tier | 
**ServiceId** | **string** | ID of a Service | 
**Status** | **string** | Status of the version set | 
**UpdatedAt** | Pointer to **time.Time** | When the version was last updated | [optional] 
**Version** | **string** | Version of the tier version set spec | 

## Methods

### NewTierVersionSetSpecSummary

`func NewTierVersionSetSpecSummary(hasSpecFile bool, productTierId string, serviceId string, status string, version string, ) *TierVersionSetSpecSummary`

NewTierVersionSetSpecSummary instantiates a new TierVersionSetSpecSummary object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewTierVersionSetSpecSummaryWithDefaults

`func NewTierVersionSetSpecSummaryWithDefaults() *TierVersionSetSpecSummary`

NewTierVersionSetSpecSummaryWithDefaults instantiates a new TierVersionSetSpecSummary object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetCreatedAt

`func (o *TierVersionSetSpecSummary) GetCreatedAt() time.Time`

GetCreatedAt returns the CreatedAt field if non-nil, zero value otherwise.

### GetCreatedAtOk

`func (o *TierVersionSetSpecSummary) GetCreatedAtOk() (*time.Time, bool)`

GetCreatedAtOk returns a tuple with the CreatedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCreatedAt

`func (o *TierVersionSetSpecSummary) SetCreatedAt(v time.Time)`

SetCreatedAt sets CreatedAt field to given value.

### HasCreatedAt

`func (o *TierVersionSetSpecSummary) HasCreatedAt() bool`

HasCreatedAt returns a boolean if a field has been set.

### GetHasSpecFile

`func (o *TierVersionSetSpecSummary) GetHasSpecFile() bool`

GetHasSpecFile returns the HasSpecFile field if non-nil, zero value otherwise.

### GetHasSpecFileOk

`func (o *TierVersionSetSpecSummary) GetHasSpecFileOk() (*bool, bool)`

GetHasSpecFileOk returns a tuple with the HasSpecFile field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetHasSpecFile

`func (o *TierVersionSetSpecSummary) SetHasSpecFile(v bool)`

SetHasSpecFile sets HasSpecFile field to given value.


### GetProductTierId

`func (o *TierVersionSetSpecSummary) GetProductTierId() string`

GetProductTierId returns the ProductTierId field if non-nil, zero value otherwise.

### GetProductTierIdOk

`func (o *TierVersionSetSpecSummary) GetProductTierIdOk() (*string, bool)`

GetProductTierIdOk returns a tuple with the ProductTierId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetProductTierId

`func (o *TierVersionSetSpecSummary) SetProductTierId(v string)`

SetProductTierId sets ProductTierId field to given value.


### GetServiceId

`func (o *TierVersionSetSpecSummary) GetServiceId() string`

GetServiceId returns the ServiceId field if non-nil, zero value otherwise.

### GetServiceIdOk

`func (o *TierVersionSetSpecSummary) GetServiceIdOk() (*string, bool)`

GetServiceIdOk returns a tuple with the ServiceId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetServiceId

`func (o *TierVersionSetSpecSummary) SetServiceId(v string)`

SetServiceId sets ServiceId field to given value.


### GetStatus

`func (o *TierVersionSetSpecSummary) GetStatus() string`

GetStatus returns the Status field if non-nil, zero value otherwise.

### GetStatusOk

`func (o *TierVersionSetSpecSummary) GetStatusOk() (*string, bool)`

GetStatusOk returns a tuple with the Status field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStatus

`func (o *TierVersionSetSpecSummary) SetStatus(v string)`

SetStatus sets Status field to given value.


### GetUpdatedAt

`func (o *TierVersionSetSpecSummary) GetUpdatedAt() time.Time`

GetUpdatedAt returns the UpdatedAt field if non-nil, zero value otherwise.

### GetUpdatedAtOk

`func (o *TierVersionSetSpecSummary) GetUpdatedAtOk() (*time.Time, bool)`

GetUpdatedAtOk returns a tuple with the UpdatedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUpdatedAt

`func (o *TierVersionSetSpecSummary) SetUpdatedAt(v time.Time)`

SetUpdatedAt sets UpdatedAt field to given value.

### HasUpdatedAt

`func (o *TierVersionSetSpecSummary) HasUpdatedAt() bool`

HasUpdatedAt returns a boolean if a field has been set.

### GetVersion

`func (o *TierVersionSetSpecSummary) GetVersion() string`

GetVersion returns the Version field if non-nil, zero value otherwise.

### GetVersionOk

`func (o *TierVersionSetSpecSummary) GetVersionOk() (*string, bool)`

GetVersionOk returns a tuple with the Version field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVersion

`func (o *TierVersionSetSpecSummary) SetVersion(v string)`

SetVersion sets Version field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


