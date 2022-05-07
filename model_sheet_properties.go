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

// SheetProperties struct for SheetProperties
type SheetProperties struct {
	SheetID *string `json:"sheetID,omitempty"`
	Title *string `json:"title,omitempty"`
	Index *int32 `json:"index,omitempty"`
}

// NewSheetProperties instantiates a new SheetProperties object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewSheetProperties() *SheetProperties {
	this := SheetProperties{}
	return &this
}

// NewSheetPropertiesWithDefaults instantiates a new SheetProperties object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewSheetPropertiesWithDefaults() *SheetProperties {
	this := SheetProperties{}
	return &this
}

// GetSheetID returns the SheetID field value if set, zero value otherwise.
func (o *SheetProperties) GetSheetID() string {
	if o == nil || o.SheetID == nil {
		var ret string
		return ret
	}
	return *o.SheetID
}

// GetSheetIDOk returns a tuple with the SheetID field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SheetProperties) GetSheetIDOk() (*string, bool) {
	if o == nil || o.SheetID == nil {
		return nil, false
	}
	return o.SheetID, true
}

// HasSheetID returns a boolean if a field has been set.
func (o *SheetProperties) HasSheetID() bool {
	if o != nil && o.SheetID != nil {
		return true
	}

	return false
}

// SetSheetID gets a reference to the given string and assigns it to the SheetID field.
func (o *SheetProperties) SetSheetID(v string) {
	o.SheetID = &v
}

// GetTitle returns the Title field value if set, zero value otherwise.
func (o *SheetProperties) GetTitle() string {
	if o == nil || o.Title == nil {
		var ret string
		return ret
	}
	return *o.Title
}

// GetTitleOk returns a tuple with the Title field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SheetProperties) GetTitleOk() (*string, bool) {
	if o == nil || o.Title == nil {
		return nil, false
	}
	return o.Title, true
}

// HasTitle returns a boolean if a field has been set.
func (o *SheetProperties) HasTitle() bool {
	if o != nil && o.Title != nil {
		return true
	}

	return false
}

// SetTitle gets a reference to the given string and assigns it to the Title field.
func (o *SheetProperties) SetTitle(v string) {
	o.Title = &v
}

// GetIndex returns the Index field value if set, zero value otherwise.
func (o *SheetProperties) GetIndex() int32 {
	if o == nil || o.Index == nil {
		var ret int32
		return ret
	}
	return *o.Index
}

// GetIndexOk returns a tuple with the Index field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SheetProperties) GetIndexOk() (*int32, bool) {
	if o == nil || o.Index == nil {
		return nil, false
	}
	return o.Index, true
}

// HasIndex returns a boolean if a field has been set.
func (o *SheetProperties) HasIndex() bool {
	if o != nil && o.Index != nil {
		return true
	}

	return false
}

// SetIndex gets a reference to the given int32 and assigns it to the Index field.
func (o *SheetProperties) SetIndex(v int32) {
	o.Index = &v
}

func (o SheetProperties) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.SheetID != nil {
		toSerialize["sheetID"] = o.SheetID
	}
	if o.Title != nil {
		toSerialize["title"] = o.Title
	}
	if o.Index != nil {
		toSerialize["index"] = o.Index
	}
	return json.Marshal(toSerialize)
}

type NullableSheetProperties struct {
	value *SheetProperties
	isSet bool
}

func (v NullableSheetProperties) Get() *SheetProperties {
	return v.value
}

func (v *NullableSheetProperties) Set(val *SheetProperties) {
	v.value = val
	v.isSet = true
}

func (v NullableSheetProperties) IsSet() bool {
	return v.isSet
}

func (v *NullableSheetProperties) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableSheetProperties(val *SheetProperties) *NullableSheetProperties {
	return &NullableSheetProperties{value: val, isSet: true}
}

func (v NullableSheetProperties) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableSheetProperties) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


