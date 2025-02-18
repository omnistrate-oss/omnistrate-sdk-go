# ListServiceProviderEvents

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**EndDate** | Pointer to **time.Time** | End time of the window in ISO 8601 format. If omitted, the filter is open-ended at the start. | [optional] 
**EnvironmentType** | Pointer to **string** | The type of service environment | [optional] 
**EventTypes** | Pointer to **[]string** | The event types to filter by | [optional] 
**InstanceID** | Pointer to **string** | ID of a Resource Instance | [optional] 
**NextPageToken** | Pointer to **string** | The next token to use for pagination | [optional] 
**PageSize** | Pointer to **int64** | The number of resources to return per page | [optional] 
**ProductTierID** | Pointer to **string** | ID of a Product Tier | [optional] 
**ServiceEnvironmentID** | Pointer to **string** | ID of a Service Environment | [optional] 
**ServiceID** | Pointer to **string** | ID of a Service | [optional] 
**StartDate** | Pointer to **time.Time** | Start time of the window in ISO 8601 format. If omitted, the filter is open-ended at the start. | [optional] 
**Token** | **string** | JWT token used to perform authorization | 

## Methods

### NewListServiceProviderEvents

`func NewListServiceProviderEvents(token string, ) *ListServiceProviderEvents`

NewListServiceProviderEvents instantiates a new ListServiceProviderEvents object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewListServiceProviderEventsWithDefaults

`func NewListServiceProviderEventsWithDefaults() *ListServiceProviderEvents`

NewListServiceProviderEventsWithDefaults instantiates a new ListServiceProviderEvents object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetEndDate

`func (o *ListServiceProviderEvents) GetEndDate() time.Time`

GetEndDate returns the EndDate field if non-nil, zero value otherwise.

### GetEndDateOk

`func (o *ListServiceProviderEvents) GetEndDateOk() (*time.Time, bool)`

GetEndDateOk returns a tuple with the EndDate field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEndDate

`func (o *ListServiceProviderEvents) SetEndDate(v time.Time)`

SetEndDate sets EndDate field to given value.

### HasEndDate

`func (o *ListServiceProviderEvents) HasEndDate() bool`

HasEndDate returns a boolean if a field has been set.

### GetEnvironmentType

`func (o *ListServiceProviderEvents) GetEnvironmentType() string`

GetEnvironmentType returns the EnvironmentType field if non-nil, zero value otherwise.

### GetEnvironmentTypeOk

`func (o *ListServiceProviderEvents) GetEnvironmentTypeOk() (*string, bool)`

GetEnvironmentTypeOk returns a tuple with the EnvironmentType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEnvironmentType

`func (o *ListServiceProviderEvents) SetEnvironmentType(v string)`

SetEnvironmentType sets EnvironmentType field to given value.

### HasEnvironmentType

`func (o *ListServiceProviderEvents) HasEnvironmentType() bool`

HasEnvironmentType returns a boolean if a field has been set.

### GetEventTypes

`func (o *ListServiceProviderEvents) GetEventTypes() []string`

GetEventTypes returns the EventTypes field if non-nil, zero value otherwise.

### GetEventTypesOk

`func (o *ListServiceProviderEvents) GetEventTypesOk() (*[]string, bool)`

GetEventTypesOk returns a tuple with the EventTypes field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEventTypes

`func (o *ListServiceProviderEvents) SetEventTypes(v []string)`

SetEventTypes sets EventTypes field to given value.

### HasEventTypes

`func (o *ListServiceProviderEvents) HasEventTypes() bool`

HasEventTypes returns a boolean if a field has been set.

### GetInstanceID

`func (o *ListServiceProviderEvents) GetInstanceID() string`

GetInstanceID returns the InstanceID field if non-nil, zero value otherwise.

### GetInstanceIDOk

`func (o *ListServiceProviderEvents) GetInstanceIDOk() (*string, bool)`

GetInstanceIDOk returns a tuple with the InstanceID field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetInstanceID

`func (o *ListServiceProviderEvents) SetInstanceID(v string)`

SetInstanceID sets InstanceID field to given value.

### HasInstanceID

`func (o *ListServiceProviderEvents) HasInstanceID() bool`

HasInstanceID returns a boolean if a field has been set.

### GetNextPageToken

`func (o *ListServiceProviderEvents) GetNextPageToken() string`

