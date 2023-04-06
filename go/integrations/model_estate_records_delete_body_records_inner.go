/*
Klarity Integrations

REST API for managing Estate Records using Klarity Integrations. You can enrich your estate by creating new kinds of estate records or extending existing ones. Before making use of the API, you must first register your External Integration in Klarity, which provides you with the required authentication credentials. Then, you use those credentials to obtain a Token that allows you to make authorized calls to Klarity’s REST API for External Integration.

API version: 0.0.5
Contact: products@nordcloud.com
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package integrations

import (
	"encoding/json"
)

// checks if the EstateRecordsDeleteBodyRecordsInner type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &EstateRecordsDeleteBodyRecordsInner{}

// EstateRecordsDeleteBodyRecordsInner struct for EstateRecordsDeleteBodyRecordsInner
type EstateRecordsDeleteBodyRecordsInner struct {
	// Nordcloud ID of a resource in Klarity.
	Nid *string `json:"nid,omitempty"`
	// Internal ID of a resource.
	Id *string `json:"id,omitempty"`
	// Friendly name of a resource (id will be used, if not provided).
	Name *string `json:"name,omitempty"`
	// Type of a resource (resourceType/extension) can contain just type, eg `jira` or with subtypes separated by \"/\" symbol.
	Type *string `json:"type,omitempty"`
	AdditionalProperties map[string]interface{}
}

type _EstateRecordsDeleteBodyRecordsInner EstateRecordsDeleteBodyRecordsInner

// NewEstateRecordsDeleteBodyRecordsInner instantiates a new EstateRecordsDeleteBodyRecordsInner object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewEstateRecordsDeleteBodyRecordsInner() *EstateRecordsDeleteBodyRecordsInner {
	this := EstateRecordsDeleteBodyRecordsInner{}
	return &this
}

// NewEstateRecordsDeleteBodyRecordsInnerWithDefaults instantiates a new EstateRecordsDeleteBodyRecordsInner object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewEstateRecordsDeleteBodyRecordsInnerWithDefaults() *EstateRecordsDeleteBodyRecordsInner {
	this := EstateRecordsDeleteBodyRecordsInner{}
	return &this
}

// GetNid returns the Nid field value if set, zero value otherwise.
func (o *EstateRecordsDeleteBodyRecordsInner) GetNid() string {
	if o == nil || IsNil(o.Nid) {
		var ret string
		return ret
	}
	return *o.Nid
}

// GetNidOk returns a tuple with the Nid field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *EstateRecordsDeleteBodyRecordsInner) GetNidOk() (*string, bool) {
	if o == nil || IsNil(o.Nid) {
		return nil, false
	}
	return o.Nid, true
}

// HasNid returns a boolean if a field has been set.
func (o *EstateRecordsDeleteBodyRecordsInner) HasNid() bool {
	if o != nil && !IsNil(o.Nid) {
		return true
	}

	return false
}

// SetNid gets a reference to the given string and assigns it to the Nid field.
func (o *EstateRecordsDeleteBodyRecordsInner) SetNid(v string) {
	o.Nid = &v
}

// GetId returns the Id field value if set, zero value otherwise.
func (o *EstateRecordsDeleteBodyRecordsInner) GetId() string {
	if o == nil || IsNil(o.Id) {
		var ret string
		return ret
	}
	return *o.Id
}

// GetIdOk returns a tuple with the Id field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *EstateRecordsDeleteBodyRecordsInner) GetIdOk() (*string, bool) {
	if o == nil || IsNil(o.Id) {
		return nil, false
	}
	return o.Id, true
}

// HasId returns a boolean if a field has been set.
func (o *EstateRecordsDeleteBodyRecordsInner) HasId() bool {
	if o != nil && !IsNil(o.Id) {
		return true
	}

	return false
}

// SetId gets a reference to the given string and assigns it to the Id field.
func (o *EstateRecordsDeleteBodyRecordsInner) SetId(v string) {
	o.Id = &v
}

// GetName returns the Name field value if set, zero value otherwise.
func (o *EstateRecordsDeleteBodyRecordsInner) GetName() string {
	if o == nil || IsNil(o.Name) {
		var ret string
		return ret
	}
	return *o.Name
}

// GetNameOk returns a tuple with the Name field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *EstateRecordsDeleteBodyRecordsInner) GetNameOk() (*string, bool) {
	if o == nil || IsNil(o.Name) {
		return nil, false
	}
	return o.Name, true
}

// HasName returns a boolean if a field has been set.
func (o *EstateRecordsDeleteBodyRecordsInner) HasName() bool {
	if o != nil && !IsNil(o.Name) {
		return true
	}

	return false
}

// SetName gets a reference to the given string and assigns it to the Name field.
func (o *EstateRecordsDeleteBodyRecordsInner) SetName(v string) {
	o.Name = &v
}

// GetType returns the Type field value if set, zero value otherwise.
func (o *EstateRecordsDeleteBodyRecordsInner) GetType() string {
	if o == nil || IsNil(o.Type) {
		var ret string
		return ret
	}
	return *o.Type
}

// GetTypeOk returns a tuple with the Type field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *EstateRecordsDeleteBodyRecordsInner) GetTypeOk() (*string, bool) {
	if o == nil || IsNil(o.Type) {
		return nil, false
	}
	return o.Type, true
}

// HasType returns a boolean if a field has been set.
func (o *EstateRecordsDeleteBodyRecordsInner) HasType() bool {
	if o != nil && !IsNil(o.Type) {
		return true
	}

	return false
}

// SetType gets a reference to the given string and assigns it to the Type field.
func (o *EstateRecordsDeleteBodyRecordsInner) SetType(v string) {
	o.Type = &v
}

func (o EstateRecordsDeleteBodyRecordsInner) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o EstateRecordsDeleteBodyRecordsInner) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.Nid) {
		toSerialize["nid"] = o.Nid
	}
	if !IsNil(o.Id) {
		toSerialize["id"] = o.Id
	}
	if !IsNil(o.Name) {
		toSerialize["name"] = o.Name
	}
	if !IsNil(o.Type) {
		toSerialize["type"] = o.Type
	}

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}

	return toSerialize, nil
}

func (o *EstateRecordsDeleteBodyRecordsInner) UnmarshalJSON(bytes []byte) (err error) {
	varEstateRecordsDeleteBodyRecordsInner := _EstateRecordsDeleteBodyRecordsInner{}

	if err = json.Unmarshal(bytes, &varEstateRecordsDeleteBodyRecordsInner); err == nil {
		*o = EstateRecordsDeleteBodyRecordsInner(varEstateRecordsDeleteBodyRecordsInner)
	}

	additionalProperties := make(map[string]interface{})

	if err = json.Unmarshal(bytes, &additionalProperties); err == nil {
		delete(additionalProperties, "nid")
		delete(additionalProperties, "id")
		delete(additionalProperties, "name")
		delete(additionalProperties, "type")
		o.AdditionalProperties = additionalProperties
	}

	return err
}

type NullableEstateRecordsDeleteBodyRecordsInner struct {
	value *EstateRecordsDeleteBodyRecordsInner
	isSet bool
}

func (v NullableEstateRecordsDeleteBodyRecordsInner) Get() *EstateRecordsDeleteBodyRecordsInner {
	return v.value
}

func (v *NullableEstateRecordsDeleteBodyRecordsInner) Set(val *EstateRecordsDeleteBodyRecordsInner) {
	v.value = val
	v.isSet = true
}

func (v NullableEstateRecordsDeleteBodyRecordsInner) IsSet() bool {
	return v.isSet
}

func (v *NullableEstateRecordsDeleteBodyRecordsInner) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableEstateRecordsDeleteBodyRecordsInner(val *EstateRecordsDeleteBodyRecordsInner) *NullableEstateRecordsDeleteBodyRecordsInner {
	return &NullableEstateRecordsDeleteBodyRecordsInner{value: val, isSet: true}
}

func (v NullableEstateRecordsDeleteBodyRecordsInner) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableEstateRecordsDeleteBodyRecordsInner) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}

