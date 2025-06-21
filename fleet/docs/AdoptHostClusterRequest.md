# AdoptHostClusterRequest

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**CloudProviderId** | **string** | ID of an CloudProvider | 
**CustomerEmail** | Pointer to **string** | Email of the customer who owns the host cluster in case this is a BYOA host cluster | [optional] 
**Description** | **string** | Description of the host cluster | 
**Id** | **string** | ID of a Host Cluster | 
**RegionId** | **string** | ID of a Region | 
**Token** | **string** | JWT token used to perform authorization | 

## Methods

### NewAdoptHostClusterRequest

`func NewAdoptHostClusterRequest(cloudProviderId string, description string, id string, regionId string, token string, ) *AdoptHostClusterRequest`

NewAdoptHostClusterRequest instantiates a new AdoptHostClusterRequest object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewAdoptHostClusterRequestWithDefaults

`func NewAdoptHostClusterRequestWithDefaults() *AdoptHostClusterRequest`

NewAdoptHostClusterRequestWithDefaults instantiates a new AdoptHostClusterRequest object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetCloudProviderId

`func (o *AdoptHostClusterRequest) GetCloudProviderId() string`

GetCloudProviderId returns the CloudProviderId field if non-nil, zero value otherwise.

### GetCloudProviderIdOk

`func (o *AdoptHostClusterRequest) GetCloudProviderIdOk() (*string, bool)`

GetCloudProviderIdOk returns a tuple with the CloudProviderId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCloudProviderId

`func (o *AdoptHostClusterRequest) SetCloudProviderId(v string)`

SetCloudProviderId sets CloudProviderId field to given value.


### GetCustomerEmail

`func (o *AdoptHostClusterRequest) GetCustomerEmail() string`

GetCustomerEmail returns the CustomerEmail field if non-nil, zero value otherwise.

### GetCustomerEmailOk

`func (o *AdoptHostClusterRequest) GetCustomerEmailOk() (*string, bool)`

GetCustomerEmailOk returns a tuple with the CustomerEmail field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCustomerEmail

`func (o *AdoptHostClusterRequest) SetCustomerEmail(v string)`

SetCustomerEmail sets CustomerEmail field to given value.

### HasCustomerEmail

`func (o *AdoptHostClusterRequest) HasCustomerEmail() bool`

HasCustomerEmail returns a boolean if a field has been set.

### GetDescription

`func (o *AdoptHostClusterRequest) GetDescription() string`

GetDescription returns the Description field if non-nil, zero value otherwise.

### GetDescriptionOk

`func (o *AdoptHostClusterRequest) GetDescriptionOk() (*string, bool)`

GetDescriptionOk returns a tuple with the Description field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDescription

`func (o *AdoptHostClusterRequest) SetDescription(v string)`

SetDescription sets Description field to given value.


### GetId

`func (o *AdoptHostClusterRequest) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *AdoptHostClusterRequest) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *AdoptHostClusterRequest) SetId(v string)`

SetId sets Id field to given value.


### GetRegionId

`func (o *AdoptHostClusterRequest) GetRegionId() string`

GetRegionId returns the RegionId field if non-nil, zero value otherwise.

### GetRegionIdOk

`func (o *AdoptHostClusterRequest) GetRegionIdOk() (*string, bool)`

GetRegionIdOk returns a tuple with the RegionId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRegionId

`func (o *AdoptHostClusterRequest) SetRegionId(v string)`

SetRegionId sets RegionId field to given value.


### GetToken

`func (o *AdoptHostClusterRequest) GetToken() string`

GetToken returns the Token field if non-nil, zero value otherwise.

### GetTokenOk

`func (o *AdoptHostClusterRequest) GetTokenOk() (*string, bool)`

GetTokenOk returns a tuple with the Token field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetToken

`func (o *AdoptHostClusterRequest) SetToken(v string)`

SetToken sets Token field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


