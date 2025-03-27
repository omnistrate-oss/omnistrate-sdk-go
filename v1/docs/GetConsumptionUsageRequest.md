# GetConsumptionUsageRequest

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**EndDate** | Pointer to **time.Time** | End time of the window in RFC 3339 format. If omitted, the filter is open-ended at the start. | [optional] 
**StartDate** | Pointer to **time.Time** | Start time of the window in RFC 3339 format. If omitted, the filter is open-ended at the start. | [optional] 
**SubscriptionID** | Pointer to **string** | ID of a Subscription | [optional] 
**Token** | **string** | JWT token used to perform authorization | 

## Methods

### NewGetConsumptionUsageRequest

`func NewGetConsumptionUsageRequest(token string, ) *GetConsumptionUsageRequest`

NewGetConsumptionUsageRequest instantiates a new GetConsumptionUsageRequest object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewGetConsumptionUsageRequestWithDefaults

`func NewGetConsumptionUsageRequestWithDefaults() *GetConsumptionUsageRequest`

NewGetConsumptionUsageRequestWithDefaults instantiates a new GetConsumptionUsageRequest object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetEndDate

`func (o *GetConsumptionUsageRequest) GetEndDate() time.Time`

GetEndDate returns the EndDate field if non-nil, zero value otherwise.

### GetEndDateOk

`func (o *GetConsumptionUsageRequest) GetEndDateOk() (*time.Time, bool)`

GetEndDateOk returns a tuple with the EndDate field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEndDate

`func (o *GetConsumptionUsageRequest) SetEndDate(v time.Time)`

SetEndDate sets EndDate field to given value.

### HasEndDate

`func (o *GetConsumptionUsageRequest) HasEndDate() bool`

HasEndDate returns a boolean if a field has been set.

### GetStartDate

`func (o *GetConsumptionUsageRequest) GetStartDate() time.Time`

GetStartDate returns the StartDate field if non-nil, zero value otherwise.

### GetStartDateOk

`func (o *GetConsumptionUsageRequest) GetStartDateOk() (*time.Time, bool)`

GetStartDateOk returns a tuple with the StartDate field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStartDate

`func (o *GetConsumptionUsageRequest) SetStartDate(v time.Time)`

SetStartDate sets StartDate field to given value.

### HasStartDate

`func (o *GetConsumptionUsageRequest) HasStartDate() bool`

HasStartDate returns a boolean if a field has been set.

### GetSubscriptionID

`func (o *GetConsumptionUsageRequest) GetSubscriptionID() string`

GetSubscriptionID returns the SubscriptionID field if non-nil, zero value otherwise.

### GetSubscriptionIDOk

`func (o *GetConsumptionUsageRequest) GetSubscriptionIDOk() (*string, bool)`

GetSubscriptionIDOk returns a tuple with the SubscriptionID field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSubscriptionID

`func (o *GetConsumptionUsageRequest) SetSubscriptionID(v string)`

SetSubscriptionID sets SubscriptionID field to given value.

### HasSubscriptionID

`func (o *GetConsumptionUsageRequest) HasSubscriptionID() bool`

HasSubscriptionID returns a boolean if a field has been set.

### GetToken

`func (o *GetConsumptionUsageRequest) GetToken() string`

GetToken returns the Token field if non-nil, zero value otherwise.

### GetTokenOk

`func (o *GetConsumptionUsageRequest) GetTokenOk() (*string, bool)`

GetTokenOk returns a tuple with the Token field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetToken

`func (o *GetConsumptionUsageRequest) SetToken(v string)`

SetToken sets Token field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


