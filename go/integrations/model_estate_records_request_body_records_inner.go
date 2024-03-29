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

// checks if the EstateRecordsRequestBodyRecordsInner type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &EstateRecordsRequestBodyRecordsInner{}

// EstateRecordsRequestBodyRecordsInner struct for EstateRecordsRequestBodyRecordsInner
type EstateRecordsRequestBodyRecordsInner struct {
	// Internal ID of a resource.
	Id string `json:"id"`
	// Friendly name of a resource (id will be used, if not provided).
	Name *string `json:"name,omitempty"`
	// Type of a resource (resourceType/extension) can contain just type, eg `jira` or with subtypes separated by \"/\" symbol.
	Type string `json:"type"`
	// Metadata contains a JSON object with information about the record. Keys of the object can by any string value, excluding the `tags` key which is reserved for the tags property. If you want to clear the metadata key, you have to provide it directly in an object with a `null` value. Maximum object size is 1MB. 
	Metadata map[string]interface{} `json:"metadata,omitempty"`
	// Tags are key => value representing tags in Klarity. If you want to clear the tag key, you have to provide it directly in an object with a `null` value. Maximum object size is 1MB. 
	Tags *map[string]string `json:"tags,omitempty"`
	Costs *Costs `json:"costs,omitempty"`
	// Determines in which period the resource will be closed. If specified, it has to be in `YYYY-MM` format and be at least current period - it can not be past period. If you need to insert Estate Records in previous period, use insertInPeriod field - when set it will treat previous period as current. If not specified, the record will exist indefinitely. E.g. a resource with `validThrough` set to `2022-05` will be active till May 2022 and start being inactive in June 2022. 
	ValidThrough *string `json:"validThrough,omitempty"`
	AdditionalProperties map[string]interface{}
}

type _EstateRecordsRequestBodyRecordsInner EstateRecordsRequestBodyRecordsInner

// NewEstateRecordsRequestBodyRecordsInner instantiates a new EstateRecordsRequestBodyRecordsInner object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewEstateRecordsRequestBodyRecordsInner(id string, type_ string) *EstateRecordsRequestBodyRecordsInner {
	this := EstateRecordsRequestBodyRecordsInner{}
	this.Id = id
	this.Type = type_
	return &this
}

// NewEstateRecordsRequestBodyRecordsInnerWithDefaults instantiates a new EstateRecordsRequestBodyRecordsInner object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewEstateRecordsRequestBodyRecordsInnerWithDefaults() *EstateRecordsRequestBodyRecordsInner {
	this := EstateRecordsRequestBodyRecordsInner{}
	return &this
}

// GetId returns the Id field value
func (o *EstateRecordsRequestBodyRecordsInner) GetId() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Id
}

// GetIdOk returns a tuple with the Id field value
// and a boolean to check if the value has been set.
func (o *EstateRecordsRequestBodyRecordsInner) GetIdOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Id, true
}

// SetId sets field value
func (o *EstateRecordsRequestBodyRecordsInner) SetId(v string) {
	o.Id = v
}

// GetName returns the Name field value if set, zero value otherwise.
func (o *EstateRecordsRequestBodyRecordsInner) GetName() string {
	if o == nil || IsNil(o.Name) {
		var ret string
		return ret
	}
	return *o.Name
}

// GetNameOk returns a tuple with the Name field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *EstateRecordsRequestBodyRecordsInner) GetNameOk() (*string, bool) {
	if o == nil || IsNil(o.Name) {
		return nil, false
	}
	return o.Name, true
}

// HasName returns a boolean if a field has been set.
func (o *EstateRecordsRequestBodyRecordsInner) HasName() bool {
	if o != nil && !IsNil(o.Name) {
		return true
	}

	return false
}

// SetName gets a reference to the given string and assigns it to the Name field.
func (o *EstateRecordsRequestBodyRecordsInner) SetName(v string) {
	o.Name = &v
}

// GetType returns the Type field value
func (o *EstateRecordsRequestBodyRecordsInner) GetType() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Type
}

// GetTypeOk returns a tuple with the Type field value
// and a boolean to check if the value has been set.
func (o *EstateRecordsRequestBodyRecordsInner) GetTypeOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Type, true
}

// SetType sets field value
func (o *EstateRecordsRequestBodyRecordsInner) SetType(v string) {
	o.Type = v
}

// GetMetadata returns the Metadata field value if set, zero value otherwise.
func (o *EstateRecordsRequestBodyRecordsInner) GetMetadata() map[string]interface{} {
	if o == nil || IsNil(o.Metadata) {
		var ret map[string]interface{}
		return ret
	}
	return o.Metadata
}

// GetMetadataOk returns a tuple with the Metadata field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *EstateRecordsRequestBodyRecordsInner) GetMetadataOk() (map[string]interface{}, bool) {
	if o == nil || IsNil(o.Metadata) {
		return map[string]interface{}{}, false
	}
	return o.Metadata, true
}

// HasMetadata returns a boolean if a field has been set.
func (o *EstateRecordsRequestBodyRecordsInner) HasMetadata() bool {
	if o != nil && !IsNil(o.Metadata) {
		return true
	}

	return false
}

// SetMetadata gets a reference to the given map[string]interface{} and assigns it to the Metadata field.
func (o *EstateRecordsRequestBodyRecordsInner) SetMetadata(v map[string]interface{}) {
	o.Metadata = v
}

