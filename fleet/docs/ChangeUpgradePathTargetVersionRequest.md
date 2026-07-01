# ChangeUpgradePathTargetVersionRequest

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**ProductTierId** | **string** | ID of a Product Tier | 
**ServiceId** | **string** | ID of a Service | 
**TargetVersion** | **string** | The new target product tier version for the scheduled upgrade path. | 
**Token** | **string** | JWT token used to perform authorization | 
**UpgradePathId** | **string** | ID of an Upgrade Path | 

## Methods

### NewChangeUpgradePathTargetVersionRequest

`func NewChangeUpgradePathTargetVersionRequest(productTierId string, serviceId string, targetVersion string, token string, upgradePathId string, ) *ChangeUpgradePathTargetVersionRequest`

NewChangeUpgradePathTargetVersionRequest instantiates a new ChangeUpgradePathTargetVersionRequest object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewChangeUpgradePathTargetVersionRequestWithDefaults

`func NewChangeUpgradePathTargetVersionRequestWithDefaults() *ChangeUpgradePathTargetVersionRequest`

NewChangeUpgradePathTargetVersionRequestWithDefaults instantiates a new ChangeUpgradePathTargetVersionRequest object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetProductTierId

`func (o *ChangeUpgradePathTargetVersionRequest) GetProductTierId() string`

GetProductTierId returns the ProductTierId field if non-nil, zero value otherwise.

### GetProductTierIdOk

`func (o *ChangeUpgradePathTargetVersionRequest) GetProductTierIdOk() (*string, bool)`

GetProductTierIdOk returns a tuple with the ProductTierId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetProductTierId

`func (o *ChangeUpgradePathTargetVersionRequest) SetProductTierId(v string)`

SetProductTierId sets ProductTierId field to given value.


### GetServiceId

`func (o *ChangeUpgradePathTargetVersionRequest) GetServiceId() string`

GetServiceId returns the ServiceId field if non-nil, zero value otherwise.

### GetServiceIdOk

`func (o *ChangeUpgradePathTargetVersionRequest) GetServiceIdOk() (*string, bool)`

GetServiceIdOk returns a tuple with the ServiceId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetServiceId

`func (o *ChangeUpgradePathTargetVersionRequest) SetServiceId(v string)`

SetServiceId sets ServiceId field to given value.


### GetTargetVersion

`func (o *ChangeUpgradePathTargetVersionRequest) GetTargetVersion() string`

GetTargetVersion returns the TargetVersion field if non-nil, zero value otherwise.

### GetTargetVersionOk

`func (o *ChangeUpgradePathTargetVersionRequest) GetTargetVersionOk() (*string, bool)`

GetTargetVersionOk returns a tuple with the TargetVersion field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTargetVersion

`func (o *ChangeUpgradePathTargetVersionRequest) SetTargetVersion(v string)`

SetTargetVersion sets TargetVersion field to given value.


### GetToken

`func (o *ChangeUpgradePathTargetVersionRequest) GetToken() string`

GetToken returns the Token field if non-nil, zero value otherwise.

### GetTokenOk

`func (o *ChangeUpgradePathTargetVersionRequest) GetTokenOk() (*string, bool)`

GetTokenOk returns a tuple with the Token field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetToken

`func (o *ChangeUpgradePathTargetVersionRequest) SetToken(v string)`

SetToken sets Token field to given value.


### GetUpgradePathId

`func (o *ChangeUpgradePathTargetVersionRequest) GetUpgradePathId() string`

GetUpgradePathId returns the UpgradePathId field if non-nil, zero value otherwise.

### GetUpgradePathIdOk

`func (o *ChangeUpgradePathTargetVersionRequest) GetUpgradePathIdOk() (*string, bool)`

GetUpgradePathIdOk returns a tuple with the UpgradePathId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUpgradePathId

`func (o *ChangeUpgradePathTargetVersionRequest) SetUpgradePathId(v string)`

SetUpgradePathId sets UpgradePathId field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


