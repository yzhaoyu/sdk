/*
Tencent Docs SDK

Tencent docs sdk contains DocAPI, SmartsheetAPI, SheetAPI, DriveAPI and OAuthAPI

API version: 0.0.1
Contact: tencentdocs@tencent.com
*/

// Code generated by Tencent Docs (https://docs.qq.com/open/wiki); DO NOT EDIT.

package openapi

import (
	"encoding/json"
)

// DeleteContentRequest struct for DeleteContentRequest
type DeleteContentRequest struct {
	Range *Range `json:"range,omitempty"`
}

// NewDeleteContentRequest instantiates a new DeleteContentRequest object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewDeleteContentRequest() *DeleteContentRequest {
	this := DeleteContentRequest{}
	return &this
}

// NewDeleteContentRequestWithDefaults instantiates a new DeleteContentRequest object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewDeleteContentRequestWithDefaults() *DeleteContentRequest {
	this := DeleteContentRequest{}
	return &this
}

// GetRange returns the Range field value if set, zero value otherwise.
func (o *DeleteContentRequest) GetRange() Range {
	if o == nil || o.Range == nil {
		var ret Range
		return ret
	}
	return *o.Range
}

// GetRangeOk returns a tuple with the Range field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *DeleteContentRequest) GetRangeOk() (*Range, bool) {
	if o == nil || o.Range == nil {
		return nil, false
	}
	return o.Range, true
}

// HasRange returns a boolean if a field has been set.
func (o *DeleteContentRequest) HasRange() bool {
	if o != nil && o.Range != nil {
		return true
	}

	return false
}

// SetRange gets a reference to the given Range and assigns it to the Range field.
func (o *DeleteContentRequest) SetRange(v Range) {
	o.Range = &v
}

func (o DeleteContentRequest) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.Range != nil {
		toSerialize["range"] = o.Range
	}
	return json.Marshal(toSerialize)
}

type NullableDeleteContentRequest struct {
	value *DeleteContentRequest
	isSet bool
}

func (v NullableDeleteContentRequest) Get() *DeleteContentRequest {
	return v.value
}

func (v *NullableDeleteContentRequest) Set(val *DeleteContentRequest) {
	v.value = val
	v.isSet = true
}

func (v NullableDeleteContentRequest) IsSet() bool {
	return v.isSet
}

func (v *NullableDeleteContentRequest) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableDeleteContentRequest(val *DeleteContentRequest) *NullableDeleteContentRequest {
	return &NullableDeleteContentRequest{value: val, isSet: true}
}

func (v NullableDeleteContentRequest) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableDeleteContentRequest) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