// GetTags returns the Tags field value if set, zero value otherwise.
func (o *EstateRecordsRequestBodyRecordsInner) GetTags() map[string]string {
	if o == nil || IsNil(o.Tags) {
		var ret map[string]string
		return ret
	}
	return *o.Tags
}

// GetTagsOk returns a tuple with the Tags field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *EstateRecordsRequestBodyRecordsInner) GetTagsOk() (*map[string]string, bool) {
	if o == nil || IsNil(o.Tags) {
		return nil, false
	}
	return o.Tags, true
}

// HasTags returns a boolean if a field has been set.
func (o *EstateRecordsRequestBodyRecordsInner) HasTags() bool {
	if o != nil && !IsNil(o.Tags) {
		return true
	}

	return false
}

// SetTags gets a reference to the given map[string]string and assigns it to the Tags field.
func (o *EstateRecordsRequestBodyRecordsInner) SetTags(v map[string]string) {
	o.Tags = &v
}

// GetCosts returns the Costs field value if set, zero value otherwise.
func (o *EstateRecordsRequestBodyRecordsInner) GetCosts() Costs {
	if o == nil || IsNil(o.Costs) {
		var ret Costs
		return ret
	}
	return *o.Costs
}

// GetCostsOk returns a tuple with the Costs field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *EstateRecordsRequestBodyRecordsInner) GetCostsOk() (*Costs, bool) {
	if o == nil || IsNil(o.Costs) {
		return nil, false
	}
	return o.Costs, true
}

// HasCosts returns a boolean if a field has been set.
func (o *EstateRecordsRequestBodyRecordsInner) HasCosts() bool {
	if o != nil && !IsNil(o.Costs) {
		return true
	}

	return false
}

// SetCosts gets a reference to the given Costs and assigns it to the Costs field.
func (o *EstateRecordsRequestBodyRecordsInner) SetCosts(v Costs) {
	o.Costs = &v
}

// GetValidThrough returns the ValidThrough field value if set, zero value otherwise.
func (o *EstateRecordsRequestBodyRecordsInner) GetValidThrough() string {
	if o == nil || IsNil(o.ValidThrough) {
		var ret string
		return ret
	}
	return *o.ValidThrough
}

// GetValidThroughOk returns a tuple with the ValidThrough field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *EstateRecordsRequestBodyRecordsInner) GetValidThroughOk() (*string, bool) {
	if o == nil || IsNil(o.ValidThrough) {
		return nil, false
	}
	return o.ValidThrough, true
}

// HasValidThrough returns a boolean if a field has been set.
func (o *EstateRecordsRequestBodyRecordsInner) HasValidThrough() bool {
	if o != nil && !IsNil(o.ValidThrough) {
		return true
	}

	return false
}

// SetValidThrough gets a reference to the given string and assigns it to the ValidThrough field.
func (o *EstateRecordsRequestBodyRecordsInner) SetValidThrough(v string) {
	o.ValidThrough = &v
}

func (o EstateRecordsRequestBodyRecordsInner) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o EstateRecordsRequestBodyRecordsInner) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["id"] = o.Id
	if !IsNil(o.Name) {
		toSerialize["name"] = o.Name
	}
	toSerialize["type"] = o.Type
	if !IsNil(o.Metadata) {
		toSerialize["metadata"] = o.Metadata
	}
	if !IsNil(o.Tags) {
		toSerialize["tags"] = o.Tags
	}
	if !IsNil(o.Costs) {
		toSerialize["costs"] = o.Costs
	}
	if !IsNil(o.ValidThrough) {
		toSerialize["validThrough"] = o.ValidThrough
	}

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}

	return toSerialize, nil
}

func (o *EstateRecordsRequestBodyRecordsInner) UnmarshalJSON(bytes []byte) (err error) {
	varEstateRecordsRequestBodyRecordsInner := _EstateRecordsRequestBodyRecordsInner{}

	if err = json.Unmarshal(bytes, &varEstateRecordsRequestBodyRecordsInner); err == nil {
		*o = EstateRecordsRequestBodyRecordsInner(varEstateRecordsRequestBodyRecordsInner)
	}

	additionalProperties := make(map[string]interface{})

	if err = json.Unmarshal(bytes, &additionalProperties); err == nil {
		delete(additionalProperties, "id")
		delete(additionalProperties, "name")
		delete(additionalProperties, "type")
		delete(additionalProperties, "metadata")
		delete(additionalProperties, "tags")
		delete(additionalProperties, "costs")
		delete(additionalProperties, "validThrough")
		o.AdditionalProperties = additionalProperties
	}

	return err
}

type NullableEstateRecordsRequestBodyRecordsInner struct {
	value *EstateRecordsRequestBodyRecordsInner
	isSet bool
}

func (v NullableEstateRecordsRequestBodyRecordsInner) Get() *EstateRecordsRequestBodyRecordsInner {
	return v.value
}

func (v *NullableEstateRecordsRequestBodyRecordsInner) Set(val *EstateRecordsRequestBodyRecordsInner) {
	v.value = val
	v.isSet = true
}

func (v NullableEstateRecordsRequestBodyRecordsInner) IsSet() bool {
	return v.isSet
}

func (v *NullableEstateRecordsRequestBodyRecordsInner) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableEstateRecordsRequestBodyRecordsInner(val *EstateRecordsRequestBodyRecordsInner) *NullableEstateRecordsRequestBodyRecordsInner {
	return &NullableEstateRecordsRequestBodyRecordsInner{value: val, isSet: true}
}

func (v NullableEstateRecordsRequestBodyRecordsInner) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableEstateRecordsRequestBodyRecordsInner) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


