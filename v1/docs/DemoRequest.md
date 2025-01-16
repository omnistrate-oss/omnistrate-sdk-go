# DemoRequest

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Company** | **string** |  | 
**Email** | **string** |  | 
**Name** | **string** |  | 
**Phone** | Pointer to **string** |  | [optional] 

## Methods

### NewDemoRequest

`func NewDemoRequest(company string, email string, name string, ) *DemoRequest`

NewDemoRequest instantiates a new DemoRequest object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewDemoRequestWithDefaults

`func NewDemoRequestWithDefaults() *DemoRequest`

NewDemoRequestWithDefaults instantiates a new DemoRequest object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetCompany

`func (o *DemoRequest) GetCompany() string`

GetCompany returns the Company field if non-nil, zero value otherwise.

### GetCompanyOk

`func (o *DemoRequest) GetCompanyOk() (*string, bool)`

GetCompanyOk returns a tuple with the Company field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCompany

`func (o *DemoRequest) SetCompany(v string)`

SetCompany sets Company field to given value.


### GetEmail

`func (o *DemoRequest) GetEmail() string`

GetEmail returns the Email field if non-nil, zero value otherwise.

### GetEmailOk

`func (o *DemoRequest) GetEmailOk() (*string, bool)`

GetEmailOk returns a tuple with the Email field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEmail

`func (o *DemoRequest) SetEmail(v string)`

SetEmail sets Email field to given value.


### GetName

`func (o *DemoRequest) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *DemoRequest) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *DemoRequest) SetName(v string)`

SetName sets Name field to given value.


### GetPhone

`func (o *DemoRequest) GetPhone() string`

GetPhone returns the Phone field if non-nil, zero value otherwise.

### GetPhoneOk

`func (o *DemoRequest) GetPhoneOk() (*string, bool)`

GetPhoneOk returns a tuple with the Phone field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPhone

`func (o *DemoRequest) SetPhone(v string)`

SetPhone sets Phone field to given value.

### HasPhone

`func (o *DemoRequest) HasPhone() bool`

HasPhone returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