GetNextPageToken returns the NextPageToken field if non-nil, zero value otherwise.

### GetNextPageTokenOk

`func (o *ListServiceProviderEvents) GetNextPageTokenOk() (*string, bool)`

GetNextPageTokenOk returns a tuple with the NextPageToken field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNextPageToken

`func (o *ListServiceProviderEvents) SetNextPageToken(v string)`

SetNextPageToken sets NextPageToken field to given value.

### HasNextPageToken

`func (o *ListServiceProviderEvents) HasNextPageToken() bool`

HasNextPageToken returns a boolean if a field has been set.

### GetPageSize

`func (o *ListServiceProviderEvents) GetPageSize() int64`

GetPageSize returns the PageSize field if non-nil, zero value otherwise.

### GetPageSizeOk

`func (o *ListServiceProviderEvents) GetPageSizeOk() (*int64, bool)`

GetPageSizeOk returns a tuple with the PageSize field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPageSize

`func (o *ListServiceProviderEvents) SetPageSize(v int64)`

SetPageSize sets PageSize field to given value.

### HasPageSize

`func (o *ListServiceProviderEvents) HasPageSize() bool`

HasPageSize returns a boolean if a field has been set.

### GetProductTierID

`func (o *ListServiceProviderEvents) GetProductTierID() string`

GetProductTierID returns the ProductTierID field if non-nil, zero value otherwise.

### GetProductTierIDOk

`func (o *ListServiceProviderEvents) GetProductTierIDOk() (*string, bool)`

GetProductTierIDOk returns a tuple with the ProductTierID field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetProductTierID

`func (o *ListServiceProviderEvents) SetProductTierID(v string)`

SetProductTierID sets ProductTierID field to given value.

### HasProductTierID

`func (o *ListServiceProviderEvents) HasProductTierID() bool`

HasProductTierID returns a boolean if a field has been set.

### GetServiceEnvironmentID

`func (o *ListServiceProviderEvents) GetServiceEnvironmentID() string`

GetServiceEnvironmentID returns the ServiceEnvironmentID field if non-nil, zero value otherwise.

### GetServiceEnvironmentIDOk

`func (o *ListServiceProviderEvents) GetServiceEnvironmentIDOk() (*string, bool)`

GetServiceEnvironmentIDOk returns a tuple with the ServiceEnvironmentID field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetServiceEnvironmentID

`func (o *ListServiceProviderEvents) SetServiceEnvironmentID(v string)`

SetServiceEnvironmentID sets ServiceEnvironmentID field to given value.

### HasServiceEnvironmentID

`func (o *ListServiceProviderEvents) HasServiceEnvironmentID() bool`

HasServiceEnvironmentID returns a boolean if a field has been set.

### GetServiceID

`func (o *ListServiceProviderEvents) GetServiceID() string`

GetServiceID returns the ServiceID field if non-nil, zero value otherwise.

### GetServiceIDOk

`func (o *ListServiceProviderEvents) GetServiceIDOk() (*string, bool)`

GetServiceIDOk returns a tuple with the ServiceID field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetServiceID

`func (o *ListServiceProviderEvents) SetServiceID(v string)`

SetServiceID sets ServiceID field to given value.

### HasServiceID

`func (o *ListServiceProviderEvents) HasServiceID() bool`

HasServiceID returns a boolean if a field has been set.

### GetStartDate

`func (o *ListServiceProviderEvents) GetStartDate() time.Time`

GetStartDate returns the StartDate field if non-nil, zero value otherwise.

### GetStartDateOk

`func (o *ListServiceProviderEvents) GetStartDateOk() (*time.Time, bool)`

GetStartDateOk returns a tuple with the StartDate field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStartDate

`func (o *ListServiceProviderEvents) SetStartDate(v time.Time)`

SetStartDate sets StartDate field to given value.

### HasStartDate

`func (o *ListServiceProviderEvents) HasStartDate() bool`

HasStartDate returns a boolean if a field has been set.

### GetToken

`func (o *ListServiceProviderEvents) GetToken() string`

GetToken returns the Token field if non-nil, zero value otherwise.

### GetTokenOk

`func (o *ListServiceProviderEvents) GetTokenOk() (*string, bool)`

GetTokenOk returns a tuple with the Token field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetToken

`func (o *ListServiceProviderEvents) SetToken(v string)`

SetToken sets Token field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


