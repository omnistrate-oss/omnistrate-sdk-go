# DescribeUsersByOrgResult

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Id** | **string** | The ID of the org | 
**OrgUsers** | [**[]OrgUsers**](OrgUsers.md) | The users associated with the corresponding org | 

## Methods

### NewDescribeUsersByOrgResult

`func NewDescribeUsersByOrgResult(id string, orgUsers []OrgUsers, ) *DescribeUsersByOrgResult`

NewDescribeUsersByOrgResult instantiates a new DescribeUsersByOrgResult object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewDescribeUsersByOrgResultWithDefaults

`func NewDescribeUsersByOrgResultWithDefaults() *DescribeUsersByOrgResult`

NewDescribeUsersByOrgResultWithDefaults instantiates a new DescribeUsersByOrgResult object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetId

`func (o *DescribeUsersByOrgResult) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *DescribeUsersByOrgResult) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *DescribeUsersByOrgResult) SetId(v string)`

SetId sets Id field to given value.


### GetOrgUsers

`func (o *DescribeUsersByOrgResult) GetOrgUsers() []OrgUsers`

GetOrgUsers returns the OrgUsers field if non-nil, zero value otherwise.

### GetOrgUsersOk

`func (o *DescribeUsersByOrgResult) GetOrgUsersOk() (*[]OrgUsers, bool)`

GetOrgUsersOk returns a tuple with the OrgUsers field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOrgUsers

`func (o *DescribeUsersByOrgResult) SetOrgUsers(v []OrgUsers)`

SetOrgUsers sets OrgUsers field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


