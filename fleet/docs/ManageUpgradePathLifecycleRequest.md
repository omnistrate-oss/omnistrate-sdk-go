# ManageUpgradePathLifecycleRequest

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Action** | **string** | The action to perform on an ongoing resource workflow | 
**ProductTierId** | **string** | ID of a Product Tier | 
**ServiceId** | **string** | ID of a Service | 
**Token** | **string** | JWT token used to perform authorization | 
**UpgradePathId** | **string** | ID of an Upgrade Path | 

## Methods

### NewManageUpgradePathLifecycleRequest

`func NewManageUpgradePathLifecycleRequest(action string, productTierId string, serviceId string, token string, upgradePathId string, ) *ManageUpgradePathLifecycleRequest`

NewManageUpgradePathLifecycleRequest instantiates a new ManageUpgradePathLifecycleRequest object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewManageUpgradePathLifecycleRequestWithDefaults

`func NewManageUpgradePathLifecycleRequestWithDefaults() *ManageUpgradePathLifecycleRequest`

NewManageUpgradePathLifecycleRequestWithDefaults instantiates a new ManageUpgradePathLifecycleRequest object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetAction

`func (o *ManageUpgradePathLifecycleRequest) GetAction() string`

GetAction returns the Action field if non-nil, zero value otherwise.

### GetActionOk

`func (o *ManageUpgradePathLifecycleRequest) GetActionOk() (*string, bool)`

GetActionOk returns a tuple with the Action field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAction

`func (o *ManageUpgradePathLifecycleRequest) SetAction(v string)`

SetAction sets Action field to given value.


### GetProductTierId

`func (o *ManageUpgradePathLifecycleRequest) GetProductTierId() string`

GetProductTierId returns the ProductTierId field if non-nil, zero value otherwise.

### GetProductTierIdOk

`func (o *ManageUpgradePathLifecycleRequest) GetProductTierIdOk() (*string, bool)`

GetProductTierIdOk returns a tuple with the ProductTierId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetProductTierId

`func (o *ManageUpgradePathLifecycleRequest) SetProductTierId(v string)`

SetProductTierId sets ProductTierId field to given value.


### GetServiceId

`func (o *ManageUpgradePathLifecycleRequest) GetServiceId() string`

GetServiceId returns the ServiceId field if non-nil, zero value otherwise.

### GetServiceIdOk

`func (o *ManageUpgradePathLifecycleRequest) GetServiceIdOk() (*string, bool)`

GetServiceIdOk returns a tuple with the ServiceId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetServiceId

`func (o *ManageUpgradePathLifecycleRequest) SetServiceId(v string)`

SetServiceId sets ServiceId field to given value.


### GetToken

`func (o *ManageUpgradePathLifecycleRequest) GetToken() string`

GetToken returns the Token field if non-nil, zero value otherwise.

### GetTokenOk

`func (o *ManageUpgradePathLifecycleRequest) GetTokenOk() (*string, bool)`

GetTokenOk returns a tuple with the Token field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetToken

`func (o *ManageUpgradePathLifecycleRequest) SetToken(v string)`

SetToken sets Token field to given value.


### GetUpgradePathId

`func (o *ManageUpgradePathLifecycleRequest) GetUpgradePathId() string`

GetUpgradePathId returns the UpgradePathId field if non-nil, zero value otherwise.

### GetUpgradePathIdOk

`func (o *ManageUpgradePathLifecycleRequest) GetUpgradePathIdOk() (*string, bool)`

GetUpgradePathIdOk returns a tuple with the UpgradePathId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUpgradePathId

`func (o *ManageUpgradePathLifecycleRequest) SetUpgradePathId(v string)`

SetUpgradePathId sets UpgradePathId field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


