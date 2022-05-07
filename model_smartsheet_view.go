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

// SmartsheetView struct for SmartsheetView
type SmartsheetView struct {
	ViewID *string `json:"viewID,omitempty"`
	ViewTitle *string `json:"viewTitle,omitempty"`
	ViewType *string `json:"viewType,omitempty"`
}

// NewSmartsheetView instantiates a new SmartsheetView object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewSmartsheetView() *SmartsheetView {
	this := SmartsheetView{}
	return &this
}

// NewSmartsheetViewWithDefaults instantiates a new SmartsheetView object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewSmartsheetViewWithDefaults() *SmartsheetView {
	this := SmartsheetView{}
	return &this
}

// GetViewID returns the ViewID field value if set, zero value otherwise.
func (o *SmartsheetView) GetViewID() string {
	if o == nil || o.ViewID == nil {
		var ret string
		return ret
	}
	return *o.ViewID
}

// GetViewIDOk returns a tuple with the ViewID field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SmartsheetView) GetViewIDOk() (*string, bool) {
	if o == nil || o.ViewID == nil {
		return nil, false
	}
	return o.ViewID, true
}

// HasViewID returns a boolean if a field has been set.
func (o *SmartsheetView) HasViewID() bool {
	if o != nil && o.ViewID != nil {
		return true
	}

	return false
}

// SetViewID gets a reference to the given string and assigns it to the ViewID field.
func (o *SmartsheetView) SetViewID(v string) {
	o.ViewID = &v
}

// GetViewTitle returns the ViewTitle field value if set, zero value otherwise.
func (o *SmartsheetView) GetViewTitle() string {
	if o == nil || o.ViewTitle == nil {
		var ret string
		return ret
	}
	return *o.ViewTitle
}

// GetViewTitleOk returns a tuple with the ViewTitle field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SmartsheetView) GetViewTitleOk() (*string, bool) {
	if o == nil || o.ViewTitle == nil {
		return nil, false
	}
	return o.ViewTitle, true
}

// HasViewTitle returns a boolean if a field has been set.
func (o *SmartsheetView) HasViewTitle() bool {
	if o != nil && o.ViewTitle != nil {
		return true
	}

	return false
}

// SetViewTitle gets a reference to the given string and assigns it to the ViewTitle field.
func (o *SmartsheetView) SetViewTitle(v string) {
	o.ViewTitle = &v
}

// GetViewType returns the ViewType field value if set, zero value otherwise.
func (o *SmartsheetView) GetViewType() string {
	if o == nil || o.ViewType == nil {
		var ret string
		return ret
	}
	return *o.ViewType
}

// GetViewTypeOk returns a tuple with the ViewType field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SmartsheetView) GetViewTypeOk() (*string, bool) {
	if o == nil || o.ViewType == nil {
		return nil, false
	}
	return o.ViewType, true
}

// HasViewType returns a boolean if a field has been set.
func (o *SmartsheetView) HasViewType() bool {
	if o != nil && o.ViewType != nil {
		return true
	}

	return false
}

// SetViewType gets a reference to the given string and assigns it to the ViewType field.
func (o *SmartsheetView) SetViewType(v string) {
	o.ViewType = &v
}

func (o SmartsheetView) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.ViewID != nil {
		toSerialize["viewID"] = o.ViewID
	}
	if o.ViewTitle != nil {
		toSerialize["viewTitle"] = o.ViewTitle
	}
	if o.ViewType != nil {
		toSerialize["viewType"] = o.ViewType
	}
	return json.Marshal(toSerialize)
}

type NullableSmartsheetView struct {
	value *SmartsheetView
	isSet bool
}

func (v NullableSmartsheetView) Get() *SmartsheetView {
	return v.value
}

func (v *NullableSmartsheetView) Set(val *SmartsheetView) {
	v.value = val
	v.isSet = true
}

func (v NullableSmartsheetView) IsSet() bool {
	return v.isSet
}

func (v *NullableSmartsheetView) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableSmartsheetView(val *SmartsheetView) *NullableSmartsheetView {
	return &NullableSmartsheetView{value: val, isSet: true}
}

func (v NullableSmartsheetView) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableSmartsheetView) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


