# CreateProxyResourceInstanceRequest2

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**CloudProvider** | Pointer to **string** | The cloud provider name | [optional] 
**Region** | Pointer to **string** | The region code | [optional] 
**RequestParams** | Pointer to **interface{}** | The request parameters | [optional] 

## Methods

### NewCreateProxyResourceInstanceRequest2

`func NewCreateProxyResourceInstanceRequest2() *CreateProxyResourceInstanceRequest2`

NewCreateProxyResourceInstanceRequest2 instantiates a new CreateProxyResourceInstanceRequest2 object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewCreateProxyResourceInstanceRequest2WithDefaults

`func NewCreateProxyResourceInstanceRequest2WithDefaults() *CreateProxyResourceInstanceRequest2`

NewCreateProxyResourceInstanceRequest2WithDefaults instantiates a new CreateProxyResourceInstanceRequest2 object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetCloudProvider

`func (o *CreateProxyResourceInstanceRequest2) GetCloudProvider() string`

GetCloudProvider returns the CloudProvider field if non-nil, zero value otherwise.

### GetCloudProviderOk

`func (o *CreateProxyResourceInstanceRequest2) GetCloudProviderOk() (*string, bool)`

GetCloudProviderOk returns a tuple with the CloudProvider field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCloudProvider

`func (o *CreateProxyResourceInstanceRequest2) SetCloudProvider(v string)`

SetCloudProvider sets CloudProvider field to given value.

### HasCloudProvider

`func (o *CreateProxyResourceInstanceRequest2) HasCloudProvider() bool`

HasCloudProvider returns a boolean if a field has been set.

### GetRegion

`func (o *CreateProxyResourceInstanceRequest2) GetRegion() string`

GetRegion returns the Region field if non-nil, zero value otherwise.

### GetRegionOk

`func (o *CreateProxyResourceInstanceRequest2) GetRegionOk() (*string, bool)`

GetRegionOk returns a tuple with the Region field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRegion

`func (o *CreateProxyResourceInstanceRequest2) SetRegion(v string)`

SetRegion sets Region field to given value.

### HasRegion

`func (o *CreateProxyResourceInstanceRequest2) HasRegion() bool`

HasRegion returns a boolean if a field has been set.

### GetRequestParams

`func (o *CreateProxyResourceInstanceRequest2) GetRequestParams() interface{}`

GetRequestParams returns the RequestParams field if non-nil, zero value otherwise.

### GetRequestParamsOk

`func (o *CreateProxyResourceInstanceRequest2) GetRequestParamsOk() (*interface{}, bool)`

GetRequestParamsOk returns a tuple with the RequestParams field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRequestParams

`func (o *CreateProxyResourceInstanceRequest2) SetRequestParams(v interface{})`

SetRequestParams sets RequestParams field to given value.

### HasRequestParams

`func (o *CreateProxyResourceInstanceRequest2) HasRequestParams() bool`

HasRequestParams returns a boolean if a field has been set.

### SetRequestParamsNil

`func (o *CreateProxyResourceInstanceRequest2) SetRequestParamsNil(b bool)`

 SetRequestParamsNil sets the value for RequestParams to be an explicit nil

### UnsetRequestParams
`func (o *CreateProxyResourceInstanceRequest2) UnsetRequestParams()`

UnsetRequestParams ensures that no value is present for RequestParams, not even an explicit nil

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


