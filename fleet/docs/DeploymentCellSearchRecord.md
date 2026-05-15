# DeploymentCellSearchRecord

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**CloudProvider** | **string** | Name of the Infra Provider | 
**CustomTags** | Pointer to [**[]CustomTag**](CustomTag.md) | The custom tags associated with resource instances in the deployment cell. | [optional] 
**Description** | **string** | The deployment cell description. | 
**Id** | **string** | The deployment cell ID. | 
**RegionCode** | **string** | The region code of the deployment cell. | 

## Methods

### NewDeploymentCellSearchRecord

`func NewDeploymentCellSearchRecord(cloudProvider string, description string, id string, regionCode string, ) *DeploymentCellSearchRecord`

NewDeploymentCellSearchRecord instantiates a new DeploymentCellSearchRecord object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewDeploymentCellSearchRecordWithDefaults

`func NewDeploymentCellSearchRecordWithDefaults() *DeploymentCellSearchRecord`

NewDeploymentCellSearchRecordWithDefaults instantiates a new DeploymentCellSearchRecord object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetCloudProvider

`func (o *DeploymentCellSearchRecord) GetCloudProvider() string`

GetCloudProvider returns the CloudProvider field if non-nil, zero value otherwise.

### GetCloudProviderOk

`func (o *DeploymentCellSearchRecord) GetCloudProviderOk() (*string, bool)`

GetCloudProviderOk returns a tuple with the CloudProvider field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCloudProvider

`func (o *DeploymentCellSearchRecord) SetCloudProvider(v string)`

SetCloudProvider sets CloudProvider field to given value.


### GetCustomTags

`func (o *DeploymentCellSearchRecord) GetCustomTags() []CustomTag`

GetCustomTags returns the CustomTags field if non-nil, zero value otherwise.

### GetCustomTagsOk

`func (o *DeploymentCellSearchRecord) GetCustomTagsOk() (*[]CustomTag, bool)`

GetCustomTagsOk returns a tuple with the CustomTags field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCustomTags

`func (o *DeploymentCellSearchRecord) SetCustomTags(v []CustomTag)`

SetCustomTags sets CustomTags field to given value.

### HasCustomTags

`func (o *DeploymentCellSearchRecord) HasCustomTags() bool`

HasCustomTags returns a boolean if a field has been set.

### GetDescription

`func (o *DeploymentCellSearchRecord) GetDescription() string`

GetDescription returns the Description field if non-nil, zero value otherwise.

### GetDescriptionOk

`func (o *DeploymentCellSearchRecord) GetDescriptionOk() (*string, bool)`

GetDescriptionOk returns a tuple with the Description field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDescription

`func (o *DeploymentCellSearchRecord) SetDescription(v string)`

SetDescription sets Description field to given value.


### GetId

`func (o *DeploymentCellSearchRecord) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *DeploymentCellSearchRecord) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *DeploymentCellSearchRecord) SetId(v string)`

SetId sets Id field to given value.


### GetRegionCode

`func (o *DeploymentCellSearchRecord) GetRegionCode() string`

GetRegionCode returns the RegionCode field if non-nil, zero value otherwise.

### GetRegionCodeOk

`func (o *DeploymentCellSearchRecord) GetRegionCodeOk() (*string, bool)`

GetRegionCodeOk returns a tuple with the RegionCode field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRegionCode

`func (o *DeploymentCellSearchRecord) SetRegionCode(v string)`

SetRegionCode sets RegionCode field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


