/*
Tencent Docs SDK

Tencent docs sdk contains DocAPI, SmartsheetAPI, SheetAPI, DriveAPI and OAuthAPI

API version: 0.0.1
Contact: tencentdocs@tencent.com
*/

// Code generated by Tencent Docs (https://docs.qq.com); DO NOT EDIT.

package openapi

import (
	"encoding/json"
)

// InsertPageBreakRequest struct for InsertPageBreakRequest
type InsertPageBreakRequest struct {
	Location *Location `json:"location,omitempty"`
}

// NewInsertPageBreakRequest instantiates a new InsertPageBreakRequest object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewInsertPageBreakRequest() *InsertPageBreakRequest {
	this := InsertPageBreakRequest{}
	return &this
}

// NewInsertPageBreakRequestWithDefaults instantiates a new InsertPageBreakRequest object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewInsertPageBreakRequestWithDefaults() *InsertPageBreakRequest {
	this := InsertPageBreakRequest{}
	return &this
}

// GetLocation returns the Location field value if set, zero value otherwise.
func (o *InsertPageBreakRequest) GetLocation() Location {
	if o == nil || o.Location == nil {
		var ret Location
		return ret
	}
	return *o.Location
}

// GetLocationOk returns a tuple with the Location field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *InsertPageBreakRequest) GetLocationOk() (*Location, bool) {
	if o == nil || o.Location == nil {
		return nil, false
	}
	return o.Location, true
}

// HasLocation returns a boolean if a field has been set.
func (o *InsertPageBreakRequest) HasLocation() bool {
	if o != nil && o.Location != nil {
		return true
	}

	return false
}

// SetLocation gets a reference to the given Location and assigns it to the Location field.
func (o *InsertPageBreakRequest) SetLocation(v Location) {
	o.Location = &v
}

func (o InsertPageBreakRequest) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.Location != nil {
		toSerialize["location"] = o.Location
	}
	return json.Marshal(toSerialize)
}

type NullableInsertPageBreakRequest struct {
	value *InsertPageBreakRequest
	isSet bool
}

func (v NullableInsertPageBreakRequest) Get() *InsertPageBreakRequest {
	return v.value
}

func (v *NullableInsertPageBreakRequest) Set(val *InsertPageBreakRequest) {
	v.value = val
	v.isSet = true
}

func (v NullableInsertPageBreakRequest) IsSet() bool {
	return v.isSet
}

func (v *NullableInsertPageBreakRequest) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableInsertPageBreakRequest(val *InsertPageBreakRequest) *NullableInsertPageBreakRequest {
	return &NullableInsertPageBreakRequest{value: val, isSet: true}
}

func (v NullableInsertPageBreakRequest) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableInsertPageBreakRequest) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


