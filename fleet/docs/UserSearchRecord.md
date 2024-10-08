# UserSearchRecord

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Email** | **string** | The user email. | 
**External** | **bool** | Is the user an external user. | 
**Id** | **string** | The user ID. | 
**Name** | **string** | The user name. | 
**OrgName** | **string** | The organization name. | 

## Methods

### NewUserSearchRecord

`func NewUserSearchRecord(email string, external bool, id string, name string, orgName string, ) *UserSearchRecord`

NewUserSearchRecord instantiates a new UserSearchRecord object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewUserSearchRecordWithDefaults

`func NewUserSearchRecordWithDefaults() *UserSearchRecord`

NewUserSearchRecordWithDefaults instantiates a new UserSearchRecord object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetEmail

`func (o *UserSearchRecord) GetEmail() string`

GetEmail returns the Email field if non-nil, zero value otherwise.

### GetEmailOk

`func (o *UserSearchRecord) GetEmailOk() (*string, bool)`

GetEmailOk returns a tuple with the Email field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEmail

`func (o *UserSearchRecord) SetEmail(v string)`

SetEmail sets Email field to given value.


### GetExternal

`func (o *UserSearchRecord) GetExternal() bool`

GetExternal returns the External field if non-nil, zero value otherwise.

### GetExternalOk

`func (o *UserSearchRecord) GetExternalOk() (*bool, bool)`

GetExternalOk returns a tuple with the External field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetExternal

`func (o *UserSearchRecord) SetExternal(v bool)`

SetExternal sets External field to given value.


### GetId

`func (o *UserSearchRecord) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *UserSearchRecord) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *UserSearchRecord) SetId(v string)`

SetId sets Id field to given value.


### GetName

`func (o *UserSearchRecord) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *UserSearchRecord) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *UserSearchRecord) SetName(v string)`

SetName sets Name field to given value.


### GetOrgName

`func (o *UserSearchRecord) GetOrgName() string`

GetOrgName returns the OrgName field if non-nil, zero value otherwise.

### GetOrgNameOk

`func (o *UserSearchRecord) GetOrgNameOk() (*string, bool)`

GetOrgNameOk returns a tuple with the OrgName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOrgName

`func (o *UserSearchRecord) SetOrgName(v string)`

SetOrgName sets OrgName field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


