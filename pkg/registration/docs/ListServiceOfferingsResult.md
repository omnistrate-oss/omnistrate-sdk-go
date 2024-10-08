# ListServiceOfferingsResult

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**NextPageToken** | Pointer to **string** | Token to use for the next request | [optional] 
**ServiceIds** | **[]string** | The service IDs | 
**Services** | Pointer to [**[]DescribeServiceOfferingResult**](DescribeServiceOfferingResult.md) | List of services | [optional] 

## Methods

### NewListServiceOfferingsResult

`func NewListServiceOfferingsResult(serviceIds []string, ) *ListServiceOfferingsResult`

NewListServiceOfferingsResult instantiates a new ListServiceOfferingsResult object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewListServiceOfferingsResultWithDefaults

`func NewListServiceOfferingsResultWithDefaults() *ListServiceOfferingsResult`

NewListServiceOfferingsResultWithDefaults instantiates a new ListServiceOfferingsResult object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetNextPageToken

`func (o *ListServiceOfferingsResult) GetNextPageToken() string`

GetNextPageToken returns the NextPageToken field if non-nil, zero value otherwise.

### GetNextPageTokenOk

`func (o *ListServiceOfferingsResult) GetNextPageTokenOk() (*string, bool)`

GetNextPageTokenOk returns a tuple with the NextPageToken field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNextPageToken

`func (o *ListServiceOfferingsResult) SetNextPageToken(v string)`

SetNextPageToken sets NextPageToken field to given value.

### HasNextPageToken

`func (o *ListServiceOfferingsResult) HasNextPageToken() bool`

HasNextPageToken returns a boolean if a field has been set.

### GetServiceIds

`func (o *ListServiceOfferingsResult) GetServiceIds() []string`

GetServiceIds returns the ServiceIds field if non-nil, zero value otherwise.

### GetServiceIdsOk

`func (o *ListServiceOfferingsResult) GetServiceIdsOk() (*[]string, bool)`

GetServiceIdsOk returns a tuple with the ServiceIds field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetServiceIds

`func (o *ListServiceOfferingsResult) SetServiceIds(v []string)`

SetServiceIds sets ServiceIds field to given value.


### GetServices

`func (o *ListServiceOfferingsResult) GetServices() []DescribeServiceOfferingResult`

GetServices returns the Services field if non-nil, zero value otherwise.

### GetServicesOk

`func (o *ListServiceOfferingsResult) GetServicesOk() (*[]DescribeServiceOfferingResult, bool)`

GetServicesOk returns a tuple with the Services field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetServices

`func (o *ListServiceOfferingsResult) SetServices(v []DescribeServiceOfferingResult)`

SetServices sets Services field to given value.

### HasServices

`func (o *ListServiceOfferingsResult) HasServices() bool`

HasServices returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


