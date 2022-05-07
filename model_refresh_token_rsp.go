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

// RefreshTokenRsp struct for RefreshTokenRsp
type RefreshTokenRsp struct {
	AccessToken *string `json:"access_token,omitempty"`
	TokenType *string `json:"token_type,omitempty"`
	ExpiresIn *int64 `json:"expires_in,omitempty"`
	UserId *string `json:"user_id,omitempty"`
	Scope *string `json:"scope,omitempty"`
}

// NewRefreshTokenRsp instantiates a new RefreshTokenRsp object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewRefreshTokenRsp() *RefreshTokenRsp {
	this := RefreshTokenRsp{}
	return &this
}

// NewRefreshTokenRspWithDefaults instantiates a new RefreshTokenRsp object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewRefreshTokenRspWithDefaults() *RefreshTokenRsp {
	this := RefreshTokenRsp{}
	return &this
}

// GetAccessToken returns the AccessToken field value if set, zero value otherwise.
func (o *RefreshTokenRsp) GetAccessToken() string {
	if o == nil || o.AccessToken == nil {
		var ret string
		return ret
	}
	return *o.AccessToken
}

// GetAccessTokenOk returns a tuple with the AccessToken field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *RefreshTokenRsp) GetAccessTokenOk() (*string, bool) {
	if o == nil || o.AccessToken == nil {
		return nil, false
	}
	return o.AccessToken, true
}

// HasAccessToken returns a boolean if a field has been set.
func (o *RefreshTokenRsp) HasAccessToken() bool {
	if o != nil && o.AccessToken != nil {
		return true
	}

	return false
}

// SetAccessToken gets a reference to the given string and assigns it to the AccessToken field.
func (o *RefreshTokenRsp) SetAccessToken(v string) {
	o.AccessToken = &v
}

// GetTokenType returns the TokenType field value if set, zero value otherwise.
func (o *RefreshTokenRsp) GetTokenType() string {
	if o == nil || o.TokenType == nil {
		var ret string
		return ret
	}
	return *o.TokenType
}

// GetTokenTypeOk returns a tuple with the TokenType field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *RefreshTokenRsp) GetTokenTypeOk() (*string, bool) {
	if o == nil || o.TokenType == nil {
		return nil, false
	}
	return o.TokenType, true
}

// HasTokenType returns a boolean if a field has been set.
func (o *RefreshTokenRsp) HasTokenType() bool {
	if o != nil && o.TokenType != nil {
		return true
	}

	return false
}

// SetTokenType gets a reference to the given string and assigns it to the TokenType field.
func (o *RefreshTokenRsp) SetTokenType(v string) {
	o.TokenType = &v
}

// GetExpiresIn returns the ExpiresIn field value if set, zero value otherwise.
func (o *RefreshTokenRsp) GetExpiresIn() int64 {
	if o == nil || o.ExpiresIn == nil {
		var ret int64
		return ret
	}
	return *o.ExpiresIn
}

// GetExpiresInOk returns a tuple with the ExpiresIn field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *RefreshTokenRsp) GetExpiresInOk() (*int64, bool) {
	if o == nil || o.ExpiresIn == nil {
		return nil, false
	}
	return o.ExpiresIn, true
}

// HasExpiresIn returns a boolean if a field has been set.
func (o *RefreshTokenRsp) HasExpiresIn() bool {
	if o != nil && o.ExpiresIn != nil {
		return true
	}

	return false
}

// SetExpiresIn gets a reference to the given int64 and assigns it to the ExpiresIn field.
func (o *RefreshTokenRsp) SetExpiresIn(v int64) {
	o.ExpiresIn = &v
}

// GetUserId returns the UserId field value if set, zero value otherwise.
func (o *RefreshTokenRsp) GetUserId() string {
	if o == nil || o.UserId == nil {
		var ret string
		return ret
	}
	return *o.UserId
}

// GetUserIdOk returns a tuple with the UserId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *RefreshTokenRsp) GetUserIdOk() (*string, bool) {
	if o == nil || o.UserId == nil {
		return nil, false
	}
	return o.UserId, true
}

// HasUserId returns a boolean if a field has been set.
func (o *RefreshTokenRsp) HasUserId() bool {
	if o != nil && o.UserId != nil {
		return true
	}

	return false
}

// SetUserId gets a reference to the given string and assigns it to the UserId field.
func (o *RefreshTokenRsp) SetUserId(v string) {
	o.UserId = &v
}

// GetScope returns the Scope field value if set, zero value otherwise.
func (o *RefreshTokenRsp) GetScope() string {
	if o == nil || o.Scope == nil {
		var ret string
		return ret
	}
	return *o.Scope
}

// GetScopeOk returns a tuple with the Scope field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *RefreshTokenRsp) GetScopeOk() (*string, bool) {
	if o == nil || o.Scope == nil {
		return nil, false
	}
	return o.Scope, true
}

// HasScope returns a boolean if a field has been set.
func (o *RefreshTokenRsp) HasScope() bool {
	if o != nil && o.Scope != nil {
		return true
	}

	return false
}

// SetScope gets a reference to the given string and assigns it to the Scope field.
func (o *RefreshTokenRsp) SetScope(v string) {
	o.Scope = &v
}

func (o RefreshTokenRsp) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.AccessToken != nil {
		toSerialize["access_token"] = o.AccessToken
	}
	if o.TokenType != nil {
		toSerialize["token_type"] = o.TokenType
	}
	if o.ExpiresIn != nil {
		toSerialize["expires_in"] = o.ExpiresIn
	}
	if o.UserId != nil {
		toSerialize["user_id"] = o.UserId
	}
	if o.Scope != nil {
		toSerialize["scope"] = o.Scope
	}
	return json.Marshal(toSerialize)
}

type NullableRefreshTokenRsp struct {
	value *RefreshTokenRsp
	isSet bool
}

func (v NullableRefreshTokenRsp) Get() *RefreshTokenRsp {
	return v.value
}

func (v *NullableRefreshTokenRsp) Set(val *RefreshTokenRsp) {
	v.value = val
	v.isSet = true
}

func (v NullableRefreshTokenRsp) IsSet() bool {
	return v.isSet
}

func (v *NullableRefreshTokenRsp) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableRefreshTokenRsp(val *RefreshTokenRsp) *NullableRefreshTokenRsp {
	return &NullableRefreshTokenRsp{value: val, isSet: true}
}

func (v NullableRefreshTokenRsp) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableRefreshTokenRsp) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


