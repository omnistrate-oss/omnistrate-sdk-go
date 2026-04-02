# DescribeAccountConfigByNebiusTenantIDRequest

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**NebiusTenantID** | **string** | The Nebius tenant ID | 
**Token** | **string** | JWT token used to perform authorization | 

## Methods

### NewDescribeAccountConfigByNebiusTenantIDRequest

`func NewDescribeAccountConfigByNebiusTenantIDRequest(nebiusTenantID string, token string, ) *DescribeAccountConfigByNebiusTenantIDRequest`

NewDescribeAccountConfigByNebiusTenantIDRequest instantiates a new DescribeAccountConfigByNebiusTenantIDRequest object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewDescribeAccountConfigByNebiusTenantIDRequestWithDefaults

`func NewDescribeAccountConfigByNebiusTenantIDRequestWithDefaults() *DescribeAccountConfigByNebiusTenantIDRequest`

NewDescribeAccountConfigByNebiusTenantIDRequestWithDefaults instantiates a new DescribeAccountConfigByNebiusTenantIDRequest object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetNebiusTenantID

`func (o *DescribeAccountConfigByNebiusTenantIDRequest) GetNebiusTenantID() string`

GetNebiusTenantID returns the NebiusTenantID field if non-nil, zero value otherwise.

### GetNebiusTenantIDOk

`func (o *DescribeAccountConfigByNebiusTenantIDRequest) GetNebiusTenantIDOk() (*string, bool)`

GetNebiusTenantIDOk returns a tuple with the NebiusTenantID field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNebiusTenantID

`func (o *DescribeAccountConfigByNebiusTenantIDRequest) SetNebiusTenantID(v string)`

SetNebiusTenantID sets NebiusTenantID field to given value.


### GetToken

`func (o *DescribeAccountConfigByNebiusTenantIDRequest) GetToken() string`

GetToken returns the Token field if non-nil, zero value otherwise.

### GetTokenOk

`func (o *DescribeAccountConfigByNebiusTenantIDRequest) GetTokenOk() (*string, bool)`

GetTokenOk returns a tuple with the Token field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetToken

`func (o *DescribeAccountConfigByNebiusTenantIDRequest) SetToken(v string)`

SetToken sets Token field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


