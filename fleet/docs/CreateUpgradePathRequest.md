# CreateUpgradePathRequest

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**ProductTierId** | **string** | ID of a Product Tier | 
**ScheduledDate** | Pointer to **string** | The future date when the upgrade is planned to be executed. Empty for immediate upgrade. | [optional] 
**ServiceId** | **string** | ID of a Service | 
**SourceVersion** | **string** | The source version of the upgrade path. | 
**TargetVersion** | **string** | The target version of the upgrade path. | 
**Token** | **string** | JWT token used to perform authorization | 
**UpgradeFilters** | **map[string]interface{}** | The filter to use to choose the instances to upgrade. | 

## Methods

### NewCreateUpgradePathRequest

`func NewCreateUpgradePathRequest(productTierId string, serviceId string, sourceVersion string, targetVersion string, token string, upgradeFilters map[string]interface{}, ) *CreateUpgradePathRequest`

NewCreateUpgradePathRequest instantiates a new CreateUpgradePathRequest object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewCreateUpgradePathRequestWithDefaults

`func NewCreateUpgradePathRequestWithDefaults() *CreateUpgradePathRequest`

NewCreateUpgradePathRequestWithDefaults instantiates a new CreateUpgradePathRequest object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetProductTierId

`func (o *CreateUpgradePathRequest) GetProductTierId() string`

GetProductTierId returns the ProductTierId field if non-nil, zero value otherwise.

### GetProductTierIdOk

`func (o *CreateUpgradePathRequest) GetProductTierIdOk() (*string, bool)`

GetProductTierIdOk returns a tuple with the ProductTierId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetProductTierId

`func (o *CreateUpgradePathRequest) SetProductTierId(v string)`

SetProductTierId sets ProductTierId field to given value.


### GetScheduledDate

`func (o *CreateUpgradePathRequest) GetScheduledDate() string`

GetScheduledDate returns the ScheduledDate field if non-nil, zero value otherwise.

### GetScheduledDateOk

`func (o *CreateUpgradePathRequest) GetScheduledDateOk() (*string, bool)`

GetScheduledDateOk returns a tuple with the ScheduledDate field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetScheduledDate

`func (o *CreateUpgradePathRequest) SetScheduledDate(v string)`

SetScheduledDate sets ScheduledDate field to given value.

### HasScheduledDate

`func (o *CreateUpgradePathRequest) HasScheduledDate() bool`

HasScheduledDate returns a boolean if a field has been set.

### GetServiceId

`func (o *CreateUpgradePathRequest) GetServiceId() string`

GetServiceId returns the ServiceId field if non-nil, zero value otherwise.

### GetServiceIdOk

`func (o *CreateUpgradePathRequest) GetServiceIdOk() (*string, bool)`

GetServiceIdOk returns a tuple with the ServiceId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetServiceId

`func (o *CreateUpgradePathRequest) SetServiceId(v string)`

SetServiceId sets ServiceId field to given value.


### GetSourceVersion

`func (o *CreateUpgradePathRequest) GetSourceVersion() string`

GetSourceVersion returns the SourceVersion field if non-nil, zero value otherwise.

### GetSourceVersionOk

`func (o *CreateUpgradePathRequest) GetSourceVersionOk() (*string, bool)`

GetSourceVersionOk returns a tuple with the SourceVersion field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSourceVersion

`func (o *CreateUpgradePathRequest) SetSourceVersion(v string)`

SetSourceVersion sets SourceVersion field to given value.


### GetTargetVersion

`func (o *CreateUpgradePathRequest) GetTargetVersion() string`

GetTargetVersion returns the TargetVersion field if non-nil, zero value otherwise.

### GetTargetVersionOk

`func (o *CreateUpgradePathRequest) GetTargetVersionOk() (*string, bool)`

GetTargetVersionOk returns a tuple with the TargetVersion field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTargetVersion

`func (o *CreateUpgradePathRequest) SetTargetVersion(v string)`

SetTargetVersion sets TargetVersion field to given value.


### GetToken

`func (o *CreateUpgradePathRequest) GetToken() string`

GetToken returns the Token field if non-nil, zero value otherwise.

### GetTokenOk

`func (o *CreateUpgradePathRequest) GetTokenOk() (*string, bool)`

GetTokenOk returns a tuple with the Token field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetToken

`func (o *CreateUpgradePathRequest) SetToken(v string)`

SetToken sets Token field to given value.


### GetUpgradeFilters

`func (o *CreateUpgradePathRequest) GetUpgradeFilters() map[string]interface{}`

GetUpgradeFilters returns the UpgradeFilters field if non-nil, zero value otherwise.

### GetUpgradeFiltersOk

`func (o *CreateUpgradePathRequest) GetUpgradeFiltersOk() (*map[string]interface{}, bool)`

GetUpgradeFiltersOk returns a tuple with the UpgradeFilters field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUpgradeFilters

`func (o *CreateUpgradePathRequest) SetUpgradeFilters(v map[string]interface{})`

SetUpgradeFilters sets UpgradeFilters field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


