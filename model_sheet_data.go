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

// SheetData struct for SheetData
type SheetData struct {
	SheetID *string `json:"sheetID,omitempty"`
	Title *string `json:"title,omitempty"`
	IsVisible *bool `json:"isVisible,omitempty"`
}

// NewSheetData instantiates a new SheetData object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewSheetData() *SheetData {
	this := SheetData{}
	return &this
}

// NewSheetDataWithDefaults instantiates a new SheetData object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewSheetDataWithDefaults() *SheetData {
	this := SheetData{}
	return &this
}

// GetSheetID returns the SheetID field value if set, zero value otherwise.
func (o *SheetData) GetSheetID() string {
	if o == nil || o.SheetID == nil {
		var ret string
		return ret
	}
	return *o.SheetID
}

// GetSheetIDOk returns a tuple with the SheetID field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SheetData) GetSheetIDOk() (*string, bool) {
	if o == nil || o.SheetID == nil {
		return nil, false
	}
	return o.SheetID, true
}

// HasSheetID returns a boolean if a field has been set.
func (o *SheetData) HasSheetID() bool {
	if o != nil && o.SheetID != nil {
		return true
	}

	return false
}

// SetSheetID gets a reference to the given string and assigns it to the SheetID field.
func (o *SheetData) SetSheetID(v string) {
	o.SheetID = &v
}

// GetTitle returns the Title field value if set, zero value otherwise.
func (o *SheetData) GetTitle() string {
	if o == nil || o.Title == nil {
		var ret string
		return ret
	}
	return *o.Title
}

// GetTitleOk returns a tuple with the Title field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SheetData) GetTitleOk() (*string, bool) {
	if o == nil || o.Title == nil {
		return nil, false
	}
	return o.Title, true
}

// HasTitle returns a boolean if a field has been set.
func (o *SheetData) HasTitle() bool {
	if o != nil && o.Title != nil {
		return true
	}

	return false
}

// SetTitle gets a reference to the given string and assigns it to the Title field.
func (o *SheetData) SetTitle(v string) {
	o.Title = &v
}

// GetIsVisible returns the IsVisible field value if set, zero value otherwise.
func (o *SheetData) GetIsVisible() bool {
	if o == nil || o.IsVisible == nil {
		var ret bool
		return ret
	}
	return *o.IsVisible
}

// GetIsVisibleOk returns a tuple with the IsVisible field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SheetData) GetIsVisibleOk() (*bool, bool) {
	if o == nil || o.IsVisible == nil {
		return nil, false
	}
	return o.IsVisible, true
}

// HasIsVisible returns a boolean if a field has been set.
func (o *SheetData) HasIsVisible() bool {
	if o != nil && o.IsVisible != nil {
		return true
	}

	return false
}

// SetIsVisible gets a reference to the given bool and assigns it to the IsVisible field.
func (o *SheetData) SetIsVisible(v bool) {
	o.IsVisible = &v
}

func (o SheetData) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.SheetID != nil {
		toSerialize["sheetID"] = o.SheetID
	}
	if o.Title != nil {
		toSerialize["title"] = o.Title
	}
	if o.IsVisible != nil {
		toSerialize["isVisible"] = o.IsVisible
	}
	return json.Marshal(toSerialize)
}

type NullableSheetData struct {
	value *SheetData
	isSet bool
}

func (v NullableSheetData) Get() *SheetData {
	return v.value
}

func (v *NullableSheetData) Set(val *SheetData) {
	v.value = val
	v.isSet = true
}

func (v NullableSheetData) IsSet() bool {
	return v.isSet
}

func (v *NullableSheetData) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableSheetData(val *SheetData) *NullableSheetData {
	return &NullableSheetData{value: val, isSet: true}
}

func (v NullableSheetData) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableSheetData) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


