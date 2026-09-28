# SubscriptionAllowedDeploymentLocation

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**CloudProvider** | **string** | Name of the Infra Provider | 
**Regions** | Pointer to **[]string** | The regions this subscription is allowed to deploy in for the cloud provider. Omit or set to an empty array to allow all product tier regions for this cloud provider. | [optional] 

## Methods

### NewSubscriptionAllowedDeploymentLocation

`func NewSubscriptionAllowedDeploymentLocation(cloudProvider string, ) *SubscriptionAllowedDeploymentLocation`

NewSubscriptionAllowedDeploymentLocation instantiates a new SubscriptionAllowedDeploymentLocation object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewSubscriptionAllowedDeploymentLocationWithDefaults

`func NewSubscriptionAllowedDeploymentLocationWithDefaults() *SubscriptionAllowedDeploymentLocation`

NewSubscriptionAllowedDeploymentLocationWithDefaults instantiates a new SubscriptionAllowedDeploymentLocation object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetCloudProvider

`func (o *SubscriptionAllowedDeploymentLocation) GetCloudProvider() string`

GetCloudProvider returns the CloudProvider field if non-nil, zero value otherwise.

### GetCloudProviderOk

`func (o *SubscriptionAllowedDeploymentLocation) GetCloudProviderOk() (*string, bool)`

GetCloudProviderOk returns a tuple with the CloudProvider field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCloudProvider

`func (o *SubscriptionAllowedDeploymentLocation) SetCloudProvider(v string)`

SetCloudProvider sets CloudProvider field to given value.


### GetRegions

`func (o *SubscriptionAllowedDeploymentLocation) GetRegions() []string`

GetRegions returns the Regions field if non-nil, zero value otherwise.

### GetRegionsOk

`func (o *SubscriptionAllowedDeploymentLocation) GetRegionsOk() (*[]string, bool)`

GetRegionsOk returns a tuple with the Regions field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRegions

`func (o *SubscriptionAllowedDeploymentLocation) SetRegions(v []string)`

SetRegions sets Regions field to given value.

### HasRegions

`func (o *SubscriptionAllowedDeploymentLocation) HasRegions() bool`

HasRegions returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


