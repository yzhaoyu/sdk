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

// InsertWebBlockRequest struct for InsertWebBlockRequest
type InsertWebBlockRequest struct {
	Location Location `json:"location"`
	// 插件 ID
	AddonID string `json:"addonID"`
	// 文档插件元素的唯一ID(6位的数字小写字母组合)
	AnchorID string `json:"anchorID"`
	// 展示文本
	DisplayText *string `json:"displayText,omitempty"`
	// 插件宽度，单位为 emu
	Width *int32 `json:"width,omitempty"`
	// 插件高度，单位为 emu
	Height *int32 `json:"height,omitempty"`
	// 原始URL
	OriginalURL string `json:"originalURL"`
	// 插件展示的URL
	EmbedURL string `json:"embedURL"`
	AddonType string `json:"addonType"`
	// 插件扩展字段
	ExtraData *string `json:"extraData,omitempty"`
	// 锁定宽高比, Default: true
	LockAspectRatio *bool `json:"lockAspectRatio,omitempty"`
}

// NewInsertWebBlockRequest instantiates a new InsertWebBlockRequest object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewInsertWebBlockRequest(location Location, addonID string, anchorID string, originalURL string, embedURL string, addonType string) *InsertWebBlockRequest {
	this := InsertWebBlockRequest{}
	this.Location = location
	this.AddonID = addonID
	this.AnchorID = anchorID
	this.OriginalURL = originalURL
	this.EmbedURL = embedURL
	this.AddonType = addonType
	return &this
}

// NewInsertWebBlockRequestWithDefaults instantiates a new InsertWebBlockRequest object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewInsertWebBlockRequestWithDefaults() *InsertWebBlockRequest {
	this := InsertWebBlockRequest{}
	return &this
}

// GetLocation returns the Location field value
func (o *InsertWebBlockRequest) GetLocation() Location {
	if o == nil {
		var ret Location
		return ret
	}

	return o.Location
}

// GetLocationOk returns a tuple with the Location field value
// and a boolean to check if the value has been set.
func (o *InsertWebBlockRequest) GetLocationOk() (*Location, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Location, true
}

// SetLocation sets field value
func (o *InsertWebBlockRequest) SetLocation(v Location) {
	o.Location = v
}

// GetAddonID returns the AddonID field value
func (o *InsertWebBlockRequest) GetAddonID() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.AddonID
}

// GetAddonIDOk returns a tuple with the AddonID field value
// and a boolean to check if the value has been set.
func (o *InsertWebBlockRequest) GetAddonIDOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.AddonID, true
}

// SetAddonID sets field value
func (o *InsertWebBlockRequest) SetAddonID(v string) {
	o.AddonID = v
}

// GetAnchorID returns the AnchorID field value
func (o *InsertWebBlockRequest) GetAnchorID() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.AnchorID
}

// GetAnchorIDOk returns a tuple with the AnchorID field value
// and a boolean to check if the value has been set.
func (o *InsertWebBlockRequest) GetAnchorIDOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.AnchorID, true
}

// SetAnchorID sets field value
func (o *InsertWebBlockRequest) SetAnchorID(v string) {
	o.AnchorID = v
}

// GetDisplayText returns the DisplayText field value if set, zero value otherwise.
func (o *InsertWebBlockRequest) GetDisplayText() string {
	if o == nil || o.DisplayText == nil {
		var ret string
		return ret
	}
	return *o.DisplayText
}

// GetDisplayTextOk returns a tuple with the DisplayText field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *InsertWebBlockRequest) GetDisplayTextOk() (*string, bool) {
	if o == nil || o.DisplayText == nil {
		return nil, false
	}
	return o.DisplayText, true
}

// HasDisplayText returns a boolean if a field has been set.
func (o *InsertWebBlockRequest) HasDisplayText() bool {
	if o != nil && o.DisplayText != nil {
		return true
	}

	return false
}

// SetDisplayText gets a reference to the given string and assigns it to the DisplayText field.
func (o *InsertWebBlockRequest) SetDisplayText(v string) {
	o.DisplayText = &v
}

// GetWidth returns the Width field value if set, zero value otherwise.
func (o *InsertWebBlockRequest) GetWidth() int32 {
	if o == nil || o.Width == nil {
		var ret int32
		return ret
	}
	return *o.Width
}

// GetWidthOk returns a tuple with the Width field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *InsertWebBlockRequest) GetWidthOk() (*int32, bool) {
	if o == nil || o.Width == nil {
		return nil, false
	}
	return o.Width, true
}

// HasWidth returns a boolean if a field has been set.
func (o *InsertWebBlockRequest) HasWidth() bool {
	if o != nil && o.Width != nil {
		return true
	}

	return false
}

// SetWidth gets a reference to the given int32 and assigns it to the Width field.
func (o *InsertWebBlockRequest) SetWidth(v int32) {
	o.Width = &v
}

// GetHeight returns the Height field value if set, zero value otherwise.
func (o *InsertWebBlockRequest) GetHeight() int32 {
	if o == nil || o.Height == nil {
		var ret int32
		return ret
	}
	return *o.Height
}

// GetHeightOk returns a tuple with the Height field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *InsertWebBlockRequest) GetHeightOk() (*int32, bool) {
	if o == nil || o.Height == nil {
		return nil, false
	}
	return o.Height, true
}

