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

// AcceptedResponseBody Response returned when processing of a request ends successfully
type AcceptedResponseBody struct {
	Message *string `json:"message,omitempty"`
	AdditionalProperties map[string]interface{}
}

type _AcceptedResponseBody AcceptedResponseBody

// NewAcceptedResponseBody instantiates a new AcceptedResponseBody object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewAcceptedResponseBody() *AcceptedResponseBody {
	this := AcceptedResponseBody{}
	return &this
}

// NewAcceptedResponseBodyWithDefaults instantiates a new AcceptedResponseBody object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewAcceptedResponseBodyWithDefaults() *AcceptedResponseBody {
	this := AcceptedResponseBody{}
	return &this
}

// GetMessage returns the Message field value if set, zero value otherwise.
func (o *AcceptedResponseBody) GetMessage() string {
	if o == nil || o.Message == nil {
		var ret string
		return ret
	}
	return *o.Message
}

// GetMessageOk returns a tuple with the Message field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AcceptedResponseBody) GetMessageOk() (*string, bool) {
	if o == nil || o.Message == nil {
		return nil, false
	}
	return o.Message, true
}

// HasMessage returns a boolean if a field has been set.
func (o *AcceptedResponseBody) HasMessage() bool {
	if o != nil && o.Message != nil {
		return true
	}

	return false
}

// SetMessage gets a reference to the given string and assigns it to the Message field.
func (o *AcceptedResponseBody) SetMessage(v string) {
	o.Message = &v
}

func (o AcceptedResponseBody) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.Message != nil {
		toSerialize["message"] = o.Message
	}

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}

	return json.Marshal(toSerialize)
}

func (o *AcceptedResponseBody) UnmarshalJSON(bytes []byte) (err error) {
	varAcceptedResponseBody := _AcceptedResponseBody{}

	if err = json.Unmarshal(bytes, &varAcceptedResponseBody); err == nil {
		*o = AcceptedResponseBody(varAcceptedResponseBody)
	}

	additionalProperties := make(map[string]interface{})

	if err = json.Unmarshal(bytes, &additionalProperties); err == nil {
		delete(additionalProperties, "message")
		o.AdditionalProperties = additionalProperties
	}

	return err
}

type NullableAcceptedResponseBody struct {
	value *AcceptedResponseBody
	isSet bool
}

func (v NullableAcceptedResponseBody) Get() *AcceptedResponseBody {
	return v.value
}

func (v *NullableAcceptedResponseBody) Set(val *AcceptedResponseBody) {
	v.value = val
	v.isSet = true
}

func (v NullableAcceptedResponseBody) IsSet() bool {
	return v.isSet
}

func (v *NullableAcceptedResponseBody) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableAcceptedResponseBody(val *AcceptedResponseBody) *NullableAcceptedResponseBody {
	return &NullableAcceptedResponseBody{value: val, isSet: true}
}

func (v NullableAcceptedResponseBody) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableAcceptedResponseBody) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


