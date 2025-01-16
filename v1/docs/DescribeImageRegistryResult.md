# DescribeImageRegistryResult

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Description** | **string** | A brief description of the Image Registry | 
**Host** | **string** | The Image Registry host | 
**Id** | **string** | ID of an Image Registry | 
**Name** | **string** | Name of the Image Registry | 
**Password** | Pointer to **string** | The password to use when authenticating to the Image Registry | [optional] 
**Username** | Pointer to **string** | The username to use when authenticating to the Image Registry | [optional] 

## Methods

### NewDescribeImageRegistryResult

`func NewDescribeImageRegistryResult(description string, host string, id string, name string, ) *DescribeImageRegistryResult`

NewDescribeImageRegistryResult instantiates a new DescribeImageRegistryResult object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewDescribeImageRegistryResultWithDefaults

`func NewDescribeImageRegistryResultWithDefaults() *DescribeImageRegistryResult`

NewDescribeImageRegistryResultWithDefaults instantiates a new DescribeImageRegistryResult object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetDescription

`func (o *DescribeImageRegistryResult) GetDescription() string`

GetDescription returns the Description field if non-nil, zero value otherwise.

### GetDescriptionOk

`func (o *DescribeImageRegistryResult) GetDescriptionOk() (*string, bool)`

GetDescriptionOk returns a tuple with the Description field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDescription

`func (o *DescribeImageRegistryResult) SetDescription(v string)`

SetDescription sets Description field to given value.


### GetHost

`func (o *DescribeImageRegistryResult) GetHost() string`

GetHost returns the Host field if non-nil, zero value otherwise.

### GetHostOk

`func (o *DescribeImageRegistryResult) GetHostOk() (*string, bool)`

GetHostOk returns a tuple with the Host field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetHost

`func (o *DescribeImageRegistryResult) SetHost(v string)`

SetHost sets Host field to given value.


### GetId

`func (o *DescribeImageRegistryResult) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *DescribeImageRegistryResult) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *DescribeImageRegistryResult) SetId(v string)`

SetId sets Id field to given value.


### GetName

`func (o *DescribeImageRegistryResult) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *DescribeImageRegistryResult) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *DescribeImageRegistryResult) SetName(v string)`

SetName sets Name field to given value.


### GetPassword

`func (o *DescribeImageRegistryResult) GetPassword() string`

GetPassword returns the Password field if non-nil, zero value otherwise.

### GetPasswordOk

`func (o *DescribeImageRegistryResult) GetPasswordOk() (*string, bool)`

GetPasswordOk returns a tuple with the Password field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPassword

`func (o *DescribeImageRegistryResult) SetPassword(v string)`

SetPassword sets Password field to given value.

### HasPassword

`func (o *DescribeImageRegistryResult) HasPassword() bool`

HasPassword returns a boolean if a field has been set.

### GetUsername

`func (o *DescribeImageRegistryResult) GetUsername() string`

GetUsername returns the Username field if non-nil, zero value otherwise.

### GetUsernameOk

`func (o *DescribeImageRegistryResult) GetUsernameOk() (*string, bool)`

GetUsernameOk returns a tuple with the Username field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUsername

`func (o *DescribeImageRegistryResult) SetUsername(v string)`

SetUsername sets Username field to given value.

### HasUsername

`func (o *DescribeImageRegistryResult) HasUsername() bool`

HasUsername returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


