# FleetSyncAccountConfigCloudNativeNetworkTarget

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**CloudNativeNetworkId** | Pointer to **string** | Optional cloud provider network ID (e.g. AWS VPC ID). When omitted, every VPC in the region is enumerated. | [optional] 
**Region** | **string** | The cloud region where the network resides | 

## Methods

### NewFleetSyncAccountConfigCloudNativeNetworkTarget

`func NewFleetSyncAccountConfigCloudNativeNetworkTarget(region string, ) *FleetSyncAccountConfigCloudNativeNetworkTarget`

NewFleetSyncAccountConfigCloudNativeNetworkTarget instantiates a new FleetSyncAccountConfigCloudNativeNetworkTarget object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewFleetSyncAccountConfigCloudNativeNetworkTargetWithDefaults

`func NewFleetSyncAccountConfigCloudNativeNetworkTargetWithDefaults() *FleetSyncAccountConfigCloudNativeNetworkTarget`

NewFleetSyncAccountConfigCloudNativeNetworkTargetWithDefaults instantiates a new FleetSyncAccountConfigCloudNativeNetworkTarget object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetCloudNativeNetworkId

`func (o *FleetSyncAccountConfigCloudNativeNetworkTarget) GetCloudNativeNetworkId() string`

GetCloudNativeNetworkId returns the CloudNativeNetworkId field if non-nil, zero value otherwise.

### GetCloudNativeNetworkIdOk

`func (o *FleetSyncAccountConfigCloudNativeNetworkTarget) GetCloudNativeNetworkIdOk() (*string, bool)`

GetCloudNativeNetworkIdOk returns a tuple with the CloudNativeNetworkId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCloudNativeNetworkId

`func (o *FleetSyncAccountConfigCloudNativeNetworkTarget) SetCloudNativeNetworkId(v string)`

SetCloudNativeNetworkId sets CloudNativeNetworkId field to given value.

### HasCloudNativeNetworkId

`func (o *FleetSyncAccountConfigCloudNativeNetworkTarget) HasCloudNativeNetworkId() bool`

HasCloudNativeNetworkId returns a boolean if a field has been set.

### GetRegion

`func (o *FleetSyncAccountConfigCloudNativeNetworkTarget) GetRegion() string`

GetRegion returns the Region field if non-nil, zero value otherwise.

### GetRegionOk

`func (o *FleetSyncAccountConfigCloudNativeNetworkTarget) GetRegionOk() (*string, bool)`

GetRegionOk returns a tuple with the Region field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRegion

`func (o *FleetSyncAccountConfigCloudNativeNetworkTarget) SetRegion(v string)`

SetRegion sets Region field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