// HasHeight returns a boolean if a field has been set.
func (o *InsertWebBlockRequest) HasHeight() bool {
	if o != nil && o.Height != nil {
		return true
	}

	return false
}

// SetHeight gets a reference to the given int32 and assigns it to the Height field.
func (o *InsertWebBlockRequest) SetHeight(v int32) {
	o.Height = &v
}

// GetOriginalURL returns the OriginalURL field value
func (o *InsertWebBlockRequest) GetOriginalURL() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.OriginalURL
}

// GetOriginalURLOk returns a tuple with the OriginalURL field value
// and a boolean to check if the value has been set.
func (o *InsertWebBlockRequest) GetOriginalURLOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.OriginalURL, true
}

// SetOriginalURL sets field value
func (o *InsertWebBlockRequest) SetOriginalURL(v string) {
	o.OriginalURL = v
}

// GetEmbedURL returns the EmbedURL field value
func (o *InsertWebBlockRequest) GetEmbedURL() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.EmbedURL
}

// GetEmbedURLOk returns a tuple with the EmbedURL field value
// and a boolean to check if the value has been set.
func (o *InsertWebBlockRequest) GetEmbedURLOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.EmbedURL, true
}

// SetEmbedURL sets field value
func (o *InsertWebBlockRequest) SetEmbedURL(v string) {
	o.EmbedURL = v
}

// GetAddonType returns the AddonType field value
func (o *InsertWebBlockRequest) GetAddonType() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.AddonType
}

// GetAddonTypeOk returns a tuple with the AddonType field value
// and a boolean to check if the value has been set.
func (o *InsertWebBlockRequest) GetAddonTypeOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.AddonType, true
}

// SetAddonType sets field value
func (o *InsertWebBlockRequest) SetAddonType(v string) {
	o.AddonType = v
}

// GetExtraData returns the ExtraData field value if set, zero value otherwise.
func (o *InsertWebBlockRequest) GetExtraData() string {
	if o == nil || o.ExtraData == nil {
		var ret string
		return ret
	}
	return *o.ExtraData
}

// GetExtraDataOk returns a tuple with the ExtraData field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *InsertWebBlockRequest) GetExtraDataOk() (*string, bool) {
	if o == nil || o.ExtraData == nil {
		return nil, false
	}
	return o.ExtraData, true
}

// HasExtraData returns a boolean if a field has been set.
func (o *InsertWebBlockRequest) HasExtraData() bool {
	if o != nil && o.ExtraData != nil {
		return true
	}

	return false
}

// SetExtraData gets a reference to the given string and assigns it to the ExtraData field.
func (o *InsertWebBlockRequest) SetExtraData(v string) {
	o.ExtraData = &v
}

// GetLockAspectRatio returns the LockAspectRatio field value if set, zero value otherwise.
func (o *InsertWebBlockRequest) GetLockAspectRatio() bool {
	if o == nil || o.LockAspectRatio == nil {
		var ret bool
		return ret
	}
	return *o.LockAspectRatio
}

// GetLockAspectRatioOk returns a tuple with the LockAspectRatio field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *InsertWebBlockRequest) GetLockAspectRatioOk() (*bool, bool) {
	if o == nil || o.LockAspectRatio == nil {
		return nil, false
	}
	return o.LockAspectRatio, true
}

// HasLockAspectRatio returns a boolean if a field has been set.
func (o *InsertWebBlockRequest) HasLockAspectRatio() bool {
	if o != nil && o.LockAspectRatio != nil {
		return true
	}

	return false
}

// SetLockAspectRatio gets a reference to the given bool and assigns it to the LockAspectRatio field.
func (o *InsertWebBlockRequest) SetLockAspectRatio(v bool) {
	o.LockAspectRatio = &v
}

func (o InsertWebBlockRequest) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if true {
		toSerialize["location"] = o.Location
	}
	if true {
		toSerialize["addonID"] = o.AddonID
	}
	if true {
		toSerialize["anchorID"] = o.AnchorID
	}
	if o.DisplayText != nil {
		toSerialize["displayText"] = o.DisplayText
	}
	if o.Width != nil {
		toSerialize["width"] = o.Width
	}
	if o.Height != nil {
		toSerialize["height"] = o.Height
	}
	if true {
		toSerialize["originalURL"] = o.OriginalURL
	}
	if true {
		toSerialize["embedURL"] = o.EmbedURL
	}
	if true {
		toSerialize["addonType"] = o.AddonType
	}
	if o.ExtraData != nil {
		toSerialize["extraData"] = o.ExtraData
	}
	if o.LockAspectRatio != nil {
		toSerialize["lockAspectRatio"] = o.LockAspectRatio
	}
	return json.Marshal(toSerialize)
}

type NullableInsertWebBlockRequest struct {
	value *InsertWebBlockRequest
	isSet bool
}

func (v NullableInsertWebBlockRequest) Get() *InsertWebBlockRequest {
	return v.value
}

func (v *NullableInsertWebBlockRequest) Set(val *InsertWebBlockRequest) {
	v.value = val
	v.isSet = true
}

func (v NullableInsertWebBlockRequest) IsSet() bool {
	return v.isSet
}

func (v *NullableInsertWebBlockRequest) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableInsertWebBlockRequest(val *InsertWebBlockRequest) *NullableInsertWebBlockRequest {
	return &NullableInsertWebBlockRequest{value: val, isSet: true}
}

func (v NullableInsertWebBlockRequest) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableInsertWebBlockRequest) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


