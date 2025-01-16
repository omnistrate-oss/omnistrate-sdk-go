# UpdateCustomNetworkRequest

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Id** | **string** | ID of a custom network | 
**Name** | Pointer to **string** | User friendly network name to help distinguish networks with same CIDRs | [optional] 
**Token** | **string** | JWT token used to perform authorization | 

## Methods

### NewUpdateCustomNetworkRequest

`func NewUpdateCustomNetworkRequest(id string, token string, ) *UpdateCustomNetworkRequest`

NewUpdateCustomNetworkRequest instantiates a new UpdateCustomNetworkRequest object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewUpdateCustomNetworkRequestWithDefaults

`func NewUpdateCustomNetworkRequestWithDefaults() *UpdateCustomNetworkRequest`

NewUpdateCustomNetworkRequestWithDefaults instantiates a new UpdateCustomNetworkRequest object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetId

`func (o *UpdateCustomNetworkRequest) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *UpdateCustomNetworkRequest) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *UpdateCustomNetworkRequest) SetId(v string)`

SetId sets Id field to given value.


### GetName

`func (o *UpdateCustomNetworkRequest) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *UpdateCustomNetworkRequest) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *UpdateCustomNetworkRequest) SetName(v string)`

SetName sets Name field to given value.

### HasName

`func (o *UpdateCustomNetworkRequest) HasName() bool`

HasName returns a boolean if a field has been set.

### GetToken

`func (o *UpdateCustomNetworkRequest) GetToken() string`

GetToken returns the Token field if non-nil, zero value otherwise.

### GetTokenOk

`func (o *UpdateCustomNetworkRequest) GetTokenOk() (*string, bool)`

GetTokenOk returns a tuple with the Token field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetToken

`func (o *UpdateCustomNetworkRequest) SetToken(v string)`

SetToken sets Token field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


