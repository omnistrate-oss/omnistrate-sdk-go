# UpdateAPIKeyMetadataRequest

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Description** | Pointer to **string** | New description. Optional; nil leaves the field unchanged. | [optional] 
**Id** | **string** | ID of an API Key | 
**Name** | Pointer to **string** | New display name. Optional; nil leaves the field unchanged. Re-validated for uniqueness within the org on change (409 on conflict). | [optional] 
**Token** | **string** | JWT token used to perform authorization | 

## Methods

### NewUpdateAPIKeyMetadataRequest

`func NewUpdateAPIKeyMetadataRequest(id string, token string, ) *UpdateAPIKeyMetadataRequest`

NewUpdateAPIKeyMetadataRequest instantiates a new UpdateAPIKeyMetadataRequest object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewUpdateAPIKeyMetadataRequestWithDefaults

`func NewUpdateAPIKeyMetadataRequestWithDefaults() *UpdateAPIKeyMetadataRequest`

NewUpdateAPIKeyMetadataRequestWithDefaults instantiates a new UpdateAPIKeyMetadataRequest object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetDescription

`func (o *UpdateAPIKeyMetadataRequest) GetDescription() string`

GetDescription returns the Description field if non-nil, zero value otherwise.

### GetDescriptionOk

`func (o *UpdateAPIKeyMetadataRequest) GetDescriptionOk() (*string, bool)`

GetDescriptionOk returns a tuple with the Description field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDescription

`func (o *UpdateAPIKeyMetadataRequest) SetDescription(v string)`

SetDescription sets Description field to given value.

### HasDescription

`func (o *UpdateAPIKeyMetadataRequest) HasDescription() bool`

HasDescription returns a boolean if a field has been set.

### GetId

`func (o *UpdateAPIKeyMetadataRequest) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *UpdateAPIKeyMetadataRequest) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *UpdateAPIKeyMetadataRequest) SetId(v string)`

SetId sets Id field to given value.


### GetName

`func (o *UpdateAPIKeyMetadataRequest) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *UpdateAPIKeyMetadataRequest) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *UpdateAPIKeyMetadataRequest) SetName(v string)`

SetName sets Name field to given value.

### HasName

`func (o *UpdateAPIKeyMetadataRequest) HasName() bool`

HasName returns a boolean if a field has been set.

### GetToken

`func (o *UpdateAPIKeyMetadataRequest) GetToken() string`

GetToken returns the Token field if non-nil, zero value otherwise.

### GetTokenOk

`func (o *UpdateAPIKeyMetadataRequest) GetTokenOk() (*string, bool)`

GetTokenOk returns a tuple with the Token field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetToken

`func (o *UpdateAPIKeyMetadataRequest) SetToken(v string)`

SetToken sets Token field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


