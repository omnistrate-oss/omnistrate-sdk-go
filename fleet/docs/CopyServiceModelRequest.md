# CopyServiceModelRequest

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**AccountConfigIds** | Pointer to **[]string** | The infrastructure account configuration ID list | [optional] 
**Description** | **string** | A brief description of the service model | 
**Features** | Pointer to [**[]ServiceModelFeatureDetail**](ServiceModelFeatureDetail.md) | Enabled service model features | [optional] 
**Name** | **string** | Name of the Service Model | 
**ServiceId** | **string** | ID of a Service | 
**SourceId** | **string** | ID of a Service Model | 
**TargetServiceModelType** | **string** | The model type encapsulating this service | 
**TargetTierType** | Pointer to **string** | ProductTierType is the type of tier for a product | [optional] 
**Token** | **string** | JWT token used to perform authorization | 

## Methods

### NewCopyServiceModelRequest

`func NewCopyServiceModelRequest(description string, name string, serviceId string, sourceId string, targetServiceModelType string, token string, ) *CopyServiceModelRequest`

NewCopyServiceModelRequest instantiates a new CopyServiceModelRequest object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewCopyServiceModelRequestWithDefaults

`func NewCopyServiceModelRequestWithDefaults() *CopyServiceModelRequest`

NewCopyServiceModelRequestWithDefaults instantiates a new CopyServiceModelRequest object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetAccountConfigIds

`func (o *CopyServiceModelRequest) GetAccountConfigIds() []string`

GetAccountConfigIds returns the AccountConfigIds field if non-nil, zero value otherwise.

### GetAccountConfigIdsOk

`func (o *CopyServiceModelRequest) GetAccountConfigIdsOk() (*[]string, bool)`

GetAccountConfigIdsOk returns a tuple with the AccountConfigIds field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAccountConfigIds

`func (o *CopyServiceModelRequest) SetAccountConfigIds(v []string)`

SetAccountConfigIds sets AccountConfigIds field to given value.

### HasAccountConfigIds

`func (o *CopyServiceModelRequest) HasAccountConfigIds() bool`

HasAccountConfigIds returns a boolean if a field has been set.

### GetDescription

`func (o *CopyServiceModelRequest) GetDescription() string`

GetDescription returns the Description field if non-nil, zero value otherwise.

### GetDescriptionOk

`func (o *CopyServiceModelRequest) GetDescriptionOk() (*string, bool)`

GetDescriptionOk returns a tuple with the Description field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDescription

`func (o *CopyServiceModelRequest) SetDescription(v string)`

SetDescription sets Description field to given value.


### GetFeatures

`func (o *CopyServiceModelRequest) GetFeatures() []ServiceModelFeatureDetail`

GetFeatures returns the Features field if non-nil, zero value otherwise.

### GetFeaturesOk

`func (o *CopyServiceModelRequest) GetFeaturesOk() (*[]ServiceModelFeatureDetail, bool)`

GetFeaturesOk returns a tuple with the Features field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFeatures

`func (o *CopyServiceModelRequest) SetFeatures(v []ServiceModelFeatureDetail)`

SetFeatures sets Features field to given value.

### HasFeatures

`func (o *CopyServiceModelRequest) HasFeatures() bool`

HasFeatures returns a boolean if a field has been set.

### GetName

`func (o *CopyServiceModelRequest) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *CopyServiceModelRequest) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *CopyServiceModelRequest) SetName(v string)`

SetName sets Name field to given value.


### GetServiceId

`func (o *CopyServiceModelRequest) GetServiceId() string`

GetServiceId returns the ServiceId field if non-nil, zero value otherwise.

### GetServiceIdOk

`func (o *CopyServiceModelRequest) GetServiceIdOk() (*string, bool)`

GetServiceIdOk returns a tuple with the ServiceId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetServiceId

`func (o *CopyServiceModelRequest) SetServiceId(v string)`

SetServiceId sets ServiceId field to given value.


### GetSourceId

`func (o *CopyServiceModelRequest) GetSourceId() string`

GetSourceId returns the SourceId field if non-nil, zero value otherwise.

### GetSourceIdOk

`func (o *CopyServiceModelRequest) GetSourceIdOk() (*string, bool)`

GetSourceIdOk returns a tuple with the SourceId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSourceId

`func (o *CopyServiceModelRequest) SetSourceId(v string)`

SetSourceId sets SourceId field to given value.


### GetTargetServiceModelType

`func (o *CopyServiceModelRequest) GetTargetServiceModelType() string`

GetTargetServiceModelType returns the TargetServiceModelType field if non-nil, zero value otherwise.

### GetTargetServiceModelTypeOk

`func (o *CopyServiceModelRequest) GetTargetServiceModelTypeOk() (*string, bool)`

GetTargetServiceModelTypeOk returns a tuple with the TargetServiceModelType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTargetServiceModelType

`func (o *CopyServiceModelRequest) SetTargetServiceModelType(v string)`

SetTargetServiceModelType sets TargetServiceModelType field to given value.


### GetTargetTierType

`func (o *CopyServiceModelRequest) GetTargetTierType() string`

GetTargetTierType returns the TargetTierType field if non-nil, zero value otherwise.

### GetTargetTierTypeOk

`func (o *CopyServiceModelRequest) GetTargetTierTypeOk() (*string, bool)`

GetTargetTierTypeOk returns a tuple with the TargetTierType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTargetTierType

`func (o *CopyServiceModelRequest) SetTargetTierType(v string)`

SetTargetTierType sets TargetTierType field to given value.

### HasTargetTierType

`func (o *CopyServiceModelRequest) HasTargetTierType() bool`

HasTargetTierType returns a boolean if a field has been set.

### GetToken

`func (o *CopyServiceModelRequest) GetToken() string`

GetToken returns the Token field if non-nil, zero value otherwise.

### GetTokenOk

`func (o *CopyServiceModelRequest) GetTokenOk() (*string, bool)`

GetTokenOk returns a tuple with the Token field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetToken

`func (o *CopyServiceModelRequest) SetToken(v string)`

SetToken sets Token field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


