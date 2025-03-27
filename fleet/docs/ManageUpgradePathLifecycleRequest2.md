# ManageUpgradePathLifecycleRequest2

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Action** | **string** | The action to perform on the upgrade path. | 
**ActionPayload** | Pointer to **map[string]interface{}** | The action payload to perform on the upgrade path. | [optional] 

## Methods

### NewManageUpgradePathLifecycleRequest2

`func NewManageUpgradePathLifecycleRequest2(action string, ) *ManageUpgradePathLifecycleRequest2`

NewManageUpgradePathLifecycleRequest2 instantiates a new ManageUpgradePathLifecycleRequest2 object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewManageUpgradePathLifecycleRequest2WithDefaults

`func NewManageUpgradePathLifecycleRequest2WithDefaults() *ManageUpgradePathLifecycleRequest2`

NewManageUpgradePathLifecycleRequest2WithDefaults instantiates a new ManageUpgradePathLifecycleRequest2 object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetAction

`func (o *ManageUpgradePathLifecycleRequest2) GetAction() string`

GetAction returns the Action field if non-nil, zero value otherwise.

### GetActionOk

`func (o *ManageUpgradePathLifecycleRequest2) GetActionOk() (*string, bool)`

GetActionOk returns a tuple with the Action field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAction

`func (o *ManageUpgradePathLifecycleRequest2) SetAction(v string)`

SetAction sets Action field to given value.


### GetActionPayload

`func (o *ManageUpgradePathLifecycleRequest2) GetActionPayload() map[string]interface{}`

GetActionPayload returns the ActionPayload field if non-nil, zero value otherwise.

### GetActionPayloadOk

`func (o *ManageUpgradePathLifecycleRequest2) GetActionPayloadOk() (*map[string]interface{}, bool)`

GetActionPayloadOk returns a tuple with the ActionPayload field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetActionPayload

`func (o *ManageUpgradePathLifecycleRequest2) SetActionPayload(v map[string]interface{})`

SetActionPayload sets ActionPayload field to given value.

### HasActionPayload

`func (o *ManageUpgradePathLifecycleRequest2) HasActionPayload() bool`

HasActionPayload returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


