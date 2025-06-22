# AdoptHostClusterRequest2

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**CloudProvider** | **string** |  | 
**CustomerEmail** | Pointer to **string** | Email of the customer who owns the host cluster in case this is a BYOA host cluster | [optional] 
**Description** | **string** | Description of the host cluster | 
**Id** | **string** | ID of the host cluster to adopt | 
**Region** | **string** | The actual region name of the host cluster | 

## Methods

### NewAdoptHostClusterRequest2

`func NewAdoptHostClusterRequest2(cloudProvider string, description string, id string, region string, ) *AdoptHostClusterRequest2`

NewAdoptHostClusterRequest2 instantiates a new AdoptHostClusterRequest2 object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewAdoptHostClusterRequest2WithDefaults

`func NewAdoptHostClusterRequest2WithDefaults() *AdoptHostClusterRequest2`

NewAdoptHostClusterRequest2WithDefaults instantiates a new AdoptHostClusterRequest2 object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetCloudProvider

`func (o *AdoptHostClusterRequest2) GetCloudProvider() string`

GetCloudProvider returns the CloudProvider field if non-nil, zero value otherwise.

### GetCloudProviderOk

`func (o *AdoptHostClusterRequest2) GetCloudProviderOk() (*string, bool)`

GetCloudProviderOk returns a tuple with the CloudProvider field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCloudProvider

`func (o *AdoptHostClusterRequest2) SetCloudProvider(v string)`

SetCloudProvider sets CloudProvider field to given value.


### GetCustomerEmail

`func (o *AdoptHostClusterRequest2) GetCustomerEmail() string`

GetCustomerEmail returns the CustomerEmail field if non-nil, zero value otherwise.

### GetCustomerEmailOk

`func (o *AdoptHostClusterRequest2) GetCustomerEmailOk() (*string, bool)`

GetCustomerEmailOk returns a tuple with the CustomerEmail field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCustomerEmail

`func (o *AdoptHostClusterRequest2) SetCustomerEmail(v string)`

SetCustomerEmail sets CustomerEmail field to given value.

### HasCustomerEmail

`func (o *AdoptHostClusterRequest2) HasCustomerEmail() bool`

HasCustomerEmail returns a boolean if a field has been set.

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


### GetRegion

`func (o *AdoptHostClusterRequest2) GetRegion() string`

GetRegion returns the Region field if non-nil, zero value otherwise.

### GetRegionOk

`func (o *AdoptHostClusterRequest2) GetRegionOk() (*string, bool)`

GetRegionOk returns a tuple with the Region field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRegion

`func (o *AdoptHostClusterRequest2) SetRegion(v string)`

SetRegion sets Region field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


