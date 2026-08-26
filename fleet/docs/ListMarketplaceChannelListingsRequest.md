# ListMarketplaceChannelListingsRequest

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Channel** | **string** | Which marketplace channel a contract came from. SUGER reaches AWS, Azure and GCP buyers through one listing. SANDBOX is the simulated channel, and is a real member of this set rather than a test mode | 
**Token** | **string** | JWT token used to perform authorization | 

## Methods

### NewListMarketplaceChannelListingsRequest

`func NewListMarketplaceChannelListingsRequest(channel string, token string, ) *ListMarketplaceChannelListingsRequest`

NewListMarketplaceChannelListingsRequest instantiates a new ListMarketplaceChannelListingsRequest object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewListMarketplaceChannelListingsRequestWithDefaults

`func NewListMarketplaceChannelListingsRequestWithDefaults() *ListMarketplaceChannelListingsRequest`

NewListMarketplaceChannelListingsRequestWithDefaults instantiates a new ListMarketplaceChannelListingsRequest object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetChannel

`func (o *ListMarketplaceChannelListingsRequest) GetChannel() string`

GetChannel returns the Channel field if non-nil, zero value otherwise.

### GetChannelOk

`func (o *ListMarketplaceChannelListingsRequest) GetChannelOk() (*string, bool)`

GetChannelOk returns a tuple with the Channel field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetChannel

`func (o *ListMarketplaceChannelListingsRequest) SetChannel(v string)`

SetChannel sets Channel field to given value.


### GetToken

`func (o *ListMarketplaceChannelListingsRequest) GetToken() string`

GetToken returns the Token field if non-nil, zero value otherwise.

### GetTokenOk

`func (o *ListMarketplaceChannelListingsRequest) GetTokenOk() (*string, bool)`

GetTokenOk returns a tuple with the Token field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetToken

`func (o *ListMarketplaceChannelListingsRequest) SetToken(v string)`

SetToken sets Token field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


