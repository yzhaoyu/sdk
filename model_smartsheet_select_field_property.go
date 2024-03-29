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

// SmartsheetSelectFieldProperty struct for SmartsheetSelectFieldProperty
type SmartsheetSelectFieldProperty struct {
	IsMultiple *bool `json:"isMultiple,omitempty"`
	IsQuickAdd *bool `json:"isQuickAdd,omitempty"`
	Options []SmartsheetOption `json:"options,omitempty"`
}

// NewSmartsheetSelectFieldProperty instantiates a new SmartsheetSelectFieldProperty object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewSmartsheetSelectFieldProperty() *SmartsheetSelectFieldProperty {
	this := SmartsheetSelectFieldProperty{}
	return &this
}

// NewSmartsheetSelectFieldPropertyWithDefaults instantiates a new SmartsheetSelectFieldProperty object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewSmartsheetSelectFieldPropertyWithDefaults() *SmartsheetSelectFieldProperty {
	this := SmartsheetSelectFieldProperty{}
	return &this
}

// GetIsMultiple returns the IsMultiple field value if set, zero value otherwise.
func (o *SmartsheetSelectFieldProperty) GetIsMultiple() bool {
	if o == nil || o.IsMultiple == nil {
		var ret bool
		return ret
	}
	return *o.IsMultiple
}

// GetIsMultipleOk returns a tuple with the IsMultiple field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SmartsheetSelectFieldProperty) GetIsMultipleOk() (*bool, bool) {
	if o == nil || o.IsMultiple == nil {
		return nil, false
	}
	return o.IsMultiple, true
}

// HasIsMultiple returns a boolean if a field has been set.
func (o *SmartsheetSelectFieldProperty) HasIsMultiple() bool {
	if o != nil && o.IsMultiple != nil {
		return true
	}

	return false
}

// SetIsMultiple gets a reference to the given bool and assigns it to the IsMultiple field.
func (o *SmartsheetSelectFieldProperty) SetIsMultiple(v bool) {
	o.IsMultiple = &v
}

// GetIsQuickAdd returns the IsQuickAdd field value if set, zero value otherwise.
func (o *SmartsheetSelectFieldProperty) GetIsQuickAdd() bool {
	if o == nil || o.IsQuickAdd == nil {
		var ret bool
		return ret
	}
	return *o.IsQuickAdd
}

// GetIsQuickAddOk returns a tuple with the IsQuickAdd field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SmartsheetSelectFieldProperty) GetIsQuickAddOk() (*bool, bool) {
	if o == nil || o.IsQuickAdd == nil {
		return nil, false
	}
	return o.IsQuickAdd, true
}

// HasIsQuickAdd returns a boolean if a field has been set.
func (o *SmartsheetSelectFieldProperty) HasIsQuickAdd() bool {
	if o != nil && o.IsQuickAdd != nil {
		return true
	}

	return false
}

// SetIsQuickAdd gets a reference to the given bool and assigns it to the IsQuickAdd field.
func (o *SmartsheetSelectFieldProperty) SetIsQuickAdd(v bool) {
	o.IsQuickAdd = &v
}

// GetOptions returns the Options field value if set, zero value otherwise.
func (o *SmartsheetSelectFieldProperty) GetOptions() []SmartsheetOption {
	if o == nil || o.Options == nil {
		var ret []SmartsheetOption
		return ret
	}
	return o.Options
}

// GetOptionsOk returns a tuple with the Options field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SmartsheetSelectFieldProperty) GetOptionsOk() ([]SmartsheetOption, bool) {
	if o == nil || o.Options == nil {
		return nil, false
	}
	return o.Options, true
}

// HasOptions returns a boolean if a field has been set.
func (o *SmartsheetSelectFieldProperty) HasOptions() bool {
	if o != nil && o.Options != nil {
		return true
	}

	return false
}

// SetOptions gets a reference to the given []SmartsheetOption and assigns it to the Options field.
func (o *SmartsheetSelectFieldProperty) SetOptions(v []SmartsheetOption) {
	o.Options = v
}

func (o SmartsheetSelectFieldProperty) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.IsMultiple != nil {
		toSerialize["isMultiple"] = o.IsMultiple
	}
	if o.IsQuickAdd != nil {
		toSerialize["isQuickAdd"] = o.IsQuickAdd
	}
	if o.Options != nil {
		toSerialize["options"] = o.Options
	}
	return json.Marshal(toSerialize)
}

type NullableSmartsheetSelectFieldProperty struct {
	value *SmartsheetSelectFieldProperty
	isSet bool
}

func (v NullableSmartsheetSelectFieldProperty) Get() *SmartsheetSelectFieldProperty {
	return v.value
}

func (v *NullableSmartsheetSelectFieldProperty) Set(val *SmartsheetSelectFieldProperty) {
	v.value = val
	v.isSet = true
}

func (v NullableSmartsheetSelectFieldProperty) IsSet() bool {
	return v.isSet
}

func (v *NullableSmartsheetSelectFieldProperty) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableSmartsheetSelectFieldProperty(val *SmartsheetSelectFieldProperty) *NullableSmartsheetSelectFieldProperty {
	return &NullableSmartsheetSelectFieldProperty{value: val, isSet: true}
}

func (v NullableSmartsheetSelectFieldProperty) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableSmartsheetSelectFieldProperty) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


