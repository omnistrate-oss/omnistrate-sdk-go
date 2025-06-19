# ListUpgradePathsResult

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**NextPageToken** | Pointer to **string** | Token to use for the next request | [optional] 
**UpgradePaths** | [**[]UpgradePath**](UpgradePath.md) | The list of upgrade paths. | 

## Methods

### NewListUpgradePathsResult

`func NewListUpgradePathsResult(upgradePaths []UpgradePath, ) *ListUpgradePathsResult`

NewListUpgradePathsResult instantiates a new ListUpgradePathsResult object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewListUpgradePathsResultWithDefaults

`func NewListUpgradePathsResultWithDefaults() *ListUpgradePathsResult`

NewListUpgradePathsResultWithDefaults instantiates a new ListUpgradePathsResult object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetNextPageToken

`func (o *ListUpgradePathsResult) GetNextPageToken() string`

GetNextPageToken returns the NextPageToken field if non-nil, zero value otherwise.

### GetNextPageTokenOk

`func (o *ListUpgradePathsResult) GetNextPageTokenOk() (*string, bool)`

GetNextPageTokenOk returns a tuple with the NextPageToken field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNextPageToken

`func (o *ListUpgradePathsResult) SetNextPageToken(v string)`

SetNextPageToken sets NextPageToken field to given value.

### HasNextPageToken

`func (o *ListUpgradePathsResult) HasNextPageToken() bool`

HasNextPageToken returns a boolean if a field has been set.

### GetUpgradePaths

`func (o *ListUpgradePathsResult) GetUpgradePaths() []UpgradePath`

GetUpgradePaths returns the UpgradePaths field if non-nil, zero value otherwise.

### GetUpgradePathsOk

`func (o *ListUpgradePathsResult) GetUpgradePathsOk() (*[]UpgradePath, bool)`

GetUpgradePathsOk returns a tuple with the UpgradePaths field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUpgradePaths

`func (o *ListUpgradePathsResult) SetUpgradePaths(v []UpgradePath)`

SetUpgradePaths sets UpgradePaths field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


