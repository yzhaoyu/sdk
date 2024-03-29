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

// Node struct for Node
type Node struct {
	Begin *int32 `json:"begin,omitempty"`
	End *int32 `json:"end,omitempty"`
	Type *string `json:"type,omitempty"`
	Children []Node `json:"children,omitempty"`
	Text *string `json:"text,omitempty"`
	Property *map[string]interface{} `json:"property,omitempty"`
}

// NewNode instantiates a new Node object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewNode() *Node {
	this := Node{}
	return &this
}

// NewNodeWithDefaults instantiates a new Node object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewNodeWithDefaults() *Node {
	this := Node{}
	return &this
}

// GetBegin returns the Begin field value if set, zero value otherwise.
func (o *Node) GetBegin() int32 {
	if o == nil || o.Begin == nil {
		var ret int32
		return ret
	}
	return *o.Begin
}

// GetBeginOk returns a tuple with the Begin field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Node) GetBeginOk() (*int32, bool) {
	if o == nil || o.Begin == nil {
		return nil, false
	}
	return o.Begin, true
}

// HasBegin returns a boolean if a field has been set.
func (o *Node) HasBegin() bool {
	if o != nil && o.Begin != nil {
		return true
	}

	return false
}

// SetBegin gets a reference to the given int32 and assigns it to the Begin field.
func (o *Node) SetBegin(v int32) {
	o.Begin = &v
}

// GetEnd returns the End field value if set, zero value otherwise.
func (o *Node) GetEnd() int32 {
	if o == nil || o.End == nil {
		var ret int32
		return ret
	}
	return *o.End
}

// GetEndOk returns a tuple with the End field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Node) GetEndOk() (*int32, bool) {
	if o == nil || o.End == nil {
		return nil, false
	}
	return o.End, true
}

// HasEnd returns a boolean if a field has been set.
func (o *Node) HasEnd() bool {
	if o != nil && o.End != nil {
		return true
	}

	return false
}

// SetEnd gets a reference to the given int32 and assigns it to the End field.
func (o *Node) SetEnd(v int32) {
	o.End = &v
}

// GetType returns the Type field value if set, zero value otherwise.
func (o *Node) GetType() string {
	if o == nil || o.Type == nil {
		var ret string
		return ret
	}
	return *o.Type
}

// GetTypeOk returns a tuple with the Type field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Node) GetTypeOk() (*string, bool) {
	if o == nil || o.Type == nil {
		return nil, false
	}
	return o.Type, true
}

// HasType returns a boolean if a field has been set.
func (o *Node) HasType() bool {
	if o != nil && o.Type != nil {
		return true
	}

	return false
}

// SetType gets a reference to the given string and assigns it to the Type field.
func (o *Node) SetType(v string) {
	o.Type = &v
}

// GetChildren returns the Children field value if set, zero value otherwise.
func (o *Node) GetChildren() []Node {
	if o == nil || o.Children == nil {
		var ret []Node
		return ret
	}
	return o.Children
}

// GetChildrenOk returns a tuple with the Children field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Node) GetChildrenOk() ([]Node, bool) {
	if o == nil || o.Children == nil {
		return nil, false
	}
	return o.Children, true
}

// HasChildren returns a boolean if a field has been set.
func (o *Node) HasChildren() bool {
	if o != nil && o.Children != nil {
		return true
	}

	return false
}

// SetChildren gets a reference to the given []Node and assigns it to the Children field.
func (o *Node) SetChildren(v []Node) {
	o.Children = v
}

// GetText returns the Text field value if set, zero value otherwise.
func (o *Node) GetText() string {
	if o == nil || o.Text == nil {
		var ret string
		return ret
	}
	return *o.Text
}

// GetTextOk returns a tuple with the Text field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Node) GetTextOk() (*string, bool) {
	if o == nil || o.Text == nil {
		return nil, false
	}
	return o.Text, true
}

// HasText returns a boolean if a field has been set.
func (o *Node) HasText() bool {
	if o != nil && o.Text != nil {
		return true
	}

	return false
}

// SetText gets a reference to the given string and assigns it to the Text field.
func (o *Node) SetText(v string) {
	o.Text = &v
}

// GetProperty returns the Property field value if set, zero value otherwise.
func (o *Node) GetProperty() map[string]interface{} {
	if o == nil || o.Property == nil {
		var ret map[string]interface{}
		return ret
	}
	return *o.Property
}

// GetPropertyOk returns a tuple with the Property field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Node) GetPropertyOk() (*map[string]interface{}, bool) {
	if o == nil || o.Property == nil {
		return nil, false
	}
	return o.Property, true
}

// HasProperty returns a boolean if a field has been set.
func (o *Node) HasProperty() bool {
	if o != nil && o.Property != nil {
		return true
	}

	return false
}

// SetProperty gets a reference to the given map[string]interface{} and assigns it to the Property field.
func (o *Node) SetProperty(v map[string]interface{}) {
	o.Property = &v
}

func (o Node) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.Begin != nil {
		toSerialize["begin"] = o.Begin
	}
	if o.End != nil {
		toSerialize["end"] = o.End
	}
	if o.Type != nil {
		toSerialize["type"] = o.Type
	}
	if o.Children != nil {
		toSerialize["children"] = o.Children
	}
	if o.Text != nil {
		toSerialize["text"] = o.Text
	}
	if o.Property != nil {
		toSerialize["property"] = o.Property
	}
	return json.Marshal(toSerialize)
}

type NullableNode struct {
	value *Node
	isSet bool
}

func (v NullableNode) Get() *Node {
	return v.value
}

func (v *NullableNode) Set(val *Node) {
	v.value = val
	v.isSet = true
}

func (v NullableNode) IsSet() bool {
	return v.isSet
}

func (v *NullableNode) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableNode(val *Node) *NullableNode {
	return &NullableNode{value: val, isSet: true}
}

func (v NullableNode) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableNode) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


