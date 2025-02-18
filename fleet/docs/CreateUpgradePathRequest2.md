# CreateUpgradePathRequest2

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**ScheduledDate** | Pointer to **string** | The future date when the upgrade is planned to be executed. Empty for immediate upgrade. | [optional] 
**SourceVersion** | **string** | The source version of the upgrade path. | 
**TargetVersion** | **string** | The target version of the upgrade path. | 
**UpgradeFilters** | **map[string][]string** | The filter to use to choose the instances to upgrade. | 

## Methods

### NewCreateUpgradePathRequest2

`func NewCreateUpgradePathRequest2(sourceVersion string, targetVersion string, upgradeFilters map[string][]string, ) *CreateUpgradePathRequest2`

NewCreateUpgradePathRequest2 instantiates a new CreateUpgradePathRequest2 object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewCreateUpgradePathRequest2WithDefaults

`func NewCreateUpgradePathRequest2WithDefaults() *CreateUpgradePathRequest2`

NewCreateUpgradePathRequest2WithDefaults instantiates a new CreateUpgradePathRequest2 object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetScheduledDate

`func (o *CreateUpgradePathRequest2) GetScheduledDate() string`

GetScheduledDate returns the ScheduledDate field if non-nil, zero value otherwise.

### GetScheduledDateOk

`func (o *CreateUpgradePathRequest2) GetScheduledDateOk() (*string, bool)`

GetScheduledDateOk returns a tuple with the ScheduledDate field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetScheduledDate

`func (o *CreateUpgradePathRequest2) SetScheduledDate(v string)`

SetScheduledDate sets ScheduledDate field to given value.

### HasScheduledDate

`func (o *CreateUpgradePathRequest2) HasScheduledDate() bool`

HasScheduledDate returns a boolean if a field has been set.

### GetSourceVersion

`func (o *CreateUpgradePathRequest2) GetSourceVersion() string`

GetSourceVersion returns the SourceVersion field if non-nil, zero value otherwise.

### GetSourceVersionOk

`func (o *CreateUpgradePathRequest2) GetSourceVersionOk() (*string, bool)`

GetSourceVersionOk returns a tuple with the SourceVersion field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSourceVersion

`func (o *CreateUpgradePathRequest2) SetSourceVersion(v string)`

SetSourceVersion sets SourceVersion field to given value.


### GetTargetVersion

`func (o *CreateUpgradePathRequest2) GetTargetVersion() string`

GetTargetVersion returns the TargetVersion field if non-nil, zero value otherwise.

### GetTargetVersionOk

`func (o *CreateUpgradePathRequest2) GetTargetVersionOk() (*string, bool)`

GetTargetVersionOk returns a tuple with the TargetVersion field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTargetVersion

`func (o *CreateUpgradePathRequest2) SetTargetVersion(v string)`

SetTargetVersion sets TargetVersion field to given value.


### GetUpgradeFilters

`func (o *CreateUpgradePathRequest2) GetUpgradeFilters() map[string][]string`

GetUpgradeFilters returns the UpgradeFilters field if non-nil, zero value otherwise.

### GetUpgradeFiltersOk

`func (o *CreateUpgradePathRequest2) GetUpgradeFiltersOk() (*map[string][]string, bool)`

GetUpgradeFiltersOk returns a tuple with the UpgradeFilters field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUpgradeFilters

`func (o *CreateUpgradePathRequest2) SetUpgradeFilters(v map[string][]string)`

SetUpgradeFilters sets UpgradeFilters field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


