# CreateUpgradePathRequestBody

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**SourceVersion** | **string** | The source version of the upgrade path. | 
**TargetVersion** | **string** | The target version of the upgrade path. | 
**UpgradeFilters** | **map[string][]string** | The filter to use to choose the instances to upgrade. | 

## Methods

### NewCreateUpgradePathRequestBody

`func NewCreateUpgradePathRequestBody(sourceVersion string, targetVersion string, upgradeFilters map[string][]string, ) *CreateUpgradePathRequestBody`

NewCreateUpgradePathRequestBody instantiates a new CreateUpgradePathRequestBody object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewCreateUpgradePathRequestBodyWithDefaults

`func NewCreateUpgradePathRequestBodyWithDefaults() *CreateUpgradePathRequestBody`

NewCreateUpgradePathRequestBodyWithDefaults instantiates a new CreateUpgradePathRequestBody object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetSourceVersion

`func (o *CreateUpgradePathRequestBody) GetSourceVersion() string`

GetSourceVersion returns the SourceVersion field if non-nil, zero value otherwise.

### GetSourceVersionOk

`func (o *CreateUpgradePathRequestBody) GetSourceVersionOk() (*string, bool)`

GetSourceVersionOk returns a tuple with the SourceVersion field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSourceVersion

`func (o *CreateUpgradePathRequestBody) SetSourceVersion(v string)`

SetSourceVersion sets SourceVersion field to given value.


### GetTargetVersion

`func (o *CreateUpgradePathRequestBody) GetTargetVersion() string`

GetTargetVersion returns the TargetVersion field if non-nil, zero value otherwise.

### GetTargetVersionOk

`func (o *CreateUpgradePathRequestBody) GetTargetVersionOk() (*string, bool)`

GetTargetVersionOk returns a tuple with the TargetVersion field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTargetVersion

`func (o *CreateUpgradePathRequestBody) SetTargetVersion(v string)`

SetTargetVersion sets TargetVersion field to given value.


### GetUpgradeFilters

`func (o *CreateUpgradePathRequestBody) GetUpgradeFilters() map[string][]string`

GetUpgradeFilters returns the UpgradeFilters field if non-nil, zero value otherwise.

### GetUpgradeFiltersOk

`func (o *CreateUpgradePathRequestBody) GetUpgradeFiltersOk() (*map[string][]string, bool)`

GetUpgradeFiltersOk returns a tuple with the UpgradeFilters field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUpgradeFilters

`func (o *CreateUpgradePathRequestBody) SetUpgradeFilters(v map[string][]string)`

SetUpgradeFilters sets UpgradeFilters field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


