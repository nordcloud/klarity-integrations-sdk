/*
Klarity Integrations

REST API for managing Estate Records using Klarity Integrations. You can enrich your estate by creating new kinds of estate records or extending existing ones. Before making use of the API, you must first register your External Integration in Klarity, which provides you with the required authentication credentials. Then, you use those credentials to obtain a Token that allows you to make authorized calls to Klarity’s REST API for External Integration.

API version: 0.0.3
Contact: products@nordcloud.com
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package integrations

import (
	"encoding/json"
)

// EstateRecordsRequestBody struct for EstateRecordsRequestBody
type EstateRecordsRequestBody struct {
	Records []EstateRecordsRequestBodyRecords `json:"records"`
	AdditionalProperties map[string]interface{}
}

type _EstateRecordsRequestBody EstateRecordsRequestBody

// NewEstateRecordsRequestBody instantiates a new EstateRecordsRequestBody object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewEstateRecordsRequestBody(records []EstateRecordsRequestBodyRecords) *EstateRecordsRequestBody {
	this := EstateRecordsRequestBody{}
	this.Records = records
	return &this
}

// NewEstateRecordsRequestBodyWithDefaults instantiates a new EstateRecordsRequestBody object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewEstateRecordsRequestBodyWithDefaults() *EstateRecordsRequestBody {
	this := EstateRecordsRequestBody{}
	return &this
}

// GetRecords returns the Records field value
func (o *EstateRecordsRequestBody) GetRecords() []EstateRecordsRequestBodyRecords {
	if o == nil {
		var ret []EstateRecordsRequestBodyRecords
		return ret
	}

	return o.Records
}

// GetRecordsOk returns a tuple with the Records field value
// and a boolean to check if the value has been set.
func (o *EstateRecordsRequestBody) GetRecordsOk() (*[]EstateRecordsRequestBodyRecords, bool) {
	if o == nil  {
		return nil, false
	}
	return &o.Records, true
}

// SetRecords sets field value
func (o *EstateRecordsRequestBody) SetRecords(v []EstateRecordsRequestBodyRecords) {
	o.Records = v
}

func (o EstateRecordsRequestBody) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if true {
		toSerialize["records"] = o.Records
	}

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}

	return json.Marshal(toSerialize)
}

func (o *EstateRecordsRequestBody) UnmarshalJSON(bytes []byte) (err error) {
	varEstateRecordsRequestBody := _EstateRecordsRequestBody{}

	if err = json.Unmarshal(bytes, &varEstateRecordsRequestBody); err == nil {
		*o = EstateRecordsRequestBody(varEstateRecordsRequestBody)
	}

	additionalProperties := make(map[string]interface{})

	if err = json.Unmarshal(bytes, &additionalProperties); err == nil {
		delete(additionalProperties, "records")
		o.AdditionalProperties = additionalProperties
	}

	return err
}

type NullableEstateRecordsRequestBody struct {
	value *EstateRecordsRequestBody
	isSet bool
}

func (v NullableEstateRecordsRequestBody) Get() *EstateRecordsRequestBody {
	return v.value
}

func (v *NullableEstateRecordsRequestBody) Set(val *EstateRecordsRequestBody) {
	v.value = val
	v.isSet = true
}

func (v NullableEstateRecordsRequestBody) IsSet() bool {
	return v.isSet
}

func (v *NullableEstateRecordsRequestBody) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableEstateRecordsRequestBody(val *EstateRecordsRequestBody) *NullableEstateRecordsRequestBody {
	return &NullableEstateRecordsRequestBody{value: val, isSet: true}
}

func (v NullableEstateRecordsRequestBody) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableEstateRecordsRequestBody) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


