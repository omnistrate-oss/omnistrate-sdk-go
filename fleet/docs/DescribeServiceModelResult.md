# DescribeServiceModelResult

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**AccountConfigIds** | Pointer to **[]string** | The infrastructure account configuration ID list | [optional] 
**ActiveAccountConfigIds** | Pointer to **map[string]interface{}** | The active infrastructure account configuration IDs per cloud provider | [optional] 
**Description** | **string** | A brief description of the service model | 
**Features** | Pointer to [**[]ServiceModelFeatureDetail**](ServiceModelFeatureDetail.md) | Enabled service model features | [optional] 
**Id** | **string** | ID of a Service Model | 
**Key** | **string** | The unique key for this service model | 
**ModelType** | **string** | The model type encapsulating this service | 
**Name** | **string** | Name of the Service Model | 
**ProductTiers** | Pointer to **[]string** | The product tiers associated with this service model | [optional] 
**ServiceApiId** | **string** | ID of a Service API | 
**ServiceId** | **string** | ID of a Service | 

## Methods

### NewDescribeServiceModelResult

`func NewDescribeServiceModelResult(description string, id string, key string, modelType string, name string, serviceApiId string, serviceId string, ) *DescribeServiceModelResult`

NewDescribeServiceModelResult instantiates a new DescribeServiceModelResult object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewDescribeServiceModelResultWithDefaults

`func NewDescribeServiceModelResultWithDefaults() *DescribeServiceModelResult`

NewDescribeServiceModelResultWithDefaults instantiates a new DescribeServiceModelResult object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetAccountConfigIds

`func (o *DescribeServiceModelResult) GetAccountConfigIds() []string`

GetAccountConfigIds returns the AccountConfigIds field if non-nil, zero value otherwise.

### GetAccountConfigIdsOk

`func (o *DescribeServiceModelResult) GetAccountConfigIdsOk() (*[]string, bool)`

GetAccountConfigIdsOk returns a tuple with the AccountConfigIds field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAccountConfigIds

`func (o *DescribeServiceModelResult) SetAccountConfigIds(v []string)`

SetAccountConfigIds sets AccountConfigIds field to given value.

### HasAccountConfigIds

`func (o *DescribeServiceModelResult) HasAccountConfigIds() bool`

HasAccountConfigIds returns a boolean if a field has been set.

### GetActiveAccountConfigIds

`func (o *DescribeServiceModelResult) GetActiveAccountConfigIds() map[string]interface{}`

GetActiveAccountConfigIds returns the ActiveAccountConfigIds field if non-nil, zero value otherwise.

### GetActiveAccountConfigIdsOk

`func (o *DescribeServiceModelResult) GetActiveAccountConfigIdsOk() (*map[string]interface{}, bool)`

GetActiveAccountConfigIdsOk returns a tuple with the ActiveAccountConfigIds field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetActiveAccountConfigIds

`func (o *DescribeServiceModelResult) SetActiveAccountConfigIds(v map[string]interface{})`

SetActiveAccountConfigIds sets ActiveAccountConfigIds field to given value.

### HasActiveAccountConfigIds

`func (o *DescribeServiceModelResult) HasActiveAccountConfigIds() bool`

HasActiveAccountConfigIds returns a boolean if a field has been set.

### GetDescription

`func (o *DescribeServiceModelResult) GetDescription() string`

GetDescription returns the Description field if non-nil, zero value otherwise.

### GetDescriptionOk

`func (o *DescribeServiceModelResult) GetDescriptionOk() (*string, bool)`

GetDescriptionOk returns a tuple with the Description field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDescription

`func (o *DescribeServiceModelResult) SetDescription(v string)`

SetDescription sets Description field to given value.


### GetFeatures

`func (o *DescribeServiceModelResult) GetFeatures() []ServiceModelFeatureDetail`

GetFeatures returns the Features field if non-nil, zero value otherwise.

### GetFeaturesOk

`func (o *DescribeServiceModelResult) GetFeaturesOk() (*[]ServiceModelFeatureDetail, bool)`

GetFeaturesOk returns a tuple with the Features field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFeatures

`func (o *DescribeServiceModelResult) SetFeatures(v []ServiceModelFeatureDetail)`

SetFeatures sets Features field to given value.

### HasFeatures

`func (o *DescribeServiceModelResult) HasFeatures() bool`

HasFeatures returns a boolean if a field has been set.

### GetId

`func (o *DescribeServiceModelResult) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *DescribeServiceModelResult) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *DescribeServiceModelResult) SetId(v string)`

SetId sets Id field to given value.


### GetKey

`func (o *DescribeServiceModelResult) GetKey() string`

GetKey returns the Key field if non-nil, zero value otherwise.

### GetKeyOk

`func (o *DescribeServiceModelResult) GetKeyOk() (*string, bool)`

GetKeyOk returns a tuple with the Key field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetKey

`func (o *DescribeServiceModelResult) SetKey(v string)`

SetKey sets Key field to given value.


### GetModelType

`func (o *DescribeServiceModelResult) GetModelType() string`

GetModelType returns the ModelType field if non-nil, zero value otherwise.

### GetModelTypeOk

`func (o *DescribeServiceModelResult) GetModelTypeOk() (*string, bool)`

GetModelTypeOk returns a tuple with the ModelType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetModelType

`func (o *DescribeServiceModelResult) SetModelType(v string)`

SetModelType sets ModelType field to given value.


### GetName

`func (o *DescribeServiceModelResult) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *DescribeServiceModelResult) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *DescribeServiceModelResult) SetName(v string)`

SetName sets Name field to given value.


### GetProductTiers

`func (o *DescribeServiceModelResult) GetProductTiers() []string`

GetProductTiers returns the ProductTiers field if non-nil, zero value otherwise.

### GetProductTiersOk

`func (o *DescribeServiceModelResult) GetProductTiersOk() (*[]string, bool)`

GetProductTiersOk returns a tuple with the ProductTiers field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetProductTiers

`func (o *DescribeServiceModelResult) SetProductTiers(v []string)`

SetProductTiers sets ProductTiers field to given value.

### HasProductTiers

`func (o *DescribeServiceModelResult) HasProductTiers() bool`

HasProductTiers returns a boolean if a field has been set.

### GetServiceApiId

`func (o *DescribeServiceModelResult) GetServiceApiId() string`

GetServiceApiId returns the ServiceApiId field if non-nil, zero value otherwise.

### GetServiceApiIdOk

`func (o *DescribeServiceModelResult) GetServiceApiIdOk() (*string, bool)`

GetServiceApiIdOk returns a tuple with the ServiceApiId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetServiceApiId

`func (o *DescribeServiceModelResult) SetServiceApiId(v string)`

SetServiceApiId sets ServiceApiId field to given value.


### GetServiceId

`func (o *DescribeServiceModelResult) GetServiceId() string`

GetServiceId returns the ServiceId field if non-nil, zero value otherwise.

### GetServiceIdOk

`func (o *DescribeServiceModelResult) GetServiceIdOk() (*string, bool)`

GetServiceIdOk returns a tuple with the ServiceId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetServiceId

`func (o *DescribeServiceModelResult) SetServiceId(v string)`

SetServiceId sets ServiceId field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


