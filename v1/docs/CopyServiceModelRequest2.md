# CopyServiceModelRequest2

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**AccountConfigIds** | Pointer to **[]string** | The infrastructure account configuration ID list | [optional] 
**Description** | **string** | A brief description of the service model | 
**Features** | Pointer to [**[]ServiceModelFeatureDetail**](ServiceModelFeatureDetail.md) | Enabled service model features | [optional] 
**Name** | **string** | Name of the Service Model | 
**TargetServiceModelType** | **string** | The model type encapsulating this service | 
**TargetTierType** | Pointer to **string** | Target product tier type | [optional] 

## Methods

### NewCopyServiceModelRequest2

`func NewCopyServiceModelRequest2(description string, name string, targetServiceModelType string, ) *CopyServiceModelRequest2`

NewCopyServiceModelRequest2 instantiates a new CopyServiceModelRequest2 object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewCopyServiceModelRequest2WithDefaults

`func NewCopyServiceModelRequest2WithDefaults() *CopyServiceModelRequest2`

NewCopyServiceModelRequest2WithDefaults instantiates a new CopyServiceModelRequest2 object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetAccountConfigIds

`func (o *CopyServiceModelRequest2) GetAccountConfigIds() []string`

GetAccountConfigIds returns the AccountConfigIds field if non-nil, zero value otherwise.

### GetAccountConfigIdsOk

`func (o *CopyServiceModelRequest2) GetAccountConfigIdsOk() (*[]string, bool)`

GetAccountConfigIdsOk returns a tuple with the AccountConfigIds field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAccountConfigIds

`func (o *CopyServiceModelRequest2) SetAccountConfigIds(v []string)`

SetAccountConfigIds sets AccountConfigIds field to given value.

### HasAccountConfigIds

`func (o *CopyServiceModelRequest2) HasAccountConfigIds() bool`

HasAccountConfigIds returns a boolean if a field has been set.

### GetDescription

`func (o *CopyServiceModelRequest2) GetDescription() string`

GetDescription returns the Description field if non-nil, zero value otherwise.

### GetDescriptionOk

`func (o *CopyServiceModelRequest2) GetDescriptionOk() (*string, bool)`

GetDescriptionOk returns a tuple with the Description field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDescription

`func (o *CopyServiceModelRequest2) SetDescription(v string)`

SetDescription sets Description field to given value.


### GetFeatures

`func (o *CopyServiceModelRequest2) GetFeatures() []ServiceModelFeatureDetail`

GetFeatures returns the Features field if non-nil, zero value otherwise.

### GetFeaturesOk

`func (o *CopyServiceModelRequest2) GetFeaturesOk() (*[]ServiceModelFeatureDetail, bool)`

GetFeaturesOk returns a tuple with the Features field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFeatures

`func (o *CopyServiceModelRequest2) SetFeatures(v []ServiceModelFeatureDetail)`

SetFeatures sets Features field to given value.

### HasFeatures

`func (o *CopyServiceModelRequest2) HasFeatures() bool`

HasFeatures returns a boolean if a field has been set.

### GetName

`func (o *CopyServiceModelRequest2) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *CopyServiceModelRequest2) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *CopyServiceModelRequest2) SetName(v string)`

SetName sets Name field to given value.


### GetTargetServiceModelType

`func (o *CopyServiceModelRequest2) GetTargetServiceModelType() string`

GetTargetServiceModelType returns the TargetServiceModelType field if non-nil, zero value otherwise.

### GetTargetServiceModelTypeOk

`func (o *CopyServiceModelRequest2) GetTargetServiceModelTypeOk() (*string, bool)`

GetTargetServiceModelTypeOk returns a tuple with the TargetServiceModelType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTargetServiceModelType

`func (o *CopyServiceModelRequest2) SetTargetServiceModelType(v string)`

SetTargetServiceModelType sets TargetServiceModelType field to given value.


### GetTargetTierType

`func (o *CopyServiceModelRequest2) GetTargetTierType() string`

GetTargetTierType returns the TargetTierType field if non-nil, zero value otherwise.

### GetTargetTierTypeOk

`func (o *CopyServiceModelRequest2) GetTargetTierTypeOk() (*string, bool)`

GetTargetTierTypeOk returns a tuple with the TargetTierType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTargetTierType

`func (o *CopyServiceModelRequest2) SetTargetTierType(v string)`

SetTargetTierType sets TargetTierType field to given value.

### HasTargetTierType

`func (o *CopyServiceModelRequest2) HasTargetTierType() bool`

HasTargetTierType returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


