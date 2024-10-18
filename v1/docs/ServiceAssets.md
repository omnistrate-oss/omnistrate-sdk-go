# ServiceAssets

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**CloudFormationURL** | Pointer to **string** | The URL for CloudFormation onboarding | [optional] 
**CloudFormationURLNoLB** | Pointer to **string** | The URL for CloudFormation onboarding without LoadBalancer policy | [optional] 

## Methods

### NewServiceAssets

`func NewServiceAssets() *ServiceAssets`

NewServiceAssets instantiates a new ServiceAssets object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewServiceAssetsWithDefaults

`func NewServiceAssetsWithDefaults() *ServiceAssets`

NewServiceAssetsWithDefaults instantiates a new ServiceAssets object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetCloudFormationURL

`func (o *ServiceAssets) GetCloudFormationURL() string`

GetCloudFormationURL returns the CloudFormationURL field if non-nil, zero value otherwise.

### GetCloudFormationURLOk

`func (o *ServiceAssets) GetCloudFormationURLOk() (*string, bool)`

GetCloudFormationURLOk returns a tuple with the CloudFormationURL field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCloudFormationURL

`func (o *ServiceAssets) SetCloudFormationURL(v string)`

SetCloudFormationURL sets CloudFormationURL field to given value.

### HasCloudFormationURL

`func (o *ServiceAssets) HasCloudFormationURL() bool`

HasCloudFormationURL returns a boolean if a field has been set.

### GetCloudFormationURLNoLB

`func (o *ServiceAssets) GetCloudFormationURLNoLB() string`

GetCloudFormationURLNoLB returns the CloudFormationURLNoLB field if non-nil, zero value otherwise.

### GetCloudFormationURLNoLBOk

`func (o *ServiceAssets) GetCloudFormationURLNoLBOk() (*string, bool)`

GetCloudFormationURLNoLBOk returns a tuple with the CloudFormationURLNoLB field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCloudFormationURLNoLB

`func (o *ServiceAssets) SetCloudFormationURLNoLB(v string)`

SetCloudFormationURLNoLB sets CloudFormationURLNoLB field to given value.

### HasCloudFormationURLNoLB

`func (o *ServiceAssets) HasCloudFormationURLNoLB() bool`

HasCloudFormationURLNoLB returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


