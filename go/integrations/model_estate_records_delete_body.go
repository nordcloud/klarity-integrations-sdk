/*
Klarity Integrations

REST API for managing Estate Records using Klarity Integrations. You can enrich your estate by creating new kinds of estate records or extending existing ones. Before making use of the API, you must first register your External Integration in Klarity, which provides you with the required authentication credentials. Then, you use those credentials to obtain a Token that allows you to make authorized calls to Klarity’s REST API for External Integration.

API version: 0.0.4
Contact: products@nordcloud.com
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package integrations

import (
	"encoding/json"
)

// EstateRecordsDeleteBody struct for EstateRecordsDeleteBody
type EstateRecordsDeleteBody struct {
	Period PeriodEnum `json:"period"`
	Records []EstateRecordsDeleteBodyRecords `json:"records"`
	AdditionalProperties map[string]interface{}
}

type _EstateRecordsDeleteBody EstateRecordsDeleteBody

// NewEstateRecordsDeleteBody instantiates a new EstateRecordsDeleteBody object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewEstateRecordsDeleteBody(period PeriodEnum, records []EstateRecordsDeleteBodyRecords) *EstateRecordsDeleteBody {
	this := EstateRecordsDeleteBody{}
	this.Period = period
	this.Records = records
	return &this
}

// NewEstateRecordsDeleteBodyWithDefaults instantiates a new EstateRecordsDeleteBody object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewEstateRecordsDeleteBodyWithDefaults() *EstateRecordsDeleteBody {
	this := EstateRecordsDeleteBody{}
	return &this
}

// GetPeriod returns the Period field value
func (o *EstateRecordsDeleteBody) GetPeriod() PeriodEnum {
	if o == nil {
		var ret PeriodEnum
		return ret
	}

	return o.Period
}

// GetPeriodOk returns a tuple with the Period field value
// and a boolean to check if the value has been set.
func (o *EstateRecordsDeleteBody) GetPeriodOk() (*PeriodEnum, bool) {
	if o == nil  {
		return nil, false
	}
	return &o.Period, true
}

// SetPeriod sets field value
func (o *EstateRecordsDeleteBody) SetPeriod(v PeriodEnum) {
	o.Period = v
}

// GetRecords returns the Records field value
func (o *EstateRecordsDeleteBody) GetRecords() []EstateRecordsDeleteBodyRecords {
	if o == nil {
		var ret []EstateRecordsDeleteBodyRecords
		return ret
	}

	return o.Records
}

// GetRecordsOk returns a tuple with the Records field value
// and a boolean to check if the value has been set.
func (o *EstateRecordsDeleteBody) GetRecordsOk() (*[]EstateRecordsDeleteBodyRecords, bool) {
	if o == nil  {
		return nil, false
	}
	return &o.Records, true
}

// SetRecords sets field value
func (o *EstateRecordsDeleteBody) SetRecords(v []EstateRecordsDeleteBodyRecords) {
	o.Records = v
}

func (o EstateRecordsDeleteBody) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if true {
		toSerialize["period"] = o.Period
	}
	if true {
		toSerialize["records"] = o.Records
	}

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}

	return json.Marshal(toSerialize)
}

func (o *EstateRecordsDeleteBody) UnmarshalJSON(bytes []byte) (err error) {
	varEstateRecordsDeleteBody := _EstateRecordsDeleteBody{}

	if err = json.Unmarshal(bytes, &varEstateRecordsDeleteBody); err == nil {
		*o = EstateRecordsDeleteBody(varEstateRecordsDeleteBody)
	}

	additionalProperties := make(map[string]interface{})

	if err = json.Unmarshal(bytes, &additionalProperties); err == nil {
		delete(additionalProperties, "period")
		delete(additionalProperties, "records")
		o.AdditionalProperties = additionalProperties
	}

	return err
}

type NullableEstateRecordsDeleteBody struct {
	value *EstateRecordsDeleteBody
	isSet bool
}

func (v NullableEstateRecordsDeleteBody) Get() *EstateRecordsDeleteBody {
	return v.value
}

func (v *NullableEstateRecordsDeleteBody) Set(val *EstateRecordsDeleteBody) {
	v.value = val
	v.isSet = true
}

func (v NullableEstateRecordsDeleteBody) IsSet() bool {
	return v.isSet
}

func (v *NullableEstateRecordsDeleteBody) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableEstateRecordsDeleteBody(val *EstateRecordsDeleteBody) *NullableEstateRecordsDeleteBody {
	return &NullableEstateRecordsDeleteBody{value: val, isSet: true}
}

func (v NullableEstateRecordsDeleteBody) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableEstateRecordsDeleteBody) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


