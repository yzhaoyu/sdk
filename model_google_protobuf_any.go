/*
Tencent Docs SDK

Tencent docs sdk contains DocAPI, SmartsheetAPI, SheetAPI, DriveAPI and OAuthAPI

API version: 0.0.1
Contact: tencentdocs@tencent.com
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package openapi

import (
	"encoding/json"
)

// GoogleProtobufAny Contains an arbitrary serialized message along with a @type that describes the type of the serialized message.
type GoogleProtobufAny struct {
	// The type of the serialized message.
	Type *string `json:"@type,omitempty"`
	AdditionalProperties map[string]interface{}
}

type _GoogleProtobufAny GoogleProtobufAny

// NewGoogleProtobufAny instantiates a new GoogleProtobufAny object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewGoogleProtobufAny() *GoogleProtobufAny {
	this := GoogleProtobufAny{}
	return &this
}

// NewGoogleProtobufAnyWithDefaults instantiates a new GoogleProtobufAny object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewGoogleProtobufAnyWithDefaults() *GoogleProtobufAny {
	this := GoogleProtobufAny{}
	return &this
}

// GetType returns the Type field value if set, zero value otherwise.
func (o *GoogleProtobufAny) GetType() string {
	if o == nil || o.Type == nil {
		var ret string
		return ret
	}
	return *o.Type
}

// GetTypeOk returns a tuple with the Type field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *GoogleProtobufAny) GetTypeOk() (*string, bool) {
	if o == nil || o.Type == nil {
		return nil, false
	}
	return o.Type, true
}

// HasType returns a boolean if a field has been set.
func (o *GoogleProtobufAny) HasType() bool {
	if o != nil && o.Type != nil {
		return true
	}

	return false
}

// SetType gets a reference to the given string and assigns it to the Type field.
func (o *GoogleProtobufAny) SetType(v string) {
	o.Type = &v
}

func (o GoogleProtobufAny) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.Type != nil {
		toSerialize["@type"] = o.Type
	}

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}

	return json.Marshal(toSerialize)
}

func (o *GoogleProtobufAny) UnmarshalJSON(bytes []byte) (err error) {
	varGoogleProtobufAny := _GoogleProtobufAny{}

	if err = json.Unmarshal(bytes, &varGoogleProtobufAny); err == nil {
		*o = GoogleProtobufAny(varGoogleProtobufAny)
	}

	additionalProperties := make(map[string]interface{})

	if err = json.Unmarshal(bytes, &additionalProperties); err == nil {
		delete(additionalProperties, "@type")
		o.AdditionalProperties = additionalProperties
	}

	return err
}

type NullableGoogleProtobufAny struct {
	value *GoogleProtobufAny
	isSet bool
}

func (v NullableGoogleProtobufAny) Get() *GoogleProtobufAny {
	return v.value
}

func (v *NullableGoogleProtobufAny) Set(val *GoogleProtobufAny) {
	v.value = val
	v.isSet = true
}

func (v NullableGoogleProtobufAny) IsSet() bool {
	return v.isSet
}

func (v *NullableGoogleProtobufAny) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableGoogleProtobufAny(val *GoogleProtobufAny) *NullableGoogleProtobufAny {
	return &NullableGoogleProtobufAny{value: val, isSet: true}
}

func (v NullableGoogleProtobufAny) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableGoogleProtobufAny) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


