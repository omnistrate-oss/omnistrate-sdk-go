# GetCurrentConsumptionUsageRequest

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**SubscriptionID** | Pointer to **string** | ID of a Subscription | [optional] 
**Token** | **string** | JWT token used to perform authorization | 

## Methods

### NewGetCurrentConsumptionUsageRequest

`func NewGetCurrentConsumptionUsageRequest(token string, ) *GetCurrentConsumptionUsageRequest`

NewGetCurrentConsumptionUsageRequest instantiates a new GetCurrentConsumptionUsageRequest object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewGetCurrentConsumptionUsageRequestWithDefaults

`func NewGetCurrentConsumptionUsageRequestWithDefaults() *GetCurrentConsumptionUsageRequest`

NewGetCurrentConsumptionUsageRequestWithDefaults instantiates a new GetCurrentConsumptionUsageRequest object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetSubscriptionID

`func (o *GetCurrentConsumptionUsageRequest) GetSubscriptionID() string`

GetSubscriptionID returns the SubscriptionID field if non-nil, zero value otherwise.

### GetSubscriptionIDOk

`func (o *GetCurrentConsumptionUsageRequest) GetSubscriptionIDOk() (*string, bool)`

GetSubscriptionIDOk returns a tuple with the SubscriptionID field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSubscriptionID

`func (o *GetCurrentConsumptionUsageRequest) SetSubscriptionID(v string)`

SetSubscriptionID sets SubscriptionID field to given value.

### HasSubscriptionID

`func (o *GetCurrentConsumptionUsageRequest) HasSubscriptionID() bool`

HasSubscriptionID returns a boolean if a field has been set.

### GetToken

`func (o *GetCurrentConsumptionUsageRequest) GetToken() string`

GetToken returns the Token field if non-nil, zero value otherwise.

### GetTokenOk

`func (o *GetCurrentConsumptionUsageRequest) GetTokenOk() (*string, bool)`

GetTokenOk returns a tuple with the Token field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetToken

`func (o *GetCurrentConsumptionUsageRequest) SetToken(v string)`

SetToken sets Token field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


