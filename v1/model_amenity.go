/*
Omnistrate Registration API

REST API for Omnistrate Service Registration

API version: 2022-09-01-00
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package v1

import (
	"encoding/json"
)

// checks if the Amenity type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &Amenity{}

// Amenity struct for Amenity
type Amenity struct {
	// A description of the amenity.
	Description *string `json:"Description,omitempty"`
	// Whether the amenity is managed by the system.
	IsManaged *bool `json:"IsManaged,omitempty"`
	// Whether the amenity can be modified.
	Modifiable *bool `json:"Modifiable,omitempty"`
	// The name of the amenity.
	Name *string `json:"Name,omitempty"`
	// The properties of the amenity.
	Properties map[string]interface{} `json:"Properties,omitempty"`
	// The type of the amenity.
	Type *string `json:"Type,omitempty"`
	AdditionalProperties map[string]interface{}
}

type _Amenity Amenity

// NewAmenity instantiates a new Amenity object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewAmenity() *Amenity {
	this := Amenity{}
	return &this
}

// NewAmenityWithDefaults instantiates a new Amenity object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewAmenityWithDefaults() *Amenity {
	this := Amenity{}
	return &this
}

// GetDescription returns the Description field value if set, zero value otherwise.
func (o *Amenity) GetDescription() string {
	if o == nil || IsNil(o.Description) {
		var ret string
		return ret
	}
	return *o.Description
}

// GetDescriptionOk returns a tuple with the Description field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Amenity) GetDescriptionOk() (*string, bool) {
	if o == nil || IsNil(o.Description) {
		return nil, false
	}
	return o.Description, true
}

// SetDescription gets a reference to the given string and assigns it to the Description field.
func (o *Amenity) SetDescription(v string) {
	o.Description = &v
}

// GetIsManaged returns the IsManaged field value if set, zero value otherwise.
func (o *Amenity) GetIsManaged() bool {
	if o == nil || IsNil(o.IsManaged) {
		var ret bool
		return ret
	}
	return *o.IsManaged
}

// GetIsManagedOk returns a tuple with the IsManaged field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Amenity) GetIsManagedOk() (*bool, bool) {
	if o == nil || IsNil(o.IsManaged) {
		return nil, false
	}
	return o.IsManaged, true
}

// SetIsManaged gets a reference to the given bool and assigns it to the IsManaged field.
func (o *Amenity) SetIsManaged(v bool) {
	o.IsManaged = &v
}

// GetModifiable returns the Modifiable field value if set, zero value otherwise.
func (o *Amenity) GetModifiable() bool {
	if o == nil || IsNil(o.Modifiable) {
		var ret bool
		return ret
	}
	return *o.Modifiable
}

// GetModifiableOk returns a tuple with the Modifiable field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Amenity) GetModifiableOk() (*bool, bool) {
	if o == nil || IsNil(o.Modifiable) {
		return nil, false
	}
	return o.Modifiable, true
}

// SetModifiable gets a reference to the given bool and assigns it to the Modifiable field.
func (o *Amenity) SetModifiable(v bool) {
	o.Modifiable = &v
}

// GetName returns the Name field value if set, zero value otherwise.
func (o *Amenity) GetName() string {
	if o == nil || IsNil(o.Name) {
		var ret string
		return ret
	}
	return *o.Name
}

// GetNameOk returns a tuple with the Name field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Amenity) GetNameOk() (*string, bool) {
	if o == nil || IsNil(o.Name) {
		return nil, false
	}
	return o.Name, true
}

// SetName gets a reference to the given string and assigns it to the Name field.
func (o *Amenity) SetName(v string) {
	o.Name = &v
}

// GetProperties returns the Properties field value if set, zero value otherwise.
func (o *Amenity) GetProperties() map[string]interface{} {
	if o == nil || IsNil(o.Properties) {
		var ret map[string]interface{}
		return ret
	}
	return o.Properties
}

// GetPropertiesOk returns a tuple with the Properties field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Amenity) GetPropertiesOk() (map[string]interface{}, bool) {
	if o == nil || IsNil(o.Properties) {
		return map[string]interface{}{}, false
	}
	return o.Properties, true
}

// SetProperties gets a reference to the given map[string]interface{} and assigns it to the Properties field.
func (o *Amenity) SetProperties(v map[string]interface{}) {
	o.Properties = v
}

// GetType returns the Type field value if set, zero value otherwise.
func (o *Amenity) GetType() string {
	if o == nil || IsNil(o.Type) {
		var ret string
		return ret
	}
	return *o.Type
}

// GetTypeOk returns a tuple with the Type field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Amenity) GetTypeOk() (*string, bool) {
	if o == nil || IsNil(o.Type) {
		return nil, false
	}
	return o.Type, true
}

// SetType gets a reference to the given string and assigns it to the Type field.
func (o *Amenity) SetType(v string) {
	o.Type = &v
}

func (o Amenity) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o Amenity) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.Description) {
		toSerialize["Description"] = o.Description
	}
	if !IsNil(o.IsManaged) {
		toSerialize["IsManaged"] = o.IsManaged
	}
	if !IsNil(o.Modifiable) {
		toSerialize["Modifiable"] = o.Modifiable
	}
	if !IsNil(o.Name) {
		toSerialize["Name"] = o.Name
	}
	if !IsNil(o.Properties) {
		toSerialize["Properties"] = o.Properties
	}
	if !IsNil(o.Type) {
		toSerialize["Type"] = o.Type
	}

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}

	return toSerialize, nil
}

func (o *Amenity) UnmarshalJSON(data []byte) (err error) {
	varAmenity := _Amenity{}

	err = json.Unmarshal(data, &varAmenity)

	if err != nil {
		return err
	}

	*o = Amenity(varAmenity)

	additionalProperties := make(map[string]interface{})

	if err = json.Unmarshal(data, &additionalProperties); err == nil {
		delete(additionalProperties, "Description")
		delete(additionalProperties, "IsManaged")
		delete(additionalProperties, "Modifiable")
		delete(additionalProperties, "Name")
		delete(additionalProperties, "Properties")
		delete(additionalProperties, "Type")
		o.AdditionalProperties = additionalProperties
	}

	return err
}

type NullableAmenity struct {
	value *Amenity
	isSet bool
}

func (v NullableAmenity) Get() *Amenity {
	return v.value
}

func (v *NullableAmenity) Set(val *Amenity) {
	v.value = val
	v.isSet = true
}

func (v NullableAmenity) IsSet() bool {
	return v.isSet
}

func (v *NullableAmenity) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableAmenity(val *Amenity) *NullableAmenity {
	return &NullableAmenity{value: val, isSet: true}
}

func (v NullableAmenity) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableAmenity) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


