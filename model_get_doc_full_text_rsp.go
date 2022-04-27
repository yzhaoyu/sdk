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

// GetDocFullTextRsp struct for GetDocFullTextRsp
type GetDocFullTextRsp struct {
	Document *Node `json:"document,omitempty"`
}

// NewGetDocFullTextRsp instantiates a new GetDocFullTextRsp object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewGetDocFullTextRsp() *GetDocFullTextRsp {
	this := GetDocFullTextRsp{}
	return &this
}

// NewGetDocFullTextRspWithDefaults instantiates a new GetDocFullTextRsp object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewGetDocFullTextRspWithDefaults() *GetDocFullTextRsp {
	this := GetDocFullTextRsp{}
	return &this
}

// GetDocument returns the Document field value if set, zero value otherwise.
func (o *GetDocFullTextRsp) GetDocument() Node {
	if o == nil || o.Document == nil {
		var ret Node
		return ret
	}
	return *o.Document
}

// GetDocumentOk returns a tuple with the Document field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *GetDocFullTextRsp) GetDocumentOk() (*Node, bool) {
	if o == nil || o.Document == nil {
		return nil, false
	}
	return o.Document, true
}

// HasDocument returns a boolean if a field has been set.
func (o *GetDocFullTextRsp) HasDocument() bool {
	if o != nil && o.Document != nil {
		return true
	}

	return false
}

// SetDocument gets a reference to the given Node and assigns it to the Document field.
func (o *GetDocFullTextRsp) SetDocument(v Node) {
	o.Document = &v
}

func (o GetDocFullTextRsp) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.Document != nil {
		toSerialize["document"] = o.Document
	}
	return json.Marshal(toSerialize)
}

type NullableGetDocFullTextRsp struct {
	value *GetDocFullTextRsp
	isSet bool
}

func (v NullableGetDocFullTextRsp) Get() *GetDocFullTextRsp {
	return v.value
}

func (v *NullableGetDocFullTextRsp) Set(val *GetDocFullTextRsp) {
	v.value = val
	v.isSet = true
}

func (v NullableGetDocFullTextRsp) IsSet() bool {
	return v.isSet
}

func (v *NullableGetDocFullTextRsp) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableGetDocFullTextRsp(val *GetDocFullTextRsp) *NullableGetDocFullTextRsp {
	return &NullableGetDocFullTextRsp{value: val, isSet: true}
}

func (v NullableGetDocFullTextRsp) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableGetDocFullTextRsp) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


