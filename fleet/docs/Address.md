# Address

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**AddressLine1** | **string** | The street address | 
**AddressLine2** | Pointer to **string** | Address Line 2 | [optional] 
**AddressLine3** | Pointer to **string** | Address Line 3 | [optional] 
**City** | **string** | Name of the City | 
**Country** | **string** | Name of the Country | 
**State** | **string** | Name of the State/Province/Region | 
**Zip** | **string** | Zip code | 

## Methods

### NewAddress

`func NewAddress(addressLine1 string, city string, country string, state string, zip string, ) *Address`

NewAddress instantiates a new Address object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewAddressWithDefaults

`func NewAddressWithDefaults() *Address`

NewAddressWithDefaults instantiates a new Address object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetAddressLine1

`func (o *Address) GetAddressLine1() string`

GetAddressLine1 returns the AddressLine1 field if non-nil, zero value otherwise.

### GetAddressLine1Ok

`func (o *Address) GetAddressLine1Ok() (*string, bool)`

GetAddressLine1Ok returns a tuple with the AddressLine1 field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAddressLine1

`func (o *Address) SetAddressLine1(v string)`

SetAddressLine1 sets AddressLine1 field to given value.


### GetAddressLine2

`func (o *Address) GetAddressLine2() string`

GetAddressLine2 returns the AddressLine2 field if non-nil, zero value otherwise.

### GetAddressLine2Ok

`func (o *Address) GetAddressLine2Ok() (*string, bool)`

GetAddressLine2Ok returns a tuple with the AddressLine2 field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAddressLine2

`func (o *Address) SetAddressLine2(v string)`

SetAddressLine2 sets AddressLine2 field to given value.

### HasAddressLine2

`func (o *Address) HasAddressLine2() bool`

HasAddressLine2 returns a boolean if a field has been set.

### GetAddressLine3

`func (o *Address) GetAddressLine3() string`

GetAddressLine3 returns the AddressLine3 field if non-nil, zero value otherwise.

### GetAddressLine3Ok

`func (o *Address) GetAddressLine3Ok() (*string, bool)`

GetAddressLine3Ok returns a tuple with the AddressLine3 field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAddressLine3

`func (o *Address) SetAddressLine3(v string)`

SetAddressLine3 sets AddressLine3 field to given value.

### HasAddressLine3

`func (o *Address) HasAddressLine3() bool`

HasAddressLine3 returns a boolean if a field has been set.

### GetCity

`func (o *Address) GetCity() string`

GetCity returns the City field if non-nil, zero value otherwise.

### GetCityOk

`func (o *Address) GetCityOk() (*string, bool)`

GetCityOk returns a tuple with the City field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCity

`func (o *Address) SetCity(v string)`

SetCity sets City field to given value.


### GetCountry

`func (o *Address) GetCountry() string`

GetCountry returns the Country field if non-nil, zero value otherwise.

### GetCountryOk

`func (o *Address) GetCountryOk() (*string, bool)`

GetCountryOk returns a tuple with the Country field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCountry

`func (o *Address) SetCountry(v string)`

SetCountry sets Country field to given value.


### GetState

`func (o *Address) GetState() string`

GetState returns the State field if non-nil, zero value otherwise.

### GetStateOk

`func (o *Address) GetStateOk() (*string, bool)`

GetStateOk returns a tuple with the State field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetState

`func (o *Address) SetState(v string)`

SetState sets State field to given value.


### GetZip

`func (o *Address) GetZip() string`

GetZip returns the Zip field if non-nil, zero value otherwise.

### GetZipOk

`func (o *Address) GetZipOk() (*string, bool)`

GetZipOk returns a tuple with the Zip field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetZip

`func (o *Address) SetZip(v string)`

SetZip sets Zip field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


