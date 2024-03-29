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

// checks if the EnrichmentsEstateRecordsRequestBodyEnrichmentsInner type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &EnrichmentsEstateRecordsRequestBodyEnrichmentsInner{}

// EnrichmentsEstateRecordsRequestBodyEnrichmentsInner struct for EnrichmentsEstateRecordsRequestBodyEnrichmentsInner
type EnrichmentsEstateRecordsRequestBodyEnrichmentsInner struct {
	Record EnrichmentRecord `json:"record"`
	// Data contains an arbitrary JSON object with enrichment of the record. Can not be empty object. Maximum object size is 1MB. 
	Data map[string]interface{} `json:"data"`
	AdditionalProperties map[string]interface{}
}

type _EnrichmentsEstateRecordsRequestBodyEnrichmentsInner EnrichmentsEstateRecordsRequestBodyEnrichmentsInner

// NewEnrichmentsEstateRecordsRequestBodyEnrichmentsInner instantiates a new EnrichmentsEstateRecordsRequestBodyEnrichmentsInner object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewEnrichmentsEstateRecordsRequestBodyEnrichmentsInner(record EnrichmentRecord, data map[string]interface{}) *EnrichmentsEstateRecordsRequestBodyEnrichmentsInner {
	this := EnrichmentsEstateRecordsRequestBodyEnrichmentsInner{}
	this.Record = record
	this.Data = data
	return &this
}

// NewEnrichmentsEstateRecordsRequestBodyEnrichmentsInnerWithDefaults instantiates a new EnrichmentsEstateRecordsRequestBodyEnrichmentsInner object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewEnrichmentsEstateRecordsRequestBodyEnrichmentsInnerWithDefaults() *EnrichmentsEstateRecordsRequestBodyEnrichmentsInner {
	this := EnrichmentsEstateRecordsRequestBodyEnrichmentsInner{}
	return &this
}

// GetRecord returns the Record field value
func (o *EnrichmentsEstateRecordsRequestBodyEnrichmentsInner) GetRecord() EnrichmentRecord {
	if o == nil {
		var ret EnrichmentRecord
		return ret
	}

	return o.Record
}

// GetRecordOk returns a tuple with the Record field value
// and a boolean to check if the value has been set.
func (o *EnrichmentsEstateRecordsRequestBodyEnrichmentsInner) GetRecordOk() (*EnrichmentRecord, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Record, true
}

// SetRecord sets field value
func (o *EnrichmentsEstateRecordsRequestBodyEnrichmentsInner) SetRecord(v EnrichmentRecord) {
	o.Record = v
}

// GetData returns the Data field value
func (o *EnrichmentsEstateRecordsRequestBodyEnrichmentsInner) GetData() map[string]interface{} {
	if o == nil {
		var ret map[string]interface{}
		return ret
	}

	return o.Data
}

// GetDataOk returns a tuple with the Data field value
// and a boolean to check if the value has been set.
func (o *EnrichmentsEstateRecordsRequestBodyEnrichmentsInner) GetDataOk() (map[string]interface{}, bool) {
	if o == nil {
		return map[string]interface{}{}, false
	}
	return o.Data, true
}

// SetData sets field value
func (o *EnrichmentsEstateRecordsRequestBodyEnrichmentsInner) SetData(v map[string]interface{}) {
	o.Data = v
}

func (o EnrichmentsEstateRecordsRequestBodyEnrichmentsInner) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o EnrichmentsEstateRecordsRequestBodyEnrichmentsInner) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["record"] = o.Record
	toSerialize["data"] = o.Data

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}

	return toSerialize, nil
}

func (o *EnrichmentsEstateRecordsRequestBodyEnrichmentsInner) UnmarshalJSON(bytes []byte) (err error) {
	varEnrichmentsEstateRecordsRequestBodyEnrichmentsInner := _EnrichmentsEstateRecordsRequestBodyEnrichmentsInner{}

	if err = json.Unmarshal(bytes, &varEnrichmentsEstateRecordsRequestBodyEnrichmentsInner); err == nil {
		*o = EnrichmentsEstateRecordsRequestBodyEnrichmentsInner(varEnrichmentsEstateRecordsRequestBodyEnrichmentsInner)
	}

	additionalProperties := make(map[string]interface{})

	if err = json.Unmarshal(bytes, &additionalProperties); err == nil {
		delete(additionalProperties, "record")
		delete(additionalProperties, "data")
		o.AdditionalProperties = additionalProperties
	}

	return err
}

type NullableEnrichmentsEstateRecordsRequestBodyEnrichmentsInner struct {
	value *EnrichmentsEstateRecordsRequestBodyEnrichmentsInner
	isSet bool
}

func (v NullableEnrichmentsEstateRecordsRequestBodyEnrichmentsInner) Get() *EnrichmentsEstateRecordsRequestBodyEnrichmentsInner {
	return v.value
}

func (v *NullableEnrichmentsEstateRecordsRequestBodyEnrichmentsInner) Set(val *EnrichmentsEstateRecordsRequestBodyEnrichmentsInner) {
	v.value = val
	v.isSet = true
}

func (v NullableEnrichmentsEstateRecordsRequestBodyEnrichmentsInner) IsSet() bool {
	return v.isSet
}

func (v *NullableEnrichmentsEstateRecordsRequestBodyEnrichmentsInner) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableEnrichmentsEstateRecordsRequestBodyEnrichmentsInner(val *EnrichmentsEstateRecordsRequestBodyEnrichmentsInner) *NullableEnrichmentsEstateRecordsRequestBodyEnrichmentsInner {
	return &NullableEnrichmentsEstateRecordsRequestBodyEnrichmentsInner{value: val, isSet: true}
}

func (v NullableEnrichmentsEstateRecordsRequestBodyEnrichmentsInner) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableEnrichmentsEstateRecordsRequestBodyEnrichmentsInner) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


