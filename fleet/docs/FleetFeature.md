# FleetFeature

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Feature** | **string** | The feature to enable. | 
**FeatureConfig** | Pointer to **map[string]string** | The configuration for the feature. | [optional] 
**Status** | **string** | The status of the feature. | 

## Methods

### NewFleetFeature

`func NewFleetFeature(feature string, status string, ) *FleetFeature`

NewFleetFeature instantiates a new FleetFeature object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewFleetFeatureWithDefaults

`func NewFleetFeatureWithDefaults() *FleetFeature`

NewFleetFeatureWithDefaults instantiates a new FleetFeature object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetFeature

`func (o *FleetFeature) GetFeature() string`

GetFeature returns the Feature field if non-nil, zero value otherwise.

### GetFeatureOk

`func (o *FleetFeature) GetFeatureOk() (*string, bool)`

GetFeatureOk returns a tuple with the Feature field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFeature

`func (o *FleetFeature) SetFeature(v string)`

SetFeature sets Feature field to given value.


### GetFeatureConfig

`func (o *FleetFeature) GetFeatureConfig() map[string]string`

GetFeatureConfig returns the FeatureConfig field if non-nil, zero value otherwise.

### GetFeatureConfigOk

`func (o *FleetFeature) GetFeatureConfigOk() (*map[string]string, bool)`

GetFeatureConfigOk returns a tuple with the FeatureConfig field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFeatureConfig

`func (o *FleetFeature) SetFeatureConfig(v map[string]string)`

SetFeatureConfig sets FeatureConfig field to given value.

### HasFeatureConfig

`func (o *FleetFeature) HasFeatureConfig() bool`

HasFeatureConfig returns a boolean if a field has been set.

### GetStatus

`func (o *FleetFeature) GetStatus() string`

GetStatus returns the Status field if non-nil, zero value otherwise.

### GetStatusOk

`func (o *FleetFeature) GetStatusOk() (*string, bool)`

GetStatusOk returns a tuple with the Status field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStatus

`func (o *FleetFeature) SetStatus(v string)`

SetStatus sets Status field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


