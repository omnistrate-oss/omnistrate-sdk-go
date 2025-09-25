# GetTierVersionSetSpecResult

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**CreatedAt** | Pointer to **time.Time** | When the version was created | [optional] 
**ProductTierId** | **string** | ID of a Product Tier | 
**ServiceId** | **string** | ID of a Service | 
**SpecContent** | **string** | JSON content containing the tier version set specification with raw YAML, configs, and secrets | 
**UpdatedAt** | Pointer to **time.Time** | When the version was last updated | [optional] 
**Version** | **string** | Version of the tier version set spec | 

## Methods

### NewGetTierVersionSetSpecResult

`func NewGetTierVersionSetSpecResult(productTierId string, serviceId string, specContent string, version string, ) *GetTierVersionSetSpecResult`

NewGetTierVersionSetSpecResult instantiates a new GetTierVersionSetSpecResult object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewGetTierVersionSetSpecResultWithDefaults

`func NewGetTierVersionSetSpecResultWithDefaults() *GetTierVersionSetSpecResult`

NewGetTierVersionSetSpecResultWithDefaults instantiates a new GetTierVersionSetSpecResult object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetCreatedAt

`func (o *GetTierVersionSetSpecResult) GetCreatedAt() time.Time`

GetCreatedAt returns the CreatedAt field if non-nil, zero value otherwise.

### GetCreatedAtOk

`func (o *GetTierVersionSetSpecResult) GetCreatedAtOk() (*time.Time, bool)`

GetCreatedAtOk returns a tuple with the CreatedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCreatedAt

`func (o *GetTierVersionSetSpecResult) SetCreatedAt(v time.Time)`

SetCreatedAt sets CreatedAt field to given value.

### HasCreatedAt

`func (o *GetTierVersionSetSpecResult) HasCreatedAt() bool`

HasCreatedAt returns a boolean if a field has been set.

### GetProductTierId

`func (o *GetTierVersionSetSpecResult) GetProductTierId() string`

GetProductTierId returns the ProductTierId field if non-nil, zero value otherwise.

### GetProductTierIdOk

`func (o *GetTierVersionSetSpecResult) GetProductTierIdOk() (*string, bool)`

GetProductTierIdOk returns a tuple with the ProductTierId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetProductTierId

`func (o *GetTierVersionSetSpecResult) SetProductTierId(v string)`

SetProductTierId sets ProductTierId field to given value.


### GetServiceId

`func (o *GetTierVersionSetSpecResult) GetServiceId() string`

GetServiceId returns the ServiceId field if non-nil, zero value otherwise.

### GetServiceIdOk

`func (o *GetTierVersionSetSpecResult) GetServiceIdOk() (*string, bool)`

GetServiceIdOk returns a tuple with the ServiceId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetServiceId

`func (o *GetTierVersionSetSpecResult) SetServiceId(v string)`

SetServiceId sets ServiceId field to given value.


### GetSpecContent

`func (o *GetTierVersionSetSpecResult) GetSpecContent() string`

GetSpecContent returns the SpecContent field if non-nil, zero value otherwise.

### GetSpecContentOk

`func (o *GetTierVersionSetSpecResult) GetSpecContentOk() (*string, bool)`

GetSpecContentOk returns a tuple with the SpecContent field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSpecContent

`func (o *GetTierVersionSetSpecResult) SetSpecContent(v string)`

SetSpecContent sets SpecContent field to given value.


### GetUpdatedAt

`func (o *GetTierVersionSetSpecResult) GetUpdatedAt() time.Time`

GetUpdatedAt returns the UpdatedAt field if non-nil, zero value otherwise.

### GetUpdatedAtOk

`func (o *GetTierVersionSetSpecResult) GetUpdatedAtOk() (*time.Time, bool)`

GetUpdatedAtOk returns a tuple with the UpdatedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUpdatedAt

`func (o *GetTierVersionSetSpecResult) SetUpdatedAt(v time.Time)`

SetUpdatedAt sets UpdatedAt field to given value.

### HasUpdatedAt

`func (o *GetTierVersionSetSpecResult) HasUpdatedAt() bool`

HasUpdatedAt returns a boolean if a field has been set.

### GetVersion

`func (o *GetTierVersionSetSpecResult) GetVersion() string`

GetVersion returns the Version field if non-nil, zero value otherwise.

### GetVersionOk

`func (o *GetTierVersionSetSpecResult) GetVersionOk() (*string, bool)`

GetVersionOk returns a tuple with the Version field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVersion

`func (o *GetTierVersionSetSpecResult) SetVersion(v string)`

SetVersion sets Version field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


