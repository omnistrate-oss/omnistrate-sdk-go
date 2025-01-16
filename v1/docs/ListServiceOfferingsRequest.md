# ListServiceOfferingsRequest

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**EnvironmentType** | Pointer to **string** | The type of service environment | [optional] 
**OrgId** | Pointer to **string** | ID of an Org | [optional] 
**Token** | **string** | JWT token used to perform authorization | 
**Visibility** | Pointer to **string** | This parameter is used to configure the visibility of the service control-plane APIs | [optional] 

## Methods

### NewListServiceOfferingsRequest

`func NewListServiceOfferingsRequest(token string, ) *ListServiceOfferingsRequest`

NewListServiceOfferingsRequest instantiates a new ListServiceOfferingsRequest object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewListServiceOfferingsRequestWithDefaults

`func NewListServiceOfferingsRequestWithDefaults() *ListServiceOfferingsRequest`

NewListServiceOfferingsRequestWithDefaults instantiates a new ListServiceOfferingsRequest object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetEnvironmentType

`func (o *ListServiceOfferingsRequest) GetEnvironmentType() string`

GetEnvironmentType returns the EnvironmentType field if non-nil, zero value otherwise.

### GetEnvironmentTypeOk

`func (o *ListServiceOfferingsRequest) GetEnvironmentTypeOk() (*string, bool)`

GetEnvironmentTypeOk returns a tuple with the EnvironmentType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEnvironmentType

`func (o *ListServiceOfferingsRequest) SetEnvironmentType(v string)`

SetEnvironmentType sets EnvironmentType field to given value.

### HasEnvironmentType

`func (o *ListServiceOfferingsRequest) HasEnvironmentType() bool`

HasEnvironmentType returns a boolean if a field has been set.

### GetOrgId

`func (o *ListServiceOfferingsRequest) GetOrgId() string`

GetOrgId returns the OrgId field if non-nil, zero value otherwise.

### GetOrgIdOk

`func (o *ListServiceOfferingsRequest) GetOrgIdOk() (*string, bool)`

GetOrgIdOk returns a tuple with the OrgId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOrgId

`func (o *ListServiceOfferingsRequest) SetOrgId(v string)`

SetOrgId sets OrgId field to given value.

### HasOrgId

`func (o *ListServiceOfferingsRequest) HasOrgId() bool`

HasOrgId returns a boolean if a field has been set.

### GetToken

`func (o *ListServiceOfferingsRequest) GetToken() string`

GetToken returns the Token field if non-nil, zero value otherwise.

### GetTokenOk

`func (o *ListServiceOfferingsRequest) GetTokenOk() (*string, bool)`

GetTokenOk returns a tuple with the Token field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetToken

`func (o *ListServiceOfferingsRequest) SetToken(v string)`

SetToken sets Token field to given value.


### GetVisibility

`func (o *ListServiceOfferingsRequest) GetVisibility() string`

GetVisibility returns the Visibility field if non-nil, zero value otherwise.

### GetVisibilityOk

`func (o *ListServiceOfferingsRequest) GetVisibilityOk() (*string, bool)`

GetVisibilityOk returns a tuple with the Visibility field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVisibility

`func (o *ListServiceOfferingsRequest) SetVisibility(v string)`

SetVisibility sets Visibility field to given value.

### HasVisibility

`func (o *ListServiceOfferingsRequest) HasVisibility() bool`

HasVisibility returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


