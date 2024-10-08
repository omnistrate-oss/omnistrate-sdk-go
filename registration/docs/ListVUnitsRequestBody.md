# ListVUnitsRequestBody

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**CloudProvider** | **string** | The infra / cloud provider name | 
**Region** | **string** | Region code specific to the cloud-provider | 

## Methods

### NewListVUnitsRequestBody

`func NewListVUnitsRequestBody(cloudProvider string, region string, ) *ListVUnitsRequestBody`

NewListVUnitsRequestBody instantiates a new ListVUnitsRequestBody object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewListVUnitsRequestBodyWithDefaults

`func NewListVUnitsRequestBodyWithDefaults() *ListVUnitsRequestBody`

NewListVUnitsRequestBodyWithDefaults instantiates a new ListVUnitsRequestBody object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetCloudProvider

`func (o *ListVUnitsRequestBody) GetCloudProvider() string`

GetCloudProvider returns the CloudProvider field if non-nil, zero value otherwise.

### GetCloudProviderOk

`func (o *ListVUnitsRequestBody) GetCloudProviderOk() (*string, bool)`

GetCloudProviderOk returns a tuple with the CloudProvider field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCloudProvider

`func (o *ListVUnitsRequestBody) SetCloudProvider(v string)`

SetCloudProvider sets CloudProvider field to given value.


### GetRegion

`func (o *ListVUnitsRequestBody) GetRegion() string`

GetRegion returns the Region field if non-nil, zero value otherwise.

### GetRegionOk

`func (o *ListVUnitsRequestBody) GetRegionOk() (*string, bool)`

GetRegionOk returns a tuple with the Region field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRegion

`func (o *ListVUnitsRequestBody) SetRegion(v string)`

SetRegion sets Region field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


