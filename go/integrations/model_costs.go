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

// Costs Costs object contains an array of costs per day and a currency. Only if there is a date passed into an array, then the cost is being added/updated for this date. If you want to clear cost for a given date, you have to `null` directly into an array. 
type Costs struct {
	// Currency in ISO-4217 format.
	Currency string `json:"currency"`
	Values []CostElement `json:"values"`
	AdditionalProperties map[string]interface{}
}

type _Costs Costs

// NewCosts instantiates a new Costs object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewCosts(currency string, values []CostElement) *Costs {
	this := Costs{}
	this.Currency = currency
	this.Values = values
	return &this
}

// NewCostsWithDefaults instantiates a new Costs object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewCostsWithDefaults() *Costs {
	this := Costs{}
	return &this
}

// GetCurrency returns the Currency field value
func (o *Costs) GetCurrency() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Currency
}

// GetCurrencyOk returns a tuple with the Currency field value
// and a boolean to check if the value has been set.
func (o *Costs) GetCurrencyOk() (*string, bool) {
	if o == nil  {
		return nil, false
	}
	return &o.Currency, true
}

// SetCurrency sets field value
func (o *Costs) SetCurrency(v string) {
	o.Currency = v
}

// GetValues returns the Values field value
func (o *Costs) GetValues() []CostElement {
	if o == nil {
		var ret []CostElement
		return ret
	}

	return o.Values
}

// GetValuesOk returns a tuple with the Values field value
// and a boolean to check if the value has been set.
func (o *Costs) GetValuesOk() (*[]CostElement, bool) {
	if o == nil  {
		return nil, false
	}
	return &o.Values, true
}

// SetValues sets field value
func (o *Costs) SetValues(v []CostElement) {
	o.Values = v
}

func (o Costs) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if true {
		toSerialize["currency"] = o.Currency
	}
	if true {
		toSerialize["values"] = o.Values
	}

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}

	return json.Marshal(toSerialize)
}

func (o *Costs) UnmarshalJSON(bytes []byte) (err error) {
	varCosts := _Costs{}

	if err = json.Unmarshal(bytes, &varCosts); err == nil {
		*o = Costs(varCosts)
	}

	additionalProperties := make(map[string]interface{})

	if err = json.Unmarshal(bytes, &additionalProperties); err == nil {
		delete(additionalProperties, "currency")
		delete(additionalProperties, "values")
		o.AdditionalProperties = additionalProperties
	}

	return err
}

type NullableCosts struct {
	value *Costs
	isSet bool
}

func (v NullableCosts) Get() *Costs {
	return v.value
}

func (v *NullableCosts) Set(val *Costs) {
	v.value = val
	v.isSet = true
}

func (v NullableCosts) IsSet() bool {
	return v.isSet
}

func (v *NullableCosts) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableCosts(val *Costs) *NullableCosts {
	return &NullableCosts{value: val, isSet: true}
}

func (v NullableCosts) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableCosts) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


