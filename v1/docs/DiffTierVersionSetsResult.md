# DiffTierVersionSetsResult

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**AnotherVersion** | **string** | The target version to compare against. | 
**ProductTierId** | **string** | ID of a Product Tier | 
**ResourceChangeSets** | **map[string]interface{}** | The difference for the service API configuration per resource | 
**ServiceId** | **string** | ID of a Service | 
**Version** | **string** | The version number for the version set. | 

## Methods

### NewDiffTierVersionSetsResult

`func NewDiffTierVersionSetsResult(anotherVersion string, productTierId string, resourceChangeSets map[string]interface{}, serviceId string, version string, ) *DiffTierVersionSetsResult`

NewDiffTierVersionSetsResult instantiates a new DiffTierVersionSetsResult object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewDiffTierVersionSetsResultWithDefaults

`func NewDiffTierVersionSetsResultWithDefaults() *DiffTierVersionSetsResult`

NewDiffTierVersionSetsResultWithDefaults instantiates a new DiffTierVersionSetsResult object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetAnotherVersion

`func (o *DiffTierVersionSetsResult) GetAnotherVersion() string`

GetAnotherVersion returns the AnotherVersion field if non-nil, zero value otherwise.

### GetAnotherVersionOk

`func (o *DiffTierVersionSetsResult) GetAnotherVersionOk() (*string, bool)`

GetAnotherVersionOk returns a tuple with the AnotherVersion field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAnotherVersion

`func (o *DiffTierVersionSetsResult) SetAnotherVersion(v string)`

SetAnotherVersion sets AnotherVersion field to given value.


### GetProductTierId

`func (o *DiffTierVersionSetsResult) GetProductTierId() string`

GetProductTierId returns the ProductTierId field if non-nil, zero value otherwise.

### GetProductTierIdOk

`func (o *DiffTierVersionSetsResult) GetProductTierIdOk() (*string, bool)`

GetProductTierIdOk returns a tuple with the ProductTierId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetProductTierId

`func (o *DiffTierVersionSetsResult) SetProductTierId(v string)`

SetProductTierId sets ProductTierId field to given value.


### GetResourceChangeSets

`func (o *DiffTierVersionSetsResult) GetResourceChangeSets() map[string]interface{}`

GetResourceChangeSets returns the ResourceChangeSets field if non-nil, zero value otherwise.

### GetResourceChangeSetsOk

`func (o *DiffTierVersionSetsResult) GetResourceChangeSetsOk() (*map[string]interface{}, bool)`

GetResourceChangeSetsOk returns a tuple with the ResourceChangeSets field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetResourceChangeSets

`func (o *DiffTierVersionSetsResult) SetResourceChangeSets(v map[string]interface{})`

SetResourceChangeSets sets ResourceChangeSets field to given value.


### GetServiceId

`func (o *DiffTierVersionSetsResult) GetServiceId() string`

GetServiceId returns the ServiceId field if non-nil, zero value otherwise.

### GetServiceIdOk

`func (o *DiffTierVersionSetsResult) GetServiceIdOk() (*string, bool)`

GetServiceIdOk returns a tuple with the ServiceId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetServiceId

`func (o *DiffTierVersionSetsResult) SetServiceId(v string)`

SetServiceId sets ServiceId field to given value.


### GetVersion

`func (o *DiffTierVersionSetsResult) GetVersion() string`

GetVersion returns the Version field if non-nil, zero value otherwise.

### GetVersionOk

`func (o *DiffTierVersionSetsResult) GetVersionOk() (*string, bool)`

GetVersionOk returns a tuple with the Version field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVersion

`func (o *DiffTierVersionSetsResult) SetVersion(v string)`

SetVersion sets Version field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


