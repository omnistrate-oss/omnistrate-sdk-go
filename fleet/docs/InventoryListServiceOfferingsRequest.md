# InventoryListServiceOfferingsRequest

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Token** | **string** | JWT token used to perform authorization | 
**Visibility** | Pointer to **string** | This parameter is used to configure the visibility of the service control-plane APIs | [optional] 

## Methods

### NewInventoryListServiceOfferingsRequest

`func NewInventoryListServiceOfferingsRequest(token string, ) *InventoryListServiceOfferingsRequest`

NewInventoryListServiceOfferingsRequest instantiates a new InventoryListServiceOfferingsRequest object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewInventoryListServiceOfferingsRequestWithDefaults

`func NewInventoryListServiceOfferingsRequestWithDefaults() *InventoryListServiceOfferingsRequest`

NewInventoryListServiceOfferingsRequestWithDefaults instantiates a new InventoryListServiceOfferingsRequest object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetToken

`func (o *InventoryListServiceOfferingsRequest) GetToken() string`

GetToken returns the Token field if non-nil, zero value otherwise.

### GetTokenOk

`func (o *InventoryListServiceOfferingsRequest) GetTokenOk() (*string, bool)`

GetTokenOk returns a tuple with the Token field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetToken

`func (o *InventoryListServiceOfferingsRequest) SetToken(v string)`

SetToken sets Token field to given value.


### GetVisibility

`func (o *InventoryListServiceOfferingsRequest) GetVisibility() string`

GetVisibility returns the Visibility field if non-nil, zero value otherwise.

### GetVisibilityOk

`func (o *InventoryListServiceOfferingsRequest) GetVisibilityOk() (*string, bool)`

GetVisibilityOk returns a tuple with the Visibility field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVisibility

`func (o *InventoryListServiceOfferingsRequest) SetVisibility(v string)`

SetVisibility sets Visibility field to given value.

### HasVisibility

`func (o *InventoryListServiceOfferingsRequest) HasVisibility() bool`

HasVisibility returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


