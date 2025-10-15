# TierVersionSetMetadata

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**ProductTierId** | **string** | ID of a Product Tier | 
**ReleaseNotes** | Pointer to **string** | Release notes for the version set. | [optional] 
**ServiceId** | **string** | ID of a Service | 
**Version** | **string** | The version number for the specific version set. | 

## Methods

### NewTierVersionSetMetadata

`func NewTierVersionSetMetadata(productTierId string, serviceId string, version string, ) *TierVersionSetMetadata`

NewTierVersionSetMetadata instantiates a new TierVersionSetMetadata object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewTierVersionSetMetadataWithDefaults

`func NewTierVersionSetMetadataWithDefaults() *TierVersionSetMetadata`

NewTierVersionSetMetadataWithDefaults instantiates a new TierVersionSetMetadata object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetProductTierId

`func (o *TierVersionSetMetadata) GetProductTierId() string`

GetProductTierId returns the ProductTierId field if non-nil, zero value otherwise.

### GetProductTierIdOk

`func (o *TierVersionSetMetadata) GetProductTierIdOk() (*string, bool)`

GetProductTierIdOk returns a tuple with the ProductTierId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetProductTierId

`func (o *TierVersionSetMetadata) SetProductTierId(v string)`

SetProductTierId sets ProductTierId field to given value.


### GetReleaseNotes

`func (o *TierVersionSetMetadata) GetReleaseNotes() string`

GetReleaseNotes returns the ReleaseNotes field if non-nil, zero value otherwise.

### GetReleaseNotesOk

`func (o *TierVersionSetMetadata) GetReleaseNotesOk() (*string, bool)`

GetReleaseNotesOk returns a tuple with the ReleaseNotes field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetReleaseNotes

`func (o *TierVersionSetMetadata) SetReleaseNotes(v string)`

SetReleaseNotes sets ReleaseNotes field to given value.

### HasReleaseNotes

`func (o *TierVersionSetMetadata) HasReleaseNotes() bool`

HasReleaseNotes returns a boolean if a field has been set.

### GetServiceId

`func (o *TierVersionSetMetadata) GetServiceId() string`

GetServiceId returns the ServiceId field if non-nil, zero value otherwise.

### GetServiceIdOk

`func (o *TierVersionSetMetadata) GetServiceIdOk() (*string, bool)`

GetServiceIdOk returns a tuple with the ServiceId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetServiceId

`func (o *TierVersionSetMetadata) SetServiceId(v string)`

SetServiceId sets ServiceId field to given value.


### GetVersion

`func (o *TierVersionSetMetadata) GetVersion() string`

GetVersion returns the Version field if non-nil, zero value otherwise.

### GetVersionOk

`func (o *TierVersionSetMetadata) GetVersionOk() (*string, bool)`

GetVersionOk returns a tuple with the Version field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVersion

`func (o *TierVersionSetMetadata) SetVersion(v string)`

SetVersion sets Version field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


