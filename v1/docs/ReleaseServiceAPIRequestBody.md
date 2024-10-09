# ReleaseServiceAPIRequestBody

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**IsPreferred** | Pointer to **bool** | Indicates whether this version set is preferred. | [optional] [default to false]
**ProductTierId** | Pointer to **string** | The product tier ID | [optional] 
**VersionSetName** | Pointer to **string** | The name of the version set to release | [optional] 
**VersionSetType** | Pointer to **string** | The version-set type of the product-tier. | [optional] 

## Methods

### NewReleaseServiceAPIRequestBody

`func NewReleaseServiceAPIRequestBody() *ReleaseServiceAPIRequestBody`

NewReleaseServiceAPIRequestBody instantiates a new ReleaseServiceAPIRequestBody object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewReleaseServiceAPIRequestBodyWithDefaults

`func NewReleaseServiceAPIRequestBodyWithDefaults() *ReleaseServiceAPIRequestBody`

NewReleaseServiceAPIRequestBodyWithDefaults instantiates a new ReleaseServiceAPIRequestBody object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetIsPreferred

`func (o *ReleaseServiceAPIRequestBody) GetIsPreferred() bool`

GetIsPreferred returns the IsPreferred field if non-nil, zero value otherwise.

### GetIsPreferredOk

`func (o *ReleaseServiceAPIRequestBody) GetIsPreferredOk() (*bool, bool)`

GetIsPreferredOk returns a tuple with the IsPreferred field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIsPreferred

`func (o *ReleaseServiceAPIRequestBody) SetIsPreferred(v bool)`

SetIsPreferred sets IsPreferred field to given value.

### HasIsPreferred

`func (o *ReleaseServiceAPIRequestBody) HasIsPreferred() bool`

HasIsPreferred returns a boolean if a field has been set.

### GetProductTierId

`func (o *ReleaseServiceAPIRequestBody) GetProductTierId() string`

GetProductTierId returns the ProductTierId field if non-nil, zero value otherwise.

### GetProductTierIdOk

`func (o *ReleaseServiceAPIRequestBody) GetProductTierIdOk() (*string, bool)`

GetProductTierIdOk returns a tuple with the ProductTierId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetProductTierId

`func (o *ReleaseServiceAPIRequestBody) SetProductTierId(v string)`

SetProductTierId sets ProductTierId field to given value.

### HasProductTierId

`func (o *ReleaseServiceAPIRequestBody) HasProductTierId() bool`

HasProductTierId returns a boolean if a field has been set.

### GetVersionSetName

`func (o *ReleaseServiceAPIRequestBody) GetVersionSetName() string`

GetVersionSetName returns the VersionSetName field if non-nil, zero value otherwise.

### GetVersionSetNameOk

`func (o *ReleaseServiceAPIRequestBody) GetVersionSetNameOk() (*string, bool)`

GetVersionSetNameOk returns a tuple with the VersionSetName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVersionSetName

`func (o *ReleaseServiceAPIRequestBody) SetVersionSetName(v string)`

SetVersionSetName sets VersionSetName field to given value.

### HasVersionSetName

`func (o *ReleaseServiceAPIRequestBody) HasVersionSetName() bool`

HasVersionSetName returns a boolean if a field has been set.

### GetVersionSetType

`func (o *ReleaseServiceAPIRequestBody) GetVersionSetType() string`

GetVersionSetType returns the VersionSetType field if non-nil, zero value otherwise.

### GetVersionSetTypeOk

`func (o *ReleaseServiceAPIRequestBody) GetVersionSetTypeOk() (*string, bool)`

GetVersionSetTypeOk returns a tuple with the VersionSetType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVersionSetType

`func (o *ReleaseServiceAPIRequestBody) SetVersionSetType(v string)`

SetVersionSetType sets VersionSetType field to given value.

### HasVersionSetType

`func (o *ReleaseServiceAPIRequestBody) HasVersionSetType() bool`

HasVersionSetType returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)

