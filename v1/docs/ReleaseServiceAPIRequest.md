# ReleaseServiceAPIRequest

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**CopySpecFromPreviousVersion** | Pointer to **bool** | When true, copies the specification from the previous version during the release operation. | [optional] [default to false]
**DryRun** | Pointer to **bool** | When true, performs a dry run of the release operation without making any actual changes to the current pending changes and the service API. | [optional] [default to false]
**Id** | **string** | ID of a Service API | 
**IsPreferred** | Pointer to **bool** | Indicates whether this version set is preferred. | [optional] [default to false]
**ProductTierId** | Pointer to **string** | ID of a Product Tier | [optional] 
**ServiceId** | **string** | ID of a Service | 
**Token** | **string** | JWT token used to perform authorization | 
**VersionSetName** | Pointer to **string** | The name of the version set to release | [optional] 
**VersionSetType** | Pointer to **string** | The version-set type of the product-tier. | [optional] 

## Methods

### NewReleaseServiceAPIRequest

`func NewReleaseServiceAPIRequest(id string, serviceId string, token string, ) *ReleaseServiceAPIRequest`

NewReleaseServiceAPIRequest instantiates a new ReleaseServiceAPIRequest object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewReleaseServiceAPIRequestWithDefaults

`func NewReleaseServiceAPIRequestWithDefaults() *ReleaseServiceAPIRequest`

NewReleaseServiceAPIRequestWithDefaults instantiates a new ReleaseServiceAPIRequest object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetCopySpecFromPreviousVersion

`func (o *ReleaseServiceAPIRequest) GetCopySpecFromPreviousVersion() bool`

GetCopySpecFromPreviousVersion returns the CopySpecFromPreviousVersion field if non-nil, zero value otherwise.

### GetCopySpecFromPreviousVersionOk

`func (o *ReleaseServiceAPIRequest) GetCopySpecFromPreviousVersionOk() (*bool, bool)`

GetCopySpecFromPreviousVersionOk returns a tuple with the CopySpecFromPreviousVersion field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCopySpecFromPreviousVersion

`func (o *ReleaseServiceAPIRequest) SetCopySpecFromPreviousVersion(v bool)`

SetCopySpecFromPreviousVersion sets CopySpecFromPreviousVersion field to given value.

### HasCopySpecFromPreviousVersion

`func (o *ReleaseServiceAPIRequest) HasCopySpecFromPreviousVersion() bool`

HasCopySpecFromPreviousVersion returns a boolean if a field has been set.

### GetDryRun

`func (o *ReleaseServiceAPIRequest) GetDryRun() bool`

GetDryRun returns the DryRun field if non-nil, zero value otherwise.

### GetDryRunOk

`func (o *ReleaseServiceAPIRequest) GetDryRunOk() (*bool, bool)`

GetDryRunOk returns a tuple with the DryRun field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDryRun

`func (o *ReleaseServiceAPIRequest) SetDryRun(v bool)`

SetDryRun sets DryRun field to given value.

### HasDryRun

`func (o *ReleaseServiceAPIRequest) HasDryRun() bool`

HasDryRun returns a boolean if a field has been set.

### GetId

`func (o *ReleaseServiceAPIRequest) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *ReleaseServiceAPIRequest) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *ReleaseServiceAPIRequest) SetId(v string)`

SetId sets Id field to given value.


### GetIsPreferred

`func (o *ReleaseServiceAPIRequest) GetIsPreferred() bool`

GetIsPreferred returns the IsPreferred field if non-nil, zero value otherwise.

### GetIsPreferredOk

`func (o *ReleaseServiceAPIRequest) GetIsPreferredOk() (*bool, bool)`

GetIsPreferredOk returns a tuple with the IsPreferred field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIsPreferred

`func (o *ReleaseServiceAPIRequest) SetIsPreferred(v bool)`

SetIsPreferred sets IsPreferred field to given value.

### HasIsPreferred

`func (o *ReleaseServiceAPIRequest) HasIsPreferred() bool`

HasIsPreferred returns a boolean if a field has been set.

### GetProductTierId

`func (o *ReleaseServiceAPIRequest) GetProductTierId() string`

GetProductTierId returns the ProductTierId field if non-nil, zero value otherwise.

### GetProductTierIdOk

`func (o *ReleaseServiceAPIRequest) GetProductTierIdOk() (*string, bool)`

GetProductTierIdOk returns a tuple with the ProductTierId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetProductTierId

`func (o *ReleaseServiceAPIRequest) SetProductTierId(v string)`

SetProductTierId sets ProductTierId field to given value.

### HasProductTierId

`func (o *ReleaseServiceAPIRequest) HasProductTierId() bool`

HasProductTierId returns a boolean if a field has been set.

### GetServiceId

`func (o *ReleaseServiceAPIRequest) GetServiceId() string`

GetServiceId returns the ServiceId field if non-nil, zero value otherwise.

### GetServiceIdOk

`func (o *ReleaseServiceAPIRequest) GetServiceIdOk() (*string, bool)`

GetServiceIdOk returns a tuple with the ServiceId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetServiceId

`func (o *ReleaseServiceAPIRequest) SetServiceId(v string)`

SetServiceId sets ServiceId field to given value.


### GetToken

`func (o *ReleaseServiceAPIRequest) GetToken() string`

GetToken returns the Token field if non-nil, zero value otherwise.

### GetTokenOk

`func (o *ReleaseServiceAPIRequest) GetTokenOk() (*string, bool)`

GetTokenOk returns a tuple with the Token field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetToken

`func (o *ReleaseServiceAPIRequest) SetToken(v string)`

SetToken sets Token field to given value.


### GetVersionSetName

`func (o *ReleaseServiceAPIRequest) GetVersionSetName() string`

GetVersionSetName returns the VersionSetName field if non-nil, zero value otherwise.

### GetVersionSetNameOk

`func (o *ReleaseServiceAPIRequest) GetVersionSetNameOk() (*string, bool)`

GetVersionSetNameOk returns a tuple with the VersionSetName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVersionSetName

`func (o *ReleaseServiceAPIRequest) SetVersionSetName(v string)`

SetVersionSetName sets VersionSetName field to given value.

### HasVersionSetName

`func (o *ReleaseServiceAPIRequest) HasVersionSetName() bool`

HasVersionSetName returns a boolean if a field has been set.

### GetVersionSetType

`func (o *ReleaseServiceAPIRequest) GetVersionSetType() string`

GetVersionSetType returns the VersionSetType field if non-nil, zero value otherwise.

### GetVersionSetTypeOk

`func (o *ReleaseServiceAPIRequest) GetVersionSetTypeOk() (*string, bool)`

GetVersionSetTypeOk returns a tuple with the VersionSetType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVersionSetType

`func (o *ReleaseServiceAPIRequest) SetVersionSetType(v string)`

SetVersionSetType sets VersionSetType field to given value.

### HasVersionSetType

`func (o *ReleaseServiceAPIRequest) HasVersionSetType() bool`

HasVersionSetType returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


