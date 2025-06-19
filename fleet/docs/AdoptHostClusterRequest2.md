# AdoptHostClusterRequest2

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**CloudProviderId** | **string** |  | 
**Description** | **string** | Description of the host cluster | 
**Id** | **string** | ID of the host cluster to adopt | 
**RegionId** | **string** | ID of the region where the host cluster is located | 
**SubscriptionId** | Pointer to **string** | The subscription ID | [optional] 

## Methods

### NewAdoptHostClusterRequest2

`func NewAdoptHostClusterRequest2(cloudProviderId string, description string, id string, regionId string, ) *AdoptHostClusterRequest2`

NewAdoptHostClusterRequest2 instantiates a new AdoptHostClusterRequest2 object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewAdoptHostClusterRequest2WithDefaults

`func NewAdoptHostClusterRequest2WithDefaults() *AdoptHostClusterRequest2`

NewAdoptHostClusterRequest2WithDefaults instantiates a new AdoptHostClusterRequest2 object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetCloudProviderId

`func (o *AdoptHostClusterRequest2) GetCloudProviderId() string`

GetCloudProviderId returns the CloudProviderId field if non-nil, zero value otherwise.

### GetCloudProviderIdOk

`func (o *AdoptHostClusterRequest2) GetCloudProviderIdOk() (*string, bool)`

GetCloudProviderIdOk returns a tuple with the CloudProviderId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCloudProviderId

`func (o *AdoptHostClusterRequest2) SetCloudProviderId(v string)`

SetCloudProviderId sets CloudProviderId field to given value.


### GetDescription

`func (o *AdoptHostClusterRequest2) GetDescription() string`

GetDescription returns the Description field if non-nil, zero value otherwise.

### GetDescriptionOk

`func (o *AdoptHostClusterRequest2) GetDescriptionOk() (*string, bool)`

GetDescriptionOk returns a tuple with the Description field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDescription

`func (o *AdoptHostClusterRequest2) SetDescription(v string)`

SetDescription sets Description field to given value.


### GetId

`func (o *AdoptHostClusterRequest2) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *AdoptHostClusterRequest2) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *AdoptHostClusterRequest2) SetId(v string)`

SetId sets Id field to given value.


### GetRegionId

`func (o *AdoptHostClusterRequest2) GetRegionId() string`

GetRegionId returns the RegionId field if non-nil, zero value otherwise.

### GetRegionIdOk

`func (o *AdoptHostClusterRequest2) GetRegionIdOk() (*string, bool)`

GetRegionIdOk returns a tuple with the RegionId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRegionId

`func (o *AdoptHostClusterRequest2) SetRegionId(v string)`

SetRegionId sets RegionId field to given value.


### GetSubscriptionId

`func (o *AdoptHostClusterRequest2) GetSubscriptionId() string`

GetSubscriptionId returns the SubscriptionId field if non-nil, zero value otherwise.

### GetSubscriptionIdOk

`func (o *AdoptHostClusterRequest2) GetSubscriptionIdOk() (*string, bool)`

GetSubscriptionIdOk returns a tuple with the SubscriptionId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSubscriptionId

`func (o *AdoptHostClusterRequest2) SetSubscriptionId(v string)`

SetSubscriptionId sets SubscriptionId field to given value.

### HasSubscriptionId

`func (o *AdoptHostClusterRequest2) HasSubscriptionId() bool`

HasSubscriptionId returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


