# CreateServiceModelRequest2

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**AccountConfigIds** | Pointer to **[]string** | The infrastructure account configuration ID list | [optional] 
**Description** | **string** | A brief description of the service model | 
**Features** | Pointer to [**[]ServiceModelFeatureDetail**](ServiceModelFeatureDetail.md) | Enabled service model features | [optional] 
**ModelType** | **string** | The model type encapsulating this service | 
**Name** | **string** | Name of the Service Model | 
**ServiceApiId** | **string** | The service API this model is for | 

## Methods

### NewCreateServiceModelRequest2

`func NewCreateServiceModelRequest2(description string, modelType string, name string, serviceApiId string, ) *CreateServiceModelRequest2`

NewCreateServiceModelRequest2 instantiates a new CreateServiceModelRequest2 object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewCreateServiceModelRequest2WithDefaults

`func NewCreateServiceModelRequest2WithDefaults() *CreateServiceModelRequest2`

NewCreateServiceModelRequest2WithDefaults instantiates a new CreateServiceModelRequest2 object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetAccountConfigIds

`func (o *CreateServiceModelRequest2) GetAccountConfigIds() []string`

GetAccountConfigIds returns the AccountConfigIds field if non-nil, zero value otherwise.

### GetAccountConfigIdsOk

`func (o *CreateServiceModelRequest2) GetAccountConfigIdsOk() (*[]string, bool)`

GetAccountConfigIdsOk returns a tuple with the AccountConfigIds field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAccountConfigIds

`func (o *CreateServiceModelRequest2) SetAccountConfigIds(v []string)`

SetAccountConfigIds sets AccountConfigIds field to given value.

### HasAccountConfigIds

`func (o *CreateServiceModelRequest2) HasAccountConfigIds() bool`

HasAccountConfigIds returns a boolean if a field has been set.

### GetDescription

`func (o *CreateServiceModelRequest2) GetDescription() string`

GetDescription returns the Description field if non-nil, zero value otherwise.

### GetDescriptionOk

`func (o *CreateServiceModelRequest2) GetDescriptionOk() (*string, bool)`

GetDescriptionOk returns a tuple with the Description field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDescription

`func (o *CreateServiceModelRequest2) SetDescription(v string)`

SetDescription sets Description field to given value.


### GetFeatures

`func (o *CreateServiceModelRequest2) GetFeatures() []ServiceModelFeatureDetail`

GetFeatures returns the Features field if non-nil, zero value otherwise.

### GetFeaturesOk

`func (o *CreateServiceModelRequest2) GetFeaturesOk() (*[]ServiceModelFeatureDetail, bool)`

GetFeaturesOk returns a tuple with the Features field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFeatures

`func (o *CreateServiceModelRequest2) SetFeatures(v []ServiceModelFeatureDetail)`

SetFeatures sets Features field to given value.

### HasFeatures

`func (o *CreateServiceModelRequest2) HasFeatures() bool`

HasFeatures returns a boolean if a field has been set.

### GetModelType

`func (o *CreateServiceModelRequest2) GetModelType() string`

GetModelType returns the ModelType field if non-nil, zero value otherwise.

### GetModelTypeOk

`func (o *CreateServiceModelRequest2) GetModelTypeOk() (*string, bool)`

GetModelTypeOk returns a tuple with the ModelType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetModelType

`func (o *CreateServiceModelRequest2) SetModelType(v string)`

SetModelType sets ModelType field to given value.


### GetName

`func (o *CreateServiceModelRequest2) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *CreateServiceModelRequest2) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *CreateServiceModelRequest2) SetName(v string)`

SetName sets Name field to given value.


### GetServiceApiId

`func (o *CreateServiceModelRequest2) GetServiceApiId() string`

GetServiceApiId returns the ServiceApiId field if non-nil, zero value otherwise.

### GetServiceApiIdOk

`func (o *CreateServiceModelRequest2) GetServiceApiIdOk() (*string, bool)`

GetServiceApiIdOk returns a tuple with the ServiceApiId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetServiceApiId

`func (o *CreateServiceModelRequest2) SetServiceApiId(v string)`

SetServiceApiId sets ServiceApiId field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


