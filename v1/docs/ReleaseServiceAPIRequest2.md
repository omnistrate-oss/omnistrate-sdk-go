# ReleaseServiceAPIRequest2

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**DryRun** | Pointer to **bool** | When true, performs a dry run of the release operation without making any actual changes to the current pending changes and the service API. | [optional] [default to false]
**IsPreferred** | Pointer to **bool** | Indicates whether this version set is preferred. | [optional] [default to false]
**ProductTierId** | Pointer to **string** | The product tier ID | [optional] 
**VersionSetDescription** | Pointer to **string** | The description of the version set to release | [optional] 
**VersionSetName** | Pointer to **string** | The name of the version set to release | [optional] 
**VersionSetType** | Pointer to **string** | The version-set type of the product-tier. | [optional] 

## Methods

### NewReleaseServiceAPIRequest2

`func NewReleaseServiceAPIRequest2() *ReleaseServiceAPIRequest2`

NewReleaseServiceAPIRequest2 instantiates a new ReleaseServiceAPIRequest2 object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewReleaseServiceAPIRequest2WithDefaults

`func NewReleaseServiceAPIRequest2WithDefaults() *ReleaseServiceAPIRequest2`

NewReleaseServiceAPIRequest2WithDefaults instantiates a new ReleaseServiceAPIRequest2 object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetDryRun

`func (o *ReleaseServiceAPIRequest2) GetDryRun() bool`

GetDryRun returns the DryRun field if non-nil, zero value otherwise.

### GetDryRunOk

`func (o *ReleaseServiceAPIRequest2) GetDryRunOk() (*bool, bool)`

GetDryRunOk returns a tuple with the DryRun field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDryRun

`func (o *ReleaseServiceAPIRequest2) SetDryRun(v bool)`

SetDryRun sets DryRun field to given value.

### HasDryRun

`func (o *ReleaseServiceAPIRequest2) HasDryRun() bool`

HasDryRun returns a boolean if a field has been set.

### GetIsPreferred

`func (o *ReleaseServiceAPIRequest2) GetIsPreferred() bool`

GetIsPreferred returns the IsPreferred field if non-nil, zero value otherwise.

### GetIsPreferredOk

`func (o *ReleaseServiceAPIRequest2) GetIsPreferredOk() (*bool, bool)`

GetIsPreferredOk returns a tuple with the IsPreferred field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIsPreferred

`func (o *ReleaseServiceAPIRequest2) SetIsPreferred(v bool)`

SetIsPreferred sets IsPreferred field to given value.

### HasIsPreferred

`func (o *ReleaseServiceAPIRequest2) HasIsPreferred() bool`

HasIsPreferred returns a boolean if a field has been set.

### GetProductTierId

`func (o *ReleaseServiceAPIRequest2) GetProductTierId() string`

GetProductTierId returns the ProductTierId field if non-nil, zero value otherwise.

### GetProductTierIdOk

`func (o *ReleaseServiceAPIRequest2) GetProductTierIdOk() (*string, bool)`

GetProductTierIdOk returns a tuple with the ProductTierId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetProductTierId

`func (o *ReleaseServiceAPIRequest2) SetProductTierId(v string)`

SetProductTierId sets ProductTierId field to given value.

### HasProductTierId

`func (o *ReleaseServiceAPIRequest2) HasProductTierId() bool`

HasProductTierId returns a boolean if a field has been set.

### GetVersionSetDescription

`func (o *ReleaseServiceAPIRequest2) GetVersionSetDescription() string`

GetVersionSetDescription returns the VersionSetDescription field if non-nil, zero value otherwise.

### GetVersionSetDescriptionOk

`func (o *ReleaseServiceAPIRequest2) GetVersionSetDescriptionOk() (*string, bool)`

GetVersionSetDescriptionOk returns a tuple with the VersionSetDescription field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVersionSetDescription

`func (o *ReleaseServiceAPIRequest2) SetVersionSetDescription(v string)`

SetVersionSetDescription sets VersionSetDescription field to given value.

### HasVersionSetDescription

`func (o *ReleaseServiceAPIRequest2) HasVersionSetDescription() bool`

HasVersionSetDescription returns a boolean if a field has been set.

### GetVersionSetName

`func (o *ReleaseServiceAPIRequest2) GetVersionSetName() string`

GetVersionSetName returns the VersionSetName field if non-nil, zero value otherwise.

### GetVersionSetNameOk

`func (o *ReleaseServiceAPIRequest2) GetVersionSetNameOk() (*string, bool)`

GetVersionSetNameOk returns a tuple with the VersionSetName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVersionSetName

`func (o *ReleaseServiceAPIRequest2) SetVersionSetName(v string)`

SetVersionSetName sets VersionSetName field to given value.

### HasVersionSetName

`func (o *ReleaseServiceAPIRequest2) HasVersionSetName() bool`

HasVersionSetName returns a boolean if a field has been set.

### GetVersionSetType

`func (o *ReleaseServiceAPIRequest2) GetVersionSetType() string`

GetVersionSetType returns the VersionSetType field if non-nil, zero value otherwise.

### GetVersionSetTypeOk

`func (o *ReleaseServiceAPIRequest2) GetVersionSetTypeOk() (*string, bool)`

GetVersionSetTypeOk returns a tuple with the VersionSetType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVersionSetType

`func (o *ReleaseServiceAPIRequest2) SetVersionSetType(v string)`

SetVersionSetType sets VersionSetType field to given value.

### HasVersionSetType

`func (o *ReleaseServiceAPIRequest2) HasVersionSetType() bool`

HasVersionSetType returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


