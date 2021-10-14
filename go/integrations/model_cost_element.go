/*
Klarity Integrations

REST API for managing Estate Records using Klarity Integrations. You can enrich your estate by creating new kinds of estate records or extending existing ones. Before making use of the API, you must first register your External Integration in Klarity, which provides you with the required authentication credentials. Then, you use those credentials to obtain a Token that allows you to make authorized calls to Klarity’s REST API for External Integration.

API version: 0.0.2
Contact: products@nordcloud.com
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package integrations

import (
	"encoding/json"
)

// CostElement struct for CostElement
type CostElement struct {
	// Date in YYYY-MM-DD format for which cost is charged.
	Date string `json:"date"`
	// Value of the cost. A dot (`.`) is used as a decimal separator. Use `null` to clear cost for given date.
	Value *string `json:"value,omitempty"`
	AdditionalProperties map[string]interface{}
}

type _CostElement CostElement

// NewCostElement instantiates a new CostElement object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewCostElement(date string) *CostElement {
	this := CostElement{}
	this.Date = date
	return &this
}

// NewCostElementWithDefaults instantiates a new CostElement object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewCostElementWithDefaults() *CostElement {
	this := CostElement{}
	return &this
}

// GetDate returns the Date field value
func (o *CostElement) GetDate() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Date
}

// GetDateOk returns a tuple with the Date field value
// and a boolean to check if the value has been set.
func (o *CostElement) GetDateOk() (*string, bool) {
	if o == nil  {
		return nil, false
	}
	return &o.Date, true
}

// SetDate sets field value
func (o *CostElement) SetDate(v string) {
	o.Date = v
}

// GetValue returns the Value field value if set, zero value otherwise.
func (o *CostElement) GetValue() string {
	if o == nil || o.Value == nil {
		var ret string
		return ret
	}
	return *o.Value
}

// GetValueOk returns a tuple with the Value field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CostElement) GetValueOk() (*string, bool) {
	if o == nil || o.Value == nil {
		return nil, false
	}
	return o.Value, true
}

// HasValue returns a boolean if a field has been set.
func (o *CostElement) HasValue() bool {
	if o != nil && o.Value != nil {
		return true
	}

	return false
}

// SetValue gets a reference to the given string and assigns it to the Value field.
func (o *CostElement) SetValue(v string) {
	o.Value = &v
}

func (o CostElement) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if true {
		toSerialize["date"] = o.Date
	}
	if o.Value != nil {
		toSerialize["value"] = o.Value
	}

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}

	return json.Marshal(toSerialize)
}

func (o *CostElement) UnmarshalJSON(bytes []byte) (err error) {
	varCostElement := _CostElement{}

	if err = json.Unmarshal(bytes, &varCostElement); err == nil {
		*o = CostElement(varCostElement)
	}

	additionalProperties := make(map[string]interface{})

	if err = json.Unmarshal(bytes, &additionalProperties); err == nil {
		delete(additionalProperties, "date")
		delete(additionalProperties, "value")
		o.AdditionalProperties = additionalProperties
	}

	return err
}

type NullableCostElement struct {
	value *CostElement
	isSet bool
}

func (v NullableCostElement) Get() *CostElement {
	return v.value
}

func (v *NullableCostElement) Set(val *CostElement) {
	v.value = val
	v.isSet = true
}

func (v NullableCostElement) IsSet() bool {
	return v.isSet
}

func (v *NullableCostElement) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableCostElement(val *CostElement) *NullableCostElement {
	return &NullableCostElement{value: val, isSet: true}
}

func (v NullableCostElement) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableCostElement) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}

